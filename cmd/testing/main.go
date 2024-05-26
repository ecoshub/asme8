package main

import (
	"emu/src/comp"
)

func main() {
	c := comp.New()
	program := playGroundJump()
	c.SetDebug(true)
	c.SetVerbose(true)
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
		0x10, 0x00, 0x06,
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
