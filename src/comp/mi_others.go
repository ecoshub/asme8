package comp

import (
	"fmt"
	"os"
)

func mFlagInst(_ *Comp, _ uint64, _ uint64) {}

func mInstBreak(c *Comp, _ uint64, _ uint64) {
	fmt.Printf("## break at: %d\n", c.programCounter)
	os.Exit(0)
}
