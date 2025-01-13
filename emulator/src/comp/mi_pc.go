package comp

func mInstProgramCounterIn(c *Comp, _ uint64) {
	c.programCounter = c.addrBus.Read()
}

func mInstProgramCounterOut(c *Comp, _ uint64) {
	c.addrBus.Write_16(c.programCounter)
}

func mInstProgramCounterInc(c *Comp, _ uint64) {
	c.programCounter++
}
