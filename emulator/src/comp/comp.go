package comp

import (
	"asme8/emulator/src/bus"
	"asme8/emulator/src/connectable"
	"asme8/emulator/src/instruction"
	"asme8/emulator/src/register"
	"asme8/emulator/src/status"
	"asme8/emulator/src/terminal"
	"asme8/emulator/utils"
	"fmt"
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

	aluDirectOut   bool
	aluEnable      bool
	aluStoreEnable bool
	bridgeEnable   bool
	bridgeDir      uint8

	busX    *bus.Bus
	busY    *bus.Bus
	dataBus *bus.Bus
	addrBus *bus.Bus
	store   uint8
	rw      uint8 // read 1 write 0
	devices []connectable.Connectable

	terminalComponents *terminal.Components

	stopChan chan struct{}
	running  bool
	debug    bool
	verbose  bool
	delay    time.Duration
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
		stopChan:     make(chan struct{}, 1),
	}
	// ConfigureBIOS(c)
	return c
}

func (c *Comp) AttachTerminalComponents(terminalComponents *terminal.Components) {
	c.terminalComponents = terminalComponents
}

func (c *Comp) SetDebug(enable bool) {
	c.debug = enable
}

func (c *Comp) SetVerbose(enable bool) {
	c.verbose = enable
}

func (c *Comp) SetDelay(delay time.Duration) {
	c.delay = delay
}

func (c *Comp) PrintRegisters() {
	fmt.Println(c.registers)
}

func (c *Comp) PrintBusses() {
	fmt.Printf("busses, data: %x, x: %x, y: %x\n", c.dataBus.Read(), c.busX.Read(), c.busY.Read())
}

func (c *Comp) ReadRegister(index uint8) uint8 {
	return c.registers.Read(index)
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
	c.running = true
	for {
		select {
		case <-t.C:
			keep := c.tick()
			if !keep {
				return
			}
		case <-c.stopChan:
			return
		}
	}
}

func (c *Comp) Stop() {
	if !c.running {
		return
	}
	c.stopChan <- struct{}{}
	c.running = false
}

func (c *Comp) tick() bool {
	defer c.clear()
	microinstructions := CONTROL_SIGNALS[c.instructionRegister][c.step]
	for i, mi := range microinstructions {
		keep := c.run(mi)
		if !keep {
			c.Logf("** [ %-16s ] mi: %d, inst: %2x, step: %d [%d]\n", MI_NAME_MAP[mi], mi, c.instructionRegister, c.step, i)
			// fmt.Printf("** [ %-16s ] mi: %d, inst: %2x, step: %d [%d]\n", MI_NAME_MAP[mi], mi, c.instructionRegister, c.step, i)
			break
		}
		if c.debug {
			if c.verbose {
				c.Logf("# [ %-16s ] pc: %04x, step: %d, inst_r: %02x, op_r: %02x, addr: %04x, data: %02x, bus_x: %02x, bus_y: %02x, rw: %x, status: %08b, regs: %s\n", MI_NAME_MAP[mi], c.programCounter, c.step, c.instructionRegister, c.operandRegister, c.addrBus.Read(), c.dataBus.Read(), c.busX.Read(), c.busY.Read(), c.rw, c.status.Flag(), c.registers)
				// fmt.Printf("# [ %-16s ] pc: %04x, step: %d, inst_r: %02x, op_r: %02x, addr: %04x, data: %02x, bus_x: %02x, bus_y: %02x, rw: %x, status: %08b, regs: %s\n", MI_NAME_MAP[mi], c.programCounter, c.step, c.instructionRegister, c.operandRegister, c.addrBus.Read(), c.dataBus.Read(), c.busX.Read(), c.busY.Read(), c.rw, c.status.Flag(), c.registers)
			}
		}
	}
	if c.debug {
		if !c.verbose {
			c.Logf("# pc: %04x, inst_r: %02x, op_r: %02x, addr: %04x, data: %02x, rw: %x, status: %08b, regs: %s\n", c.programCounter, c.instructionRegister, c.operandRegister, c.addrBus.Read(), c.dataBus.Read(), c.rw, c.status.Flag(), c.registers)
			// fmt.Printf("# pc: %04x, inst_r: %02x, op_r: %02x, addr: %04x, data: %02x, rw: %x, status: %08b, regs: %s\n", c.programCounter, c.instructionRegister, c.operandRegister, c.addrBus.Read(), c.dataBus.Read(), c.rw, c.status.Flag(), c.registers)
		} else {
			c.Logf("--")
			// fmt.Println("--")
		}
	}
	if len(microinstructions) == 0 || c.instructionRegister == instruction.Type(MI_BRK) {
		if c.debug {
			c.Logf(" ## BREAK ## ")
			// fmt.Println(" ## BREAK ## ")
		}
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

func (c *Comp) Reset() {
	c.Stop()
	for _, dev := range c.devices {
		dev.Clear()
	}
	c.registers.Write(register.IndexRegA, 0)
	c.registers.Write(register.IndexRegB, 0)
	c.registers.Write(register.IndexRegC, 0)
	c.registers.Write(register.IndexRegD, 0)
	c.status.Clear()
	c.step = 0
	c.instructionRegister = 0
	c.operandRegister = 0
	c.programCounter = 0
	c.memoryAddressRegister = 0
	c.stackPointer = StackStart
	c.memoryAddressRegister = 0
	c.store = 0
	c.rw = utils.IO_READ
	c.stopChan = make(chan struct{}, 1)
	c.running = false
	c.clear()
}

func (c *Comp) clear() {
	c.dataBus.Clear()
	c.busX.Clear()
	c.busY.Clear()
	c.addrBus.Clear()
	c.aluEnable = false
	c.aluStoreEnable = false
	c.aluDirectOut = false
	c.bridgeEnable = false
	c.bridgeDir = BRIDGE_DIR_IN
	c.rw = utils.IO_READ
}
