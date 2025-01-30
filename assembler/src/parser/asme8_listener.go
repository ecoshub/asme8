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

	// EnterInst_stack_imm is called when entering the inst_stack_imm production.
	EnterInst_stack_imm(c *Inst_stack_immContext)

	// EnterInst_index_imm is called when entering the inst_index_imm production.
	EnterInst_index_imm(c *Inst_index_immContext)

	// EnterInst_reg_imm_variable is called when entering the inst_reg_imm_variable production.
	EnterInst_reg_imm_variable(c *Inst_reg_imm_variableContext)

	// EnterInst_ptr_reg is called when entering the inst_ptr_reg production.
	EnterInst_ptr_reg(c *Inst_ptr_regContext)

	// EnterInst_reg_ptr is called when entering the inst_reg_ptr production.
	EnterInst_reg_ptr(c *Inst_reg_ptrContext)

	// EnterInst_reg_ptr_offset is called when entering the inst_reg_ptr_offset production.
	EnterInst_reg_ptr_offset(c *Inst_reg_ptr_offsetContext)

	// EnterInst_indirect_reg_stack is called when entering the inst_indirect_reg_stack production.
	EnterInst_indirect_reg_stack(c *Inst_indirect_reg_stackContext)

	// EnterInst_indirect_stack_register is called when entering the inst_indirect_stack_register production.
	EnterInst_indirect_stack_register(c *Inst_indirect_stack_registerContext)

	// EnterInst_indirect_reg_index is called when entering the inst_indirect_reg_index production.
	EnterInst_indirect_reg_index(c *Inst_indirect_reg_indexContext)

	// EnterInst_indirect_index_register is called when entering the inst_indirect_index_register production.
	EnterInst_indirect_index_register(c *Inst_indirect_index_registerContext)

	// EnterInst_ptr_offset_reg is called when entering the inst_ptr_offset_reg production.
	EnterInst_ptr_offset_reg(c *Inst_ptr_offset_regContext)

	// EnterInst_single_reg is called when entering the inst_single_reg production.
	EnterInst_single_reg(c *Inst_single_regContext)

	// EnterInst_implied_stack is called when entering the inst_implied_stack production.
	EnterInst_implied_stack(c *Inst_implied_stackContext)

	// EnterInst_implied_index is called when entering the inst_implied_index production.
	EnterInst_implied_index(c *Inst_implied_indexContext)

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

	// EnterStack is called when entering the stack production.
	EnterStack(c *StackContext)

	// EnterIndex is called when entering the index production.
	EnterIndex(c *IndexContext)

	// EnterPtr is called when entering the ptr production.
	EnterPtr(c *PtrContext)

	// EnterPtr_virtual_offset is called when entering the ptr_virtual_offset production.
	EnterPtr_virtual_offset(c *Ptr_virtual_offsetContext)

	// EnterPtr_offset is called when entering the ptr_offset production.
	EnterPtr_offset(c *Ptr_offsetContext)

	// EnterStack_offset is called when entering the stack_offset production.
	EnterStack_offset(c *Stack_offsetContext)

	// EnterIndex_offset is called when entering the index_offset production.
	EnterIndex_offset(c *Index_offsetContext)

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

	// ExitInst_stack_imm is called when exiting the inst_stack_imm production.
	ExitInst_stack_imm(c *Inst_stack_immContext)

	// ExitInst_index_imm is called when exiting the inst_index_imm production.
	ExitInst_index_imm(c *Inst_index_immContext)

	// ExitInst_reg_imm_variable is called when exiting the inst_reg_imm_variable production.
	ExitInst_reg_imm_variable(c *Inst_reg_imm_variableContext)

	// ExitInst_ptr_reg is called when exiting the inst_ptr_reg production.
	ExitInst_ptr_reg(c *Inst_ptr_regContext)

	// ExitInst_reg_ptr is called when exiting the inst_reg_ptr production.
	ExitInst_reg_ptr(c *Inst_reg_ptrContext)

	// ExitInst_reg_ptr_offset is called when exiting the inst_reg_ptr_offset production.
	ExitInst_reg_ptr_offset(c *Inst_reg_ptr_offsetContext)

	// ExitInst_indirect_reg_stack is called when exiting the inst_indirect_reg_stack production.
	ExitInst_indirect_reg_stack(c *Inst_indirect_reg_stackContext)

	// ExitInst_indirect_stack_register is called when exiting the inst_indirect_stack_register production.
	ExitInst_indirect_stack_register(c *Inst_indirect_stack_registerContext)

	// ExitInst_indirect_reg_index is called when exiting the inst_indirect_reg_index production.
	ExitInst_indirect_reg_index(c *Inst_indirect_reg_indexContext)

	// ExitInst_indirect_index_register is called when exiting the inst_indirect_index_register production.
	ExitInst_indirect_index_register(c *Inst_indirect_index_registerContext)

	// ExitInst_ptr_offset_reg is called when exiting the inst_ptr_offset_reg production.
	ExitInst_ptr_offset_reg(c *Inst_ptr_offset_regContext)

	// ExitInst_single_reg is called when exiting the inst_single_reg production.
	ExitInst_single_reg(c *Inst_single_regContext)

	// ExitInst_implied_stack is called when exiting the inst_implied_stack production.
	ExitInst_implied_stack(c *Inst_implied_stackContext)

	// ExitInst_implied_index is called when exiting the inst_implied_index production.
	ExitInst_implied_index(c *Inst_implied_indexContext)

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

	// ExitStack is called when exiting the stack production.
	ExitStack(c *StackContext)

	// ExitIndex is called when exiting the index production.
	ExitIndex(c *IndexContext)

	// ExitPtr is called when exiting the ptr production.
	ExitPtr(c *PtrContext)

	// ExitPtr_virtual_offset is called when exiting the ptr_virtual_offset production.
	ExitPtr_virtual_offset(c *Ptr_virtual_offsetContext)

	// ExitPtr_offset is called when exiting the ptr_offset production.
	ExitPtr_offset(c *Ptr_offsetContext)

	// ExitStack_offset is called when exiting the stack_offset production.
	ExitStack_offset(c *Stack_offsetContext)

	// ExitIndex_offset is called when exiting the index_offset production.
	ExitIndex_offset(c *Index_offsetContext)

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
