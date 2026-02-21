package terminal

import (
	"github.com/ecoshub/termium/utils/ansi"
)

type Terminal struct {
	Keyboard   *Keyboard
	Screen     *Screen
	Components *Components
	running    bool
}

func New(sizeBufferSize int, isSerial bool, interruptFunc func()) (*Terminal, error) {
	components, err := NewSetup()
	if err != nil {
		return nil, err
	}
	s, err := NewScreen(components, sizeBufferSize, isSerial)
	if err != nil {
		return nil, err
	}
	return &Terminal{
		Keyboard:   NewKeyboard(components.Screen.CommandPalette, interruptFunc),
		Screen:     s,
		Components: components,
	}, nil
}

func (t *Terminal) Run() {
	if t.running {
		return
	}
	t.running = true
	t.Components.Screen.Start()
	t.running = false
}

func ResetScreen() {
	print(ansi.ResetAllModes)
	print(ansi.MakeCursorVisible)
}
