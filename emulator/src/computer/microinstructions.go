package computer

const (
	MI_HLT uint64 = iota
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
	MI_CL_C
	MI_CL_IE
	MI_STI
	MI_CL_IRQ
	MI_STATUS_OUT
	MI_STATUS_IN
	MI_VEC_0_OUT
	MI_VEC_1_OUT
	MI_VEC_2_OUT
	MI_VEC_3_OUT
	TEST_MI_IRQ_HIGH
)

type miFunc func(c *Computer, command uint64)

var (
	microinstructionFunctions map[uint64]miFunc = map[uint64]miFunc{
		MI_HLT:                   mInstHalt,
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
		MI_CL_C:                  mInstClearCarryFlag,
		MI_CL_IE:                 mInstClearInterruptEnableFlag,
		MI_STI:                   mInstSetInterruptEnableFlag,
		MI_CL_IRQ: func(c *Computer, command uint64) {
			c.irqLine = false
			c.irqAckFlipFlop = false
		},
		MI_STATUS_OUT: mInstStatusRegisterOut,
		MI_STATUS_IN:  mInstStatusRegisterIn,
		MI_VEC_0_OUT:  mInstVec0Out,
		MI_VEC_1_OUT:  mInstVec1Out,
		MI_VEC_2_OUT:  mInstVec2Out,
		MI_VEC_3_OUT:  mInstVec3Out,
		TEST_MI_IRQ_HIGH: func(c *Computer, command uint64) {
			c.irqLine = true
		},
	}
)
