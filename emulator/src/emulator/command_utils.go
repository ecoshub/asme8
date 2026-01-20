package emulator

import (
	"asme8/emulator/src/emulator/breakpoint"
	"fmt"
	"strconv"
	"strings"
)

func splitNumberCommand(command string, prefixes ...string) (int64, bool, error) {
	param, exists, err := splitStringCommand(command, prefixes...)
	if err != nil {
		return 0, exists, err
	}

	var n int64
	if strings.HasPrefix(param, "0x") {
		param = strings.TrimPrefix(param, "0x")
		n, err = strconv.ParseInt(param, 16, 64)
		if err != nil {
			return 0, true, fmt.Errorf("parse hex error. command: %s, err: %s", command, err)
		}
		return n, true, nil
	}
	n, err = strconv.ParseInt(param, 10, 64)
	if err != nil {
		return 0, true, fmt.Errorf("parse number error. command: %s, err: %s", command, err)
	}
	return n, true, nil
}

func splitStringCommand(command string, prefixes ...string) (string, bool, error) {
	for _, prefix := range prefixes {
		if strings.HasPrefix(command, prefix+" ") {
			tokens := strings.Split(command, " ")
			param := tokens[1]
			return param, true, nil
		}
	}
	return "", false, fmt.Errorf("error unknown params. command: '%s'", command)
}

func listBreakpoints(bpm *breakpoint.Manager) []string {
	bps := bpm.GetBreakpoints()
	lines := make([]string, 0, len(bps))
	if len(bps) == 0 {
		lines = append(lines, "no breakpoint found")
		return lines
	}
	lines = append(lines, "breakpoints:")
	for _, bp := range bps {
		lines = append(lines, fmt.Sprintf(" ● 0x%04x", bp))
	}
	return lines
}
