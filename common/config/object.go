package config

import (
	"asme8/common/object"
	"os"
)

func ReadObjectFiles(paths ...string) ([]*object.ELF, error) {
	elfFiles := make([]*object.ELF, len(paths))
	for i, path := range paths {
		elf, err := readObjectFile(path)
		if err != nil {
			return nil, err
		}
		elfFiles[i] = elf
	}
	return elfFiles, nil
}

func readObjectFile(path string) (*object.ELF, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	elf := &object.ELF{}
	err = elf.Decode(file)
	if err != nil {
		return nil, err
	}

	return elf, nil
}
