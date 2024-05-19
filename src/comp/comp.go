package comp

import (
	"emu/src/mem"
	"emu/src/register"
	"fmt"
)

const (
	DefaultMemorySize = 1 << 16
)

type Comp struct {
	registers register.Module
	ram       mem.Mem
	step      uint8

	instructionRegister uint8
	operandRegister     uint8

	programCounter        uint16
	memoryAddressRegister uint16
	memoryDataRegister    uint8

	busX    uint8
	busY    uint8
	dataBus uint8
	addrBus uint16

	debug   bool
	verbose bool
}

func New() *Comp {
	return &Comp{
		registers: register.NewModule(),
		ram:       mem.New(DefaultMemorySize),
	}
}

func (c *Comp) Put(offset int, program []uint8) {
	copy(c.ram[offset:offset+len(program)], program[:])
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
	for {
		command := control[c.instructionRegister][c.step]
		c.run(command)
		if c.debug {
			if c.verbose {
				fmt.Printf("# command: %028b, step: %d, r_inst: %02x, r_op:%02x, bus_a: %04x, bus_d: %02x, bus_x: %02x, bus_y: %02x, registers: %s\n", command, c.step, c.instructionRegister, c.operandRegister, c.addrBus, c.dataBus, c.busX, c.busY, c.registers)
			}
			fmt.Printf("# step: %d, r_inst: %02x, r_op:%02x, registers: %s\n", c.step, c.instructionRegister, c.operandRegister, c.registers)
			if c.step == 0 {
				fmt.Println(" --- ")
			}
		}
		c.clearBusses()
	}
}

// I am not a fun of that...
func (c *Comp) clearBusses() {
	c.dataBus = 0
	c.busX = 0
	c.busY = 0
	c.addrBus = 0
}

func (c *Comp) SetDebug(val bool) {
	c.debug = val
}

func (c *Comp) SetVerbose(val bool) {
	c.verbose = val
}

func (c *Comp) PrintRegisters() {
	fmt.Println(c.registers)
}

func (c *Comp) PrintBusses() {
	fmt.Printf("busses, data: %x, x: %x, y:%x\n", c.dataBus, c.busX, c.busY)
}
