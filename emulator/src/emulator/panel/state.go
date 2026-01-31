package panel

import (
	"asme8/common/instruction"
	"asme8/emulator/src/computer"
	"fmt"

	"github.com/ecoshub/termium/component/panel"
	"github.com/ecoshub/termium/component/style"
)

type StatePanel struct {
	computer   *computer.Computer
	statePanel *panel.Stack
}

func NewStatePanel(computer *computer.Computer, statePanel *panel.Stack) *StatePanel {
	return &StatePanel{
		computer:   computer,
		statePanel: statePanel,
	}
}

func (sp *StatePanel) Render() {

	ir := sp.computer.GetInstructionRegister()
	pc := sp.computer.GetProgramCounter()
	step := sp.computer.GetStep()
	registers := sp.computer.GetRegisters()
	status := sp.computer.GetStatusRegister()
	operandRegister := sp.computer.GetOperandRegister()
	memoryDataRegister := sp.computer.GetMemoryDataRegister()
	memoryAddressRegister := sp.computer.GetMemoryAddressRegister()
	addrBus := sp.computer.GetAddrBus().Read_16()
	inputBus := sp.computer.GetInputBus().Read_16()
	aluBus := sp.computer.GetAluBus().Read_16()
	outputBus := sp.computer.GetOutputBus().Read_16()
	rw := sp.computer.GetRw()
	stackPointer := sp.computer.GetStackPointer()
	bridgeEnable := sp.computer.GetBridgeEnable()

	stepLen := len(computer.CONTROL_SIGNALS[ir])
	visualStep := stepLen
	if stepLen != 0 {
		visualStep = ((int(step) + stepLen - 1) % stepLen) + 1
	}

	registerStatus := ""
	registerStatus += fmt.Sprintf("%02x ", registers[instruction.REGISTER_OPCODE_A])
	registerStatus += fmt.Sprintf("%02x ", registers[instruction.REGISTER_OPCODE_B])
	registerStatus += fmt.Sprintf("%02x ", registers[instruction.REGISTER_OPCODE_C])
	registerStatus += fmt.Sprintf("%02x ", registers[instruction.REGISTER_OPCODE_D])
	registerStatus += fmt.Sprintf("%02x ", registers[instruction.REGISTER_OPCODE_E])
	registerStatus += fmt.Sprintf("%02x ", registers[instruction.REGISTER_OPCODE_IPH])
	registerStatus += fmt.Sprintf("%02x ", registers[instruction.REGISTER_OPCODE_IPL])
	registerStatus += fmt.Sprintf("%02x ", registers[instruction.REGISTER_OPCODE_BPH])
	registerStatus += fmt.Sprintf("%02x ", registers[instruction.REGISTER_OPCODE_BPL])

	if sp.statePanel == nil {
		be := 0
		if bridgeEnable {
			be = 1
		}
		fmt.Printf("# %04x    %01d       %02x   %02x   %02x   %04x  %04x    %02x          %02x    %02x       %d   %x        %08b        %s\n", pc, step, ir, operandRegister, memoryDataRegister, memoryAddressRegister, addrBus, inputBus, aluBus, outputBus, rw, be, status.Flag(), registerStatus)
		return
	}

	// NOTE: can not try to convert for 'multi write'. do not try it please...
	sp.logfFlagIndexWithStyle(0, DefaultStyle6, "# Registers:")
	sp.logfFlagIndexWithStyle(1, DefaultStyle6, "Z S P C O I X X")
	sp.logfFlagIndexWithStyle(2, DefaultStyle6, "%s", status)
	sp.logfFlagIndexWithStyle(3, DefaultStyle6, "")
	sp.logfFlagIndexWithStyle(4, DefaultStyle7, "A  B  C  D  E  IH IL BH BL")
	sp.logfFlagIndexWithStyle(5, DefaultStyle7, "%s", registerStatus)
	sp.logfFlagIndexWithStyle(6, DefaultStyle7, "")
	sp.logfFlagIndexWithStyle(7, DefaultStyle1, "IR   : %02x [%s]", ir, instruction.INST_HUMAN_READABLE[ir])
	sp.logfFlagIndexWithStyle(8, DefaultStyle1, "STEP : %d/%d", visualStep, stepLen)
	sp.logfFlagIndexWithStyle(9, DefaultStyle1, "PC   : %04x", pc)
	sp.logfFlagIndexWithStyle(10, DefaultStyle1, "SP   : %04x", stackPointer)
	sp.logfFlagIndexWithStyle(11, DefaultStyle1, "MAR  : %02x", memoryAddressRegister)
	sp.logfFlagIndexWithStyle(12, DefaultStyle1, "MDR  : %02x", memoryDataRegister)
	sp.logfFlagIndexWithStyle(13, DefaultStyle1, "OR   : %02x", operandRegister)
	sp.logfFlagIndexWithStyle(14, DefaultStyle1, "")
	sp.logfFlagIndexWithStyle(15, DefaultStyle2, "# Buses:")
	sp.logfFlagIndexWithStyle(16, DefaultStyle2, "ADDR   : %04x", addrBus)
	sp.logfFlagIndexWithStyle(17, DefaultStyle2, "INPUT  : %02x", inputBus)
	sp.logfFlagIndexWithStyle(18, DefaultStyle2, "OUTPUT : %02x", outputBus)
	sp.logfFlagIndexWithStyle(19, DefaultStyle2, "ALU    : %02x", aluBus)
	sp.logfFlagIndexWithStyle(20, DefaultStyle2, "RW     : %01x", rw)
	sp.logfFlagIndexWithStyle(21, DefaultStyle3, "")
	sp.logfFlagIndexWithStyle(22, DefaultStyle3, "# Bridge:")
	sp.logfFlagIndexWithStyle(23, DefaultStyle3, "ENB : %v", bridgeEnable)
}

func (sp *StatePanel) logfFlagIndexWithStyle(index int, sty *style.Style, format string, a ...any) {
	if sp.statePanel == nil {
		fmt.Printf(format, a...)
		return
	}
	sp.statePanel.Write(index, fmt.Sprintf(format, a...), sty)
}

func (sp *StatePanel) PrintStateHeader() {
	fmt.Printf("# pc      step    ir   or   mdr  mar   addr  input_bus  alu_bus  out_bus  rw  be flags(xxIOCPSZ)  regs(A  B  C  D  E  IH IL BH BL)\n")
}
