package comp

const (
	MI_BRK uint64 = iota
	MI_BRIDGE_ENABLE
	MI_PC_IN_ADDR
	MI_PC_OUT_ADDR
	MI_PC_L_OUT
	MI_PC_H_OUT
	MI_PC_INC
	MI_SP_OUT_ADDR
	MI_SP_L_OUT
	MI_SP_H_OUT
	MI_SP_L_IN
	MI_SP_H_IN
	MI_SP_INC
	MI_SP_DEC
	MI_MAR_OUT_ADDR
	MI_MAR_L_IN
	MI_MAR_H_IN
	MI_MDR_IN
	MI_MDR_OUT
	MI_MDR_OUT_ALU
	MI_ALU_OUT
	MI_ALU_ADD
	MI_ALU_ADC
	MI_ALU_SUB
	MI_ALU_SBB
	MI_ALU_CMP
	MI_ALU_XOR
	MI_ALU_AND
	MI_ALU_OR
	MI_ALU_NOT
	MI_ALU_SHL
	MI_ALU_SHR
	MI_ALU_ROL
	MI_ALU_ROR
	MI_OP_1_LOW_IN
	MI_OP_1_HIGH_IN
	MI_OP_1_LOW_OUT
	MI_OP_1_HIGH_OUT
	MI_OP_1_LOW_OUT_SPECIAL
	MI_OP_1_HIGH_OUT_SPECIAL
	MI_OP_2_LOW_IN
	MI_OP_2_HIGH_IN
	MI_OP_2_LOW_OUT
	MI_OP_2_HIGH_OUT
	MI_OP_2_LOW_OUT_SPECIAL
	MI_OP_2_HIGH_OUT_SPECIAL
	MI_ONES_OUT_ALU
	MI_ONES_OUT
	MI_IO_READ
	MI_IO_WRITE
	MI_INSTRUCTION_REG_IN
	MI_OPERAND_REG_IN
	MI_STEP_INC
	MI_STEP_CLR
	MI_JMP_CTRL
	MI_CLC
)

var (
	MI_NAME_MAP map[uint64]string = map[uint64]string{
		MI_BRK:                "BRK",
		MI_BRIDGE_ENABLE:      "BRIDGE_ENABLE",
		MI_PC_IN_ADDR:         "PC_IN_ADDR",
		MI_PC_OUT_ADDR:        "PC_OUT_ADDR",
		MI_PC_L_OUT:           "PC_L_OUT",
		MI_PC_H_OUT:           "PC_H_OUT",
		MI_PC_INC:             "PC_INC",
		MI_SP_OUT_ADDR:        "SP_OUT_ADDR",
		MI_SP_L_OUT:           "SP_L_OUT",
		MI_SP_H_OUT:           "SP_H_OUT",
		MI_SP_L_IN:            "SP_L_IN",
		MI_SP_H_IN:            "SP_H_IN",
		MI_MAR_OUT_ADDR:       "MAR_OUT_ADDR",
		MI_MAR_L_IN:           "MAR_L_IN",
		MI_MAR_H_IN:           "MAR_H_IN",
		MI_MDR_IN:             "MDR_IN",
		MI_MDR_OUT:            "MDR_OUT",
		MI_ALU_OUT:            "ALU_OUT",
		MI_ALU_ADD:            "ALU_ADD",
		MI_ALU_ADC:            "ALU_ADC",
		MI_ALU_SUB:            "ALU_SUB",
		MI_ALU_SBB:            "ALU_SBB",
		MI_ALU_AND:            "ALU_AND",
		MI_ALU_OR:             "ALU_OR",
		MI_ALU_NOT:            "ALU_NOT",
		MI_ALU_SHL:            "ALU_SHL",
		MI_ALU_SHR:            "ALU_SHR",
		MI_ALU_ROL:            "ALU_ROL",
		MI_ALU_ROR:            "ALU_ROR",
		MI_ALU_CMP:            "ALU_CMP",
		MI_ONES_OUT_ALU:       "ONES_OUT_ALU",
		MI_ONES_OUT:           "ONES_OUT",
		MI_IO_READ:            "IO_READ",
		MI_IO_WRITE:           "IO_WRITE",
		MI_INSTRUCTION_REG_IN: "INST_REG_IN",
		MI_OPERAND_REG_IN:     "OP_REG_IN",
		MI_STEP_INC:           "STEP_INC",
		MI_STEP_CLR:           "STEP_CLR",
		MI_JMP_CTRL:           "JMP_CTRL",
		MI_CLC:                "MI_CLC",
	}
)

type miFunc func(c *Comp, command uint64)

var (
	microinstructionFunctions map[uint64]miFunc = map[uint64]miFunc{
		MI_BRK:                   mInstBreak,
		MI_BRIDGE_ENABLE:         mInstBridgeEnable,
		MI_PC_IN_ADDR:            mInstProgramCounterInAddr,
		MI_PC_OUT_ADDR:           mInstProgramCounterOutAddr,
		MI_PC_L_OUT:              mInstProgramCounterLowOut,
		MI_PC_H_OUT:              mInstProgramCounterHighOut,
		MI_PC_INC:                mInstProgramCounterInc,
		MI_SP_OUT_ADDR:           mInstStackPointerOutAddr,
		MI_SP_L_OUT:              mInstStackPointerLowOut,
		MI_SP_H_OUT:              mInstStackPointerHighOut,
		MI_SP_L_IN:               mInstStackPointerLowIn,
		MI_SP_H_IN:               mInstStackPointerHighIn,
		MI_SP_INC:                mInstStackPointerInc,
		MI_SP_DEC:                mInstStackPointerDec,
		MI_MAR_OUT_ADDR:          mInstMemoryAddressRegisterOutAddr,
		MI_MAR_L_IN:              mInstMemoryAddressRegisterLowIn,
		MI_MAR_H_IN:              mInstMemoryAddressRegisterHighIn,
		MI_MDR_IN:                mInstMemoryDataRegisterIn,
		MI_MDR_OUT:               mInstMemoryDataRegisterOut,
		MI_MDR_OUT_ALU:           mInstMemoryDataRegisterOutAlu,
		MI_ALU_OUT:               mInstAluOut,
		MI_ALU_ADD:               mInstAluAdd,
		MI_ALU_ADC:               mInstAluAdc,
		MI_ALU_SUB:               mInstAluSub,
		MI_ALU_SBB:               mInstAluSbb,
		MI_ALU_CMP:               mInstAluCmp,
		MI_ALU_XOR:               mInstAluXor,
		MI_ALU_AND:               mInstAluAnd,
		MI_ALU_OR:                mInstAluOr,
		MI_ALU_NOT:               mInstAluNot,
		MI_ALU_SHL:               mInstAluShl,
		MI_ALU_SHR:               mInstAluShr,
		MI_ALU_ROL:               mInstAluRol,
		MI_ALU_ROR:               mInstAluRor,
		MI_OP_1_LOW_IN:           OperandOperation(OP_1 | OP_LOW | OP_INPUT),
		MI_OP_1_HIGH_IN:          OperandOperation(OP_1 | OP_HIGH | OP_INPUT),
		MI_OP_1_LOW_OUT:          OperandOperation(OP_1 | OP_LOW | OP_OUTPUT),
		MI_OP_1_HIGH_OUT:         OperandOperation(OP_1 | OP_HIGH | OP_OUTPUT),
		MI_OP_1_LOW_OUT_SPECIAL:  OperandOperation(OP_1 | OP_LOW | OP_OUTPUT_SPECIAL),
		MI_OP_1_HIGH_OUT_SPECIAL: OperandOperation(OP_1 | OP_HIGH | OP_OUTPUT_SPECIAL),
		MI_OP_2_LOW_IN:           OperandOperation(OP_2 | OP_LOW | OP_INPUT),
		MI_OP_2_HIGH_IN:          OperandOperation(OP_2 | OP_HIGH | OP_INPUT),
		MI_OP_2_LOW_OUT:          OperandOperation(OP_2 | OP_LOW | OP_OUTPUT),
		MI_OP_2_HIGH_OUT:         OperandOperation(OP_2 | OP_HIGH | OP_OUTPUT),
		MI_OP_2_LOW_OUT_SPECIAL:  OperandOperation(OP_2 | OP_LOW | OP_OUTPUT_SPECIAL),
		MI_OP_2_HIGH_OUT_SPECIAL: OperandOperation(OP_2 | OP_HIGH | OP_OUTPUT_SPECIAL),
		MI_ONES_OUT_ALU:          mInstOnesOutAlu,
		MI_ONES_OUT:              mInstOnesOut,
		MI_IO_READ:               mInstSetRead,
		MI_IO_WRITE:              mInstSetWrite,
		MI_INSTRUCTION_REG_IN:    mInstInstructionRegisterIn,
		MI_OPERAND_REG_IN:        mInstOperandRegisterIn,
		MI_STEP_INC:              mInstStepInc,
		MI_STEP_CLR:              mInstStepClr,
		MI_JMP_CTRL:              mInstStatusControl,
		MI_CLC:                   mInstClearCarryFlag,
	}
)
