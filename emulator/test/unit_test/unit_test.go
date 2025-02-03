package unit_test

import (
	"asme8/common/instruction"
	"asme8/emulator/src/status"
	"asme8/emulator/test"
	"testing"
)

var (
	Tests = []*test.TestCase{
		{
			Name: "mov immediate reg a",
			// mov a, 0x10
			Program: []uint8{instruction.INST_MOV_REG_IMM, instruction.REGISTER_OPCODE_A, 0x10},
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
			Program: []uint8{instruction.INST_MOV_REG_IMM, instruction.REGISTER_OPCODE_A, 0xff},
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
				instruction.INST_MOV_REG_IMM, instruction.REGISTER_OPCODE_A, 0x10,
				instruction.INST_MOV_REG_REG, instruction.REGISTER_OPCODE_A<<1 | instruction.REGISTER_OPCODE_B,
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
				instruction.INST_MOV_REG_IMM, instruction.REGISTER_OPCODE_A, 0xff,
				instruction.INST_MOV_REG_REG, instruction.REGISTER_OPCODE_A<<4 | instruction.REGISTER_OPCODE_B,
				instruction.INST_MOV_REG_REG, instruction.REGISTER_OPCODE_A<<4 | instruction.REGISTER_OPCODE_C,
				instruction.INST_MOV_REG_REG, instruction.REGISTER_OPCODE_A<<4 | instruction.REGISTER_OPCODE_D,
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
				instruction.INST_MOV_REG_IMM, instruction.REGISTER_OPCODE_A, 0xff,
				instruction.INST_MOV_REG_MEM, instruction.REGISTER_OPCODE_A, 0x00, 0x90,
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
				instruction.INST_MOV_REG_IMM, instruction.REGISTER_OPCODE_A, 0x90,
				instruction.INST_MOV_REG_MEM, instruction.REGISTER_OPCODE_A, 0x00, 0x90,
				instruction.INST_MOV_REG_IMM, instruction.REGISTER_OPCODE_A, 0x91,
				instruction.INST_MOV_REG_MEM, instruction.REGISTER_OPCODE_A, 0x01, 0x90,
				instruction.INST_MOV_REG_IMM, instruction.REGISTER_OPCODE_A, 0x92,
				instruction.INST_MOV_REG_MEM, instruction.REGISTER_OPCODE_A, 0x02, 0x90,
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
				instruction.INST_MOV_REG_IMM, instruction.REGISTER_OPCODE_A, 0x90,
				instruction.INST_MOV_REG_MEM, instruction.REGISTER_OPCODE_A, 0x00, 0x90,
				instruction.INST_MOV_MEM_REG, instruction.REGISTER_OPCODE_B, 0x00, 0x90,
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
				instruction.INST_MOV_REG_IMM, instruction.REGISTER_OPCODE_A, 0x10,
				instruction.INST_MOV_REG_IMM, instruction.REGISTER_OPCODE_B, 0x2,
				instruction.INST_MOV_REG_MEM, instruction.REGISTER_OPCODE_A, 0x02, 0x90,
				instruction.INST_MOV_MEM_REG_OFFSET, instruction.REGISTER_OPCODE_C<<4 | instruction.REGISTER_OPCODE_B, 0x00, 0x90,
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
				instruction.INST_MOV_REG_IMM, instruction.REGISTER_OPCODE_A, 0x10,
				instruction.INST_MOV_REG_IMM, instruction.REGISTER_OPCODE_B, 0x2,
				instruction.INST_MOV_REG_MEM, instruction.REGISTER_OPCODE_A, 0x01, 0x91,
				instruction.INST_MOV_MEM_REG_OFFSET, instruction.REGISTER_OPCODE_C<<4 | instruction.REGISTER_OPCODE_B, 0xff, 0x90,
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
				instruction.INST_MOV_REG_IMM, instruction.REGISTER_OPCODE_A, 0x10,
				instruction.INST_MOV_REG_IMM, instruction.REGISTER_OPCODE_B, 0x2,
				instruction.INST_MOV_REG_IMM, instruction.REGISTER_OPCODE_C, 0x0,
				instruction.INST_MOV_REG_MEM_OFFSET, instruction.REGISTER_OPCODE_A<<4 | instruction.REGISTER_OPCODE_B, 0xff, 0x90,
				instruction.INST_MOV_MEM_REG_OFFSET, instruction.REGISTER_OPCODE_C<<4 | instruction.REGISTER_OPCODE_B, 0xff, 0x90,
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
				instruction.INST_MOV_REG_IMM, instruction.REGISTER_OPCODE_A, 0x10,
				instruction.INST_BRK_IMPL,
				instruction.INST_MOV_REG_IMM, instruction.REGISTER_OPCODE_A, 0x20,
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
				instruction.INST_MOV_REG_IMM, instruction.REGISTER_OPCODE_A, 0x10,
				instruction.INST_JMP_IMM, 0x0a, 0x00,
				instruction.INST_BRK_IMPL,
				instruction.INST_MOV_REG_IMM, instruction.REGISTER_OPCODE_A, 0x20,
				instruction.INST_MOV_REG_IMM, instruction.REGISTER_OPCODE_A, 0x30,
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
				instruction.INST_MOV_REG_IMM, instruction.REGISTER_OPCODE_A, 0x10,
				instruction.INST_CMP_REG_IMM, instruction.REGISTER_OPCODE_A, 0x10,
				instruction.INST_JZ_IMM, 0x0d, 0x00,
				instruction.INST_MOV_REG_IMM, instruction.REGISTER_OPCODE_A, 0x20,
				instruction.INST_BRK_IMPL,
				instruction.INST_MOV_REG_IMM, instruction.REGISTER_OPCODE_A, 0x30,
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
				instruction.INST_MOV_REG_IMM, instruction.REGISTER_OPCODE_A, 0x10,
				instruction.INST_CMP_REG_IMM, instruction.REGISTER_OPCODE_A, 0x20,
				instruction.INST_JZ_IMM, 0x0d, 0x00,
				instruction.INST_MOV_REG_IMM, instruction.REGISTER_OPCODE_A, 0x20,
				instruction.INST_BRK_IMPL,
				instruction.INST_MOV_REG_IMM, instruction.REGISTER_OPCODE_A, 0x30,
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
				instruction.INST_MOV_REG_IMM, instruction.REGISTER_OPCODE_A, 0x10,
				instruction.INST_CMP_REG_IMM, instruction.REGISTER_OPCODE_A, 0x10,
				instruction.INST_JNZ_IMM, 0x0d, 0x00,
				instruction.INST_MOV_REG_IMM, instruction.REGISTER_OPCODE_A, 0x20,
				instruction.INST_BRK_IMPL,
				instruction.INST_MOV_REG_IMM, instruction.REGISTER_OPCODE_A, 0x30,
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
				instruction.INST_MOV_REG_IMM, instruction.REGISTER_OPCODE_A, 0x20,
				instruction.INST_CMP_REG_IMM, instruction.REGISTER_OPCODE_A, 0x10,
				instruction.INST_JNZ_IMM, 0x0d, 0x00,
				instruction.INST_MOV_REG_IMM, instruction.REGISTER_OPCODE_A, 0x20,
				instruction.INST_BRK_IMPL,
				instruction.INST_MOV_REG_IMM, instruction.REGISTER_OPCODE_A, 0x30,
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
				instruction.INST_MOV_REG_IMM, instruction.REGISTER_OPCODE_A, 0x10,
				instruction.INST_MOV_REG_IMM, instruction.REGISTER_OPCODE_B, 0x10,
				instruction.INST_CMP_REG_REG, instruction.REGISTER_OPCODE_A<<4 | instruction.REGISTER_OPCODE_B,
				instruction.INST_JZ_IMM, 0x0f, 0x00,
				instruction.INST_MOV_REG_IMM, instruction.REGISTER_OPCODE_A, 0x20,
				instruction.INST_BRK_IMPL,
				instruction.INST_MOV_REG_IMM, instruction.REGISTER_OPCODE_A, 0x30,
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
				instruction.INST_MOV_REG_IMM, instruction.REGISTER_OPCODE_A, 0x10,
				instruction.INST_MOV_REG_IMM, instruction.REGISTER_OPCODE_B, 0x20,
				instruction.INST_CMP_REG_REG, instruction.REGISTER_OPCODE_A<<4 | instruction.REGISTER_OPCODE_B,
				instruction.INST_JZ_IMM, 0x0f, 0x00,
				instruction.INST_MOV_REG_IMM, instruction.REGISTER_OPCODE_A, 0x20,
				instruction.INST_BRK_IMPL,
				instruction.INST_MOV_REG_IMM, instruction.REGISTER_OPCODE_A, 0x30,
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
			Program: []uint8{instruction.INST_ADD_REG_IMM, instruction.REGISTER_OPCODE_A, 0x10},
			Expect: &test.Expect{
				Registers: []*test.ExpectRegister{
					{Index: instruction.REGISTER_OPCODE_A, Data: 0x10},
				},
			},
		},
		{
			Name: "add reg imm",
			// mov a, 0x20
			// add a, 0x10
			Program: []uint8{
				instruction.INST_MOV_REG_IMM, instruction.REGISTER_OPCODE_A, 0x20,
				instruction.INST_ADD_REG_IMM, instruction.REGISTER_OPCODE_A, 0x10,
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
				instruction.INST_MOV_REG_IMM, instruction.REGISTER_OPCODE_A, 0xff,
				instruction.INST_ADD_REG_IMM, instruction.REGISTER_OPCODE_A, 0x01,
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
				instruction.INST_MOV_REG_IMM, instruction.REGISTER_OPCODE_A, 0xff,
				instruction.INST_MOV_REG_IMM, instruction.REGISTER_OPCODE_B, 0x11,
				instruction.INST_ADD_REG_REG, instruction.REGISTER_OPCODE_B<<4 | instruction.REGISTER_OPCODE_A,
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
				instruction.INST_MOV_REG_IMM, instruction.REGISTER_OPCODE_A, 0xff,
				instruction.INST_MOV_REG_IMM, instruction.REGISTER_OPCODE_B, 0x11,
				instruction.INST_ADD_REG_REG, instruction.REGISTER_OPCODE_A<<4 | instruction.REGISTER_OPCODE_B,
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
			Program: []uint8{instruction.INST_SUB_REG_IMM, instruction.REGISTER_OPCODE_A, 0x10},
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
				instruction.INST_MOV_REG_IMM, instruction.REGISTER_OPCODE_A, 0x20,
				instruction.INST_SUB_REG_IMM, instruction.REGISTER_OPCODE_A, 0x10,
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
				instruction.INST_MOV_REG_IMM, instruction.REGISTER_OPCODE_A, 0xff,
				instruction.INST_SUB_REG_IMM, instruction.REGISTER_OPCODE_A, 0x01,
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
				instruction.INST_MOV_REG_IMM, instruction.REGISTER_OPCODE_A, 0xff,
				instruction.INST_MOV_REG_IMM, instruction.REGISTER_OPCODE_B, 0x11,
				instruction.INST_SUB_REG_REG, instruction.REGISTER_OPCODE_B<<4 | instruction.REGISTER_OPCODE_A,
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
				instruction.INST_MOV_REG_IMM, instruction.REGISTER_OPCODE_A, 0xff,
				instruction.INST_MOV_REG_IMM, instruction.REGISTER_OPCODE_B, 0x11,
				instruction.INST_SUB_REG_REG, instruction.REGISTER_OPCODE_A<<4 | instruction.REGISTER_OPCODE_B,
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
				instruction.INST_MOV_REG_IMM, instruction.REGISTER_OPCODE_A, 0x10,
				instruction.INST_PUSH_REG, instruction.REGISTER_OPCODE_A,
				instruction.INST_POP_REG, instruction.REGISTER_OPCODE_B,
			},
			Expect: &test.Expect{
				Registers: []*test.ExpectRegister{
					{Index: instruction.REGISTER_OPCODE_A, Data: 0x10},
					{Index: instruction.REGISTER_OPCODE_B, Data: 0x10},
				},
			},
		},
		{
			Name: "push pop reg imm",
			// push 0x10
			// push 0x20
			// push 0x30
			// push 0x40
			// pop d
			// pop c
			// pop b
			// pop a
			Program: []uint8{
				instruction.INST_PUSH_IMM, 0x10,
				instruction.INST_PUSH_IMM, 0x20,
				instruction.INST_PUSH_IMM, 0x30,
				instruction.INST_PUSH_IMM, 0x40,
				instruction.INST_POP_REG, instruction.REGISTER_OPCODE_D,
				instruction.INST_POP_REG, instruction.REGISTER_OPCODE_C,
				instruction.INST_POP_REG, instruction.REGISTER_OPCODE_B,
				instruction.INST_POP_REG, instruction.REGISTER_OPCODE_A,
			},
			Expect: &test.Expect{
				Registers: []*test.ExpectRegister{
					{Index: instruction.REGISTER_OPCODE_A, Data: 0x10},
					{Index: instruction.REGISTER_OPCODE_B, Data: 0x20},
					{Index: instruction.REGISTER_OPCODE_C, Data: 0x30},
					{Index: instruction.REGISTER_OPCODE_D, Data: 0x40},
				},
			},
		},
		{
			Name: "inc reg",
			// mov a, 0x20
			// inc a
			Program: []uint8{
				instruction.INST_MOV_REG_IMM, instruction.REGISTER_OPCODE_A, 0x20,
				instruction.INST_INC_REG, instruction.REGISTER_OPCODE_A,
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
				instruction.INST_MOV_REG_IMM, instruction.REGISTER_OPCODE_A, 0xff,
				instruction.INST_INC_REG, instruction.REGISTER_OPCODE_A,
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
				instruction.INST_MOV_REG_IMM, instruction.REGISTER_OPCODE_A, 0x20,
				instruction.INST_DEC_REG, instruction.REGISTER_OPCODE_A,
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
				instruction.INST_MOV_REG_IMM, instruction.REGISTER_OPCODE_A, 0x00,
				instruction.INST_DEC_REG, instruction.REGISTER_OPCODE_A,
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
				instruction.INST_CALL_IMPL, 0x07, 0x00,
				instruction.INST_MOV_REG_IMM, instruction.REGISTER_OPCODE_A, 0x10,
				instruction.INST_BRK_IMPL,
				instruction.INST_MOV_REG_IMM, instruction.REGISTER_OPCODE_A, 0x20,
				instruction.INST_MOV_REG_IMM, instruction.REGISTER_OPCODE_B, 0x20,
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
			Name: "xor reg reg",
			// mov a, 0b00001111
			// mov b, 0b11001100
			// xor b, a
			Program: []uint8{
				instruction.INST_MOV_REG_IMM, instruction.REGISTER_OPCODE_A, 0b00001111,
				instruction.INST_MOV_REG_IMM, instruction.REGISTER_OPCODE_B, 0b11001100,
				instruction.INST_XOR_REG_REG, instruction.REGISTER_OPCODE_A<<4 | instruction.REGISTER_OPCODE_B,
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
				instruction.INST_MOV_REG_IMM, instruction.REGISTER_OPCODE_A, 0b00001111,
				instruction.INST_XOR_REG_REG, instruction.REGISTER_OPCODE_A<<4 | instruction.REGISTER_OPCODE_A,
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
				instruction.INST_MOV_REG_IMM, instruction.REGISTER_OPCODE_A, 0b11001100,
				instruction.INST_XOR_REG_IMM, instruction.REGISTER_OPCODE_A, 0b00001111,
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
				instruction.INST_MOV_REG_IMM, instruction.REGISTER_OPCODE_A, 0b00001111,
				instruction.INST_MOV_REG_IMM, instruction.REGISTER_OPCODE_B, 0b11001100,
				instruction.INST_AND_REG_REG, instruction.REGISTER_OPCODE_A<<4 | instruction.REGISTER_OPCODE_B,
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
				instruction.INST_MOV_REG_IMM, instruction.REGISTER_OPCODE_A, 0b11001100,
				instruction.INST_AND_REG_IMM, instruction.REGISTER_OPCODE_A, 0b00001111,
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
				instruction.INST_MOV_REG_IMM, instruction.REGISTER_OPCODE_A, 0b00001111,
				instruction.INST_MOV_REG_IMM, instruction.REGISTER_OPCODE_B, 0b11001100,
				instruction.INST_OR_REG_REG, instruction.REGISTER_OPCODE_A<<4 | instruction.REGISTER_OPCODE_B,
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
				instruction.INST_MOV_REG_IMM, instruction.REGISTER_OPCODE_A, 0b11001100,
				instruction.INST_OR_REG_IMM, instruction.REGISTER_OPCODE_A, 0b00001111,
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
				instruction.INST_MOV_REG_IMM, instruction.REGISTER_OPCODE_A, 0b11001100,
				instruction.INST_NOT_REG, instruction.REGISTER_OPCODE_A,
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
				instruction.INST_MOV_REG_IMM, instruction.REGISTER_OPCODE_A, 0b11001100,
				instruction.INST_SHL_REG, instruction.REGISTER_OPCODE_A,
			},
			Expect: &test.Expect{
				Registers: []*test.ExpectRegister{
					{Index: instruction.REGISTER_OPCODE_A, Data: 0b10011000},
				},
				Status: &test.ExpectStatusData{
					Data: 0 | status.STATUS_FLAG_CARRY | status.STATUS_FLAG_SIGN,
				},
			},
		},
		{
			Name: "shl reg zero",
			// mov a, 0b00110011
			// shl a,
			Program: []uint8{
				instruction.INST_MOV_REG_IMM, instruction.REGISTER_OPCODE_A, 0b00110011,
				instruction.INST_SHL_REG, instruction.REGISTER_OPCODE_A,
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
				instruction.INST_MOV_REG_IMM, instruction.REGISTER_OPCODE_A, 0b11001100,
				instruction.INST_SHR_REG, instruction.REGISTER_OPCODE_A,
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
				instruction.INST_MOV_REG_IMM, instruction.REGISTER_OPCODE_A, 0b00110011,
				instruction.INST_SHR_REG, instruction.REGISTER_OPCODE_A,
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
				instruction.INST_MOV_REG_IMM, instruction.REGISTER_OPCODE_A, 0b11001100,
				instruction.INST_ROR_REG, instruction.REGISTER_OPCODE_A,
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
				instruction.INST_MOV_REG_IMM, instruction.REGISTER_OPCODE_A, 0b00110011,
				instruction.INST_ROR_REG, instruction.REGISTER_OPCODE_A,
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
				instruction.INST_MOV_REG_IMM, instruction.REGISTER_OPCODE_A, 0b11001100,
				instruction.INST_ROL_REG, instruction.REGISTER_OPCODE_A,
			},
			Expect: &test.Expect{
				Registers: []*test.ExpectRegister{
					{Index: instruction.REGISTER_OPCODE_A, Data: 0b10011000},
				},
				Status: &test.ExpectStatusData{
					Data: 0 | status.STATUS_FLAG_CARRY | status.STATUS_FLAG_SIGN,
				},
			},
		},
		{
			Name: "rol zero carry",
			// mov a, 0b00110011
			// rol a,
			Program: []uint8{
				instruction.INST_MOV_REG_IMM, instruction.REGISTER_OPCODE_A, 0b00110011,
				instruction.INST_ROL_REG, instruction.REGISTER_OPCODE_A,
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
			// mov ipl, 0xf0
			// add [0x9000], ipl
			// mov a, 0x30
			// add [0x9000], a
			Name: "add reg mem",
			Program: []uint8{
				instruction.INST_MOV_REG_IMM, instruction.REGISTER_OPCODE_IPL, 0xf0,
				instruction.INST_ADD_REG_MEM, instruction.REGISTER_OPCODE_IPL, 0x00, 0x90,
				instruction.INST_MOV_REG_IMM, instruction.REGISTER_OPCODE_A, 0x30,
				instruction.INST_ADD_REG_MEM, instruction.REGISTER_OPCODE_A, 0x00, 0x90,
			},
			Expect: &test.Expect{
				Data: []*test.ExpectData{
					{Type: test.DEV_TYPE_RAM, Addr: 0x9000, Data: 0x20},
				},
			},
		},
		{
			// mov a, 0xff
			// mov [0x9000], a
			// mov a, 0x01
			// add [0x9000], a
			// mov a, 0x00
			// adc [0x9001], a
			Name: "adc reg mem",
			Program: []uint8{
				instruction.INST_MOV_REG_IMM, instruction.REGISTER_OPCODE_A, 0xff,
				instruction.INST_ADD_REG_MEM, instruction.REGISTER_OPCODE_A, 0x00, 0x90,
				instruction.INST_MOV_REG_IMM, instruction.REGISTER_OPCODE_A, 0x01,
				instruction.INST_ADD_REG_MEM, instruction.REGISTER_OPCODE_A, 0x00, 0x90,
				instruction.INST_MOV_REG_IMM, instruction.REGISTER_OPCODE_A, 0x00,
				instruction.INST_ADC_REG_MEM, instruction.REGISTER_OPCODE_A, 0x01, 0x90,
			},
			Expect: &test.Expect{
				Data: []*test.ExpectData{
					{Type: test.DEV_TYPE_RAM, Addr: 0x9000, Data: 0x00},
					{Type: test.DEV_TYPE_RAM, Addr: 0x9001, Data: 0x01},
				},
			},
		},
		{
			// mov a, 0xff
			// mov [0x9000], a
			// mov a, 0x01
			// mov [0x9001], a
			// mov ipl, 0x01
			// mov iph, 0x01
			// add ipl, [0x9000]
			// adc iph, [0x9001]
			Name: "adc mem reg",
			Program: []uint8{
				instruction.INST_MOV_REG_IMM, instruction.REGISTER_OPCODE_A, 0xff,
				instruction.INST_MOV_REG_MEM, instruction.REGISTER_OPCODE_A, 0x00, 0x90,
				instruction.INST_MOV_REG_IMM, instruction.REGISTER_OPCODE_A, 0x01,
				instruction.INST_MOV_REG_MEM, instruction.REGISTER_OPCODE_A, 0x01, 0x90,
				instruction.INST_MOV_REG_IMM, instruction.REGISTER_OPCODE_IPL, 0x01,
				instruction.INST_MOV_REG_IMM, instruction.REGISTER_OPCODE_IPH, 0x01,
				instruction.INST_ADD_MEM_REG, instruction.REGISTER_OPCODE_IPL, 0x00, 0x90,
				instruction.INST_ADC_MEM_REG, instruction.REGISTER_OPCODE_IPH, 0x01, 0x90,
			},
			Expect: &test.Expect{
				Data: []*test.ExpectData{
					{Type: test.DEV_TYPE_RAM, Addr: 0x9000, Data: 0xff},
					{Type: test.DEV_TYPE_RAM, Addr: 0x9001, Data: 0x01},
				},
				Registers: []*test.ExpectRegister{
					{Index: instruction.REGISTER_OPCODE_IPL, Data: 0x00},
					{Index: instruction.REGISTER_OPCODE_IPH, Data: 0x03},
				},
			},
		},
		{
			// mov a, 0x30
			// mov [0x9000], a
			// add a, [0x9000]
			Name: "add mem reg",
			Program: []uint8{
				instruction.INST_MOV_REG_IMM, instruction.REGISTER_OPCODE_A, 0x30,
				instruction.INST_MOV_REG_MEM, instruction.REGISTER_OPCODE_A, 0x00, 0x90,
				instruction.INST_ADD_MEM_REG, instruction.REGISTER_OPCODE_A, 0x00, 0x90,
			},
			Expect: &test.Expect{
				Registers: []*test.ExpectRegister{
					{Index: instruction.REGISTER_OPCODE_A, Data: 0x60},
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
				instruction.INST_MOV_REG_IMM, instruction.REGISTER_OPCODE_A, 0xff,
				instruction.INST_MOV_REG_IMM, instruction.REGISTER_OPCODE_B, 0x02,
				instruction.INST_ADD_REG_REG, instruction.REGISTER_OPCODE_B<<4 | instruction.REGISTER_OPCODE_A,
				instruction.INST_JC_IMM, 0x0f, 0x00,
				instruction.INST_MOV_REG_IMM, instruction.REGISTER_OPCODE_C, 0x20,
				instruction.INST_BRK_IMPL,
				instruction.INST_MOV_REG_IMM, instruction.REGISTER_OPCODE_C, 0x30,
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
				instruction.INST_MOV_REG_IMM, instruction.REGISTER_OPCODE_A, 0x01,
				instruction.INST_MOV_REG_IMM, instruction.REGISTER_OPCODE_B, 0x02,
				instruction.INST_ADD_REG_REG, instruction.REGISTER_OPCODE_B<<4 | instruction.REGISTER_OPCODE_A,
				instruction.INST_JC_IMM, 0x0f, 0x00,
				instruction.INST_MOV_REG_IMM, instruction.REGISTER_OPCODE_C, 0x20,
				instruction.INST_BRK_IMPL,
				instruction.INST_MOV_REG_IMM, instruction.REGISTER_OPCODE_C, 0x30,
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
				instruction.INST_MOV_REG_IMM, instruction.REGISTER_OPCODE_A, 0xff,
				instruction.INST_MOV_REG_IMM, instruction.REGISTER_OPCODE_B, 0x02,
				instruction.INST_ADD_REG_REG, instruction.REGISTER_OPCODE_B<<4 | instruction.REGISTER_OPCODE_A,
				instruction.INST_CLC_IMPL,
				instruction.INST_JC_IMM, 0x10, 0x00,
				instruction.INST_MOV_REG_IMM, instruction.REGISTER_OPCODE_C, 0x20,
				instruction.INST_BRK_IMPL,
				instruction.INST_MOV_REG_IMM, instruction.REGISTER_OPCODE_C, 0x30,
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
				instruction.INST_MOV_REG_IMM, instruction.REGISTER_OPCODE_A, 0x10,
				instruction.INST_CMP_REG_IMM, instruction.REGISTER_OPCODE_A, 0x20,
				instruction.INST_JS_IMM, 0x0d, 0x00,
				instruction.INST_MOV_REG_IMM, instruction.REGISTER_OPCODE_A, 0x20,
				instruction.INST_BRK_IMPL,
				instruction.INST_MOV_REG_IMM, instruction.REGISTER_OPCODE_A, 0x30,
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
				instruction.INST_MOV_REG_IMM, instruction.REGISTER_OPCODE_A, 0x20,
				instruction.INST_CMP_REG_IMM, instruction.REGISTER_OPCODE_A, 0x10,
				instruction.INST_JS_IMM, 0x0d, 0x00,
				instruction.INST_MOV_REG_IMM, instruction.REGISTER_OPCODE_A, 0x20,
				instruction.INST_BRK_IMPL,
				instruction.INST_MOV_REG_IMM, instruction.REGISTER_OPCODE_A, 0x30,
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
				instruction.INST_MOV_REG_IMM, instruction.REGISTER_OPCODE_A, 0x20,
				instruction.INST_PUSH_REG, instruction.REGISTER_OPCODE_A,
				instruction.INST_PUSH_SP,
			},
			Expect: &test.Expect{
				Data: []*test.ExpectData{
					{Type: test.DEV_TYPE_RAM, Addr: 0x20ff - 1, Data: 0xfe},
					{Type: test.DEV_TYPE_RAM, Addr: 0x20ff - 2, Data: 0x20},
				},
			},
		},
		{
			Name: "add sp",
			// add sp, 4
			// mov a, 0x10
			// push a
			Program: []uint8{
				instruction.INST_ADD_SP_IMM, 0x04,
				instruction.INST_MOV_REG_IMM, instruction.REGISTER_OPCODE_A, 0x10,
				instruction.INST_PUSH_REG, instruction.REGISTER_OPCODE_A,
			},
			Expect: &test.Expect{
				Data: []*test.ExpectData{
					{Type: test.DEV_TYPE_RAM, Addr: 0x20ff + 4, Data: 0x10},
				},
			},
		},
		{
			Name: "sub sp",
			// sub sp, 4
			// mov a, 0x10
			// push a
			Program: []uint8{
				instruction.INST_SUB_SP_IMM, 0x04,
				instruction.INST_MOV_REG_IMM, instruction.REGISTER_OPCODE_A, 0x10,
				instruction.INST_PUSH_REG, instruction.REGISTER_OPCODE_A,
			},
			Expect: &test.Expect{
				Data: []*test.ExpectData{
					{Type: test.DEV_TYPE_RAM, Addr: 0x20ff - 4, Data: 0x10},
				},
			},
		},
		{
			Name: "mov reg sp",
			// mov b, 0x10
			// mov [sp], b
			Program: []uint8{
				instruction.INST_MOV_REG_IMM, instruction.REGISTER_OPCODE_B, 0x10,
				instruction.INST_MOV_REG_SP, instruction.REGISTER_OPCODE_B,
			},
			Expect: &test.Expect{
				Data: []*test.ExpectData{
					{Type: test.DEV_TYPE_RAM, Addr: 0x20ff, Data: 0x10},
				},
			},
		},
		{
			Name: "mov sp reg",
			// mov b, 0x10
			// mov [sp], b
			// mov a, [sp]
			Program: []uint8{
				instruction.INST_MOV_REG_IMM, instruction.REGISTER_OPCODE_B, 0x10,
				instruction.INST_MOV_REG_SP, instruction.REGISTER_OPCODE_B,
				instruction.INST_MOV_SP_REG, instruction.REGISTER_OPCODE_A,
			},
			Expect: &test.Expect{
				Registers: []*test.ExpectRegister{
					{Index: instruction.REGISTER_OPCODE_A, Data: 0x10},
					{Index: instruction.REGISTER_OPCODE_B, Data: 0x10},
				},
			},
		},
		{
			Name: "mov reg sp offset",
			// mov b, 0x10
			// mov [sp+2], b
			Program: []uint8{
				instruction.INST_MOV_REG_IMM, instruction.REGISTER_OPCODE_B, 0x10,
				instruction.INST_MOV_REG_SP_OFFSET, instruction.REGISTER_OPCODE_B, 0x02,
			},
			Expect: &test.Expect{
				Data: []*test.ExpectData{
					{Type: test.DEV_TYPE_RAM, Addr: 0x20ff + 2, Data: 0x10},
				},
			},
		},
		{
			Name: "mov sp reg offset",
			// mov b, 0x10
			// mov [sp+2], b
			// mov a, [sp+2]
			Program: []uint8{
				instruction.INST_MOV_REG_IMM, instruction.REGISTER_OPCODE_B, 0x10,
				instruction.INST_MOV_REG_SP_OFFSET, instruction.REGISTER_OPCODE_B, 0x02,
				instruction.INST_MOV_SP_REG_OFFSET, instruction.REGISTER_OPCODE_A, 0x02,
			},
			Expect: &test.Expect{
				Registers: []*test.ExpectRegister{
					{Index: instruction.REGISTER_OPCODE_A, Data: 0x10},
					{Index: instruction.REGISTER_OPCODE_B, Data: 0x10},
				},
			},
		},
		{
			Name: "mov reg sp offset reg",
			// mov a, 0x02
			// mov b, 0x10
			// mov [sp+a], b
			Program: []uint8{
				instruction.INST_MOV_REG_IMM, instruction.REGISTER_OPCODE_A, 0x02,
				instruction.INST_MOV_REG_IMM, instruction.REGISTER_OPCODE_B, 0x10,
				instruction.INST_MOV_REG_SP_OFFSET_REG, instruction.REGISTER_OPCODE_A<<4 | instruction.REGISTER_OPCODE_B,
			},
			Expect: &test.Expect{
				Data: []*test.ExpectData{
					{Type: test.DEV_TYPE_RAM, Addr: 0x20ff + 2, Data: 0x10},
				},
			},
		},
		{
			Name: "mov sp reg offset reg",
			// mov a, 0x02
			// mov b, 0x10
			// mov [sp+a], b
			// mov b, [sp+a]
			Program: []uint8{
				instruction.INST_MOV_REG_IMM, instruction.REGISTER_OPCODE_A, 0x02,
				instruction.INST_MOV_REG_IMM, instruction.REGISTER_OPCODE_B, 0x10,
				instruction.INST_MOV_REG_SP_OFFSET_REG, instruction.REGISTER_OPCODE_A<<4 | instruction.REGISTER_OPCODE_B,
				instruction.INST_MOV_SP_REG_OFFSET_REG, instruction.REGISTER_OPCODE_A<<4 | instruction.REGISTER_OPCODE_B,
			},
			Expect: &test.Expect{
				Registers: []*test.ExpectRegister{
					{Index: instruction.REGISTER_OPCODE_B, Data: 0x10},
				},
			},
		},
		{
			Name: "mov ipl iph imm",
			// mov ipl, 0x10
			// mov iph, 0x20
			// push ip
			Program: []uint8{
				instruction.INST_MOV_REG_IMM, instruction.REGISTER_OPCODE_IPL, 0x20,
				instruction.INST_MOV_REG_IMM, instruction.REGISTER_OPCODE_IPH, 0x10,
				instruction.INST_PUSH_IP,
			},
			Expect: &test.Expect{
				Registers: []*test.ExpectRegister{
					{Index: instruction.REGISTER_OPCODE_IPL, Data: 0x20},
					{Index: instruction.REGISTER_OPCODE_IPH, Data: 0x10},
				},
				Data: []*test.ExpectData{
					{Type: test.DEV_TYPE_RAM, Addr: 0x20ff, Data: 0x20},
					{Type: test.DEV_TYPE_RAM, Addr: 0x20ff - 1, Data: 0x10},
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
				instruction.INST_MOV_REG_IMM, instruction.REGISTER_OPCODE_A, 0x10,
				instruction.INST_MOV_REG_IMM, instruction.REGISTER_OPCODE_B, 0x20,
				instruction.INST_MOV_REG_REG, instruction.REGISTER_OPCODE_A<<4 | instruction.REGISTER_OPCODE_IPL,
				instruction.INST_MOV_REG_REG, instruction.REGISTER_OPCODE_B<<4 | instruction.REGISTER_OPCODE_IPH,
				instruction.INST_PUSH_IP,
			},
			Expect: &test.Expect{
				Registers: []*test.ExpectRegister{
					{Index: instruction.REGISTER_OPCODE_A, Data: 0x10},
					{Index: instruction.REGISTER_OPCODE_B, Data: 0x20},
					{Index: instruction.REGISTER_OPCODE_IPL, Data: 0x10},
					{Index: instruction.REGISTER_OPCODE_IPH, Data: 0x20},
				},
				Data: []*test.ExpectData{
					{Type: test.DEV_TYPE_RAM, Addr: 0x20ff, Data: 0x10},
					{Type: test.DEV_TYPE_RAM, Addr: 0x20ff - 1, Data: 0x20},
				},
			},
		},
		{
			Name: "push ip",
			// mov ipl, 0x30
			// mov iph, 0x20
			// push ip
			Program: []uint8{
				instruction.INST_MOV_REG_IMM, instruction.REGISTER_OPCODE_IPL, 0x30,
				instruction.INST_MOV_REG_IMM, instruction.REGISTER_OPCODE_IPH, 0x20,
				instruction.INST_PUSH_IP,
			},
			Expect: &test.Expect{
				Data: []*test.ExpectData{
					{Type: test.DEV_TYPE_RAM, Addr: 0x20ff, Data: 0x30},
					{Type: test.DEV_TYPE_RAM, Addr: 0x20ff - 1, Data: 0x20},
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
				instruction.INST_MOV_REG_IMM, instruction.REGISTER_OPCODE_IPL, 0x30,
				instruction.INST_MOV_REG_IMM, instruction.REGISTER_OPCODE_IPH, 0x20,
				instruction.INST_PUSH_IP,
				instruction.INST_MOV_IP_IMM_16, 0x00, 0x40,
				instruction.INST_POP_IP,
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
			Program: []uint8{
				instruction.INST_MOV_REG_IMM, instruction.REGISTER_OPCODE_IPL, 0x20,
				instruction.INST_MOV_REG_IMM, instruction.REGISTER_OPCODE_IPH, 0x10,
				instruction.INST_ADD_IP_IMM, 0x04,
				instruction.INST_MOV_REG_IMM, instruction.REGISTER_OPCODE_A, 0x10,
				instruction.INST_PUSH_REG, instruction.REGISTER_OPCODE_A,
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
				instruction.INST_MOV_REG_IMM, instruction.REGISTER_OPCODE_IPL, 0xff,
				instruction.INST_MOV_REG_IMM, instruction.REGISTER_OPCODE_IPH, 0x10,
				instruction.INST_ADD_IP_IMM, 0x04,
				instruction.INST_MOV_REG_IMM, instruction.REGISTER_OPCODE_A, 0x10,
				instruction.INST_PUSH_REG, instruction.REGISTER_OPCODE_A,
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
				instruction.INST_MOV_REG_IMM, instruction.REGISTER_OPCODE_IPL, 0x25,
				instruction.INST_MOV_REG_IMM, instruction.REGISTER_OPCODE_IPH, 0x10,
				instruction.INST_SUB_IP_IMM, 0x04,
				instruction.INST_MOV_REG_IMM, instruction.REGISTER_OPCODE_A, 0x10,
				instruction.INST_PUSH_REG, instruction.REGISTER_OPCODE_A,
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
				instruction.INST_MOV_REG_IMM, instruction.REGISTER_OPCODE_IPL, 0x20,
				instruction.INST_MOV_REG_IMM, instruction.REGISTER_OPCODE_IPH, 0x10,
				instruction.INST_SUB_IP_IMM, 0x04,
				instruction.INST_MOV_REG_IMM, instruction.REGISTER_OPCODE_A, 0x10,
				instruction.INST_PUSH_REG, instruction.REGISTER_OPCODE_A,
			},
			Expect: &test.Expect{
				Registers: []*test.ExpectRegister{
					{Index: instruction.REGISTER_OPCODE_IPL, Data: 0x1c},
					{Index: instruction.REGISTER_OPCODE_IPH, Data: 0x10},
				},
			},
		},
		{
			Name: "mov ip reg",
			// mov ipl, 0x20
			// mov iph, 0x80
			// mov b, 0x10
			// mov [ip], b
			// mov a, [ip]
			Program: []uint8{
				instruction.INST_MOV_REG_IMM, instruction.REGISTER_OPCODE_IPL, 0x20,
				instruction.INST_MOV_REG_IMM, instruction.REGISTER_OPCODE_IPH, 0x80,
				instruction.INST_MOV_REG_IMM, instruction.REGISTER_OPCODE_B, 0x10,
				instruction.INST_MOV_REG_IP, instruction.REGISTER_OPCODE_B,
				instruction.INST_MOV_IP_REG, instruction.REGISTER_OPCODE_A,
			},
			Expect: &test.Expect{
				Registers: []*test.ExpectRegister{
					{Index: instruction.REGISTER_OPCODE_A, Data: 0x10},
					{Index: instruction.REGISTER_OPCODE_B, Data: 0x10},
					{Index: instruction.REGISTER_OPCODE_IPL, Data: 0x20},
					{Index: instruction.REGISTER_OPCODE_IPH, Data: 0x80},
				},
			},
		},
		{
			Name: "mov reg ip offset",
			// mov ipl, 0x20
			// mov iph, 0x80
			// mov b, 0x10
			// mov [ip+2], b
			// mov a, [ip+2]
			Program: []uint8{
				instruction.INST_MOV_REG_IMM, instruction.REGISTER_OPCODE_IPL, 0x20,
				instruction.INST_MOV_REG_IMM, instruction.REGISTER_OPCODE_IPH, 0x80,
				instruction.INST_MOV_REG_IMM, instruction.REGISTER_OPCODE_B, 0x10,
				instruction.INST_MOV_REG_IP_OFFSET, instruction.REGISTER_OPCODE_B, 0x2,
				instruction.INST_MOV_IP_REG_OFFSET, instruction.REGISTER_OPCODE_A, 0x2,
			},
			Expect: &test.Expect{
				Registers: []*test.ExpectRegister{
					{Index: instruction.REGISTER_OPCODE_A, Data: 0x10},
					{Index: instruction.REGISTER_OPCODE_B, Data: 0x10},
					{Index: instruction.REGISTER_OPCODE_IPL, Data: 0x20},
					{Index: instruction.REGISTER_OPCODE_IPH, Data: 0x80},
				},
			},
		},
		{
			Name: "mov reg ip offset reg",
			// mov ipl, 0x20
			// mov iph, 0x80
			// mov b, 0x10
			// mov c, 0x5
			// mov [ip+c], b
			// mov a, [ip+c]
			Program: []uint8{
				instruction.INST_MOV_REG_IMM, instruction.REGISTER_OPCODE_IPL, 0x20,
				instruction.INST_MOV_REG_IMM, instruction.REGISTER_OPCODE_IPH, 0x80,
				instruction.INST_MOV_REG_IMM, instruction.REGISTER_OPCODE_B, 0x10,
				instruction.INST_MOV_REG_IMM, instruction.REGISTER_OPCODE_C, 0x5,
				instruction.INST_MOV_REG_IP_OFFSET_REG, instruction.REGISTER_OPCODE_C<<4 | instruction.REGISTER_OPCODE_B,
				instruction.INST_MOV_IP_REG_OFFSET_REG, instruction.REGISTER_OPCODE_C<<4 | instruction.REGISTER_OPCODE_A,
			},
			Expect: &test.Expect{
				Registers: []*test.ExpectRegister{
					{Index: instruction.REGISTER_OPCODE_A, Data: 0x10},
					{Index: instruction.REGISTER_OPCODE_B, Data: 0x10},
					{Index: instruction.REGISTER_OPCODE_C, Data: 0x5},
					{Index: instruction.REGISTER_OPCODE_IPL, Data: 0x20},
					{Index: instruction.REGISTER_OPCODE_IPH, Data: 0x80},
				},
			},
		},
		{
			Name: "add ptr imm",
			// mov a, 4
			// mov [0x8000], a
			// add [0x8000], 5
			// mov b, [0x8000]
			Program: []uint8{
				instruction.INST_MOV_REG_IMM, instruction.REGISTER_OPCODE_A, 0x4,
				instruction.INST_MOV_REG_MEM, instruction.REGISTER_OPCODE_A, 0x00, 0x80,
				instruction.INST_ADD_MEM_IMM, 0x5, 0x00, 0x80,
				instruction.INST_MOV_MEM_REG, instruction.REGISTER_OPCODE_B, 0x00, 0x80,
			},
			Expect: &test.Expect{
				Data: []*test.ExpectData{
					{Type: test.DEV_TYPE_RAM, Addr: 0x8000, Data: 0x9},
				},
				Registers: []*test.ExpectRegister{
					{Index: instruction.REGISTER_OPCODE_A, Data: 0x4},
					{Index: instruction.REGISTER_OPCODE_B, Data: 0x9},
				},
			},
		},
		{
			Name: "adc ptr imm",
			// mov ip, 0x12ab
			Program: []uint8{
				instruction.INST_MOV_IP_IMM_16, 0xab, 0x12,
			},
			Expect: &test.Expect{
				Registers: []*test.ExpectRegister{
					{Index: instruction.REGISTER_OPCODE_IPL, Data: 0xab},
					{Index: instruction.REGISTER_OPCODE_IPH, Data: 0x12},
				},
			},
		},
		{
			Name: "ip imm 16",
			// mov a, 0xff
			Program: []uint8{
				instruction.INST_MOV_REG_IMM, instruction.REGISTER_OPCODE_A, 0xff,
				instruction.INST_MOV_REG_MEM, instruction.REGISTER_OPCODE_A, 0x00, 0x80,
				instruction.INST_ADD_MEM_IMM, 0x5, 0x00, 0x80,
				instruction.INST_ADC_MEM_IMM, 0x0, 0x01, 0x80,
				instruction.INST_MOV_MEM_REG, instruction.REGISTER_OPCODE_B, 0x00, 0x80,
			},
			Expect: &test.Expect{
				Data: []*test.ExpectData{
					{Type: test.DEV_TYPE_RAM, Addr: 0x8000, Data: 0x04},
					{Type: test.DEV_TYPE_RAM, Addr: 0x8001, Data: 0x01},
				},
				Registers: []*test.ExpectRegister{
					{Index: instruction.REGISTER_OPCODE_A, Data: 0xff},
					{Index: instruction.REGISTER_OPCODE_B, Data: 0x4},
				},
			},
		},
		{
			Name: "mov mem imm",
			// mov [0x9000], 0x10
			// mov [0x9001], 0x20
			// mov [0x9002], 0xff
			// mov [0x9003], 0x00
			Program: []uint8{
				instruction.INST_MOV_MEM_IMM, 0x10, 0x00, 0x90,
				instruction.INST_MOV_MEM_IMM, 0x20, 0x01, 0x90,
				instruction.INST_MOV_MEM_IMM, 0xff, 0x02, 0x90,
				instruction.INST_MOV_MEM_IMM, 0x00, 0x03, 0x90,
			},
			Expect: &test.Expect{
				Data: []*test.ExpectData{
					{Type: test.DEV_TYPE_RAM, Addr: 0x9000, Data: 0x10},
					{Type: test.DEV_TYPE_RAM, Addr: 0x9001, Data: 0x20},
					{Type: test.DEV_TYPE_RAM, Addr: 0x9002, Data: 0xff},
					{Type: test.DEV_TYPE_RAM, Addr: 0x9003, Data: 0x00},
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
