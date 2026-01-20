package computer

func mInstStackPointerOutAddr(c *Computer, _ uint64) {
	c.addrBus.Write_16(c.stackPointer)
}

func mInstStackPointerLowIn(c *Computer, _ uint64) {
	upper := c.stackPointer & 0xff00
	c.stackPointer = upper | c.inputBus.Read_16()
}

func mInstStackPointerHighIn(c *Computer, _ uint64) {
	lower := c.stackPointer & 0x00ff
	c.stackPointer = lower | (c.inputBus.Read_16() << 8)
}

func mInstStackPointerLowOut(c *Computer, _ uint64) {
	c.outputBus.Write_8(uint8(c.stackPointer))
	triggerBridge(c)
}

func mInstStackPointerHighOut(c *Computer, _ uint64) {
	c.outputBus.Write_8(uint8(c.stackPointer >> 8))
	triggerBridge(c)
}

func mInstStackPointerInc(c *Computer, _ uint64) {
	c.stackPointer++
}

func mInstStackPointerDec(c *Computer, _ uint64) {
	c.stackPointer--
}
