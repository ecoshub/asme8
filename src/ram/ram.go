package ram

import (
	"emu/src/bus"
	"emu/src/connectable"
	"emu/src/utils"
)

var _ connectable.Connectable = &Ram{}

type Ram struct {
	addressBus *bus.Bus
	dataBus    *bus.Bus
	size       uint16
	rangeStart uint16
	rangeEnd   uint16
	data       []byte
}

func New(size uint16) *Ram {
	return &Ram{
		size: size,
		data: make([]byte, size),
	}
}

func (r *Ram) Load(offset int, data []byte) {
	copy(r.data[offset:offset+len(data)], data[:])
}

func (r *Ram) Attach(addrBus, dataBus *bus.Bus, rangeStart, rangeEnd uint16) {
	r.addressBus = addrBus
	r.dataBus = dataBus
	r.rangeStart = rangeStart
	r.rangeEnd = rangeEnd
}

func (r *Ram) Tick(rw uint8) {
	if rw == utils.IO_READ {
		r.ReadRequest()
		return
	}
	if rw == utils.IO_WRITE {
		r.WriteRequest()
	}
}

func (r *Ram) ReadRequest() {
	val := r.addressBus.Read()
	if !connectable.IsMyRange(r.rangeStart, r.rangeEnd, val) {
		return
	}
	addr := val - r.rangeStart
	r.dataBus.Write_8(r.data[addr])
}

func (r *Ram) WriteRequest() {
	addr := r.addressBus.Read()
	if !connectable.IsMyRange(r.rangeStart, r.rangeEnd, addr) {
		return
	}
	addr = addr - r.rangeStart
	data := r.dataBus.Read()
	r.data[addr] = uint8(data)
}
