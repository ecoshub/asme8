package main

import (
	"asme8/common/config"
	"asme8/linker/src/core"
	"flag"
	"fmt"
)

var (
	flagMemoryConfigPath  = flag.String("memory-config", "", "Path to the memory config file")
	flagSegmentConfigPath = flag.String("segment-config", "", "Path to the segment config file")
	flagOutput            = flag.String("output", "a.bin", "output file name")
	flagPrintSymbols      = flag.Bool("print", false, "print all symbol tables")
)

func main() {
	flag.Parse()

	// Get remaining arguments as .o files
	objectFilePaths := flag.Args()

	// Check if config path is provided

	if *flagMemoryConfigPath == "" {
		fmt.Println("Memory config path required")
		return
	}

	if *flagSegmentConfigPath == "" {
		fmt.Println("Segment config path required")
		return
	}

	memoryConfig, err := config.ParseMemoryConfig(*flagMemoryConfigPath)
	if err != nil {
		fmt.Printf("Memory config parse failed. err: %s", err.Error())
		return
	}

	segmentConfig, err := config.ParseSegmentConfig(*flagSegmentConfigPath)
	if err != nil {
		fmt.Printf("Segment config parse failed. err: %s", err.Error())
		return
	}

	objectFiles, err := config.ReadObjectFiles(objectFilePaths...)
	if err != nil {
		fmt.Println(err)
		return
	}

	l := core.NewLinker(memoryConfig, segmentConfig, objectFiles...)
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
