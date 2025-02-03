package terminal

import (
	"asme8/emulator/src/bus"
	"asme8/emulator/src/connectable"
	"asme8/emulator/utils"
	"fmt"
)

var _ connectable.Connectable = &Screen{}

type Screen struct {
	name       string
	addressBus *bus.Bus
	dataBus    *bus.Bus
	addrStart  uint16
	addrSize   uint16
	buffer     *ScreenBuffer
	dataBuffer []uint8
	components *Components
}

func NewScreen(components *Components, size int) (*Screen, error) {
	return &Screen{
		name:       "VIDEO",
		buffer:     NewScreenBuffer(size),
		dataBuffer: make([]uint8, 8),
		components: components,
	}, nil
}

func (s *Screen) GetName() string {
	return s.name
}

func (s *Screen) GetRange() (uint16, uint16) {
	return s.addrStart, s.addrStart + s.addrSize
}

// Attach implements connectable.Connectable.
func (s *Screen) Attach(addrBus *bus.Bus, dataBus *bus.Bus, rangeStart uint16, size uint16) {
	s.addressBus = addrBus
	s.dataBus = dataBus
	s.addrStart = rangeStart
	s.addrSize = size
}

// Tick implements connectable.Connectable.
func (s *Screen) Tick(rw uint8) {
	if rw == utils.IO_WRITE {
		s.WriteRequest()
	}
}

func (s *Screen) WriteRequest() {
	addr := s.addressBus.Read_16()
	if !connectable.IsMyRange(s.addrStart, s.addrSize, addr) {
		return
	}
	addr = addr - s.addrStart
	data := s.dataBus.Read_16()
	s.buffer.write(addr, rune(data))
	s.components.SysLogPanel.Push(fmt.Sprintf("push to screen. addr: %04x, data: %02x [%s]", addr, data, string(data)))
	s.components.MainPanel.Write(int(addr), rune(data))
}

func (s *Screen) Read(addr uint16) uint8 {
	if !connectable.IsMyRange(s.addrStart, s.addrSize, addr) {
		return 0
	}
	addr = addr - s.addrStart
	return uint8(s.buffer.read(addr))
}

// Clear implements connectable.Connectable.
func (s *Screen) Clear() {
	s.buffer.clear()
	s.components.MainPanel.Clear()
	s.components.SysLogPanel.Clear()
	s.components.FlagPanel.Clear()
	s.components.MemoryPanel.Clear()
	s.components.CodePanel.Clear()
}
