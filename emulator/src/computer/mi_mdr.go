package computer

func mInstMemoryDataRegisterIn(c *Computer, _ uint64) {
	c.memoryDataRegister = c.inputBus.Read_8()
}

func mInstMemoryDataRegisterOutAlu(c *Computer, _ uint64) {
	c.aluBus.Write_8(c.memoryDataRegister)
}
