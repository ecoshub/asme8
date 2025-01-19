package comp

import (
	"asme8/emulator/src/instruction"
	"asme8/emulator/utils"
	"fmt"

	"github.com/ecoshub/termium/component/style"
)

var (
	DefaultStyle1       = &style.Style{ForegroundColor: 38}
	DefaultStyle2       = &style.Style{ForegroundColor: 204}
	DefaultStyle3       = &style.Style{ForegroundColor: 227}
	DefaultStyle4       = &style.Style{ForegroundColor: 116}
	DefaultStyle5       = &style.Style{ForegroundColor: 35}
	DefaultHelpStyle    = &style.Style{ForegroundColor: 247}
	DefaultWarningStyle = &style.Style{ForegroundColor: 226}
	DefaultBreakStyle   = &style.Style{ForegroundColor: 162}
)

func (c *Comp) Log(str string) {
	if c.terminalComponents == nil || c.terminalComponents.SysLogPanel == nil {
		fmt.Println(str)
		return
	}
	c.terminalComponents.SysLogPanel.Push(str)
}

func (c *Comp) LogWithStyle(str string, sty *style.Style) {
	if c.terminalComponents == nil || c.terminalComponents.SysLogPanel == nil {
		if c.debug {
			fmt.Println(str)
		}
		return
	}
	c.terminalComponents.SysLogPanel.Push(str, sty)
}

func (c *Comp) Logf(format string, a ...any) {
	if c.terminalComponents == nil || c.terminalComponents.SysLogPanel == nil {
		if c.debug {
			fmt.Printf(format, a...)
		}
		return
	}
	c.terminalComponents.SysLogPanel.Push(fmt.Sprintf(format, a...))
}

func (c *Comp) LogfFlag(format string, a ...any) {
	if c.terminalComponents == nil || c.terminalComponents.MemoryPanel == nil {
		if c.debug {
			fmt.Printf(format, a...)
		}
		return
	}
	c.terminalComponents.MemoryPanel.Push(fmt.Sprintf(format, a...))
}

func (c *Comp) LogfFlagIndex(index int, format string, a ...any) {
	if c.terminalComponents == nil || c.terminalComponents.FlagPanel == nil {
		if c.debug {
			fmt.Printf(format, a...)
		}
		return
	}
	c.terminalComponents.FlagPanel.Write(index, fmt.Sprintf(format, a...))
}

func (c *Comp) LogfFlagIndexWithStyle(index int, sty *style.Style, format string, a ...any) {
	if c.terminalComponents == nil || c.terminalComponents.FlagPanel == nil {
		if c.debug {
			fmt.Printf(format, a...)
		}
		return
	}
	c.terminalComponents.FlagPanel.Write(index, fmt.Sprintf(format, a...), sty)
}

func (c *Comp) LogState() {
	if c.terminalComponents == nil || c.terminalComponents.FlagPanel == nil {
		c.Logf("# pc: %04x, step: %d, inst_r: %02x, op_r: %02x, addr: %04x, data: %02x, bus_x: %02x, bus_y: %02x, rw: %x, status: %08b, regs: %s\n", c.programCounter, c.step, c.instructionRegister, c.operandRegister, c.addrBus.Read_16(), c.dataBus.Read_16(), c.busX.Read_16(), c.busY.Read_16(), c.rw, c.status.Flag(), c.registers)
		return
	}
	stepLen := len(CONTROL_SIGNALS[c.instructionRegister])
	visualStep := ((int(c.step) + stepLen - 1) % stepLen) + 1
	// NOTE: can not convertible for multi write do not try it
	c.LogfFlagIndexWithStyle(0, DefaultStyle1, "# REGISTER:")
	c.LogfFlagIndexWithStyle(1, DefaultStyle1, "flags     : Z S P C O I X X")
	c.LogfFlagIndexWithStyle(2, DefaultStyle1, "flag vals : %s", c.status)
	c.LogfFlagIndexWithStyle(3, DefaultStyle1, "reg       :  A  B  C  D")
	c.LogfFlagIndexWithStyle(4, DefaultStyle1, "reg vals  : %s", c.registers)
	c.LogfFlagIndexWithStyle(5, DefaultStyle1, "STEP  : %d/%d", visualStep, stepLen)
	c.LogfFlagIndexWithStyle(6, DefaultStyle1, "IR    : %02x [%s]", c.instructionRegister, instruction.INST_MNEMONICS[c.instructionRegister])
	c.LogfFlagIndexWithStyle(7, DefaultStyle1, "PC    : %04x", c.programCounter)
	c.LogfFlagIndexWithStyle(8, DefaultStyle1, "SP    : %04x", c.stackPointer)
	c.LogfFlagIndexWithStyle(9, DefaultStyle1, "MAR   : %02x", c.memoryAddressRegister)
	c.LogfFlagIndexWithStyle(10, DefaultStyle1, "OR    : %02x", c.operandRegister)
	c.LogfFlagIndexWithStyle(11, DefaultStyle2, "# BUS:")
	c.LogfFlagIndexWithStyle(12, DefaultStyle2, "ADDR : %04x", c.addrBus.Read_16())
	c.LogfFlagIndexWithStyle(13, DefaultStyle2, "DATA : %02x", c.dataBus.Read_16())
	c.LogfFlagIndexWithStyle(14, DefaultStyle2, "RW   : %02x", c.rw)
	c.LogfFlagIndexWithStyle(15, DefaultStyle2, "X    : %02x", c.busX.Read_16())
	c.LogfFlagIndexWithStyle(16, DefaultStyle2, "Y    : %02x", c.busY.Read_16())
	c.LogfFlagIndexWithStyle(17, DefaultStyle3, "# BRIDGE:")
	c.LogfFlagIndexWithStyle(18, DefaultStyle3, "DIR : %02x", c.bridgeDir)
	c.LogfFlagIndexWithStyle(19, DefaultStyle3, "ENB : %v", c.bridgeEnable)
	c.LogfFlagIndexWithStyle(20, DefaultStyle4, "# OTHER:")
	c.LogfFlagIndexWithStyle(21, DefaultStyle4, "STR : %02x", c.store)
}

func (c *Comp) LogMemory() {
	if c.terminalComponents == nil || c.terminalComponents.MemoryPanel == nil {
		return
	}
	buffer := make([]uint8, 0x10000)
	for _, dev := range c.devices {
		start, end := dev.GetRange()
		for i := start; i < end; i++ {
			buffer[i] = dev.Read(i)
		}
	}
	height := c.terminalComponents.MemoryPanel.Config.Height
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
			logLines = append(logLines, fmt.Sprintf(" %04x: %s", index, utils.ToHexArray(buffer[index:end], false)))
		}
		lineCount++
		if height <= lineCount {
			break
		}
	}

	c.terminalComponents.MemoryPanel.WriteMultiStyle(logLines, DefaultStyle5)
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
