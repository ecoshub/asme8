package comp

func mInstProgramCounterIn(c *Comp, _ uint64) {
	c.programCounter = c.addrBus.Read()
}

func mInstProgramCounterLowOutData(c *Comp, _ uint64) {
	c.dataBus.Write_8(uint8(c.programCounter))
	triggerBridge(c)
}

func mInstProgramCounterHighOutData(c *Comp, _ uint64) {
	c.dataBus.Write_8(uint8(c.programCounter >> 8))
	triggerBridge(c)
}

func mInstProgramCounterOut(c *Comp, _ uint64) {
	c.addrBus.Write_16(c.programCounter)
}

func mInstProgramCounterInc(c *Comp, _ uint64) {
	c.programCounter++
}
