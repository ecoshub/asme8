package main

import (
	"asme8/assembler/src/assembler"
	"asme8/emulator/src/comp"
	"asme8/emulator/src/ram"
	"asme8/emulator/src/rom"
	"asme8/emulator/src/terminal"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	flagFileBin        = flag.String("file-bin", "", "bin file path")
	flagFileAsm        = flag.String("file-asm", "", "asm file path")
	flagDebug          = flag.Bool("debug", false, "enable debug mode")
	flagVerbose        = flag.Bool("verbose", false, "enable verbosity")
	flagDelay          = flag.Duration("delay", 10*time.Millisecond, "delay between instruction execution cycle")
	flagEnableTerminal = flag.Bool("enable-term", false, "enable terminal. screen, panels and keyboard")
)

func main() {
	flag.Parse()

	var program []uint8
	var err error

	if *flagFileBin != "" {
		program, err = os.ReadFile(*flagFileBin)
		if err != nil {
			fmt.Println("executable read error", err)
			return
		}
	} else {
		if *flagFileAsm != "" {
			assembler, err := assembler.AssembleFile(*flagFileAsm)
			if err != nil {
				fmt.Println(err)
				return
			}
			program, err = assembler.Out()
			if err != nil {
				fmt.Println(err)
				return
			}
		} else {
			fmt.Println("Error. No input file. Please provide an executable (bin), or and assemble file (asm)")
			flag.PrintDefaults()
			os.Exit(0)
		}
	}

	rom := rom.New(0x8000)
	rom.Load(0, program)

	// NOTE: for testing offset memory read logic. can be removed
	ram := ram.New(0x8000)
	ram.Load(0x10ff, []byte{0x90, 0x91, 0x92, 0x93, 0x94})

	c := comp.New()

	var term *terminal.Terminal
	if *flagEnableTerminal {
		term, err = terminal.New()
		if err != nil {
			fmt.Println(err)
			return
		}
		c.AttachTerminalComponents(term.Components)

		c.ConnectDevice(term.Screen, 0x7000, 0x70ff)
		c.ConnectDevice(term.Keyboard, 0x7100, 0x7101)
		*flagDebug = true
	}

	c.ConnectDevice(rom, 0x0000, 0x6fff)
	c.ConnectDevice(ram, 0x8000, 0xffff)

	c.SetDebug(*flagDebug)
	c.SetVerbose(*flagVerbose)
	c.SetDelay(*flagDelay)

	if *flagEnableTerminal {
		go term.Run()
	}

	chDone := make(chan struct{}, 1)
	go func() {
		c.Run()
		chDone <- struct{}{}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	select {
	case <-stop:
	case <-chDone:
	}

	terminal.ResetScreen()

}
