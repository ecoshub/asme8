package instruction

type Type = uint8

const (
	INST_BRK_IMPL                       Type = iota // example: brk
	INST_NOP_IMPL                                   // example: nop
	INST_JMP_REG16                                  // example: jmp ip
	INST_JMP_IMPL_IMM16                             // example: jmp 0x1234
	INST_JS_IMPL_IMM16                              // example: js 0x1234
	INST_JNS_IMPL_IMM16                             // example: jns 0x1234
	INST_JZ_IMPL_IMM16                              // example: jz 0x1234
	INST_JNZ_IMPL_IMM16                             // example: jnz 0x1234
	INST_JC_IMPL_IMM16                              // example: jc 0x1234
	INST_JNC_IMPL_IMM16                             // example: jnc 0x1234
	INST_CMP_REG8_REG8                              // example: cmp a, b
	INST_CMP_REG8_IMM8                              // example: cmp a, 0x10
	INST_ADD_REG8_REG8                              // example: add a, b
	INST_ADD_REG8_IMM8                              // example: add a, 0x10
	INST_ADD_SP_IMM8                                // example: add sp, 0x10
	INST_ADD_REG16_IMM8                             // example: add ip, 0x10		(not necessary, use 'add ipl, 0x10' and 'adc iph, 0')
	INST_ADC_REG8_REG8                              // example: adc a, b
	INST_ADC_REG8_IMM8                              // example: adc a, 0x10
	INST_SUB_REG8_REG8                              // example: sub a, b
	INST_SUB_REG8_IMM8                              // example: sub a, 0x10
	INST_SUB_SP_IMM8                                // example: sub sp, 0x10
	INST_SUB_REG16_IMM8                             // example: sub ip, 0x10		(not necessary, use 'sub ipl, 0x10' and 'sbb iph, 0')
	INST_SBB_REG8_REG8                              // example: sbb a, b
	INST_SBB_REG8_IMM8                              // example: sbb a, 0x10
	INST_INC_REG8                                   // example: inc a 				(not necessary, use 'add a, 1')
	INST_INC_REG16                                  // example: inc ip 				(not necessary, use 'add ipl, 1' and 'adc iph, 0')
	INST_DEC_REG8                                   // example: dec a 				(not necessary, use 'sub a, 1')
	INST_DEC_REG16                                  // example: dec ip 				(not necessary, use 'sub ipl, 1' and 'sbb iph, 0')
	INST_MOV_REG8_REG8                              // example: mov a, b
	INST_MOV_REG8_IMM8                              // example: mov a, 0x10
	INST_MOV_MEM_TO_REG_DIRECT                      // example: mov a, [0x1234]
	INST_MOV_REG_TO_MEM_DIRECT                      // example: mov [0x1234], a
	INST_MOV_MEM_TO_REG_INDEXED                     // example: mov a, [0x1234+b]
	INST_MOV_REG_TO_MEM_INDEXED                     // example: mov [0x1234+b], a
	INST_MOV_REG16_IMM16                            // example: mov ip, 0x1234
	INST_MOV_REG16_REG16                            // example: mov ip, bp 			(not necessary, use 'mov ipl, bpl' and 'mov iph, bph')
	INST_MOV_REG16_SP                               // example: mov sp, ip
	INST_MOV_SP_REG16                               // example: mov ip, sp
	INST_MOV_REG_TO_MEM_INDIRECT                    // example: mov [ip], a
	INST_MOV_MEM_TO_REG_INDIRECT                    // example: mov a, [ip]
	INST_MOV_REG_TO_MEM_INDIRECT_OFFSET             // example: mov [ip+8], a
	INST_MOV_MEM_TO_REG_INDIRECT_OFFSET             // example: mov a, [ip+8]
	INST_MOV_REG_TO_MEM_REG16_INDEXED               // example: mov [ip+b], a
	INST_MOV_MEM_TO_REG_REG16_INDEXED               // example: mov a, [ip+b]
	INST_XOR_REG8_REG8                              // example: xor a, b
	INST_XOR_REG8_IMM8                              // example: xor a, 0x10
	INST_AND_REG8_REG8                              // example: and a, b
	INST_AND_REG8_IMM8                              // example: and a, 0x10
	INST_OR_REG8_REG8                               // example: or a, b
	INST_OR_REG8_IMM8                               // example: or a, 0x10
	INST_NOT_REG8                                   // example: not a 				(not necessary, use 'xor a, 0xff')
	INST_SHL_REG8                                   // example: shl a
	INST_SHR_REG8                                   // example: shr a
	INST_ROL_REG8                                   // example: rol a
	INST_ROR_REG8                                   // example: ror a
	INST_CLC_IMPL                                   // example: clc
	INST_PUSH_SP                                    // example: push sp
	INST_PUSH_REG8                                  // example: push a
	INST_PUSH_REG16                                 // example: push ip 			(not necessary, use 'push ipl' and 'push iph')
	INST_PUSH_SR                                    // example: push sr
	INST_POP_SP                                     // example: pop sp
	INST_POP_REG8                                   // example: pop a
	INST_POP_REG16                                  // example: pop ip				(not necessary, use 'pop iph' and 'pop ipl')
	INST_POP_SR                                     // example: pop sr
	INST_CALL_IMM16                                 // example: call 0x1234
	INST_CALL_REG16                                 // example: call ip
	INST_RET_IMPL                                   // example: ret
	INST_RTI_IMPL                                   // example: rti
)

/*

Unique addressing modes
implied_imm16 ...................: <inst> 0x1234
reg_to_mem_indexed ..............: <inst> [0x1234+b], a
reg_to_mem_direct ...............: <inst> [0x1234], a
reg_to_mem_indirect_offset ......: <inst> [ip+8], a
reg_to_mem_reg16_indexed ........: <inst> [ip+b], a
reg_to_mem_indirect .............: <inst> [ip], a
implied_reg8 ....................: <inst> a
reg8_imm8 .......................: <inst> a, 0x10
mem_to_reg_indexed ..............: <inst> a, [0x1234+b]
mem_to_reg_direct ...............: <inst> a, [0x1234]
mem_to_reg_indirect_offset ......: <inst> a, [ip+8]
mem_to_reg_reg16_indexed ........: <inst> a, [ip+b]
mem_to_reg_indirect .............: <inst> a, [ip]
reg8_reg8 .......................: <inst> a, b
implied .........................: <inst> brk
implied_reg16 ...................: <inst> ip
reg16_imm8 ......................: <inst> ip, 0x10
reg16_imm16 .....................: <inst> ip, 0x1234
reg16_reg16 .....................: <inst> ip, bp
reg16_stack .....................: <inst> ip, sp
implied_stack ...................: <inst> sp
stack_imm8 ......................: <inst> sp, 0x10
stack_reg16 .....................: <inst> sp, ip
*/

const (
	INST_MNEMONICS_ADC  string = "adc"
	INST_MNEMONICS_ADD  string = "add"
	INST_MNEMONICS_AND  string = "and"
	INST_MNEMONICS_BRK  string = "brk"
	INST_MNEMONICS_CALL string = "call"
	INST_MNEMONICS_CLC  string = "clc"
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
	INST_MNEMONICS_NOT  string = "not"
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
		INST_BRK_IMPL:                       "BRK_IMPL",
		INST_NOP_IMPL:                       "NOP_IMPL",
		INST_JMP_REG16:                      "JMP_R16",
		INST_JMP_IMPL_IMM16:                 "JMP_IMPL_IMM16",
		INST_JS_IMPL_IMM16:                  "JS_IMPL_IMM16",
		INST_JNS_IMPL_IMM16:                 "JNS_IMPL_IMM16",
		INST_JZ_IMPL_IMM16:                  "JZ_IMPL_IMM16",
		INST_JNZ_IMPL_IMM16:                 "JNZ_IMPL_IMM16",
		INST_JC_IMPL_IMM16:                  "JC_IMPL_IMM16",
		INST_JNC_IMPL_IMM16:                 "JNC_IMPL_IMM16",
		INST_CMP_REG8_REG8:                  "CMP_R8_R8",
		INST_CMP_REG8_IMM8:                  "CMP_R8_IMM8",
		INST_ADD_REG8_REG8:                  "ADD_R8_R8",
		INST_ADD_REG8_IMM8:                  "ADD_R8_IMM8",
		INST_ADD_SP_IMM8:                    "ADD_SP_IMM8",
		INST_ADD_REG16_IMM8:                 "ADD_R16_IMM8",
		INST_ADC_REG8_REG8:                  "ADC_R8_R8",
		INST_ADC_REG8_IMM8:                  "ADC_R8_IMM8",
		INST_SUB_REG8_REG8:                  "SUB_R8_R8",
		INST_SUB_REG8_IMM8:                  "SUB_R8_IMM8",
		INST_SUB_SP_IMM8:                    "SUB_SP_IMM8",
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
		INST_MOV_REG16_SP:                   "MV_R16_SP",
		INST_MOV_SP_REG16:                   "MV_SP_R16",
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
		INST_NOT_REG8:                       "NOT_R8",
		INST_SHL_REG8:                       "SHL_R8",
		INST_SHR_REG8:                       "SHR_R8",
		INST_ROL_REG8:                       "ROL_R8",
		INST_ROR_REG8:                       "ROR_R8",
		INST_CLC_IMPL:                       "CLC_IMPL",
		INST_PUSH_SP:                        "PUSH_SP",
		INST_PUSH_REG8:                      "PUSH_R8",
		INST_PUSH_REG16:                     "PUSH_R16",
		INST_PUSH_SR:                        "PUSH_SR",
		INST_POP_SP:                         "POP_SP",
		INST_POP_REG8:                       "POP_R8",
		INST_POP_REG16:                      "POP_R16",
		INST_POP_SR:                         "POP_SR",
		INST_CALL_IMM16:                     "CALL_IMM16",
		INST_CALL_REG16:                     "CALL_R16",
		INST_RET_IMPL:                       "RET_IMPL",
		INST_RTI_IMPL:                       "RTI_IMPL",
	}
)
