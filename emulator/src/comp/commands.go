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
			c.LogWithStyle("no program load. load program with 'load-bin' or 'load-asm' command", DefaultWarningStyle)
			return
		}
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
	case "hz":
		c.Logf("HZ: %d", time.Second/c.delay)
		return
	case "exit":
		fallthrough
	case "quit":
		os.Exit(0)
		return
	}

	exists, path, err := SplitStringCommand(command, "load-asm")
	if exists {
		if err != nil {
			c.Log(err.Error())
			return
		}
		program, err := utils.AssembleProgram(path)
		if err != nil {
			c.Log(err.Error())
			return
		}
		c.rom.Load(0, program)
		c.programLoaded = true
		c.Restart(true)
		c.Logf("-> Program loaded in to ROM. Start addr is 0x%04x and its %d bytes", 0, len(program))
		c.pushToCommandPalletHistory(command)
		return
	}

	exists, path, err = SplitStringCommand(command, "load-bin")
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
		c.rom.Load(0, program)
		c.programLoaded = true
		c.Restart(true)
		c.Logf("-> Program loaded in to ROM. Start addr is 0x%04x and its %d bytes", 0, len(program))
		c.pushToCommandPalletHistory(command)
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

	exists, param, err := SplitNumberCommand(command, "hz")
	if exists {
		if err != nil {
			c.Log(err.Error())
			return
		}
		delay := time.Second / time.Duration(param)
		from := c.delay
		c.SetDelay(delay)
		c.Restart(true)
		c.Logf("delay changed from %s to %s and computer restarted", from, delay)
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

func SplitStringCommand(command string, prefix string) (bool, string, error) {
	if !strings.HasPrefix(command, prefix+" ") {
		return false, "", fmt.Errorf("error unknown command. command: %s", command)
	}
	tokens := strings.Split(command, prefix+" ")
	param := tokens[1]
	return true, param, nil
}

func (c *Comp) pushToCommandPalletHistory(command string) {
	c.terminalComponents.Screen.CommandPalette.AddToHistory(command)
}

func (c *Comp) Help() {
	c.LogWithStyle("help:", DefaultHelpStyle)
	c.LogWithStyle("- s .....................: start/stop clock", DefaultHelpStyle)
	c.LogWithStyle("- t | tick | tick <n> ...: advance clock 1 or n time", DefaultHelpStyle)
	c.LogWithStyle("- b <n> .................: add breakpoint to address n", DefaultHelpStyle)
	c.LogWithStyle("- rb <n> ................: remove breakpoint to address n", DefaultHelpStyle)
	c.LogWithStyle("- lsb ...................: list breakpoints", DefaultHelpStyle)
	c.LogWithStyle("- mem | mem <n> .........: get or set memory starting with address n", DefaultHelpStyle)
	c.LogWithStyle("- clear .................: clear the log panel", DefaultHelpStyle)
	c.LogWithStyle("- r | restart ...........: restart the emulator", DefaultHelpStyle)
	c.LogWithStyle("- exit | quit ...........: exit emulator", DefaultHelpStyle)
	c.LogWithStyle("- hz | hz <n> ...........: get or set clock speed in hertz (hz)", DefaultHelpStyle)
	c.LogWithStyle("- load-asm | load-bin ...: load program using '.asm' file or '.bin' file", DefaultHelpStyle)
	c.LogWithStyle("- help ..................: prints this dialog box", DefaultHelpStyle)
	c.LogWithStyle("notes:", DefaultHelpStyle)
	c.LogWithStyle("-  n values can be number of hex with e '0x' prefix", DefaultHelpStyle)
	c.LogWithStyle("-  use ctrl+D to direct keyboard input into emulator", DefaultHelpStyle)
	c.LogWithStyle("-  load command always load program starting with addr 0x000", DefaultHelpStyle)
	c.LogWithStyle("", DefaultHelpStyle)
}
