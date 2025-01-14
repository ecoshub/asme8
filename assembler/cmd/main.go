package main

import (
	"asme8/assembler/src/parser"
	"encoding/binary"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/antlr4-go/antlr/v4"
)

var (
	flagOutput = flag.String("output", "a.bin", "output file name")
	flagFile   = flag.String("file", "", "input file")
	flagPrint  = flag.Bool("print", false, "print the code and the machine code")
)

func main() {
	flag.Parse()

	if *flagFile == "" {
		fmt.Println("no input file given. --input <file_name>")
		return
	}

	input, err := antlr.NewFileStream(*flagFile)
	if err != nil {
		fmt.Println("file stream error", err)
		return
	}
	lexer := parser.NewAsmE8Lexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewAsmE8Parser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	tree := p.Instruction()
	l := parser.NewAssembler()
	antlr.ParseTreeWalkerDefault.Walk(l, tree)

	out, err := l.Out()
	if err != nil {
		fmt.Println(err)
		return
	}

	f, err := os.Create(*flagOutput)
	if err != nil {
		log.Fatal("Couldn't open file", err)
	}
	defer f.Close()

	if *flagPrint {
		fmt.Println(l)
	}

	err = binary.Write(f, binary.BigEndian, out)
	if err != nil {
		log.Fatal("Write failed")
	}

	fmt.Printf("assemble success. %d bytes assembled. output file: '%s'\n", len(out), *flagOutput)

}
