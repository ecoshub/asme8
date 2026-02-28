package computer

func mInstVec0OutLow(c *Computer, _ uint64) {
	c.addrBus.Write_16(IntVec0AddrLow)
}

func mInstVec0OutHigh(c *Computer, _ uint64) {
	c.addrBus.Write_16(IntVec0AddrHigh)
}
