package main

import (
	"asme8/assembler/src/assembler"
	"asme8/assembler/src/utils"
	"encoding/binary"
	"flag"
	"fmt"
	"log"
	"os"
)

var (
	flagOutput      = flag.String("output", "a.bin", "output file name")
	flagFile        = flag.String("file", "", "input file")
	flagPrint       = flag.Bool("print", false, "print the code and the machine code")
	flagGlobals     = flag.String("globals", "", "globals file path")
	flagSegmentAddr = flag.Uint("segment-addr", 0, "segment start addr")
)

func main() {
	var globals utils.FlagArray
	flag.Var(&globals, "global", "global list")
	flag.Parse()

	if *flagGlobals != "" {
		gg, err := utils.ReadArrayFile(*flagGlobals)
		if err != nil {
			fmt.Println("globals file read error. err", err)
			return
		}
		globals = append(globals, gg...)
	}

	if *flagFile == "" {
		fmt.Println("no input file given. --input <file_name>")
		return
	}

	out, err := assembler.AssembleFile(&assembler.Options{
		FilePath:    *flagFile,
		PrintDetail: *flagPrint,
		Globals:     globals,
		SegmentAddr: uint16(*flagSegmentAddr),
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	f, err := os.Create(*flagOutput)
	if err != nil {
		log.Fatal("Couldn't open file", err)
	}
	defer f.Close()

	err = binary.Write(f, binary.BigEndian, out)
	if err != nil {
		fmt.Println("file write error", err)
		return
	}

	fmt.Printf("assemble success. %d bytes assembled. output file: '%s'\n", len(out), *flagOutput)

}
