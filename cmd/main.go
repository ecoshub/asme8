package main

import (
	"emu/src/comp"
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

	// peripheral.TerminalClear()
	// peripheral.TerminalCharOut(5, 20, '<')
	// peripheral.TerminalCharErase(5, 20)
	// peripheral.ListenKeys(func(e keyboard.KeyEvent) {
	// 	fmt.Println(e.Err, e.Key, string(e.Rune))
	// })
	// fmt.Println()
	// select {}

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

	c := comp.New()
	c.SetDebug(*flagDebug)
	c.SetVerbose(*flagVerbose)
	c.SetDelayMS(*flagDelay)
	c.Load(0, program)
	c.Run()
}
