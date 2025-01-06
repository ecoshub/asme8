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

	c := comp.New()
	program := playGroundJump()
	if *flagFileBin != "" {
		var err error
		program, err = os.ReadFile(*flagFileBin)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	fmt.Println(program)

	c.SetDebug(true)
	c.SetVerbose(true)
	c.SetDelayMS(100)
	c.Put(0, program)
	c.Run()
}

func playGroundJump() []byte {
	return []byte{
		// mov a, #0
		0x41, 0x00, 0x00,
		// mov b, #1
		0x41, 0x01, 0x01,
		// add a, b
		0x30, 0x10,
		// jmp #6
		0x10, 0x06, 0x00,
	}
}

func playGround() []byte {
	return []byte{
		// mov a, #30
		0x41, 0x00, 0x30,
		// mov b, #31
		0x41, 0x01, 0x31,
		// mov c, a
		0x40, 0x20,
		// mov c, a
		0x30, 0x21,
	}
}
