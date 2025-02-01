package config

type MemoryConfig struct {
	Configs []*Memory
}

type SegmentConfig struct {
	Configs []*Segment
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
