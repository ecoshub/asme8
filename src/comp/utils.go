package comp

func (c *Comp) PutToDataBus(val uint8, command uint64) {
	c.dataBus = val
	if command&MI_BRIDGE_ENABLE > 0 {
		if command&MI_BRIDGE_DIR_IN > 0 {
			c.busY = c.dataBus
		}
	}
}

func (c *Comp) PutToBusY(val uint8, command uint64) {
	c.busY = val
	if command&MI_BRIDGE_ENABLE > 0 {
		if command&MI_BRIDGE_DIR_OUT > 0 {
			c.dataBus = c.busY
		}
	}
}
