package symbols

const (
	/*
		; mov, add, adc, sub, sbb, and, or, xor, cmp
		mov
		    mov bl, cl          ; [1] r8 → r8
		    mov d, a            ; [1] r16 → r16
		    mov dl, #10         ; [8] imm8 → r8
		    mov c, #10fe        ; [8] imm16 → r16
		    mov [#10fe], al     ; [8] r8 → m8
		    mov bh, [#10fe]     ; [8] m8 → r8
		    mov cl, [a]         ; [1] ptr16 → r8
		    mov [b], ch         ; [1] r8 → ptr16
		; 9 operation (36) opcode

		; jmp, jmpz, jmpc, jmpv, jmpnz, jmpnc, jmpnv, call
		jmp
		    jmp #10fe           ; [1] imm16
		    jmp #10fe + 0x7f    ; [1] imm16 + offset
		    jmp #10fe + 0x7fff  ; [1] imm16 + offset
		    jmp sp              ; [1] r16
		; 7 operation (4) opcode


		; ret, rfi, hlt, nop
		ret
		    hlt
		    nop
		; 4 operation (4) opcode

	*/

	REGISTER_SYMBOL_A string = "a"
	REGISTER_SYMBOL_B string = "b"
	REGISTER_SYMBOL_C string = "c"
	REGISTER_SYMBOL_D string = "d"

	INST_BRK  string = "brk"
	INST_MOV  string = "mov"
	INST_ADD  string = "add"
	INST_SUB  string = "sub"
	INST_CMP  string = "cmp"
	INST_JMP  string = "jmp"
	INST_JZ   string = "jz"
	INST_NOP  string = "nop"
	INST_PUSH string = "push"
	INST_POP  string = "pop"
	INST_JSR  string = "jsr"
	INST_RTS  string = "rts"

	ADDRESSING_MODE_NONE uint8 = iota
	ADDRESSING_MODE_IMPL
	ADDRESSING_MODE_INM_8
	ADDRESSING_MODE_IMPL_REG
	ADDRESSING_MODE_RR_8
	ADDRESSING_MODE_INM_16
	ADDRESSING_MODE_REG_MEM
	ADDRESSING_MODE_MEM_REG
	ADDRESSING_MODE_PTR
	ADDRESSING_MODE_MEM_REG_OFFSET
	ADDRESSING_MODE_REG_MEM_OFFSET
)

var (
	REGISTER_OPCODE = map[string]uint8{
		REGISTER_SYMBOL_A: 0,
		REGISTER_SYMBOL_B: 1,
		REGISTER_SYMBOL_C: 2,
		REGISTER_SYMBOL_D: 3,
	}

	_MNEMONIC_TO_OPCODE = map[string]map[uint8]uint8{
		INST_BRK: {
			ADDRESSING_MODE_NONE:    0,
			ADDRESSING_MODE_INM_8:   0,
			ADDRESSING_MODE_RR_8:    0,
			ADDRESSING_MODE_INM_16:  0,
			ADDRESSING_MODE_REG_MEM: 0,
			ADDRESSING_MODE_MEM_REG: 0,
		},
		INST_MOV: {
			// mov b, a
			ADDRESSING_MODE_RR_8: 0x40,
			// mov a, 4
			ADDRESSING_MODE_INM_8: 0x41,
			// mov a, [8000]
			ADDRESSING_MODE_MEM_REG: 0x42,
			// mov [8000], a
			ADDRESSING_MODE_REG_MEM: 0x43,
			// mov a, [8000+b]
			ADDRESSING_MODE_MEM_REG_OFFSET: 0x44,
			// mov [8000+b], a
			ADDRESSING_MODE_REG_MEM_OFFSET: 0x45,
		},
		INST_CMP: {
			// cmp b, a
			ADDRESSING_MODE_RR_8: 0x3a,
			// cmp a, 4
			ADDRESSING_MODE_INM_8: 0x3b,
		},
		INST_ADD: {
			// add b, a
			ADDRESSING_MODE_RR_8: 0x30,
			// add a, 4
			ADDRESSING_MODE_INM_8: 0x31,
		},
		INST_SUB: {
			// sub b, a
			ADDRESSING_MODE_RR_8: 0x35,
			// sub a, 4
			ADDRESSING_MODE_INM_8: 0x36,
		},
		INST_JMP: {
			// jmp 0x5500
			ADDRESSING_MODE_INM_16: 0x10,
		},
		INST_JZ: {
			// jmp 0x5500
			ADDRESSING_MODE_INM_16: 0x15,
		},
		INST_NOP: {
			ADDRESSING_MODE_IMPL: 0xff,
		},
		INST_PUSH: {
			// push a
			ADDRESSING_MODE_IMPL_REG: 0xf0,
			// push 0x10
			ADDRESSING_MODE_INM_8: 0xf1,
		},
		INST_POP: {
			// pop a
			ADDRESSING_MODE_IMPL_REG: 0xf2,
		},
		INST_JSR: {
			// jsr
			ADDRESSING_MODE_IMPL: 0xf3,
		},
		INST_RTS: {
			// rts
			ADDRESSING_MODE_IMPL: 0xf4,
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
