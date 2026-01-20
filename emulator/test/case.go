package test

import (
	"asme8/emulator/src/connectable"
	"asme8/emulator/utils"
	"testing"
)

const (
	DEV_TYPE_RAM uint8 = 1
	DEV_TYPE_ROM uint8 = 2
)

type TestCase struct {
	Name    string
	Program []byte
	Expect  *Expect
	Disable bool
	Debug   bool
}

type Expect struct {
	Registers []*ExpectRegister
	Data      []*ExpectData
	Status    *ExpectStatusData
}

type ExpectRegister struct {
	Index uint8
	Data  uint8
}

type ExpectData struct {
	Type uint8
	Addr uint16
	Data uint8
}

type ExpectStatusData struct {
	Data uint8
}

func RunCase(t *testing.T, tc *TestCase) {
	t.Run(tc.Name, func(tt *testing.T) {
		c := GetComp()
		c.Reset(false, false)
		c.Config.Program = tc.Program
		c.LoadProgram()
		tt.Logf("program: %s", utils.ToHexArray(tc.Program))
		c.Run()

		for _, r := range tc.Expect.Registers {
			RegTest(tt, r.Index, r.Data)
		}
		for _, d := range tc.Expect.Data {
			var dev connectable.Connectable
			switch d.Type {
			case DEV_TYPE_RAM:
				dev = MainRAM
			case DEV_TYPE_ROM:
				dev = MainROM
			default:
				tt.Fatalf("unexpected device type. type: %d", d.Type)
			}
			DataTest(tt, dev, d.Addr, d.Data)
		}
		if tc.Expect.Status != nil {
			if c.GetStatusRegister().Flag() != tc.Expect.Status.Data {
				t.Fatalf("unexpected flag status. expected: 0b%08b, got: 0b%08b", tc.Expect.Status.Data, c.GetStatusRegister())
			}
		}
	})
}
