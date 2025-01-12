package video

type Buffer struct {
	width  uint8
	height uint8
	size   int
	buffer []rune
}

func NewBuffer() *Buffer {
	size := int(SCREEN_WIDTH) * int(SCREEN_HEIGHT)
	return &Buffer{
		width:  SCREEN_WIDTH,
		height: SCREEN_HEIGHT,
		size:   size,
		buffer: make([]rune, size),
	}
}

func (b *Buffer) write(addr uint16, data rune) {
	b.buffer[addr] = data
}

func (b *Buffer) read(addr uint16) rune {
	return b.buffer[addr]
}
