package comp

type instruction = uint8
type step = uint8

const (
	INST_BREAK       byte = 0x00
	INST_JMP         byte = 0x10
	INST_ADD_RR      byte = 0x30
	INTS_MOV_REG_REG byte = 0x40
	INST_MOV_INM_8   byte = 0x41
)

const (
	M_INST_FETCH = M_INST_PC_OUT | M_INST_RAM_OUT | M_INST_INST_REG_IN | M_INST_PC_INC | M_INST_STEP_INC
)

var (
	control map[instruction]map[step]uint64 = map[instruction]map[step]uint64{
		INST_BREAK: {
			0: M_INST_BRK,
			1: M_INST_BRK,
			2: M_INST_BRK,
			3: M_INST_BRK,
			4: M_INST_BRK,
			5: M_INST_BRK,
			6: M_INST_BRK,
			7: M_INST_BRK,
		},
		INST_MOV_INM_8: {
			0: M_INST_FETCH,
			1: M_INST_PC_OUT | M_INST_RAM_OUT | M_INST_OP_REG_IN | M_INST_PC_INC | M_INST_STEP_INC,
			2: M_INST_PC_OUT | M_INST_RAM_OUT | M_INST_REG_1_IN | M_INST_PC_INC | M_INST_STEP_CLR,
		},
		INTS_MOV_REG_REG: {
			0: M_INST_FETCH,
			1: M_INST_PC_OUT | M_INST_RAM_OUT | M_INST_OP_REG_IN | M_INST_PC_INC | M_INST_STEP_INC,
			2: M_INST_REG_1_OUT | M_INST_REG_2_IN | M_INST_STEP_CLR,
		},
	}
)
