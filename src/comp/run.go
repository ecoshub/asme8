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
	pc        uint16
	step      byte
	inst      byte
	operand   byte
	mar       uint16
	registers register.Module
	mem       mem.Mem
	addrBus   uint16
	dataBus   byte
}

func New() *Comp {
	return &Comp{
		registers: register.NewModule(),
		mem:       mem.New(DefaultMemorySize),
	}
}

func (c *Comp) Put(offset int, program []byte) {
	copy(c.mem[offset:offset+len(program)], program[:])
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
		fmt.Printf("step: %d, inst: 0x%x\n", c.step, c.inst)
		if c.step == 0 {
			c.PrintRegisters()
			c.run(M_INST_FETCH)
			continue
		}
		command := control[c.inst][c.step]
		breakx.Point(c.inst, c.step, command)
		c.run(command)
	}
}

func (c *Comp) PrintRegisters() {
	fmt.Println(c.registers)
}
