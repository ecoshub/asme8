package register

import "fmt"

const (
	RegA int = iota
	RegB int = iota
	RegC int = iota
	RegD int = iota
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
	str := "# registers: "
	for i, r := range m {
		str += fmt.Sprintf("0x%x", r)
		if i != len(m)-1 {
			str += " "
		}
	}
	return str
}
