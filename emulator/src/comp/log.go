package comp

import "fmt"

func (c *Comp) Logf(format string, a ...any) {
	if c.terminalComponents == nil {
		fmt.Printf(format, a...)
		return
	}
	c.terminalComponents.DebugPanel.Push(fmt.Sprintf(format, a...))
}
