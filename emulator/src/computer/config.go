package computer

import (
	"asme8/common/config"
	"time"
)

type Config struct {
	Program        []byte
	MemoryConfig   *config.MemoryConfig
	EnableDebug    bool
	EnableTestMode bool
	IsHeadless     bool
	Delay          time.Duration
}
