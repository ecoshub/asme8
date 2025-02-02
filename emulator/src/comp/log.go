package comp

import (
	"asme8/common/instruction"
	"asme8/emulator/utils"
	"fmt"
	"strings"

	"github.com/ecoshub/termium/component/style"
)

var (
	DefaultStyle1       = &style.Style{ForegroundColor: 38}
	DefaultStyle2       = &style.Style{ForegroundColor: 204}
	DefaultStyle3       = &style.Style{ForegroundColor: 227}
	DefaultStyle4       = &style.Style{ForegroundColor: 116}
	DefaultStyle5       = &style.Style{ForegroundColor: 35}
	DefaultStyle6       = &style.Style{ForegroundColor: 202}
	DefaultStyle7       = &style.Style{ForegroundColor: 157}
	DefaultHelpStyle    = &style.Style{ForegroundColor: 247}
	DefaultWarningStyle = &style.Style{ForegroundColor: 226}
	DefaultBreakStyle   = &style.Style{ForegroundColor: 162}
)

func (c *Comp) Log(str string) {
	if c.terminal == nil || c.terminal.Components.SysLogPanel == nil {
		fmt.Println(str)
		return
	}
	c.terminal.Components.SysLogPanel.Push(str)
}

func (c *Comp) LogWithStyle(str string, sty *style.Style) {
	if c.terminal == nil || c.terminal.Components.SysLogPanel == nil {
		if c.Config.Debug {
			fmt.Println(str)
		}
		return
	}
	c.terminal.Components.SysLogPanel.Push(str, sty)
}

func (c *Comp) Logf(format string, a ...any) {
	if c.terminal == nil || c.terminal.Components.SysLogPanel == nil {
		if c.Config.Debug {
			if !strings.HasSuffix(format, "\n") {
				format += "\n"
			}
			fmt.Printf(format, a...)
		}
		return
	}
	c.terminal.Components.SysLogPanel.Push(fmt.Sprintf(format, a...))
}

func (c *Comp) LogfFlag(format string, a ...any) {
	if c.terminal == nil || c.terminal.Components.MemoryPanel == nil {
		if c.Config.Debug {
			fmt.Printf(format, a...)
		}
		return
	}
	c.terminal.Components.MemoryPanel.Push(fmt.Sprintf(format, a...))
}

func (c *Comp) LogfFlagIndex(index int, format string, a ...any) {
	if c.terminal == nil || c.terminal.Components.FlagPanel == nil {
		if c.Config.Debug {
			fmt.Printf(format, a...)
		}
		return
	}
	c.terminal.Components.FlagPanel.Write(index, fmt.Sprintf(format, a...))
}

func (c *Comp) LogfFlagIndexWithStyle(index int, sty *style.Style, format string, a ...any) {
	if c.terminal == nil || c.terminal.Components.FlagPanel == nil {
		if c.Config.Debug {
			fmt.Printf(format, a...)
		}
		return
	}
	c.terminal.Components.FlagPanel.Write(index, fmt.Sprintf(format, a...), sty)
}

func (c *Comp) LogState() {
	if c.terminal == nil || c.terminal.Components.FlagPanel == nil {
		c.Logf("# pc: %04x, step: %d, inst_r: %02x, op_r: %02x, mdr: %02x, mar: %04x, addr: %04x, input_bus: %02x, alu_bus: %02x, output_bus: %02x, rw: %x, status: %08b, regs: %s\n", c.programCounter, c.step, c.instructionRegister, c.operandRegister, c.memoryDataRegister, c.memoryAddressRegister, c.addrBus.Read_16(), c.inputBus.Read_16(), c.aluBus.Read_16(), c.outputBus.Read_16(), c.rw, c.status.Flag(), c.registers)
		return
	}
	stepLen := len(CONTROL_SIGNALS[c.instructionRegister])
	visualStep := ((int(c.step) + stepLen - 1) % stepLen) + 1
	// NOTE: can not convertible for multi write do not try it
	c.LogfFlagIndexWithStyle(0, DefaultStyle6, "# Registers:")
	c.LogfFlagIndexWithStyle(1, DefaultStyle6, "Z S P C O I X X")
	c.LogfFlagIndexWithStyle(2, DefaultStyle6, "%s", c.status)
	c.LogfFlagIndexWithStyle(3, DefaultStyle7, "A  B  C  D  IH IL")
	c.LogfFlagIndexWithStyle(4, DefaultStyle7, "%s", c.registers)
	c.LogfFlagIndexWithStyle(5, DefaultStyle1, "IR    : %02x [%s]", c.instructionRegister, instruction.INST_HUMAN_READABLE[c.instructionRegister])
	c.LogfFlagIndexWithStyle(6, DefaultStyle1, "STEP  : %d/%d", visualStep, stepLen)
	c.LogfFlagIndexWithStyle(7, DefaultStyle1, "PC    : %04x", c.programCounter)
	c.LogfFlagIndexWithStyle(8, DefaultStyle1, "SP    : %04x", c.stackPointer)
	c.LogfFlagIndexWithStyle(9, DefaultStyle1, "MAR   : %02x", c.memoryAddressRegister)
	c.LogfFlagIndexWithStyle(10, DefaultStyle1, "MDR   : %02x", c.memoryDataRegister)
	c.LogfFlagIndexWithStyle(11, DefaultStyle1, "OR    : %02x", c.operandRegister)
	c.LogfFlagIndexWithStyle(12, DefaultStyle2, "# Buses:")
	c.LogfFlagIndexWithStyle(13, DefaultStyle2, "ADDR   : %04x", c.addrBus.Read_16())
	c.LogfFlagIndexWithStyle(16, DefaultStyle2, "INPUT  : %02x", c.inputBus.Read_16())
	c.LogfFlagIndexWithStyle(15, DefaultStyle2, "OUTPUT : %02x", c.outputBus.Read_16())
	c.LogfFlagIndexWithStyle(14, DefaultStyle2, "ALU    : %02x", c.aluBus.Read_16())
	c.LogfFlagIndexWithStyle(17, DefaultStyle2, "RW     : %01x", c.rw)
	c.LogfFlagIndexWithStyle(18, DefaultStyle3, "# Bridge:")
	c.LogfFlagIndexWithStyle(19, DefaultStyle3, "ENB : %v", c.bridgeEnable)
}

func (c *Comp) LogMemory() {
	if c.terminal == nil || c.terminal.Components.MemoryPanel == nil {
		return
	}
	buffer := make([]uint8, 0x10000)
	for _, dev := range c.devices {
		start, end := dev.GetRange()
		for i := start; i < end; i++ {
			buffer[i] = dev.Read(i)
		}
	}
	height := c.terminal.Components.MemoryPanel.Config.Height
	lineCount := 0
	logLines := make([]string, 0, height)
	for i := 0; i < 0x10000; i += 8 {
		index := i + int(c.inspectionMemoryOffset)
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

	c.terminal.Components.MemoryPanel.WriteMultiStyle(logLines, DefaultStyle5)
}

func (c *Comp) ListBreakPoints() {
	if len(c.breakPoints) == 0 {
		c.Logf("No breakpoints set")
		return
	}
	c.Logf("Breakpoints:")
	for _, bp := range c.breakPoints {
		c.Logf("● 0x%04x", bp)
	}
}
