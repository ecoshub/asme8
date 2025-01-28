package comp

import "fmt"

func mInstProgramCounterOutAddr(c *Comp, _ uint64) {
	c.addrBus.Write_16(c.programCounter)
}

func mInstProgramCounterInAddr(c *Comp, _ uint64) {
	c.programCounter = c.addrBus.Read_16()
	checkBreakPoint(c)
}

func mInstProgramCounterLowIn(c *Comp, _ uint64) {
	upper := c.programCounter & 0xff00
	c.programCounter = upper | c.inputBus.Read_16()
}

func mInstProgramCounterHighIn(c *Comp, _ uint64) {
	lower := c.programCounter & 0x00ff
	c.programCounter = lower | (c.inputBus.Read_16() << 8)
	checkBreakPoint(c)
}

func mInstProgramCounterLowOut(c *Comp, _ uint64) {
	c.outputBus.Write_8(uint8(c.programCounter))
	triggerBridge(c)
}

func mInstProgramCounterHighOut(c *Comp, _ uint64) {
	c.outputBus.Write_8(uint8(c.programCounter >> 8))
	triggerBridge(c)
}

func mInstProgramCounterInc(c *Comp, _ uint64) {
	c.programCounter++
	checkBreakPoint(c)
}

func checkBreakPoint(c *Comp) {
	for _, bp := range c.breakPoints {
		if bp == c.programCounter {
			c.pause = true
			c.Log("Stop")
			c.LogWithStyle(fmt.Sprintf("● Breakpoint triggered. addr: 0x%04x", bp), DefaultBreakStyle)
			return
		}
	}
}
