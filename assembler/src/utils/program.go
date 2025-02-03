package utils

import (
	"asme8/assembler/src/assembler"
	"fmt"
	"os"
)

func ResolveProgram(binFilePath, asmFilePath string) ([]byte, string, error) {
	if binFilePath != "" {
		program, err := ReadProgram(binFilePath)
		if err != nil {
			return nil, "", err
		}
		return program, "", nil
	}
	if asmFilePath != "" {
		program, code, err := AssembleProgram(asmFilePath)
		if err != nil {
			return nil, "", err
		}
		return program, code, nil
	}
	return []byte{}, "", nil
}

func ReadProgram(binFilePath string) ([]uint8, error) {
	program, err := os.ReadFile(binFilePath)
	if err != nil {
		return nil, fmt.Errorf("executable read error. err: %s", err)
	}
	return program, nil
}

func AssembleProgram(asmFilePath string) ([]uint8, string, error) {
	program, code, err := assembler.AssembleFile(&assembler.Options{
		FilePath: asmFilePath,
		Mode:     assembler.ASM_MODE_EXE,
	})
	if err != nil {
		return nil, "", err
	}

	return program, code, nil
}
