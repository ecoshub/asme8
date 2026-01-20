package emulator

import (
	"asme8/emulator/src/computer"
	"asme8/emulator/src/emulator/breakpoint"
	"asme8/emulator/src/emulator/panel"
	"asme8/emulator/src/terminal"
	"asme8/emulator/utils"
)

type Config struct {
	EnableDebug bool
}

type Emulator struct {
	computer          *computer.Computer
	terminal          *terminal.Terminal
	breakpointManager *breakpoint.Manager
	memoryPanel       *panel.MemoryPanel
	statePanel        *panel.StatePanel
	codePanel         *panel.CodePanel
	sysLogPanel       *panel.SystemLogPanel
	commandPalette    *CommandPalette
	conf              *Config
}

func New(computer *computer.Computer, conf *Config) *Emulator {
	terminal := computer.GetTerminal()
	devices := computer.GetDevices()
	return &Emulator{
		conf:              conf,
		computer:          computer,
		terminal:          terminal,
		breakpointManager: breakpoint.NewManager(),
		memoryPanel:       panel.NewMemoryPanel(terminal.Components.MemoryPanel, devices),
		statePanel:        panel.NewStatePanel(computer, terminal.Components.FlagPanel),
		sysLogPanel:       panel.NewSystemLogPanel(terminal.Components.SysLogPanel),
		commandPalette:    NewCommandPalette(terminal.Components.Screen.CommandPalette),
		codePanel:         &panel.CodePanel{},
	}
}

func (e *Emulator) CreateCodePanel(codeLines map[uint16]string, hasSymbols bool) {
	e.codePanel = panel.NewCodePanel(codeLines, e.breakpointManager, e.terminal.Components.CodePanel, e.terminal.Components.CodeRulerPanel)
	if !hasSymbols {
		e.sysLogPanel.Log("- Symbol file not found. Can not render code panel.")
		e.sysLogPanel.Log("- To attach a .sym file use --symbol-file <path>")
		e.codePanel.Clear()
		e.codePanel.WriteNoCode()
		return
	} else {
		e.codePanel.Render(0, false)
	}
}

func (e *Emulator) renderPanels() {
	e.statePanel.Render()
	e.memoryPanel.Render()
	e.codePanel.Render(0, true)
}

func (e *Emulator) Run() {
	e.computer.AttachTickEventHandle(e.tickHandle)
	e.computer.AttachBreakEventHandle(e.breakHandle)
	e.terminal.Components.Screen.CommandPalette.ListenKeyEventEnter(e.handleCommand)
	e.terminal.Keyboard.AttachPipeChange(e.pipeChangeHandler)
	e.computer.GetInputBus().AttachBusChange(e.busChangeHandle)
	e.renderPanels()
	e.commandHelp("")
	go e.terminal.Run()
	e.computer.Run()
}

func (e *Emulator) pipeChangeHandler(pipeInput bool) {
	if pipeInput {
		e.sysLogPanel.LogWithStyle("<< Keyboard input directed to emulator.", panel.DefaultStyle1)
	} else {
		e.sysLogPanel.LogWithStyle(">> Keyboard input directed to terminal.", panel.DefaultStyle1)
	}
}

func (e *Emulator) breakHandle(pc uint16) {
	e.sysLogPanel.Log(" # code hit the 'break' computer halted")
}

func (e *Emulator) tickHandle(pc uint16, step uint8) {
	e.statePanel.Render()
	e.CheckBreakPoint(pc)
	e.codePanel.Render(pc, true)
}

func (e *Emulator) busChangeHandle(rw uint8) {
	if rw == utils.IO_WRITE {
		e.memoryPanel.Render()
	}
}

func (e *Emulator) CheckBreakPoint(pc uint16) {
	if !e.breakpointManager.CheckAndHitBreakpoint(pc) {
		return
	}
	e.computer.SetPause(true)
	e.sysLogPanel.Logf("break point!")
}
