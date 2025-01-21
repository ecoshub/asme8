package utils

import (
	"asme8/assembler/src/assembler"
	"fmt"
	"os"
)

func ResolveProgram(binFilePath, asmFilePath string) ([]byte, error) {
	if binFilePath != "" {
		program, err := ReadProgram(binFilePath)
		if err != nil {
			return nil, err
		}
		return program, nil
	}
	if asmFilePath != "" {
		program, err := AssembleProgram(asmFilePath)
		if err != nil {
			return nil, err
		}
		return program, nil
	}
	return []byte{}, nil
}

func ReadProgram(binFilePath string) ([]uint8, error) {
	program, err := os.ReadFile(binFilePath)
	if err != nil {
		return nil, fmt.Errorf("executable read error. err: %s", err)
	}
	return program, nil
}

func AssembleProgram(asmFilePath string) ([]uint8, error) {
	fmt.Println("path", asmFilePath)
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
