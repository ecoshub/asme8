package panel

import (
	"fmt"
	"strings"

	"github.com/ecoshub/termium/component/panel"
	"github.com/ecoshub/termium/component/style"
)

type SystemLogPanel struct {
	sysLogPanel *panel.Stack
}

func NewSystemLogPanelCLI() *SystemLogPanel {
	return &SystemLogPanel{}
}

func NewSystemLogPanel(sysLogPanel *panel.Stack) *SystemLogPanel {
	return &SystemLogPanel{
		sysLogPanel: sysLogPanel,
	}
}

func (sp *SystemLogPanel) Log(str string) {
	if sp.sysLogPanel == nil {
		fmt.Println(str)
		return
	}
	sp.sysLogPanel.Push(str)
}

func (sp *SystemLogPanel) LogWithStyle(str string, sty *style.Style) {
	if sp.sysLogPanel == nil {
		fmt.Println(str)
		return
	}
	sp.sysLogPanel.Push(str, sty)
}

func (sp *SystemLogPanel) Logf(format string, a ...any) {
	if sp.sysLogPanel == nil {
		if !strings.HasSuffix(format, "\n") {
			format += "\n"
		}
		fmt.Printf(format, a...)
		return
	}
	sp.sysLogPanel.Push(fmt.Sprintf(format, a...))
}

func (sp *SystemLogPanel) Clear() {
	sp.sysLogPanel.Flush()
}
