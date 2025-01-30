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

func New(sizeBufferSize int) (*Terminal, error) {
	components, err := NewSetup()
	if err != nil {
		return nil, err
	}
	k := NewKeyboard(components.Screen.CommandPalette)
	k.AttachPipeChange(func(pipeInput bool) {
		if pipeInput {
			components.SysLogPanel.Push(">> Keyboard input directed to emulator ( use CTRL + D to switch)")
		} else {
			components.SysLogPanel.Push("<< Keyboard input directed to command pallet ( use CTRL + D to switch)")
		}
	})
	s, err := NewScreen(components, sizeBufferSize)
	if err != nil {
		return nil, err
	}
	return &Terminal{
		Keyboard:   k,
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

func (t *Terminal) Clear() {
	t.Screen.Clear()
}

func ResetScreen() {
	print(ansi.ResetAllModes)
	print(ansi.MakeCursorVisible)
}
