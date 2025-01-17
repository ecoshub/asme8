package comp

func (c *Comp) AddBreakPoint(no int) {
	c.breakPoints = append(c.breakPoints, uint16(no))
}

func (c *Comp) RemoveBreakPoint(no int) {
	breakPoints := make([]uint16, 0, len(c.breakPoints))
	for _, bp := range c.breakPoints {
		if bp == uint16(no) {
			continue
		}
		breakPoints = append(breakPoints, bp)
	}
	c.breakPoints = breakPoints
}
