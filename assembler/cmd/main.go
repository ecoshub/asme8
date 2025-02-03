package main

import (
	"asme8/assembler/src/assembler"
	"flag"
	"fmt"
	"path/filepath"
	"strings"
)

var (
	flagOutput      = flag.String("output", "a.bin", "output file name")
	flagFile        = flag.String("file", "", "input file")
	flagPrint       = flag.Bool("print", false, "print the code and the machine code")
	flagSegmentAddr = flag.Uint("segment-addr", 0, "segment start addr")
	flagMode        = flag.String("mode", "exe", "assembly mode exe | elf")
	flagOutCodePath = flag.String("output-code", "", "write indexed code to a path")
)

func main() {
	flag.Parse()

	if *flagFile == "" {
		fmt.Println("no input file given. --input <file_name>")
		return
	}

	mode := assembler.ASM_MODE(0)
	switch *flagMode {
	case "exe":
		mode = assembler.ASM_MODE_EXE
	case "elf":
		if *flagOutput == "a.bin" {
			ext := filepath.Ext(*flagFile)
			base := filepath.Base(*flagFile)
			*flagOutput = strings.TrimSuffix(base, ext) + ".o"
		}
		mode = assembler.ASM_MODE_ELF
	}

	if *flagOutCodePath == "" {
		ext := filepath.Ext(*flagFile)
		base := filepath.Base(*flagFile)
		*flagOutCodePath = strings.TrimSuffix(base, ext) + ".sym"
	}

	out, code, err := assembler.AssembleFile(&assembler.Options{
		FilePath:       *flagFile,
		PrintDetail:    *flagPrint,
		SegmentAddr:    uint16(*flagSegmentAddr),
		Mode:           mode,
		OutputCodePath: *flagOutCodePath,
		OutputFilePath: *flagOutput,
	})

	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = assembler.OutCode(*flagOutCodePath, code)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("assemble success. %d bytes assembled. output file: '%s'\n", len(out), *flagOutput)

}
