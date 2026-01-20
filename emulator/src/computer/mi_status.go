package computer

import (
	"asme8/common/instruction"
	"asme8/emulator/src/status"
)

func mInstStatusRegisterOut(c *Computer, mi uint64) {
	c.outputBus.Write_8(c.status.Flag())
}

func mInstStatusRegisterIn(c *Computer, mi uint64) {
	c.status.Set(c.inputBus.Read_8())
}

func mInstClearCarryFlag(c *Computer, mi uint64) {
	c.status.ClearCarryFlag()
}

func mInstStatusControl(c *Computer, mi uint64) {
	statusMask := uint8(0)
	not := false
	switch c.instructionRegister {
	case instruction.INST_JZ_IMPL_IMM16:
		statusMask = status.STATUS_FLAG_ZERO
	case instruction.INST_JNZ_IMPL_IMM16:
		statusMask = status.STATUS_FLAG_ZERO
		not = true
	case instruction.INST_JC_IMPL_IMM16:
		statusMask = status.STATUS_FLAG_CARRY
	case instruction.INST_JNC_IMPL_IMM16:
		statusMask = status.STATUS_FLAG_CARRY
		not = true
	case instruction.INST_JS_IMPL_IMM16:
		statusMask = status.STATUS_FLAG_SIGN
	case instruction.INST_JNS_IMPL_IMM16:
		statusMask = status.STATUS_FLAG_SIGN
		not = true
	}

	if c.status.IsSet(statusMask) == !not {
		mInstStepInc(c, mi)
		return
	}
	mInstProgramCounterInc(c, mi)
	mInstProgramCounterInc(c, mi)
	mInstStepClr(c, mi)
}

func setFlags(s *status.StatusRegister, result uint16) {
	if result == 0 {
		s.SetZeroFlag()
	} else {
		s.ClearZeroFlag()
	}

	if int16(result) < 0 {
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
