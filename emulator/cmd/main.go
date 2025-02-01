package main

import (
	"asme8/assembler/src/utils"
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
	flagFileBin  = flag.String("load-bin", "", "load bin file to rom")
	flagFileAsm  = flag.String("load-asm", "", "assemble and load asm file to rom")
	flagDelay    = flag.Duration("delay", 10*time.Millisecond, "delay between instruction execution cycle")
	flagHeadless = flag.Bool("headless", false, "run computer as 'headless'")
)

func main() {
	flag.Parse()

	var program []uint8
	var err error

	program, err = utils.ResolveProgram(*flagFileBin, *flagFileAsm)
	if err != nil {
		fmt.Println(err)
		return
	}

	if *flagHeadless {
		HeadlessMode(program)
		return
	}

	EmulatorMode(program)

}

func EmulatorMode(program []byte) {

	// create the rom that holds the program
	rom := rom.New(0x2000)

	// create new raw
	ram := ram.New(0xd748)

	term, err := terminal.New(0x780)
	if err != nil {
		fmt.Println(err)
		return
	}

	c := comp.New(rom)

	// screen components such as panels attaching to computer
	c.AttachTerminalComponents(term.Components)

	// print help strings on to panel
	c.Help()

	// connecting rom at 0x0000 through 0x7fff
	c.ConnectDevice(rom, 0x0000, 0x2000)

	// connecting ram at 0x8000 through 0xffff
	c.ConnectDevice(ram, 0x20ff, 0xd748)

	// 80 x 24 bytes need to hold the screen buffer
	c.ConnectDevice(term.Screen, 0xf847, 0x780)

	// keyboard needs only 8 bytes to hold char and state
	c.ConnectDevice(term.Keyboard, 0xffe7, 8)

	// setting base clock speed
	c.SetDelay(*flagDelay)

	// start emulator with computer pause mode
	c.SetPause(true)

	// load program in to rom starting from 0x0000
	programStartAddr := 0
	if len(program) != 0 {
		rom.Load(programStartAddr, program)
		c.ProgramLoaded()
		c.Logf("-> Program loaded in to ROM. Start addr is 0x%04x and its %d bytes", programStartAddr, len(program))
	}

	// running terminal (screen and keyboard)
	go term.Run()

	// running computer
	go c.Run()
	c.LogWithStyle("TIP: type 's' and press 'ENTER' to start the program", comp.DefaultStyle6)

	// standard main thread blocker
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop

	// restore terminal values before termination
	terminal.ResetScreen()

}

func HeadlessMode(program []byte) {

	// create the rom that holds the program
	rom := rom.New(0x2000)

	// create new raw
	ram := ram.New(0xd748)

	c := comp.New(rom)

	c.ConnectDevice(rom, 0x0000, 0x2000)

	c.ConnectDevice(ram, 0x20ff, 0xd748)

	// setting base clock speed
	c.SetDelay(*flagDelay)

	// setting debug flag to output logs in to stdout
	c.SetDebug(true)

	// start as running
	c.SetPause(false)

	// load program in to rom starting from 0x0000
	programStartAddr := 0
	if len(program) != 0 {
		rom.Load(programStartAddr, program)
		c.ProgramLoaded()
		c.Logf("-> Program loaded in to ROM. Start addr is 0x%04x and its %d bytes", programStartAddr, len(program))
	}
	// running computer
	go c.Run()

	// standard main thread blocker
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop
}
