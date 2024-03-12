package main

import (
	"emu/src/comp"
)

func main() {
	c := comp.New()
	program := []byte{
		0x41, 0x00, 0x31,
		0x40, 0x10,
	}
	c.Put(0, program)
	c.Run()
}
