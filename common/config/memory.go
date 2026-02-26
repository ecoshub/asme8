package config

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	MemoryTypeReadOnly  string = "ro"
	MemoryTypeReadWrite string = "rw"
)

type Memory struct {
	Name  string
	Start *NullableValue
	Size  uint16
	Type  string
}

func ResolveMemoryLayout(ml []*Memory) error {
	offset := uint16(0)
	for _, m := range ml {
		if m.Start == nil {
			m.Start = &NullableValue{Value: offset}
			offset += m.Size
		}
		// fmt.Printf("%s: start=0x%04x, size=0x%04x, type=%s\n", m.Name, m.Start.Value, m.Size.Value, m.Type)
	}
	return nil
}

func parseMemory(line string) (*Memory, error) {
	parts := strings.Split(line, ":")
	if len(parts) != 2 {
		return nil, fmt.Errorf("invalid memory line: %s", line)
	}

	name := strings.TrimSpace(parts[0])
	paramLine := strings.TrimSuffix(parts[1], ";")
	params := strings.Split(paramLine, ",")
	if len(params) < 1 {
		return nil, fmt.Errorf("invalid memory parameters: %s", parts[1])
	}

	memory := &Memory{Name: name}
	for _, param := range params {
		kv := strings.Split(param, "=")
		if len(kv) != 2 {
			return nil, fmt.Errorf("invalid memory parameter: %s", param)
		}
		key := strings.TrimSpace(kv[0])
		value := strings.TrimSpace(kv[1])

		switch key {
		case "start":
			start, err := strconv.ParseUint(value, 0, 16)
			if err != nil {
				return nil, err
			}
			memory.Start = &NullableValue{Value: uint16(start)}
		case "size":
			size, err := strconv.ParseUint(value, 0, 16)
			if err != nil {
				return nil, err
			}
			memory.Size = uint16(size)
		case "type":
			memory.Type = value
		default:
			return nil, fmt.Errorf("unknown memory parameter: %s", key)
		}
	}

	return memory, nil
}
