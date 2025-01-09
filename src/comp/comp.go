package comp

import (
	"emu/src/instruction"
	"emu/src/mem"
	"emu/src/register"
	"emu/src/status"
	"fmt"
	"time"
)

const (
	DefaultMemorySize = 1 << 16

	StackStart uint16 = 0x7fff
)

type Comp struct {
	registers register.Module
	ram       *mem.Mem
	status    *status.StatusRegister
	step      uint8

	instructionRegister instruction.Type
	operandRegister     uint8

	programCounter        uint16
	memoryAddressRegister uint16
	stackPointer          uint16
	memoryDataRegister    uint8

	busX    uint8
	busY    uint8
	dataBus uint8
	addrBus uint16

	debug   bool
	verbose bool
	delay   time.Duration
}

func New() *Comp {
	c := &Comp{
		registers:    register.NewModule(),
		status:       status.NewStatusRegister(),
		stackPointer: StackStart,
		ram:          mem.New(DefaultMemorySize),
	}
	ConfigureBIOS(c)
	return c
}

func (c *Comp) Load(offset int, program []uint8) {
	c.ram.Load(offset, program)
}

func (c *Comp) SetDebug(val bool) {
	c.debug = val
	c.ram.SetDebug(val)
}

func (c *Comp) SetVerbose(val bool) {
	c.verbose = val
}

func (c *Comp) SetDelayMS(delay time.Duration) {
	c.delay = delay
}

func (c *Comp) PrintRegisters() {
	fmt.Println(c.registers)
}

func (c *Comp) PrintBusses() {
	fmt.Printf("busses, data: %x, x: %x, y:%x\n", c.dataBus, c.busX, c.busY)
}

func (c *Comp) run(command uint64) {
	for i := 0; i < 64; i++ {
		mask := uint64(1) << i
		if command&mask > 0 {
			microInstructions[mask](c, command, mask)
		}
	}
}

func (c *Comp) Run() {
	t := time.NewTicker(c.delay)
	for range t.C {
		command := control[c.instructionRegister][c.step]
		if c.debug {
			if c.step == 0 {
				fmt.Println(" --- ")
			}
			if c.verbose {
				fmt.Printf("# command: %028b, step: %d, r_inst: %02x, r_op:%02x, bus_a: %04x, bus_d: %02x, bus_x: %02x, bus_y: %02x, registers: %s\n", command, c.step, c.instructionRegister, c.operandRegister, c.addrBus, c.dataBus, c.busX, c.busY, c.registers)
			}
			fmt.Printf("# pc: %02x, step: %d, r_inst: %02x, r_op:%02x, registers: %s, status: %08b\n", c.programCounter, c.step, c.instructionRegister, c.operandRegister, c.registers, c.status.Flag())
		}
		if command == MI_BRK {
			if c.debug {
				fmt.Println(" ** BREAK ** ")
			}
			break
		}
		c.run(command)
		c.clearBusses()
	}
}

func (c *Comp) clearBusses() {
	c.dataBus = 0
	c.busX = 0
	c.busY = 0
	c.addrBus = 0
}
