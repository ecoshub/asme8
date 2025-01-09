package comp

import "emu/src/utils"

func mInstSetRead(c *Comp, _ uint64, _ uint64) {
	c.rw = utils.IO_READ
	c.tickDevices()
}

func mInstSetWrite(c *Comp, _ uint64, _ uint64) {
	c.rw = utils.IO_WRITE
	c.tickDevices()
}
