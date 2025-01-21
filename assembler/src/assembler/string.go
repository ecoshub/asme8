package assembler

import (
	"asme8/assembler/src/types"
	"asme8/emulator/utils"
	"fmt"
)

func (a *Assembler) CreatePrintable() string {
	buffer := "\n"
	if len(a.variables) > 0 {
		buffer += "; variables:\n"
		buffer += "-----------------\n"
		for k, v := range a.variables {
			buffer += fmt.Sprintf("%-18s = 0x%02x\n", k, v.Val.GetValue())
		}
		buffer += "\n"
	}
	if len(a.labels) > 0 {
		buffer += "; labels:\n"
		buffer += "-----------------\n"
		for k, v := range a.labels {
			buffer += fmt.Sprintf("; %-16s = 0x%04x\n", k, v.Offset)
		}
		buffer += "\n"
	}
	ops := ""
	for i, c := range a.out {
		if isNotInDirective(i, a.directives) {
			ops += fmt.Sprintf("%02x", c)
			if i < int(a.max) {
				ops += " "
			}
		}
		ok, label := isLabel(i, a.labels)
		if ok {
			buffer += fmt.Sprintf("%s:\n", label)
		}
		ok, directive := isDirective(i, a.directives)
		if ok {
			if directive.Type == ".org" {
				buffer += fmt.Sprintf("%s\n", directive.Raw)
			} else {
				var arr string
				if directive.Type == ".byte" {
					arr = ToHexArray_1byte(directive.Values, true)
				}
				if directive.Type == ".word" {
					arr = ToHexArray_2byte(directive.Values, true)
				}
				buffer += fmt.Sprintf("%-32s   ; %v\n", directive.Raw, arr)
			}
			ops = ""
		}
		ok, li := isLineEnd(i, a.Coder.linesEndings)
		if ok {
			buffer += fmt.Sprintf("   %-32s; %s\n", a.Coder.instructions[li], ops)
			ops = ""
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

func isLabel(index int, labels map[string]*types.Label) (bool, string) {
	for label, l := range labels {
		if int(l.Offset) == index {
			return true, label
		}
	}
	return false, ""
}

func isDirective(index int, directives map[uint16]*types.Directive) (bool, *types.Directive) {
	for si, d := range directives {
		if int(si) == index {
			return true, d
		}
	}
	return false, nil
}

func isNotInDirective(index int, directives map[uint16]*types.Directive) bool {
	for _, d := range directives {
		if d.Type == ".org" {
			continue
		}
		if d.Position > uint16(index+1) || (d.Position+d.Offset) < uint16(index+1) {
			continue
		} else {
			return false
		}
	}
	return true
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
