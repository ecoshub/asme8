package comp

func mInstProgramCounterIn(c *Comp, _ uint64, _ uint64) {
	c.programCounter = c.addrBus.Read()
}

func mInstProgramCounterOut(c *Comp, _ uint64, _ uint64) {
	c.addrBus.Write(c.programCounter)
}

func mInstProgramCounterInc(c *Comp, _ uint64, _ uint64) {
	c.programCounter++
}
