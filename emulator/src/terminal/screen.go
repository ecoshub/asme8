package terminal

import (
	"asme8/emulator/src/bus"
	"asme8/emulator/src/connectable"
	"asme8/emulator/utils"
)

var _ connectable.Connectable = &Screen{}

type Screen struct {
	addressBus *bus.Bus
	dataBus    *bus.Bus
	rangeStart uint16
	rangeEnd   uint16
	buffer     *ScreenBuffer
	components *Components
}

// Clear implements connectable.Connectable.
func (v *Screen) Clear() {
	v.buffer.clear()
	v.components.MainPanel.Clear()
	v.components.DebugPanel.Clear()
	v.components.MemoryPanel.Clear()
	v.components.FlagPanel.Clear()
}

func NewScreen(components *Components) (*Screen, error) {
	return &Screen{
		buffer:     NewScreenBuffer(),
		components: components,
	}, nil
}

// Attach implements connectable.Connectable.
func (v *Screen) Attach(addrBus *bus.Bus, dataBus *bus.Bus, rangeStart uint16, rangeEnd uint16) {
	v.addressBus = addrBus
	v.dataBus = dataBus
	v.rangeStart = rangeStart
	v.rangeEnd = rangeEnd
}

// Tick implements connectable.Connectable.
func (v *Screen) Tick(rw uint8) {
	if rw == utils.IO_WRITE {
		v.WriteRequest()
	}
}

func (v *Screen) WriteRequest() {
	addr := v.addressBus.Read()
	if !connectable.IsMyRange(v.rangeStart, v.rangeEnd, addr) {
		return
	}
	addr = addr - v.rangeStart
	data := v.dataBus.Read()
	v.buffer.write(addr, rune(data))
	v.components.MainPanel.Write(int(addr), uint8(data))
}

func (v *Screen) Read(addr uint16) uint8 {
	if !connectable.IsMyRange(v.rangeStart, v.rangeEnd, addr) {
		return 0
	}
	addr = addr - v.rangeStart
	return uint8(v.buffer.read(addr))
}
