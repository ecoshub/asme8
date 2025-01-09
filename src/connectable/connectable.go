package connectable

import "emu/src/bus"

type Connectable interface {
	Attach(addrBus, dataBus *bus.Bus, rangeStart, rangeEnd uint16)
	Tick(rw uint8)
}

func IsMyRange(from, to, addr uint16) bool {
	if addr < from || addr > to {
		return false
	}
	return true
}
