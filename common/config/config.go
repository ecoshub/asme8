package config

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Config struct {
	MemoryConfig  *MemoryConfig
	SegmentConfig *SegmentConfig
}

type MemoryConfig struct {
	Configs []*Memory
}

type SegmentConfig struct {
	Configs []*Segment
}

func ParseConfig(filePath string) (*Config, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var section string

	c := &Config{
		SegmentConfig: &SegmentConfig{
			Configs: make([]*Segment, 0, 4),
		},
		MemoryConfig: &MemoryConfig{
			Configs: make([]*Memory, 0, 4),
		},
	}

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
			c.MemoryConfig.Configs = append(c.MemoryConfig.Configs, m)
		case "SEGMENTS":
			s, err := parseSegment(line)
			if err != nil {
				return nil, err
			}
			c.SegmentConfig.Configs = append(c.SegmentConfig.Configs, s)
		}
	}

	totalMemory := int(0)
	lastStart := int(0)
	for _, mc := range c.MemoryConfig.Configs {
		if mc.Size == nil && mc.Start == nil {
			return nil, fmt.Errorf("malformed memory definition. memory: %s", mc.Name)
		}
		if mc.Size != nil {
			totalMemory += int(mc.Size.Value)
		} else {
			totalMemory += (int(mc.Start.Value) - lastStart)
			lastStart = int(mc.Start.Value)
		}
	}
	if totalMemory > 0x10000 {
		return nil, fmt.Errorf("total allocated memory exceeding 64k")
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
	for _, c := range c.Configs {
		if c.Name == name {
			return c, true
		}
	}
	return nil, false
}
