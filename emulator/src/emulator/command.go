package emulator

import (
	"asme8/assembler/src/utils"
	"asme8/emulator/src/emulator/panel"
	"asme8/emulator/src/terminal"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/ecoshub/termium/component/style"
)

const (
	CommandHelp                     string = "help"
	CommandShortStartStop           string = "s"
	CommandShortTick                string = "t"
	CommandTick                     string = "tick"
	CommandShortClear               string = "c"
	CommandClear                    string = "clear"
	CommandShortRestart             string = "r"
	CommandRestart                  string = "restart"
	CommandShortQuit                string = "q"
	CommandExit                     string = "exit"
	CommandQuit                     string = "quit"
	CommandShortAddBreakpoint       string = "b"
	CommandShortAddRemoveBreakpoint string = "rb"
	CommandShortMemory              string = "mem"
	CommandShortHertz               string = "hz"
	CommandShortHelp                string = "h"
	CommandShortListBreakpoints     string = "lb"
	CommandShortClearBreakpoints    string = "cb"
	CommandLoadAsmFile              string = "load-asm"
	CommandLoadBinFile              string = "load-bin"
	CommandShortScrollCodePanelDown string = ">"
	CommandShortScrollCodePanelUp   string = "<"
	CommandShortScrollCodePanelTop  string = "<<"
	CommandShortScrollCodePanelBot  string = ">>"
	CommandToggleCodeTracking       string = "track"
	CommandShortToggleCodeTracking  string = "tr"
)

func (e *Emulator) handleCommand(input string) {
	if input == "" {
		lastCommand := e.commandPalette.GetLastCommand()
		if lastCommand != "" {
			e.handleCommand(lastCommand)
			return
		}
	}

	tokens := strings.Split(input, " ")
	cmd := tokens[0]
	switch cmd {
	case CommandShortStartStop:
		e.commandStartStop(input)
	case CommandShortTick, CommandTick:
		e.commandTick(input)
	case CommandShortClear, CommandClear:
		e.commandClear(input)
	case CommandShortRestart, CommandRestart:
		e.commandRestart(input)
	case CommandShortQuit, CommandQuit, CommandExit:
		e.commandExit(input)
	case CommandShortAddBreakpoint:
		e.commandAddBreakpoint(input)
	case CommandShortAddRemoveBreakpoint:
		e.commandRemoveBreakpoint(input)
	case CommandShortMemory:
		e.commandMemory(input)
	case CommandShortHertz:
		e.commandChangeClockSpeed(input)
	case CommandHelp, CommandShortHelp:
		e.commandHelp(input)
	case CommandShortListBreakpoints:
		e.commandListBreakpoints(input)
	case CommandShortClearBreakpoints:
		e.commandClearBreakpoints(input)
	case CommandLoadAsmFile:
		e.commandLoadAsmFile(input)
	case CommandLoadBinFile:
		e.commandLoadBinFile(input)
	case CommandShortScrollCodePanelDown:
		e.commandScrollCodePanelDown(input)
	case CommandShortScrollCodePanelUp:
		e.commandScrollCodePanelUp(input)
	case CommandShortScrollCodePanelTop:
		e.commandScrollCodePanelTop(input)
	case CommandShortScrollCodePanelBot:
		e.commandScrollCodePanelBot(input)
	case CommandShortToggleCodeTracking, CommandToggleCodeTracking:
		e.commandToggleCodeTracking(input)
	default:
		e.sysLogPanel.LogWithStyle(fmt.Sprintf("command not found. command: '%s'", input), style.DefaultStyleError)
	}
	e.commandPalette.AddToCommandHistory(input)
}

func (e *Emulator) commandToggleCodeTracking(_ string) {
	tracking := e.codePanel.ToggleTrackExecution()
	if tracking {
		e.sysLogPanel.LogWithStyle("tracking enabled", style.DefaultStyleInfo)
	} else {
		e.sysLogPanel.LogWithStyle("tracking disabled", style.DefaultStyleInfo)
	}
}

func (e *Emulator) setTrackExecution(tracking bool) {
	if e.conf.Headless {
		return
	}
	if e.codePanel.GetTrackExecution() == tracking {
		return
	}
	e.codePanel.SetTrackExecution(tracking)
	if tracking {
		e.sysLogPanel.LogWithStyle("tracking enabled", style.DefaultStyleInfo)
	} else {
		e.sysLogPanel.LogWithStyle("tracking disabled", style.DefaultStyleInfo)
	}
}

func (e *Emulator) commandScrollCodePanelDown(_ string) {
	e.setTrackExecution(false)
	e.codePanel.SetSkipRequest(1)
	e.codePanel.Render(e.computer.GetProgramCounter(), true)
	skip := e.codePanel.GetSkip()
	e.sysLogPanel.LogWithStyle(fmt.Sprintf(" -- page %d --", skip+1), style.DefaultStyleInfo)
}

func (e *Emulator) commandScrollCodePanelUp(_ string) {
	e.setTrackExecution(false)
	e.codePanel.SetSkipRequest(-1)
	e.codePanel.Render(e.computer.GetProgramCounter(), true)
	skip := e.codePanel.GetSkip()
	e.sysLogPanel.LogWithStyle(fmt.Sprintf(" -- page %d --", skip+1), style.DefaultStyleInfo)
}

func (e *Emulator) commandScrollCodePanelTop(_ string) {
	e.setTrackExecution(false)
	e.codePanel.SetSkipTop()
	e.codePanel.Render(e.computer.GetProgramCounter(), true)
	skip := e.codePanel.GetSkip()
	e.sysLogPanel.LogWithStyle(fmt.Sprintf(" -- page %d --", skip+1), style.DefaultStyleInfo)
}

func (e *Emulator) commandScrollCodePanelBot(_ string) {
	e.setTrackExecution(false)
	e.codePanel.SetSkipBot()
	e.codePanel.Render(e.computer.GetProgramCounter(), true)
	skip := e.codePanel.GetSkip()
	e.sysLogPanel.LogWithStyle(fmt.Sprintf(" -- page %d --", skip+1), style.DefaultStyleInfo)
}

func (e *Emulator) commandLoadAsmFile(input string) {
	param, ok, err := splitStringCommand(input, CommandLoadAsmFile)
	if err != nil {
		e.sysLogPanel.LogWithStyle(err.Error(), style.DefaultStyleError)
		return
	}

	if !ok {
		e.sysLogPanel.LogWithStyle(fmt.Sprintf("command not found. command: '%s'", input), style.DefaultStyleError)
		return
	}

	program, code, err := utils.AssembleProgram(param)
	if err != nil {
		e.sysLogPanel.LogWithStyle(fmt.Sprintf("program assembly error. err: %s", err), style.DefaultStyleError)
		return
	}

	codeLines, hasSymbols, err := ResolveSymbolFromFile(code)
	e.CreateCodePanel(codeLines, hasSymbols)

	e.computer.Config.Program = program
	e.computer.LoadProgram()
	e.computer.SetProgramLoaded()

	e.computer.Restart(true)
	e.renderPanels()

	e.sysLogPanel.LogWithStyle(fmt.Sprintf("-> Program loaded in to ROM. Start addr is 0x%04x and its %d bytes", 0, len(program)), style.DefaultStyleInfo)
}

func (e *Emulator) commandLoadBinFile(input string) {
	param, ok, err := splitStringCommand(input, CommandLoadBinFile)
	if err != nil {
		e.sysLogPanel.LogWithStyle(err.Error(), style.DefaultStyleError)
		return
	}

	if !ok {
		e.sysLogPanel.LogWithStyle(fmt.Sprintf("command not found. command: '%s'", input), style.DefaultStyleError)
		return
	}

	program, err := utils.ReadProgram(param)
	if err != nil {
		e.sysLogPanel.LogWithStyle(fmt.Sprintf("program assembly error. err: %s", err), style.DefaultStyleError)
		return
	}

	codeLines, hasSymbols, err := ReadSymbolFileAndCreateCodePanel(param, "")
	e.CreateCodePanel(codeLines, hasSymbols)

	e.computer.Config.Program = program
	e.computer.LoadProgram()
	e.computer.SetProgramLoaded()

	e.computer.Restart(true)
	e.renderPanels()

	e.sysLogPanel.LogWithStyle(fmt.Sprintf("program loaded in to ROM. Start addr is 0x%04x and its %d bytes.", 0, len(program)), style.DefaultStyleInfo)
	if hasSymbols {
		e.sysLogPanel.LogWithStyle("also symbols found (<filename>.sym) and load.", style.DefaultStyleInfo)
	}
}

func (e *Emulator) commandClearBreakpoints(_ string) {
	e.breakpointManager.Clear()
	e.sysLogPanel.LogWithStyle("all breakpoints cleared", style.DefaultStyleInfo)
}

func (e *Emulator) commandListBreakpoints(_ string) {
	lines := listBreakpoints(e.breakpointManager)
	for _, line := range lines {
		e.sysLogPanel.LogWithStyle(line, panel.BreakStyle)
	}
}

func (e *Emulator) commandChangeClockSpeed(input string) {
	if input == CommandShortHertz {
		e.sysLogPanel.LogWithStyle(fmt.Sprintf("cpu emulation speed is %d Hz", time.Second/e.computer.Config.Delay), panel.DefaultStyle1)
		return
	}
	number, ok, err := splitNumberCommand(input, CommandShortHertz)
	if err != nil {
		e.sysLogPanel.LogWithStyle(err.Error(), style.DefaultStyleError)
		return
	}

	if !ok {
		e.sysLogPanel.LogWithStyle(fmt.Sprintf("command not found. command: '%s'", input), style.DefaultStyleError)
		return
	}

	delay := time.Second / time.Duration(number)
	e.computer.SetDelay(delay)
	e.computer.Restart(true)
	e.renderPanels()
	e.sysLogPanel.Clear()
	lastDelayValue := e.computer.Config.Delay
	e.computer.Config.Delay = delay
	e.sysLogPanel.LogWithStyle(fmt.Sprintf("cpu emulation speed changed from %d Hz to %d Hz.", time.Second/lastDelayValue, time.Second/delay), panel.DefaultStyle1)
}

func (e *Emulator) commandMemory(input string) {
	if input == CommandShortMemory+" STACK" || input == CommandShortMemory+" stack" {
		number := e.computer.GetStartOfStack()
		e.memoryPanel.SetOffset(number)
		e.sysLogPanel.LogWithStyle(fmt.Sprintf("now memory page points to 0x%x [stack]", number), panel.DefaultStyle1)
		return
	}

	if input == CommandShortMemory+" RAM" || input == CommandShortMemory+" ram" {
		number := e.computer.GetRamStart()
		e.memoryPanel.SetOffset(number)
		e.sysLogPanel.LogWithStyle(fmt.Sprintf("now memory page points to 0x%x [ram]", number), panel.DefaultStyle1)
		return
	}

	number, ok, err := splitNumberCommand(input, CommandShortMemory)
	if err != nil {
		e.sysLogPanel.LogWithStyle(err.Error(), style.DefaultStyleError)
		return
	}

	if !ok {
		e.sysLogPanel.LogWithStyle(fmt.Sprintf("command not found. command: '%s'", input), style.DefaultStyleError)
		return
	}

	e.memoryPanel.SetOffset(uint16(number))
	e.sysLogPanel.LogWithStyle(fmt.Sprintf("now memory page points to 0x%x", number), panel.DefaultStyle1)
}

func (e *Emulator) commandRemoveBreakpoint(input string) {
	number, ok, err := splitNumberCommand(input, CommandShortAddRemoveBreakpoint)
	if err != nil {
		e.sysLogPanel.LogWithStyle(err.Error(), style.DefaultStyleError)
		return
	}

	if !ok {
		e.sysLogPanel.LogWithStyle(fmt.Sprintf("command not found. command: '%s'", input), style.DefaultStyleError)
		return
	}

	offset := e.codePanel.UseCodeOffsetMap(uint16(number))
	e.breakpointManager.Remove(offset)

	e.sysLogPanel.LogWithStyle(fmt.Sprintf("○ Breakpoint removed 0x%x", offset), panel.BreakStyle)
	e.codePanel.Render(e.computer.GetProgramCounter(), true)
}

func (e *Emulator) commandAddBreakpoint(input string) {
	number, ok, err := splitNumberCommand(input, CommandShortAddBreakpoint)
	if err != nil {
		e.sysLogPanel.LogWithStyle(err.Error(), style.DefaultStyleError)
		return
	}

	if !ok {
		e.sysLogPanel.LogWithStyle(fmt.Sprintf("command not found. command: '%s'", input), style.DefaultStyleError)
		return
	}

	offset := e.codePanel.UseCodeOffsetMap(uint16(number))
	exists := e.breakpointManager.AddBreakpoint(offset)
	if exists {
		e.sysLogPanel.LogWithStyle("breakpoint already exists.", style.DefaultStyleInfo)
		return
	}

	if offset != uint16(number) {
		e.sysLogPanel.LogWithStyle(fmt.Sprintf("● Breakpoint added 0x%x [fixed to instruction start]", offset), panel.BreakStyle)
	} else {
		e.sysLogPanel.LogWithStyle(fmt.Sprintf("● Breakpoint added 0x%x", offset), panel.BreakStyle)
	}

	e.codePanel.Render(e.computer.GetProgramCounter(), true)
}

func (e *Emulator) commandExit(_ string) {
	e.computer.Stop()
	// static reset function
	terminal.ResetScreen()
	os.Exit(0)
}

func (e *Emulator) commandRestart(_ string) {
	e.computer.Restart(true)
	e.renderPanels()
	e.commandPalette.ClearCommandHistory()
	e.sysLogPanel.Clear()
}

func (e *Emulator) commandClear(_ string) {
	e.sysLogPanel.Clear()
	e.commandPalette.ClearCommandHistory()
}

func (e *Emulator) commandTick(_ string) {
	if !e.computer.IsRunning() {
		e.sysLogPanel.LogWithStyle("# program is in break state. try restart using 'r'", panel.WarningStyle)
		return
	}
	e.computer.Tick()
}

func (e *Emulator) commandStartStop(_ string) {
	if !e.computer.IsProgramLoaded() {
		e.sysLogPanel.LogWithStyle("no program load. load program with 'load-bin' or 'load-asm' command", panel.WarningStyle)
		return
	}
	if e.computer.IsPause() {
		e.sysLogPanel.LogWithStyle("Running", style.DefaultStyleInfo)
		e.computer.SetPause(false)
		return
	}
	e.sysLogPanel.LogWithStyle("Pause", style.DefaultStyleInfo)
	e.computer.SetPause(true)
}

func (e *Emulator) commandHelp(_ string) {
	e.sysLogPanel.LogWithStyle("help:", panel.HelpStyle)
	e.sysLogPanel.LogWithStyle("- help | h  ...........: display this help dialog", panel.HelpStyle)
	e.sysLogPanel.LogWithStyle("- s ...................: start or stop the clock", panel.HelpStyle)
	e.sysLogPanel.LogWithStyle("- t ...................: advance the clock by 1 cycle", panel.HelpStyle)
	e.sysLogPanel.LogWithStyle("- b <n> ...............: add a breakpoint at address n", panel.HelpStyle)
	e.sysLogPanel.LogWithStyle("- rb <n> ..............: remove the breakpoint at address n", panel.HelpStyle)
	e.sysLogPanel.LogWithStyle("- lb ..................: list all breakpoints", panel.HelpStyle)
	e.sysLogPanel.LogWithStyle("- cb ..................: clear all breakpoints", panel.HelpStyle)
	e.sysLogPanel.LogWithStyle("- mem <n> .............: set memory page starting at address n (or 'stack'/'ram')", panel.HelpStyle)
	e.sysLogPanel.LogWithStyle("- r | restart .........: restart the emulator", panel.HelpStyle)
	e.sysLogPanel.LogWithStyle("- hz | hz <n> .........: get or set clock speed in hertz (Hz)", panel.HelpStyle)
	e.sysLogPanel.LogWithStyle("- load-asm | load-bin .: load a program from a '.asm' or '.bin' file", panel.HelpStyle)
	e.sysLogPanel.LogWithStyle("- q | quit | exit .....: exit the emulator", panel.HelpStyle)
	e.sysLogPanel.LogWithStyle("- c | clear ...........: clear the log panel", panel.HelpStyle)
	e.sysLogPanel.LogWithStyle("- > | < | << | >> .....: scroll code panel down/up/top/bottom", panel.HelpStyle)
	e.sysLogPanel.LogWithStyle("- track | tr ..........: toggle code execution tracking", panel.HelpStyle)
	e.sysLogPanel.LogWithStyle("", panel.HelpStyle)
	e.sysLogPanel.LogWithStyle("tips:", panel.HelpStyle)
	e.sysLogPanel.LogWithStyle("- <n> values can be decimal or hexadecimal with a '0x' prefix", panel.HelpStyle)
	e.sysLogPanel.LogWithStyle("- use Ctrl+D to direct keyboard input into the emulator", panel.HelpStyle)
	e.sysLogPanel.LogWithStyle("- 'load-bin' auto-loads a matching .sym file", panel.HelpStyle)
}
