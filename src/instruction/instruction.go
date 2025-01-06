package instruction

type Type = uint8

const (
	INST_BREAK       byte = 0x00
	INST_JMP_INM     byte = 0x10
	INST_JZ_INM      byte = 0x15
	INST_ADD_RR      byte = 0x30
	INST_SUB_RR      byte = 0x35
	INST_CMP_RR      byte = 0x3a
	INST_CMP_INM     byte = 0x3b
	INTS_MOV_REG_REG byte = 0x40
	INST_MOV_INM_8   byte = 0x41
	INST_NOP         byte = 0xff
)
