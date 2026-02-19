package assembler

import (
	"asme8/assembler/src/types"
	"asme8/emulator/utils"
	"fmt"
	"strings"
)

func (a *Assembler) CreateSymbolsPrint() string {
	builder := strings.Builder{}
	builder.WriteString("Assemble Completed Successfully\n\n")

	if len(a.labels) > 0 {
		builder.WriteString(" Symbols:\n----------------------------\n")
		for k, v := range a.labels {
			builder.WriteString(fmt.Sprintf(" %-18s <%04x>\n", k, v.Offset))
		}
		builder.WriteString("\n")
	}

	if len(a.variables) > 0 {
		builder.WriteString(" Variables:\n----------------------------\n")
		for k, v := range a.variables {
			builder.WriteString(fmt.Sprintf(" %s=0x%02x\n", k, v.Val.GetValue()))
		}
		builder.WriteString("\n")
	}

	builder.WriteString("\n")
	return builder.String()
}

// emulator uses this format so do not change output model
func (a *Assembler) CodeString() string {
	directiveOffsets := a.buildDirectivesOffsetMap()

	builder := strings.Builder{}
	index := 0
	op := strings.Builder{}
	for i := 0; i < len(a.out); i++ {
		code, ok := a.machineCode[uint16(i)]
		if ok {
			op.WriteString(fmt.Sprintf("%02x ", code))
		}
		isLabel, label := findLabelByIndex(i, a.labels)
		if isLabel {
			builder.WriteString(fmt.Sprintf("<%04x> %s:\n", index, label.Text))
		}
		isDirective, directive := isDirective(i, a.directives)
		if isDirective {
			oldIndex := index
			s := ""
			switch directive.Type {
			case DirectiveOrg:
				val := directive.Values[0].GetValue()
				index = int(val)
				builder.WriteString(fmt.Sprintf("<%04x> %s: \n", oldIndex, directive.Raw))
			case DirectiveReserveByte:
				val := directive.Values[0].GetValue()
				index += int(val)
				if val > 8 {
					s = fmt.Sprintf("00 .. .. .. .. .. .. 00 (%d)", val)
				} else {
					for i := 0; i < int(val); i++ {
						s += "00 "
					}
				}
				builder.WriteString(fmt.Sprintf("<%04x> %-34s; %s\n", oldIndex, directive.Raw, s))
			case DirectiveASCII:
				fallthrough
			case DirectiveASCIIZ:
				fallthrough
			case DirectiveByte:
				index += totalOffset(directive.Values)
				arr := ToHexArray_1byte(directive.Values, true)
				s = arr
				builder.WriteString(fmt.Sprintf("<%04x> %-34s; %s\n", oldIndex, directive.Raw, s))
			case DirectiveWord:
				index += totalOffset(directive.Values)
				arr := ToHexArray_2byte(directive.Values, true)
				s = arr
				builder.WriteString(fmt.Sprintf("<%04x> %-34s; %s\n", oldIndex, directive.Raw, s))
			}
		}
		// isBlankLine := isBlank(i, a.Coder.blanksLines)
		isBeginning, _ := isLineBegin(i, a.Coder.lineBeginnings)
		if isBeginning {
			builder.WriteString(fmt.Sprintf("<%04x> ", index))
		}
		_, isDirectiveByte := directiveOffsets[uint16(i)]
		if !isDirectiveByte {
			index++
		} else {
			op.Reset()
		}
		ok, le := isLineEnding(i, a.Coder.lineEndings)
		if ok {
			line := a.Coder.instructions[le]
			builder.WriteString(fmt.Sprintf("    %-30s; %s\n", line, op.String()))
			op.Reset()
		}
	}
	return builder.String()
}

func (a *Assembler) buildDirectivesOffsetMap() map[uint16]struct{} {
	directiveOffsets := make(map[uint16]struct{})
	for i := 0; i < len(a.out); i++ {
		isDirective, directive := isDirective(i, a.directives)
		if !isDirective {
			continue
		}
		switch directive.Type {
		case DirectiveOrg:
			val := directive.Values[0].GetValue()
			i = int(val)
			continue
		case DirectiveReserveByte:
			offset := directive.Values[0].GetValue()
			for j := 0; j < int(offset); j++ {
				directiveOffsets[uint16(i+j)] = struct{}{}
			}
			i += int(offset)
			i--
			continue
		case DirectiveASCII:
			fallthrough
		case DirectiveASCIIZ:
			fallthrough
		case DirectiveByte:
			fallthrough
		case DirectiveWord:
			offset := totalOffset(directive.Values)
			for j := 0; j < int(offset); j++ {
				directiveOffsets[uint16(i+j)] = struct{}{}
			}
			i += int(offset)
			i--
			continue
		}
	}
	return directiveOffsets
}

func totalOffset(values []*types.Value) int {
	offset := 0
	for _, v := range values {
		offset += v.GetSize() / 8
	}
	return offset
}

func isBlank(index int, blanks []uint16) bool {
	for _, le := range blanks {
		if int(le)-1 == index {
			return true
		}
	}
	return false
}

func isLineBegin(index int, lineBeginnings []uint16) (bool, int) {
	for li, le := range lineBeginnings {
		if int(le) == index {
			return true, li
		}
	}
	return false, 0
}

func isLineEnding(index int, lineEndings []uint16) (bool, int) {
	for li, le := range lineEndings {
		if int(le)-1 == index {
			return true, li
		}
	}
	return false, 0
}

func findLabelByIndex(index int, labels map[string]*types.Label) (bool, *types.Label) {
	for _, l := range labels {
		if int(l.Offset) == index {
			return true, l
		}
	}
	return false, nil
}

func isDirective(index int, directives map[uint16]*types.Directive) (bool, *types.Directive) {
	for si, d := range directives {
		if int(si) == index {
			return true, d
		}
	}
	return false, nil
}

func ToHexArray_2byte(arr []*types.Value, noParenthesis ...bool) string {
	arr8 := make([]byte, len(arr)*2)
	offset := 0
	for i := 0; i < len(arr); i++ {
		arr8[offset] = uint8(arr[i].GetLowByte())
		offset++
		arr8[offset] = uint8(arr[i].GetHighByte())
		offset++
	}
	return utils.ToHexArray(arr8, noParenthesis...)
}

func ToHexArray_1byte(arr []*types.Value, noParenthesis ...bool) string {
	arr8 := make([]byte, len(arr))
	for i := range arr8 {
		arr8[i] = uint8(arr[i].GetLowByte())
	}
	return utils.ToHexArray(arr8, noParenthesis...)
}
