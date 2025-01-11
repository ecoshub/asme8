package comp

import (
	"emu/src/bus"
	"emu/src/connectable"
	"emu/src/instruction"
	"emu/src/register"
	"emu/src/status"
	"fmt"
	"os"
	"time"
)

const (
	DefaultMemorySize = 1 << 16

	StackStart uint16 = 0x7fff
)

type Comp struct {
	registers register.Module
	status    *status.StatusRegister
	step      uint8

	instructionRegister instruction.Type
	operandRegister     uint8

	programCounter        uint16
	memoryAddressRegister uint16
	stackPointer          uint16
	memoryDataRegister    uint8

	busX    *bus.Bus
	dataBus *bus.Bus
	addrBus *bus.Bus
	store   uint8
	rw      uint8 // read 1 write 0
	devices []connectable.Connectable

	debug   bool
	verbose bool
	delay   time.Duration
}

func New() *Comp {
	c := &Comp{
		registers:    register.NewModule(),
		status:       status.NewStatusRegister(),
		stackPointer: StackStart,
		busX:         bus.New(),
		dataBus:      bus.New(),
		addrBus:      bus.New(),
		devices:      make([]connectable.Connectable, 0, 1),
	}
	// ConfigureBIOS(c)
	return c
}

func (c *Comp) SetDebug(enable bool) {
	c.debug = enable
}

func (c *Comp) SetVerbose(enable bool) {
	c.verbose = enable
}

func (c *Comp) SetDelayMS(delay time.Duration) {
	c.delay = delay
}

func (c *Comp) PrintRegisters() {
	fmt.Println(c.registers)
}

func (c *Comp) PrintBusses() {
	fmt.Printf("busses, data: %x, x: %x\n", c.dataBus.Read(), c.busX.Read())
}

func (c *Comp) ConnectDevice(dev connectable.Connectable, rangeStart uint16, rangeEnd uint16) {
	if dev == nil {
		return
	}
	dev.Attach(c.addrBus, c.dataBus, rangeStart, rangeEnd)
	c.devices = append(c.devices, dev)
}

func (c *Comp) Run() {
	t := time.NewTicker(c.delay)
	for range t.C {
		keep := c.tick()
		if !keep {
			return
		}
	}
}

func (c *Comp) tick() bool {
	defer c.clearBusses()

	command := control[c.instructionRegister][c.step]
	if c.debug {
		if c.step == 0 {
			fmt.Println(" --- ")
		}
		if c.verbose {
			fmt.Printf("# pc: %02x, step: %d, inst: %02x, operand:%02x, addr: %04x, data: %02x, bus_x: %02x, status: %08b, registers: %s\n", c.programCounter, c.step, c.instructionRegister, c.operandRegister, c.addrBus.Read(), c.dataBus.Read(), c.busX.Read(), c.status.Flag(), c.registers)
		}
	}
	if command == MI_BRK {
		fmt.Println(" # BREAK # ")
		os.Exit(0)
		return false
	}
	c.run(command)
	return true
}

func (c *Comp) run(command uint64) {
	for i := 0; i < 64; i++ {
		mask := uint64(1) << i
		if command&mask > 0 {
			microInstructions[mask](c, command, mask)
		}
	}
}

func (c *Comp) tickDevices() {
	for _, dev := range c.devices {
		dev.Tick(c.rw)
	}
}

func (c *Comp) clearBusses() {
	c.dataBus.Clear()
	c.busX.Clear()
	c.addrBus.Clear()
}
