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
    add a, 0x10
	add b, 0x20
	add c, 0x30
	add d, 0x40
    adc a, 0x10
	adc b, 0x20
	adc c, 0x30
	adc d, 0x40
    sub a, 0x10
	sub b, 0x20
	sub c, 0x30
	sub d, 0x40
    sbb a, 0x10
	sbb b, 0x20
	sbb c, 0x30
	sbb d, 0x40
    cmp a, 0x10
	cmp b, 0x20
	cmp c, 0x30
	cmp d, 0x40
    and a, 0x10
	and b, 0x20
	and c, 0x30
	and d, 0x40
    xor a, 0x10
	xor b, 0x20
	xor c, 0x30
	xor d, 0x40
    or a, 0x10
	or b, 0x20
	or c, 0x30
	or d, 0x40
`,
			Expected: []uint8{
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x10,
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_B, 0x20,
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_C, 0x30,
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_D, 0x40,
				instruction.INST_ADD_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x10,
				instruction.INST_ADD_REG8_IMM8, instruction.REGISTER_OPCODE_B, 0x20,
				instruction.INST_ADD_REG8_IMM8, instruction.REGISTER_OPCODE_C, 0x30,
				instruction.INST_ADD_REG8_IMM8, instruction.REGISTER_OPCODE_D, 0x40,
				instruction.INST_ADC_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x10,
				instruction.INST_ADC_REG8_IMM8, instruction.REGISTER_OPCODE_B, 0x20,
				instruction.INST_ADC_REG8_IMM8, instruction.REGISTER_OPCODE_C, 0x30,
				instruction.INST_ADC_REG8_IMM8, instruction.REGISTER_OPCODE_D, 0x40,
				instruction.INST_SUB_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x10,
				instruction.INST_SUB_REG8_IMM8, instruction.REGISTER_OPCODE_B, 0x20,
				instruction.INST_SUB_REG8_IMM8, instruction.REGISTER_OPCODE_C, 0x30,
				instruction.INST_SUB_REG8_IMM8, instruction.REGISTER_OPCODE_D, 0x40,
				instruction.INST_SBB_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x10,
				instruction.INST_SBB_REG8_IMM8, instruction.REGISTER_OPCODE_B, 0x20,
				instruction.INST_SBB_REG8_IMM8, instruction.REGISTER_OPCODE_C, 0x30,
				instruction.INST_SBB_REG8_IMM8, instruction.REGISTER_OPCODE_D, 0x40,
				instruction.INST_CMP_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x10,
				instruction.INST_CMP_REG8_IMM8, instruction.REGISTER_OPCODE_B, 0x20,
				instruction.INST_CMP_REG8_IMM8, instruction.REGISTER_OPCODE_C, 0x30,
				instruction.INST_CMP_REG8_IMM8, instruction.REGISTER_OPCODE_D, 0x40,
				instruction.INST_AND_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x10,
				instruction.INST_AND_REG8_IMM8, instruction.REGISTER_OPCODE_B, 0x20,
				instruction.INST_AND_REG8_IMM8, instruction.REGISTER_OPCODE_C, 0x30,
				instruction.INST_AND_REG8_IMM8, instruction.REGISTER_OPCODE_D, 0x40,
				instruction.INST_XOR_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x10,
				instruction.INST_XOR_REG8_IMM8, instruction.REGISTER_OPCODE_B, 0x20,
				instruction.INST_XOR_REG8_IMM8, instruction.REGISTER_OPCODE_C, 0x30,
				instruction.INST_XOR_REG8_IMM8, instruction.REGISTER_OPCODE_D, 0x40,
				instruction.INST_OR_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x10,
				instruction.INST_OR_REG8_IMM8, instruction.REGISTER_OPCODE_B, 0x20,
				instruction.INST_OR_REG8_IMM8, instruction.REGISTER_OPCODE_C, 0x30,
				instruction.INST_OR_REG8_IMM8, instruction.REGISTER_OPCODE_D, 0x40,
			},
		},
		{
			Name: "reg8 imm8 tag",
			Code: `
VALUE_1=0x10
VALUE_2=0x20
VALUE_3=0x30
VALUE_4=0x40

    mov a, VALUE_1
    mov b, VALUE_2
    mov c, VALUE_3
    mov d, VALUE_4
`,
			Expected: []byte{
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x10,
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_B, 0x20,
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_C, 0x30,
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_D, 0x40,
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

	shl b
	shr c
	rol d
	inc a
	dec b
	push c
	pop d

`,
			Expected: []byte{
				instruction.INST_SHL_REG8, instruction.REGISTER_OPCODE_B,
				instruction.INST_SHR_REG8, instruction.REGISTER_OPCODE_C,
				instruction.INST_ROL_REG8, instruction.REGISTER_OPCODE_D,
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
    add a, b
	add b, c
	add c, d
    adc a, b
	adc b, c
	adc c, d
    sub a, b
	sub b, c
	sub c, d
    sbb a, b
	sbb b, c
	sbb c, d
    cmp a, b
	cmp b, c
	cmp c, d
    and a, b
	and b, c
	and c, d
    xor a, b
	xor b, c
	xor c, d
    or a, b
	or b, c
	or c, d
	mov ip, bp
	mov bp, ip

`,
			Expected: []byte{
				instruction.INST_MOV_REG8_REG8, instruction.REGISTER_OPCODE_B<<4 | instruction.REGISTER_OPCODE_A,
				instruction.INST_MOV_REG8_REG8, instruction.REGISTER_OPCODE_C<<4 | instruction.REGISTER_OPCODE_B,
				instruction.INST_MOV_REG8_REG8, instruction.REGISTER_OPCODE_D<<4 | instruction.REGISTER_OPCODE_C,
				instruction.INST_ADD_REG8_REG8, instruction.REGISTER_OPCODE_B<<4 | instruction.REGISTER_OPCODE_A,
				instruction.INST_ADD_REG8_REG8, instruction.REGISTER_OPCODE_C<<4 | instruction.REGISTER_OPCODE_B,
				instruction.INST_ADD_REG8_REG8, instruction.REGISTER_OPCODE_D<<4 | instruction.REGISTER_OPCODE_C,
				instruction.INST_ADC_REG8_REG8, instruction.REGISTER_OPCODE_B<<4 | instruction.REGISTER_OPCODE_A,
				instruction.INST_ADC_REG8_REG8, instruction.REGISTER_OPCODE_C<<4 | instruction.REGISTER_OPCODE_B,
				instruction.INST_ADC_REG8_REG8, instruction.REGISTER_OPCODE_D<<4 | instruction.REGISTER_OPCODE_C,
				instruction.INST_SUB_REG8_REG8, instruction.REGISTER_OPCODE_B<<4 | instruction.REGISTER_OPCODE_A,
				instruction.INST_SUB_REG8_REG8, instruction.REGISTER_OPCODE_C<<4 | instruction.REGISTER_OPCODE_B,
				instruction.INST_SUB_REG8_REG8, instruction.REGISTER_OPCODE_D<<4 | instruction.REGISTER_OPCODE_C,
				instruction.INST_SBB_REG8_REG8, instruction.REGISTER_OPCODE_B<<4 | instruction.REGISTER_OPCODE_A,
				instruction.INST_SBB_REG8_REG8, instruction.REGISTER_OPCODE_C<<4 | instruction.REGISTER_OPCODE_B,
				instruction.INST_SBB_REG8_REG8, instruction.REGISTER_OPCODE_D<<4 | instruction.REGISTER_OPCODE_C,
				instruction.INST_CMP_REG8_REG8, instruction.REGISTER_OPCODE_B<<4 | instruction.REGISTER_OPCODE_A,
				instruction.INST_CMP_REG8_REG8, instruction.REGISTER_OPCODE_C<<4 | instruction.REGISTER_OPCODE_B,
				instruction.INST_CMP_REG8_REG8, instruction.REGISTER_OPCODE_D<<4 | instruction.REGISTER_OPCODE_C,
				instruction.INST_AND_REG8_REG8, instruction.REGISTER_OPCODE_B<<4 | instruction.REGISTER_OPCODE_A,
				instruction.INST_AND_REG8_REG8, instruction.REGISTER_OPCODE_C<<4 | instruction.REGISTER_OPCODE_B,
				instruction.INST_AND_REG8_REG8, instruction.REGISTER_OPCODE_D<<4 | instruction.REGISTER_OPCODE_C,
				instruction.INST_XOR_REG8_REG8, instruction.REGISTER_OPCODE_B<<4 | instruction.REGISTER_OPCODE_A,
				instruction.INST_XOR_REG8_REG8, instruction.REGISTER_OPCODE_C<<4 | instruction.REGISTER_OPCODE_B,
				instruction.INST_XOR_REG8_REG8, instruction.REGISTER_OPCODE_D<<4 | instruction.REGISTER_OPCODE_C,
				instruction.INST_OR_REG8_REG8, instruction.REGISTER_OPCODE_B<<4 | instruction.REGISTER_OPCODE_A,
				instruction.INST_OR_REG8_REG8, instruction.REGISTER_OPCODE_C<<4 | instruction.REGISTER_OPCODE_B,
				instruction.INST_OR_REG8_REG8, instruction.REGISTER_OPCODE_D<<4 | instruction.REGISTER_OPCODE_C,
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
	hlt
	nop

`,
			Expected: []byte{
				instruction.INST_RET_IMPL,
				instruction.INST_RTI_IMPL,
				instruction.INST_CLC_IMPL,
				instruction.INST_HLT_IMPL,
				instruction.INST_NOP_IMPL,
			},
		},
		{
			Name: "reg16 imm8",
			Code: `

	add ip, 0x10
	add bp, 0x20
	add sp, 0x30
	sub ip, 0x40
	sub bp, 0x50
	sub sp, 0x60

`,
			Expected: []byte{
				instruction.INST_ADD_REG16_IMM8, instruction.REGISTER_OPCODE_IP, 0x10,
				instruction.INST_ADD_REG16_IMM8, instruction.REGISTER_OPCODE_BP, 0x20,
				instruction.INST_ADD_REG16_IMM8, instruction.REGISTER_OPCODE_SP, 0x30,
				instruction.INST_SUB_REG16_IMM8, instruction.REGISTER_OPCODE_IP, 0x40,
				instruction.INST_SUB_REG16_IMM8, instruction.REGISTER_OPCODE_BP, 0x50,
				instruction.INST_SUB_REG16_IMM8, instruction.REGISTER_OPCODE_SP, 0x60,
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
				instruction.INST_PUSH_REG16, instruction.REGISTER_OPCODE_SP,
				instruction.INST_POP_REG16, instruction.REGISTER_OPCODE_SP,
			},
		},
		{
			Name: "stack imm8",
			Code: `

	add sp, 0x10
	sub sp, 0x10


`,
			Expected: []byte{
				instruction.INST_ADD_REG16_IMM8, instruction.REGISTER_OPCODE_SP, 0x10,
				instruction.INST_SUB_REG16_IMM8, instruction.REGISTER_OPCODE_SP, 0x10,
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
				instruction.INST_ADD_REG16_IMM8, instruction.REGISTER_OPCODE_SP, 0x10,
				instruction.INST_SUB_REG16_IMM8, instruction.REGISTER_OPCODE_SP, 0x10,
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
				instruction.INST_MOV_REG16_REG16, instruction.REGISTER_OPCODE_SP | instruction.REGISTER_OPCODE_IP<<4,
				instruction.INST_MOV_REG16_REG16, instruction.REGISTER_OPCODE_SP | instruction.REGISTER_OPCODE_BP<<4,
				instruction.INST_MOV_REG16_REG16, instruction.REGISTER_OPCODE_IP | instruction.REGISTER_OPCODE_SP<<4,
				instruction.INST_MOV_REG16_REG16, instruction.REGISTER_OPCODE_BP | instruction.REGISTER_OPCODE_SP<<4,
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
				instruction.INST_MOV_MEM_TO_REG_DIRECT, instruction.REGISTER_OPCODE_A, 0x34, 0x12,
				instruction.INST_MOV_REG_TO_MEM_DIRECT, instruction.REGISTER_OPCODE_B, 0x34, 0x12,
				instruction.INST_MOV_MEM_TO_REG_DIRECT, instruction.REGISTER_OPCODE_A, 0xcd, 0xab,
				instruction.INST_MOV_REG_TO_MEM_DIRECT, instruction.REGISTER_OPCODE_B, 0xcd, 0xab,
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
				instruction.INST_MOV_MEM_TO_REG_INDEXED, instruction.REGISTER_OPCODE_A<<4 | instruction.REGISTER_OPCODE_B, 0x34, 0x12,
				instruction.INST_MOV_REG_TO_MEM_INDEXED, instruction.REGISTER_OPCODE_B<<4 | instruction.REGISTER_OPCODE_A, 0x34, 0x12,
				instruction.INST_MOV_MEM_TO_REG_INDEXED, instruction.REGISTER_OPCODE_A<<4 | instruction.REGISTER_OPCODE_B, 0xcd, 0xab,
				instruction.INST_MOV_REG_TO_MEM_INDEXED, instruction.REGISTER_OPCODE_B<<4 | instruction.REGISTER_OPCODE_A, 0xcd, 0xab,
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
				instruction.INST_MOV_MEM_TO_REG_INDIRECT, instruction.REGISTER_OPCODE_A<<4 | instruction.REGISTER_OPCODE_IP,
				instruction.INST_MOV_REG_TO_MEM_INDIRECT, instruction.REGISTER_OPCODE_IP<<4 | instruction.REGISTER_OPCODE_B,
				instruction.INST_MOV_MEM_TO_REG_INDIRECT, instruction.REGISTER_OPCODE_A<<4 | instruction.REGISTER_OPCODE_BP,
				instruction.INST_MOV_REG_TO_MEM_INDIRECT, instruction.REGISTER_OPCODE_BP<<4 | instruction.REGISTER_OPCODE_B,
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
				instruction.INST_MOV_MEM_TO_REG_INDIRECT_OFFSET, instruction.REGISTER_OPCODE_A<<4 | instruction.REGISTER_OPCODE_IP, 0x04,
				instruction.INST_MOV_REG_TO_MEM_INDIRECT_OFFSET, instruction.REGISTER_OPCODE_IP<<4 | instruction.REGISTER_OPCODE_B, 0x08,
				instruction.INST_MOV_MEM_TO_REG_INDIRECT_OFFSET, instruction.REGISTER_OPCODE_A<<4 | instruction.REGISTER_OPCODE_BP, 0x04,
				instruction.INST_MOV_REG_TO_MEM_INDIRECT_OFFSET, instruction.REGISTER_OPCODE_BP<<4 | instruction.REGISTER_OPCODE_B, 0x08,
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
				instruction.INST_MOV_MEM_TO_REG_REG16_INDEXED, instruction.REGISTER_OPCODE_IP<<4 | instruction.REGISTER_OPCODE_C, instruction.REGISTER_OPCODE_A,
				instruction.INST_MOV_REG_TO_MEM_REG16_INDEXED, instruction.REGISTER_OPCODE_IP<<4 | instruction.REGISTER_OPCODE_D, instruction.REGISTER_OPCODE_B,
				instruction.INST_MOV_MEM_TO_REG_REG16_INDEXED, instruction.REGISTER_OPCODE_BP<<4 | instruction.REGISTER_OPCODE_C, instruction.REGISTER_OPCODE_A,
				instruction.INST_MOV_REG_TO_MEM_REG16_INDEXED, instruction.REGISTER_OPCODE_BP<<4 | instruction.REGISTER_OPCODE_D, instruction.REGISTER_OPCODE_B,
			},
		},
		{
			Name: "status register",
			Code: `

	push sr
	pop sr

`,
			Expected: []byte{
				instruction.INST_PUSH_SR,
				instruction.INST_POP_SR,
			},
		},
		{
			Name: "index pointer low-high register",
			Code: `

    mov a, 0xff
    mov [0x3010], a
    mov ipl, 0x10
    mov iph, 0x30
    mov b, [ip]

`,
			Expected: []byte{
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0xff,
				instruction.INST_MOV_REG_TO_MEM_DIRECT, instruction.REGISTER_OPCODE_A, 0x10, 0x30,
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_IPL, 0x10,
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_IPH, 0x30,
				instruction.INST_MOV_MEM_TO_REG_INDIRECT, instruction.REGISTER_OPCODE_B<<4 | instruction.REGISTER_OPCODE_IP,
			},
		},
		{
			Name: "directive .byte",
			Code: `
.byte 0x41, 0x42, 0x43

`,
			Expected: []byte{
				0x41, 0x42, 0x43,
			},
		},
		{
			Name: "directive .word",
			Code: `
.word 0x1000, 0x2000

`,
			Expected: []byte{
				0x00, 0x10, 0x00, 0x20,
			},
		},
		{
			Name: "directive .resb",
			Code: `
.resb 3

start:
    mov a, 0x10

`,
			Expected: []byte{
				0x00, 0x00, 0x00,
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x10,
			},
		},
		{
			Name: "directive .asciiz",
			Code: `
.asciiz "Hi"

`,
			Expected: []byte{
				'H', 'i', 0x00,
			},
		},
		{
			Name: "directive .ascii",
			Code: `
.ascii "OK"

`,
			Expected: []byte{
				'O', 'K',
			},
		},
		{
			Name: "variable simple",
			Code: `
CONST=0x42

start:
    mov a, CONST

`,
			Expected: []byte{
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x42,
			},
		},
		{
			Name: "label reference in jump",
			Code: `
start:
    mov a, 0x05
    cmp a, 0x00
    jz done
    mov a, 0x01

done:
    hlt

`,
			Expected: []byte{
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x05,
				instruction.INST_CMP_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x00,
				instruction.INST_JZ_IMM16, 0x0c, 0x00,
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x01,
				instruction.INST_HLT_IMPL,
			},
		},
		{
			Name: "label with code sections",
			Code: `
code_start:
    mov a, 0x11

code_end:
    mov a, 0x22

`,
			Expected: []byte{
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x11,
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x22,
			},
		},
		{
			Name: "directives with variables",
			Code: `
INIT=0xAA

.resb 2
.byte INIT, INIT

start:
    mov a, INIT

`,
			Expected: []byte{
				0x00, 0x00,
				0xAA, 0xAA,
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0xAA,
			},
		},
		{
			Name: "directive in middle of code",
			Code: `
start:
    mov a, 0x10
    mov b, 0x20

.byte 0x99, 0x88

end:
    hlt

`,
			Expected: []byte{
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x10,
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_B, 0x20,
				0x99, 0x88,
				instruction.INST_HLT_IMPL,
			},
		},
		{
			Name: "resb at end of code block",
			Code: `
start:
    mov a, 0x01
    mov b, 0x02

.resb 4

`,
			Expected: []byte{
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x01,
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_B, 0x02,
				0x00, 0x00, 0x00, 0x00,
			},
		},
		{
			Name: "multiple directives interleaved with code",
			Code: `
CODE=0x01
DATA=0xAA

start:
    mov a, CODE

.byte DATA

    mov b, CODE

.word 0x1234

end:
    hlt

`,
			Expected: []byte{
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x01,
				0xAA,
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_B, 0x01,
				0x34, 0x12,
				instruction.INST_HLT_IMPL,
			},
		},
		{
			Name: "directive after label",
			Code: `
data_section:
.byte 0x11, 0x22, 0x33

code_section:
    mov a, 0x44

`,
			Expected: []byte{
				0x11, 0x22, 0x33,
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x44,
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
