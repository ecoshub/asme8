package comp

import (
	"asme8/assembler/src/utils"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func (c *Comp) HandleCommands(command string) {
	switch command {
	case "s":
		if !c.programLoaded {
			c.LogWithStyle("no program load. load program with 'load-bin' or 'load-asm' command", WarningStyle)
			return
		}
		if c.pause {
			c.Log("Starts running...")
		} else {
			c.Log("Stop")
		}
		c.pause = !c.pause
		// c.pushToCommandPalletHistory(command)
		return
	case "cyc":
		fallthrough
	case "cycle":
		c.Logf("total cycle: %d", c.cycleCount)
		// c.pushToCommandPalletHistory(command)
		return
	case "ccyc":
		fallthrough
	case "clear cycle":
		c.cycleCount = 0
		c.Logf("total cycle: %d", c.cycleCount)
		// c.pushToCommandPalletHistory(command)
		return
	case "t":
		fallthrough
	case "tick":
		if !c.running {
			c.Logf("! program is in break state. try restart using 'r'")
			return
		}
		c.singleTicker <- struct{}{}
		// c.pushToCommandPalletHistory(command)
		return
	case "mem":
		c.LogMemory()
		c.pushToCommandPalletHistory(command)
		return
	case "f":
		c.scrollFocus = "nan"
		c.Log("Scroll focus 'Disabled'")
		c.pushToCommandPalletHistory(command)
		return
	case "h":
		fallthrough
	case "help":
		// c.Help()
		return
	// case "lbp":
	// c.ListBreakPoints()
	// return
	// case "cbp":
	// 	c.ClearBreakPoints()
	// 	c.terminal.Components.CodeRulerPanel.Clear()
	// 	c.LogCodePanel(true)
	// 	return
	case "c":
		fallthrough
	case "clear":
		c.terminal.Components.SysLogPanel.Clear()
		c.pushToCommandPalletHistory(command)
		return
	case "r":
		fallthrough
	case "restart":
		c.Restart(true)
		return
	case "hz":
		c.Logf("HZ: %d", time.Second/c.Config.Delay)
		c.pushToCommandPalletHistory(command)
		return
	case "exit":
		fallthrough
	case "quit":
		os.Exit(0)
		return
	}

	path, exists, err := SplitStringCommand(command, "load-asm")
	if exists {
		if err != nil {
			c.Log(err.Error())
			return
		}
		program, code, err := utils.AssembleProgram(path)
		if err != nil {
			c.Log(err.Error())
			return
		}

		c.ResolveSymbolFromFile(code)

		c.Config.Program = program
		c.LoadProgram()

		c.Restart(true)
		c.Logf("-> Program loaded in to ROM. Start addr is 0x%04x and its %d bytes", 0, len(program))
		// c.pushToCommandPalletHistory(command)
		return
	}

	path, exists, err = SplitStringCommand(command, "load-bin")
	if exists {
		if err != nil {
			c.Log(err.Error())
			return
		}
		program, err := utils.ReadProgram(path)
		if err != nil {
			c.Log(err.Error())
			return
		}

		c.Config.Program = program
		c.LoadProgram()

		c.Restart(true)
		c.Logf("-> Program loaded in to ROM. Start addr is 0x%04x and its %d bytes", 0, len(program))
		// c.pushToCommandPalletHistory(command)
		return
	}

	n, exists, err := SplitNumberCommand(command, "b")
	if exists {
		if err != nil {
			c.Log(err.Error())
			return
		}
		ok := c.AddBreakPoint(int(n))
		if !ok {
			c.Logf("# Breakpoint already exists. point: 0x%x", n)
			return
		}
		c.Logf("● Breakpoint added 0x%x", n)
		c.terminal.Components.CodeRulerPanel.Clear()
		c.LogCodePanel(true)
		c.pushToCommandPalletHistory(command)
		return
	}

	n, exists, err = SplitNumberCommand(command, "rb")
	if exists {
		if err != nil {
			c.Log(err.Error())
			return
		}
		ok := c.RemoveBreakPoint(int(n))
		if !ok {
			return
		}
		c.Logf("○ Breakpoint removed 0x%x", n)
		c.terminal.Components.CodeRulerPanel.Clear()
		c.LogCodePanel(true)
		c.pushToCommandPalletHistory(command)
		return
	}

	cmd, exists, err := SplitStringCommand(command, "mem")
	if exists {
		if err != nil {
			c.Log(err.Error())
			return
		}
		if cmd == "STACK" {
			c.inspectionMemoryOffset = StackStart
			c.inspectionMemoryCustomOffset = 0
			c.LogMemory()
			c.Logf("Memory panel now points to addr 0x%04x", StackStart)
			c.pushToCommandPalletHistory(command)
			return
		}
		dev, ok := c.GetDevice(cmd)
		if ok {
			start, _ := dev.GetRange()
			c.inspectionMemoryOffset = start
			c.inspectionMemoryCustomOffset = 0
			c.LogMemory()
			c.Logf("Memory panel now points to addr 0x%04x", start)
			c.pushToCommandPalletHistory(command)
			return
		}
	}

	cmd, exists, err = SplitStringCommand(command, "f")
	if exists {
		if err != nil {
			c.Log(err.Error())
			return
		}
		switch cmd {
		case "mem":
			c.Log("Scroll focus changed to 'mem'")
			c.scrollFocus = "mem"
		case "code":
			c.Log("Scroll focus changed to 'code'")
			c.scrollFocus = "code"
		case "nan":
			c.Log("Scroll focus 'Disabled'")
			c.scrollFocus = "nan"
		}
		c.pushToCommandPalletHistory(command)
		return
	}

	n, exists, err = SplitNumberCommand(command, "mem")
	if exists {
		if err != nil {
			c.Log(err.Error())
			return
		}
		n = (n / 8) * 8
		c.inspectionMemoryOffset = uint16(n)
		c.inspectionMemoryCustomOffset = 0
		c.LogMemory()
		c.Logf("Memory panel now points to addr 0x%04x", n)
		c.pushToCommandPalletHistory(command)
		return
	}

	n, exists, err = SplitNumberCommand(command, "tick")
	if exists {
		if err != nil {
			c.Log(err.Error())
			return
		}
		for i := 0; i < int(n); i++ {
			c.tick()
			time.Sleep(c.Config.Delay)
		}
		c.Logf("clock tick %d time", n)
		c.pushToCommandPalletHistory(command)
		return
	}

	param, exists, err := SplitNumberCommand(command, "hz")
	if exists {
		if err != nil {
			c.Log(err.Error())
			return
		}
		delay := time.Second / time.Duration(param)
		from := c.Config.Delay
		c.SetDelay(delay)
		c.Restart(true)
		c.Logf("delay changed from %s to %s and computer restarted", from, delay)
		c.pushToCommandPalletHistory(command)
		return
	}

	c.LogWithStyle(fmt.Sprintf("# unknown command '%s'", command), WarningStyle)
}

func SplitNumberCommand(command string, prefixes ...string) (int64, bool, error) {
	param, exists, err := SplitStringCommand(command, prefixes...)
	if err != nil {
		return 0, exists, err
	}

	var n int64
	if strings.HasPrefix(param, "0x") {
		param = strings.TrimPrefix(param, "0x")
		n, err = strconv.ParseInt(param, 16, 64)
		if err != nil {
			return 0, true, fmt.Errorf("parse hex error. command: %s, err: %s", command, err)
		}
		return n, true, nil
	}
	n, err = strconv.ParseInt(param, 10, 64)
	if err != nil {
		return 0, true, fmt.Errorf("parse number error. command: %s, err: %s", command, err)
	}
	return n, true, nil
}

func SplitStringCommand(command string, prefixes ...string) (string, bool, error) {
	for _, prefix := range prefixes {
		if strings.HasPrefix(command, prefix+" ") {
			tokens := strings.Split(command, " ")
			param := tokens[1]
			return param, true, nil
		}
	}
	return "", false, fmt.Errorf("error unknown command. command: %s", command)
}

// func (c *Comp) pushToCommandPalletHistory(command string) {
// 	if command == c.lastCommand {
// 		return
// 	}
// 	c.terminal.Components.Screen.CommandPalette.AddToHistory(command)
// }

func (c *Comp) Help() {
	c.LogWithStyle("help:", HelpStyle)
	c.LogWithStyle("- s ...................: start/stop clock", HelpStyle)
	c.LogWithStyle("- t | tick | tick <n> .: advance clock 1 or n time", HelpStyle)
	c.LogWithStyle("- b <n> ...............: add breakpoint to address n", HelpStyle)
	c.LogWithStyle("- rb <n> ..............: remove breakpoint to address n", HelpStyle)
	c.LogWithStyle("- lbp .................: list breakpoints", HelpStyle)
	c.LogWithStyle("- cbp .................: clear all breakpoints", HelpStyle)
	c.LogWithStyle("- mem | mem <n> .......: get or set memory starting with address n", HelpStyle)
	c.LogWithStyle("- r | restart .........: restart the emulator", HelpStyle)
	c.LogWithStyle("- f <mem/code/nan> ....: focus scroll. use left and right arrows", HelpStyle)
	c.LogWithStyle("- hz | hz <n> .........: get or set clock speed in hertz (hz)", HelpStyle)
	c.LogWithStyle("- load-asm | load-bin .: load program using '.asm' file or '.bin' file", HelpStyle)
	c.LogWithStyle("- cycle | cyc .........: print total cycle run", HelpStyle)
	c.LogWithStyle("- clear cycle | ccyc ..: reset total cycle", HelpStyle)
	c.LogWithStyle("- page | page <n> .....: toggle page mode, set page offset", HelpStyle)
	c.LogWithStyle("- exit | quit .........: exit emulator", HelpStyle)
	c.LogWithStyle("- clear ...............: clear the log panel", HelpStyle)
	c.LogWithStyle("- help ................: prints this dialog box", HelpStyle)
	c.LogWithStyle("* note:  n values can be number of hex with e '0x' prefix", HelpStyle)
	c.LogWithStyle("* note:  use ctrl+D to direct keyboard input into emulator", HelpStyle)
	c.LogWithStyle("", HelpStyle)
}
