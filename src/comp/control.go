package comp

type instruction = uint8
type step = uint8

const (
	MI_FETCH = MI_PC_OUT | MI_RAM_OUT | MI_INST_REG_IN | MI_PC_INC | MI_STEP_INC
)

var (
	control map[instruction]map[step]uint64 = map[instruction]map[step]uint64{
		INST_BREAK: {
			0: MI_FETCH,
			1: MI_BRK,
			2: MI_BRK,
			3: MI_BRK,
			4: MI_BRK,
			5: MI_BRK,
			6: MI_BRK,
			7: MI_BRK,
		},
		INST_MOV_INM_8: {
			0: MI_FETCH,                                                        // fetch next instruction
			1: MI_PC_OUT | MI_RAM_OUT | MI_OP_REG_IN | MI_PC_INC | MI_STEP_INC, // load operation register
			2: MI_PC_OUT | MI_RAM_OUT | MI_OPERAND_REG_1_IN | MI_PC_INC | MI_STEP_CLR,
		},
		INTS_MOV_REG_REG: {
			0: MI_FETCH,                                                        // fetch next instruction
			1: MI_PC_OUT | MI_RAM_OUT | MI_OP_REG_IN | MI_PC_INC | MI_STEP_INC, // load operation register
			2: MI_BRIDGE_ENABLE | MI_BRIDGE_DIR_OUT | MI_OP_REG_1_OUT_Y | MI_OP_REG_2_IN | MI_STEP_CLR,
		},
		INST_ADD_RR: {
			0: MI_FETCH,                                                        // fetch next instruction
			1: MI_PC_OUT | MI_RAM_OUT | MI_OP_REG_IN | MI_PC_INC | MI_STEP_INC, // load operation register
		},
	}
)
