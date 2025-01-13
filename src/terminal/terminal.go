package terminal

import (
	"emu/src/bus"
	"emu/src/connectable"
	"emu/src/utils"

	"github.com/ecoshub/termium/utils/ansi"
)

var _ connectable.Connectable = &Terminal{}

type Terminal struct {
	addressBus *bus.Bus
	dataBus    *bus.Bus
	rangeStart uint16
	rangeEnd   uint16
}

func New(size uint16) *Terminal {
	return &Terminal{}
}

func (t *Terminal) Attach(addrBus, dataBus *bus.Bus, rangeStart, rangeEnd uint16) {
	t.addressBus = addrBus
	t.dataBus = dataBus
	t.rangeStart = rangeStart
	t.rangeEnd = rangeEnd
}

func (t *Terminal) Tick(rw uint8) {
	if rw == utils.IO_READ {
		t.ReadRequest()
		return
	}
	if rw == utils.IO_WRITE {
		t.WriteRequest()
	}
}

func (t *Terminal) ReadRequest() {

}

func (t *Terminal) WriteRequest() {

}

func (t *Terminal) Read(addr uint16) uint8 {
	return 0
}

func (t *Terminal) Clear() {}

func ResetTerminal() {
	print(ansi.ResetAllModes)
	print(ansi.MakeCursorVisible)
}
