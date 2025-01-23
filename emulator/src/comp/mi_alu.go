package comp

import (
	"asme8/emulator/src/instruction"
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

func mInstStatusControl(c *Comp, mi uint64) {
	statusMask := uint8(0)
	isNot := false
	switch c.instructionRegister {
	case instruction.INST_JZ_INM:
		statusMask = status.STATUS_FLAG_ZERO
	case instruction.INST_JNZ_INM:
		statusMask = status.STATUS_FLAG_ZERO
		isNot = true
	}
	if c.status.IsSet(statusMask) == !isNot {
		mInstStepInc(c, mi)
		return
	}
	mInstProgramCounterInc(c, mi)
	mInstProgramCounterInc(c, mi)
	mInstStepClr(c, mi)
}

func setFlags(s *status.StatusRegister, op uint8, a uint8, b uint8, withCarry bool) {

	var result uint16
	var carry uint8

	if withCarry && s.IsSet(status.STATUS_FLAG_CARRY) {
		carry = 1
	}

	if op == OPERATION_PLUS {
		result = uint16(a) + uint16(b) + uint16(carry)
	}
	if op == OPERATION_MINUS {
		result = uint16(a) - uint16(b) - uint16(carry)
	}

	if op == OPERATION_PLUS {
		if result > 0xff {
			s.SetCarryFlag()
		} else {
			s.ClearCarryFlag()
		}
	}

	if op == OPERATION_MINUS {
		if a < b {
		} else {
			s.ClearCarryFlag()
		}
	}

	if result == 0 {
		s.SetZeroFlag()
	} else {
		s.ClearZeroFlag()
	}

	if (result&0x80)>>7 == 1 {
		s.SetSignFlag()
	} else {
		s.ClearSignFlag()
	}

	if parity(uint8(result)) {
		s.SetParityFlag()
	} else {
		s.ClearParityFlag()
	}

}

// Helper function to calculate parity (even or odd number of 1 bits)
func parity(value uint8) bool {
	// Return true if number of 1 bits is even
	count := 0
	for value > 0 {
		if value&1 == 1 {
			count++
		}
		value >>= 1
	}
	return count%2 == 0
}
