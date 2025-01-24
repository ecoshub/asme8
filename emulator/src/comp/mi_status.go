package comp

import (
	"asme8/emulator/src/instruction"
	"asme8/emulator/src/status"
)

func mInstClearCarryFlag(c *Comp, mi uint64) {
	c.status.ClearCarryFlag()
}

func mInstStatusControl(c *Comp, mi uint64) {
	statusMask := uint8(0)
	not := false
	switch c.instructionRegister {
	case instruction.INST_JZ_INM:
		statusMask = status.STATUS_FLAG_ZERO
	case instruction.INST_JNZ_INM:
		statusMask = status.STATUS_FLAG_ZERO
		not = true
	case instruction.INST_JC_INM:
		statusMask = status.STATUS_FLAG_CARRY
	}
	if c.status.IsSet(statusMask) == !not {
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
