package computer

func mInstProgramCounterOutAddr(c *Computer, _ uint64) {
	c.addrBus.Write_16(c.programCounter)
}

func mInstProgramCounterInAddr(c *Computer, _ uint64) {
	c.programCounter = c.addrBus.Read_16()
}

func mInstProgramCounterLowOut(c *Computer, _ uint64) {
	c.outputBus.Write_8(uint8(c.programCounter))
	triggerBridge(c)
}

func mInstProgramCounterHighOut(c *Computer, _ uint64) {
	c.outputBus.Write_8(uint8(c.programCounter >> 8))
	triggerBridge(c)
}

func mInstProgramCounterInc(c *Computer, _ uint64) {
	c.programCounter++
}
