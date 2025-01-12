package video

import (
	"emu/src/peripheral"
	"time"

	"github.com/ecoshub/termium/utils/ansi"
)

const (
	SCREEN_WIDTH  uint8 = 80
	SCREEN_HEIGHT uint8 = 24
)

const (
	RESET_CHAR rune = ' '
)

const (
	REFRESH_RATE_HZ float64 = 100
)

func (v *Video) Reset() {
	print(ansi.MakeCursorInvisible)
	print(ansi.ResetBlink)
	peripheral.TerminalClear()
	for i := 0; i < int(v.buffer.height); i++ {
		for j := 0; j < int(v.buffer.width); j++ {
			peripheral.TerminalCharOut(i+1, j+1, RESET_CHAR)
			addr := uint16(i*int(v.buffer.height) + j)
			v.buffer.write(addr, RESET_CHAR)
		}
	}
	peripheral.TerminalGoToFirstBlock()
}

func (v *Video) Run() {
	t := time.NewTicker(time.Duration(1.0/REFRESH_RATE_HZ*1000) * time.Millisecond)
	v.running = true
	for {
		select {
		case <-t.C:
			v.refresh()
		case <-v.stopChan:
			return
		}
	}
}

func (v *Video) Stop() {
	if !v.running {
		return
	}
	v.stopChan <- struct{}{}
	v.running = false
}

func (v *Video) refresh() {
	for i := 0; i < v.buffer.size; i++ {
		row := i / int(v.buffer.width)
		column := i - row*int(v.buffer.width)
		r := v.buffer.buffer[i]
		peripheral.TerminalCharOut(row+1, column+1, r)
	}
}
