package comp

func mInstBridgeEnable(c *Comp, _ uint64) {
	c.bridgeEnable = true
}

func triggerBridge(c *Comp) {
	if !c.bridgeEnable {
		return
	}
	c.inputBus.Write_8(c.outputBus.Read_8())
}
