package mem

import "fmt"

type PortEvent uint8

const (
	PORT_W PortEvent = 0
	PORT_R PortEvent = 1
)

type PortListener func(addr uint16, data uint8, rw PortEvent)

func (m *Mem) AttachPort(addr uint16) {
	m.ports = append(m.ports, addr)
}

func (m *Mem) AttachPortListener(f PortListener) {
	if f == nil {
		return
	}
	m.listener = f
}

func (m *Mem) invokePort(addr uint16, data uint8, rw PortEvent) {
	if m.listener == nil {
		return
	}
	for _, p := range m.ports {
		if p == addr {
			if m.debug {
				fmt.Printf(" >> PORT LISTEN INVOKE. addr: %04x, data: %02x, rw: %x\n", addr, data, rw)
			}
			m.listener(addr, data, rw)
			break
		}
	}
}
