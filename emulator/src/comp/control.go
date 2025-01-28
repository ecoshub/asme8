package comp

import (
	"asme8/emulator/src/instruction"
)

type step = uint8

var (
	MI_CMD_FETCH_INST    = []uint64{MI_BRIDGE_ENABLE, MI_PC_OUT_ADDR, MI_IO_READ, MI_INST_REG_IN, MI_PC_INC, MI_STEP_INC}
	MI_CMD_FETCH_OPERAND = []uint64{MI_BRIDGE_ENABLE, MI_PC_OUT_ADDR, MI_IO_READ, MI_OP_REG_IN, MI_PC_INC, MI_STEP_INC}
)

var (
	GENERIC_CONDITIONAL_JMP_AND_NOT_JMP map[step][]uint64 = map[step][]uint64{
		0: MI_CMD_FETCH_INST,
		1: {MI_JMP_CTRL},
		2: {MI_BRIDGE_ENABLE, MI_PC_OUT_ADDR, MI_IO_READ, MI_MAR_L_IN, MI_PC_INC, MI_STEP_INC},
		3: {MI_BRIDGE_ENABLE, MI_PC_OUT_ADDR, MI_IO_READ, MI_MAR_H_IN, MI_PC_INC, MI_STEP_INC},
		4: {MI_MAR_OUT_ADDR, MI_PC_IN_ADDR, MI_STEP_CLR},
	}
)

var (
	CONTROL_SIGNALS map[instruction.Type]map[step][]uint64 = map[instruction.Type]map[step][]uint64{
		instruction.INST_BRK: {
			0: MI_CMD_FETCH_INST,
			1: {MI_BRK},
		},
		instruction.INST_NOP: {
			0: MI_CMD_FETCH_INST,
			1: {MI_STEP_CLR},
		},
		instruction.INST_MOV_INM: {
			0: MI_CMD_FETCH_INST,
			1: MI_CMD_FETCH_OPERAND,
			2: {MI_BRIDGE_ENABLE, MI_PC_OUT_ADDR, MI_IO_READ, MI_REG_OP_1_IN, MI_PC_INC, MI_STEP_CLR},
		},
		instruction.INST_MOV_REG_REG: {
			0: MI_CMD_FETCH_INST,
			1: MI_CMD_FETCH_OPERAND,
			2: {MI_BRIDGE_ENABLE, MI_REG_OP_2_OUT, MI_REG_OP_1_IN, MI_STEP_CLR},
		},
		instruction.INST_MOV_REG_MEM: {
			0: MI_CMD_FETCH_INST,
			1: MI_CMD_FETCH_OPERAND,
			2: {MI_BRIDGE_ENABLE, MI_PC_OUT_ADDR, MI_IO_READ, MI_MAR_L_IN, MI_PC_INC, MI_STEP_INC},
			3: {MI_BRIDGE_ENABLE, MI_PC_OUT_ADDR, MI_IO_READ, MI_MAR_H_IN, MI_PC_INC, MI_STEP_INC},
			4: {MI_REG_OP_1_OUT, MI_MAR_OUT_ADDR, MI_IO_WRITE, MI_STEP_CLR},
		},
		instruction.INST_MOV_MEM_REG: {
			0: MI_CMD_FETCH_INST,
			1: MI_CMD_FETCH_OPERAND,
			2: {MI_BRIDGE_ENABLE, MI_PC_OUT_ADDR, MI_IO_READ, MI_MAR_L_IN, MI_PC_INC, MI_STEP_INC},
			3: {MI_BRIDGE_ENABLE, MI_PC_OUT_ADDR, MI_IO_READ, MI_MAR_H_IN, MI_PC_INC, MI_STEP_INC},
			4: {MI_BRIDGE_ENABLE, MI_MAR_OUT_ADDR, MI_IO_READ, MI_REG_OP_1_IN, MI_STEP_CLR},
		},
		instruction.INST_MOV_MEM_REG_OFFSET: {
			0: MI_CMD_FETCH_INST,
			1: MI_CMD_FETCH_OPERAND,
			2: {MI_PC_OUT_ADDR, MI_IO_READ, MI_REG_OP_1_OUT_ALU, MI_ALU_OUT, MI_ALU_ADD, MI_MAR_L_IN, MI_PC_INC, MI_STEP_INC},
			3: {MI_PC_OUT_ADDR, MI_IO_READ, MI_ALU_OUT, MI_ALU_ADC, MI_MAR_H_IN, MI_PC_INC, MI_STEP_INC},
			4: {MI_BRIDGE_ENABLE, MI_MAR_OUT_ADDR, MI_IO_READ, MI_REG_OP_2_IN, MI_STEP_CLR},
		},
		instruction.INST_MOV_REG_MEM_OFFSET: {
			0: MI_CMD_FETCH_INST,
			1: MI_CMD_FETCH_OPERAND,
			2: {MI_PC_OUT_ADDR, MI_IO_READ, MI_REG_OP_1_OUT_ALU, MI_ALU_OUT, MI_ALU_ADD, MI_MAR_L_IN, MI_PC_INC, MI_STEP_INC},
			3: {MI_PC_OUT_ADDR, MI_IO_READ, MI_ALU_OUT, MI_ALU_ADC, MI_MAR_H_IN, MI_PC_INC, MI_STEP_INC},
			4: {MI_REG_OP_2_OUT, MI_MAR_OUT_ADDR, MI_IO_WRITE, MI_STEP_CLR},
		},
		instruction.INST_JMP_INM: {
			0: MI_CMD_FETCH_INST,
			1: {MI_BRIDGE_ENABLE, MI_PC_OUT_ADDR, MI_IO_READ, MI_MAR_L_IN, MI_PC_INC, MI_STEP_INC},
			2: {MI_BRIDGE_ENABLE, MI_PC_OUT_ADDR, MI_IO_READ, MI_MAR_H_IN, MI_PC_INC, MI_STEP_INC},
			3: {MI_MAR_OUT_ADDR, MI_PC_IN_ADDR, MI_STEP_CLR},
		},
		instruction.INST_CMP_REG_INM: {
			0: MI_CMD_FETCH_INST,
			1: MI_CMD_FETCH_OPERAND,
			2: {MI_PC_OUT_ADDR, MI_IO_READ, MI_REG_OP_1_OUT_ALU, MI_ALU_CMP, MI_PC_INC, MI_STEP_CLR},
		},
		instruction.INST_CMP_RR: {
			0: MI_CMD_FETCH_INST,
			1: MI_CMD_FETCH_OPERAND,
			2: {MI_REG_OP_1_OUT_ALU, MI_REG_OP_2_OUT, MI_ALU_CMP, MI_STEP_CLR},
		},
		instruction.INST_ADD_REG_INM: {
			0: MI_CMD_FETCH_INST,
			1: MI_CMD_FETCH_OPERAND,
			2: {MI_PC_OUT_ADDR, MI_IO_READ, MI_REG_OP_1_OUT_ALU, MI_ALU_OUT, MI_ALU_ADD, MI_REG_OP_1_IN, MI_PC_INC, MI_STEP_CLR},
		},
		instruction.INST_ADD_RR: {
			0: MI_CMD_FETCH_INST,
			1: MI_CMD_FETCH_OPERAND,
			2: {MI_REG_OP_1_OUT_ALU, MI_REG_OP_2_OUT, MI_ALU_OUT, MI_ALU_ADD, MI_REG_OP_1_IN, MI_STEP_CLR},
		},
		instruction.INST_SUB_REG_INM: {
			0: MI_CMD_FETCH_INST,
			1: MI_CMD_FETCH_OPERAND,
			2: {MI_PC_OUT_ADDR, MI_IO_READ, MI_REG_OP_1_OUT_ALU, MI_ALU_OUT, MI_ALU_SUB, MI_REG_OP_1_IN, MI_PC_INC, MI_STEP_CLR},
		},
		instruction.INST_SUB_RR: {
			0: MI_CMD_FETCH_INST,
			1: MI_CMD_FETCH_OPERAND,
			2: {MI_REG_OP_1_OUT_ALU, MI_REG_OP_2_OUT, MI_ALU_OUT, MI_ALU_SUB, MI_REG_OP_1_IN, MI_STEP_CLR},
		},
		instruction.INST_PUSH_REG: {
			0: MI_CMD_FETCH_INST,
			1: MI_CMD_FETCH_OPERAND,
			2: {MI_REG_OP_1_OUT, MI_SP_OUT_ADDR, MI_IO_WRITE, MI_SP_DEC, MI_STEP_CLR},
		},
		instruction.INST_POP_REG: {
			0: MI_CMD_FETCH_INST,
			1: MI_CMD_FETCH_OPERAND,
			2: {MI_BRIDGE_ENABLE, MI_SP_INC, MI_SP_OUT_ADDR, MI_IO_READ, MI_REG_OP_1_IN, MI_STEP_CLR},
		},
		instruction.INST_PUSH_INM: {
			0: MI_CMD_FETCH_INST,
			1: {MI_BRIDGE_ENABLE, MI_PC_OUT_ADDR, MI_IO_READ, MI_MDR_IN, MI_PC_INC, MI_STEP_INC},
			2: {MI_MDR_OUT, MI_SP_OUT_ADDR, MI_IO_WRITE, MI_SP_DEC, MI_STEP_CLR},
		},
		instruction.INST_PUSH_SP: {
			0: MI_CMD_FETCH_INST,
			1: MI_CMD_FETCH_OPERAND,
			2: {MI_SP_L_OUT, MI_SP_OUT_ADDR, MI_IO_WRITE, MI_SP_DEC, MI_STEP_INC},
			3: {MI_SP_H_OUT, MI_SP_OUT_ADDR, MI_IO_WRITE, MI_SP_DEC, MI_STEP_CLR},
		},
		instruction.INST_POP_SP: {
			0: MI_CMD_FETCH_INST,
			1: MI_CMD_FETCH_OPERAND,
			2: {MI_BRIDGE_ENABLE, MI_SP_INC, MI_SP_OUT_ADDR, MI_IO_READ, MI_SP_H_IN, MI_STEP_INC},
			3: {MI_BRIDGE_ENABLE, MI_SP_INC, MI_SP_OUT_ADDR, MI_IO_READ, MI_SP_L_IN, MI_STEP_CLR},
		},
		instruction.INST_INC: {
			0: MI_CMD_FETCH_INST,
			1: MI_CMD_FETCH_OPERAND,
			2: {MI_REG_OP_1_OUT_ALU, MI_ONES_OUT, MI_ALU_OUT, MI_ALU_ADD, MI_REG_OP_1_IN, MI_STEP_CLR},
		},
		instruction.INST_DEC: {
			0: MI_CMD_FETCH_INST,
			1: MI_CMD_FETCH_OPERAND,
			2: {MI_REG_OP_1_OUT_ALU, MI_ONES_OUT, MI_ALU_OUT, MI_ALU_SUB, MI_REG_OP_1_IN, MI_STEP_CLR},
		},
		instruction.INST_JSR: {
			0: MI_CMD_FETCH_INST,
			1: {MI_BRIDGE_ENABLE, MI_PC_OUT_ADDR, MI_IO_READ, MI_MAR_L_IN, MI_PC_INC, MI_STEP_INC},
			2: {MI_BRIDGE_ENABLE, MI_PC_OUT_ADDR, MI_IO_READ, MI_MAR_H_IN, MI_PC_INC, MI_STEP_INC},
			3: {MI_PC_L_OUT, MI_SP_OUT_ADDR, MI_IO_WRITE, MI_SP_DEC, MI_STEP_INC},
			4: {MI_PC_H_OUT, MI_SP_OUT_ADDR, MI_IO_WRITE, MI_SP_DEC, MI_STEP_INC},
			5: {MI_MAR_OUT_ADDR, MI_PC_IN_ADDR, MI_STEP_CLR},
		},
		instruction.INST_RTS: {
			0: MI_CMD_FETCH_INST,
			1: {MI_BRIDGE_ENABLE, MI_SP_INC, MI_SP_OUT_ADDR, MI_IO_READ, MI_MAR_H_IN, MI_STEP_INC},
			2: {MI_BRIDGE_ENABLE, MI_SP_INC, MI_SP_OUT_ADDR, MI_IO_READ, MI_MAR_L_IN, MI_STEP_INC},
			3: {MI_MAR_OUT_ADDR, MI_PC_IN_ADDR, MI_STEP_CLR},
		},
		instruction.INST_XOR_RR: {
			0: MI_CMD_FETCH_INST,
			1: MI_CMD_FETCH_OPERAND,
			2: {MI_REG_OP_1_OUT_ALU, MI_REG_OP_2_OUT, MI_ALU_OUT, MI_ALU_XOR, MI_REG_OP_1_IN, MI_STEP_CLR},
		},
		instruction.INST_ADD_REG_MEM: {
			0: MI_CMD_FETCH_INST,
			1: MI_CMD_FETCH_OPERAND,
			2: {MI_BRIDGE_ENABLE, MI_PC_OUT_ADDR, MI_IO_READ, MI_MAR_L_IN, MI_PC_INC, MI_STEP_INC},
			3: {MI_BRIDGE_ENABLE, MI_PC_OUT_ADDR, MI_IO_READ, MI_MAR_H_IN, MI_PC_INC, MI_STEP_INC},
			4: {MI_REG_OP_1_OUT_ALU, MI_MAR_OUT_ADDR, MI_IO_READ, MI_ALU_OUT, MI_ALU_ADD, MI_MDR_IN, MI_STEP_INC},
			5: {MI_MDR_OUT, MI_MAR_OUT_ADDR, MI_IO_WRITE, MI_STEP_CLR},
		},
		instruction.INST_ADD_MEM_REG: {
			0: MI_CMD_FETCH_INST,
			1: MI_CMD_FETCH_OPERAND,
			2: {MI_BRIDGE_ENABLE, MI_PC_OUT_ADDR, MI_IO_READ, MI_MAR_L_IN, MI_PC_INC, MI_STEP_INC},
			3: {MI_BRIDGE_ENABLE, MI_PC_OUT_ADDR, MI_IO_READ, MI_MAR_H_IN, MI_PC_INC, MI_STEP_INC},
			4: {MI_REG_OP_1_OUT_ALU, MI_MAR_OUT_ADDR, MI_IO_READ, MI_ALU_OUT, MI_ALU_ADD, MI_REG_OP_1_IN, MI_STEP_CLR},
		},
		instruction.INST_ADD_SP: {
			0: MI_CMD_FETCH_INST,
			1: {MI_BRIDGE_ENABLE, MI_PC_OUT_ADDR, MI_IO_READ, MI_MDR_IN, MI_PC_INC, MI_STEP_INC},
			2: {MI_SP_L_OUT, MI_MDR_OUT_ALU, MI_ALU_OUT, MI_ALU_ADD, MI_SP_L_IN, MI_STEP_INC},
			3: {MI_SP_H_OUT, MI_ALU_OUT, MI_ALU_ADC, MI_SP_H_IN, MI_STEP_CLR},
		},
		instruction.INST_SUB_SP: {
			0: MI_CMD_FETCH_INST,
			1: {MI_BRIDGE_ENABLE, MI_SP_L_OUT, MI_MDR_IN, MI_STEP_INC},
			2: {MI_PC_OUT_ADDR, MI_IO_READ, MI_MDR_OUT_ALU, MI_ALU_OUT, MI_ALU_SUB, MI_SP_L_IN, MI_PC_INC, MI_STEP_INC},
			3: {MI_BRIDGE_ENABLE, MI_SP_H_OUT, MI_MDR_IN, MI_STEP_INC},
			4: {MI_MDR_OUT_ALU, MI_ALU_OUT, MI_ALU_SBB, MI_SP_H_IN, MI_STEP_CLR},
		},
		instruction.INST_CLC: {
			0: MI_CMD_FETCH_INST,
			1: {MI_CLC, MI_STEP_CLR},
		},
		instruction.INST_JZ_INM:  GENERIC_CONDITIONAL_JMP_AND_NOT_JMP,
		instruction.INST_JNZ_INM: GENERIC_CONDITIONAL_JMP_AND_NOT_JMP,
		instruction.INST_JS_INM:  GENERIC_CONDITIONAL_JMP_AND_NOT_JMP,
		instruction.INST_JNS_INM: GENERIC_CONDITIONAL_JMP_AND_NOT_JMP,
		instruction.INST_JC_INM:  GENERIC_CONDITIONAL_JMP_AND_NOT_JMP,
	}
)
