package main

import (
	"emu/src/comp"
	"flag"
	"fmt"
	"os"
)

var (
	flagFileBin = flag.String("file-bin", "", "bin file of program")
)

func main() {
	flag.Parse()

	if *flagFileBin == "" {
		fmt.Println("error no input file. please provide a executable (bin)")
		return
	}

	program, err := os.ReadFile(*flagFileBin)
	if err != nil {
		fmt.Println(err)
		return
	}

	c := comp.New()
	c.SetDebug(true)
	c.SetVerbose(true)
	c.SetDelayMS(100)
	c.Put(0, program)
	c.Run()
}
