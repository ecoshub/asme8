package comp

func mInstMemoryAddressRegisterOutAddr(c *Comp, _ uint64) {
	c.addrBus.Write_16(c.memoryAddressRegister)
}

func mInstMemoryAddressRegisterLowIn(c *Comp, _ uint64) {
	upper := c.memoryAddressRegister & 0xff00
	c.memoryAddressRegister = upper | c.inputBus.Read_16()
}

func mInstMemoryAddressRegisterHighIn(c *Comp, _ uint64) {
	lower := c.memoryAddressRegister & 0x00ff
	c.memoryAddressRegister = lower | (c.inputBus.Read_16() << 8)
}
