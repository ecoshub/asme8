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
	Globals     []string
	SegmentAddr uint16
}

func AssembleFile(options *Options) ([]byte, error) {
	if options == nil {
		return nil, errors.New("must provide at least file path with option")
	}

	input, err := antlr.NewFileStream(options.FilePath)
	if err != nil {
		return nil, fmt.Errorf("file stream error. err: %s", err)
	}

	globals, err := FormatGlobals(options.Globals)
	if err != nil {
		return nil, fmt.Errorf("global read error. err: %s", err)
	}

	el := &error_listener.CustomErrorListener{}
	lexer := parser.NewAsmE8Lexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(el)

	p := parser.NewAsmE8Parser(stream)
	p.RemoveErrorListeners()
	p.AddErrorListener(el)
	tree := p.Instruction()

	assembler := New()
	assembler.AttachErrorListener(el)
	assembler.AttachGlobals(globals)
	antlr.ParseTreeWalkerDefault.Walk(assembler, tree)

	err = el.GetError()
	if err != nil {
		return nil, err
	}

	out, err := assembler.Assemble()
	if err != nil {
		return nil, err
	}

	if options.PrintDetail {
		prt := assembler.CreatePrintable()
		fmt.Println(prt)
	}

	return out, nil
}
