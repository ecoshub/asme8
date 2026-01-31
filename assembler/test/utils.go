package test

import (
	"asme8/assembler/src/assembler"
	"asme8/assembler/src/error_listener"
	"asme8/assembler/src/parser"
	"fmt"
	"testing"

	"github.com/antlr4-go/antlr/v4"
)

var (
	MainAssembler *assembler.Assembler
)

func OutTest(t *testing.T, expected, got []byte) {
	if len(expected) != len(got) {
		t.Fatalf("out is wrong. expected:\n %s, got:\n %s", ToHexArray(expected), ToHexArray(got))
	}
	for i := 0; i < len(got); i++ {
		if expected[i] != got[i] {
			t.Fatalf("out is wrong. index: %d, expected:\n %s, got:\n %s", i, ToHexArray(expected), ToHexArray(got))
		}
	}
}

func GetAssembler() *assembler.Assembler {
	if MainAssembler != nil {
		return MainAssembler
	}
	MainAssembler = assembler.New(assembler.ASM_MODE_EXE)
	return MainAssembler
}

func Assemble(code string) ([]byte, error) {

	assembler := GetAssembler()
	assembler.Reset()

	input := antlr.NewInputStream(code)
	el := &error_listener.CustomErrorListener{}
	lexer := parser.NewAsmE8Lexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(el)

	p := parser.NewAsmE8Parser(stream)
	p.RemoveErrorListeners()
	p.AddErrorListener(el)
	tree := p.Instruction()

	assembler.AttachErrorListener(el)
	antlr.ParseTreeWalkerDefault.Walk(assembler, tree)

	err := el.GetError()
	if err != nil {
		return nil, err
	}

	out, _, err := assembler.Assemble()
	if err != nil {
		return nil, err
	}
	return out, nil
}

func ToHexArray(arr []uint8) string {
	out := "["
	for i, b := range arr {
		out += fmt.Sprintf("%02x", b)
		if i != len(arr)-1 {
			out += " "
		}
	}
	out += "]"
	return out
}
