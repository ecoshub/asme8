package assembler

import (
	"bytes"
	"fmt"
	"os"

	"github.com/ecoshub/machine"
)

const (
	AssemblerMachineFilePath string = "src/assembler/ass.mac"
)

type Assembler struct {
	data               []byte
	operations         []*Operation
	tempOperation      *Operation
	sectionDataDefined bool
	sectionBssDefined  bool
	sectionTextDefined bool
	global             string
	sm                 *machine.StateMachine
}

func New(data []byte) (*Assembler, error) {
	file, err := os.ReadFile(AssemblerMachineFilePath)
	if err != nil {
		return nil, err
	}
	sm, err := machine.Build(file)
	if err != nil {
		return nil, err
	}
	data = bytes.Replace(data, []byte{'\t'}, []byte{' ', ' ', ' ', ' '}, -1)
	return &Assembler{
		data:          data,
		sm:            sm,
		operations:    make([]*Operation, 0, 2),
		tempOperation: &Operation{},
	}, nil
}

func (a *Assembler) Test() {
	a.sm.SetDebug(true)
	a.setRecorders()
	err := a.sm.Match(string(a.data))
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, op := range a.operations {
		fmt.Printf(">%s<>%s<>%s<\n", op.mnemonic, op.operand1, op.operand2)
	}
	fmt.Println("OK")
}

func (a *Assembler) Assemble() []byte {
	// return []byte{}
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

func (a *Assembler) parseLine() {
}
