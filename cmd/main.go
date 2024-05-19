package main

import (
	"emu/src/comp"
)

func main() {
	c := comp.New()
	program := []byte{
		// mov a, #30
		0x41, 0x00, 0x30,
		// mov b, #31
		0x41, 0x01, 0x31,
		// mov c, a
		0x40, 0x02,
		// add c, a
		0x30, 0x21,
	}
	c.SetDebug(true)
	c.SetVerbose(true)
	c.Put(0, program)
	c.Run()
}
