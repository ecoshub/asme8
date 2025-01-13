package comp

func mFlagInst(c *Comp, _ uint64) {
	c.aluEnable = true
}

func mInstBreak(c *Comp, _ uint64) {}
