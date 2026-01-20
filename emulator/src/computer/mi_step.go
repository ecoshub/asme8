package computer

func mInstStepInc(c *Computer, _ uint64) {
	c.step++
}

func mInstStepClr(c *Computer, _ uint64) {
	c.step = 0
}
