package assembler

import (
	"asme8/assembler/src/error_listener"
	"asme8/assembler/src/parser"
	"asme8/assembler/src/types"
	"asme8/common/instruction"
	"asme8/common/object"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type ASM_MODE uint8

const (
	VERSION uint16 = 1
)

const (
	ASM_MODE_NONE ASM_MODE = 0
	ASM_MODE_EXE  ASM_MODE = 1
	ASM_MODE_ELF  ASM_MODE = 2
)

var _ parser.AsmE8Listener = &Assembler{}

type Assembler struct {
	mode                ASM_MODE
	offset              uint16
	currentInstruction  string
	currentValue        *types.Value
	reference           *types.Variable
	currentValueList    []*types.Value
	currentTag          *types.Tag
	currentRegisters    []*types.Register
	labels              map[string]*types.Label
	variables           map[string]*types.Variable
	unresolvedVariables map[string]*types.Variable
	missingSymbols      map[uint16]*types.Tag
	machineCode         map[uint16]uint8
	directives          map[uint16]*types.Directive
	errorListener       *error_listener.CustomErrorListener
	symbolTracker       *object.Tracker
	hasVirtualOffset    bool
	out                 []uint8
	topAddress          uint16

	Coder *Coder
}

type Coder struct {
	linesEndings []uint16
	skip         []uint16
	blanksLines  []uint16
	instructions []string
	lastLine     string
}

func New(mode ASM_MODE) *Assembler {
	return &Assembler{
		mode:                mode,
		currentValueList:    []*types.Value{},
		currentRegisters:    []*types.Register{},
		labels:              make(map[string]*types.Label),
		variables:           make(map[string]*types.Variable),
		unresolvedVariables: make(map[string]*types.Variable),
		machineCode:         make(map[uint16]uint8),
		missingSymbols:      make(map[uint16]*types.Tag),
		directives:          make(map[uint16]*types.Directive),
		symbolTracker:       object.NewTracker(),
		Coder: &Coder{
			linesEndings: make([]uint16, 0, 16),
			blanksLines:  make([]uint16, 0, 16),
			skip:         make([]uint16, 0, 16),
			instructions: make([]string, 0, 16),
		},
	}
}

func (a *Assembler) Assemble() ([]uint8, int, error) {
	switch a.mode {
	case ASM_MODE_ELF:
		a.RestoreMissingSymbols()
		a.ResolveReferenceVariables()
		out := a.assemble()
		err := a.errorListener.GetError()
		if err != nil {
			return nil, 0, err
		}
		a.symbolTracker.AttachBin(out)
		a.symbolTracker.AttachCode(a.CodeString())
		elf := object.ELF{
			Header:  object.NewStdHeader(VERSION),
			Tracker: a.symbolTracker,
		}
		enc, err := elf.Encode()
		if err != nil {
			return nil, 0, err
		}
		return enc, len(out), nil
	case ASM_MODE_EXE:
		errs := a.ResolveReferenceVariables()
		if len(errs) > 0 {
			return nil, 0, error_listener.WrapErrors(errs...)
		}
		errs = a.RestoreMissingSymbols()
		if len(errs) > 0 {
			return nil, 0, error_listener.WrapErrors(errs...)
		}
		out := a.assemble()
		err := a.errorListener.GetError()
		if err != nil {
			return nil, 0, err
		}
		return out, len(out), nil
	}
	return nil, 0, errors.New("unknown assembly mode")
}

func (a *Assembler) Reset() {
	a.offset = 0
	a.currentInstruction = ""
	a.currentValue = nil
	a.reference = nil
	a.currentTag = nil
	a.hasVirtualOffset = false
	a.out = make([]uint8, 0, 32)
	a.topAddress = 0
	a.currentValueList = []*types.Value{}
	a.currentRegisters = []*types.Register{}
	a.labels = make(map[string]*types.Label)
	a.variables = make(map[string]*types.Variable)
	a.unresolvedVariables = make(map[string]*types.Variable)
	a.machineCode = make(map[uint16]uint8)
	a.missingSymbols = make(map[uint16]*types.Tag)
	a.directives = make(map[uint16]*types.Directive)
	a.symbolTracker = object.NewTracker()
	a.Coder = &Coder{
		linesEndings: make([]uint16, 0, 16),
		blanksLines:  make([]uint16, 0, 16),
		skip:         make([]uint16, 0, 16),
		instructions: make([]string, 0, 16),
	}
}

func (a *Assembler) AttachErrorListener(errorListener *error_listener.CustomErrorListener) {
	a.errorListener = errorListener
}

func (a *Assembler) assemble() []uint8 {
	topAddress := uint16(0)
	for offset := range a.machineCode {
		if offset >= topAddress {
			topAddress = offset
		}
	}
	if len(a.machineCode) == 0 {
		a.out = []byte{}
		a.topAddress = 0
		return a.out
	}
	out := make([]uint8, topAddress+1)
	for i := uint16(0); i < topAddress+1; i++ {
		b, ok := a.machineCode[i]
		if !ok {
			a.Coder.skip = append(a.Coder.skip, i)
			out[i] = 0
			continue
		}
		out[i] = b
	}
	a.topAddress = topAddress
	a.out = out
	return a.out
}

func (a *Assembler) AppendMachineCode(code uint8) {
	a.machineCode[a.offset] = code
	a.offset++
}

// func (a *Assembler) ParseIndexImmediate(line, column int) {
// 	if a.currentValue.GetSize() == 8 {
// 		a.GetVariableOrTagMissing(1, 8)
// 		opcode := a.GetOrFailOpCode(a.currentInstruction, instruction.ADDRESSING_MODE_IP_IMM, line, column)
// 		a.AppendMachineCode(opcode)
// 		a.AppendMachineCode(a.currentValue.GetLowByte())
// 		return
// 	}
// 	a.GetVariableOrTagMissing(1, 16)
// 	opcode := a.GetOrFailOpCode(a.currentInstruction, instruction.ADDRESSING_MODE_IP_IMM_16, line, column)
// 	a.AppendMachineCode(opcode)
// 	a.AppendMachineCode(a.currentValue.GetLowByte())
// 	a.AppendMachineCode(a.currentValue.GetHighByte())
// }

func (a *Assembler) ParseStackImm8(line, column int) {
	opcode := a.GetOrFailOpCode(a.currentInstruction, instruction.ADDR_MODE_STACK_IMM8, line, column)
	a.AppendMachineCode(opcode)
	a.AppendMachineCode(a.currentValue.GetLowByte())
}

func (a *Assembler) ParseReg8Imm8(line, column int) {
	opcode := a.GetOrFailOpCode(a.currentInstruction, instruction.ADDR_MODE_REG8_IMM8, line, column)
	a.AppendMachineCode(opcode)
	a.AppendMachineCode(a.currentRegisters[0].GetCode())
	a.AppendMachineCode(a.currentValue.GetLowByte())
}

func (a *Assembler) ParseReg8Reg8(line, column int) {
	var val uint8 = a.currentRegisters[0].GetCode() | (a.currentRegisters[1].GetCode() << 4)
	opcode := a.GetOrFailOpCode(a.currentInstruction, instruction.ADDR_MODE_REG8_REG8, line, column)
	a.AppendMachineCode(opcode)
	a.AppendMachineCode(val)
}

func (a *Assembler) ParseReg16Reg16(line, column int) {
	var val uint8 = a.currentRegisters[0].GetCode() | (a.currentRegisters[1].GetCode() << 4)
	opcode := a.GetOrFailOpCode(a.currentInstruction, instruction.ADDR_MODE_REG16_REG16, line, column)
	a.AppendMachineCode(opcode)
	a.AppendMachineCode(val)
}

func (a *Assembler) ParseReg16Imm(line, column int) {
	reg := a.currentRegisters[0].GetCode()
	if a.currentInstruction == instruction.INST_MNEMONICS_ADD || a.currentInstruction == instruction.INST_MNEMONICS_SUB {
		opcode := a.GetOrFailOpCode(a.currentInstruction, instruction.ADDR_MODE_REG16_IMM8, line, column)
		a.AppendMachineCode(opcode)
		a.AppendMachineCode(reg)
		a.AppendMachineCode(a.currentValue.GetLowByte())
		return
	}
	opcode := a.GetOrFailOpCode(a.currentInstruction, instruction.ADDR_MODE_REG16_IMM16, line, column)
	a.AppendMachineCode(opcode)
	a.AppendMachineCode(reg)
	a.AppendMachineCode(a.currentValue.GetLowByte())
	a.AppendMachineCode(a.currentValue.GetHighByte())
}

func (a *Assembler) ParseDirectRegToMem(line, column int) {
	opcode := a.GetOrFailOpCode(a.currentInstruction, instruction.ADDR_MODE_REG_TO_MEM_DIRECT, line, column)
	a.AppendMachineCode(opcode)
	a.AppendMachineCode(a.currentRegisters[0].GetCode())
	a.AppendMachineCode(a.currentValue.GetLowByte())
	a.AppendMachineCode(a.currentValue.GetHighByte())
}

func (a *Assembler) ParseDirectMemToReg(line, column int) {
	opcode := a.GetOrFailOpCode(a.currentInstruction, instruction.ADDR_MODE_MEM_TO_REG_DIRECT, line, column)
	a.AppendMachineCode(opcode)
	a.AppendMachineCode(a.currentRegisters[0].GetCode())
	a.AppendMachineCode(a.currentValue.GetLowByte())
	a.AppendMachineCode(a.currentValue.GetHighByte())
}

func (a *Assembler) ParseIndirectRegToMem(line, column int) {
	var val uint8 = a.currentRegisters[1].GetCode() | (a.currentRegisters[0].GetCode() << 4)
	opcode := a.GetOrFailOpCode(a.currentInstruction, instruction.ADDR_MODE_REG_TO_MEM_INDIRECT, line, column)
	a.AppendMachineCode(opcode)
	a.AppendMachineCode(val)
}

func (a *Assembler) ParseIndirectMemToReg(line, column int) {
	var val uint8 = a.currentRegisters[1].GetCode() | (a.currentRegisters[0].GetCode() << 4)
	opcode := a.GetOrFailOpCode(a.currentInstruction, instruction.ADDR_MODE_MEM_TO_REG_INDIRECT, line, column)
	a.AppendMachineCode(opcode)
	a.AppendMachineCode(val)
}

func (a *Assembler) ParseIndirectOffsetRegToMem(line, column int) {
	var val uint8 = a.currentRegisters[1].GetCode() | (a.currentRegisters[0].GetCode() << 4)
	opcode := a.GetOrFailOpCode(a.currentInstruction, instruction.ADDR_MODE_REG_TO_MEM_INDIRECT_OFFSET, line, column)
	a.AppendMachineCode(opcode)
	a.AppendMachineCode(val)
	a.AppendMachineCode(a.currentValue.GetLowByte())
}

func (a *Assembler) ParseIndirectOffsetMemToReg(line, column int) {
	var val uint8 = a.currentRegisters[1].GetCode() | (a.currentRegisters[0].GetCode() << 4)
	opcode := a.GetOrFailOpCode(a.currentInstruction, instruction.ADDR_MODE_MEM_TO_REG_INDIRECT_OFFSET, line, column)
	a.AppendMachineCode(opcode)
	a.AppendMachineCode(val)
	a.AppendMachineCode(a.currentValue.GetLowByte())
}

func (a *Assembler) ParseDirectOffsetRegToMem(line, column int) {
	a.GetVariableOrTagMissing(2, 16)
	var val uint8 = a.currentRegisters[1].GetCode() | (a.currentRegisters[0].GetCode() << 4)
	opcode := a.GetOrFailOpCode(a.currentInstruction, instruction.ADDR_MODE_REG_TO_MEM_INDIRECT_OFFSET, line, column)
	a.AppendMachineCode(opcode)
	a.AppendMachineCode(val)
	a.AppendMachineCode(a.currentValue.GetLowByte())
	a.AppendMachineCode(a.currentValue.GetHighByte())
}

func (a *Assembler) ParseDirectOffsetMemToReg(line, column int) {
	a.GetVariableOrTagMissing(2, 16)
	var val uint8 = a.currentRegisters[0].GetCode() | (a.currentRegisters[1].GetCode() << 4)
	opcode := a.GetOrFailOpCode(a.currentInstruction, instruction.ADDR_MODE_MEM_TO_REG_INDIRECT_OFFSET, line, column)
	a.AppendMachineCode(opcode)
	a.AppendMachineCode(val)
	a.AppendMachineCode(a.currentValue.GetLowByte())
	a.AppendMachineCode(a.currentValue.GetHighByte())
}

func (a *Assembler) ParseIndexedRegToMem(line, column int) {
	a.GetVariableOrTagMissing(2, 16)
	var val uint8 = a.currentRegisters[0].GetCode() | (a.currentRegisters[1].GetCode() << 4)
	opcode := a.GetOrFailOpCode(a.currentInstruction, instruction.ADDR_MODE_REG_TO_MEM_INDEXED, line, column)
	a.AppendMachineCode(opcode)
	a.AppendMachineCode(val)
	a.AppendMachineCode(a.currentValue.GetLowByte())
	a.AppendMachineCode(a.currentValue.GetHighByte())
}

func (a *Assembler) ParseIndexedMemToReg(line, column int) {
	a.GetVariableOrTagMissing(2, 16)
	var val uint8 = a.currentRegisters[1].GetCode() | (a.currentRegisters[0].GetCode() << 4)
	opcode := a.GetOrFailOpCode(a.currentInstruction, instruction.ADDR_MODE_MEM_TO_REG_INDEXED, line, column)
	a.AppendMachineCode(opcode)
	a.AppendMachineCode(val)
	a.AppendMachineCode(a.currentValue.GetLowByte())
	a.AppendMachineCode(a.currentValue.GetHighByte())
}

func (a *Assembler) ParseIndexedReg16RegToMem(line, column int) {
	var val uint8 = a.currentRegisters[1].GetCode() | (a.currentRegisters[0].GetCode() << 4)
	opcode := a.GetOrFailOpCode(a.currentInstruction, instruction.ADDR_MODE_REG_TO_MEM_REG16_INDEXED, line, column)
	a.AppendMachineCode(opcode)
	a.AppendMachineCode(val)
	a.AppendMachineCode(a.currentRegisters[2].GetCode())
}

func (a *Assembler) ParseIndexedReg16MemToReg(line, column int) {
	var val uint8 = a.currentRegisters[2].GetCode() | (a.currentRegisters[1].GetCode() << 4)
	opcode := a.GetOrFailOpCode(a.currentInstruction, instruction.ADDR_MODE_MEM_TO_REG_REG16_INDEXED, line, column)
	a.AppendMachineCode(opcode)
	a.AppendMachineCode(val)
	a.AppendMachineCode(a.currentRegisters[0].GetCode())
}

func (a *Assembler) ParseImpliedReg8(line, column int) {
	opcode := a.GetOrFailOpCode(a.currentInstruction, instruction.ADDR_MODE_IMPLIED_REG8, line, column)
	a.AppendMachineCode(opcode)
	a.AppendMachineCode(a.currentRegisters[0].GetCode())
}

func (a *Assembler) ParseImpliedReg16(line, column int) {
	opcode := a.GetOrFailOpCode(a.currentInstruction, instruction.ADDR_MODE_IMPLIED_REG16, line, column)
	a.AppendMachineCode(opcode)
	a.AppendMachineCode(a.currentRegisters[0].GetCode())
}

func (a *Assembler) ParseImpliedStack(line, column int) {
	opcode := a.GetOrFailOpCode(a.currentInstruction, instruction.ADDR_MODE_IMPLIED_STACK, line, column)
	a.AppendMachineCode(opcode)
}

func (a *Assembler) ParseStackReg16(line, column int) {
	opcode := a.GetOrFailOpCode(a.currentInstruction, instruction.ADDR_MODE_STACK_REG16, line, column)
	a.AppendMachineCode(opcode)
	a.AppendMachineCode(a.currentRegisters[1].GetCode())
}

func (a *Assembler) ParseReg16Stack(line, column int) {
	opcode := a.GetOrFailOpCode(a.currentInstruction, instruction.ADDR_MODE_REG16_STACK, line, column)
	a.AppendMachineCode(opcode)
	a.AppendMachineCode(a.currentRegisters[0].GetCode())
}

func (a *Assembler) ParseImplied(line, column int) {
	opcode := a.GetOrFailOpCode(a.currentInstruction, instruction.ADDR_MODE_IMPLIED, line, column)
	a.AppendMachineCode(opcode)
}

func (a *Assembler) ParseImpliedImm16(line, column int) {
	opcode := a.GetOrFailOpCode(a.currentInstruction, instruction.ADDR_MODE_IMPLIED_IMM16, line, column)
	a.AppendMachineCode(opcode)
	a.AppendMachineCode(a.currentValue.GetLowByte())
	a.AppendMachineCode(a.currentValue.GetHighByte())
}

func (a *Assembler) ParseImpliedTag(line, column int) {
	v, exists := a.FindVariable(a.currentTag.Text)
	if exists {
		a.currentValue = v
		a.ParseImpliedImm16(line, column)
		return
	}

	label, exist := a.FindLabel(a.currentTag.Text)
	if exist {
		a.symbolTracker.SymbolHit(label.Text, a.offset+1, 16, 0)
		// fmt.Printf("[*] SYMBOL: text:>%16s<, offset: %04x, line: %4d, column: %4d\n", label.Text, a.offset+1, label.Line, label.Column)
		a.currentValue = types.NewValue(int64(label.Offset))
		a.ParseImpliedImm16(line, column)
		return
	}

	// fmt.Printf("? symbol not found. tag:>%s<, offset: %d\n", a.currentTag.Text, a.offset)
	a.currentTag.Size = 16
	a.missingSymbols[a.offset+1] = a.currentTag
	a.currentValue = types.NewValue(0)

	a.ParseImpliedImm16(line, column)
}

func (a *Assembler) ParseLabel(text string, line, column int, disablePrint bool) {
	label := types.NewLabel(text, line, column, a.offset)
	label.DisablePrint = disablePrint
	a.labels[text] = label

	a.symbolTracker.SetIndex(text, a.offset)
	a.symbolTracker.SetType(a.currentTag.Text, object.SYMBOL_TYPE_LABEL)
}

func (a *Assembler) ParsePtr(text string, line, column int) {
	if a.hasVirtualOffset {
		a.ParsePtrVirtualOffset(text, line, column)
		a.hasVirtualOffset = false
		return
	}

	if text == "" {
		return
	}

	v, exists := a.FindVariable(text)
	if exists {
		a.symbolTracker.SymbolHit(a.currentTag.Text, a.offset+2, 16, 0)
		// fmt.Printf("[v] SYMBOL: text:>%16s<, offset: %04x, line: %4d, column: %4d\n", a.currentTag.Text, a.offset+2, 0, 0)
		a.currentValue = v
		return
	}

	number := strings.TrimPrefix(text, "[")
	number = strings.TrimSuffix(number, "]")
	val, ok := types.ParseValue(number)
	if ok {
		a.currentValue = val
		return
	}

	// fmt.Printf("? symbol not found. tag:>%s<, offset: %d\n", a.currentTag.Text, a.offset)
	a.currentTag.Size = 16
	a.missingSymbols[a.offset+2] = a.currentTag
	a.currentValue = types.NewValue(0)
}

func (a *Assembler) GetOrFailOpCode(inst string, mode uint8, line, column int) uint8 {
	opcode, ok := instruction.GetOpCode(inst, mode)
	if !ok {
		msg := fmt.Sprintf("opcode not found for this instruction and addressing mode. inst: '%s', mode: %0d", a.currentInstruction, mode)
		if a.errorListener == nil {
			panic(msg)
		} else {
			a.errorListener.NewSimpleError(msg, line, column)
			return 0
		}
	}
	return opcode
}

func (a *Assembler) ParsePtrVirtualOffset(text string, line, column int) {
	text = text[1 : len(text)-1]
	sign := 1
	tokens := strings.Split(text, "+")
	if len(tokens) == 1 {
		sign = -1
		tokens = strings.Split(text, "-")
	}
	tag := tokens[0]
	offsetString := tokens[1]
	offset, _ := strconv.Atoi(offsetString)
	// fmt.Printf("tag: >%s<, offset(%s): %d, op: %d\n", tag, offsetString, offset, sign)
	// t := types.Tag{Text: tag, Line: line, Column: column, OptionalOffset: offset}
	offset = offset * sign
	v, exists := a.FindVariable(tag)
	if exists {
		a.symbolTracker.SymbolHit(tag, a.offset+2, 16, uint16(offset))
		// fmt.Printf("[V] SYMBOL: text:>%16s<, offset: %04x, line: %4d, column: %4d, optionalOffset: %d\n", tag, a.offset+2, line, column, offset)
		a.currentValue = types.NewValue(int64(v.GetValue()) + int64(offset))
		// fmt.Println("val", a.currentValue.GetValue())
		return
	}
	// fmt.Printf("??? text:>%16s<, offset: %04x, line: %4d, column: %4d, optionalOffset: %d\n", tag, a.offset+2, line, column, offset)
	a.missingSymbols[a.offset+2] = &types.Tag{Text: tag, Line: line, Column: column, OptionalOffset: offset, Size: 16}
	a.currentValue = types.NewValue(0)
}

func (a *Assembler) ParseAscii(text string) {
	text = text[1 : len(text)-1]
	list := make([]*types.Value, 0, len(text))
	for i := range text {
		v := types.NewValue(int64(text[i]))
		list = append(list, v)
	}
	a.currentValueList = list
}

func (a *Assembler) ParseDirective(text string, line, column int) {
	if strings.HasPrefix(text, ".org") {
		if len(a.currentValueList) == 0 {
			msg := fmt.Sprintf("value is unknown. value: '%s'", text)
			if a.errorListener == nil {
				panic(msg)
			} else {
				a.errorListener.NewSimpleError(msg, line, column)
			}
			return
		}
		val := a.currentValueList[0]
		a.directives[a.offset] = &types.Directive{Raw: text, Type: ".org", Position: val.GetValue(), Offset: 0xffff, Values: a.currentValueList}
		a.offset = val.GetValue()
		return
	}
	if strings.HasPrefix(text, ".resb") {
		val := a.currentValueList[0]
		a.directives[a.offset] = &types.Directive{Raw: text, Type: ".resb", Position: a.offset, Offset: val.GetValue(), Values: a.currentValueList}
		for i := 0; i < int(val.GetValue()); i++ {
			a.AppendMachineCode(0)
		}
		return
	}
	if strings.HasPrefix(text, ".byte") {
		a.directives[a.offset] = &types.Directive{Raw: text, Type: ".byte", Position: a.offset, Offset: uint16(len(a.currentValueList)), Values: a.currentValueList}
		for _, v := range a.currentValueList {
			a.AppendMachineCode(v.GetLowByte())
		}
		return
	}
	if strings.HasPrefix(text, ".word") {
		a.directives[a.offset] = &types.Directive{Raw: text, Type: ".word", Position: a.offset, Offset: uint16(len(a.currentValueList) * 2), Values: a.currentValueList}
		for _, v := range a.currentValueList {
			a.AppendMachineCode(v.GetLowByte())
			a.AppendMachineCode(v.GetHighByte())
		}
		return
	}
	if strings.HasPrefix(text, ".asciiz") {
		a.currentValueList = append(a.currentValueList, types.NewValue(0))
		a.directives[a.offset] = &types.Directive{Raw: text, Type: ".asciiz", Position: a.offset, Offset: uint16(len(a.currentValueList)), Values: a.currentValueList}
		for _, v := range a.currentValueList {
			a.AppendMachineCode(v.GetLowByte())
		}
		return
	}
	if strings.HasPrefix(text, ".ascii") {
		a.directives[a.offset] = &types.Directive{Raw: text, Type: ".ascii", Position: a.offset, Offset: uint16(len(a.currentValueList)), Values: a.currentValueList}
		for _, v := range a.currentValueList {
			a.AppendMachineCode(v.GetLowByte())
		}
		return
	}
}

func (a *Assembler) ParseTag(text string, line, column int) {
	a.currentTag = &types.Tag{Text: text, Line: line, Column: column}
}

func (a *Assembler) ParseRegister(text string) {
	a.currentRegisters = append(a.currentRegisters, types.ParseRegister(text))
}

func (a *Assembler) ParseVariable(name string, val *types.Value) {
	v := types.NewVariable(name, val)
	a.variables[name] = v

	a.symbolTracker.SetIndex(name, v.Val.GetValue())
	a.symbolTracker.SetType(name, object.SYMBOL_TYPE_VAR)
}

func (a *Assembler) ParseReference(text string) {
	refName := text
	offset := 0
	sign := 1
	if strings.Contains(text, "+") {
		tokens := strings.Split(text, "+")
		refName = tokens[0]
		sign = 1
	}
	if strings.Contains(text, "-") {
		tokens := strings.Split(text, "-")
		refName = tokens[0]
		sign = -1
	}
	// parse the offset if there is
	if a.currentValue != nil {
		offset = int(a.currentValue.GetValue())
		offset *= sign
	}
	a.reference = types.NewVariable(refName, types.NewValue(0))
	a.reference.Offset = offset
	// fmt.Println("REFERENCE CREATED", a.currentTag.Text, refName, offset)
}

func (a *Assembler) ParseVariableReference() {
	v, ok := a.variables[a.reference.Name]
	if ok {
		val := types.NewValue(int64(v.Val.GetValue() + uint16(a.reference.Offset)))
		v := types.NewVariable(a.currentTag.Text, val)
		v.Unresolved = false
		v.ReferenceName = a.reference.Name
		v.Offset = a.reference.Offset
		a.variables[a.currentTag.Text] = v
		a.symbolTracker.SetIndex(a.currentTag.Text, v.Val.GetValue())
		a.symbolTracker.SetType(a.currentTag.Text, object.SYMBOL_TYPE_VAR)
		return
	}
	vr := types.NewVariable(a.currentTag.Text, types.NewValue(0))
	vr.Unresolved = true
	vr.ReferenceName = a.reference.Name
	vr.Offset = a.reference.Offset
	a.unresolvedVariables[a.currentTag.Text] = vr
	a.symbolTracker.SetReference(a.currentTag.Text, a.reference.Name, int32(a.reference.Offset))
	// fmt.Println("REFERENCE VARIABLE SET", a.currentTag.Text, a.reference.Name, a.reference.Offset)
}

func (a *Assembler) ParseInstruction(text string) {
	a.currentInstruction = strings.ToLower(text)
}

// func (a *Assembler) ParsePtrImm(text string, line, column int) {
// 	opcode := a.GetOrFailOpCode(a.currentInstruction, instruction.ADDRESSING_MODE_MEM_IMM, line, column)
// 	a.AppendMachineCode(opcode)
// 	a.AppendMachineCode(a.currentValueList[0].GetLowByte())
// 	a.AppendMachineCode(a.currentValue.GetLowByte())
// 	a.AppendMachineCode(a.currentValue.GetHighByte())
// }

func (a *Assembler) ParseValue(text string, line, column int) {
	var ok bool

	val, ok := types.ParseValue(text)
	if !ok {
		a.missingSymbols[a.offset] = &types.Tag{Text: text, Line: line, Column: column, Size: 0}
	} else {
		if a.currentValue != nil {
			a.currentValueList = append(a.currentValueList, val)
		} else {
			a.currentValue = val
		}
	}
}

func (a *Assembler) ParseValueList(text string, line, column int) {
	tokens := strings.Split(text, ", ")
	valueList := make([]*types.Value, 0, len(tokens))
	for _, tok := range tokens {
		tok = strings.TrimSpace(tok)
		val, ok := types.ParseValue(tok)
		if ok {
			valueList = append(valueList, val)
			continue
		}
		variable, ok := a.FindVariable(tok)
		if ok {
			a.symbolTracker.SymbolHit(tok, a.offset, uint8(variable.GetSize()), 0)
			valueList = append(valueList, variable)
		} else {
			a.missingSymbols[a.offset] = &types.Tag{Text: tok, Line: line, Column: column}
		}
	}
	a.currentValueList = valueList
}

func (a *Assembler) ParseSegment(text string) {
	segment := strings.TrimPrefix(text, ".segment \"")
	segment = strings.TrimSuffix(segment, "\"")
	a.symbolTracker.SetSegment(segment)
}

func (a *Assembler) ParseAccessModifier(text string) {
	if strings.HasPrefix(text, "    extern") {
		a.symbolTracker.SetExtern(a.currentTag.Text)
	}
	if strings.HasPrefix(text, "    global") {
		a.symbolTracker.SetGlobal(a.currentTag.Text)
	}
}

func (a *Assembler) ResetCurrentValues() {
	a.currentRegisters = []*types.Register{}
	a.currentValue = nil
	a.currentTag = &types.Tag{}
	a.currentInstruction = ""
	a.hasVirtualOffset = false
	a.currentValueList = make([]*types.Value, 0, 2)
}

func (a *Assembler) GetVariableOrTagMissing(offsetPlus uint16, requiredSize uint8) {
	if a.currentTag.Text == "" {
		return
	}
	variable, exists := a.FindVariable(a.currentTag.Text)
	if exists {
		a.symbolTracker.SymbolHit(a.currentTag.Text, a.offset+offsetPlus, requiredSize, 0)
		a.currentValue = types.NewValue(int64(variable.GetValue()))
		return
	}

	a.currentTag.Size = int8(requiredSize)
	a.missingSymbols[a.offset+offsetPlus] = a.currentTag
	a.currentValue = types.NewValue(0)
}

func (a *Assembler) RestoreMissingSymbols() []error {
	errs := make([]error, 0, 2)
	for offset, symbol := range a.missingSymbols {
		err := a.RestoreSymbol(symbol, offset)
		if err != nil {
			errs = append(errs, err)
		}
	}
	return errs
}

func (a *Assembler) ResolveReferenceVariables() []error {
	errs := make([]error, 0, 2)
	for _, uv := range a.unresolvedVariables {
		if !uv.Unresolved {
			continue
		}
		// fmt.Println("UNRESOLVED VARIABLE FOUND", uv.Name)
		found := false
		var ref *types.Variable
		for _, v := range a.variables {
			if uv.ReferenceName == v.Name {
				found = true
				ref = v
				break
			}
		}
		if !found {
			err := fmt.Errorf("variable reference not found. variable: %s, reference: %s", uv.Name, uv.ReferenceName)
			errs = append(errs, err)
			continue
			// return fmt.Errorf("variable reference not found. variable: %s, reference: %s", uv.Name, uv.ReferenceName)
		}
		value := int(ref.Val.GetValue()) + uv.Offset
		// fmt.Printf("VARIABLE RESOLVED. variable: %s, value: %04x, offset: %d. final_value: %04x\n", uv.Name, ref.Val.GetValue(), uv.Offset, value)
		val := types.NewValue(int64(value))
		a.variables[uv.Name] = types.NewVariable(uv.Name, val)

		a.symbolTracker.SetIndex(uv.Name, val.GetValue())
		a.symbolTracker.SetType(uv.Name, object.SYMBOL_TYPE_VAR)
		a.symbolTracker.SetReference(uv.Name, "", 0)
	}
	return errs
}

func (a *Assembler) RestoreSymbol(symbol *types.Tag, offset uint16) error {
	// fmt.Printf("* searching symbol. symbol:>%v<\n", symbol)
	im, exists := a.FindVariable(symbol.Text)
	if !exists {
		label, exists := a.FindLabel(symbol.Text)
		if !exists {
			// fmt.Println("symbol not found", symbol.Text, "size", uint8(symbol.Size))
			size := a.symbolTracker.GetSymbolSize(symbol.Text, uint8(symbol.Size))
			a.symbolTracker.SetMissing(symbol.Text, offset, size, uint16(symbol.OptionalOffset))
			// fmt.Printf("[x] SYMBOL: text:>%16s<, offset: %04x, line: %4d, column: %4d, optionalOffset: %d, resolved: %v\n", symbol.Text, offset, 0, 0, 0, false)
			return fmt.Errorf("line %d:%d symbol not found. symbol: '%s'", symbol.Line, symbol.Column, symbol.Text)
		}
		// fmt.Printf("* symbol found it is a label. tag:>%s<, label_offset: %d\n", a.currentTag.Text, label.Offset)
		im = types.NewValue(int64(label.Offset))
		a.symbolTracker.SymbolHit(label.Text, offset, uint8(symbol.Size), uint16(symbol.OptionalOffset))
		// fmt.Printf("[?] SYMBOL: text:>%16s<, offset: %04x, line: %4d, column: %4d, optionalOffset: %d, resolved: %v\n", label.Text, offset, label.Line, label.Column, symbol.OptionalOffset, true)
	} else {
		a.symbolTracker.SymbolHit(symbol.Text, offset, uint8(symbol.Size), uint16(symbol.OptionalOffset))
		// fmt.Printf("[!] SYMBOL: text:>%16s<, offset: %04x, line: %4d, column: %4d, optionalOffset: %d, resolved: %v\n", symbol.Text, offset, 0, 0, symbol.OptionalOffset, true)
	}
	im.Add(uint16(symbol.OptionalOffset))
	// fmt.Printf("* restoring symbol. symbol:>%s<, val: %04x, offset: %d\n", symbol.Text, im.GetValue(), offset)
	a.machineCode[offset] = im.GetLowByte()
	if im.GetSize() == 16 {
		a.machineCode[offset+1] = im.GetHighByte()
	}
	return nil
}

func (a *Assembler) FindVariable(name string) (*types.Value, bool) {
	for _, v := range a.variables {
		if v.Name == name {
			return v.Val, true
		}
	}
	return nil, false
}

func (a *Assembler) FindLabel(name string) (*types.Label, bool) {
	label, ok := a.labels[name]
	return label, ok
}
