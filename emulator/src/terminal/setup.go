package terminal

import (
	"github.com/ecoshub/termium/component/config"
	"github.com/ecoshub/termium/component/palette"
	"github.com/ecoshub/termium/component/panel"
	"github.com/ecoshub/termium/component/screen"
	"github.com/ecoshub/termium/component/style"
	"github.com/ecoshub/termium/utils"
)

const (
	SCREEN_WIDTH  int = 80
	SCREEN_HEIGHT int = 24
)

const (
	SysLogPanelHeight int = 5
)

type Components struct {
	Screen      *screen.Screen
	MainPanel   *panel.Raw
	MemoryPanel *panel.Stack
	FlagPanel   *panel.Stack
	DebugPanel  *panel.Stack
	SysLogPanel *panel.Stack
}

func NewSetup() (*Components, error) {

	s, err := screen.New(&screen.Config{
		CommandPaletteConfig: &palette.Config{
			Prompt: "> ",
			Style: &style.Style{
				ForegroundColor: 83,
			},
		},
	})
	if err != nil {
		return nil, err
	}

	mainPanel := panel.NewRawPanel(&config.Config{
		Width:  SCREEN_WIDTH,
		Height: SCREEN_HEIGHT,
		ContentStyle: &style.Style{
			ForegroundColor: 195,
			BackgroundColor: 232,
		},
	})

	sysLogPanel := panel.NewStackPanel(&config.Config{
		Width:  utils.TerminalWith,
		Height: SysLogPanelHeight,
		Title:  "System Logs:",
		TitleStyle: &style.Style{
			ForegroundColor: 27,
			BackgroundColor: 236,
		},
		RenderTitle: true,
		ContentStyle: &style.Style{
			ForegroundColor: 27,
			BackgroundColor: 236,
		},
	})

	debugPanel := panel.NewStackPanel(&config.Config{
		Width:  utils.TerminalWith,
		Height: utils.TerminalHeight - SCREEN_HEIGHT - SysLogPanelHeight - 1,
		Title:  "Debug Logs:",
		TitleStyle: &style.Style{
			ForegroundColor: 70,
			BackgroundColor: 235,
		},
		RenderTitle: true,
		ContentStyle: &style.Style{
			ForegroundColor: 70,
			BackgroundColor: 235,
		},
	})

	mainPanelX := (utils.TerminalWith - mainPanel.Config.Width) / 2

	memoryPanel := panel.NewStackPanel(&config.Config{
		Width:  mainPanelX,
		Height: SCREEN_HEIGHT,
		Title:  "Memory:",
		TitleStyle: &style.Style{
			ForegroundColor: 123,
			BackgroundColor: 234,
		},
		RenderTitle: true,
		ContentStyle: &style.Style{
			ForegroundColor: 123,
			BackgroundColor: 234,
		},
	})

	flagPanel := panel.NewStackPanel(&config.Config{
		Width:  mainPanelX,
		Height: SCREEN_HEIGHT,
		Title:  "Flags:",
		TitleStyle: &style.Style{
			ForegroundColor: 204,
			BackgroundColor: 234,
		},
		RenderTitle: true,
		ContentStyle: &style.Style{
			ForegroundColor: 204,
			BackgroundColor: 234,
		},
	})

	s.Add(mainPanel, mainPanelX+1, 0)
	s.Add(debugPanel, 0, SCREEN_HEIGHT)
	s.Add(sysLogPanel, 0, SCREEN_HEIGHT+debugPanel.Config.Height)
	s.Add(memoryPanel, 0, 0)
	s.Add(flagPanel, mainPanelX+SCREEN_WIDTH, 0)

	return &Components{
		Screen:      s,
		MainPanel:   mainPanel,
		DebugPanel:  debugPanel,
		MemoryPanel: memoryPanel,
		FlagPanel:   flagPanel,
		SysLogPanel: sysLogPanel,
	}, nil
}
