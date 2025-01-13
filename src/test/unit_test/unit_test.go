package unit_test

import (
	"emu/src/instruction"
	"emu/src/register"
	"emu/src/test"
	"testing"
)

var (
	Tests = []*test.TestCase{
		{
			Name: "mov immediate reg a",
			// mov a, 0x10
			Program: []uint8{instruction.INST_MOV_INM_8, register.IndexRegA, 0x10},
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
			Program: []uint8{instruction.INST_MOV_INM_8, register.IndexRegA, 0xff},
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
				instruction.INST_MOV_INM_8, register.IndexRegA, 0x10,
				instruction.INTS_MOV_REG_REG, register.IndexRegA<<1 | register.IndexRegB,
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
				instruction.INST_MOV_INM_8, register.IndexRegA, 0xff,
				instruction.INTS_MOV_REG_REG, register.IndexRegA<<4 | register.IndexRegB,
				instruction.INTS_MOV_REG_REG, register.IndexRegA<<4 | register.IndexRegC,
				instruction.INTS_MOV_REG_REG, register.IndexRegA<<4 | register.IndexRegD,
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
				instruction.INST_MOV_INM_8, register.IndexRegA, 0xff,
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
			Name: "mov reg data",
			Program: []uint8{
				instruction.INST_MOV_INM_8, register.IndexRegA, 0x90,
				instruction.INST_MOV_REG_MEM, register.IndexRegA, 0x00, 0x90,
				instruction.INST_MOV_INM_8, register.IndexRegA, 0x91,
				instruction.INST_MOV_REG_MEM, register.IndexRegA, 0x01, 0x90,
				instruction.INST_MOV_INM_8, register.IndexRegA, 0x92,
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
			Name: "mov reg data",
			Program: []uint8{
				instruction.INST_MOV_INM_8, register.IndexRegA, 0x90,
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
			Name: "mov reg data offset",
			Program: []uint8{
				instruction.INST_MOV_INM_8, register.IndexRegA, 0x10,
				instruction.INST_MOV_INM_8, register.IndexRegB, 0x2,
				instruction.INST_MOV_REG_MEM, register.IndexRegA, 0x02, 0x90,
				instruction.INST_MOV_REG_MEM_OFFSET, register.IndexRegC<<4 | register.IndexRegB, 0x00, 0x90,
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
			Name: "mov reg data offset",
			Program: []uint8{
				instruction.INST_MOV_INM_8, register.IndexRegA, 0x10,
				instruction.INST_MOV_INM_8, register.IndexRegB, 0x2,
				instruction.INST_MOV_REG_MEM, register.IndexRegA, 0x01, 0x91,
				instruction.INST_MOV_REG_MEM_OFFSET, register.IndexRegC<<4 | register.IndexRegB, 0xff, 0x90,
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
	}
)

func TestCore(t *testing.T) {
	for _, tc := range Tests {
		test.RunCase(t, tc)
	}
}
