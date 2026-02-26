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
			Name: "mov_immediate_reg_a",
			// mov a, 0x10
			// hlt
			Program: []uint8{
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x10,
				instruction.INST_HLT_IMPL,
			},
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
			Name: "mov_immediate_reg_a",
			// mov a, 0xff
			// hlt
			Program: []uint8{
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0xff,
				instruction.INST_HLT_IMPL,
			},
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
			Name: "mov_reg_reg",
			// mov a, 0xff
			// mov b, a
			// hlt
			Program: []uint8{
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x10,
				instruction.INST_MOV_REG8_REG8, instruction.REGISTER_OPCODE_A<<1 | instruction.REGISTER_OPCODE_B,
				instruction.INST_HLT_IMPL,
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
			Name: "mov_reg_reg",
			// mov a, 0xff
			// mov b, a
			// mov c, a
			// mov d, a
			// hlt
			Program: []uint8{
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0xff,
				instruction.INST_MOV_REG8_REG8, instruction.REGISTER_OPCODE_A<<4 | instruction.REGISTER_OPCODE_B,
				instruction.INST_MOV_REG8_REG8, instruction.REGISTER_OPCODE_A<<4 | instruction.REGISTER_OPCODE_C,
				instruction.INST_MOV_REG8_REG8, instruction.REGISTER_OPCODE_A<<4 | instruction.REGISTER_OPCODE_D,
				instruction.INST_HLT_IMPL,
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
			Name: "mov_reg_data",
			// mov a, 0xff
			// mov [0x9000], a
			// hlt
			Program: []uint8{
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0xff,
				instruction.INST_MOV_REG_TO_MEM_DIRECT, instruction.REGISTER_OPCODE_A, 0x00, 0x90,
				instruction.INST_HLT_IMPL,
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
			Name: "mov_reg_data_last_bytes",
			// mov a, 0xff
			// mov [0xffff], a
			// mov [0xfffe], a
			// mov [0xfffd], a
			// hlt
			Program: []uint8{
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0xff,
				instruction.INST_MOV_REG_TO_MEM_DIRECT, instruction.REGISTER_OPCODE_A, 0xff, 0xff,
				instruction.INST_MOV_REG_TO_MEM_DIRECT, instruction.REGISTER_OPCODE_A, 0xfe, 0xff,
				instruction.INST_MOV_REG_TO_MEM_DIRECT, instruction.REGISTER_OPCODE_A, 0xfd, 0xff,
				instruction.INST_HLT_IMPL,
			},
			Expect: &test.Expect{
				Registers: []*test.ExpectRegister{
					{Index: instruction.REGISTER_OPCODE_A, Data: 0xff},
				},
				Data: []*test.ExpectData{
					{Type: test.DEV_TYPE_RAM_2, Addr: 0xffff, Data: 0xff},
					{Type: test.DEV_TYPE_RAM_2, Addr: 0xfffe, Data: 0xff},
					{Type: test.DEV_TYPE_RAM_2, Addr: 0xfffd, Data: 0xff},
				},
			},
		},
		{
			Name: "mov_reg_mem",
			// mov a, 0x90
			// mov [0x9000], a
			// mov a, 0x91
			// mov [0x9001], a
			// mov a, 0x92
			// mov [0x9002], a
			// hlt
			Program: []uint8{
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x90,
				instruction.INST_MOV_REG_TO_MEM_DIRECT, instruction.REGISTER_OPCODE_A, 0x00, 0x90,
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x91,
				instruction.INST_MOV_REG_TO_MEM_DIRECT, instruction.REGISTER_OPCODE_A, 0x01, 0x90,
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x92,
				instruction.INST_MOV_REG_TO_MEM_DIRECT, instruction.REGISTER_OPCODE_A, 0x02, 0x90,
				instruction.INST_HLT_IMPL,
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
			Name: "mov_reg_mem",
			// mov a, 0x90
			// mov [0x9000], a
			// mov b, [0x9000]
			// hlt
			Program: []uint8{
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x90,
				instruction.INST_MOV_REG_TO_MEM_DIRECT, instruction.REGISTER_OPCODE_A, 0x00, 0x90,
				instruction.INST_MOV_MEM_TO_REG_DIRECT, instruction.REGISTER_OPCODE_B, 0x00, 0x90,
				instruction.INST_HLT_IMPL,
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
			Name: "mov_mem_reg_data_offset",
			// mov a, 0x10
			// mov b, 0x2
			// mov [0x9002], a
			// mov c, [0x9000+b]
			// hlt
			Program: []uint8{
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x10,
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_B, 0x2,
				instruction.INST_MOV_REG_TO_MEM_DIRECT, instruction.REGISTER_OPCODE_A, 0x02, 0x90,
				instruction.INST_MOV_MEM_TO_REG_INDEXED, instruction.REGISTER_OPCODE_C<<4 | instruction.REGISTER_OPCODE_B, 0x00, 0x90,
				instruction.INST_HLT_IMPL,
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
			Name: "mov_mem_reg_data_offset",
			// mov a, 0x10
			// mov b, 0x2
			// mov [0x9101], a
			// mov c, [0x90ff+b]
			// hlt
			Program: []uint8{
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x10,
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_B, 0x2,
				instruction.INST_MOV_REG_TO_MEM_DIRECT, instruction.REGISTER_OPCODE_A, 0x01, 0x91,
				instruction.INST_MOV_MEM_TO_REG_INDEXED, instruction.REGISTER_OPCODE_C<<4 | instruction.REGISTER_OPCODE_B, 0xff, 0x90,
				instruction.INST_HLT_IMPL,
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
			Name: "mov_reg_mem_data_offset",
			// mov a, 0x10
			// mov b, 0x2
			// mov [0x90ff+b], a
			// mov c, [0x90ff+b]
			// hlt
			Program: []uint8{
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x10,
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_B, 0x2,
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_C, 0x0,
				instruction.INST_MOV_REG_TO_MEM_INDEXED, instruction.REGISTER_OPCODE_A<<4 | instruction.REGISTER_OPCODE_B, 0xff, 0x90,
				instruction.INST_MOV_MEM_TO_REG_INDEXED, instruction.REGISTER_OPCODE_C<<4 | instruction.REGISTER_OPCODE_B, 0xff, 0x90,
				instruction.INST_HLT_IMPL,
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
			Name: "halt",
			// mov a, 0x10
			// hlt
			// mov a, 0x20
			// hlt
			Program: []uint8{
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x10,
				instruction.INST_HLT_IMPL,
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x20,
				instruction.INST_HLT_IMPL,
			},
			Expect: &test.Expect{
				Registers: []*test.ExpectRegister{
					{Index: instruction.REGISTER_OPCODE_A, Data: 0x10},
				},
			},
		},
		{
			Name: "jump_imm",
			//     mov a, 0x10
			//     jmp here
			//     mov a, 0x20
			//     hlt
			// here:
			//     mov a, 0x30
			// hlt
			Program: []uint8{
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x10,
				instruction.INST_JMP_IMM16, 0x0a, 0x00,
				instruction.INST_HLT_IMPL,
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x20,
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x30,
				instruction.INST_HLT_IMPL,
			},
			Expect: &test.Expect{
				Registers: []*test.ExpectRegister{
					{Index: instruction.REGISTER_OPCODE_A, Data: 0x30},
				},
			},
		},
		{
			Name: "jump_reg16",
			//     mov a, 0x10
			//     mov ipl, 0x0f
			//     mov iph, 0x00
			//     jmp ip
			//     mov a, 0x20
			//     hlt
			// here:
			//     mov a, 0x30
			//     hlt
			Program: []uint8{
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x10,
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_IPL, 0x0f,
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_IPH, 0x00,
				instruction.INST_JMP_REG16, instruction.REGISTER_OPCODE_IP,
				instruction.INST_HLT_IMPL,
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x20,
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x30,
				instruction.INST_HLT_IMPL,
			},
			Expect: &test.Expect{
				Registers: []*test.ExpectRegister{
					{Index: instruction.REGISTER_OPCODE_A, Data: 0x30},
				},
			},
		},
		{
			Name: "cmp_and_jz",
			// start:
			//     mov a, 0x10
			//     cmp a, 0x10
			//     jz here
			//     mov a, 0x20
			//     hlt
			// here:
			//     mov a, 0x30
			//     hlt
			Program: []uint8{
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x10,
				instruction.INST_CMP_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x10,
				instruction.INST_JZ_IMM16, 0x0d, 0x00,
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x20,
				instruction.INST_HLT_IMPL,
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x30,
				instruction.INST_HLT_IMPL,
			},
			Expect: &test.Expect{
				Registers: []*test.ExpectRegister{
					{Index: instruction.REGISTER_OPCODE_A, Data: 0x30},
				},
			},
		},
		{
			Name: "cmp_and_jz_no_jump",
			// start:
			//     mov a, 0x10
			//     cmp a, 0x10
			//     jz here
			//     mov a, 0x20
			//     hlt
			// here:
			//     mov a, 0x30
			// hlt
			Program: []uint8{
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x10,
				instruction.INST_CMP_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x20,
				instruction.INST_JZ_IMM16, 0x0d, 0x00,
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x20,
				instruction.INST_HLT_IMPL,
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x30,
				instruction.INST_HLT_IMPL,
			},
			Expect: &test.Expect{
				Registers: []*test.ExpectRegister{
					{Index: instruction.REGISTER_OPCODE_A, Data: 0x20},
				},
			},
		},
		{
			Name: "jnz_jump",
			// start:
			//     mov a, 0x10
			// 	   cmp a, 0x10
			//     jnz here
			//     mov a, 0x20
			//     hlt
			// here:
			//     mov a, 0x30
			// hlt
			Program: []uint8{
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x10,
				instruction.INST_CMP_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x10,
				instruction.INST_JNZ_IMM16, 0x0d, 0x00,
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x20,
				instruction.INST_HLT_IMPL,
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x30,
				instruction.INST_HLT_IMPL,
			},
			Expect: &test.Expect{
				Registers: []*test.ExpectRegister{
					{Index: instruction.REGISTER_OPCODE_A, Data: 0x20},
				},
			},
		},
		{
			Name: "jnz_no_jump",
			// start:
			//     mov a, 0x20
			// 	   cmp a, 0x10
			//     jnz here
			//     mov a, 0x20
			//     hlt
			// here:
			//     mov a, 0x30
			// hlt
			Program: []uint8{
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x20,
				instruction.INST_CMP_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x10,
				instruction.INST_JNZ_IMM16, 0x0d, 0x00,
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x20,
				instruction.INST_HLT_IMPL,
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x30,
				instruction.INST_HLT_IMPL,
			},
			Expect: &test.Expect{
				Registers: []*test.ExpectRegister{
					{Index: instruction.REGISTER_OPCODE_A, Data: 0x30},
				},
			},
		},
		{
			Name: "cmp_rr_and_jz",
			// start:
			//     mov a, 0x10
			//     mov b, 0x10
			//     cmp a, b
			//     jz here
			//     mov a, 0x20
			//     hlt
			// here:
			//     mov a, 0x30
			// hlt
			Program: []uint8{
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x10,
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_B, 0x10,
				instruction.INST_CMP_REG8_REG8, instruction.REGISTER_OPCODE_A<<4 | instruction.REGISTER_OPCODE_B,
				instruction.INST_JZ_IMM16, 0x0f, 0x00,
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x20,
				instruction.INST_HLT_IMPL,
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x30,
				instruction.INST_HLT_IMPL,
			},
			Expect: &test.Expect{
				Registers: []*test.ExpectRegister{
					{Index: instruction.REGISTER_OPCODE_A, Data: 0x30},
				},
			},
		},
		{
			Name: "cmp_rr_and_jz_no_jump",
			// start:
			//     mov a, 0x10
			//     mov b, 0x20
			//     cmp a, b
			//     jz here
			//     mov a, 0x20
			//     hlt
			// here:
			//     mov a, 0x30
			//     hlt
			Program: []uint8{
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x10,
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_B, 0x20,
				instruction.INST_CMP_REG8_REG8, instruction.REGISTER_OPCODE_A<<4 | instruction.REGISTER_OPCODE_B,
				instruction.INST_JZ_IMM16, 0x0f, 0x00,
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x20,
				instruction.INST_HLT_IMPL,
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x30,
				instruction.INST_HLT_IMPL,
			},
			Expect: &test.Expect{
				Registers: []*test.ExpectRegister{
					{Index: instruction.REGISTER_OPCODE_A, Data: 0x20},
				},
			},
		},
		{
			Name: "add_reg_imm",
			// add a, 0x10
			// hlt
			Program: []uint8{
				instruction.INST_ADD_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x10,
				instruction.INST_HLT_IMPL,
			},
			Expect: &test.Expect{
				Registers: []*test.ExpectRegister{
					{Index: instruction.REGISTER_OPCODE_A, Data: 0x10},
				},
			},
		},
		{
			Name: "adc_reg_imm",
			// mov a, 0xff
			// mov b, 0x10
			// add a, b
			// adc c, 0x10
			// hlt
			Program: []uint8{
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0xff,
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_B, 0x10,
				instruction.INST_ADD_REG8_REG8, instruction.REGISTER_OPCODE_B<<4 | instruction.REGISTER_OPCODE_A,
				instruction.INST_ADC_REG8_IMM8, instruction.REGISTER_OPCODE_C, 0x10,
				instruction.INST_HLT_IMPL,
			},
			Expect: &test.Expect{
				Registers: []*test.ExpectRegister{
					{Index: instruction.REGISTER_OPCODE_C, Data: 0x11},
				},
			},
		},
		{
			Name: "add_reg_imm",
			// mov a, 0x20
			// add a, 0x10
			// hlt
			Program: []uint8{
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x20,
				instruction.INST_ADD_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x10,
				instruction.INST_HLT_IMPL,
			},
			Expect: &test.Expect{
				Registers: []*test.ExpectRegister{
					{Index: instruction.REGISTER_OPCODE_A, Data: 0x30},
				},
			},
		},
		{
			Name: "add_reg_imm",
			// mov a, 0xff
			// add a, 0x01
			// hlt
			Program: []uint8{
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0xff,
				instruction.INST_ADD_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x01,
				instruction.INST_HLT_IMPL,
			},
			Expect: &test.Expect{
				Registers: []*test.ExpectRegister{
					{Index: instruction.REGISTER_OPCODE_A, Data: 0x0},
				},
			},
		},
		{
			Name: "add_reg_reg",
			// mov a, 0xff
			// mov b, 0x11
			// add a, b
			// hlt
			Program: []uint8{
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0xff,
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_B, 0x11,
				instruction.INST_ADD_REG8_REG8, instruction.REGISTER_OPCODE_B<<4 | instruction.REGISTER_OPCODE_A,
				instruction.INST_HLT_IMPL,
			},
			Expect: &test.Expect{
				Registers: []*test.ExpectRegister{
					{Index: instruction.REGISTER_OPCODE_A, Data: 0x10},
					{Index: instruction.REGISTER_OPCODE_B, Data: 0x11},
				},
			},
		},
		{
			Name: "add_reg_reg",
			// mov a, 0xff
			// mov b, 0x11
			// add b, a
			// hlt
			Program: []uint8{
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0xff,
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_B, 0x11,
				instruction.INST_ADD_REG8_REG8, instruction.REGISTER_OPCODE_A<<4 | instruction.REGISTER_OPCODE_B,
				instruction.INST_HLT_IMPL,
			},
			Expect: &test.Expect{
				Registers: []*test.ExpectRegister{
					{Index: instruction.REGISTER_OPCODE_A, Data: 0xff},
					{Index: instruction.REGISTER_OPCODE_B, Data: 0x10},
				},
			},
		},
		{
			Name: "sub_reg_imm",
			// sub a, 0x10
			// hlt
			Program: []uint8{
				instruction.INST_SUB_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x10,
				instruction.INST_HLT_IMPL,
			},
			Expect: &test.Expect{
				Registers: []*test.ExpectRegister{
					{Index: instruction.REGISTER_OPCODE_A, Data: 0xf0},
				},
			},
		},
		{
			Name: "sub_reg_imm",
			// mov a, 0x20
			// sub a, 0x10
			// hlt
			Program: []uint8{
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x20,
				instruction.INST_SUB_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x10,
				instruction.INST_HLT_IMPL,
			},
			Expect: &test.Expect{
				Registers: []*test.ExpectRegister{
					{Index: instruction.REGISTER_OPCODE_A, Data: 0x10},
				},
			},
		},
		{
			Name: "sub_reg_imm",
			// mov a, 0xff
			// sub a, 0x01
			// hlt
			Program: []uint8{
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0xff,
				instruction.INST_SUB_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x01,
				instruction.INST_HLT_IMPL,
			},
			Expect: &test.Expect{
				Registers: []*test.ExpectRegister{
					{Index: instruction.REGISTER_OPCODE_A, Data: 0xfe},
				},
			},
		},
		{
			Name: "sub_reg_reg",
			// mov a, 0xff
			// mov b, 0x11
			// sub a, b
			// hlt
			Program: []uint8{
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0xff,
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_B, 0x11,
				instruction.INST_SUB_REG8_REG8, instruction.REGISTER_OPCODE_B<<4 | instruction.REGISTER_OPCODE_A,
				instruction.INST_HLT_IMPL,
			},
			Expect: &test.Expect{
				Registers: []*test.ExpectRegister{
					{Index: instruction.REGISTER_OPCODE_A, Data: 0xee},
					{Index: instruction.REGISTER_OPCODE_B, Data: 0x11},
				},
			},
		},
		{
			Name: "sub_reg_reg",
			// mov a, 0xff
			// mov b, 0x11
			// sub b, a
			// hlt
			Program: []uint8{
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0xff,
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_B, 0x11,
				instruction.INST_SUB_REG8_REG8, instruction.REGISTER_OPCODE_A<<4 | instruction.REGISTER_OPCODE_B,
				instruction.INST_HLT_IMPL,
			},
			Expect: &test.Expect{
				Registers: []*test.ExpectRegister{
					{Index: instruction.REGISTER_OPCODE_A, Data: 0xff},
					{Index: instruction.REGISTER_OPCODE_B, Data: 0x12},
				},
			},
		},
		{
			Name: "push_pop_reg",
			// mov a, 0x10
			// push a
			// pop b
			// hlt
			Program: []uint8{
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x10,
				instruction.INST_PUSH_REG8, instruction.REGISTER_OPCODE_A,
				instruction.INST_POP_REG8, instruction.REGISTER_OPCODE_B,
				instruction.INST_HLT_IMPL,
			},
			Expect: &test.Expect{
				Registers: []*test.ExpectRegister{
					{Index: instruction.REGISTER_OPCODE_A, Data: 0x10},
					{Index: instruction.REGISTER_OPCODE_B, Data: 0x10},
				},
			},
		},
		{
			Name: "inc_reg",
			// mov a, 0x20
			// inc a
			// hlt
			Program: []uint8{
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x20,
				instruction.INST_INC_REG8, instruction.REGISTER_OPCODE_A,
				instruction.INST_HLT_IMPL,
			},
			Expect: &test.Expect{
				Registers: []*test.ExpectRegister{
					{Index: instruction.REGISTER_OPCODE_A, Data: 0x21},
				},
			},
		},
		{
			Name: "inc_reg",
			// mov a, 0xff
			// inc a
			// hlt
			Program: []uint8{
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0xff,
				instruction.INST_INC_REG8, instruction.REGISTER_OPCODE_A,
				instruction.INST_HLT_IMPL,
			},
			Expect: &test.Expect{
				Registers: []*test.ExpectRegister{
					{Index: instruction.REGISTER_OPCODE_A, Data: 0x00},
				},
			},
		},
		{
			Name: "dec_reg",
			// mov a, 0x20
			// dec a
			// hlt
			Program: []uint8{
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x20,
				instruction.INST_DEC_REG8, instruction.REGISTER_OPCODE_A,
				instruction.INST_HLT_IMPL,
			},
			Expect: &test.Expect{
				Registers: []*test.ExpectRegister{
					{Index: instruction.REGISTER_OPCODE_A, Data: 0x1f},
				},
			},
		},
		{
			Name: "dec_reg",
			// mov a, 0x00
			// dec a
			// hlt
			Program: []uint8{
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x00,
				instruction.INST_DEC_REG8, instruction.REGISTER_OPCODE_A,
				instruction.INST_HLT_IMPL,
			},
			Expect: &test.Expect{
				Registers: []*test.ExpectRegister{
					{Index: instruction.REGISTER_OPCODE_A, Data: 0xff},
				},
			},
		},
		{
			Name: "call_and_ret",
			// start:
			//     call hello
			//     mov a, 0x10
			//     hlt
			// hello:
			//     mov a, 0x20
			//     mov b, 0x20
			//     ret
			//     hlt
			Program: []uint8{
				instruction.INST_CALL_IMM16, 0x07, 0x00,
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x10,
				instruction.INST_HLT_IMPL,
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x20,
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_B, 0x20,
				instruction.INST_RET_IMPL,
				instruction.INST_HLT_IMPL,
			},
			Expect: &test.Expect{
				Registers: []*test.ExpectRegister{
					{Index: instruction.REGISTER_OPCODE_A, Data: 0x10},
					{Index: instruction.REGISTER_OPCODE_B, Data: 0x20},
				},
			},
		},
		{
			Name: "call_reg16_and_ret",
			// start:
			//     mov ipl, 0x0c
			//     mov iph, 0x00
			//     call ip
			//     mov a, 0x10
			//     hlt
			//     mov a, 0x20
			//     mov b, 0x20
			//     ret
			//     hlt
			Program: []uint8{
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_IPL, 0x0c,
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_IPH, 0x00,
				instruction.INST_CALL_REG16, instruction.REGISTER_OPCODE_IP,
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x10,
				instruction.INST_HLT_IMPL,
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x20,
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_B, 0x20,
				instruction.INST_RET_IMPL,
				instruction.INST_HLT_IMPL,
			},
			Expect: &test.Expect{
				Registers: []*test.ExpectRegister{
					{Index: instruction.REGISTER_OPCODE_A, Data: 0x10},
					{Index: instruction.REGISTER_OPCODE_B, Data: 0x20},
				},
			},
		},
		{
			Name: "xor_reg_reg",
			// mov a, 0b00001111
			// mov b, 0b11001100
			// xor b, a
			// hlt
			Program: []uint8{
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0b00001111,
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_B, 0b11001100,
				instruction.INST_XOR_REG8_REG8, instruction.REGISTER_OPCODE_A<<4 | instruction.REGISTER_OPCODE_B,
				instruction.INST_HLT_IMPL,
			},
			Expect: &test.Expect{
				Registers: []*test.ExpectRegister{
					{Index: instruction.REGISTER_OPCODE_A, Data: 0b00001111},
					{Index: instruction.REGISTER_OPCODE_B, Data: 0b11000011},
				},
			},
		},
		{
			Name: "xor_reg_reg_clear",
			// mov a, 0b00001111
			// xor a, a
			// hlt
			Program: []uint8{
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0b00001111,
				instruction.INST_XOR_REG8_REG8, instruction.REGISTER_OPCODE_A<<4 | instruction.REGISTER_OPCODE_A,
				instruction.INST_HLT_IMPL,
			},
			Expect: &test.Expect{
				Registers: []*test.ExpectRegister{
					{Index: instruction.REGISTER_OPCODE_A, Data: 0},
				},
			},
		},
		{
			Name: "xor_reg_imm",
			// mov a, 0b11001100
			// xor a, 0b00001111
			// hlt
			Program: []uint8{
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0b11001100,
				instruction.INST_XOR_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0b00001111,
				instruction.INST_HLT_IMPL,
			},
			Expect: &test.Expect{
				Registers: []*test.ExpectRegister{
					{Index: instruction.REGISTER_OPCODE_A, Data: 0b11000011},
				},
			},
		},
		{
			Name: "and_reg_reg",
			// mov a, 0b00001111
			// mov b, 0b11001100
			// and b, a
			// hlt
			Program: []uint8{
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0b00001111,
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_B, 0b11001100,
				instruction.INST_AND_REG8_REG8, instruction.REGISTER_OPCODE_A<<4 | instruction.REGISTER_OPCODE_B,
				instruction.INST_HLT_IMPL,
			},
			Expect: &test.Expect{
				Registers: []*test.ExpectRegister{
					{Index: instruction.REGISTER_OPCODE_A, Data: 0b00001111},
					{Index: instruction.REGISTER_OPCODE_B, Data: 0b00001100},
				},
			},
		},
		{
			Name: "and_reg_imm",
			// mov a, 0b11001100
			// and a, 0b00001111
			// hlt
			Program: []uint8{
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0b11001100,
				instruction.INST_AND_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0b00001111,
				instruction.INST_HLT_IMPL,
			},
			Expect: &test.Expect{
				Registers: []*test.ExpectRegister{
					{Index: instruction.REGISTER_OPCODE_A, Data: 0b00001100},
				},
			},
		},
		{
			Name: "or_reg_reg",
			// mov a, 0b00001111
			// mov b, 0b11001100
			// or b, a
			// hlt
			Program: []uint8{
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0b00001111,
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_B, 0b11001100,
				instruction.INST_OR_REG8_REG8, instruction.REGISTER_OPCODE_A<<4 | instruction.REGISTER_OPCODE_B,
				instruction.INST_HLT_IMPL,
			},
			Expect: &test.Expect{
				Registers: []*test.ExpectRegister{
					{Index: instruction.REGISTER_OPCODE_A, Data: 0b00001111},
					{Index: instruction.REGISTER_OPCODE_B, Data: 0b11001111},
				},
			},
		},
		{
			Name: "or_reg_imm",
			// mov a, 0b11001100
			// or a, 0b00001111
			// hlt
			Program: []uint8{
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0b11001100,
				instruction.INST_OR_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0b00001111,
				instruction.INST_HLT_IMPL,
			},
			Expect: &test.Expect{
				Registers: []*test.ExpectRegister{
					{Index: instruction.REGISTER_OPCODE_A, Data: 0b11001111},
				},
			},
		},
		{
			Name: "shl_reg_one",
			// mov a, 0b11001100
			// shl a,
			// hlt
			Program: []uint8{
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0b11001100,
				instruction.INST_SHL_REG8, instruction.REGISTER_OPCODE_A,
				instruction.INST_HLT_IMPL,
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
			Name: "shl_reg_zero",
			// mov a, 0b00110011
			// shl a,
			// hlt
			Program: []uint8{
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0b00110011,
				instruction.INST_SHL_REG8, instruction.REGISTER_OPCODE_A,
				instruction.INST_HLT_IMPL,
			},
			Expect: &test.Expect{
				Registers: []*test.ExpectRegister{
					{Index: instruction.REGISTER_OPCODE_A, Data: 0b01100110},
				},
				Status: &test.ExpectStatusData{
					Data: 0,
				},
			},
		},
		{
			Name: "shr_reg_zero",
			// mov a, 0b11001100
			// shr a,
			// hlt
			Program: []uint8{
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0b11001100,
				instruction.INST_SHR_REG8, instruction.REGISTER_OPCODE_A,
				instruction.INST_HLT_IMPL,
			},
			Expect: &test.Expect{
				Registers: []*test.ExpectRegister{
					{Index: instruction.REGISTER_OPCODE_A, Data: 0b01100110},
				},
				Status: &test.ExpectStatusData{
					Data: 0,
				},
			},
		},
		{
			Name: "shr_reg_one",
			// mov a, 0b00110011
			// shr a,
			// hlt
			Program: []uint8{
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0b00110011,
				instruction.INST_SHR_REG8, instruction.REGISTER_OPCODE_A,
				instruction.INST_HLT_IMPL,
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
			Name: "ror_no_carry",
			// mov a, 0b11001100
			// ror a,
			// hlt
			Program: []uint8{
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0b11001100,
				instruction.INST_ROR_REG8, instruction.REGISTER_OPCODE_A,
				instruction.INST_HLT_IMPL,
			},
			Expect: &test.Expect{
				Registers: []*test.ExpectRegister{
					{Index: instruction.REGISTER_OPCODE_A, Data: 0b01100110},
				},
				Status: &test.ExpectStatusData{
					Data: 0,
				},
			},
		},
		{
			Name: "ror_zero_carry",
			// mov a, 0b00110011
			// ror a,
			// hlt
			Program: []uint8{
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0b00110011,
				instruction.INST_ROR_REG8, instruction.REGISTER_OPCODE_A,
				instruction.INST_HLT_IMPL,
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
			Name: "rol_no_carry",
			// mov a, 0b11001100
			// rol a,
			// hlt
			Program: []uint8{
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0b11001100,
				instruction.INST_ROL_REG8, instruction.REGISTER_OPCODE_A,
				instruction.INST_HLT_IMPL,
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
			Name: "rol_zero_parity",
			// mov a, 0b00110011
			// rol a,
			// hlt
			Program: []uint8{
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0b00110011,
				instruction.INST_ROL_REG8, instruction.REGISTER_OPCODE_A,
				instruction.INST_HLT_IMPL,
			},
			Expect: &test.Expect{
				Registers: []*test.ExpectRegister{
					{Index: instruction.REGISTER_OPCODE_A, Data: 0b01100110},
				},
				Status: &test.ExpectStatusData{
					Data: 0,
				},
			},
		},
		{
			Name: "jc",
			// start:
			// 	   mov a, 0xff
			// 	   mov b, 0x02
			//     add a, b
			//     jc here
			//     mov c, 0x20
			//     hlt
			// here:
			//     mov c, 0x30
			// hlt
			Program: []uint8{
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0xff,
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_B, 0x02,
				instruction.INST_ADD_REG8_REG8, instruction.REGISTER_OPCODE_B<<4 | instruction.REGISTER_OPCODE_A,
				instruction.INST_JC_IMM16, 0x0f, 0x00,
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_C, 0x20,
				instruction.INST_HLT_IMPL,
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_C, 0x30,
				instruction.INST_HLT_IMPL,
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
			Name: "jc_no_jump",
			// start:
			// 	   mov a, 0x01
			// 	   mov b, 0x02
			//     add a, b
			//     jc here
			//     mov c, 0x20
			//     hlt
			// here:
			//     mov c, 0x30
			// hlt
			Program: []uint8{
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x01,
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_B, 0x02,
				instruction.INST_ADD_REG8_REG8, instruction.REGISTER_OPCODE_B<<4 | instruction.REGISTER_OPCODE_A,
				instruction.INST_JC_IMM16, 0x0f, 0x00,
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_C, 0x20,
				instruction.INST_HLT_IMPL,
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_C, 0x30,
				instruction.INST_HLT_IMPL,
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
			Name: "jc_with_clear_carry",
			// start:
			// 	   mov a, 0xff
			// 	   mov b, 0x02
			//     add a, b
			//     clc
			//     jc here
			//     mov c, 0x20
			//     hlt
			// here:
			//     mov c, 0x30
			// hlt
			Program: []uint8{
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0xff,
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_B, 0x02,
				instruction.INST_ADD_REG8_REG8, instruction.REGISTER_OPCODE_B<<4 | instruction.REGISTER_OPCODE_A,
				instruction.INST_CLC_IMPL,
				instruction.INST_JC_IMM16, 0x10, 0x00,
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_C, 0x20,
				instruction.INST_HLT_IMPL,
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_C, 0x30,
				instruction.INST_HLT_IMPL,
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
			Name: "cmp_and_js",
			// start:
			//     mov a, 0x10
			//     cmp a, 0x20
			//     js here
			//     mov a, 0x20
			//     hlt
			// here:
			//     mov a, 0x30
			// hlt
			Program: []uint8{
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x10,
				instruction.INST_CMP_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x20,
				instruction.INST_JS_IMM16, 0x0d, 0x00,
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x20,
				instruction.INST_HLT_IMPL,
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x30,
				instruction.INST_HLT_IMPL,
			},
			Expect: &test.Expect{
				Registers: []*test.ExpectRegister{
					{Index: instruction.REGISTER_OPCODE_A, Data: 0x30},
				},
			},
		},
		{
			Name: "cmp_and_js_no_jump",
			// start:
			//     mov a, 0x20
			//     cmp a, 0x10
			//     js here
			//     mov a, 0x20
			//     hlt
			// here:
			//     mov a, 0x30
			// hlt
			Program: []uint8{
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x20,
				instruction.INST_CMP_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x10,
				instruction.INST_JS_IMM16, 0x0d, 0x00,
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x20,
				instruction.INST_HLT_IMPL,
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x30,
				instruction.INST_HLT_IMPL,
			},
			Expect: &test.Expect{
				Registers: []*test.ExpectRegister{
					{Index: instruction.REGISTER_OPCODE_A, Data: 0x20},
				},
			},
		},
		{
			Name: "push_pop_sp",
			// mov a, 0x20
			// push a
			// push sp
			// hlt
			Program: []uint8{
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x20,
				instruction.INST_PUSH_REG8, instruction.REGISTER_OPCODE_A,
				instruction.INST_PUSH_REG16, instruction.REGISTER_OPCODE_SP,
				instruction.INST_HLT_IMPL,
			},
			Expect: &test.Expect{
				Data: []*test.ExpectData{
					{Type: test.DEV_TYPE_RAM, Addr: TEST_STACK_START - 1, Data: 0xfe},
					{Type: test.DEV_TYPE_RAM, Addr: TEST_STACK_START - 2, Data: 0x20},
				},
			},
		},
		{
			Name: "add_sp",
			// add sp, 4
			// mov a, 0x10
			// push a
			// hlt
			Program: []uint8{
				instruction.INST_ADD_REG16_IMM8, instruction.REGISTER_OPCODE_SP, 0x04,
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x10,
				instruction.INST_PUSH_REG8, instruction.REGISTER_OPCODE_A,
				instruction.INST_HLT_IMPL,
			},
			Expect: &test.Expect{
				Data: []*test.ExpectData{
					{Type: test.DEV_TYPE_RAM, Addr: TEST_STACK_START + 4, Data: 0x10},
				},
			},
		},
		{
			Name: "sub_sp",
			// sub sp, 4
			// mov a, 0x10
			// push a
			// hlt
			Program: []uint8{
				instruction.INST_SUB_REG16_IMM8, instruction.REGISTER_OPCODE_SP, 0x04,
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x10,
				instruction.INST_PUSH_REG8, instruction.REGISTER_OPCODE_A,
				instruction.INST_HLT_IMPL,
			},
			Expect: &test.Expect{
				Data: []*test.ExpectData{
					{Type: test.DEV_TYPE_RAM, Addr: TEST_STACK_START - 4, Data: 0x10},
				},
			},
		},
		{
			Name: "mov_ipl_iph_imm",
			// mov ipl, 0x10
			// mov iph, 0x20
			// push ip
			// mov ipl, 0x11
			// mov iph, 0x21
			// pop ip
			// hlt
			Program: []uint8{
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_IPL, 0x20,
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_IPH, 0x10,
				instruction.INST_PUSH_REG16, instruction.REGISTER_OPCODE_IP,
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_IPL, 0x21,
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_IPH, 0x12,
				instruction.INST_POP_REG16, instruction.REGISTER_OPCODE_IP,
				instruction.INST_HLT_IMPL,
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
			Name: "mov_ipl_iph_reg",
			// mov a, 0x10
			// mov b, 0x20
			// mov ipl, a
			// mov iph, b
			// push ip
			// hlt
			Program: []uint8{
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x10,
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_B, 0x20,
				instruction.INST_MOV_REG8_REG8, instruction.REGISTER_OPCODE_A<<4 | instruction.REGISTER_OPCODE_IPL,
				instruction.INST_MOV_REG8_REG8, instruction.REGISTER_OPCODE_B<<4 | instruction.REGISTER_OPCODE_IPH,
				instruction.INST_PUSH_REG16, instruction.REGISTER_OPCODE_IP,
				instruction.INST_HLT_IMPL,
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
			Name: "push_ip",
			// mov ipl, 0x30
			// mov iph, 0x20
			// push ip
			// hlt
			Program: []uint8{
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_IPL, 0x30,
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_IPH, 0x20,
				instruction.INST_PUSH_REG16, instruction.REGISTER_OPCODE_IP,
				instruction.INST_HLT_IMPL,
			},
			Expect: &test.Expect{
				Data: []*test.ExpectData{
					{Type: test.DEV_TYPE_RAM, Addr: TEST_STACK_START, Data: 0x30},
					{Type: test.DEV_TYPE_RAM, Addr: TEST_STACK_START - 1, Data: 0x20},
				},
			},
		},
		{
			Name: "push_pop_ip",
			// mov ipl, 0x30
			// mov iph, 0x20
			// push ip
			// mov ip, 0x4000
			// pop ip
			// hlt
			Program: []uint8{
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_IPL, 0x30,
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_IPH, 0x20,
				instruction.INST_PUSH_REG16, instruction.REGISTER_OPCODE_IP,
				instruction.INST_MOV_REG16_IMM16, instruction.REGISTER_OPCODE_IP, 0x00, 0x40,
				instruction.INST_POP_REG16, instruction.REGISTER_OPCODE_IP,
				instruction.INST_HLT_IMPL,
			},
			Expect: &test.Expect{
				Registers: []*test.ExpectRegister{
					{Index: instruction.REGISTER_OPCODE_IPL, Data: 0x30},
					{Index: instruction.REGISTER_OPCODE_IPH, Data: 0x20},
				},
			},
		},
		{
			Name: "add_ip",
			// mov ipl, 0xff
			// mov iph, 0x10
			// add ip, 4
			// hlt
			Program: []uint8{
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_IPL, 0xff,
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_IPH, 0x10,
				instruction.INST_ADD_REG16_IMM8, instruction.REGISTER_OPCODE_IP, 0x04,
				instruction.INST_HLT_IMPL,
			},
			Expect: &test.Expect{
				Registers: []*test.ExpectRegister{
					{Index: instruction.REGISTER_OPCODE_IPL, Data: 0x03},
					{Index: instruction.REGISTER_OPCODE_IPH, Data: 0x11},
				},
			},
		},
		{
			Name: "add_ip_carry",
			// mov ipl, 0xff
			// mov iph, 0x10
			// add ip, 4
			// hlt
			Program: []uint8{
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_IPL, 0xff,
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_IPH, 0x10,
				instruction.INST_ADD_REG16_IMM8, instruction.REGISTER_OPCODE_IP, 0x04,
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x10,
				instruction.INST_PUSH_REG8, instruction.REGISTER_OPCODE_A,
				instruction.INST_HLT_IMPL,
			},
			Expect: &test.Expect{
				Registers: []*test.ExpectRegister{
					{Index: instruction.REGISTER_OPCODE_IPL, Data: 0x03},
					{Index: instruction.REGISTER_OPCODE_IPH, Data: 0x11},
				},
			},
		},
		{
			Name: "sub_ip",
			// mov ipl, 0x25
			// mov iph, 0x10
			// sub ip, 4
			// hlt
			Program: []uint8{
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_IPL, 0x25,
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_IPH, 0x10,
				instruction.INST_SUB_REG16_IMM8, instruction.REGISTER_OPCODE_IP, 0x04,
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x10,
				instruction.INST_PUSH_REG8, instruction.REGISTER_OPCODE_A,
				instruction.INST_HLT_IMPL,
			},
			Expect: &test.Expect{
				Registers: []*test.ExpectRegister{
					{Index: instruction.REGISTER_OPCODE_IPL, Data: 0x21},
					{Index: instruction.REGISTER_OPCODE_IPH, Data: 0x10},
				},
			},
		},
		{
			Name: "sub_ip_borrow",
			// mov ipl, 0x20
			// mov iph, 0x10
			// sub ip, 4
			// hlt
			Program: []uint8{
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_IPL, 0x20,
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_IPH, 0x10,
				instruction.INST_SUB_REG16_IMM8, instruction.REGISTER_OPCODE_IP, 0x04,
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x10,
				instruction.INST_PUSH_REG8, instruction.REGISTER_OPCODE_A,
				instruction.INST_HLT_IMPL,
			},
			Expect: &test.Expect{
				Registers: []*test.ExpectRegister{
					{Index: instruction.REGISTER_OPCODE_IPL, Data: 0x1c},
					{Index: instruction.REGISTER_OPCODE_IPH, Data: 0x10},
				},
			},
		},
		{
			Name: "mov_reg_imm16",
			// mov ip, 0x1234
			// hlt
			Program: []uint8{
				instruction.INST_MOV_REG16_IMM16, instruction.REGISTER_OPCODE_IP, 0x34, 0x12,
				instruction.INST_HLT_IMPL,
			},
			Expect: &test.Expect{
				Registers: []*test.ExpectRegister{
					{Index: instruction.REGISTER_OPCODE_IPL, Data: 0x34},
					{Index: instruction.REGISTER_OPCODE_IPH, Data: 0x12},
				},
			},
		},
		{
			Name: "mov_addr_reg_imm",
			// mov ipl, 0x34
			// mov bpl, 0x34
			// mov iph, 0x12
			// mov bph, 0x12
			// hlt
			Program: []uint8{
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_IPL, 0x34,
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_BPL, 0x34,
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_IPH, 0x12,
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_BPH, 0x12,
				instruction.INST_HLT_IMPL,
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
			Name: "mov_addr_reg_reg",
			// mov a, 0x34
			// mov ipl, a
			// mov bpl, a
			// mov a, 0x12
			// mov iph, a
			// mov bph, a
			// hlt
			Program: []uint8{
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x34,
				instruction.INST_MOV_REG8_REG8, instruction.REGISTER_OPCODE_A<<1 | instruction.REGISTER_OPCODE_IPL,
				instruction.INST_MOV_REG8_REG8, instruction.REGISTER_OPCODE_A<<1 | instruction.REGISTER_OPCODE_BPL,
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x12,
				instruction.INST_MOV_REG8_REG8, instruction.REGISTER_OPCODE_A<<1 | instruction.REGISTER_OPCODE_IPH,
				instruction.INST_MOV_REG8_REG8, instruction.REGISTER_OPCODE_A<<1 | instruction.REGISTER_OPCODE_BPH,
				instruction.INST_HLT_IMPL,
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
			Name: "mov_addr_reg_addr_reg",
			// mov ip, 0x1234
			// mov bp, ip
			// hlt
			Program: []uint8{
				instruction.INST_MOV_REG16_IMM16, instruction.REGISTER_OPCODE_IP, 0x34, 0x12,
				instruction.INST_MOV_REG16_REG16, instruction.REGISTER_OPCODE_BP<<4 | instruction.REGISTER_OPCODE_IP,
				instruction.INST_HLT_IMPL,
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
			Name: "mov_indirect_reg",
			// mov a, 0x30
			// mov [0x8000], a
			// mov ip, 0x8000
			// mov b, [ip]
			// hlt
			Program: []uint8{
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x30,
				instruction.INST_MOV_REG_TO_MEM_DIRECT, instruction.REGISTER_OPCODE_A, 0x00, 0x80,
				instruction.INST_MOV_REG16_IMM16, instruction.REGISTER_OPCODE_IP, 0x00, 0x80,
				instruction.INST_MOV_MEM_TO_REG_INDIRECT, instruction.REGISTER_OPCODE_B<<4 | instruction.REGISTER_OPCODE_IP,
				instruction.INST_HLT_IMPL,
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
			Name: "mov_indirect_reg_last_byte",
			// mov a, 0x30
			// mov [0xffff], a
			// mov ip, 0xffff
			// mov b, [ip]
			// hlt
			Program: []uint8{
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x30,
				instruction.INST_MOV_REG_TO_MEM_DIRECT, instruction.REGISTER_OPCODE_A, 0xfe, 0xff,
				instruction.INST_MOV_REG16_IMM16, instruction.REGISTER_OPCODE_IP, 0xfe, 0xff,
				instruction.INST_MOV_MEM_TO_REG_INDIRECT, instruction.REGISTER_OPCODE_B<<4 | instruction.REGISTER_OPCODE_IP,
				instruction.INST_HLT_IMPL,
			},
			Expect: &test.Expect{
				Registers: []*test.ExpectRegister{
					{Index: instruction.REGISTER_OPCODE_A, Data: 0x30},
					{Index: instruction.REGISTER_OPCODE_B, Data: 0x30},
				},
				Data: []*test.ExpectData{
					{Type: test.DEV_TYPE_RAM_2, Addr: 0xfffe, Data: 0x30},
				},
			},
		},
		{
			Name: "mov_indirect_offset_reg",
			// mov a, 0x30
			// mov [0x8008], a
			// mov ip, 0x8000
			// mov b, [ip+8]
			// hlt
			Program: []uint8{
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x30,
				instruction.INST_MOV_REG_TO_MEM_DIRECT, instruction.REGISTER_OPCODE_A, 0x08, 0x80,
				instruction.INST_MOV_REG16_IMM16, instruction.REGISTER_OPCODE_IP, 0x00, 0x80,
				instruction.INST_MOV_MEM_TO_REG_INDIRECT_OFFSET, instruction.REGISTER_OPCODE_B<<4 | instruction.REGISTER_OPCODE_IP, 0x08,
				instruction.INST_HLT_IMPL,
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
			Name: "mov_indexed_reg",
			// mov a, 0x30
			// mov b, 0x08
			// mov [0x8008], a
			// mov ip, 0x8000
			// mov c, [ip+b]
			// hlt
			Program: []uint8{
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x30,
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_B, 0x08,
				instruction.INST_MOV_REG_TO_MEM_DIRECT, instruction.REGISTER_OPCODE_A, 0x08, 0x80,
				instruction.INST_MOV_REG16_IMM16, instruction.REGISTER_OPCODE_IP, 0x00, 0x80,
				instruction.INST_MOV_MEM_TO_REG_REG16_INDEXED, instruction.REGISTER_OPCODE_IP<<4 | instruction.REGISTER_OPCODE_B, instruction.REGISTER_OPCODE_C,
				instruction.INST_HLT_IMPL,
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
			Name: "mov_reg_indirect",
			// mov a, 0x30
			// mov ip, 0x8000
			// mov [ip], a
			// hlt
			Program: []uint8{
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x30,
				instruction.INST_MOV_REG16_IMM16, instruction.REGISTER_OPCODE_IP, 0x00, 0x80,
				instruction.INST_MOV_REG_TO_MEM_INDIRECT, instruction.REGISTER_OPCODE_IP<<4 | instruction.REGISTER_OPCODE_A,
				instruction.INST_HLT_IMPL,
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
			Name: "mov_reg_indirect_offset",
			// mov a, 0x30
			// mov bp, 0x8000
			// mov [bp+8], a
			// hlt
			Program: []uint8{
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x30,
				instruction.INST_MOV_REG16_IMM16, instruction.REGISTER_OPCODE_BP, 0x00, 0x80,
				instruction.INST_MOV_REG_TO_MEM_INDIRECT_OFFSET, instruction.REGISTER_OPCODE_BP<<4 | instruction.REGISTER_OPCODE_A, 0x08,
				instruction.INST_HLT_IMPL,
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
			Name: "mov_reg_indexed",
			// mov a, 0x30
			// mov b, 0x08
			// mov bp, 0x8000
			// mov [bp+b], a
			// hlt
			Program: []uint8{
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x30,
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_B, 0x08,
				instruction.INST_MOV_REG16_IMM16, instruction.REGISTER_OPCODE_BP, 0x00, 0x80,
				instruction.INST_MOV_REG_TO_MEM_REG16_INDEXED, instruction.REGISTER_OPCODE_BP<<4 | instruction.REGISTER_OPCODE_B, instruction.REGISTER_OPCODE_A,
				instruction.INST_HLT_IMPL,
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
			Name: "mov_sp_address_reg",
			// mov bp, sp
			// mov ip, sp
			// hlt
			Program: []uint8{
				instruction.INST_MOV_REG16_REG16, instruction.REGISTER_OPCODE_BP<<4 | instruction.REGISTER_OPCODE_SP,
				instruction.INST_MOV_REG16_REG16, instruction.REGISTER_OPCODE_IP<<4 | instruction.REGISTER_OPCODE_SP,
				instruction.INST_HLT_IMPL,
			},
			Expect: &test.Expect{
				Registers: []*test.ExpectRegister{
					{Index: instruction.REGISTER_OPCODE_IPL, Data: uint8(TEST_STACK_START)},
					{Index: instruction.REGISTER_OPCODE_IPH, Data: uint8((TEST_STACK_START) >> 8)},
					{Index: instruction.REGISTER_OPCODE_BPL, Data: uint8(TEST_STACK_START)},
					{Index: instruction.REGISTER_OPCODE_BPH, Data: uint8((TEST_STACK_START) >> 8)},
				},
			},
		},
		{
			Name: "mov_address_reg_sp",
			// mov ip, 0x8000
			// mov bp, 0x8000
			// mov sp, bp
			// mov sp, ip
			// hlt
			Program: []uint8{
				instruction.INST_MOV_REG16_IMM16, instruction.REGISTER_OPCODE_IP, 0x00, 0x80,
				instruction.INST_MOV_REG16_IMM16, instruction.REGISTER_OPCODE_BP, 0x00, 0x80,
				instruction.INST_MOV_REG16_REG16, instruction.REGISTER_OPCODE_SP<<4 | instruction.REGISTER_OPCODE_BP,
				instruction.INST_MOV_REG16_REG16, instruction.REGISTER_OPCODE_SP<<4 | instruction.REGISTER_OPCODE_IP,
				instruction.INST_HLT_IMPL,
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
			Name: "add_address_register",
			// mov ip, 0x80ff
			// add ip, 8
			// hlt
			Program: []uint8{
				instruction.INST_MOV_REG16_IMM16, instruction.REGISTER_OPCODE_IP, 0xff, 0x80,
				instruction.INST_ADD_REG16_IMM8, instruction.REGISTER_OPCODE_IP, 0x08,
				instruction.INST_HLT_IMPL,
			},
			Expect: &test.Expect{
				Registers: []*test.ExpectRegister{
					{Index: instruction.REGISTER_OPCODE_IPL, Data: 0x07},
					{Index: instruction.REGISTER_OPCODE_IPH, Data: 0x81},
				},
			},
		},
		{
			Name: "sub_address_register",
			// mov bp, 0x8107
			// sub bp, 0x08
			// hlt
			Program: []uint8{
				instruction.INST_MOV_REG16_IMM16, instruction.REGISTER_OPCODE_BP, 0x07, 0x81,
				instruction.INST_SUB_REG16_IMM8, instruction.REGISTER_OPCODE_BP, 0x08,
				instruction.INST_HLT_IMPL,
			},
			Expect: &test.Expect{
				Registers: []*test.ExpectRegister{
					{Index: instruction.REGISTER_OPCODE_BPL, Data: 0xff},
					{Index: instruction.REGISTER_OPCODE_BPH, Data: 0x80},
				},
			},
		},
		{
			Name: "inc_address_register",
			// mov bp, 0x80ff
			// inc bp
			// hlt
			Program: []uint8{
				instruction.INST_MOV_REG16_IMM16, instruction.REGISTER_OPCODE_BP, 0xff, 0x80,
				instruction.INST_INC_REG16, instruction.REGISTER_OPCODE_BP,
				instruction.INST_HLT_IMPL,
			},
			Expect: &test.Expect{
				Registers: []*test.ExpectRegister{
					{Index: instruction.REGISTER_OPCODE_BPL, Data: 0x00},
					{Index: instruction.REGISTER_OPCODE_BPH, Data: 0x81},
				},
			},
		},
		{
			Name: "dec_address_register",
			// mov bp, 0x8100
			// dec bp
			// hlt
			Program: []uint8{
				instruction.INST_MOV_REG16_IMM16, instruction.REGISTER_OPCODE_BP, 0x00, 0x81,
				instruction.INST_DEC_REG16, instruction.REGISTER_OPCODE_BP,
				instruction.INST_HLT_IMPL,
			},
			Expect: &test.Expect{
				Registers: []*test.ExpectRegister{
					{Index: instruction.REGISTER_OPCODE_BPL, Data: 0xff},
					{Index: instruction.REGISTER_OPCODE_BPH, Data: 0x80},
				},
			},
		},
		{
			Name: "index_pointer_reconstruction",
			// mov a, 0xff
			// mov [0x3010], a
			// mov ipl, 0x10
			// mov iph, 0x30
			// mov b, [ip]
			// hlt
			Program: []uint8{
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0xff,
				instruction.INST_MOV_REG_TO_MEM_DIRECT, instruction.REGISTER_OPCODE_A, 0x10, 0x30,
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_IPL, 0x10,
				instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_IPH, 0x30,
				instruction.INST_MOV_MEM_TO_REG_INDIRECT, instruction.REGISTER_OPCODE_B<<4 | instruction.REGISTER_OPCODE_IP,
				instruction.INST_HLT_IMPL,
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
		{
			Name: "ip_set_to_last_bytes",
			// mov ip, 0x1020
			// mov [0xfffe], ipl
			// mov [0xffff], iph
			// hlt
			Program: []uint8{
				instruction.INST_MOV_REG16_IMM16, instruction.REGISTER_OPCODE_IP, 0x20, 0x10,
				instruction.INST_MOV_REG_TO_MEM_DIRECT, instruction.REGISTER_OPCODE_IPL, 0xfe, 0xff,
				instruction.INST_MOV_REG_TO_MEM_DIRECT, instruction.REGISTER_OPCODE_IPH, 0xff, 0xff,
				instruction.INST_HLT_IMPL,
			},
			Expect: &test.Expect{
				Data: []*test.ExpectData{
					{Type: test.DEV_TYPE_RAM_2, Addr: 0xfffe, Data: 0x20},
					{Type: test.DEV_TYPE_RAM_2, Addr: 0xffff, Data: 0x10},
				},
				Registers: []*test.ExpectRegister{
					{Index: instruction.REGISTER_OPCODE_IPL, Data: 0x20},
					{Index: instruction.REGISTER_OPCODE_IPH, Data: 0x10},
				},
				Status: &test.ExpectStatusData{
					Data: 0,
				},
			},
		},
		{
			Name: "ip_set_to_last_bytes_2",
			// mov ip, 0x1020
			// mov [0xfffe], iph
			// mov [0xffff], ipl
			// hlt
			Program: []uint8{
				instruction.INST_MOV_REG16_IMM16, instruction.REGISTER_OPCODE_IP, 0x20, 0x10,
				instruction.INST_MOV_REG_TO_MEM_DIRECT, instruction.REGISTER_OPCODE_IPH, 0xfe, 0xff,
				instruction.INST_MOV_REG_TO_MEM_DIRECT, instruction.REGISTER_OPCODE_IPL, 0xff, 0xff,
				instruction.INST_HLT_IMPL,
			},
			Expect: &test.Expect{
				Data: []*test.ExpectData{
					{Type: test.DEV_TYPE_RAM_2, Addr: 0xfffe, Data: 0x10},
					{Type: test.DEV_TYPE_RAM_2, Addr: 0xffff, Data: 0x20},
				},
				Registers: []*test.ExpectRegister{
					{Index: instruction.REGISTER_OPCODE_IPL, Data: 0x20},
					{Index: instruction.REGISTER_OPCODE_IPH, Data: 0x10},
				},
				Status: &test.ExpectStatusData{
					Data: 0,
				},
			},
		},
		{
			Name: "set_interrupt_enable",
			// sti
			// hlt
			Program: []uint8{
				instruction.INST_STI_IMPL,
				instruction.INST_HLT_IMPL,
			},
			Expect: &test.Expect{
				Status: &test.ExpectStatusData{
					Data: status.STATUS_FLAG_INTERRUPT_ENABLE,
				},
			},
		},
		{
			Name: "clear_interrupt_enable",
			// sti
			// hlt
			Program: []uint8{
				instruction.INST_CLI_IMPL,
				instruction.INST_HLT_IMPL,
			},
			Expect: &test.Expect{
				Status: &test.ExpectStatusData{
					Data: 0,
				},
			},
		},
		{
			Name: "set/clear_interrupt_enable",
			// sti
			// cli
			// hlt
			Program: []uint8{
				instruction.INST_STI_IMPL,
				instruction.INST_CLI_IMPL,
				instruction.INST_HLT_IMPL,
			},
			Expect: &test.Expect{
				Status: &test.ExpectStatusData{
					Data: 0,
				},
			},
		},
		// {
		// 	Name: "interrupt_fail_stats_not_set",
		// 	// start:
		// 	//     mov a, 0x10
		// 	//     mov b, 0x11
		// 	//     hlt

		// 	// .org 0x10
		// 	//     mov a, 0x31
		// 	//     mov b, 0x32
		// 	//     rti
		// 	// hlt
		// 	Program: []uint8{
		// 		instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x10,
		// 		instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_B, 0x11,
		// 		instruction.TEST_INST_SET_IRQ,
		// 		0,
		// 		0,
		// 		0,
		// 		0,
		// 		0,
		// 		0,
		// 		0,
		// 		0,
		// 		0,
		// 		instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x31,
		// 		instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_B, 0x32,
		// 		instruction.INST_RTI_IMPL,
		// 		instruction.INST_HLT_IMPL,
		// 	},
		// 	Expect: &test.Expect{
		// 		Registers: []*test.ExpectRegister{
		// 			{Index: instruction.REGISTER_OPCODE_A, Data: 0x10},
		// 			{Index: instruction.REGISTER_OPCODE_B, Data: 0x11},
		// 		},
		// 		Status: &test.ExpectStatusData{
		// 			Data: 0,
		// 		},
		// 	},
		// },
		// {
		// 	Name: "interrupt_success",
		// 	//     sti
		// 	// start:
		// 	//     mov a, 0x10
		// 	//     mov b, 0x11
		// 	//     hlt

		// 	// .org 0x10
		// 	//     mov a, 0x31
		// 	//     mov b, 0x32
		// 	//     rti
		// 	// hlt
		// 	Program: []uint8{
		// 		instruction.INST_STI_IMPL,
		// 		instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x10,
		// 		instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_B, 0x11,
		// 		instruction.TEST_INST_SET_IRQ,
		// 		0,
		// 		0,
		// 		0,
		// 		0,
		// 		0,
		// 		0,
		// 		0,
		// 		0,
		// 		instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x31,
		// 		instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_B, 0x32,
		// 		instruction.INST_RTI_IMPL,
		// 		instruction.INST_HLT_IMPL,
		// 	},
		// 	Expect: &test.Expect{
		// 		Registers: []*test.ExpectRegister{
		// 			{Index: instruction.REGISTER_OPCODE_A, Data: 0x31},
		// 			{Index: instruction.REGISTER_OPCODE_B, Data: 0x32},
		// 		},
		// 		Status: &test.ExpectStatusData{
		// 			Data: status.STATUS_FLAG_INTERRUPT_ENABLE,
		// 		},
		// 	},
		// },
		// {
		// 	Name: "interrupt_success_2",
		// 	//     sti
		// 	// start:
		// 	//     mov a, 0x10
		// 	//     mov b, 0x11
		// 	//     hlt

		// 	// .org 0x10
		// 	//     mov a, 0x31
		// 	//     mov b, 0x32
		// 	//     rti
		// 	// hlt
		// 	Program: []uint8{
		// 		instruction.INST_STI_IMPL,
		// 		instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_A, 0x10,
		// 		instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_B, 0x11,
		// 		instruction.TEST_INST_SET_IRQ,
		// 		0,
		// 		0,
		// 		0,
		// 		0,
		// 		0,
		// 		0,
		// 		0,
		// 		0,
		// 		instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_C, 0x12,
		// 		instruction.INST_MOV_REG8_IMM8, instruction.REGISTER_OPCODE_D, 0x13,
		// 		instruction.INST_RTI_IMPL,
		// 		instruction.INST_HLT_IMPL,
		// 	},
		// 	Expect: &test.Expect{
		// 		Registers: []*test.ExpectRegister{
		// 			{Index: instruction.REGISTER_OPCODE_A, Data: 0x10},
		// 			{Index: instruction.REGISTER_OPCODE_B, Data: 0x11},
		// 			{Index: instruction.REGISTER_OPCODE_C, Data: 0x12},
		// 			{Index: instruction.REGISTER_OPCODE_D, Data: 0x13},
		// 		},
		// 		Status: &test.ExpectStatusData{
		// 			Data: status.STATUS_FLAG_INTERRUPT_ENABLE,
		// 		},
		// },
		// },
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
		if tc.Disable {
			continue
		}
		test.RunCase(t, tc)
	}
}
