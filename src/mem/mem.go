package mem

import "fmt"

type Mem struct {
	mem []byte
}

func New(size int) *Mem {
	return &Mem{
		mem: make([]byte, size),
	}
}

func (m *Mem) Read(addr uint16) byte {
	data := m.mem[addr]
	fmt.Printf("MEM READ. addr: %04x, data: %02x\n", addr, data)
	return data
}

func (m *Mem) Write(addr uint16, data byte) {
	fmt.Printf("MEM WRITE. addr: %04x, data: %02x\n", addr, data)
	m.mem[addr] = data
}

func (m *Mem) Load(offset int, program []byte) {
	copy(m.mem[offset:offset+len(program)], program[:])
}
