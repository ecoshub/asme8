package comp

import (
	"asme8/emulator/src/instruction"
)

func mInstInstructionRegisterIn(c *Comp, _ uint64) {
	c.instructionRegister = instruction.Type(c.outputBus.Read_8())
}

func mInstOperandRegisterIn(c *Comp, _ uint64) {
	c.operandRegister = c.outputBus.Read_8()
}
