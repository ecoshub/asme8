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

	// EnterAddr_mode_imm16 is called when entering the addr_mode_imm16 production.
	EnterAddr_mode_imm16(c *Addr_mode_imm16Context)

	// EnterAddr_mode_imm16_tag is called when entering the addr_mode_imm16_tag production.
	EnterAddr_mode_imm16_tag(c *Addr_mode_imm16_tagContext)

	// EnterAddr_mode_reg_to_mem_indexed is called when entering the addr_mode_reg_to_mem_indexed production.
	EnterAddr_mode_reg_to_mem_indexed(c *Addr_mode_reg_to_mem_indexedContext)

	// EnterAddr_mode_reg_to_mem_direct is called when entering the addr_mode_reg_to_mem_direct production.
	EnterAddr_mode_reg_to_mem_direct(c *Addr_mode_reg_to_mem_directContext)

	// EnterAddr_mode_reg_to_mem_indirect_offset is called when entering the addr_mode_reg_to_mem_indirect_offset production.
	EnterAddr_mode_reg_to_mem_indirect_offset(c *Addr_mode_reg_to_mem_indirect_offsetContext)

	// EnterAddr_mode_reg_to_mem_reg16_indexed is called when entering the addr_mode_reg_to_mem_reg16_indexed production.
	EnterAddr_mode_reg_to_mem_reg16_indexed(c *Addr_mode_reg_to_mem_reg16_indexedContext)

	// EnterAddr_mode_reg_to_mem_indirect is called when entering the addr_mode_reg_to_mem_indirect production.
	EnterAddr_mode_reg_to_mem_indirect(c *Addr_mode_reg_to_mem_indirectContext)

	// EnterAddr_mode_implied_reg8 is called when entering the addr_mode_implied_reg8 production.
	EnterAddr_mode_implied_reg8(c *Addr_mode_implied_reg8Context)

	// EnterAddr_mode_reg8_imm8 is called when entering the addr_mode_reg8_imm8 production.
	EnterAddr_mode_reg8_imm8(c *Addr_mode_reg8_imm8Context)

	// EnterAddr_mode_reg8_imm8_tag is called when entering the addr_mode_reg8_imm8_tag production.
	EnterAddr_mode_reg8_imm8_tag(c *Addr_mode_reg8_imm8_tagContext)

	// EnterAddr_mode_mem_to_reg_indexed is called when entering the addr_mode_mem_to_reg_indexed production.
	EnterAddr_mode_mem_to_reg_indexed(c *Addr_mode_mem_to_reg_indexedContext)

	// EnterAddr_mode_mem_to_reg_direct is called when entering the addr_mode_mem_to_reg_direct production.
	EnterAddr_mode_mem_to_reg_direct(c *Addr_mode_mem_to_reg_directContext)

	// EnterAddr_mode_mem_to_reg_indirect_offset is called when entering the addr_mode_mem_to_reg_indirect_offset production.
	EnterAddr_mode_mem_to_reg_indirect_offset(c *Addr_mode_mem_to_reg_indirect_offsetContext)

	// EnterAddr_mode_mem_to_reg_reg16_indexed is called when entering the addr_mode_mem_to_reg_reg16_indexed production.
	EnterAddr_mode_mem_to_reg_reg16_indexed(c *Addr_mode_mem_to_reg_reg16_indexedContext)

	// EnterAddr_mode_mem_to_reg_indirect is called when entering the addr_mode_mem_to_reg_indirect production.
	EnterAddr_mode_mem_to_reg_indirect(c *Addr_mode_mem_to_reg_indirectContext)

	// EnterAddr_mode_reg8_reg8 is called when entering the addr_mode_reg8_reg8 production.
	EnterAddr_mode_reg8_reg8(c *Addr_mode_reg8_reg8Context)

	// EnterAddr_mode_implied is called when entering the addr_mode_implied production.
	EnterAddr_mode_implied(c *Addr_mode_impliedContext)

	// EnterAddr_mode_implied_reg16 is called when entering the addr_mode_implied_reg16 production.
	EnterAddr_mode_implied_reg16(c *Addr_mode_implied_reg16Context)

	// EnterAddr_mode_reg16_imm is called when entering the addr_mode_reg16_imm production.
	EnterAddr_mode_reg16_imm(c *Addr_mode_reg16_immContext)

	// EnterAddr_mode_reg16_imm_tag is called when entering the addr_mode_reg16_imm_tag production.
	EnterAddr_mode_reg16_imm_tag(c *Addr_mode_reg16_imm_tagContext)

	// EnterAddr_mode_reg16_reg16 is called when entering the addr_mode_reg16_reg16 production.
	EnterAddr_mode_reg16_reg16(c *Addr_mode_reg16_reg16Context)

	// EnterAddr_mode_reg16_stack is called when entering the addr_mode_reg16_stack production.
	EnterAddr_mode_reg16_stack(c *Addr_mode_reg16_stackContext)

	// EnterAddr_mode_implied_stack is called when entering the addr_mode_implied_stack production.
	EnterAddr_mode_implied_stack(c *Addr_mode_implied_stackContext)

	// EnterAddr_mode_implied_status_register is called when entering the addr_mode_implied_status_register production.
	EnterAddr_mode_implied_status_register(c *Addr_mode_implied_status_registerContext)

	// EnterAddr_mode_stack_imm8 is called when entering the addr_mode_stack_imm8 production.
	EnterAddr_mode_stack_imm8(c *Addr_mode_stack_imm8Context)

	// EnterAddr_mode_stack_imm8_tag is called when entering the addr_mode_stack_imm8_tag production.
	EnterAddr_mode_stack_imm8_tag(c *Addr_mode_stack_imm8_tagContext)

	// EnterAddr_mode_stack_reg16 is called when entering the addr_mode_stack_reg16 production.
	EnterAddr_mode_stack_reg16(c *Addr_mode_stack_reg16Context)

	// EnterMnemonic is called when entering the mnemonic production.
	EnterMnemonic(c *MnemonicContext)

	// EnterReg8 is called when entering the reg8 production.
	EnterReg8(c *Reg8Context)

	// EnterReg16 is called when entering the reg16 production.
	EnterReg16(c *Reg16Context)

	// EnterStack is called when entering the stack production.
	EnterStack(c *StackContext)

	// EnterStatus_register is called when entering the status_register production.
	EnterStatus_register(c *Status_registerContext)

	// EnterDirect is called when entering the direct production.
	EnterDirect(c *DirectContext)

	// EnterDirect_virtual_offset is called when entering the direct_virtual_offset production.
	EnterDirect_virtual_offset(c *Direct_virtual_offsetContext)

	// EnterDirect_offset is called when entering the direct_offset production.
	EnterDirect_offset(c *Direct_offsetContext)

	// EnterIndirect_offset is called when entering the indirect_offset production.
	EnterIndirect_offset(c *Indirect_offsetContext)

	// EnterReg16_indexed is called when entering the reg16_indexed production.
	EnterReg16_indexed(c *Reg16_indexedContext)

	// EnterIndirect is called when entering the indirect production.
	EnterIndirect(c *IndirectContext)

	// EnterVariable is called when entering the variable production.
	EnterVariable(c *VariableContext)

	// EnterVariable_reference is called when entering the variable_reference production.
	EnterVariable_reference(c *Variable_referenceContext)

	// EnterReference is called when entering the reference production.
	EnterReference(c *ReferenceContext)

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

	// EnterAscii is called when entering the ascii production.
	EnterAscii(c *AsciiContext)

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

	// ExitAddr_mode_imm16 is called when exiting the addr_mode_imm16 production.
	ExitAddr_mode_imm16(c *Addr_mode_imm16Context)

	// ExitAddr_mode_imm16_tag is called when exiting the addr_mode_imm16_tag production.
	ExitAddr_mode_imm16_tag(c *Addr_mode_imm16_tagContext)

	// ExitAddr_mode_reg_to_mem_indexed is called when exiting the addr_mode_reg_to_mem_indexed production.
	ExitAddr_mode_reg_to_mem_indexed(c *Addr_mode_reg_to_mem_indexedContext)

	// ExitAddr_mode_reg_to_mem_direct is called when exiting the addr_mode_reg_to_mem_direct production.
	ExitAddr_mode_reg_to_mem_direct(c *Addr_mode_reg_to_mem_directContext)

	// ExitAddr_mode_reg_to_mem_indirect_offset is called when exiting the addr_mode_reg_to_mem_indirect_offset production.
	ExitAddr_mode_reg_to_mem_indirect_offset(c *Addr_mode_reg_to_mem_indirect_offsetContext)

	// ExitAddr_mode_reg_to_mem_reg16_indexed is called when exiting the addr_mode_reg_to_mem_reg16_indexed production.
	ExitAddr_mode_reg_to_mem_reg16_indexed(c *Addr_mode_reg_to_mem_reg16_indexedContext)

	// ExitAddr_mode_reg_to_mem_indirect is called when exiting the addr_mode_reg_to_mem_indirect production.
	ExitAddr_mode_reg_to_mem_indirect(c *Addr_mode_reg_to_mem_indirectContext)

	// ExitAddr_mode_implied_reg8 is called when exiting the addr_mode_implied_reg8 production.
	ExitAddr_mode_implied_reg8(c *Addr_mode_implied_reg8Context)

	// ExitAddr_mode_reg8_imm8 is called when exiting the addr_mode_reg8_imm8 production.
	ExitAddr_mode_reg8_imm8(c *Addr_mode_reg8_imm8Context)

	// ExitAddr_mode_reg8_imm8_tag is called when exiting the addr_mode_reg8_imm8_tag production.
	ExitAddr_mode_reg8_imm8_tag(c *Addr_mode_reg8_imm8_tagContext)

	// ExitAddr_mode_mem_to_reg_indexed is called when exiting the addr_mode_mem_to_reg_indexed production.
	ExitAddr_mode_mem_to_reg_indexed(c *Addr_mode_mem_to_reg_indexedContext)

	// ExitAddr_mode_mem_to_reg_direct is called when exiting the addr_mode_mem_to_reg_direct production.
	ExitAddr_mode_mem_to_reg_direct(c *Addr_mode_mem_to_reg_directContext)

	// ExitAddr_mode_mem_to_reg_indirect_offset is called when exiting the addr_mode_mem_to_reg_indirect_offset production.
	ExitAddr_mode_mem_to_reg_indirect_offset(c *Addr_mode_mem_to_reg_indirect_offsetContext)

	// ExitAddr_mode_mem_to_reg_reg16_indexed is called when exiting the addr_mode_mem_to_reg_reg16_indexed production.
	ExitAddr_mode_mem_to_reg_reg16_indexed(c *Addr_mode_mem_to_reg_reg16_indexedContext)

	// ExitAddr_mode_mem_to_reg_indirect is called when exiting the addr_mode_mem_to_reg_indirect production.
	ExitAddr_mode_mem_to_reg_indirect(c *Addr_mode_mem_to_reg_indirectContext)

	// ExitAddr_mode_reg8_reg8 is called when exiting the addr_mode_reg8_reg8 production.
	ExitAddr_mode_reg8_reg8(c *Addr_mode_reg8_reg8Context)

	// ExitAddr_mode_implied is called when exiting the addr_mode_implied production.
	ExitAddr_mode_implied(c *Addr_mode_impliedContext)

	// ExitAddr_mode_implied_reg16 is called when exiting the addr_mode_implied_reg16 production.
	ExitAddr_mode_implied_reg16(c *Addr_mode_implied_reg16Context)

	// ExitAddr_mode_reg16_imm is called when exiting the addr_mode_reg16_imm production.
	ExitAddr_mode_reg16_imm(c *Addr_mode_reg16_immContext)

	// ExitAddr_mode_reg16_imm_tag is called when exiting the addr_mode_reg16_imm_tag production.
	ExitAddr_mode_reg16_imm_tag(c *Addr_mode_reg16_imm_tagContext)

	// ExitAddr_mode_reg16_reg16 is called when exiting the addr_mode_reg16_reg16 production.
	ExitAddr_mode_reg16_reg16(c *Addr_mode_reg16_reg16Context)

	// ExitAddr_mode_reg16_stack is called when exiting the addr_mode_reg16_stack production.
	ExitAddr_mode_reg16_stack(c *Addr_mode_reg16_stackContext)

	// ExitAddr_mode_implied_stack is called when exiting the addr_mode_implied_stack production.
	ExitAddr_mode_implied_stack(c *Addr_mode_implied_stackContext)

	// ExitAddr_mode_implied_status_register is called when exiting the addr_mode_implied_status_register production.
	ExitAddr_mode_implied_status_register(c *Addr_mode_implied_status_registerContext)

	// ExitAddr_mode_stack_imm8 is called when exiting the addr_mode_stack_imm8 production.
	ExitAddr_mode_stack_imm8(c *Addr_mode_stack_imm8Context)

	// ExitAddr_mode_stack_imm8_tag is called when exiting the addr_mode_stack_imm8_tag production.
	ExitAddr_mode_stack_imm8_tag(c *Addr_mode_stack_imm8_tagContext)

	// ExitAddr_mode_stack_reg16 is called when exiting the addr_mode_stack_reg16 production.
	ExitAddr_mode_stack_reg16(c *Addr_mode_stack_reg16Context)

	// ExitMnemonic is called when exiting the mnemonic production.
	ExitMnemonic(c *MnemonicContext)

	// ExitReg8 is called when exiting the reg8 production.
	ExitReg8(c *Reg8Context)

	// ExitReg16 is called when exiting the reg16 production.
	ExitReg16(c *Reg16Context)

	// ExitStack is called when exiting the stack production.
	ExitStack(c *StackContext)

	// ExitStatus_register is called when exiting the status_register production.
	ExitStatus_register(c *Status_registerContext)

	// ExitDirect is called when exiting the direct production.
	ExitDirect(c *DirectContext)

	// ExitDirect_virtual_offset is called when exiting the direct_virtual_offset production.
	ExitDirect_virtual_offset(c *Direct_virtual_offsetContext)

	// ExitDirect_offset is called when exiting the direct_offset production.
	ExitDirect_offset(c *Direct_offsetContext)

	// ExitIndirect_offset is called when exiting the indirect_offset production.
	ExitIndirect_offset(c *Indirect_offsetContext)

	// ExitReg16_indexed is called when exiting the reg16_indexed production.
	ExitReg16_indexed(c *Reg16_indexedContext)

	// ExitIndirect is called when exiting the indirect production.
	ExitIndirect(c *IndirectContext)

	// ExitVariable is called when exiting the variable production.
	ExitVariable(c *VariableContext)

	// ExitVariable_reference is called when exiting the variable_reference production.
	ExitVariable_reference(c *Variable_referenceContext)

	// ExitReference is called when exiting the reference production.
	ExitReference(c *ReferenceContext)

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

	// ExitAscii is called when exiting the ascii production.
	ExitAscii(c *AsciiContext)

	// ExitTag is called when exiting the tag production.
	ExitTag(c *TagContext)
}
