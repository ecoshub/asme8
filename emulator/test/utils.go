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

	var err error
	MainTestComputer, err = computer.New(&computer.Config{
		MemoryConfig: &config.MemoryConfig{
			Configs: []*config.Memory{
				{Name: "ROM", Size: config.NewNullable(0x2000), Type: "ro"},
				{Name: "RAM", Size: config.NewNullable(0xdfff), Type: "rw"},
			},
		},
		Headless:       true,
		Delay:          time.Nanosecond,
		EnableTestMode: true,
	})
	if err != nil {
		if err.Error() != "program not found" {
			panic(err)
		}
	}
	MainROM, _ = MainTestComputer.GetRom()
	MainRAM, _ = MainTestComputer.GetRam()
	return MainTestComputer
}
