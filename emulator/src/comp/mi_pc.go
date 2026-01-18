package comp

func mInstProgramCounterOutAddr(c *Comp, _ uint64) {
	c.addrBus.Write_16(c.programCounter)
	// c.LogCodePanel(false)
	// c.checkBreakPoint()
}

func mInstProgramCounterInAddr(c *Comp, _ uint64) {
	c.programCounter = c.addrBus.Read_16()
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

// func (c *Comp) checkBreakPoint() {
// 	ok := c.IsBreakPoint(c.programCounter)
// 	if !ok {
// 		c.breakpointHit = false
// 		return
// 	}
// 	if c.breakpointHit {
// 		if c.lastBreakpoint == c.programCounter {
// 			return
// 		}
// 	}
// 	c.pause = true
// 	c.lastBreakpoint = c.programCounter
// 	c.breakpointHit = true
// 	c.Log("Stop")
// 	c.LogWithStyle(fmt.Sprintf("● Breakpoint triggered. addr: 0x%04x", c.programCounter), BreakStyle)

// 	c.forcePageEnable = false
// 	c.LogCodePanel(true)
// }

// func (c *Comp) IsBreakPoint(offset uint16) bool {
// 	for _, bp := range c.breakPoints {
// 		if bp == offset {
// 			return true
// 		}
// 	}
// 	return false
// }
