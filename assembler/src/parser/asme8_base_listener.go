// Code generated from AsmE8.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // AsmE8

import "github.com/antlr4-go/antlr/v4"

// BaseAsmE8Listener is a complete listener for a parse tree produced by AsmE8Parser.
type BaseAsmE8Listener struct{}

var _ AsmE8Listener = &BaseAsmE8Listener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseAsmE8Listener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseAsmE8Listener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseAsmE8Listener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseAsmE8Listener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProg is called when production prog is entered.
func (s *BaseAsmE8Listener) EnterProg(ctx *ProgContext) {}

// ExitProg is called when production prog is exited.
func (s *BaseAsmE8Listener) ExitProg(ctx *ProgContext) {}

// EnterInstruction is called when production instruction is entered.
func (s *BaseAsmE8Listener) EnterInstruction(ctx *InstructionContext) {}

// ExitInstruction is called when production instruction is exited.
func (s *BaseAsmE8Listener) ExitInstruction(ctx *InstructionContext) {}

// EnterLine is called when production line is entered.
func (s *BaseAsmE8Listener) EnterLine(ctx *LineContext) {}

// ExitLine is called when production line is exited.
func (s *BaseAsmE8Listener) ExitLine(ctx *LineContext) {}

// EnterLabel is called when production label is entered.
func (s *BaseAsmE8Listener) EnterLabel(ctx *LabelContext) {}

// ExitLabel is called when production label is exited.
func (s *BaseAsmE8Listener) ExitLabel(ctx *LabelContext) {}

// EnterInst is called when production inst is entered.
func (s *BaseAsmE8Listener) EnterInst(ctx *InstContext) {}

// ExitInst is called when production inst is exited.
func (s *BaseAsmE8Listener) ExitInst(ctx *InstContext) {}

// EnterAddr_mode_imm16 is called when production addr_mode_imm16 is entered.
func (s *BaseAsmE8Listener) EnterAddr_mode_imm16(ctx *Addr_mode_imm16Context) {}

// ExitAddr_mode_imm16 is called when production addr_mode_imm16 is exited.
func (s *BaseAsmE8Listener) ExitAddr_mode_imm16(ctx *Addr_mode_imm16Context) {}

// EnterAddr_mode_imm16_tag is called when production addr_mode_imm16_tag is entered.
func (s *BaseAsmE8Listener) EnterAddr_mode_imm16_tag(ctx *Addr_mode_imm16_tagContext) {}

// ExitAddr_mode_imm16_tag is called when production addr_mode_imm16_tag is exited.
func (s *BaseAsmE8Listener) ExitAddr_mode_imm16_tag(ctx *Addr_mode_imm16_tagContext) {}

// EnterAddr_mode_reg_to_mem_indexed is called when production addr_mode_reg_to_mem_indexed is entered.
func (s *BaseAsmE8Listener) EnterAddr_mode_reg_to_mem_indexed(ctx *Addr_mode_reg_to_mem_indexedContext) {
}

// ExitAddr_mode_reg_to_mem_indexed is called when production addr_mode_reg_to_mem_indexed is exited.
func (s *BaseAsmE8Listener) ExitAddr_mode_reg_to_mem_indexed(ctx *Addr_mode_reg_to_mem_indexedContext) {
}

// EnterAddr_mode_reg_to_mem_direct is called when production addr_mode_reg_to_mem_direct is entered.
func (s *BaseAsmE8Listener) EnterAddr_mode_reg_to_mem_direct(ctx *Addr_mode_reg_to_mem_directContext) {
}

// ExitAddr_mode_reg_to_mem_direct is called when production addr_mode_reg_to_mem_direct is exited.
func (s *BaseAsmE8Listener) ExitAddr_mode_reg_to_mem_direct(ctx *Addr_mode_reg_to_mem_directContext) {
}

// EnterAddr_mode_reg_to_mem_indirect_offset is called when production addr_mode_reg_to_mem_indirect_offset is entered.
func (s *BaseAsmE8Listener) EnterAddr_mode_reg_to_mem_indirect_offset(ctx *Addr_mode_reg_to_mem_indirect_offsetContext) {
}

// ExitAddr_mode_reg_to_mem_indirect_offset is called when production addr_mode_reg_to_mem_indirect_offset is exited.
func (s *BaseAsmE8Listener) ExitAddr_mode_reg_to_mem_indirect_offset(ctx *Addr_mode_reg_to_mem_indirect_offsetContext) {
}

// EnterAddr_mode_reg_to_mem_reg16_indexed is called when production addr_mode_reg_to_mem_reg16_indexed is entered.
func (s *BaseAsmE8Listener) EnterAddr_mode_reg_to_mem_reg16_indexed(ctx *Addr_mode_reg_to_mem_reg16_indexedContext) {
}

// ExitAddr_mode_reg_to_mem_reg16_indexed is called when production addr_mode_reg_to_mem_reg16_indexed is exited.
func (s *BaseAsmE8Listener) ExitAddr_mode_reg_to_mem_reg16_indexed(ctx *Addr_mode_reg_to_mem_reg16_indexedContext) {
}

// EnterAddr_mode_reg_to_mem_indirect is called when production addr_mode_reg_to_mem_indirect is entered.
func (s *BaseAsmE8Listener) EnterAddr_mode_reg_to_mem_indirect(ctx *Addr_mode_reg_to_mem_indirectContext) {
}

// ExitAddr_mode_reg_to_mem_indirect is called when production addr_mode_reg_to_mem_indirect is exited.
func (s *BaseAsmE8Listener) ExitAddr_mode_reg_to_mem_indirect(ctx *Addr_mode_reg_to_mem_indirectContext) {
}

// EnterAddr_mode_implied_reg8 is called when production addr_mode_implied_reg8 is entered.
func (s *BaseAsmE8Listener) EnterAddr_mode_implied_reg8(ctx *Addr_mode_implied_reg8Context) {}

// ExitAddr_mode_implied_reg8 is called when production addr_mode_implied_reg8 is exited.
func (s *BaseAsmE8Listener) ExitAddr_mode_implied_reg8(ctx *Addr_mode_implied_reg8Context) {}

// EnterAddr_mode_reg8_imm8 is called when production addr_mode_reg8_imm8 is entered.
func (s *BaseAsmE8Listener) EnterAddr_mode_reg8_imm8(ctx *Addr_mode_reg8_imm8Context) {}

// ExitAddr_mode_reg8_imm8 is called when production addr_mode_reg8_imm8 is exited.
func (s *BaseAsmE8Listener) ExitAddr_mode_reg8_imm8(ctx *Addr_mode_reg8_imm8Context) {}

// EnterAddr_mode_reg8_imm8_tag is called when production addr_mode_reg8_imm8_tag is entered.
func (s *BaseAsmE8Listener) EnterAddr_mode_reg8_imm8_tag(ctx *Addr_mode_reg8_imm8_tagContext) {}

// ExitAddr_mode_reg8_imm8_tag is called when production addr_mode_reg8_imm8_tag is exited.
func (s *BaseAsmE8Listener) ExitAddr_mode_reg8_imm8_tag(ctx *Addr_mode_reg8_imm8_tagContext) {}

// EnterAddr_mode_mem_to_reg_indexed is called when production addr_mode_mem_to_reg_indexed is entered.
func (s *BaseAsmE8Listener) EnterAddr_mode_mem_to_reg_indexed(ctx *Addr_mode_mem_to_reg_indexedContext) {
}

// ExitAddr_mode_mem_to_reg_indexed is called when production addr_mode_mem_to_reg_indexed is exited.
func (s *BaseAsmE8Listener) ExitAddr_mode_mem_to_reg_indexed(ctx *Addr_mode_mem_to_reg_indexedContext) {
}

// EnterAddr_mode_mem_to_reg_direct is called when production addr_mode_mem_to_reg_direct is entered.
func (s *BaseAsmE8Listener) EnterAddr_mode_mem_to_reg_direct(ctx *Addr_mode_mem_to_reg_directContext) {
}

// ExitAddr_mode_mem_to_reg_direct is called when production addr_mode_mem_to_reg_direct is exited.
func (s *BaseAsmE8Listener) ExitAddr_mode_mem_to_reg_direct(ctx *Addr_mode_mem_to_reg_directContext) {
}

// EnterAddr_mode_mem_to_reg_indirect_offset is called when production addr_mode_mem_to_reg_indirect_offset is entered.
func (s *BaseAsmE8Listener) EnterAddr_mode_mem_to_reg_indirect_offset(ctx *Addr_mode_mem_to_reg_indirect_offsetContext) {
}

// ExitAddr_mode_mem_to_reg_indirect_offset is called when production addr_mode_mem_to_reg_indirect_offset is exited.
func (s *BaseAsmE8Listener) ExitAddr_mode_mem_to_reg_indirect_offset(ctx *Addr_mode_mem_to_reg_indirect_offsetContext) {
}

// EnterAddr_mode_mem_to_reg_reg16_indexed is called when production addr_mode_mem_to_reg_reg16_indexed is entered.
func (s *BaseAsmE8Listener) EnterAddr_mode_mem_to_reg_reg16_indexed(ctx *Addr_mode_mem_to_reg_reg16_indexedContext) {
}

// ExitAddr_mode_mem_to_reg_reg16_indexed is called when production addr_mode_mem_to_reg_reg16_indexed is exited.
func (s *BaseAsmE8Listener) ExitAddr_mode_mem_to_reg_reg16_indexed(ctx *Addr_mode_mem_to_reg_reg16_indexedContext) {
}

// EnterAddr_mode_mem_to_reg_indirect is called when production addr_mode_mem_to_reg_indirect is entered.
func (s *BaseAsmE8Listener) EnterAddr_mode_mem_to_reg_indirect(ctx *Addr_mode_mem_to_reg_indirectContext) {
}

// ExitAddr_mode_mem_to_reg_indirect is called when production addr_mode_mem_to_reg_indirect is exited.
func (s *BaseAsmE8Listener) ExitAddr_mode_mem_to_reg_indirect(ctx *Addr_mode_mem_to_reg_indirectContext) {
}

// EnterAddr_mode_reg8_reg8 is called when production addr_mode_reg8_reg8 is entered.
func (s *BaseAsmE8Listener) EnterAddr_mode_reg8_reg8(ctx *Addr_mode_reg8_reg8Context) {}

// ExitAddr_mode_reg8_reg8 is called when production addr_mode_reg8_reg8 is exited.
func (s *BaseAsmE8Listener) ExitAddr_mode_reg8_reg8(ctx *Addr_mode_reg8_reg8Context) {}

// EnterAddr_mode_implied is called when production addr_mode_implied is entered.
func (s *BaseAsmE8Listener) EnterAddr_mode_implied(ctx *Addr_mode_impliedContext) {}

// ExitAddr_mode_implied is called when production addr_mode_implied is exited.
func (s *BaseAsmE8Listener) ExitAddr_mode_implied(ctx *Addr_mode_impliedContext) {}

// EnterAddr_mode_implied_reg16 is called when production addr_mode_implied_reg16 is entered.
func (s *BaseAsmE8Listener) EnterAddr_mode_implied_reg16(ctx *Addr_mode_implied_reg16Context) {}

// ExitAddr_mode_implied_reg16 is called when production addr_mode_implied_reg16 is exited.
func (s *BaseAsmE8Listener) ExitAddr_mode_implied_reg16(ctx *Addr_mode_implied_reg16Context) {}

// EnterAddr_mode_reg16_imm is called when production addr_mode_reg16_imm is entered.
func (s *BaseAsmE8Listener) EnterAddr_mode_reg16_imm(ctx *Addr_mode_reg16_immContext) {}

// ExitAddr_mode_reg16_imm is called when production addr_mode_reg16_imm is exited.
func (s *BaseAsmE8Listener) ExitAddr_mode_reg16_imm(ctx *Addr_mode_reg16_immContext) {}

// EnterAddr_mode_reg16_imm_tag is called when production addr_mode_reg16_imm_tag is entered.
func (s *BaseAsmE8Listener) EnterAddr_mode_reg16_imm_tag(ctx *Addr_mode_reg16_imm_tagContext) {}

// ExitAddr_mode_reg16_imm_tag is called when production addr_mode_reg16_imm_tag is exited.
func (s *BaseAsmE8Listener) ExitAddr_mode_reg16_imm_tag(ctx *Addr_mode_reg16_imm_tagContext) {}

// EnterAddr_mode_reg16_reg16 is called when production addr_mode_reg16_reg16 is entered.
func (s *BaseAsmE8Listener) EnterAddr_mode_reg16_reg16(ctx *Addr_mode_reg16_reg16Context) {}

// ExitAddr_mode_reg16_reg16 is called when production addr_mode_reg16_reg16 is exited.
func (s *BaseAsmE8Listener) ExitAddr_mode_reg16_reg16(ctx *Addr_mode_reg16_reg16Context) {}

// EnterAddr_mode_reg16_stack is called when production addr_mode_reg16_stack is entered.
func (s *BaseAsmE8Listener) EnterAddr_mode_reg16_stack(ctx *Addr_mode_reg16_stackContext) {}

// ExitAddr_mode_reg16_stack is called when production addr_mode_reg16_stack is exited.
func (s *BaseAsmE8Listener) ExitAddr_mode_reg16_stack(ctx *Addr_mode_reg16_stackContext) {}

// EnterAddr_mode_implied_stack is called when production addr_mode_implied_stack is entered.
func (s *BaseAsmE8Listener) EnterAddr_mode_implied_stack(ctx *Addr_mode_implied_stackContext) {}

// ExitAddr_mode_implied_stack is called when production addr_mode_implied_stack is exited.
func (s *BaseAsmE8Listener) ExitAddr_mode_implied_stack(ctx *Addr_mode_implied_stackContext) {}

// EnterAddr_mode_implied_status_register is called when production addr_mode_implied_status_register is entered.
func (s *BaseAsmE8Listener) EnterAddr_mode_implied_status_register(ctx *Addr_mode_implied_status_registerContext) {
}

// ExitAddr_mode_implied_status_register is called when production addr_mode_implied_status_register is exited.
func (s *BaseAsmE8Listener) ExitAddr_mode_implied_status_register(ctx *Addr_mode_implied_status_registerContext) {
}

// EnterAddr_mode_stack_imm8 is called when production addr_mode_stack_imm8 is entered.
func (s *BaseAsmE8Listener) EnterAddr_mode_stack_imm8(ctx *Addr_mode_stack_imm8Context) {}

// ExitAddr_mode_stack_imm8 is called when production addr_mode_stack_imm8 is exited.
func (s *BaseAsmE8Listener) ExitAddr_mode_stack_imm8(ctx *Addr_mode_stack_imm8Context) {}

// EnterAddr_mode_stack_imm8_tag is called when production addr_mode_stack_imm8_tag is entered.
func (s *BaseAsmE8Listener) EnterAddr_mode_stack_imm8_tag(ctx *Addr_mode_stack_imm8_tagContext) {}

// ExitAddr_mode_stack_imm8_tag is called when production addr_mode_stack_imm8_tag is exited.
func (s *BaseAsmE8Listener) ExitAddr_mode_stack_imm8_tag(ctx *Addr_mode_stack_imm8_tagContext) {}

// EnterAddr_mode_stack_reg16 is called when production addr_mode_stack_reg16 is entered.
func (s *BaseAsmE8Listener) EnterAddr_mode_stack_reg16(ctx *Addr_mode_stack_reg16Context) {}

// ExitAddr_mode_stack_reg16 is called when production addr_mode_stack_reg16 is exited.
func (s *BaseAsmE8Listener) ExitAddr_mode_stack_reg16(ctx *Addr_mode_stack_reg16Context) {}

// EnterMnemonic is called when production mnemonic is entered.
func (s *BaseAsmE8Listener) EnterMnemonic(ctx *MnemonicContext) {}

// ExitMnemonic is called when production mnemonic is exited.
func (s *BaseAsmE8Listener) ExitMnemonic(ctx *MnemonicContext) {}

// EnterReg8 is called when production reg8 is entered.
func (s *BaseAsmE8Listener) EnterReg8(ctx *Reg8Context) {}

// ExitReg8 is called when production reg8 is exited.
func (s *BaseAsmE8Listener) ExitReg8(ctx *Reg8Context) {}

// EnterReg16 is called when production reg16 is entered.
func (s *BaseAsmE8Listener) EnterReg16(ctx *Reg16Context) {}

// ExitReg16 is called when production reg16 is exited.
func (s *BaseAsmE8Listener) ExitReg16(ctx *Reg16Context) {}

// EnterStack is called when production stack is entered.
func (s *BaseAsmE8Listener) EnterStack(ctx *StackContext) {}

// ExitStack is called when production stack is exited.
func (s *BaseAsmE8Listener) ExitStack(ctx *StackContext) {}

// EnterStatus_register is called when production status_register is entered.
func (s *BaseAsmE8Listener) EnterStatus_register(ctx *Status_registerContext) {}

// ExitStatus_register is called when production status_register is exited.
func (s *BaseAsmE8Listener) ExitStatus_register(ctx *Status_registerContext) {}

// EnterDirect is called when production direct is entered.
func (s *BaseAsmE8Listener) EnterDirect(ctx *DirectContext) {}

// ExitDirect is called when production direct is exited.
func (s *BaseAsmE8Listener) ExitDirect(ctx *DirectContext) {}

// EnterDirect_virtual_offset is called when production direct_virtual_offset is entered.
func (s *BaseAsmE8Listener) EnterDirect_virtual_offset(ctx *Direct_virtual_offsetContext) {}

// ExitDirect_virtual_offset is called when production direct_virtual_offset is exited.
func (s *BaseAsmE8Listener) ExitDirect_virtual_offset(ctx *Direct_virtual_offsetContext) {}

// EnterDirect_offset is called when production direct_offset is entered.
func (s *BaseAsmE8Listener) EnterDirect_offset(ctx *Direct_offsetContext) {}

// ExitDirect_offset is called when production direct_offset is exited.
func (s *BaseAsmE8Listener) ExitDirect_offset(ctx *Direct_offsetContext) {}

// EnterIndirect_offset is called when production indirect_offset is entered.
func (s *BaseAsmE8Listener) EnterIndirect_offset(ctx *Indirect_offsetContext) {}

// ExitIndirect_offset is called when production indirect_offset is exited.
func (s *BaseAsmE8Listener) ExitIndirect_offset(ctx *Indirect_offsetContext) {}

// EnterReg16_indexed is called when production reg16_indexed is entered.
func (s *BaseAsmE8Listener) EnterReg16_indexed(ctx *Reg16_indexedContext) {}

// ExitReg16_indexed is called when production reg16_indexed is exited.
func (s *BaseAsmE8Listener) ExitReg16_indexed(ctx *Reg16_indexedContext) {}

// EnterIndirect is called when production indirect is entered.
func (s *BaseAsmE8Listener) EnterIndirect(ctx *IndirectContext) {}

// ExitIndirect is called when production indirect is exited.
func (s *BaseAsmE8Listener) ExitIndirect(ctx *IndirectContext) {}

// EnterVariable is called when production variable is entered.
func (s *BaseAsmE8Listener) EnterVariable(ctx *VariableContext) {}

// ExitVariable is called when production variable is exited.
func (s *BaseAsmE8Listener) ExitVariable(ctx *VariableContext) {}

// EnterVariable_reference is called when production variable_reference is entered.
func (s *BaseAsmE8Listener) EnterVariable_reference(ctx *Variable_referenceContext) {}

// ExitVariable_reference is called when production variable_reference is exited.
func (s *BaseAsmE8Listener) ExitVariable_reference(ctx *Variable_referenceContext) {}

// EnterReference is called when production reference is entered.
func (s *BaseAsmE8Listener) EnterReference(ctx *ReferenceContext) {}

// ExitReference is called when production reference is exited.
func (s *BaseAsmE8Listener) ExitReference(ctx *ReferenceContext) {}

// EnterDirectives is called when production directives is entered.
func (s *BaseAsmE8Listener) EnterDirectives(ctx *DirectivesContext) {}

// ExitDirectives is called when production directives is exited.
func (s *BaseAsmE8Listener) ExitDirectives(ctx *DirectivesContext) {}

// EnterSegment is called when production segment is entered.
func (s *BaseAsmE8Listener) EnterSegment(ctx *SegmentContext) {}

// ExitSegment is called when production segment is exited.
func (s *BaseAsmE8Listener) ExitSegment(ctx *SegmentContext) {}

// EnterAccess is called when production access is entered.
func (s *BaseAsmE8Listener) EnterAccess(ctx *AccessContext) {}

// ExitAccess is called when production access is exited.
func (s *BaseAsmE8Listener) ExitAccess(ctx *AccessContext) {}

// EnterImm_list is called when production imm_list is entered.
func (s *BaseAsmE8Listener) EnterImm_list(ctx *Imm_listContext) {}

// ExitImm_list is called when production imm_list is exited.
func (s *BaseAsmE8Listener) ExitImm_list(ctx *Imm_listContext) {}

// EnterImm is called when production imm is entered.
func (s *BaseAsmE8Listener) EnterImm(ctx *ImmContext) {}

// ExitImm is called when production imm is exited.
func (s *BaseAsmE8Listener) ExitImm(ctx *ImmContext) {}

// EnterAscii is called when production ascii is entered.
func (s *BaseAsmE8Listener) EnterAscii(ctx *AsciiContext) {}

// ExitAscii is called when production ascii is exited.
func (s *BaseAsmE8Listener) ExitAscii(ctx *AsciiContext) {}

// EnterTag is called when production tag is entered.
func (s *BaseAsmE8Listener) EnterTag(ctx *TagContext) {}

// ExitTag is called when production tag is exited.
func (s *BaseAsmE8Listener) ExitTag(ctx *TagContext) {}
