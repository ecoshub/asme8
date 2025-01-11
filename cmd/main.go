package main

import (
	"emu/src/comp"
	"emu/src/peripheral"
	"emu/src/ram"
	"emu/src/rom"
	"flag"
	"fmt"
	"os"
	"time"
)

var (
	flagFileBin = flag.String("file-bin", "", "bin file of program")
	flagDebug   = flag.Bool("debug", false, "enable debug mode")
	flagVerbose = flag.Bool("verbose", false, "enable verbosity")
	flagDelay   = flag.Duration("delay", 10*time.Millisecond, "delay between instruction execution cycle")
)

func main() {
	flag.Parse()

	if !*flagDebug {
		peripheral.TerminalClear()
		peripheral.TerminalGoToFirstBlock()
	}

	if *flagFileBin == "" {
		fmt.Println("error no input file. please provide a executable (bin)")
		return
	}

	program, err := os.ReadFile(*flagFileBin)
	if err != nil {
		fmt.Println("executable read error", err)
		return
	}

	rom := rom.New(0x8000)
	rom.Load(0, program)

	ram := ram.New(0x8000)

	c := comp.New()
	c.ConnectDevice(rom, 0x0000, 0x8000)
	c.ConnectDevice(ram, 0x8000, 0xffff)
	c.SetDebug(*flagDebug)
	c.SetVerbose(*flagVerbose)
	c.SetDelayMS(*flagDelay)
	c.Run()

}
