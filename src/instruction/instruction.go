package instruction

type Type uint8

const (
	INST_BREAK       Type = 0x00
	INST_JMP_INM     Type = 0x10
	INST_JZ_INM      Type = 0x15
	INST_ADD_RR      Type = 0x30
	INST_SUB_RR      Type = 0x35
	INST_CMP_RR      Type = 0x3a
	INST_CMP_INM     Type = 0x3b
	INTS_MOV_REG_REG Type = 0x40
	INST_MOV_INM_8   Type = 0x41
	INST_MOV_MEM_REG Type = 0x42
	INST_MOV_REG_MEM Type = 0x43
	INST_NOP         Type = 0xff
)
