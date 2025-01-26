package config

import (
	"asme8/linker/src/memory"
	"asme8/linker/src/segment"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func (c *Config) ParseConfig(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var section string

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		if line == "}" {
			section = ""
			continue
		}

		if strings.HasPrefix(line, "MEMORY") {
			section = "MEMORY"
			continue
		} else if strings.HasPrefix(line, "SEGMENTS") {
			section = "SEGMENTS"
			continue
		}

		switch section {
		case "MEMORY":
			err := c.parseMemory(line)
			if err != nil {
				return err
			}
		case "SEGMENTS":
			err := c.parseSegment(line)
			if err != nil {
				return err
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

func (c *Config) parseMemory(line string) error {
	parts := strings.Split(line, ":")
	if len(parts) != 2 {
		return fmt.Errorf("invalid memory line: %s", line)
	}

	name := strings.TrimSpace(parts[0])
	paramLine := strings.TrimSuffix(parts[1], ";")
	params := strings.Split(paramLine, ",")
	if len(params) != 3 {
		return fmt.Errorf("invalid memory parameters: %s", parts[1])
	}

	memory := &memory.Memory{Name: name}
	for _, param := range params {
		kv := strings.Split(param, "=")
		if len(kv) != 2 {
			return fmt.Errorf("invalid memory parameter: %s", param)
		}
		key := strings.TrimSpace(kv[0])
		value := strings.TrimSpace(kv[1])

		switch key {
		case "start":
			start, err := strconv.ParseUint(value, 0, 16)
			if err != nil {
				return err
			}
			memory.Start = uint16(start)
		case "size":
			size, err := strconv.ParseUint(value, 0, 16)
			if err != nil {
				return err
			}
			memory.Size = uint16(size)
		case "type":
			memory.Type = value
		default:
			return fmt.Errorf("unknown memory parameter: %s", key)
		}
	}

	c.Memory = append(c.Memory, memory)
	return nil
}

func (c *Config) parseSegment(line string) error {
	parts := strings.Split(line, ":")
	if len(parts) != 2 {
		return fmt.Errorf("invalid segment line: %s", line)
	}

	name := strings.TrimSpace(parts[0])
	paramLine := strings.TrimSuffix(parts[1], ";")
	params := strings.Split(paramLine, ",")
	if len(params) < 2 {
		return fmt.Errorf("invalid segment parameters: %s", parts[1])
	}

	segment := &segment.Segment{Name: name}
	for _, param := range params {
		kv := strings.Split(param, "=")
		if len(kv) != 2 {
			return fmt.Errorf("invalid segment parameter: %s", param)
		}
		key := strings.TrimSpace(kv[0])
		value := strings.TrimSpace(kv[1])

		switch key {
		case "load":
			segment.Load = value
		case "start":
			start, err := strconv.ParseUint(value, 0, 16)
			if err != nil {
				return err
			}
			segment.Start = uint16(start)
		case "type":
			segment.Type = value
		default:
			return fmt.Errorf("unknown segment parameter: %s", key)
		}
	}

	c.Segments = append(c.Segments, segment)
	return nil
}
