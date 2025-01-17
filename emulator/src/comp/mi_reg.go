package comp

func mInstReg1In(c *Comp, _ uint64) {
	c.registers.Write(c.operandRegister&0xf, uint8(c.dataBus.Read_16()))
}

func mInstReg2In(c *Comp, _ uint64) {
	c.registers.Write((c.operandRegister&0xf0)>>4, uint8(c.dataBus.Read_16()))
}

func mInstReg1OutX(c *Comp, _ uint64) {
	c.busX.Write_8(c.registers.Read(c.operandRegister & 0xf))
}

func mInstReg1OutY(c *Comp, _ uint64) {
	c.busY.Write_8(c.registers.Read(c.operandRegister & 0xf))
	triggerBridge(c)
}

func mInstReg2OutX(c *Comp, _ uint64) {
	c.busX.Write_8(c.registers.Read((c.operandRegister & 0xf0) >> 4))
}

func mInstReg2OutY(c *Comp, _ uint64) {
	c.busY.Write_8(c.registers.Read((c.operandRegister & 0xf0) >> 4))
	triggerBridge(c)
}

func mInstOnesOutX(c *Comp, _ uint64) {
	c.busX.Write_8(1)
}

func mInstOnesOutY(c *Comp, _ uint64) {
	c.busY.Write_8(1)
	triggerBridge(c)
}
