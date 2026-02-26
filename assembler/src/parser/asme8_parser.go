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
		"'add'", "'and'", "'htl'", "'call'", "'clc'", "'int'", "'sti'", "'cli'",
		"'cmp'", "'dec'", "'inc'", "'jc'", "'jmp'", "'jnc'", "'jns'", "'jnz'",
		"'js'", "'jz'", "'mov'", "'nop'", "'or'", "'pop'", "'push'", "'ret'",
		"'rol'", "'ror'", "'rti'", "'sbb'", "'shl'", "'shr'", "'sub'", "'xor'",
		"'hlt'", "'a'", "'b'", "'c'", "'d'", "'e'", "'ipl'", "'iph'", "'bpl'",
		"'bph'", "'ip'", "'bp'", "'sp'", "'sr'", "'['", "']'", "'+'", "'-'",
		"'='", "'.asciiz '", "'.ascii '", "'.byte '", "'.word '", "'.resb '",
		"'.org '", "'.segment '", "'global '", "'extern '",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "LINE_COMMENT", "SUFFIX_WHITESPACE", "WHITE_SPACE", "BINARY", "INT",
		"CHAR", "ANY_STR", "HEX", "STR",
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
		4, 1, 77, 438, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
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
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 142, 8, 2, 1, 3, 1, 3, 1, 3,
		1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4,
		1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4,
		1, 4, 1, 4, 1, 4, 3, 4, 174, 8, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6,
		1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8,
		1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1,
		10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12,
		1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1,
		14, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 16,
		1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1,
		17, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 19, 1, 19, 1, 19, 1, 19,
		1, 19, 1, 19, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 21, 1, 21, 1,
		22, 1, 22, 1, 22, 1, 22, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 24,
		1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1,
		25, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 27, 1, 27, 1, 27, 1, 27,
		1, 28, 1, 28, 1, 28, 1, 28, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1,
		30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31,
		1, 31, 1, 32, 1, 32, 1, 33, 1, 33, 1, 34, 1, 34, 1, 35, 1, 35, 1, 36, 1,
		36, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 3, 37,
		341, 8, 37, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 39, 1, 39, 1,
		39, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 3, 39,
		361, 8, 39, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 41, 1, 41, 1,
		41, 1, 41, 1, 41, 1, 41, 1, 42, 1, 42, 1, 42, 1, 42, 1, 43, 1, 43, 1, 43,
		1, 43, 1, 44, 1, 44, 1, 44, 1, 44, 1, 45, 1, 45, 1, 45, 1, 45, 3, 45, 391,
		8, 45, 1, 46, 1, 46, 1, 46, 1, 46, 1, 46, 1, 46, 1, 46, 1, 46, 1, 46, 1,
		46, 1, 46, 1, 46, 3, 46, 405, 8, 46, 1, 47, 1, 47, 1, 47, 1, 48, 1, 48,
		1, 48, 1, 48, 1, 48, 1, 48, 3, 48, 416, 8, 48, 1, 49, 1, 49, 3, 49, 420,
		8, 49, 1, 49, 1, 49, 1, 49, 3, 49, 425, 8, 49, 5, 49, 427, 8, 49, 10, 49,
		12, 49, 430, 9, 49, 1, 50, 1, 50, 1, 51, 1, 51, 1, 52, 1, 52, 1, 52, 0,
		0, 53, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34,
		36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 70,
		72, 74, 76, 78, 80, 82, 84, 86, 88, 90, 92, 94, 96, 98, 100, 102, 104,
		0, 6, 1, 0, 1, 3, 1, 0, 8, 41, 1, 0, 42, 50, 1, 0, 51, 52, 1, 0, 57, 58,
		2, 0, 72, 74, 76, 76, 437, 0, 106, 1, 0, 0, 0, 2, 125, 1, 0, 0, 0, 4, 141,
		1, 0, 0, 0, 6, 143, 1, 0, 0, 0, 8, 173, 1, 0, 0, 0, 10, 175, 1, 0, 0, 0,
		12, 179, 1, 0, 0, 0, 14, 183, 1, 0, 0, 0, 16, 189, 1, 0, 0, 0, 18, 195,
		1, 0, 0, 0, 20, 201, 1, 0, 0, 0, 22, 207, 1, 0, 0, 0, 24, 213, 1, 0, 0,
		0, 26, 217, 1, 0, 0, 0, 28, 223, 1, 0, 0, 0, 30, 229, 1, 0, 0, 0, 32, 235,
		1, 0, 0, 0, 34, 241, 1, 0, 0, 0, 36, 247, 1, 0, 0, 0, 38, 253, 1, 0, 0,
		0, 40, 259, 1, 0, 0, 0, 42, 265, 1, 0, 0, 0, 44, 267, 1, 0, 0, 0, 46, 271,
		1, 0, 0, 0, 48, 277, 1, 0, 0, 0, 50, 283, 1, 0, 0, 0, 52, 289, 1, 0, 0,
		0, 54, 295, 1, 0, 0, 0, 56, 299, 1, 0, 0, 0, 58, 303, 1, 0, 0, 0, 60, 309,
		1, 0, 0, 0, 62, 315, 1, 0, 0, 0, 64, 321, 1, 0, 0, 0, 66, 323, 1, 0, 0,
		0, 68, 325, 1, 0, 0, 0, 70, 327, 1, 0, 0, 0, 72, 329, 1, 0, 0, 0, 74, 340,
		1, 0, 0, 0, 76, 342, 1, 0, 0, 0, 78, 360, 1, 0, 0, 0, 80, 362, 1, 0, 0,
		0, 82, 368, 1, 0, 0, 0, 84, 374, 1, 0, 0, 0, 86, 378, 1, 0, 0, 0, 88, 382,
		1, 0, 0, 0, 90, 390, 1, 0, 0, 0, 92, 404, 1, 0, 0, 0, 94, 406, 1, 0, 0,
		0, 96, 415, 1, 0, 0, 0, 98, 419, 1, 0, 0, 0, 100, 431, 1, 0, 0, 0, 102,
		433, 1, 0, 0, 0, 104, 435, 1, 0, 0, 0, 106, 107, 3, 2, 1, 0, 107, 108,
		5, 0, 0, 1, 108, 1, 1, 0, 0, 0, 109, 111, 3, 4, 2, 0, 110, 109, 1, 0, 0,
		0, 111, 112, 1, 0, 0, 0, 112, 110, 1, 0, 0, 0, 112, 113, 1, 0, 0, 0, 113,
		115, 1, 0, 0, 0, 114, 116, 5, 69, 0, 0, 115, 114, 1, 0, 0, 0, 115, 116,
		1, 0, 0, 0, 116, 126, 1, 0, 0, 0, 117, 119, 3, 4, 2, 0, 118, 117, 1, 0,
		0, 0, 119, 120, 1, 0, 0, 0, 120, 118, 1, 0, 0, 0, 120, 121, 1, 0, 0, 0,
		121, 123, 1, 0, 0, 0, 122, 124, 5, 70, 0, 0, 123, 122, 1, 0, 0, 0, 123,
		124, 1, 0, 0, 0, 124, 126, 1, 0, 0, 0, 125, 110, 1, 0, 0, 0, 125, 118,
		1, 0, 0, 0, 126, 3, 1, 0, 0, 0, 127, 142, 3, 88, 44, 0, 128, 142, 3, 86,
		43, 0, 129, 130, 7, 0, 0, 0, 130, 142, 3, 8, 4, 0, 131, 142, 3, 8, 4, 0,
		132, 142, 3, 94, 47, 0, 133, 142, 3, 96, 48, 0, 134, 135, 3, 6, 3, 0, 135,
		136, 5, 4, 0, 0, 136, 137, 3, 92, 46, 0, 137, 142, 1, 0, 0, 0, 138, 142,
		3, 92, 46, 0, 139, 142, 3, 6, 3, 0, 140, 142, 5, 5, 0, 0, 141, 127, 1,
		0, 0, 0, 141, 128, 1, 0, 0, 0, 141, 129, 1, 0, 0, 0, 141, 131, 1, 0, 0,
		0, 141, 132, 1, 0, 0, 0, 141, 133, 1, 0, 0, 0, 141, 134, 1, 0, 0, 0, 141,
		138, 1, 0, 0, 0, 141, 139, 1, 0, 0, 0, 141, 140, 1, 0, 0, 0, 142, 5, 1,
		0, 0, 0, 143, 144, 3, 104, 52, 0, 144, 145, 5, 6, 0, 0, 145, 7, 1, 0, 0,
		0, 146, 174, 3, 10, 5, 0, 147, 174, 3, 12, 6, 0, 148, 174, 3, 14, 7, 0,
		149, 174, 3, 16, 8, 0, 150, 174, 3, 18, 9, 0, 151, 174, 3, 20, 10, 0, 152,
		174, 3, 22, 11, 0, 153, 174, 3, 24, 12, 0, 154, 174, 3, 26, 13, 0, 155,
		174, 3, 28, 14, 0, 156, 174, 3, 30, 15, 0, 157, 174, 3, 32, 16, 0, 158,
		174, 3, 34, 17, 0, 159, 174, 3, 36, 18, 0, 160, 174, 3, 38, 19, 0, 161,
		174, 3, 40, 20, 0, 162, 174, 3, 42, 21, 0, 163, 174, 3, 44, 22, 0, 164,
		174, 3, 46, 23, 0, 165, 174, 3, 48, 24, 0, 166, 174, 3, 50, 25, 0, 167,
		174, 3, 52, 26, 0, 168, 174, 3, 54, 27, 0, 169, 174, 3, 56, 28, 0, 170,
		174, 3, 58, 29, 0, 171, 174, 3, 60, 30, 0, 172, 174, 3, 62, 31, 0, 173,
		146, 1, 0, 0, 0, 173, 147, 1, 0, 0, 0, 173, 148, 1, 0, 0, 0, 173, 149,
		1, 0, 0, 0, 173, 150, 1, 0, 0, 0, 173, 151, 1, 0, 0, 0, 173, 152, 1, 0,
		0, 0, 173, 153, 1, 0, 0, 0, 173, 154, 1, 0, 0, 0, 173, 155, 1, 0, 0, 0,
		173, 156, 1, 0, 0, 0, 173, 157, 1, 0, 0, 0, 173, 158, 1, 0, 0, 0, 173,
		159, 1, 0, 0, 0, 173, 160, 1, 0, 0, 0, 173, 161, 1, 0, 0, 0, 173, 162,
		1, 0, 0, 0, 173, 163, 1, 0, 0, 0, 173, 164, 1, 0, 0, 0, 173, 165, 1, 0,
		0, 0, 173, 166, 1, 0, 0, 0, 173, 167, 1, 0, 0, 0, 173, 168, 1, 0, 0, 0,
		173, 169, 1, 0, 0, 0, 173, 170, 1, 0, 0, 0, 173, 171, 1, 0, 0, 0, 173,
		172, 1, 0, 0, 0, 174, 9, 1, 0, 0, 0, 175, 176, 3, 64, 32, 0, 176, 177,
		5, 4, 0, 0, 177, 178, 3, 100, 50, 0, 178, 11, 1, 0, 0, 0, 179, 180, 3,
		64, 32, 0, 180, 181, 5, 4, 0, 0, 181, 182, 3, 104, 52, 0, 182, 13, 1, 0,
		0, 0, 183, 184, 3, 64, 32, 0, 184, 185, 5, 4, 0, 0, 185, 186, 3, 78, 39,
		0, 186, 187, 5, 7, 0, 0, 187, 188, 3, 66, 33, 0, 188, 15, 1, 0, 0, 0, 189,
		190, 3, 64, 32, 0, 190, 191, 5, 4, 0, 0, 191, 192, 3, 74, 37, 0, 192, 193,
		5, 7, 0, 0, 193, 194, 3, 66, 33, 0, 194, 17, 1, 0, 0, 0, 195, 196, 3, 64,
		32, 0, 196, 197, 5, 4, 0, 0, 197, 198, 3, 80, 40, 0, 198, 199, 5, 7, 0,
		0, 199, 200, 3, 66, 33, 0, 200, 19, 1, 0, 0, 0, 201, 202, 3, 64, 32, 0,
		202, 203, 5, 4, 0, 0, 203, 204, 3, 82, 41, 0, 204, 205, 5, 7, 0, 0, 205,
		206, 3, 66, 33, 0, 206, 21, 1, 0, 0, 0, 207, 208, 3, 64, 32, 0, 208, 209,
		5, 4, 0, 0, 209, 210, 3, 84, 42, 0, 210, 211, 5, 7, 0, 0, 211, 212, 3,
		66, 33, 0, 212, 23, 1, 0, 0, 0, 213, 214, 3, 64, 32, 0, 214, 215, 5, 4,
		0, 0, 215, 216, 3, 66, 33, 0, 216, 25, 1, 0, 0, 0, 217, 218, 3, 64, 32,
		0, 218, 219, 5, 4, 0, 0, 219, 220, 3, 66, 33, 0, 220, 221, 5, 7, 0, 0,
		221, 222, 3, 100, 50, 0, 222, 27, 1, 0, 0, 0, 223, 224, 3, 64, 32, 0, 224,
		225, 5, 4, 0, 0, 225, 226, 3, 66, 33, 0, 226, 227, 5, 7, 0, 0, 227, 228,
		3, 104, 52, 0, 228, 29, 1, 0, 0, 0, 229, 230, 3, 64, 32, 0, 230, 231, 5,
		4, 0, 0, 231, 232, 3, 66, 33, 0, 232, 233, 5, 7, 0, 0, 233, 234, 3, 78,
		39, 0, 234, 31, 1, 0, 0, 0, 235, 236, 3, 64, 32, 0, 236, 237, 5, 4, 0,
		0, 237, 238, 3, 66, 33, 0, 238, 239, 5, 7, 0, 0, 239, 240, 3, 74, 37, 0,
		240, 33, 1, 0, 0, 0, 241, 242, 3, 64, 32, 0, 242, 243, 5, 4, 0, 0, 243,
		244, 3, 66, 33, 0, 244, 245, 5, 7, 0, 0, 245, 246, 3, 80, 40, 0, 246, 35,
		1, 0, 0, 0, 247, 248, 3, 64, 32, 0, 248, 249, 5, 4, 0, 0, 249, 250, 3,
		66, 33, 0, 250, 251, 5, 7, 0, 0, 251, 252, 3, 82, 41, 0, 252, 37, 1, 0,
		0, 0, 253, 254, 3, 64, 32, 0, 254, 255, 5, 4, 0, 0, 255, 256, 3, 66, 33,
		0, 256, 257, 5, 7, 0, 0, 257, 258, 3, 84, 42, 0, 258, 39, 1, 0, 0, 0, 259,
		260, 3, 64, 32, 0, 260, 261, 5, 4, 0, 0, 261, 262, 3, 66, 33, 0, 262, 263,
		5, 7, 0, 0, 263, 264, 3, 66, 33, 0, 264, 41, 1, 0, 0, 0, 265, 266, 3, 64,
		32, 0, 266, 43, 1, 0, 0, 0, 267, 268, 3, 64, 32, 0, 268, 269, 5, 4, 0,
		0, 269, 270, 3, 68, 34, 0, 270, 45, 1, 0, 0, 0, 271, 272, 3, 64, 32, 0,
		272, 273, 5, 4, 0, 0, 273, 274, 3, 68, 34, 0, 274, 275, 5, 7, 0, 0, 275,
		276, 3, 100, 50, 0, 276, 47, 1, 0, 0, 0, 277, 278, 3, 64, 32, 0, 278, 279,
		5, 4, 0, 0, 279, 280, 3, 68, 34, 0, 280, 281, 5, 7, 0, 0, 281, 282, 3,
		104, 52, 0, 282, 49, 1, 0, 0, 0, 283, 284, 3, 64, 32, 0, 284, 285, 5, 4,
		0, 0, 285, 286, 3, 68, 34, 0, 286, 287, 5, 7, 0, 0, 287, 288, 3, 68, 34,
		0, 288, 51, 1, 0, 0, 0, 289, 290, 3, 64, 32, 0, 290, 291, 5, 4, 0, 0, 291,
		292, 3, 68, 34, 0, 292, 293, 5, 7, 0, 0, 293, 294, 3, 70, 35, 0, 294, 53,
		1, 0, 0, 0, 295, 296, 3, 64, 32, 0, 296, 297, 5, 4, 0, 0, 297, 298, 3,
		70, 35, 0, 298, 55, 1, 0, 0, 0, 299, 300, 3, 64, 32, 0, 300, 301, 5, 4,
		0, 0, 301, 302, 3, 72, 36, 0, 302, 57, 1, 0, 0, 0, 303, 304, 3, 64, 32,
		0, 304, 305, 5, 4, 0, 0, 305, 306, 3, 70, 35, 0, 306, 307, 5, 7, 0, 0,
		307, 308, 3, 100, 50, 0, 308, 59, 1, 0, 0, 0, 309, 310, 3, 64, 32, 0, 310,
		311, 5, 4, 0, 0, 311, 312, 3, 70, 35, 0, 312, 313, 5, 7, 0, 0, 313, 314,
		3, 104, 52, 0, 314, 61, 1, 0, 0, 0, 315, 316, 3, 64, 32, 0, 316, 317, 5,
		4, 0, 0, 317, 318, 3, 70, 35, 0, 318, 319, 5, 7, 0, 0, 319, 320, 3, 68,
		34, 0, 320, 63, 1, 0, 0, 0, 321, 322, 7, 1, 0, 0, 322, 65, 1, 0, 0, 0,
		323, 324, 7, 2, 0, 0, 324, 67, 1, 0, 0, 0, 325, 326, 7, 3, 0, 0, 326, 69,
		1, 0, 0, 0, 327, 328, 5, 53, 0, 0, 328, 71, 1, 0, 0, 0, 329, 330, 5, 54,
		0, 0, 330, 73, 1, 0, 0, 0, 331, 332, 5, 55, 0, 0, 332, 333, 3, 100, 50,
		0, 333, 334, 5, 56, 0, 0, 334, 341, 1, 0, 0, 0, 335, 336, 5, 55, 0, 0,
		336, 337, 3, 104, 52, 0, 337, 338, 5, 56, 0, 0, 338, 341, 1, 0, 0, 0, 339,
		341, 3, 76, 38, 0, 340, 331, 1, 0, 0, 0, 340, 335, 1, 0, 0, 0, 340, 339,
		1, 0, 0, 0, 341, 75, 1, 0, 0, 0, 342, 343, 5, 55, 0, 0, 343, 344, 5, 77,
		0, 0, 344, 345, 7, 4, 0, 0, 345, 346, 5, 73, 0, 0, 346, 347, 5, 56, 0,
		0, 347, 77, 1, 0, 0, 0, 348, 349, 5, 55, 0, 0, 349, 350, 3, 100, 50, 0,
		350, 351, 5, 57, 0, 0, 351, 352, 3, 66, 33, 0, 352, 353, 5, 56, 0, 0, 353,
		361, 1, 0, 0, 0, 354, 355, 5, 55, 0, 0, 355, 356, 3, 104, 52, 0, 356, 357,
		5, 57, 0, 0, 357, 358, 3, 66, 33, 0, 358, 359, 5, 56, 0, 0, 359, 361, 1,
		0, 0, 0, 360, 348, 1, 0, 0, 0, 360, 354, 1, 0, 0, 0, 361, 79, 1, 0, 0,
		0, 362, 363, 5, 55, 0, 0, 363, 364, 3, 68, 34, 0, 364, 365, 5, 57, 0, 0,
		365, 366, 3, 100, 50, 0, 366, 367, 5, 56, 0, 0, 367, 81, 1, 0, 0, 0, 368,
		369, 5, 55, 0, 0, 369, 370, 3, 68, 34, 0, 370, 371, 5, 57, 0, 0, 371, 372,
		3, 66, 33, 0, 372, 373, 5, 56, 0, 0, 373, 83, 1, 0, 0, 0, 374, 375, 5,
		55, 0, 0, 375, 376, 3, 68, 34, 0, 376, 377, 5, 56, 0, 0, 377, 85, 1, 0,
		0, 0, 378, 379, 3, 104, 52, 0, 379, 380, 5, 59, 0, 0, 380, 381, 3, 100,
		50, 0, 381, 87, 1, 0, 0, 0, 382, 383, 3, 104, 52, 0, 383, 384, 5, 59, 0,
		0, 384, 385, 3, 90, 45, 0, 385, 89, 1, 0, 0, 0, 386, 391, 5, 77, 0, 0,
		387, 388, 5, 77, 0, 0, 388, 389, 7, 4, 0, 0, 389, 391, 3, 100, 50, 0, 390,
		386, 1, 0, 0, 0, 390, 387, 1, 0, 0, 0, 391, 91, 1, 0, 0, 0, 392, 393, 5,
		60, 0, 0, 393, 405, 3, 102, 51, 0, 394, 395, 5, 61, 0, 0, 395, 405, 3,
		102, 51, 0, 396, 397, 5, 62, 0, 0, 397, 405, 3, 98, 49, 0, 398, 399, 5,
		63, 0, 0, 399, 405, 3, 98, 49, 0, 400, 401, 5, 64, 0, 0, 401, 405, 3, 98,
		49, 0, 402, 403, 5, 65, 0, 0, 403, 405, 3, 98, 49, 0, 404, 392, 1, 0, 0,
		0, 404, 394, 1, 0, 0, 0, 404, 396, 1, 0, 0, 0, 404, 398, 1, 0, 0, 0, 404,
		400, 1, 0, 0, 0, 404, 402, 1, 0, 0, 0, 405, 93, 1, 0, 0, 0, 406, 407, 5,
		66, 0, 0, 407, 408, 5, 75, 0, 0, 408, 95, 1, 0, 0, 0, 409, 410, 5, 2, 0,
		0, 410, 411, 5, 67, 0, 0, 411, 416, 3, 104, 52, 0, 412, 413, 5, 2, 0, 0,
		413, 414, 5, 68, 0, 0, 414, 416, 3, 104, 52, 0, 415, 409, 1, 0, 0, 0, 415,
		412, 1, 0, 0, 0, 416, 97, 1, 0, 0, 0, 417, 420, 3, 100, 50, 0, 418, 420,
		3, 104, 52, 0, 419, 417, 1, 0, 0, 0, 419, 418, 1, 0, 0, 0, 420, 428, 1,
		0, 0, 0, 421, 424, 5, 7, 0, 0, 422, 425, 3, 100, 50, 0, 423, 425, 3, 104,
		52, 0, 424, 422, 1, 0, 0, 0, 424, 423, 1, 0, 0, 0, 425, 427, 1, 0, 0, 0,
		426, 421, 1, 0, 0, 0, 427, 430, 1, 0, 0, 0, 428, 426, 1, 0, 0, 0, 428,
		429, 1, 0, 0, 0, 429, 99, 1, 0, 0, 0, 430, 428, 1, 0, 0, 0, 431, 432, 7,
		5, 0, 0, 432, 101, 1, 0, 0, 0, 433, 434, 5, 75, 0, 0, 434, 103, 1, 0, 0,
		0, 435, 436, 5, 77, 0, 0, 436, 105, 1, 0, 0, 0, 15, 112, 115, 120, 123,
		125, 141, 173, 340, 360, 390, 404, 415, 419, 424, 428,
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
	AsmE8ParserT__65             = 66
	AsmE8ParserT__66             = 67
	AsmE8ParserT__67             = 68
	AsmE8ParserLINE_COMMENT      = 69
	AsmE8ParserSUFFIX_WHITESPACE = 70
	AsmE8ParserWHITE_SPACE       = 71
	AsmE8ParserBINARY            = 72
	AsmE8ParserINT               = 73
	AsmE8ParserCHAR              = 74
	AsmE8ParserANY_STR           = 75
	AsmE8ParserHEX               = 76
	AsmE8ParserSTR               = 77
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

		for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&-1152917106560336082) != 0) || ((int64((_la-64)) & ^0x3f) == 0 && ((int64(1)<<(_la-64))&8199) != 0) {
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

		for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&-1152917106560336082) != 0) || ((int64((_la-64)) & ^0x3f) == 0 && ((int64(1)<<(_la-64))&8199) != 0) {
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

	p.SetState(141)
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
			p.Inst()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(132)
			p.Segment()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(133)
			p.Access()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(134)
			p.Label()
		}
		{
			p.SetState(135)
			p.Match(AsmE8ParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(136)
			p.Directives()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(138)
			p.Directives()
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(139)
			p.Label()
		}

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(140)
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
		p.SetState(143)
		p.Tag()
	}
	{
		p.SetState(144)
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
	p.SetState(173)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(146)
			p.Addr_mode_imm16()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(147)
			p.Addr_mode_imm16_tag()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(148)
			p.Addr_mode_reg_to_mem_indexed()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(149)
			p.Addr_mode_reg_to_mem_direct()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(150)
			p.Addr_mode_reg_to_mem_indirect_offset()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(151)
			p.Addr_mode_reg_to_mem_reg16_indexed()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(152)
			p.Addr_mode_reg_to_mem_indirect()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(153)
			p.Addr_mode_implied_reg8()
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(154)
			p.Addr_mode_reg8_imm8()
		}

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(155)
			p.Addr_mode_reg8_imm8_tag()
		}

	case 11:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(156)
			p.Addr_mode_mem_to_reg_indexed()
		}

	case 12:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(157)
			p.Addr_mode_mem_to_reg_direct()
		}

	case 13:
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(158)
			p.Addr_mode_mem_to_reg_indirect_offset()
		}

	case 14:
		p.EnterOuterAlt(localctx, 14)
		{
			p.SetState(159)
			p.Addr_mode_mem_to_reg_reg16_indexed()
		}

	case 15:
		p.EnterOuterAlt(localctx, 15)
		{
			p.SetState(160)
			p.Addr_mode_mem_to_reg_indirect()
		}

	case 16:
		p.EnterOuterAlt(localctx, 16)
		{
			p.SetState(161)
			p.Addr_mode_reg8_reg8()
		}

	case 17:
		p.EnterOuterAlt(localctx, 17)
		{
			p.SetState(162)
			p.Addr_mode_implied()
		}

	case 18:
		p.EnterOuterAlt(localctx, 18)
		{
			p.SetState(163)
			p.Addr_mode_implied_reg16()
		}

	case 19:
		p.EnterOuterAlt(localctx, 19)
		{
			p.SetState(164)
			p.Addr_mode_reg16_imm()
		}

	case 20:
		p.EnterOuterAlt(localctx, 20)
		{
			p.SetState(165)
			p.Addr_mode_reg16_imm_tag()
		}

	case 21:
		p.EnterOuterAlt(localctx, 21)
		{
			p.SetState(166)
			p.Addr_mode_reg16_reg16()
		}

	case 22:
		p.EnterOuterAlt(localctx, 22)
		{
			p.SetState(167)
			p.Addr_mode_reg16_stack()
		}

	case 23:
		p.EnterOuterAlt(localctx, 23)
		{
			p.SetState(168)
			p.Addr_mode_implied_stack()
		}

	case 24:
		p.EnterOuterAlt(localctx, 24)
		{
			p.SetState(169)
			p.Addr_mode_implied_status_register()
		}

	case 25:
		p.EnterOuterAlt(localctx, 25)
		{
			p.SetState(170)
			p.Addr_mode_stack_imm8()
		}

	case 26:
		p.EnterOuterAlt(localctx, 26)
		{
			p.SetState(171)
			p.Addr_mode_stack_imm8_tag()
		}

	case 27:
		p.EnterOuterAlt(localctx, 27)
		{
			p.SetState(172)
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
		p.SetState(175)
		p.Mnemonic()
	}
	{
		p.SetState(176)
		p.Match(AsmE8ParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(177)
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
		p.SetState(179)
		p.Mnemonic()
	}
	{
		p.SetState(180)
		p.Match(AsmE8ParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(181)
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
		p.Direct_offset()
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
		p.Direct()
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
		p.Indirect_offset()
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
		p.Reg16_indexed()
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
		p.Indirect()
	}
	{
		p.SetState(210)
		p.Match(AsmE8ParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(211)
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
		p.SetState(213)
		p.Mnemonic()
	}
	{
		p.SetState(214)
		p.Match(AsmE8ParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(215)
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
		p.SetState(259)
		p.Mnemonic()
	}
	{
		p.SetState(260)
		p.Match(AsmE8ParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(261)
		p.Reg8()
	}
	{
		p.SetState(262)
		p.Match(AsmE8ParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(263)
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
		p.SetState(265)
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
		p.SetState(267)
		p.Mnemonic()
	}
	{
		p.SetState(268)
		p.Match(AsmE8ParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(269)
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
		p.Reg16()
	}
	{
		p.SetState(292)
		p.Match(AsmE8ParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(293)
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
		p.SetState(295)
		p.Mnemonic()
	}
	{
		p.SetState(296)
		p.Match(AsmE8ParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(297)
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
		p.SetState(303)
		p.Mnemonic()
	}
	{
		p.SetState(304)
		p.Match(AsmE8ParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(305)
		p.Stack()
	}
	{
		p.SetState(306)
		p.Match(AsmE8ParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(307)
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
		p.SetState(309)
		p.Mnemonic()
	}
	{
		p.SetState(310)
		p.Match(AsmE8ParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(311)
		p.Stack()
	}
	{
		p.SetState(312)
		p.Match(AsmE8ParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(313)
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
		p.SetState(315)
		p.Mnemonic()
	}
	{
		p.SetState(316)
		p.Match(AsmE8ParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(317)
		p.Stack()
	}
	{
		p.SetState(318)
		p.Match(AsmE8ParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(319)
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
		p.SetState(321)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4398046510848) != 0) {
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
		p.SetState(323)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2247401767174144) != 0) {
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
		p.SetState(325)
		_la = p.GetTokenStream().LA(1)

		if !(_la == AsmE8ParserT__50 || _la == AsmE8ParserT__51) {
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
		p.SetState(327)
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
		p.SetState(329)
		p.Match(AsmE8ParserT__53)
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
	p.SetState(340)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 7, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(331)
			p.Match(AsmE8ParserT__54)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(332)
			p.Imm()
		}
		{
			p.SetState(333)
			p.Match(AsmE8ParserT__55)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(335)
			p.Match(AsmE8ParserT__54)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(336)
			p.Tag()
		}
		{
			p.SetState(337)
			p.Match(AsmE8ParserT__55)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(339)
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
		p.SetState(342)
		p.Match(AsmE8ParserT__54)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(343)
		p.Match(AsmE8ParserSTR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(344)
		_la = p.GetTokenStream().LA(1)

		if !(_la == AsmE8ParserT__56 || _la == AsmE8ParserT__57) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(345)
		p.Match(AsmE8ParserINT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(346)
		p.Match(AsmE8ParserT__55)
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
	p.SetState(360)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(348)
			p.Match(AsmE8ParserT__54)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(349)
			p.Imm()
		}
		{
			p.SetState(350)
			p.Match(AsmE8ParserT__56)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(351)
			p.Reg8()
		}
		{
			p.SetState(352)
			p.Match(AsmE8ParserT__55)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(354)
			p.Match(AsmE8ParserT__54)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(355)
			p.Tag()
		}
		{
			p.SetState(356)
			p.Match(AsmE8ParserT__56)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(357)
			p.Reg8()
		}
		{
			p.SetState(358)
			p.Match(AsmE8ParserT__55)
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
		p.SetState(362)
		p.Match(AsmE8ParserT__54)
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
		p.Match(AsmE8ParserT__56)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(365)
		p.Imm()
	}
	{
		p.SetState(366)
		p.Match(AsmE8ParserT__55)
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
		p.SetState(368)
		p.Match(AsmE8ParserT__54)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(369)
		p.Reg16()
	}
	{
		p.SetState(370)
		p.Match(AsmE8ParserT__56)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(371)
		p.Reg8()
	}
	{
		p.SetState(372)
		p.Match(AsmE8ParserT__55)
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
		p.SetState(374)
		p.Match(AsmE8ParserT__54)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(375)
		p.Reg16()
	}
	{
		p.SetState(376)
		p.Match(AsmE8ParserT__55)
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
		p.SetState(378)
		p.Tag()
	}
	{
		p.SetState(379)
		p.Match(AsmE8ParserT__58)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(380)
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
		p.SetState(382)
		p.Tag()
	}
	{
		p.SetState(383)
		p.Match(AsmE8ParserT__58)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(384)
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

	p.SetState(390)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(386)
			p.Match(AsmE8ParserSTR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(387)
			p.Match(AsmE8ParserSTR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(388)
			_la = p.GetTokenStream().LA(1)

			if !(_la == AsmE8ParserT__56 || _la == AsmE8ParserT__57) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(389)
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
	p.SetState(404)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case AsmE8ParserT__59:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(392)
			p.Match(AsmE8ParserT__59)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(393)
			p.Ascii()
		}

	case AsmE8ParserT__60:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(394)
			p.Match(AsmE8ParserT__60)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(395)
			p.Ascii()
		}

	case AsmE8ParserT__61:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(396)
			p.Match(AsmE8ParserT__61)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(397)
			p.Imm_list()
		}

	case AsmE8ParserT__62:
		p.EnterOuterAlt(localctx, 4)
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
			p.Imm_list()
		}

	case AsmE8ParserT__63:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(400)
			p.Match(AsmE8ParserT__63)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(401)
			p.Imm_list()
		}

	case AsmE8ParserT__64:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(402)
			p.Match(AsmE8ParserT__64)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(403)
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
		p.SetState(406)
		p.Match(AsmE8ParserT__65)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(407)
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
	p.SetState(415)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(409)
			p.Match(AsmE8ParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(410)
			p.Match(AsmE8ParserT__66)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(411)
			p.Tag()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(412)
			p.Match(AsmE8ParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(413)
			p.Match(AsmE8ParserT__67)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(414)
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
	p.SetState(419)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case AsmE8ParserBINARY, AsmE8ParserINT, AsmE8ParserCHAR, AsmE8ParserHEX:
		{
			p.SetState(417)
			p.Imm()
		}

	case AsmE8ParserSTR:
		{
			p.SetState(418)
			p.Tag()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.SetState(428)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == AsmE8ParserT__6 {
		{
			p.SetState(421)
			p.Match(AsmE8ParserT__6)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(424)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case AsmE8ParserBINARY, AsmE8ParserINT, AsmE8ParserCHAR, AsmE8ParserHEX:
			{
				p.SetState(422)
				p.Imm()
			}

		case AsmE8ParserSTR:
			{
				p.SetState(423)
				p.Tag()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(430)
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
		p.SetState(431)
		_la = p.GetTokenStream().LA(1)

		if !((int64((_la-72)) & ^0x3f) == 0 && ((int64(1)<<(_la-72))&23) != 0) {
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
		p.SetState(433)
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
		p.SetState(435)
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
