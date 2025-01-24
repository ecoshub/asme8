package assembler

import (
	"asme8/assembler/src/parser"

	"github.com/antlr4-go/antlr/v4"
)

var _ parser.AsmE8Listener = &Assembler{}

func (a *Assembler) ExitDirectives(c *parser.DirectivesContext) {
	text := c.GetText()
	a.ParseDirective(text)
}

func (a *Assembler) ExitImm(c *parser.ImmContext) {
	text := c.GetText()
	a.ParseValue(text)
}

func (a *Assembler) ExitImm_list(c *parser.Imm_listContext) {
	text := c.GetText()
	a.ParseValueList(text)
}

func (a *Assembler) ExitInst_ptr_offset_reg(c *parser.Inst_ptr_offset_regContext) {
	a.ParsePointerRegisterOffset()
}

func (a *Assembler) ExitInst_ptr_reg(c *parser.Inst_ptr_regContext) {
	a.ParsePointerRegister()
}

func (a *Assembler) ExitInst_reg_imm(c *parser.Inst_reg_immContext) {
	a.ParseRegisterImmediate()
}

func (a *Assembler) ExitInst_reg_imm_variable(c *parser.Inst_reg_imm_variableContext) {
	a.ParseRegisterImmediateVariable()
}

func (a *Assembler) ExitInst_reg_ptr(c *parser.Inst_reg_ptrContext) {
	a.ParseRegisterPointer()
}

func (a *Assembler) ExitInst_reg_ptr_offset(c *parser.Inst_reg_ptr_offsetContext) {
	a.ParseRegisterPointerOffset()
}

func (a *Assembler) ExitInst_reg_reg(c *parser.Inst_reg_regContext) {
	a.ParseRegisterRegister()
}

func (a *Assembler) ExitInst_single(c *parser.Inst_singleContext) {
	a.ParseImplied()
}

func (a *Assembler) ExitInst_single_imm(c *parser.Inst_single_immContext) {
	a.ParseImpliedImmediate()
}

func (a *Assembler) ExitInst_single_reg(c *parser.Inst_single_regContext) {
	a.ParseImpliedRegister()
}

func (a *Assembler) ExitInst_single_tag(c *parser.Inst_single_tagContext) {
	a.ParseImpliedTag()
}

func (a *Assembler) ExitLabel(c *parser.LabelContext) {
	line := c.GetStart().GetLine()
	column := c.RuleIndex
	a.ParseLabel(a.currentTag.Text, line, column, false)
}

func (a *Assembler) ExitLine(c *parser.LineContext) {
	text := c.GetText()
	a.ResetCurrentValues()
	a.CodeParseLastLine(text)
}

func (a *Assembler) ExitMnemonic(c *parser.MnemonicContext) {
	a.currentInstruction = c.GetText()
}

func (a *Assembler) ExitPtr(c *parser.PtrContext) {
	text := c.GetText()
	line := c.GetStart().GetLine()
	column := c.RuleIndex
	a.ParsePtr(text, line, column)
}

func (a *Assembler) ExitPtr_virtual_offset(c *parser.Ptr_virtual_offsetContext) {
	a.hasVirtualOffset = true
}

func (a *Assembler) ExitReg(c *parser.RegContext) {
	text := c.GetText()
	a.ParseRegister(text)
}

func (a *Assembler) ExitTag(c *parser.TagContext) {
	text := c.GetText()
	line := c.GetStart().GetLine()
	column := c.RuleIndex
	a.ParseTag(text, line, column)
}

func (a *Assembler) ExitVariable(c *parser.VariableContext) {
	a.ParseVariable()
}

func (a *Assembler) ExitInst(c *parser.InstContext) {
	text := c.GetText()
	a.CodeParseExitInst(text)
}

func (a *Assembler) VisitErrorNode(node antlr.ErrorNode) {}

func (a *Assembler) VisitTerminal(node antlr.TerminalNode) {}

func (a *Assembler) EnterDirectives(c *parser.DirectivesContext) {}

func (a *Assembler) EnterEveryRule(ctx antlr.ParserRuleContext) {}

func (a *Assembler) EnterImm(c *parser.ImmContext) {}

func (a *Assembler) EnterImm_list(c *parser.Imm_listContext) {}

func (a *Assembler) EnterInst(c *parser.InstContext) {}

func (a *Assembler) EnterInst_ptr_offset_reg(c *parser.Inst_ptr_offset_regContext) {}

func (a *Assembler) EnterInst_ptr_reg(c *parser.Inst_ptr_regContext) {}

func (a *Assembler) EnterInst_reg_imm(c *parser.Inst_reg_immContext) {}

func (a *Assembler) EnterInst_reg_imm_variable(c *parser.Inst_reg_imm_variableContext) {}

func (a *Assembler) EnterInst_reg_ptr(c *parser.Inst_reg_ptrContext) {}

func (a *Assembler) EnterInst_reg_ptr_offset(c *parser.Inst_reg_ptr_offsetContext) {}

func (a *Assembler) EnterInst_reg_reg(c *parser.Inst_reg_regContext) {}

func (a *Assembler) EnterInst_single(c *parser.Inst_singleContext) {}

func (a *Assembler) EnterInst_single_imm(c *parser.Inst_single_immContext) {}

func (a *Assembler) EnterInst_single_reg(c *parser.Inst_single_regContext) {}

func (a *Assembler) EnterInst_single_tag(c *parser.Inst_single_tagContext) {}

func (a *Assembler) EnterInstruction(c *parser.InstructionContext) {}

func (a *Assembler) EnterLabel(c *parser.LabelContext) {}

func (a *Assembler) EnterLine(c *parser.LineContext) {}

func (a *Assembler) EnterMnemonic(c *parser.MnemonicContext) {}

func (a *Assembler) EnterProg(c *parser.ProgContext) {}

func (a *Assembler) EnterPtr(c *parser.PtrContext) {}

func (a *Assembler) EnterPtr_offset(c *parser.Ptr_offsetContext) {}

func (a *Assembler) EnterReg(c *parser.RegContext) {}

func (a *Assembler) EnterTag(c *parser.TagContext) {}

func (a *Assembler) EnterVariable(c *parser.VariableContext) {}

// NOTE no need. immediate and ptr rules parsing before this
func (a *Assembler) ExitPtr_offset(c *parser.Ptr_offsetContext) {}

func (a *Assembler) ExitInstruction(c *parser.InstructionContext) {}

func (a *Assembler) ExitEveryRule(ctx antlr.ParserRuleContext) {}

func (a *Assembler) ExitProg(c *parser.ProgContext) {}

func (a *Assembler) EnterPtr_virtual_offset(c *parser.Ptr_virtual_offsetContext) {}
