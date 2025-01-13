package parser

import (
	"asme8/assembler/src/symbols"
	"fmt"
	"strconv"
	"strings"

	"github.com/antlr4-go/antlr/v4"
)

// Assembler is a complete listener for a parse tree produced by instParser.
type Assembler struct {
	offset        int
	labels        map[string]uint16
	out           []byte
	ops           []uint8
	lastInst      string
	lastTag       string
	missingLabels map[uint16]string
	inmSize       uint8
	inmRaw        uint16
	variables     map[string]uint16
	ptrVariable   bool
}

func NewAssembler() *Assembler {
	return &Assembler{
		labels:        make(map[string]uint16),
		out:           make([]byte, 0, 32),
		ops:           make([]uint8, 0, 2),
		missingLabels: make(map[uint16]string),
		variables:     make(map[string]uint16),
	}
}

var _ AsmE8Listener = &Assembler{}

func (a *Assembler) Out() ([]byte, error) {
	for o, l := range a.missingLabels {
		val, ok := a.labels[l]
		if !ok {
			return nil, fmt.Errorf("error unknown label '%s' at %d", l, o)
		}
		a.out[o] = uint8(val)
		a.out[o+1] = uint8(val >> 8)
	}
	return a.out, nil
}

func (a *Assembler) AddOpCode(opCode uint8) {
	a.out = append(a.out, opCode)
	a.offset++
}

// EnterEveryRule implements AsmE8Listener.
func (a *Assembler) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// EnterInm implements AsmE8Listener.
func (a *Assembler) EnterInm(c *InmContext) {}

// EnterInst implements AsmE8Listener.
func (a *Assembler) EnterInst(c *InstContext) {}

// EnterVariable implements AsmE8Listener.
func (a *Assembler) EnterVariable(c *VariableContext) {}

// EnterInst_reg_inm implements AsmE8Listener.
func (a *Assembler) EnterInst_reg_inm(c *Inst_reg_inmContext) {}

// EnterInst_reg_reg implements AsmE8Listener.
func (a *Assembler) EnterInst_reg_reg(c *Inst_reg_regContext) {}

// EnterInst_single implements AsmE8Listener.
func (a *Assembler) EnterInst_single(c *Inst_singleContext) {}

// EnterInst_single_inm implements AsmE8Listener.
func (a *Assembler) EnterInst_single_inm(c *Inst_single_inmContext) {}

// EnterInst_single_reg implements AsmE8Listener.
func (a *Assembler) EnterInst_single_reg(c *Inst_single_regContext) {}

// EnterInst_single_tag implements AsmE8Listener.
func (a *Assembler) EnterInst_single_tag(c *Inst_single_tagContext) {}

// EnterInst_ptr_reg implements AsmE8Listener.
func (a *Assembler) EnterInst_ptr_reg(c *Inst_ptr_regContext) {}

// EnterInst_reg_ptr implements AsmE8Listener.
func (a *Assembler) EnterInst_reg_ptr(c *Inst_reg_ptrContext) {}

// EnterInst_reg_ptr_offset implements AsmE8Listener.
func (a *Assembler) EnterInst_reg_ptr_offset(c *Inst_reg_ptr_offsetContext) {}

// EnterInst_ptr_offset_reg implements AsmE8Listener.
func (a *Assembler) EnterInst_ptr_offset_reg(c *Inst_ptr_offset_regContext) {}

// EnterPtr_offset implements AsmE8Listener.
func (a *Assembler) EnterPtr_offset(c *Ptr_offsetContext) {}

// EnterInstruction implements AsmE8Listener.
func (a *Assembler) EnterInstruction(c *InstructionContext) {}

// EnterLabel implements AsmE8Listener.
func (a *Assembler) EnterLabel(c *LabelContext) {}

// EnterMnemonic implements AsmE8Listener.
func (a *Assembler) EnterMnemonic(c *MnemonicContext) {}

// EnterProg implements AsmE8Listener.
func (a *Assembler) EnterProg(c *ProgContext) {}

// EnterPtr implements AsmE8Listener.
func (a *Assembler) EnterPtr(c *PtrContext) {}

// EnterReg implements AsmE8Listener.
func (a *Assembler) EnterReg(c *RegContext) {}

// EnterTag implements AsmE8Listener.
func (a *Assembler) EnterTag(c *TagContext) {}

// ExitEveryRule implements AsmE8Listener.
func (a *Assembler) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// ExitInst implements AsmE8Listener.
func (a *Assembler) ExitInst(c *InstContext) {}

// ExitInst_reg_inm implements AsmE8Listener.
func (a *Assembler) ExitInst_reg_inm(c *Inst_reg_inmContext) {
	opcode := symbols.GetOpCode(strings.ToLower(a.lastInst), symbols.ADDRESSING_MODE_INM_8)
	a.AddOpCode(opcode)
	a.AddOpCode(a.ops[0])
	a.AddOpCode(a.ops[1])
}

// ExitInst_reg_reg implements AsmE8Listener.
func (a *Assembler) ExitInst_reg_reg(c *Inst_reg_regContext) {
	opcode := symbols.GetOpCode(strings.ToLower(a.lastInst), symbols.ADDRESSING_MODE_RR_8)
	a.AddOpCode(opcode)
	var val uint8 = a.ops[0] | (a.ops[1] << 4)
	a.AddOpCode(val)
}

// ExitInst_single implements AsmE8Listener.
func (a *Assembler) ExitInst_single(c *Inst_singleContext) {
	opcode := symbols.GetOpCode(strings.ToLower(a.lastInst), symbols.ADDRESSING_MODE_IMPL)
	a.AddOpCode(opcode)
}

// ExitInst_single_inm implements AsmE8Listener.
func (a *Assembler) ExitInst_single_inm(c *Inst_single_inmContext) {
	var mode uint8
	if a.inmSize == 16 {
		mode = symbols.ADDRESSING_MODE_INM_16
	} else {
		mode = symbols.ADDRESSING_MODE_INM_8
	}
	opcode := symbols.GetOpCode(strings.ToLower(a.lastInst), mode)
	a.AddOpCode(opcode)
	a.AddOpCode(a.ops[0])
	if a.inmSize == 16 {
		a.AddOpCode(a.ops[1])
	}
}

// ExitInst_single_reg implements AsmE8Listener.
func (a *Assembler) ExitInst_single_reg(c *Inst_single_regContext) {
	opcode := symbols.GetOpCode(strings.ToLower(a.lastInst), symbols.ADDRESSING_MODE_IMPL_REG)
	a.AddOpCode(opcode)
	a.AddOpCode(a.ops[0])
}

// ExitInst_single_tag implements AsmE8Listener.
func (a *Assembler) ExitInst_single_tag(c *Inst_single_tagContext) {
	opcode := symbols.GetOpCode(strings.ToLower(a.lastInst), symbols.ADDRESSING_MODE_INM_16)
	val, exists := a.labels[a.lastTag]
	if !exists {
		a.missingLabels[uint16(a.offset+1)] = a.lastTag
	}
	a.AddOpCode(opcode)
	a.AddOpCode(uint8(val))
	a.AddOpCode(uint8(val >> 8))
}

// ExitInst_ptr_reg implements AsmE8Listener.
func (a *Assembler) ExitInst_ptr_reg(c *Inst_ptr_regContext) {
	opcode := symbols.GetOpCode(strings.ToLower(a.lastInst), symbols.ADDRESSING_MODE_REG_MEM)
	a.AddOpCode(opcode)
	a.AddOpCode(a.ops[2])
	a.AddOpCode(a.ops[0])
	a.AddOpCode(a.ops[1])
}

// ExitInst_reg_ptr implements AsmE8Listener.
func (a *Assembler) ExitInst_reg_ptr(c *Inst_reg_ptrContext) {
	opcode := symbols.GetOpCode(strings.ToLower(a.lastInst), symbols.ADDRESSING_MODE_MEM_REG)
	a.AddOpCode(opcode)
	a.AddOpCode(a.ops[0])
	a.AddOpCode(a.ops[1])
	a.AddOpCode(a.ops[2])
}

// ExitInst_reg_ptr_offset implements AsmE8Listener.
func (a *Assembler) ExitInst_reg_ptr_offset(c *Inst_reg_ptr_offsetContext) {
	opcode := symbols.GetOpCode(strings.ToLower(a.lastInst), symbols.ADDRESSING_MODE_MEM_REG_OFFSET)
	a.AddOpCode(opcode)
	if !a.ptrVariable {
		var val uint8 = a.ops[3] | (a.ops[0] << 4)
		a.AddOpCode(val)
		a.AddOpCode(a.ops[1])
		a.AddOpCode(a.ops[2])
	} else {
		var val uint8 = a.ops[1] | (a.ops[0] << 4)
		a.AddOpCode(val)
		a.AddOpCode(a.ops[2])
		a.AddOpCode(a.ops[3])
	}
}

// ExitInst_ptr_offset_reg implements AsmE8Listener.
func (a *Assembler) ExitInst_ptr_offset_reg(c *Inst_ptr_offset_regContext) {
	opcode := symbols.GetOpCode(strings.ToLower(a.lastInst), symbols.ADDRESSING_MODE_REG_MEM_OFFSET)
	a.AddOpCode(opcode)
	if !a.ptrVariable {
		var val uint8 = a.ops[2] | (a.ops[3] << 4)
		a.AddOpCode(val)
		a.AddOpCode(a.ops[0])
		a.AddOpCode(a.ops[1])
	} else {
		var val uint8 = a.ops[0] | (a.ops[1] << 4)
		a.AddOpCode(val)
		a.AddOpCode(a.ops[3])
		a.AddOpCode(a.ops[2])
	}
}

// ExitPtr_offset implements AsmE8Listener.
func (a *Assembler) ExitPtr_offset(c *Ptr_offsetContext) {
	if a.lastTag == "" {
		return
	}
	a.setInm(int64(a.variables[a.lastTag]))
	a.ptrVariable = true
}

// ExitInstruction implements AsmE8Listener.
func (a *Assembler) ExitInstruction(c *InstructionContext) {}

// ExitVariable implements AsmE8Listener.
func (a *Assembler) ExitVariable(c *VariableContext) {
	a.variables[a.lastTag] = a.inmRaw
}

// ExitLabel implements AsmE8Listener.
func (a *Assembler) ExitLabel(c *LabelContext) {
	a.labels[a.lastTag] = uint16(a.offset)
}

// ExitMnemonic implements AsmE8Listener.
func (a *Assembler) ExitMnemonic(c *MnemonicContext) {
	a.lastInst = c.GetText()
}

// ExitProg implements AsmE8Listener.
func (a *Assembler) ExitProg(c *ProgContext) {}

// ExitPtr implements AsmE8Listener.
func (a *Assembler) ExitPtr(c *PtrContext) {
	if a.lastTag == "" {
		return
	}
	a.setInm(int64(a.variables[a.lastTag]))
	a.ptrVariable = true
}

// VisitErrorNode implements AsmE8Listener.
func (a *Assembler) VisitErrorNode(node antlr.ErrorNode) {}

// VisitTerminal implements AsmE8Listener.
func (a *Assembler) VisitTerminal(node antlr.TerminalNode) {}

// ExitTag implements instListener.
func (a *Assembler) ExitTag(c *TagContext) {
	a.lastTag = c.GetText()
}

// EnterLine implements instListener.
func (a *Assembler) EnterLine(c *LineContext) {
	a.ops = make([]uint8, 0, 2)
}

// ExitInm implements instListener.
func (a *Assembler) ExitInm(c *InmContext) {
	text := c.GetText()
	a.parseInm(text)
}

// ExitLine implements instListener.
func (a *Assembler) ExitLine(c *LineContext) {
	a.inmRaw = 0
	a.inmSize = 0
	a.lastTag = ""
	a.ptrVariable = false
}

// ExitReg implements instListener.
func (a *Assembler) ExitReg(c *RegContext) {
	regOP := symbols.REGISTER_OPCODE[strings.ToLower(c.GetText())]
	a.ops = append(a.ops, regOP)
}

func (a *Assembler) parseInm(text string) {
	var val int64
	if strings.HasPrefix(text, "0x") {
		text = strings.TrimPrefix(text, "0x")
		val, _ = strconv.ParseInt(text, 16, 64)
	} else {
		val, _ = strconv.ParseInt(text, 10, 64)
	}
	a.setInm(val)
}

func (a *Assembler) setInm(val int64) {
	a.inmRaw = uint16(val)
	if val > 0xff {
		a.inmSize = 16
	} else {
		a.inmSize = 8
	}
	a.ops = append(a.ops, uint8(val))
	a.ops = append(a.ops, uint8(val>>8))
}
