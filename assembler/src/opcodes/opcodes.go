package opcodes

const (
	REGISTER_SYMBOL_A  string = "a"
	REGISTER_SYMBOL_B  string = "b"
	REGISTER_SYMBOL_C  string = "c"
	REGISTER_SYMBOL_D  string = "d"
	REGISTER_SYMBOL_SP string = "sp"

	INST_BRK  string = "brk"
	INST_MOV  string = "mov"
	INST_ADD  string = "add"
	INST_ADC  string = "adc"
	INST_SUB  string = "sub"
	INST_SBB  string = "sbb"
	INST_XOR  string = "xor"
	INST_AND  string = "and"
	INST_OR   string = "or"
	INST_NOT  string = "not"
	INST_SHL  string = "shl"
	INST_SHR  string = "shr"
	INST_ROL  string = "rol"
	INST_ROR  string = "ror"
	INST_CMP  string = "cmp"
	INST_JMP  string = "jmp"
	INST_JZ   string = "jz"
	INST_JNZ  string = "jnz"
	INST_JC   string = "jc"
	INST_JS   string = "js"
	INST_JNS  string = "jns"
	INST_CLC  string = "clc"
	INST_PUSH string = "push"
	INST_POP  string = "pop"
	INST_INC  string = "inc"
	INST_DEC  string = "dec"
	INST_JSR  string = "jsr"
	INST_RTS  string = "rts"
	INST_NOP  string = "nop"

	ADDRESSING_MODE_NONE uint8 = iota
	ADDRESSING_MODE_IMPL
	ADDRESSING_MODE_IMPL_REG
	ADDRESSING_MODE_REG_IMM_8
	ADDRESSING_MODE_IMPL_IMM_16
	ADDRESSING_MODE_REG_REG
	ADDRESSING_MODE_REG_MEM
	ADDRESSING_MODE_MEM_REG
	ADDRESSING_MODE_MEM_REG_OFFSET
	ADDRESSING_MODE_REG_MEM_OFFSET
	ADDRESSING_MODE_IMPL_STACK
	ADDRESSING_MODE_STACK_IMM
	ADDRESSING_MODE_STACK_REG
	ADDRESSING_MODE_REG_STACK
	ADDRESSING_MODE_STACK_REG_OFFSET
	ADDRESSING_MODE_REG_STACK_OFFSET
	ADDRESSING_MODE_STACK_REG_OFFSET_REG
	ADDRESSING_MODE_REG_STACK_OFFSET_REG
)

var (
	REGISTER_OPCODE = map[string]uint8{
		REGISTER_SYMBOL_A:  0,
		REGISTER_SYMBOL_B:  1,
		REGISTER_SYMBOL_C:  2,
		REGISTER_SYMBOL_D:  3,
		REGISTER_SYMBOL_SP: 4,
	}

	_MNEMONIC_TO_OPCODE = map[string]map[uint8]uint8{
		INST_BRK: {
			ADDRESSING_MODE_IMPL: 0,
		},
		INST_MOV: {
			// mov b, a
			ADDRESSING_MODE_REG_REG: 0x40,
			// mov a, 4
			ADDRESSING_MODE_REG_IMM_8: 0x41,
			// mov a, [8000]
			ADDRESSING_MODE_MEM_REG: 0x42,
			// mov [8000], a
			ADDRESSING_MODE_REG_MEM: 0x43,
			// mov a, [8000+b]
			ADDRESSING_MODE_MEM_REG_OFFSET: 0x44,
			// mov [8000+b], a
			ADDRESSING_MODE_REG_MEM_OFFSET: 0x45,
			// mov a, [sp]
			ADDRESSING_MODE_REG_STACK: 0x46,
			// mov [sp], a
			ADDRESSING_MODE_STACK_REG: 0x47,
			// mov a, [sp+4]
			ADDRESSING_MODE_REG_STACK_OFFSET: 0x48,
			// mov [sp+4], a
			ADDRESSING_MODE_STACK_REG_OFFSET: 0x49,
			// mov a, [sp+b]
			ADDRESSING_MODE_REG_STACK_OFFSET_REG: 0x4a,
			// mov [sp+b], a
			ADDRESSING_MODE_STACK_REG_OFFSET_REG: 0x4b,
		},
		INST_XOR: {
			// xor b, a
			ADDRESSING_MODE_REG_REG: 0x50,
			// xor b, 0x10
			ADDRESSING_MODE_REG_IMM_8: 0x51,
		},
		INST_AND: {
			// and b, a
			ADDRESSING_MODE_REG_REG: 0x52,
			// and b, 0x10
			ADDRESSING_MODE_REG_IMM_8: 0x53,
		},
		INST_OR: {
			// or b, a
			ADDRESSING_MODE_REG_REG: 0x54,
			// or b, 0x10
			ADDRESSING_MODE_REG_IMM_8: 0x55,
		},
		INST_NOT: {
			// not a
			ADDRESSING_MODE_IMPL_REG: 0x56,
		},
		INST_SHL: {
			// shl a
			ADDRESSING_MODE_IMPL_REG: 0x57,
		},
		INST_SHR: {
			// shr a
			ADDRESSING_MODE_IMPL_REG: 0x58,
		},
		INST_ROL: {
			// rol a
			ADDRESSING_MODE_IMPL_REG: 0x59,
		},
		INST_ROR: {
			// ror a
			ADDRESSING_MODE_IMPL_REG: 0x5a,
		},
		INST_CMP: {
			// cmp b, a
			ADDRESSING_MODE_REG_REG: 0x1a,
			// cmp a, 4
			ADDRESSING_MODE_REG_IMM_8: 0x1b,
		},
		INST_ADD: {
			// add b, a
			ADDRESSING_MODE_REG_REG: 0x30,
			// add a, 4
			ADDRESSING_MODE_REG_IMM_8: 0x31,
			// add [8000], a
			ADDRESSING_MODE_REG_MEM: 0x3a,
			// add sp, 4
			ADDRESSING_MODE_STACK_IMM: 0x3c,
		},
		INST_ADC: {
			// adc b, a
			ADDRESSING_MODE_REG_REG: 0x32,
			// adc a, 4
			ADDRESSING_MODE_REG_IMM_8: 0x33,
		},
		INST_SUB: {
			// sub b, a
			ADDRESSING_MODE_REG_REG: 0x34,
			// sub a, 4
			ADDRESSING_MODE_REG_IMM_8: 0x35,
			// sub sp, 4
			ADDRESSING_MODE_STACK_IMM: 0x3d,
		},
		INST_SBB: {
			// sbb b, a
			ADDRESSING_MODE_REG_REG: 0x36,
			// sbb a, 4
			ADDRESSING_MODE_REG_IMM_8: 0x37,
		},
		INST_INC: {
			// inc a
			ADDRESSING_MODE_IMPL_REG: 0x38,
		},
		INST_DEC: {
			// dec a
			ADDRESSING_MODE_IMPL_REG: 0x39,
		},
		INST_JMP: {
			// jmp 0x5500
			ADDRESSING_MODE_IMPL_IMM_16: 0x10,
		},
		INST_JS: {
			// js 0x6000
			ADDRESSING_MODE_IMPL_IMM_16: 0x13,
		},
		INST_JNS: {
			// jns 0x6000
			ADDRESSING_MODE_IMPL_IMM_16: 0x14,
		},
		INST_JZ: {
			// jz 0x6000
			ADDRESSING_MODE_IMPL_IMM_16: 0x15,
		},
		INST_JNZ: {
			// jnz 0x6000
			ADDRESSING_MODE_IMPL_IMM_16: 0x16,
		},
		INST_JC: {
			// jc 0x6000
			ADDRESSING_MODE_IMPL_IMM_16: 0x17,
		},
		INST_NOP: {
			ADDRESSING_MODE_IMPL: 0xff,
		},
		INST_PUSH: {
			// push sp
			ADDRESSING_MODE_IMPL_STACK: 0xee,
			// push a
			ADDRESSING_MODE_IMPL_REG: 0xf0,
			// push 0x10
			ADDRESSING_MODE_REG_IMM_8: 0xf1,
		},
		INST_POP: {
			// pop sp
			ADDRESSING_MODE_IMPL_STACK: 0xef,
			// pop a
			ADDRESSING_MODE_IMPL_REG: 0xf2,
		},
		INST_JSR: {
			// jsr
			ADDRESSING_MODE_IMPL_IMM_16: 0xf3,
		},
		INST_RTS: {
			// rts
			ADDRESSING_MODE_IMPL: 0xf4,
		},
		INST_CLC: {
			// clc
			ADDRESSING_MODE_IMPL: 0xe0,
		},
	}
)

func GetOpCode(inst string, addressingMode uint8) uint8 {
	modes, ok := _MNEMONIC_TO_OPCODE[inst]
	if !ok {
		return 0
	}
	op, ok := modes[addressingMode]
	if !ok {
		return 0
	}
	return op
}
