package terminal

type ScreenBuffer struct {
	width  int
	height int
	size   int
	buffer []rune
}

func NewScreenBuffer() *ScreenBuffer {
	size := SCREEN_WIDTH * SCREEN_HEIGHT
	return &ScreenBuffer{
		width:  SCREEN_WIDTH,
		height: SCREEN_HEIGHT,
		size:   size,
		buffer: make([]rune, size),
	}
}

func (b *ScreenBuffer) write(addr uint16, data rune) {
	b.buffer[addr] = data
}

func (b *ScreenBuffer) read(addr uint16) rune {
	return b.buffer[addr]
}

func (b *ScreenBuffer) clear() {
	b.buffer = make([]rune, b.size)
}
