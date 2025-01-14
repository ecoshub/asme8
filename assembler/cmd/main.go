package main

import (
	"asme8/assembler/src/assembler"
	"encoding/binary"
	"flag"
	"fmt"
	"log"
	"os"
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

	assembler, err := assembler.AssembleFile(*flagFile)
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
		fmt.Println(assembler)
	}

	out, err := assembler.Out()
	if err != nil {
		fmt.Println(err)
		return
	}

	err = binary.Write(f, binary.BigEndian, out)
	if err != nil {
		fmt.Println("file write error", err)
		return
	}

	fmt.Printf("assemble success. %d bytes assembled. output file: '%s'\n", len(out), *flagOutput)

}
