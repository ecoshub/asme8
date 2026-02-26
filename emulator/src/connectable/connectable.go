package connectable

import (
	"asme8/emulator/src/bus"
)

type Connectable interface {
	GetName() string
	Attach(addrBus, outputBus *bus.Bus, rangeStart, rangeEnd uint16)
	Tick(rw uint8)
	Read(addr uint16) uint8
	Clear()
	GetRange() (uint16, uint16)
	Load(offset int, data []byte)
}

func IsMyRange(from, size, addr uint16) bool {
	if addr < from {
		return false
	}
	if addr > from+size-1 {
		return false
	}
	return true
}
