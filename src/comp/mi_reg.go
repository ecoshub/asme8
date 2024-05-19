package comp

func mInstReg1In(c *Comp, _ uint64, _ uint64) {
	c.registers.Write(c.operandRegister&0xf, c.dataBus)
}

func mInstReg2In(c *Comp, _ uint64, _ uint64) {
	c.registers.Write((c.operandRegister&0xf0)>>4, c.dataBus)
}

func mInstReg1OutX(c *Comp, _ uint64, _ uint64) {
	c.busX = c.registers.Read(c.operandRegister & 0xf)
}

func mInstReg1OutY(c *Comp, command uint64, _ uint64) {
	c.PutToBusY(c.registers.Read(c.operandRegister&0xf), command)
}

func mInstReg2OutX(c *Comp, _ uint64, _ uint64) {
	c.busX = c.registers.Read((c.operandRegister & 0xf0) >> 4)
}

func mInstReg2OutY(c *Comp, command uint64, _ uint64) {
	c.PutToBusY(c.registers.Read((c.operandRegister&0xf0)>>4), command)
}

func mInstOnesOutX(c *Comp, _ uint64, _ uint64) {
	c.busX = 1
}

func mInstOnesOutY(c *Comp, command uint64, _ uint64) {
	c.PutToBusY(1, command)
}
