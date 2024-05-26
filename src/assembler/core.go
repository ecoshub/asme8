package assembler

import (
	"github.com/ecoshub/machine"
)

var (
	mnemonic string
	op1      string
	op2      string
)

func (a *Assembler) setRecorders() {

	a.sm.Listen(func(s *machine.State, r rune, changed bool) {
		// fmt.Printf("r:%d[%s], s:%s\n", r, string(r), s.GetTag())
	})

	a.sm.Record("_main_.mnemonic", "_main_.mnemonic_end", func(value string) {
		a.tempOperation.mnemonic = value
	})

	a.sm.Record("_main_.op_1", "_main_.op_1_end", func(value string) {
		a.tempOperation.operand1 = value
	})

	a.sm.Record("_main_.op_2", "_main_.line_end", func(value string) {
		a.tempOperation.operand2 = value
	})

	a.sm.Listen(func(s *machine.State, r rune, changed bool) {
		if s.GetTag() == "_main_.line_end" {
			// fmt.Printf("mnem:>%s<,op1:>%s<,op2:>%s<\n", mnemonic, op1, op2)
			a.operations = append(a.operations, a.tempOperation)
			a.tempOperation = &Operation{}
		}
	})

}
