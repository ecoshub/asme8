package comp

const (
	// NOTE ORDER IS IMPORTANT!
	// *_OUT microinstructions must lesser mask than *_IN microinstructions
	// because those micro instructions run by sequence and must output instructions run before
	// input ones
	MI_BRK = 0

	// OUT SECTION
	MI_PC_OUT uint64 = 1 << iota
	MI_SP_OUT
	MI_SPL_OUT_X
	MI_SPH_OUT_X
	MI_MAR_OUT
	MI_MDR_OUT
	MI_IO_READ
	MI_ALU_OUT
	MI_REG_OPERAND_1_OUT_X
	MI_REG_OPERAND_1_OUT_DATA
	MI_REG_OPERAND_2_OUT_X
	MI_REG_OPERAND_2_OUT_DATA
	MI_ONES_OUT_X
	MI_ONES_OUT_Y

	// FLAG SECTION
	MI_ALU_ENABLE
	MI_ALU_ADD
	MI_ALU_ADC
	MI_ALU_SUB
	MI_ALU_CMP

	// MUST BE LAST
	MI_ALU_STORE

	// IN SECTION
	MI_IO_WRITE
	MI_INST_REG_IN
	MI_OPERAND_REG_IN
	MI_REG_OPERAND_1_IN
	MI_REG_OPERAND_2_IN
	MI_MDR_IN
	MI_MAR_L_IN
	MI_MAR_H_IN
	MI_SP_L_IN
	MI_SP_H_IN

	// OTHER SECTION
	MI_PC_IN
	MI_PC_INC
	MI_SP_INC
	MI_STEP_INC
	MI_STEP_CLR

	// STATUS SECTION
	MI_JMP_CTRL
)

type runnerFunc func(c *Comp, command uint64, _ uint64)

var (
	microInstructions map[uint64]runnerFunc = map[uint64]runnerFunc{
		MI_BRK:                    mInstBreak,
		MI_ALU_ENABLE:             mFlagInst,
		MI_ALU_STORE:              mInstAluStore,
		MI_ALU_OUT:                mInstAluOut,
		MI_ALU_ADD:                mInstAluAdd,
		MI_ALU_ADC:                mInstAluAdc,
		MI_ALU_SUB:                mInstAluSub,
		MI_ALU_CMP:                mInstAluCmp,
		MI_JMP_CTRL:               mInstStatusControl,
		MI_PC_OUT:                 mInstProgramCounterOut,
		MI_SP_OUT:                 mInstStackPointerOut,
		MI_SPL_OUT_X:              mInstStackPointerLowOutX,
		MI_SPH_OUT_X:              mInstStackPointerHighOutX,
		MI_IO_READ:                mInstSetRead,
		MI_IO_WRITE:               mInstSetWrite,
		MI_INST_REG_IN:            mInstInstructionRegisterIn,
		MI_OPERAND_REG_IN:         mInstOperandRegisterIn,
		MI_PC_INC:                 mInstProgramCounterInc,
		MI_SP_INC:                 mInstStackPointerInc,
		MI_STEP_INC:               mInstStepInc,
		MI_STEP_CLR:               mInstStepClr,
		MI_REG_OPERAND_1_IN:       mInstReg1In,
		MI_REG_OPERAND_2_IN:       mInstReg2In,
		MI_REG_OPERAND_1_OUT_X:    mInstReg1OutX,
		MI_REG_OPERAND_1_OUT_DATA: mInstReg1OutData,
		MI_REG_OPERAND_2_OUT_X:    mInstReg2OutX,
		MI_REG_OPERAND_2_OUT_DATA: mInstReg2OutData,
		MI_MDR_IN:                 mInstMemoryDataRegisterIn,
		MI_MDR_OUT:                mInstMemoryDataRegisterOut,
		MI_ONES_OUT_X:             mInstOnesOutX,
		MI_ONES_OUT_Y:             mInstOnesOutData,
		MI_MAR_OUT:                mInstMarOut,
		MI_MAR_L_IN:               mInstMarLowIn,
		MI_MAR_H_IN:               mInstMarHighIn,
		MI_PC_IN:                  mInstProgramCounterIn,
	}
)
