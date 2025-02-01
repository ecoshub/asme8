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
		"", "'\\t'", "'    '", "'  '", "' '", "'\\n'", "':'", "', '", "'mov'",
		"'add'", "'adc'", "'sub'", "'sbb'", "'cmp'", "'inc'", "'dec'", "'xor'",
		"'and'", "'or'", "'not'", "'shl'", "'shr'", "'rol'", "'ror'", "'jmp'",
		"'jz'", "'jnz'", "'jc'", "'js'", "'jns'", "'call'", "'ret'", "'brk'",
		"'nop'", "'push'", "'pop'", "'clc'", "'a'", "'b'", "'c'", "'d'", "'ipl'",
		"'iph'", "'sp'", "'ip'", "'['", "']'", "'+'", "'-'", "'='", "'.byte '",
		"'.word '", "'.resb '", "'.org '", "'.segment '", "'\"'", "'global '",
		"'extern '",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "LINE_COMMENT", "SUFFIX_WHITESPACE", "WHITE_SPACE",
		"BINARY", "INT", "CHAR", "HEX", "STR",
	}
	staticData.RuleNames = []string{
		"prog", "instruction", "line", "label", "inst", "inst_reg_reg", "inst_reg_imm",
		"inst_stack_imm", "inst_index_register_imm_variable", "inst_index_imm",
		"inst_reg_imm_variable", "inst_ptr_reg", "inst_ptr_imm", "inst_reg_ptr",
		"inst_reg_ptr_offset", "inst_indirect_reg_stack", "inst_indirect_stack_register",
		"inst_indirect_reg_index", "inst_indirect_index_register", "inst_ptr_offset_reg",
		"inst_single_reg", "inst_implied_stack", "inst_implied_index", "inst_single_imm",
		"inst_single_tag", "inst_single", "mnemonic", "reg", "stack", "index",
		"ptr", "ptr_virtual_offset", "ptr_offset", "stack_offset", "index_offset",
		"variable", "variable_reference", "reference", "directives", "segment",
		"access", "imm_list", "imm", "tag",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 65, 393, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36, 7, 36,
		2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7, 41, 2,
		42, 7, 42, 2, 43, 7, 43, 1, 0, 1, 0, 1, 0, 1, 1, 4, 1, 93, 8, 1, 11, 1,
		12, 1, 94, 1, 1, 3, 1, 98, 8, 1, 1, 1, 4, 1, 101, 8, 1, 11, 1, 12, 1, 102,
		1, 1, 3, 1, 106, 8, 1, 3, 1, 108, 8, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1,
		2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 123, 8, 2, 1, 3, 1,
		3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1,
		4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 3, 4, 149,
		8, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6,
		1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8,
		1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1,
		10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12,
		1, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 14, 1,
		14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15,
		1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 1, 17, 1,
		17, 1, 17, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 19, 1, 19, 1, 19,
		1, 19, 1, 19, 1, 19, 1, 20, 1, 20, 1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 1,
		21, 1, 22, 1, 22, 1, 22, 1, 22, 1, 23, 1, 23, 1, 23, 1, 23, 1, 24, 1, 24,
		1, 24, 1, 24, 1, 25, 1, 25, 1, 26, 1, 26, 1, 27, 1, 27, 1, 28, 1, 28, 1,
		29, 1, 29, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30,
		3, 30, 280, 8, 30, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 32, 1,
		32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32,
		3, 32, 300, 8, 32, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1,
		33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 3, 33, 318,
		8, 33, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1,
		34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 3, 34, 336, 8, 34, 1, 35,
		1, 35, 1, 35, 1, 35, 1, 36, 1, 36, 1, 36, 1, 36, 1, 37, 1, 37, 1, 37, 1,
		37, 3, 37, 350, 8, 37, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38,
		1, 38, 3, 38, 360, 8, 38, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 1, 40, 1,
		40, 1, 40, 1, 40, 1, 40, 1, 40, 3, 40, 373, 8, 40, 1, 41, 1, 41, 3, 41,
		377, 8, 41, 1, 41, 1, 41, 1, 41, 3, 41, 382, 8, 41, 5, 41, 384, 8, 41,
		10, 41, 12, 41, 387, 9, 41, 1, 42, 1, 42, 1, 43, 1, 43, 1, 43, 0, 0, 44,
		0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36,
		38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 70, 72,
		74, 76, 78, 80, 82, 84, 86, 0, 5, 1, 0, 1, 3, 1, 0, 8, 36, 1, 0, 37, 42,
		1, 0, 47, 48, 1, 0, 61, 64, 396, 0, 88, 1, 0, 0, 0, 2, 107, 1, 0, 0, 0,
		4, 122, 1, 0, 0, 0, 6, 124, 1, 0, 0, 0, 8, 148, 1, 0, 0, 0, 10, 150, 1,
		0, 0, 0, 12, 156, 1, 0, 0, 0, 14, 162, 1, 0, 0, 0, 16, 168, 1, 0, 0, 0,
		18, 174, 1, 0, 0, 0, 20, 180, 1, 0, 0, 0, 22, 186, 1, 0, 0, 0, 24, 192,
		1, 0, 0, 0, 26, 198, 1, 0, 0, 0, 28, 204, 1, 0, 0, 0, 30, 210, 1, 0, 0,
		0, 32, 216, 1, 0, 0, 0, 34, 222, 1, 0, 0, 0, 36, 228, 1, 0, 0, 0, 38, 234,
		1, 0, 0, 0, 40, 240, 1, 0, 0, 0, 42, 244, 1, 0, 0, 0, 44, 248, 1, 0, 0,
		0, 46, 252, 1, 0, 0, 0, 48, 256, 1, 0, 0, 0, 50, 260, 1, 0, 0, 0, 52, 262,
		1, 0, 0, 0, 54, 264, 1, 0, 0, 0, 56, 266, 1, 0, 0, 0, 58, 268, 1, 0, 0,
		0, 60, 279, 1, 0, 0, 0, 62, 281, 1, 0, 0, 0, 64, 299, 1, 0, 0, 0, 66, 317,
		1, 0, 0, 0, 68, 335, 1, 0, 0, 0, 70, 337, 1, 0, 0, 0, 72, 341, 1, 0, 0,
		0, 74, 349, 1, 0, 0, 0, 76, 359, 1, 0, 0, 0, 78, 361, 1, 0, 0, 0, 80, 372,
		1, 0, 0, 0, 82, 376, 1, 0, 0, 0, 84, 388, 1, 0, 0, 0, 86, 390, 1, 0, 0,
		0, 88, 89, 3, 2, 1, 0, 89, 90, 5, 0, 0, 1, 90, 1, 1, 0, 0, 0, 91, 93, 3,
		4, 2, 0, 92, 91, 1, 0, 0, 0, 93, 94, 1, 0, 0, 0, 94, 92, 1, 0, 0, 0, 94,
		95, 1, 0, 0, 0, 95, 97, 1, 0, 0, 0, 96, 98, 5, 58, 0, 0, 97, 96, 1, 0,
		0, 0, 97, 98, 1, 0, 0, 0, 98, 108, 1, 0, 0, 0, 99, 101, 3, 4, 2, 0, 100,
		99, 1, 0, 0, 0, 101, 102, 1, 0, 0, 0, 102, 100, 1, 0, 0, 0, 102, 103, 1,
		0, 0, 0, 103, 105, 1, 0, 0, 0, 104, 106, 5, 59, 0, 0, 105, 104, 1, 0, 0,
		0, 105, 106, 1, 0, 0, 0, 106, 108, 1, 0, 0, 0, 107, 92, 1, 0, 0, 0, 107,
		100, 1, 0, 0, 0, 108, 3, 1, 0, 0, 0, 109, 123, 3, 72, 36, 0, 110, 123,
		3, 70, 35, 0, 111, 112, 7, 0, 0, 0, 112, 123, 3, 8, 4, 0, 113, 123, 3,
		78, 39, 0, 114, 123, 3, 80, 40, 0, 115, 116, 3, 6, 3, 0, 116, 117, 5, 4,
		0, 0, 117, 118, 3, 76, 38, 0, 118, 123, 1, 0, 0, 0, 119, 123, 3, 76, 38,
		0, 120, 123, 3, 6, 3, 0, 121, 123, 5, 5, 0, 0, 122, 109, 1, 0, 0, 0, 122,
		110, 1, 0, 0, 0, 122, 111, 1, 0, 0, 0, 122, 113, 1, 0, 0, 0, 122, 114,
		1, 0, 0, 0, 122, 115, 1, 0, 0, 0, 122, 119, 1, 0, 0, 0, 122, 120, 1, 0,
		0, 0, 122, 121, 1, 0, 0, 0, 123, 5, 1, 0, 0, 0, 124, 125, 3, 86, 43, 0,
		125, 126, 5, 6, 0, 0, 126, 7, 1, 0, 0, 0, 127, 149, 3, 10, 5, 0, 128, 149,
		3, 12, 6, 0, 129, 149, 3, 14, 7, 0, 130, 149, 3, 18, 9, 0, 131, 149, 3,
		20, 10, 0, 132, 149, 3, 22, 11, 0, 133, 149, 3, 24, 12, 0, 134, 149, 3,
		26, 13, 0, 135, 149, 3, 28, 14, 0, 136, 149, 3, 30, 15, 0, 137, 149, 3,
		32, 16, 0, 138, 149, 3, 34, 17, 0, 139, 149, 3, 36, 18, 0, 140, 149, 3,
		38, 19, 0, 141, 149, 3, 40, 20, 0, 142, 149, 3, 42, 21, 0, 143, 149, 3,
		44, 22, 0, 144, 149, 3, 46, 23, 0, 145, 149, 3, 48, 24, 0, 146, 149, 3,
		50, 25, 0, 147, 149, 3, 16, 8, 0, 148, 127, 1, 0, 0, 0, 148, 128, 1, 0,
		0, 0, 148, 129, 1, 0, 0, 0, 148, 130, 1, 0, 0, 0, 148, 131, 1, 0, 0, 0,
		148, 132, 1, 0, 0, 0, 148, 133, 1, 0, 0, 0, 148, 134, 1, 0, 0, 0, 148,
		135, 1, 0, 0, 0, 148, 136, 1, 0, 0, 0, 148, 137, 1, 0, 0, 0, 148, 138,
		1, 0, 0, 0, 148, 139, 1, 0, 0, 0, 148, 140, 1, 0, 0, 0, 148, 141, 1, 0,
		0, 0, 148, 142, 1, 0, 0, 0, 148, 143, 1, 0, 0, 0, 148, 144, 1, 0, 0, 0,
		148, 145, 1, 0, 0, 0, 148, 146, 1, 0, 0, 0, 148, 147, 1, 0, 0, 0, 149,
		9, 1, 0, 0, 0, 150, 151, 3, 52, 26, 0, 151, 152, 5, 4, 0, 0, 152, 153,
		3, 54, 27, 0, 153, 154, 5, 7, 0, 0, 154, 155, 3, 54, 27, 0, 155, 11, 1,
		0, 0, 0, 156, 157, 3, 52, 26, 0, 157, 158, 5, 4, 0, 0, 158, 159, 3, 54,
		27, 0, 159, 160, 5, 7, 0, 0, 160, 161, 3, 84, 42, 0, 161, 13, 1, 0, 0,
		0, 162, 163, 3, 52, 26, 0, 163, 164, 5, 4, 0, 0, 164, 165, 3, 56, 28, 0,
		165, 166, 5, 7, 0, 0, 166, 167, 3, 84, 42, 0, 167, 15, 1, 0, 0, 0, 168,
		169, 3, 52, 26, 0, 169, 170, 5, 4, 0, 0, 170, 171, 3, 58, 29, 0, 171, 172,
		5, 7, 0, 0, 172, 173, 3, 86, 43, 0, 173, 17, 1, 0, 0, 0, 174, 175, 3, 52,
		26, 0, 175, 176, 5, 4, 0, 0, 176, 177, 3, 58, 29, 0, 177, 178, 5, 7, 0,
		0, 178, 179, 3, 84, 42, 0, 179, 19, 1, 0, 0, 0, 180, 181, 3, 52, 26, 0,
		181, 182, 5, 4, 0, 0, 182, 183, 3, 54, 27, 0, 183, 184, 5, 7, 0, 0, 184,
		185, 3, 86, 43, 0, 185, 21, 1, 0, 0, 0, 186, 187, 3, 52, 26, 0, 187, 188,
		5, 4, 0, 0, 188, 189, 3, 60, 30, 0, 189, 190, 5, 7, 0, 0, 190, 191, 3,
		54, 27, 0, 191, 23, 1, 0, 0, 0, 192, 193, 3, 52, 26, 0, 193, 194, 5, 4,
		0, 0, 194, 195, 3, 60, 30, 0, 195, 196, 5, 7, 0, 0, 196, 197, 3, 84, 42,
		0, 197, 25, 1, 0, 0, 0, 198, 199, 3, 52, 26, 0, 199, 200, 5, 4, 0, 0, 200,
		201, 3, 54, 27, 0, 201, 202, 5, 7, 0, 0, 202, 203, 3, 60, 30, 0, 203, 27,
		1, 0, 0, 0, 204, 205, 3, 52, 26, 0, 205, 206, 5, 4, 0, 0, 206, 207, 3,
		54, 27, 0, 207, 208, 5, 7, 0, 0, 208, 209, 3, 64, 32, 0, 209, 29, 1, 0,
		0, 0, 210, 211, 3, 52, 26, 0, 211, 212, 5, 4, 0, 0, 212, 213, 3, 54, 27,
		0, 213, 214, 5, 7, 0, 0, 214, 215, 3, 66, 33, 0, 215, 31, 1, 0, 0, 0, 216,
		217, 3, 52, 26, 0, 217, 218, 5, 4, 0, 0, 218, 219, 3, 66, 33, 0, 219, 220,
		5, 7, 0, 0, 220, 221, 3, 54, 27, 0, 221, 33, 1, 0, 0, 0, 222, 223, 3, 52,
		26, 0, 223, 224, 5, 4, 0, 0, 224, 225, 3, 54, 27, 0, 225, 226, 5, 7, 0,
		0, 226, 227, 3, 68, 34, 0, 227, 35, 1, 0, 0, 0, 228, 229, 3, 52, 26, 0,
		229, 230, 5, 4, 0, 0, 230, 231, 3, 68, 34, 0, 231, 232, 5, 7, 0, 0, 232,
		233, 3, 54, 27, 0, 233, 37, 1, 0, 0, 0, 234, 235, 3, 52, 26, 0, 235, 236,
		5, 4, 0, 0, 236, 237, 3, 64, 32, 0, 237, 238, 5, 7, 0, 0, 238, 239, 3,
		54, 27, 0, 239, 39, 1, 0, 0, 0, 240, 241, 3, 52, 26, 0, 241, 242, 5, 4,
		0, 0, 242, 243, 3, 54, 27, 0, 243, 41, 1, 0, 0, 0, 244, 245, 3, 52, 26,
		0, 245, 246, 5, 4, 0, 0, 246, 247, 3, 56, 28, 0, 247, 43, 1, 0, 0, 0, 248,
		249, 3, 52, 26, 0, 249, 250, 5, 4, 0, 0, 250, 251, 3, 58, 29, 0, 251, 45,
		1, 0, 0, 0, 252, 253, 3, 52, 26, 0, 253, 254, 5, 4, 0, 0, 254, 255, 3,
		84, 42, 0, 255, 47, 1, 0, 0, 0, 256, 257, 3, 52, 26, 0, 257, 258, 5, 4,
		0, 0, 258, 259, 3, 86, 43, 0, 259, 49, 1, 0, 0, 0, 260, 261, 3, 52, 26,
		0, 261, 51, 1, 0, 0, 0, 262, 263, 7, 1, 0, 0, 263, 53, 1, 0, 0, 0, 264,
		265, 7, 2, 0, 0, 265, 55, 1, 0, 0, 0, 266, 267, 5, 43, 0, 0, 267, 57, 1,
		0, 0, 0, 268, 269, 5, 44, 0, 0, 269, 59, 1, 0, 0, 0, 270, 271, 5, 45, 0,
		0, 271, 272, 3, 84, 42, 0, 272, 273, 5, 46, 0, 0, 273, 280, 1, 0, 0, 0,
		274, 275, 5, 45, 0, 0, 275, 276, 3, 86, 43, 0, 276, 277, 5, 46, 0, 0, 277,
		280, 1, 0, 0, 0, 278, 280, 3, 62, 31, 0, 279, 270, 1, 0, 0, 0, 279, 274,
		1, 0, 0, 0, 279, 278, 1, 0, 0, 0, 280, 61, 1, 0, 0, 0, 281, 282, 5, 45,
		0, 0, 282, 283, 5, 65, 0, 0, 283, 284, 7, 3, 0, 0, 284, 285, 5, 62, 0,
		0, 285, 286, 5, 46, 0, 0, 286, 63, 1, 0, 0, 0, 287, 288, 5, 45, 0, 0, 288,
		289, 3, 84, 42, 0, 289, 290, 5, 47, 0, 0, 290, 291, 3, 54, 27, 0, 291,
		292, 5, 46, 0, 0, 292, 300, 1, 0, 0, 0, 293, 294, 5, 45, 0, 0, 294, 295,
		3, 86, 43, 0, 295, 296, 5, 47, 0, 0, 296, 297, 3, 54, 27, 0, 297, 298,
		5, 46, 0, 0, 298, 300, 1, 0, 0, 0, 299, 287, 1, 0, 0, 0, 299, 293, 1, 0,
		0, 0, 300, 65, 1, 0, 0, 0, 301, 302, 5, 45, 0, 0, 302, 303, 3, 56, 28,
		0, 303, 304, 5, 46, 0, 0, 304, 318, 1, 0, 0, 0, 305, 306, 5, 45, 0, 0,
		306, 307, 3, 56, 28, 0, 307, 308, 5, 47, 0, 0, 308, 309, 3, 84, 42, 0,
		309, 310, 5, 46, 0, 0, 310, 318, 1, 0, 0, 0, 311, 312, 5, 45, 0, 0, 312,
		313, 3, 56, 28, 0, 313, 314, 5, 47, 0, 0, 314, 315, 3, 54, 27, 0, 315,
		316, 5, 46, 0, 0, 316, 318, 1, 0, 0, 0, 317, 301, 1, 0, 0, 0, 317, 305,
		1, 0, 0, 0, 317, 311, 1, 0, 0, 0, 318, 67, 1, 0, 0, 0, 319, 320, 5, 45,
		0, 0, 320, 321, 3, 58, 29, 0, 321, 322, 5, 46, 0, 0, 322, 336, 1, 0, 0,
		0, 323, 324, 5, 45, 0, 0, 324, 325, 3, 58, 29, 0, 325, 326, 5, 47, 0, 0,
		326, 327, 3, 84, 42, 0, 327, 328, 5, 46, 0, 0, 328, 336, 1, 0, 0, 0, 329,
		330, 5, 45, 0, 0, 330, 331, 3, 58, 29, 0, 331, 332, 5, 47, 0, 0, 332, 333,
		3, 54, 27, 0, 333, 334, 5, 46, 0, 0, 334, 336, 1, 0, 0, 0, 335, 319, 1,
		0, 0, 0, 335, 323, 1, 0, 0, 0, 335, 329, 1, 0, 0, 0, 336, 69, 1, 0, 0,
		0, 337, 338, 3, 86, 43, 0, 338, 339, 5, 49, 0, 0, 339, 340, 3, 84, 42,
		0, 340, 71, 1, 0, 0, 0, 341, 342, 3, 86, 43, 0, 342, 343, 5, 49, 0, 0,
		343, 344, 3, 74, 37, 0, 344, 73, 1, 0, 0, 0, 345, 350, 5, 65, 0, 0, 346,
		347, 5, 65, 0, 0, 347, 348, 7, 3, 0, 0, 348, 350, 3, 84, 42, 0, 349, 345,
		1, 0, 0, 0, 349, 346, 1, 0, 0, 0, 350, 75, 1, 0, 0, 0, 351, 352, 5, 50,
		0, 0, 352, 360, 3, 82, 41, 0, 353, 354, 5, 51, 0, 0, 354, 360, 3, 82, 41,
		0, 355, 356, 5, 52, 0, 0, 356, 360, 3, 82, 41, 0, 357, 358, 5, 53, 0, 0,
		358, 360, 3, 82, 41, 0, 359, 351, 1, 0, 0, 0, 359, 353, 1, 0, 0, 0, 359,
		355, 1, 0, 0, 0, 359, 357, 1, 0, 0, 0, 360, 77, 1, 0, 0, 0, 361, 362, 5,
		54, 0, 0, 362, 363, 5, 55, 0, 0, 363, 364, 5, 65, 0, 0, 364, 365, 5, 55,
		0, 0, 365, 79, 1, 0, 0, 0, 366, 367, 5, 2, 0, 0, 367, 368, 5, 56, 0, 0,
		368, 373, 3, 86, 43, 0, 369, 370, 5, 2, 0, 0, 370, 371, 5, 57, 0, 0, 371,
		373, 3, 86, 43, 0, 372, 366, 1, 0, 0, 0, 372, 369, 1, 0, 0, 0, 373, 81,
		1, 0, 0, 0, 374, 377, 3, 84, 42, 0, 375, 377, 3, 86, 43, 0, 376, 374, 1,
		0, 0, 0, 376, 375, 1, 0, 0, 0, 377, 385, 1, 0, 0, 0, 378, 381, 5, 7, 0,
		0, 379, 382, 3, 84, 42, 0, 380, 382, 3, 86, 43, 0, 381, 379, 1, 0, 0, 0,
		381, 380, 1, 0, 0, 0, 382, 384, 1, 0, 0, 0, 383, 378, 1, 0, 0, 0, 384,
		387, 1, 0, 0, 0, 385, 383, 1, 0, 0, 0, 385, 386, 1, 0, 0, 0, 386, 83, 1,
		0, 0, 0, 387, 385, 1, 0, 0, 0, 388, 389, 7, 4, 0, 0, 389, 85, 1, 0, 0,
		0, 390, 391, 5, 65, 0, 0, 391, 87, 1, 0, 0, 0, 17, 94, 97, 102, 105, 107,
		122, 148, 279, 299, 317, 335, 349, 359, 372, 376, 381, 385,
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
	AsmE8ParserLINE_COMMENT      = 58
	AsmE8ParserSUFFIX_WHITESPACE = 59
	AsmE8ParserWHITE_SPACE       = 60
	AsmE8ParserBINARY            = 61
	AsmE8ParserINT               = 62
	AsmE8ParserCHAR              = 63
	AsmE8ParserHEX               = 64
	AsmE8ParserSTR               = 65
)

// AsmE8Parser rules.
const (
	AsmE8ParserRULE_prog                             = 0
	AsmE8ParserRULE_instruction                      = 1
	AsmE8ParserRULE_line                             = 2
	AsmE8ParserRULE_label                            = 3
	AsmE8ParserRULE_inst                             = 4
	AsmE8ParserRULE_inst_reg_reg                     = 5
	AsmE8ParserRULE_inst_reg_imm                     = 6
	AsmE8ParserRULE_inst_stack_imm                   = 7
	AsmE8ParserRULE_inst_index_register_imm_variable = 8
	AsmE8ParserRULE_inst_index_imm                   = 9
	AsmE8ParserRULE_inst_reg_imm_variable            = 10
	AsmE8ParserRULE_inst_ptr_reg                     = 11
	AsmE8ParserRULE_inst_ptr_imm                     = 12
	AsmE8ParserRULE_inst_reg_ptr                     = 13
	AsmE8ParserRULE_inst_reg_ptr_offset              = 14
	AsmE8ParserRULE_inst_indirect_reg_stack          = 15
	AsmE8ParserRULE_inst_indirect_stack_register     = 16
	AsmE8ParserRULE_inst_indirect_reg_index          = 17
	AsmE8ParserRULE_inst_indirect_index_register     = 18
	AsmE8ParserRULE_inst_ptr_offset_reg              = 19
	AsmE8ParserRULE_inst_single_reg                  = 20
	AsmE8ParserRULE_inst_implied_stack               = 21
	AsmE8ParserRULE_inst_implied_index               = 22
	AsmE8ParserRULE_inst_single_imm                  = 23
	AsmE8ParserRULE_inst_single_tag                  = 24
	AsmE8ParserRULE_inst_single                      = 25
	AsmE8ParserRULE_mnemonic                         = 26
	AsmE8ParserRULE_reg                              = 27
	AsmE8ParserRULE_stack                            = 28
	AsmE8ParserRULE_index                            = 29
	AsmE8ParserRULE_ptr                              = 30
	AsmE8ParserRULE_ptr_virtual_offset               = 31
	AsmE8ParserRULE_ptr_offset                       = 32
	AsmE8ParserRULE_stack_offset                     = 33
	AsmE8ParserRULE_index_offset                     = 34
	AsmE8ParserRULE_variable                         = 35
	AsmE8ParserRULE_variable_reference               = 36
	AsmE8ParserRULE_reference                        = 37
	AsmE8ParserRULE_directives                       = 38
	AsmE8ParserRULE_segment                          = 39
	AsmE8ParserRULE_access                           = 40
	AsmE8ParserRULE_imm_list                         = 41
	AsmE8ParserRULE_imm                              = 42
	AsmE8ParserRULE_tag                              = 43
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
		p.SetState(88)
		p.Instruction()
	}
	{
		p.SetState(89)
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

	p.SetState(107)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(92)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&34902897112121390) != 0) || _la == AsmE8ParserSTR {
			{
				p.SetState(91)
				p.Line()
			}

			p.SetState(94)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		p.SetState(97)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == AsmE8ParserLINE_COMMENT {
			{
				p.SetState(96)
				p.Match(AsmE8ParserLINE_COMMENT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(100)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&34902897112121390) != 0) || _la == AsmE8ParserSTR {
			{
				p.SetState(99)
				p.Line()
			}

			p.SetState(102)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		p.SetState(105)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == AsmE8ParserSUFFIX_WHITESPACE {
			{
				p.SetState(104)
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

	p.SetState(122)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(109)
			p.Variable_reference()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(110)
			p.Variable()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(111)
			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&14) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(112)
			p.Inst()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(113)
			p.Segment()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(114)
			p.Access()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(115)
			p.Label()
		}
		{
			p.SetState(116)
			p.Match(AsmE8ParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(117)
			p.Directives()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(119)
			p.Directives()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(120)
			p.Label()
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(121)
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
		p.SetState(124)
		p.Tag()
	}
	{
		p.SetState(125)
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
	Inst_reg_reg() IInst_reg_regContext
	Inst_reg_imm() IInst_reg_immContext
	Inst_stack_imm() IInst_stack_immContext
	Inst_index_imm() IInst_index_immContext
	Inst_reg_imm_variable() IInst_reg_imm_variableContext
	Inst_ptr_reg() IInst_ptr_regContext
	Inst_ptr_imm() IInst_ptr_immContext
	Inst_reg_ptr() IInst_reg_ptrContext
	Inst_reg_ptr_offset() IInst_reg_ptr_offsetContext
	Inst_indirect_reg_stack() IInst_indirect_reg_stackContext
	Inst_indirect_stack_register() IInst_indirect_stack_registerContext
	Inst_indirect_reg_index() IInst_indirect_reg_indexContext
	Inst_indirect_index_register() IInst_indirect_index_registerContext
	Inst_ptr_offset_reg() IInst_ptr_offset_regContext
	Inst_single_reg() IInst_single_regContext
	Inst_implied_stack() IInst_implied_stackContext
	Inst_implied_index() IInst_implied_indexContext
	Inst_single_imm() IInst_single_immContext
	Inst_single_tag() IInst_single_tagContext
	Inst_single() IInst_singleContext
	Inst_index_register_imm_variable() IInst_index_register_imm_variableContext

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

func (s *InstContext) Inst_reg_reg() IInst_reg_regContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInst_reg_regContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInst_reg_regContext)
}

func (s *InstContext) Inst_reg_imm() IInst_reg_immContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInst_reg_immContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInst_reg_immContext)
}

func (s *InstContext) Inst_stack_imm() IInst_stack_immContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInst_stack_immContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInst_stack_immContext)
}

func (s *InstContext) Inst_index_imm() IInst_index_immContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInst_index_immContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInst_index_immContext)
}

func (s *InstContext) Inst_reg_imm_variable() IInst_reg_imm_variableContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInst_reg_imm_variableContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInst_reg_imm_variableContext)
}

func (s *InstContext) Inst_ptr_reg() IInst_ptr_regContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInst_ptr_regContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInst_ptr_regContext)
}

func (s *InstContext) Inst_ptr_imm() IInst_ptr_immContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInst_ptr_immContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInst_ptr_immContext)
}

func (s *InstContext) Inst_reg_ptr() IInst_reg_ptrContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInst_reg_ptrContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInst_reg_ptrContext)
}

func (s *InstContext) Inst_reg_ptr_offset() IInst_reg_ptr_offsetContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInst_reg_ptr_offsetContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInst_reg_ptr_offsetContext)
}

func (s *InstContext) Inst_indirect_reg_stack() IInst_indirect_reg_stackContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInst_indirect_reg_stackContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInst_indirect_reg_stackContext)
}

func (s *InstContext) Inst_indirect_stack_register() IInst_indirect_stack_registerContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInst_indirect_stack_registerContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInst_indirect_stack_registerContext)
}

func (s *InstContext) Inst_indirect_reg_index() IInst_indirect_reg_indexContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInst_indirect_reg_indexContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInst_indirect_reg_indexContext)
}

func (s *InstContext) Inst_indirect_index_register() IInst_indirect_index_registerContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInst_indirect_index_registerContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInst_indirect_index_registerContext)
}

func (s *InstContext) Inst_ptr_offset_reg() IInst_ptr_offset_regContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInst_ptr_offset_regContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInst_ptr_offset_regContext)
}

func (s *InstContext) Inst_single_reg() IInst_single_regContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInst_single_regContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInst_single_regContext)
}

func (s *InstContext) Inst_implied_stack() IInst_implied_stackContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInst_implied_stackContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInst_implied_stackContext)
}

func (s *InstContext) Inst_implied_index() IInst_implied_indexContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInst_implied_indexContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInst_implied_indexContext)
}

func (s *InstContext) Inst_single_imm() IInst_single_immContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInst_single_immContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInst_single_immContext)
}

func (s *InstContext) Inst_single_tag() IInst_single_tagContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInst_single_tagContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInst_single_tagContext)
}

func (s *InstContext) Inst_single() IInst_singleContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInst_singleContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInst_singleContext)
}

func (s *InstContext) Inst_index_register_imm_variable() IInst_index_register_imm_variableContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInst_index_register_imm_variableContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInst_index_register_imm_variableContext)
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
	p.SetState(148)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(127)
			p.Inst_reg_reg()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(128)
			p.Inst_reg_imm()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(129)
			p.Inst_stack_imm()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(130)
			p.Inst_index_imm()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(131)
			p.Inst_reg_imm_variable()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(132)
			p.Inst_ptr_reg()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(133)
			p.Inst_ptr_imm()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(134)
			p.Inst_reg_ptr()
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(135)
			p.Inst_reg_ptr_offset()
		}

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(136)
			p.Inst_indirect_reg_stack()
		}

	case 11:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(137)
			p.Inst_indirect_stack_register()
		}

	case 12:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(138)
			p.Inst_indirect_reg_index()
		}

	case 13:
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(139)
			p.Inst_indirect_index_register()
		}

	case 14:
		p.EnterOuterAlt(localctx, 14)
		{
			p.SetState(140)
			p.Inst_ptr_offset_reg()
		}

	case 15:
		p.EnterOuterAlt(localctx, 15)
		{
			p.SetState(141)
			p.Inst_single_reg()
		}

	case 16:
		p.EnterOuterAlt(localctx, 16)
		{
			p.SetState(142)
			p.Inst_implied_stack()
		}

	case 17:
		p.EnterOuterAlt(localctx, 17)
		{
			p.SetState(143)
			p.Inst_implied_index()
		}

	case 18:
		p.EnterOuterAlt(localctx, 18)
		{
			p.SetState(144)
			p.Inst_single_imm()
		}

	case 19:
		p.EnterOuterAlt(localctx, 19)
		{
			p.SetState(145)
			p.Inst_single_tag()
		}

	case 20:
		p.EnterOuterAlt(localctx, 20)
		{
			p.SetState(146)
			p.Inst_single()
		}

	case 21:
		p.EnterOuterAlt(localctx, 21)
		{
			p.SetState(147)
			p.Inst_index_register_imm_variable()
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

// IInst_reg_regContext is an interface to support dynamic dispatch.
type IInst_reg_regContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Mnemonic() IMnemonicContext
	AllReg() []IRegContext
	Reg(i int) IRegContext

	// IsInst_reg_regContext differentiates from other interfaces.
	IsInst_reg_regContext()
}

type Inst_reg_regContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInst_reg_regContext() *Inst_reg_regContext {
	var p = new(Inst_reg_regContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsmE8ParserRULE_inst_reg_reg
	return p
}

func InitEmptyInst_reg_regContext(p *Inst_reg_regContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsmE8ParserRULE_inst_reg_reg
}

func (*Inst_reg_regContext) IsInst_reg_regContext() {}

func NewInst_reg_regContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Inst_reg_regContext {
	var p = new(Inst_reg_regContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AsmE8ParserRULE_inst_reg_reg

	return p
}

func (s *Inst_reg_regContext) GetParser() antlr.Parser { return s.parser }

func (s *Inst_reg_regContext) Mnemonic() IMnemonicContext {
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

func (s *Inst_reg_regContext) AllReg() []IRegContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IRegContext); ok {
			len++
		}
	}

	tst := make([]IRegContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IRegContext); ok {
			tst[i] = t.(IRegContext)
			i++
		}
	}

	return tst
}

func (s *Inst_reg_regContext) Reg(i int) IRegContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRegContext); ok {
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

	return t.(IRegContext)
}

func (s *Inst_reg_regContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Inst_reg_regContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Inst_reg_regContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsmE8Listener); ok {
		listenerT.EnterInst_reg_reg(s)
	}
}

func (s *Inst_reg_regContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsmE8Listener); ok {
		listenerT.ExitInst_reg_reg(s)
	}
}

func (p *AsmE8Parser) Inst_reg_reg() (localctx IInst_reg_regContext) {
	localctx = NewInst_reg_regContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, AsmE8ParserRULE_inst_reg_reg)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(150)
		p.Mnemonic()
	}
	{
		p.SetState(151)
		p.Match(AsmE8ParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(152)
		p.Reg()
	}
	{
		p.SetState(153)
		p.Match(AsmE8ParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(154)
		p.Reg()
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

// IInst_reg_immContext is an interface to support dynamic dispatch.
type IInst_reg_immContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Mnemonic() IMnemonicContext
	Reg() IRegContext
	Imm() IImmContext

	// IsInst_reg_immContext differentiates from other interfaces.
	IsInst_reg_immContext()
}

type Inst_reg_immContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInst_reg_immContext() *Inst_reg_immContext {
	var p = new(Inst_reg_immContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsmE8ParserRULE_inst_reg_imm
	return p
}

func InitEmptyInst_reg_immContext(p *Inst_reg_immContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsmE8ParserRULE_inst_reg_imm
}

func (*Inst_reg_immContext) IsInst_reg_immContext() {}

func NewInst_reg_immContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Inst_reg_immContext {
	var p = new(Inst_reg_immContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AsmE8ParserRULE_inst_reg_imm

	return p
}

func (s *Inst_reg_immContext) GetParser() antlr.Parser { return s.parser }

func (s *Inst_reg_immContext) Mnemonic() IMnemonicContext {
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

func (s *Inst_reg_immContext) Reg() IRegContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRegContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRegContext)
}

func (s *Inst_reg_immContext) Imm() IImmContext {
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

func (s *Inst_reg_immContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Inst_reg_immContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Inst_reg_immContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsmE8Listener); ok {
		listenerT.EnterInst_reg_imm(s)
	}
}

func (s *Inst_reg_immContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsmE8Listener); ok {
		listenerT.ExitInst_reg_imm(s)
	}
}

func (p *AsmE8Parser) Inst_reg_imm() (localctx IInst_reg_immContext) {
	localctx = NewInst_reg_immContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, AsmE8ParserRULE_inst_reg_imm)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(156)
		p.Mnemonic()
	}
	{
		p.SetState(157)
		p.Match(AsmE8ParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(158)
		p.Reg()
	}
	{
		p.SetState(159)
		p.Match(AsmE8ParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(160)
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

// IInst_stack_immContext is an interface to support dynamic dispatch.
type IInst_stack_immContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Mnemonic() IMnemonicContext
	Stack() IStackContext
	Imm() IImmContext

	// IsInst_stack_immContext differentiates from other interfaces.
	IsInst_stack_immContext()
}

type Inst_stack_immContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInst_stack_immContext() *Inst_stack_immContext {
	var p = new(Inst_stack_immContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsmE8ParserRULE_inst_stack_imm
	return p
}

func InitEmptyInst_stack_immContext(p *Inst_stack_immContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsmE8ParserRULE_inst_stack_imm
}

func (*Inst_stack_immContext) IsInst_stack_immContext() {}

func NewInst_stack_immContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Inst_stack_immContext {
	var p = new(Inst_stack_immContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AsmE8ParserRULE_inst_stack_imm

	return p
}

func (s *Inst_stack_immContext) GetParser() antlr.Parser { return s.parser }

func (s *Inst_stack_immContext) Mnemonic() IMnemonicContext {
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

func (s *Inst_stack_immContext) Stack() IStackContext {
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

func (s *Inst_stack_immContext) Imm() IImmContext {
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

func (s *Inst_stack_immContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Inst_stack_immContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Inst_stack_immContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsmE8Listener); ok {
		listenerT.EnterInst_stack_imm(s)
	}
}

func (s *Inst_stack_immContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsmE8Listener); ok {
		listenerT.ExitInst_stack_imm(s)
	}
}

func (p *AsmE8Parser) Inst_stack_imm() (localctx IInst_stack_immContext) {
	localctx = NewInst_stack_immContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, AsmE8ParserRULE_inst_stack_imm)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(162)
		p.Mnemonic()
	}
	{
		p.SetState(163)
		p.Match(AsmE8ParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(164)
		p.Stack()
	}
	{
		p.SetState(165)
		p.Match(AsmE8ParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(166)
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

// IInst_index_register_imm_variableContext is an interface to support dynamic dispatch.
type IInst_index_register_imm_variableContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Mnemonic() IMnemonicContext
	Index() IIndexContext
	Tag() ITagContext

	// IsInst_index_register_imm_variableContext differentiates from other interfaces.
	IsInst_index_register_imm_variableContext()
}

type Inst_index_register_imm_variableContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInst_index_register_imm_variableContext() *Inst_index_register_imm_variableContext {
	var p = new(Inst_index_register_imm_variableContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsmE8ParserRULE_inst_index_register_imm_variable
	return p
}

func InitEmptyInst_index_register_imm_variableContext(p *Inst_index_register_imm_variableContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsmE8ParserRULE_inst_index_register_imm_variable
}

func (*Inst_index_register_imm_variableContext) IsInst_index_register_imm_variableContext() {}

func NewInst_index_register_imm_variableContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Inst_index_register_imm_variableContext {
	var p = new(Inst_index_register_imm_variableContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AsmE8ParserRULE_inst_index_register_imm_variable

	return p
}

func (s *Inst_index_register_imm_variableContext) GetParser() antlr.Parser { return s.parser }

func (s *Inst_index_register_imm_variableContext) Mnemonic() IMnemonicContext {
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

func (s *Inst_index_register_imm_variableContext) Index() IIndexContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIndexContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIndexContext)
}

func (s *Inst_index_register_imm_variableContext) Tag() ITagContext {
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

func (s *Inst_index_register_imm_variableContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Inst_index_register_imm_variableContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Inst_index_register_imm_variableContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsmE8Listener); ok {
		listenerT.EnterInst_index_register_imm_variable(s)
	}
}

func (s *Inst_index_register_imm_variableContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsmE8Listener); ok {
		listenerT.ExitInst_index_register_imm_variable(s)
	}
}

func (p *AsmE8Parser) Inst_index_register_imm_variable() (localctx IInst_index_register_imm_variableContext) {
	localctx = NewInst_index_register_imm_variableContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, AsmE8ParserRULE_inst_index_register_imm_variable)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(168)
		p.Mnemonic()
	}
	{
		p.SetState(169)
		p.Match(AsmE8ParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(170)
		p.Index()
	}
	{
		p.SetState(171)
		p.Match(AsmE8ParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(172)
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

// IInst_index_immContext is an interface to support dynamic dispatch.
type IInst_index_immContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Mnemonic() IMnemonicContext
	Index() IIndexContext
	Imm() IImmContext

	// IsInst_index_immContext differentiates from other interfaces.
	IsInst_index_immContext()
}

type Inst_index_immContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInst_index_immContext() *Inst_index_immContext {
	var p = new(Inst_index_immContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsmE8ParserRULE_inst_index_imm
	return p
}

func InitEmptyInst_index_immContext(p *Inst_index_immContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsmE8ParserRULE_inst_index_imm
}

func (*Inst_index_immContext) IsInst_index_immContext() {}

func NewInst_index_immContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Inst_index_immContext {
	var p = new(Inst_index_immContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AsmE8ParserRULE_inst_index_imm

	return p
}

func (s *Inst_index_immContext) GetParser() antlr.Parser { return s.parser }

func (s *Inst_index_immContext) Mnemonic() IMnemonicContext {
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

func (s *Inst_index_immContext) Index() IIndexContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIndexContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIndexContext)
}

func (s *Inst_index_immContext) Imm() IImmContext {
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

func (s *Inst_index_immContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Inst_index_immContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Inst_index_immContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsmE8Listener); ok {
		listenerT.EnterInst_index_imm(s)
	}
}

func (s *Inst_index_immContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsmE8Listener); ok {
		listenerT.ExitInst_index_imm(s)
	}
}

func (p *AsmE8Parser) Inst_index_imm() (localctx IInst_index_immContext) {
	localctx = NewInst_index_immContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, AsmE8ParserRULE_inst_index_imm)
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
		p.Index()
	}
	{
		p.SetState(177)
		p.Match(AsmE8ParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(178)
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

// IInst_reg_imm_variableContext is an interface to support dynamic dispatch.
type IInst_reg_imm_variableContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Mnemonic() IMnemonicContext
	Reg() IRegContext
	Tag() ITagContext

	// IsInst_reg_imm_variableContext differentiates from other interfaces.
	IsInst_reg_imm_variableContext()
}

type Inst_reg_imm_variableContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInst_reg_imm_variableContext() *Inst_reg_imm_variableContext {
	var p = new(Inst_reg_imm_variableContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsmE8ParserRULE_inst_reg_imm_variable
	return p
}

func InitEmptyInst_reg_imm_variableContext(p *Inst_reg_imm_variableContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsmE8ParserRULE_inst_reg_imm_variable
}

func (*Inst_reg_imm_variableContext) IsInst_reg_imm_variableContext() {}

func NewInst_reg_imm_variableContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Inst_reg_imm_variableContext {
	var p = new(Inst_reg_imm_variableContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AsmE8ParserRULE_inst_reg_imm_variable

	return p
}

func (s *Inst_reg_imm_variableContext) GetParser() antlr.Parser { return s.parser }

func (s *Inst_reg_imm_variableContext) Mnemonic() IMnemonicContext {
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

func (s *Inst_reg_imm_variableContext) Reg() IRegContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRegContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRegContext)
}

func (s *Inst_reg_imm_variableContext) Tag() ITagContext {
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

func (s *Inst_reg_imm_variableContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Inst_reg_imm_variableContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Inst_reg_imm_variableContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsmE8Listener); ok {
		listenerT.EnterInst_reg_imm_variable(s)
	}
}

func (s *Inst_reg_imm_variableContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsmE8Listener); ok {
		listenerT.ExitInst_reg_imm_variable(s)
	}
}

func (p *AsmE8Parser) Inst_reg_imm_variable() (localctx IInst_reg_imm_variableContext) {
	localctx = NewInst_reg_imm_variableContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, AsmE8ParserRULE_inst_reg_imm_variable)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(180)
		p.Mnemonic()
	}
	{
		p.SetState(181)
		p.Match(AsmE8ParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(182)
		p.Reg()
	}
	{
		p.SetState(183)
		p.Match(AsmE8ParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(184)
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

// IInst_ptr_regContext is an interface to support dynamic dispatch.
type IInst_ptr_regContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Mnemonic() IMnemonicContext
	Ptr() IPtrContext
	Reg() IRegContext

	// IsInst_ptr_regContext differentiates from other interfaces.
	IsInst_ptr_regContext()
}

type Inst_ptr_regContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInst_ptr_regContext() *Inst_ptr_regContext {
	var p = new(Inst_ptr_regContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsmE8ParserRULE_inst_ptr_reg
	return p
}

func InitEmptyInst_ptr_regContext(p *Inst_ptr_regContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsmE8ParserRULE_inst_ptr_reg
}

func (*Inst_ptr_regContext) IsInst_ptr_regContext() {}

func NewInst_ptr_regContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Inst_ptr_regContext {
	var p = new(Inst_ptr_regContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AsmE8ParserRULE_inst_ptr_reg

	return p
}

func (s *Inst_ptr_regContext) GetParser() antlr.Parser { return s.parser }

func (s *Inst_ptr_regContext) Mnemonic() IMnemonicContext {
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

func (s *Inst_ptr_regContext) Ptr() IPtrContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPtrContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPtrContext)
}

func (s *Inst_ptr_regContext) Reg() IRegContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRegContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRegContext)
}

func (s *Inst_ptr_regContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Inst_ptr_regContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Inst_ptr_regContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsmE8Listener); ok {
		listenerT.EnterInst_ptr_reg(s)
	}
}

func (s *Inst_ptr_regContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsmE8Listener); ok {
		listenerT.ExitInst_ptr_reg(s)
	}
}

func (p *AsmE8Parser) Inst_ptr_reg() (localctx IInst_ptr_regContext) {
	localctx = NewInst_ptr_regContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, AsmE8ParserRULE_inst_ptr_reg)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(186)
		p.Mnemonic()
	}
	{
		p.SetState(187)
		p.Match(AsmE8ParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(188)
		p.Ptr()
	}
	{
		p.SetState(189)
		p.Match(AsmE8ParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(190)
		p.Reg()
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

// IInst_ptr_immContext is an interface to support dynamic dispatch.
type IInst_ptr_immContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Mnemonic() IMnemonicContext
	Ptr() IPtrContext
	Imm() IImmContext

	// IsInst_ptr_immContext differentiates from other interfaces.
	IsInst_ptr_immContext()
}

type Inst_ptr_immContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInst_ptr_immContext() *Inst_ptr_immContext {
	var p = new(Inst_ptr_immContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsmE8ParserRULE_inst_ptr_imm
	return p
}

func InitEmptyInst_ptr_immContext(p *Inst_ptr_immContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsmE8ParserRULE_inst_ptr_imm
}

func (*Inst_ptr_immContext) IsInst_ptr_immContext() {}

func NewInst_ptr_immContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Inst_ptr_immContext {
	var p = new(Inst_ptr_immContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AsmE8ParserRULE_inst_ptr_imm

	return p
}

func (s *Inst_ptr_immContext) GetParser() antlr.Parser { return s.parser }

func (s *Inst_ptr_immContext) Mnemonic() IMnemonicContext {
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

func (s *Inst_ptr_immContext) Ptr() IPtrContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPtrContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPtrContext)
}

func (s *Inst_ptr_immContext) Imm() IImmContext {
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

func (s *Inst_ptr_immContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Inst_ptr_immContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Inst_ptr_immContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsmE8Listener); ok {
		listenerT.EnterInst_ptr_imm(s)
	}
}

func (s *Inst_ptr_immContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsmE8Listener); ok {
		listenerT.ExitInst_ptr_imm(s)
	}
}

func (p *AsmE8Parser) Inst_ptr_imm() (localctx IInst_ptr_immContext) {
	localctx = NewInst_ptr_immContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, AsmE8ParserRULE_inst_ptr_imm)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(192)
		p.Mnemonic()
	}
	{
		p.SetState(193)
		p.Match(AsmE8ParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(194)
		p.Ptr()
	}
	{
		p.SetState(195)
		p.Match(AsmE8ParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(196)
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

// IInst_reg_ptrContext is an interface to support dynamic dispatch.
type IInst_reg_ptrContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Mnemonic() IMnemonicContext
	Reg() IRegContext
	Ptr() IPtrContext

	// IsInst_reg_ptrContext differentiates from other interfaces.
	IsInst_reg_ptrContext()
}

type Inst_reg_ptrContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInst_reg_ptrContext() *Inst_reg_ptrContext {
	var p = new(Inst_reg_ptrContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsmE8ParserRULE_inst_reg_ptr
	return p
}

func InitEmptyInst_reg_ptrContext(p *Inst_reg_ptrContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsmE8ParserRULE_inst_reg_ptr
}

func (*Inst_reg_ptrContext) IsInst_reg_ptrContext() {}

func NewInst_reg_ptrContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Inst_reg_ptrContext {
	var p = new(Inst_reg_ptrContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AsmE8ParserRULE_inst_reg_ptr

	return p
}

func (s *Inst_reg_ptrContext) GetParser() antlr.Parser { return s.parser }

func (s *Inst_reg_ptrContext) Mnemonic() IMnemonicContext {
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

func (s *Inst_reg_ptrContext) Reg() IRegContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRegContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRegContext)
}

func (s *Inst_reg_ptrContext) Ptr() IPtrContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPtrContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPtrContext)
}

func (s *Inst_reg_ptrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Inst_reg_ptrContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Inst_reg_ptrContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsmE8Listener); ok {
		listenerT.EnterInst_reg_ptr(s)
	}
}

func (s *Inst_reg_ptrContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsmE8Listener); ok {
		listenerT.ExitInst_reg_ptr(s)
	}
}

func (p *AsmE8Parser) Inst_reg_ptr() (localctx IInst_reg_ptrContext) {
	localctx = NewInst_reg_ptrContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, AsmE8ParserRULE_inst_reg_ptr)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(198)
		p.Mnemonic()
	}
	{
		p.SetState(199)
		p.Match(AsmE8ParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(200)
		p.Reg()
	}
	{
		p.SetState(201)
		p.Match(AsmE8ParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(202)
		p.Ptr()
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

// IInst_reg_ptr_offsetContext is an interface to support dynamic dispatch.
type IInst_reg_ptr_offsetContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Mnemonic() IMnemonicContext
	Reg() IRegContext
	Ptr_offset() IPtr_offsetContext

	// IsInst_reg_ptr_offsetContext differentiates from other interfaces.
	IsInst_reg_ptr_offsetContext()
}

type Inst_reg_ptr_offsetContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInst_reg_ptr_offsetContext() *Inst_reg_ptr_offsetContext {
	var p = new(Inst_reg_ptr_offsetContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsmE8ParserRULE_inst_reg_ptr_offset
	return p
}

func InitEmptyInst_reg_ptr_offsetContext(p *Inst_reg_ptr_offsetContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsmE8ParserRULE_inst_reg_ptr_offset
}

func (*Inst_reg_ptr_offsetContext) IsInst_reg_ptr_offsetContext() {}

func NewInst_reg_ptr_offsetContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Inst_reg_ptr_offsetContext {
	var p = new(Inst_reg_ptr_offsetContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AsmE8ParserRULE_inst_reg_ptr_offset

	return p
}

func (s *Inst_reg_ptr_offsetContext) GetParser() antlr.Parser { return s.parser }

func (s *Inst_reg_ptr_offsetContext) Mnemonic() IMnemonicContext {
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

func (s *Inst_reg_ptr_offsetContext) Reg() IRegContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRegContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRegContext)
}

func (s *Inst_reg_ptr_offsetContext) Ptr_offset() IPtr_offsetContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPtr_offsetContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPtr_offsetContext)
}

func (s *Inst_reg_ptr_offsetContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Inst_reg_ptr_offsetContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Inst_reg_ptr_offsetContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsmE8Listener); ok {
		listenerT.EnterInst_reg_ptr_offset(s)
	}
}

func (s *Inst_reg_ptr_offsetContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsmE8Listener); ok {
		listenerT.ExitInst_reg_ptr_offset(s)
	}
}

func (p *AsmE8Parser) Inst_reg_ptr_offset() (localctx IInst_reg_ptr_offsetContext) {
	localctx = NewInst_reg_ptr_offsetContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, AsmE8ParserRULE_inst_reg_ptr_offset)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(204)
		p.Mnemonic()
	}
	{
		p.SetState(205)
		p.Match(AsmE8ParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(206)
		p.Reg()
	}
	{
		p.SetState(207)
		p.Match(AsmE8ParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(208)
		p.Ptr_offset()
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

// IInst_indirect_reg_stackContext is an interface to support dynamic dispatch.
type IInst_indirect_reg_stackContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Mnemonic() IMnemonicContext
	Reg() IRegContext
	Stack_offset() IStack_offsetContext

	// IsInst_indirect_reg_stackContext differentiates from other interfaces.
	IsInst_indirect_reg_stackContext()
}

type Inst_indirect_reg_stackContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInst_indirect_reg_stackContext() *Inst_indirect_reg_stackContext {
	var p = new(Inst_indirect_reg_stackContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsmE8ParserRULE_inst_indirect_reg_stack
	return p
}

func InitEmptyInst_indirect_reg_stackContext(p *Inst_indirect_reg_stackContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsmE8ParserRULE_inst_indirect_reg_stack
}

func (*Inst_indirect_reg_stackContext) IsInst_indirect_reg_stackContext() {}

func NewInst_indirect_reg_stackContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Inst_indirect_reg_stackContext {
	var p = new(Inst_indirect_reg_stackContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AsmE8ParserRULE_inst_indirect_reg_stack

	return p
}

func (s *Inst_indirect_reg_stackContext) GetParser() antlr.Parser { return s.parser }

func (s *Inst_indirect_reg_stackContext) Mnemonic() IMnemonicContext {
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

func (s *Inst_indirect_reg_stackContext) Reg() IRegContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRegContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRegContext)
}

func (s *Inst_indirect_reg_stackContext) Stack_offset() IStack_offsetContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStack_offsetContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStack_offsetContext)
}

func (s *Inst_indirect_reg_stackContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Inst_indirect_reg_stackContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Inst_indirect_reg_stackContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsmE8Listener); ok {
		listenerT.EnterInst_indirect_reg_stack(s)
	}
}

func (s *Inst_indirect_reg_stackContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsmE8Listener); ok {
		listenerT.ExitInst_indirect_reg_stack(s)
	}
}

func (p *AsmE8Parser) Inst_indirect_reg_stack() (localctx IInst_indirect_reg_stackContext) {
	localctx = NewInst_indirect_reg_stackContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, AsmE8ParserRULE_inst_indirect_reg_stack)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(210)
		p.Mnemonic()
	}
	{
		p.SetState(211)
		p.Match(AsmE8ParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(212)
		p.Reg()
	}
	{
		p.SetState(213)
		p.Match(AsmE8ParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(214)
		p.Stack_offset()
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

// IInst_indirect_stack_registerContext is an interface to support dynamic dispatch.
type IInst_indirect_stack_registerContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Mnemonic() IMnemonicContext
	Stack_offset() IStack_offsetContext
	Reg() IRegContext

	// IsInst_indirect_stack_registerContext differentiates from other interfaces.
	IsInst_indirect_stack_registerContext()
}

type Inst_indirect_stack_registerContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInst_indirect_stack_registerContext() *Inst_indirect_stack_registerContext {
	var p = new(Inst_indirect_stack_registerContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsmE8ParserRULE_inst_indirect_stack_register
	return p
}

func InitEmptyInst_indirect_stack_registerContext(p *Inst_indirect_stack_registerContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsmE8ParserRULE_inst_indirect_stack_register
}

func (*Inst_indirect_stack_registerContext) IsInst_indirect_stack_registerContext() {}

func NewInst_indirect_stack_registerContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Inst_indirect_stack_registerContext {
	var p = new(Inst_indirect_stack_registerContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AsmE8ParserRULE_inst_indirect_stack_register

	return p
}

func (s *Inst_indirect_stack_registerContext) GetParser() antlr.Parser { return s.parser }

func (s *Inst_indirect_stack_registerContext) Mnemonic() IMnemonicContext {
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

func (s *Inst_indirect_stack_registerContext) Stack_offset() IStack_offsetContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStack_offsetContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStack_offsetContext)
}

func (s *Inst_indirect_stack_registerContext) Reg() IRegContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRegContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRegContext)
}

func (s *Inst_indirect_stack_registerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Inst_indirect_stack_registerContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Inst_indirect_stack_registerContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsmE8Listener); ok {
		listenerT.EnterInst_indirect_stack_register(s)
	}
}

func (s *Inst_indirect_stack_registerContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsmE8Listener); ok {
		listenerT.ExitInst_indirect_stack_register(s)
	}
}

func (p *AsmE8Parser) Inst_indirect_stack_register() (localctx IInst_indirect_stack_registerContext) {
	localctx = NewInst_indirect_stack_registerContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, AsmE8ParserRULE_inst_indirect_stack_register)
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
		p.Stack_offset()
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
		p.Reg()
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

// IInst_indirect_reg_indexContext is an interface to support dynamic dispatch.
type IInst_indirect_reg_indexContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Mnemonic() IMnemonicContext
	Reg() IRegContext
	Index_offset() IIndex_offsetContext

	// IsInst_indirect_reg_indexContext differentiates from other interfaces.
	IsInst_indirect_reg_indexContext()
}

type Inst_indirect_reg_indexContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInst_indirect_reg_indexContext() *Inst_indirect_reg_indexContext {
	var p = new(Inst_indirect_reg_indexContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsmE8ParserRULE_inst_indirect_reg_index
	return p
}

func InitEmptyInst_indirect_reg_indexContext(p *Inst_indirect_reg_indexContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsmE8ParserRULE_inst_indirect_reg_index
}

func (*Inst_indirect_reg_indexContext) IsInst_indirect_reg_indexContext() {}

func NewInst_indirect_reg_indexContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Inst_indirect_reg_indexContext {
	var p = new(Inst_indirect_reg_indexContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AsmE8ParserRULE_inst_indirect_reg_index

	return p
}

func (s *Inst_indirect_reg_indexContext) GetParser() antlr.Parser { return s.parser }

func (s *Inst_indirect_reg_indexContext) Mnemonic() IMnemonicContext {
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

func (s *Inst_indirect_reg_indexContext) Reg() IRegContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRegContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRegContext)
}

func (s *Inst_indirect_reg_indexContext) Index_offset() IIndex_offsetContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIndex_offsetContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIndex_offsetContext)
}

func (s *Inst_indirect_reg_indexContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Inst_indirect_reg_indexContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Inst_indirect_reg_indexContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsmE8Listener); ok {
		listenerT.EnterInst_indirect_reg_index(s)
	}
}

func (s *Inst_indirect_reg_indexContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsmE8Listener); ok {
		listenerT.ExitInst_indirect_reg_index(s)
	}
}

func (p *AsmE8Parser) Inst_indirect_reg_index() (localctx IInst_indirect_reg_indexContext) {
	localctx = NewInst_indirect_reg_indexContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, AsmE8ParserRULE_inst_indirect_reg_index)
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
		p.Reg()
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
		p.Index_offset()
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

// IInst_indirect_index_registerContext is an interface to support dynamic dispatch.
type IInst_indirect_index_registerContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Mnemonic() IMnemonicContext
	Index_offset() IIndex_offsetContext
	Reg() IRegContext

	// IsInst_indirect_index_registerContext differentiates from other interfaces.
	IsInst_indirect_index_registerContext()
}

type Inst_indirect_index_registerContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInst_indirect_index_registerContext() *Inst_indirect_index_registerContext {
	var p = new(Inst_indirect_index_registerContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsmE8ParserRULE_inst_indirect_index_register
	return p
}

func InitEmptyInst_indirect_index_registerContext(p *Inst_indirect_index_registerContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsmE8ParserRULE_inst_indirect_index_register
}

func (*Inst_indirect_index_registerContext) IsInst_indirect_index_registerContext() {}

func NewInst_indirect_index_registerContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Inst_indirect_index_registerContext {
	var p = new(Inst_indirect_index_registerContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AsmE8ParserRULE_inst_indirect_index_register

	return p
}

func (s *Inst_indirect_index_registerContext) GetParser() antlr.Parser { return s.parser }

func (s *Inst_indirect_index_registerContext) Mnemonic() IMnemonicContext {
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

func (s *Inst_indirect_index_registerContext) Index_offset() IIndex_offsetContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIndex_offsetContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIndex_offsetContext)
}

func (s *Inst_indirect_index_registerContext) Reg() IRegContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRegContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRegContext)
}

func (s *Inst_indirect_index_registerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Inst_indirect_index_registerContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Inst_indirect_index_registerContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsmE8Listener); ok {
		listenerT.EnterInst_indirect_index_register(s)
	}
}

func (s *Inst_indirect_index_registerContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsmE8Listener); ok {
		listenerT.ExitInst_indirect_index_register(s)
	}
}

func (p *AsmE8Parser) Inst_indirect_index_register() (localctx IInst_indirect_index_registerContext) {
	localctx = NewInst_indirect_index_registerContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, AsmE8ParserRULE_inst_indirect_index_register)
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
		p.Index_offset()
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
		p.Reg()
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

// IInst_ptr_offset_regContext is an interface to support dynamic dispatch.
type IInst_ptr_offset_regContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Mnemonic() IMnemonicContext
	Ptr_offset() IPtr_offsetContext
	Reg() IRegContext

	// IsInst_ptr_offset_regContext differentiates from other interfaces.
	IsInst_ptr_offset_regContext()
}

type Inst_ptr_offset_regContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInst_ptr_offset_regContext() *Inst_ptr_offset_regContext {
	var p = new(Inst_ptr_offset_regContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsmE8ParserRULE_inst_ptr_offset_reg
	return p
}

func InitEmptyInst_ptr_offset_regContext(p *Inst_ptr_offset_regContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsmE8ParserRULE_inst_ptr_offset_reg
}

func (*Inst_ptr_offset_regContext) IsInst_ptr_offset_regContext() {}

func NewInst_ptr_offset_regContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Inst_ptr_offset_regContext {
	var p = new(Inst_ptr_offset_regContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AsmE8ParserRULE_inst_ptr_offset_reg

	return p
}

func (s *Inst_ptr_offset_regContext) GetParser() antlr.Parser { return s.parser }

func (s *Inst_ptr_offset_regContext) Mnemonic() IMnemonicContext {
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

func (s *Inst_ptr_offset_regContext) Ptr_offset() IPtr_offsetContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPtr_offsetContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPtr_offsetContext)
}

func (s *Inst_ptr_offset_regContext) Reg() IRegContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRegContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRegContext)
}

func (s *Inst_ptr_offset_regContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Inst_ptr_offset_regContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Inst_ptr_offset_regContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsmE8Listener); ok {
		listenerT.EnterInst_ptr_offset_reg(s)
	}
}

func (s *Inst_ptr_offset_regContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsmE8Listener); ok {
		listenerT.ExitInst_ptr_offset_reg(s)
	}
}

func (p *AsmE8Parser) Inst_ptr_offset_reg() (localctx IInst_ptr_offset_regContext) {
	localctx = NewInst_ptr_offset_regContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, AsmE8ParserRULE_inst_ptr_offset_reg)
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
		p.Ptr_offset()
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
		p.Reg()
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

// IInst_single_regContext is an interface to support dynamic dispatch.
type IInst_single_regContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Mnemonic() IMnemonicContext
	Reg() IRegContext

	// IsInst_single_regContext differentiates from other interfaces.
	IsInst_single_regContext()
}

type Inst_single_regContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInst_single_regContext() *Inst_single_regContext {
	var p = new(Inst_single_regContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsmE8ParserRULE_inst_single_reg
	return p
}

func InitEmptyInst_single_regContext(p *Inst_single_regContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsmE8ParserRULE_inst_single_reg
}

func (*Inst_single_regContext) IsInst_single_regContext() {}

func NewInst_single_regContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Inst_single_regContext {
	var p = new(Inst_single_regContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AsmE8ParserRULE_inst_single_reg

	return p
}

func (s *Inst_single_regContext) GetParser() antlr.Parser { return s.parser }

func (s *Inst_single_regContext) Mnemonic() IMnemonicContext {
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

func (s *Inst_single_regContext) Reg() IRegContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRegContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRegContext)
}

func (s *Inst_single_regContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Inst_single_regContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Inst_single_regContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsmE8Listener); ok {
		listenerT.EnterInst_single_reg(s)
	}
}

func (s *Inst_single_regContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsmE8Listener); ok {
		listenerT.ExitInst_single_reg(s)
	}
}

func (p *AsmE8Parser) Inst_single_reg() (localctx IInst_single_regContext) {
	localctx = NewInst_single_regContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, AsmE8ParserRULE_inst_single_reg)
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
		p.Reg()
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

// IInst_implied_stackContext is an interface to support dynamic dispatch.
type IInst_implied_stackContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Mnemonic() IMnemonicContext
	Stack() IStackContext

	// IsInst_implied_stackContext differentiates from other interfaces.
	IsInst_implied_stackContext()
}

type Inst_implied_stackContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInst_implied_stackContext() *Inst_implied_stackContext {
	var p = new(Inst_implied_stackContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsmE8ParserRULE_inst_implied_stack
	return p
}

func InitEmptyInst_implied_stackContext(p *Inst_implied_stackContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsmE8ParserRULE_inst_implied_stack
}

func (*Inst_implied_stackContext) IsInst_implied_stackContext() {}

func NewInst_implied_stackContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Inst_implied_stackContext {
	var p = new(Inst_implied_stackContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AsmE8ParserRULE_inst_implied_stack

	return p
}

func (s *Inst_implied_stackContext) GetParser() antlr.Parser { return s.parser }

func (s *Inst_implied_stackContext) Mnemonic() IMnemonicContext {
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

func (s *Inst_implied_stackContext) Stack() IStackContext {
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

func (s *Inst_implied_stackContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Inst_implied_stackContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Inst_implied_stackContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsmE8Listener); ok {
		listenerT.EnterInst_implied_stack(s)
	}
}

func (s *Inst_implied_stackContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsmE8Listener); ok {
		listenerT.ExitInst_implied_stack(s)
	}
}

func (p *AsmE8Parser) Inst_implied_stack() (localctx IInst_implied_stackContext) {
	localctx = NewInst_implied_stackContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, AsmE8ParserRULE_inst_implied_stack)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(244)
		p.Mnemonic()
	}
	{
		p.SetState(245)
		p.Match(AsmE8ParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(246)
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

// IInst_implied_indexContext is an interface to support dynamic dispatch.
type IInst_implied_indexContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Mnemonic() IMnemonicContext
	Index() IIndexContext

	// IsInst_implied_indexContext differentiates from other interfaces.
	IsInst_implied_indexContext()
}

type Inst_implied_indexContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInst_implied_indexContext() *Inst_implied_indexContext {
	var p = new(Inst_implied_indexContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsmE8ParserRULE_inst_implied_index
	return p
}

func InitEmptyInst_implied_indexContext(p *Inst_implied_indexContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsmE8ParserRULE_inst_implied_index
}

func (*Inst_implied_indexContext) IsInst_implied_indexContext() {}

func NewInst_implied_indexContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Inst_implied_indexContext {
	var p = new(Inst_implied_indexContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AsmE8ParserRULE_inst_implied_index

	return p
}

func (s *Inst_implied_indexContext) GetParser() antlr.Parser { return s.parser }

func (s *Inst_implied_indexContext) Mnemonic() IMnemonicContext {
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

func (s *Inst_implied_indexContext) Index() IIndexContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIndexContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIndexContext)
}

func (s *Inst_implied_indexContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Inst_implied_indexContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Inst_implied_indexContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsmE8Listener); ok {
		listenerT.EnterInst_implied_index(s)
	}
}

func (s *Inst_implied_indexContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsmE8Listener); ok {
		listenerT.ExitInst_implied_index(s)
	}
}

func (p *AsmE8Parser) Inst_implied_index() (localctx IInst_implied_indexContext) {
	localctx = NewInst_implied_indexContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, AsmE8ParserRULE_inst_implied_index)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(248)
		p.Mnemonic()
	}
	{
		p.SetState(249)
		p.Match(AsmE8ParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(250)
		p.Index()
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

// IInst_single_immContext is an interface to support dynamic dispatch.
type IInst_single_immContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Mnemonic() IMnemonicContext
	Imm() IImmContext

	// IsInst_single_immContext differentiates from other interfaces.
	IsInst_single_immContext()
}

type Inst_single_immContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInst_single_immContext() *Inst_single_immContext {
	var p = new(Inst_single_immContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsmE8ParserRULE_inst_single_imm
	return p
}

func InitEmptyInst_single_immContext(p *Inst_single_immContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsmE8ParserRULE_inst_single_imm
}

func (*Inst_single_immContext) IsInst_single_immContext() {}

func NewInst_single_immContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Inst_single_immContext {
	var p = new(Inst_single_immContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AsmE8ParserRULE_inst_single_imm

	return p
}

func (s *Inst_single_immContext) GetParser() antlr.Parser { return s.parser }

func (s *Inst_single_immContext) Mnemonic() IMnemonicContext {
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

func (s *Inst_single_immContext) Imm() IImmContext {
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

func (s *Inst_single_immContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Inst_single_immContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Inst_single_immContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsmE8Listener); ok {
		listenerT.EnterInst_single_imm(s)
	}
}

func (s *Inst_single_immContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsmE8Listener); ok {
		listenerT.ExitInst_single_imm(s)
	}
}

func (p *AsmE8Parser) Inst_single_imm() (localctx IInst_single_immContext) {
	localctx = NewInst_single_immContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, AsmE8ParserRULE_inst_single_imm)
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

// IInst_single_tagContext is an interface to support dynamic dispatch.
type IInst_single_tagContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Mnemonic() IMnemonicContext
	Tag() ITagContext

	// IsInst_single_tagContext differentiates from other interfaces.
	IsInst_single_tagContext()
}

type Inst_single_tagContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInst_single_tagContext() *Inst_single_tagContext {
	var p = new(Inst_single_tagContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsmE8ParserRULE_inst_single_tag
	return p
}

func InitEmptyInst_single_tagContext(p *Inst_single_tagContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsmE8ParserRULE_inst_single_tag
}

func (*Inst_single_tagContext) IsInst_single_tagContext() {}

func NewInst_single_tagContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Inst_single_tagContext {
	var p = new(Inst_single_tagContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AsmE8ParserRULE_inst_single_tag

	return p
}

func (s *Inst_single_tagContext) GetParser() antlr.Parser { return s.parser }

func (s *Inst_single_tagContext) Mnemonic() IMnemonicContext {
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

func (s *Inst_single_tagContext) Tag() ITagContext {
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

func (s *Inst_single_tagContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Inst_single_tagContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Inst_single_tagContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsmE8Listener); ok {
		listenerT.EnterInst_single_tag(s)
	}
}

func (s *Inst_single_tagContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsmE8Listener); ok {
		listenerT.ExitInst_single_tag(s)
	}
}

func (p *AsmE8Parser) Inst_single_tag() (localctx IInst_single_tagContext) {
	localctx = NewInst_single_tagContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, AsmE8ParserRULE_inst_single_tag)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(256)
		p.Mnemonic()
	}
	{
		p.SetState(257)
		p.Match(AsmE8ParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(258)
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

// IInst_singleContext is an interface to support dynamic dispatch.
type IInst_singleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Mnemonic() IMnemonicContext

	// IsInst_singleContext differentiates from other interfaces.
	IsInst_singleContext()
}

type Inst_singleContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInst_singleContext() *Inst_singleContext {
	var p = new(Inst_singleContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsmE8ParserRULE_inst_single
	return p
}

func InitEmptyInst_singleContext(p *Inst_singleContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsmE8ParserRULE_inst_single
}

func (*Inst_singleContext) IsInst_singleContext() {}

func NewInst_singleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Inst_singleContext {
	var p = new(Inst_singleContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AsmE8ParserRULE_inst_single

	return p
}

func (s *Inst_singleContext) GetParser() antlr.Parser { return s.parser }

func (s *Inst_singleContext) Mnemonic() IMnemonicContext {
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

func (s *Inst_singleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Inst_singleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Inst_singleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsmE8Listener); ok {
		listenerT.EnterInst_single(s)
	}
}

func (s *Inst_singleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsmE8Listener); ok {
		listenerT.ExitInst_single(s)
	}
}

func (p *AsmE8Parser) Inst_single() (localctx IInst_singleContext) {
	localctx = NewInst_singleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, AsmE8ParserRULE_inst_single)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(260)
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
	p.EnterRule(localctx, 52, AsmE8ParserRULE_mnemonic)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(262)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&137438953216) != 0) {
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

// IRegContext is an interface to support dynamic dispatch.
type IRegContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsRegContext differentiates from other interfaces.
	IsRegContext()
}

type RegContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRegContext() *RegContext {
	var p = new(RegContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsmE8ParserRULE_reg
	return p
}

func InitEmptyRegContext(p *RegContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsmE8ParserRULE_reg
}

func (*RegContext) IsRegContext() {}

func NewRegContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RegContext {
	var p = new(RegContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AsmE8ParserRULE_reg

	return p
}

func (s *RegContext) GetParser() antlr.Parser { return s.parser }
func (s *RegContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RegContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RegContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsmE8Listener); ok {
		listenerT.EnterReg(s)
	}
}

func (s *RegContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsmE8Listener); ok {
		listenerT.ExitReg(s)
	}
}

func (p *AsmE8Parser) Reg() (localctx IRegContext) {
	localctx = NewRegContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, AsmE8ParserRULE_reg)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(264)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&8658654068736) != 0) {
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
	p.EnterRule(localctx, 56, AsmE8ParserRULE_stack)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(266)
		p.Match(AsmE8ParserT__42)
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

// IIndexContext is an interface to support dynamic dispatch.
type IIndexContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsIndexContext differentiates from other interfaces.
	IsIndexContext()
}

type IndexContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIndexContext() *IndexContext {
	var p = new(IndexContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsmE8ParserRULE_index
	return p
}

func InitEmptyIndexContext(p *IndexContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsmE8ParserRULE_index
}

func (*IndexContext) IsIndexContext() {}

func NewIndexContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IndexContext {
	var p = new(IndexContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AsmE8ParserRULE_index

	return p
}

func (s *IndexContext) GetParser() antlr.Parser { return s.parser }
func (s *IndexContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IndexContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IndexContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsmE8Listener); ok {
		listenerT.EnterIndex(s)
	}
}

func (s *IndexContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsmE8Listener); ok {
		listenerT.ExitIndex(s)
	}
}

func (p *AsmE8Parser) Index() (localctx IIndexContext) {
	localctx = NewIndexContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, AsmE8ParserRULE_index)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(268)
		p.Match(AsmE8ParserT__43)
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

// IPtrContext is an interface to support dynamic dispatch.
type IPtrContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Imm() IImmContext
	Tag() ITagContext
	Ptr_virtual_offset() IPtr_virtual_offsetContext

	// IsPtrContext differentiates from other interfaces.
	IsPtrContext()
}

type PtrContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPtrContext() *PtrContext {
	var p = new(PtrContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsmE8ParserRULE_ptr
	return p
}

func InitEmptyPtrContext(p *PtrContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsmE8ParserRULE_ptr
}

func (*PtrContext) IsPtrContext() {}

func NewPtrContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PtrContext {
	var p = new(PtrContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AsmE8ParserRULE_ptr

	return p
}

func (s *PtrContext) GetParser() antlr.Parser { return s.parser }

func (s *PtrContext) Imm() IImmContext {
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

func (s *PtrContext) Tag() ITagContext {
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

func (s *PtrContext) Ptr_virtual_offset() IPtr_virtual_offsetContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPtr_virtual_offsetContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPtr_virtual_offsetContext)
}

func (s *PtrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PtrContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PtrContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsmE8Listener); ok {
		listenerT.EnterPtr(s)
	}
}

func (s *PtrContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsmE8Listener); ok {
		listenerT.ExitPtr(s)
	}
}

func (p *AsmE8Parser) Ptr() (localctx IPtrContext) {
	localctx = NewPtrContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, AsmE8ParserRULE_ptr)
	p.SetState(279)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 7, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(270)
			p.Match(AsmE8ParserT__44)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(271)
			p.Imm()
		}
		{
			p.SetState(272)
			p.Match(AsmE8ParserT__45)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(274)
			p.Match(AsmE8ParserT__44)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(275)
			p.Tag()
		}
		{
			p.SetState(276)
			p.Match(AsmE8ParserT__45)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(278)
			p.Ptr_virtual_offset()
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

// IPtr_virtual_offsetContext is an interface to support dynamic dispatch.
type IPtr_virtual_offsetContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	STR() antlr.TerminalNode
	INT() antlr.TerminalNode

	// IsPtr_virtual_offsetContext differentiates from other interfaces.
	IsPtr_virtual_offsetContext()
}

type Ptr_virtual_offsetContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPtr_virtual_offsetContext() *Ptr_virtual_offsetContext {
	var p = new(Ptr_virtual_offsetContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsmE8ParserRULE_ptr_virtual_offset
	return p
}

func InitEmptyPtr_virtual_offsetContext(p *Ptr_virtual_offsetContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsmE8ParserRULE_ptr_virtual_offset
}

func (*Ptr_virtual_offsetContext) IsPtr_virtual_offsetContext() {}

func NewPtr_virtual_offsetContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Ptr_virtual_offsetContext {
	var p = new(Ptr_virtual_offsetContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AsmE8ParserRULE_ptr_virtual_offset

	return p
}

func (s *Ptr_virtual_offsetContext) GetParser() antlr.Parser { return s.parser }

func (s *Ptr_virtual_offsetContext) STR() antlr.TerminalNode {
	return s.GetToken(AsmE8ParserSTR, 0)
}

func (s *Ptr_virtual_offsetContext) INT() antlr.TerminalNode {
	return s.GetToken(AsmE8ParserINT, 0)
}

func (s *Ptr_virtual_offsetContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Ptr_virtual_offsetContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Ptr_virtual_offsetContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsmE8Listener); ok {
		listenerT.EnterPtr_virtual_offset(s)
	}
}

func (s *Ptr_virtual_offsetContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsmE8Listener); ok {
		listenerT.ExitPtr_virtual_offset(s)
	}
}

func (p *AsmE8Parser) Ptr_virtual_offset() (localctx IPtr_virtual_offsetContext) {
	localctx = NewPtr_virtual_offsetContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, AsmE8ParserRULE_ptr_virtual_offset)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(281)
		p.Match(AsmE8ParserT__44)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(282)
		p.Match(AsmE8ParserSTR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(283)
		_la = p.GetTokenStream().LA(1)

		if !(_la == AsmE8ParserT__46 || _la == AsmE8ParserT__47) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(284)
		p.Match(AsmE8ParserINT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(285)
		p.Match(AsmE8ParserT__45)
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

// IPtr_offsetContext is an interface to support dynamic dispatch.
type IPtr_offsetContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Imm() IImmContext
	Reg() IRegContext
	Tag() ITagContext

	// IsPtr_offsetContext differentiates from other interfaces.
	IsPtr_offsetContext()
}

type Ptr_offsetContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPtr_offsetContext() *Ptr_offsetContext {
	var p = new(Ptr_offsetContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsmE8ParserRULE_ptr_offset
	return p
}

func InitEmptyPtr_offsetContext(p *Ptr_offsetContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsmE8ParserRULE_ptr_offset
}

func (*Ptr_offsetContext) IsPtr_offsetContext() {}

func NewPtr_offsetContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Ptr_offsetContext {
	var p = new(Ptr_offsetContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AsmE8ParserRULE_ptr_offset

	return p
}

func (s *Ptr_offsetContext) GetParser() antlr.Parser { return s.parser }

func (s *Ptr_offsetContext) Imm() IImmContext {
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

func (s *Ptr_offsetContext) Reg() IRegContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRegContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRegContext)
}

func (s *Ptr_offsetContext) Tag() ITagContext {
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

func (s *Ptr_offsetContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Ptr_offsetContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Ptr_offsetContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsmE8Listener); ok {
		listenerT.EnterPtr_offset(s)
	}
}

func (s *Ptr_offsetContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsmE8Listener); ok {
		listenerT.ExitPtr_offset(s)
	}
}

func (p *AsmE8Parser) Ptr_offset() (localctx IPtr_offsetContext) {
	localctx = NewPtr_offsetContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, AsmE8ParserRULE_ptr_offset)
	p.SetState(299)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(287)
			p.Match(AsmE8ParserT__44)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(288)
			p.Imm()
		}
		{
			p.SetState(289)
			p.Match(AsmE8ParserT__46)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(290)
			p.Reg()
		}
		{
			p.SetState(291)
			p.Match(AsmE8ParserT__45)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(293)
			p.Match(AsmE8ParserT__44)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(294)
			p.Tag()
		}
		{
			p.SetState(295)
			p.Match(AsmE8ParserT__46)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(296)
			p.Reg()
		}
		{
			p.SetState(297)
			p.Match(AsmE8ParserT__45)
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

// IStack_offsetContext is an interface to support dynamic dispatch.
type IStack_offsetContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Stack() IStackContext
	Imm() IImmContext
	Reg() IRegContext

	// IsStack_offsetContext differentiates from other interfaces.
	IsStack_offsetContext()
}

type Stack_offsetContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStack_offsetContext() *Stack_offsetContext {
	var p = new(Stack_offsetContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsmE8ParserRULE_stack_offset
	return p
}

func InitEmptyStack_offsetContext(p *Stack_offsetContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsmE8ParserRULE_stack_offset
}

func (*Stack_offsetContext) IsStack_offsetContext() {}

func NewStack_offsetContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Stack_offsetContext {
	var p = new(Stack_offsetContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AsmE8ParserRULE_stack_offset

	return p
}

func (s *Stack_offsetContext) GetParser() antlr.Parser { return s.parser }

func (s *Stack_offsetContext) Stack() IStackContext {
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

func (s *Stack_offsetContext) Imm() IImmContext {
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

func (s *Stack_offsetContext) Reg() IRegContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRegContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRegContext)
}

func (s *Stack_offsetContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Stack_offsetContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Stack_offsetContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsmE8Listener); ok {
		listenerT.EnterStack_offset(s)
	}
}

func (s *Stack_offsetContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsmE8Listener); ok {
		listenerT.ExitStack_offset(s)
	}
}

func (p *AsmE8Parser) Stack_offset() (localctx IStack_offsetContext) {
	localctx = NewStack_offsetContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, AsmE8ParserRULE_stack_offset)
	p.SetState(317)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(301)
			p.Match(AsmE8ParserT__44)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(302)
			p.Stack()
		}
		{
			p.SetState(303)
			p.Match(AsmE8ParserT__45)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(305)
			p.Match(AsmE8ParserT__44)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(306)
			p.Stack()
		}
		{
			p.SetState(307)
			p.Match(AsmE8ParserT__46)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(308)
			p.Imm()
		}
		{
			p.SetState(309)
			p.Match(AsmE8ParserT__45)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(311)
			p.Match(AsmE8ParserT__44)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(312)
			p.Stack()
		}
		{
			p.SetState(313)
			p.Match(AsmE8ParserT__46)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(314)
			p.Reg()
		}
		{
			p.SetState(315)
			p.Match(AsmE8ParserT__45)
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

// IIndex_offsetContext is an interface to support dynamic dispatch.
type IIndex_offsetContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Index() IIndexContext
	Imm() IImmContext
	Reg() IRegContext

	// IsIndex_offsetContext differentiates from other interfaces.
	IsIndex_offsetContext()
}

type Index_offsetContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIndex_offsetContext() *Index_offsetContext {
	var p = new(Index_offsetContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsmE8ParserRULE_index_offset
	return p
}

func InitEmptyIndex_offsetContext(p *Index_offsetContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsmE8ParserRULE_index_offset
}

func (*Index_offsetContext) IsIndex_offsetContext() {}

func NewIndex_offsetContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Index_offsetContext {
	var p = new(Index_offsetContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AsmE8ParserRULE_index_offset

	return p
}

func (s *Index_offsetContext) GetParser() antlr.Parser { return s.parser }

func (s *Index_offsetContext) Index() IIndexContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIndexContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIndexContext)
}

func (s *Index_offsetContext) Imm() IImmContext {
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

func (s *Index_offsetContext) Reg() IRegContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRegContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRegContext)
}

func (s *Index_offsetContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Index_offsetContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Index_offsetContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsmE8Listener); ok {
		listenerT.EnterIndex_offset(s)
	}
}

func (s *Index_offsetContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsmE8Listener); ok {
		listenerT.ExitIndex_offset(s)
	}
}

func (p *AsmE8Parser) Index_offset() (localctx IIndex_offsetContext) {
	localctx = NewIndex_offsetContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, AsmE8ParserRULE_index_offset)
	p.SetState(335)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 10, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(319)
			p.Match(AsmE8ParserT__44)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(320)
			p.Index()
		}
		{
			p.SetState(321)
			p.Match(AsmE8ParserT__45)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(323)
			p.Match(AsmE8ParserT__44)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(324)
			p.Index()
		}
		{
			p.SetState(325)
			p.Match(AsmE8ParserT__46)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(326)
			p.Imm()
		}
		{
			p.SetState(327)
			p.Match(AsmE8ParserT__45)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(329)
			p.Match(AsmE8ParserT__44)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(330)
			p.Index()
		}
		{
			p.SetState(331)
			p.Match(AsmE8ParserT__46)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(332)
			p.Reg()
		}
		{
			p.SetState(333)
			p.Match(AsmE8ParserT__45)
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
	p.EnterRule(localctx, 70, AsmE8ParserRULE_variable)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(337)
		p.Tag()
	}
	{
		p.SetState(338)
		p.Match(AsmE8ParserT__48)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(339)
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
	p.EnterRule(localctx, 72, AsmE8ParserRULE_variable_reference)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(341)
		p.Tag()
	}
	{
		p.SetState(342)
		p.Match(AsmE8ParserT__48)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(343)
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
	p.EnterRule(localctx, 74, AsmE8ParserRULE_reference)
	var _la int

	p.SetState(349)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(345)
			p.Match(AsmE8ParserSTR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(346)
			p.Match(AsmE8ParserSTR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(347)
			_la = p.GetTokenStream().LA(1)

			if !(_la == AsmE8ParserT__46 || _la == AsmE8ParserT__47) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(348)
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
	p.EnterRule(localctx, 76, AsmE8ParserRULE_directives)
	p.SetState(359)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case AsmE8ParserT__49:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(351)
			p.Match(AsmE8ParserT__49)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(352)
			p.Imm_list()
		}

	case AsmE8ParserT__50:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(353)
			p.Match(AsmE8ParserT__50)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(354)
			p.Imm_list()
		}

	case AsmE8ParserT__51:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(355)
			p.Match(AsmE8ParserT__51)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(356)
			p.Imm_list()
		}

	case AsmE8ParserT__52:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(357)
			p.Match(AsmE8ParserT__52)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(358)
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
	STR() antlr.TerminalNode

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

func (s *SegmentContext) STR() antlr.TerminalNode {
	return s.GetToken(AsmE8ParserSTR, 0)
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
	p.EnterRule(localctx, 78, AsmE8ParserRULE_segment)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(361)
		p.Match(AsmE8ParserT__53)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
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
		p.Match(AsmE8ParserSTR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(364)
		p.Match(AsmE8ParserT__54)
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
	p.EnterRule(localctx, 80, AsmE8ParserRULE_access)
	p.SetState(372)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 13, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(366)
			p.Match(AsmE8ParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(367)
			p.Match(AsmE8ParserT__55)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(368)
			p.Tag()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(369)
			p.Match(AsmE8ParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
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
	p.EnterRule(localctx, 82, AsmE8ParserRULE_imm_list)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(376)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case AsmE8ParserBINARY, AsmE8ParserINT, AsmE8ParserCHAR, AsmE8ParserHEX:
		{
			p.SetState(374)
			p.Imm()
		}

	case AsmE8ParserSTR:
		{
			p.SetState(375)
			p.Tag()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.SetState(385)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == AsmE8ParserT__6 {
		{
			p.SetState(378)
			p.Match(AsmE8ParserT__6)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(381)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case AsmE8ParserBINARY, AsmE8ParserINT, AsmE8ParserCHAR, AsmE8ParserHEX:
			{
				p.SetState(379)
				p.Imm()
			}

		case AsmE8ParserSTR:
			{
				p.SetState(380)
				p.Tag()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(387)
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
	p.EnterRule(localctx, 84, AsmE8ParserRULE_imm)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(388)
		_la = p.GetTokenStream().LA(1)

		if !((int64((_la-61)) & ^0x3f) == 0 && ((int64(1)<<(_la-61))&15) != 0) {
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
	p.EnterRule(localctx, 86, AsmE8ParserRULE_tag)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(390)
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
