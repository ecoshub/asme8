package parser

import "fmt"

func (a *Assembler) String() string {
	a.process()
	buffer := "\n"
	if len(a.variables) > 0 {
		for k, v := range a.variables {
			buffer += fmt.Sprintf("%-18s = 0x%02x\n", k, v)
		}
		buffer += "\n"
	}
	if len(a.labels) > 0 {
		buffer += "; # offsets:\n"
		for k, v := range a.labels {
			buffer += fmt.Sprintf("; %-16s = 0x%02x\n", k, v)
		}
		buffer += "\n"
	}
	ops := ""
	for i, c := range a.out {
		ops += fmt.Sprintf("%02x", c)
		if i < len(a.out) {
			ops += " "
		}
		ok, li := isLineEnd(i, a.linesEndings)
		if ok {
			buffer += fmt.Sprintf("   %-32s; %s\n", a.lines[li], ops)
			ops = ""
		}
		ok, label := isLabel(i, a.labels)
		if ok {
			buffer += fmt.Sprintf("%s:\n", label)
		}
	}
	return buffer
}
