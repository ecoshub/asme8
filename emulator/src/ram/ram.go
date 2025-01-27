package ram

import (
	"asme8/emulator/src/bus"
	"asme8/emulator/src/connectable"
	"asme8/emulator/utils"
)

var _ connectable.Connectable = &Ram{}

type Ram struct {
	name       string
	addressBus *bus.Bus
	dataBus    *bus.Bus
	size       uint16
	addrStart  uint16
	addrSize   uint16
	data       []byte
}

func New(size uint16) *Ram {
	return &Ram{
		name: "RAM",
		size: size,
		data: make([]byte, size),
	}
}

func (r *Ram) Load(offset int, data []byte) {
	copy(r.data[offset:offset+len(data)], data[:])
}

func (r *Ram) Attach(addrBus, dataBus *bus.Bus, rangeStart, size uint16) {
	r.addressBus = addrBus
	r.dataBus = dataBus
	r.addrStart = rangeStart
	r.addrSize = size
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
	val := r.addressBus.Read_16()
	if !connectable.IsMyRange(r.addrStart, r.addrSize, val) {
		return
	}
	addr := val - r.addrStart
	r.dataBus.Write_8(r.data[addr])
}

func (r *Ram) WriteRequest() {
	addr := r.addressBus.Read_16()
	// fmt.Printf("WriteRequest arrived. addr: %04x, start: %04x, size: %04x, range: %04x-%04x\n", addr, r.addrStart, r.addrSize, r.addrStart, r.addrStart+r.addrSize)
	if !connectable.IsMyRange(r.addrStart, r.addrSize, addr) {
		return
	}
	addr = addr - r.addrStart
	data := r.dataBus.Read_16()
	// fmt.Printf("Writing. addr: %04x, data: %02x\n", addr, data)
	r.data[addr] = uint8(data)
}

func (r *Ram) GetName() string {
	return r.name
}

func (r *Ram) GetRange() (uint16, uint16) {
	return r.addrStart, r.addrSize
}

func (r *Ram) Read(addr uint16) uint8 {
	if !connectable.IsMyRange(r.addrStart, r.addrSize, addr) {
		return 0
	}
	addr = addr - r.addrStart
	return r.data[addr]
}

func (r *Ram) Clear() {
	r.data = make([]byte, r.size)
}
