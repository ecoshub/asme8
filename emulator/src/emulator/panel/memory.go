package panel

import (
	"asme8/emulator/src/connectable"
	"asme8/emulator/utils"
	"fmt"

	"github.com/ecoshub/termium/component/panel"
)

type MemoryPanel struct {
	enabled     bool
	memoryPanel *panel.Stack
	devices     []connectable.Connectable
	offset      int
}

func NewMemoryPanel(memoryPanel *panel.Stack, devices []connectable.Connectable) *MemoryPanel {
	return &MemoryPanel{
		enabled:     memoryPanel != nil,
		memoryPanel: memoryPanel,
		devices:     devices,
	}
}

func (mp *MemoryPanel) IncOffset(value int) {
	mp.offset += value
	mp.Render()
}

func (mp *MemoryPanel) SetOffset(value int) {
	mp.offset = value
	mp.Render()
}

func (mp *MemoryPanel) Render() {
	if !mp.enabled {
		return
	}
	buffer := make([]uint8, 0x10000)
	for _, dev := range mp.devices {
		start, end := dev.GetRange()
		for i := start; i < end; i++ {
			buffer[i] = dev.Read(i)
		}
	}
	height := mp.memoryPanel.Config.Height
	lineCount := 0
	logLines := make([]string, 0, height)
	for i := 0; i < 0x10000; i += 8 {
		index := i + mp.offset
		if index > 0xffff {
			logLines = append(logLines, "")
		} else {
			end := (index + 8)
			if index+8 > 0xffff {
				end = 0xffff + 1
			}
			line := utils.ToHexArray(buffer[index:end], false)
			logLines = append(logLines, fmt.Sprintf(" %04x: %s", index, line))
		}
		lineCount++
		if height <= lineCount {
			break
		}
	}

	mp.memoryPanel.WriteMultiStyle(logLines, DefaultStyle5)
}
