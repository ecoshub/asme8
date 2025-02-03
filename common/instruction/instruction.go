package instruction

type Type = uint8

const (
	INST_BRK_IMPL              Type = 0x00
	INST_NOP_IMPL              Type = 0x01
	INST_JMP_IMM               Type = 0x02
	INST_JS_IMM                Type = 0x03
	INST_JNS_IMM               Type = 0x04
	INST_JZ_IMM                Type = 0x05
	INST_JNZ_IMM               Type = 0x06
	INST_JC_IMM                Type = 0x07
	INST_CMP_REG_REG           Type = 0x08
	INST_CMP_REG_IMM           Type = 0x09
	INST_MOV_REG_IP            Type = 0x0a
	INST_MOV_IP_REG            Type = 0x0b
	INST_MOV_REG_IP_OFFSET     Type = 0x0c
	INST_MOV_IP_REG_OFFSET     Type = 0x0d
	INST_MOV_REG_IP_OFFSET_REG Type = 0x0e
	INST_MOV_IP_REG_OFFSET_REG Type = 0x0f
	INST_MOV_IP_IMM_16         Type = 0x10
	INST_MOV_MEM_IMM           Type = 0x11
	INST_ADD_REG_REG           Type = 0x12
	INST_ADD_REG_IMM           Type = 0x13
	INST_ADC_REG_REG           Type = 0x14
	INST_ADC_REG_MEM           Type = 0x15
	INST_ADC_MEM_REG           Type = 0x16
	INST_ADC_REG_IMM           Type = 0x17
	INST_SUB_REG_REG           Type = 0x18
	INST_SUB_REG_IMM           Type = 0x19
	INST_SBB_REG_REG           Type = 0x1a
	INST_SBB_REG_IMM           Type = 0x1b
	INST_INC_REG               Type = 0x1c
	INST_DEC_REG               Type = 0x1d
	INST_ADD_REG_MEM           Type = 0x1e
	INST_ADD_MEM_IMM           Type = 0x1f
	INST_ADC_MEM_IMM           Type = 0x20
	INST_ADD_MEM_REG           Type = 0x21
	INST_ADD_SP_IMM            Type = 0x22
	INST_SUB_SP_IMM            Type = 0x23
	INST_ADD_IP_IMM            Type = 0x24
	INST_SUB_IP_IMM            Type = 0x25
	INST_MOV_REG_REG           Type = 0x26
	INST_MOV_REG_IMM           Type = 0x27
	INST_MOV_MEM_REG           Type = 0x28
	INST_MOV_REG_MEM           Type = 0x29
	INST_MOV_MEM_REG_OFFSET    Type = 0x2a
	INST_MOV_REG_MEM_OFFSET    Type = 0x2b
	INST_MOV_REG_SP            Type = 0x2c
	INST_MOV_SP_REG            Type = 0x2d
	INST_MOV_REG_SP_OFFSET     Type = 0x2e
	INST_MOV_SP_REG_OFFSET     Type = 0x2f
	INST_MOV_REG_SP_OFFSET_REG Type = 0x30
	INST_MOV_SP_REG_OFFSET_REG Type = 0x31
	INST_XOR_REG_REG           Type = 0x32
	INST_XOR_REG_IMM           Type = 0x33
	INST_AND_REG_REG           Type = 0x34
	INST_AND_REG_IMM           Type = 0x35
	INST_OR_REG_REG            Type = 0x36
	INST_OR_REG_IMM            Type = 0x37
	INST_NOT_REG               Type = 0x38
	INST_SHL_REG               Type = 0x39
	INST_SHR_REG               Type = 0x3a
	INST_ROL_REG               Type = 0x3b
	INST_ROR_REG               Type = 0x3c
	INST_CLC_IMPL              Type = 0x3d
	INST_PUSH_IP               Type = 0x3e
	INST_POP_IP                Type = 0x3f
	INST_PUSH_SP               Type = 0x40
	INST_POP_SP                Type = 0x41
	INST_PUSH_REG              Type = 0x42
	INST_PUSH_IMM              Type = 0x43
	INST_POP_REG               Type = 0x44
	INST_CALL_IMPL             Type = 0x45
	INST_RET_IMPL              Type = 0x46
)

const (
	INST_MNEMONICS_BRK  string = "brk"
	INST_MNEMONICS_MOV  string = "mov"
	INST_MNEMONICS_ADD  string = "add"
	INST_MNEMONICS_ADC  string = "adc"
	INST_MNEMONICS_SUB  string = "sub"
	INST_MNEMONICS_SBB  string = "sbb"
	INST_MNEMONICS_XOR  string = "xor"
	INST_MNEMONICS_AND  string = "and"
	INST_MNEMONICS_OR   string = "or"
	INST_MNEMONICS_NOT  string = "not"
	INST_MNEMONICS_SHL  string = "shl"
	INST_MNEMONICS_SHR  string = "shr"
	INST_MNEMONICS_ROL  string = "rol"
	INST_MNEMONICS_ROR  string = "ror"
	INST_MNEMONICS_CMP  string = "cmp"
	INST_MNEMONICS_JMP  string = "jmp"
	INST_MNEMONICS_JZ   string = "jz"
	INST_MNEMONICS_JNZ  string = "jnz"
	INST_MNEMONICS_JC   string = "jc"
	INST_MNEMONICS_JS   string = "js"
	INST_MNEMONICS_JNS  string = "jns"
	INST_MNEMONICS_CLC  string = "clc"
	INST_MNEMONICS_PUSH string = "push"
	INST_MNEMONICS_POP  string = "pop"
	INST_MNEMONICS_INC  string = "inc"
	INST_MNEMONICS_DEC  string = "dec"
	INST_MNEMONICS_CALL string = "call"
	INST_MNEMONICS_RET  string = "ret"
	INST_MNEMONICS_NOP  string = "nop"
)

var (
	INST_HUMAN_READABLE = map[uint8]string{
		0x00: "BRK_IMPL",
		0x01: "NOP_IMPL",
		0x02: "JMP_IMM",
		0x03: "JS_IMM",
		0x04: "JNS_IMM",
		0x05: "JZ_IMM",
		0x06: "JNZ_IMM",
		0x07: "JC_IMM",
		0x08: "CMP_REG_REG",
		0x09: "CMP_REG_IMM",
		0x0a: "MOV_REG_IP",
		0x0b: "MOV_IP_REG",
		0x0c: "MOV_REG_IP_OFFSET",
		0x0d: "MOV_IP_REG_OFFSET",
		0x0e: "MOV_REG_IP_OFFSET_REG",
		0x0f: "MOV_IP_REG_OFFSET_REG",
		0x10: "MOV_IP_IMM_16",
		0x11: "ADD_REG_REG",
		0x12: "ADD_REG_IMM",
		0x13: "ADC_REG_REG",
		0x14: "ADC_REG_MEM",
		0x15: "ADC_MEM_REG",
		0x16: "ADC_REG_IMM",
		0x17: "SUB_REG_REG",
		0x18: "SUB_REG_IMM",
		0x19: "SBB_REG_REG",
		0x1a: "SBB_REG_IMM",
		0x1b: "INC_REG",
		0x1c: "DEC_REG",
		0x1d: "ADD_REG_MEM",
		0x1e: "ADD_MEM_IMM",
		0x1f: "ADC_MEM_IMM",
		0x20: "ADD_MEM_REG",
		0x21: "ADD_SP_IMM",
		0x22: "SUB_SP_IMM",
		0x23: "ADD_IP_IMM",
		0x24: "SUB_IP_IMM",
		0x25: "MOV_REG_REG",
		0x26: "MOV_REG_IMM",
		0x27: "MOV_MEM_REG",
		0x28: "MOV_REG_MEM",
		0x29: "MOV_MEM_REG_OFFSET",
		0x2a: "MOV_REG_MEM_OFFSET",
		0x2b: "MOV_REG_SP",
		0x2c: "MOV_SP_REG",
		0x2d: "MOV_REG_SP_OFFSET",
		0x2e: "MOV_SP_REG_OFFSET",
		0x2f: "MOV_REG_SP_OFFSET_REG",
		0x30: "MOV_SP_REG_OFFSET_REG",
		0x31: "XOR_REG_REG",
		0x32: "XOR_REG_IMM",
		0x33: "AND_REG_REG",
		0x34: "AND_REG_IMM",
		0x35: "OR_REG_REG",
		0x36: "OR_REG_IMM",
		0x37: "NOT_REG",
		0x38: "SHL_REG",
		0x39: "SHR_REG",
		0x3a: "ROL_REG",
		0x3b: "ROR_REG",
		0x3c: "CLC_IMPL",
		0x3d: "PUSH_IP",
		0x3e: "POP_IP",
		0x3f: "PUSH_SP",
		0x40: "POP_SP",
		0x41: "PUSH_REG",
		0x42: "PUSH_IMM",
		0x43: "POP_REG",
		0x44: "CALL_IMPL",
		0x45: "RET_IMPL",
	}
)
