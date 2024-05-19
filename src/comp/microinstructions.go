package comp

import (
	"fmt"
	"os"
)

const (
	M_INST_BRK uint64 = 1 << iota
	M_INST_PC_OUT
	M_INST_PC_INC
	M_INST_RAM_OUT
	M_INST_INST_REG_IN
	M_INST_OP_REG_IN
	M_INST_STEP_INC
	M_INST_STEP_CLR
	M_INST_REG_1_IN
	M_INST_REG_2_IN
	M_INST_REG_1_OUT
	M_INST_REG_2_OUT
	M_INST_MDR_IN
	M_INST_MDR_OUT
)

var (
	microInstructions map[uint64]func(c *Comp) = map[uint64]func(c *Comp){
		M_INST_BRK:         mInstBreak,
		M_INST_PC_OUT:      mInstProgramCounterOut,
		M_INST_RAM_OUT:     mInstRamOut,
		M_INST_INST_REG_IN: mInstInstructionRegisterIn,
		M_INST_OP_REG_IN:   mInstOperandRegisterIn,
		M_INST_PC_INC:      mInstProgramCounterIncrement,
		M_INST_STEP_INC:    mInstStepIncrement,
		M_INST_STEP_CLR:    mInstStepClear,
		M_INST_REG_1_IN:    mInstReg1In,
		M_INST_REG_1_OUT:   mInstReg1Out,
		M_INST_REG_2_IN:    mInstReg2In,
		M_INST_REG_2_OUT:   mInstReg2Out,
		M_INST_MDR_IN:      mInstMdrIn,
		M_INST_MDR_OUT:     mInstMdrOut,
	}
)

func mInstProgramCounterOut(c *Comp)       { c.addrBus = c.programCounter }
func mInstRamOut(c *Comp)                  { c.dataBus = c.ram.Read(c.addrBus) }
func mInstInstructionRegisterIn(c *Comp)   { c.instructionRegister = c.dataBus }
func mInstOperandRegisterIn(c *Comp)       { c.operandRegister = c.dataBus }
func mInstProgramCounterIncrement(c *Comp) { c.programCounter++ }
func mInstStepIncrement(c *Comp)           { c.step++ }
func mInstStepClear(c *Comp)               { c.step = 0 }
func mInstReg1In(c *Comp)                  { c.registers.Write(c.operandRegister&0xf, c.dataBus) }
func mInstReg1Out(c *Comp)                 { c.dataBus = c.registers.Read(c.operandRegister & 0xf) }
func mInstReg2In(c *Comp)                  { c.registers.Write((c.operandRegister&0xf0)>>4, c.dataBus) }
func mInstReg2Out(c *Comp)                 { c.dataBus = c.registers.Read((c.operandRegister & 0xf0) >> 4) }
func mInstBreak(c *Comp)                   { fmt.Printf("## break at: %d\n", c.programCounter); os.Exit(0) }
func mInstMdrIn(c *Comp)                   { c.memoryDataRegister = c.dataBus }
func mInstMdrOut(c *Comp)                  { c.memoryDataRegister = c.dataBus }
