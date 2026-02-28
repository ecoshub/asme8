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
		fmt.Println("Linker config path required. use --config <config_path>")
		return
	}

	conf, err := config.ParseConfig(*flagConfigPath)
	if err != nil {
		fmt.Printf("Config parse failed. err: %s\n", err.Error())
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

	if l.IsMemoryClean() {
		fmt.Println("Warning: The linker could not proceed because the 'SEGMENTS' configuration is missing or invalid.")
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

	fmt.Printf("Link success. files: %v, output file: %s\n", objectFilePaths, *flagOutput)

}
