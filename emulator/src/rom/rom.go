package rom

import (
	"asme8/emulator/src/bus"
	"asme8/emulator/src/connectable"
	"asme8/emulator/utils"
)

var _ connectable.Connectable = &Rom{}

type Rom struct {
	name       string
	addressBus *bus.Bus
	dataBus    *bus.Bus
	size       uint16
	addrStart  uint16
	addrSize   uint16
	data       []byte
}

func New(size uint16) *Rom {
	return &Rom{
		name: "ROM",
		size: size,
		data: make([]byte, size),
	}
}

func (r *Rom) SetName(name string) {
	r.name = name
}

func (r *Rom) Load(offset int, data []byte) {
	copy(r.data[offset:len(data)], data[:])
}

func (r *Rom) GetData() []byte {
	return r.data
}

func (r *Rom) Attach(addrBus, dataBus *bus.Bus, rangeStart, size uint16) {
	r.addressBus = addrBus
	r.dataBus = dataBus
	r.addrStart = rangeStart
	r.addrSize = size
}

func (r *Rom) Tick(rw uint8) {
	if rw == utils.IO_READ {
		r.ReadRequest()
		return
	}
}

func (r *Rom) ReadRequest() {
	val := r.addressBus.Read_16()
	if !connectable.IsMyRange(r.addrStart, r.addrSize, val) {
		return
	}
	addr := val - r.addrStart
	r.dataBus.Write_8(r.data[addr])
}

func (r *Rom) Read(addr uint16) uint8 {
	if !connectable.IsMyRange(r.addrStart, r.addrSize, addr) {
		return 0
	}
	addr = addr - r.addrStart
	return r.data[addr]
}

func (r *Rom) GetName() string {
	return r.name
}

func (r *Rom) GetRange() (uint16, uint16) {
	return r.addrStart, r.addrStart + r.addrSize
}

func (r *Rom) Clear() {
	r.data = make([]byte, r.size)
}
