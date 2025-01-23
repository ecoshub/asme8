package assembler

import (
	"asme8/assembler/src/error_listener"
	"asme8/assembler/src/parser"
	"errors"
	"fmt"

	"github.com/antlr4-go/antlr/v4"
)

type Options struct {
	FilePath    string
	PrintDetail bool
}

func AssembleFile(options *Options) ([]byte, error) {
	if options == nil {
		return nil, errors.New("must provide at least file path with option")
	}
	input, err := antlr.NewFileStream(options.FilePath)
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

	assembler := New(cel)
	antlr.ParseTreeWalkerDefault.Walk(assembler, tree)

	out, err := assembler.Assemble()
	if err != nil {
		return nil, err
	}

	err = cel.GetError()
	if err != nil {
		return nil, err
	}

	if options.PrintDetail {
		prt := assembler.CreatePrintable()
		fmt.Println(prt)
	}

	return out, nil
}
