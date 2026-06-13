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
	flagFileBin    = flag.String("load-bin", "", "load bin file to emulator")
	flagFileAsm    = flag.String("load-asm", "", "assemble and load asm file to emulator")
	flagDelay      = flag.Duration("delay", 10*time.Millisecond, "delay between instruction execution cycle")
	flagHeadless   = flag.Bool("headless", false, "run computer as 'headless'")
	flagConfigPath = flag.String("linker-config", "", "Path to linker configfile")
	flagSymbolFile = flag.String("symbol-file", "", "Path to indexed code (symbol) file")
)

func main() {
	flag.Parse()

	var conf *config.Config
	var err error
	if *flagConfigPath == "" {
		conf = config.DefaultConfig
	} else {
		conf, err = config.ParseConfigPath(*flagConfigPath)
		if err != nil {
			fmt.Printf("config parse failed\npath: %s\nerr: %s\n", *flagConfigPath, err.Error())
			return
		}
	}

	program, code, path, err := utils.ResolveProgram(*flagFileBin, *flagFileAsm)
	if err != nil {
		fmt.Println(err)
		return
	}

	c, err := computer.New(&computer.Config{
		MemorySegment: conf,
		Headless:      *flagHeadless,
		Delay:         *flagDelay,
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	err = c.LoadProgram(program)
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
