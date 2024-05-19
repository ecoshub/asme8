package main

import (
	"emu/src/comp"
)

func main() {
	c := comp.New()
	program := []byte{
		0x41, 0x00, 0x30,
		0x41, 0x01, 0x31,
		0x40, 0x20,
		0x30, 0x21,
	}
	c.SetDebug(true)
	c.SetVerbose(true)
	c.Put(0, program)
	c.Run()
}
