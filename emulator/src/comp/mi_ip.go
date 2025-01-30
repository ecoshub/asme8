package comp

import "asme8/common/instruction"

func mInstIndexPointerOutAddr(c *Comp, _ uint64) {
	ipl := c.registers.Read_16(instruction.REGISTER_OPCODE_IPL)
	iph := c.registers.Read_16(instruction.REGISTER_OPCODE_IPH)
	val := ipl
	val |= (iph << 8)
	c.addrBus.Write_16(val)
}

func mInstIndexPointerInAddr(c *Comp, _ uint64) {
	data := c.addrBus.Read_16()
	c.registers.Write(instruction.REGISTER_OPCODE_IPL, uint8(data))
	c.registers.Write(instruction.REGISTER_OPCODE_IPH, uint8(data<<8))
}

func mInstIndexPointerLowIn(c *Comp, _ uint64) {
	data := c.inputBus.Read_8()
	c.registers.Write(instruction.REGISTER_OPCODE_IPL, uint8(data))
}

func mInstIndexPointerHighIn(c *Comp, _ uint64) {
	data := c.inputBus.Read_8()
	c.registers.Write(instruction.REGISTER_OPCODE_IPH, uint8(data))
}

func mInstIndexPointerLowOut(c *Comp, _ uint64) {
	data := c.registers.Read_8(instruction.REGISTER_OPCODE_IPL)
	c.outputBus.Write_8(data)
	triggerBridge(c)
}

func mInstIndexPointerHighOut(c *Comp, _ uint64) {
	data := c.registers.Read_8(instruction.REGISTER_OPCODE_IPH)
	c.outputBus.Write_8(data)
	triggerBridge(c)
}
