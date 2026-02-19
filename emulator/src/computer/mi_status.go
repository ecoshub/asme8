package computer

import (
	"asme8/emulator/src/status"
)

func mInstStatusRegisterOut(c *Computer, _ uint64) {
	c.outputBus.Write_8(c.status.Flag())
}

func mInstStatusRegisterIn(c *Computer, _ uint64) {
	c.status.Set(c.inputBus.Read_8())
}

func mInstClearCarryFlag(c *Computer, _ uint64) {
	c.status.ClearCarryFlag()
}

func mInstClearInterruptEnableFlag(c *Computer, _ uint64) {
	c.status.ClearInterruptEnableFlag()
}

func mInstSetInterruptEnableFlag(c *Computer, _ uint64) {
	c.status.SetInterruptEnableFlag()
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
