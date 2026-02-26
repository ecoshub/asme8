package computer

func mInstStackPointerOutAddr(c *Computer, _ uint64) {
	stackPointer := c.GetStackPointer()
	c.addrBus.Write_16(stackPointer)
}

func mInstStackPointerLowIn(c *Computer, _ uint64) {
	stackPointer := c.GetStackPointer()
	upper := stackPointer & 0xff00
	stackPointer = upper | c.inputBus.Read_16()
	c.SetStackPointer(stackPointer)
}

func mInstStackPointerHighIn(c *Computer, _ uint64) {
	stackPointer := c.GetStackPointer()
	lower := stackPointer & 0x00ff
	stackPointer = lower | (c.inputBus.Read_16() << 8)
	c.SetStackPointer(stackPointer)
}

func mInstStackPointerLowOut(c *Computer, _ uint64) {
	stackPointer := c.GetStackPointer()
	c.outputBus.Write_8(uint8(stackPointer))
	triggerBridge(c)
}

func mInstStackPointerHighOut(c *Computer, _ uint64) {
	stackPointer := c.GetStackPointer()
	c.outputBus.Write_8(uint8(stackPointer >> 8))
	triggerBridge(c)
}

func mInstStackPointerInc(c *Computer, _ uint64) {
	stackPointer := c.GetStackPointer()
	stackPointer += 1
	c.SetStackPointer(stackPointer)
}

func mInstStackPointerDec(c *Computer, _ uint64) {
	stackPointer := c.GetStackPointer()
	stackPointer -= 1
	c.SetStackPointer(stackPointer)
}
