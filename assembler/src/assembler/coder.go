package assembler

import (
	"encoding/binary"
	"fmt"
	"os"
)

func (a *Assembler) CodeParseEnterInst(text string) {
	a.Coder.lineBeginnings = append(a.Coder.lineBeginnings, a.offset)
}

func (a *Assembler) CodeParseExitInst(text string) {
	a.Coder.instructions = append(a.Coder.instructions, text)
	a.Coder.lineEndings = append(a.Coder.lineEndings, a.offset)
}

func (a *Assembler) CodeParseLastLine(text string) {
	if text == "\n" && a.Coder.lastLine == "\n" {
		a.Coder.blanksLines = append(a.Coder.blanksLines, a.offset)
	}
	a.Coder.lastLine = text
}

func OutCode(path string, code string) (int, error) {
	f, err := os.Create(path)
	if err != nil {
		return 0, fmt.Errorf("open file error. err: %s", err)
	}
	defer f.Close()

	err = os.WriteFile(path, []byte(code), 0644)
	if err != nil {
		return 0, fmt.Errorf("file write error. err: %s", err)
	}

	return len(code), nil
}

func (a *Assembler) WriteBinaryFile(path string, out []byte) error {
	if path == "" {
		return nil
	}

	f, err := os.Create(path)
	if err != nil {
		fmt.Println(path)
		return err
	}
	defer f.Close()

	err = binary.Write(f, binary.BigEndian, out)
	if err != nil {
		return err
	}
	return nil
}
