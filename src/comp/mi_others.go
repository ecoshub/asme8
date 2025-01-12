package comp

import (
	"fmt"
	"os"
)

func mFlagInst(c *Comp, _ uint64) {
	c.aluEnable = true
}

func mInstBreak(c *Comp, _ uint64) {
	fmt.Printf("## break at: %d\n", c.programCounter)
	os.Exit(0)
}
