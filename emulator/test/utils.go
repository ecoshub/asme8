package test

import (
	"asme8/common/config"
	"asme8/emulator/src/computer"
	"asme8/emulator/src/connectable"
	"asme8/emulator/src/ram"
	"asme8/emulator/src/rom"
	"testing"
	"time"
)

var (
	MainTestComputer *computer.Computer
	MainROM          *rom.Rom
	MainRAM          *ram.Ram
	SecondaryRAM     *ram.Ram
)

const (
	TEST_INTERRUPT_VEC_1_ADDR_LOW         = 0x10
	TEST_INTERRUPT_VEC_1_ADDR_HIGH        = 0x11
	TEST_STACK_START               uint16 = 0x20ff
)

func RegTest(t *testing.T, index uint8, inm uint8) {
	c := GetComp()
	regVal := c.ReadRegister(index)
	if regVal != inm {
		t.Fatalf("register value is wrong. expected: 0x%x, got: 0x%x, reg_index: %d", inm, regVal, index)
	}
}

func DataTest(t *testing.T, dev connectable.Connectable, addr uint16, data uint8) {
	val := dev.Read(addr)
	if val != data {
		t.Fatalf("data is wrong. addr: 0x%04x, expected: 0x%x, got: 0x%x", addr, data, val)
	}
}

func GetComp() *computer.Computer {
	if MainTestComputer != nil {
		return MainTestComputer
	}

	computer.IntVec0AddrLow = TEST_INTERRUPT_VEC_1_ADDR_LOW
	computer.IntVec1AddrLow = TEST_INTERRUPT_VEC_1_ADDR_HIGH
	var err error
	MainTestComputer, err = computer.New(&computer.Config{
		MemoryConfig: &config.MemoryConfig{
			Configs: []*config.Memory{
				{Name: "ROM", Size: 0x2000, Type: config.MemoryTypeReadOnly},
				{Name: "RAM", Size: 0xdfed, Type: config.MemoryTypeReadWrite},
				{Name: "SERIAL", Size: 0x3, Type: config.MemoryTypeReadWrite},
				{Name: "RAM_2", Size: 0x10, Type: config.MemoryTypeReadWrite},
			},
		},
		Headless: true,
		Delay:    time.Nanosecond,
		TestMode: true,
	})
	if err != nil {
		if err.Error() != "program not found" {
			panic(err)
		}
	}
	MainTestComputer.SetStackStart(TEST_STACK_START)
	MainROM, _ = MainTestComputer.GetRom()
	MainRAM, _ = MainTestComputer.GetRam()
	dev, _ := MainTestComputer.GetDevice("RAM_2")
	var ok bool
	SecondaryRAM, ok = dev.(*ram.Ram)
	if !ok {
		panic("secondary ram part not found")
	}
	return MainTestComputer
}
