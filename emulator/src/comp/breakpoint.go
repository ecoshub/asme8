package comp

func (c *Comp) AddBreakPoint(no int) bool {
	for _, bp := range c.breakPoints {
		if bp == uint16(no) {
			return false
		}
	}
	c.breakPoints = append(c.breakPoints, uint16(no))
	return true
}

func (c *Comp) RemoveBreakPoint(no int) bool {
	breakPoints := make([]uint16, 0, len(c.breakPoints))
	found := false
	for _, bp := range c.breakPoints {
		if bp == uint16(no) {
			found = true
			continue
		}
		breakPoints = append(breakPoints, bp)
	}
	c.breakPoints = breakPoints
	return found
}
