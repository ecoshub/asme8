package comp

import (
	"asme8/emulator/src/status"
)

const (
	OPERATION_UNKNOWN uint8 = 0
	OPERATION_PLUS    uint8 = 1
	OPERATION_MINUS   uint8 = 2
	OPERATION_XOR     uint8 = 3
)

func mInstAluOut(c *Comp, _ uint64) {
	c.aluOut = true
}

func mInstAluAdd(c *Comp, mi uint64) {
	AluOP(c, mi, OPERATION_PLUS, false)
}

func mInstAluAdc(c *Comp, mi uint64) {
	AluOP(c, mi, OPERATION_PLUS, true)
}

func mInstAluSub(c *Comp, mi uint64) {
	AluOP(c, mi, OPERATION_MINUS, false)
}

func mInstAluSbb(c *Comp, mi uint64) {
	AluOP(c, mi, OPERATION_MINUS, true)
}

func mInstAluCmp(c *Comp, mi uint64) {
	setFlags(c.status, OPERATION_MINUS, c.aluBus.Read_8(), c.outputBus.Read_8(), false)
}

func mInstAluXor(c *Comp, mi uint64) {
	AluOP(c, mi, OPERATION_XOR, false)
}

func AluOP(c *Comp, mi uint64, op uint8, withCarry bool) {
	result := doOperation(c.status, op, c.aluBus.Read_8(), c.outputBus.Read_8(), withCarry)
	if !c.aluOut {
		return
	}
	// fmt.Printf("ALU OP. op: %d, op_1: %02x, op_2: %02x, carry: %5v, result: %02x, flags: %s\n", op, c.aluBus.Read_8(), c.outputBus.Read_8(), withCarry, result, c.status)
	c.inputBus.Write_8(result)
	triggerBridge(c)
}

func doOperation(s *status.StatusRegister, op uint8, a uint8, b uint8, withCarry bool) uint8 {
	carry := uint8(0)
	if withCarry && s.IsSet(status.STATUS_FLAG_CARRY) {
		carry = 1
	}
	setFlags(s, op, a, b, withCarry)
	switch op {
	case OPERATION_MINUS:
		return a - b - carry
	case OPERATION_PLUS:
		return a + b + carry
	case OPERATION_XOR:
		return a ^ b
	default:
		return 0
	}
}
