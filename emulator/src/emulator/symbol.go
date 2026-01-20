package emulator

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

func ReadSymbolFileAndCreateCodePanel(filePath, symbolFilePath string) (map[uint16]string, bool, error) {
	possiblePathInUse := false
	if symbolFilePath == "" {
		base := filepath.Base(filePath)
		tokens := strings.Split(base, ".")
		if len(tokens) != 2 {
			return map[uint16]string{}, false, errors.New("sym file not found")
		}
		rest := strings.TrimSuffix(filePath, base)
		base = tokens[0]
		possiblePath := filepath.Join(rest, base+".sym")
		symbolFilePath = possiblePath
		possiblePathInUse = true
	}

	file, err := os.ReadFile(symbolFilePath)
	if err != nil {
		if !possiblePathInUse {
			return map[uint16]string{}, false, err
		}
	}

	return ResolveSymbolFromFile(string(file))
}

func ResolveSymbolFromFile(code string) (map[uint16]string, bool, error) {
	if code == "" {
		return nil, false, nil
	}
	codeLines := make(map[uint16]string, 64)
	lines := strings.Split(code, "\n")
	for _, l := range lines {
		if len(l) < 6 {
			continue
		}
		hex := l[1:5]
		str := l[6:]
		v, err := strconv.ParseInt(hex, 16, 32)
		if err != nil {
			return nil, false, fmt.Errorf("code resolve error. machine code parse error. line: '%s', code: '%s', err: %s", l, hex, err)
		}
		s, exists := codeLines[uint16(v)]
		if exists {
			codeLines[uint16(v)] = s + " " + strings.TrimSpace(str)
		} else {
			codeLines[uint16(v)] = l
		}
	}

	for offset, l := range codeLines {
		index := strings.Index(l, ";")
		if index == -1 {
			continue
		}
		c := l[:index]
		c = strings.TrimSpace(c)
		m := l[index+1:]
		m = strings.TrimSpace(m)
		codeLines[offset] = fmt.Sprintf("%-52s%s", c, m)
	}

	return codeLines, true, nil
}

func SortCodeLines(m map[uint16]string) []uint16 {
	keys := make([]uint16, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}

	sort.Slice(keys, func(i, j int) bool {
		return m[keys[i]] < m[keys[j]]
	})

	return keys
}
