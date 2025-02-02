package comp

import (
	"asme8/common/config"
	"time"
)

type Config struct {
	Debug          bool
	MemoryConfig   *config.MemoryConfig
	Headless       bool
	Program        []byte
	Delay          time.Duration
	Test           bool
	SymbolFilePath string
}
