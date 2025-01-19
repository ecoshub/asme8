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

type Components struct {
	Screen      *screen.Screen
	MainPanel   *panel.Raw
	FlagPanel   *panel.Stack
	MemoryPanel *panel.Stack
	SysLogPanel *panel.Stack
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
		Width:  s.TerminalWidth,
		Height: s.TerminalHeight - SCREEN_HEIGHT - 2,
		ContentStyle: &style.Style{
			ForegroundColor: 26,
		},
	})

	mainPanelX := (s.TerminalWidth - mainPanel.Config.Width) / 2

	flagPanel := panel.NewStackPanel(&config.Config{
		Width:  mainPanelX,
		Height: SCREEN_HEIGHT,
		ContentStyle: &style.Style{
			ForegroundColor: 123,
		},
	})

	memoryPanel := panel.NewStackPanel(&config.Config{
		Width:  mainPanelX - 2,
		Height: SCREEN_HEIGHT,
		ContentStyle: &style.Style{
			ForegroundColor: 204,
		},
	})

	s.Add(mainPanel, mainPanelX+1, 0)
	s.Add(sysLogPanel, 0, SCREEN_HEIGHT+1)
	s.Add(memoryPanel, mainPanelX+SCREEN_WIDTH+2, 0)
	s.Add(flagPanel, 0, 0)

	return &Components{
		Screen:      s,
		MainPanel:   mainPanel,
		FlagPanel:   flagPanel,
		MemoryPanel: memoryPanel,
		SysLogPanel: sysLogPanel,
	}, nil
}
