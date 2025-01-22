package unit_test

import (
	"asme8/emulator/src/instruction"
	"asme8/emulator/src/register"
	"asme8/emulator/test"
	"testing"
)

var (
	Tests = []*test.TestCase{
		{
			Name: "mov immediate reg a",
			// mov a, 0x10
			Program: []uint8{instruction.INST_MOV_INM, register.IndexRegA, 0x10},
			Expect: &test.Expect{
				Registers: []*test.RegData{
					{Index: register.IndexRegA, Data: 0x10},
					{Index: register.IndexRegB, Data: 0x00},
					{Index: register.IndexRegC, Data: 0x00},
					{Index: register.IndexRegD, Data: 0x00},
				},
			},
		},
		{
			Name: "mov immediate reg a",
			// mov a, 0xff
			Program: []uint8{instruction.INST_MOV_INM, register.IndexRegA, 0xff},
			Expect: &test.Expect{
				Registers: []*test.RegData{
					{Index: register.IndexRegA, Data: 0xff},
					{Index: register.IndexRegB, Data: 0x00},
					{Index: register.IndexRegC, Data: 0x00},
					{Index: register.IndexRegD, Data: 0x00},
				},
			},
		},
		{
			// mov a, 0xff
			// mov b, a
			Name: "mov reg reg",
			Program: []uint8{
				instruction.INST_MOV_INM, register.IndexRegA, 0x10,
				instruction.INST_MOV_REG_REG, register.IndexRegA<<1 | register.IndexRegB,
			},
			Expect: &test.Expect{
				Registers: []*test.RegData{
					{Index: register.IndexRegA, Data: 0x10},
					{Index: register.IndexRegB, Data: 0x10},
					{Index: register.IndexRegC, Data: 0x00},
					{Index: register.IndexRegD, Data: 0x00},
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
				instruction.INST_MOV_INM, register.IndexRegA, 0xff,
				instruction.INST_MOV_REG_REG, register.IndexRegA<<4 | register.IndexRegB,
				instruction.INST_MOV_REG_REG, register.IndexRegA<<4 | register.IndexRegC,
				instruction.INST_MOV_REG_REG, register.IndexRegA<<4 | register.IndexRegD,
			},
			Expect: &test.Expect{
				Registers: []*test.RegData{
					{Index: register.IndexRegA, Data: 0xff},
					{Index: register.IndexRegB, Data: 0xff},
					{Index: register.IndexRegC, Data: 0xff},
					{Index: register.IndexRegD, Data: 0xff},
				},
			},
		},
		{
			// mov a, 0xff
			// mov [0x9000], a
			Name: "mov reg data",
			Program: []uint8{
				instruction.INST_MOV_INM, register.IndexRegA, 0xff,
				instruction.INST_MOV_REG_MEM, register.IndexRegA, 0x00, 0x90,
			},
			Expect: &test.Expect{
				Registers: []*test.RegData{
					{Index: register.IndexRegA, Data: 0xff},
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
				instruction.INST_MOV_INM, register.IndexRegA, 0x90,
				instruction.INST_MOV_REG_MEM, register.IndexRegA, 0x00, 0x90,
				instruction.INST_MOV_INM, register.IndexRegA, 0x91,
				instruction.INST_MOV_REG_MEM, register.IndexRegA, 0x01, 0x90,
				instruction.INST_MOV_INM, register.IndexRegA, 0x92,
				instruction.INST_MOV_REG_MEM, register.IndexRegA, 0x02, 0x90,
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
				instruction.INST_MOV_INM, register.IndexRegA, 0x90,
				instruction.INST_MOV_REG_MEM, register.IndexRegA, 0x00, 0x90,
				instruction.INST_MOV_MEM_REG, register.IndexRegB, 0x00, 0x90,
			},
			Expect: &test.Expect{
				Registers: []*test.RegData{
					{Index: register.IndexRegB, Data: 0x90},
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
				instruction.INST_MOV_INM, register.IndexRegA, 0x10,
				instruction.INST_MOV_INM, register.IndexRegB, 0x2,
				instruction.INST_MOV_REG_MEM, register.IndexRegA, 0x02, 0x90,
				instruction.INST_MOV_MEM_REG_OFFSET, register.IndexRegC<<4 | register.IndexRegB, 0x00, 0x90,
			},
			Expect: &test.Expect{
				Registers: []*test.RegData{
					{Index: register.IndexRegA, Data: 0x10},
					{Index: register.IndexRegB, Data: 0x2},
					{Index: register.IndexRegC, Data: 0x10},
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
				instruction.INST_MOV_INM, register.IndexRegA, 0x10,
				instruction.INST_MOV_INM, register.IndexRegB, 0x2,
				instruction.INST_MOV_REG_MEM, register.IndexRegA, 0x01, 0x91,
				instruction.INST_MOV_MEM_REG_OFFSET, register.IndexRegC<<4 | register.IndexRegB, 0xff, 0x90,
			},
			Expect: &test.Expect{
				Registers: []*test.RegData{
					{Index: register.IndexRegA, Data: 0x10},
					{Index: register.IndexRegB, Data: 0x2},
					{Index: register.IndexRegC, Data: 0x10},
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
				instruction.INST_MOV_INM, register.IndexRegA, 0x10,
				instruction.INST_MOV_INM, register.IndexRegB, 0x2,
				instruction.INST_MOV_INM, register.IndexRegC, 0x0,
				instruction.INST_MOV_REG_MEM_OFFSET, register.IndexRegA<<4 | register.IndexRegB, 0xff, 0x90,
				instruction.INST_MOV_MEM_REG_OFFSET, register.IndexRegC<<4 | register.IndexRegB, 0xff, 0x90,
			},
			Expect: &test.Expect{
				Registers: []*test.RegData{
					{Index: register.IndexRegA, Data: 0x10},
					{Index: register.IndexRegB, Data: 0x2},
					{Index: register.IndexRegC, Data: 0x10},
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
				instruction.INST_MOV_INM, register.IndexRegA, 0x10,
				instruction.INST_BRK,
				instruction.INST_MOV_INM, register.IndexRegA, 0x20,
			},
			Expect: &test.Expect{
				Registers: []*test.RegData{
					{Index: register.IndexRegA, Data: 0x10},
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
			Name: "jump inm",
			Program: []uint8{
				instruction.INST_MOV_INM, register.IndexRegA, 0x10,
				instruction.INST_JMP_INM, 0x0a, 0x00,
				instruction.INST_BRK,
				instruction.INST_MOV_INM, register.IndexRegA, 0x20,
				instruction.INST_MOV_INM, register.IndexRegA, 0x30,
			},
			Expect: &test.Expect{
				Registers: []*test.RegData{
					{Index: register.IndexRegA, Data: 0x30},
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
				instruction.INST_MOV_INM, register.IndexRegA, 0x10,
				instruction.INST_CMP_REG_INM, register.IndexRegA, 0x10,
				instruction.INST_JZ_INM, 0x0d, 0x00,
				instruction.INST_MOV_INM, register.IndexRegA, 0x20,
				instruction.INST_BRK,
				instruction.INST_MOV_INM, register.IndexRegA, 0x30,
			},
			Expect: &test.Expect{
				Registers: []*test.RegData{
					{Index: register.IndexRegA, Data: 0x30},
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
				instruction.INST_MOV_INM, register.IndexRegA, 0x10,
				instruction.INST_CMP_REG_INM, register.IndexRegA, 0x20,
				instruction.INST_JZ_INM, 0x0d, 0x00,
				instruction.INST_MOV_INM, register.IndexRegA, 0x20,
				instruction.INST_BRK,
				instruction.INST_MOV_INM, register.IndexRegA, 0x30,
			},
			Expect: &test.Expect{
				Registers: []*test.RegData{
					{Index: register.IndexRegA, Data: 0x20},
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
				instruction.INST_MOV_INM, register.IndexRegA, 0x10,
				instruction.INST_CMP_REG_INM, register.IndexRegA, 0x10,
				instruction.INST_JNZ_INM, 0x0d, 0x00,
				instruction.INST_MOV_INM, register.IndexRegA, 0x20,
				instruction.INST_BRK,
				instruction.INST_MOV_INM, register.IndexRegA, 0x30,
			},
			Expect: &test.Expect{
				Registers: []*test.RegData{
					{Index: register.IndexRegA, Data: 0x20},
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
				instruction.INST_MOV_INM, register.IndexRegA, 0x20,
				instruction.INST_CMP_REG_INM, register.IndexRegA, 0x10,
				instruction.INST_JNZ_INM, 0x0d, 0x00,
				instruction.INST_MOV_INM, register.IndexRegA, 0x20,
				instruction.INST_BRK,
				instruction.INST_MOV_INM, register.IndexRegA, 0x30,
			},
			Expect: &test.Expect{
				Registers: []*test.RegData{
					{Index: register.IndexRegA, Data: 0x30},
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
				instruction.INST_MOV_INM, register.IndexRegA, 0x10,
				instruction.INST_MOV_INM, register.IndexRegB, 0x10,
				instruction.INST_CMP_RR, register.IndexRegA<<4 | register.IndexRegB,
				instruction.INST_JZ_INM, 0x0f, 0x00,
				instruction.INST_MOV_INM, register.IndexRegA, 0x20,
				instruction.INST_BRK,
				instruction.INST_MOV_INM, register.IndexRegA, 0x30,
			},
			Expect: &test.Expect{
				Registers: []*test.RegData{
					{Index: register.IndexRegA, Data: 0x30},
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
				instruction.INST_MOV_INM, register.IndexRegA, 0x10,
				instruction.INST_MOV_INM, register.IndexRegB, 0x20,
				instruction.INST_CMP_RR, register.IndexRegA<<4 | register.IndexRegB,
				instruction.INST_JZ_INM, 0x0f, 0x00,
				instruction.INST_MOV_INM, register.IndexRegA, 0x20,
				instruction.INST_BRK,
				instruction.INST_MOV_INM, register.IndexRegA, 0x30,
			},
			Expect: &test.Expect{
				Registers: []*test.RegData{
					{Index: register.IndexRegA, Data: 0x20},
				},
			},
		},
		{
			Name: "add reg inm",
			// add a, 0x10
			Program: []uint8{instruction.INST_ADD_REG_INM, register.IndexRegA, 0x10},
			Expect: &test.Expect{
				Registers: []*test.RegData{
					{Index: register.IndexRegA, Data: 0x10},
				},
			},
		},
		{
			Name: "add reg inm",
			// mov a, 0x20
			// add a, 0x10
			Program: []uint8{
				instruction.INST_MOV_INM, register.IndexRegA, 0x20,
				instruction.INST_ADD_REG_INM, register.IndexRegA, 0x10,
			},
			Expect: &test.Expect{
				Registers: []*test.RegData{
					{Index: register.IndexRegA, Data: 0x30},
				},
			},
		},
		{
			Name: "add reg inm",
			// mov a, 0xff
			// add a, 0x01
			Program: []uint8{
				instruction.INST_MOV_INM, register.IndexRegA, 0xff,
				instruction.INST_ADD_REG_INM, register.IndexRegA, 0x01,
			},
			Expect: &test.Expect{
				Registers: []*test.RegData{
					{Index: register.IndexRegA, Data: 0x0},
				},
			},
		},
		{
			Name: "add reg reg",
			// mov a, 0xff
			// mov b, 0x11
			// add a, b
			Program: []uint8{
				instruction.INST_MOV_INM, register.IndexRegA, 0xff,
				instruction.INST_MOV_INM, register.IndexRegB, 0x11,
				instruction.INST_ADD_RR, register.IndexRegB<<4 | register.IndexRegA,
			},
			Expect: &test.Expect{
				Registers: []*test.RegData{
					{Index: register.IndexRegA, Data: 0x10},
					{Index: register.IndexRegB, Data: 0x11},
				},
			},
		},
		{
			Name: "add reg reg",
			// mov a, 0xff
			// mov b, 0x11
			// add b, a
			Program: []uint8{
				instruction.INST_MOV_INM, register.IndexRegA, 0xff,
				instruction.INST_MOV_INM, register.IndexRegB, 0x11,
				instruction.INST_ADD_RR, register.IndexRegA<<4 | register.IndexRegB,
			},
			Expect: &test.Expect{
				Registers: []*test.RegData{
					{Index: register.IndexRegA, Data: 0xff},
					{Index: register.IndexRegB, Data: 0x10},
				},
			},
		},
		{
			Name: "sub reg inm",
			// sub a, 0x10
			Program: []uint8{instruction.INST_SUB_REG_INM, register.IndexRegA, 0x10},
			Expect: &test.Expect{
				Registers: []*test.RegData{
					{Index: register.IndexRegA, Data: 0xf0},
				},
			},
		},
		{
			Name: "sub reg inm",
			// mov a, 0x20
			// sub a, 0x10
			Program: []uint8{
				instruction.INST_MOV_INM, register.IndexRegA, 0x20,
				instruction.INST_SUB_REG_INM, register.IndexRegA, 0x10,
			},
			Expect: &test.Expect{
				Registers: []*test.RegData{
					{Index: register.IndexRegA, Data: 0x10},
				},
			},
		},
		{
			Name: "sub reg inm",
			// mov a, 0xff
			// sub a, 0x01
			Program: []uint8{
				instruction.INST_MOV_INM, register.IndexRegA, 0xff,
				instruction.INST_SUB_REG_INM, register.IndexRegA, 0x01,
			},
			Expect: &test.Expect{
				Registers: []*test.RegData{
					{Index: register.IndexRegA, Data: 0xfe},
				},
			},
		},
		{
			Name: "sub reg reg",
			// mov a, 0xff
			// mov b, 0x11
			// sub a, b
			Program: []uint8{
				instruction.INST_MOV_INM, register.IndexRegA, 0xff,
				instruction.INST_MOV_INM, register.IndexRegB, 0x11,
				instruction.INST_SUB_RR, register.IndexRegB<<4 | register.IndexRegA,
			},
			Expect: &test.Expect{
				Registers: []*test.RegData{
					{Index: register.IndexRegA, Data: 0xee},
					{Index: register.IndexRegB, Data: 0x11},
				},
			},
		},
		{
			Name: "sub reg reg",
			// mov a, 0xff
			// mov b, 0x11
			// sub b, a
			Program: []uint8{
				instruction.INST_MOV_INM, register.IndexRegA, 0xff,
				instruction.INST_MOV_INM, register.IndexRegB, 0x11,
				instruction.INST_SUB_RR, register.IndexRegA<<4 | register.IndexRegB,
			},
			Expect: &test.Expect{
				Registers: []*test.RegData{
					{Index: register.IndexRegA, Data: 0xff},
					{Index: register.IndexRegB, Data: 0x12},
				},
			},
		},
		{
			Name: "push pop reg",
			// mov a, 0x10
			// push a
			// pop b
			Program: []uint8{
				instruction.INST_MOV_INM, register.IndexRegA, 0x10,
				instruction.INST_PUSH_REG, register.IndexRegA,
				instruction.INST_POP_REG, register.IndexRegB,
			},
			Expect: &test.Expect{
				Registers: []*test.RegData{
					{Index: register.IndexRegA, Data: 0x10},
					{Index: register.IndexRegB, Data: 0x10},
				},
			},
		},
		{
			Name: "push pop reg inm",
			// push 0x10
			// push 0x20
			// push 0x30
			// push 0x40
			// pop d
			// pop c
			// pop b
			// pop a
			Program: []uint8{
				instruction.INST_PUSH_INM, 0x10,
				instruction.INST_PUSH_INM, 0x20,
				instruction.INST_PUSH_INM, 0x30,
				instruction.INST_PUSH_INM, 0x40,
				instruction.INST_POP_REG, register.IndexRegD,
				instruction.INST_POP_REG, register.IndexRegC,
				instruction.INST_POP_REG, register.IndexRegB,
				instruction.INST_POP_REG, register.IndexRegA,
			},
			Expect: &test.Expect{
				Registers: []*test.RegData{
					{Index: register.IndexRegA, Data: 0x10},
					{Index: register.IndexRegB, Data: 0x20},
					{Index: register.IndexRegC, Data: 0x30},
					{Index: register.IndexRegD, Data: 0x40},
				},
			},
		},
		{
			Name: "inc reg",
			// mov a, 0x20
			// inc a
			Program: []uint8{
				instruction.INST_MOV_INM, register.IndexRegA, 0x20,
				instruction.INST_INC, register.IndexRegA,
			},
			Expect: &test.Expect{
				Registers: []*test.RegData{
					{Index: register.IndexRegA, Data: 0x21},
				},
			},
		},
		{
			Name: "inc reg",
			// mov a, 0xff
			// inc a
			Program: []uint8{
				instruction.INST_MOV_INM, register.IndexRegA, 0xff,
				instruction.INST_INC, register.IndexRegA,
			},
			Expect: &test.Expect{
				Registers: []*test.RegData{
					{Index: register.IndexRegA, Data: 0x00},
				},
			},
		},
		{
			Name: "dec reg",
			// mov a, 0x20
			// dec a
			Program: []uint8{
				instruction.INST_MOV_INM, register.IndexRegA, 0x20,
				instruction.INST_DEC, register.IndexRegA,
			},
			Expect: &test.Expect{
				Registers: []*test.RegData{
					{Index: register.IndexRegA, Data: 0x1f},
				},
			},
		},
		{
			Name: "dec reg",
			// mov a, 0x00
			// dec a
			Program: []uint8{
				instruction.INST_MOV_INM, register.IndexRegA, 0x00,
				instruction.INST_DEC, register.IndexRegA,
			},
			Expect: &test.Expect{
				Registers: []*test.RegData{
					{Index: register.IndexRegA, Data: 0xff},
				},
			},
		},
		{
			Name: "jsr and rts",
			// start:
			//     jsr hello
			//     mov a, 0x10
			//     brk
			// hello:
			//     mov a, 0x20
			//     mov b, 0x20
			//     rts
			//     brk
			Program: []uint8{
				instruction.INST_JSR, 0x07, 0x00,
				instruction.INST_MOV_INM, register.IndexRegA, 0x10,
				instruction.INST_BRK,
				instruction.INST_MOV_INM, register.IndexRegA, 0x20,
				instruction.INST_MOV_INM, register.IndexRegB, 0x20,
				instruction.INST_RTS,
				instruction.INST_BRK,
			},
			Expect: &test.Expect{
				Registers: []*test.RegData{
					{Index: register.IndexRegA, Data: 0x10},
					{Index: register.IndexRegB, Data: 0x20},
				},
			},
		},
	}
)

func TestCore(t *testing.T) {
	// c := test.GetComp()
	// c.SetDebug(true)
	for _, tc := range Tests {
		test.RunCase(t, tc)
	}
}
