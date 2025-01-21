package assembler

import (
	"asme8/assembler/src/error_listener"
	"asme8/assembler/src/parser"
	"fmt"

	"github.com/antlr4-go/antlr/v4"
)

func AssembleFile(filePath string) (*parser.Assembler, error) {
	input, err := antlr.NewFileStream(filePath)
	if err != nil {
		return nil, fmt.Errorf("file stream error. err: %s", err)
	}

	cel := &error_listener.CustomErrorListener{}
	lexer := parser.NewAsmE8Lexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(cel)

	p := parser.NewAsmE8Parser(stream)
	p.RemoveErrorListeners()
	p.AddErrorListener(cel)
	tree := p.Instruction()

	assembler := parser.NewAssembler(cel)
	antlr.ParseTreeWalkerDefault.Walk(assembler, tree)

	return assembler, cel.GetError()
}
