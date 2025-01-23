package comp

func mInstReg1In(c *Comp, _ uint64) {
	c.registers.Write(c.operandRegister&0xf, uint8(c.inputBus.Read_16()))
}

func mInstReg2In(c *Comp, _ uint64) {
	c.registers.Write((c.operandRegister&0xf0)>>4, uint8(c.inputBus.Read_16()))
}

func mInstReg1OutALU(c *Comp, _ uint64) {
	c.aluBus.Write_8(c.registers.Read(c.operandRegister & 0xf))
}

func mInstReg1Out(c *Comp, _ uint64) {
	c.outputBus.Write_8(c.registers.Read(c.operandRegister & 0xf))
	triggerBridge(c)
}

func mInstReg2OutAlu(c *Comp, _ uint64) {
	c.aluBus.Write_8(c.registers.Read((c.operandRegister & 0xf0) >> 4))
}

func mInstReg2Out(c *Comp, _ uint64) {
	c.outputBus.Write_8(c.registers.Read((c.operandRegister & 0xf0) >> 4))
	triggerBridge(c)
}

func mInstOnesOutAlu(c *Comp, _ uint64) {
	c.aluBus.Write_8(1)
}

func mInstOnesOut(c *Comp, _ uint64) {
	c.outputBus.Write_8(1)
	triggerBridge(c)
}
