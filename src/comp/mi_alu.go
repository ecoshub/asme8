package comp

func mInstAluAdd(c *Comp, command uint64, _ uint64) {
	if command&MI_ALU_ENABLE <= 0 {
		return
	}

	c.dataBus = c.busX + c.busY
}
