package mem

type Mem struct {
	mem []byte
}

func New(size int) *Mem {
	return &Mem{
		mem: make([]byte, size),
	}
}

func (m *Mem) Read(addr uint16) byte {
	return m.mem[addr]
}

func (m *Mem) Write(addr uint16, data byte) {
	m.mem[addr] = data
}

func (m *Mem) Load(offset int, program []byte) {
	copy(m.mem[offset:offset+len(program)], program[:])
}
