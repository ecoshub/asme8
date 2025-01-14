package parser

import (
	"asme8/emulator/utils"
	"fmt"
)

func (a *Assembler) String() string {
	a.process()
	buffer := "\n"
	if len(a.variables) > 0 {
		buffer += "; variables:\n"
		buffer += "-----------------\n"
		for k, v := range a.variables {
			buffer += fmt.Sprintf("%-18s = 0x%02x\n", k, v)
		}
		buffer += "\n"
	}
	if len(a.labels) > 0 {
		buffer += "; labels:\n"
		buffer += "-----------------\n"
		for k, v := range a.labels {
			buffer += fmt.Sprintf("; %-16s = 0x%02x\n", k, v)
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
			if directive.code == ".org" {
				buffer += fmt.Sprintf("%s\n", directive.raw)
			} else {
				var arr string
				if directive.code == ".byte" {
					arr = utils.ToHexArray_1byte(directive.inm, true)
				}
				if directive.code == ".word" {
					arr = utils.ToHexArray_2byte(directive.inm, true)
				}
				buffer += fmt.Sprintf("%-32s   ; %v\n", directive.raw, arr)
			}
			ops = ""
		}
		ok, li := isLineEnd(i, a.linesEndings)
		if ok {
			buffer += fmt.Sprintf("   %-32s; %s\n", a.instructions[li], ops)
			ops = ""
		}
		ok = isBlank(i, a.blanks)
		if ok {
			buffer += "\n"
		}
	}
	return buffer
}

func isBlank(index int, blanks []uint64) bool {
	for _, le := range blanks {
		if int(le-1) == index {
			return true
		}
	}
	return false
}

func isLineEnd(index int, linesEndings []uint64) (bool, int) {
	for li, le := range linesEndings {
		if int(le-1) == index {
			return true, li
		}
	}
	return false, 0
}

func isLabel(index int, labels map[string]uint64) (bool, string) {
	for label, li := range labels {
		if int(li) == index {
			return true, label
		}
	}
	return false, ""
}

func isDirective(index int, directives map[uint16]*Directive) (bool, *Directive) {
	for si, d := range directives {
		if int(si) == index {
			return true, d
		}
	}
	return false, nil
}

func isNotInDirective(index int, directives map[uint16]*Directive) bool {
	for _, d := range directives {
		if d.code == ".org" {
			continue
		}
		if d.position > uint64(index+1) || (d.position+d.offset) < uint64(index+1) {
			continue
		} else {
			return false
		}
	}
	return true
}
