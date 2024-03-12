package comp

const (
	instJump         byte = 0x10
	instAddRegReg    byte = 0x30
	INTS_MOV_REG_REG byte = 0x40
	INST_MOV_INM_8   byte = 0x41
	INS_BREAK        byte = 0x00
)

const (
	CMD_FETCH = CMD_PC_OUT | CMD_RAM_OUT | CMD_INST_REG_IN | CMD_PC_INC | CMD_STEP_INC
)

var (
	control map[uint8]map[uint8]uint64 = map[uint8]map[uint8]uint64{
		0: {
			INS_BREAK:        CMD_BRK,
			INST_MOV_INM_8:   CMD_FETCH,
			INTS_MOV_REG_REG: CMD_FETCH,
		},
		1: {
			INS_BREAK:        CMD_BRK,
			INST_MOV_INM_8:   CMD_PC_OUT | CMD_RAM_OUT | CMD_OP_REG_IN | CMD_PC_INC | CMD_STEP_INC,
			INTS_MOV_REG_REG: CMD_PC_OUT | CMD_RAM_OUT | CMD_OP_REG_IN | CMD_PC_INC | CMD_STEP_INC,
		},
		2: {
			INS_BREAK:        CMD_BRK,
			INST_MOV_INM_8:   CMD_PC_OUT | CMD_RAM_OUT | CMD_REG_1_IN | CMD_PC_INC | CMD_STEP_CLR,
			INTS_MOV_REG_REG: CMD_REG_1_OUT | CMD_REG_2_IN | CMD_STEP_CLR,
		},
	}
)
