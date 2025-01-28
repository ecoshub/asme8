package comp

func mInstMemoryDataRegisterIn(c *Comp, _ uint64) {
	c.memoryDataRegister = c.inputBus.Read_8()
}

func mInstMemoryDataRegisterOut(c *Comp, _ uint64) {
	c.outputBus.Write_8(c.memoryDataRegister)
	triggerBridge(c)
}

func mInstMemoryDataRegisterOutAlu(c *Comp, _ uint64) {
	c.aluBus.Write_8(c.memoryDataRegister)
}
