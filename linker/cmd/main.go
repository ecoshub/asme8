package main

import (
	"asme8/common/config"
	"asme8/linker/src/core"
	"flag"
	"fmt"
	"path/filepath"
	"strings"
)

var (
	flagConfigPath   = flag.String("config", "", "Path to the memory config file")
	flagOutput       = flag.String("output", "a.bin", "output file name")
	flagPrintSymbols = flag.Bool("print", false, "print all symbol tables")
	flagOutCodePath  = flag.String("output-code", "a.sym", "write indexed code to a path")
)

func main() {
	flag.Parse()

	// Get remaining arguments as .o files
	objectFilePaths := flag.Args()

	// Check if config path is provided
	if *flagConfigPath == "" {
		fmt.Println("Linker config path required")
		return
	}

	conf, err := config.ParseConfig(*flagConfigPath)
	if err != nil {
		fmt.Printf("Memory config parse failed. err: %s", err.Error())
		return
	}

	objectFiles, err := config.ReadObjectFiles(objectFilePaths...)
	if err != nil {
		fmt.Println(err)
		return
	}

	l := core.NewLinker(conf, objectFiles...)

	err = l.Link()
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = l.Out(*flagOutput)
	if err != nil {
		fmt.Println(err)
		return
	}

	if *flagPrintSymbols {
		l.PrintSymbols(false)
	}

	if *flagOutput != "a.bin" {
		ext := filepath.Ext(*flagOutput)
		base := filepath.Base(*flagOutput)
		*flagOutCodePath = strings.TrimSuffix(base, ext) + ".sym"
	}

	_, err = l.OutCode(*flagOutCodePath)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("link success. files: %v, output file: %s\n", objectFilePaths, *flagOutput)

}
