package comp

import (
	"fmt"
	"os"
)

const (
	// NOTE ORDER IS IMPORTANT!
	// *_OUT microinstructions must lesser mask than *_IN microinstructions
	MI_BRK uint64 = 1 << iota
	MI_PC_OUT
	MI_RAM_OUT
	MI_OPERAND_REG_1_OUT_X
	MI_OP_REG_1_OUT_Y
	MI_OPERAND_REG_2_OUT_X
	MI_OPERAND_REG_2_OUT_Y
	MI_MDR_OUT
	MI_ONES_OUT_X
	MI_ONES_OUT_Y
	MI_BRIDGE_ENABLE
	MI_BRIDGE_DIR_OUT
	MI_BRIDGE_DIR_IN
	MI_INST_REG_IN
	MI_OP_REG_IN
	MI_OPERAND_REG_1_IN
	MI_OP_REG_2_IN
	MI_MDR_IN
	MI_PC_INC
	MI_STEP_INC
	MI_STEP_CLR
)

type runnerFunc func(c *Comp, command uint64, _ uint64)

var (
	microInstructions map[uint64]runnerFunc = map[uint64]runnerFunc{
		MI_BRK:                 mInstBreak,
		MI_BRIDGE_ENABLE:       nop,
		MI_BRIDGE_DIR_IN:       nop,
		MI_BRIDGE_DIR_OUT:      nop,
		MI_PC_OUT:              mInstProgramCounterOut,
		MI_RAM_OUT:             mInstRamOut,
		MI_INST_REG_IN:         mInstInstructionRegisterIn,
		MI_OP_REG_IN:           mInstOperandRegisterIn,
		MI_PC_INC:              mInstProgramCounterIncrement,
		MI_STEP_INC:            mInstStepIncrement,
		MI_STEP_CLR:            mInstStepClear,
		MI_OPERAND_REG_1_IN:    mInstReg1In,
		MI_OP_REG_2_IN:         mInstReg2In,
		MI_OPERAND_REG_1_OUT_X: mInstReg1OutX,
		MI_OP_REG_1_OUT_Y:      mInstReg1OutY,
		MI_OPERAND_REG_2_OUT_X: mInstReg2OutX,
		MI_OPERAND_REG_2_OUT_Y: mInstReg2OutY,
		MI_MDR_IN:              mInstMdrIn,
		MI_MDR_OUT:             mInstMdrOut,
		MI_ONES_OUT_X:          mInstOnesOutX,
		MI_ONES_OUT_Y:          mInstOnesOutY,
	}
)

func nop(_ *Comp, _ uint64, _ uint64) {}

func mInstInstructionRegisterIn(c *Comp, _ uint64, _ uint64) {
	c.instructionRegister = c.dataBus
}

func mInstOperandRegisterIn(c *Comp, _ uint64, _ uint64) {
	c.operandRegister = c.dataBus
}

func mInstProgramCounterOut(c *Comp, _ uint64, _ uint64) {
	c.addrBus = c.programCounter
}

func mInstProgramCounterIncrement(c *Comp, _ uint64, _ uint64) {
	c.programCounter++
}

func mInstStepIncrement(c *Comp, _ uint64, _ uint64) {
	c.step++
}

func mInstStepClear(c *Comp, _ uint64, _ uint64) {
	c.step = 0
}

func mInstRamOut(c *Comp, command uint64, _ uint64) {
	c.PutToDataBus(c.ram.Read(c.addrBus), command)
}

func mInstReg1In(c *Comp, _ uint64, _ uint64) {
	c.registers.Write(c.operandRegister&0xf, c.dataBus)
}

func mInstReg2In(c *Comp, _ uint64, _ uint64) {
	c.registers.Write((c.operandRegister&0xf0)>>4, c.dataBus)
}

func mInstReg1OutX(c *Comp, _ uint64, _ uint64) {
	c.busX = c.registers.Read(c.operandRegister & 0xf)
}

func mInstReg1OutY(c *Comp, command uint64, _ uint64) {
	c.PutToBusY(c.registers.Read(c.operandRegister&0xf), command)
}

func mInstReg2OutX(c *Comp, _ uint64, _ uint64) {
	c.busX = c.registers.Read((c.operandRegister & 0xf0) >> 4)
}

func mInstReg2OutY(c *Comp, command uint64, _ uint64) {
	c.PutToBusY(c.registers.Read((c.operandRegister&0xf0)>>4), command)
}

func mInstBreak(c *Comp, _ uint64, _ uint64) {
	fmt.Printf("## break at: %d\n", c.programCounter)
	os.Exit(0)
}

func mInstMdrIn(c *Comp, _ uint64, _ uint64) {
	c.memoryDataRegister = c.dataBus
}

func mInstMdrOut(c *Comp, _ uint64, _ uint64) {
	c.memoryDataRegister = c.dataBus
}

func mInstOnesOutX(c *Comp, _ uint64, _ uint64) {
	c.busX = 1
}

func mInstOnesOutY(c *Comp, command uint64, _ uint64) {
	c.PutToBusY(1, command)
}
