package terminal

import (
	"asme8/emulator/src/bus"
	"asme8/emulator/src/connectable"
	"asme8/emulator/utils"
)

var _ connectable.Connectable = &Screen{}

type Screen struct {
	name       string
	addressBus *bus.Bus
	dataBus    *bus.Bus
	rangeStart uint16
	rangeEnd   uint16
	buffer     *ScreenBuffer
	components *Components
}

func NewScreen(components *Components) (*Screen, error) {
	return &Screen{
		name:       "SCREEN",
		buffer:     NewScreenBuffer(),
		components: components,
	}, nil
}

func (s *Screen) GetName() string {
	return s.name
}

func (s *Screen) GetRange() (uint16, uint16) {
	return s.rangeStart, s.rangeEnd
}

// Attach implements connectable.Connectable.
func (s *Screen) Attach(addrBus *bus.Bus, dataBus *bus.Bus, rangeStart uint16, rangeEnd uint16) {
	s.addressBus = addrBus
	s.dataBus = dataBus
	s.rangeStart = rangeStart
	s.rangeEnd = rangeEnd
}

// Tick implements connectable.Connectable.
func (s *Screen) Tick(rw uint8) {
	if rw == utils.IO_WRITE {
		s.WriteRequest()
	}
}

func (s *Screen) WriteRequest() {
	addr := s.addressBus.Read_16()
	if !connectable.IsMyRange(s.rangeStart, s.rangeEnd, addr) {
		return
	}
	addr = addr - s.rangeStart
	data := s.dataBus.Read_16()
	s.buffer.write(addr, rune(data))
	s.components.MainPanel.Write(int(addr), rune(data))
}

func (s *Screen) Read(addr uint16) uint8 {
	if !connectable.IsMyRange(s.rangeStart, s.rangeEnd, addr) {
		return 0
	}
	addr = addr - s.rangeStart
	return uint8(s.buffer.read(addr))
}

// Clear implements connectable.Connectable.
func (s *Screen) Clear() {
	s.buffer.clear()
	s.components.MainPanel.Clear()
	s.components.SysLogPanel.Clear()
	s.components.FlagPanel.Clear()
	s.components.MemoryPanel.Clear()
}
