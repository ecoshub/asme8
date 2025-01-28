package config

import (
	"asme8/linker/src/memory"
	"asme8/linker/src/segment"
)

type Config struct {
	Memory   []*memory.Memory
	Segments []*segment.Segment
}

func NewConfig() *Config {
	return &Config{
		Memory:   []*memory.Memory{},
		Segments: []*segment.Segment{},
	}
}

func (c *Config) GetMemoryConfig(name string) (*memory.Memory, bool) {
	for _, m := range c.Memory {
		if m.Name == name {
			return m, true
		}
	}
	return nil, false
}

func (c *Config) GetSegmentConfig(name string) (*segment.Segment, bool) {
	for _, c := range c.Segments {
		if c.Name == name {
			return c, true
		}
	}
	return nil, false
}
