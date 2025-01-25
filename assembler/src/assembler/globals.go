package assembler

import (
	"asme8/assembler/src/types"
	"fmt"
	"strings"
)

func FormatGlobals(globals []string) (map[string]*types.Variable, error) {
	variables := make(map[string]*types.Variable, len(globals))
	for _, kv := range globals {
		tokens := strings.Split(kv, "=")
		if len(tokens) != 2 {
			return nil, fmt.Errorf("malformed global '%s'", kv)
		}
		key := strings.TrimSpace(tokens[0])
		valRaw := strings.TrimSpace(tokens[1])
		val, ok := types.ParseValue(valRaw)
		if !ok {
			return nil, fmt.Errorf("invalid value for global '%s'", kv)
		}
		variables[key] = &types.Variable{Name: key, Val: val}
	}
	return variables, nil
}
