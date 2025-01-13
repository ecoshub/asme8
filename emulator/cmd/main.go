package main

import (
	"asme8/emulator/src/comp"
	"asme8/emulator/src/keyboard"
	"asme8/emulator/src/ram"
	"asme8/emulator/src/rom"
	"asme8/emulator/src/video"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	flagFileBin        = flag.String("file-bin", "", "bin file of program")
	flagDebug          = flag.Bool("debug", false, "enable debug mode")
	flagVerbose        = flag.Bool("verbose", false, "enable verbosity")
	flagDelay          = flag.Duration("delay", 10*time.Millisecond, "delay between instruction execution cycle")
	flagEnableVideo    = flag.Bool("enable-video", false, "enable video output")
	flagEnableKeyboard = flag.Bool("enable-keyboard", false, "enable keyboard output")
)

func main() {

	flag.Parse()

	if *flagFileBin == "" {
		fmt.Println("error no input file. please provide a executable (bin)")
		return
	}

	program, err := os.ReadFile(*flagFileBin)
	if err != nil {
		fmt.Println("executable read error", err)
		return
	}

	keyboard := keyboard.New()

	rom := rom.New(0x8000)
	rom.Load(0, program)

	// NOTE: for testing offset memory read logic. can be removed
	ram := ram.New(0x8000)
	ram.Load(0x10ff, []byte{0x90, 0x91, 0x92, 0x93, 0x94})

	vram := video.New()

	c := comp.New()
	c.ConnectDevice(rom, 0x0000, 0x6fff)
	c.ConnectDevice(vram, 0x7000, 0x70ff)
	c.ConnectDevice(keyboard, 0x7100, 0x7101)
	c.ConnectDevice(ram, 0x8000, 0xffff)

	c.SetDebug(*flagDebug)
	c.SetVerbose(*flagVerbose)
	c.SetDelay(*flagDelay)

	if *flagEnableVideo {
		vram.Reset()
		go vram.Run()
	}

	if *flagEnableKeyboard {
		go keyboard.ListenKeys()
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

	vram.ResetScreen()

}
