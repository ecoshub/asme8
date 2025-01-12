package instruction

type Type uint8

const (
	INST_BREAK              Type = 0x00
	INST_JMP_INM            Type = 0x10
	INST_JZ_INM             Type = 0x15
	INST_ADD_RR             Type = 0x30
	INST_ADD_INM            Type = 0x31
	INST_SUB_RR             Type = 0x35
	INST_CMP_RR             Type = 0x3a
	INST_CMP_INM            Type = 0x3b
	INTS_MOV_REG_REG        Type = 0x40
	INST_MOV_INM_8          Type = 0x41
	INST_MOV_MEM_REG        Type = 0x42
	INST_MOV_REG_MEM        Type = 0x43
	INST_MOV_REG_MEM_OFFSET Type = 0x44
	INST_NOP                Type = 0xff
	INST_PUSH_REG           Type = 0xf0
	INST_PUSH_INM           Type = 0xf1
	INST_POP_REG            Type = 0xf2
	INST_JSR                Type = 0xf3
	INST_RTS                Type = 0xf4
)
