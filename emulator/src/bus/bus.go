package bus

type Bus struct {
	val    uint16
	events []func()
}

func New() *Bus {
	return &Bus{val: 0}
}

func (b *Bus) Write_16(val uint16) {
	defer b.runChangeEvent()
	b.val = val
}

func (b *Bus) Write_8(val uint8) {
	b.Write_16(uint16(val))
}

func (b *Bus) Read_16() uint16 {
	defer b.runChangeEvent()
	return b.val
}

func (b *Bus) Read_8() uint8 {
	return uint8(b.Read_16())
}

func (b *Bus) Clear() {
	b.val = 0
}

func (b *Bus) AttachBusChange(f func()) {
	if f == nil {
		return
	}
	b.events = append(b.events, f)
}

func (b *Bus) runChangeEvent() {
	for _, f := range b.events {
		if f == nil {
			continue
		}
		f()
	}
}
