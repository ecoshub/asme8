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

	StackStart uint16 = 0x80ff
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

	aluDirectOut bool
	aluEnable    bool
	bridgeEnable bool
	bridgeDir    uint8

	busX    *bus.Bus
	busY    *bus.Bus
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
		busY:         bus.New(),
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
	fmt.Printf("busses, data: %x, x: %x, y: %x\n", c.dataBus.Read(), c.busX.Read(), c.busY.Read())
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
	defer c.clear()
	microinstructions := CONTROL_SIGNALS[c.instructionRegister][c.step]
	for i, mi := range microinstructions {
		keep := c.run(mi)
		if !keep {
			fmt.Printf("** [ %-16s ] mi: %d, inst: %2x, step: %d [%d]\n", MI_NAME_MAP[mi], mi, c.instructionRegister, c.step, i)
			break
		}
		if c.debug {
			if c.verbose {
				fmt.Printf("# [ %-16s ] pc: %04x, step: %d, inst_r: %02x, op_r: %02x, addr: %04x, data: %02x, bus_x: %02x, bus_y: %02x, rw: %x, status: %08b, regs: %s\n", MI_NAME_MAP[mi], c.programCounter, c.step, c.instructionRegister, c.operandRegister, c.addrBus.Read(), c.dataBus.Read(), c.busX.Read(), c.busY.Read(), c.rw, c.status.Flag(), c.registers)
			}
		}
	}
	if c.debug {
		if !c.verbose {
			fmt.Printf("# pc: %04x, inst_r: %02x, op_r: %02x, addr: %04x, data: %02x, rw: %x, status: %08b, regs: %s\n", c.programCounter, c.instructionRegister, c.operandRegister, c.addrBus.Read(), c.dataBus.Read(), c.rw, c.status.Flag(), c.registers)
		} else {
			fmt.Println("--")
		}
	}
	if len(microinstructions) == 0 || c.instructionRegister == instruction.Type(MI_BRK) {
		fmt.Println(" # BREAK # ")
		os.Exit(0)
		return false
	}
	return true
}

func (c *Comp) run(mi uint64) bool {
	f, ok := microinstructionFunctions[mi]
	if ok {
		f(c, mi)
		return true
	}
	return false
}

func (c *Comp) tickDevices() {
	for _, dev := range c.devices {
		dev.Tick(c.rw)
	}
}

func (c *Comp) clear() {
	c.dataBus.Clear()
	c.busX.Clear()
	c.busY.Clear()
	c.addrBus.Clear()
	c.bridgeDir = BRIDGE_DIR_IN
	c.bridgeEnable = false
	c.aluEnable = false
	c.aluDirectOut = false
}
