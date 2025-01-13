package rom

import (
	"emu/src/bus"
	"emu/src/connectable"
	"emu/src/utils"
)

var _ connectable.Connectable = &Rom{}

type Rom struct {
	addressBus *bus.Bus
	dataBus    *bus.Bus
	size       uint16
	rangeStart uint16
	rangeEnd   uint16
	data       []byte
}

func New(size uint16) *Rom {
	return &Rom{
		size: size,
		data: make([]byte, size),
	}
}

func (r *Rom) Load(offset int, data []byte) {
	copy(r.data[offset:len(data)], data[:])
}

func (r *Rom) GetData() []byte {
	return r.data
}

func (r *Rom) Attach(addrBus, dataBus *bus.Bus, rangeStart, rangeEnd uint16) {
	r.addressBus = addrBus
	r.dataBus = dataBus
	r.rangeStart = rangeStart
	r.rangeEnd = rangeEnd
}

func (r *Rom) Tick(rw uint8) {
	if rw == utils.IO_READ {
		r.ReadRequest()
		return
	}
}

func (r *Rom) ReadRequest() {
	val := r.addressBus.Read()
	if !connectable.IsMyRange(r.rangeStart, r.rangeEnd, val) {
		return
	}
	addr := val - r.rangeStart
	r.dataBus.Write_8(r.data[addr])
}

func (r *Rom) Read(addr uint16) uint8 {
	if !connectable.IsMyRange(r.rangeStart, r.rangeEnd, addr) {
		return 0
	}
	addr = addr - r.rangeStart
	return r.data[addr]
}

func (r *Rom) Clear() {
	r.data = make([]byte, r.size)
}
