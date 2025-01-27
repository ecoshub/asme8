package assembler

import (
	"asme8/assembler/src/error_listener"
	"asme8/assembler/src/opcodes"
	"asme8/assembler/src/types"
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

type Assembler struct {
	mode               ASM_MODE
	offset             uint16
	currentInstruction string
	currentValue       *types.Value
	currentValueList   []*types.Value
	currentTag         *types.Tag
	currentRegisters   []*types.Register
	labels             map[string]*types.Label
	globals            map[string]*types.Variable
	variables          map[string]*types.Variable
	missingSymbols     map[uint16]*types.Tag
	machineCode        map[uint16]uint8
	directives         map[uint16]*types.Directive
	errorListener      *error_listener.CustomErrorListener
	symbolTracker      *object.Tracker
	hasVirtualOffset   bool
	out                []uint8
	topAddress         uint16

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
		mode:             mode,
		currentValueList: []*types.Value{},
		currentRegisters: []*types.Register{},
		labels:           make(map[string]*types.Label),
		variables:        make(map[string]*types.Variable),
		globals:          make(map[string]*types.Variable),
		machineCode:      make(map[uint16]uint8),
		missingSymbols:   make(map[uint16]*types.Tag),
		directives:       make(map[uint16]*types.Directive),
		symbolTracker:    object.NewTracker(),
		Coder: &Coder{
			linesEndings: make([]uint16, 0, 16),
			blanksLines:  make([]uint16, 0, 16),
			skip:         make([]uint16, 0, 16),
			instructions: make([]string, 0, 16),
		},
	}
}

func (a *Assembler) Assemble() ([]uint8, error) {
	out := a.assemble()
	err := a.errorListener.GetError()
	if err != nil {
		return nil, err
	}
	switch a.mode {
	case ASM_MODE_ELF:
		a.RestoreMissingSymbols()
		a.symbolTracker.AttachBin(out)
		a.symbolTracker.AttachCode(a.CodeString())
		elf := object.ELF{
			Header:  object.NewStdHeader(VERSION),
			Tracker: a.symbolTracker,
		}
		return elf.Encode()
	case ASM_MODE_EXE:
		errs := a.RestoreMissingSymbols()
		if len(errs) > 0 {
			return nil, error_listener.WrapErrors(errs...)
		}
		return out, nil
	}
	return nil, errors.New("unknown assembly mode")
}

func (a *Assembler) AttachErrorListener(errorListener *error_listener.CustomErrorListener) {
	a.errorListener = errorListener
}

func (a *Assembler) AttachGlobals(globals map[string]*types.Variable) {
	a.globals = globals
}

func (a *Assembler) assemble() []uint8 {
	topAddress := uint16(0)
	for offset := range a.machineCode {
		if offset >= topAddress {
			topAddress = offset
		}
	}
	if topAddress == 0 {
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

func (a *Assembler) ParseRegisterImmediate() {
	opcode := opcodes.GetOpCode(a.currentInstruction, opcodes.ADDRESSING_MODE_REG_INM_8)
	a.AppendMachineCode(opcode)
	a.AppendMachineCode(a.currentRegisters[0].GetCode())
	a.AppendMachineCode(a.currentValue.GetLowByte())
}

func (a *Assembler) ParseRegisterRegister() {
	var val uint8 = a.currentRegisters[0].GetCode() | (a.currentRegisters[1].GetCode() << 4)
	opcode := opcodes.GetOpCode(a.currentInstruction, opcodes.ADDRESSING_MODE_REG_REG)
	a.AppendMachineCode(opcode)
	a.AppendMachineCode(val)
}

func (a *Assembler) ParseRegisterPointer() {
	opcode := opcodes.GetOpCode(a.currentInstruction, opcodes.ADDRESSING_MODE_MEM_REG)
	a.AppendMachineCode(opcode)
	a.AppendMachineCode(a.currentRegisters[0].GetCode())
	a.AppendMachineCode(a.currentValue.GetLowByte())
	a.AppendMachineCode(a.currentValue.GetHighByte())
}

func (a *Assembler) ParsePointerRegister() {
	opcode := opcodes.GetOpCode(a.currentInstruction, opcodes.ADDRESSING_MODE_REG_MEM)
	a.AppendMachineCode(opcode)
	a.AppendMachineCode(a.currentRegisters[0].GetCode())
	a.AppendMachineCode(a.currentValue.GetLowByte())
	a.AppendMachineCode(a.currentValue.GetHighByte())
}

func (a *Assembler) ParseRegisterPointerOffset() {
	a.GetVariableOrTagMissing(2, 16)
	var val uint8 = a.currentRegisters[1].GetCode() | (a.currentRegisters[0].GetCode() << 4)
	opcode := opcodes.GetOpCode(a.currentInstruction, opcodes.ADDRESSING_MODE_MEM_REG_OFFSET)
	a.AppendMachineCode(opcode)
	a.AppendMachineCode(val)
	a.AppendMachineCode(a.currentValue.GetLowByte())
	a.AppendMachineCode(a.currentValue.GetHighByte())
}

func (a *Assembler) ParsePointerRegisterOffset() {
	a.GetVariableOrTagMissing(2, 16)
	var val uint8 = a.currentRegisters[0].GetCode() | (a.currentRegisters[1].GetCode() << 4)
	opcode := opcodes.GetOpCode(a.currentInstruction, opcodes.ADDRESSING_MODE_REG_MEM_OFFSET)
	a.AppendMachineCode(opcode)
	a.AppendMachineCode(val)
	a.AppendMachineCode(a.currentValue.GetLowByte())
	a.AppendMachineCode(a.currentValue.GetHighByte())
}

func (a *Assembler) ParseRegisterImmediateVariable() {
	a.GetVariableOrTagMissing(2, 8)
	a.ParseRegisterImmediate()
}

func (a *Assembler) ParseImpliedRegister() {
	opcode := opcodes.GetOpCode(a.currentInstruction, opcodes.ADDRESSING_MODE_IMPL_REG)
	a.AppendMachineCode(opcode)
	a.AppendMachineCode(a.currentRegisters[0].GetCode())
}

func (a *Assembler) ParseImplied() {
	opcode := opcodes.GetOpCode(a.currentInstruction, opcodes.ADDRESSING_MODE_IMPL)
	a.AppendMachineCode(opcode)
}

func (a *Assembler) ParseImpliedImmediate() {
	opcode := opcodes.GetOpCode(a.currentInstruction, opcodes.ADDRESSING_MODE_IMPL_INM_16)
	a.AppendMachineCode(opcode)
	a.AppendMachineCode(a.currentValue.GetLowByte())
	a.AppendMachineCode(a.currentValue.GetHighByte())
}

func (a *Assembler) ParseImpliedTag() {
	v, exists := a.FindVariable(a.currentTag.Text)
	if exists {
		a.currentValue = v
		a.ParseImpliedImmediate()
		return
	}

	label, exist := a.FindLabel(a.currentTag.Text)
	if exist {
		a.symbolTracker.SymbolHit(label.Text, a.offset+1, 16, 0)
		// fmt.Printf("[*] SYMBOL: text:>%16s<, offset: %04x, line: %4d, column: %4d\n", label.Text, a.offset+1, label.Line, label.Column)
		a.currentValue = types.NewValue(int64(label.Offset))
		a.ParseImpliedImmediate()
		return
	}

	// fmt.Printf("? symbol not found. tag:>%s<, offset: %d\n", a.currentTag.Text, a.offset)
	a.currentTag.Size = 16
	a.missingSymbols[a.offset+1] = a.currentTag
	a.currentValue = types.NewValue(0)

	a.ParseImpliedImmediate()
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

	// fmt.Printf("? symbol not found. tag:>%s<, offset: %d\n", a.currentTag.Text, a.offset)
	a.currentTag.Size = 16
	a.missingSymbols[a.offset+2] = a.currentTag
	a.currentValue = types.NewValue(0)
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
}

func (a *Assembler) ParseTag(text string, line, column int) {
	a.currentTag = &types.Tag{Text: text, Line: line, Column: column}
}

func (a *Assembler) ParseRegister(text string) {
	a.currentRegisters = append(a.currentRegisters, types.ParseRegister(text))
}

func (a *Assembler) ParseVariable() {
	v := types.NewVariable(a.currentTag.Text, a.currentValue.Copy())
	a.variables[a.currentTag.Text] = v

	a.symbolTracker.SetIndex(a.currentTag.Text, v.Val.GetValue())
	a.symbolTracker.SetType(a.currentTag.Text, object.SYMBOL_TYPE_VAR)
}

func (a *Assembler) ParseInstruction(text string) {
	a.currentInstruction = strings.ToLower(text)
}

func (a *Assembler) ParseValue(text string, line, column int) {
	var ok bool
	a.currentValue, ok = types.ParseValue(text)
	if !ok {
		a.missingSymbols[a.offset] = &types.Tag{Text: text, Line: line, Column: column, Size: 0}
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
	// fmt.Printf("?? symbol not found. symbol:>%s<, offset: %d\n", a.currentTag, a.offset)
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
	for _, v := range a.globals {
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
