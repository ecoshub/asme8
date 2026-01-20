package computer

import (
	"asme8/common/instruction"
)

func mInstInstructionRegisterIn(c *Computer, _ uint64) {
	c.instructionRegister = instruction.Type(c.outputBus.Read_8())
}

func mInstOperandRegisterIn(c *Computer, _ uint64) {
	c.operandRegister = c.outputBus.Read_8()
}
