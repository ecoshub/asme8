package computer

func mInstMemoryAddressRegisterOutAddr(c *Computer, _ uint64) {
	c.addrBus.Write_16(c.memoryAddressRegister)
}

func mInstMemoryAddressRegisterLowIn(c *Computer, _ uint64) {
	upper := c.memoryAddressRegister & 0xff00
	c.memoryAddressRegister = upper | c.inputBus.Read_16()
}

func mInstMemoryAddressRegisterHighIn(c *Computer, _ uint64) {
	lower := c.memoryAddressRegister & 0x00ff
	c.memoryAddressRegister = lower | (c.inputBus.Read_16() << 8)
}
