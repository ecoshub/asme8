package main

import (
	"asme8/assembler/src/utils"
	"asme8/common/config"
	"asme8/emulator/src/comp"
	"flag"
	"fmt"
	"time"
)

var (
	flagFileBin          = flag.String("load-bin", "", "load bin file to rom")
	flagFileAsm          = flag.String("load-asm", "", "assemble and load asm file to rom")
	flagDelay            = flag.Duration("delay", 10*time.Millisecond, "delay between instruction execution cycle")
	flagHeadless         = flag.Bool("headless", false, "run computer as 'headless'")
	flagMemoryConfigPath = flag.String("memory-config", "", "Path to the memory config file")
	flagSymbolFile       = flag.String("symbol-file", "", "Path to indexed code (symbol) file")
)

func main() {
	flag.Parse()

	if *flagMemoryConfigPath == "" {
		fmt.Println("Memory config path required")
		return
	}

	memoryConfig, err := config.ParseMemoryConfig(*flagMemoryConfigPath)
	if err != nil {
		fmt.Printf("Memory config parse failed. err: %s", err.Error())
		return
	}

	var program []uint8
	var code string
	program, code, err = utils.ResolveProgram(*flagFileBin, *flagFileAsm)
	if err != nil {
		fmt.Println(err)
		return
	}

	c, err := comp.New(&comp.Config{
		MemoryConfig:   memoryConfig,
		Headless:       *flagHeadless,
		Program:        program,
		Delay:          *flagDelay,
		SymbolFilePath: *flagSymbolFile,
		SymbolFile:     code,
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	c.Run()

}
