package emulator

import "github.com/ecoshub/termium/component/palette"

type CommandPalette struct {
	commandPalette *palette.Palette
	lastCommand    string
}

func NewCommandPalette(commandPalette *palette.Palette) *CommandPalette {
	return &CommandPalette{
		commandPalette: commandPalette,
	}
}

func (s *CommandPalette) AddToCommandHistory(input string) {
	if s.lastCommand == input {
		return
	}
	s.commandPalette.AddToHistory(input)
	s.lastCommand = input
}

func (s *CommandPalette) ClearCommandHistory() {
	s.commandPalette.ClearHistory()
	s.lastCommand = ""
}

func (s *CommandPalette) GetLastCommand() string {
	return s.lastCommand
}
