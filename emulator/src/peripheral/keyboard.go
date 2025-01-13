package peripheral

import (
	"os"

	"github.com/eiannone/keyboard"
)

func ListenKeys(f func(e keyboard.KeyEvent)) {
	if f == nil {
		return
	}
	keyEvents, err := keyboard.GetKeys(3)
	if err != nil {
		return
	}
	go func(keyEvents <-chan keyboard.KeyEvent) {
		for event := range keyEvents {
			switch event.Key {
			case keyboard.KeyCtrlC:
				os.Exit(0)
			default:
				f(event)
			}
		}
	}(keyEvents)
}
