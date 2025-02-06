package assembler

import (
	"asme8/assembler/src/types"
	"asme8/emulator/utils"
	"fmt"
)

func (a *Assembler) CreatePrintable() string {
	buffer := "\n"
	if len(a.labels) > 0 {
		buffer += "<0000> ; symbols:\n"
		buffer += "<0000> -----------------\n"
		for k, v := range a.labels {
			buffer += fmt.Sprintf("<0000> ; %-18s <%04x>\n", k, v.Offset)
		}
		buffer += "<0000> \n"
	}
	if len(a.variables) > 0 {
		buffer += "<0000> ; symbols (variables):\n"
		buffer += "<0000> -----------------\n"
		for k, v := range a.variables {
			buffer += fmt.Sprintf("<0000> %s=0x%02x\n", k, v.Val.GetValue())
		}
		buffer += "<0000> \n"
	}
	ops := ""
	for i, c := range a.out {
		if !isSkip(i, a.Coder.skip) {
			ops += fmt.Sprintf("%02x", c)
			ops += " "
		}
		ok, label := isLabel(i, a.labels)
		if ok {
			if !label.DisablePrint {
				buffer += fmt.Sprintf("<%04x>  %s:\n", i, label.Text)
			}
		}
		ok, directive := isDirective(i, a.directives)
		if ok {
			switch directive.Type {
			case ".org":
				buffer += fmt.Sprintf("<%04x> %-40s ; (location: %04x)\n", i, directive.Raw, directive.Values[0].GetValue())
			case ".resb":
				buffer += fmt.Sprintf("<%04x> %-40s ; 00 ... (%d bytes)\n", i, directive.Raw, directive.Values[0].GetValue())
			case ".ascii":
				fallthrough
			case ".asciiz":
				fallthrough
			case ".byte":
				arr := ToHexArray_1byte(directive.Values, true)
				buffer += fmt.Sprintf("<%04x> %-40s ; %v\n", i, directive.Raw, arr)
			case ".word":
				arr := ToHexArray_2byte(directive.Values, true)
				buffer += fmt.Sprintf("<%04x> %-40s ; %v\n", i, directive.Raw, arr)
			}
			ops = ""
		}
		ok, li := isLineEnd(i, a.Coder.linesEndings)
		if ok {
			buffer += fmt.Sprintf("<%04x>     %-36s ; %s\n", i, a.Coder.instructions[li], ops)
			ops = ""
		}
		ok = isBlank(i, a.Coder.blanksLines)
		if ok {
			buffer += fmt.Sprintf("<%04x> \n", i)
		}
	}
	return buffer
}

func (a *Assembler) CodeString() string {
	buffer := ""
	ops := ""
	index := 0
	lastIndex := 0
	for i, c := range a.out {
		if !isSkip(i, a.Coder.skip) {
			ops += fmt.Sprintf("%02x", c)
			ops += " "
			index++
		}
		ok, label := isLabel(i, a.labels)
		if ok {
			if !label.DisablePrint {
				buffer += fmt.Sprintf("<%04x> %s:\n", lastIndex, label.Text)
			}
		}
		ok, directive := isDirective(i, a.directives)
		if ok {
			switch directive.Type {
			case ".org":
				buffer += fmt.Sprintf("<%04x> %-40s; (location: %04x)\n", lastIndex, directive.Raw, directive.Values[0].GetValue())
			case ".resb":
				buffer += fmt.Sprintf("<%04x> %-40s; 00 ... (%d bytes)\n", lastIndex, directive.Raw, directive.Values[0].GetValue())
			case ".ascii":
				fallthrough
			case ".asciiz":
				fallthrough
			case ".byte":
				arr := ToHexArray_1byte(directive.Values, true)
				buffer += fmt.Sprintf("<%04x> %-40s; %v\n", lastIndex, directive.Raw, arr)
			case ".word":
				arr := ToHexArray_2byte(directive.Values, true)
				buffer += fmt.Sprintf("<%04x> %-40s; %v\n", lastIndex, directive.Raw, arr)
			}
			ops = ""
		}
		ok, li := isLineEnd(i, a.Coder.linesEndings)
		if ok {
			buffer += fmt.Sprintf("<%04x>    %-36s; %s\n", lastIndex, a.Coder.instructions[li], ops)
			ops = ""
			lastIndex = index
		}
		ok = isBlank(i, a.Coder.blanksLines)
		if ok {
			buffer += "\n"
		}
	}
	return buffer
}

func isBlank(index int, blanks []uint16) bool {
	for _, le := range blanks {
		if int(le-1) == index {
			return true
		}
	}
	return false
}

func isLineEnd(index int, linesEndings []uint16) (bool, int) {
	for li, le := range linesEndings {
		if int(le-1) == index {
			return true, li
		}
	}
	return false, 0
}

func isLabel(index int, labels map[string]*types.Label) (bool, *types.Label) {
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

func isSkip(index int, skip []uint16) bool {
	for i := 0; i < len(skip); i++ {
		if uint16(index) == skip[i] {
			return true
		}
	}
	return false
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
