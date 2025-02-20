package register

import "fmt"

type Module []uint8

func NewModule(size int) Module {
	return make([]uint8, size)
}

func (m Module) Read_8(index uint8) uint8 {
	return m[index]
}

func (m Module) Read_16(index uint8) uint16 {
	return uint16(m[index])
}

func (m Module) Write(index uint8, val uint8) {
	m[index] = val
}

func (m Module) Clear() {
	for i := range m {
		m[i] = 0
	}
}

func (m Module) String() string {
	str := ""
	for i, r := range m[:6] {
		str += fmt.Sprintf("%02x", r)
		if i != len(m)-1 {
			str += " "
		}
	}
	return str
}
