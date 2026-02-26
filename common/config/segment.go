package config

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Segment struct {
	Name  string
	Load  string
	Start *NullableValue
}

func ParseSegmentConfig(filePath string) (*SegmentConfig, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	c := &SegmentConfig{
		Configs: make([]*Segment, 0, 2),
	}

	scanner := bufio.NewScanner(file)
	var section string

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		if line == "}" {
			section = ""
			return c, nil
		}

		if strings.HasPrefix(line, "SEGMENTS") {
			section = "SEGMENTS"
			continue
		}

		if section != "SEGMENTS" {
			return nil, fmt.Errorf("unexpected section. section: %s", section)
		}

		seg, err := parseSegment(line)
		if err != nil {
			return nil, err
		}
		c.Configs = append(c.Configs, seg)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return c, nil
}

func parseSegment(line string) (*Segment, error) {
	parts := strings.Split(line, ":")
	if len(parts) != 2 {
		return nil, fmt.Errorf("invalid segment line: %s", line)
	}

	name := strings.TrimSpace(parts[0])
	paramLine := strings.TrimSuffix(parts[1], ";")
	params := strings.Split(paramLine, ",")
	if len(params) < 1 {
		return nil, fmt.Errorf("invalid segment parameters: %s", parts[1])
	}

	segment := &Segment{Name: name}
	for _, param := range params {
		kv := strings.Split(param, "=")
		if len(kv) != 2 {
			return nil, fmt.Errorf("invalid segment parameter: %s", param)
		}
		key := strings.TrimSpace(kv[0])
		value := strings.TrimSpace(kv[1])

		switch key {
		case "load":
			segment.Load = value
		case "start":
			start, err := strconv.ParseUint(value, 0, 16)
			if err != nil {
				return nil, err
			}
			segment.Start = &NullableValue{Value: uint16(start)}
		default:
			return nil, fmt.Errorf("unknown segment parameter: %s", key)
		}
	}

	return segment, nil
}
