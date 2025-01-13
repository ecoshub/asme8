package connectable

import "asme8/emulator/src/bus"

type Connectable interface {
	Attach(addrBus, dataBus *bus.Bus, rangeStart, rangeEnd uint16)
	Tick(rw uint8)
	Read(addr uint16) uint8
	Clear()
}

func IsMyRange(from, to, addr uint16) bool {
	if addr < from || addr > to {
		return false
	}
	return true
}
