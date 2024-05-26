package main

import (
	"emu/src/assembler"
	"os"
)

func main() {
	exampleFilePath := "src/assembler/example.as"
	file, err := os.ReadFile(exampleFilePath)
	if err != nil {
		panic(err)
	}
	ass, err := assembler.New(file)
	if err != nil {
		panic(err)
	}
	ass.Test()
	// return
	// program := ass.Assemble()
	// c := comp.New()
	// c.Put(0, program)
	// c.SetDebug(true)
	// // c.SetVerbose(true)
	// c.Run()
}
