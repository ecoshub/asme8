package comp

import "emu/src/utils"

func mInstSetRead(c *Comp, _ uint64) {
	c.rw = utils.IO_READ
	triggerBridge(c)
	c.tickDevices()
	triggerBridge(c)
}

func mInstSetWrite(c *Comp, _ uint64) {
	triggerBridge(c)
	c.rw = utils.IO_WRITE
	c.tickDevices()
}
