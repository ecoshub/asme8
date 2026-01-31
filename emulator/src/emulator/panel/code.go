package panel

import (
	"asme8/emulator/src/emulator/breakpoint"
	"strings"

	"github.com/ecoshub/termium/component/panel"
	"github.com/ecoshub/termium/component/style"
)

type CodePanel struct {
	enable          bool
	codeLines       map[uint16]string
	codePanel       *panel.Base
	codeRulerPanel  *panel.Raw
	bpm             *breakpoint.Manager
	lastRenderedPc  uint16
	codeOffsetMap   map[uint16]uint16
	lineToOffsetMap map[int]uint16
	skipRequest     int
	skip            int
	trackExecution  bool
}

func NewCodePanel(codeLines map[uint16]string, bpm *breakpoint.Manager, codePanel *panel.Base, codeRulerPanel *panel.Raw) *CodePanel {
	return &CodePanel{
		enable:          true,
		codeLines:       codeLines,
		codePanel:       codePanel,
		codeRulerPanel:  codeRulerPanel,
		bpm:             bpm,
		codeOffsetMap:   createCodeOffsetMap(codeLines),
		lineToOffsetMap: createLineToCodeOffsetMap(codeLines),
		trackExecution:  true,
	}
}

func (cp *CodePanel) SetTrackExecution(value bool) {
	cp.trackExecution = value
}

func (cp *CodePanel) GetTrackExecution() bool {
	return cp.trackExecution
}

func (cp *CodePanel) ToggleTrackExecution() bool {
	cp.trackExecution = !cp.trackExecution
	return cp.trackExecution
}

func (cp *CodePanel) SetSkipRequest(value int) {
	cp.skipRequest = value
}

func (cp *CodePanel) SetSkipTop() {
	cp.skip = 0
	cp.skipRequest = 0
}

func (cp *CodePanel) SetSkipBot() {
	totalPage := len(cp.codeLines) / cp.codePanel.Config.Height
	cp.skip = totalPage
	cp.skipRequest = 0
}

func (cp *CodePanel) GetSkip() int {
	return cp.skip
}

func (cp *CodePanel) Clear() {
	cp.codePanel.Clear()
}

func (cp *CodePanel) Render(pc uint16, renderPC bool) {
	if !cp.enable {
		return
	}

	if cp.trackExecution {
		cp.trackPage(pc)
		return
	}
	cp.staticPage(pc)
}

func (cp *CodePanel) renderCodePanel(pc uint16) {
	height := cp.codePanel.Config.Height
	startPage := (height * cp.skip)
	start := int(cp.lineToOffsetMap[startPage])
	// get lines and breakpoints
	lines, breakpoints := cp.processLines(start, height)

	// render standard lines
	cp.renderStandartLines(lines)

	// clear debug ruler
	cp.codeRulerPanel.Clear()

	// render breakpoints
	cp.renderBreakpoints(breakpoints)

	pcShown, lineIndexOfPc := cp.isExecutedLineShown(start, height, pc)
	if pcShown {
		// render current executed line
		cp.renderExecutedLine(lineIndexOfPc, pc)
	}
}

func (cp *CodePanel) trackPage(pc uint16) {
	page, ok := cp.findPage(pc)
	if !ok {
		return
	}

	cp.skip = page
	cp.renderCodePanel(pc)
}

func (cp *CodePanel) staticPage(pc uint16) {
	cp.processSkipRequest()
	cp.renderCodePanel(pc)
}

func (cp *CodePanel) findPage(pc uint16) (int, bool) {
	// panel height
	height := cp.codePanel.Config.Height

	lineCount := 0
	for i := 0; i < 0x10000; i++ {
		offset := uint16(i)
		_, exists := cp.codeLines[offset]
		if !exists {
			continue
		}
		if offset == cp.codeOffsetMap[pc] {
			return lineCount / height, true
		}
		lineCount++
	}
	return 0, false
}

func (cp *CodePanel) processSkipRequest() {
	totalPage := len(cp.codeLines) / cp.codePanel.Config.Height
	if cp.skipRequest == 0 {
		return
	}
	if cp.skip+cp.skipRequest < 0 {
		cp.skip = 0
	} else if cp.skip+cp.skipRequest > totalPage {
		cp.skip = totalPage
	} else {
		cp.skip = cp.skip + cp.skipRequest
	}
	cp.skipRequest = 0
}

func (cp *CodePanel) processLines(start, height int) ([]string, map[int]string) {
	lines := make([]string, height)
	breakPoints := make(map[int]string, 2)
	lineCount := 0
	for i := start; i < 0x10000; i++ {
		offset := uint16(i)
		line, exists := cp.codeLines[offset]
		if !exists {
			continue
		}
		if lineCount+1 > height {
			break
		}
		if cp.bpm.IsBreakPoint(offset) {
			breakPoints[lineCount] = line
		}
		lines[lineCount] = line
		lineCount++
	}
	return lines, breakPoints
}

func (cp *CodePanel) isExecutedLineShown(start, height int, pc uint16) (bool, int) {
	lineCount := 0
	for i := start; i < 0x10000; i++ {
		offset := uint16(i)
		_, exists := cp.codeLines[offset]
		if !exists {
			continue
		}
		if lineCount+1 > height {
			break
		}
		if offset == cp.codeOffsetMap[pc] {
			return true, lineCount
		}
		lineCount++
	}
	return false, 0
}

func (cp *CodePanel) renderStandartLines(lines []string) {
	for i, line := range lines {
		if strings.Contains(line, ":") {
			cp.codePanel.Write(i, line, DefaultStyle4)
			continue
		}
		cp.codePanel.Write(i, line, CodeStyle)
	}
}

func (cp *CodePanel) renderBreakpoints(breakPoints map[int]string) {
	for index, line := range breakPoints {
		cp.codePanel.Write(index, line, BreakStyle)
		cp.codeRulerPanel.Write(index, '●', BreakStyle)
	}
}

func (cp *CodePanel) renderExecutedLine(lineIndexOfPc int, pc uint16) {
	cp.codePanel.Write(lineIndexOfPc, cp.codeLines[cp.codeOffsetMap[pc]], HighlightedCodeStyle)
	cp.lastRenderedPc = cp.codeOffsetMap[pc]
}

func (cp *CodePanel) WriteNoCode() {
	cp.codePanel.Write(0, "No symbols found", style.DefaultStyleEvent)
}

func (cp *CodePanel) UseCodeOffsetMap(point uint16) uint16 {
	return cp.codeOffsetMap[point]
}

// creates a map for point all micro instructions to start of instruction
func createCodeOffsetMap(codeLines map[uint16]string) map[uint16]uint16 {
	offsetMap := make(map[uint16]uint16, 0xffff)
	lastOffset := uint16(0)
	for i := 0; i < 0x10000; i++ {
		offset := uint16(i)
		_, exists := codeLines[offset]
		if !exists {
			offsetMap[offset] = lastOffset
		} else {
			offsetMap[offset] = offset
			lastOffset = offset
		}
	}
	return offsetMap
}

func createLineToCodeOffsetMap(codeLines map[uint16]string) map[int]uint16 {
	lineToOffsetMap := make(map[int]uint16, len(codeLines))
	lineNumber := 0
	for i := 0; i < 0x10000; i++ {
		offset := uint16(i)
		_, exists := codeLines[offset]
		if !exists {
			continue
		}
		lineToOffsetMap[lineNumber] = offset
		lineNumber++
	}
	return lineToOffsetMap
}
