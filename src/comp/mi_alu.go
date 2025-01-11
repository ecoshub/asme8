package comp

import (
	"emu/src/instruction"
	"emu/src/status"
	"fmt"
)

func mInstAluStore(c *Comp, command uint64, _ uint64) {
	if command&MI_ALU_ENABLE <= 0 {
		return
	}
	c.store = uint8(c.dataBus.Read())
}

func mInstAluOut(c *Comp, command uint64, _ uint64) {
	c.dataBus.Write(uint16(c.store))
}

func mInstAluAdd(c *Comp, command uint64, _ uint64) {
	if command&MI_ALU_ENABLE <= 0 {
		return
	}
	c.dataBus.Write(c.busX.Read() + c.dataBus.Read())
	setFlags(c.status, c.busX.Read()+c.dataBus.Read())
}

func mInstAluSub(c *Comp, command uint64, _ uint64) {
	if command&MI_ALU_ENABLE <= 0 {
		return
	}
	c.dataBus.Write(c.busX.Read() - c.dataBus.Read())
	setFlags(c.status, c.busX.Read()-c.dataBus.Read())
}

func mInstAluCmp(c *Comp, command uint64, _ uint64) {
	if command&MI_ALU_ENABLE <= 0 {
		return
	}
	setFlags(c.status, c.busX.Read()-c.dataBus.Read())
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
