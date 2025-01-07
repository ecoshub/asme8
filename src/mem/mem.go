package mem

import "fmt"

type Mem struct {
	mem      []uint8
	ports    []uint16
	listener PortListener
	debug    bool
}

func New(size int) *Mem {
	return &Mem{
		mem:   make([]uint8, size),
		ports: make([]uint16, 0, 2),
	}
}

func (m *Mem) Read(addr uint16) uint8 {
	data := m.mem[addr]
	defer m.invokePort(addr, data, PORT_R)
	if m.debug {
		fmt.Printf("MEM READ. addr: %04x, data: %02x\n", addr, data)
	}
	return data
}

func (m *Mem) Write(addr uint16, data uint8) {
	defer m.invokePort(addr, data, PORT_W)
	if m.debug {
		fmt.Printf("MEM WRITE. addr: %04x, data: %02x\n", addr, data)
	}
	m.mem[addr] = data
}

func (m *Mem) Load(offset int, program []uint8) {
	copy(m.mem[offset:offset+len(program)], program[:])
}

func (m *Mem) SetDebug(enable bool) {
	m.debug = enable
}
