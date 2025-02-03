package terminal

import (
	"github.com/ecoshub/termium/component/config"
	"github.com/ecoshub/termium/component/palette"
	"github.com/ecoshub/termium/component/panel"
	"github.com/ecoshub/termium/component/screen"
	"github.com/ecoshub/termium/component/style"
)

const (
	SCREEN_WIDTH  int = 80
	SCREEN_HEIGHT int = 24
)

const (
	FlagPanelWidth   = 31
	MemoryPanelWidth = 31
	sysPanelWidth    = 71
	codePanelWidth   = 70
)

type Components struct {
	Screen      *screen.Screen
	MainPanel   *panel.Raw
	FlagPanel   *panel.Stack
	MemoryPanel *panel.Stack
	SysLogPanel *panel.Stack
	CodePanel   *panel.Base
}

func NewSetup() (*Components, error) {

	s, err := screen.New(&screen.Config{
		FPSLimit: 60,
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
			BackgroundColor: 234,
		},
	})

	sysLogPanel := panel.NewStackPanel(&config.Config{
		Width:  sysPanelWidth,
		Height: s.TerminalHeight - SCREEN_HEIGHT - 2,
		ContentStyle: &style.Style{
			ForegroundColor: 26,
		},
	})

	flagPanel := panel.NewStackPanel(&config.Config{
		Width:  FlagPanelWidth,
		Height: SCREEN_HEIGHT,
		ContentStyle: &style.Style{
			ForegroundColor: 123,
		},
	})

	memoryPanel := panel.NewStackPanel(&config.Config{
		Width:  MemoryPanelWidth,
		Height: SCREEN_HEIGHT,
		ContentStyle: &style.Style{
			ForegroundColor: 204,
		},
	})

	codePanel := panel.NewBasicPanel(&config.Config{
		Width:  codePanelWidth,
		Height: s.TerminalHeight - SCREEN_HEIGHT - 2,
		ContentStyle: &style.Style{
			ForegroundColor: 26,
		},
	})

	s.Add(mainPanel, FlagPanelWidth+1, 0)
	s.Add(sysLogPanel, 0, SCREEN_HEIGHT+1)
	s.Add(codePanel, sysPanelWidth+1, SCREEN_HEIGHT+1)
	s.Add(memoryPanel, SCREEN_WIDTH+FlagPanelWidth+1, 0)
	s.Add(flagPanel, 0, 0)

	return &Components{
		Screen:      s,
		MainPanel:   mainPanel,
		FlagPanel:   flagPanel,
		MemoryPanel: memoryPanel,
		SysLogPanel: sysLogPanel,
		CodePanel:   codePanel,
	}, nil
}
