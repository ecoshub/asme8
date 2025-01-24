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

// EnterInst_reg_imm_variable is called when production inst_reg_imm_variable is entered.
func (s *BaseAsmE8Listener) EnterInst_reg_imm_variable(ctx *Inst_reg_imm_variableContext) {}

// ExitInst_reg_imm_variable is called when production inst_reg_imm_variable is exited.
func (s *BaseAsmE8Listener) ExitInst_reg_imm_variable(ctx *Inst_reg_imm_variableContext) {}

// EnterInst_ptr_reg is called when production inst_ptr_reg is entered.
func (s *BaseAsmE8Listener) EnterInst_ptr_reg(ctx *Inst_ptr_regContext) {}

// ExitInst_ptr_reg is called when production inst_ptr_reg is exited.
func (s *BaseAsmE8Listener) ExitInst_ptr_reg(ctx *Inst_ptr_regContext) {}

// EnterInst_reg_ptr is called when production inst_reg_ptr is entered.
func (s *BaseAsmE8Listener) EnterInst_reg_ptr(ctx *Inst_reg_ptrContext) {}

// ExitInst_reg_ptr is called when production inst_reg_ptr is exited.
func (s *BaseAsmE8Listener) ExitInst_reg_ptr(ctx *Inst_reg_ptrContext) {}

// EnterInst_reg_ptr_offset is called when production inst_reg_ptr_offset is entered.
func (s *BaseAsmE8Listener) EnterInst_reg_ptr_offset(ctx *Inst_reg_ptr_offsetContext) {}

// ExitInst_reg_ptr_offset is called when production inst_reg_ptr_offset is exited.
func (s *BaseAsmE8Listener) ExitInst_reg_ptr_offset(ctx *Inst_reg_ptr_offsetContext) {}

// EnterInst_ptr_offset_reg is called when production inst_ptr_offset_reg is entered.
func (s *BaseAsmE8Listener) EnterInst_ptr_offset_reg(ctx *Inst_ptr_offset_regContext) {}

// ExitInst_ptr_offset_reg is called when production inst_ptr_offset_reg is exited.
func (s *BaseAsmE8Listener) ExitInst_ptr_offset_reg(ctx *Inst_ptr_offset_regContext) {}

// EnterInst_single_reg is called when production inst_single_reg is entered.
func (s *BaseAsmE8Listener) EnterInst_single_reg(ctx *Inst_single_regContext) {}

// ExitInst_single_reg is called when production inst_single_reg is exited.
func (s *BaseAsmE8Listener) ExitInst_single_reg(ctx *Inst_single_regContext) {}

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

// EnterVariable is called when production variable is entered.
func (s *BaseAsmE8Listener) EnterVariable(ctx *VariableContext) {}

// ExitVariable is called when production variable is exited.
func (s *BaseAsmE8Listener) ExitVariable(ctx *VariableContext) {}

// EnterDirectives is called when production directives is entered.
func (s *BaseAsmE8Listener) EnterDirectives(ctx *DirectivesContext) {}

// ExitDirectives is called when production directives is exited.
func (s *BaseAsmE8Listener) ExitDirectives(ctx *DirectivesContext) {}

// EnterImm_list is called when production imm_list is entered.
func (s *BaseAsmE8Listener) EnterImm_list(ctx *Imm_listContext) {}

// ExitImm_list is called when production imm_list is exited.
func (s *BaseAsmE8Listener) ExitImm_list(ctx *Imm_listContext) {}

// EnterImm is called when production imm is entered.
func (s *BaseAsmE8Listener) EnterImm(ctx *ImmContext) {}

// ExitImm is called when production imm is exited.
func (s *BaseAsmE8Listener) ExitImm(ctx *ImmContext) {}

// EnterTag is called when production tag is entered.
func (s *BaseAsmE8Listener) EnterTag(ctx *TagContext) {}

// ExitTag is called when production tag is exited.
func (s *BaseAsmE8Listener) ExitTag(ctx *TagContext) {}
