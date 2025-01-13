package keyboard

import (
	"asme8/emulator/src/bus"
	"asme8/emulator/src/connectable"
	"asme8/emulator/utils"
	"os"

	"github.com/ecoshub/termium/utils/ansi"
	"github.com/eiannone/keyboard"
)

var _ connectable.Connectable = &Keyboard{}

type Keyboard struct {
	addressBus *bus.Bus
	dataBus    *bus.Bus
	rangeStart uint16
	rangeEnd   uint16
	input      uint8
	consumed   bool
}

func New() *Keyboard {
	return &Keyboard{}
}

func (k *Keyboard) Listen() {
	k.ListenKeys()
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
		if k.consumed {
			k.dataBus.Write_8(0)
			return
		}
		k.dataBus.Write_8(1)
		return
	}
	if addr == k.rangeStart+1 {
		if k.consumed {
			k.dataBus.Write_8(0)
			return
		}
		k.dataBus.Write_8(k.input)
		k.consumed = true
	}
}

func (k *Keyboard) WriteRequest() {}

func (k *Keyboard) ListenKeys() {
	keyEvents, err := keyboard.GetKeys(1)
	if err != nil {
		return
	}
	func(keyEvents <-chan keyboard.KeyEvent) {
		for event := range keyEvents {
			k.input = uint8(event.Key)
			if event.Key == 0 {
				k.input = uint8(event.Rune)
			}
			// fmt.Println("input", k.input, string(k.input))
			k.consumed = false
			if event.Key == keyboard.KeyCtrlC {
				print(ansi.ResetAllModes)
				// print(ansi.MakeCursorVisible)
				os.Exit(0)
			}
		}
	}(keyEvents)
}
