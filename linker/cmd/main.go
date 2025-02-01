package main

import (
	"asme8/linker/src/config"
	"asme8/linker/src/core"
	"asme8/linker/src/object"
	"flag"
	"fmt"
	"log"
)

var (
	flagConfigPath   = flag.String("config", "", "Path to the linker config file")
	flagOutput       = flag.String("output", "a.bin", "output file name")
	flagPrintSymbols = flag.Bool("print", false, "print all symbol tables")
)

func main() {
	flag.Parse()

	// Get remaining arguments as .o files
	objectFilePaths := flag.Args()

	// Check if config path is provided
	if *flagConfigPath == "" {
		fmt.Println("Config path is required")
		return
	}

	// Load config
	conf := config.NewConfig()
	err := conf.ParseConfig(*flagConfigPath)
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	objectFiles, err := object.ReadObjectFiles(objectFilePaths...)
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
		l.PrintSymbols()
	}

	fmt.Printf("link success. files: %v, output file: %s\n", objectFilePaths, *flagOutput)

}
