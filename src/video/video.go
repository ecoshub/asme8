package video

import (
	"emu/src/bus"
	"emu/src/connectable"
	"emu/src/utils"
)

var _ connectable.Connectable = &Video{}

type Video struct {
	addressBus *bus.Bus
	dataBus    *bus.Bus
	rangeStart uint16
	rangeEnd   uint16
	buffer     *Buffer
	running    bool
	stopChan   chan struct{}
}

func New() *Video {
	return &Video{
		buffer:   NewBuffer(),
		stopChan: make(chan struct{}, 1),
	}
}

// Attach implements connectable.Connectable.
func (v *Video) Attach(addrBus *bus.Bus, dataBus *bus.Bus, rangeStart uint16, rangeEnd uint16) {
	v.addressBus = addrBus
	v.dataBus = dataBus
	v.rangeStart = rangeStart
	v.rangeEnd = rangeEnd
}

// Tick implements connectable.Connectable.
func (v *Video) Tick(rw uint8) {
	if rw == utils.IO_WRITE {
		v.WriteRequest()
	}
}

func (v *Video) WriteRequest() {
	addr := v.addressBus.Read()
	if !connectable.IsMyRange(v.rangeStart, v.rangeEnd, addr) {
		return
	}
	addr = addr - v.rangeStart
	data := v.dataBus.Read()
	v.buffer.write(addr, rune(data))
}
