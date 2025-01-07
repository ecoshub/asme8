package comp

func mInstRamOut(c *Comp, command uint64, _ uint64) {
	c.PutToDataBus(c.ram.Read(c.addrBus), command)
}

func mInstRamIn(c *Comp, command uint64, _ uint64) {
	c.ram.Write(c.addrBus, c.dataBus)
}
