package assembler

import (
	"asme8/assembler/src/parser"

	"github.com/antlr4-go/antlr/v4"
)

func (a *Assembler) EnterAccess(c *parser.AccessContext) {}

func (a *Assembler) EnterAddr_mode_imm16(c *parser.Addr_mode_imm16Context) {}

func (a *Assembler) EnterAddr_mode_implied(c *parser.Addr_mode_impliedContext) {}

func (a *Assembler) EnterAddr_mode_implied_reg16(c *parser.Addr_mode_implied_reg16Context) {}

func (a *Assembler) EnterAddr_mode_implied_reg8(c *parser.Addr_mode_implied_reg8Context) {}

func (a *Assembler) EnterAddr_mode_implied_stack(c *parser.Addr_mode_implied_stackContext) {}

func (a *Assembler) EnterAddr_mode_mem_to_reg_direct(c *parser.Addr_mode_mem_to_reg_directContext) {}

func (a *Assembler) EnterAddr_mode_mem_to_reg_indexed(c *parser.Addr_mode_mem_to_reg_indexedContext) {
}

func (a *Assembler) EnterAddr_mode_mem_to_reg_indirect(c *parser.Addr_mode_mem_to_reg_indirectContext) {
}

func (a *Assembler) EnterAddr_mode_mem_to_reg_indirect_offset(c *parser.Addr_mode_mem_to_reg_indirect_offsetContext) {
}

func (a *Assembler) EnterAddr_mode_mem_to_reg_reg16_indexed(c *parser.Addr_mode_mem_to_reg_reg16_indexedContext) {
}

func (a *Assembler) EnterAddr_mode_reg16_imm(c *parser.Addr_mode_reg16_immContext) {}

func (a *Assembler) EnterAddr_mode_reg16_imm_tag(c *parser.Addr_mode_reg16_imm_tagContext) {}

func (a *Assembler) EnterAddr_mode_reg16_reg16(c *parser.Addr_mode_reg16_reg16Context) {}

func (a *Assembler) EnterAddr_mode_reg16_stack(c *parser.Addr_mode_reg16_stackContext) {}

func (a *Assembler) EnterAddr_mode_reg8_imm8(c *parser.Addr_mode_reg8_imm8Context) {}

func (a *Assembler) EnterAddr_mode_reg8_imm8_tag(c *parser.Addr_mode_reg8_imm8_tagContext) {}

func (a *Assembler) EnterAddr_mode_reg8_reg8(c *parser.Addr_mode_reg8_reg8Context) {}

func (a *Assembler) EnterAddr_mode_reg_to_mem_direct(c *parser.Addr_mode_reg_to_mem_directContext) {}

func (a *Assembler) EnterAddr_mode_reg_to_mem_indexed(c *parser.Addr_mode_reg_to_mem_indexedContext) {
}

func (a *Assembler) EnterAddr_mode_reg_to_mem_indirect(c *parser.Addr_mode_reg_to_mem_indirectContext) {
}

func (a *Assembler) EnterAddr_mode_reg_to_mem_indirect_offset(c *parser.Addr_mode_reg_to_mem_indirect_offsetContext) {
}

func (a *Assembler) EnterAddr_mode_reg_to_mem_reg16_indexed(c *parser.Addr_mode_reg_to_mem_reg16_indexedContext) {
}

func (a *Assembler) EnterAddr_mode_stack_imm8(c *parser.Addr_mode_stack_imm8Context) {}

func (a *Assembler) EnterAddr_mode_stack_imm8_tag(c *parser.Addr_mode_stack_imm8_tagContext) {}

func (a *Assembler) EnterAddr_mode_stack_reg16(c *parser.Addr_mode_stack_reg16Context) {}

func (a *Assembler) EnterAscii(c *parser.AsciiContext) {}

func (a *Assembler) EnterDirect(c *parser.DirectContext) {}

func (a *Assembler) EnterDirect_offset(c *parser.Direct_offsetContext) {}

func (a *Assembler) EnterDirect_virtual_offset(c *parser.Direct_virtual_offsetContext) {}

func (a *Assembler) EnterDirectives(c *parser.DirectivesContext) {}

func (a *Assembler) EnterEveryRule(ctx antlr.ParserRuleContext) {}

func (a *Assembler) EnterImm(c *parser.ImmContext) {}

func (a *Assembler) EnterImm_list(c *parser.Imm_listContext) {}

func (a *Assembler) EnterIndirect(c *parser.IndirectContext) {}

func (a *Assembler) EnterIndirect_offset(c *parser.Indirect_offsetContext) {}

func (a *Assembler) EnterInst(c *parser.InstContext) {
	text := c.GetText()
	a.CodeParseEnterInst(text)
}

func (a *Assembler) EnterInstruction(c *parser.InstructionContext) {}

func (a *Assembler) EnterLabel(c *parser.LabelContext) {}

func (a *Assembler) EnterLine(c *parser.LineContext) {}

func (a *Assembler) EnterMnemonic(c *parser.MnemonicContext) {}

func (a *Assembler) EnterProg(c *parser.ProgContext) {}

func (a *Assembler) EnterReference(c *parser.ReferenceContext) {}

func (a *Assembler) EnterReg16(c *parser.Reg16Context) {}

func (a *Assembler) EnterReg16_indexed(c *parser.Reg16_indexedContext) {}

func (a *Assembler) EnterReg8(c *parser.Reg8Context) {}

func (a *Assembler) EnterSegment(c *parser.SegmentContext) {}

func (a *Assembler) EnterStack(c *parser.StackContext) {}

func (a *Assembler) EnterTag(c *parser.TagContext) {}

func (a *Assembler) EnterVariable(c *parser.VariableContext) {}

func (a *Assembler) EnterVariable_reference(c *parser.Variable_referenceContext) {}

func (a *Assembler) EnterAddr_mode_imm16_tag(c *parser.Addr_mode_imm16_tagContext) {}

func (a *Assembler) EnterAddr_mode_implied_status_register(c *parser.Addr_mode_implied_status_registerContext) {
}

func (a *Assembler) EnterStatus_register(c *parser.Status_registerContext) {}

func (a *Assembler) ExitAddr_mode_implied_status_register(c *parser.Addr_mode_implied_status_registerContext) {
	line := c.GetStart().GetLine()
	column := c.RuleIndex
	a.ParseImpliedStatusRegister(line, column)
}

func (a *Assembler) ExitStatus_register(c *parser.Status_registerContext) {}

func (a *Assembler) ExitAddr_mode_imm16_tag(c *parser.Addr_mode_imm16_tagContext) {
	line := c.GetStart().GetLine()
	column := c.RuleIndex
	a.ParseImpliedTag(line, column)
}

func (a *Assembler) ExitAccess(c *parser.AccessContext) {
	text := c.GetText()
	a.ParseAccessModifier(text)
}

func (a *Assembler) ExitAddr_mode_imm16(c *parser.Addr_mode_imm16Context) {
	line := c.GetStart().GetLine()
	column := c.RuleIndex
	a.ParseImpliedImm16(line, column)
}

func (a *Assembler) ExitAddr_mode_implied(c *parser.Addr_mode_impliedContext) {
	line := c.GetStart().GetLine()
	column := c.RuleIndex
	a.ParseImplied(line, column)
}

func (a *Assembler) ExitAddr_mode_implied_reg16(c *parser.Addr_mode_implied_reg16Context) {
	line := c.GetStart().GetLine()
	column := c.RuleIndex
	a.ParseImpliedReg16(line, column)
}

func (a *Assembler) ExitAddr_mode_implied_reg8(c *parser.Addr_mode_implied_reg8Context) {
	line := c.GetStart().GetLine()
	column := c.RuleIndex
	a.ParseImpliedReg8(line, column)
}

func (a *Assembler) ExitAddr_mode_implied_stack(c *parser.Addr_mode_implied_stackContext) {
	line := c.GetStart().GetLine()
	column := c.RuleIndex
	a.ParseImpliedStack(line, column)
}

func (a *Assembler) ExitAddr_mode_mem_to_reg_direct(c *parser.Addr_mode_mem_to_reg_directContext) {
	line := c.GetStart().GetLine()
	column := c.RuleIndex
	a.ParseDirectMemToReg(line, column)
}

func (a *Assembler) ExitAddr_mode_mem_to_reg_indexed(c *parser.Addr_mode_mem_to_reg_indexedContext) {
	line := c.GetStart().GetLine()
	column := c.RuleIndex
	a.ParseIndexedMemToReg(line, column)
}

func (a *Assembler) ExitAddr_mode_mem_to_reg_indirect(c *parser.Addr_mode_mem_to_reg_indirectContext) {
	line := c.GetStart().GetLine()
	column := c.RuleIndex
	a.ParseIndirectMemToReg(line, column)
}

func (a *Assembler) ExitAddr_mode_mem_to_reg_indirect_offset(c *parser.Addr_mode_mem_to_reg_indirect_offsetContext) {
	line := c.GetStart().GetLine()
	column := c.RuleIndex
	a.ParseIndirectOffsetMemToReg(line, column)
}

func (a *Assembler) ExitAddr_mode_mem_to_reg_reg16_indexed(c *parser.Addr_mode_mem_to_reg_reg16_indexedContext) {
	line := c.GetStart().GetLine()
	column := c.RuleIndex
	a.ParseIndexedReg16MemToReg(line, column)
}

func (a *Assembler) ExitAddr_mode_reg16_imm(c *parser.Addr_mode_reg16_immContext) {
	line := c.GetStart().GetLine()
	column := c.RuleIndex
	a.ParseReg16Imm(line, column)
}

func (a *Assembler) ExitAddr_mode_reg16_imm_tag(c *parser.Addr_mode_reg16_imm_tagContext) {
	line := c.GetStart().GetLine()
	column := c.RuleIndex
	a.GetVariableOrTagMissing(2, 16)
	a.ParseReg16Imm(line, column)
}

func (a *Assembler) ExitAddr_mode_reg16_reg16(c *parser.Addr_mode_reg16_reg16Context) {
	line := c.GetStart().GetLine()
	column := c.RuleIndex
	a.ParseReg16Reg16(line, column)
}

func (a *Assembler) ExitAddr_mode_reg16_stack(c *parser.Addr_mode_reg16_stackContext) {
	line := c.GetStart().GetLine()
	column := c.RuleIndex
	a.ParseReg16Stack(line, column)
}

func (a *Assembler) ExitAddr_mode_reg8_imm8(c *parser.Addr_mode_reg8_imm8Context) {
	line := c.GetStart().GetLine()
	column := c.RuleIndex
	a.ParseReg8Imm8(line, column)
}

func (a *Assembler) ExitAddr_mode_reg8_imm8_tag(c *parser.Addr_mode_reg8_imm8_tagContext) {
	line := c.GetStart().GetLine()
	column := c.RuleIndex
	a.GetVariableOrTagMissing(2, 8)
	a.ParseReg8Imm8(line, column)
}

func (a *Assembler) ExitAddr_mode_reg8_reg8(c *parser.Addr_mode_reg8_reg8Context) {
	line := c.GetStart().GetLine()
	column := c.RuleIndex
	a.ParseReg8Reg8(line, column)
}

func (a *Assembler) ExitAddr_mode_reg_to_mem_direct(c *parser.Addr_mode_reg_to_mem_directContext) {
	line := c.GetStart().GetLine()
	column := c.RuleIndex
	a.ParseDirectRegToMem(line, column)
}

func (a *Assembler) ExitAddr_mode_reg_to_mem_indexed(c *parser.Addr_mode_reg_to_mem_indexedContext) {
	line := c.GetStart().GetLine()
	column := c.RuleIndex
	a.ParseIndexedRegToMem(line, column)
}

func (a *Assembler) ExitAddr_mode_reg_to_mem_indirect(c *parser.Addr_mode_reg_to_mem_indirectContext) {
	line := c.GetStart().GetLine()
	column := c.RuleIndex
	a.ParseIndirectRegToMem(line, column)
}

func (a *Assembler) ExitAddr_mode_reg_to_mem_indirect_offset(c *parser.Addr_mode_reg_to_mem_indirect_offsetContext) {
	line := c.GetStart().GetLine()
	column := c.RuleIndex
	a.ParseIndirectOffsetRegToMem(line, column)
}

func (a *Assembler) ExitAddr_mode_reg_to_mem_reg16_indexed(c *parser.Addr_mode_reg_to_mem_reg16_indexedContext) {
	line := c.GetStart().GetLine()
	column := c.RuleIndex
	a.ParseIndexedReg16RegToMem(line, column)
}

func (a *Assembler) ExitAddr_mode_stack_imm8(c *parser.Addr_mode_stack_imm8Context) {
	line := c.GetStart().GetLine()
	column := c.RuleIndex
	a.ParseStackImm8(line, column)
}

func (a *Assembler) ExitAddr_mode_stack_imm8_tag(c *parser.Addr_mode_stack_imm8_tagContext) {
	line := c.GetStart().GetLine()
	column := c.RuleIndex
	a.GetVariableOrTagMissing(2, 8)
	a.ParseStackImm8(line, column)
}

func (a *Assembler) ExitAddr_mode_stack_reg16(c *parser.Addr_mode_stack_reg16Context) {
	line := c.GetStart().GetLine()
	column := c.RuleIndex
	a.ParseStackReg16(line, column)
}

func (a *Assembler) ExitAscii(c *parser.AsciiContext) {
	text := c.GetText()
	a.ParseAscii(text)
}

func (a *Assembler) ExitDirect(c *parser.DirectContext) {
	text := c.GetText()
	line := c.GetStart().GetLine()
	column := c.RuleIndex
	a.ParsePtr(text, line, column)
}

func (a *Assembler) ExitDirect_offset(c *parser.Direct_offsetContext) {}

func (a *Assembler) ExitDirect_virtual_offset(c *parser.Direct_virtual_offsetContext) {
	a.hasVirtualOffset = true
}

func (a *Assembler) ExitDirectives(c *parser.DirectivesContext) {
	text := c.GetText()
	line := c.GetStart().GetLine()
	column := c.RuleIndex
	a.ParseDirective(text, line, column)
}

func (a *Assembler) ExitEveryRule(ctx antlr.ParserRuleContext) {}

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

func (a *Assembler) ExitIndirect(c *parser.IndirectContext) {}

func (a *Assembler) ExitIndirect_offset(c *parser.Indirect_offsetContext) {}

func (a *Assembler) ExitInst(c *parser.InstContext) {
	text := c.GetText()
	a.CodeParseExitInst(text)
}

func (a *Assembler) ExitInstruction(c *parser.InstructionContext) {}

func (a *Assembler) ExitLabel(c *parser.LabelContext) {
	line := c.GetStart().GetLine()
	column := c.RuleIndex
	a.ParseLabel(a.currentTag.Text, line, column)
}

func (a *Assembler) ExitLine(c *parser.LineContext) {
	text := c.GetText()
	a.ResetCurrentValues()
	a.CodeParseLastLine(text)
}

func (a *Assembler) ExitMnemonic(c *parser.MnemonicContext) {
	a.currentInstruction = c.GetText()
}

func (a *Assembler) ExitProg(c *parser.ProgContext) {}

func (a *Assembler) ExitReference(c *parser.ReferenceContext) {
	text := c.GetText()
	a.ParseReference(text)
}

func (a *Assembler) ExitReg16(c *parser.Reg16Context) {
	text := c.GetText()
	a.ParseRegister(text)
}

func (a *Assembler) ExitReg16_indexed(c *parser.Reg16_indexedContext) {}

func (a *Assembler) ExitReg8(c *parser.Reg8Context) {
	text := c.GetText()
	a.ParseRegister(text)
}

func (a *Assembler) ExitSegment(c *parser.SegmentContext) {
	text := c.GetText()
	a.ParseSegment(text)
}

func (a *Assembler) ExitStack(c *parser.StackContext) {
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
	a.ParseVariable(a.currentTag.Text, a.currentValue.Copy())
}

func (a *Assembler) ExitVariable_reference(c *parser.Variable_referenceContext) {
	a.ParseVariableReference()
}

func (a *Assembler) VisitErrorNode(node antlr.ErrorNode) {}

func (a *Assembler) VisitTerminal(node antlr.TerminalNode) {}
