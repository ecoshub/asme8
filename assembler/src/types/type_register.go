package types

import (
	"asme8/assembler/src/opcodes"
	"strings"
)

type Register struct {
	identifier string
	code       uint8
}

func NewRegister(identifier string, code uint8) *Register {
	return &Register{identifier: identifier, code: code}
}

func ParseRegister(text string) *Register {
	identifier := strings.ToLower(text)
	return &Register{
		identifier: identifier,
		code:       opcodes.REGISTER_OPCODE[identifier],
	}
}

func (r *Register) GetCode() uint8 {
	return r.code
}
