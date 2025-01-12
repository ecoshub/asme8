package comp

const (
	MI_BRK            uint64 = 0
	MI_BRIDGE_ENABLE  uint64 = 1
	MI_BRIDGE_DIR_OUT uint64 = 2
	MI_PC_OUT         uint64 = 3
	MI_SP_OUT         uint64 = 4
	MI_MAR_OUT        uint64 = 5
	MI_MDR_OUT        uint64 = 6
	MI_ALU_OUT        uint64 = 7
	MI_REG_OP_1_OUT_X uint64 = 8
	MI_REG_OP_1_OUT_Y uint64 = 9
	MI_REG_OP_2_OUT_X uint64 = 10
	MI_REG_OP_2_OUT_Y uint64 = 11
	MI_ONES_OUT_X     uint64 = 12
	MI_ONES_OUT_Y     uint64 = 13
	MI_IO_READ        uint64 = 14
	MI_ALU_ENABLE     uint64 = 15
	MI_ALU_ADD        uint64 = 16
	MI_ALU_ADC        uint64 = 17
	MI_ALU_SUB        uint64 = 18
	MI_ALU_CMP        uint64 = 19
	MI_ALU_STR        uint64 = 20
	MI_ALU_STR_OUT    uint64 = 21
	MI_IO_WRITE       uint64 = 22
	MI_INST_REG_IN    uint64 = 23
	MI_OP_REG_IN      uint64 = 24
	MI_REG_OP_1_IN    uint64 = 25
	MI_REG_OP_2_IN    uint64 = 26
	MI_MDR_IN         uint64 = 27
	MI_MAR_L_IN       uint64 = 28
	MI_MAR_H_IN       uint64 = 29
	MI_SP_L_IN        uint64 = 30
	MI_SP_H_IN        uint64 = 31
	MI_PC_IN          uint64 = 32
	MI_PC_INC         uint64 = 33
	MI_SP_INC         uint64 = 34
	MI_SP_DEC         uint64 = 35
	MI_STEP_INC       uint64 = 36
	MI_STEP_CLR       uint64 = 37
	MI_JMP_CTRL       uint64 = 38
)

var (
	MI_NAME_MAP map[uint64]string = map[uint64]string{
		MI_BRK:            "BRK",
		MI_BRIDGE_ENABLE:  "BRIDGE_ENABLE",
		MI_BRIDGE_DIR_OUT: "BRIDGE_DIR_OUT",
		MI_PC_OUT:         "PC_OUT",
		MI_SP_OUT:         "SP_OUT",
		MI_MAR_OUT:        "MAR_OUT",
		MI_MDR_OUT:        "MDR_OUT",
		MI_ALU_OUT:        "ALU_OUT",
		MI_ALU_STR_OUT:    "MI_ALU_STR_OUT",
		MI_REG_OP_1_OUT_X: "REG_OP_1_OUT_X",
		MI_REG_OP_1_OUT_Y: "REG_OP_1_OUT_Y",
		MI_REG_OP_2_OUT_X: "REG_OP_2_OUT_X",
		MI_REG_OP_2_OUT_Y: "REG_OP_2_OUT_Y",
		MI_ONES_OUT_X:     "ONES_OUT_X",
		MI_ONES_OUT_Y:     "ONES_OUT_Y",
		MI_IO_READ:        "IO_READ",
		MI_ALU_ENABLE:     "ALU_ENABLE",
		MI_ALU_ADD:        "ALU_ADD",
		MI_ALU_ADC:        "ALU_ADC",
		MI_ALU_SUB:        "ALU_SUB",
		MI_ALU_CMP:        "ALU_CMP",
		MI_ALU_STR:        "ALU_STORE",
		MI_IO_WRITE:       "IO_WRITE",
		MI_INST_REG_IN:    "INST_REG_IN",
		MI_OP_REG_IN:      "OP_REG_IN",
		MI_REG_OP_1_IN:    "REG_OP_1_IN",
		MI_REG_OP_2_IN:    "REG_OP_2_IN",
		MI_MDR_IN:         "MDR_IN",
		MI_MAR_L_IN:       "MAR_L_IN",
		MI_MAR_H_IN:       "MAR_H_IN",
		MI_SP_L_IN:        "SP_L_IN",
		MI_SP_H_IN:        "SP_H_IN",
		MI_PC_IN:          "PC_IN",
		MI_PC_INC:         "PC_INC",
		MI_SP_INC:         "SP_INC",
		MI_SP_DEC:         "SP_DEC",
		MI_STEP_INC:       "STEP_INC",
		MI_STEP_CLR:       "STEP_CLR",
		MI_JMP_CTRL:       "JMP_CTRL",
	}
)

type miFunc func(c *Comp, command uint64)

var (
	microinstructionFunctions map[uint64]miFunc = map[uint64]miFunc{
		MI_BRK:            mInstBreak,
		MI_BRIDGE_ENABLE:  mInstBridgeEnable,
		MI_BRIDGE_DIR_OUT: mInstBridgeDIR_OUT,
		MI_ALU_ENABLE:     mFlagInst,
		MI_ALU_STR:        mInstAluStore,
		MI_ALU_OUT:        mInstAluStoreOut,
		MI_ALU_STR_OUT:    mInstAluStoreOut,
		MI_ALU_ADD:        mInstAluAdd,
		MI_ALU_ADC:        mInstAluAdc,
		MI_ALU_SUB:        mInstAluSub,
		MI_ALU_CMP:        mInstAluCmp,
		MI_JMP_CTRL:       mInstStatusControl,
		MI_PC_OUT:         mInstProgramCounterOut,
		MI_SP_OUT:         mInstStackPointerOut,
		MI_IO_READ:        mInstSetRead,
		MI_IO_WRITE:       mInstSetWrite,
		MI_INST_REG_IN:    mInstInstructionRegisterIn,
		MI_OP_REG_IN:      mInstOperandRegisterIn,
		MI_PC_INC:         mInstProgramCounterInc,
		MI_SP_INC:         mInstStackPointerInc,
		MI_SP_DEC:         mInstStackPointerDec,
		MI_STEP_INC:       mInstStepInc,
		MI_STEP_CLR:       mInstStepClr,
		MI_REG_OP_1_IN:    mInstReg1In,
		MI_REG_OP_2_IN:    mInstReg2In,
		MI_REG_OP_1_OUT_X: mInstReg1OutX,
		MI_REG_OP_1_OUT_Y: mInstReg1OutY,
		MI_REG_OP_2_OUT_X: mInstReg2OutX,
		MI_REG_OP_2_OUT_Y: mInstReg2OutY,
		MI_MDR_IN:         mInstMemoryDataRegisterIn,
		MI_MDR_OUT:        mInstMemoryDataRegisterOut,
		MI_ONES_OUT_X:     mInstOnesOutX,
		MI_ONES_OUT_Y:     mInstOnesOutY,
		MI_MAR_OUT:        mInstMarOut,
		MI_MAR_L_IN:       mInstMarLowIn,
		MI_MAR_H_IN:       mInstMarHighIn,
		MI_SP_L_IN:        mInstStackPointerLowIn,
		MI_SP_H_IN:        mInstStackPointerHighIn,
		MI_PC_IN:          mInstProgramCounterIn,
	}
)
