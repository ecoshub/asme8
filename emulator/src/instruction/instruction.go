package instruction

type Type = uint8

const (
	// Special
	INST_BRK Type = 0x00
	INST_NOP Type = 0xff

	// Branching
	INST_JMP_IMM     Type = 0x10
	_                Type = 0x11
	_                Type = 0x12
	INST_JS_IMM      Type = 0x13
	INST_JNS_IMM     Type = 0x14
	INST_JZ_IMM      Type = 0x15
	INST_JNZ_IMM     Type = 0x16
	INST_JC_IMM      Type = 0x17
	_                Type = 0x18
	_                Type = 0x19
	INST_CMP_RR      Type = 0x1a
	INST_CMP_REG_IMM Type = 0x1b
	_                Type = 0x1c
	_                Type = 0x1d
	_                Type = 0x1e
	_                Type = 0x1f

	_ Type = 0x20
	_ Type = 0x21
	_ Type = 0x22

	// Arithmetic operations
	INST_ADD_RR      Type = 0x30
	INST_ADD_REG_IMM Type = 0x31
	INST_ADC_RR      Type = 0x32
	INST_ADC_REG_IMM Type = 0x33
	INST_SUB_RR      Type = 0x34
	INST_SUB_REG_IMM Type = 0x35
	INST_SBB_RR      Type = 0x36
	INST_SBB_REG_IMM Type = 0x37
	INST_INC         Type = 0x38
	INST_DEC         Type = 0x39
	INST_ADD_REG_MEM Type = 0x3a
	INST_ADD_MEM_REG Type = 0x3b
	INST_ADD_SP      Type = 0x3c
	INST_SUB_SP      Type = 0x3d
	_                Type = 0x3e
	_                Type = 0x3f

	// Move
	INST_MOV_REG_REG           Type = 0x40
	INST_MOV_IMM               Type = 0x41
	INST_MOV_MEM_REG           Type = 0x42
	INST_MOV_REG_MEM           Type = 0x43
	INST_MOV_MEM_REG_OFFSET    Type = 0x44
	INST_MOV_REG_MEM_OFFSET    Type = 0x45
	INST_MOV_REG_SP            Type = 0x46
	INST_MOV_SP_REG            Type = 0x47
	INST_MOV_REG_SP_OFFSET     Type = 0x48
	INST_MOV_SP_REG_OFFSET     Type = 0x49
	INST_MOV_REG_SP_OFFSET_REG Type = 0x4a
	INST_MOV_SP_REG_OFFSET_REG Type = 0x4b
	_                          Type = 0x4c
	_                          Type = 0x4d
	_                          Type = 0x4e
	_                          Type = 0x4f

	// Logic operations
	INST_XOR_RR      Type = 0x50
	INST_XOR_REG_IMM Type = 0x51 // NOTE: NOT IMPLEMENTED YET
	INST_AND_RR      Type = 0x52 // NOTE: NOT IMPLEMENTED YET
	INST_AND_REG_IMM Type = 0x53 // NOTE: NOT IMPLEMENTED YET
	INST_OR_RR       Type = 0x54 // NOTE: NOT IMPLEMENTED YET
	INST_OR_REG_IMM  Type = 0x55 // NOTE: NOT IMPLEMENTED YET
	INST_NOT_REG     Type = 0x56 // NOTE: NOT IMPLEMENTED YET

	// Shift operations
	INST_SHL_REG Type = 0x57 // NOTE: NOT IMPLEMENTED YET
	INST_SHR_REG Type = 0x58 // NOTE: NOT IMPLEMENTED YET
	INST_ROL_REG Type = 0x59 // NOTE: NOT IMPLEMENTED YET
	INST_ROR_REG Type = 0x5a // NOTE: NOT IMPLEMENTED YET
	INST_ASL_REG Type = 0x5b // NOTE: NOT IMPLEMENTED YET
	INST_ASR_REG Type = 0x5c // NOTE: NOT IMPLEMENTED YET

	// Flag (status) operations
	INST_CLC Type = 0xe0

	// Stack base
	INST_PUSH_SP  Type = 0xee
	INST_POP_SP   Type = 0xef
	INST_PUSH_REG Type = 0xf0
	INST_PUSH_IMM Type = 0xf1
	INST_POP_REG  Type = 0xf2
	INST_JSR      Type = 0xf3 // NOT IMPLEMENTED
	INST_RTS      Type = 0xf4 // NOT IMPLEMENTED
)

var (
	INST_MNEMONICS = map[uint8]string{
		0x00: "BRK",
		0xff: "NOP",
		0x10: "JMP_IMM",
		0x13: "JS_IMM",
		0x14: "JNS_IMM",
		0x15: "JZ_IMM",
		0x16: "JNZ_IMM",
		0x1a: "CMP_RR",
		0x1b: "CMP_RI",
		0x30: "ADD_RR",
		0x31: "ADD_RI",
		0x32: "ADC_RR",
		0x33: "ADC_RI",
		0x34: "SUB_RR",
		0x35: "SUB_RI",
		0x36: "SBB_RR",
		0x37: "SBB_RI",
		0x38: "INC",
		0x39: "DEC",
		0x40: "MOV_RR",
		0x41: "MOV_IMM",
		0x42: "MOV_MR",
		0x43: "MOV_RM",
		0x44: "MOV_MR_OFF",
		0x45: "MOV_RM_OFF",
		0xee: "PUSH_SP",
		0xef: "POP_SP",
		0xf0: "PUSH_REG",
		0xf1: "PUSH_IMM",
		0xf2: "POP_REG",
		0xf3: "JSR",
		0xf4: "RTS",
	}
)
