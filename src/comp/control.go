package comp

import (
	"emu/src/instruction"
)

type step = uint8

const (
	MI_CMD_FETCH_INST    = MI_PC_OUT | MI_IO_READ | MI_INST_REG_IN | MI_PC_INC | MI_STEP_INC
	MI_CMD_FETCH_OPERAND = MI_PC_OUT | MI_IO_READ | MI_OPERAND_REG_IN | MI_PC_INC | MI_STEP_INC
)

var (
	control map[instruction.Type]map[step]uint64 = map[instruction.Type]map[step]uint64{
		instruction.INST_BREAK: {
			0: MI_CMD_FETCH_INST,
			1: MI_BRK,
			2: MI_BRK,
			3: MI_BRK,
			4: MI_BRK,
			5: MI_BRK,
			6: MI_BRK,
			7: MI_BRK,
		},
		instruction.INST_JMP_INM: {
			0: MI_CMD_FETCH_INST, // fetch next instruction
			1: MI_PC_OUT | MI_IO_READ | MI_MAR_L_IN | MI_PC_INC | MI_STEP_INC,
			2: MI_PC_OUT | MI_IO_READ | MI_MAR_H_IN | MI_PC_INC | MI_STEP_INC,
			3: MI_MAR_OUT | MI_PC_IN | MI_STEP_CLR,
		},
		instruction.INST_JZ_INM: {
			0: MI_CMD_FETCH_INST, // fetch next instruction
			1: MI_JMP_CTRL | MI_STEP_INC,
			2: MI_PC_OUT | MI_IO_READ | MI_MAR_L_IN | MI_PC_INC | MI_STEP_INC,
			3: MI_PC_OUT | MI_IO_READ | MI_MAR_H_IN | MI_PC_INC | MI_STEP_INC,
			4: MI_MAR_OUT | MI_PC_IN | MI_STEP_CLR,
		},
		instruction.INST_MOV_INM_8: {
			0: MI_CMD_FETCH_INST,    // fetch next instruction
			1: MI_CMD_FETCH_OPERAND, // load operation register
			2: MI_PC_OUT | MI_IO_READ | MI_REG_OPERAND_1_IN | MI_PC_INC | MI_STEP_CLR,
		},
		instruction.INTS_MOV_REG_REG: {
			0: MI_CMD_FETCH_INST,    // fetch next instruction
			1: MI_CMD_FETCH_OPERAND, // load operation register
			2: MI_REG_OPERAND_2_OUT_DATA | MI_REG_OPERAND_1_IN | MI_STEP_CLR,
		},
		instruction.INST_MOV_MEM_REG: {
			0: MI_CMD_FETCH_INST,    // fetch next instruction
			1: MI_CMD_FETCH_OPERAND, // load operation register
			2: MI_PC_OUT | MI_IO_READ | MI_MAR_L_IN | MI_PC_INC | MI_STEP_INC,
			3: MI_PC_OUT | MI_IO_READ | MI_MAR_H_IN | MI_PC_INC | MI_STEP_INC,
			4: MI_MAR_OUT | MI_IO_READ | MI_REG_OPERAND_1_IN | MI_STEP_CLR,
		},
		instruction.INST_MOV_REG_MEM: {
			0: MI_CMD_FETCH_INST,    // fetch next instruction
			1: MI_CMD_FETCH_OPERAND, // load operation register
			2: MI_PC_OUT | MI_IO_READ | MI_MAR_L_IN | MI_PC_INC | MI_STEP_INC,
			3: MI_PC_OUT | MI_IO_READ | MI_MAR_H_IN | MI_PC_INC | MI_STEP_INC,
			4: MI_REG_OPERAND_1_OUT_DATA | MI_MAR_OUT | MI_IO_WRITE | MI_STEP_CLR,
		},
		instruction.INST_ADD_INM: {
			0: MI_CMD_FETCH_INST,    // fetch next instruction
			1: MI_CMD_FETCH_OPERAND, // load operation register
			2: MI_PC_OUT | MI_IO_READ | MI_REG_OPERAND_1_OUT_X | MI_ALU_ENABLE | MI_ALU_ADD | MI_ALU_STORE | MI_PC_INC | MI_STEP_INC,
			3: MI_ALU_OUT | MI_REG_OPERAND_1_IN | MI_STEP_CLR,
		},
		instruction.INST_ADD_RR: {
			0: MI_CMD_FETCH_INST,    // fetch next instruction
			1: MI_CMD_FETCH_OPERAND, // load operation register
			2: MI_REG_OPERAND_1_OUT_X | MI_REG_OPERAND_2_OUT_DATA | MI_ALU_ENABLE | MI_ALU_ADD | MI_REG_OPERAND_1_IN | MI_STEP_CLR,
		},
		instruction.INST_SUB_RR: {
			0: MI_CMD_FETCH_INST,    // fetch next instruction
			1: MI_CMD_FETCH_OPERAND, // load operation register
			2: MI_REG_OPERAND_1_OUT_X | MI_REG_OPERAND_2_OUT_DATA | MI_ALU_ENABLE | MI_ALU_SUB | MI_REG_OPERAND_1_IN | MI_STEP_CLR,
		},
		instruction.INST_CMP_RR: {
			0: MI_CMD_FETCH_INST,    // fetch next instruction
			1: MI_CMD_FETCH_OPERAND, // load operation register
			2: MI_REG_OPERAND_1_OUT_X | MI_REG_OPERAND_2_OUT_DATA | MI_ALU_ENABLE | MI_ALU_CMP | MI_STEP_CLR,
		},
		instruction.INST_CMP_INM: {
			0: MI_CMD_FETCH_INST,    // fetch next instruction
			1: MI_CMD_FETCH_OPERAND, // load operation register
			2: MI_PC_OUT | MI_IO_READ | MI_REG_OPERAND_1_OUT_X | MI_ALU_ENABLE | MI_ALU_CMP | MI_PC_INC | MI_STEP_CLR,
		},
		instruction.INST_NOP: {
			0: MI_CMD_FETCH_INST, // fetch next instruction
			1: MI_STEP_CLR,
		},
	}
)
