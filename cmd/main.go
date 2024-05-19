package main

import (
	"emu/src/comp"
)

func main() {
	c := comp.New()
	program := []byte{
		0x41, 0x00, 0x30,
		0x40, 0x10,
		0x00,
	}
	c.SetDebug(true)
	c.SetVerbose(true)
	c.Put(0, program)
	c.Run()
}
