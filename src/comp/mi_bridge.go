package comp

const (
	BRIDGE_DIR_OUT uint8 = 0
	BRIDGE_DIR_IN  uint8 = 1
)

func mInstBridgeDIR_OUT(c *Comp, _ uint64, _ uint64) {
	c.bridgeDir = BRIDGE_DIR_OUT
}

func mInstBridgeEnable(c *Comp, _ uint64, _ uint64) {
	c.bridgeEnable = true
}

func triggerBridge(c *Comp) {
	if !c.bridgeEnable {
		return
	}
	if c.bridgeDir == BRIDGE_DIR_IN {
		c.busY.Write_8(c.dataBus.Read_8())
		return
	}
	if c.bridgeDir == BRIDGE_DIR_OUT {
		c.dataBus.Write_8(c.busY.Read_8())
		return
	}
}
