package main

import (
	"asme8/assembler/src/assembler"
	"asme8/emulator/src/comp"
	"asme8/emulator/src/ram"
	"asme8/emulator/src/rom"
	"asme8/emulator/src/terminal"
	"errors"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	flagFileBin = flag.String("file-bin", "", "bin file path")
	flagFileAsm = flag.String("file-asm", "", "asm file path")
	flagDelay   = flag.Duration("delay", 10*time.Millisecond, "delay between instruction execution cycle")
)

func main() {
	flag.Parse()

	var program []uint8
	var err error

	if *flagFileBin == "" && *flagFileAsm == "" {
		err := errors.New("error. No input file. Please provide an executable (bin), or and assemble file (asm)")
		fmt.Println(err)
		flag.PrintDefaults()
		return
	}

	program, err = ResolveProgram(*flagFileBin, *flagFileAsm)
	if err != nil {
		fmt.Println(err)
		return
	}

	c := comp.New()
	// create the rom that holds the program
	rom := rom.New(0x8000)

	// create new raw
	ram := ram.New(0x8000)

	term, err := terminal.New()
	if err != nil {
		fmt.Println(err)
		return
	}

	// screen components such as panels attaching to computer
	c.AttachTerminalComponents(term.Components)

	// print help strings on to panel
	c.Help()

	// connecting rom at 0x0000 through 0x7fff
	c.ConnectDevice(rom, 0x0000, 0x7fff)

	// connecting ram at 0x8000 through 0xffff
	c.ConnectDevice(ram, 0x8000, 0xffff)

	// keyboard needs only 2 bytes to hold char and state
	c.ConnectDevice(term.Keyboard, 0x6ffe, 0x6fff)

	// connecting terminal screen in at addr 0x7000 through 0x7780
	// 80 x 24 bytes need to hold the screen buffer
	c.ConnectDevice(term.Screen, 0x7000, 0x7780)

	// setting base clock speed
	c.SetDelay(*flagDelay)

	// start emulator with computer pause mode
	c.SetPause(true)

	// load program in to rom starting from 0x0000
	programStartAddr := 0
	rom.Load(programStartAddr, program)
	c.Logf("-> Program loaded in to ROM. Start addr is 0x%04x and its %d bytes", programStartAddr, len(program))

	// running terminal (screen and keyboard)
	go term.Run()

	// running computer
	go c.Run()
	c.Log("? prompt 's' to start the program")

	// standard main thread blocker
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop

	// restore terminal values before termination
	terminal.ResetScreen()

}

func ResolveProgram(binFilePath, asmFilePath string) ([]uint8, error) {
	if *flagFileBin != "" {
		program, err := ReadProgram(*flagFileBin)
		if err != nil {
			return nil, err
		}
		return program, nil
	}
	if *flagFileAsm != "" {
		program, err := AssembleProgram(*flagFileAsm)
		if err != nil {
			return nil, err
		}
		return program, nil
	}
	return nil, errors.New("must provide a .asm or .bin file to start")
}

func ReadProgram(binFilePath string) ([]uint8, error) {
	program, err := os.ReadFile(binFilePath)
	if err != nil {
		return nil, fmt.Errorf("executable read error. err: %s", err)
	}
	return program, nil
}

func AssembleProgram(asmFilePath string) ([]uint8, error) {
	assembler, err := assembler.AssembleFile(asmFilePath)
	if err != nil {
		return nil, err
	}
	program, err := assembler.Out()
	if err != nil {
		return nil, fmt.Errorf("file assemble error. err: %s", err)
	}
	return program, nil
}
