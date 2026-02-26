package instruction

type Type = uint8

const (
	INST_NOP_IMPL                       Type = 0x00 // example: nop
	INST_JMP_REG16                      Type = 0x01 // example: jmp ip
	INST_JMP_IMM16                      Type = 0x02 // example: jmp 0x1234
	INST_JS_IMM16                       Type = 0x03 // example: js 0x1234
	INST_JNS_IMM16                      Type = 0x04 // example: jns 0x1234
	INST_JZ_IMM16                       Type = 0x05 // example: jz 0x1234
	INST_JNZ_IMM16                      Type = 0x06 // example: jnz 0x1234
	INST_JC_IMM16                       Type = 0x07 // example: jc 0x1234
	INST_JNC_IMM16                      Type = 0x08 // example: jnc 0x1234
	INST_CMP_REG8_REG8                  Type = 0x09 // example: cmp a, b
	INST_CMP_REG8_IMM8                  Type = 0x0a // example: cmp a, 0x10
	INST_ADD_REG8_REG8                  Type = 0x0b // example: add a, b
	INST_ADD_REG8_IMM8                  Type = 0x0c // example: add a, 0x10
	INST_ADD_REG16_IMM8                 Type = 0x0d // example: add ip, 0x10		(not necessary, use 'add ipl, 0x10' and 'adc iph, 0')
	INST_ADC_REG8_REG8                  Type = 0x0e // example: adc a, b
	INST_ADC_REG8_IMM8                  Type = 0x0f // example: adc a, 0x10
	INST_SUB_REG8_REG8                  Type = 0x10 // example: sub a, b
	INST_SUB_REG8_IMM8                  Type = 0x11 // example: sub a, 0x10
	INST_SUB_REG16_IMM8                 Type = 0x12 // example: sub ip, 0x10		(not necessary, use 'sub ipl, 0x10' and 'sbb iph, 0')
	INST_SBB_REG8_REG8                  Type = 0x13 // example: sbb a, b
	INST_SBB_REG8_IMM8                  Type = 0x14 // example: sbb a, 0x10
	INST_INC_REG8                       Type = 0x15 // example: inc a 				(not necessary, use 'add a, 1')
	INST_INC_REG16                      Type = 0x16 // example: inc ip 				(not necessary, use 'add ipl, 1' and 'adc iph, 0')
	INST_DEC_REG8                       Type = 0x17 // example: dec a 				(not necessary, use 'sub a, 1')
	INST_DEC_REG16                      Type = 0x18 // example: dec ip 				(not necessary, use 'sub ipl, 1' and 'sbb iph, 0')
	INST_MOV_REG8_REG8                  Type = 0x19 // example: mov a, b
	INST_MOV_REG8_IMM8                  Type = 0x1a // example: mov a, 0x10
	INST_MOV_MEM_TO_REG_DIRECT          Type = 0x1b // example: mov a, [0x1234]
	INST_MOV_REG_TO_MEM_DIRECT          Type = 0x1c // example: mov [0x1234], a
	INST_MOV_MEM_TO_REG_INDEXED         Type = 0x1d // example: mov a, [0x1234+b]
	INST_MOV_REG_TO_MEM_INDEXED         Type = 0x1e // example: mov [0x1234+b], a
	INST_MOV_REG16_IMM16                Type = 0x1f // example: mov ip, 0x1234
	INST_MOV_REG16_REG16                Type = 0x20 // example: mov ip, bp 			(not necessary, use 'mov ipl, bpl' and 'mov iph, bph')
	INST_MOV_REG_TO_MEM_INDIRECT        Type = 0x21 // example: mov [ip], a
	INST_MOV_MEM_TO_REG_INDIRECT        Type = 0x22 // example: mov a, [ip]
	INST_MOV_REG_TO_MEM_INDIRECT_OFFSET Type = 0x23 // example: mov [ip+8], a
	INST_MOV_MEM_TO_REG_INDIRECT_OFFSET Type = 0x24 // example: mov a, [ip+8]
	INST_MOV_REG_TO_MEM_REG16_INDEXED   Type = 0x25 // example: mov [ip+b], a
	INST_MOV_MEM_TO_REG_REG16_INDEXED   Type = 0x26 // example: mov a, [ip+b]
	INST_XOR_REG8_REG8                  Type = 0x27 // example: xor a, b
	INST_XOR_REG8_IMM8                  Type = 0x28 // example: xor a, 0x10
	INST_AND_REG8_REG8                  Type = 0x29 // example: and a, b
	INST_AND_REG8_IMM8                  Type = 0x2a // example: and a, 0x10
	INST_OR_REG8_REG8                   Type = 0x2b // example: or a, b
	INST_OR_REG8_IMM8                   Type = 0x2c // example: or a, 0x10
	_                                   Type = 0x2d
	INST_SHL_REG8                       Type = 0x2e // example: shl a
	INST_SHR_REG8                       Type = 0x2f // example: shr a
	INST_ROL_REG8                       Type = 0x30 // example: rol a
	INST_ROR_REG8                       Type = 0x31 // example: ror a
	INST_CLC_IMPL                       Type = 0x32 // example: clc
	INST_CLI_IMPL                       Type = 0x33 // example: cli
	INST_STI_IMPL                       Type = 0x34 // example: sti
	INST_PUSH_REG8                      Type = 0x35 // example: push a
	INST_PUSH_REG16                     Type = 0x36 // example: push ip 			(not necessary, use 'push ipl' and 'push iph')
	INST_PUSH_SR                        Type = 0x37 // example: push sr
	INST_POP_REG8                       Type = 0x38 // example: pop a
	INST_POP_REG16                      Type = 0x39 // example: pop ip				(not necessary, use 'pop iph' and 'pop ipl')
	INST_POP_SR                         Type = 0x3a // example: pop sr
	INST_CALL_IMM16                     Type = 0x3b // example: call 0x1234
	INST_CALL_REG16                     Type = 0x3c // example: call ip
	INST_RET_IMPL                       Type = 0x3d // example: ret
	INST_RTI_IMPL                       Type = 0x3e // example: rti
	INST_HLT_IMPL                       Type = 0x3f // example: hlt
)

const (
	TEST_INST_SET_IRQ Type = 0xff
)

const (
	INST_MNEMONICS_ADC  string = "adc"
	INST_MNEMONICS_ADD  string = "add"
	INST_MNEMONICS_AND  string = "and"
	INST_MNEMONICS_HLT  string = "hlt"
	INST_MNEMONICS_CALL string = "call"
	INST_MNEMONICS_INT  string = "int"
	INST_MNEMONICS_CLC  string = "clc"
	INST_MNEMONICS_CLI  string = "cli"
	INST_MNEMONICS_STI  string = "sti"
	INST_MNEMONICS_CMP  string = "cmp"
	INST_MNEMONICS_DEC  string = "dec"
	INST_MNEMONICS_INC  string = "inc"
	INST_MNEMONICS_JC   string = "jc"
	INST_MNEMONICS_JMP  string = "jmp"
	INST_MNEMONICS_JNC  string = "jnc"
	INST_MNEMONICS_JNS  string = "jns"
	INST_MNEMONICS_JNZ  string = "jnz"
	INST_MNEMONICS_JS   string = "js"
	INST_MNEMONICS_JZ   string = "jz"
	INST_MNEMONICS_MOV  string = "mov"
	INST_MNEMONICS_NOP  string = "nop"
	INST_MNEMONICS_OR   string = "or"
	INST_MNEMONICS_POP  string = "pop"
	INST_MNEMONICS_PUSH string = "push"
	INST_MNEMONICS_RET  string = "ret"
	INST_MNEMONICS_ROL  string = "rol"
	INST_MNEMONICS_ROR  string = "ror"
	INST_MNEMONICS_RTI  string = "rti"
	INST_MNEMONICS_SBB  string = "sbb"
	INST_MNEMONICS_SHL  string = "shl"
	INST_MNEMONICS_SHR  string = "shr"
	INST_MNEMONICS_SUB  string = "sub"
	INST_MNEMONICS_XOR  string = "xor"
)

var (
	INST_HUMAN_READABLE = map[uint8]string{
		INST_HLT_IMPL:                       "HLT_IMPL",
		INST_NOP_IMPL:                       "NOP_IMPL",
		INST_JMP_REG16:                      "JMP_R16",
		INST_JMP_IMM16:                      "JMP_IMPL_IMM16",
		INST_JS_IMM16:                       "JS_IMPL_IMM16",
		INST_JNS_IMM16:                      "JNS_IMPL_IMM16",
		INST_JZ_IMM16:                       "JZ_IMPL_IMM16",
		INST_JNZ_IMM16:                      "JNZ_IMPL_IMM16",
		INST_JC_IMM16:                       "JC_IMPL_IMM16",
		INST_JNC_IMM16:                      "JNC_IMPL_IMM16",
		INST_CMP_REG8_REG8:                  "CMP_R8_R8",
		INST_CMP_REG8_IMM8:                  "CMP_R8_IMM8",
		INST_ADD_REG8_REG8:                  "ADD_R8_R8",
		INST_ADD_REG8_IMM8:                  "ADD_R8_IMM8",
		INST_ADD_REG16_IMM8:                 "ADD_R16_IMM8",
		INST_ADC_REG8_REG8:                  "ADC_R8_R8",
		INST_ADC_REG8_IMM8:                  "ADC_R8_IMM8",
		INST_SUB_REG8_REG8:                  "SUB_R8_R8",
		INST_SUB_REG8_IMM8:                  "SUB_R8_IMM8",
		INST_SUB_REG16_IMM8:                 "SUB_R16_IMM8",
		INST_SBB_REG8_REG8:                  "SBB_R8_R8",
		INST_SBB_REG8_IMM8:                  "SBB_R8_IMM8",
		INST_INC_REG8:                       "INC_R8",
		INST_INC_REG16:                      "INC_R16",
		INST_DEC_REG8:                       "DEC_R8",
		INST_DEC_REG16:                      "DEC_R16",
		INST_MOV_REG8_REG8:                  "MV_R8_R8",
		INST_MOV_REG8_IMM8:                  "MV_R8_IMM8",
		INST_MOV_MEM_TO_REG_DIRECT:          "MV_MtR_DRT",
		INST_MOV_REG_TO_MEM_DIRECT:          "MV_RtM_DRT",
		INST_MOV_MEM_TO_REG_INDEXED:         "MV_MtR_INDX",
		INST_MOV_REG_TO_MEM_INDEXED:         "MV_RtM_INDX",
		INST_MOV_REG16_IMM16:                "MV_R16_IMM16",
		INST_MOV_REG16_REG16:                "MV_R16_R16",
		INST_MOV_REG_TO_MEM_INDIRECT:        "MV_RtM_INDR",
		INST_MOV_MEM_TO_REG_INDIRECT:        "MV_MtR_INDR",
		INST_MOV_REG_TO_MEM_INDIRECT_OFFSET: "MV_RtM_INDR_OFF",
		INST_MOV_MEM_TO_REG_INDIRECT_OFFSET: "MV_MtR_INDR_OFF",
		INST_MOV_REG_TO_MEM_REG16_INDEXED:   "MV_RtM_R16_INDX",
		INST_MOV_MEM_TO_REG_REG16_INDEXED:   "MV_MtR_R16_INDX",
		INST_XOR_REG8_REG8:                  "XOR_R8_R8",
		INST_XOR_REG8_IMM8:                  "XOR_R8_IMM8",
		INST_AND_REG8_REG8:                  "AND_R8_R8",
		INST_AND_REG8_IMM8:                  "AND_R8_IMM8",
		INST_OR_REG8_REG8:                   "OR_R8_R8",
		INST_OR_REG8_IMM8:                   "OR_R8_IMM8",
		INST_SHL_REG8:                       "SHL_R8",
		INST_SHR_REG8:                       "SHR_R8",
		INST_ROL_REG8:                       "ROL_R8",
		INST_ROR_REG8:                       "ROR_R8",
		INST_CLC_IMPL:                       "CLC_IMPL",
		INST_CLI_IMPL:                       "CLI_IMPL",
		INST_STI_IMPL:                       "STI_IMPL",
		INST_PUSH_REG8:                      "PUSH_R8",
		INST_PUSH_REG16:                     "PUSH_R16",
		INST_PUSH_SR:                        "PUSH_SR",
		INST_POP_REG8:                       "POP_R8",
		INST_POP_REG16:                      "POP_R16",
		INST_POP_SR:                         "POP_SR",
		INST_CALL_IMM16:                     "CALL_IMM16",
		INST_CALL_REG16:                     "CALL_R16",
		INST_RET_IMPL:                       "RET_IMPL",
		INST_RTI_IMPL:                       "RTI_IMPL",
		TEST_INST_SET_IRQ:                   "TEST_IRQ",
	}
)
