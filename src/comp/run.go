package comp

import (
	"emu/src/mem"
	"emu/src/register"
	"fmt"

	"github.com/ecoshub/breakx"
)

const (
	DefaultMemorySize = 1 << 16
)

type Comp struct {
	registers             register.Module
	ram                   mem.Mem
	step                  uint8
	instructionRegister   uint8
	operandRegister       uint8
	programCounter        uint16
	memoryAddressRegister uint16
	addrBus               uint16
	memoryDataRegister    uint8
	dataBus               uint8
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
			microInstructions[mask](c)
		}
	}
}

func (c *Comp) Run() {
	for {
		fmt.Printf("step: %d, inst: 0x%x\n", c.step, c.instructionRegister)
		if c.step == 0 {
			c.PrintRegisters()
			c.run(M_INST_FETCH)
			continue
		}
		command := control[c.instructionRegister][c.step]
		breakx.Point(c.instructionRegister, c.step, command)
		c.run(command)
	}
}

func (c *Comp) PrintRegisters() {
	fmt.Println(c.registers)
}
