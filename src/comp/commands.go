package comp

import (
	"fmt"
	"os"
)

const (
	CMD_PC_OUT uint64 = 1 << iota
	CMD_PC_INC
	CMD_RAM_OUT
	CMD_INST_REG_IN
	CMD_OP_REG_IN
	CMD_STEP_INC
	CMD_STEP_CLR
	CMD_REG_1_IN
	CMD_REG_2_IN
	CMD_REG_1_OUT
	CMD_REG_2_OUT
	CMD_BRK
)

var (
	functions map[uint64]func(c *Comp) = map[uint64]func(c *Comp){
		CMD_PC_OUT:      CommandProgramCounterOut,
		CMD_RAM_OUT:     CommandRamOut,
		CMD_INST_REG_IN: CommandInstructionRegisterIn,
		CMD_OP_REG_IN:   CommandOperandRegisterIn,
		CMD_PC_INC:      CommandProgramCounterIncrement,
		CMD_STEP_INC:    CommandStepIncrement,
		CMD_STEP_CLR:    CommandStepClear,
		CMD_REG_1_IN:    CommandReg1In,
		CMD_REG_1_OUT:   CommandReg1Out,
		CMD_REG_2_IN:    CommandReg2In,
		CMD_REG_2_OUT:   CommandReg2Out,
		CMD_BRK:         CommandBrk,
	}
)

func CommandProgramCounterOut(c *Comp)       { c.memBus = c.pc }
func CommandRamOut(c *Comp)                  { c.dataBus = c.mem.Read(c.memBus) }
func CommandInstructionRegisterIn(c *Comp)   { c.inst = c.dataBus }
func CommandOperandRegisterIn(c *Comp)       { c.operand = c.dataBus }
func CommandProgramCounterIncrement(c *Comp) { c.pc++ }
func CommandStepIncrement(c *Comp)           { c.step++ }
func CommandStepClear(c *Comp)               { c.step = 0 }
func CommandReg1In(c *Comp)                  { c.registers.Write(c.operand&0xf, c.dataBus) }
func CommandReg1Out(c *Comp)                 { c.dataBus = c.registers.Read(c.operand & 0xf) }
func CommandReg2In(c *Comp)                  { c.registers.Write((c.operand&0xf0)>>4, c.dataBus) }
func CommandReg2Out(c *Comp)                 { c.dataBus = c.registers.Read((c.operand & 0xf0) >> 4) }
func CommandBrk(c *Comp)                     { fmt.Printf("## break at: %d\n", c.pc); os.Exit(0) }
