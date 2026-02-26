package core

import (
	"asme8/common/object"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func (l *Linker) ResolveCode() error {
	all := make([]string, 0, 32)
	sorted := SortSegmentOffset(l.segmentOffsets)
	for _, seg := range sorted {
		offset := l.segmentOffsets[seg]
		o, ok := l.GetObjectFile(seg)
		if !ok {
			return fmt.Errorf("fatal error. segment defined in offset is not present in object files. (impossible). segment: %s", seg)
		}
		code := o.Tracker.GetCode()
		lines, err := ResolveCode(code, offset)
		if err != nil {
			return err
		}
		all = append(all, lines...)
	}
	corrected, err := l.CorrectDisassembly(all)
	if err != nil {
		return err
	}
	l.code = corrected
	return nil
}

func ResolveCode(code string, offset uint16) ([]string, error) {
	lines := make([]string, 0, 32)
	tokens := strings.Split(code, "\n")
	lastOffset := offset
	// resolves code offset (<0000> part)
	for _, t := range tokens {
		newOffset := uint16(0)
		l := ""
		t = strings.TrimSpace(t)
		// offset representation char count ("<0000>" len 6)
		v, ok, err := parseOffsetText(t)
		if err != nil {
			return nil, err
		}
		if ok {
			newOffset = uint16(v) + offset
			lastOffset = newOffset
			l = t[6:]
		} else {
			newOffset = lastOffset
			l = t
		}
		line := fmt.Sprintf("<%04x>%s", newOffset, l)
		lines = append(lines, line)
	}
	return lines, nil
}

func (l *Linker) CorrectDisassembly(lines []string) ([]string, error) {
	corrected := make([]string, len(lines))
	for i, line := range lines {
		index := strings.Index(line, ";")
		if index == -1 {
			corrected[i] = line
			continue
		}
		// assembly directives skipping
		if len(line) > 7 {
			if line[7] == '.' {
				corrected[i] = line
				continue
			}
		}
		var mainOffset int64
		var err error
		if len(line) > 6 {
			mainOffsetText := line[:6]
			mainOffset, _, err = parseOffsetText(mainOffsetText)
			if err != nil {
				return nil, err
			}
		}
		machineCodePart := line[index+2:]
		machineCodePart = strings.TrimSpace(machineCodePart)
		machineCodesArray := strings.Split(machineCodePart, " ")
		machineCodes, err := stringHexArrayToNumber(machineCodesArray)
		if err != nil {
			return nil, err
		}
		restored := make([]uint8, len(machineCodes))
		for i := 0; i < len(machineCodes); i++ {
			index := int(mainOffset) + i
			b := l.memory[index]
			restored[i] = b
		}
		newLine := strings.Join(hexArrayToStringHex(restored), " ")
		corrected[i] = line[:index] + "; " + newLine
	}
	return corrected, nil
}

func parseOffsetText(text string) (int64, bool, error) {
	if len(text) < 6 {
		return 0, false, nil
	}
	number := text[1:5]
	v, err := strconv.ParseInt(number, 16, 32)
	if err != nil {
		return 0, false, fmt.Errorf("code resolve errors. machine code parse error. line: %s, code: %s, err: %s", text, number, err)
	}
	return v, true, nil
}

func hexArrayToStringHex(arr []uint8) []string {
	n := make([]string, len(arr))
	for i := range arr {
		n[i] = fmt.Sprintf("%02x", arr[i])
	}
	return n
}

func stringHexArrayToNumber(arr []string) ([]uint8, error) {
	n := make([]uint8, len(arr))
	for i := range arr {
		l := arr[i]
		v, err := strconv.ParseInt(l, 16, 16)
		if err != nil {
			// do something
			return nil, fmt.Errorf("wrong string hex input. input:>%s<", l)
		}
		n[i] = uint8(v)
	}
	return n, nil
}

func (l *Linker) GetObjectFile(seg string) (*object.ELF, bool) {
	for _, o := range l.objects {
		if o.Tracker.GetSegment() == seg {
			return o, true
		}
	}
	return nil, false
}

func SortSegmentOffset(m map[string]uint16) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}

	sort.Slice(keys, func(i, j int) bool {
		return m[keys[i]] < m[keys[j]]
	})

	return keys
}
