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
	offset       uint16
	inm          uint16
	inmSize      uint8
	lastInst     string
	lastTag      string
	codes        map[uint16]uint8
	regs         []uint8
	labels       map[string]uint16
	variables    map[string]uint16
	missingTags  map[uint16]string
	directives   map[uint16]*Directive
	linesEndings []uint16
	instructions []string
	inmList      []uint16
	out          []uint8
	max          uint16
	blanks       []uint16
	lastLine     string

	processDone bool
}

func NewAssembler() *Assembler {
	return &Assembler{
		codes:        make(map[uint16]uint8, 32),
		regs:         make([]uint8, 0, 2),
		labels:       make(map[string]uint16),
		missingTags:  make(map[uint16]string),
		variables:    make(map[string]uint16),
		directives:   make(map[uint16]*Directive),
		linesEndings: make([]uint16, 0, 10),
		blanks:       make([]uint16, 0, 10),
		instructions: make([]string, 0, 256),
	}
}

var _ AsmE8Listener = &Assembler{}

func (a *Assembler) process() error {
	if a.processDone {
		return nil
	}
	for o, l := range a.missingTags {
		val, ok := a.labels[l]
		if !ok {
			val, ok = a.variables[l]
			if !ok {
				return fmt.Errorf("error unknown label '%s' at %d", l, o)
			}
		}
		a.codes[o] = uint8(val)
		a.codes[o+1] = uint8(val >> 8)
	}
	max := uint16(0)
	for offset := range a.codes {
		if offset >= max {
			max = offset
		}
	}
	ordered := make([]uint8, max+1)
	for i := uint16(0); i < max+1; i++ {
		b := a.codes[i]
		ordered[i] = b
	}
	a.out = ordered
	a.max = max
	a.processDone = true
	return nil
}

func (a *Assembler) Out() ([]byte, error) {
	a.process()
	return a.out, nil
}

func (a *Assembler) AddOpCode(opCode uint8) {
	a.codes[a.offset] = opCode
	a.offset++
}

// ExitInst_reg_inm_variable implements AsmE8Listener.
func (a *Assembler) ExitInst_reg_inm_variable(c *Inst_reg_inm_variableContext) {
	a.setInm(a.variables[a.lastTag])
	a.ExitInst_reg_inm_core()
}

// ExitInst_reg_inm implements AsmE8Listener.
func (a *Assembler) ExitInst_reg_inm(c *Inst_reg_inmContext) {
	a.ExitInst_reg_inm_core()
}

func (a *Assembler) ExitInst_reg_inm_core() {
	opcode := symbols.GetOpCode(strings.ToLower(a.lastInst), symbols.ADDRESSING_MODE_REG_INM_8)
	a.AddOpCode(opcode)
	a.AddOpCode(a.regs[0])
	a.AddOpCode(uint8(a.inm))
}

// ExitInst_reg_reg implements AsmE8Listener.
func (a *Assembler) ExitInst_reg_reg(c *Inst_reg_regContext) {
	opcode := symbols.GetOpCode(strings.ToLower(a.lastInst), symbols.ADDRESSING_MODE_REG_REG)
	a.AddOpCode(opcode)
	var val uint8 = a.regs[0] | (a.regs[1] << 4)
	a.AddOpCode(val)
}

// ExitInst_single implements AsmE8Listener.
func (a *Assembler) ExitInst_single(c *Inst_singleContext) {
	opcode := symbols.GetOpCode(strings.ToLower(a.lastInst), symbols.ADDRESSING_MODE_IMPL)
	a.AddOpCode(opcode)
}

// ExitInst_single_inm implements AsmE8Listener.
func (a *Assembler) ExitInst_single_inm(c *Inst_single_inmContext) {
	a.ExitInst_single_inm_core()
}

func (a *Assembler) ExitInst_single_inm_core() {
	var mode uint8
	if a.inmSize == 16 {
		mode = symbols.ADDRESSING_MODE_IMPL_INM_16
	} else {
		mode = symbols.ADDRESSING_MODE_REG_INM_8
	}
	opcode := symbols.GetOpCode(strings.ToLower(a.lastInst), mode)
	a.AddOpCode(opcode)
	a.AddOpCode(uint8(a.inm))
	if a.inmSize == 16 {
		a.AddOpCode(uint8(a.inm >> 8))
	}
}

// ExitInst_single_reg implements AsmE8Listener.
func (a *Assembler) ExitInst_single_reg(c *Inst_single_regContext) {
	opcode := symbols.GetOpCode(strings.ToLower(a.lastInst), symbols.ADDRESSING_MODE_IMPL_REG)
	a.AddOpCode(opcode)
	a.AddOpCode(a.regs[0])
}

// ExitInst_single_tag implements AsmE8Listener.
func (a *Assembler) ExitInst_single_tag(c *Inst_single_tagContext) {
	val, exists := a.variables[a.lastTag]
	if exists {
		a.setInm(val)
		a.ExitInst_single_inm_core()
		return
	}

	opcode := symbols.GetOpCode(strings.ToLower(a.lastInst), symbols.ADDRESSING_MODE_IMPL_INM_16)
	a.AddOpCode(opcode)
	val64, exists := a.labels[a.lastTag]
	if !exists {
		a.missingTags[a.offset] = a.lastTag
	}
	a.AddOpCode(uint8(val64))
	a.AddOpCode(uint8(val64 >> 8))
}

// ExitInst_ptr_reg implements AsmE8Listener.
func (a *Assembler) ExitInst_ptr_reg(c *Inst_ptr_regContext) {
	opcode := symbols.GetOpCode(strings.ToLower(a.lastInst), symbols.ADDRESSING_MODE_REG_MEM)
	a.AddOpCode(opcode)
	a.AddOpCode(a.regs[0])
	a.AddOpCode(uint8(a.inm))
	a.AddOpCode(uint8(a.inm >> 8))
}

// ExitInst_reg_ptr implements AsmE8Listener.
func (a *Assembler) ExitInst_reg_ptr(c *Inst_reg_ptrContext) {
	opcode := symbols.GetOpCode(strings.ToLower(a.lastInst), symbols.ADDRESSING_MODE_MEM_REG)
	a.AddOpCode(opcode)
	a.AddOpCode(a.regs[0])
	a.AddOpCode(uint8(a.inm))
	a.AddOpCode(uint8(a.inm >> 8))
}

// ExitInst_reg_ptr_offset implements AsmE8Listener.
func (a *Assembler) ExitInst_reg_ptr_offset(c *Inst_reg_ptr_offsetContext) {
	opcode := symbols.GetOpCode(strings.ToLower(a.lastInst), symbols.ADDRESSING_MODE_MEM_REG_OFFSET)
	a.AddOpCode(opcode)
	var val uint8 = a.regs[1] | (a.regs[0] << 4)
	a.AddOpCode(val)
	if a.lastTag != "" && a.inm == 0 {
		val, ok := a.GetTagValue(a.lastTag)
		if !ok {
			a.missingTags[a.offset] = a.lastTag
		} else {
			a.setInm(val)
		}
	}
	a.AddOpCode(uint8(a.inm))
	a.AddOpCode(uint8(a.inm >> 8))
}

// ExitInst_ptr_offset_reg implements AsmE8Listener.
func (a *Assembler) ExitInst_ptr_offset_reg(c *Inst_ptr_offset_regContext) {
	opcode := symbols.GetOpCode(strings.ToLower(a.lastInst), symbols.ADDRESSING_MODE_REG_MEM_OFFSET)
	a.AddOpCode(opcode)
	var val uint8 = a.regs[0] | (a.regs[1] << 4)
	a.AddOpCode(val)
	if a.lastTag != "" && a.inm == 0 {
		val, ok := a.GetTagValue(a.lastTag)
		if !ok {
			a.missingTags[a.offset] = a.lastTag
		} else {
			a.setInm(val)
		}
	}
	a.AddOpCode(uint8(a.inm))
	a.AddOpCode(uint8(a.inm >> 8))
}

func (a *Assembler) GetTagValue(tag string) (uint16, bool) {
	val, exists := a.variables[tag]
	if exists {
		return val, true
	}
	val, exists = a.labels[tag]
	if exists {
		return val, true
	}
	return 0, false
}

// ExitPtr_offset implements AsmE8Listener.
func (a *Assembler) ExitPtr_offset(c *Ptr_offsetContext) {
	if a.lastTag == "" {
		return
	}
	a.setInm(a.variables[a.lastTag])
}

// ExitVariable implements AsmE8Listener.
func (a *Assembler) ExitVariable(c *VariableContext) {
	a.variables[a.lastTag] = a.inm
}

// ExitLabel implements AsmE8Listener.
func (a *Assembler) ExitLabel(c *LabelContext) {
	a.labels[a.lastTag] = a.offset
}

// ExitMnemonic implements AsmE8Listener.
func (a *Assembler) ExitMnemonic(c *MnemonicContext) {
	a.lastInst = c.GetText()
}

// ExitPtr implements AsmE8Listener.
func (a *Assembler) ExitPtr(c *PtrContext) {
	if a.lastTag == "" {
		return
	}
	a.setInm(a.variables[a.lastTag])
}

// ExitTag implements instListener.
func (a *Assembler) ExitTag(c *TagContext) {
	a.lastTag = c.GetText()
}

// ExitInm implements instListener.
func (a *Assembler) ExitInm(c *InmContext) {
	text := c.GetText()
	val := a.parseInm(text)
	a.setInm(val)
}

// ExitLine implements instListener.
func (a *Assembler) ExitLine(c *LineContext) {
	a.inm = 0
	a.inmSize = 0
	a.lastTag = ""
	a.regs = make([]uint8, 0, 2)

	text := c.GetText()
	if text == "\n" && a.lastLine == "\n" {
		a.blanks = append(a.blanks, a.offset)
	}
	a.lastLine = text
}

// ExitReg implements instListener.
func (a *Assembler) ExitReg(c *RegContext) {
	regOP := symbols.REGISTER_OPCODE[strings.ToLower(c.GetText())]
	a.regs = append(a.regs, regOP)
}

// ExitDirectives implements AsmE8Listener.
func (a *Assembler) ExitDirectives(c *DirectivesContext) {
	text := c.GetText()
	if strings.HasPrefix(text, ".org") {
		offset := a.inmList[0]
		a.directives[a.offset] = &Directive{raw: text, code: ".org", position: a.offset, offset: offset, single: true, inm: a.inmList}
		a.offset = offset
		return
	}
	if strings.HasPrefix(text, ".byte") {
		a.directives[a.offset] = &Directive{raw: text, code: ".byte", position: a.offset, offset: uint16(len(a.inmList)), single: false, inm: a.inmList}
		for _, b := range a.inmList {
			a.AddOpCode(uint8(b))
		}
		return
	}
	if strings.HasPrefix(text, ".word") {
		a.directives[a.offset] = &Directive{raw: text, code: ".word", position: a.offset, offset: uint16(len(a.inmList) * 2), single: false, inm: a.inmList}
		for _, b := range a.inmList {
			a.AddOpCode(uint8(b))
			a.AddOpCode(uint8(b >> 8))
		}
		return
	}
}

// ExitInm_list implements AsmE8Listener.
func (a *Assembler) ExitInm_list(c *Inm_listContext) {
	list := c.GetText()
	tokens := strings.Split(list, ",")
	inmList := make([]uint16, 0, len(tokens))
	for _, tok := range tokens {
		tok = strings.TrimSpace(tok)
		val, ok := a.variables[tok]
		if ok {
			inmList = append(inmList, val)
			continue
		}
		val64 := a.parseInm(tok)
		inmList = append(inmList, uint16(val64))
	}
	a.inmList = inmList
}

func (a *Assembler) parseInm(text string) uint16 {
	var val int64

	if strings.HasPrefix(text, "'") && strings.HasSuffix(text, "'") {
		text = strings.TrimPrefix(text, "'")
		text = strings.TrimSuffix(text, "'")
		val := uint16(text[0])
		return val
	}

	if strings.HasPrefix(text, "\"") && strings.HasSuffix(text, "\"") {
		text = strings.TrimPrefix(text, "\"")
		text = strings.TrimSuffix(text, "\"")
		val := uint16(text[0])
		return val
	}

	if strings.HasPrefix(text, "0x") {
		text = strings.TrimPrefix(text, "0x")
		val, _ = strconv.ParseInt(text, 16, 64)
		return uint16(val)
	}

	if strings.HasPrefix(text, "0b") {
		text = strings.TrimPrefix(text, "0b")
		val, _ = strconv.ParseInt(text, 2, 64)
		return uint16(val)
	}

	val, _ = strconv.ParseInt(text, 10, 64)
	return uint16(val)
}

func (a *Assembler) setInm(val uint16) {
	a.inm = val
	if val > 0xff {
		a.inmSize = 16
	} else {
		a.inmSize = 8
	}
}

// ExitInst implements AsmE8Listener.
func (a *Assembler) ExitInst(c *InstContext) {
	a.linesEndings = append(a.linesEndings, a.offset)
	a.instructions = append(a.instructions, c.GetText())
}

// EnterInst_reg_inm_variable implements AsmE8Listener.
func (a *Assembler) EnterInst_reg_inm_variable(c *Inst_reg_inm_variableContext) {}

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

// ExitInstruction implements AsmE8Listener.
func (a *Assembler) ExitInstruction(c *InstructionContext) {}

// ExitProg implements AsmE8Listener.
func (a *Assembler) ExitProg(c *ProgContext) {}

// VisitErrorNode implements AsmE8Listener.
func (a *Assembler) VisitErrorNode(node antlr.ErrorNode) {}

// VisitTerminal implements AsmE8Listener.
func (a *Assembler) VisitTerminal(node antlr.TerminalNode) {}

// EnterLine implements instListener.
func (a *Assembler) EnterLine(c *LineContext) {}

// EnterDirectives implements AsmE8Listener.
func (a *Assembler) EnterDirectives(c *DirectivesContext) {}

// EnterInm_list implements AsmE8Listener.
func (a *Assembler) EnterInm_list(c *Inm_listContext) {}
