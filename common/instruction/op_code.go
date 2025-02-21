package instruction

var (
	_MNEMONIC_TO_OPCODE = map[string]map[uint8]uint8{
		INST_MNEMONICS_MOV: {
			// mov b, a
			ADDR_MODE_REG8_REG8: INST_MOV_REG8_REG8,
			// mov a, 4
			ADDR_MODE_REG8_IMM8: INST_MOV_REG8_IMM8,
			// mov a, [8000]
			ADDR_MODE_MEM_TO_REG_DIRECT: INST_MOV_MEM_TO_REG_DIRECT,
			// mov [8000], a
			ADDR_MODE_REG_TO_MEM_DIRECT: INST_MOV_REG_TO_MEM_DIRECT,
			// mov a, [8000+b]
			ADDR_MODE_MEM_TO_REG_INDEXED: INST_MOV_MEM_TO_REG_INDEXED,
			// mov [8000+b], a
			ADDR_MODE_REG_TO_MEM_INDEXED: INST_MOV_REG_TO_MEM_INDEXED,
			// mov a, [ip]
			ADDR_MODE_MEM_TO_REG_INDIRECT: INST_MOV_MEM_TO_REG_INDIRECT,
			// mov [ip], a
			ADDR_MODE_REG_TO_MEM_INDIRECT: INST_MOV_REG_TO_MEM_INDIRECT,
			// mov a, [ip+4]
			ADDR_MODE_MEM_TO_REG_INDIRECT_OFFSET: INST_MOV_MEM_TO_REG_INDIRECT_OFFSET,
			// mov [ip+4], a
			ADDR_MODE_REG_TO_MEM_INDIRECT_OFFSET: INST_MOV_REG_TO_MEM_INDIRECT_OFFSET,
			// mov a, [ip+b]
			ADDR_MODE_MEM_TO_REG_REG16_INDEXED: INST_MOV_MEM_TO_REG_REG16_INDEXED,
			// mov [ip+b], a
			ADDR_MODE_REG_TO_MEM_REG16_INDEXED: INST_MOV_REG_TO_MEM_REG16_INDEXED,
			// mov ip, sp
			ADDR_MODE_REG16_STACK: INST_MOV_SP_REG16,
			// mov sp, bp
			ADDR_MODE_STACK_REG16: INST_MOV_REG16_SP,
			// mov bp, ip
			ADDR_MODE_REG16_REG16: INST_MOV_REG16_REG16,
			// mov bp, 0x1234
			ADDR_MODE_REG16_IMM16: INST_MOV_REG16_IMM16,
		},
		INST_MNEMONICS_XOR: {
			// xor b, a
			ADDR_MODE_REG8_REG8: INST_XOR_REG8_REG8,
			// xor b, 0x10
			ADDR_MODE_REG8_IMM8: INST_XOR_REG8_IMM8,
		},
		INST_MNEMONICS_AND: {
			// and b, a
			ADDR_MODE_REG8_REG8: INST_AND_REG8_REG8,
			// and b, 0x10
			ADDR_MODE_REG8_IMM8: INST_AND_REG8_IMM8,
		},
		INST_MNEMONICS_OR: {
			// or b, a
			ADDR_MODE_REG8_REG8: INST_OR_REG8_REG8,
			// or b, 0x10
			ADDR_MODE_REG8_IMM8: INST_OR_REG8_IMM8,
		},
		INST_MNEMONICS_NOT: {
			// not a
			ADDR_MODE_IMPL_REG8: INST_NOT_REG8,
		},
		INST_MNEMONICS_SHL: {
			// shl a
			ADDR_MODE_IMPL_REG8: INST_SHL_REG8,
		},
		INST_MNEMONICS_SHR: {
			// shr a
			ADDR_MODE_IMPL_REG8: INST_SHR_REG8,
		},
		INST_MNEMONICS_ROL: {
			// rol a
			ADDR_MODE_IMPL_REG8: INST_ROL_REG8,
		},
		INST_MNEMONICS_ROR: {
			// ror a
			ADDR_MODE_IMPL_REG8: INST_ROR_REG8,
		},
		INST_MNEMONICS_CMP: {
			// cmp b, a
			ADDR_MODE_REG8_REG8: INST_CMP_REG8_REG8,
			// cmp a, 4
			ADDR_MODE_REG8_IMM8: INST_CMP_REG8_IMM8,
		},
		INST_MNEMONICS_ADD: {
			// add b, a
			ADDR_MODE_REG8_REG8: INST_ADD_REG8_REG8,
			// add a, 4
			ADDR_MODE_REG8_IMM8: INST_ADD_REG8_IMM8,
			// add sp, 4
			ADDR_MODE_STACK_IMM8: INST_ADD_SP_IMM8,
			// add bp, 0x10
			ADDR_MODE_REG16_IMM8: INST_ADD_REG16_IMM8,
		},
		INST_MNEMONICS_ADC: {
			// adc b, a
			ADDR_MODE_REG8_REG8: INST_ADC_REG8_REG8,
			// adc a, 4
			ADDR_MODE_REG8_IMM8: INST_ADC_REG8_IMM8,
		},
		INST_MNEMONICS_SUB: {
			// sub b, a
			ADDR_MODE_REG8_REG8: INST_SUB_REG8_REG8,
			// sub a, 4
			ADDR_MODE_REG8_IMM8: INST_SUB_REG8_IMM8,
			// sub sp, 4
			ADDR_MODE_STACK_IMM8: INST_SUB_SP_IMM8,
			// sub bp, 0x10
			ADDR_MODE_REG16_IMM8: INST_SUB_REG16_IMM8,
		},
		INST_MNEMONICS_SBB: {
			// sbb b, a
			ADDR_MODE_REG8_REG8: INST_SBB_REG8_REG8,
			// sbb a, 4
			ADDR_MODE_REG8_IMM8: INST_SBB_REG8_IMM8,
		},
		INST_MNEMONICS_INC: {
			// inc a
			ADDR_MODE_IMPL_REG8: INST_INC_REG8,
			// inc bp
			ADDR_MODE_IMPL_REG16: INST_INC_REG16,
		},
		INST_MNEMONICS_DEC: {
			// dec a
			ADDR_MODE_IMPL_REG8: INST_DEC_REG8,
			// dec bp
			ADDR_MODE_IMPL_REG16: INST_DEC_REG16,
		},
		INST_MNEMONICS_JMP: {
			// jmp 0x5500
			ADDR_MODE_IMPL_IMM16: INST_JMP_IMPL_IMM16,
		},
		INST_MNEMONICS_JS: {
			// js 0x6000
			ADDR_MODE_IMPL_IMM16: INST_JS_IMPL_IMM16,
		},
		INST_MNEMONICS_JNS: {
			// jns 0x6000
			ADDR_MODE_IMPL_IMM16: INST_JNS_IMPL_IMM16,
		},
		INST_MNEMONICS_JZ: {
			// jz 0x6000
			ADDR_MODE_IMPL_IMM16: INST_JZ_IMPL_IMM16,
		},
		INST_MNEMONICS_JNZ: {
			// jnz 0x6000
			ADDR_MODE_IMPL_IMM16: INST_JNZ_IMPL_IMM16,
		},
		INST_MNEMONICS_JC: {
			// jc 0x6000
			ADDR_MODE_IMPL_IMM16: INST_JC_IMPL_IMM16,
		},
		INST_MNEMONICS_JNC: {
			// jc 0x6000
			ADDR_MODE_IMPL_IMM16: INST_JNC_IMPL_IMM16,
		},
		INST_MNEMONICS_PUSH: {
			// push a
			ADDR_MODE_IMPL_REG8: INST_PUSH_REG8,
			// push ip
			ADDR_MODE_IMPL_REG16: INST_PUSH_REG16,
			// push sp
			ADDR_MODE_IMPL_STACK: INST_PUSH_SP,
			// push sr
			ADDR_MODE_IMPL_STATUS_REG: INST_PUSH_SR,
		},
		INST_MNEMONICS_POP: {
			// pop a
			ADDR_MODE_IMPL_REG8: INST_POP_REG8,
			// pop ip
			ADDR_MODE_IMPL_REG16: INST_POP_REG16,
			// pop sp
			ADDR_MODE_IMPL_STACK: INST_POP_SP,
			// pop sr
			ADDR_MODE_IMPL_STATUS_REG: INST_POP_SR,
		},
		INST_MNEMONICS_CALL: {
			// call ip
			ADDR_MODE_IMPL_REG16: INST_CALL_REG16,
			// call 0x1234
			ADDR_MODE_IMPL_IMM16: INST_CALL_IMM16,
		},
		INST_MNEMONICS_RET: {
			// ret
			ADDR_MODE_IMPL: INST_RET_IMPL,
		},
		INST_MNEMONICS_RTI: {
			// rti
			ADDR_MODE_IMPL: INST_RTI_IMPL,
		},
		INST_MNEMONICS_CLC: {
			// clc
			ADDR_MODE_IMPL: INST_CLC_IMPL,
		},
		INST_MNEMONICS_BRK: {
			// brk
			ADDR_MODE_IMPL: INST_BRK_IMPL,
		},
		INST_MNEMONICS_NOP: {
			// nop
			ADDR_MODE_IMPL: INST_NOP_IMPL,
		},
	}
)

func GetOpCode(inst string, addressingMode uint8) (uint8, bool) {
	modes, ok := _MNEMONIC_TO_OPCODE[inst]
	if !ok {
		return 0, false
	}
	op, ok := modes[addressingMode]
	if !ok {
		return 0, false
	}
	return op, true
}
