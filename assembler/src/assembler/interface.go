package assembler

import (
	"asme8/assembler/src/parser"
	"asme8/common/instruction"

	"github.com/antlr4-go/antlr/v4"
)

func (a *Assembler) ExitInst_index_register_imm_variable(c *parser.Inst_index_register_imm_variableContext) {
	a.GetVariableOrTagMissing(1, 16)
	line := c.GetStart().GetLine()
	column := c.RuleIndex
	opcode := a.GetOrFailOpCode(a.currentInstruction, instruction.ADDRESSING_MODE_IP_IMM_16, line, column)
	a.AppendMachineCode(opcode)
	a.AppendMachineCode(a.currentValue.GetLowByte())
	a.AppendMachineCode(a.currentValue.GetHighByte())
}

func (a *Assembler) ExitInst_ptr_imm(c *parser.Inst_ptr_immContext) {
	text := c.GetText()
	line := c.GetStart().GetLine()
	column := c.RuleIndex
	a.ParsePtrImm(text, line, column)
}

func (a *Assembler) ExitSegment(c *parser.SegmentContext) {
	text := c.GetText()
	a.ParseSegment(text)
}

func (a *Assembler) ExitAccess(c *parser.AccessContext) {
	text := c.GetText()
	a.ParseAccessModifier(text)
}

func (a *Assembler) ExitDirectives(c *parser.DirectivesContext) {
	text := c.GetText()
	line := c.GetStart().GetLine()
	column := c.RuleIndex
	a.ParseDirective(text, line, column)
}

func (a *Assembler) ExitImm(c *parser.ImmContext) {
	text := c.GetText()
	line := c.GetStart().GetLine()
	column := c.RuleIndex
	a.ParseValue(text, line, column)
}

func (a *Assembler) ExitImm_list(c *parser.Imm_listContext) {
	text := c.GetText()
	line := c.GetStart().GetLine()
	column := c.RuleIndex
	a.ParseValueList(text, line, column)
}

func (a *Assembler) ExitIndex(c *parser.IndexContext) {
	text := c.GetText()
	a.ParseRegister(text)
}

func (a *Assembler) ExitIndex_offset(c *parser.Index_offsetContext) {}

func (a *Assembler) ExitInst_implied_index(c *parser.Inst_implied_indexContext) {
	line := c.GetStart().GetLine()
	column := c.RuleIndex
	a.ParseImpliedIndex(line, column)
}

func (a *Assembler) ExitInst_index_imm(c *parser.Inst_index_immContext) {
	line := c.GetStart().GetLine()
	column := c.RuleIndex
	a.ParseIndexImmediate(line, column)
}

func (a *Assembler) ExitInst_indirect_index_register(c *parser.Inst_indirect_index_registerContext) {
	line := c.GetStart().GetLine()
	column := c.RuleIndex
	a.ParseIndexRegister(true, line, column)
}

func (a *Assembler) ExitInst_indirect_reg_index(c *parser.Inst_indirect_reg_indexContext) {
	line := c.GetStart().GetLine()
	column := c.RuleIndex
	a.ParseIndexRegister(false, line, column)
}

func (a *Assembler) ExitInst_implied_stack(c *parser.Inst_implied_stackContext) {
	line := c.GetStart().GetLine()
	column := c.RuleIndex
	a.ParseImpliedStack(line, column)
}

func (a *Assembler) ExitStack(c *parser.StackContext) {
	text := c.GetText()
	a.ParseRegister(text)
}

func (a *Assembler) ExitInst_ptr_offset_reg(c *parser.Inst_ptr_offset_regContext) {
	line := c.GetStart().GetLine()
	column := c.RuleIndex
	a.ParsePointerRegisterOffset(line, column)
}

func (a *Assembler) ExitInst_ptr_reg(c *parser.Inst_ptr_regContext) {
	line := c.GetStart().GetLine()
	column := c.RuleIndex
	a.ParsePointerRegister(line, column)
}

func (a *Assembler) ExitInst_reg_imm(c *parser.Inst_reg_immContext) {
	line := c.GetStart().GetLine()
	column := c.RuleIndex
	a.ParseRegisterImmediate(line, column)
}

func (a *Assembler) ExitInst_stack_imm(c *parser.Inst_stack_immContext) {
	line := c.GetStart().GetLine()
	column := c.RuleIndex
	a.ParseStackImmediate(line, column)
}

func (a *Assembler) ExitInst_reg_imm_variable(c *parser.Inst_reg_imm_variableContext) {
	line := c.GetStart().GetLine()
	column := c.RuleIndex
	a.ParseRegisterImmediateVariable(line, column)
}

func (a *Assembler) ExitInst_reg_ptr(c *parser.Inst_reg_ptrContext) {
	line := c.GetStart().GetLine()
	column := c.RuleIndex
	a.ParseRegisterPointer(line, column)
}

func (a *Assembler) ExitInst_reg_ptr_offset(c *parser.Inst_reg_ptr_offsetContext) {
	line := c.GetStart().GetLine()
	column := c.RuleIndex
	a.ParseRegisterPointerOffset(line, column)
}

func (a *Assembler) ExitInst_indirect_reg_stack(c *parser.Inst_indirect_reg_stackContext) {
	line := c.GetStart().GetLine()
	column := c.RuleIndex
	a.ParseStackRegister(false, line, column)
}

func (a *Assembler) ExitInst_indirect_stack_register(c *parser.Inst_indirect_stack_registerContext) {
	line := c.GetStart().GetLine()
	column := c.RuleIndex
	a.ParseStackRegister(true, line, column)
}

func (a *Assembler) ExitStack_offset(c *parser.Stack_offsetContext) {}

func (a *Assembler) ExitInst_reg_reg(c *parser.Inst_reg_regContext) {
	line := c.GetStart().GetLine()
	column := c.RuleIndex
	a.ParseRegisterRegister(line, column)
}

func (a *Assembler) ExitInst_single(c *parser.Inst_singleContext) {
	line := c.GetStart().GetLine()
	column := c.RuleIndex
	a.ParseImplied(line, column)
}

func (a *Assembler) ExitInst_single_imm(c *parser.Inst_single_immContext) {
	line := c.GetStart().GetLine()
	column := c.RuleIndex
	a.ParseImpliedImmediate(line, column)
}

func (a *Assembler) ExitInst_single_reg(c *parser.Inst_single_regContext) {
	line := c.GetStart().GetLine()
	column := c.RuleIndex
	a.ParseImpliedRegister(line, column)
}

func (a *Assembler) ExitInst_single_tag(c *parser.Inst_single_tagContext) {
	line := c.GetStart().GetLine()
	column := c.RuleIndex
	a.ParseImpliedTag(line, column)
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

func (a *Assembler) ExitReference(c *parser.ReferenceContext) {
	text := c.GetText()
	a.ParseReference(text)
}

func (a *Assembler) ExitVariable_reference(c *parser.Variable_referenceContext) {
	a.ParseVariableReference()
}

func (a *Assembler) ExitVariable(c *parser.VariableContext) {
	a.ParseVariable(a.currentTag.Text, a.currentValue.Copy())
}

func (a *Assembler) ExitInst(c *parser.InstContext) {
	text := c.GetText()
	a.CodeParseExitInst(text)
}

func (a *Assembler) EnterReference(c *parser.ReferenceContext) {}

func (a *Assembler) EnterVariable_reference(c *parser.Variable_referenceContext) {}

func (a *Assembler) EnterInst_index_register_imm_variable(c *parser.Inst_index_register_imm_variableContext) {
}

func (a *Assembler) EnterInst_ptr_imm(c *parser.Inst_ptr_immContext) {}

func (a *Assembler) EnterIndex(c *parser.IndexContext) {}

func (a *Assembler) EnterIndex_offset(c *parser.Index_offsetContext) {}

func (a *Assembler) EnterInst_implied_index(c *parser.Inst_implied_indexContext) {}

func (a *Assembler) EnterInst_index_imm(c *parser.Inst_index_immContext) {}

func (a *Assembler) EnterInst_indirect_index_register(c *parser.Inst_indirect_index_registerContext) {
}

func (a *Assembler) EnterInst_indirect_reg_index(c *parser.Inst_indirect_reg_indexContext) {}

func (a *Assembler) EnterInst_indirect_reg_stack(c *parser.Inst_indirect_reg_stackContext) {}

func (a *Assembler) EnterInst_indirect_stack_register(c *parser.Inst_indirect_stack_registerContext) {
}

func (a *Assembler) EnterStack_offset(c *parser.Stack_offsetContext) {}

func (a *Assembler) EnterInst_stack_imm(c *parser.Inst_stack_immContext) {}

func (a *Assembler) EnterInst_implied_stack(c *parser.Inst_implied_stackContext) {}

func (a *Assembler) EnterStack(c *parser.StackContext) {}

func (a *Assembler) EnterSegment(c *parser.SegmentContext) {}

func (a *Assembler) EnterAccess(c *parser.AccessContext) {}

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
