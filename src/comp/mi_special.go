package comp

func mInstInstructionRegisterIn(c *Comp, _ uint64, _ uint64) {
	c.instructionRegister = c.dataBus
}

func mInstOperandRegisterIn(c *Comp, _ uint64, _ uint64) {
	c.operandRegister = c.dataBus
}

func mInstMemoryDataRegisterIn(c *Comp, _ uint64, _ uint64) {
	c.memoryDataRegister = c.dataBus
}

func mInstMemoryDataRegisterOut(c *Comp, _ uint64, _ uint64) {
	c.memoryDataRegister = c.dataBus
}

func mInstMarLowIn(c *Comp, _ uint64, _ uint64) {
	upper := c.memoryAddressRegister & 0xff00
	c.memoryAddressRegister = upper | uint16(c.dataBus)
}

func mInstMarHighIn(c *Comp, _ uint64, _ uint64) {
	lower := c.memoryAddressRegister & 0x00ff
	c.memoryAddressRegister = lower | (uint16(c.dataBus) << 8)
}

func mInstMarOut(c *Comp, _ uint64, _ uint64) {
	c.addrBus = c.memoryAddressRegister
}
