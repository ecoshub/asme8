package comp

func mInstRamOut(c *Comp, command uint64, _ uint64) {
	c.PutToDataBus(c.ram.Read(c.addrBus), command)
}
