package comp

import "fmt"

func mInstProgramCounterIn(c *Comp, _ uint64) {
	c.programCounter = c.addrBus.Read_16()
	checkBreakPoint(c)
}

func mInstProgramCounterLowOutData(c *Comp, _ uint64) {
	c.dataBus.Write_8(uint8(c.programCounter))
	triggerBridge(c)
	checkBreakPoint(c)
}

func mInstProgramCounterHighOutData(c *Comp, _ uint64) {
	c.dataBus.Write_8(uint8(c.programCounter >> 8))
	triggerBridge(c)
	checkBreakPoint(c)
}

func mInstProgramCounterOut(c *Comp, _ uint64) {
	c.addrBus.Write_16(c.programCounter)
	checkBreakPoint(c)
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
