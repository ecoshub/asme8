package comp

import (
	"emu/src/instruction"
	"emu/src/status"
	"fmt"
)

func mInstAluAdd(c *Comp, command uint64, _ uint64) {
	if command&MI_ALU_ENABLE <= 0 {
		return
	}
	c.dataBus = c.busX + c.busY
	setFlags(c.status, uint16(c.busX)+uint16(c.busY))
}

func mInstAluSub(c *Comp, command uint64, _ uint64) {
	if command&MI_ALU_ENABLE <= 0 {
		return
	}
	c.dataBus = c.busX - c.busY
	setFlags(c.status, uint16(c.busX)-uint16(c.busY))
}

func mInstAluCmp(c *Comp, command uint64, _ uint64) {
	if command&MI_ALU_ENABLE <= 0 {
		return
	}
	// fmt.Println("c.busX, c.busY", c.busX, c.busY)
	setFlags(c.status, uint16(c.busX)-uint16(c.busY))
}

func mInstStatusControl(c *Comp, command, mask uint64) {
	statusMask := uint8(0)
	switch c.instructionRegister {
	case instruction.INST_JZ_INM:
		statusMask = status.STATUS_FLAG_ZERO
	}
	if c.status.IsSet(statusMask) {
		if c.debug {
			fmt.Println(" > jump for condition")
		}
		return
	}
	mInstProgramCounterInc(c, command, mask)
	mInstProgramCounterInc(c, command, mask)
	mInstStepClr(c, command, mask)
	if c.debug {
		fmt.Println(" > no jump")
	}
}

func setFlags(status *status.StatusRegister, result uint16) {
	// fmt.Printf("result %x\n", result)
	status.ClearOverflowFlag()
	status.ClearCarryFlag()
	status.ClearZeroFlag()
	status.ClearSignFlag()
	status.ClearParityFlag()

	if result > 0xff {
		status.SetOverflowFlag()
	}

	if result == 0x100 {
		status.SetCarryFlag()
	}

	if result == 0 {
		status.SetZeroFlag()
	}

	if (result&0x80)>>7 == 1 {
		status.SetSignFlag()
	}

	if result&1 == 0 {
		status.SetParityFlag()
	}
}
