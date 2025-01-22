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
}

type Expect struct {
	Registers []*RegData
	Data      []*ExpectData
}

type RegData struct {
	Index uint8
	Data  uint8
}

type ExpectData struct {
	Type uint8
	Addr uint16
	Data uint8
}

func RunCase(t *testing.T, tc *TestCase) {
	t.Run(tc.Name, func(tt *testing.T) {
		c := GetComp()
		c.Reset(false, false)
		MainROM.Load(0, tc.Program)
		c.ProgramLoaded()
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
	})
}
