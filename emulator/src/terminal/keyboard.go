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
	name             string
	addressBus       *bus.Bus
	dataBus          *bus.Bus
	addrStart        uint16
	addrSize         uint16
	input            uint8
	ready            bool
	cp               *palette.Palette
	pipeInput        bool
	pipeChangedEvent func(pipeInput bool)
	scrollEvent      func(c uint16)
}

func NewKeyboard(cp *palette.Palette, interruptFunc func()) *Keyboard {
	k := &Keyboard{
		name:      "KEYBOARD",
		cp:        cp,
		pipeInput: false,
	}
	cp.AttachKeyEventHandler(func(event keyboard.KeyEvent) {
		switch event.Key {
		case keyboard.KeyCtrlD:
			k.pipeInput = !k.pipeInput
			cp.SetBaseListener(k.pipeInput)
			if k.pipeChangedEvent != nil {
				k.pipeChangedEvent(k.pipeInput)
			}
			return
		case keyboard.KeyArrowLeft, keyboard.KeyArrowRight:
			if k.scrollEvent != nil {
				k.scrollEvent(uint16(event.Key))
			}
			return
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
		if interruptFunc != nil {
			interruptFunc()
		}
	})
	return k
}

// Attach implements connectable.Connectable.
func (k *Keyboard) Attach(addrBus *bus.Bus, dataBus *bus.Bus, rangeStart uint16, size uint16) {
	k.addressBus = addrBus
	k.dataBus = dataBus
	k.addrStart = rangeStart
	k.addrSize = size
}

func (k *Keyboard) GetName() string {
	return k.name
}

func (k *Keyboard) GetRange() (uint16, uint16) {
	return k.addrStart, uint16(uint64(k.addrStart) + uint64(k.addrSize) - 1)
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
	addr := k.addressBus.Read_16()
	if !connectable.IsMyRange(k.addrStart, k.addrSize, addr) {
		return
	}
	if addr == k.addrStart {
		if k.ready {
			k.dataBus.Write_8(1)
			return
		}
		k.dataBus.Write_8(0)
		return
	}
	if addr == k.addrStart+1 {
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

func (k *Keyboard) Load(offset int, data []byte) {}

func (k *Keyboard) AttachPipeChange(f func(pipeInput bool)) {
	if f == nil {
		return
	}
	k.pipeChangedEvent = f
}

func (k *Keyboard) AttachScrollEvent(f func(c uint16)) {
	if f == nil {
		return
	}
	k.scrollEvent = f
}

func (k *Keyboard) SetPipeInput(value bool) {
	k.pipeInput = value
}

func (k *Keyboard) GetPipeInput() bool {
	return k.pipeInput
}
