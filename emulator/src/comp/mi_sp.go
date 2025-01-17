package comp

func mInstStackPointerOut(c *Comp, _ uint64) {
	c.addrBus.Write_16(c.stackPointer)
}

func mInstStackPointerLowIn(c *Comp, _ uint64) {
	upper := c.stackPointer & 0xff00
	c.stackPointer = upper | c.dataBus.Read_16()
}

func mInstStackPointerHighIn(c *Comp, _ uint64) {
	lower := c.stackPointer & 0x00ff
	c.stackPointer = lower | (c.dataBus.Read_16() << 8)
}

func mInstStackPointerInc(c *Comp, _ uint64) {
	c.stackPointer++
}

func mInstStackPointerDec(c *Comp, _ uint64) {
	c.stackPointer--
}
