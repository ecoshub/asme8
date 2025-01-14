package assembler

import (
	"asme8/assembler/src/parser"
	"fmt"

	"github.com/antlr4-go/antlr/v4"
)

func AssembleFile(filePath string) (*parser.Assembler, error) {
	input, err := antlr.NewFileStream(filePath)
	if err != nil {
		return nil, fmt.Errorf("file stream error. err: %s", err)
	}

	lexer := parser.NewAsmE8Lexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewAsmE8Parser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	tree := p.Instruction()
	assembler := parser.NewAssembler()
	antlr.ParseTreeWalkerDefault.Walk(assembler, tree)

	return assembler, nil
}
