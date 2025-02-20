package comp

import (
	"asme8/common/instruction"
)

const (
	OP_LOW uint8 = 1 << iota
	OP_HIGH
	OP_1
	OP_2
	OP_INPUT
	OP_OUTPUT
	OP_OUTPUT_SPECIAL
)

func OperandOperation(command uint8) func(c *Comp, _ uint64) {
	return func(c *Comp, _ uint64) {
		var reg uint8
		if command&OP_1 > 0 {
			reg = c.operandRegister & 0xf
		}
		if command&OP_2 > 0 {
			reg = c.operandRegister & 0xf0 >> 4
		}
		var addr bool
		if command&OP_LOW > 0 {
			if reg == instruction.REGISTER_OPCODE_IP {
				reg = instruction.REGISTER_OPCODE_IPL
				addr = true
			}
			if reg == instruction.REGISTER_OPCODE_BP {
				reg = instruction.REGISTER_OPCODE_BPL
				addr = true
			}
		}
		if command&OP_HIGH > 0 {
			if reg == instruction.REGISTER_OPCODE_IP {
				reg = instruction.REGISTER_OPCODE_IPH
				addr = true
			}
			if reg == instruction.REGISTER_OPCODE_BP {
				reg = instruction.REGISTER_OPCODE_BPH
				addr = true
			}
		}
		if command&OP_INPUT > 0 {
			c.registers.Write(reg, c.inputBus.Read_8())
			return
		}
		if command&OP_OUTPUT > 0 {
			c.outputBus.Write_8(c.registers.Read_8(reg))
			triggerBridge(c)
			return
		}
		if command&OP_OUTPUT_SPECIAL > 0 {
			if addr {
				value := c.addrBus.Read_16()
				if command&OP_LOW > 0 {
					value = value & 0xf0
					regValue := c.registers.Read_16(reg)
					value = value | regValue
				}
				if command&OP_HIGH > 0 {
					value = value & 0xf
					regValue := c.registers.Read_16(reg) << 8
					value = value | regValue
				}
				c.addrBus.Write_16(value)
			} else {
				c.aluBus.Write_8(c.registers.Read_8(reg))
			}
		}
	}
}

func mInstOnesOutAlu(c *Comp, _ uint64) {
	c.aluBus.Write_8(1)
}

func mInstOnesOut(c *Comp, _ uint64) {
	c.outputBus.Write_8(1)
	triggerBridge(c)
}
