package computer

import (
	"asme8/common/config"
	"time"
)

type Config struct {
	Program      []byte
	MemoryConfig *config.MemoryConfig
	Debug        bool
	TestMode     bool
	Headless     bool
	Delay        time.Duration
}
