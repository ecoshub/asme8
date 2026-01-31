package assembler

import (
	"asme8/assembler/src/error_listener"
	"asme8/assembler/src/parser"
	"asme8/assembler/src/types"
	"errors"
	"fmt"

	"github.com/antlr4-go/antlr/v4"
)

type Options struct {
	Mode           ASM_MODE
	FilePath       string
	OutputCodePath string
	OutputFilePath string
	PrintDetail    bool
	SegmentAddr    uint16
}

type VariablePair struct {
	Key string
	Val uint16
}

func AssembleFile(options *Options, variables ...*VariablePair) ([]byte, string, int, error) {
	if options == nil {
		return nil, "", 0, errors.New("must provide at least file path with option")
	}

	input, err := antlr.NewFileStream(options.FilePath)
	if err != nil {
		return nil, "", 0, err
	}

	if options.Mode == ASM_MODE_NONE {
		return nil, "", 0, errors.New("must provide an assembly mode")
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

	assembler := New(options.Mode)
	assembler.AttachErrorListener(el)

	for _, v := range variables {
		val := types.NewValue(int64(v.Val))
		assembler.ParseVariable(v.Key, val)
	}

	antlr.ParseTreeWalkerDefault.Walk(assembler, tree)

	err = el.GetError()
	if err != nil {
		return nil, "", 0, err
	}

	out, length, err := assembler.Assemble()
	if err != nil {
		return nil, "", 0, err
	}

	if options.PrintDetail {
		if options.Mode == ASM_MODE_EXE {
			prt := assembler.CreatePrintable()
			fmt.Println(prt)
		}
		if options.Mode == ASM_MODE_ELF {
			assembler.symbolTracker.Print()
		}
	}

	err = assembler.WriteBinaryFile(options.OutputFilePath, out)
	if err != nil {
		return nil, "", 0, err
	}

	return out, assembler.CodeString(), length, nil
}
