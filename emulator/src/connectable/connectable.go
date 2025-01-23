package connectable

import "asme8/emulator/src/bus"

type Connectable interface {
	GetName() string
	Attach(addrBus, outputBus *bus.Bus, rangeStart, rangeEnd uint16)
	Tick(rw uint8)
	Read(addr uint16) uint8
	Clear()
	GetRange() (uint16, uint16)
}

func IsMyRange(from, to, addr uint16) bool {
	if addr < from || addr > to {
		return false
	}
	return true
}
