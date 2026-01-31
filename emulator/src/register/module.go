package register

type Module []uint8

func NewModule(size int) Module {
	return make([]uint8, size)
}

func (m Module) Read_8(index uint8) uint8 {
	return m[index]
}

func (m Module) Write(index uint8, val uint8) {
	m[index] = val
}

func (m Module) Clear() {
	for i := range m {
		m[i] = 0
	}
}
