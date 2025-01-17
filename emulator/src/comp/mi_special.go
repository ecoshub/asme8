package comp

import "asme8/emulator/src/instruction"

func mInstInstructionRegisterIn(c *Comp, _ uint64) {
	c.instructionRegister = instruction.Type(c.dataBus.Read_16())
}

func mInstOperandRegisterIn(c *Comp, _ uint64) {
	c.operandRegister = uint8(c.dataBus.Read_16())
}

func mInstMarLowIn(c *Comp, _ uint64) {
	upper := c.memoryAddressRegister & 0xff00
	c.memoryAddressRegister = upper | c.dataBus.Read_16()
}

func mInstMarHighIn(c *Comp, _ uint64) {
	lower := c.memoryAddressRegister & 0x00ff
	c.memoryAddressRegister = lower | (c.dataBus.Read_16() << 8)
}

func mInstMarOut(c *Comp, _ uint64) {
	c.addrBus.Write_16(c.memoryAddressRegister)
}
