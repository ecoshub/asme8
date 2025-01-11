package bus

type BusEvent func(rw uint8, val uint16)

type Bus struct {
	val uint16
}

func New() *Bus {
	return &Bus{val: 0}
}

func (b *Bus) Write_16(val uint16) {
	b.val = val
}

func (b *Bus) Write_8(val uint8) {
	b.val = uint16(val)
}

func (b *Bus) Read() uint16 {
	return b.val
}

func (b *Bus) Read_8() uint8 {
	return uint8(b.val)
}

func (b *Bus) Clear() {
	b.val = 0
}
