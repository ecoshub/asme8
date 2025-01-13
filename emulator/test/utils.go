package test

import (
	"asme8/emulator/src/comp"
	"asme8/emulator/src/connectable"
	"asme8/emulator/src/ram"
	"asme8/emulator/src/rom"
	"testing"
	"time"
)

var (
	MainTestComputer *comp.Comp
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
		t.Fatalf("data read is wrong. expected: 0x%x, got: 0x%x", data, val)
	}
}

func GetComp() *comp.Comp {
	if MainTestComputer != nil {
		return MainTestComputer
	}

	MainTestComputer = comp.New()
	MainROM = rom.New(0x8000)
	MainRAM = ram.New(0x8000)
	MainTestComputer.ConnectDevice(MainROM, 0x0000, 0x8000)
	MainTestComputer.ConnectDevice(MainRAM, 0x8000, 0xffff)
	MainTestComputer.SetDelay(time.Nanosecond)
	return MainTestComputer
}
