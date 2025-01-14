package instruction

type Type = uint8

const (
	// Special
	INST_BREAK Type = 0x00
	INST_NOP   Type = 0xff

	// Branching
	INST_JMP_INM     Type = 0x10
	_                Type = 0x11
	_                Type = 0x12
	_                Type = 0x13
	_                Type = 0x14
	INST_JZ_INM      Type = 0x15
	_                Type = 0x16
	_                Type = 0x17
	_                Type = 0x18
	_                Type = 0x19
	INST_CMP_RR      Type = 0x1a
	INST_CMP_REG_INM Type = 0x1b
	_                Type = 0x1c
	_                Type = 0x1d

	// Arithmetic
	INST_ADD_RR      Type = 0x30
	INST_ADD_REG_INM Type = 0x31
	INST_ADC_RR      Type = 0x32
	INST_ADC_REG_INM Type = 0x33
	INST_SUB_RR      Type = 0x34
	INST_SUB_REG_INM Type = 0x35
	INST_SBB_RR      Type = 0x36
	INST_SBB_REG_INM Type = 0x37
	INST_INC         Type = 0x38
	INST_DEC         Type = 0x39
	_                Type = 0x3a
	_                Type = 0x3b
	_                Type = 0x3c
	_                Type = 0x3d
	_                Type = 0x3e
	_                Type = 0x3f

	// Move
	INTS_MOV_REG_REG        Type = 0x40
	INST_MOV_INM_8          Type = 0x41
	INST_MOV_MEM_REG        Type = 0x42
	INST_MOV_REG_MEM        Type = 0x43
	INST_MOV_MEM_REG_OFFSET Type = 0x44
	INST_MOV_REG_MEM_OFFSET Type = 0x45
	_                       Type = 0x46
	_                       Type = 0x47
	_                       Type = 0x48
	_                       Type = 0x49
	_                       Type = 0x4a
	_                       Type = 0x4b
	_                       Type = 0x4c
	_                       Type = 0x4d
	_                       Type = 0x4e
	_                       Type = 0x4f

	// Stack base
	INST_PUSH_REG Type = 0xf0
	INST_PUSH_INM Type = 0xf1
	INST_POP_REG  Type = 0xf2
	INST_JSR      Type = 0xf3 // NOT IMPLEMENTED
	INST_RTS      Type = 0xf4 // NOT IMPLEMENTED

)
