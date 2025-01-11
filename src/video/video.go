package video

import (
	"emu/src/bus"
	"emu/src/connectable"
	"emu/src/utils"
)

const (
	SCREEN_WIDTH  uint8 = 80
	SCREEN_HEIGHT uint8 = 24
)

const (
	RESET_CHAR rune = ' '
)

const (
	REFRESH_RATE_HZ float64 = 10
)

var _ connectable.Connectable = &Video{}

type Video struct {
	addressBus *bus.Bus
	dataBus    *bus.Bus
	rangeStart uint16
	rangeEnd   uint16
	buffer     *Buffer
}

func New() *Video {
	size := int(SCREEN_WIDTH) * int(SCREEN_HEIGHT)
	buffer := &Buffer{
		width:  SCREEN_WIDTH,
		height: SCREEN_HEIGHT,
		size:   size,
		buffer: make([]rune, size),
	}
	return &Video{
		buffer: buffer,
	}
}

func (v *Video) Reset() {
	v.buffer.reset()
}

func (v *Video) Run() {
	v.buffer.run()
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
	v.buffer.write(addr, data)
}
