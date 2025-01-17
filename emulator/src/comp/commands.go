package comp

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func (c *Comp) HandleCommands(command string) {
	switch command {
	case "s":
		if c.pause {
			c.Log("Starts running...")
		} else {
			c.Log("Stop")
		}
		c.pause = !c.pause
		c.pushToCommandPalletHistory(command)
		return
	case "t":
		fallthrough
	case "tick":
		c.singleTicker <- struct{}{}
		c.pushToCommandPalletHistory(command)
		return
	case "mem":
		c.LogMemory()
		c.pushToCommandPalletHistory(command)
		return
	case "h":
		fallthrough
	case "help":
		c.Help()
		return
	case "lsb":
		c.ListBreakPoints()
		return
	case "c":
		fallthrough
	case "clear":
		c.terminalComponents.SysLogPanel.Clear()
		return
	case "r":
		fallthrough
	case "restart":
		c.Restart(true)
		return
	case "exit":
		fallthrough
	case "quit":
		os.Exit(0)
		return
	}

	exists, n, err := SplitNumberCommand(command, "b")
	if exists {
		if err != nil {
			c.Log(err.Error())
			return
		}
		c.AddBreakPoint(int(n))
		c.Logf("● Breakpoint added 0x%x", n)
		c.pushToCommandPalletHistory(command)
		return
	}

	exists, n, err = SplitNumberCommand(command, "rb")
	if exists {
		if err != nil {
			c.Log(err.Error())
			return
		}
		c.RemoveBreakPoint(int(n))
		c.Logf("○ Breakpoint removed 0x%x", n)
		c.pushToCommandPalletHistory(command)
		return
	}

	exists, n, err = SplitNumberCommand(command, "mem")
	if exists {
		if err != nil {
			c.Log(err.Error())
			return
		}
		n = (n / 8) * 8
		c.inspectionMemoryOffset = uint16(n)
		c.LogMemory()
		c.Logf("Memory panel now points to addr 0x%04x", n)
		c.pushToCommandPalletHistory(command)
		return
	}

	exists, n, err = SplitNumberCommand(command, "tick")
	if exists {
		if err != nil {
			c.Log(err.Error())
			return
		}
		for i := 0; i < int(n); i++ {
			c.tick()
			time.Sleep(c.delay)
		}
		c.Logf("clock tick %d time", n)
		c.pushToCommandPalletHistory(command)
		return
	}

	c.LogWithStyle(fmt.Sprintf("# unknown command '%s'", command), DefaultWarningStyle)
}

func SplitNumberCommand(command string, prefix string) (bool, int64, error) {
	if !strings.HasPrefix(command, prefix+" ") {
		return false, 0, fmt.Errorf("error unknown command. command: %s", command)
	}
	tokens := strings.Split(command, prefix+" ")
	number := tokens[1]
	var n int64
	var err error
	if strings.HasPrefix(number, "0x") {
		number = strings.TrimPrefix(number, "0x")
		n, err = strconv.ParseInt(number, 16, 64)
		if err != nil {
			return true, 0, fmt.Errorf("parse hex error. command: %s, err: %s", command, err)
		}
		return true, n, nil
	}
	n, err = strconv.ParseInt(number, 10, 64)
	if err != nil {
		return true, 0, fmt.Errorf("parse number error. command: %s, err: %s", command, err)
	}
	return true, n, nil
}

func (c *Comp) pushToCommandPalletHistory(command string) {
	c.terminalComponents.Screen.CommandPalette.AddToHistory(command)
}

func (c *Comp) Help() {
	c.LogWithStyle("help:", DefaultHelpStyle)
	c.LogWithStyle("'s' ...............: start/stop", DefaultHelpStyle)
	c.LogWithStyle("'tick' | 't' ......: advance clock 1 time", DefaultHelpStyle)
	c.LogWithStyle("'tick <n>' ........: advance clock n time", DefaultHelpStyle)
	c.LogWithStyle("'b <n>' ...........: add breakpoint to address n", DefaultHelpStyle)
	c.LogWithStyle("'rb <n>' ..........: remove breakpoint to address n", DefaultHelpStyle)
	c.LogWithStyle("'lsb' .............: list breakpoints", DefaultHelpStyle)
	c.LogWithStyle("'mem' .............: refresh memory panel", DefaultHelpStyle)
	c.LogWithStyle("'mem <n>' .........: print memory starting with address n", DefaultHelpStyle)
	c.LogWithStyle("'clear' ...........: clear the log panel", DefaultHelpStyle)
	c.LogWithStyle("'restart' | 'r' ...: restart the emulator", DefaultHelpStyle)
	c.LogWithStyle("'exit' | 'quit' ...: exit emulator", DefaultHelpStyle)
	c.LogWithStyle("'help' ............: prints this dialog box", DefaultHelpStyle)
	c.LogWithStyle("", DefaultHelpStyle)
	c.LogWithStyle("note:", DefaultHelpStyle)
	c.LogWithStyle("-  n values can be number of hex with e '0x' prefix", DefaultHelpStyle)
	c.LogWithStyle("-  use ctrl+D to direct keyboard input into emulator", DefaultHelpStyle)
	c.LogWithStyle("", DefaultHelpStyle)
}
