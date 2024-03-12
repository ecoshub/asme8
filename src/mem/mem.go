package mem

type Mem []byte

func New(size int) Mem {
	return make([]byte, size)
}

func (m Mem) Read(addr uint16) byte {
	return m[addr]
}

func (m Mem) Write(addr uint16, data byte) {
	m[addr] = data
}
