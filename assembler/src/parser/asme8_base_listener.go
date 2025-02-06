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

// EnterInst_reg_reg is called when production inst_reg_reg is entered.
func (s *BaseAsmE8Listener) EnterInst_reg_reg(ctx *Inst_reg_regContext) {}

// ExitInst_reg_reg is called when production inst_reg_reg is exited.
func (s *BaseAsmE8Listener) ExitInst_reg_reg(ctx *Inst_reg_regContext) {}

// EnterInst_reg_imm is called when production inst_reg_imm is entered.
func (s *BaseAsmE8Listener) EnterInst_reg_imm(ctx *Inst_reg_immContext) {}

// ExitInst_reg_imm is called when production inst_reg_imm is exited.
func (s *BaseAsmE8Listener) ExitInst_reg_imm(ctx *Inst_reg_immContext) {}

// EnterInst_stack_imm is called when production inst_stack_imm is entered.
func (s *BaseAsmE8Listener) EnterInst_stack_imm(ctx *Inst_stack_immContext) {}

// ExitInst_stack_imm is called when production inst_stack_imm is exited.
func (s *BaseAsmE8Listener) ExitInst_stack_imm(ctx *Inst_stack_immContext) {}

// EnterInst_index_register_imm_variable is called when production inst_index_register_imm_variable is entered.
func (s *BaseAsmE8Listener) EnterInst_index_register_imm_variable(ctx *Inst_index_register_imm_variableContext) {
}

// ExitInst_index_register_imm_variable is called when production inst_index_register_imm_variable is exited.
func (s *BaseAsmE8Listener) ExitInst_index_register_imm_variable(ctx *Inst_index_register_imm_variableContext) {
}

// EnterInst_index_imm is called when production inst_index_imm is entered.
func (s *BaseAsmE8Listener) EnterInst_index_imm(ctx *Inst_index_immContext) {}

// ExitInst_index_imm is called when production inst_index_imm is exited.
func (s *BaseAsmE8Listener) ExitInst_index_imm(ctx *Inst_index_immContext) {}

// EnterInst_reg_imm_variable is called when production inst_reg_imm_variable is entered.
func (s *BaseAsmE8Listener) EnterInst_reg_imm_variable(ctx *Inst_reg_imm_variableContext) {}

// ExitInst_reg_imm_variable is called when production inst_reg_imm_variable is exited.
func (s *BaseAsmE8Listener) ExitInst_reg_imm_variable(ctx *Inst_reg_imm_variableContext) {}

// EnterInst_ptr_reg is called when production inst_ptr_reg is entered.
func (s *BaseAsmE8Listener) EnterInst_ptr_reg(ctx *Inst_ptr_regContext) {}

// ExitInst_ptr_reg is called when production inst_ptr_reg is exited.
func (s *BaseAsmE8Listener) ExitInst_ptr_reg(ctx *Inst_ptr_regContext) {}

// EnterInst_ptr_imm is called when production inst_ptr_imm is entered.
func (s *BaseAsmE8Listener) EnterInst_ptr_imm(ctx *Inst_ptr_immContext) {}

// ExitInst_ptr_imm is called when production inst_ptr_imm is exited.
func (s *BaseAsmE8Listener) ExitInst_ptr_imm(ctx *Inst_ptr_immContext) {}

// EnterInst_reg_ptr is called when production inst_reg_ptr is entered.
func (s *BaseAsmE8Listener) EnterInst_reg_ptr(ctx *Inst_reg_ptrContext) {}

// ExitInst_reg_ptr is called when production inst_reg_ptr is exited.
func (s *BaseAsmE8Listener) ExitInst_reg_ptr(ctx *Inst_reg_ptrContext) {}

// EnterInst_reg_ptr_offset is called when production inst_reg_ptr_offset is entered.
func (s *BaseAsmE8Listener) EnterInst_reg_ptr_offset(ctx *Inst_reg_ptr_offsetContext) {}

// ExitInst_reg_ptr_offset is called when production inst_reg_ptr_offset is exited.
func (s *BaseAsmE8Listener) ExitInst_reg_ptr_offset(ctx *Inst_reg_ptr_offsetContext) {}

// EnterInst_indirect_reg_stack is called when production inst_indirect_reg_stack is entered.
func (s *BaseAsmE8Listener) EnterInst_indirect_reg_stack(ctx *Inst_indirect_reg_stackContext) {}

// ExitInst_indirect_reg_stack is called when production inst_indirect_reg_stack is exited.
func (s *BaseAsmE8Listener) ExitInst_indirect_reg_stack(ctx *Inst_indirect_reg_stackContext) {}

// EnterInst_indirect_stack_register is called when production inst_indirect_stack_register is entered.
func (s *BaseAsmE8Listener) EnterInst_indirect_stack_register(ctx *Inst_indirect_stack_registerContext) {
}

// ExitInst_indirect_stack_register is called when production inst_indirect_stack_register is exited.
func (s *BaseAsmE8Listener) ExitInst_indirect_stack_register(ctx *Inst_indirect_stack_registerContext) {
}

// EnterInst_indirect_reg_index is called when production inst_indirect_reg_index is entered.
func (s *BaseAsmE8Listener) EnterInst_indirect_reg_index(ctx *Inst_indirect_reg_indexContext) {}

// ExitInst_indirect_reg_index is called when production inst_indirect_reg_index is exited.
func (s *BaseAsmE8Listener) ExitInst_indirect_reg_index(ctx *Inst_indirect_reg_indexContext) {}

// EnterInst_indirect_index_register is called when production inst_indirect_index_register is entered.
func (s *BaseAsmE8Listener) EnterInst_indirect_index_register(ctx *Inst_indirect_index_registerContext) {
}

// ExitInst_indirect_index_register is called when production inst_indirect_index_register is exited.
func (s *BaseAsmE8Listener) ExitInst_indirect_index_register(ctx *Inst_indirect_index_registerContext) {
}

// EnterInst_ptr_offset_reg is called when production inst_ptr_offset_reg is entered.
func (s *BaseAsmE8Listener) EnterInst_ptr_offset_reg(ctx *Inst_ptr_offset_regContext) {}

// ExitInst_ptr_offset_reg is called when production inst_ptr_offset_reg is exited.
func (s *BaseAsmE8Listener) ExitInst_ptr_offset_reg(ctx *Inst_ptr_offset_regContext) {}

// EnterInst_single_reg is called when production inst_single_reg is entered.
func (s *BaseAsmE8Listener) EnterInst_single_reg(ctx *Inst_single_regContext) {}

// ExitInst_single_reg is called when production inst_single_reg is exited.
func (s *BaseAsmE8Listener) ExitInst_single_reg(ctx *Inst_single_regContext) {}

// EnterInst_implied_stack is called when production inst_implied_stack is entered.
func (s *BaseAsmE8Listener) EnterInst_implied_stack(ctx *Inst_implied_stackContext) {}

// ExitInst_implied_stack is called when production inst_implied_stack is exited.
func (s *BaseAsmE8Listener) ExitInst_implied_stack(ctx *Inst_implied_stackContext) {}

// EnterInst_implied_index is called when production inst_implied_index is entered.
func (s *BaseAsmE8Listener) EnterInst_implied_index(ctx *Inst_implied_indexContext) {}

// ExitInst_implied_index is called when production inst_implied_index is exited.
func (s *BaseAsmE8Listener) ExitInst_implied_index(ctx *Inst_implied_indexContext) {}

// EnterInst_single_imm is called when production inst_single_imm is entered.
func (s *BaseAsmE8Listener) EnterInst_single_imm(ctx *Inst_single_immContext) {}

// ExitInst_single_imm is called when production inst_single_imm is exited.
func (s *BaseAsmE8Listener) ExitInst_single_imm(ctx *Inst_single_immContext) {}

// EnterInst_single_tag is called when production inst_single_tag is entered.
func (s *BaseAsmE8Listener) EnterInst_single_tag(ctx *Inst_single_tagContext) {}

// ExitInst_single_tag is called when production inst_single_tag is exited.
func (s *BaseAsmE8Listener) ExitInst_single_tag(ctx *Inst_single_tagContext) {}

// EnterInst_single is called when production inst_single is entered.
func (s *BaseAsmE8Listener) EnterInst_single(ctx *Inst_singleContext) {}

// ExitInst_single is called when production inst_single is exited.
func (s *BaseAsmE8Listener) ExitInst_single(ctx *Inst_singleContext) {}

// EnterMnemonic is called when production mnemonic is entered.
func (s *BaseAsmE8Listener) EnterMnemonic(ctx *MnemonicContext) {}

// ExitMnemonic is called when production mnemonic is exited.
func (s *BaseAsmE8Listener) ExitMnemonic(ctx *MnemonicContext) {}

// EnterReg is called when production reg is entered.
func (s *BaseAsmE8Listener) EnterReg(ctx *RegContext) {}

// ExitReg is called when production reg is exited.
func (s *BaseAsmE8Listener) ExitReg(ctx *RegContext) {}

// EnterStack is called when production stack is entered.
func (s *BaseAsmE8Listener) EnterStack(ctx *StackContext) {}

// ExitStack is called when production stack is exited.
func (s *BaseAsmE8Listener) ExitStack(ctx *StackContext) {}

// EnterIndex is called when production index is entered.
func (s *BaseAsmE8Listener) EnterIndex(ctx *IndexContext) {}

// ExitIndex is called when production index is exited.
func (s *BaseAsmE8Listener) ExitIndex(ctx *IndexContext) {}

// EnterPtr is called when production ptr is entered.
func (s *BaseAsmE8Listener) EnterPtr(ctx *PtrContext) {}

// ExitPtr is called when production ptr is exited.
func (s *BaseAsmE8Listener) ExitPtr(ctx *PtrContext) {}

// EnterPtr_virtual_offset is called when production ptr_virtual_offset is entered.
func (s *BaseAsmE8Listener) EnterPtr_virtual_offset(ctx *Ptr_virtual_offsetContext) {}

// ExitPtr_virtual_offset is called when production ptr_virtual_offset is exited.
func (s *BaseAsmE8Listener) ExitPtr_virtual_offset(ctx *Ptr_virtual_offsetContext) {}

// EnterPtr_offset is called when production ptr_offset is entered.
func (s *BaseAsmE8Listener) EnterPtr_offset(ctx *Ptr_offsetContext) {}

// ExitPtr_offset is called when production ptr_offset is exited.
func (s *BaseAsmE8Listener) ExitPtr_offset(ctx *Ptr_offsetContext) {}

// EnterStack_offset is called when production stack_offset is entered.
func (s *BaseAsmE8Listener) EnterStack_offset(ctx *Stack_offsetContext) {}

// ExitStack_offset is called when production stack_offset is exited.
func (s *BaseAsmE8Listener) ExitStack_offset(ctx *Stack_offsetContext) {}

// EnterIndex_offset is called when production index_offset is entered.
func (s *BaseAsmE8Listener) EnterIndex_offset(ctx *Index_offsetContext) {}

// ExitIndex_offset is called when production index_offset is exited.
func (s *BaseAsmE8Listener) ExitIndex_offset(ctx *Index_offsetContext) {}

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
