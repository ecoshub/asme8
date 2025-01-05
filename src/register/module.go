package register

import "fmt"

const (
	RegA int = iota
	RegB
	RegC
	RegD
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
