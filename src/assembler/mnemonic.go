package assembler

import (
	"emu/src/instruction"
)

var (
	mnemonics = map[string]byte{
		"mov": instruction.INST_ADD_RR,
	}
)
