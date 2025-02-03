package comp

import "fmt"

func mInstProgramCounterOutAddr(c *Comp, _ uint64) {
	c.addrBus.Write_16(c.programCounter)
	c.LogCodePanel(false)
	c.checkBreakPoint()
}

func mInstProgramCounterInAddr(c *Comp, _ uint64) {
	c.programCounter = c.addrBus.Read_16()
}

func mInstProgramCounterLowIn(c *Comp, _ uint64) {
	upper := c.programCounter & 0xff00
	c.programCounter = upper | c.inputBus.Read_16()
}

func mInstProgramCounterHighIn(c *Comp, _ uint64) {
	lower := c.programCounter & 0x00ff
	c.programCounter = lower | (c.inputBus.Read_16() << 8)
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
}

func (c *Comp) checkBreakPoint() {
	ok := c.IsBreakPoint(c.programCounter)
	if !ok {
		c.breakpointHit = false
		return
	}
	if c.breakpointHit {
		if c.lastBreakpoint == c.programCounter {
			return
		}
	}
	c.pause = true
	c.lastBreakpoint = c.programCounter
	c.breakpointHit = true
	c.Log("Stop")
	c.LogWithStyle(fmt.Sprintf("● Breakpoint triggered. addr: 0x%04x", c.programCounter), BreakStyle)

	// for _, bp := range c.breakPoints {
	// 	if bp == c.programCounter {
	// 		if c.breakpointHit {
	// 			if c.lastBreakpoint == bp {
	// 				return
	// 			}
	// 		}
	// 		c.pause = true
	// 		c.lastBreakpoint = bp
	// 		c.breakpointHit = true
	// 		c.Log("Stop")
	// 		c.LogWithStyle(fmt.Sprintf("● Breakpoint triggered. addr: 0x%04x", bp), DefaultBreakStyle)
	// 		return
	// 	}
	// }
	// c.breakpointHit = false
}

func (c *Comp) IsBreakPoint(offset uint16) bool {
	for _, bp := range c.breakPoints {
		if bp == offset {
			return true
		}
	}
	return false
}
