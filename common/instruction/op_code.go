package instruction

var (
	_MNEMONIC_TO_OPCODE = map[string]map[uint8]uint8{
		INST_MNEMONICS_BRK: {
			// brk
			ADDRESSING_MODE_IMPL: 0,
		},
		INST_MNEMONICS_MOV: {
			// mov b, a
			ADDRESSING_MODE_REG_REG: INST_MOV_REG_REG,
			// mov a, 4
			ADDRESSING_MODE_REG_IMM: INST_MOV_REG_IMM,
			// mov a, [8000]
			ADDRESSING_MODE_MEM_REG: INST_MOV_MEM_REG,
			// mov [8000], a
			ADDRESSING_MODE_REG_MEM: INST_MOV_REG_MEM,
			// mov a, [8000+b]
			ADDRESSING_MODE_MEM_REG_OFFSET: INST_MOV_MEM_REG_OFFSET,
			// mov [8000+b], a
			ADDRESSING_MODE_REG_MEM_OFFSET: INST_MOV_REG_MEM_OFFSET,
			// mov a, [sp]
			ADDRESSING_MODE_REG_SP: INST_MOV_REG_SP,
			// mov [sp], a
			ADDRESSING_MODE_SP_REG: INST_MOV_SP_REG,
			// mov a, [sp+4]
			ADDRESSING_MODE_REG_SP_OFFSET: INST_MOV_REG_SP_OFFSET,
			// mov [sp+4], a
			ADDRESSING_MODE_SP_REG_OFFSET: INST_MOV_SP_REG_OFFSET,
			// mov a, [sp+b]
			ADDRESSING_MODE_REG_SP_OFFSET_REG: INST_MOV_REG_SP_OFFSET_REG,
			// mov [sp+b], a
			ADDRESSING_MODE_SP_REG_OFFSET_REG: INST_MOV_SP_REG_OFFSET_REG,
			// mov a, [ip]
			ADDRESSING_MODE_REG_IP: INST_MOV_REG_IP,
			// mov [ip], a
			ADDRESSING_MODE_IP_REG: INST_MOV_IP_REG,
			// mov a, [ip+4]
			ADDRESSING_MODE_REG_IP_OFFSET: INST_MOV_REG_IP_OFFSET,
			// mov [ip+4], a
			ADDRESSING_MODE_IP_REG_OFFSET: INST_MOV_IP_REG_OFFSET,
			// mov a, [ip+b]
			ADDRESSING_MODE_REG_IP_OFFSET_REG: INST_MOV_REG_IP_OFFSET_REG,
			// mov [ip+b], a
			ADDRESSING_MODE_IP_REG_OFFSET_REG: INST_MOV_IP_REG_OFFSET_REG,
		},
		INST_MNEMONICS_XOR: {
			// xor b, a
			ADDRESSING_MODE_REG_REG: INST_XOR_REG_REG,
			// xor b, 0x10
			ADDRESSING_MODE_REG_IMM: INST_XOR_REG_IMM,
		},
		INST_MNEMONICS_AND: {
			// and b, a
			ADDRESSING_MODE_REG_REG: INST_AND_REG_REG,
			// and b, 0x10
			ADDRESSING_MODE_REG_IMM: INST_AND_REG_IMM,
		},
		INST_MNEMONICS_OR: {
			// or b, a
			ADDRESSING_MODE_REG_REG: INST_OR_REG_REG,
			// or b, 0x10
			ADDRESSING_MODE_REG_IMM: INST_OR_REG_IMM,
		},
		INST_MNEMONICS_NOT: {
			// not a
			ADDRESSING_MODE_IMPL_REG: INST_NOT_REG,
		},
		INST_MNEMONICS_SHL: {
			// shl a
			ADDRESSING_MODE_IMPL_REG: INST_SHL_REG,
		},
		INST_MNEMONICS_SHR: {
			// shr a
			ADDRESSING_MODE_IMPL_REG: INST_SHR_REG,
		},
		INST_MNEMONICS_ROL: {
			// rol a
			ADDRESSING_MODE_IMPL_REG: INST_ROL_REG,
		},
		INST_MNEMONICS_ROR: {
			// ror a
			ADDRESSING_MODE_IMPL_REG: INST_ROR_REG,
		},
		INST_MNEMONICS_CMP: {
			// cmp b, a
			ADDRESSING_MODE_REG_REG: INST_CMP_REG_REG,
			// cmp a, 4
			ADDRESSING_MODE_REG_IMM: INST_CMP_REG_IMM,
		},
		INST_MNEMONICS_ADC: {
			// adc b, a
			ADDRESSING_MODE_REG_REG: INST_ADC_REG_REG,
			// adc a, 4
			ADDRESSING_MODE_REG_IMM: INST_ADC_REG_IMM,
			// adc [8000], 2
			ADDRESSING_MODE_PTR_IMM: INST_ADC_PTR_IMM,
		},
		INST_MNEMONICS_SBB: {
			// sbb b, a
			ADDRESSING_MODE_REG_REG: INST_SBB_REG_REG,
			// sbb a, 4
			ADDRESSING_MODE_REG_IMM: INST_SBB_REG_IMM,
		},
		INST_MNEMONICS_ADD: {
			// add b, a
			ADDRESSING_MODE_REG_REG: INST_ADD_REG_REG,
			// add a, 4
			ADDRESSING_MODE_REG_IMM: INST_ADD_REG_IMM,
			// add [8000], a
			ADDRESSING_MODE_REG_MEM: INST_ADD_REG_MEM,
			// add sp, 4
			ADDRESSING_MODE_SP_IMM: INST_ADD_SP_IMM,
			// add ip, 4
			ADDRESSING_MODE_IP_IMM: INST_ADD_IP_IMM,
			// add [8000], 2
			ADDRESSING_MODE_PTR_IMM: INST_ADD_PTR_IMM,
		},
		INST_MNEMONICS_SUB: {
			// sub b, a
			ADDRESSING_MODE_REG_REG: INST_SUB_REG_REG,
			// sub a, 4
			ADDRESSING_MODE_REG_IMM: INST_SUB_REG_IMM,
			// sub sp, 4
			ADDRESSING_MODE_SP_IMM: INST_SUB_SP_IMM,
			// sub ip, 4
			ADDRESSING_MODE_IP_IMM: INST_SUB_IP_IMM,
		},
		INST_MNEMONICS_INC: {
			// inc a
			ADDRESSING_MODE_IMPL_REG: INST_INC_REG,
		},
		INST_MNEMONICS_DEC: {
			// dec a
			ADDRESSING_MODE_IMPL_REG: INST_DEC_REG,
		},
		INST_MNEMONICS_JMP: {
			// jmp 0x5500
			ADDRESSING_MODE_IMPL_IMM_16: INST_JMP_IMM,
		},
		INST_MNEMONICS_JS: {
			// js 0x6000
			ADDRESSING_MODE_IMPL_IMM_16: INST_JS_IMM,
		},
		INST_MNEMONICS_JNS: {
			// jns 0x6000
			ADDRESSING_MODE_IMPL_IMM_16: INST_JNS_IMM,
		},
		INST_MNEMONICS_JZ: {
			// jz 0x6000
			ADDRESSING_MODE_IMPL_IMM_16: INST_JZ_IMM,
		},
		INST_MNEMONICS_JNZ: {
			// jnz 0x6000
			ADDRESSING_MODE_IMPL_IMM_16: INST_JNZ_IMM,
		},
		INST_MNEMONICS_JC: {
			// jc 0x6000
			ADDRESSING_MODE_IMPL_IMM_16: INST_JC_IMM,
		},
		INST_MNEMONICS_NOP: {
			ADDRESSING_MODE_IMPL: INST_NOP_IMPL,
		},
		INST_MNEMONICS_PUSH: {
			// push sp
			ADDRESSING_MODE_IMPL_SP: INST_PUSH_SP,
			// push ip
			ADDRESSING_MODE_IMPL_IP: INST_PUSH_IP,
			// push a
			ADDRESSING_MODE_IMPL_REG: INST_PUSH_REG,
			// push 0x10
			ADDRESSING_MODE_REG_IMM: INST_PUSH_IMM,
		},
		INST_MNEMONICS_POP: {
			// pop sp
			ADDRESSING_MODE_IMPL_SP: INST_POP_SP,
			// pop ip
			ADDRESSING_MODE_IMPL_IP: INST_POP_IP,
			// pop a
			ADDRESSING_MODE_IMPL_REG: INST_POP_REG,
		},
		INST_MNEMONICS_CALL: {
			// call
			ADDRESSING_MODE_IMPL_IMM_16: INST_CALL_IMPL,
		},
		INST_MNEMONICS_RET: {
			// ret
			ADDRESSING_MODE_IMPL: INST_RET_IMPL,
		},
		INST_MNEMONICS_CLC: {
			// clc
			ADDRESSING_MODE_IMPL: INST_CLC_IMPL,
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
