package computer

func mInstBridgeEnable(c *Computer, _ uint64) {
	c.bridgeEnable = true
}

func triggerBridge(c *Computer) {
	if !c.bridgeEnable {
		return
	}
	c.inputBus.Write_8(c.outputBus.Read_8())
}
