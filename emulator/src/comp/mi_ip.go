package comp

func mInstIndexPointerOutAddr(c *Comp, _ uint64) {
	c.addrBus.Write_16(c.indexPointer)
}

func mInstIndexPointerInAddr(c *Comp, _ uint64) {
	c.indexPointer = c.addrBus.Read_16()
}

func mInstIndexPointerLowIn(c *Comp, _ uint64) {
	upper := c.indexPointer & 0xff00
	c.indexPointer = upper | c.inputBus.Read_16()
}

func mInstIndexPointerHighIn(c *Comp, _ uint64) {
	lower := c.indexPointer & 0x00ff
	c.indexPointer = lower | (c.inputBus.Read_16() << 8)
}

func mInstIndexPointerLowOut(c *Comp, _ uint64) {
	c.outputBus.Write_8(uint8(c.indexPointer))
	triggerBridge(c)
}

func mInstIndexPointerHighOut(c *Comp, _ uint64) {
	c.outputBus.Write_8(uint8(c.indexPointer >> 8))
	triggerBridge(c)
}
