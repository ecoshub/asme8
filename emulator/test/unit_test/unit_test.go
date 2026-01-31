package unit_test

import (
	"asme8/common/instruction"
	"asme8/emulator/src/status"
	"asme8/emulator/test"
	"testing"
)

var (
	TEST_STACK_START uint16 = 0x20ff
)

var (
	Tests = []*test.TestCase{
		{
			Name: "mov immediate reg a",
			// mov a, 0x10
			Program: []uint8{instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x10},
			Expect: &test.Expect{
				Registers: []*test.ExpectRegister{
					{Index: instruction.REGISTER_OPCODE_A, Data: 0x10},
					{Index: instruction.REGISTER_OPCODE_B, Data: 0x00},
					{Index: instruction.REGISTER_OPCODE_C, Data: 0x00},
					{Index: instruction.REGISTER_OPCODE_D, Data: 0x00},
				},
			},
		},
		{
			Name: "mov immediate reg a",
			// mov a, 0xff
			Program: []uint8{instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0xff},
			Expect: &test.Expect{
				Registers: []*test.ExpectRegister{
					{Index: instruction.REGISTER_OPCODE_A, Data: 0xff},
					{Index: instruction.REGISTER_OPCODE_B, Data: 0x00},
					{Index: instruction.REGISTER_OPCODE_C, Data: 0x00},
					{Index: instruction.REGISTER_OPCODE_D, Data: 0x00},
				},
			},
		},
		{
			// mov a, 0xff
			// mov b, a
			Name: "mov reg reg",
			Program: []uint8{
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x10,
				instruction.INST_MOV_REG8_REG8, instruction.REGISTER_OPCODE_A<<1 | instruction.REGISTER_OPCODE_B,
			},
			Expect: &test.Expect{
				Registers: []*test.ExpectRegister{
					{Index: instruction.REGISTER_OPCODE_A, Data: 0x10},
					{Index: instruction.REGISTER_OPCODE_B, Data: 0x10},
					{Index: instruction.REGISTER_OPCODE_C, Data: 0x00},
					{Index: instruction.REGISTER_OPCODE_D, Data: 0x00},
				},
			},
		},
		{
			// mov a, 0xff
			// mov b, a
			// mov c, a
			// mov d, a
			Name: "mov reg reg",
			Program: []uint8{
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0xff,
				instruction.INST_MOV_REG8_REG8, instruction.REGISTER_OPCODE_A<<4 | instruction.REGISTER_OPCODE_B,
				instruction.INST_MOV_REG8_REG8, instruction.REGISTER_OPCODE_A<<4 | instruction.REGISTER_OPCODE_C,
				instruction.INST_MOV_REG8_REG8, instruction.REGISTER_OPCODE_A<<4 | instruction.REGISTER_OPCODE_D,
			},
			Expect: &test.Expect{
				Registers: []*test.ExpectRegister{
					{Index: instruction.REGISTER_OPCODE_A, Data: 0xff},
					{Index: instruction.REGISTER_OPCODE_B, Data: 0xff},
					{Index: instruction.REGISTER_OPCODE_C, Data: 0xff},
					{Index: instruction.REGISTER_OPCODE_D, Data: 0xff},
				},
			},
		},
		{
			// mov a, 0xff
			// mov [0x9000], a
			Name: "mov reg data",
			Program: []uint8{
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0xff,
				instruction.INST_MOV_REG_TO_MEM_DIRECT, instruction.REGISTER_OPCODE_A, 0x00, 0x90,
			},
			Expect: &test.Expect{
				Registers: []*test.ExpectRegister{
					{Index: instruction.REGISTER_OPCODE_A, Data: 0xff},
				},
				Data: []*test.ExpectData{
					{Type: test.DEV_TYPE_RAM, Addr: 0x9000, Data: 0xff},
				},
			},
		},
		{
			// mov a, 0x90
			// mov [0x9000], a
			// mov a, 0x91
			// mov [0x9001], a
			// mov a, 0x92
			// mov [0x9002], a
			Name: "mov reg mem",
			Program: []uint8{
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x90,
				instruction.INST_MOV_REG_TO_MEM_DIRECT, instruction.REGISTER_OPCODE_A, 0x00, 0x90,
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x91,
				instruction.INST_MOV_REG_TO_MEM_DIRECT, instruction.REGISTER_OPCODE_A, 0x01, 0x90,
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x92,
				instruction.INST_MOV_REG_TO_MEM_DIRECT, instruction.REGISTER_OPCODE_A, 0x02, 0x90,
			},
			Expect: &test.Expect{
				Data: []*test.ExpectData{
					{Type: test.DEV_TYPE_RAM, Addr: 0x9000, Data: 0x90},
					{Type: test.DEV_TYPE_RAM, Addr: 0x9001, Data: 0x91},
					{Type: test.DEV_TYPE_RAM, Addr: 0x9002, Data: 0x92},
				},
			},
		},
		{
			// mov a, 0x90
			// mov [0x9000], a
			// mov b, [0x9000]
			Name: "mov reg mem",
			Program: []uint8{
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x90,
				instruction.INST_MOV_REG_TO_MEM_DIRECT, instruction.REGISTER_OPCODE_A, 0x00, 0x90,
				instruction.INST_MOV_MEM_TO_REG_DIRECT, instruction.REGISTER_OPCODE_B, 0x00, 0x90,
			},
			Expect: &test.Expect{
				Registers: []*test.ExpectRegister{
					{Index: instruction.REGISTER_OPCODE_B, Data: 0x90},
				},
				Data: []*test.ExpectData{
					{Type: test.DEV_TYPE_RAM, Addr: 0x9000, Data: 0x90},
				},
			},
		},
		{
			// mov a, 0x10
			// mov b, 0x2
			// mov [0x9002], a
			// mov c, [0x9000+b]
			Name: "mov mem reg data offset",
			Program: []uint8{
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x10,
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_B, 0x2,
				instruction.INST_MOV_REG_TO_MEM_DIRECT, instruction.REGISTER_OPCODE_A, 0x02, 0x90,
				instruction.INST_MOV_MEM_TO_REG_INDEXED, instruction.REGISTER_OPCODE_C<<4 | instruction.REGISTER_OPCODE_B, 0x00, 0x90,
			},
			Expect: &test.Expect{
				Registers: []*test.ExpectRegister{
					{Index: instruction.REGISTER_OPCODE_A, Data: 0x10},
					{Index: instruction.REGISTER_OPCODE_B, Data: 0x2},
					{Index: instruction.REGISTER_OPCODE_C, Data: 0x10},
				},
				Data: []*test.ExpectData{
					{Type: test.DEV_TYPE_RAM, Addr: 0x9002, Data: 0x10},
				},
			},
		},
		{
			// mov a, 0x10
			// mov b, 0x2
			// mov [0x9101], a
			// mov c, [0x90ff+b]
			Name: "mov mem reg data offset",
			Program: []uint8{
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x10,
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_B, 0x2,
				instruction.INST_MOV_REG_TO_MEM_DIRECT, instruction.REGISTER_OPCODE_A, 0x01, 0x91,
				instruction.INST_MOV_MEM_TO_REG_INDEXED, instruction.REGISTER_OPCODE_C<<4 | instruction.REGISTER_OPCODE_B, 0xff, 0x90,
			},
			Expect: &test.Expect{
				Registers: []*test.ExpectRegister{
					{Index: instruction.REGISTER_OPCODE_A, Data: 0x10},
					{Index: instruction.REGISTER_OPCODE_B, Data: 0x2},
					{Index: instruction.REGISTER_OPCODE_C, Data: 0x10},
				},
				Data: []*test.ExpectData{
					{Type: test.DEV_TYPE_RAM, Addr: 0x9101, Data: 0x10},
				},
			},
		},
		{
			// mov a, 0x10
			// mov b, 0x2
			// mov [0x90ff+b], a
			// mov c, [0x90ff+b]
			Name: "mov reg mem data offset",
			Program: []uint8{
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x10,
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_B, 0x2,
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_C, 0x0,
				instruction.INST_MOV_REG_TO_MEM_INDEXED, instruction.REGISTER_OPCODE_A<<4 | instruction.REGISTER_OPCODE_B, 0xff, 0x90,
				instruction.INST_MOV_MEM_TO_REG_INDEXED, instruction.REGISTER_OPCODE_C<<4 | instruction.REGISTER_OPCODE_B, 0xff, 0x90,
			},
			Expect: &test.Expect{
				Registers: []*test.ExpectRegister{
					{Index: instruction.REGISTER_OPCODE_A, Data: 0x10},
					{Index: instruction.REGISTER_OPCODE_B, Data: 0x2},
					{Index: instruction.REGISTER_OPCODE_C, Data: 0x10},
				},
				Data: []*test.ExpectData{
					{Type: test.DEV_TYPE_RAM, Addr: 0x9101, Data: 0x10},
				},
			},
		},
		{
			// mov a, 0x10
			// brk
			// mov a, 0x20
			Name: "break",
			Program: []uint8{
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x10,
				instruction.INST_BRK_IMPL,
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x20,
			},
			Expect: &test.Expect{
				Registers: []*test.ExpectRegister{
					{Index: instruction.REGISTER_OPCODE_A, Data: 0x10},
				},
			},
		},
		{
			//     mov a, 0x10
			//     jmp here
			//     mov a, 0x20
			//     brk
			// here:
			//     mov a, 0x30
			Name: "jump imm",
			Program: []uint8{
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x10,
				instruction.INST_JMP_IMPL_IMM16, 0x0a, 0x00,
				instruction.INST_BRK_IMPL,
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x20,
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x30,
			},
			Expect: &test.Expect{
				Registers: []*test.ExpectRegister{
					{Index: instruction.REGISTER_OPCODE_A, Data: 0x30},
				},
			},
		},
		{
			//    mov a, 0x10
			//    mov ipl, 0x0f
			//    mov iph, 0x00
			//    jmp ip
			//    mov a, 0x20
			//    brk
			// here:
			//    mov a, 0x30
			Name: "jump reg16",
			Program: []uint8{
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x10,
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_IPL, 0x0f,
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_IPH, 0x00,
				instruction.INST_JMP_REG16, instruction.REGISTER_OPCODE_IP,
				instruction.INST_BRK_IMPL,
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x20,
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x30,
			},
			Expect: &test.Expect{
				Registers: []*test.ExpectRegister{
					{Index: instruction.REGISTER_OPCODE_A, Data: 0x30},
				},
			},
		},
		{
			// start:
			//     mov a, 0x10
			//     cmp a, 0x10
			//     jz here
			//     mov a, 0x20
			//     brk
			// here:
			//     mov a, 0x30
			Name: "cmp and jz",
			Program: []uint8{
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x10,
				instruction.INST_CMP_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x10,
				instruction.INST_JZ_IMPL_IMM16, 0x0d, 0x00,
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x20,
				instruction.INST_BRK_IMPL,
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x30,
			},
			Expect: &test.Expect{
				Registers: []*test.ExpectRegister{
					{Index: instruction.REGISTER_OPCODE_A, Data: 0x30},
				},
			},
		},
		{
			// start:
			//     mov a, 0x10
			//     cmp a, 0x10
			//     jz here
			//     mov a, 0x20
			//     brk
			// here:
			//     mov a, 0x30
			Name: "cmp and jz no jump",
			Program: []uint8{
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x10,
				instruction.INST_CMP_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x20,
				instruction.INST_JZ_IMPL_IMM16, 0x0d, 0x00,
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x20,
				instruction.INST_BRK_IMPL,
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x30,
			},
			Expect: &test.Expect{
				Registers: []*test.ExpectRegister{
					{Index: instruction.REGISTER_OPCODE_A, Data: 0x20},
				},
			},
		},
		{
			// start:
			//     mov a, 0x10
			// 	   cmp a, 0x10
			//     jnz here
			//     mov a, 0x20
			//     brk
			// here:
			//     mov a, 0x30
			Name: "jnz jump",
			Program: []uint8{
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x10,
				instruction.INST_CMP_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x10,
				instruction.INST_JNZ_IMPL_IMM16, 0x0d, 0x00,
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x20,
				instruction.INST_BRK_IMPL,
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x30,
			},
			Expect: &test.Expect{
				Registers: []*test.ExpectRegister{
					{Index: instruction.REGISTER_OPCODE_A, Data: 0x20},
				},
			},
		},
		{
			// start:
			//     mov a, 0x20
			// 	   cmp a, 0x10
			//     jnz here
			//     mov a, 0x20
			//     brk
			// here:
			//     mov a, 0x30
			Name: "jnz no jump",
			Program: []uint8{
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x20,
				instruction.INST_CMP_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x10,
				instruction.INST_JNZ_IMPL_IMM16, 0x0d, 0x00,
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x20,
				instruction.INST_BRK_IMPL,
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x30,
			},
			Expect: &test.Expect{
				Registers: []*test.ExpectRegister{
					{Index: instruction.REGISTER_OPCODE_A, Data: 0x30},
				},
			},
		},
		{
			// start:
			//     mov a, 0x10
			//     mov b, 0x10
			//     cmp a, b
			//     jz here
			//     mov a, 0x20
			//     brk
			// here:
			//     mov a, 0x30
			Name: "cmp rr and jz",
			Program: []uint8{
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x10,
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_B, 0x10,
				instruction.INST_CMP_REG8_REG8, instruction.REGISTER_OPCODE_A<<4 | instruction.REGISTER_OPCODE_B,
				instruction.INST_JZ_IMPL_IMM16, 0x0f, 0x00,
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x20,
				instruction.INST_BRK_IMPL,
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x30,
			},
			Expect: &test.Expect{
				Registers: []*test.ExpectRegister{
					{Index: instruction.REGISTER_OPCODE_A, Data: 0x30},
				},
			},
		},
		{
			// start:
			//     mov a, 0x10
			//     mov b, 0x20
			//     cmp a, b
			//     jz here
			//     mov a, 0x20
			//     brk
			// here:
			//     mov a, 0x30
			Name: "cmp rr and jz no jump",
			Program: []uint8{
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x10,
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_B, 0x20,
				instruction.INST_CMP_REG8_REG8, instruction.REGISTER_OPCODE_A<<4 | instruction.REGISTER_OPCODE_B,
				instruction.INST_JZ_IMPL_IMM16, 0x0f, 0x00,
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x20,
				instruction.INST_BRK_IMPL,
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x30,
			},
			Expect: &test.Expect{
				Registers: []*test.ExpectRegister{
					{Index: instruction.REGISTER_OPCODE_A, Data: 0x20},
				},
			},
		},
		{
			Name: "add reg imm",
			// add a, 0x10
			Program: []uint8{instruction.INST_ADD_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x10},
			Expect: &test.Expect{
				Registers: []*test.ExpectRegister{
					{Index: instruction.REGISTER_OPCODE_A, Data: 0x10},
				},
			},
		},
		{
			Name: "adc reg imm",
			// mov a, 0xff
			// mov b, 0x10
			// add a, b
			// adc c, 0x10
			Program: []uint8{
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0xff,
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_B, 0x10,
				instruction.INST_ADD_REG8_REG8, instruction.REGISTER_OPCODE_B<<4 | instruction.REGISTER_OPCODE_A,
				instruction.INST_ADC_REG8_IMM8, instruction.REGISTER_OPCODE_C, 0x10,
			},
			Expect: &test.Expect{
				Registers: []*test.ExpectRegister{
					{Index: instruction.REGISTER_OPCODE_C, Data: 0x11},
				},
			},
		},
		{
			Name: "add reg imm",
			// mov a, 0x20
			// add a, 0x10
			Program: []uint8{
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x20,
				instruction.INST_ADD_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x10,
			},
			Expect: &test.Expect{
				Registers: []*test.ExpectRegister{
					{Index: instruction.REGISTER_OPCODE_A, Data: 0x30},
				},
			},
		},
		{
			Name: "add reg imm",
			// mov a, 0xff
			// add a, 0x01
			Program: []uint8{
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0xff,
				instruction.INST_ADD_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x01,
			},
			Expect: &test.Expect{
				Registers: []*test.ExpectRegister{
					{Index: instruction.REGISTER_OPCODE_A, Data: 0x0},
				},
			},
		},
		{
			Name: "add reg reg",
			// mov a, 0xff
			// mov b, 0x11
			// add a, b
			Program: []uint8{
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0xff,
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_B, 0x11,
				instruction.INST_ADD_REG8_REG8, instruction.REGISTER_OPCODE_B<<4 | instruction.REGISTER_OPCODE_A,
			},
			Expect: &test.Expect{
				Registers: []*test.ExpectRegister{
					{Index: instruction.REGISTER_OPCODE_A, Data: 0x10},
					{Index: instruction.REGISTER_OPCODE_B, Data: 0x11},
				},
			},
		},
		{
			Name: "add reg reg",
			// mov a, 0xff
			// mov b, 0x11
			// add b, a
			Program: []uint8{
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0xff,
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_B, 0x11,
				instruction.INST_ADD_REG8_REG8, instruction.REGISTER_OPCODE_A<<4 | instruction.REGISTER_OPCODE_B,
			},
			Expect: &test.Expect{
				Registers: []*test.ExpectRegister{
					{Index: instruction.REGISTER_OPCODE_A, Data: 0xff},
					{Index: instruction.REGISTER_OPCODE_B, Data: 0x10},
				},
			},
		},
		{
			Name: "sub reg imm",
			// sub a, 0x10
			Program: []uint8{instruction.INST_SUB_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x10},
			Expect: &test.Expect{
				Registers: []*test.ExpectRegister{
					{Index: instruction.REGISTER_OPCODE_A, Data: 0xf0},
				},
			},
		},
		{
			Name: "sub reg imm",
			// mov a, 0x20
			// sub a, 0x10
			Program: []uint8{
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x20,
				instruction.INST_SUB_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x10,
			},
			Expect: &test.Expect{
				Registers: []*test.ExpectRegister{
					{Index: instruction.REGISTER_OPCODE_A, Data: 0x10},
				},
			},
		},
		{
			Name: "sub reg imm",
			// mov a, 0xff
			// sub a, 0x01
			Program: []uint8{
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0xff,
				instruction.INST_SUB_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x01,
			},
			Expect: &test.Expect{
				Registers: []*test.ExpectRegister{
					{Index: instruction.REGISTER_OPCODE_A, Data: 0xfe},
				},
			},
		},
		{
			Name: "sub reg reg",
			// mov a, 0xff
			// mov b, 0x11
			// sub a, b
			Program: []uint8{
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0xff,
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_B, 0x11,
				instruction.INST_SUB_REG8_REG8, instruction.REGISTER_OPCODE_B<<4 | instruction.REGISTER_OPCODE_A,
			},
			Expect: &test.Expect{
				Registers: []*test.ExpectRegister{
					{Index: instruction.REGISTER_OPCODE_A, Data: 0xee},
					{Index: instruction.REGISTER_OPCODE_B, Data: 0x11},
				},
			},
		},
		{
			Name: "sub reg reg",
			// mov a, 0xff
			// mov b, 0x11
			// sub b, a
			Program: []uint8{
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0xff,
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_B, 0x11,
				instruction.INST_SUB_REG8_REG8, instruction.REGISTER_OPCODE_A<<4 | instruction.REGISTER_OPCODE_B,
			},
			Expect: &test.Expect{
				Registers: []*test.ExpectRegister{
					{Index: instruction.REGISTER_OPCODE_A, Data: 0xff},
					{Index: instruction.REGISTER_OPCODE_B, Data: 0x12},
				},
			},
		},
		{
			Name: "push pop reg",
			// mov a, 0x10
			// push a
			// pop b
			Program: []uint8{
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x10,
				instruction.INST_PUSH_REG8, instruction.REGISTER_OPCODE_A,
				instruction.INST_POP_REG8, instruction.REGISTER_OPCODE_B,
			},
			Expect: &test.Expect{
				Registers: []*test.ExpectRegister{
					{Index: instruction.REGISTER_OPCODE_A, Data: 0x10},
					{Index: instruction.REGISTER_OPCODE_B, Data: 0x10},
				},
			},
		},
		{
			Name: "inc reg",
			// mov a, 0x20
			// inc a
			Program: []uint8{
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x20,
				instruction.INST_INC_REG8, instruction.REGISTER_OPCODE_A,
			},
			Expect: &test.Expect{
				Registers: []*test.ExpectRegister{
					{Index: instruction.REGISTER_OPCODE_A, Data: 0x21},
				},
			},
		},
		{
			Name: "inc reg",
			// mov a, 0xff
			// inc a
			Program: []uint8{
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0xff,
				instruction.INST_INC_REG8, instruction.REGISTER_OPCODE_A,
			},
			Expect: &test.Expect{
				Registers: []*test.ExpectRegister{
					{Index: instruction.REGISTER_OPCODE_A, Data: 0x00},
				},
			},
		},
		{
			Name: "dec reg",
			// mov a, 0x20
			// dec a
			Program: []uint8{
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x20,
				instruction.INST_DEC_REG8, instruction.REGISTER_OPCODE_A,
			},
			Expect: &test.Expect{
				Registers: []*test.ExpectRegister{
					{Index: instruction.REGISTER_OPCODE_A, Data: 0x1f},
				},
			},
		},
		{
			Name: "dec reg",
			// mov a, 0x00
			// dec a
			Program: []uint8{
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x00,
				instruction.INST_DEC_REG8, instruction.REGISTER_OPCODE_A,
			},
			Expect: &test.Expect{
				Registers: []*test.ExpectRegister{
					{Index: instruction.REGISTER_OPCODE_A, Data: 0xff},
				},
			},
		},
		{
			Name: "call and ret",
			// start:
			//     call hello
			//     mov a, 0x10
			//     brk
			// hello:
			//     mov a, 0x20
			//     mov b, 0x20
			//     ret
			//     brk
			Program: []uint8{
				instruction.INST_CALL_IMM16, 0x07, 0x00,
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x10,
				instruction.INST_BRK_IMPL,
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x20,
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_B, 0x20,
				instruction.INST_RET_IMPL,
				instruction.INST_BRK_IMPL,
			},
			Expect: &test.Expect{
				Registers: []*test.ExpectRegister{
					{Index: instruction.REGISTER_OPCODE_A, Data: 0x10},
					{Index: instruction.REGISTER_OPCODE_B, Data: 0x20},
				},
			},
		},
		{
			Name: "call reg16 and ret",
			// start:
			//     mov ipl, 0x0c
			//     mov iph, 0x00
			//     call ip
			//     mov a, 0x10
			//     brk
			//     mov a, 0x20
			//     mov b, 0x20
			//     ret
			Program: []uint8{
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_IPL, 0x0c,
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_IPH, 0x00,
				instruction.INST_CALL_REG16, instruction.REGISTER_OPCODE_IP,
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x10,
				instruction.INST_BRK_IMPL,
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x20,
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_B, 0x20,
				instruction.INST_RET_IMPL,
			},
			Expect: &test.Expect{
				Registers: []*test.ExpectRegister{
					{Index: instruction.REGISTER_OPCODE_A, Data: 0x10},
					{Index: instruction.REGISTER_OPCODE_B, Data: 0x20},
				},
			},
		},
		{
			Name: "xor reg reg",
			// mov a, 0b00001111
			// mov b, 0b11001100
			// xor b, a
			Program: []uint8{
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0b00001111,
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_B, 0b11001100,
				instruction.INST_XOR_REG8_REG8, instruction.REGISTER_OPCODE_A<<4 | instruction.REGISTER_OPCODE_B,
			},
			Expect: &test.Expect{
				Registers: []*test.ExpectRegister{
					{Index: instruction.REGISTER_OPCODE_A, Data: 0b00001111},
					{Index: instruction.REGISTER_OPCODE_B, Data: 0b11000011},
				},
			},
		},
		{
			Name: "xor reg reg clear",
			// mov a, 0b00001111
			// xor a, a
			Program: []uint8{
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0b00001111,
				instruction.INST_XOR_REG8_REG8, instruction.REGISTER_OPCODE_A<<4 | instruction.REGISTER_OPCODE_A,
			},
			Expect: &test.Expect{
				Registers: []*test.ExpectRegister{
					{Index: instruction.REGISTER_OPCODE_A, Data: 0},
				},
			},
		},
		{
			Name: "xor reg imm",
			// mov a, 0b11001100
			// xor a, 0b00001111
			Program: []uint8{
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0b11001100,
				instruction.INST_XOR_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0b00001111,
			},
			Expect: &test.Expect{
				Registers: []*test.ExpectRegister{
					{Index: instruction.REGISTER_OPCODE_A, Data: 0b11000011},
				},
			},
		},
		{
			Name: "and reg reg",
			// mov a, 0b00001111
			// mov b, 0b11001100
			// and b, a
			Program: []uint8{
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0b00001111,
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_B, 0b11001100,
				instruction.INST_AND_REG8_REG8, instruction.REGISTER_OPCODE_A<<4 | instruction.REGISTER_OPCODE_B,
			},
			Expect: &test.Expect{
				Registers: []*test.ExpectRegister{
					{Index: instruction.REGISTER_OPCODE_A, Data: 0b00001111},
					{Index: instruction.REGISTER_OPCODE_B, Data: 0b00001100},
				},
			},
		},
		{
			Name: "and reg imm",
			// mov a, 0b11001100
			// and a, 0b00001111
			Program: []uint8{
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0b11001100,
				instruction.INST_AND_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0b00001111,
			},
			Expect: &test.Expect{
				Registers: []*test.ExpectRegister{
					{Index: instruction.REGISTER_OPCODE_A, Data: 0b00001100},
				},
			},
		},
		{
			Name: "or reg reg",
			// mov a, 0b00001111
			// mov b, 0b11001100
			// or b, a
			Program: []uint8{
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0b00001111,
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_B, 0b11001100,
				instruction.INST_OR_REG8_REG8, instruction.REGISTER_OPCODE_A<<4 | instruction.REGISTER_OPCODE_B,
			},
			Expect: &test.Expect{
				Registers: []*test.ExpectRegister{
					{Index: instruction.REGISTER_OPCODE_A, Data: 0b00001111},
					{Index: instruction.REGISTER_OPCODE_B, Data: 0b11001111},
				},
			},
		},
		{
			Name: "or reg imm",
			// mov a, 0b11001100
			// or a, 0b00001111
			Program: []uint8{
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0b11001100,
				instruction.INST_OR_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0b00001111,
			},
			Expect: &test.Expect{
				Registers: []*test.ExpectRegister{
					{Index: instruction.REGISTER_OPCODE_A, Data: 0b11001111},
				},
			},
		},
		{
			Name: "not reg",
			// mov a, 0b11001100
			// not a,
			Program: []uint8{
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0b11001100,
				instruction.INST_NOT_REG8, instruction.REGISTER_OPCODE_A,
			},
			Expect: &test.Expect{
				Registers: []*test.ExpectRegister{
					{Index: instruction.REGISTER_OPCODE_A, Data: 0b00110011},
				},
			},
		},
		{
			Name: "shl reg one",
			// mov a, 0b11001100
			// shl a,
			Program: []uint8{
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0b11001100,
				instruction.INST_SHL_REG8, instruction.REGISTER_OPCODE_A,
			},
			Expect: &test.Expect{
				Registers: []*test.ExpectRegister{
					{Index: instruction.REGISTER_OPCODE_A, Data: 0b10011000},
				},
				Status: &test.ExpectStatusData{
					Data: 0 | status.STATUS_FLAG_CARRY,
				},
			},
		},
		{
			Name: "shl reg zero",
			// mov a, 0b00110011
			// shl a,
			Program: []uint8{
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0b00110011,
				instruction.INST_SHL_REG8, instruction.REGISTER_OPCODE_A,
			},
			Expect: &test.Expect{
				Registers: []*test.ExpectRegister{
					{Index: instruction.REGISTER_OPCODE_A, Data: 0b01100110},
				},
				Status: &test.ExpectStatusData{
					Data: 0 | status.STATUS_FLAG_PARITY,
				},
			},
		},
		{
			Name: "shr reg zero",
			// mov a, 0b11001100
			// shr a,
			Program: []uint8{
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0b11001100,
				instruction.INST_SHR_REG8, instruction.REGISTER_OPCODE_A,
			},
			Expect: &test.Expect{
				Registers: []*test.ExpectRegister{
					{Index: instruction.REGISTER_OPCODE_A, Data: 0b01100110},
				},
				Status: &test.ExpectStatusData{
					Data: 0 | status.STATUS_FLAG_PARITY,
				},
			},
		},
		{
			Name: "shr reg one",
			// mov a, 0b00110011
			// shr a,
			Program: []uint8{
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0b00110011,
				instruction.INST_SHR_REG8, instruction.REGISTER_OPCODE_A,
			},
			Expect: &test.Expect{
				Registers: []*test.ExpectRegister{
					{Index: instruction.REGISTER_OPCODE_A, Data: 0b00011001},
				},
				Status: &test.ExpectStatusData{
					Data: 0 | status.STATUS_FLAG_CARRY,
				},
			},
		},
		{
			Name: "ror no carry",
			// mov a, 0b11001100
			// ror a,
			Program: []uint8{
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0b11001100,
				instruction.INST_ROR_REG8, instruction.REGISTER_OPCODE_A,
			},
			Expect: &test.Expect{
				Registers: []*test.ExpectRegister{
					{Index: instruction.REGISTER_OPCODE_A, Data: 0b01100110},
				},
				Status: &test.ExpectStatusData{
					Data: 0 | status.STATUS_FLAG_PARITY,
				},
			},
		},
		{
			Name: "ror zero carry",
			// mov a, 0b00110011
			// ror a,
			Program: []uint8{
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0b00110011,
				instruction.INST_ROR_REG8, instruction.REGISTER_OPCODE_A,
			},
			Expect: &test.Expect{
				Registers: []*test.ExpectRegister{
					{Index: instruction.REGISTER_OPCODE_A, Data: 0b00011001},
				},
				Status: &test.ExpectStatusData{
					Data: 0 | status.STATUS_FLAG_CARRY,
				},
			},
		},
		{
			Name: "rol no carry",
			// mov a, 0b11001100
			// rol a,
			Program: []uint8{
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0b11001100,
				instruction.INST_ROL_REG8, instruction.REGISTER_OPCODE_A,
			},
			Expect: &test.Expect{
				Registers: []*test.ExpectRegister{
					{Index: instruction.REGISTER_OPCODE_A, Data: 0b10011000},
				},
				Status: &test.ExpectStatusData{
					Data: 0 | status.STATUS_FLAG_CARRY,
				},
			},
		},
		{
			Name: "rol zero carry",
			// mov a, 0b00110011
			// rol a,
			Program: []uint8{
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0b00110011,
				instruction.INST_ROL_REG8, instruction.REGISTER_OPCODE_A,
			},
			Expect: &test.Expect{
				Registers: []*test.ExpectRegister{
					{Index: instruction.REGISTER_OPCODE_A, Data: 0b01100110},
				},
				Status: &test.ExpectStatusData{
					Data: 0 | status.STATUS_FLAG_PARITY,
				},
			},
		},
		{
			// start:
			// 	   mov a, 0xff
			// 	   mov b, 0x02
			//     add a, b
			//     jc here
			//     mov c, 0x20
			//     brk
			// here:
			//     mov c, 0x30
			Name: "jc",
			Program: []uint8{
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0xff,
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_B, 0x02,
				instruction.INST_ADD_REG8_REG8, instruction.REGISTER_OPCODE_B<<4 | instruction.REGISTER_OPCODE_A,
				instruction.INST_JC_IMPL_IMM16, 0x0f, 0x00,
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_C, 0x20,
				instruction.INST_BRK_IMPL,
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_C, 0x30,
			},
			Expect: &test.Expect{
				Registers: []*test.ExpectRegister{
					{Index: instruction.REGISTER_OPCODE_A, Data: 0x01},
					{Index: instruction.REGISTER_OPCODE_B, Data: 0x02},
					{Index: instruction.REGISTER_OPCODE_C, Data: 0x30},
				},
			},
		},
		{
			// start:
			// 	   mov a, 0x01
			// 	   mov b, 0x02
			//     add a, b
			//     jc here
			//     mov c, 0x20
			//     brk
			// here:
			//     mov c, 0x30
			Name: "jc no jump",
			Program: []uint8{
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x01,
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_B, 0x02,
				instruction.INST_ADD_REG8_REG8, instruction.REGISTER_OPCODE_B<<4 | instruction.REGISTER_OPCODE_A,
				instruction.INST_JC_IMPL_IMM16, 0x0f, 0x00,
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_C, 0x20,
				instruction.INST_BRK_IMPL,
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_C, 0x30,
			},
			Expect: &test.Expect{
				Registers: []*test.ExpectRegister{
					{Index: instruction.REGISTER_OPCODE_A, Data: 0x03},
					{Index: instruction.REGISTER_OPCODE_B, Data: 0x02},
					{Index: instruction.REGISTER_OPCODE_C, Data: 0x20},
				},
			},
		},
		{
			// start:
			// 	   mov a, 0xff
			// 	   mov b, 0x02
			//     add a, b
			//     clc
			//     jc here
			//     mov c, 0x20
			//     brk
			// here:
			//     mov c, 0x30
			Name: "jc with clear carry",
			Program: []uint8{
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0xff,
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_B, 0x02,
				instruction.INST_ADD_REG8_REG8, instruction.REGISTER_OPCODE_B<<4 | instruction.REGISTER_OPCODE_A,
				instruction.INST_CLC_IMPL,
				instruction.INST_JC_IMPL_IMM16, 0x10, 0x00,
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_C, 0x20,
				instruction.INST_BRK_IMPL,
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_C, 0x30,
			},
			Expect: &test.Expect{
				Registers: []*test.ExpectRegister{
					{Index: instruction.REGISTER_OPCODE_A, Data: 0x01},
					{Index: instruction.REGISTER_OPCODE_B, Data: 0x02},
					{Index: instruction.REGISTER_OPCODE_C, Data: 0x20},
				},
			},
		},
		{
			// start:
			//     mov a, 0x10
			//     cmp a, 0x20
			//     js here
			//     mov a, 0x20
			//     brk
			// here:
			//     mov a, 0x30
			Name: "cmp and js",
			Program: []uint8{
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x10,
				instruction.INST_CMP_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x20,
				instruction.INST_JS_IMPL_IMM16, 0x0d, 0x00,
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x20,
				instruction.INST_BRK_IMPL,
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x30,
			},
			Expect: &test.Expect{
				Registers: []*test.ExpectRegister{
					{Index: instruction.REGISTER_OPCODE_A, Data: 0x30},
				},
			},
		},
		{
			// start:
			//     mov a, 0x20
			//     cmp a, 0x10
			//     js here
			//     mov a, 0x20
			//     brk
			// here:
			//     mov a, 0x30
			Name: "cmp and js no jump",
			Program: []uint8{
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x20,
				instruction.INST_CMP_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x10,
				instruction.INST_JS_IMPL_IMM16, 0x0d, 0x00,
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x20,
				instruction.INST_BRK_IMPL,
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x30,
			},
			Expect: &test.Expect{
				Registers: []*test.ExpectRegister{
					{Index: instruction.REGISTER_OPCODE_A, Data: 0x20},
				},
			},
		},
		{
			Name: "push pop sp",
			// mov a, 0x20
			// push a
			// push sp
			Program: []uint8{
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x20,
				instruction.INST_PUSH_REG8, instruction.REGISTER_OPCODE_A,
				instruction.INST_PUSH_SP,
			},
			Expect: &test.Expect{
				Data: []*test.ExpectData{
					{Type: test.DEV_TYPE_RAM, Addr: TEST_STACK_START - 1, Data: 0xfe},
					{Type: test.DEV_TYPE_RAM, Addr: TEST_STACK_START - 2, Data: 0x20},
				},
			},
		},
		{
			Name: "add sp",
			// add sp, 4
			// mov a, 0x10
			// push a
			Program: []uint8{
				instruction.INST_ADD_SP_IMM8, 0x04,
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x10,
				instruction.INST_PUSH_REG8, instruction.REGISTER_OPCODE_A,
			},
			Expect: &test.Expect{
				Data: []*test.ExpectData{
					{Type: test.DEV_TYPE_RAM, Addr: TEST_STACK_START + 4, Data: 0x10},
				},
			},
		},
		{
			Name: "sub sp",
			// sub sp, 4
			// mov a, 0x10
			// push a
			Program: []uint8{
				instruction.INST_SUB_SP_IMM8, 0x04,
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x10,
				instruction.INST_PUSH_REG8, instruction.REGISTER_OPCODE_A,
			},
			Expect: &test.Expect{
				Data: []*test.ExpectData{
					{Type: test.DEV_TYPE_RAM, Addr: TEST_STACK_START - 4, Data: 0x10},
				},
			},
		},
		{
			Name: "mov ipl iph imm",
			// mov ipl, 0x10
			// mov iph, 0x20
			// push ip
			// mov ipl, 0x11
			// mov iph, 0x21
			// pop ip
			Program: []uint8{
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_IPL, 0x20,
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_IPH, 0x10,
				instruction.INST_PUSH_REG16, instruction.REGISTER_OPCODE_IP,
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_IPL, 0x21,
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_IPH, 0x12,
				instruction.INST_POP_REG16, instruction.REGISTER_OPCODE_IP,
			},
			Expect: &test.Expect{
				Registers: []*test.ExpectRegister{
					{Index: instruction.REGISTER_OPCODE_IPL, Data: 0x20},
					{Index: instruction.REGISTER_OPCODE_IPH, Data: 0x10},
				},
				Data: []*test.ExpectData{
					{Type: test.DEV_TYPE_RAM, Addr: TEST_STACK_START, Data: 0x20},
					{Type: test.DEV_TYPE_RAM, Addr: TEST_STACK_START - 1, Data: 0x10},
				},
			},
		},
		{
			Name: "mov ipl iph reg",
			// mov a, 0x10
			// mov b, 0x20
			// mov ipl, a
			// mov iph, b
			// push ip
			Program: []uint8{
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x10,
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_B, 0x20,
				instruction.INST_MOV_REG8_REG8, instruction.REGISTER_OPCODE_A<<4 | instruction.REGISTER_OPCODE_IPL,
				instruction.INST_MOV_REG8_REG8, instruction.REGISTER_OPCODE_B<<4 | instruction.REGISTER_OPCODE_IPH,
				instruction.INST_PUSH_REG16, instruction.REGISTER_OPCODE_IP,
			},
			Expect: &test.Expect{
				Registers: []*test.ExpectRegister{
					{Index: instruction.REGISTER_OPCODE_A, Data: 0x10},
					{Index: instruction.REGISTER_OPCODE_B, Data: 0x20},
					{Index: instruction.REGISTER_OPCODE_IPL, Data: 0x10},
					{Index: instruction.REGISTER_OPCODE_IPH, Data: 0x20},
				},
				Data: []*test.ExpectData{
					{Type: test.DEV_TYPE_RAM, Addr: TEST_STACK_START, Data: 0x10},
					{Type: test.DEV_TYPE_RAM, Addr: TEST_STACK_START - 1, Data: 0x20},
				},
			},
		},
		{
			Name: "push ip",
			// mov ipl, 0x30
			// mov iph, 0x20
			// push ip
			Program: []uint8{
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_IPL, 0x30,
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_IPH, 0x20,
				instruction.INST_PUSH_REG16, instruction.REGISTER_OPCODE_IP,
			},
			Expect: &test.Expect{
				Data: []*test.ExpectData{
					{Type: test.DEV_TYPE_RAM, Addr: TEST_STACK_START, Data: 0x30},
					{Type: test.DEV_TYPE_RAM, Addr: TEST_STACK_START - 1, Data: 0x20},
				},
			},
		},
		{
			Name: "push pop ip",
			// mov ipl, 0x30
			// mov iph, 0x20
			// push ip
			// mov ip, 0x4000
			// pop ip
			Program: []uint8{
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_IPL, 0x30,
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_IPH, 0x20,
				instruction.INST_PUSH_REG16, instruction.REGISTER_OPCODE_IP,
				instruction.INST_MOV_REG16_IMM16, 0x00, 0x40,
				instruction.INST_POP_REG16, instruction.REGISTER_OPCODE_IP,
			},
			Expect: &test.Expect{
				Registers: []*test.ExpectRegister{
					{Index: instruction.REGISTER_OPCODE_IPL, Data: 0x30},
					{Index: instruction.REGISTER_OPCODE_IPH, Data: 0x20},
				},
			},
		},
		{
			Name: "add ip",
			// mov ipl, 0x20
			// mov iph, 0x10
			// add ip, 4
			// mov a, 0x10
			// push a
			Program: []uint8{
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_IPL, 0x20,
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_IPH, 0x10,
				instruction.INST_ADD_REG16_IMM8, instruction.REGISTER_OPCODE_IP, 0x04,
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x10,
				instruction.INST_PUSH_REG8, instruction.REGISTER_OPCODE_A,
			},
			Expect: &test.Expect{
				Registers: []*test.ExpectRegister{
					{Index: instruction.REGISTER_OPCODE_IPL, Data: 0x24},
					{Index: instruction.REGISTER_OPCODE_IPH, Data: 0x10},
				},
			},
		},
		{
			Name: "add ip carry",
			// mov ipl, 0xff
			// mov iph, 0x10
			// add ip, 4
			Program: []uint8{
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_IPL, 0xff,
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_IPH, 0x10,
				instruction.INST_ADD_REG16_IMM8, instruction.REGISTER_OPCODE_IP, 0x04,
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x10,
				instruction.INST_PUSH_REG8, instruction.REGISTER_OPCODE_A,
			},
			Expect: &test.Expect{
				Registers: []*test.ExpectRegister{
					{Index: instruction.REGISTER_OPCODE_IPL, Data: 0x03},
					{Index: instruction.REGISTER_OPCODE_IPH, Data: 0x11},
				},
			},
		},
		{
			Name: "sub ip",
			// mov ipl, 0x25
			// mov iph, 0x10
			// sub ip, 4
			Program: []uint8{
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_IPL, 0x25,
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_IPH, 0x10,
				instruction.INST_SUB_REG16_IMM8, instruction.REGISTER_OPCODE_IP, 0x04,
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x10,
				instruction.INST_PUSH_REG8, instruction.REGISTER_OPCODE_A,
			},
			Expect: &test.Expect{
				Registers: []*test.ExpectRegister{
					{Index: instruction.REGISTER_OPCODE_IPL, Data: 0x21},
					{Index: instruction.REGISTER_OPCODE_IPH, Data: 0x10},
				},
			},
		},
		{
			Name: "sub ip borrow",
			// mov ipl, 0x20
			// mov iph, 0x10
			// sub ip, 4
			Program: []uint8{
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_IPL, 0x20,
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_IPH, 0x10,
				instruction.INST_SUB_REG16_IMM8, instruction.REGISTER_OPCODE_IP, 0x04,
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x10,
				instruction.INST_PUSH_REG8, instruction.REGISTER_OPCODE_A,
			},
			Expect: &test.Expect{
				Registers: []*test.ExpectRegister{
					{Index: instruction.REGISTER_OPCODE_IPL, Data: 0x1c},
					{Index: instruction.REGISTER_OPCODE_IPH, Data: 0x10},
				},
			},
		},
		{
			Name: "mov reg imm16",
			// mov ip, 0x1234
			Program: []uint8{
				instruction.INST_MOV_REG16_IMM16, instruction.REGISTER_OPCODE_IP, 0x34, 0x12,
			},
			Expect: &test.Expect{
				Registers: []*test.ExpectRegister{
					{Index: instruction.REGISTER_OPCODE_IPL, Data: 0x34},
					{Index: instruction.REGISTER_OPCODE_IPH, Data: 0x12},
				},
			},
		},
		{
			Name: "mov addr reg imm",
			// mov ipl, 0x34
			// mov bpl, 0x34
			// mov iph, 0x12
			// mov bph, 0x12
			Program: []uint8{
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_IPL, 0x34,
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_BPL, 0x34,
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_IPH, 0x12,
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_BPH, 0x12,
			},
			Expect: &test.Expect{
				Registers: []*test.ExpectRegister{
					{Index: instruction.REGISTER_OPCODE_IPL, Data: 0x34},
					{Index: instruction.REGISTER_OPCODE_IPH, Data: 0x12},
					{Index: instruction.REGISTER_OPCODE_BPL, Data: 0x34},
					{Index: instruction.REGISTER_OPCODE_BPH, Data: 0x12},
				},
			},
		},
		{
			Name: "mov addr reg reg",
			// mov a, 0x34
			// mov ipl, a
			// mov bpl, a
			// mov a, 0x12
			// mov iph, a
			// mov bph, a
			Program: []uint8{
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x34,
				instruction.INST_MOV_REG8_REG8, instruction.REGISTER_OPCODE_A<<1 | instruction.REGISTER_OPCODE_IPL,
				instruction.INST_MOV_REG8_REG8, instruction.REGISTER_OPCODE_A<<1 | instruction.REGISTER_OPCODE_BPL,
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x12,
				instruction.INST_MOV_REG8_REG8, instruction.REGISTER_OPCODE_A<<1 | instruction.REGISTER_OPCODE_IPH,
				instruction.INST_MOV_REG8_REG8, instruction.REGISTER_OPCODE_A<<1 | instruction.REGISTER_OPCODE_BPH,
			},
			Expect: &test.Expect{
				Registers: []*test.ExpectRegister{
					{Index: instruction.REGISTER_OPCODE_IPL, Data: 0x34},
					{Index: instruction.REGISTER_OPCODE_IPH, Data: 0x12},
					{Index: instruction.REGISTER_OPCODE_BPL, Data: 0x34},
					{Index: instruction.REGISTER_OPCODE_BPH, Data: 0x12},
				},
			},
		},
		{
			Name: "mov addr reg addr reg",
			// mov ip, 0x1234
			// mov bp, ip
			Program: []uint8{
				instruction.INST_MOV_REG16_IMM16, instruction.REGISTER_OPCODE_IP, 0x34, 0x12,
				instruction.INST_MOV_REG16_REG16, instruction.REGISTER_OPCODE_BP<<4 | instruction.REGISTER_OPCODE_IP,
			},
			Expect: &test.Expect{
				Registers: []*test.ExpectRegister{
					{Index: instruction.REGISTER_OPCODE_IPL, Data: 0x34},
					{Index: instruction.REGISTER_OPCODE_IPH, Data: 0x12},
					{Index: instruction.REGISTER_OPCODE_BPL, Data: 0x34},
					{Index: instruction.REGISTER_OPCODE_BPH, Data: 0x12},
				},
			},
		},
		{
			Name: "mov indirect reg",
			// mov a, 0x30
			// mov [0x8000], a
			// mov ip, 0x8000
			// mov b, [ip]
			Program: []uint8{
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x30,
				instruction.INST_MOV_REG_TO_MEM_DIRECT, instruction.REGISTER_OPCODE_A, 0x00, 0x80,
				instruction.INST_MOV_REG16_IMM16, instruction.REGISTER_OPCODE_IP, 0x00, 0x80,
				instruction.INST_MOV_MEM_TO_REG_INDIRECT, instruction.REGISTER_OPCODE_B<<4 | instruction.REGISTER_OPCODE_IP,
			},
			Expect: &test.Expect{
				Registers: []*test.ExpectRegister{
					{Index: instruction.REGISTER_OPCODE_A, Data: 0x30},
					{Index: instruction.REGISTER_OPCODE_B, Data: 0x30},
				},
				Data: []*test.ExpectData{
					{Type: test.DEV_TYPE_RAM, Addr: 0x8000, Data: 0x30},
				},
			},
		},
		{
			Name: "mov indirect offset reg",
			// mov a, 0x30
			// mov [0x8008], a
			// mov ip, 0x8000
			// mov b, [ip+8]
			Program: []uint8{
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x30,
				instruction.INST_MOV_REG_TO_MEM_DIRECT, instruction.REGISTER_OPCODE_A, 0x08, 0x80,
				instruction.INST_MOV_REG16_IMM16, instruction.REGISTER_OPCODE_IP, 0x00, 0x80,
				instruction.INST_MOV_MEM_TO_REG_INDIRECT_OFFSET, instruction.REGISTER_OPCODE_B<<4 | instruction.REGISTER_OPCODE_IP, 0x08,
			},
			Expect: &test.Expect{
				Registers: []*test.ExpectRegister{
					{Index: instruction.REGISTER_OPCODE_A, Data: 0x30},
					{Index: instruction.REGISTER_OPCODE_B, Data: 0x30},
				},
				Data: []*test.ExpectData{
					{Type: test.DEV_TYPE_RAM, Addr: 0x8008, Data: 0x30},
				},
			},
		},
		{
			Name: "mov indexed reg",
			// mov a, 0x30
			// mov b, 0x08
			// mov [0x8008], a
			// mov ip, 0x8000
			// mov c, [ip+b]
			Program: []uint8{
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x30,
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_B, 0x08,
				instruction.INST_MOV_REG_TO_MEM_DIRECT, instruction.REGISTER_OPCODE_A, 0x08, 0x80,
				instruction.INST_MOV_REG16_IMM16, instruction.REGISTER_OPCODE_IP, 0x00, 0x80,
				instruction.INST_MOV_MEM_TO_REG_REG16_INDEXED, instruction.REGISTER_OPCODE_IP<<4 | instruction.REGISTER_OPCODE_B, instruction.REGISTER_OPCODE_C,
			},
			Expect: &test.Expect{
				Registers: []*test.ExpectRegister{
					{Index: instruction.REGISTER_OPCODE_A, Data: 0x30},
					{Index: instruction.REGISTER_OPCODE_B, Data: 0x08},
					{Index: instruction.REGISTER_OPCODE_C, Data: 0x30},
				},
				Data: []*test.ExpectData{
					{Type: test.DEV_TYPE_RAM, Addr: 0x8008, Data: 0x30},
				},
			},
		},
		{
			Name: "mov reg indirect",
			// mov a, 0x30
			// mov ip, 0x8000
			// mov [ip], a
			Program: []uint8{
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x30,
				instruction.INST_MOV_REG16_IMM16, instruction.REGISTER_OPCODE_IP, 0x00, 0x80,
				instruction.INST_MOV_REG_TO_MEM_INDIRECT, instruction.REGISTER_OPCODE_IP<<4 | instruction.REGISTER_OPCODE_A,
			},
			Expect: &test.Expect{
				Registers: []*test.ExpectRegister{
					{Index: instruction.REGISTER_OPCODE_A, Data: 0x30},
					{Index: instruction.REGISTER_OPCODE_IPL, Data: 0x00},
					{Index: instruction.REGISTER_OPCODE_IPH, Data: 0x80},
				},
				Data: []*test.ExpectData{
					{Type: test.DEV_TYPE_RAM, Addr: 0x8000, Data: 0x30},
				},
			},
		},
		{
			Name: "mov reg indirect offset",
			// mov a, 0x30
			// mov bp, 0x8000
			// mov [bp+8], a
			Program: []uint8{
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x30,
				instruction.INST_MOV_REG16_IMM16, instruction.REGISTER_OPCODE_BP, 0x00, 0x80,
				instruction.INST_MOV_REG_TO_MEM_INDIRECT_OFFSET, instruction.REGISTER_OPCODE_BP<<4 | instruction.REGISTER_OPCODE_A, 0x08,
			},
			Expect: &test.Expect{
				Registers: []*test.ExpectRegister{
					{Index: instruction.REGISTER_OPCODE_A, Data: 0x30},
					{Index: instruction.REGISTER_OPCODE_BPL, Data: 0x00},
					{Index: instruction.REGISTER_OPCODE_BPH, Data: 0x80},
				},
				Data: []*test.ExpectData{
					{Type: test.DEV_TYPE_RAM, Addr: 0x8008, Data: 0x30},
				},
			},
		},
		{
			Name: "mov reg indexed",
			// mov a, 0x30
			// mov b, 0x08
			// mov bp, 0x8000
			// mov [bp+b], a
			Program: []uint8{
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x30,
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_B, 0x08,
				instruction.INST_MOV_REG16_IMM16, instruction.REGISTER_OPCODE_BP, 0x00, 0x80,
				instruction.INST_MOV_REG_TO_MEM_REG16_INDEXED, instruction.REGISTER_OPCODE_BP<<4 | instruction.REGISTER_OPCODE_B, instruction.REGISTER_OPCODE_A,
			},
			Expect: &test.Expect{
				Registers: []*test.ExpectRegister{
					{Index: instruction.REGISTER_OPCODE_A, Data: 0x30},
					{Index: instruction.REGISTER_OPCODE_B, Data: 0x08},
					{Index: instruction.REGISTER_OPCODE_BPL, Data: 0x00},
					{Index: instruction.REGISTER_OPCODE_BPH, Data: 0x80},
				},
				Data: []*test.ExpectData{
					{Type: test.DEV_TYPE_RAM, Addr: 0x8008, Data: 0x30},
				},
			},
		},
		{
			Name: "mov sp address reg",
			// mov bp, sp
			// mov ip, sp
			Program: []uint8{
				instruction.INST_MOV_SP_REG16, instruction.REGISTER_OPCODE_BP,
				instruction.INST_MOV_SP_REG16, instruction.REGISTER_OPCODE_IP,
			},
			Expect: &test.Expect{
				Registers: []*test.ExpectRegister{
					{Index: instruction.REGISTER_OPCODE_IPL, Data: uint8(TEST_STACK_START)},
					{Index: instruction.REGISTER_OPCODE_IPH, Data: uint8((TEST_STACK_START & 0xff00) >> 8)},
					{Index: instruction.REGISTER_OPCODE_BPL, Data: uint8(TEST_STACK_START)},
					{Index: instruction.REGISTER_OPCODE_BPH, Data: uint8((TEST_STACK_START & 0xff00) >> 8)},
				},
			},
		},
		{
			Name: "mov address reg sp",
			// mov ip, 0x8000
			// mov bp, 0x8000
			// mov sp, bp
			// mov sp, ip
			Program: []uint8{
				instruction.INST_MOV_REG16_IMM16, instruction.REGISTER_OPCODE_IP, 0x00, 0x80,
				instruction.INST_MOV_REG16_IMM16, instruction.REGISTER_OPCODE_BP, 0x00, 0x80,
				instruction.INST_MOV_REG16_SP, instruction.REGISTER_OPCODE_BP,
				instruction.INST_MOV_REG16_SP, instruction.REGISTER_OPCODE_IP,
			},
			Expect: &test.Expect{
				Registers: []*test.ExpectRegister{
					{Index: instruction.REGISTER_OPCODE_IPL, Data: 0x00},
					{Index: instruction.REGISTER_OPCODE_IPH, Data: 0x80},
					{Index: instruction.REGISTER_OPCODE_BPL, Data: 0x00},
					{Index: instruction.REGISTER_OPCODE_BPH, Data: 0x80},
				},
			},
		},
		{
			Name: "add address register",
			// mov ip, 0x80ff
			// add ip, 8
			Program: []uint8{
				instruction.INST_MOV_REG16_IMM16, instruction.REGISTER_OPCODE_IP, 0xff, 0x80,
				instruction.INST_ADD_REG16_IMM8, instruction.REGISTER_OPCODE_IP, 0x08,
			},
			Expect: &test.Expect{
				Registers: []*test.ExpectRegister{
					{Index: instruction.REGISTER_OPCODE_IPL, Data: 0x07},
					{Index: instruction.REGISTER_OPCODE_IPH, Data: 0x81},
				},
			},
		},
		{
			Name: "sub address register",
			// mov bp, 0x8107
			// sub bp, 0x08
			Program: []uint8{
				instruction.INST_MOV_REG16_IMM16, instruction.REGISTER_OPCODE_BP, 0x07, 0x81,
				instruction.INST_SUB_REG16_IMM8, instruction.REGISTER_OPCODE_BP, 0x08,
			},
			Expect: &test.Expect{
				Registers: []*test.ExpectRegister{
					{Index: instruction.REGISTER_OPCODE_BPL, Data: 0xff},
					{Index: instruction.REGISTER_OPCODE_BPH, Data: 0x80},
				},
			},
		},
		{
			Name: "inc address register",
			// mov bp, 0x80ff
			// inc bp
			Program: []uint8{
				instruction.INST_MOV_REG16_IMM16, instruction.REGISTER_OPCODE_BP, 0xff, 0x80,
				instruction.INST_INC_REG16, instruction.REGISTER_OPCODE_BP,
			},
			Expect: &test.Expect{
				Registers: []*test.ExpectRegister{
					{Index: instruction.REGISTER_OPCODE_BPL, Data: 0x00},
					{Index: instruction.REGISTER_OPCODE_BPH, Data: 0x81},
				},
			},
		},
		{
			Name: "dec address register",
			// mov bp, 0x8100
			// dec bp
			Program: []uint8{
				instruction.INST_MOV_REG16_IMM16, instruction.REGISTER_OPCODE_BP, 0x00, 0x81,
				instruction.INST_DEC_REG16, instruction.REGISTER_OPCODE_BP,
			},
			Expect: &test.Expect{
				Registers: []*test.ExpectRegister{
					{Index: instruction.REGISTER_OPCODE_BPL, Data: 0xff},
					{Index: instruction.REGISTER_OPCODE_BPH, Data: 0x80},
				},
			},
		},
		{
			Name: "index pointer reconstruction",
			// mov a, 0xff
			// mov [0x3010], a
			// mov ipl, 0x10
			// mov iph, 0x30
			// mov b, [ip]
			Program: []uint8{
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0xff,
				instruction.INST_MOV_REG_TO_MEM_DIRECT, instruction.REGISTER_OPCODE_A, 0x10, 0x30,
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_IPL, 0x10,
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_IPH, 0x30,
				instruction.INST_MOV_MEM_TO_REG_INDIRECT, instruction.REGISTER_OPCODE_B<<4 | instruction.REGISTER_OPCODE_IP,
			},
			Expect: &test.Expect{
				Registers: []*test.ExpectRegister{
					{Index: instruction.REGISTER_OPCODE_IPL, Data: 0x10},
					{Index: instruction.REGISTER_OPCODE_IPH, Data: 0x30},
					{Index: instruction.REGISTER_OPCODE_A, Data: 0xff},
					{Index: instruction.REGISTER_OPCODE_B, Data: 0xff},
				},
				Data: []*test.ExpectData{
					{Type: test.DEV_TYPE_RAM, Addr: 0x3010, Data: 0xff},
				},
			},
		},
	}
)

func TestCore(t *testing.T) {
	// Run only test cases with debug flag
	for _, tc := range Tests {
		if tc.Debug {
			c := test.GetComp()
			c.SetDebug(true)
			test.RunCase(t, tc)
			return
		}
	}
	// Run all test case unless disabled
	for _, tc := range Tests {
		if tc.Disable {
			continue
		}
		test.RunCase(t, tc)
	}
}
