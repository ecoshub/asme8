// Code generated from AsmE8.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // AsmE8

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type AsmE8Parser struct {
	*antlr.BaseParser
}

var AsmE8ParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func asme8ParserInit() {
	staticData := &AsmE8ParserStaticData
	staticData.LiteralNames = []string{
		"", "'\\t'", "'    '", "'  '", "' '", "'\\n'", "':'", "', '", "'adc'",
		"'add'", "'and'", "'brk'", "'call'", "'clc'", "'cmp'", "'dec'", "'inc'",
		"'jc'", "'jmp'", "'jnc'", "'jns'", "'jnz'", "'js'", "'jz'", "'mov'",
		"'nop'", "'not'", "'or'", "'pop'", "'push'", "'ret'", "'rol'", "'ror'",
		"'rti'", "'sbb'", "'shl'", "'shr'", "'sub'", "'xor'", "'a'", "'b'",
		"'c'", "'d'", "'e'", "'ipl'", "'iph'", "'bpl'", "'bph'", "'ip'", "'bp'",
		"'sp'", "'['", "']'", "'+'", "'-'", "'='", "'.asciiz '", "'.ascii '",
		"'.byte '", "'.word '", "'.resb '", "'.org '", "'.segment '", "'global '",
		"'extern '",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "LINE_COMMENT",
		"SUFFIX_WHITESPACE", "WHITE_SPACE", "BINARY", "INT", "CHAR", "ANY_STR",
		"HEX", "STR",
	}
	staticData.RuleNames = []string{
		"prog", "instruction", "line", "label", "inst", "addr_mode_imm16", "addr_mode_imm16_tag",
		"addr_mode_reg_to_mem_indexed", "addr_mode_reg_to_mem_direct", "addr_mode_reg_to_mem_indirect_offset",
		"addr_mode_reg_to_mem_reg16_indexed", "addr_mode_reg_to_mem_indirect",
		"addr_mode_implied_reg8", "addr_mode_reg8_imm8", "addr_mode_reg8_imm8_tag",
		"addr_mode_mem_to_reg_indexed", "addr_mode_mem_to_reg_direct", "addr_mode_mem_to_reg_indirect_offset",
		"addr_mode_mem_to_reg_reg16_indexed", "addr_mode_mem_to_reg_indirect",
		"addr_mode_reg8_reg8", "addr_mode_implied", "addr_mode_implied_reg16",
		"addr_mode_reg16_imm", "addr_mode_reg16_imm_tag", "addr_mode_reg16_reg16",
		"addr_mode_reg16_stack", "addr_mode_implied_stack", "addr_mode_stack_imm8",
		"addr_mode_stack_imm8_tag", "addr_mode_stack_reg16", "mnemonic", "reg8",
		"reg16", "stack", "direct", "direct_virtual_offset", "direct_offset",
		"indirect_offset", "reg16_indexed", "indirect", "variable", "variable_reference",
		"reference", "directives", "segment", "access", "imm_list", "imm", "ascii",
		"tag",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 73, 426, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36, 7, 36,
		2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7, 41, 2,
		42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 2, 46, 7, 46, 2, 47,
		7, 47, 2, 48, 7, 48, 2, 49, 7, 49, 2, 50, 7, 50, 1, 0, 1, 0, 1, 0, 1, 1,
		4, 1, 107, 8, 1, 11, 1, 12, 1, 108, 1, 1, 3, 1, 112, 8, 1, 1, 1, 4, 1,
		115, 8, 1, 11, 1, 12, 1, 116, 1, 1, 3, 1, 120, 8, 1, 3, 1, 122, 8, 1, 1,
		2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1,
		2, 3, 2, 137, 8, 2, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1,
		4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1,
		4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 3, 4, 168, 8, 4, 1,
		5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1,
		7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1,
		9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11,
		1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 1,
		13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15,
		1, 15, 1, 15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1,
		17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18,
		1, 18, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 20, 1, 20, 1, 20, 1,
		20, 1, 20, 1, 20, 1, 21, 1, 21, 1, 22, 1, 22, 1, 22, 1, 22, 1, 23, 1, 23,
		1, 23, 1, 23, 1, 23, 1, 23, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1,
		25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26,
		1, 26, 1, 27, 1, 27, 1, 27, 1, 27, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1,
		28, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 30, 1, 30, 1, 30, 1, 30,
		1, 30, 1, 30, 1, 31, 1, 31, 1, 32, 1, 32, 1, 33, 1, 33, 1, 34, 1, 34, 1,
		35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 3, 35, 329,
		8, 35, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 37, 1, 37, 1, 37, 1,
		37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 3, 37, 349,
		8, 37, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 39, 1, 39, 1, 39, 1,
		39, 1, 39, 1, 39, 1, 40, 1, 40, 1, 40, 1, 40, 1, 41, 1, 41, 1, 41, 1, 41,
		1, 42, 1, 42, 1, 42, 1, 42, 1, 43, 1, 43, 1, 43, 1, 43, 3, 43, 379, 8,
		43, 1, 44, 1, 44, 1, 44, 1, 44, 1, 44, 1, 44, 1, 44, 1, 44, 1, 44, 1, 44,
		1, 44, 1, 44, 3, 44, 393, 8, 44, 1, 45, 1, 45, 1, 45, 1, 46, 1, 46, 1,
		46, 1, 46, 1, 46, 1, 46, 3, 46, 404, 8, 46, 1, 47, 1, 47, 3, 47, 408, 8,
		47, 1, 47, 1, 47, 1, 47, 3, 47, 413, 8, 47, 5, 47, 415, 8, 47, 10, 47,
		12, 47, 418, 9, 47, 1, 48, 1, 48, 1, 49, 1, 49, 1, 50, 1, 50, 1, 50, 0,
		0, 51, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34,
		36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 70,
		72, 74, 76, 78, 80, 82, 84, 86, 88, 90, 92, 94, 96, 98, 100, 0, 6, 1, 0,
		1, 3, 1, 0, 8, 38, 1, 0, 39, 47, 1, 0, 48, 49, 1, 0, 53, 54, 2, 0, 68,
		70, 72, 72, 425, 0, 102, 1, 0, 0, 0, 2, 121, 1, 0, 0, 0, 4, 136, 1, 0,
		0, 0, 6, 138, 1, 0, 0, 0, 8, 167, 1, 0, 0, 0, 10, 169, 1, 0, 0, 0, 12,
		173, 1, 0, 0, 0, 14, 177, 1, 0, 0, 0, 16, 183, 1, 0, 0, 0, 18, 189, 1,
		0, 0, 0, 20, 195, 1, 0, 0, 0, 22, 201, 1, 0, 0, 0, 24, 207, 1, 0, 0, 0,
		26, 211, 1, 0, 0, 0, 28, 217, 1, 0, 0, 0, 30, 223, 1, 0, 0, 0, 32, 229,
		1, 0, 0, 0, 34, 235, 1, 0, 0, 0, 36, 241, 1, 0, 0, 0, 38, 247, 1, 0, 0,
		0, 40, 253, 1, 0, 0, 0, 42, 259, 1, 0, 0, 0, 44, 261, 1, 0, 0, 0, 46, 265,
		1, 0, 0, 0, 48, 271, 1, 0, 0, 0, 50, 277, 1, 0, 0, 0, 52, 283, 1, 0, 0,
		0, 54, 289, 1, 0, 0, 0, 56, 293, 1, 0, 0, 0, 58, 299, 1, 0, 0, 0, 60, 305,
		1, 0, 0, 0, 62, 311, 1, 0, 0, 0, 64, 313, 1, 0, 0, 0, 66, 315, 1, 0, 0,
		0, 68, 317, 1, 0, 0, 0, 70, 328, 1, 0, 0, 0, 72, 330, 1, 0, 0, 0, 74, 348,
		1, 0, 0, 0, 76, 350, 1, 0, 0, 0, 78, 356, 1, 0, 0, 0, 80, 362, 1, 0, 0,
		0, 82, 366, 1, 0, 0, 0, 84, 370, 1, 0, 0, 0, 86, 378, 1, 0, 0, 0, 88, 392,
		1, 0, 0, 0, 90, 394, 1, 0, 0, 0, 92, 403, 1, 0, 0, 0, 94, 407, 1, 0, 0,
		0, 96, 419, 1, 0, 0, 0, 98, 421, 1, 0, 0, 0, 100, 423, 1, 0, 0, 0, 102,
		103, 3, 2, 1, 0, 103, 104, 5, 0, 0, 1, 104, 1, 1, 0, 0, 0, 105, 107, 3,
		4, 2, 0, 106, 105, 1, 0, 0, 0, 107, 108, 1, 0, 0, 0, 108, 106, 1, 0, 0,
		0, 108, 109, 1, 0, 0, 0, 109, 111, 1, 0, 0, 0, 110, 112, 5, 65, 0, 0, 111,
		110, 1, 0, 0, 0, 111, 112, 1, 0, 0, 0, 112, 122, 1, 0, 0, 0, 113, 115,
		3, 4, 2, 0, 114, 113, 1, 0, 0, 0, 115, 116, 1, 0, 0, 0, 116, 114, 1, 0,
		0, 0, 116, 117, 1, 0, 0, 0, 117, 119, 1, 0, 0, 0, 118, 120, 5, 66, 0, 0,
		119, 118, 1, 0, 0, 0, 119, 120, 1, 0, 0, 0, 120, 122, 1, 0, 0, 0, 121,
		106, 1, 0, 0, 0, 121, 114, 1, 0, 0, 0, 122, 3, 1, 0, 0, 0, 123, 137, 3,
		84, 42, 0, 124, 137, 3, 82, 41, 0, 125, 126, 7, 0, 0, 0, 126, 137, 3, 8,
		4, 0, 127, 137, 3, 90, 45, 0, 128, 137, 3, 92, 46, 0, 129, 130, 3, 6, 3,
		0, 130, 131, 5, 4, 0, 0, 131, 132, 3, 88, 44, 0, 132, 137, 1, 0, 0, 0,
		133, 137, 3, 88, 44, 0, 134, 137, 3, 6, 3, 0, 135, 137, 5, 5, 0, 0, 136,
		123, 1, 0, 0, 0, 136, 124, 1, 0, 0, 0, 136, 125, 1, 0, 0, 0, 136, 127,
		1, 0, 0, 0, 136, 128, 1, 0, 0, 0, 136, 129, 1, 0, 0, 0, 136, 133, 1, 0,
		0, 0, 136, 134, 1, 0, 0, 0, 136, 135, 1, 0, 0, 0, 137, 5, 1, 0, 0, 0, 138,
		139, 3, 100, 50, 0, 139, 140, 5, 6, 0, 0, 140, 7, 1, 0, 0, 0, 141, 168,
		3, 10, 5, 0, 142, 168, 3, 12, 6, 0, 143, 168, 3, 14, 7, 0, 144, 168, 3,
		16, 8, 0, 145, 168, 3, 18, 9, 0, 146, 168, 3, 20, 10, 0, 147, 168, 3, 22,
		11, 0, 148, 168, 3, 24, 12, 0, 149, 168, 3, 26, 13, 0, 150, 168, 3, 28,
		14, 0, 151, 168, 3, 30, 15, 0, 152, 168, 3, 32, 16, 0, 153, 168, 3, 34,
		17, 0, 154, 168, 3, 36, 18, 0, 155, 168, 3, 38, 19, 0, 156, 168, 3, 40,
		20, 0, 157, 168, 3, 42, 21, 0, 158, 168, 3, 44, 22, 0, 159, 168, 3, 46,
		23, 0, 160, 168, 3, 48, 24, 0, 161, 168, 3, 50, 25, 0, 162, 168, 3, 52,
		26, 0, 163, 168, 3, 54, 27, 0, 164, 168, 3, 56, 28, 0, 165, 168, 3, 58,
		29, 0, 166, 168, 3, 60, 30, 0, 167, 141, 1, 0, 0, 0, 167, 142, 1, 0, 0,
		0, 167, 143, 1, 0, 0, 0, 167, 144, 1, 0, 0, 0, 167, 145, 1, 0, 0, 0, 167,
		146, 1, 0, 0, 0, 167, 147, 1, 0, 0, 0, 167, 148, 1, 0, 0, 0, 167, 149,
		1, 0, 0, 0, 167, 150, 1, 0, 0, 0, 167, 151, 1, 0, 0, 0, 167, 152, 1, 0,
		0, 0, 167, 153, 1, 0, 0, 0, 167, 154, 1, 0, 0, 0, 167, 155, 1, 0, 0, 0,
		167, 156, 1, 0, 0, 0, 167, 157, 1, 0, 0, 0, 167, 158, 1, 0, 0, 0, 167,
		159, 1, 0, 0, 0, 167, 160, 1, 0, 0, 0, 167, 161, 1, 0, 0, 0, 167, 162,
		1, 0, 0, 0, 167, 163, 1, 0, 0, 0, 167, 164, 1, 0, 0, 0, 167, 165, 1, 0,
		0, 0, 167, 166, 1, 0, 0, 0, 168, 9, 1, 0, 0, 0, 169, 170, 3, 62, 31, 0,
		170, 171, 5, 4, 0, 0, 171, 172, 3, 96, 48, 0, 172, 11, 1, 0, 0, 0, 173,
		174, 3, 62, 31, 0, 174, 175, 5, 4, 0, 0, 175, 176, 3, 100, 50, 0, 176,
		13, 1, 0, 0, 0, 177, 178, 3, 62, 31, 0, 178, 179, 5, 4, 0, 0, 179, 180,
		3, 74, 37, 0, 180, 181, 5, 7, 0, 0, 181, 182, 3, 64, 32, 0, 182, 15, 1,
		0, 0, 0, 183, 184, 3, 62, 31, 0, 184, 185, 5, 4, 0, 0, 185, 186, 3, 70,
		35, 0, 186, 187, 5, 7, 0, 0, 187, 188, 3, 64, 32, 0, 188, 17, 1, 0, 0,
		0, 189, 190, 3, 62, 31, 0, 190, 191, 5, 4, 0, 0, 191, 192, 3, 76, 38, 0,
		192, 193, 5, 7, 0, 0, 193, 194, 3, 64, 32, 0, 194, 19, 1, 0, 0, 0, 195,
		196, 3, 62, 31, 0, 196, 197, 5, 4, 0, 0, 197, 198, 3, 78, 39, 0, 198, 199,
		5, 7, 0, 0, 199, 200, 3, 64, 32, 0, 200, 21, 1, 0, 0, 0, 201, 202, 3, 62,
		31, 0, 202, 203, 5, 4, 0, 0, 203, 204, 3, 80, 40, 0, 204, 205, 5, 7, 0,
		0, 205, 206, 3, 64, 32, 0, 206, 23, 1, 0, 0, 0, 207, 208, 3, 62, 31, 0,
		208, 209, 5, 4, 0, 0, 209, 210, 3, 64, 32, 0, 210, 25, 1, 0, 0, 0, 211,
		212, 3, 62, 31, 0, 212, 213, 5, 4, 0, 0, 213, 214, 3, 64, 32, 0, 214, 215,
		5, 7, 0, 0, 215, 216, 3, 96, 48, 0, 216, 27, 1, 0, 0, 0, 217, 218, 3, 62,
		31, 0, 218, 219, 5, 4, 0, 0, 219, 220, 3, 64, 32, 0, 220, 221, 5, 7, 0,
		0, 221, 222, 3, 100, 50, 0, 222, 29, 1, 0, 0, 0, 223, 224, 3, 62, 31, 0,
		224, 225, 5, 4, 0, 0, 225, 226, 3, 64, 32, 0, 226, 227, 5, 7, 0, 0, 227,
		228, 3, 74, 37, 0, 228, 31, 1, 0, 0, 0, 229, 230, 3, 62, 31, 0, 230, 231,
		5, 4, 0, 0, 231, 232, 3, 64, 32, 0, 232, 233, 5, 7, 0, 0, 233, 234, 3,
		70, 35, 0, 234, 33, 1, 0, 0, 0, 235, 236, 3, 62, 31, 0, 236, 237, 5, 4,
		0, 0, 237, 238, 3, 64, 32, 0, 238, 239, 5, 7, 0, 0, 239, 240, 3, 76, 38,
		0, 240, 35, 1, 0, 0, 0, 241, 242, 3, 62, 31, 0, 242, 243, 5, 4, 0, 0, 243,
		244, 3, 64, 32, 0, 244, 245, 5, 7, 0, 0, 245, 246, 3, 78, 39, 0, 246, 37,
		1, 0, 0, 0, 247, 248, 3, 62, 31, 0, 248, 249, 5, 4, 0, 0, 249, 250, 3,
		64, 32, 0, 250, 251, 5, 7, 0, 0, 251, 252, 3, 80, 40, 0, 252, 39, 1, 0,
		0, 0, 253, 254, 3, 62, 31, 0, 254, 255, 5, 4, 0, 0, 255, 256, 3, 64, 32,
		0, 256, 257, 5, 7, 0, 0, 257, 258, 3, 64, 32, 0, 258, 41, 1, 0, 0, 0, 259,
		260, 3, 62, 31, 0, 260, 43, 1, 0, 0, 0, 261, 262, 3, 62, 31, 0, 262, 263,
		5, 4, 0, 0, 263, 264, 3, 66, 33, 0, 264, 45, 1, 0, 0, 0, 265, 266, 3, 62,
		31, 0, 266, 267, 5, 4, 0, 0, 267, 268, 3, 66, 33, 0, 268, 269, 5, 7, 0,
		0, 269, 270, 3, 96, 48, 0, 270, 47, 1, 0, 0, 0, 271, 272, 3, 62, 31, 0,
		272, 273, 5, 4, 0, 0, 273, 274, 3, 66, 33, 0, 274, 275, 5, 7, 0, 0, 275,
		276, 3, 100, 50, 0, 276, 49, 1, 0, 0, 0, 277, 278, 3, 62, 31, 0, 278, 279,
		5, 4, 0, 0, 279, 280, 3, 66, 33, 0, 280, 281, 5, 7, 0, 0, 281, 282, 3,
		66, 33, 0, 282, 51, 1, 0, 0, 0, 283, 284, 3, 62, 31, 0, 284, 285, 5, 4,
		0, 0, 285, 286, 3, 66, 33, 0, 286, 287, 5, 7, 0, 0, 287, 288, 3, 68, 34,
		0, 288, 53, 1, 0, 0, 0, 289, 290, 3, 62, 31, 0, 290, 291, 5, 4, 0, 0, 291,
		292, 3, 68, 34, 0, 292, 55, 1, 0, 0, 0, 293, 294, 3, 62, 31, 0, 294, 295,
		5, 4, 0, 0, 295, 296, 3, 68, 34, 0, 296, 297, 5, 7, 0, 0, 297, 298, 3,
		96, 48, 0, 298, 57, 1, 0, 0, 0, 299, 300, 3, 62, 31, 0, 300, 301, 5, 4,
		0, 0, 301, 302, 3, 68, 34, 0, 302, 303, 5, 7, 0, 0, 303, 304, 3, 100, 50,
		0, 304, 59, 1, 0, 0, 0, 305, 306, 3, 62, 31, 0, 306, 307, 5, 4, 0, 0, 307,
		308, 3, 68, 34, 0, 308, 309, 5, 7, 0, 0, 309, 310, 3, 66, 33, 0, 310, 61,
		1, 0, 0, 0, 311, 312, 7, 1, 0, 0, 312, 63, 1, 0, 0, 0, 313, 314, 7, 2,
		0, 0, 314, 65, 1, 0, 0, 0, 315, 316, 7, 3, 0, 0, 316, 67, 1, 0, 0, 0, 317,
		318, 5, 50, 0, 0, 318, 69, 1, 0, 0, 0, 319, 320, 5, 51, 0, 0, 320, 321,
		3, 96, 48, 0, 321, 322, 5, 52, 0, 0, 322, 329, 1, 0, 0, 0, 323, 324, 5,
		51, 0, 0, 324, 325, 3, 100, 50, 0, 325, 326, 5, 52, 0, 0, 326, 329, 1,
		0, 0, 0, 327, 329, 3, 72, 36, 0, 328, 319, 1, 0, 0, 0, 328, 323, 1, 0,
		0, 0, 328, 327, 1, 0, 0, 0, 329, 71, 1, 0, 0, 0, 330, 331, 5, 51, 0, 0,
		331, 332, 5, 73, 0, 0, 332, 333, 7, 4, 0, 0, 333, 334, 5, 69, 0, 0, 334,
		335, 5, 52, 0, 0, 335, 73, 1, 0, 0, 0, 336, 337, 5, 51, 0, 0, 337, 338,
		3, 96, 48, 0, 338, 339, 5, 53, 0, 0, 339, 340, 3, 64, 32, 0, 340, 341,
		5, 52, 0, 0, 341, 349, 1, 0, 0, 0, 342, 343, 5, 51, 0, 0, 343, 344, 3,
		100, 50, 0, 344, 345, 5, 53, 0, 0, 345, 346, 3, 64, 32, 0, 346, 347, 5,
		52, 0, 0, 347, 349, 1, 0, 0, 0, 348, 336, 1, 0, 0, 0, 348, 342, 1, 0, 0,
		0, 349, 75, 1, 0, 0, 0, 350, 351, 5, 51, 0, 0, 351, 352, 3, 66, 33, 0,
		352, 353, 5, 53, 0, 0, 353, 354, 3, 96, 48, 0, 354, 355, 5, 52, 0, 0, 355,
		77, 1, 0, 0, 0, 356, 357, 5, 51, 0, 0, 357, 358, 3, 66, 33, 0, 358, 359,
		5, 53, 0, 0, 359, 360, 3, 64, 32, 0, 360, 361, 5, 52, 0, 0, 361, 79, 1,
		0, 0, 0, 362, 363, 5, 51, 0, 0, 363, 364, 3, 66, 33, 0, 364, 365, 5, 52,
		0, 0, 365, 81, 1, 0, 0, 0, 366, 367, 3, 100, 50, 0, 367, 368, 5, 55, 0,
		0, 368, 369, 3, 96, 48, 0, 369, 83, 1, 0, 0, 0, 370, 371, 3, 100, 50, 0,
		371, 372, 5, 55, 0, 0, 372, 373, 3, 86, 43, 0, 373, 85, 1, 0, 0, 0, 374,
		379, 5, 73, 0, 0, 375, 376, 5, 73, 0, 0, 376, 377, 7, 4, 0, 0, 377, 379,
		3, 96, 48, 0, 378, 374, 1, 0, 0, 0, 378, 375, 1, 0, 0, 0, 379, 87, 1, 0,
		0, 0, 380, 381, 5, 56, 0, 0, 381, 393, 3, 98, 49, 0, 382, 383, 5, 57, 0,
		0, 383, 393, 3, 98, 49, 0, 384, 385, 5, 58, 0, 0, 385, 393, 3, 94, 47,
		0, 386, 387, 5, 59, 0, 0, 387, 393, 3, 94, 47, 0, 388, 389, 5, 60, 0, 0,
		389, 393, 3, 94, 47, 0, 390, 391, 5, 61, 0, 0, 391, 393, 3, 94, 47, 0,
		392, 380, 1, 0, 0, 0, 392, 382, 1, 0, 0, 0, 392, 384, 1, 0, 0, 0, 392,
		386, 1, 0, 0, 0, 392, 388, 1, 0, 0, 0, 392, 390, 1, 0, 0, 0, 393, 89, 1,
		0, 0, 0, 394, 395, 5, 62, 0, 0, 395, 396, 5, 71, 0, 0, 396, 91, 1, 0, 0,
		0, 397, 398, 5, 2, 0, 0, 398, 399, 5, 63, 0, 0, 399, 404, 3, 100, 50, 0,
		400, 401, 5, 2, 0, 0, 401, 402, 5, 64, 0, 0, 402, 404, 3, 100, 50, 0, 403,
		397, 1, 0, 0, 0, 403, 400, 1, 0, 0, 0, 404, 93, 1, 0, 0, 0, 405, 408, 3,
		96, 48, 0, 406, 408, 3, 100, 50, 0, 407, 405, 1, 0, 0, 0, 407, 406, 1,
		0, 0, 0, 408, 416, 1, 0, 0, 0, 409, 412, 5, 7, 0, 0, 410, 413, 3, 96, 48,
		0, 411, 413, 3, 100, 50, 0, 412, 410, 1, 0, 0, 0, 412, 411, 1, 0, 0, 0,
		413, 415, 1, 0, 0, 0, 414, 409, 1, 0, 0, 0, 415, 418, 1, 0, 0, 0, 416,
		414, 1, 0, 0, 0, 416, 417, 1, 0, 0, 0, 417, 95, 1, 0, 0, 0, 418, 416, 1,
		0, 0, 0, 419, 420, 7, 5, 0, 0, 420, 97, 1, 0, 0, 0, 421, 422, 5, 71, 0,
		0, 422, 99, 1, 0, 0, 0, 423, 424, 5, 73, 0, 0, 424, 101, 1, 0, 0, 0, 15,
		108, 111, 116, 119, 121, 136, 167, 328, 348, 378, 392, 403, 407, 412, 416,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// AsmE8ParserInit initializes any static state used to implement AsmE8Parser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewAsmE8Parser(). You can call this function if you wish to initialize the static state ahead
// of time.
func AsmE8ParserInit() {
	staticData := &AsmE8ParserStaticData
	staticData.once.Do(asme8ParserInit)
}

// NewAsmE8Parser produces a new parser instance for the optional input antlr.TokenStream.
func NewAsmE8Parser(input antlr.TokenStream) *AsmE8Parser {
	AsmE8ParserInit()
	this := new(AsmE8Parser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &AsmE8ParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "AsmE8.g4"

	return this
}

// AsmE8Parser tokens.
const (
	AsmE8ParserEOF               = antlr.TokenEOF
	AsmE8ParserT__0              = 1
	AsmE8ParserT__1              = 2
	AsmE8ParserT__2              = 3
	AsmE8ParserT__3              = 4
	AsmE8ParserT__4              = 5
	AsmE8ParserT__5              = 6
	AsmE8ParserT__6              = 7
	AsmE8ParserT__7              = 8
	AsmE8ParserT__8              = 9
	AsmE8ParserT__9              = 10
	AsmE8ParserT__10             = 11
	AsmE8ParserT__11             = 12
	AsmE8ParserT__12             = 13
	AsmE8ParserT__13             = 14
	AsmE8ParserT__14             = 15
	AsmE8ParserT__15             = 16
	AsmE8ParserT__16             = 17
	AsmE8ParserT__17             = 18
	AsmE8ParserT__18             = 19
	AsmE8ParserT__19             = 20
	AsmE8ParserT__20             = 21
	AsmE8ParserT__21             = 22
	AsmE8ParserT__22             = 23
	AsmE8ParserT__23             = 24
	AsmE8ParserT__24             = 25
	AsmE8ParserT__25             = 26
	AsmE8ParserT__26             = 27
	AsmE8ParserT__27             = 28
	AsmE8ParserT__28             = 29
	AsmE8ParserT__29             = 30
	AsmE8ParserT__30             = 31
	AsmE8ParserT__31             = 32
	AsmE8ParserT__32             = 33
	AsmE8ParserT__33             = 34
	AsmE8ParserT__34             = 35
	AsmE8ParserT__35             = 36
	AsmE8ParserT__36             = 37
	AsmE8ParserT__37             = 38
	AsmE8ParserT__38             = 39
	AsmE8ParserT__39             = 40
	AsmE8ParserT__40             = 41
	AsmE8ParserT__41             = 42
	AsmE8ParserT__42             = 43
	AsmE8ParserT__43             = 44
	AsmE8ParserT__44             = 45
	AsmE8ParserT__45             = 46
	AsmE8ParserT__46             = 47
	AsmE8ParserT__47             = 48
	AsmE8ParserT__48             = 49
	AsmE8ParserT__49             = 50
	AsmE8ParserT__50             = 51
	AsmE8ParserT__51             = 52
	AsmE8ParserT__52             = 53
	AsmE8ParserT__53             = 54
	AsmE8ParserT__54             = 55
	AsmE8ParserT__55             = 56
	AsmE8ParserT__56             = 57
	AsmE8ParserT__57             = 58
	AsmE8ParserT__58             = 59
	AsmE8ParserT__59             = 60
	AsmE8ParserT__60             = 61
	AsmE8ParserT__61             = 62
	AsmE8ParserT__62             = 63
	AsmE8ParserT__63             = 64
	AsmE8ParserLINE_COMMENT      = 65
	AsmE8ParserSUFFIX_WHITESPACE = 66
	AsmE8ParserWHITE_SPACE       = 67
	AsmE8ParserBINARY            = 68
	AsmE8ParserINT               = 69
	AsmE8ParserCHAR              = 70
	AsmE8ParserANY_STR           = 71
	AsmE8ParserHEX               = 72
	AsmE8ParserSTR               = 73
)

// AsmE8Parser rules.
const (
	AsmE8ParserRULE_prog                                 = 0
	AsmE8ParserRULE_instruction                          = 1
	AsmE8ParserRULE_line                                 = 2
	AsmE8ParserRULE_label                                = 3
	AsmE8ParserRULE_inst                                 = 4
	AsmE8ParserRULE_addr_mode_imm16                      = 5
	AsmE8ParserRULE_addr_mode_imm16_tag                  = 6
	AsmE8ParserRULE_addr_mode_reg_to_mem_indexed         = 7
	AsmE8ParserRULE_addr_mode_reg_to_mem_direct          = 8
	AsmE8ParserRULE_addr_mode_reg_to_mem_indirect_offset = 9
	AsmE8ParserRULE_addr_mode_reg_to_mem_reg16_indexed   = 10
	AsmE8ParserRULE_addr_mode_reg_to_mem_indirect        = 11
	AsmE8ParserRULE_addr_mode_implied_reg8               = 12
	AsmE8ParserRULE_addr_mode_reg8_imm8                  = 13
	AsmE8ParserRULE_addr_mode_reg8_imm8_tag              = 14
	AsmE8ParserRULE_addr_mode_mem_to_reg_indexed         = 15
	AsmE8ParserRULE_addr_mode_mem_to_reg_direct          = 16
	AsmE8ParserRULE_addr_mode_mem_to_reg_indirect_offset = 17
	AsmE8ParserRULE_addr_mode_mem_to_reg_reg16_indexed   = 18
	AsmE8ParserRULE_addr_mode_mem_to_reg_indirect        = 19
	AsmE8ParserRULE_addr_mode_reg8_reg8                  = 20
	AsmE8ParserRULE_addr_mode_implied                    = 21
	AsmE8ParserRULE_addr_mode_implied_reg16              = 22
	AsmE8ParserRULE_addr_mode_reg16_imm                  = 23
	AsmE8ParserRULE_addr_mode_reg16_imm_tag              = 24
	AsmE8ParserRULE_addr_mode_reg16_reg16                = 25
	AsmE8ParserRULE_addr_mode_reg16_stack                = 26
	AsmE8ParserRULE_addr_mode_implied_stack              = 27
	AsmE8ParserRULE_addr_mode_stack_imm8                 = 28
	AsmE8ParserRULE_addr_mode_stack_imm8_tag             = 29
	AsmE8ParserRULE_addr_mode_stack_reg16                = 30
	AsmE8ParserRULE_mnemonic                             = 31
	AsmE8ParserRULE_reg8                                 = 32
	AsmE8ParserRULE_reg16                                = 33
	AsmE8ParserRULE_stack                                = 34
	AsmE8ParserRULE_direct                               = 35
	AsmE8ParserRULE_direct_virtual_offset                = 36
	AsmE8ParserRULE_direct_offset                        = 37
	AsmE8ParserRULE_indirect_offset                      = 38
	AsmE8ParserRULE_reg16_indexed                        = 39
	AsmE8ParserRULE_indirect                             = 40
	AsmE8ParserRULE_variable                             = 41
	AsmE8ParserRULE_variable_reference                   = 42
	AsmE8ParserRULE_reference                            = 43
	AsmE8ParserRULE_directives                           = 44
	AsmE8ParserRULE_segment                              = 45
	AsmE8ParserRULE_access                               = 46
	AsmE8ParserRULE_imm_list                             = 47
	AsmE8ParserRULE_imm                                  = 48
	AsmE8ParserRULE_ascii                                = 49
	AsmE8ParserRULE_tag                                  = 50
)

// IProgContext is an interface to support dynamic dispatch.
type IProgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Instruction() IInstructionContext
	EOF() antlr.TerminalNode

	// IsProgContext differentiates from other interfaces.
	IsProgContext()
}

type ProgContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgContext() *ProgContext {
	var p = new(ProgContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsmE8ParserRULE_prog
	return p
}

func InitEmptyProgContext(p *ProgContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsmE8ParserRULE_prog
}

func (*ProgContext) IsProgContext() {}

func NewProgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgContext {
	var p = new(ProgContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AsmE8ParserRULE_prog

	return p
}

func (s *ProgContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgContext) Instruction() IInstructionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInstructionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInstructionContext)
}

func (s *ProgContext) EOF() antlr.TerminalNode {
	return s.GetToken(AsmE8ParserEOF, 0)
}

func (s *ProgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsmE8Listener); ok {
		listenerT.EnterProg(s)
	}
}

func (s *ProgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsmE8Listener); ok {
		listenerT.ExitProg(s)
	}
}

func (p *AsmE8Parser) Prog() (localctx IProgContext) {
	localctx = NewProgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, AsmE8ParserRULE_prog)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(102)
		p.Instruction()
	}
	{
		p.SetState(103)
		p.Match(AsmE8ParserEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IInstructionContext is an interface to support dynamic dispatch.
type IInstructionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllLine() []ILineContext
	Line(i int) ILineContext
	LINE_COMMENT() antlr.TerminalNode
	SUFFIX_WHITESPACE() antlr.TerminalNode

	// IsInstructionContext differentiates from other interfaces.
	IsInstructionContext()
}

type InstructionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInstructionContext() *InstructionContext {
	var p = new(InstructionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsmE8ParserRULE_instruction
	return p
}

func InitEmptyInstructionContext(p *InstructionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsmE8ParserRULE_instruction
}

func (*InstructionContext) IsInstructionContext() {}

func NewInstructionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InstructionContext {
	var p = new(InstructionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AsmE8ParserRULE_instruction

	return p
}

func (s *InstructionContext) GetParser() antlr.Parser { return s.parser }

func (s *InstructionContext) AllLine() []ILineContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ILineContext); ok {
			len++
		}
	}

	tst := make([]ILineContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ILineContext); ok {
			tst[i] = t.(ILineContext)
			i++
		}
	}

	return tst
}

func (s *InstructionContext) Line(i int) ILineContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILineContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILineContext)
}

func (s *InstructionContext) LINE_COMMENT() antlr.TerminalNode {
	return s.GetToken(AsmE8ParserLINE_COMMENT, 0)
}

func (s *InstructionContext) SUFFIX_WHITESPACE() antlr.TerminalNode {
	return s.GetToken(AsmE8ParserSUFFIX_WHITESPACE, 0)
}

func (s *InstructionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InstructionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InstructionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsmE8Listener); ok {
		listenerT.EnterInstruction(s)
	}
}

func (s *InstructionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsmE8Listener); ok {
		listenerT.ExitInstruction(s)
	}
}

func (p *AsmE8Parser) Instruction() (localctx IInstructionContext) {
	localctx = NewInstructionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, AsmE8ParserRULE_instruction)
	var _la int

	p.SetState(121)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(106)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&9151314442816847918) != 0) || _la == AsmE8ParserSTR {
			{
				p.SetState(105)
				p.Line()
			}

			p.SetState(108)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		p.SetState(111)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == AsmE8ParserLINE_COMMENT {
			{
				p.SetState(110)
				p.Match(AsmE8ParserLINE_COMMENT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(114)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&9151314442816847918) != 0) || _la == AsmE8ParserSTR {
			{
				p.SetState(113)
				p.Line()
			}

			p.SetState(116)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		p.SetState(119)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == AsmE8ParserSUFFIX_WHITESPACE {
			{
				p.SetState(118)
				p.Match(AsmE8ParserSUFFIX_WHITESPACE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILineContext is an interface to support dynamic dispatch.
type ILineContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Variable_reference() IVariable_referenceContext
	Variable() IVariableContext
	Inst() IInstContext
	Segment() ISegmentContext
	Access() IAccessContext
	Label() ILabelContext
	Directives() IDirectivesContext

	// IsLineContext differentiates from other interfaces.
	IsLineContext()
}

type LineContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLineContext() *LineContext {
	var p = new(LineContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsmE8ParserRULE_line
	return p
}

func InitEmptyLineContext(p *LineContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsmE8ParserRULE_line
}

func (*LineContext) IsLineContext() {}

func NewLineContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LineContext {
	var p = new(LineContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AsmE8ParserRULE_line

	return p
}

func (s *LineContext) GetParser() antlr.Parser { return s.parser }

func (s *LineContext) Variable_reference() IVariable_referenceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVariable_referenceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVariable_referenceContext)
}

func (s *LineContext) Variable() IVariableContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVariableContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVariableContext)
}

func (s *LineContext) Inst() IInstContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInstContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInstContext)
}

func (s *LineContext) Segment() ISegmentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISegmentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISegmentContext)
}

func (s *LineContext) Access() IAccessContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAccessContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAccessContext)
}

func (s *LineContext) Label() ILabelContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILabelContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILabelContext)
}

func (s *LineContext) Directives() IDirectivesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDirectivesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDirectivesContext)
}

func (s *LineContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LineContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LineContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsmE8Listener); ok {
		listenerT.EnterLine(s)
	}
}

func (s *LineContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsmE8Listener); ok {
		listenerT.ExitLine(s)
	}
}

func (p *AsmE8Parser) Line() (localctx ILineContext) {
	localctx = NewLineContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, AsmE8ParserRULE_line)
	var _la int

	p.SetState(136)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(123)
			p.Variable_reference()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(124)
			p.Variable()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(125)
			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&14) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(126)
			p.Inst()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(127)
			p.Segment()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(128)
			p.Access()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(129)
			p.Label()
		}
		{
			p.SetState(130)
			p.Match(AsmE8ParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(131)
			p.Directives()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(133)
			p.Directives()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(134)
			p.Label()
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(135)
			p.Match(AsmE8ParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILabelContext is an interface to support dynamic dispatch.
type ILabelContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Tag() ITagContext

	// IsLabelContext differentiates from other interfaces.
	IsLabelContext()
}

type LabelContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLabelContext() *LabelContext {
	var p = new(LabelContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsmE8ParserRULE_label
	return p
}

func InitEmptyLabelContext(p *LabelContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsmE8ParserRULE_label
}

func (*LabelContext) IsLabelContext() {}

func NewLabelContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LabelContext {
	var p = new(LabelContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AsmE8ParserRULE_label

	return p
}

func (s *LabelContext) GetParser() antlr.Parser { return s.parser }

func (s *LabelContext) Tag() ITagContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITagContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITagContext)
}

func (s *LabelContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LabelContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LabelContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsmE8Listener); ok {
		listenerT.EnterLabel(s)
	}
}

func (s *LabelContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsmE8Listener); ok {
		listenerT.ExitLabel(s)
	}
}

func (p *AsmE8Parser) Label() (localctx ILabelContext) {
	localctx = NewLabelContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, AsmE8ParserRULE_label)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(138)
		p.Tag()
	}
	{
		p.SetState(139)
		p.Match(AsmE8ParserT__5)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IInstContext is an interface to support dynamic dispatch.
type IInstContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Addr_mode_imm16() IAddr_mode_imm16Context
	Addr_mode_imm16_tag() IAddr_mode_imm16_tagContext
	Addr_mode_reg_to_mem_indexed() IAddr_mode_reg_to_mem_indexedContext
	Addr_mode_reg_to_mem_direct() IAddr_mode_reg_to_mem_directContext
	Addr_mode_reg_to_mem_indirect_offset() IAddr_mode_reg_to_mem_indirect_offsetContext
	Addr_mode_reg_to_mem_reg16_indexed() IAddr_mode_reg_to_mem_reg16_indexedContext
	Addr_mode_reg_to_mem_indirect() IAddr_mode_reg_to_mem_indirectContext
	Addr_mode_implied_reg8() IAddr_mode_implied_reg8Context
	Addr_mode_reg8_imm8() IAddr_mode_reg8_imm8Context
	Addr_mode_reg8_imm8_tag() IAddr_mode_reg8_imm8_tagContext
	Addr_mode_mem_to_reg_indexed() IAddr_mode_mem_to_reg_indexedContext
	Addr_mode_mem_to_reg_direct() IAddr_mode_mem_to_reg_directContext
	Addr_mode_mem_to_reg_indirect_offset() IAddr_mode_mem_to_reg_indirect_offsetContext
	Addr_mode_mem_to_reg_reg16_indexed() IAddr_mode_mem_to_reg_reg16_indexedContext
	Addr_mode_mem_to_reg_indirect() IAddr_mode_mem_to_reg_indirectContext
	Addr_mode_reg8_reg8() IAddr_mode_reg8_reg8Context
	Addr_mode_implied() IAddr_mode_impliedContext
	Addr_mode_implied_reg16() IAddr_mode_implied_reg16Context
	Addr_mode_reg16_imm() IAddr_mode_reg16_immContext
	Addr_mode_reg16_imm_tag() IAddr_mode_reg16_imm_tagContext
	Addr_mode_reg16_reg16() IAddr_mode_reg16_reg16Context
	Addr_mode_reg16_stack() IAddr_mode_reg16_stackContext
	Addr_mode_implied_stack() IAddr_mode_implied_stackContext
	Addr_mode_stack_imm8() IAddr_mode_stack_imm8Context
	Addr_mode_stack_imm8_tag() IAddr_mode_stack_imm8_tagContext
	Addr_mode_stack_reg16() IAddr_mode_stack_reg16Context

	// IsInstContext differentiates from other interfaces.
	IsInstContext()
}

type InstContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInstContext() *InstContext {
	var p = new(InstContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsmE8ParserRULE_inst
	return p
}

func InitEmptyInstContext(p *InstContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsmE8ParserRULE_inst
}

func (*InstContext) IsInstContext() {}

func NewInstContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InstContext {
	var p = new(InstContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AsmE8ParserRULE_inst

	return p
}

func (s *InstContext) GetParser() antlr.Parser { return s.parser }

func (s *InstContext) Addr_mode_imm16() IAddr_mode_imm16Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAddr_mode_imm16Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAddr_mode_imm16Context)
}

func (s *InstContext) Addr_mode_imm16_tag() IAddr_mode_imm16_tagContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAddr_mode_imm16_tagContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAddr_mode_imm16_tagContext)
}

func (s *InstContext) Addr_mode_reg_to_mem_indexed() IAddr_mode_reg_to_mem_indexedContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAddr_mode_reg_to_mem_indexedContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAddr_mode_reg_to_mem_indexedContext)
}

func (s *InstContext) Addr_mode_reg_to_mem_direct() IAddr_mode_reg_to_mem_directContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAddr_mode_reg_to_mem_directContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAddr_mode_reg_to_mem_directContext)
}

func (s *InstContext) Addr_mode_reg_to_mem_indirect_offset() IAddr_mode_reg_to_mem_indirect_offsetContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAddr_mode_reg_to_mem_indirect_offsetContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAddr_mode_reg_to_mem_indirect_offsetContext)
}

func (s *InstContext) Addr_mode_reg_to_mem_reg16_indexed() IAddr_mode_reg_to_mem_reg16_indexedContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAddr_mode_reg_to_mem_reg16_indexedContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAddr_mode_reg_to_mem_reg16_indexedContext)
}

func (s *InstContext) Addr_mode_reg_to_mem_indirect() IAddr_mode_reg_to_mem_indirectContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAddr_mode_reg_to_mem_indirectContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAddr_mode_reg_to_mem_indirectContext)
}

func (s *InstContext) Addr_mode_implied_reg8() IAddr_mode_implied_reg8Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAddr_mode_implied_reg8Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAddr_mode_implied_reg8Context)
}

func (s *InstContext) Addr_mode_reg8_imm8() IAddr_mode_reg8_imm8Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAddr_mode_reg8_imm8Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAddr_mode_reg8_imm8Context)
}

func (s *InstContext) Addr_mode_reg8_imm8_tag() IAddr_mode_reg8_imm8_tagContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAddr_mode_reg8_imm8_tagContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAddr_mode_reg8_imm8_tagContext)
}

func (s *InstContext) Addr_mode_mem_to_reg_indexed() IAddr_mode_mem_to_reg_indexedContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAddr_mode_mem_to_reg_indexedContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAddr_mode_mem_to_reg_indexedContext)
}

func (s *InstContext) Addr_mode_mem_to_reg_direct() IAddr_mode_mem_to_reg_directContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAddr_mode_mem_to_reg_directContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAddr_mode_mem_to_reg_directContext)
}

func (s *InstContext) Addr_mode_mem_to_reg_indirect_offset() IAddr_mode_mem_to_reg_indirect_offsetContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAddr_mode_mem_to_reg_indirect_offsetContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAddr_mode_mem_to_reg_indirect_offsetContext)
}

func (s *InstContext) Addr_mode_mem_to_reg_reg16_indexed() IAddr_mode_mem_to_reg_reg16_indexedContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAddr_mode_mem_to_reg_reg16_indexedContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAddr_mode_mem_to_reg_reg16_indexedContext)
}

func (s *InstContext) Addr_mode_mem_to_reg_indirect() IAddr_mode_mem_to_reg_indirectContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAddr_mode_mem_to_reg_indirectContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAddr_mode_mem_to_reg_indirectContext)
}

func (s *InstContext) Addr_mode_reg8_reg8() IAddr_mode_reg8_reg8Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAddr_mode_reg8_reg8Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAddr_mode_reg8_reg8Context)
}

func (s *InstContext) Addr_mode_implied() IAddr_mode_impliedContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAddr_mode_impliedContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAddr_mode_impliedContext)
}

func (s *InstContext) Addr_mode_implied_reg16() IAddr_mode_implied_reg16Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAddr_mode_implied_reg16Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAddr_mode_implied_reg16Context)
}

func (s *InstContext) Addr_mode_reg16_imm() IAddr_mode_reg16_immContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAddr_mode_reg16_immContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAddr_mode_reg16_immContext)
}

func (s *InstContext) Addr_mode_reg16_imm_tag() IAddr_mode_reg16_imm_tagContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAddr_mode_reg16_imm_tagContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAddr_mode_reg16_imm_tagContext)
}

func (s *InstContext) Addr_mode_reg16_reg16() IAddr_mode_reg16_reg16Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAddr_mode_reg16_reg16Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAddr_mode_reg16_reg16Context)
}

func (s *InstContext) Addr_mode_reg16_stack() IAddr_mode_reg16_stackContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAddr_mode_reg16_stackContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAddr_mode_reg16_stackContext)
}

func (s *InstContext) Addr_mode_implied_stack() IAddr_mode_implied_stackContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAddr_mode_implied_stackContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAddr_mode_implied_stackContext)
}

func (s *InstContext) Addr_mode_stack_imm8() IAddr_mode_stack_imm8Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAddr_mode_stack_imm8Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAddr_mode_stack_imm8Context)
}

func (s *InstContext) Addr_mode_stack_imm8_tag() IAddr_mode_stack_imm8_tagContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAddr_mode_stack_imm8_tagContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAddr_mode_stack_imm8_tagContext)
}

func (s *InstContext) Addr_mode_stack_reg16() IAddr_mode_stack_reg16Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAddr_mode_stack_reg16Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAddr_mode_stack_reg16Context)
}

func (s *InstContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InstContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InstContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsmE8Listener); ok {
		listenerT.EnterInst(s)
	}
}

func (s *InstContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsmE8Listener); ok {
		listenerT.ExitInst(s)
	}
}

func (p *AsmE8Parser) Inst() (localctx IInstContext) {
	localctx = NewInstContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, AsmE8ParserRULE_inst)
	p.SetState(167)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(141)
			p.Addr_mode_imm16()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(142)
			p.Addr_mode_imm16_tag()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(143)
			p.Addr_mode_reg_to_mem_indexed()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(144)
			p.Addr_mode_reg_to_mem_direct()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(145)
			p.Addr_mode_reg_to_mem_indirect_offset()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(146)
			p.Addr_mode_reg_to_mem_reg16_indexed()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(147)
			p.Addr_mode_reg_to_mem_indirect()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(148)
			p.Addr_mode_implied_reg8()
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(149)
			p.Addr_mode_reg8_imm8()
		}

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(150)
			p.Addr_mode_reg8_imm8_tag()
		}

	case 11:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(151)
			p.Addr_mode_mem_to_reg_indexed()
		}

	case 12:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(152)
			p.Addr_mode_mem_to_reg_direct()
		}

	case 13:
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(153)
			p.Addr_mode_mem_to_reg_indirect_offset()
		}

	case 14:
		p.EnterOuterAlt(localctx, 14)
		{
			p.SetState(154)
			p.Addr_mode_mem_to_reg_reg16_indexed()
		}

	case 15:
		p.EnterOuterAlt(localctx, 15)
		{
			p.SetState(155)
			p.Addr_mode_mem_to_reg_indirect()
		}

	case 16:
		p.EnterOuterAlt(localctx, 16)
		{
			p.SetState(156)
			p.Addr_mode_reg8_reg8()
		}

	case 17:
		p.EnterOuterAlt(localctx, 17)
		{
			p.SetState(157)
			p.Addr_mode_implied()
		}

	case 18:
		p.EnterOuterAlt(localctx, 18)
		{
			p.SetState(158)
			p.Addr_mode_implied_reg16()
		}

	case 19:
		p.EnterOuterAlt(localctx, 19)
		{
			p.SetState(159)
			p.Addr_mode_reg16_imm()
		}

	case 20:
		p.EnterOuterAlt(localctx, 20)
		{
			p.SetState(160)
			p.Addr_mode_reg16_imm_tag()
		}

	case 21:
		p.EnterOuterAlt(localctx, 21)
		{
			p.SetState(161)
			p.Addr_mode_reg16_reg16()
		}

	case 22:
		p.EnterOuterAlt(localctx, 22)
		{
			p.SetState(162)
			p.Addr_mode_reg16_stack()
		}

	case 23:
		p.EnterOuterAlt(localctx, 23)
		{
			p.SetState(163)
			p.Addr_mode_implied_stack()
		}

	case 24:
		p.EnterOuterAlt(localctx, 24)
		{
			p.SetState(164)
			p.Addr_mode_stack_imm8()
		}

	case 25:
		p.EnterOuterAlt(localctx, 25)
		{
			p.SetState(165)
			p.Addr_mode_stack_imm8_tag()
		}

	case 26:
		p.EnterOuterAlt(localctx, 26)
		{
			p.SetState(166)
			p.Addr_mode_stack_reg16()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAddr_mode_imm16Context is an interface to support dynamic dispatch.
type IAddr_mode_imm16Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Mnemonic() IMnemonicContext
	Imm() IImmContext

	// IsAddr_mode_imm16Context differentiates from other interfaces.
	IsAddr_mode_imm16Context()
}

type Addr_mode_imm16Context struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAddr_mode_imm16Context() *Addr_mode_imm16Context {
	var p = new(Addr_mode_imm16Context)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsmE8ParserRULE_addr_mode_imm16
	return p
}

func InitEmptyAddr_mode_imm16Context(p *Addr_mode_imm16Context) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsmE8ParserRULE_addr_mode_imm16
}

func (*Addr_mode_imm16Context) IsAddr_mode_imm16Context() {}

func NewAddr_mode_imm16Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Addr_mode_imm16Context {
	var p = new(Addr_mode_imm16Context)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AsmE8ParserRULE_addr_mode_imm16

	return p
}

func (s *Addr_mode_imm16Context) GetParser() antlr.Parser { return s.parser }

func (s *Addr_mode_imm16Context) Mnemonic() IMnemonicContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMnemonicContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMnemonicContext)
}

func (s *Addr_mode_imm16Context) Imm() IImmContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IImmContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IImmContext)
}

func (s *Addr_mode_imm16Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Addr_mode_imm16Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Addr_mode_imm16Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsmE8Listener); ok {
		listenerT.EnterAddr_mode_imm16(s)
	}
}

func (s *Addr_mode_imm16Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsmE8Listener); ok {
		listenerT.ExitAddr_mode_imm16(s)
	}
}

func (p *AsmE8Parser) Addr_mode_imm16() (localctx IAddr_mode_imm16Context) {
	localctx = NewAddr_mode_imm16Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, AsmE8ParserRULE_addr_mode_imm16)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(169)
		p.Mnemonic()
	}
	{
		p.SetState(170)
		p.Match(AsmE8ParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(171)
		p.Imm()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAddr_mode_imm16_tagContext is an interface to support dynamic dispatch.
type IAddr_mode_imm16_tagContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Mnemonic() IMnemonicContext
	Tag() ITagContext

	// IsAddr_mode_imm16_tagContext differentiates from other interfaces.
	IsAddr_mode_imm16_tagContext()
}

type Addr_mode_imm16_tagContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAddr_mode_imm16_tagContext() *Addr_mode_imm16_tagContext {
	var p = new(Addr_mode_imm16_tagContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsmE8ParserRULE_addr_mode_imm16_tag
	return p
}

func InitEmptyAddr_mode_imm16_tagContext(p *Addr_mode_imm16_tagContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsmE8ParserRULE_addr_mode_imm16_tag
}

func (*Addr_mode_imm16_tagContext) IsAddr_mode_imm16_tagContext() {}

func NewAddr_mode_imm16_tagContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Addr_mode_imm16_tagContext {
	var p = new(Addr_mode_imm16_tagContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AsmE8ParserRULE_addr_mode_imm16_tag

	return p
}

func (s *Addr_mode_imm16_tagContext) GetParser() antlr.Parser { return s.parser }

func (s *Addr_mode_imm16_tagContext) Mnemonic() IMnemonicContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMnemonicContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMnemonicContext)
}

func (s *Addr_mode_imm16_tagContext) Tag() ITagContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITagContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITagContext)
}

func (s *Addr_mode_imm16_tagContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Addr_mode_imm16_tagContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Addr_mode_imm16_tagContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsmE8Listener); ok {
		listenerT.EnterAddr_mode_imm16_tag(s)
	}
}

func (s *Addr_mode_imm16_tagContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsmE8Listener); ok {
		listenerT.ExitAddr_mode_imm16_tag(s)
	}
}

func (p *AsmE8Parser) Addr_mode_imm16_tag() (localctx IAddr_mode_imm16_tagContext) {
	localctx = NewAddr_mode_imm16_tagContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, AsmE8ParserRULE_addr_mode_imm16_tag)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(173)
		p.Mnemonic()
	}
	{
		p.SetState(174)
		p.Match(AsmE8ParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(175)
		p.Tag()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAddr_mode_reg_to_mem_indexedContext is an interface to support dynamic dispatch.
type IAddr_mode_reg_to_mem_indexedContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Mnemonic() IMnemonicContext
	Direct_offset() IDirect_offsetContext
	Reg8() IReg8Context

	// IsAddr_mode_reg_to_mem_indexedContext differentiates from other interfaces.
	IsAddr_mode_reg_to_mem_indexedContext()
}

type Addr_mode_reg_to_mem_indexedContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAddr_mode_reg_to_mem_indexedContext() *Addr_mode_reg_to_mem_indexedContext {
	var p = new(Addr_mode_reg_to_mem_indexedContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsmE8ParserRULE_addr_mode_reg_to_mem_indexed
	return p
}

func InitEmptyAddr_mode_reg_to_mem_indexedContext(p *Addr_mode_reg_to_mem_indexedContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsmE8ParserRULE_addr_mode_reg_to_mem_indexed
}

func (*Addr_mode_reg_to_mem_indexedContext) IsAddr_mode_reg_to_mem_indexedContext() {}

func NewAddr_mode_reg_to_mem_indexedContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Addr_mode_reg_to_mem_indexedContext {
	var p = new(Addr_mode_reg_to_mem_indexedContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AsmE8ParserRULE_addr_mode_reg_to_mem_indexed

	return p
}

func (s *Addr_mode_reg_to_mem_indexedContext) GetParser() antlr.Parser { return s.parser }

func (s *Addr_mode_reg_to_mem_indexedContext) Mnemonic() IMnemonicContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMnemonicContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMnemonicContext)
}

func (s *Addr_mode_reg_to_mem_indexedContext) Direct_offset() IDirect_offsetContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDirect_offsetContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDirect_offsetContext)
}

func (s *Addr_mode_reg_to_mem_indexedContext) Reg8() IReg8Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReg8Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReg8Context)
}

func (s *Addr_mode_reg_to_mem_indexedContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Addr_mode_reg_to_mem_indexedContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Addr_mode_reg_to_mem_indexedContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsmE8Listener); ok {
		listenerT.EnterAddr_mode_reg_to_mem_indexed(s)
	}
}

func (s *Addr_mode_reg_to_mem_indexedContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsmE8Listener); ok {
		listenerT.ExitAddr_mode_reg_to_mem_indexed(s)
	}
}

func (p *AsmE8Parser) Addr_mode_reg_to_mem_indexed() (localctx IAddr_mode_reg_to_mem_indexedContext) {
	localctx = NewAddr_mode_reg_to_mem_indexedContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, AsmE8ParserRULE_addr_mode_reg_to_mem_indexed)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(177)
		p.Mnemonic()
	}
	{
		p.SetState(178)
		p.Match(AsmE8ParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(179)
		p.Direct_offset()
	}
	{
		p.SetState(180)
		p.Match(AsmE8ParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(181)
		p.Reg8()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAddr_mode_reg_to_mem_directContext is an interface to support dynamic dispatch.
type IAddr_mode_reg_to_mem_directContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Mnemonic() IMnemonicContext
	Direct() IDirectContext
	Reg8() IReg8Context

	// IsAddr_mode_reg_to_mem_directContext differentiates from other interfaces.
	IsAddr_mode_reg_to_mem_directContext()
}

type Addr_mode_reg_to_mem_directContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAddr_mode_reg_to_mem_directContext() *Addr_mode_reg_to_mem_directContext {
	var p = new(Addr_mode_reg_to_mem_directContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsmE8ParserRULE_addr_mode_reg_to_mem_direct
	return p
}

func InitEmptyAddr_mode_reg_to_mem_directContext(p *Addr_mode_reg_to_mem_directContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsmE8ParserRULE_addr_mode_reg_to_mem_direct
}

func (*Addr_mode_reg_to_mem_directContext) IsAddr_mode_reg_to_mem_directContext() {}

func NewAddr_mode_reg_to_mem_directContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Addr_mode_reg_to_mem_directContext {
	var p = new(Addr_mode_reg_to_mem_directContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AsmE8ParserRULE_addr_mode_reg_to_mem_direct

	return p
}

func (s *Addr_mode_reg_to_mem_directContext) GetParser() antlr.Parser { return s.parser }

func (s *Addr_mode_reg_to_mem_directContext) Mnemonic() IMnemonicContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMnemonicContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMnemonicContext)
}

func (s *Addr_mode_reg_to_mem_directContext) Direct() IDirectContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDirectContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDirectContext)
}

func (s *Addr_mode_reg_to_mem_directContext) Reg8() IReg8Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReg8Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReg8Context)
}

func (s *Addr_mode_reg_to_mem_directContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Addr_mode_reg_to_mem_directContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Addr_mode_reg_to_mem_directContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsmE8Listener); ok {
		listenerT.EnterAddr_mode_reg_to_mem_direct(s)
	}
}

func (s *Addr_mode_reg_to_mem_directContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsmE8Listener); ok {
		listenerT.ExitAddr_mode_reg_to_mem_direct(s)
	}
}

func (p *AsmE8Parser) Addr_mode_reg_to_mem_direct() (localctx IAddr_mode_reg_to_mem_directContext) {
	localctx = NewAddr_mode_reg_to_mem_directContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, AsmE8ParserRULE_addr_mode_reg_to_mem_direct)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(183)
		p.Mnemonic()
	}
	{
		p.SetState(184)
		p.Match(AsmE8ParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(185)
		p.Direct()
	}
	{
		p.SetState(186)
		p.Match(AsmE8ParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(187)
		p.Reg8()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAddr_mode_reg_to_mem_indirect_offsetContext is an interface to support dynamic dispatch.
type IAddr_mode_reg_to_mem_indirect_offsetContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Mnemonic() IMnemonicContext
	Indirect_offset() IIndirect_offsetContext
	Reg8() IReg8Context

	// IsAddr_mode_reg_to_mem_indirect_offsetContext differentiates from other interfaces.
	IsAddr_mode_reg_to_mem_indirect_offsetContext()
}

type Addr_mode_reg_to_mem_indirect_offsetContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAddr_mode_reg_to_mem_indirect_offsetContext() *Addr_mode_reg_to_mem_indirect_offsetContext {
	var p = new(Addr_mode_reg_to_mem_indirect_offsetContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsmE8ParserRULE_addr_mode_reg_to_mem_indirect_offset
	return p
}

func InitEmptyAddr_mode_reg_to_mem_indirect_offsetContext(p *Addr_mode_reg_to_mem_indirect_offsetContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsmE8ParserRULE_addr_mode_reg_to_mem_indirect_offset
}

func (*Addr_mode_reg_to_mem_indirect_offsetContext) IsAddr_mode_reg_to_mem_indirect_offsetContext() {}

func NewAddr_mode_reg_to_mem_indirect_offsetContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Addr_mode_reg_to_mem_indirect_offsetContext {
	var p = new(Addr_mode_reg_to_mem_indirect_offsetContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AsmE8ParserRULE_addr_mode_reg_to_mem_indirect_offset

	return p
}

func (s *Addr_mode_reg_to_mem_indirect_offsetContext) GetParser() antlr.Parser { return s.parser }

func (s *Addr_mode_reg_to_mem_indirect_offsetContext) Mnemonic() IMnemonicContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMnemonicContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMnemonicContext)
}

func (s *Addr_mode_reg_to_mem_indirect_offsetContext) Indirect_offset() IIndirect_offsetContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIndirect_offsetContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIndirect_offsetContext)
}

func (s *Addr_mode_reg_to_mem_indirect_offsetContext) Reg8() IReg8Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReg8Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReg8Context)
}

func (s *Addr_mode_reg_to_mem_indirect_offsetContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Addr_mode_reg_to_mem_indirect_offsetContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Addr_mode_reg_to_mem_indirect_offsetContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsmE8Listener); ok {
		listenerT.EnterAddr_mode_reg_to_mem_indirect_offset(s)
	}
}

func (s *Addr_mode_reg_to_mem_indirect_offsetContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsmE8Listener); ok {
		listenerT.ExitAddr_mode_reg_to_mem_indirect_offset(s)
	}
}

func (p *AsmE8Parser) Addr_mode_reg_to_mem_indirect_offset() (localctx IAddr_mode_reg_to_mem_indirect_offsetContext) {
	localctx = NewAddr_mode_reg_to_mem_indirect_offsetContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, AsmE8ParserRULE_addr_mode_reg_to_mem_indirect_offset)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(189)
		p.Mnemonic()
	}
	{
		p.SetState(190)
		p.Match(AsmE8ParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(191)
		p.Indirect_offset()
	}
	{
		p.SetState(192)
		p.Match(AsmE8ParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(193)
		p.Reg8()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAddr_mode_reg_to_mem_reg16_indexedContext is an interface to support dynamic dispatch.
type IAddr_mode_reg_to_mem_reg16_indexedContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Mnemonic() IMnemonicContext
	Reg16_indexed() IReg16_indexedContext
	Reg8() IReg8Context

	// IsAddr_mode_reg_to_mem_reg16_indexedContext differentiates from other interfaces.
	IsAddr_mode_reg_to_mem_reg16_indexedContext()
}

type Addr_mode_reg_to_mem_reg16_indexedContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAddr_mode_reg_to_mem_reg16_indexedContext() *Addr_mode_reg_to_mem_reg16_indexedContext {
	var p = new(Addr_mode_reg_to_mem_reg16_indexedContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsmE8ParserRULE_addr_mode_reg_to_mem_reg16_indexed
	return p
}

func InitEmptyAddr_mode_reg_to_mem_reg16_indexedContext(p *Addr_mode_reg_to_mem_reg16_indexedContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsmE8ParserRULE_addr_mode_reg_to_mem_reg16_indexed
}

func (*Addr_mode_reg_to_mem_reg16_indexedContext) IsAddr_mode_reg_to_mem_reg16_indexedContext() {}

func NewAddr_mode_reg_to_mem_reg16_indexedContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Addr_mode_reg_to_mem_reg16_indexedContext {
	var p = new(Addr_mode_reg_to_mem_reg16_indexedContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AsmE8ParserRULE_addr_mode_reg_to_mem_reg16_indexed

	return p
}

func (s *Addr_mode_reg_to_mem_reg16_indexedContext) GetParser() antlr.Parser { return s.parser }

func (s *Addr_mode_reg_to_mem_reg16_indexedContext) Mnemonic() IMnemonicContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMnemonicContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMnemonicContext)
}

func (s *Addr_mode_reg_to_mem_reg16_indexedContext) Reg16_indexed() IReg16_indexedContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReg16_indexedContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReg16_indexedContext)
}

func (s *Addr_mode_reg_to_mem_reg16_indexedContext) Reg8() IReg8Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReg8Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReg8Context)
}

func (s *Addr_mode_reg_to_mem_reg16_indexedContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Addr_mode_reg_to_mem_reg16_indexedContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Addr_mode_reg_to_mem_reg16_indexedContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsmE8Listener); ok {
		listenerT.EnterAddr_mode_reg_to_mem_reg16_indexed(s)
	}
}

func (s *Addr_mode_reg_to_mem_reg16_indexedContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsmE8Listener); ok {
		listenerT.ExitAddr_mode_reg_to_mem_reg16_indexed(s)
	}
}

func (p *AsmE8Parser) Addr_mode_reg_to_mem_reg16_indexed() (localctx IAddr_mode_reg_to_mem_reg16_indexedContext) {
	localctx = NewAddr_mode_reg_to_mem_reg16_indexedContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, AsmE8ParserRULE_addr_mode_reg_to_mem_reg16_indexed)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(195)
		p.Mnemonic()
	}
	{
		p.SetState(196)
		p.Match(AsmE8ParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(197)
		p.Reg16_indexed()
	}
	{
		p.SetState(198)
		p.Match(AsmE8ParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(199)
		p.Reg8()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAddr_mode_reg_to_mem_indirectContext is an interface to support dynamic dispatch.
type IAddr_mode_reg_to_mem_indirectContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Mnemonic() IMnemonicContext
	Indirect() IIndirectContext
	Reg8() IReg8Context

	// IsAddr_mode_reg_to_mem_indirectContext differentiates from other interfaces.
	IsAddr_mode_reg_to_mem_indirectContext()
}

type Addr_mode_reg_to_mem_indirectContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAddr_mode_reg_to_mem_indirectContext() *Addr_mode_reg_to_mem_indirectContext {
	var p = new(Addr_mode_reg_to_mem_indirectContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsmE8ParserRULE_addr_mode_reg_to_mem_indirect
	return p
}

func InitEmptyAddr_mode_reg_to_mem_indirectContext(p *Addr_mode_reg_to_mem_indirectContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsmE8ParserRULE_addr_mode_reg_to_mem_indirect
}

func (*Addr_mode_reg_to_mem_indirectContext) IsAddr_mode_reg_to_mem_indirectContext() {}

func NewAddr_mode_reg_to_mem_indirectContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Addr_mode_reg_to_mem_indirectContext {
	var p = new(Addr_mode_reg_to_mem_indirectContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AsmE8ParserRULE_addr_mode_reg_to_mem_indirect

	return p
}

func (s *Addr_mode_reg_to_mem_indirectContext) GetParser() antlr.Parser { return s.parser }

func (s *Addr_mode_reg_to_mem_indirectContext) Mnemonic() IMnemonicContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMnemonicContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMnemonicContext)
}

func (s *Addr_mode_reg_to_mem_indirectContext) Indirect() IIndirectContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIndirectContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIndirectContext)
}

func (s *Addr_mode_reg_to_mem_indirectContext) Reg8() IReg8Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReg8Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReg8Context)
}

func (s *Addr_mode_reg_to_mem_indirectContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Addr_mode_reg_to_mem_indirectContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Addr_mode_reg_to_mem_indirectContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsmE8Listener); ok {
		listenerT.EnterAddr_mode_reg_to_mem_indirect(s)
	}
}

func (s *Addr_mode_reg_to_mem_indirectContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsmE8Listener); ok {
		listenerT.ExitAddr_mode_reg_to_mem_indirect(s)
	}
}

func (p *AsmE8Parser) Addr_mode_reg_to_mem_indirect() (localctx IAddr_mode_reg_to_mem_indirectContext) {
	localctx = NewAddr_mode_reg_to_mem_indirectContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, AsmE8ParserRULE_addr_mode_reg_to_mem_indirect)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(201)
		p.Mnemonic()
	}
	{
		p.SetState(202)
		p.Match(AsmE8ParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(203)
		p.Indirect()
	}
	{
		p.SetState(204)
		p.Match(AsmE8ParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(205)
		p.Reg8()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAddr_mode_implied_reg8Context is an interface to support dynamic dispatch.
type IAddr_mode_implied_reg8Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Mnemonic() IMnemonicContext
	Reg8() IReg8Context

	// IsAddr_mode_implied_reg8Context differentiates from other interfaces.
	IsAddr_mode_implied_reg8Context()
}

type Addr_mode_implied_reg8Context struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAddr_mode_implied_reg8Context() *Addr_mode_implied_reg8Context {
	var p = new(Addr_mode_implied_reg8Context)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsmE8ParserRULE_addr_mode_implied_reg8
	return p
}

func InitEmptyAddr_mode_implied_reg8Context(p *Addr_mode_implied_reg8Context) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsmE8ParserRULE_addr_mode_implied_reg8
}

func (*Addr_mode_implied_reg8Context) IsAddr_mode_implied_reg8Context() {}

func NewAddr_mode_implied_reg8Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Addr_mode_implied_reg8Context {
	var p = new(Addr_mode_implied_reg8Context)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AsmE8ParserRULE_addr_mode_implied_reg8

	return p
}

func (s *Addr_mode_implied_reg8Context) GetParser() antlr.Parser { return s.parser }

func (s *Addr_mode_implied_reg8Context) Mnemonic() IMnemonicContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMnemonicContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMnemonicContext)
}

func (s *Addr_mode_implied_reg8Context) Reg8() IReg8Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReg8Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReg8Context)
}

func (s *Addr_mode_implied_reg8Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Addr_mode_implied_reg8Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Addr_mode_implied_reg8Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsmE8Listener); ok {
		listenerT.EnterAddr_mode_implied_reg8(s)
	}
}

func (s *Addr_mode_implied_reg8Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsmE8Listener); ok {
		listenerT.ExitAddr_mode_implied_reg8(s)
	}
}

func (p *AsmE8Parser) Addr_mode_implied_reg8() (localctx IAddr_mode_implied_reg8Context) {
	localctx = NewAddr_mode_implied_reg8Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, AsmE8ParserRULE_addr_mode_implied_reg8)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(207)
		p.Mnemonic()
	}
	{
		p.SetState(208)
		p.Match(AsmE8ParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(209)
		p.Reg8()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAddr_mode_reg8_imm8Context is an interface to support dynamic dispatch.
type IAddr_mode_reg8_imm8Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Mnemonic() IMnemonicContext
	Reg8() IReg8Context
	Imm() IImmContext

	// IsAddr_mode_reg8_imm8Context differentiates from other interfaces.
	IsAddr_mode_reg8_imm8Context()
}

type Addr_mode_reg8_imm8Context struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAddr_mode_reg8_imm8Context() *Addr_mode_reg8_imm8Context {
	var p = new(Addr_mode_reg8_imm8Context)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsmE8ParserRULE_addr_mode_reg8_imm8
	return p
}

func InitEmptyAddr_mode_reg8_imm8Context(p *Addr_mode_reg8_imm8Context) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsmE8ParserRULE_addr_mode_reg8_imm8
}

func (*Addr_mode_reg8_imm8Context) IsAddr_mode_reg8_imm8Context() {}

func NewAddr_mode_reg8_imm8Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Addr_mode_reg8_imm8Context {
	var p = new(Addr_mode_reg8_imm8Context)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AsmE8ParserRULE_addr_mode_reg8_imm8

	return p
}

func (s *Addr_mode_reg8_imm8Context) GetParser() antlr.Parser { return s.parser }

func (s *Addr_mode_reg8_imm8Context) Mnemonic() IMnemonicContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMnemonicContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMnemonicContext)
}

func (s *Addr_mode_reg8_imm8Context) Reg8() IReg8Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReg8Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReg8Context)
}

func (s *Addr_mode_reg8_imm8Context) Imm() IImmContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IImmContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IImmContext)
}

func (s *Addr_mode_reg8_imm8Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Addr_mode_reg8_imm8Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Addr_mode_reg8_imm8Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsmE8Listener); ok {
		listenerT.EnterAddr_mode_reg8_imm8(s)
	}
}

func (s *Addr_mode_reg8_imm8Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsmE8Listener); ok {
		listenerT.ExitAddr_mode_reg8_imm8(s)
	}
}

func (p *AsmE8Parser) Addr_mode_reg8_imm8() (localctx IAddr_mode_reg8_imm8Context) {
	localctx = NewAddr_mode_reg8_imm8Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, AsmE8ParserRULE_addr_mode_reg8_imm8)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(211)
		p.Mnemonic()
	}
	{
		p.SetState(212)
		p.Match(AsmE8ParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(213)
		p.Reg8()
	}
	{
		p.SetState(214)
		p.Match(AsmE8ParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(215)
		p.Imm()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAddr_mode_reg8_imm8_tagContext is an interface to support dynamic dispatch.
type IAddr_mode_reg8_imm8_tagContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Mnemonic() IMnemonicContext
	Reg8() IReg8Context
	Tag() ITagContext

	// IsAddr_mode_reg8_imm8_tagContext differentiates from other interfaces.
	IsAddr_mode_reg8_imm8_tagContext()
}

type Addr_mode_reg8_imm8_tagContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAddr_mode_reg8_imm8_tagContext() *Addr_mode_reg8_imm8_tagContext {
	var p = new(Addr_mode_reg8_imm8_tagContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsmE8ParserRULE_addr_mode_reg8_imm8_tag
	return p
}

func InitEmptyAddr_mode_reg8_imm8_tagContext(p *Addr_mode_reg8_imm8_tagContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsmE8ParserRULE_addr_mode_reg8_imm8_tag
}

func (*Addr_mode_reg8_imm8_tagContext) IsAddr_mode_reg8_imm8_tagContext() {}

func NewAddr_mode_reg8_imm8_tagContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Addr_mode_reg8_imm8_tagContext {
	var p = new(Addr_mode_reg8_imm8_tagContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AsmE8ParserRULE_addr_mode_reg8_imm8_tag

	return p
}

func (s *Addr_mode_reg8_imm8_tagContext) GetParser() antlr.Parser { return s.parser }

func (s *Addr_mode_reg8_imm8_tagContext) Mnemonic() IMnemonicContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMnemonicContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMnemonicContext)
}

func (s *Addr_mode_reg8_imm8_tagContext) Reg8() IReg8Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReg8Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReg8Context)
}

func (s *Addr_mode_reg8_imm8_tagContext) Tag() ITagContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITagContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITagContext)
}

func (s *Addr_mode_reg8_imm8_tagContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Addr_mode_reg8_imm8_tagContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Addr_mode_reg8_imm8_tagContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsmE8Listener); ok {
		listenerT.EnterAddr_mode_reg8_imm8_tag(s)
	}
}

func (s *Addr_mode_reg8_imm8_tagContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsmE8Listener); ok {
		listenerT.ExitAddr_mode_reg8_imm8_tag(s)
	}
}

func (p *AsmE8Parser) Addr_mode_reg8_imm8_tag() (localctx IAddr_mode_reg8_imm8_tagContext) {
	localctx = NewAddr_mode_reg8_imm8_tagContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, AsmE8ParserRULE_addr_mode_reg8_imm8_tag)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(217)
		p.Mnemonic()
	}
	{
		p.SetState(218)
		p.Match(AsmE8ParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(219)
		p.Reg8()
	}
	{
		p.SetState(220)
		p.Match(AsmE8ParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(221)
		p.Tag()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAddr_mode_mem_to_reg_indexedContext is an interface to support dynamic dispatch.
type IAddr_mode_mem_to_reg_indexedContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Mnemonic() IMnemonicContext
	Reg8() IReg8Context
	Direct_offset() IDirect_offsetContext

	// IsAddr_mode_mem_to_reg_indexedContext differentiates from other interfaces.
	IsAddr_mode_mem_to_reg_indexedContext()
}

type Addr_mode_mem_to_reg_indexedContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAddr_mode_mem_to_reg_indexedContext() *Addr_mode_mem_to_reg_indexedContext {
	var p = new(Addr_mode_mem_to_reg_indexedContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsmE8ParserRULE_addr_mode_mem_to_reg_indexed
	return p
}

func InitEmptyAddr_mode_mem_to_reg_indexedContext(p *Addr_mode_mem_to_reg_indexedContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsmE8ParserRULE_addr_mode_mem_to_reg_indexed
}

func (*Addr_mode_mem_to_reg_indexedContext) IsAddr_mode_mem_to_reg_indexedContext() {}

func NewAddr_mode_mem_to_reg_indexedContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Addr_mode_mem_to_reg_indexedContext {
	var p = new(Addr_mode_mem_to_reg_indexedContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AsmE8ParserRULE_addr_mode_mem_to_reg_indexed

	return p
}

func (s *Addr_mode_mem_to_reg_indexedContext) GetParser() antlr.Parser { return s.parser }

func (s *Addr_mode_mem_to_reg_indexedContext) Mnemonic() IMnemonicContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMnemonicContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMnemonicContext)
}

func (s *Addr_mode_mem_to_reg_indexedContext) Reg8() IReg8Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReg8Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReg8Context)
}

func (s *Addr_mode_mem_to_reg_indexedContext) Direct_offset() IDirect_offsetContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDirect_offsetContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDirect_offsetContext)
}

func (s *Addr_mode_mem_to_reg_indexedContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Addr_mode_mem_to_reg_indexedContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Addr_mode_mem_to_reg_indexedContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsmE8Listener); ok {
		listenerT.EnterAddr_mode_mem_to_reg_indexed(s)
	}
}

func (s *Addr_mode_mem_to_reg_indexedContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsmE8Listener); ok {
		listenerT.ExitAddr_mode_mem_to_reg_indexed(s)
	}
}

func (p *AsmE8Parser) Addr_mode_mem_to_reg_indexed() (localctx IAddr_mode_mem_to_reg_indexedContext) {
	localctx = NewAddr_mode_mem_to_reg_indexedContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, AsmE8ParserRULE_addr_mode_mem_to_reg_indexed)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(223)
		p.Mnemonic()
	}
	{
		p.SetState(224)
		p.Match(AsmE8ParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(225)
		p.Reg8()
	}
	{
		p.SetState(226)
		p.Match(AsmE8ParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(227)
		p.Direct_offset()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAddr_mode_mem_to_reg_directContext is an interface to support dynamic dispatch.
type IAddr_mode_mem_to_reg_directContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Mnemonic() IMnemonicContext
	Reg8() IReg8Context
	Direct() IDirectContext

	// IsAddr_mode_mem_to_reg_directContext differentiates from other interfaces.
	IsAddr_mode_mem_to_reg_directContext()
}

type Addr_mode_mem_to_reg_directContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAddr_mode_mem_to_reg_directContext() *Addr_mode_mem_to_reg_directContext {
	var p = new(Addr_mode_mem_to_reg_directContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsmE8ParserRULE_addr_mode_mem_to_reg_direct
	return p
}

func InitEmptyAddr_mode_mem_to_reg_directContext(p *Addr_mode_mem_to_reg_directContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsmE8ParserRULE_addr_mode_mem_to_reg_direct
}

func (*Addr_mode_mem_to_reg_directContext) IsAddr_mode_mem_to_reg_directContext() {}

func NewAddr_mode_mem_to_reg_directContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Addr_mode_mem_to_reg_directContext {
	var p = new(Addr_mode_mem_to_reg_directContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AsmE8ParserRULE_addr_mode_mem_to_reg_direct

	return p
}

func (s *Addr_mode_mem_to_reg_directContext) GetParser() antlr.Parser { return s.parser }

func (s *Addr_mode_mem_to_reg_directContext) Mnemonic() IMnemonicContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMnemonicContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMnemonicContext)
}

func (s *Addr_mode_mem_to_reg_directContext) Reg8() IReg8Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReg8Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReg8Context)
}

func (s *Addr_mode_mem_to_reg_directContext) Direct() IDirectContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDirectContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDirectContext)
}

func (s *Addr_mode_mem_to_reg_directContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Addr_mode_mem_to_reg_directContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Addr_mode_mem_to_reg_directContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsmE8Listener); ok {
		listenerT.EnterAddr_mode_mem_to_reg_direct(s)
	}
}

func (s *Addr_mode_mem_to_reg_directContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsmE8Listener); ok {
		listenerT.ExitAddr_mode_mem_to_reg_direct(s)
	}
}

func (p *AsmE8Parser) Addr_mode_mem_to_reg_direct() (localctx IAddr_mode_mem_to_reg_directContext) {
	localctx = NewAddr_mode_mem_to_reg_directContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, AsmE8ParserRULE_addr_mode_mem_to_reg_direct)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(229)
		p.Mnemonic()
	}
	{
		p.SetState(230)
		p.Match(AsmE8ParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(231)
		p.Reg8()
	}
	{
		p.SetState(232)
		p.Match(AsmE8ParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(233)
		p.Direct()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAddr_mode_mem_to_reg_indirect_offsetContext is an interface to support dynamic dispatch.
type IAddr_mode_mem_to_reg_indirect_offsetContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Mnemonic() IMnemonicContext
	Reg8() IReg8Context
	Indirect_offset() IIndirect_offsetContext

	// IsAddr_mode_mem_to_reg_indirect_offsetContext differentiates from other interfaces.
	IsAddr_mode_mem_to_reg_indirect_offsetContext()
}

type Addr_mode_mem_to_reg_indirect_offsetContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAddr_mode_mem_to_reg_indirect_offsetContext() *Addr_mode_mem_to_reg_indirect_offsetContext {
	var p = new(Addr_mode_mem_to_reg_indirect_offsetContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsmE8ParserRULE_addr_mode_mem_to_reg_indirect_offset
	return p
}

func InitEmptyAddr_mode_mem_to_reg_indirect_offsetContext(p *Addr_mode_mem_to_reg_indirect_offsetContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsmE8ParserRULE_addr_mode_mem_to_reg_indirect_offset
}

func (*Addr_mode_mem_to_reg_indirect_offsetContext) IsAddr_mode_mem_to_reg_indirect_offsetContext() {}

func NewAddr_mode_mem_to_reg_indirect_offsetContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Addr_mode_mem_to_reg_indirect_offsetContext {
	var p = new(Addr_mode_mem_to_reg_indirect_offsetContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AsmE8ParserRULE_addr_mode_mem_to_reg_indirect_offset

	return p
}

func (s *Addr_mode_mem_to_reg_indirect_offsetContext) GetParser() antlr.Parser { return s.parser }

func (s *Addr_mode_mem_to_reg_indirect_offsetContext) Mnemonic() IMnemonicContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMnemonicContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMnemonicContext)
}

func (s *Addr_mode_mem_to_reg_indirect_offsetContext) Reg8() IReg8Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReg8Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReg8Context)
}

func (s *Addr_mode_mem_to_reg_indirect_offsetContext) Indirect_offset() IIndirect_offsetContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIndirect_offsetContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIndirect_offsetContext)
}

func (s *Addr_mode_mem_to_reg_indirect_offsetContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Addr_mode_mem_to_reg_indirect_offsetContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Addr_mode_mem_to_reg_indirect_offsetContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsmE8Listener); ok {
		listenerT.EnterAddr_mode_mem_to_reg_indirect_offset(s)
	}
}

func (s *Addr_mode_mem_to_reg_indirect_offsetContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsmE8Listener); ok {
		listenerT.ExitAddr_mode_mem_to_reg_indirect_offset(s)
	}
}

func (p *AsmE8Parser) Addr_mode_mem_to_reg_indirect_offset() (localctx IAddr_mode_mem_to_reg_indirect_offsetContext) {
	localctx = NewAddr_mode_mem_to_reg_indirect_offsetContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, AsmE8ParserRULE_addr_mode_mem_to_reg_indirect_offset)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(235)
		p.Mnemonic()
	}
	{
		p.SetState(236)
		p.Match(AsmE8ParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(237)
		p.Reg8()
	}
	{
		p.SetState(238)
		p.Match(AsmE8ParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(239)
		p.Indirect_offset()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAddr_mode_mem_to_reg_reg16_indexedContext is an interface to support dynamic dispatch.
type IAddr_mode_mem_to_reg_reg16_indexedContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Mnemonic() IMnemonicContext
	Reg8() IReg8Context
	Reg16_indexed() IReg16_indexedContext

	// IsAddr_mode_mem_to_reg_reg16_indexedContext differentiates from other interfaces.
	IsAddr_mode_mem_to_reg_reg16_indexedContext()
}

type Addr_mode_mem_to_reg_reg16_indexedContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAddr_mode_mem_to_reg_reg16_indexedContext() *Addr_mode_mem_to_reg_reg16_indexedContext {
	var p = new(Addr_mode_mem_to_reg_reg16_indexedContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsmE8ParserRULE_addr_mode_mem_to_reg_reg16_indexed
	return p
}

func InitEmptyAddr_mode_mem_to_reg_reg16_indexedContext(p *Addr_mode_mem_to_reg_reg16_indexedContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsmE8ParserRULE_addr_mode_mem_to_reg_reg16_indexed
}

func (*Addr_mode_mem_to_reg_reg16_indexedContext) IsAddr_mode_mem_to_reg_reg16_indexedContext() {}

func NewAddr_mode_mem_to_reg_reg16_indexedContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Addr_mode_mem_to_reg_reg16_indexedContext {
	var p = new(Addr_mode_mem_to_reg_reg16_indexedContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AsmE8ParserRULE_addr_mode_mem_to_reg_reg16_indexed

	return p
}

func (s *Addr_mode_mem_to_reg_reg16_indexedContext) GetParser() antlr.Parser { return s.parser }

func (s *Addr_mode_mem_to_reg_reg16_indexedContext) Mnemonic() IMnemonicContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMnemonicContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMnemonicContext)
}

func (s *Addr_mode_mem_to_reg_reg16_indexedContext) Reg8() IReg8Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReg8Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReg8Context)
}

func (s *Addr_mode_mem_to_reg_reg16_indexedContext) Reg16_indexed() IReg16_indexedContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReg16_indexedContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReg16_indexedContext)
}

func (s *Addr_mode_mem_to_reg_reg16_indexedContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Addr_mode_mem_to_reg_reg16_indexedContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Addr_mode_mem_to_reg_reg16_indexedContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsmE8Listener); ok {
		listenerT.EnterAddr_mode_mem_to_reg_reg16_indexed(s)
	}
}

func (s *Addr_mode_mem_to_reg_reg16_indexedContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsmE8Listener); ok {
		listenerT.ExitAddr_mode_mem_to_reg_reg16_indexed(s)
	}
}

func (p *AsmE8Parser) Addr_mode_mem_to_reg_reg16_indexed() (localctx IAddr_mode_mem_to_reg_reg16_indexedContext) {
	localctx = NewAddr_mode_mem_to_reg_reg16_indexedContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, AsmE8ParserRULE_addr_mode_mem_to_reg_reg16_indexed)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(241)
		p.Mnemonic()
	}
	{
		p.SetState(242)
		p.Match(AsmE8ParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(243)
		p.Reg8()
	}
	{
		p.SetState(244)
		p.Match(AsmE8ParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(245)
		p.Reg16_indexed()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAddr_mode_mem_to_reg_indirectContext is an interface to support dynamic dispatch.
type IAddr_mode_mem_to_reg_indirectContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Mnemonic() IMnemonicContext
	Reg8() IReg8Context
	Indirect() IIndirectContext

	// IsAddr_mode_mem_to_reg_indirectContext differentiates from other interfaces.
	IsAddr_mode_mem_to_reg_indirectContext()
}

type Addr_mode_mem_to_reg_indirectContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAddr_mode_mem_to_reg_indirectContext() *Addr_mode_mem_to_reg_indirectContext {
	var p = new(Addr_mode_mem_to_reg_indirectContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsmE8ParserRULE_addr_mode_mem_to_reg_indirect
	return p
}

func InitEmptyAddr_mode_mem_to_reg_indirectContext(p *Addr_mode_mem_to_reg_indirectContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsmE8ParserRULE_addr_mode_mem_to_reg_indirect
}

func (*Addr_mode_mem_to_reg_indirectContext) IsAddr_mode_mem_to_reg_indirectContext() {}

func NewAddr_mode_mem_to_reg_indirectContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Addr_mode_mem_to_reg_indirectContext {
	var p = new(Addr_mode_mem_to_reg_indirectContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AsmE8ParserRULE_addr_mode_mem_to_reg_indirect

	return p
}

func (s *Addr_mode_mem_to_reg_indirectContext) GetParser() antlr.Parser { return s.parser }

func (s *Addr_mode_mem_to_reg_indirectContext) Mnemonic() IMnemonicContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMnemonicContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMnemonicContext)
}

func (s *Addr_mode_mem_to_reg_indirectContext) Reg8() IReg8Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReg8Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReg8Context)
}

func (s *Addr_mode_mem_to_reg_indirectContext) Indirect() IIndirectContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIndirectContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIndirectContext)
}

func (s *Addr_mode_mem_to_reg_indirectContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Addr_mode_mem_to_reg_indirectContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Addr_mode_mem_to_reg_indirectContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsmE8Listener); ok {
		listenerT.EnterAddr_mode_mem_to_reg_indirect(s)
	}
}

func (s *Addr_mode_mem_to_reg_indirectContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsmE8Listener); ok {
		listenerT.ExitAddr_mode_mem_to_reg_indirect(s)
	}
}

func (p *AsmE8Parser) Addr_mode_mem_to_reg_indirect() (localctx IAddr_mode_mem_to_reg_indirectContext) {
	localctx = NewAddr_mode_mem_to_reg_indirectContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, AsmE8ParserRULE_addr_mode_mem_to_reg_indirect)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(247)
		p.Mnemonic()
	}
	{
		p.SetState(248)
		p.Match(AsmE8ParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(249)
		p.Reg8()
	}
	{
		p.SetState(250)
		p.Match(AsmE8ParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(251)
		p.Indirect()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAddr_mode_reg8_reg8Context is an interface to support dynamic dispatch.
type IAddr_mode_reg8_reg8Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Mnemonic() IMnemonicContext
	AllReg8() []IReg8Context
	Reg8(i int) IReg8Context

	// IsAddr_mode_reg8_reg8Context differentiates from other interfaces.
	IsAddr_mode_reg8_reg8Context()
}

type Addr_mode_reg8_reg8Context struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAddr_mode_reg8_reg8Context() *Addr_mode_reg8_reg8Context {
	var p = new(Addr_mode_reg8_reg8Context)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsmE8ParserRULE_addr_mode_reg8_reg8
	return p
}

func InitEmptyAddr_mode_reg8_reg8Context(p *Addr_mode_reg8_reg8Context) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsmE8ParserRULE_addr_mode_reg8_reg8
}

func (*Addr_mode_reg8_reg8Context) IsAddr_mode_reg8_reg8Context() {}

func NewAddr_mode_reg8_reg8Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Addr_mode_reg8_reg8Context {
	var p = new(Addr_mode_reg8_reg8Context)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AsmE8ParserRULE_addr_mode_reg8_reg8

	return p
}

func (s *Addr_mode_reg8_reg8Context) GetParser() antlr.Parser { return s.parser }

func (s *Addr_mode_reg8_reg8Context) Mnemonic() IMnemonicContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMnemonicContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMnemonicContext)
}

func (s *Addr_mode_reg8_reg8Context) AllReg8() []IReg8Context {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IReg8Context); ok {
			len++
		}
	}

	tst := make([]IReg8Context, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IReg8Context); ok {
			tst[i] = t.(IReg8Context)
			i++
		}
	}

	return tst
}

func (s *Addr_mode_reg8_reg8Context) Reg8(i int) IReg8Context {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReg8Context); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReg8Context)
}

func (s *Addr_mode_reg8_reg8Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Addr_mode_reg8_reg8Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Addr_mode_reg8_reg8Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsmE8Listener); ok {
		listenerT.EnterAddr_mode_reg8_reg8(s)
	}
}

func (s *Addr_mode_reg8_reg8Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsmE8Listener); ok {
		listenerT.ExitAddr_mode_reg8_reg8(s)
	}
}

func (p *AsmE8Parser) Addr_mode_reg8_reg8() (localctx IAddr_mode_reg8_reg8Context) {
	localctx = NewAddr_mode_reg8_reg8Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, AsmE8ParserRULE_addr_mode_reg8_reg8)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(253)
		p.Mnemonic()
	}
	{
		p.SetState(254)
		p.Match(AsmE8ParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(255)
		p.Reg8()
	}
	{
		p.SetState(256)
		p.Match(AsmE8ParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(257)
		p.Reg8()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAddr_mode_impliedContext is an interface to support dynamic dispatch.
type IAddr_mode_impliedContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Mnemonic() IMnemonicContext

	// IsAddr_mode_impliedContext differentiates from other interfaces.
	IsAddr_mode_impliedContext()
}

type Addr_mode_impliedContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAddr_mode_impliedContext() *Addr_mode_impliedContext {
	var p = new(Addr_mode_impliedContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsmE8ParserRULE_addr_mode_implied
	return p
}

func InitEmptyAddr_mode_impliedContext(p *Addr_mode_impliedContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsmE8ParserRULE_addr_mode_implied
}

func (*Addr_mode_impliedContext) IsAddr_mode_impliedContext() {}

func NewAddr_mode_impliedContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Addr_mode_impliedContext {
	var p = new(Addr_mode_impliedContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AsmE8ParserRULE_addr_mode_implied

	return p
}

func (s *Addr_mode_impliedContext) GetParser() antlr.Parser { return s.parser }

func (s *Addr_mode_impliedContext) Mnemonic() IMnemonicContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMnemonicContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMnemonicContext)
}

func (s *Addr_mode_impliedContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Addr_mode_impliedContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Addr_mode_impliedContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsmE8Listener); ok {
		listenerT.EnterAddr_mode_implied(s)
	}
}

func (s *Addr_mode_impliedContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsmE8Listener); ok {
		listenerT.ExitAddr_mode_implied(s)
	}
}

func (p *AsmE8Parser) Addr_mode_implied() (localctx IAddr_mode_impliedContext) {
	localctx = NewAddr_mode_impliedContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, AsmE8ParserRULE_addr_mode_implied)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(259)
		p.Mnemonic()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAddr_mode_implied_reg16Context is an interface to support dynamic dispatch.
type IAddr_mode_implied_reg16Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Mnemonic() IMnemonicContext
	Reg16() IReg16Context

	// IsAddr_mode_implied_reg16Context differentiates from other interfaces.
	IsAddr_mode_implied_reg16Context()
}

type Addr_mode_implied_reg16Context struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAddr_mode_implied_reg16Context() *Addr_mode_implied_reg16Context {
	var p = new(Addr_mode_implied_reg16Context)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsmE8ParserRULE_addr_mode_implied_reg16
	return p
}

func InitEmptyAddr_mode_implied_reg16Context(p *Addr_mode_implied_reg16Context) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsmE8ParserRULE_addr_mode_implied_reg16
}

func (*Addr_mode_implied_reg16Context) IsAddr_mode_implied_reg16Context() {}

func NewAddr_mode_implied_reg16Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Addr_mode_implied_reg16Context {
	var p = new(Addr_mode_implied_reg16Context)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AsmE8ParserRULE_addr_mode_implied_reg16

	return p
}

func (s *Addr_mode_implied_reg16Context) GetParser() antlr.Parser { return s.parser }

func (s *Addr_mode_implied_reg16Context) Mnemonic() IMnemonicContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMnemonicContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMnemonicContext)
}

func (s *Addr_mode_implied_reg16Context) Reg16() IReg16Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReg16Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReg16Context)
}

func (s *Addr_mode_implied_reg16Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Addr_mode_implied_reg16Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Addr_mode_implied_reg16Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsmE8Listener); ok {
		listenerT.EnterAddr_mode_implied_reg16(s)
	}
}

func (s *Addr_mode_implied_reg16Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsmE8Listener); ok {
		listenerT.ExitAddr_mode_implied_reg16(s)
	}
}

func (p *AsmE8Parser) Addr_mode_implied_reg16() (localctx IAddr_mode_implied_reg16Context) {
	localctx = NewAddr_mode_implied_reg16Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, AsmE8ParserRULE_addr_mode_implied_reg16)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(261)
		p.Mnemonic()
	}
	{
		p.SetState(262)
		p.Match(AsmE8ParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(263)
		p.Reg16()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAddr_mode_reg16_immContext is an interface to support dynamic dispatch.
type IAddr_mode_reg16_immContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Mnemonic() IMnemonicContext
	Reg16() IReg16Context
	Imm() IImmContext

	// IsAddr_mode_reg16_immContext differentiates from other interfaces.
	IsAddr_mode_reg16_immContext()
}

type Addr_mode_reg16_immContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAddr_mode_reg16_immContext() *Addr_mode_reg16_immContext {
	var p = new(Addr_mode_reg16_immContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsmE8ParserRULE_addr_mode_reg16_imm
	return p
}

func InitEmptyAddr_mode_reg16_immContext(p *Addr_mode_reg16_immContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsmE8ParserRULE_addr_mode_reg16_imm
}

func (*Addr_mode_reg16_immContext) IsAddr_mode_reg16_immContext() {}

func NewAddr_mode_reg16_immContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Addr_mode_reg16_immContext {
	var p = new(Addr_mode_reg16_immContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AsmE8ParserRULE_addr_mode_reg16_imm

	return p
}

func (s *Addr_mode_reg16_immContext) GetParser() antlr.Parser { return s.parser }

func (s *Addr_mode_reg16_immContext) Mnemonic() IMnemonicContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMnemonicContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMnemonicContext)
}

func (s *Addr_mode_reg16_immContext) Reg16() IReg16Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReg16Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReg16Context)
}

func (s *Addr_mode_reg16_immContext) Imm() IImmContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IImmContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IImmContext)
}

func (s *Addr_mode_reg16_immContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Addr_mode_reg16_immContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Addr_mode_reg16_immContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsmE8Listener); ok {
		listenerT.EnterAddr_mode_reg16_imm(s)
	}
}

func (s *Addr_mode_reg16_immContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsmE8Listener); ok {
		listenerT.ExitAddr_mode_reg16_imm(s)
	}
}

func (p *AsmE8Parser) Addr_mode_reg16_imm() (localctx IAddr_mode_reg16_immContext) {
	localctx = NewAddr_mode_reg16_immContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, AsmE8ParserRULE_addr_mode_reg16_imm)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(265)
		p.Mnemonic()
	}
	{
		p.SetState(266)
		p.Match(AsmE8ParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(267)
		p.Reg16()
	}
	{
		p.SetState(268)
		p.Match(AsmE8ParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(269)
		p.Imm()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAddr_mode_reg16_imm_tagContext is an interface to support dynamic dispatch.
type IAddr_mode_reg16_imm_tagContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Mnemonic() IMnemonicContext
	Reg16() IReg16Context
	Tag() ITagContext

	// IsAddr_mode_reg16_imm_tagContext differentiates from other interfaces.
	IsAddr_mode_reg16_imm_tagContext()
}

type Addr_mode_reg16_imm_tagContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAddr_mode_reg16_imm_tagContext() *Addr_mode_reg16_imm_tagContext {
	var p = new(Addr_mode_reg16_imm_tagContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsmE8ParserRULE_addr_mode_reg16_imm_tag
	return p
}

func InitEmptyAddr_mode_reg16_imm_tagContext(p *Addr_mode_reg16_imm_tagContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsmE8ParserRULE_addr_mode_reg16_imm_tag
}

func (*Addr_mode_reg16_imm_tagContext) IsAddr_mode_reg16_imm_tagContext() {}

func NewAddr_mode_reg16_imm_tagContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Addr_mode_reg16_imm_tagContext {
	var p = new(Addr_mode_reg16_imm_tagContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AsmE8ParserRULE_addr_mode_reg16_imm_tag

	return p
}

func (s *Addr_mode_reg16_imm_tagContext) GetParser() antlr.Parser { return s.parser }

func (s *Addr_mode_reg16_imm_tagContext) Mnemonic() IMnemonicContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMnemonicContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMnemonicContext)
}

func (s *Addr_mode_reg16_imm_tagContext) Reg16() IReg16Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReg16Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReg16Context)
}

func (s *Addr_mode_reg16_imm_tagContext) Tag() ITagContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITagContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITagContext)
}

func (s *Addr_mode_reg16_imm_tagContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Addr_mode_reg16_imm_tagContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Addr_mode_reg16_imm_tagContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsmE8Listener); ok {
		listenerT.EnterAddr_mode_reg16_imm_tag(s)
	}
}

func (s *Addr_mode_reg16_imm_tagContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsmE8Listener); ok {
		listenerT.ExitAddr_mode_reg16_imm_tag(s)
	}
}

func (p *AsmE8Parser) Addr_mode_reg16_imm_tag() (localctx IAddr_mode_reg16_imm_tagContext) {
	localctx = NewAddr_mode_reg16_imm_tagContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, AsmE8ParserRULE_addr_mode_reg16_imm_tag)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(271)
		p.Mnemonic()
	}
	{
		p.SetState(272)
		p.Match(AsmE8ParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(273)
		p.Reg16()
	}
	{
		p.SetState(274)
		p.Match(AsmE8ParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(275)
		p.Tag()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAddr_mode_reg16_reg16Context is an interface to support dynamic dispatch.
type IAddr_mode_reg16_reg16Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Mnemonic() IMnemonicContext
	AllReg16() []IReg16Context
	Reg16(i int) IReg16Context

	// IsAddr_mode_reg16_reg16Context differentiates from other interfaces.
	IsAddr_mode_reg16_reg16Context()
}

type Addr_mode_reg16_reg16Context struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAddr_mode_reg16_reg16Context() *Addr_mode_reg16_reg16Context {
	var p = new(Addr_mode_reg16_reg16Context)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsmE8ParserRULE_addr_mode_reg16_reg16
	return p
}

func InitEmptyAddr_mode_reg16_reg16Context(p *Addr_mode_reg16_reg16Context) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsmE8ParserRULE_addr_mode_reg16_reg16
}

func (*Addr_mode_reg16_reg16Context) IsAddr_mode_reg16_reg16Context() {}

func NewAddr_mode_reg16_reg16Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Addr_mode_reg16_reg16Context {
	var p = new(Addr_mode_reg16_reg16Context)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AsmE8ParserRULE_addr_mode_reg16_reg16

	return p
}

func (s *Addr_mode_reg16_reg16Context) GetParser() antlr.Parser { return s.parser }

func (s *Addr_mode_reg16_reg16Context) Mnemonic() IMnemonicContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMnemonicContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMnemonicContext)
}

func (s *Addr_mode_reg16_reg16Context) AllReg16() []IReg16Context {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IReg16Context); ok {
			len++
		}
	}

	tst := make([]IReg16Context, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IReg16Context); ok {
			tst[i] = t.(IReg16Context)
			i++
		}
	}

	return tst
}

func (s *Addr_mode_reg16_reg16Context) Reg16(i int) IReg16Context {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReg16Context); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReg16Context)
}

func (s *Addr_mode_reg16_reg16Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Addr_mode_reg16_reg16Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Addr_mode_reg16_reg16Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsmE8Listener); ok {
		listenerT.EnterAddr_mode_reg16_reg16(s)
	}
}

func (s *Addr_mode_reg16_reg16Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsmE8Listener); ok {
		listenerT.ExitAddr_mode_reg16_reg16(s)
	}
}

func (p *AsmE8Parser) Addr_mode_reg16_reg16() (localctx IAddr_mode_reg16_reg16Context) {
	localctx = NewAddr_mode_reg16_reg16Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, AsmE8ParserRULE_addr_mode_reg16_reg16)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(277)
		p.Mnemonic()
	}
	{
		p.SetState(278)
		p.Match(AsmE8ParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(279)
		p.Reg16()
	}
	{
		p.SetState(280)
		p.Match(AsmE8ParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(281)
		p.Reg16()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAddr_mode_reg16_stackContext is an interface to support dynamic dispatch.
type IAddr_mode_reg16_stackContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Mnemonic() IMnemonicContext
	Reg16() IReg16Context
	Stack() IStackContext

	// IsAddr_mode_reg16_stackContext differentiates from other interfaces.
	IsAddr_mode_reg16_stackContext()
}

type Addr_mode_reg16_stackContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAddr_mode_reg16_stackContext() *Addr_mode_reg16_stackContext {
	var p = new(Addr_mode_reg16_stackContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsmE8ParserRULE_addr_mode_reg16_stack
	return p
}

func InitEmptyAddr_mode_reg16_stackContext(p *Addr_mode_reg16_stackContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsmE8ParserRULE_addr_mode_reg16_stack
}

func (*Addr_mode_reg16_stackContext) IsAddr_mode_reg16_stackContext() {}

func NewAddr_mode_reg16_stackContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Addr_mode_reg16_stackContext {
	var p = new(Addr_mode_reg16_stackContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AsmE8ParserRULE_addr_mode_reg16_stack

	return p
}

func (s *Addr_mode_reg16_stackContext) GetParser() antlr.Parser { return s.parser }

func (s *Addr_mode_reg16_stackContext) Mnemonic() IMnemonicContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMnemonicContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMnemonicContext)
}

func (s *Addr_mode_reg16_stackContext) Reg16() IReg16Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReg16Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReg16Context)
}

func (s *Addr_mode_reg16_stackContext) Stack() IStackContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStackContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStackContext)
}

func (s *Addr_mode_reg16_stackContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Addr_mode_reg16_stackContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Addr_mode_reg16_stackContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsmE8Listener); ok {
		listenerT.EnterAddr_mode_reg16_stack(s)
	}
}

func (s *Addr_mode_reg16_stackContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsmE8Listener); ok {
		listenerT.ExitAddr_mode_reg16_stack(s)
	}
}

func (p *AsmE8Parser) Addr_mode_reg16_stack() (localctx IAddr_mode_reg16_stackContext) {
	localctx = NewAddr_mode_reg16_stackContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, AsmE8ParserRULE_addr_mode_reg16_stack)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(283)
		p.Mnemonic()
	}
	{
		p.SetState(284)
		p.Match(AsmE8ParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(285)
		p.Reg16()
	}
	{
		p.SetState(286)
		p.Match(AsmE8ParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(287)
		p.Stack()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAddr_mode_implied_stackContext is an interface to support dynamic dispatch.
type IAddr_mode_implied_stackContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Mnemonic() IMnemonicContext
	Stack() IStackContext

	// IsAddr_mode_implied_stackContext differentiates from other interfaces.
	IsAddr_mode_implied_stackContext()
}

type Addr_mode_implied_stackContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAddr_mode_implied_stackContext() *Addr_mode_implied_stackContext {
	var p = new(Addr_mode_implied_stackContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsmE8ParserRULE_addr_mode_implied_stack
	return p
}

func InitEmptyAddr_mode_implied_stackContext(p *Addr_mode_implied_stackContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsmE8ParserRULE_addr_mode_implied_stack
}

func (*Addr_mode_implied_stackContext) IsAddr_mode_implied_stackContext() {}

func NewAddr_mode_implied_stackContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Addr_mode_implied_stackContext {
	var p = new(Addr_mode_implied_stackContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AsmE8ParserRULE_addr_mode_implied_stack

	return p
}

func (s *Addr_mode_implied_stackContext) GetParser() antlr.Parser { return s.parser }

func (s *Addr_mode_implied_stackContext) Mnemonic() IMnemonicContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMnemonicContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMnemonicContext)
}

func (s *Addr_mode_implied_stackContext) Stack() IStackContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStackContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStackContext)
}

func (s *Addr_mode_implied_stackContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Addr_mode_implied_stackContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Addr_mode_implied_stackContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsmE8Listener); ok {
		listenerT.EnterAddr_mode_implied_stack(s)
	}
}

func (s *Addr_mode_implied_stackContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsmE8Listener); ok {
		listenerT.ExitAddr_mode_implied_stack(s)
	}
}

func (p *AsmE8Parser) Addr_mode_implied_stack() (localctx IAddr_mode_implied_stackContext) {
	localctx = NewAddr_mode_implied_stackContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, AsmE8ParserRULE_addr_mode_implied_stack)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(289)
		p.Mnemonic()
	}
	{
		p.SetState(290)
		p.Match(AsmE8ParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(291)
		p.Stack()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAddr_mode_stack_imm8Context is an interface to support dynamic dispatch.
type IAddr_mode_stack_imm8Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Mnemonic() IMnemonicContext
	Stack() IStackContext
	Imm() IImmContext

	// IsAddr_mode_stack_imm8Context differentiates from other interfaces.
	IsAddr_mode_stack_imm8Context()
}

type Addr_mode_stack_imm8Context struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAddr_mode_stack_imm8Context() *Addr_mode_stack_imm8Context {
	var p = new(Addr_mode_stack_imm8Context)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsmE8ParserRULE_addr_mode_stack_imm8
	return p
}

func InitEmptyAddr_mode_stack_imm8Context(p *Addr_mode_stack_imm8Context) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsmE8ParserRULE_addr_mode_stack_imm8
}

func (*Addr_mode_stack_imm8Context) IsAddr_mode_stack_imm8Context() {}

func NewAddr_mode_stack_imm8Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Addr_mode_stack_imm8Context {
	var p = new(Addr_mode_stack_imm8Context)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AsmE8ParserRULE_addr_mode_stack_imm8

	return p
}

func (s *Addr_mode_stack_imm8Context) GetParser() antlr.Parser { return s.parser }

func (s *Addr_mode_stack_imm8Context) Mnemonic() IMnemonicContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMnemonicContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMnemonicContext)
}

func (s *Addr_mode_stack_imm8Context) Stack() IStackContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStackContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStackContext)
}

func (s *Addr_mode_stack_imm8Context) Imm() IImmContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IImmContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IImmContext)
}

func (s *Addr_mode_stack_imm8Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Addr_mode_stack_imm8Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Addr_mode_stack_imm8Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsmE8Listener); ok {
		listenerT.EnterAddr_mode_stack_imm8(s)
	}
}

func (s *Addr_mode_stack_imm8Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsmE8Listener); ok {
		listenerT.ExitAddr_mode_stack_imm8(s)
	}
}

func (p *AsmE8Parser) Addr_mode_stack_imm8() (localctx IAddr_mode_stack_imm8Context) {
	localctx = NewAddr_mode_stack_imm8Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, AsmE8ParserRULE_addr_mode_stack_imm8)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(293)
		p.Mnemonic()
	}
	{
		p.SetState(294)
		p.Match(AsmE8ParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(295)
		p.Stack()
	}
	{
		p.SetState(296)
		p.Match(AsmE8ParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(297)
		p.Imm()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAddr_mode_stack_imm8_tagContext is an interface to support dynamic dispatch.
type IAddr_mode_stack_imm8_tagContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Mnemonic() IMnemonicContext
	Stack() IStackContext
	Tag() ITagContext

	// IsAddr_mode_stack_imm8_tagContext differentiates from other interfaces.
	IsAddr_mode_stack_imm8_tagContext()
}

type Addr_mode_stack_imm8_tagContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAddr_mode_stack_imm8_tagContext() *Addr_mode_stack_imm8_tagContext {
	var p = new(Addr_mode_stack_imm8_tagContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsmE8ParserRULE_addr_mode_stack_imm8_tag
	return p
}

func InitEmptyAddr_mode_stack_imm8_tagContext(p *Addr_mode_stack_imm8_tagContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsmE8ParserRULE_addr_mode_stack_imm8_tag
}

func (*Addr_mode_stack_imm8_tagContext) IsAddr_mode_stack_imm8_tagContext() {}

func NewAddr_mode_stack_imm8_tagContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Addr_mode_stack_imm8_tagContext {
	var p = new(Addr_mode_stack_imm8_tagContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AsmE8ParserRULE_addr_mode_stack_imm8_tag

	return p
}

func (s *Addr_mode_stack_imm8_tagContext) GetParser() antlr.Parser { return s.parser }

func (s *Addr_mode_stack_imm8_tagContext) Mnemonic() IMnemonicContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMnemonicContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMnemonicContext)
}

func (s *Addr_mode_stack_imm8_tagContext) Stack() IStackContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStackContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStackContext)
}

func (s *Addr_mode_stack_imm8_tagContext) Tag() ITagContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITagContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITagContext)
}

func (s *Addr_mode_stack_imm8_tagContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Addr_mode_stack_imm8_tagContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Addr_mode_stack_imm8_tagContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsmE8Listener); ok {
		listenerT.EnterAddr_mode_stack_imm8_tag(s)
	}
}

func (s *Addr_mode_stack_imm8_tagContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsmE8Listener); ok {
		listenerT.ExitAddr_mode_stack_imm8_tag(s)
	}
}

func (p *AsmE8Parser) Addr_mode_stack_imm8_tag() (localctx IAddr_mode_stack_imm8_tagContext) {
	localctx = NewAddr_mode_stack_imm8_tagContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, AsmE8ParserRULE_addr_mode_stack_imm8_tag)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(299)
		p.Mnemonic()
	}
	{
		p.SetState(300)
		p.Match(AsmE8ParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(301)
		p.Stack()
	}
	{
		p.SetState(302)
		p.Match(AsmE8ParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(303)
		p.Tag()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAddr_mode_stack_reg16Context is an interface to support dynamic dispatch.
type IAddr_mode_stack_reg16Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Mnemonic() IMnemonicContext
	Stack() IStackContext
	Reg16() IReg16Context

	// IsAddr_mode_stack_reg16Context differentiates from other interfaces.
	IsAddr_mode_stack_reg16Context()
}

type Addr_mode_stack_reg16Context struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAddr_mode_stack_reg16Context() *Addr_mode_stack_reg16Context {
	var p = new(Addr_mode_stack_reg16Context)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsmE8ParserRULE_addr_mode_stack_reg16
	return p
}

func InitEmptyAddr_mode_stack_reg16Context(p *Addr_mode_stack_reg16Context) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsmE8ParserRULE_addr_mode_stack_reg16
}

func (*Addr_mode_stack_reg16Context) IsAddr_mode_stack_reg16Context() {}

func NewAddr_mode_stack_reg16Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Addr_mode_stack_reg16Context {
	var p = new(Addr_mode_stack_reg16Context)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AsmE8ParserRULE_addr_mode_stack_reg16

	return p
}

func (s *Addr_mode_stack_reg16Context) GetParser() antlr.Parser { return s.parser }

func (s *Addr_mode_stack_reg16Context) Mnemonic() IMnemonicContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMnemonicContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMnemonicContext)
}

func (s *Addr_mode_stack_reg16Context) Stack() IStackContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStackContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStackContext)
}

func (s *Addr_mode_stack_reg16Context) Reg16() IReg16Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReg16Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReg16Context)
}

func (s *Addr_mode_stack_reg16Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Addr_mode_stack_reg16Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Addr_mode_stack_reg16Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsmE8Listener); ok {
		listenerT.EnterAddr_mode_stack_reg16(s)
	}
}

func (s *Addr_mode_stack_reg16Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsmE8Listener); ok {
		listenerT.ExitAddr_mode_stack_reg16(s)
	}
}

func (p *AsmE8Parser) Addr_mode_stack_reg16() (localctx IAddr_mode_stack_reg16Context) {
	localctx = NewAddr_mode_stack_reg16Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, AsmE8ParserRULE_addr_mode_stack_reg16)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(305)
		p.Mnemonic()
	}
	{
		p.SetState(306)
		p.Match(AsmE8ParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(307)
		p.Stack()
	}
	{
		p.SetState(308)
		p.Match(AsmE8ParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(309)
		p.Reg16()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMnemonicContext is an interface to support dynamic dispatch.
type IMnemonicContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsMnemonicContext differentiates from other interfaces.
	IsMnemonicContext()
}

type MnemonicContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMnemonicContext() *MnemonicContext {
	var p = new(MnemonicContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsmE8ParserRULE_mnemonic
	return p
}

func InitEmptyMnemonicContext(p *MnemonicContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsmE8ParserRULE_mnemonic
}

func (*MnemonicContext) IsMnemonicContext() {}

func NewMnemonicContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MnemonicContext {
	var p = new(MnemonicContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AsmE8ParserRULE_mnemonic

	return p
}

func (s *MnemonicContext) GetParser() antlr.Parser { return s.parser }
func (s *MnemonicContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MnemonicContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MnemonicContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsmE8Listener); ok {
		listenerT.EnterMnemonic(s)
	}
}

func (s *MnemonicContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsmE8Listener); ok {
		listenerT.ExitMnemonic(s)
	}
}

func (p *AsmE8Parser) Mnemonic() (localctx IMnemonicContext) {
	localctx = NewMnemonicContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, AsmE8ParserRULE_mnemonic)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(311)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&549755813632) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IReg8Context is an interface to support dynamic dispatch.
type IReg8Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsReg8Context differentiates from other interfaces.
	IsReg8Context()
}

type Reg8Context struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReg8Context() *Reg8Context {
	var p = new(Reg8Context)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsmE8ParserRULE_reg8
	return p
}

func InitEmptyReg8Context(p *Reg8Context) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsmE8ParserRULE_reg8
}

func (*Reg8Context) IsReg8Context() {}

func NewReg8Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Reg8Context {
	var p = new(Reg8Context)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AsmE8ParserRULE_reg8

	return p
}

func (s *Reg8Context) GetParser() antlr.Parser { return s.parser }
func (s *Reg8Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Reg8Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Reg8Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsmE8Listener); ok {
		listenerT.EnterReg8(s)
	}
}

func (s *Reg8Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsmE8Listener); ok {
		listenerT.ExitReg8(s)
	}
}

func (p *AsmE8Parser) Reg8() (localctx IReg8Context) {
	localctx = NewReg8Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, AsmE8ParserRULE_reg8)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(313)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&280925220896768) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IReg16Context is an interface to support dynamic dispatch.
type IReg16Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsReg16Context differentiates from other interfaces.
	IsReg16Context()
}

type Reg16Context struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReg16Context() *Reg16Context {
	var p = new(Reg16Context)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsmE8ParserRULE_reg16
	return p
}

func InitEmptyReg16Context(p *Reg16Context) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsmE8ParserRULE_reg16
}

func (*Reg16Context) IsReg16Context() {}

func NewReg16Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Reg16Context {
	var p = new(Reg16Context)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AsmE8ParserRULE_reg16

	return p
}

func (s *Reg16Context) GetParser() antlr.Parser { return s.parser }
func (s *Reg16Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Reg16Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Reg16Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsmE8Listener); ok {
		listenerT.EnterReg16(s)
	}
}

func (s *Reg16Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsmE8Listener); ok {
		listenerT.ExitReg16(s)
	}
}

func (p *AsmE8Parser) Reg16() (localctx IReg16Context) {
	localctx = NewReg16Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, AsmE8ParserRULE_reg16)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(315)
		_la = p.GetTokenStream().LA(1)

		if !(_la == AsmE8ParserT__47 || _la == AsmE8ParserT__48) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStackContext is an interface to support dynamic dispatch.
type IStackContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsStackContext differentiates from other interfaces.
	IsStackContext()
}

type StackContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStackContext() *StackContext {
	var p = new(StackContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsmE8ParserRULE_stack
	return p
}

func InitEmptyStackContext(p *StackContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsmE8ParserRULE_stack
}

func (*StackContext) IsStackContext() {}

func NewStackContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StackContext {
	var p = new(StackContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AsmE8ParserRULE_stack

	return p
}

func (s *StackContext) GetParser() antlr.Parser { return s.parser }
func (s *StackContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StackContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StackContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsmE8Listener); ok {
		listenerT.EnterStack(s)
	}
}

func (s *StackContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsmE8Listener); ok {
		listenerT.ExitStack(s)
	}
}

func (p *AsmE8Parser) Stack() (localctx IStackContext) {
	localctx = NewStackContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, AsmE8ParserRULE_stack)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(317)
		p.Match(AsmE8ParserT__49)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDirectContext is an interface to support dynamic dispatch.
type IDirectContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Imm() IImmContext
	Tag() ITagContext
	Direct_virtual_offset() IDirect_virtual_offsetContext

	// IsDirectContext differentiates from other interfaces.
	IsDirectContext()
}

type DirectContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDirectContext() *DirectContext {
	var p = new(DirectContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsmE8ParserRULE_direct
	return p
}

func InitEmptyDirectContext(p *DirectContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsmE8ParserRULE_direct
}

func (*DirectContext) IsDirectContext() {}

func NewDirectContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DirectContext {
	var p = new(DirectContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AsmE8ParserRULE_direct

	return p
}

func (s *DirectContext) GetParser() antlr.Parser { return s.parser }

func (s *DirectContext) Imm() IImmContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IImmContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IImmContext)
}

func (s *DirectContext) Tag() ITagContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITagContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITagContext)
}

func (s *DirectContext) Direct_virtual_offset() IDirect_virtual_offsetContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDirect_virtual_offsetContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDirect_virtual_offsetContext)
}

func (s *DirectContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DirectContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DirectContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsmE8Listener); ok {
		listenerT.EnterDirect(s)
	}
}

func (s *DirectContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsmE8Listener); ok {
		listenerT.ExitDirect(s)
	}
}

func (p *AsmE8Parser) Direct() (localctx IDirectContext) {
	localctx = NewDirectContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, AsmE8ParserRULE_direct)
	p.SetState(328)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 7, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(319)
			p.Match(AsmE8ParserT__50)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(320)
			p.Imm()
		}
		{
			p.SetState(321)
			p.Match(AsmE8ParserT__51)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(323)
			p.Match(AsmE8ParserT__50)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(324)
			p.Tag()
		}
		{
			p.SetState(325)
			p.Match(AsmE8ParserT__51)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(327)
			p.Direct_virtual_offset()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDirect_virtual_offsetContext is an interface to support dynamic dispatch.
type IDirect_virtual_offsetContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	STR() antlr.TerminalNode
	INT() antlr.TerminalNode

	// IsDirect_virtual_offsetContext differentiates from other interfaces.
	IsDirect_virtual_offsetContext()
}

type Direct_virtual_offsetContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDirect_virtual_offsetContext() *Direct_virtual_offsetContext {
	var p = new(Direct_virtual_offsetContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsmE8ParserRULE_direct_virtual_offset
	return p
}

func InitEmptyDirect_virtual_offsetContext(p *Direct_virtual_offsetContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsmE8ParserRULE_direct_virtual_offset
}

func (*Direct_virtual_offsetContext) IsDirect_virtual_offsetContext() {}

func NewDirect_virtual_offsetContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Direct_virtual_offsetContext {
	var p = new(Direct_virtual_offsetContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AsmE8ParserRULE_direct_virtual_offset

	return p
}

func (s *Direct_virtual_offsetContext) GetParser() antlr.Parser { return s.parser }

func (s *Direct_virtual_offsetContext) STR() antlr.TerminalNode {
	return s.GetToken(AsmE8ParserSTR, 0)
}

func (s *Direct_virtual_offsetContext) INT() antlr.TerminalNode {
	return s.GetToken(AsmE8ParserINT, 0)
}

func (s *Direct_virtual_offsetContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Direct_virtual_offsetContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Direct_virtual_offsetContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsmE8Listener); ok {
		listenerT.EnterDirect_virtual_offset(s)
	}
}

func (s *Direct_virtual_offsetContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsmE8Listener); ok {
		listenerT.ExitDirect_virtual_offset(s)
	}
}

func (p *AsmE8Parser) Direct_virtual_offset() (localctx IDirect_virtual_offsetContext) {
	localctx = NewDirect_virtual_offsetContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, AsmE8ParserRULE_direct_virtual_offset)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(330)
		p.Match(AsmE8ParserT__50)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(331)
		p.Match(AsmE8ParserSTR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(332)
		_la = p.GetTokenStream().LA(1)

		if !(_la == AsmE8ParserT__52 || _la == AsmE8ParserT__53) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(333)
		p.Match(AsmE8ParserINT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(334)
		p.Match(AsmE8ParserT__51)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDirect_offsetContext is an interface to support dynamic dispatch.
type IDirect_offsetContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Imm() IImmContext
	Reg8() IReg8Context
	Tag() ITagContext

	// IsDirect_offsetContext differentiates from other interfaces.
	IsDirect_offsetContext()
}

type Direct_offsetContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDirect_offsetContext() *Direct_offsetContext {
	var p = new(Direct_offsetContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsmE8ParserRULE_direct_offset
	return p
}

func InitEmptyDirect_offsetContext(p *Direct_offsetContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsmE8ParserRULE_direct_offset
}

func (*Direct_offsetContext) IsDirect_offsetContext() {}

func NewDirect_offsetContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Direct_offsetContext {
	var p = new(Direct_offsetContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AsmE8ParserRULE_direct_offset

	return p
}

func (s *Direct_offsetContext) GetParser() antlr.Parser { return s.parser }

func (s *Direct_offsetContext) Imm() IImmContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IImmContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IImmContext)
}

func (s *Direct_offsetContext) Reg8() IReg8Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReg8Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReg8Context)
}

func (s *Direct_offsetContext) Tag() ITagContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITagContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITagContext)
}

func (s *Direct_offsetContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Direct_offsetContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Direct_offsetContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsmE8Listener); ok {
		listenerT.EnterDirect_offset(s)
	}
}

func (s *Direct_offsetContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsmE8Listener); ok {
		listenerT.ExitDirect_offset(s)
	}
}

func (p *AsmE8Parser) Direct_offset() (localctx IDirect_offsetContext) {
	localctx = NewDirect_offsetContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, AsmE8ParserRULE_direct_offset)
	p.SetState(348)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(336)
			p.Match(AsmE8ParserT__50)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(337)
			p.Imm()
		}
		{
			p.SetState(338)
			p.Match(AsmE8ParserT__52)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(339)
			p.Reg8()
		}
		{
			p.SetState(340)
			p.Match(AsmE8ParserT__51)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(342)
			p.Match(AsmE8ParserT__50)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(343)
			p.Tag()
		}
		{
			p.SetState(344)
			p.Match(AsmE8ParserT__52)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(345)
			p.Reg8()
		}
		{
			p.SetState(346)
			p.Match(AsmE8ParserT__51)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIndirect_offsetContext is an interface to support dynamic dispatch.
type IIndirect_offsetContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Reg16() IReg16Context
	Imm() IImmContext

	// IsIndirect_offsetContext differentiates from other interfaces.
	IsIndirect_offsetContext()
}

type Indirect_offsetContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIndirect_offsetContext() *Indirect_offsetContext {
	var p = new(Indirect_offsetContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsmE8ParserRULE_indirect_offset
	return p
}

func InitEmptyIndirect_offsetContext(p *Indirect_offsetContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsmE8ParserRULE_indirect_offset
}

func (*Indirect_offsetContext) IsIndirect_offsetContext() {}

func NewIndirect_offsetContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Indirect_offsetContext {
	var p = new(Indirect_offsetContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AsmE8ParserRULE_indirect_offset

	return p
}

func (s *Indirect_offsetContext) GetParser() antlr.Parser { return s.parser }

func (s *Indirect_offsetContext) Reg16() IReg16Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReg16Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReg16Context)
}

func (s *Indirect_offsetContext) Imm() IImmContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IImmContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IImmContext)
}

func (s *Indirect_offsetContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Indirect_offsetContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Indirect_offsetContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsmE8Listener); ok {
		listenerT.EnterIndirect_offset(s)
	}
}

func (s *Indirect_offsetContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsmE8Listener); ok {
		listenerT.ExitIndirect_offset(s)
	}
}

func (p *AsmE8Parser) Indirect_offset() (localctx IIndirect_offsetContext) {
	localctx = NewIndirect_offsetContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, AsmE8ParserRULE_indirect_offset)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(350)
		p.Match(AsmE8ParserT__50)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(351)
		p.Reg16()
	}
	{
		p.SetState(352)
		p.Match(AsmE8ParserT__52)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(353)
		p.Imm()
	}
	{
		p.SetState(354)
		p.Match(AsmE8ParserT__51)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IReg16_indexedContext is an interface to support dynamic dispatch.
type IReg16_indexedContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Reg16() IReg16Context
	Reg8() IReg8Context

	// IsReg16_indexedContext differentiates from other interfaces.
	IsReg16_indexedContext()
}

type Reg16_indexedContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReg16_indexedContext() *Reg16_indexedContext {
	var p = new(Reg16_indexedContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsmE8ParserRULE_reg16_indexed
	return p
}

func InitEmptyReg16_indexedContext(p *Reg16_indexedContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsmE8ParserRULE_reg16_indexed
}

func (*Reg16_indexedContext) IsReg16_indexedContext() {}

func NewReg16_indexedContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Reg16_indexedContext {
	var p = new(Reg16_indexedContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AsmE8ParserRULE_reg16_indexed

	return p
}

func (s *Reg16_indexedContext) GetParser() antlr.Parser { return s.parser }

func (s *Reg16_indexedContext) Reg16() IReg16Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReg16Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReg16Context)
}

func (s *Reg16_indexedContext) Reg8() IReg8Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReg8Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReg8Context)
}

func (s *Reg16_indexedContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Reg16_indexedContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Reg16_indexedContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsmE8Listener); ok {
		listenerT.EnterReg16_indexed(s)
	}
}

func (s *Reg16_indexedContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsmE8Listener); ok {
		listenerT.ExitReg16_indexed(s)
	}
}

func (p *AsmE8Parser) Reg16_indexed() (localctx IReg16_indexedContext) {
	localctx = NewReg16_indexedContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, AsmE8ParserRULE_reg16_indexed)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(356)
		p.Match(AsmE8ParserT__50)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(357)
		p.Reg16()
	}
	{
		p.SetState(358)
		p.Match(AsmE8ParserT__52)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(359)
		p.Reg8()
	}
	{
		p.SetState(360)
		p.Match(AsmE8ParserT__51)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIndirectContext is an interface to support dynamic dispatch.
type IIndirectContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Reg16() IReg16Context

	// IsIndirectContext differentiates from other interfaces.
	IsIndirectContext()
}

type IndirectContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIndirectContext() *IndirectContext {
	var p = new(IndirectContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsmE8ParserRULE_indirect
	return p
}

func InitEmptyIndirectContext(p *IndirectContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsmE8ParserRULE_indirect
}

func (*IndirectContext) IsIndirectContext() {}

func NewIndirectContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IndirectContext {
	var p = new(IndirectContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AsmE8ParserRULE_indirect

	return p
}

func (s *IndirectContext) GetParser() antlr.Parser { return s.parser }

func (s *IndirectContext) Reg16() IReg16Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReg16Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReg16Context)
}

func (s *IndirectContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IndirectContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IndirectContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsmE8Listener); ok {
		listenerT.EnterIndirect(s)
	}
}

func (s *IndirectContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsmE8Listener); ok {
		listenerT.ExitIndirect(s)
	}
}

func (p *AsmE8Parser) Indirect() (localctx IIndirectContext) {
	localctx = NewIndirectContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 80, AsmE8ParserRULE_indirect)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(362)
		p.Match(AsmE8ParserT__50)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(363)
		p.Reg16()
	}
	{
		p.SetState(364)
		p.Match(AsmE8ParserT__51)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IVariableContext is an interface to support dynamic dispatch.
type IVariableContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Tag() ITagContext
	Imm() IImmContext

	// IsVariableContext differentiates from other interfaces.
	IsVariableContext()
}

type VariableContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariableContext() *VariableContext {
	var p = new(VariableContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsmE8ParserRULE_variable
	return p
}

func InitEmptyVariableContext(p *VariableContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsmE8ParserRULE_variable
}

func (*VariableContext) IsVariableContext() {}

func NewVariableContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableContext {
	var p = new(VariableContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AsmE8ParserRULE_variable

	return p
}

func (s *VariableContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableContext) Tag() ITagContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITagContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITagContext)
}

func (s *VariableContext) Imm() IImmContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IImmContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IImmContext)
}

func (s *VariableContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsmE8Listener); ok {
		listenerT.EnterVariable(s)
	}
}

func (s *VariableContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsmE8Listener); ok {
		listenerT.ExitVariable(s)
	}
}

func (p *AsmE8Parser) Variable() (localctx IVariableContext) {
	localctx = NewVariableContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 82, AsmE8ParserRULE_variable)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(366)
		p.Tag()
	}
	{
		p.SetState(367)
		p.Match(AsmE8ParserT__54)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(368)
		p.Imm()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IVariable_referenceContext is an interface to support dynamic dispatch.
type IVariable_referenceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Tag() ITagContext
	Reference() IReferenceContext

	// IsVariable_referenceContext differentiates from other interfaces.
	IsVariable_referenceContext()
}

type Variable_referenceContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariable_referenceContext() *Variable_referenceContext {
	var p = new(Variable_referenceContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsmE8ParserRULE_variable_reference
	return p
}

func InitEmptyVariable_referenceContext(p *Variable_referenceContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsmE8ParserRULE_variable_reference
}

func (*Variable_referenceContext) IsVariable_referenceContext() {}

func NewVariable_referenceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Variable_referenceContext {
	var p = new(Variable_referenceContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AsmE8ParserRULE_variable_reference

	return p
}

func (s *Variable_referenceContext) GetParser() antlr.Parser { return s.parser }

func (s *Variable_referenceContext) Tag() ITagContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITagContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITagContext)
}

func (s *Variable_referenceContext) Reference() IReferenceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReferenceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReferenceContext)
}

func (s *Variable_referenceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Variable_referenceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Variable_referenceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsmE8Listener); ok {
		listenerT.EnterVariable_reference(s)
	}
}

func (s *Variable_referenceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsmE8Listener); ok {
		listenerT.ExitVariable_reference(s)
	}
}

func (p *AsmE8Parser) Variable_reference() (localctx IVariable_referenceContext) {
	localctx = NewVariable_referenceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 84, AsmE8ParserRULE_variable_reference)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(370)
		p.Tag()
	}
	{
		p.SetState(371)
		p.Match(AsmE8ParserT__54)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(372)
		p.Reference()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IReferenceContext is an interface to support dynamic dispatch.
type IReferenceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	STR() antlr.TerminalNode
	Imm() IImmContext

	// IsReferenceContext differentiates from other interfaces.
	IsReferenceContext()
}

type ReferenceContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReferenceContext() *ReferenceContext {
	var p = new(ReferenceContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsmE8ParserRULE_reference
	return p
}

func InitEmptyReferenceContext(p *ReferenceContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsmE8ParserRULE_reference
}

func (*ReferenceContext) IsReferenceContext() {}

func NewReferenceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReferenceContext {
	var p = new(ReferenceContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AsmE8ParserRULE_reference

	return p
}

func (s *ReferenceContext) GetParser() antlr.Parser { return s.parser }

func (s *ReferenceContext) STR() antlr.TerminalNode {
	return s.GetToken(AsmE8ParserSTR, 0)
}

func (s *ReferenceContext) Imm() IImmContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IImmContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IImmContext)
}

func (s *ReferenceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReferenceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReferenceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsmE8Listener); ok {
		listenerT.EnterReference(s)
	}
}

func (s *ReferenceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsmE8Listener); ok {
		listenerT.ExitReference(s)
	}
}

func (p *AsmE8Parser) Reference() (localctx IReferenceContext) {
	localctx = NewReferenceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 86, AsmE8ParserRULE_reference)
	var _la int

	p.SetState(378)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(374)
			p.Match(AsmE8ParserSTR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(375)
			p.Match(AsmE8ParserSTR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(376)
			_la = p.GetTokenStream().LA(1)

			if !(_la == AsmE8ParserT__52 || _la == AsmE8ParserT__53) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(377)
			p.Imm()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDirectivesContext is an interface to support dynamic dispatch.
type IDirectivesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Ascii() IAsciiContext
	Imm_list() IImm_listContext

	// IsDirectivesContext differentiates from other interfaces.
	IsDirectivesContext()
}

type DirectivesContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDirectivesContext() *DirectivesContext {
	var p = new(DirectivesContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsmE8ParserRULE_directives
	return p
}

func InitEmptyDirectivesContext(p *DirectivesContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsmE8ParserRULE_directives
}

func (*DirectivesContext) IsDirectivesContext() {}

func NewDirectivesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DirectivesContext {
	var p = new(DirectivesContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AsmE8ParserRULE_directives

	return p
}

func (s *DirectivesContext) GetParser() antlr.Parser { return s.parser }

func (s *DirectivesContext) Ascii() IAsciiContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAsciiContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAsciiContext)
}

func (s *DirectivesContext) Imm_list() IImm_listContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IImm_listContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IImm_listContext)
}

func (s *DirectivesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DirectivesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DirectivesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsmE8Listener); ok {
		listenerT.EnterDirectives(s)
	}
}

func (s *DirectivesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsmE8Listener); ok {
		listenerT.ExitDirectives(s)
	}
}

func (p *AsmE8Parser) Directives() (localctx IDirectivesContext) {
	localctx = NewDirectivesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 88, AsmE8ParserRULE_directives)
	p.SetState(392)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case AsmE8ParserT__55:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(380)
			p.Match(AsmE8ParserT__55)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(381)
			p.Ascii()
		}

	case AsmE8ParserT__56:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(382)
			p.Match(AsmE8ParserT__56)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(383)
			p.Ascii()
		}

	case AsmE8ParserT__57:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(384)
			p.Match(AsmE8ParserT__57)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(385)
			p.Imm_list()
		}

	case AsmE8ParserT__58:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(386)
			p.Match(AsmE8ParserT__58)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(387)
			p.Imm_list()
		}

	case AsmE8ParserT__59:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(388)
			p.Match(AsmE8ParserT__59)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(389)
			p.Imm_list()
		}

	case AsmE8ParserT__60:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(390)
			p.Match(AsmE8ParserT__60)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(391)
			p.Imm_list()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISegmentContext is an interface to support dynamic dispatch.
type ISegmentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ANY_STR() antlr.TerminalNode

	// IsSegmentContext differentiates from other interfaces.
	IsSegmentContext()
}

type SegmentContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySegmentContext() *SegmentContext {
	var p = new(SegmentContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsmE8ParserRULE_segment
	return p
}

func InitEmptySegmentContext(p *SegmentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsmE8ParserRULE_segment
}

func (*SegmentContext) IsSegmentContext() {}

func NewSegmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SegmentContext {
	var p = new(SegmentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AsmE8ParserRULE_segment

	return p
}

func (s *SegmentContext) GetParser() antlr.Parser { return s.parser }

func (s *SegmentContext) ANY_STR() antlr.TerminalNode {
	return s.GetToken(AsmE8ParserANY_STR, 0)
}

func (s *SegmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SegmentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SegmentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsmE8Listener); ok {
		listenerT.EnterSegment(s)
	}
}

func (s *SegmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsmE8Listener); ok {
		listenerT.ExitSegment(s)
	}
}

func (p *AsmE8Parser) Segment() (localctx ISegmentContext) {
	localctx = NewSegmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 90, AsmE8ParserRULE_segment)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(394)
		p.Match(AsmE8ParserT__61)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(395)
		p.Match(AsmE8ParserANY_STR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAccessContext is an interface to support dynamic dispatch.
type IAccessContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Tag() ITagContext

	// IsAccessContext differentiates from other interfaces.
	IsAccessContext()
}

type AccessContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAccessContext() *AccessContext {
	var p = new(AccessContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsmE8ParserRULE_access
	return p
}

func InitEmptyAccessContext(p *AccessContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsmE8ParserRULE_access
}

func (*AccessContext) IsAccessContext() {}

func NewAccessContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AccessContext {
	var p = new(AccessContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AsmE8ParserRULE_access

	return p
}

func (s *AccessContext) GetParser() antlr.Parser { return s.parser }

func (s *AccessContext) Tag() ITagContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITagContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITagContext)
}

func (s *AccessContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AccessContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AccessContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsmE8Listener); ok {
		listenerT.EnterAccess(s)
	}
}

func (s *AccessContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsmE8Listener); ok {
		listenerT.ExitAccess(s)
	}
}

func (p *AsmE8Parser) Access() (localctx IAccessContext) {
	localctx = NewAccessContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 92, AsmE8ParserRULE_access)
	p.SetState(403)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(397)
			p.Match(AsmE8ParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(398)
			p.Match(AsmE8ParserT__62)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(399)
			p.Tag()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(400)
			p.Match(AsmE8ParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(401)
			p.Match(AsmE8ParserT__63)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(402)
			p.Tag()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IImm_listContext is an interface to support dynamic dispatch.
type IImm_listContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllImm() []IImmContext
	Imm(i int) IImmContext
	AllTag() []ITagContext
	Tag(i int) ITagContext

	// IsImm_listContext differentiates from other interfaces.
	IsImm_listContext()
}

type Imm_listContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyImm_listContext() *Imm_listContext {
	var p = new(Imm_listContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsmE8ParserRULE_imm_list
	return p
}

func InitEmptyImm_listContext(p *Imm_listContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsmE8ParserRULE_imm_list
}

func (*Imm_listContext) IsImm_listContext() {}

func NewImm_listContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Imm_listContext {
	var p = new(Imm_listContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AsmE8ParserRULE_imm_list

	return p
}

func (s *Imm_listContext) GetParser() antlr.Parser { return s.parser }

func (s *Imm_listContext) AllImm() []IImmContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IImmContext); ok {
			len++
		}
	}

	tst := make([]IImmContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IImmContext); ok {
			tst[i] = t.(IImmContext)
			i++
		}
	}

	return tst
}

func (s *Imm_listContext) Imm(i int) IImmContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IImmContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IImmContext)
}

func (s *Imm_listContext) AllTag() []ITagContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITagContext); ok {
			len++
		}
	}

	tst := make([]ITagContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITagContext); ok {
			tst[i] = t.(ITagContext)
			i++
		}
	}

	return tst
}

func (s *Imm_listContext) Tag(i int) ITagContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITagContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITagContext)
}

func (s *Imm_listContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Imm_listContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Imm_listContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsmE8Listener); ok {
		listenerT.EnterImm_list(s)
	}
}

func (s *Imm_listContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsmE8Listener); ok {
		listenerT.ExitImm_list(s)
	}
}

func (p *AsmE8Parser) Imm_list() (localctx IImm_listContext) {
	localctx = NewImm_listContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 94, AsmE8ParserRULE_imm_list)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(407)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case AsmE8ParserBINARY, AsmE8ParserINT, AsmE8ParserCHAR, AsmE8ParserHEX:
		{
			p.SetState(405)
			p.Imm()
		}

	case AsmE8ParserSTR:
		{
			p.SetState(406)
			p.Tag()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.SetState(416)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == AsmE8ParserT__6 {
		{
			p.SetState(409)
			p.Match(AsmE8ParserT__6)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(412)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case AsmE8ParserBINARY, AsmE8ParserINT, AsmE8ParserCHAR, AsmE8ParserHEX:
			{
				p.SetState(410)
				p.Imm()
			}

		case AsmE8ParserSTR:
			{
				p.SetState(411)
				p.Tag()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(418)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IImmContext is an interface to support dynamic dispatch.
type IImmContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	INT() antlr.TerminalNode
	HEX() antlr.TerminalNode
	BINARY() antlr.TerminalNode
	CHAR() antlr.TerminalNode

	// IsImmContext differentiates from other interfaces.
	IsImmContext()
}

type ImmContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyImmContext() *ImmContext {
	var p = new(ImmContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsmE8ParserRULE_imm
	return p
}

func InitEmptyImmContext(p *ImmContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsmE8ParserRULE_imm
}

func (*ImmContext) IsImmContext() {}

func NewImmContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ImmContext {
	var p = new(ImmContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AsmE8ParserRULE_imm

	return p
}

func (s *ImmContext) GetParser() antlr.Parser { return s.parser }

func (s *ImmContext) INT() antlr.TerminalNode {
	return s.GetToken(AsmE8ParserINT, 0)
}

func (s *ImmContext) HEX() antlr.TerminalNode {
	return s.GetToken(AsmE8ParserHEX, 0)
}

func (s *ImmContext) BINARY() antlr.TerminalNode {
	return s.GetToken(AsmE8ParserBINARY, 0)
}

func (s *ImmContext) CHAR() antlr.TerminalNode {
	return s.GetToken(AsmE8ParserCHAR, 0)
}

func (s *ImmContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ImmContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ImmContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsmE8Listener); ok {
		listenerT.EnterImm(s)
	}
}

func (s *ImmContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsmE8Listener); ok {
		listenerT.ExitImm(s)
	}
}

func (p *AsmE8Parser) Imm() (localctx IImmContext) {
	localctx = NewImmContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 96, AsmE8ParserRULE_imm)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(419)
		_la = p.GetTokenStream().LA(1)

		if !((int64((_la-68)) & ^0x3f) == 0 && ((int64(1)<<(_la-68))&23) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAsciiContext is an interface to support dynamic dispatch.
type IAsciiContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ANY_STR() antlr.TerminalNode

	// IsAsciiContext differentiates from other interfaces.
	IsAsciiContext()
}

type AsciiContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAsciiContext() *AsciiContext {
	var p = new(AsciiContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsmE8ParserRULE_ascii
	return p
}

func InitEmptyAsciiContext(p *AsciiContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsmE8ParserRULE_ascii
}

func (*AsciiContext) IsAsciiContext() {}

func NewAsciiContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AsciiContext {
	var p = new(AsciiContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AsmE8ParserRULE_ascii

	return p
}

func (s *AsciiContext) GetParser() antlr.Parser { return s.parser }

func (s *AsciiContext) ANY_STR() antlr.TerminalNode {
	return s.GetToken(AsmE8ParserANY_STR, 0)
}

func (s *AsciiContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AsciiContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AsciiContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsmE8Listener); ok {
		listenerT.EnterAscii(s)
	}
}

func (s *AsciiContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsmE8Listener); ok {
		listenerT.ExitAscii(s)
	}
}

func (p *AsmE8Parser) Ascii() (localctx IAsciiContext) {
	localctx = NewAsciiContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 98, AsmE8ParserRULE_ascii)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(421)
		p.Match(AsmE8ParserANY_STR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITagContext is an interface to support dynamic dispatch.
type ITagContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	STR() antlr.TerminalNode

	// IsTagContext differentiates from other interfaces.
	IsTagContext()
}

type TagContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTagContext() *TagContext {
	var p = new(TagContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsmE8ParserRULE_tag
	return p
}

func InitEmptyTagContext(p *TagContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsmE8ParserRULE_tag
}

func (*TagContext) IsTagContext() {}

func NewTagContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TagContext {
	var p = new(TagContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AsmE8ParserRULE_tag

	return p
}

func (s *TagContext) GetParser() antlr.Parser { return s.parser }

func (s *TagContext) STR() antlr.TerminalNode {
	return s.GetToken(AsmE8ParserSTR, 0)
}

func (s *TagContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TagContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TagContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsmE8Listener); ok {
		listenerT.EnterTag(s)
	}
}

func (s *TagContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsmE8Listener); ok {
		listenerT.ExitTag(s)
	}
}

func (p *AsmE8Parser) Tag() (localctx ITagContext) {
	localctx = NewTagContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 100, AsmE8ParserRULE_tag)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(423)
		p.Match(AsmE8ParserSTR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}
