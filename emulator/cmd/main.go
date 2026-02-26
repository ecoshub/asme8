package main

import (
	"asme8/assembler/src/utils"
	"asme8/common/config"
	"asme8/emulator/src/computer"
	"asme8/emulator/src/emulator"
	"flag"
	"fmt"
	"time"
)

var (
	flagFileBin    = flag.String("load-bin", "", "load bin file to rom")
	flagFileAsm    = flag.String("load-asm", "", "assemble and load asm file to rom")
	flagDelay      = flag.Duration("delay", 10*time.Millisecond, "delay between instruction execution cycle")
	flagHeadless   = flag.Bool("headless", false, "run computer as 'headless'")
	flagConfigPath = flag.String("config", "default_config", "Path to the memory config file")
	flagSymbolFile = flag.String("symbol-file", "", "Path to indexed code (symbol) file")
)

func main() {
	flag.Parse()

	if *flagConfigPath == "" {
		fmt.Println("config path required for memory mapping")
		return
	}

	conf, err := config.ParseConfig(*flagConfigPath)
	if err != nil {
		fmt.Printf("config parse failed. path: %s, err: %s", *flagConfigPath, err.Error())
		return
	}

	var program []uint8
	program, code, path, err := utils.ResolveProgram(*flagFileBin, *flagFileAsm)
	if err != nil {
		fmt.Println(err)
		return
	}

	c, err := computer.New(&computer.Config{
		MemoryConfig: conf.MemoryConfig,
		Headless:     *flagHeadless,
		Program:      program,
		Delay:        *flagDelay,
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	var codeLines map[uint16]string
	var hasSymbols bool
	if code != "" {
		codeLines, hasSymbols, err = emulator.ResolveSymbolFromFile(code)
	} else {
		codeLines, hasSymbols, err = emulator.ReadSymbolFileAndCreateCodePanel(path, *flagSymbolFile)
	}
	if err != nil {
		fmt.Println(err)
		return
	}
	e := emulator.New(c, &emulator.Config{
		Headless: *flagHeadless,
	})
	e.CreateCodePanel(codeLines, hasSymbols)
	e.Run()
}
