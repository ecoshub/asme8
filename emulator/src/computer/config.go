package computer

import (
	"asme8/common/config"
	"time"
)

type Config struct {
	MemorySegment *config.Config
	Debug         bool
	TestMode      bool
	Headless      bool
	Delay         time.Duration
}
