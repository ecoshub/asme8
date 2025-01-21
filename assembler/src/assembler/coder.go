package assembler

func (a *Assembler) CodeParseExitInst(text string) {
	a.Coder.linesEndings = append(a.Coder.linesEndings, a.offset)
	a.Coder.instructions = append(a.Coder.instructions, text)
}

func (a *Assembler) CodeParseLastLine(text string) {
	if text == "\n" && a.Coder.lastLine == "\n" {
		a.Coder.blanksLines = append(a.Coder.blanksLines, a.offset)
	}
	a.Coder.lastLine = text
}
