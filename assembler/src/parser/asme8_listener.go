// Code generated from AsmE8.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // AsmE8

import "github.com/antlr4-go/antlr/v4"

// AsmE8Listener is a complete listener for a parse tree produced by AsmE8Parser.
type AsmE8Listener interface {
	antlr.ParseTreeListener

	// EnterProg is called when entering the prog production.
	EnterProg(c *ProgContext)

	// EnterInstruction is called when entering the instruction production.
	EnterInstruction(c *InstructionContext)

	// EnterLine is called when entering the line production.
	EnterLine(c *LineContext)

	// EnterLabel is called when entering the label production.
	EnterLabel(c *LabelContext)

	// EnterInst is called when entering the inst production.
	EnterInst(c *InstContext)

	// EnterInst_reg_reg is called when entering the inst_reg_reg production.
	EnterInst_reg_reg(c *Inst_reg_regContext)

	// EnterInst_reg_imm is called when entering the inst_reg_imm production.
	EnterInst_reg_imm(c *Inst_reg_immContext)

	// EnterInst_reg_imm_variable is called when entering the inst_reg_imm_variable production.
	EnterInst_reg_imm_variable(c *Inst_reg_imm_variableContext)

	// EnterInst_ptr_reg is called when entering the inst_ptr_reg production.
	EnterInst_ptr_reg(c *Inst_ptr_regContext)

	// EnterInst_reg_ptr is called when entering the inst_reg_ptr production.
	EnterInst_reg_ptr(c *Inst_reg_ptrContext)

	// EnterInst_reg_ptr_offset is called when entering the inst_reg_ptr_offset production.
	EnterInst_reg_ptr_offset(c *Inst_reg_ptr_offsetContext)

	// EnterInst_ptr_offset_reg is called when entering the inst_ptr_offset_reg production.
	EnterInst_ptr_offset_reg(c *Inst_ptr_offset_regContext)

	// EnterInst_single_reg is called when entering the inst_single_reg production.
	EnterInst_single_reg(c *Inst_single_regContext)

	// EnterInst_single_imm is called when entering the inst_single_imm production.
	EnterInst_single_imm(c *Inst_single_immContext)

	// EnterInst_single_tag is called when entering the inst_single_tag production.
	EnterInst_single_tag(c *Inst_single_tagContext)

	// EnterInst_single is called when entering the inst_single production.
	EnterInst_single(c *Inst_singleContext)

	// EnterMnemonic is called when entering the mnemonic production.
	EnterMnemonic(c *MnemonicContext)

	// EnterReg is called when entering the reg production.
	EnterReg(c *RegContext)

	// EnterPtr is called when entering the ptr production.
	EnterPtr(c *PtrContext)

	// EnterPtr_virtual_offset is called when entering the ptr_virtual_offset production.
	EnterPtr_virtual_offset(c *Ptr_virtual_offsetContext)

	// EnterPtr_offset is called when entering the ptr_offset production.
	EnterPtr_offset(c *Ptr_offsetContext)

	// EnterVariable is called when entering the variable production.
	EnterVariable(c *VariableContext)

	// EnterDirectives is called when entering the directives production.
	EnterDirectives(c *DirectivesContext)

	// EnterSegment is called when entering the segment production.
	EnterSegment(c *SegmentContext)

	// EnterAccess is called when entering the access production.
	EnterAccess(c *AccessContext)

	// EnterImm_list is called when entering the imm_list production.
	EnterImm_list(c *Imm_listContext)

	// EnterImm is called when entering the imm production.
	EnterImm(c *ImmContext)

	// EnterTag is called when entering the tag production.
	EnterTag(c *TagContext)

	// ExitProg is called when exiting the prog production.
	ExitProg(c *ProgContext)

	// ExitInstruction is called when exiting the instruction production.
	ExitInstruction(c *InstructionContext)

	// ExitLine is called when exiting the line production.
	ExitLine(c *LineContext)

	// ExitLabel is called when exiting the label production.
	ExitLabel(c *LabelContext)

	// ExitInst is called when exiting the inst production.
	ExitInst(c *InstContext)

	// ExitInst_reg_reg is called when exiting the inst_reg_reg production.
	ExitInst_reg_reg(c *Inst_reg_regContext)

	// ExitInst_reg_imm is called when exiting the inst_reg_imm production.
	ExitInst_reg_imm(c *Inst_reg_immContext)

	// ExitInst_reg_imm_variable is called when exiting the inst_reg_imm_variable production.
	ExitInst_reg_imm_variable(c *Inst_reg_imm_variableContext)

	// ExitInst_ptr_reg is called when exiting the inst_ptr_reg production.
	ExitInst_ptr_reg(c *Inst_ptr_regContext)

	// ExitInst_reg_ptr is called when exiting the inst_reg_ptr production.
	ExitInst_reg_ptr(c *Inst_reg_ptrContext)

	// ExitInst_reg_ptr_offset is called when exiting the inst_reg_ptr_offset production.
	ExitInst_reg_ptr_offset(c *Inst_reg_ptr_offsetContext)

	// ExitInst_ptr_offset_reg is called when exiting the inst_ptr_offset_reg production.
	ExitInst_ptr_offset_reg(c *Inst_ptr_offset_regContext)

	// ExitInst_single_reg is called when exiting the inst_single_reg production.
	ExitInst_single_reg(c *Inst_single_regContext)

	// ExitInst_single_imm is called when exiting the inst_single_imm production.
	ExitInst_single_imm(c *Inst_single_immContext)

	// ExitInst_single_tag is called when exiting the inst_single_tag production.
	ExitInst_single_tag(c *Inst_single_tagContext)

	// ExitInst_single is called when exiting the inst_single production.
	ExitInst_single(c *Inst_singleContext)

	// ExitMnemonic is called when exiting the mnemonic production.
	ExitMnemonic(c *MnemonicContext)

	// ExitReg is called when exiting the reg production.
	ExitReg(c *RegContext)

	// ExitPtr is called when exiting the ptr production.
	ExitPtr(c *PtrContext)

	// ExitPtr_virtual_offset is called when exiting the ptr_virtual_offset production.
	ExitPtr_virtual_offset(c *Ptr_virtual_offsetContext)

	// ExitPtr_offset is called when exiting the ptr_offset production.
	ExitPtr_offset(c *Ptr_offsetContext)

	// ExitVariable is called when exiting the variable production.
	ExitVariable(c *VariableContext)

	// ExitDirectives is called when exiting the directives production.
	ExitDirectives(c *DirectivesContext)

	// ExitSegment is called when exiting the segment production.
	ExitSegment(c *SegmentContext)

	// ExitAccess is called when exiting the access production.
	ExitAccess(c *AccessContext)

	// ExitImm_list is called when exiting the imm_list production.
	ExitImm_list(c *Imm_listContext)

	// ExitImm is called when exiting the imm production.
	ExitImm(c *ImmContext)

	// ExitTag is called when exiting the tag production.
	ExitTag(c *TagContext)
}
