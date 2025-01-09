package comp

import "emu/src/instruction"

func mInstInstructionRegisterIn(c *Comp, _ uint64, _ uint64) {
	c.instructionRegister = instruction.Type(c.dataBus.Read())
}

func mInstOperandRegisterIn(c *Comp, _ uint64, _ uint64) {
	c.operandRegister = uint8(c.dataBus.Read())
}

func mInstMemoryDataRegisterIn(c *Comp, _ uint64, _ uint64) {
	c.memoryDataRegister = uint8(c.dataBus.Read())
}

func mInstMemoryDataRegisterOut(c *Comp, _ uint64, _ uint64) {
	c.memoryDataRegister = uint8(c.dataBus.Read())
}

func mInstMarLowIn(c *Comp, _ uint64, _ uint64) {
	upper := c.memoryAddressRegister & 0xff00
	c.memoryAddressRegister = upper | c.dataBus.Read()
}

func mInstMarHighIn(c *Comp, _ uint64, _ uint64) {
	lower := c.memoryAddressRegister & 0x00ff
	c.memoryAddressRegister = lower | (c.dataBus.Read() << 8)
}

func mInstMarOut(c *Comp, _ uint64, _ uint64) {
	c.addrBus.Write(c.memoryAddressRegister)
}
