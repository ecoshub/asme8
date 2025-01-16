package terminal

import (
	"asme8/emulator/src/bus"
	"asme8/emulator/src/connectable"
	"asme8/emulator/utils"

	"github.com/ecoshub/termium/component/palette"
	"github.com/eiannone/keyboard"
)

var _ connectable.Connectable = &Keyboard{}

type Keyboard struct {
	addressBus  *bus.Bus
	dataBus     *bus.Bus
	rangeStart  uint16
	rangeEnd    uint16
	input       uint8
	ready       bool
	cp          *palette.Palette
	pipeInput   bool
	pipeChanged func(pipeInput bool)
}

func NewKeyboard(cp *palette.Palette) *Keyboard {
	k := &Keyboard{
		cp:        cp,
		pipeInput: false,
	}
	cp.AttachKeyEventHandler(func(event keyboard.KeyEvent) {
		if event.Key == keyboard.KeyCtrlD {
			k.pipeInput = !k.pipeInput
			cp.SetBaseListener(k.pipeInput)
			if k.pipeChanged != nil {
				k.pipeChanged(k.pipeInput)
			}
		}
		if !k.pipeInput {
			return
		}
		input := uint8(event.Key)
		if event.Key == 0 {
			input = uint8(event.Rune)
		}
		k.input = uint8(input)
		k.ready = true
	})
	return k
}

// Attach implements connectable.Connectable.
func (k *Keyboard) Attach(addrBus *bus.Bus, dataBus *bus.Bus, rangeStart uint16, rangeEnd uint16) {
	k.addressBus = addrBus
	k.dataBus = dataBus
	k.rangeStart = rangeStart
	k.rangeEnd = rangeEnd
}

// Clear implements connectable.Connectable.
func (k *Keyboard) Clear() {
	k.input = 0
}

// Read implements connectable.Connectable.
func (k *Keyboard) Read(addr uint16) uint8 {
	return k.input
}

// Tick implements connectable.Connectable.
func (k *Keyboard) Tick(rw uint8) {
	if rw == utils.IO_READ {
		k.ReadRequest()
		return
	}
	if rw == utils.IO_WRITE {
		k.WriteRequest()
	}
}

func (k *Keyboard) ReadRequest() {
	addr := k.addressBus.Read()
	if !connectable.IsMyRange(k.rangeStart, k.rangeEnd, addr) {
		return
	}
	if addr == k.rangeStart {
		if k.ready {
			k.dataBus.Write_8(1)
			return
		}
		k.dataBus.Write_8(0)
		return
	}
	if addr == k.rangeStart+1 {
		if k.ready {
			k.dataBus.Write_8(k.input)
			k.ready = false
			return
		}
		k.dataBus.Write_8(0)
		return
	}
}

func (k *Keyboard) WriteRequest() {}

func (k *Keyboard) AttachPipeChange(f func(pipeInput bool)) {
	if f == nil {
		return
	}
	k.pipeChanged = f
}
