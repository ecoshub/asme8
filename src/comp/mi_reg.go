package comp

func mInstReg1In(c *Comp, _ uint64, _ uint64) {
	c.registers.Write(c.operandRegister&0xf, uint8(c.dataBus.Read()))
}

func mInstReg2In(c *Comp, _ uint64, _ uint64) {
	c.registers.Write((c.operandRegister&0xf0)>>4, uint8(c.dataBus.Read()))
}

func mInstReg1OutX(c *Comp, _ uint64, _ uint64) {
	c.busX.Write_8(c.registers.Read(c.operandRegister & 0xf))
}

func mInstReg1OutData(c *Comp, command uint64, _ uint64) {
	c.dataBus.Write_8(c.registers.Read(c.operandRegister & 0xf))
}

func mInstReg2OutX(c *Comp, _ uint64, _ uint64) {
	c.busX.Write_8(c.registers.Read((c.operandRegister & 0xf0) >> 4))
}

func mInstReg2OutData(c *Comp, command uint64, _ uint64) {
	c.dataBus.Write_8(c.registers.Read((c.operandRegister & 0xf0) >> 4))
}

func mInstOnesOutX(c *Comp, _ uint64, _ uint64) {
	c.busX.Write_8(1)
}

func mInstOnesOutData(c *Comp, command uint64, _ uint64) {
	c.dataBus.Write_8(1)
}
