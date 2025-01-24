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
		"'jmp'", "'jz'", "'jnz'", "'jsr'", "'rts'", "'brk'", "'nop'", "'push'",
		"'pop'", "'a'", "'b'", "'c'", "'d'", "'['", "']'", "'+'", "'-'", "'='",
		"'.byte '", "'.word '", "'.resb '",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "LINE_COMMENT", "SUFFIX_WHITESPACE", "WHITE_SPACE",
		"BINARY", "INT", "CHAR", "HEX", "STR",
	}
	staticData.RuleNames = []string{
		"prog", "instruction", "line", "label", "inst", "inst_reg_reg", "inst_reg_imm",
		"inst_reg_imm_variable", "inst_ptr_reg", "inst_reg_ptr", "inst_reg_ptr_offset",
		"inst_ptr_offset_reg", "inst_single_reg", "inst_single_imm", "inst_single_tag",
		"inst_single", "mnemonic", "reg", "ptr", "ptr_virtual_offset", "ptr_offset",
		"variable", "directives", "imm_list", "imm", "tag",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 45, 265, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 1, 0,
		1, 0, 1, 0, 1, 1, 4, 1, 57, 8, 1, 11, 1, 12, 1, 58, 1, 1, 3, 1, 62, 8,
		1, 1, 1, 4, 1, 65, 8, 1, 11, 1, 12, 1, 66, 1, 1, 3, 1, 70, 8, 1, 3, 1,
		72, 8, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3,
		2, 84, 8, 2, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4,
		1, 4, 1, 4, 1, 4, 1, 4, 3, 4, 100, 8, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5,
		1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7,
		1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9,
		1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1,
		11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 1, 13,
		1, 14, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 16, 1, 16, 1, 17, 1, 17, 1,
		18, 1, 18, 3, 18, 164, 8, 18, 1, 18, 1, 18, 3, 18, 168, 8, 18, 1, 18, 1,
		18, 1, 18, 1, 18, 3, 18, 174, 8, 18, 1, 18, 1, 18, 3, 18, 178, 8, 18, 1,
		18, 1, 18, 1, 18, 3, 18, 183, 8, 18, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19,
		1, 19, 1, 20, 1, 20, 3, 20, 193, 8, 20, 1, 20, 1, 20, 3, 20, 197, 8, 20,
		1, 20, 1, 20, 3, 20, 201, 8, 20, 1, 20, 1, 20, 3, 20, 205, 8, 20, 1, 20,
		1, 20, 1, 20, 1, 20, 3, 20, 211, 8, 20, 1, 20, 1, 20, 3, 20, 215, 8, 20,
		1, 20, 1, 20, 3, 20, 219, 8, 20, 1, 20, 1, 20, 3, 20, 223, 8, 20, 1, 20,
		1, 20, 3, 20, 227, 8, 20, 1, 21, 1, 21, 3, 21, 231, 8, 21, 1, 21, 1, 21,
		3, 21, 235, 8, 21, 1, 21, 1, 21, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1,
		22, 3, 22, 245, 8, 22, 1, 23, 1, 23, 3, 23, 249, 8, 23, 1, 23, 1, 23, 1,
		23, 3, 23, 254, 8, 23, 5, 23, 256, 8, 23, 10, 23, 12, 23, 259, 9, 23, 1,
		24, 1, 24, 1, 25, 1, 25, 1, 25, 0, 0, 26, 0, 2, 4, 6, 8, 10, 12, 14, 16,
		18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 0,
		5, 1, 0, 1, 3, 1, 0, 8, 25, 1, 0, 26, 29, 1, 0, 32, 33, 1, 0, 41, 44, 280,
		0, 52, 1, 0, 0, 0, 2, 71, 1, 0, 0, 0, 4, 83, 1, 0, 0, 0, 6, 85, 1, 0, 0,
		0, 8, 99, 1, 0, 0, 0, 10, 101, 1, 0, 0, 0, 12, 107, 1, 0, 0, 0, 14, 113,
		1, 0, 0, 0, 16, 119, 1, 0, 0, 0, 18, 125, 1, 0, 0, 0, 20, 131, 1, 0, 0,
		0, 22, 137, 1, 0, 0, 0, 24, 143, 1, 0, 0, 0, 26, 147, 1, 0, 0, 0, 28, 151,
		1, 0, 0, 0, 30, 155, 1, 0, 0, 0, 32, 157, 1, 0, 0, 0, 34, 159, 1, 0, 0,
		0, 36, 182, 1, 0, 0, 0, 38, 184, 1, 0, 0, 0, 40, 226, 1, 0, 0, 0, 42, 228,
		1, 0, 0, 0, 44, 244, 1, 0, 0, 0, 46, 248, 1, 0, 0, 0, 48, 260, 1, 0, 0,
		0, 50, 262, 1, 0, 0, 0, 52, 53, 3, 2, 1, 0, 53, 54, 5, 0, 0, 1, 54, 1,
		1, 0, 0, 0, 55, 57, 3, 4, 2, 0, 56, 55, 1, 0, 0, 0, 57, 58, 1, 0, 0, 0,
		58, 56, 1, 0, 0, 0, 58, 59, 1, 0, 0, 0, 59, 61, 1, 0, 0, 0, 60, 62, 5,
		38, 0, 0, 61, 60, 1, 0, 0, 0, 61, 62, 1, 0, 0, 0, 62, 72, 1, 0, 0, 0, 63,
		65, 3, 4, 2, 0, 64, 63, 1, 0, 0, 0, 65, 66, 1, 0, 0, 0, 66, 64, 1, 0, 0,
		0, 66, 67, 1, 0, 0, 0, 67, 69, 1, 0, 0, 0, 68, 70, 5, 39, 0, 0, 69, 68,
		1, 0, 0, 0, 69, 70, 1, 0, 0, 0, 70, 72, 1, 0, 0, 0, 71, 56, 1, 0, 0, 0,
		71, 64, 1, 0, 0, 0, 72, 3, 1, 0, 0, 0, 73, 84, 3, 42, 21, 0, 74, 75, 7,
		0, 0, 0, 75, 84, 3, 8, 4, 0, 76, 77, 3, 6, 3, 0, 77, 78, 5, 4, 0, 0, 78,
		79, 3, 44, 22, 0, 79, 84, 1, 0, 0, 0, 80, 84, 3, 44, 22, 0, 81, 84, 3,
		6, 3, 0, 82, 84, 5, 5, 0, 0, 83, 73, 1, 0, 0, 0, 83, 74, 1, 0, 0, 0, 83,
		76, 1, 0, 0, 0, 83, 80, 1, 0, 0, 0, 83, 81, 1, 0, 0, 0, 83, 82, 1, 0, 0,
		0, 84, 5, 1, 0, 0, 0, 85, 86, 3, 50, 25, 0, 86, 87, 5, 6, 0, 0, 87, 7,
		1, 0, 0, 0, 88, 100, 3, 10, 5, 0, 89, 100, 3, 12, 6, 0, 90, 100, 3, 14,
		7, 0, 91, 100, 3, 24, 12, 0, 92, 100, 3, 26, 13, 0, 93, 100, 3, 28, 14,
		0, 94, 100, 3, 30, 15, 0, 95, 100, 3, 16, 8, 0, 96, 100, 3, 18, 9, 0, 97,
		100, 3, 20, 10, 0, 98, 100, 3, 22, 11, 0, 99, 88, 1, 0, 0, 0, 99, 89, 1,
		0, 0, 0, 99, 90, 1, 0, 0, 0, 99, 91, 1, 0, 0, 0, 99, 92, 1, 0, 0, 0, 99,
		93, 1, 0, 0, 0, 99, 94, 1, 0, 0, 0, 99, 95, 1, 0, 0, 0, 99, 96, 1, 0, 0,
		0, 99, 97, 1, 0, 0, 0, 99, 98, 1, 0, 0, 0, 100, 9, 1, 0, 0, 0, 101, 102,
		3, 32, 16, 0, 102, 103, 5, 4, 0, 0, 103, 104, 3, 34, 17, 0, 104, 105, 5,
		7, 0, 0, 105, 106, 3, 34, 17, 0, 106, 11, 1, 0, 0, 0, 107, 108, 3, 32,
		16, 0, 108, 109, 5, 4, 0, 0, 109, 110, 3, 34, 17, 0, 110, 111, 5, 7, 0,
		0, 111, 112, 3, 48, 24, 0, 112, 13, 1, 0, 0, 0, 113, 114, 3, 32, 16, 0,
		114, 115, 5, 4, 0, 0, 115, 116, 3, 34, 17, 0, 116, 117, 5, 7, 0, 0, 117,
		118, 3, 50, 25, 0, 118, 15, 1, 0, 0, 0, 119, 120, 3, 32, 16, 0, 120, 121,
		5, 4, 0, 0, 121, 122, 3, 36, 18, 0, 122, 123, 5, 7, 0, 0, 123, 124, 3,
		34, 17, 0, 124, 17, 1, 0, 0, 0, 125, 126, 3, 32, 16, 0, 126, 127, 5, 4,
		0, 0, 127, 128, 3, 34, 17, 0, 128, 129, 5, 7, 0, 0, 129, 130, 3, 36, 18,
		0, 130, 19, 1, 0, 0, 0, 131, 132, 3, 32, 16, 0, 132, 133, 5, 4, 0, 0, 133,
		134, 3, 34, 17, 0, 134, 135, 5, 7, 0, 0, 135, 136, 3, 40, 20, 0, 136, 21,
		1, 0, 0, 0, 137, 138, 3, 32, 16, 0, 138, 139, 5, 4, 0, 0, 139, 140, 3,
		40, 20, 0, 140, 141, 5, 7, 0, 0, 141, 142, 3, 34, 17, 0, 142, 23, 1, 0,
		0, 0, 143, 144, 3, 32, 16, 0, 144, 145, 5, 4, 0, 0, 145, 146, 3, 34, 17,
		0, 146, 25, 1, 0, 0, 0, 147, 148, 3, 32, 16, 0, 148, 149, 5, 4, 0, 0, 149,
		150, 3, 48, 24, 0, 150, 27, 1, 0, 0, 0, 151, 152, 3, 32, 16, 0, 152, 153,
		5, 4, 0, 0, 153, 154, 3, 50, 25, 0, 154, 29, 1, 0, 0, 0, 155, 156, 3, 32,
		16, 0, 156, 31, 1, 0, 0, 0, 157, 158, 7, 1, 0, 0, 158, 33, 1, 0, 0, 0,
		159, 160, 7, 2, 0, 0, 160, 35, 1, 0, 0, 0, 161, 163, 5, 30, 0, 0, 162,
		164, 5, 40, 0, 0, 163, 162, 1, 0, 0, 0, 163, 164, 1, 0, 0, 0, 164, 165,
		1, 0, 0, 0, 165, 167, 3, 48, 24, 0, 166, 168, 5, 40, 0, 0, 167, 166, 1,
		0, 0, 0, 167, 168, 1, 0, 0, 0, 168, 169, 1, 0, 0, 0, 169, 170, 5, 31, 0,
		0, 170, 183, 1, 0, 0, 0, 171, 173, 5, 30, 0, 0, 172, 174, 5, 40, 0, 0,
		173, 172, 1, 0, 0, 0, 173, 174, 1, 0, 0, 0, 174, 175, 1, 0, 0, 0, 175,
		177, 3, 50, 25, 0, 176, 178, 5, 40, 0, 0, 177, 176, 1, 0, 0, 0, 177, 178,
		1, 0, 0, 0, 178, 179, 1, 0, 0, 0, 179, 180, 5, 31, 0, 0, 180, 183, 1, 0,
		0, 0, 181, 183, 3, 38, 19, 0, 182, 161, 1, 0, 0, 0, 182, 171, 1, 0, 0,
		0, 182, 181, 1, 0, 0, 0, 183, 37, 1, 0, 0, 0, 184, 185, 5, 30, 0, 0, 185,
		186, 5, 45, 0, 0, 186, 187, 7, 3, 0, 0, 187, 188, 5, 42, 0, 0, 188, 189,
		5, 31, 0, 0, 189, 39, 1, 0, 0, 0, 190, 192, 5, 30, 0, 0, 191, 193, 5, 40,
		0, 0, 192, 191, 1, 0, 0, 0, 192, 193, 1, 0, 0, 0, 193, 194, 1, 0, 0, 0,
		194, 196, 3, 48, 24, 0, 195, 197, 5, 40, 0, 0, 196, 195, 1, 0, 0, 0, 196,
		197, 1, 0, 0, 0, 197, 198, 1, 0, 0, 0, 198, 200, 5, 32, 0, 0, 199, 201,
		5, 40, 0, 0, 200, 199, 1, 0, 0, 0, 200, 201, 1, 0, 0, 0, 201, 202, 1, 0,
		0, 0, 202, 204, 3, 34, 17, 0, 203, 205, 5, 40, 0, 0, 204, 203, 1, 0, 0,
		0, 204, 205, 1, 0, 0, 0, 205, 206, 1, 0, 0, 0, 206, 207, 5, 31, 0, 0, 207,
		227, 1, 0, 0, 0, 208, 210, 5, 30, 0, 0, 209, 211, 5, 40, 0, 0, 210, 209,
		1, 0, 0, 0, 210, 211, 1, 0, 0, 0, 211, 212, 1, 0, 0, 0, 212, 214, 3, 50,
		25, 0, 213, 215, 5, 40, 0, 0, 214, 213, 1, 0, 0, 0, 214, 215, 1, 0, 0,
		0, 215, 216, 1, 0, 0, 0, 216, 218, 5, 32, 0, 0, 217, 219, 5, 40, 0, 0,
		218, 217, 1, 0, 0, 0, 218, 219, 1, 0, 0, 0, 219, 220, 1, 0, 0, 0, 220,
		222, 3, 34, 17, 0, 221, 223, 5, 40, 0, 0, 222, 221, 1, 0, 0, 0, 222, 223,
		1, 0, 0, 0, 223, 224, 1, 0, 0, 0, 224, 225, 5, 31, 0, 0, 225, 227, 1, 0,
		0, 0, 226, 190, 1, 0, 0, 0, 226, 208, 1, 0, 0, 0, 227, 41, 1, 0, 0, 0,
		228, 230, 3, 50, 25, 0, 229, 231, 5, 40, 0, 0, 230, 229, 1, 0, 0, 0, 230,
		231, 1, 0, 0, 0, 231, 232, 1, 0, 0, 0, 232, 234, 5, 34, 0, 0, 233, 235,
		5, 40, 0, 0, 234, 233, 1, 0, 0, 0, 234, 235, 1, 0, 0, 0, 235, 236, 1, 0,
		0, 0, 236, 237, 3, 48, 24, 0, 237, 43, 1, 0, 0, 0, 238, 239, 5, 35, 0,
		0, 239, 245, 3, 46, 23, 0, 240, 241, 5, 36, 0, 0, 241, 245, 3, 46, 23,
		0, 242, 243, 5, 37, 0, 0, 243, 245, 3, 46, 23, 0, 244, 238, 1, 0, 0, 0,
		244, 240, 1, 0, 0, 0, 244, 242, 1, 0, 0, 0, 245, 45, 1, 0, 0, 0, 246, 249,
		3, 48, 24, 0, 247, 249, 3, 50, 25, 0, 248, 246, 1, 0, 0, 0, 248, 247, 1,
		0, 0, 0, 249, 257, 1, 0, 0, 0, 250, 253, 5, 7, 0, 0, 251, 254, 3, 48, 24,
		0, 252, 254, 3, 50, 25, 0, 253, 251, 1, 0, 0, 0, 253, 252, 1, 0, 0, 0,
		254, 256, 1, 0, 0, 0, 255, 250, 1, 0, 0, 0, 256, 259, 1, 0, 0, 0, 257,
		255, 1, 0, 0, 0, 257, 258, 1, 0, 0, 0, 258, 47, 1, 0, 0, 0, 259, 257, 1,
		0, 0, 0, 260, 261, 7, 4, 0, 0, 261, 49, 1, 0, 0, 0, 262, 263, 5, 45, 0,
		0, 263, 51, 1, 0, 0, 0, 27, 58, 61, 66, 69, 71, 83, 99, 163, 167, 173,
		177, 182, 192, 196, 200, 204, 210, 214, 218, 222, 226, 230, 234, 244, 248,
		253, 257,
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
	AsmE8ParserLINE_COMMENT      = 38
	AsmE8ParserSUFFIX_WHITESPACE = 39
	AsmE8ParserWHITE_SPACE       = 40
	AsmE8ParserBINARY            = 41
	AsmE8ParserINT               = 42
	AsmE8ParserCHAR              = 43
	AsmE8ParserHEX               = 44
	AsmE8ParserSTR               = 45
)

// AsmE8Parser rules.
const (
	AsmE8ParserRULE_prog                  = 0
	AsmE8ParserRULE_instruction           = 1
	AsmE8ParserRULE_line                  = 2
	AsmE8ParserRULE_label                 = 3
	AsmE8ParserRULE_inst                  = 4
	AsmE8ParserRULE_inst_reg_reg          = 5
	AsmE8ParserRULE_inst_reg_imm          = 6
	AsmE8ParserRULE_inst_reg_imm_variable = 7
	AsmE8ParserRULE_inst_ptr_reg          = 8
	AsmE8ParserRULE_inst_reg_ptr          = 9
	AsmE8ParserRULE_inst_reg_ptr_offset   = 10
	AsmE8ParserRULE_inst_ptr_offset_reg   = 11
	AsmE8ParserRULE_inst_single_reg       = 12
	AsmE8ParserRULE_inst_single_imm       = 13
	AsmE8ParserRULE_inst_single_tag       = 14
	AsmE8ParserRULE_inst_single           = 15
	AsmE8ParserRULE_mnemonic              = 16
	AsmE8ParserRULE_reg                   = 17
	AsmE8ParserRULE_ptr                   = 18
	AsmE8ParserRULE_ptr_virtual_offset    = 19
	AsmE8ParserRULE_ptr_offset            = 20
	AsmE8ParserRULE_variable              = 21
	AsmE8ParserRULE_directives            = 22
	AsmE8ParserRULE_imm_list              = 23
	AsmE8ParserRULE_imm                   = 24
	AsmE8ParserRULE_tag                   = 25
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
		p.SetState(52)
		p.Instruction()
	}
	{
		p.SetState(53)
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

	p.SetState(71)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(56)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&35424890257454) != 0) {
			{
				p.SetState(55)
				p.Line()
			}

			p.SetState(58)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		p.SetState(61)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == AsmE8ParserLINE_COMMENT {
			{
				p.SetState(60)
				p.Match(AsmE8ParserLINE_COMMENT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(64)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&35424890257454) != 0) {
			{
				p.SetState(63)
				p.Line()
			}

			p.SetState(66)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		p.SetState(69)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == AsmE8ParserSUFFIX_WHITESPACE {
			{
				p.SetState(68)
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
	Variable() IVariableContext
	Inst() IInstContext
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

	p.SetState(83)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(73)
			p.Variable()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(74)
			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&14) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(75)
			p.Inst()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(76)
			p.Label()
		}
		{
			p.SetState(77)
			p.Match(AsmE8ParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(78)
			p.Directives()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(80)
			p.Directives()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(81)
			p.Label()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(82)
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
		p.SetState(85)
		p.Tag()
	}
	{
		p.SetState(86)
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
	Inst_reg_imm_variable() IInst_reg_imm_variableContext
	Inst_single_reg() IInst_single_regContext
	Inst_single_imm() IInst_single_immContext
	Inst_single_tag() IInst_single_tagContext
	Inst_single() IInst_singleContext
	Inst_ptr_reg() IInst_ptr_regContext
	Inst_reg_ptr() IInst_reg_ptrContext
	Inst_reg_ptr_offset() IInst_reg_ptr_offsetContext
	Inst_ptr_offset_reg() IInst_ptr_offset_regContext

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
	p.SetState(99)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(88)
			p.Inst_reg_reg()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(89)
			p.Inst_reg_imm()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(90)
			p.Inst_reg_imm_variable()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(91)
			p.Inst_single_reg()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(92)
			p.Inst_single_imm()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(93)
			p.Inst_single_tag()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(94)
			p.Inst_single()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(95)
			p.Inst_ptr_reg()
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(96)
			p.Inst_reg_ptr()
		}

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(97)
			p.Inst_reg_ptr_offset()
		}

	case 11:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(98)
			p.Inst_ptr_offset_reg()
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
		p.SetState(101)
		p.Mnemonic()
	}
	{
		p.SetState(102)
		p.Match(AsmE8ParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(103)
		p.Reg()
	}
	{
		p.SetState(104)
		p.Match(AsmE8ParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(105)
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
		p.SetState(107)
		p.Mnemonic()
	}
	{
		p.SetState(108)
		p.Match(AsmE8ParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(109)
		p.Reg()
	}
	{
		p.SetState(110)
		p.Match(AsmE8ParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(111)
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
	p.EnterRule(localctx, 14, AsmE8ParserRULE_inst_reg_imm_variable)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(113)
		p.Mnemonic()
	}
	{
		p.SetState(114)
		p.Match(AsmE8ParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(115)
		p.Reg()
	}
	{
		p.SetState(116)
		p.Match(AsmE8ParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(117)
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
	p.EnterRule(localctx, 16, AsmE8ParserRULE_inst_ptr_reg)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(119)
		p.Mnemonic()
	}
	{
		p.SetState(120)
		p.Match(AsmE8ParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(121)
		p.Ptr()
	}
	{
		p.SetState(122)
		p.Match(AsmE8ParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(123)
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
	p.EnterRule(localctx, 18, AsmE8ParserRULE_inst_reg_ptr)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(125)
		p.Mnemonic()
	}
	{
		p.SetState(126)
		p.Match(AsmE8ParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(127)
		p.Reg()
	}
	{
		p.SetState(128)
		p.Match(AsmE8ParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(129)
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
	p.EnterRule(localctx, 20, AsmE8ParserRULE_inst_reg_ptr_offset)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(131)
		p.Mnemonic()
	}
	{
		p.SetState(132)
		p.Match(AsmE8ParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(133)
		p.Reg()
	}
	{
		p.SetState(134)
		p.Match(AsmE8ParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(135)
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
	p.EnterRule(localctx, 22, AsmE8ParserRULE_inst_ptr_offset_reg)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(137)
		p.Mnemonic()
	}
	{
		p.SetState(138)
		p.Match(AsmE8ParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(139)
		p.Ptr_offset()
	}
	{
		p.SetState(140)
		p.Match(AsmE8ParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(141)
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
	p.EnterRule(localctx, 24, AsmE8ParserRULE_inst_single_reg)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(143)
		p.Mnemonic()
	}
	{
		p.SetState(144)
		p.Match(AsmE8ParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(145)
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
	p.EnterRule(localctx, 26, AsmE8ParserRULE_inst_single_imm)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(147)
		p.Mnemonic()
	}
	{
		p.SetState(148)
		p.Match(AsmE8ParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(149)
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
	p.EnterRule(localctx, 28, AsmE8ParserRULE_inst_single_tag)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(151)
		p.Mnemonic()
	}
	{
		p.SetState(152)
		p.Match(AsmE8ParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(153)
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
	p.EnterRule(localctx, 30, AsmE8ParserRULE_inst_single)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(155)
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
	p.EnterRule(localctx, 32, AsmE8ParserRULE_mnemonic)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(157)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&67108608) != 0) {
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
	p.EnterRule(localctx, 34, AsmE8ParserRULE_reg)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(159)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1006632960) != 0) {
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

// IPtrContext is an interface to support dynamic dispatch.
type IPtrContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Imm() IImmContext
	AllWHITE_SPACE() []antlr.TerminalNode
	WHITE_SPACE(i int) antlr.TerminalNode
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

func (s *PtrContext) AllWHITE_SPACE() []antlr.TerminalNode {
	return s.GetTokens(AsmE8ParserWHITE_SPACE)
}

func (s *PtrContext) WHITE_SPACE(i int) antlr.TerminalNode {
	return s.GetToken(AsmE8ParserWHITE_SPACE, i)
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
	p.EnterRule(localctx, 36, AsmE8ParserRULE_ptr)
	var _la int

	p.SetState(182)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(161)
			p.Match(AsmE8ParserT__29)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(163)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == AsmE8ParserWHITE_SPACE {
			{
				p.SetState(162)
				p.Match(AsmE8ParserWHITE_SPACE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(165)
			p.Imm()
		}
		p.SetState(167)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == AsmE8ParserWHITE_SPACE {
			{
				p.SetState(166)
				p.Match(AsmE8ParserWHITE_SPACE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(169)
			p.Match(AsmE8ParserT__30)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(171)
			p.Match(AsmE8ParserT__29)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(173)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == AsmE8ParserWHITE_SPACE {
			{
				p.SetState(172)
				p.Match(AsmE8ParserWHITE_SPACE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(175)
			p.Tag()
		}
		p.SetState(177)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == AsmE8ParserWHITE_SPACE {
			{
				p.SetState(176)
				p.Match(AsmE8ParserWHITE_SPACE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(179)
			p.Match(AsmE8ParserT__30)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(181)
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
	p.EnterRule(localctx, 38, AsmE8ParserRULE_ptr_virtual_offset)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(184)
		p.Match(AsmE8ParserT__29)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(185)
		p.Match(AsmE8ParserSTR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(186)
		_la = p.GetTokenStream().LA(1)

		if !(_la == AsmE8ParserT__31 || _la == AsmE8ParserT__32) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(187)
		p.Match(AsmE8ParserINT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(188)
		p.Match(AsmE8ParserT__30)
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
	AllWHITE_SPACE() []antlr.TerminalNode
	WHITE_SPACE(i int) antlr.TerminalNode
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

func (s *Ptr_offsetContext) AllWHITE_SPACE() []antlr.TerminalNode {
	return s.GetTokens(AsmE8ParserWHITE_SPACE)
}

func (s *Ptr_offsetContext) WHITE_SPACE(i int) antlr.TerminalNode {
	return s.GetToken(AsmE8ParserWHITE_SPACE, i)
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
	p.EnterRule(localctx, 40, AsmE8ParserRULE_ptr_offset)
	var _la int

	p.SetState(226)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 20, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(190)
			p.Match(AsmE8ParserT__29)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(192)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == AsmE8ParserWHITE_SPACE {
			{
				p.SetState(191)
				p.Match(AsmE8ParserWHITE_SPACE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(194)
			p.Imm()
		}
		p.SetState(196)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == AsmE8ParserWHITE_SPACE {
			{
				p.SetState(195)
				p.Match(AsmE8ParserWHITE_SPACE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(198)
			p.Match(AsmE8ParserT__31)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(200)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == AsmE8ParserWHITE_SPACE {
			{
				p.SetState(199)
				p.Match(AsmE8ParserWHITE_SPACE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(202)
			p.Reg()
		}
		p.SetState(204)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == AsmE8ParserWHITE_SPACE {
			{
				p.SetState(203)
				p.Match(AsmE8ParserWHITE_SPACE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(206)
			p.Match(AsmE8ParserT__30)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(208)
			p.Match(AsmE8ParserT__29)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(210)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == AsmE8ParserWHITE_SPACE {
			{
				p.SetState(209)
				p.Match(AsmE8ParserWHITE_SPACE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(212)
			p.Tag()
		}
		p.SetState(214)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == AsmE8ParserWHITE_SPACE {
			{
				p.SetState(213)
				p.Match(AsmE8ParserWHITE_SPACE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(216)
			p.Match(AsmE8ParserT__31)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(218)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == AsmE8ParserWHITE_SPACE {
			{
				p.SetState(217)
				p.Match(AsmE8ParserWHITE_SPACE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(220)
			p.Reg()
		}
		p.SetState(222)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == AsmE8ParserWHITE_SPACE {
			{
				p.SetState(221)
				p.Match(AsmE8ParserWHITE_SPACE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(224)
			p.Match(AsmE8ParserT__30)
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
	AllWHITE_SPACE() []antlr.TerminalNode
	WHITE_SPACE(i int) antlr.TerminalNode

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

func (s *VariableContext) AllWHITE_SPACE() []antlr.TerminalNode {
	return s.GetTokens(AsmE8ParserWHITE_SPACE)
}

func (s *VariableContext) WHITE_SPACE(i int) antlr.TerminalNode {
	return s.GetToken(AsmE8ParserWHITE_SPACE, i)
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
	p.EnterRule(localctx, 42, AsmE8ParserRULE_variable)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(228)
		p.Tag()
	}
	p.SetState(230)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == AsmE8ParserWHITE_SPACE {
		{
			p.SetState(229)
			p.Match(AsmE8ParserWHITE_SPACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(232)
		p.Match(AsmE8ParserT__33)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(234)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == AsmE8ParserWHITE_SPACE {
		{
			p.SetState(233)
			p.Match(AsmE8ParserWHITE_SPACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(236)
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
	p.EnterRule(localctx, 44, AsmE8ParserRULE_directives)
	p.SetState(244)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case AsmE8ParserT__34:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(238)
			p.Match(AsmE8ParserT__34)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(239)
			p.Imm_list()
		}

	case AsmE8ParserT__35:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(240)
			p.Match(AsmE8ParserT__35)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(241)
			p.Imm_list()
		}

	case AsmE8ParserT__36:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(242)
			p.Match(AsmE8ParserT__36)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(243)
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
	p.EnterRule(localctx, 46, AsmE8ParserRULE_imm_list)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(248)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case AsmE8ParserBINARY, AsmE8ParserINT, AsmE8ParserCHAR, AsmE8ParserHEX:
		{
			p.SetState(246)
			p.Imm()
		}

	case AsmE8ParserSTR:
		{
			p.SetState(247)
			p.Tag()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.SetState(257)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == AsmE8ParserT__6 {
		{
			p.SetState(250)
			p.Match(AsmE8ParserT__6)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(253)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case AsmE8ParserBINARY, AsmE8ParserINT, AsmE8ParserCHAR, AsmE8ParserHEX:
			{
				p.SetState(251)
				p.Imm()
			}

		case AsmE8ParserSTR:
			{
				p.SetState(252)
				p.Tag()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(259)
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
	p.EnterRule(localctx, 48, AsmE8ParserRULE_imm)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(260)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&32985348833280) != 0) {
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
	p.EnterRule(localctx, 50, AsmE8ParserRULE_tag)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(262)
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
