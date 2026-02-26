package computer

func mInstVec0Out(c *Computer, _ uint64) {
	c.outputBus.Write_8(IntVec0AddrLow)
	triggerBridge(c)
}

func mInstVec1Out(c *Computer, _ uint64) {
	c.outputBus.Write_8(IntVec1AddrLow)
	triggerBridge(c)
}

func mInstVec2Out(c *Computer, _ uint64) {
	c.outputBus.Write_8(IntVec2AddrLow)
	triggerBridge(c)
}

func mInstVec3Out(c *Computer, _ uint64) {
	c.outputBus.Write_8(IntVec3AddrLow)
	triggerBridge(c)
}
