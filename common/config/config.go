package config

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

type Config struct {
	Memory  *MemoryConfig
	Segment *SegmentConfig
}

type MemoryConfig struct {
	Configs []*Memory
}

type SegmentConfig struct {
	Configs []*Segment
}

var (
	DefaultConfig = &Config{
		Memory: &MemoryConfig{
			Configs: []*Memory{
				{Name: "ROM", Start: NewNullable(0x0000), Size: 0x2000, Type: MemoryTypeReadOnly},
				{Name: "RAM", Start: NewNullable(0x2000), Size: 0xdfed, Type: MemoryTypeReadWrite},
				{Name: "SERIAL", Start: NewNullable(0xffed), Size: 0x3, Type: MemoryTypeSerial},
				{Name: "VEC", Start: NewNullable(0xfff0), Size: 0x10, Type: MemoryTypeReadOnly},
			},
		},
	}
)

func ParseConfigPath(filePath string) (*Config, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var section string

	c := &Config{
		Segment: &SegmentConfig{
			Configs: make([]*Segment, 0, 4),
		},
		Memory: &MemoryConfig{
			Configs: make([]*Memory, 0, 4),
		},
	}

	memoryNames := make(map[string]struct{})
	segmentNames := make(map[string]struct{})
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
			m, err := parseMemory(line)
			if err != nil {
				return nil, err
			}
			if _, ok := memoryNames[m.Name]; ok {
				return nil, fmt.Errorf("duplicate memory. name: '%s'", m.Name)
			}
			memoryNames[m.Name] = struct{}{}
			c.Memory.Configs = append(c.Memory.Configs, m)
		case "SEGMENTS":
			s, err := parseSegment(line)
			if err != nil {
				return nil, err
			}
			if _, ok := segmentNames[s.Name]; ok {
				return nil, fmt.Errorf("duplicate segment. name: '%s'", s.Name)
			}
			segmentNames[s.Name] = struct{}{}
			c.Segment.Configs = append(c.Segment.Configs, s)
		}
	}

	totalMemory := int(0)
	for _, mc := range c.Memory.Configs {
		if mc.Size == 0 {
			return nil, fmt.Errorf("malformed memory definition. memory: %s", mc.Name)
		}
		totalMemory += int(mc.Size)
	}

	if totalMemory > 0x10000 {
		return nil, fmt.Errorf("total allocated memory exceeding 64k, size: %d", totalMemory)
	}

	RomCount := 0
	for _, conf := range c.Memory.Configs {
		if strings.HasPrefix(conf.Name, "ROM") {
			RomCount += 1
		}
	}

	if RomCount > 1 && len(c.Segment.Configs) == 0 {
		return nil, errors.New("missing 'SEGMENTS'. more than one 'ROM' defined but no 'SEGMENTS' defined to map segments to memory partitions")
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return c, nil
}

func (c *MemoryConfig) GetMemoryConfig(name string) (*Memory, bool) {
	for _, m := range c.Configs {
		if m.Name == name {
			return m, true
		}
	}
	return nil, false
}

func (c *SegmentConfig) GetSegmentConfig(name string) (*Segment, bool) {
	for _, s := range c.Configs {
		if s.Name == name {
			return s, true
		}
	}
	return nil, false
}
