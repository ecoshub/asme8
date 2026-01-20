package computer

import (
	"asme8/emulator/utils"
)

func mInstSetRead(c *Computer, mi uint64) {
	c.rw = utils.IO_READ
	triggerBridge(c)
	c.tickDevices()
	triggerBridge(c)
}

func mInstSetWrite(c *Computer, _ uint64) {
	triggerBridge(c)
	c.rw = utils.IO_WRITE
	c.tickDevices()
}
