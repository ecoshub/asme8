package assembler

import (
	"asme8/assembler/src/error_listener"
	"asme8/assembler/src/symbols"
	"asme8/assembler/src/types"
	"fmt"
	"strconv"
	"strings"
)

type Assembler struct {
	Coder              *Coder
	offset             uint16
	currentInstruction string
	currentValue       *types.Value
	currentValueList   []*types.Value
	currentTag         *types.Tag
	currentRegisters   []*types.Register
	labels             map[string]*types.Label
	variables          map[string]*types.Variable
	missingSymbols     map[uint16]*types.Tag
	machineCode        map[uint16]uint8
	directives         map[uint16]*types.Directive
	errorListener      *error_listener.CustomErrorListener
	hasVirtualOffset   bool
	out                []uint8
	max                uint16
}

type Coder struct {
	linesEndings []uint16
	blanksLines  []uint16
	instructions []string
	lastLine     string
}

func New(errorListener *error_listener.CustomErrorListener) *Assembler {
	return &Assembler{
		currentValueList: []*types.Value{},
		currentRegisters: []*types.Register{},
		labels:           make(map[string]*types.Label),
		variables:        make(map[string]*types.Variable),
		machineCode:      make(map[uint16]uint8),
		missingSymbols:   make(map[uint16]*types.Tag),
		directives:       make(map[uint16]*types.Directive),
		errorListener:    errorListener,
		Coder: &Coder{
			linesEndings: []uint16{},
			blanksLines:  []uint16{},
			instructions: []string{},
		},
	}
}

func (a *Assembler) Assemble() ([]uint8, error) {
	errs := a.RestoreMissingSymbols()
	if len(errs) > 0 {
		return nil, error_listener.WrapErrors(errs...)
	}
	max := uint16(0)
	for offset := range a.machineCode {
		if offset >= max {
			max = offset
		}
	}
	out := make([]uint8, max+1)
	for i := uint16(0); i < max+1; i++ {
		b := a.machineCode[i]
		out[i] = b
	}
	a.max = max
	a.out = out
	return out, nil
}

func (a *Assembler) AppendMachineCode(code uint8) {
	a.machineCode[a.offset] = code
	a.offset++
}

func (a *Assembler) ParseRegisterImmediate() {
	opcode := symbols.GetOpCode(a.currentInstruction, symbols.ADDRESSING_MODE_REG_INM_8)
	a.AppendMachineCode(opcode)
	a.AppendMachineCode(a.currentRegisters[0].GetCode())
	a.AppendMachineCode(a.currentValue.GetLowByte())
}

func (a *Assembler) ParseRegisterRegister() {
	var val uint8 = a.currentRegisters[0].GetCode() | (a.currentRegisters[1].GetCode() << 4)
	opcode := symbols.GetOpCode(a.currentInstruction, symbols.ADDRESSING_MODE_REG_REG)
	a.AppendMachineCode(opcode)
	a.AppendMachineCode(val)
}

func (a *Assembler) ParseRegisterPointer() {
	opcode := symbols.GetOpCode(a.currentInstruction, symbols.ADDRESSING_MODE_MEM_REG)
	a.AppendMachineCode(opcode)
	a.AppendMachineCode(a.currentRegisters[0].GetCode())
	a.AppendMachineCode(a.currentValue.GetLowByte())
	a.AppendMachineCode(a.currentValue.GetHighByte())
}

func (a *Assembler) ParsePointerRegister() {
	opcode := symbols.GetOpCode(a.currentInstruction, symbols.ADDRESSING_MODE_REG_MEM)
	a.AppendMachineCode(opcode)
	a.AppendMachineCode(a.currentRegisters[0].GetCode())
	a.AppendMachineCode(a.currentValue.GetLowByte())
	a.AppendMachineCode(a.currentValue.GetHighByte())
}

func (a *Assembler) ParseRegisterPointerOffset() {
	a.GetVariableOrTagMissing(2)
	var val uint8 = a.currentRegisters[1].GetCode() | (a.currentRegisters[0].GetCode() << 4)
	opcode := symbols.GetOpCode(a.currentInstruction, symbols.ADDRESSING_MODE_MEM_REG_OFFSET)
	a.AppendMachineCode(opcode)
	a.AppendMachineCode(val)
	a.AppendMachineCode(a.currentValue.GetLowByte())
	a.AppendMachineCode(a.currentValue.GetHighByte())
}

func (a *Assembler) ParsePointerRegisterOffset() {
	a.GetVariableOrTagMissing(2)
	var val uint8 = a.currentRegisters[0].GetCode() | (a.currentRegisters[1].GetCode() << 4)
	opcode := symbols.GetOpCode(a.currentInstruction, symbols.ADDRESSING_MODE_REG_MEM_OFFSET)
	a.AppendMachineCode(opcode)
	a.AppendMachineCode(val)
	a.AppendMachineCode(a.currentValue.GetLowByte())
	a.AppendMachineCode(a.currentValue.GetHighByte())
}

func (a *Assembler) ParseRegisterImmediateVariable() {
	a.GetVariableOrTagMissing(2)
	a.ParseRegisterImmediate()
}

func (a *Assembler) ParseImpliedRegister() {
	opcode := symbols.GetOpCode(a.currentInstruction, symbols.ADDRESSING_MODE_IMPL_REG)
	a.AppendMachineCode(opcode)
	a.AppendMachineCode(a.currentRegisters[0].GetCode())
}

func (a *Assembler) ParseImplied() {
	opcode := symbols.GetOpCode(a.currentInstruction, symbols.ADDRESSING_MODE_IMPL)
	a.AppendMachineCode(opcode)
}

func (a *Assembler) ParseImpliedImmediate() {
	opcode := symbols.GetOpCode(a.currentInstruction, symbols.ADDRESSING_MODE_IMPL_INM_16)
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
		a.currentValue = types.NewValue(int64(label.Offset))
		a.ParseImpliedImmediate()
		return
	}

	// fmt.Printf("? symbol not found. tag:>%s<, offset: %d\n", a.currentTag, a.offset)
	a.missingSymbols[a.offset+1] = a.currentTag
	a.currentValue = types.NewValue(0)

	a.ParseImpliedImmediate()
}

func (a *Assembler) ParseLabel(text string, line, column int, disablePrint bool) {
	label := types.NewLabel(text, line, column, a.offset)
	label.DisablePrint = disablePrint
	a.labels[text] = label
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
		a.currentValue = v
		return
	}

	// fmt.Printf("? symbol not found. tag:>%s<, offset: %d\n", a.currentTag, a.offset)
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
		a.currentValue = types.NewValue(int64(v.GetValue()) + int64(offset))
		// fmt.Println("val", a.currentValue.GetValue())
		return
	}
	a.missingSymbols[a.offset+2] = &types.Tag{Text: tag, Line: line, Column: column, OptionalOffset: offset}
	a.currentValue = types.NewValue(0)
}

func (a *Assembler) ParseDirective(text string) {
	if strings.HasPrefix(text, ".resb") {
		fmt.Println("directive added to", a.offset, "text:", text)
		val := a.currentValueList[0]
		a.directives[a.offset] = &types.Directive{Raw: text, Type: ".resb", Position: a.offset, Offset: val.GetValue(), Values: a.currentValueList}
		for i := 0; i < int(val.GetValue()); i++ {
			a.AppendMachineCode(0)
		}
		return
	}
	if strings.HasPrefix(text, ".byte") {
		fmt.Println("directive added to", a.offset, "text:", text)
		a.directives[a.offset] = &types.Directive{Raw: text, Type: ".byte", Position: a.offset, Offset: uint16(len(a.currentValueList)), Values: a.currentValueList}
		for _, v := range a.currentValueList {
			a.AppendMachineCode(v.GetLowByte())
		}
		return
	}
	if strings.HasPrefix(text, ".word") {
		fmt.Println("directive added to", a.offset, "text:", text)
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
}

func (a *Assembler) ParseInstruction(text string) {
	a.currentInstruction = strings.ToLower(text)
}

func (a *Assembler) ParseValue(text string) {
	a.currentValue = types.ParseValue(text)
}

func (a *Assembler) ParseValueList(text string) {
	tokens := strings.Split(text, ", ")
	valueList := make([]*types.Value, 0, len(tokens))
	for _, tok := range tokens {
		tok = strings.TrimSpace(tok)
		variable, ok := a.variables[tok]
		if ok {
			valueList = append(valueList, variable.Val)
			continue
		}
		val := types.ParseValue(tok)
		valueList = append(valueList, val)
	}
	a.currentValueList = valueList
}

func (a *Assembler) ResetCurrentValues() {
	a.currentRegisters = []*types.Register{}
	a.currentValue = nil
	a.currentTag = &types.Tag{}
	a.currentInstruction = ""
	a.hasVirtualOffset = false
}

func (a *Assembler) GetVariableOrTagMissing(offsetPlus uint16) {
	if a.currentTag.Text == "" {
		return
	}
	variable, exists := a.FindVariable(a.currentTag.Text)
	if exists {
		a.currentValue = types.NewValue(int64(variable.GetValue()))
		return
	}
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
			// fmt.Println("symbol not found", symbol.Text)
			return fmt.Errorf("line %d:%d symbol not found. symbol: '%s'", symbol.Line, symbol.Column, symbol.Text)
		}
		// fmt.Printf("* symbol found it is a label. tag:>%s<, label_offset: %d\n", a.currentTag.Text, label.Offset)
		im = types.NewValue(int64(label.Offset))
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
