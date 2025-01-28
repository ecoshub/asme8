package comp

import (
	"asme8/emulator/src/status"
)

const (
	OPERATION_UNKNOWN uint8 = iota
	OPERATION_PLUS
	OPERATION_MINUS
	OPERATION_CMP
	OPERATION_XOR
	OPERATION_AND
	OPERATION_OR
	OPERATION_NOT
	OPERATION_SHL
	OPERATION_SHR
	OPERATION_ROL
	OPERATION_ROR
)

func mInstAluOut(c *Comp, _ uint64) {
	c.aluOut = true
}

func mInstAluAdd(c *Comp, mi uint64) {
	a := c.aluBus.Read_8()
	b := c.outputBus.Read_8()
	result := doOperation(c, OPERATION_PLUS, a, b, false)
	if !c.aluOut {
		return
	}
	c.inputBus.Write_8(result)
	triggerBridge(c)
}

func mInstAluAdc(c *Comp, mi uint64) {
	a := c.aluBus.Read_8()
	b := c.outputBus.Read_8()
	result := doOperation(c, OPERATION_PLUS, a, b, true)
	if !c.aluOut {
		return
	}
	c.inputBus.Write_8(result)
	triggerBridge(c)
}

func mInstAluSub(c *Comp, mi uint64) {
	a := c.aluBus.Read_8()
	b := c.outputBus.Read_8()
	result := doOperation(c, OPERATION_MINUS, a, b, false)
	if !c.aluOut {
		return
	}
	c.inputBus.Write_8(result)
	triggerBridge(c)
}

func mInstAluSbb(c *Comp, mi uint64) {
	a := c.aluBus.Read_8()
	b := c.outputBus.Read_8()
	result := doOperation(c, OPERATION_MINUS, a, b, true)
	if !c.aluOut {
		return
	}
	c.inputBus.Write_8(result)
	triggerBridge(c)
}

func mInstAluCmp(c *Comp, mi uint64) {
	a := c.aluBus.Read_8()
	b := c.outputBus.Read_8()
	doOperation(c, OPERATION_CMP, a, b, false)
}

func mInstAluXor(c *Comp, mi uint64) {
	a := c.aluBus.Read_8()
	b := c.outputBus.Read_8()
	result := doOperation(c, OPERATION_XOR, a, b, false)
	if !c.aluOut {
		return
	}
	c.inputBus.Write_8(result)
	triggerBridge(c)
}

func mInstAluAnd(c *Comp, mi uint64) {
	a := c.aluBus.Read_8()
	b := c.outputBus.Read_8()
	result := doOperation(c, OPERATION_AND, a, b, false)
	if !c.aluOut {
		return
	}
	c.inputBus.Write_8(result)
	triggerBridge(c)
}

func mInstAluOr(c *Comp, mi uint64) {
	a := c.aluBus.Read_8()
	b := c.outputBus.Read_8()
	result := doOperation(c, OPERATION_OR, a, b, false)
	if !c.aluOut {
		return
	}
	c.inputBus.Write_8(result)
	triggerBridge(c)
}

func mInstAluNot(c *Comp, mi uint64) {
	a := c.aluBus.Read_8()
	b := c.outputBus.Read_8()
	result := doOperation(c, OPERATION_NOT, a, b, false)
	if !c.aluOut {
		return
	}
	c.inputBus.Write_8(result)
	triggerBridge(c)
}

func mInstAluShl(c *Comp, mi uint64) {
	a := c.aluBus.Read_8()
	b := c.outputBus.Read_8()
	result := doOperation(c, OPERATION_SHL, a, b, true)
	if !c.aluOut {
		return
	}
	c.inputBus.Write_8(result)
	triggerBridge(c)
}

func mInstAluShr(c *Comp, mi uint64) {
	a := c.aluBus.Read_8()
	b := c.outputBus.Read_8()
	result := doOperation(c, OPERATION_SHR, a, b, true)
	if !c.aluOut {
		return
	}
	c.inputBus.Write_8(result)
	triggerBridge(c)
}

func mInstAluRol(c *Comp, mi uint64) {
	a := c.aluBus.Read_8()
	b := c.outputBus.Read_8()
	result := doOperation(c, OPERATION_ROL, a, b, true)
	if !c.aluOut {
		return
	}
	c.inputBus.Write_8(result)
	triggerBridge(c)
}

func mInstAluRor(c *Comp, mi uint64) {
	a := c.aluBus.Read_8()
	b := c.outputBus.Read_8()
	result := doOperation(c, OPERATION_ROR, a, b, true)
	if !c.aluOut {
		return
	}
	c.inputBus.Write_8(result)
	triggerBridge(c)
}

func doOperation(c *Comp, op uint8, a uint8, b uint8, useCarry bool) uint8 {
	carry := uint8(0)
	if useCarry && c.status.IsSet(status.STATUS_FLAG_CARRY) {
		carry = 1
	}
	// fmt.Printf("ALU OP. op: %d, op_1: %08b, op_2: %08b, use_carry: %5v, carry: %d, result: %08b, flags: %s\n", op, c.aluBus.Read_8(), c.outputBus.Read_8(), useCarry, carry, result, c.status)
	switch op {
	case OPERATION_MINUS:
		result := opSub(c.status, a, b, carry)
		setFlags(c.status, uint16(result))
		return uint8(result)
	case OPERATION_PLUS:
		result := opAdd(c.status, a, b, carry)
		setFlags(c.status, uint16(result))
		return uint8(result)
	case OPERATION_CMP:
		result := opSub(c.status, a, b, carry)
		setFlags(c.status, uint16(result))
		return uint8(result)
	case OPERATION_XOR:
		result := opXOR(c.status, a, b, carry)
		setFlags(c.status, uint16(result))
		return uint8(result)
	case OPERATION_AND:
		result := opAND(c.status, a, b, carry)
		setFlags(c.status, uint16(result))
		return uint8(result)
	case OPERATION_OR:
		result := opOR(c.status, a, b, carry)
		setFlags(c.status, uint16(result))
		return uint8(result)
	case OPERATION_NOT:
		result := opNOT(c.status, a, b, carry)
		setFlags(c.status, uint16(result))
		return uint8(result)
	case OPERATION_SHL:
		result := opSHL(c.status, a, b, carry)
		setFlags(c.status, uint16(result))
		return uint8(result)
	case OPERATION_SHR:
		result := opSHR(c.status, a, b, carry)
		setFlags(c.status, uint16(result))
		return uint8(result)
	case OPERATION_ROL:
		result := opROL(c.status, a, b, carry)
		setFlags(c.status, uint16(result))
		return uint8(result)
	case OPERATION_ROR:
		result := opROR(c.status, a, b, carry)
		setFlags(c.status, uint16(result))
		return uint8(result)
	}
	return 0
}

func opSub(s *status.StatusRegister, a uint8, b uint8, carry uint8) uint16 {
	result := uint16(a) - uint16(b) - uint16(carry)
	if a < b {
		s.SetCarryFlag()
	} else {
		s.ClearCarryFlag()
	}
	return result
}

func opAdd(s *status.StatusRegister, a uint8, b uint8, carry uint8) uint16 {
	result := uint16(a) + uint16(b) + uint16(carry)
	if result > 0xff {
		s.SetCarryFlag()
	} else {
		s.ClearCarryFlag()
	}
	return result
}

func opXOR(s *status.StatusRegister, a uint8, b uint8, carry uint8) uint16 {
	result := uint16(a) ^ uint16(b) - uint16(carry)
	s.ClearCarryFlag()
	return result
}

func opAND(s *status.StatusRegister, a uint8, b uint8, carry uint8) uint16 {
	result := uint16(a)&uint16(b) - uint16(carry)
	s.ClearCarryFlag()
	return result
}

func opOR(s *status.StatusRegister, a uint8, b uint8, carry uint8) uint16 {
	result := uint16(a) | uint16(b) - uint16(carry)
	s.ClearCarryFlag()
	return result
}

func opNOT(_ *status.StatusRegister, a uint8, _ uint8, _ uint8) uint16 {
	result := uint16(^a)
	return result
}

func opSHL(s *status.StatusRegister, a uint8, _ uint8, _ uint8) uint16 {
	willCarry := (a & 0b1000000) > 0
	result := a << 1
	if willCarry {
		s.SetCarryFlag()
	} else {
		s.ClearCarryFlag()
	}
	return uint16(result)
}

func opSHR(s *status.StatusRegister, a uint8, _ uint8, _ uint8) uint16 {
	willCarry := (a & 0b0000001) > 0
	result := a >> 1
	if willCarry {
		s.SetCarryFlag()
	} else {
		s.ClearCarryFlag()
	}
	return uint16(result)
}

func opROL(s *status.StatusRegister, a uint8, _ uint8, carry uint8) uint16 {
	willCarry := (a & 0b1000000) > 0
	result := a << 1
	result = result | carry
	if willCarry {
		s.SetCarryFlag()
	} else {
		s.ClearCarryFlag()
	}
	return uint16(result)
}

func opROR(s *status.StatusRegister, a uint8, _ uint8, carry uint8) uint16 {
	willCarry := (a & 0b0000001) > 0
	result := a >> 1
	result = result | (carry << 7)
	if willCarry {
		s.SetCarryFlag()
	} else {
		s.ClearCarryFlag()
	}
	return uint16(result)
}
