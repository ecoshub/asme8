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
		"'sp'", "'sr'", "'['", "']'", "'+'", "'-'", "'='", "'.asciiz '", "'.ascii '",
		"'.byte '", "'.word '", "'.resb '", "'.org '", "'.segment '", "'global '",
		"'extern '",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "LINE_COMMENT",
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
		"addr_mode_reg16_stack", "addr_mode_implied_stack", "addr_mode_implied_status_register",
		"addr_mode_stack_imm8", "addr_mode_stack_imm8_tag", "addr_mode_stack_reg16",
		"mnemonic", "reg8", "reg16", "stack", "status_register", "direct", "direct_virtual_offset",
		"direct_offset", "indirect_offset", "reg16_indexed", "indirect", "variable",
		"variable_reference", "reference", "directives", "segment", "access",
		"imm_list", "imm", "ascii", "tag",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 74, 437, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36, 7, 36,
		2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7, 41, 2,
		42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 2, 46, 7, 46, 2, 47,
		7, 47, 2, 48, 7, 48, 2, 49, 7, 49, 2, 50, 7, 50, 2, 51, 7, 51, 2, 52, 7,
		52, 1, 0, 1, 0, 1, 0, 1, 1, 4, 1, 111, 8, 1, 11, 1, 12, 1, 112, 1, 1, 3,
		1, 116, 8, 1, 1, 1, 4, 1, 119, 8, 1, 11, 1, 12, 1, 120, 1, 1, 3, 1, 124,
		8, 1, 3, 1, 126, 8, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 141, 8, 2, 1, 3, 1, 3, 1, 3, 1, 4,
		1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4,
		1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4,
		1, 4, 1, 4, 3, 4, 173, 8, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6,
		1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8,
		1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1,
		10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12,
		1, 12, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1,
		14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 16, 1, 16,
		1, 16, 1, 16, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1,
		18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19,
		1, 19, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 21, 1, 21, 1, 22, 1,
		22, 1, 22, 1, 22, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 24, 1, 24,
		1, 24, 1, 24, 1, 24, 1, 24, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1,
		26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 27, 1, 27, 1, 27, 1, 27, 1, 28,
		1, 28, 1, 28, 1, 28, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 30, 1,
		30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31,
		1, 32, 1, 32, 1, 33, 1, 33, 1, 34, 1, 34, 1, 35, 1, 35, 1, 36, 1, 36, 1,
		37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 3, 37, 340,
		8, 37, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 39, 1, 39, 1, 39, 1,
		39, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 3, 39, 360,
		8, 39, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 41, 1, 41, 1, 41, 1,
		41, 1, 41, 1, 41, 1, 42, 1, 42, 1, 42, 1, 42, 1, 43, 1, 43, 1, 43, 1, 43,
		1, 44, 1, 44, 1, 44, 1, 44, 1, 45, 1, 45, 1, 45, 1, 45, 3, 45, 390, 8,
		45, 1, 46, 1, 46, 1, 46, 1, 46, 1, 46, 1, 46, 1, 46, 1, 46, 1, 46, 1, 46,
		1, 46, 1, 46, 3, 46, 404, 8, 46, 1, 47, 1, 47, 1, 47, 1, 48, 1, 48, 1,
		48, 1, 48, 1, 48, 1, 48, 3, 48, 415, 8, 48, 1, 49, 1, 49, 3, 49, 419, 8,
		49, 1, 49, 1, 49, 1, 49, 3, 49, 424, 8, 49, 5, 49, 426, 8, 49, 10, 49,
		12, 49, 429, 9, 49, 1, 50, 1, 50, 1, 51, 1, 51, 1, 52, 1, 52, 1, 52, 0,
		0, 53, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34,
		36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 70,
		72, 74, 76, 78, 80, 82, 84, 86, 88, 90, 92, 94, 96, 98, 100, 102, 104,
		0, 6, 1, 0, 1, 3, 1, 0, 8, 38, 1, 0, 39, 47, 1, 0, 48, 49, 1, 0, 54, 55,
		2, 0, 69, 71, 73, 73, 435, 0, 106, 1, 0, 0, 0, 2, 125, 1, 0, 0, 0, 4, 140,
		1, 0, 0, 0, 6, 142, 1, 0, 0, 0, 8, 172, 1, 0, 0, 0, 10, 174, 1, 0, 0, 0,
		12, 178, 1, 0, 0, 0, 14, 182, 1, 0, 0, 0, 16, 188, 1, 0, 0, 0, 18, 194,
		1, 0, 0, 0, 20, 200, 1, 0, 0, 0, 22, 206, 1, 0, 0, 0, 24, 212, 1, 0, 0,
		0, 26, 216, 1, 0, 0, 0, 28, 222, 1, 0, 0, 0, 30, 228, 1, 0, 0, 0, 32, 234,
		1, 0, 0, 0, 34, 240, 1, 0, 0, 0, 36, 246, 1, 0, 0, 0, 38, 252, 1, 0, 0,
		0, 40, 258, 1, 0, 0, 0, 42, 264, 1, 0, 0, 0, 44, 266, 1, 0, 0, 0, 46, 270,
		1, 0, 0, 0, 48, 276, 1, 0, 0, 0, 50, 282, 1, 0, 0, 0, 52, 288, 1, 0, 0,
		0, 54, 294, 1, 0, 0, 0, 56, 298, 1, 0, 0, 0, 58, 302, 1, 0, 0, 0, 60, 308,
		1, 0, 0, 0, 62, 314, 1, 0, 0, 0, 64, 320, 1, 0, 0, 0, 66, 322, 1, 0, 0,
		0, 68, 324, 1, 0, 0, 0, 70, 326, 1, 0, 0, 0, 72, 328, 1, 0, 0, 0, 74, 339,
		1, 0, 0, 0, 76, 341, 1, 0, 0, 0, 78, 359, 1, 0, 0, 0, 80, 361, 1, 0, 0,
		0, 82, 367, 1, 0, 0, 0, 84, 373, 1, 0, 0, 0, 86, 377, 1, 0, 0, 0, 88, 381,
		1, 0, 0, 0, 90, 389, 1, 0, 0, 0, 92, 403, 1, 0, 0, 0, 94, 405, 1, 0, 0,
		0, 96, 414, 1, 0, 0, 0, 98, 418, 1, 0, 0, 0, 100, 430, 1, 0, 0, 0, 102,
		432, 1, 0, 0, 0, 104, 434, 1, 0, 0, 0, 106, 107, 3, 2, 1, 0, 107, 108,
		5, 0, 0, 1, 108, 1, 1, 0, 0, 0, 109, 111, 3, 4, 2, 0, 110, 109, 1, 0, 0,
		0, 111, 112, 1, 0, 0, 0, 112, 110, 1, 0, 0, 0, 112, 113, 1, 0, 0, 0, 113,
		115, 1, 0, 0, 0, 114, 116, 5, 66, 0, 0, 115, 114, 1, 0, 0, 0, 115, 116,
		1, 0, 0, 0, 116, 126, 1, 0, 0, 0, 117, 119, 3, 4, 2, 0, 118, 117, 1, 0,
		0, 0, 119, 120, 1, 0, 0, 0, 120, 118, 1, 0, 0, 0, 120, 121, 1, 0, 0, 0,
		121, 123, 1, 0, 0, 0, 122, 124, 5, 67, 0, 0, 123, 122, 1, 0, 0, 0, 123,
		124, 1, 0, 0, 0, 124, 126, 1, 0, 0, 0, 125, 110, 1, 0, 0, 0, 125, 118,
		1, 0, 0, 0, 126, 3, 1, 0, 0, 0, 127, 141, 3, 88, 44, 0, 128, 141, 3, 86,
		43, 0, 129, 130, 7, 0, 0, 0, 130, 141, 3, 8, 4, 0, 131, 141, 3, 94, 47,
		0, 132, 141, 3, 96, 48, 0, 133, 134, 3, 6, 3, 0, 134, 135, 5, 4, 0, 0,
		135, 136, 3, 92, 46, 0, 136, 141, 1, 0, 0, 0, 137, 141, 3, 92, 46, 0, 138,
		141, 3, 6, 3, 0, 139, 141, 5, 5, 0, 0, 140, 127, 1, 0, 0, 0, 140, 128,
		1, 0, 0, 0, 140, 129, 1, 0, 0, 0, 140, 131, 1, 0, 0, 0, 140, 132, 1, 0,
		0, 0, 140, 133, 1, 0, 0, 0, 140, 137, 1, 0, 0, 0, 140, 138, 1, 0, 0, 0,
		140, 139, 1, 0, 0, 0, 141, 5, 1, 0, 0, 0, 142, 143, 3, 104, 52, 0, 143,
		144, 5, 6, 0, 0, 144, 7, 1, 0, 0, 0, 145, 173, 3, 10, 5, 0, 146, 173, 3,
		12, 6, 0, 147, 173, 3, 14, 7, 0, 148, 173, 3, 16, 8, 0, 149, 173, 3, 18,
		9, 0, 150, 173, 3, 20, 10, 0, 151, 173, 3, 22, 11, 0, 152, 173, 3, 24,
		12, 0, 153, 173, 3, 26, 13, 0, 154, 173, 3, 28, 14, 0, 155, 173, 3, 30,
		15, 0, 156, 173, 3, 32, 16, 0, 157, 173, 3, 34, 17, 0, 158, 173, 3, 36,
		18, 0, 159, 173, 3, 38, 19, 0, 160, 173, 3, 40, 20, 0, 161, 173, 3, 42,
		21, 0, 162, 173, 3, 44, 22, 0, 163, 173, 3, 46, 23, 0, 164, 173, 3, 48,
		24, 0, 165, 173, 3, 50, 25, 0, 166, 173, 3, 52, 26, 0, 167, 173, 3, 54,
		27, 0, 168, 173, 3, 56, 28, 0, 169, 173, 3, 58, 29, 0, 170, 173, 3, 60,
		30, 0, 171, 173, 3, 62, 31, 0, 172, 145, 1, 0, 0, 0, 172, 146, 1, 0, 0,
		0, 172, 147, 1, 0, 0, 0, 172, 148, 1, 0, 0, 0, 172, 149, 1, 0, 0, 0, 172,
		150, 1, 0, 0, 0, 172, 151, 1, 0, 0, 0, 172, 152, 1, 0, 0, 0, 172, 153,
		1, 0, 0, 0, 172, 154, 1, 0, 0, 0, 172, 155, 1, 0, 0, 0, 172, 156, 1, 0,
		0, 0, 172, 157, 1, 0, 0, 0, 172, 158, 1, 0, 0, 0, 172, 159, 1, 0, 0, 0,
		172, 160, 1, 0, 0, 0, 172, 161, 1, 0, 0, 0, 172, 162, 1, 0, 0, 0, 172,
		163, 1, 0, 0, 0, 172, 164, 1, 0, 0, 0, 172, 165, 1, 0, 0, 0, 172, 166,
		1, 0, 0, 0, 172, 167, 1, 0, 0, 0, 172, 168, 1, 0, 0, 0, 172, 169, 1, 0,
		0, 0, 172, 170, 1, 0, 0, 0, 172, 171, 1, 0, 0, 0, 173, 9, 1, 0, 0, 0, 174,
		175, 3, 64, 32, 0, 175, 176, 5, 4, 0, 0, 176, 177, 3, 100, 50, 0, 177,
		11, 1, 0, 0, 0, 178, 179, 3, 64, 32, 0, 179, 180, 5, 4, 0, 0, 180, 181,
		3, 104, 52, 0, 181, 13, 1, 0, 0, 0, 182, 183, 3, 64, 32, 0, 183, 184, 5,
		4, 0, 0, 184, 185, 3, 78, 39, 0, 185, 186, 5, 7, 0, 0, 186, 187, 3, 66,
		33, 0, 187, 15, 1, 0, 0, 0, 188, 189, 3, 64, 32, 0, 189, 190, 5, 4, 0,
		0, 190, 191, 3, 74, 37, 0, 191, 192, 5, 7, 0, 0, 192, 193, 3, 66, 33, 0,
		193, 17, 1, 0, 0, 0, 194, 195, 3, 64, 32, 0, 195, 196, 5, 4, 0, 0, 196,
		197, 3, 80, 40, 0, 197, 198, 5, 7, 0, 0, 198, 199, 3, 66, 33, 0, 199, 19,
		1, 0, 0, 0, 200, 201, 3, 64, 32, 0, 201, 202, 5, 4, 0, 0, 202, 203, 3,
		82, 41, 0, 203, 204, 5, 7, 0, 0, 204, 205, 3, 66, 33, 0, 205, 21, 1, 0,
		0, 0, 206, 207, 3, 64, 32, 0, 207, 208, 5, 4, 0, 0, 208, 209, 3, 84, 42,
		0, 209, 210, 5, 7, 0, 0, 210, 211, 3, 66, 33, 0, 211, 23, 1, 0, 0, 0, 212,
		213, 3, 64, 32, 0, 213, 214, 5, 4, 0, 0, 214, 215, 3, 66, 33, 0, 215, 25,
		1, 0, 0, 0, 216, 217, 3, 64, 32, 0, 217, 218, 5, 4, 0, 0, 218, 219, 3,
		66, 33, 0, 219, 220, 5, 7, 0, 0, 220, 221, 3, 100, 50, 0, 221, 27, 1, 0,
		0, 0, 222, 223, 3, 64, 32, 0, 223, 224, 5, 4, 0, 0, 224, 225, 3, 66, 33,
		0, 225, 226, 5, 7, 0, 0, 226, 227, 3, 104, 52, 0, 227, 29, 1, 0, 0, 0,
		228, 229, 3, 64, 32, 0, 229, 230, 5, 4, 0, 0, 230, 231, 3, 66, 33, 0, 231,
		232, 5, 7, 0, 0, 232, 233, 3, 78, 39, 0, 233, 31, 1, 0, 0, 0, 234, 235,
		3, 64, 32, 0, 235, 236, 5, 4, 0, 0, 236, 237, 3, 66, 33, 0, 237, 238, 5,
		7, 0, 0, 238, 239, 3, 74, 37, 0, 239, 33, 1, 0, 0, 0, 240, 241, 3, 64,
		32, 0, 241, 242, 5, 4, 0, 0, 242, 243, 3, 66, 33, 0, 243, 244, 5, 7, 0,
		0, 244, 245, 3, 80, 40, 0, 245, 35, 1, 0, 0, 0, 246, 247, 3, 64, 32, 0,
		247, 248, 5, 4, 0, 0, 248, 249, 3, 66, 33, 0, 249, 250, 5, 7, 0, 0, 250,
		251, 3, 82, 41, 0, 251, 37, 1, 0, 0, 0, 252, 253, 3, 64, 32, 0, 253, 254,
		5, 4, 0, 0, 254, 255, 3, 66, 33, 0, 255, 256, 5, 7, 0, 0, 256, 257, 3,
		84, 42, 0, 257, 39, 1, 0, 0, 0, 258, 259, 3, 64, 32, 0, 259, 260, 5, 4,
		0, 0, 260, 261, 3, 66, 33, 0, 261, 262, 5, 7, 0, 0, 262, 263, 3, 66, 33,
		0, 263, 41, 1, 0, 0, 0, 264, 265, 3, 64, 32, 0, 265, 43, 1, 0, 0, 0, 266,
		267, 3, 64, 32, 0, 267, 268, 5, 4, 0, 0, 268, 269, 3, 68, 34, 0, 269, 45,
		1, 0, 0, 0, 270, 271, 3, 64, 32, 0, 271, 272, 5, 4, 0, 0, 272, 273, 3,
		68, 34, 0, 273, 274, 5, 7, 0, 0, 274, 275, 3, 100, 50, 0, 275, 47, 1, 0,
		0, 0, 276, 277, 3, 64, 32, 0, 277, 278, 5, 4, 0, 0, 278, 279, 3, 68, 34,
		0, 279, 280, 5, 7, 0, 0, 280, 281, 3, 104, 52, 0, 281, 49, 1, 0, 0, 0,
		282, 283, 3, 64, 32, 0, 283, 284, 5, 4, 0, 0, 284, 285, 3, 68, 34, 0, 285,
		286, 5, 7, 0, 0, 286, 287, 3, 68, 34, 0, 287, 51, 1, 0, 0, 0, 288, 289,
		3, 64, 32, 0, 289, 290, 5, 4, 0, 0, 290, 291, 3, 68, 34, 0, 291, 292, 5,
		7, 0, 0, 292, 293, 3, 70, 35, 0, 293, 53, 1, 0, 0, 0, 294, 295, 3, 64,
		32, 0, 295, 296, 5, 4, 0, 0, 296, 297, 3, 70, 35, 0, 297, 55, 1, 0, 0,
		0, 298, 299, 3, 64, 32, 0, 299, 300, 5, 4, 0, 0, 300, 301, 3, 72, 36, 0,
		301, 57, 1, 0, 0, 0, 302, 303, 3, 64, 32, 0, 303, 304, 5, 4, 0, 0, 304,
		305, 3, 70, 35, 0, 305, 306, 5, 7, 0, 0, 306, 307, 3, 100, 50, 0, 307,
		59, 1, 0, 0, 0, 308, 309, 3, 64, 32, 0, 309, 310, 5, 4, 0, 0, 310, 311,
		3, 70, 35, 0, 311, 312, 5, 7, 0, 0, 312, 313, 3, 104, 52, 0, 313, 61, 1,
		0, 0, 0, 314, 315, 3, 64, 32, 0, 315, 316, 5, 4, 0, 0, 316, 317, 3, 70,
		35, 0, 317, 318, 5, 7, 0, 0, 318, 319, 3, 68, 34, 0, 319, 63, 1, 0, 0,
		0, 320, 321, 7, 1, 0, 0, 321, 65, 1, 0, 0, 0, 322, 323, 7, 2, 0, 0, 323,
		67, 1, 0, 0, 0, 324, 325, 7, 3, 0, 0, 325, 69, 1, 0, 0, 0, 326, 327, 5,
		50, 0, 0, 327, 71, 1, 0, 0, 0, 328, 329, 5, 51, 0, 0, 329, 73, 1, 0, 0,
		0, 330, 331, 5, 52, 0, 0, 331, 332, 3, 100, 50, 0, 332, 333, 5, 53, 0,
		0, 333, 340, 1, 0, 0, 0, 334, 335, 5, 52, 0, 0, 335, 336, 3, 104, 52, 0,
		336, 337, 5, 53, 0, 0, 337, 340, 1, 0, 0, 0, 338, 340, 3, 76, 38, 0, 339,
		330, 1, 0, 0, 0, 339, 334, 1, 0, 0, 0, 339, 338, 1, 0, 0, 0, 340, 75, 1,
		0, 0, 0, 341, 342, 5, 52, 0, 0, 342, 343, 5, 74, 0, 0, 343, 344, 7, 4,
		0, 0, 344, 345, 5, 70, 0, 0, 345, 346, 5, 53, 0, 0, 346, 77, 1, 0, 0, 0,
		347, 348, 5, 52, 0, 0, 348, 349, 3, 100, 50, 0, 349, 350, 5, 54, 0, 0,
		350, 351, 3, 66, 33, 0, 351, 352, 5, 53, 0, 0, 352, 360, 1, 0, 0, 0, 353,
		354, 5, 52, 0, 0, 354, 355, 3, 104, 52, 0, 355, 356, 5, 54, 0, 0, 356,
		357, 3, 66, 33, 0, 357, 358, 5, 53, 0, 0, 358, 360, 1, 0, 0, 0, 359, 347,
		1, 0, 0, 0, 359, 353, 1, 0, 0, 0, 360, 79, 1, 0, 0, 0, 361, 362, 5, 52,
		0, 0, 362, 363, 3, 68, 34, 0, 363, 364, 5, 54, 0, 0, 364, 365, 3, 100,
		50, 0, 365, 366, 5, 53, 0, 0, 366, 81, 1, 0, 0, 0, 367, 368, 5, 52, 0,
		0, 368, 369, 3, 68, 34, 0, 369, 370, 5, 54, 0, 0, 370, 371, 3, 66, 33,
		0, 371, 372, 5, 53, 0, 0, 372, 83, 1, 0, 0, 0, 373, 374, 5, 52, 0, 0, 374,
		375, 3, 68, 34, 0, 375, 376, 5, 53, 0, 0, 376, 85, 1, 0, 0, 0, 377, 378,
		3, 104, 52, 0, 378, 379, 5, 56, 0, 0, 379, 380, 3, 100, 50, 0, 380, 87,
		1, 0, 0, 0, 381, 382, 3, 104, 52, 0, 382, 383, 5, 56, 0, 0, 383, 384, 3,
		90, 45, 0, 384, 89, 1, 0, 0, 0, 385, 390, 5, 74, 0, 0, 386, 387, 5, 74,
		0, 0, 387, 388, 7, 4, 0, 0, 388, 390, 3, 100, 50, 0, 389, 385, 1, 0, 0,
		0, 389, 386, 1, 0, 0, 0, 390, 91, 1, 0, 0, 0, 391, 392, 5, 57, 0, 0, 392,
		404, 3, 102, 51, 0, 393, 394, 5, 58, 0, 0, 394, 404, 3, 102, 51, 0, 395,
		396, 5, 59, 0, 0, 396, 404, 3, 98, 49, 0, 397, 398, 5, 60, 0, 0, 398, 404,
		3, 98, 49, 0, 399, 400, 5, 61, 0, 0, 400, 404, 3, 98, 49, 0, 401, 402,
		5, 62, 0, 0, 402, 404, 3, 98, 49, 0, 403, 391, 1, 0, 0, 0, 403, 393, 1,
		0, 0, 0, 403, 395, 1, 0, 0, 0, 403, 397, 1, 0, 0, 0, 403, 399, 1, 0, 0,
		0, 403, 401, 1, 0, 0, 0, 404, 93, 1, 0, 0, 0, 405, 406, 5, 63, 0, 0, 406,
		407, 5, 72, 0, 0, 407, 95, 1, 0, 0, 0, 408, 409, 5, 2, 0, 0, 409, 410,
		5, 64, 0, 0, 410, 415, 3, 104, 52, 0, 411, 412, 5, 2, 0, 0, 412, 413, 5,
		65, 0, 0, 413, 415, 3, 104, 52, 0, 414, 408, 1, 0, 0, 0, 414, 411, 1, 0,
		0, 0, 415, 97, 1, 0, 0, 0, 416, 419, 3, 100, 50, 0, 417, 419, 3, 104, 52,
		0, 418, 416, 1, 0, 0, 0, 418, 417, 1, 0, 0, 0, 419, 427, 1, 0, 0, 0, 420,
		423, 5, 7, 0, 0, 421, 424, 3, 100, 50, 0, 422, 424, 3, 104, 52, 0, 423,
		421, 1, 0, 0, 0, 423, 422, 1, 0, 0, 0, 424, 426, 1, 0, 0, 0, 425, 420,
		1, 0, 0, 0, 426, 429, 1, 0, 0, 0, 427, 425, 1, 0, 0, 0, 427, 428, 1, 0,
		0, 0, 428, 99, 1, 0, 0, 0, 429, 427, 1, 0, 0, 0, 430, 431, 7, 5, 0, 0,
		431, 101, 1, 0, 0, 0, 432, 433, 5, 72, 0, 0, 433, 103, 1, 0, 0, 0, 434,
		435, 5, 74, 0, 0, 435, 105, 1, 0, 0, 0, 15, 112, 115, 120, 123, 125, 140,
		172, 339, 359, 389, 403, 414, 418, 423, 427,
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
	AsmE8ParserT__64             = 65
	AsmE8ParserLINE_COMMENT      = 66
	AsmE8ParserSUFFIX_WHITESPACE = 67
	AsmE8ParserWHITE_SPACE       = 68
	AsmE8ParserBINARY            = 69
	AsmE8ParserINT               = 70
	AsmE8ParserCHAR              = 71
	AsmE8ParserANY_STR           = 72
	AsmE8ParserHEX               = 73
	AsmE8ParserSTR               = 74
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
	AsmE8ParserRULE_addr_mode_implied_status_register    = 28
	AsmE8ParserRULE_addr_mode_stack_imm8                 = 29
	AsmE8ParserRULE_addr_mode_stack_imm8_tag             = 30
	AsmE8ParserRULE_addr_mode_stack_reg16                = 31
	AsmE8ParserRULE_mnemonic                             = 32
	AsmE8ParserRULE_reg8                                 = 33
	AsmE8ParserRULE_reg16                                = 34
	AsmE8ParserRULE_stack                                = 35
	AsmE8ParserRULE_status_register                      = 36
	AsmE8ParserRULE_direct                               = 37
	AsmE8ParserRULE_direct_virtual_offset                = 38
	AsmE8ParserRULE_direct_offset                        = 39
	AsmE8ParserRULE_indirect_offset                      = 40
	AsmE8ParserRULE_reg16_indexed                        = 41
	AsmE8ParserRULE_indirect                             = 42
	AsmE8ParserRULE_variable                             = 43
	AsmE8ParserRULE_variable_reference                   = 44
	AsmE8ParserRULE_reference                            = 45
	AsmE8ParserRULE_directives                           = 46
	AsmE8ParserRULE_segment                              = 47
	AsmE8ParserRULE_access                               = 48
	AsmE8ParserRULE_imm_list                             = 49
	AsmE8ParserRULE_imm                                  = 50
	AsmE8ParserRULE_ascii                                = 51
	AsmE8ParserRULE_tag                                  = 52
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
		p.SetState(106)
		p.Instruction()
	}
	{
		p.SetState(107)
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

	p.SetState(125)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(110)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&-144115188075855826) != 0) || _la == AsmE8ParserSTR {
			{
				p.SetState(109)
				p.Line()
			}

			p.SetState(112)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		p.SetState(115)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == AsmE8ParserLINE_COMMENT {
			{
				p.SetState(114)
				p.Match(AsmE8ParserLINE_COMMENT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(118)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&-144115188075855826) != 0) || _la == AsmE8ParserSTR {
			{
				p.SetState(117)
				p.Line()
			}

			p.SetState(120)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		p.SetState(123)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == AsmE8ParserSUFFIX_WHITESPACE {
			{
				p.SetState(122)
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

	p.SetState(140)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(127)
			p.Variable_reference()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(128)
			p.Variable()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(129)
			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&14) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(130)
			p.Inst()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(131)
			p.Segment()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(132)
			p.Access()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(133)
			p.Label()
		}
		{
			p.SetState(134)
			p.Match(AsmE8ParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(135)
			p.Directives()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(137)
			p.Directives()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(138)
			p.Label()
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(139)
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
		p.SetState(142)
		p.Tag()
	}
	{
		p.SetState(143)
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
	Addr_mode_implied_status_register() IAddr_mode_implied_status_registerContext
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

func (s *InstContext) Addr_mode_implied_status_register() IAddr_mode_implied_status_registerContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAddr_mode_implied_status_registerContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAddr_mode_implied_status_registerContext)
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
	p.SetState(172)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(145)
			p.Addr_mode_imm16()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(146)
			p.Addr_mode_imm16_tag()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(147)
			p.Addr_mode_reg_to_mem_indexed()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(148)
			p.Addr_mode_reg_to_mem_direct()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(149)
			p.Addr_mode_reg_to_mem_indirect_offset()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(150)
			p.Addr_mode_reg_to_mem_reg16_indexed()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(151)
			p.Addr_mode_reg_to_mem_indirect()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(152)
			p.Addr_mode_implied_reg8()
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(153)
			p.Addr_mode_reg8_imm8()
		}

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(154)
			p.Addr_mode_reg8_imm8_tag()
		}

	case 11:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(155)
			p.Addr_mode_mem_to_reg_indexed()
		}

	case 12:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(156)
			p.Addr_mode_mem_to_reg_direct()
		}

	case 13:
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(157)
			p.Addr_mode_mem_to_reg_indirect_offset()
		}

	case 14:
		p.EnterOuterAlt(localctx, 14)
		{
			p.SetState(158)
			p.Addr_mode_mem_to_reg_reg16_indexed()
		}

	case 15:
		p.EnterOuterAlt(localctx, 15)
		{
			p.SetState(159)
			p.Addr_mode_mem_to_reg_indirect()
		}

	case 16:
		p.EnterOuterAlt(localctx, 16)
		{
			p.SetState(160)
			p.Addr_mode_reg8_reg8()
		}

	case 17:
		p.EnterOuterAlt(localctx, 17)
		{
			p.SetState(161)
			p.Addr_mode_implied()
		}

	case 18:
		p.EnterOuterAlt(localctx, 18)
		{
			p.SetState(162)
			p.Addr_mode_implied_reg16()
		}

	case 19:
		p.EnterOuterAlt(localctx, 19)
		{
			p.SetState(163)
			p.Addr_mode_reg16_imm()
		}

	case 20:
		p.EnterOuterAlt(localctx, 20)
		{
			p.SetState(164)
			p.Addr_mode_reg16_imm_tag()
		}

	case 21:
		p.EnterOuterAlt(localctx, 21)
		{
			p.SetState(165)
			p.Addr_mode_reg16_reg16()
		}

	case 22:
		p.EnterOuterAlt(localctx, 22)
		{
			p.SetState(166)
			p.Addr_mode_reg16_stack()
		}

	case 23:
		p.EnterOuterAlt(localctx, 23)
		{
			p.SetState(167)
			p.Addr_mode_implied_stack()
		}

	case 24:
		p.EnterOuterAlt(localctx, 24)
		{
			p.SetState(168)
			p.Addr_mode_implied_status_register()
		}

	case 25:
		p.EnterOuterAlt(localctx, 25)
		{
			p.SetState(169)
			p.Addr_mode_stack_imm8()
		}

	case 26:
		p.EnterOuterAlt(localctx, 26)
		{
			p.SetState(170)
			p.Addr_mode_stack_imm8_tag()
		}

	case 27:
		p.EnterOuterAlt(localctx, 27)
		{
			p.SetState(171)
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
		p.SetState(174)
		p.Mnemonic()
	}
	{
		p.SetState(175)
		p.Match(AsmE8ParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(176)
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
		p.SetState(178)
		p.Mnemonic()
	}
	{
		p.SetState(179)
		p.Match(AsmE8ParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(180)
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
		p.SetState(182)
		p.Mnemonic()
	}
	{
		p.SetState(183)
		p.Match(AsmE8ParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(184)
		p.Direct_offset()
	}
	{
		p.SetState(185)
		p.Match(AsmE8ParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(186)
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
		p.SetState(188)
		p.Mnemonic()
	}
	{
		p.SetState(189)
		p.Match(AsmE8ParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(190)
		p.Direct()
	}
	{
		p.SetState(191)
		p.Match(AsmE8ParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(192)
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
		p.SetState(194)
		p.Mnemonic()
	}
	{
		p.SetState(195)
		p.Match(AsmE8ParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(196)
		p.Indirect_offset()
	}
	{
		p.SetState(197)
		p.Match(AsmE8ParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(198)
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
		p.SetState(200)
		p.Mnemonic()
	}
	{
		p.SetState(201)
		p.Match(AsmE8ParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(202)
		p.Reg16_indexed()
	}
	{
		p.SetState(203)
		p.Match(AsmE8ParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(204)
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
		p.SetState(206)
		p.Mnemonic()
	}
	{
		p.SetState(207)
		p.Match(AsmE8ParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(208)
		p.Indirect()
	}
	{
		p.SetState(209)
		p.Match(AsmE8ParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(210)
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
		p.SetState(212)
		p.Mnemonic()
	}
	{
		p.SetState(213)
		p.Match(AsmE8ParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(214)
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
		p.SetState(216)
		p.Mnemonic()
	}
	{
		p.SetState(217)
		p.Match(AsmE8ParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(218)
		p.Reg8()
	}
	{
		p.SetState(219)
		p.Match(AsmE8ParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(220)
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
		p.SetState(222)
		p.Mnemonic()
	}
	{
		p.SetState(223)
		p.Match(AsmE8ParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(224)
		p.Reg8()
	}
	{
		p.SetState(225)
		p.Match(AsmE8ParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(226)
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
		p.SetState(228)
		p.Mnemonic()
	}
	{
		p.SetState(229)
		p.Match(AsmE8ParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(230)
		p.Reg8()
	}
	{
		p.SetState(231)
		p.Match(AsmE8ParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(232)
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
		p.SetState(234)
		p.Mnemonic()
	}
	{
		p.SetState(235)
		p.Match(AsmE8ParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(236)
		p.Reg8()
	}
	{
		p.SetState(237)
		p.Match(AsmE8ParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(238)
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
		p.SetState(240)
		p.Mnemonic()
	}
	{
		p.SetState(241)
		p.Match(AsmE8ParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(242)
		p.Reg8()
	}
	{
		p.SetState(243)
		p.Match(AsmE8ParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(244)
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
		p.SetState(246)
		p.Mnemonic()
	}
	{
		p.SetState(247)
		p.Match(AsmE8ParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(248)
		p.Reg8()
	}
	{
		p.SetState(249)
		p.Match(AsmE8ParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(250)
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
		p.SetState(252)
		p.Mnemonic()
	}
	{
		p.SetState(253)
		p.Match(AsmE8ParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(254)
		p.Reg8()
	}
	{
		p.SetState(255)
		p.Match(AsmE8ParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(256)
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
		p.SetState(258)
		p.Mnemonic()
	}
	{
		p.SetState(259)
		p.Match(AsmE8ParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(260)
		p.Reg8()
	}
	{
		p.SetState(261)
		p.Match(AsmE8ParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(262)
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
		p.SetState(264)
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
		p.SetState(266)
		p.Mnemonic()
	}
	{
		p.SetState(267)
		p.Match(AsmE8ParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(268)
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
		p.SetState(270)
		p.Mnemonic()
	}
	{
		p.SetState(271)
		p.Match(AsmE8ParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(272)
		p.Reg16()
	}
	{
		p.SetState(273)
		p.Match(AsmE8ParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(274)
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
		p.SetState(276)
		p.Mnemonic()
	}
	{
		p.SetState(277)
		p.Match(AsmE8ParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(278)
		p.Reg16()
	}
	{
		p.SetState(279)
		p.Match(AsmE8ParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(280)
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
		p.SetState(282)
		p.Mnemonic()
	}
	{
		p.SetState(283)
		p.Match(AsmE8ParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(284)
		p.Reg16()
	}
	{
		p.SetState(285)
		p.Match(AsmE8ParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(286)
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
		p.SetState(288)
		p.Mnemonic()
	}
	{
		p.SetState(289)
		p.Match(AsmE8ParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(290)
		p.Reg16()
	}
	{
		p.SetState(291)
		p.Match(AsmE8ParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(292)
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
		p.SetState(294)
		p.Mnemonic()
	}
	{
		p.SetState(295)
		p.Match(AsmE8ParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(296)
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

// IAddr_mode_implied_status_registerContext is an interface to support dynamic dispatch.
type IAddr_mode_implied_status_registerContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Mnemonic() IMnemonicContext
	Status_register() IStatus_registerContext

	// IsAddr_mode_implied_status_registerContext differentiates from other interfaces.
	IsAddr_mode_implied_status_registerContext()
}

type Addr_mode_implied_status_registerContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAddr_mode_implied_status_registerContext() *Addr_mode_implied_status_registerContext {
	var p = new(Addr_mode_implied_status_registerContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsmE8ParserRULE_addr_mode_implied_status_register
	return p
}

func InitEmptyAddr_mode_implied_status_registerContext(p *Addr_mode_implied_status_registerContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsmE8ParserRULE_addr_mode_implied_status_register
}

func (*Addr_mode_implied_status_registerContext) IsAddr_mode_implied_status_registerContext() {}

func NewAddr_mode_implied_status_registerContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Addr_mode_implied_status_registerContext {
	var p = new(Addr_mode_implied_status_registerContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AsmE8ParserRULE_addr_mode_implied_status_register

	return p
}

func (s *Addr_mode_implied_status_registerContext) GetParser() antlr.Parser { return s.parser }

func (s *Addr_mode_implied_status_registerContext) Mnemonic() IMnemonicContext {
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

func (s *Addr_mode_implied_status_registerContext) Status_register() IStatus_registerContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatus_registerContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatus_registerContext)
}

func (s *Addr_mode_implied_status_registerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Addr_mode_implied_status_registerContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Addr_mode_implied_status_registerContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsmE8Listener); ok {
		listenerT.EnterAddr_mode_implied_status_register(s)
	}
}

func (s *Addr_mode_implied_status_registerContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsmE8Listener); ok {
		listenerT.ExitAddr_mode_implied_status_register(s)
	}
}

func (p *AsmE8Parser) Addr_mode_implied_status_register() (localctx IAddr_mode_implied_status_registerContext) {
	localctx = NewAddr_mode_implied_status_registerContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, AsmE8ParserRULE_addr_mode_implied_status_register)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(298)
		p.Mnemonic()
	}
	{
		p.SetState(299)
		p.Match(AsmE8ParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(300)
		p.Status_register()
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
	p.EnterRule(localctx, 58, AsmE8ParserRULE_addr_mode_stack_imm8)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(302)
		p.Mnemonic()
	}
	{
		p.SetState(303)
		p.Match(AsmE8ParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(304)
		p.Stack()
	}
	{
		p.SetState(305)
		p.Match(AsmE8ParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(306)
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
	p.EnterRule(localctx, 60, AsmE8ParserRULE_addr_mode_stack_imm8_tag)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(308)
		p.Mnemonic()
	}
	{
		p.SetState(309)
		p.Match(AsmE8ParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(310)
		p.Stack()
	}
	{
		p.SetState(311)
		p.Match(AsmE8ParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(312)
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
	p.EnterRule(localctx, 62, AsmE8ParserRULE_addr_mode_stack_reg16)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(314)
		p.Mnemonic()
	}
	{
		p.SetState(315)
		p.Match(AsmE8ParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(316)
		p.Stack()
	}
	{
		p.SetState(317)
		p.Match(AsmE8ParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(318)
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
	p.EnterRule(localctx, 64, AsmE8ParserRULE_mnemonic)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(320)
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
	p.EnterRule(localctx, 66, AsmE8ParserRULE_reg8)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(322)
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
	p.EnterRule(localctx, 68, AsmE8ParserRULE_reg16)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(324)
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
	p.EnterRule(localctx, 70, AsmE8ParserRULE_stack)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(326)
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

// IStatus_registerContext is an interface to support dynamic dispatch.
type IStatus_registerContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsStatus_registerContext differentiates from other interfaces.
	IsStatus_registerContext()
}

type Status_registerContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatus_registerContext() *Status_registerContext {
	var p = new(Status_registerContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsmE8ParserRULE_status_register
	return p
}

func InitEmptyStatus_registerContext(p *Status_registerContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsmE8ParserRULE_status_register
}

func (*Status_registerContext) IsStatus_registerContext() {}

func NewStatus_registerContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Status_registerContext {
	var p = new(Status_registerContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AsmE8ParserRULE_status_register

	return p
}

func (s *Status_registerContext) GetParser() antlr.Parser { return s.parser }
func (s *Status_registerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Status_registerContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Status_registerContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsmE8Listener); ok {
		listenerT.EnterStatus_register(s)
	}
}

func (s *Status_registerContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsmE8Listener); ok {
		listenerT.ExitStatus_register(s)
	}
}

func (p *AsmE8Parser) Status_register() (localctx IStatus_registerContext) {
	localctx = NewStatus_registerContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, AsmE8ParserRULE_status_register)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(328)
		p.Match(AsmE8ParserT__50)
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
	p.EnterRule(localctx, 74, AsmE8ParserRULE_direct)
	p.SetState(339)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 7, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(330)
			p.Match(AsmE8ParserT__51)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(331)
			p.Imm()
		}
		{
			p.SetState(332)
			p.Match(AsmE8ParserT__52)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(334)
			p.Match(AsmE8ParserT__51)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(335)
			p.Tag()
		}
		{
			p.SetState(336)
			p.Match(AsmE8ParserT__52)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(338)
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
	p.EnterRule(localctx, 76, AsmE8ParserRULE_direct_virtual_offset)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(341)
		p.Match(AsmE8ParserT__51)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(342)
		p.Match(AsmE8ParserSTR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(343)
		_la = p.GetTokenStream().LA(1)

		if !(_la == AsmE8ParserT__53 || _la == AsmE8ParserT__54) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(344)
		p.Match(AsmE8ParserINT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(345)
		p.Match(AsmE8ParserT__52)
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
	p.EnterRule(localctx, 78, AsmE8ParserRULE_direct_offset)
	p.SetState(359)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(347)
			p.Match(AsmE8ParserT__51)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(348)
			p.Imm()
		}
		{
			p.SetState(349)
			p.Match(AsmE8ParserT__53)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(350)
			p.Reg8()
		}
		{
			p.SetState(351)
			p.Match(AsmE8ParserT__52)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(353)
			p.Match(AsmE8ParserT__51)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(354)
			p.Tag()
		}
		{
			p.SetState(355)
			p.Match(AsmE8ParserT__53)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(356)
			p.Reg8()
		}
		{
			p.SetState(357)
			p.Match(AsmE8ParserT__52)
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
	p.EnterRule(localctx, 80, AsmE8ParserRULE_indirect_offset)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(361)
		p.Match(AsmE8ParserT__51)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(362)
		p.Reg16()
	}
	{
		p.SetState(363)
		p.Match(AsmE8ParserT__53)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(364)
		p.Imm()
	}
	{
		p.SetState(365)
		p.Match(AsmE8ParserT__52)
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
	p.EnterRule(localctx, 82, AsmE8ParserRULE_reg16_indexed)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(367)
		p.Match(AsmE8ParserT__51)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(368)
		p.Reg16()
	}
	{
		p.SetState(369)
		p.Match(AsmE8ParserT__53)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(370)
		p.Reg8()
	}
	{
		p.SetState(371)
		p.Match(AsmE8ParserT__52)
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
	p.EnterRule(localctx, 84, AsmE8ParserRULE_indirect)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(373)
		p.Match(AsmE8ParserT__51)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(374)
		p.Reg16()
	}
	{
		p.SetState(375)
		p.Match(AsmE8ParserT__52)
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
	p.EnterRule(localctx, 86, AsmE8ParserRULE_variable)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(377)
		p.Tag()
	}
	{
		p.SetState(378)
		p.Match(AsmE8ParserT__55)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(379)
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
	p.EnterRule(localctx, 88, AsmE8ParserRULE_variable_reference)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(381)
		p.Tag()
	}
	{
		p.SetState(382)
		p.Match(AsmE8ParserT__55)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(383)
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
	p.EnterRule(localctx, 90, AsmE8ParserRULE_reference)
	var _la int

	p.SetState(389)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(385)
			p.Match(AsmE8ParserSTR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(386)
			p.Match(AsmE8ParserSTR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(387)
			_la = p.GetTokenStream().LA(1)

			if !(_la == AsmE8ParserT__53 || _la == AsmE8ParserT__54) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(388)
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
	p.EnterRule(localctx, 92, AsmE8ParserRULE_directives)
	p.SetState(403)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case AsmE8ParserT__56:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(391)
			p.Match(AsmE8ParserT__56)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(392)
			p.Ascii()
		}

	case AsmE8ParserT__57:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(393)
			p.Match(AsmE8ParserT__57)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(394)
			p.Ascii()
		}

	case AsmE8ParserT__58:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(395)
			p.Match(AsmE8ParserT__58)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(396)
			p.Imm_list()
		}

	case AsmE8ParserT__59:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(397)
			p.Match(AsmE8ParserT__59)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(398)
			p.Imm_list()
		}

	case AsmE8ParserT__60:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(399)
			p.Match(AsmE8ParserT__60)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(400)
			p.Imm_list()
		}

	case AsmE8ParserT__61:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(401)
			p.Match(AsmE8ParserT__61)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(402)
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
	p.EnterRule(localctx, 94, AsmE8ParserRULE_segment)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(405)
		p.Match(AsmE8ParserT__62)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(406)
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
	p.EnterRule(localctx, 96, AsmE8ParserRULE_access)
	p.SetState(414)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(408)
			p.Match(AsmE8ParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(409)
			p.Match(AsmE8ParserT__63)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(410)
			p.Tag()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(411)
			p.Match(AsmE8ParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(412)
			p.Match(AsmE8ParserT__64)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(413)
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
	p.EnterRule(localctx, 98, AsmE8ParserRULE_imm_list)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(418)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case AsmE8ParserBINARY, AsmE8ParserINT, AsmE8ParserCHAR, AsmE8ParserHEX:
		{
			p.SetState(416)
			p.Imm()
		}

	case AsmE8ParserSTR:
		{
			p.SetState(417)
			p.Tag()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.SetState(427)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == AsmE8ParserT__6 {
		{
			p.SetState(420)
			p.Match(AsmE8ParserT__6)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(423)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case AsmE8ParserBINARY, AsmE8ParserINT, AsmE8ParserCHAR, AsmE8ParserHEX:
			{
				p.SetState(421)
				p.Imm()
			}

		case AsmE8ParserSTR:
			{
				p.SetState(422)
				p.Tag()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(429)
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
	p.EnterRule(localctx, 100, AsmE8ParserRULE_imm)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(430)
		_la = p.GetTokenStream().LA(1)

		if !((int64((_la-69)) & ^0x3f) == 0 && ((int64(1)<<(_la-69))&23) != 0) {
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
	p.EnterRule(localctx, 102, AsmE8ParserRULE_ascii)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(432)
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
	p.EnterRule(localctx, 104, AsmE8ParserRULE_tag)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(434)
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
