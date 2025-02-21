package unit_test

import (
	"asme8/assembler/test"
	"asme8/common/instruction"
	"testing"
)

var (
	Tests = []*test.TestCase{
		{
			Name: "reg8 imm8",
			Code: `
    mov a, 0x10
	mov b, 0x20
	mov c, 0x30
	mov d, 0x40
	mov e, 0x50
    add a, 0x10
	add b, 0x20
	add c, 0x30
	add d, 0x40
	add e, 0x50
    adc a, 0x10
	adc b, 0x20
	adc c, 0x30
	adc d, 0x40
	adc e, 0x50
    sub a, 0x10
	sub b, 0x20
	sub c, 0x30
	sub d, 0x40
	sub e, 0x50
    sbb a, 0x10
	sbb b, 0x20
	sbb c, 0x30
	sbb d, 0x40
	sbb e, 0x50
    cmp a, 0x10
	cmp b, 0x20
	cmp c, 0x30
	cmp d, 0x40
	cmp e, 0x50
    and a, 0x10
	and b, 0x20
	and c, 0x30
	and d, 0x40
	and e, 0x50
    xor a, 0x10
	xor b, 0x20
	xor c, 0x30
	xor d, 0x40
	xor e, 0x50
    or a, 0x10
	or b, 0x20
	or c, 0x30
	or d, 0x40
	or e, 0x50
`,
			Expected: []uint8{
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x10,
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_B, 0x20,
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_C, 0x30,
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_D, 0x40,
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_E, 0x50,
				instruction.INST_ADD_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x10,
				instruction.INST_ADD_REG8_IMM8, instruction.REGISTER_OPCODE_B, 0x20,
				instruction.INST_ADD_REG8_IMM8, instruction.REGISTER_OPCODE_C, 0x30,
				instruction.INST_ADD_REG8_IMM8, instruction.REGISTER_OPCODE_D, 0x40,
				instruction.INST_ADD_REG8_IMM8, instruction.REGISTER_OPCODE_E, 0x50,
				instruction.INST_ADC_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x10,
				instruction.INST_ADC_REG8_IMM8, instruction.REGISTER_OPCODE_B, 0x20,
				instruction.INST_ADC_REG8_IMM8, instruction.REGISTER_OPCODE_C, 0x30,
				instruction.INST_ADC_REG8_IMM8, instruction.REGISTER_OPCODE_D, 0x40,
				instruction.INST_ADC_REG8_IMM8, instruction.REGISTER_OPCODE_E, 0x50,
				instruction.INST_SUB_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x10,
				instruction.INST_SUB_REG8_IMM8, instruction.REGISTER_OPCODE_B, 0x20,
				instruction.INST_SUB_REG8_IMM8, instruction.REGISTER_OPCODE_C, 0x30,
				instruction.INST_SUB_REG8_IMM8, instruction.REGISTER_OPCODE_D, 0x40,
				instruction.INST_SUB_REG8_IMM8, instruction.REGISTER_OPCODE_E, 0x50,
				instruction.INST_SBB_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x10,
				instruction.INST_SBB_REG8_IMM8, instruction.REGISTER_OPCODE_B, 0x20,
				instruction.INST_SBB_REG8_IMM8, instruction.REGISTER_OPCODE_C, 0x30,
				instruction.INST_SBB_REG8_IMM8, instruction.REGISTER_OPCODE_D, 0x40,
				instruction.INST_SBB_REG8_IMM8, instruction.REGISTER_OPCODE_E, 0x50,
				instruction.INST_CMP_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x10,
				instruction.INST_CMP_REG8_IMM8, instruction.REGISTER_OPCODE_B, 0x20,
				instruction.INST_CMP_REG8_IMM8, instruction.REGISTER_OPCODE_C, 0x30,
				instruction.INST_CMP_REG8_IMM8, instruction.REGISTER_OPCODE_D, 0x40,
				instruction.INST_CMP_REG8_IMM8, instruction.REGISTER_OPCODE_E, 0x50,
				instruction.INST_AND_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x10,
				instruction.INST_AND_REG8_IMM8, instruction.REGISTER_OPCODE_B, 0x20,
				instruction.INST_AND_REG8_IMM8, instruction.REGISTER_OPCODE_C, 0x30,
				instruction.INST_AND_REG8_IMM8, instruction.REGISTER_OPCODE_D, 0x40,
				instruction.INST_AND_REG8_IMM8, instruction.REGISTER_OPCODE_E, 0x50,
				instruction.INST_XOR_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x10,
				instruction.INST_XOR_REG8_IMM8, instruction.REGISTER_OPCODE_B, 0x20,
				instruction.INST_XOR_REG8_IMM8, instruction.REGISTER_OPCODE_C, 0x30,
				instruction.INST_XOR_REG8_IMM8, instruction.REGISTER_OPCODE_D, 0x40,
				instruction.INST_XOR_REG8_IMM8, instruction.REGISTER_OPCODE_E, 0x50,
				instruction.INST_OR_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x10,
				instruction.INST_OR_REG8_IMM8, instruction.REGISTER_OPCODE_B, 0x20,
				instruction.INST_OR_REG8_IMM8, instruction.REGISTER_OPCODE_C, 0x30,
				instruction.INST_OR_REG8_IMM8, instruction.REGISTER_OPCODE_D, 0x40,
				instruction.INST_OR_REG8_IMM8, instruction.REGISTER_OPCODE_E, 0x50,
			},
		},
		{
			Name: "reg8 imm8 tag",
			Code: `
VALUE_1=0x10
VALUE_2=0x20
VALUE_3=0x30
VALUE_4=0x40
VALUE_5=0x50

    mov a, VALUE_1
    mov b, VALUE_2
    mov c, VALUE_3
    mov d, VALUE_4
    mov e, VALUE_5
`,
			Expected: []byte{
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x10,
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_B, 0x20,
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_C, 0x30,
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_D, 0x40,
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_E, 0x50,
			},
		},
		{
			Name: "impl imm16",
			Code: `
    jmp 0x1234
    jz 0x1234
    jnz 0x1234
    js 0x1234
    jns 0x1234
    jc 0x1234
    jnc 0x1234
`,
			Expected: []byte{
				instruction.INST_JMP_IMM16, 0x34, 0x12,
				instruction.INST_JZ_IMM16, 0x34, 0x12,
				instruction.INST_JNZ_IMM16, 0x34, 0x12,
				instruction.INST_JS_IMM16, 0x34, 0x12,
				instruction.INST_JNS_IMM16, 0x34, 0x12,
				instruction.INST_JC_IMM16, 0x34, 0x12,
				instruction.INST_JNC_IMM16, 0x34, 0x12,
			},
		},
		{
			Name: "impl imm16 tag",
			Code: `

ADDR=0x1234

    jmp ADDR
    jz ADDR
    jnz ADDR
    js ADDR
    jns ADDR
    jc ADDR
    jnc ADDR

`,
			Expected: []byte{
				instruction.INST_JMP_IMM16, 0x34, 0x12,
				instruction.INST_JZ_IMM16, 0x34, 0x12,
				instruction.INST_JNZ_IMM16, 0x34, 0x12,
				instruction.INST_JS_IMM16, 0x34, 0x12,
				instruction.INST_JNS_IMM16, 0x34, 0x12,
				instruction.INST_JC_IMM16, 0x34, 0x12,
				instruction.INST_JNC_IMM16, 0x34, 0x12,
			},
		},
		{
			Name: "impl reg8",
			Code: `

    not a
	shl b
	shr c
	rol d
	ror e
	inc a
	dec b
	push c
	pop d

`,
			Expected: []byte{
				instruction.INST_NOT_REG8, instruction.REGISTER_OPCODE_A,
				instruction.INST_SHL_REG8, instruction.REGISTER_OPCODE_B,
				instruction.INST_SHR_REG8, instruction.REGISTER_OPCODE_C,
				instruction.INST_ROL_REG8, instruction.REGISTER_OPCODE_D,
				instruction.INST_ROR_REG8, instruction.REGISTER_OPCODE_E,
				instruction.INST_INC_REG8, instruction.REGISTER_OPCODE_A,
				instruction.INST_DEC_REG8, instruction.REGISTER_OPCODE_B,
				instruction.INST_PUSH_REG8, instruction.REGISTER_OPCODE_C,
				instruction.INST_POP_REG8, instruction.REGISTER_OPCODE_D,
			},
		},
		{
			Name: "impl reg16",
			Code: `

    inc ip
    inc bp
	dec ip
	dec bp
	push ip
	push bp
	pop ip
	pop bp
	call ip
	call bp

`,
			Expected: []byte{
				instruction.INST_INC_REG16, instruction.REGISTER_OPCODE_IP,
				instruction.INST_INC_REG16, instruction.REGISTER_OPCODE_BP,
				instruction.INST_DEC_REG16, instruction.REGISTER_OPCODE_IP,
				instruction.INST_DEC_REG16, instruction.REGISTER_OPCODE_BP,
				instruction.INST_PUSH_REG16, instruction.REGISTER_OPCODE_IP,
				instruction.INST_PUSH_REG16, instruction.REGISTER_OPCODE_BP,
				instruction.INST_POP_REG16, instruction.REGISTER_OPCODE_IP,
				instruction.INST_POP_REG16, instruction.REGISTER_OPCODE_BP,
				instruction.INST_CALL_REG16, instruction.REGISTER_OPCODE_IP,
				instruction.INST_CALL_REG16, instruction.REGISTER_OPCODE_BP,
			},
		},
		{
			Name: "reg8 reg8 and reg16 reg 16",
			Code: `
    mov a, b
	mov b, c
	mov c, d
	mov d, e
	mov e, a
    add a, b
	add b, c
	add c, d
	add d, e
	add e, a
    adc a, b
	adc b, c
	adc c, d
	adc d, e
	adc e, a
    sub a, b
	sub b, c
	sub c, d
	sub d, e
	sub e, a
    sbb a, b
	sbb b, c
	sbb c, d
	sbb d, e
	sbb e, a
    cmp a, b
	cmp b, c
	cmp c, d
	cmp d, e
	cmp e, a
    and a, b
	and b, c
	and c, d
	and d, e
	and e, a
    xor a, b
	xor b, c
	xor c, d
	xor d, e
	xor e, a
    or a, b
	or b, c
	or c, d
	or d, e
	or e, a
	mov ip, bp
	mov bp, ip

`,
			Expected: []byte{
				instruction.INST_MOV_REG8_REG8, instruction.REGISTER_OPCODE_B<<4 | instruction.REGISTER_OPCODE_A,
				instruction.INST_MOV_REG8_REG8, instruction.REGISTER_OPCODE_C<<4 | instruction.REGISTER_OPCODE_B,
				instruction.INST_MOV_REG8_REG8, instruction.REGISTER_OPCODE_D<<4 | instruction.REGISTER_OPCODE_C,
				instruction.INST_MOV_REG8_REG8, instruction.REGISTER_OPCODE_E<<4 | instruction.REGISTER_OPCODE_D,
				instruction.INST_MOV_REG8_REG8, instruction.REGISTER_OPCODE_A<<4 | instruction.REGISTER_OPCODE_E,
				instruction.INST_ADD_REG8_REG8, instruction.REGISTER_OPCODE_B<<4 | instruction.REGISTER_OPCODE_A,
				instruction.INST_ADD_REG8_REG8, instruction.REGISTER_OPCODE_C<<4 | instruction.REGISTER_OPCODE_B,
				instruction.INST_ADD_REG8_REG8, instruction.REGISTER_OPCODE_D<<4 | instruction.REGISTER_OPCODE_C,
				instruction.INST_ADD_REG8_REG8, instruction.REGISTER_OPCODE_E<<4 | instruction.REGISTER_OPCODE_D,
				instruction.INST_ADD_REG8_REG8, instruction.REGISTER_OPCODE_A<<4 | instruction.REGISTER_OPCODE_E,
				instruction.INST_ADC_REG8_REG8, instruction.REGISTER_OPCODE_B<<4 | instruction.REGISTER_OPCODE_A,
				instruction.INST_ADC_REG8_REG8, instruction.REGISTER_OPCODE_C<<4 | instruction.REGISTER_OPCODE_B,
				instruction.INST_ADC_REG8_REG8, instruction.REGISTER_OPCODE_D<<4 | instruction.REGISTER_OPCODE_C,
				instruction.INST_ADC_REG8_REG8, instruction.REGISTER_OPCODE_E<<4 | instruction.REGISTER_OPCODE_D,
				instruction.INST_ADC_REG8_REG8, instruction.REGISTER_OPCODE_A<<4 | instruction.REGISTER_OPCODE_E,
				instruction.INST_SUB_REG8_REG8, instruction.REGISTER_OPCODE_B<<4 | instruction.REGISTER_OPCODE_A,
				instruction.INST_SUB_REG8_REG8, instruction.REGISTER_OPCODE_C<<4 | instruction.REGISTER_OPCODE_B,
				instruction.INST_SUB_REG8_REG8, instruction.REGISTER_OPCODE_D<<4 | instruction.REGISTER_OPCODE_C,
				instruction.INST_SUB_REG8_REG8, instruction.REGISTER_OPCODE_E<<4 | instruction.REGISTER_OPCODE_D,
				instruction.INST_SUB_REG8_REG8, instruction.REGISTER_OPCODE_A<<4 | instruction.REGISTER_OPCODE_E,
				instruction.INST_SBB_REG8_REG8, instruction.REGISTER_OPCODE_B<<4 | instruction.REGISTER_OPCODE_A,
				instruction.INST_SBB_REG8_REG8, instruction.REGISTER_OPCODE_C<<4 | instruction.REGISTER_OPCODE_B,
				instruction.INST_SBB_REG8_REG8, instruction.REGISTER_OPCODE_D<<4 | instruction.REGISTER_OPCODE_C,
				instruction.INST_SBB_REG8_REG8, instruction.REGISTER_OPCODE_E<<4 | instruction.REGISTER_OPCODE_D,
				instruction.INST_SBB_REG8_REG8, instruction.REGISTER_OPCODE_A<<4 | instruction.REGISTER_OPCODE_E,
				instruction.INST_CMP_REG8_REG8, instruction.REGISTER_OPCODE_B<<4 | instruction.REGISTER_OPCODE_A,
				instruction.INST_CMP_REG8_REG8, instruction.REGISTER_OPCODE_C<<4 | instruction.REGISTER_OPCODE_B,
				instruction.INST_CMP_REG8_REG8, instruction.REGISTER_OPCODE_D<<4 | instruction.REGISTER_OPCODE_C,
				instruction.INST_CMP_REG8_REG8, instruction.REGISTER_OPCODE_E<<4 | instruction.REGISTER_OPCODE_D,
				instruction.INST_CMP_REG8_REG8, instruction.REGISTER_OPCODE_A<<4 | instruction.REGISTER_OPCODE_E,
				instruction.INST_AND_REG8_REG8, instruction.REGISTER_OPCODE_B<<4 | instruction.REGISTER_OPCODE_A,
				instruction.INST_AND_REG8_REG8, instruction.REGISTER_OPCODE_C<<4 | instruction.REGISTER_OPCODE_B,
				instruction.INST_AND_REG8_REG8, instruction.REGISTER_OPCODE_D<<4 | instruction.REGISTER_OPCODE_C,
				instruction.INST_AND_REG8_REG8, instruction.REGISTER_OPCODE_E<<4 | instruction.REGISTER_OPCODE_D,
				instruction.INST_AND_REG8_REG8, instruction.REGISTER_OPCODE_A<<4 | instruction.REGISTER_OPCODE_E,
				instruction.INST_XOR_REG8_REG8, instruction.REGISTER_OPCODE_B<<4 | instruction.REGISTER_OPCODE_A,
				instruction.INST_XOR_REG8_REG8, instruction.REGISTER_OPCODE_C<<4 | instruction.REGISTER_OPCODE_B,
				instruction.INST_XOR_REG8_REG8, instruction.REGISTER_OPCODE_D<<4 | instruction.REGISTER_OPCODE_C,
				instruction.INST_XOR_REG8_REG8, instruction.REGISTER_OPCODE_E<<4 | instruction.REGISTER_OPCODE_D,
				instruction.INST_XOR_REG8_REG8, instruction.REGISTER_OPCODE_A<<4 | instruction.REGISTER_OPCODE_E,
				instruction.INST_OR_REG8_REG8, instruction.REGISTER_OPCODE_B<<4 | instruction.REGISTER_OPCODE_A,
				instruction.INST_OR_REG8_REG8, instruction.REGISTER_OPCODE_C<<4 | instruction.REGISTER_OPCODE_B,
				instruction.INST_OR_REG8_REG8, instruction.REGISTER_OPCODE_D<<4 | instruction.REGISTER_OPCODE_C,
				instruction.INST_OR_REG8_REG8, instruction.REGISTER_OPCODE_E<<4 | instruction.REGISTER_OPCODE_D,
				instruction.INST_OR_REG8_REG8, instruction.REGISTER_OPCODE_A<<4 | instruction.REGISTER_OPCODE_E,
				instruction.INST_MOV_REG16_REG16, instruction.REGISTER_OPCODE_BP<<4 | instruction.REGISTER_OPCODE_IP,
				instruction.INST_MOV_REG16_REG16, instruction.REGISTER_OPCODE_IP<<4 | instruction.REGISTER_OPCODE_BP,
			},
		},
		{
			Name: "impl",
			Code: `
	ret
	rti
	clc
	brk
	nop

`,
			Expected: []byte{
				instruction.INST_RET_IMPL,
				instruction.INST_RTI_IMPL,
				instruction.INST_CLC_IMPL,
				instruction.INST_BRK_IMPL,
				instruction.INST_NOP_IMPL,
			},
		},
		{
			Name: "reg16 imm8",
			Code: `

	add ip, 0x10
	add bp, 0x20
	sub ip, 0x30
	sub bp, 0x40

`,
			Expected: []byte{
				instruction.INST_ADD_REG16_IMM8, instruction.REGISTER_OPCODE_IP, 0x10,
				instruction.INST_ADD_REG16_IMM8, instruction.REGISTER_OPCODE_BP, 0x20,
				instruction.INST_SUB_REG16_IMM8, instruction.REGISTER_OPCODE_IP, 0x30,
				instruction.INST_SUB_REG16_IMM8, instruction.REGISTER_OPCODE_BP, 0x40,
			},
		},
		{
			Name: "reg16 imm16",
			Code: `

	mov ip, 0x1234
	mov bp, 0x1234

`,
			Expected: []byte{
				instruction.INST_MOV_REG16_IMM16, instruction.REGISTER_OPCODE_IP, 0x34, 0x12,
				instruction.INST_MOV_REG16_IMM16, instruction.REGISTER_OPCODE_BP, 0x34, 0x12,
			},
		},
		{
			Name: "reg16 imm8 tag",
			Code: `

VALUE_1=0x10
VALUE_2=0x20
VALUE_3=0x30
VALUE_4=0x40

	add ip, VALUE_1
	add bp, VALUE_2
	sub ip, VALUE_3
	sub bp, VALUE_4

`,
			Expected: []byte{
				instruction.INST_ADD_REG16_IMM8, instruction.REGISTER_OPCODE_IP, 0x10,
				instruction.INST_ADD_REG16_IMM8, instruction.REGISTER_OPCODE_BP, 0x20,
				instruction.INST_SUB_REG16_IMM8, instruction.REGISTER_OPCODE_IP, 0x30,
				instruction.INST_SUB_REG16_IMM8, instruction.REGISTER_OPCODE_BP, 0x40,
			},
		},
		{
			Name: "reg16 imm16 tag",
			Code: `

ADDR=0x1234

	mov ip, ADDR
	mov bp, ADDR

`,
			Expected: []byte{
				instruction.INST_MOV_REG16_IMM16, instruction.REGISTER_OPCODE_IP, 0x34, 0x12,
				instruction.INST_MOV_REG16_IMM16, instruction.REGISTER_OPCODE_BP, 0x34, 0x12,
			},
		},
		{
			Name: "implied stack",
			Code: `

	push sp
	pop sp

`,
			Expected: []byte{
				instruction.INST_PUSH_SP,
				instruction.INST_POP_SP,
			},
		},
		{
			Name: "stack imm8",
			Code: `

	add sp, 0x10
	sub sp, 0x10


`,
			Expected: []byte{
				instruction.INST_ADD_SP_IMM8, 0x10,
				instruction.INST_SUB_SP_IMM8, 0x10,
			},
		},
		{
			Name: "stack imm8 tag",
			Code: `

VALUE=0x10

	add sp, VALUE
	sub sp, VALUE


`,
			Expected: []byte{
				instruction.INST_ADD_SP_IMM8, 0x10,
				instruction.INST_SUB_SP_IMM8, 0x10,
			},
		},
		{
			Name: "stack reg16",
			Code: `

	mov sp, ip
	mov sp, bp
	mov ip, sp
	mov bp, sp

`,
			Expected: []byte{
				instruction.INST_MOV_REG16_SP, instruction.REGISTER_OPCODE_IP,
				instruction.INST_MOV_REG16_SP, instruction.REGISTER_OPCODE_BP,
				instruction.INST_MOV_SP_REG16, instruction.REGISTER_OPCODE_IP,
				instruction.INST_MOV_SP_REG16, instruction.REGISTER_OPCODE_BP,
			},
		},
		{
			Name: "direct",
			Code: `

ADDR=0xabcd

	mov a, [0x1234]
	mov [0x1234], b
	mov a, [ADDR]
	mov [ADDR], b

`,
			Expected: []byte{
				instruction.INST_MOV_DIRECT_REG8, instruction.REGISTER_OPCODE_A, 0x34, 0x12,
				instruction.INST_MOV_REG8_DIRECT, instruction.REGISTER_OPCODE_B, 0x34, 0x12,
				instruction.INST_MOV_DIRECT_REG8, instruction.REGISTER_OPCODE_A, 0xcd, 0xab,
				instruction.INST_MOV_REG8_DIRECT, instruction.REGISTER_OPCODE_B, 0xcd, 0xab,
			},
		},
		{
			Name: "indexed",
			Code: `

ADDR=0xabcd

	mov a, [0x1234+b]
	mov [0x1234+a], b
	mov a, [ADDR+b]
	mov [ADDR+a], b

`,
			Expected: []byte{
				instruction.INST_MOV_MEM_INDEXED_REG8, instruction.REGISTER_OPCODE_A<<4 | instruction.REGISTER_OPCODE_B, 0x34, 0x12,
				instruction.INST_MOV_REG8_MEM_INDEXED, instruction.REGISTER_OPCODE_B<<4 | instruction.REGISTER_OPCODE_A, 0x34, 0x12,
				instruction.INST_MOV_MEM_INDEXED_REG8, instruction.REGISTER_OPCODE_A<<4 | instruction.REGISTER_OPCODE_B, 0xcd, 0xab,
				instruction.INST_MOV_REG8_MEM_INDEXED, instruction.REGISTER_OPCODE_B<<4 | instruction.REGISTER_OPCODE_A, 0xcd, 0xab,
			},
		},
		{
			Name: "indirect",
			Code: `

	mov a, [ip]
	mov [ip], b
	mov a, [bp]
	mov [bp], b

`,
			Expected: []byte{
				instruction.INST_MOV_INDIRECT_REG8, instruction.REGISTER_OPCODE_A<<4 | instruction.REGISTER_OPCODE_IP,
				instruction.INST_MOV_REG8_INDIRECT, instruction.REGISTER_OPCODE_IP<<4 | instruction.REGISTER_OPCODE_B,
				instruction.INST_MOV_INDIRECT_REG8, instruction.REGISTER_OPCODE_A<<4 | instruction.REGISTER_OPCODE_BP,
				instruction.INST_MOV_REG8_INDIRECT, instruction.REGISTER_OPCODE_BP<<4 | instruction.REGISTER_OPCODE_B,
			},
		},
		{
			Name: "indirect offset",
			Code: `

	mov a, [ip+4]
	mov [ip+8], b
	mov a, [bp+4]
	mov [bp+8], b

`,
			Expected: []byte{
				instruction.INST_MOV_INDIRECT_OFFSET_REG8, instruction.REGISTER_OPCODE_A<<4 | instruction.REGISTER_OPCODE_IP, 0x04,
				instruction.INST_MOV_REG8_INDIRECT_OFFSET, instruction.REGISTER_OPCODE_IP<<4 | instruction.REGISTER_OPCODE_B, 0x08,
				instruction.INST_MOV_INDIRECT_OFFSET_REG8, instruction.REGISTER_OPCODE_A<<4 | instruction.REGISTER_OPCODE_BP, 0x04,
				instruction.INST_MOV_REG8_INDIRECT_OFFSET, instruction.REGISTER_OPCODE_BP<<4 | instruction.REGISTER_OPCODE_B, 0x08,
			},
		},
		{
			Name: "indexed reg16",
			Code: `

	mov a, [ip+c]
	mov [ip+d], b
	mov a, [bp+c]
	mov [bp+d], b

`,
			Expected: []byte{
				instruction.INST_MOV_INDEXED_REG8, instruction.REGISTER_OPCODE_IP<<4 | instruction.REGISTER_OPCODE_C, instruction.REGISTER_OPCODE_A,
				instruction.INST_MOV_REG8_INDEXED, instruction.REGISTER_OPCODE_IP<<4 | instruction.REGISTER_OPCODE_D, instruction.REGISTER_OPCODE_B,
				instruction.INST_MOV_INDEXED_REG8, instruction.REGISTER_OPCODE_BP<<4 | instruction.REGISTER_OPCODE_C, instruction.REGISTER_OPCODE_A,
				instruction.INST_MOV_REG8_INDEXED, instruction.REGISTER_OPCODE_BP<<4 | instruction.REGISTER_OPCODE_D, instruction.REGISTER_OPCODE_B,
			},
		},
	}
)

func TestCore(t *testing.T) {
	// Run only test cases with debug flag
	for _, tc := range Tests {
		if tc.Debug {
			test.RunCase(t, tc)
			return
		}
	}
	// Run all test case unless disabled
	for _, tc := range Tests {
		test.RunCase(t, tc)
	}
}
