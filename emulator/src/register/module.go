package register

import "fmt"

const (
	IndexRegA uint8 = 0
	IndexRegB uint8 = 1
	IndexRegC uint8 = 2
	IndexRegD uint8 = 3
)

type Module []byte

func NewModule() Module {
	return make([]byte, 4)
}

func (m Module) Read(index uint8) byte {
	return m[index]
}

func (m Module) Write(index uint8, val byte) {
	m[index] = val
}

func (m Module) String() string {
	str := ""
	for i, r := range m {
		str += fmt.Sprintf("%02x", r)
		if i != len(m)-1 {
			str += " "
		}
	}
	return str
}
