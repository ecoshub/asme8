package peripheral

import (
	"github.com/ecoshub/termium/utils/ansi"
)

func TerminalClear() {
	print(ansi.ClearScreen)
}

func TerminalGoToFirstBlock() {
	print(ansi.GoToFirstBlock)
}

func TerminalCharOut(row, column int, char rune) {
	ansi.GotoRowAndColumn(row, column)
	print(string(char))
}

func TerminalCharErase(row, column int) {
	ansi.GotoRowAndColumn(row, column)
	print(ansi.EscapeChar + "[0K")
}
