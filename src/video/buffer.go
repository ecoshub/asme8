package video

import (
	"emu/src/peripheral"
	"time"

	"github.com/ecoshub/termium/utils/ansi"
)

type Buffer struct {
	width  uint8
	height uint8
	size   int
	buffer []rune
}

func (b *Buffer) reset() {
	print(ansi.MakeCursorInvisible)
	print(ansi.ResetBlink)
	peripheral.TerminalClear()
	for i := 0; i < int(b.height); i++ {
		for j := 0; j < int(b.width); j++ {
			peripheral.TerminalCharOut(i+1, j+1, RESET_CHAR)
			b.buffer[i*int(b.height)+j] = RESET_CHAR
		}
	}
	peripheral.TerminalGoToFirstBlock()
}

func (b *Buffer) refresh() {
	for i := 0; i < b.size; i++ {
		row := i / int(b.width)
		column := i - row*int(b.width)
		r := b.buffer[i]
		peripheral.TerminalCharOut(row+1, column+1, r)
	}
}

func (b *Buffer) run() {
	t := time.NewTicker(time.Duration(1.0/REFRESH_RATE_HZ*1000) * time.Millisecond)
	for range t.C {
		b.refresh()
	}
}

func (b *Buffer) write(addr uint16, data uint16) {
	b.buffer[addr] = rune(data)
}
