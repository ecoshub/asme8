package comp

func mInstStepInc(c *Comp, _ uint64) {
	c.step++
}

func mInstStepClr(c *Comp, _ uint64) {
	c.step = 0
}
