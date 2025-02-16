package comp

import (
	"asme8/common/instruction"
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
		instruction.INST_BRK_IMPL: {
			0: MI_CMD_FETCH_INST,
			1: {MI_BRK},
		},
		instruction.INST_NOP_IMPL: {
			0: MI_CMD_FETCH_INST,
			1: {MI_STEP_CLR},
		},
		instruction.INST_MOV_REG_IMM: {
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
		instruction.INST_MOV_REG_SP: {
			0: MI_CMD_FETCH_INST,
			1: MI_CMD_FETCH_OPERAND,
			2: {MI_REG_OP_1_OUT, MI_SP_OUT_ADDR, MI_IO_WRITE, MI_STEP_CLR},
		},
		instruction.INST_MOV_SP_REG: {
			0: MI_CMD_FETCH_INST,
			1: MI_CMD_FETCH_OPERAND,
			2: {MI_BRIDGE_ENABLE, MI_SP_OUT_ADDR, MI_IO_READ, MI_REG_OP_1_IN, MI_STEP_CLR},
		},
		instruction.INST_MOV_REG_SP_OFFSET: {
			0: MI_CMD_FETCH_INST,
			1: MI_CMD_FETCH_OPERAND,
			2: {MI_BRIDGE_ENABLE, MI_PC_OUT_ADDR, MI_IO_READ, MI_MDR_IN, MI_PC_INC, MI_STEP_INC},
			3: {MI_SP_L_OUT, MI_MDR_OUT_ALU, MI_ALU_OUT, MI_ALU_ADD, MI_MAR_L_IN, MI_STEP_INC},
			4: {MI_SP_H_OUT, MI_ALU_OUT, MI_ALU_ADC, MI_MAR_H_IN, MI_STEP_INC},
			5: {MI_MAR_OUT_ADDR, MI_REG_OP_1_OUT, MI_IO_WRITE, MI_STEP_CLR},
		},
		instruction.INST_MOV_SP_REG_OFFSET: {
			0: MI_CMD_FETCH_INST,
			1: MI_CMD_FETCH_OPERAND,
			2: {MI_BRIDGE_ENABLE, MI_PC_OUT_ADDR, MI_IO_READ, MI_MDR_IN, MI_PC_INC, MI_STEP_INC},
			3: {MI_SP_L_OUT, MI_MDR_OUT_ALU, MI_ALU_OUT, MI_ALU_ADD, MI_MAR_L_IN, MI_STEP_INC},
			4: {MI_SP_H_OUT, MI_ALU_OUT, MI_ALU_ADC, MI_MAR_H_IN, MI_STEP_INC},
			5: {MI_BRIDGE_ENABLE, MI_MAR_OUT_ADDR, MI_IO_READ, MI_REG_OP_1_IN, MI_STEP_CLR},
		},
		instruction.INST_MOV_REG_SP_OFFSET_REG: {
			0: MI_CMD_FETCH_INST,
			1: MI_CMD_FETCH_OPERAND,
			2: {MI_SP_L_OUT, MI_REG_OP_2_OUT_ALU, MI_ALU_OUT, MI_ALU_ADD, MI_MAR_L_IN, MI_STEP_INC},
			3: {MI_SP_H_OUT, MI_ALU_OUT, MI_ALU_ADC, MI_MAR_H_IN, MI_STEP_INC},
			4: {MI_MAR_OUT_ADDR, MI_REG_OP_1_OUT, MI_IO_WRITE, MI_STEP_CLR},
		},
		instruction.INST_MOV_SP_REG_OFFSET_REG: {
			0: MI_CMD_FETCH_INST,
			1: MI_CMD_FETCH_OPERAND,
			2: {MI_SP_L_OUT, MI_REG_OP_2_OUT_ALU, MI_ALU_OUT, MI_ALU_ADD, MI_MAR_L_IN, MI_STEP_INC},
			3: {MI_SP_H_OUT, MI_ALU_OUT, MI_ALU_ADC, MI_MAR_H_IN, MI_STEP_INC},
			4: {MI_BRIDGE_ENABLE, MI_MAR_OUT_ADDR, MI_IO_READ, MI_REG_OP_1_IN, MI_STEP_CLR},
		},
		instruction.INST_MOV_MEM_IMM: {
			0: MI_CMD_FETCH_INST,
			1: {MI_BRIDGE_ENABLE, MI_PC_OUT_ADDR, MI_IO_READ, MI_MDR_IN, MI_PC_INC, MI_STEP_INC},
			2: {MI_BRIDGE_ENABLE, MI_PC_OUT_ADDR, MI_IO_READ, MI_MAR_L_IN, MI_PC_INC, MI_STEP_INC},
			3: {MI_BRIDGE_ENABLE, MI_PC_OUT_ADDR, MI_IO_READ, MI_MAR_H_IN, MI_PC_INC, MI_STEP_INC},
			4: {MI_MDR_OUT, MI_MAR_OUT_ADDR, MI_IO_WRITE, MI_STEP_CLR},
		},
		instruction.INST_JMP_IMM: {
			0: MI_CMD_FETCH_INST,
			1: {MI_BRIDGE_ENABLE, MI_PC_OUT_ADDR, MI_IO_READ, MI_MAR_L_IN, MI_PC_INC, MI_STEP_INC},
			2: {MI_BRIDGE_ENABLE, MI_PC_OUT_ADDR, MI_IO_READ, MI_MAR_H_IN, MI_PC_INC, MI_STEP_INC},
			3: {MI_MAR_OUT_ADDR, MI_PC_IN_ADDR, MI_STEP_CLR},
		},
		instruction.INST_CMP_REG_IMM: {
			0: MI_CMD_FETCH_INST,
			1: MI_CMD_FETCH_OPERAND,
			2: {MI_PC_OUT_ADDR, MI_IO_READ, MI_REG_OP_1_OUT_ALU, MI_ALU_CMP, MI_PC_INC, MI_STEP_CLR},
		},
		instruction.INST_CMP_REG_REG: {
			0: MI_CMD_FETCH_INST,
			1: MI_CMD_FETCH_OPERAND,
			2: {MI_REG_OP_1_OUT_ALU, MI_REG_OP_2_OUT, MI_ALU_CMP, MI_STEP_CLR},
		},
		instruction.INST_ADD_REG_IMM: {
			0: MI_CMD_FETCH_INST,
			1: MI_CMD_FETCH_OPERAND,
			2: {MI_PC_OUT_ADDR, MI_IO_READ, MI_REG_OP_1_OUT_ALU, MI_ALU_OUT, MI_ALU_ADD, MI_REG_OP_1_IN, MI_PC_INC, MI_STEP_CLR},
		},
		instruction.INST_ADD_REG_REG: {
			0: MI_CMD_FETCH_INST,
			1: MI_CMD_FETCH_OPERAND,
			2: {MI_REG_OP_1_OUT_ALU, MI_REG_OP_2_OUT, MI_ALU_OUT, MI_ALU_ADD, MI_REG_OP_1_IN, MI_STEP_CLR},
		},
		instruction.INST_ADC_REG_REG: {
			0: MI_CMD_FETCH_INST,
			1: MI_CMD_FETCH_OPERAND,
			2: {MI_REG_OP_1_OUT_ALU, MI_REG_OP_2_OUT, MI_ALU_OUT, MI_ALU_ADC, MI_REG_OP_1_IN, MI_STEP_CLR},
		},
		instruction.INST_ADC_REG_IMM: {
			0: MI_CMD_FETCH_INST,
			1: MI_CMD_FETCH_OPERAND,
			2: {MI_PC_OUT_ADDR, MI_IO_READ, MI_REG_OP_1_OUT_ALU, MI_ALU_OUT, MI_ALU_ADC, MI_REG_OP_1_IN, MI_PC_INC, MI_STEP_CLR},
		},
		instruction.INST_ADC_REG_MEM: {
			0: MI_CMD_FETCH_INST,
			1: MI_CMD_FETCH_OPERAND,
			2: {MI_BRIDGE_ENABLE, MI_PC_OUT_ADDR, MI_IO_READ, MI_MAR_L_IN, MI_PC_INC, MI_STEP_INC},
			3: {MI_BRIDGE_ENABLE, MI_PC_OUT_ADDR, MI_IO_READ, MI_MAR_H_IN, MI_PC_INC, MI_STEP_INC},
			4: {MI_REG_OP_1_OUT_ALU, MI_MAR_OUT_ADDR, MI_IO_READ, MI_ALU_OUT, MI_ALU_ADC, MI_MDR_IN, MI_STEP_INC},
			5: {MI_MDR_OUT, MI_MAR_OUT_ADDR, MI_IO_WRITE, MI_STEP_CLR},
		},
		instruction.INST_ADC_MEM_REG: {
			0: MI_CMD_FETCH_INST,
			1: MI_CMD_FETCH_OPERAND,
			2: {MI_BRIDGE_ENABLE, MI_PC_OUT_ADDR, MI_IO_READ, MI_MAR_L_IN, MI_PC_INC, MI_STEP_INC},
			3: {MI_BRIDGE_ENABLE, MI_PC_OUT_ADDR, MI_IO_READ, MI_MAR_H_IN, MI_PC_INC, MI_STEP_INC},
			4: {MI_REG_OP_1_OUT_ALU, MI_MAR_OUT_ADDR, MI_IO_READ, MI_ALU_OUT, MI_ALU_ADC, MI_MDR_IN, MI_STEP_INC},
			5: {MI_BRIDGE_ENABLE, MI_MDR_OUT, MI_MAR_OUT_ADDR, MI_REG_OP_1_IN, MI_STEP_CLR},
		},
		instruction.INST_SUB_REG_IMM: {
			0: MI_CMD_FETCH_INST,
			1: MI_CMD_FETCH_OPERAND,
			2: {MI_PC_OUT_ADDR, MI_IO_READ, MI_REG_OP_1_OUT_ALU, MI_ALU_OUT, MI_ALU_SUB, MI_REG_OP_1_IN, MI_PC_INC, MI_STEP_CLR},
		},
		instruction.INST_SUB_REG_REG: {
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
		instruction.INST_PUSH_IMM: {
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
		instruction.INST_INC_REG: {
			0: MI_CMD_FETCH_INST,
			1: MI_CMD_FETCH_OPERAND,
			2: {MI_REG_OP_1_OUT_ALU, MI_ONES_OUT, MI_ALU_OUT, MI_ALU_ADD, MI_REG_OP_1_IN, MI_STEP_CLR},
		},
		instruction.INST_DEC_REG: {
			0: MI_CMD_FETCH_INST,
			1: MI_CMD_FETCH_OPERAND,
			2: {MI_REG_OP_1_OUT_ALU, MI_ONES_OUT, MI_ALU_OUT, MI_ALU_SUB, MI_REG_OP_1_IN, MI_STEP_CLR},
		},
		instruction.INST_CALL_IMPL: {
			0: MI_CMD_FETCH_INST,
			1: {MI_BRIDGE_ENABLE, MI_PC_OUT_ADDR, MI_IO_READ, MI_MAR_L_IN, MI_PC_INC, MI_STEP_INC},
			2: {MI_BRIDGE_ENABLE, MI_PC_OUT_ADDR, MI_IO_READ, MI_MAR_H_IN, MI_PC_INC, MI_STEP_INC},
			3: {MI_PC_L_OUT, MI_SP_OUT_ADDR, MI_IO_WRITE, MI_SP_DEC, MI_STEP_INC},
			4: {MI_PC_H_OUT, MI_SP_OUT_ADDR, MI_IO_WRITE, MI_SP_DEC, MI_STEP_INC},
			5: {MI_MAR_OUT_ADDR, MI_PC_IN_ADDR, MI_STEP_CLR},
		},
		instruction.INST_RET_IMPL: {
			0: MI_CMD_FETCH_INST,
			1: {MI_BRIDGE_ENABLE, MI_SP_INC, MI_SP_OUT_ADDR, MI_IO_READ, MI_MAR_H_IN, MI_STEP_INC},
			2: {MI_BRIDGE_ENABLE, MI_SP_INC, MI_SP_OUT_ADDR, MI_IO_READ, MI_MAR_L_IN, MI_STEP_INC},
			3: {MI_MAR_OUT_ADDR, MI_PC_IN_ADDR, MI_STEP_CLR},
		},
		instruction.INST_XOR_REG_REG: {
			0: MI_CMD_FETCH_INST,
			1: MI_CMD_FETCH_OPERAND,
			2: {MI_REG_OP_1_OUT_ALU, MI_REG_OP_2_OUT, MI_ALU_OUT, MI_ALU_XOR, MI_REG_OP_1_IN, MI_STEP_CLR},
		},
		instruction.INST_XOR_REG_IMM: {
			0: MI_CMD_FETCH_INST,
			1: MI_CMD_FETCH_OPERAND,
			2: {MI_PC_OUT_ADDR, MI_IO_READ, MI_REG_OP_1_OUT_ALU, MI_ALU_OUT, MI_ALU_XOR, MI_REG_OP_1_IN, MI_PC_INC, MI_STEP_CLR},
		},
		instruction.INST_AND_REG_REG: {
			0: MI_CMD_FETCH_INST,
			1: MI_CMD_FETCH_OPERAND,
			2: {MI_REG_OP_1_OUT_ALU, MI_REG_OP_2_OUT, MI_ALU_OUT, MI_ALU_AND, MI_REG_OP_1_IN, MI_STEP_CLR},
		},
		instruction.INST_AND_REG_IMM: {
			0: MI_CMD_FETCH_INST,
			1: MI_CMD_FETCH_OPERAND,
			2: {MI_PC_OUT_ADDR, MI_IO_READ, MI_REG_OP_1_OUT_ALU, MI_ALU_OUT, MI_ALU_AND, MI_REG_OP_1_IN, MI_PC_INC, MI_STEP_CLR},
		},
		instruction.INST_OR_REG_REG: {
			0: MI_CMD_FETCH_INST,
			1: MI_CMD_FETCH_OPERAND,
			2: {MI_REG_OP_1_OUT_ALU, MI_REG_OP_2_OUT, MI_ALU_OUT, MI_ALU_OR, MI_REG_OP_1_IN, MI_STEP_CLR},
		},
		instruction.INST_OR_REG_IMM: {
			0: MI_CMD_FETCH_INST,
			1: MI_CMD_FETCH_OPERAND,
			2: {MI_PC_OUT_ADDR, MI_IO_READ, MI_REG_OP_1_OUT_ALU, MI_ALU_OUT, MI_ALU_OR, MI_REG_OP_1_IN, MI_PC_INC, MI_STEP_CLR},
		},
		instruction.INST_NOT_REG: {
			0: MI_CMD_FETCH_INST,
			1: MI_CMD_FETCH_OPERAND,
			2: {MI_REG_OP_1_OUT_ALU, MI_ALU_OUT, MI_ALU_NOT, MI_REG_OP_1_IN, MI_STEP_CLR},
		},
		instruction.INST_SHL_REG: {
			0: MI_CMD_FETCH_INST,
			1: MI_CMD_FETCH_OPERAND,
			2: {MI_REG_OP_1_OUT_ALU, MI_ALU_OUT, MI_ALU_SHL, MI_REG_OP_1_IN, MI_STEP_CLR},
		},
		instruction.INST_SHR_REG: {
			0: MI_CMD_FETCH_INST,
			1: MI_CMD_FETCH_OPERAND,
			2: {MI_REG_OP_1_OUT_ALU, MI_ALU_OUT, MI_ALU_SHR, MI_REG_OP_1_IN, MI_STEP_CLR},
		},
		instruction.INST_ROL_REG: {
			0: MI_CMD_FETCH_INST,
			1: MI_CMD_FETCH_OPERAND,
			2: {MI_REG_OP_1_OUT_ALU, MI_ALU_OUT, MI_ALU_ROL, MI_REG_OP_1_IN, MI_STEP_CLR},
		},
		instruction.INST_ROR_REG: {
			0: MI_CMD_FETCH_INST,
			1: MI_CMD_FETCH_OPERAND,
			2: {MI_REG_OP_1_OUT_ALU, MI_ALU_OUT, MI_ALU_ROR, MI_REG_OP_1_IN, MI_STEP_CLR},
		},
		instruction.INST_ADD_REG_MEM: {
			0: MI_CMD_FETCH_INST,
			1: MI_CMD_FETCH_OPERAND,
			2: {MI_BRIDGE_ENABLE, MI_PC_OUT_ADDR, MI_IO_READ, MI_MAR_L_IN, MI_PC_INC, MI_STEP_INC},
			3: {MI_BRIDGE_ENABLE, MI_PC_OUT_ADDR, MI_IO_READ, MI_MAR_H_IN, MI_PC_INC, MI_STEP_INC},
			4: {MI_REG_OP_1_OUT_ALU, MI_MAR_OUT_ADDR, MI_IO_READ, MI_ALU_OUT, MI_ALU_ADD, MI_MDR_IN, MI_STEP_INC},
			5: {MI_MDR_OUT, MI_MAR_OUT_ADDR, MI_IO_WRITE, MI_STEP_CLR},
		},
		instruction.INST_ADD_MEM_IMM: {
			0: MI_CMD_FETCH_INST,
			1: {MI_BRIDGE_ENABLE, MI_PC_OUT_ADDR, MI_IO_READ, MI_MDR_IN, MI_PC_INC, MI_STEP_INC},
			2: {MI_BRIDGE_ENABLE, MI_PC_OUT_ADDR, MI_IO_READ, MI_MAR_L_IN, MI_PC_INC, MI_STEP_INC},
			3: {MI_BRIDGE_ENABLE, MI_PC_OUT_ADDR, MI_IO_READ, MI_MAR_H_IN, MI_PC_INC, MI_STEP_INC},
			4: {MI_MDR_OUT_ALU, MI_MAR_OUT_ADDR, MI_IO_READ, MI_ALU_OUT, MI_ALU_ADD, MI_MDR_IN, MI_STEP_INC},
			5: {MI_MDR_OUT, MI_MAR_OUT_ADDR, MI_IO_WRITE, MI_STEP_CLR},
		},
		instruction.INST_ADC_MEM_IMM: {
			0: MI_CMD_FETCH_INST,
			1: {MI_BRIDGE_ENABLE, MI_PC_OUT_ADDR, MI_IO_READ, MI_MDR_IN, MI_PC_INC, MI_STEP_INC},
			2: {MI_BRIDGE_ENABLE, MI_PC_OUT_ADDR, MI_IO_READ, MI_MAR_L_IN, MI_PC_INC, MI_STEP_INC},
			3: {MI_BRIDGE_ENABLE, MI_PC_OUT_ADDR, MI_IO_READ, MI_MAR_H_IN, MI_PC_INC, MI_STEP_INC},
			4: {MI_MDR_OUT_ALU, MI_MAR_OUT_ADDR, MI_IO_READ, MI_ALU_OUT, MI_ALU_ADC, MI_MDR_IN, MI_STEP_INC},
			5: {MI_MDR_OUT, MI_MAR_OUT_ADDR, MI_IO_WRITE, MI_STEP_CLR},
		},
		instruction.INST_ADD_MEM_REG: {
			0: MI_CMD_FETCH_INST,
			1: MI_CMD_FETCH_OPERAND,
			2: {MI_BRIDGE_ENABLE, MI_PC_OUT_ADDR, MI_IO_READ, MI_MAR_L_IN, MI_PC_INC, MI_STEP_INC},
			3: {MI_BRIDGE_ENABLE, MI_PC_OUT_ADDR, MI_IO_READ, MI_MAR_H_IN, MI_PC_INC, MI_STEP_INC},
			4: {MI_REG_OP_1_OUT_ALU, MI_MAR_OUT_ADDR, MI_IO_READ, MI_ALU_OUT, MI_ALU_ADD, MI_REG_OP_1_IN, MI_STEP_CLR},
		},
		instruction.INST_ADD_SP_IMM: {
			0: MI_CMD_FETCH_INST,
			1: {MI_BRIDGE_ENABLE, MI_PC_OUT_ADDR, MI_IO_READ, MI_MDR_IN, MI_PC_INC, MI_STEP_INC},
			2: {MI_SP_L_OUT, MI_MDR_OUT_ALU, MI_ALU_OUT, MI_ALU_ADD, MI_SP_L_IN, MI_STEP_INC},
			3: {MI_SP_H_OUT, MI_ALU_OUT, MI_ALU_ADC, MI_SP_H_IN, MI_STEP_CLR},
		},
		instruction.INST_SUB_SP_IMM: {
			0: MI_CMD_FETCH_INST,
			1: {MI_BRIDGE_ENABLE, MI_SP_L_OUT, MI_MDR_IN, MI_STEP_INC},
			2: {MI_PC_OUT_ADDR, MI_IO_READ, MI_MDR_OUT_ALU, MI_ALU_OUT, MI_ALU_SUB, MI_SP_L_IN, MI_PC_INC, MI_STEP_INC},
			3: {MI_BRIDGE_ENABLE, MI_SP_H_OUT, MI_MDR_IN, MI_STEP_INC},
			4: {MI_MDR_OUT_ALU, MI_ALU_OUT, MI_ALU_SBB, MI_SP_H_IN, MI_STEP_CLR},
		},
		instruction.INST_MOV_REG_IP: {
			0: MI_CMD_FETCH_INST,
			1: MI_CMD_FETCH_OPERAND,
			2: {MI_REG_OP_1_OUT, MI_IP_OUT_ADDR, MI_IO_WRITE, MI_STEP_CLR},
		},
		instruction.INST_MOV_IP_IMM_16: {
			0: MI_CMD_FETCH_INST,
			1: {MI_BRIDGE_ENABLE, MI_PC_OUT_ADDR, MI_IO_READ, MI_IP_L_IN, MI_PC_INC, MI_STEP_INC},
			2: {MI_BRIDGE_ENABLE, MI_PC_OUT_ADDR, MI_IO_READ, MI_IP_H_IN, MI_PC_INC, MI_STEP_CLR},
		},
		instruction.INST_MOV_IP_REG: {
			0: MI_CMD_FETCH_INST,
			1: MI_CMD_FETCH_OPERAND,
			2: {MI_BRIDGE_ENABLE, MI_IP_OUT_ADDR, MI_IO_READ, MI_REG_OP_1_IN, MI_STEP_CLR},
		},
		instruction.INST_MOV_REG_IP_OFFSET: {
			0: MI_CMD_FETCH_INST,
			1: MI_CMD_FETCH_OPERAND,
			2: {MI_BRIDGE_ENABLE, MI_PC_OUT_ADDR, MI_IO_READ, MI_MDR_IN, MI_PC_INC, MI_STEP_INC},
			3: {MI_IP_L_OUT, MI_MDR_OUT_ALU, MI_ALU_OUT, MI_ALU_ADD, MI_MAR_L_IN, MI_STEP_INC},
			4: {MI_IP_H_OUT, MI_ALU_OUT, MI_ALU_ADC, MI_MAR_H_IN, MI_STEP_INC},
			5: {MI_MAR_OUT_ADDR, MI_REG_OP_1_OUT, MI_IO_WRITE, MI_STEP_CLR},
		},
		instruction.INST_MOV_IP_REG_OFFSET: {
			0: MI_CMD_FETCH_INST,
			1: MI_CMD_FETCH_OPERAND,
			2: {MI_BRIDGE_ENABLE, MI_PC_OUT_ADDR, MI_IO_READ, MI_MDR_IN, MI_PC_INC, MI_STEP_INC},
			3: {MI_IP_L_OUT, MI_MDR_OUT_ALU, MI_ALU_OUT, MI_ALU_ADD, MI_MAR_L_IN, MI_STEP_INC},
			4: {MI_IP_H_OUT, MI_ALU_OUT, MI_ALU_ADC, MI_MAR_H_IN, MI_STEP_INC},
			5: {MI_BRIDGE_ENABLE, MI_MAR_OUT_ADDR, MI_IO_READ, MI_REG_OP_1_IN, MI_STEP_CLR},
		},
		instruction.INST_MOV_REG_IP_OFFSET_REG: {
			0: MI_CMD_FETCH_INST,
			1: MI_CMD_FETCH_OPERAND,
			2: {MI_IP_L_OUT, MI_REG_OP_2_OUT_ALU, MI_ALU_OUT, MI_ALU_ADD, MI_MAR_L_IN, MI_STEP_INC},
			3: {MI_IP_H_OUT, MI_ALU_OUT, MI_ALU_ADC, MI_MAR_H_IN, MI_STEP_INC},
			4: {MI_MAR_OUT_ADDR, MI_REG_OP_1_OUT, MI_IO_WRITE, MI_STEP_CLR},
		},
		instruction.INST_MOV_IP_REG_OFFSET_REG: {
			0: MI_CMD_FETCH_INST,
			1: MI_CMD_FETCH_OPERAND,
			2: {MI_IP_L_OUT, MI_REG_OP_2_OUT_ALU, MI_ALU_OUT, MI_ALU_ADD, MI_MAR_L_IN, MI_STEP_INC},
			3: {MI_IP_H_OUT, MI_ALU_OUT, MI_ALU_ADC, MI_MAR_H_IN, MI_STEP_INC},
			4: {MI_BRIDGE_ENABLE, MI_MAR_OUT_ADDR, MI_IO_READ, MI_REG_OP_1_IN, MI_STEP_CLR},
		},
		instruction.INST_PUSH_IP: {
			0: MI_CMD_FETCH_INST,
			1: {MI_IP_L_OUT, MI_SP_OUT_ADDR, MI_IO_WRITE, MI_SP_DEC, MI_STEP_INC},
			2: {MI_IP_H_OUT, MI_SP_OUT_ADDR, MI_IO_WRITE, MI_SP_DEC, MI_STEP_CLR},
		},
		instruction.INST_POP_IP: {
			0: MI_CMD_FETCH_INST,
			1: {MI_BRIDGE_ENABLE, MI_SP_INC, MI_SP_OUT_ADDR, MI_IO_READ, MI_IP_H_IN, MI_STEP_INC},
			2: {MI_BRIDGE_ENABLE, MI_SP_INC, MI_SP_OUT_ADDR, MI_IO_READ, MI_IP_L_IN, MI_STEP_CLR},
		},
		instruction.INST_ADD_IP_IMM: {
			0: MI_CMD_FETCH_INST,
			1: {MI_BRIDGE_ENABLE, MI_PC_OUT_ADDR, MI_IO_READ, MI_MDR_IN, MI_PC_INC, MI_STEP_INC},
			2: {MI_IP_L_OUT, MI_MDR_OUT_ALU, MI_ALU_OUT, MI_ALU_ADD, MI_IP_L_IN, MI_STEP_INC},
			3: {MI_IP_H_OUT, MI_ALU_OUT, MI_ALU_ADC, MI_IP_H_IN, MI_STEP_CLR},
		},
		instruction.INST_SUB_IP_IMM: {
			0: MI_CMD_FETCH_INST,
			1: {MI_BRIDGE_ENABLE, MI_IP_L_OUT, MI_MDR_IN, MI_STEP_INC},
			2: {MI_PC_OUT_ADDR, MI_IO_READ, MI_MDR_OUT_ALU, MI_ALU_OUT, MI_ALU_SUB, MI_IP_L_IN, MI_PC_INC, MI_STEP_INC},
			3: {MI_BRIDGE_ENABLE, MI_IP_H_OUT, MI_MDR_IN, MI_STEP_INC},
			4: {MI_MDR_OUT_ALU, MI_ALU_OUT, MI_ALU_SBB, MI_IP_H_IN, MI_STEP_CLR},
		},
		instruction.INST_CLC_IMPL: {
			0: MI_CMD_FETCH_INST,
			1: {MI_CLC, MI_STEP_CLR},
		},
		instruction.INST_JMP_IP: {
			0: MI_CMD_FETCH_INST,
			1: {MI_IP_OUT_ADDR, MI_PC_IN_ADDR, MI_STEP_CLR},
		},
		instruction.INST_JZ_IMM:  GENERIC_CONDITIONAL_JMP_AND_NOT_JMP,
		instruction.INST_JNZ_IMM: GENERIC_CONDITIONAL_JMP_AND_NOT_JMP,
		instruction.INST_JS_IMM:  GENERIC_CONDITIONAL_JMP_AND_NOT_JMP,
		instruction.INST_JNS_IMM: GENERIC_CONDITIONAL_JMP_AND_NOT_JMP,
		instruction.INST_JC_IMM:  GENERIC_CONDITIONAL_JMP_AND_NOT_JMP,
	}
)
