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
		"", "'\\t'", "'    '", "'  '", "'\\n'", "':'", "' '", "', '", "'mov'",
		"'add'", "'adc'", "'sub'", "'sbb'", "'cmp'", "'inc'", "'dec'", "'jmp'",
		"'jz'", "'jnz'", "'jsr'", "'rts'", "'brk'", "'nop'", "'push'", "'pop'",
		"'a'", "'b'", "'c'", "'d'", "'['", "']'", "'+'", "'='", "'.org '", "'.byte '",
		"'.word '", "','", "'''", "'' ''",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "LINE_COMMENT", "INT", "CHAR", "STR", "HEX", "BINARY",
		"WHITE_SPACE",
	}
	staticData.RuleNames = []string{
		"prog", "instruction", "line", "label", "inst", "inst_reg_reg", "inst_reg_inm",
		"inst_reg_inm_variable", "inst_ptr_reg", "inst_reg_ptr", "inst_reg_ptr_offset",
		"inst_ptr_offset_reg", "inst_single_reg", "inst_single_inm", "inst_single_tag",
		"inst_single", "mnemonic", "reg", "ptr", "ptr_offset", "variable", "directives",
		"inm_list", "inm", "tag",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 45, 246, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 1, 0, 1, 0, 1, 0,
		1, 1, 4, 1, 55, 8, 1, 11, 1, 12, 1, 56, 1, 1, 3, 1, 60, 8, 1, 1, 2, 1,
		2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 68, 8, 2, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4,
		1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 3, 4, 84, 8, 4, 1,
		5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1,
		7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1,
		9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1,
		10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 12,
		1, 13, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1,
		16, 1, 16, 1, 17, 1, 17, 1, 18, 1, 18, 3, 18, 148, 8, 18, 1, 18, 1, 18,
		3, 18, 152, 8, 18, 1, 18, 1, 18, 1, 18, 1, 18, 3, 18, 158, 8, 18, 1, 18,
		1, 18, 3, 18, 162, 8, 18, 1, 18, 1, 18, 3, 18, 166, 8, 18, 1, 19, 1, 19,
		3, 19, 170, 8, 19, 1, 19, 1, 19, 3, 19, 174, 8, 19, 1, 19, 1, 19, 3, 19,
		178, 8, 19, 1, 19, 1, 19, 3, 19, 182, 8, 19, 1, 19, 1, 19, 1, 19, 1, 19,
		3, 19, 188, 8, 19, 1, 19, 1, 19, 3, 19, 192, 8, 19, 1, 19, 1, 19, 3, 19,
		196, 8, 19, 1, 19, 1, 19, 3, 19, 200, 8, 19, 1, 19, 1, 19, 3, 19, 204,
		8, 19, 1, 20, 1, 20, 3, 20, 208, 8, 20, 1, 20, 1, 20, 3, 20, 212, 8, 20,
		1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 3, 21, 222, 8,
		21, 1, 22, 1, 22, 1, 22, 1, 22, 3, 22, 228, 8, 22, 5, 22, 230, 8, 22, 10,
		22, 12, 22, 233, 9, 22, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23,
		3, 23, 242, 8, 23, 1, 24, 1, 24, 1, 24, 0, 0, 25, 0, 2, 4, 6, 8, 10, 12,
		14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48,
		0, 4, 1, 0, 1, 3, 1, 0, 8, 24, 1, 0, 25, 28, 2, 0, 7, 7, 36, 36, 260, 0,
		50, 1, 0, 0, 0, 2, 54, 1, 0, 0, 0, 4, 67, 1, 0, 0, 0, 6, 69, 1, 0, 0, 0,
		8, 83, 1, 0, 0, 0, 10, 85, 1, 0, 0, 0, 12, 91, 1, 0, 0, 0, 14, 97, 1, 0,
		0, 0, 16, 103, 1, 0, 0, 0, 18, 109, 1, 0, 0, 0, 20, 115, 1, 0, 0, 0, 22,
		121, 1, 0, 0, 0, 24, 127, 1, 0, 0, 0, 26, 131, 1, 0, 0, 0, 28, 135, 1,
		0, 0, 0, 30, 139, 1, 0, 0, 0, 32, 141, 1, 0, 0, 0, 34, 143, 1, 0, 0, 0,
		36, 165, 1, 0, 0, 0, 38, 203, 1, 0, 0, 0, 40, 205, 1, 0, 0, 0, 42, 221,
		1, 0, 0, 0, 44, 223, 1, 0, 0, 0, 46, 241, 1, 0, 0, 0, 48, 243, 1, 0, 0,
		0, 50, 51, 3, 2, 1, 0, 51, 52, 5, 0, 0, 1, 52, 1, 1, 0, 0, 0, 53, 55, 3,
		4, 2, 0, 54, 53, 1, 0, 0, 0, 55, 56, 1, 0, 0, 0, 56, 54, 1, 0, 0, 0, 56,
		57, 1, 0, 0, 0, 57, 59, 1, 0, 0, 0, 58, 60, 5, 39, 0, 0, 59, 58, 1, 0,
		0, 0, 59, 60, 1, 0, 0, 0, 60, 3, 1, 0, 0, 0, 61, 68, 3, 40, 20, 0, 62,
		63, 7, 0, 0, 0, 63, 68, 3, 8, 4, 0, 64, 68, 3, 6, 3, 0, 65, 68, 3, 42,
		21, 0, 66, 68, 5, 4, 0, 0, 67, 61, 1, 0, 0, 0, 67, 62, 1, 0, 0, 0, 67,
		64, 1, 0, 0, 0, 67, 65, 1, 0, 0, 0, 67, 66, 1, 0, 0, 0, 68, 5, 1, 0, 0,
		0, 69, 70, 3, 48, 24, 0, 70, 71, 5, 5, 0, 0, 71, 7, 1, 0, 0, 0, 72, 84,
		3, 10, 5, 0, 73, 84, 3, 12, 6, 0, 74, 84, 3, 14, 7, 0, 75, 84, 3, 24, 12,
		0, 76, 84, 3, 26, 13, 0, 77, 84, 3, 28, 14, 0, 78, 84, 3, 30, 15, 0, 79,
		84, 3, 16, 8, 0, 80, 84, 3, 18, 9, 0, 81, 84, 3, 20, 10, 0, 82, 84, 3,
		22, 11, 0, 83, 72, 1, 0, 0, 0, 83, 73, 1, 0, 0, 0, 83, 74, 1, 0, 0, 0,
		83, 75, 1, 0, 0, 0, 83, 76, 1, 0, 0, 0, 83, 77, 1, 0, 0, 0, 83, 78, 1,
		0, 0, 0, 83, 79, 1, 0, 0, 0, 83, 80, 1, 0, 0, 0, 83, 81, 1, 0, 0, 0, 83,
		82, 1, 0, 0, 0, 84, 9, 1, 0, 0, 0, 85, 86, 3, 32, 16, 0, 86, 87, 5, 6,
		0, 0, 87, 88, 3, 34, 17, 0, 88, 89, 5, 7, 0, 0, 89, 90, 3, 34, 17, 0, 90,
		11, 1, 0, 0, 0, 91, 92, 3, 32, 16, 0, 92, 93, 5, 6, 0, 0, 93, 94, 3, 34,
		17, 0, 94, 95, 5, 7, 0, 0, 95, 96, 3, 46, 23, 0, 96, 13, 1, 0, 0, 0, 97,
		98, 3, 32, 16, 0, 98, 99, 5, 6, 0, 0, 99, 100, 3, 34, 17, 0, 100, 101,
		5, 7, 0, 0, 101, 102, 3, 48, 24, 0, 102, 15, 1, 0, 0, 0, 103, 104, 3, 32,
		16, 0, 104, 105, 5, 6, 0, 0, 105, 106, 3, 36, 18, 0, 106, 107, 5, 7, 0,
		0, 107, 108, 3, 34, 17, 0, 108, 17, 1, 0, 0, 0, 109, 110, 3, 32, 16, 0,
		110, 111, 5, 6, 0, 0, 111, 112, 3, 34, 17, 0, 112, 113, 5, 7, 0, 0, 113,
		114, 3, 36, 18, 0, 114, 19, 1, 0, 0, 0, 115, 116, 3, 32, 16, 0, 116, 117,
		5, 6, 0, 0, 117, 118, 3, 34, 17, 0, 118, 119, 5, 7, 0, 0, 119, 120, 3,
		38, 19, 0, 120, 21, 1, 0, 0, 0, 121, 122, 3, 32, 16, 0, 122, 123, 5, 6,
		0, 0, 123, 124, 3, 38, 19, 0, 124, 125, 5, 7, 0, 0, 125, 126, 3, 34, 17,
		0, 126, 23, 1, 0, 0, 0, 127, 128, 3, 32, 16, 0, 128, 129, 5, 6, 0, 0, 129,
		130, 3, 34, 17, 0, 130, 25, 1, 0, 0, 0, 131, 132, 3, 32, 16, 0, 132, 133,
		5, 6, 0, 0, 133, 134, 3, 46, 23, 0, 134, 27, 1, 0, 0, 0, 135, 136, 3, 32,
		16, 0, 136, 137, 5, 6, 0, 0, 137, 138, 3, 48, 24, 0, 138, 29, 1, 0, 0,
		0, 139, 140, 3, 32, 16, 0, 140, 31, 1, 0, 0, 0, 141, 142, 7, 1, 0, 0, 142,
		33, 1, 0, 0, 0, 143, 144, 7, 2, 0, 0, 144, 35, 1, 0, 0, 0, 145, 147, 5,
		29, 0, 0, 146, 148, 5, 45, 0, 0, 147, 146, 1, 0, 0, 0, 147, 148, 1, 0,
		0, 0, 148, 149, 1, 0, 0, 0, 149, 151, 3, 46, 23, 0, 150, 152, 5, 45, 0,
		0, 151, 150, 1, 0, 0, 0, 151, 152, 1, 0, 0, 0, 152, 153, 1, 0, 0, 0, 153,
		154, 5, 30, 0, 0, 154, 166, 1, 0, 0, 0, 155, 157, 5, 29, 0, 0, 156, 158,
		5, 45, 0, 0, 157, 156, 1, 0, 0, 0, 157, 158, 1, 0, 0, 0, 158, 159, 1, 0,
		0, 0, 159, 161, 3, 48, 24, 0, 160, 162, 5, 45, 0, 0, 161, 160, 1, 0, 0,
		0, 161, 162, 1, 0, 0, 0, 162, 163, 1, 0, 0, 0, 163, 164, 5, 30, 0, 0, 164,
		166, 1, 0, 0, 0, 165, 145, 1, 0, 0, 0, 165, 155, 1, 0, 0, 0, 166, 37, 1,
		0, 0, 0, 167, 169, 5, 29, 0, 0, 168, 170, 5, 45, 0, 0, 169, 168, 1, 0,
		0, 0, 169, 170, 1, 0, 0, 0, 170, 171, 1, 0, 0, 0, 171, 173, 3, 46, 23,
		0, 172, 174, 5, 45, 0, 0, 173, 172, 1, 0, 0, 0, 173, 174, 1, 0, 0, 0, 174,
		175, 1, 0, 0, 0, 175, 177, 5, 31, 0, 0, 176, 178, 5, 45, 0, 0, 177, 176,
		1, 0, 0, 0, 177, 178, 1, 0, 0, 0, 178, 179, 1, 0, 0, 0, 179, 181, 3, 34,
		17, 0, 180, 182, 5, 45, 0, 0, 181, 180, 1, 0, 0, 0, 181, 182, 1, 0, 0,
		0, 182, 183, 1, 0, 0, 0, 183, 184, 5, 30, 0, 0, 184, 204, 1, 0, 0, 0, 185,
		187, 5, 29, 0, 0, 186, 188, 5, 45, 0, 0, 187, 186, 1, 0, 0, 0, 187, 188,
		1, 0, 0, 0, 188, 189, 1, 0, 0, 0, 189, 191, 3, 48, 24, 0, 190, 192, 5,
		45, 0, 0, 191, 190, 1, 0, 0, 0, 191, 192, 1, 0, 0, 0, 192, 193, 1, 0, 0,
		0, 193, 195, 5, 31, 0, 0, 194, 196, 5, 45, 0, 0, 195, 194, 1, 0, 0, 0,
		195, 196, 1, 0, 0, 0, 196, 197, 1, 0, 0, 0, 197, 199, 3, 34, 17, 0, 198,
		200, 5, 45, 0, 0, 199, 198, 1, 0, 0, 0, 199, 200, 1, 0, 0, 0, 200, 201,
		1, 0, 0, 0, 201, 202, 5, 30, 0, 0, 202, 204, 1, 0, 0, 0, 203, 167, 1, 0,
		0, 0, 203, 185, 1, 0, 0, 0, 204, 39, 1, 0, 0, 0, 205, 207, 3, 48, 24, 0,
		206, 208, 5, 45, 0, 0, 207, 206, 1, 0, 0, 0, 207, 208, 1, 0, 0, 0, 208,
		209, 1, 0, 0, 0, 209, 211, 5, 32, 0, 0, 210, 212, 5, 45, 0, 0, 211, 210,
		1, 0, 0, 0, 211, 212, 1, 0, 0, 0, 212, 213, 1, 0, 0, 0, 213, 214, 3, 46,
		23, 0, 214, 41, 1, 0, 0, 0, 215, 216, 5, 33, 0, 0, 216, 222, 3, 44, 22,
		0, 217, 218, 5, 34, 0, 0, 218, 222, 3, 44, 22, 0, 219, 220, 5, 35, 0, 0,
		220, 222, 3, 44, 22, 0, 221, 215, 1, 0, 0, 0, 221, 217, 1, 0, 0, 0, 221,
		219, 1, 0, 0, 0, 222, 43, 1, 0, 0, 0, 223, 231, 3, 46, 23, 0, 224, 227,
		7, 3, 0, 0, 225, 228, 3, 46, 23, 0, 226, 228, 3, 48, 24, 0, 227, 225, 1,
		0, 0, 0, 227, 226, 1, 0, 0, 0, 228, 230, 1, 0, 0, 0, 229, 224, 1, 0, 0,
		0, 230, 233, 1, 0, 0, 0, 231, 229, 1, 0, 0, 0, 231, 232, 1, 0, 0, 0, 232,
		45, 1, 0, 0, 0, 233, 231, 1, 0, 0, 0, 234, 242, 5, 40, 0, 0, 235, 242,
		5, 43, 0, 0, 236, 242, 5, 44, 0, 0, 237, 238, 5, 37, 0, 0, 238, 239, 5,
		41, 0, 0, 239, 242, 5, 37, 0, 0, 240, 242, 5, 38, 0, 0, 241, 234, 1, 0,
		0, 0, 241, 235, 1, 0, 0, 0, 241, 236, 1, 0, 0, 0, 241, 237, 1, 0, 0, 0,
		241, 240, 1, 0, 0, 0, 242, 47, 1, 0, 0, 0, 243, 244, 5, 42, 0, 0, 244,
		49, 1, 0, 0, 0, 24, 56, 59, 67, 83, 147, 151, 157, 161, 165, 169, 173,
		177, 181, 187, 191, 195, 199, 203, 207, 211, 221, 227, 231, 241,
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
	AsmE8ParserEOF          = antlr.TokenEOF
	AsmE8ParserT__0         = 1
	AsmE8ParserT__1         = 2
	AsmE8ParserT__2         = 3
	AsmE8ParserT__3         = 4
	AsmE8ParserT__4         = 5
	AsmE8ParserT__5         = 6
	AsmE8ParserT__6         = 7
	AsmE8ParserT__7         = 8
	AsmE8ParserT__8         = 9
	AsmE8ParserT__9         = 10
	AsmE8ParserT__10        = 11
	AsmE8ParserT__11        = 12
	AsmE8ParserT__12        = 13
	AsmE8ParserT__13        = 14
	AsmE8ParserT__14        = 15
	AsmE8ParserT__15        = 16
	AsmE8ParserT__16        = 17
	AsmE8ParserT__17        = 18
	AsmE8ParserT__18        = 19
	AsmE8ParserT__19        = 20
	AsmE8ParserT__20        = 21
	AsmE8ParserT__21        = 22
	AsmE8ParserT__22        = 23
	AsmE8ParserT__23        = 24
	AsmE8ParserT__24        = 25
	AsmE8ParserT__25        = 26
	AsmE8ParserT__26        = 27
	AsmE8ParserT__27        = 28
	AsmE8ParserT__28        = 29
	AsmE8ParserT__29        = 30
	AsmE8ParserT__30        = 31
	AsmE8ParserT__31        = 32
	AsmE8ParserT__32        = 33
	AsmE8ParserT__33        = 34
	AsmE8ParserT__34        = 35
	AsmE8ParserT__35        = 36
	AsmE8ParserT__36        = 37
	AsmE8ParserT__37        = 38
	AsmE8ParserLINE_COMMENT = 39
	AsmE8ParserINT          = 40
	AsmE8ParserCHAR         = 41
	AsmE8ParserSTR          = 42
	AsmE8ParserHEX          = 43
	AsmE8ParserBINARY       = 44
	AsmE8ParserWHITE_SPACE  = 45
)

// AsmE8Parser rules.
const (
	AsmE8ParserRULE_prog                  = 0
	AsmE8ParserRULE_instruction           = 1
	AsmE8ParserRULE_line                  = 2
	AsmE8ParserRULE_label                 = 3
	AsmE8ParserRULE_inst                  = 4
	AsmE8ParserRULE_inst_reg_reg          = 5
	AsmE8ParserRULE_inst_reg_inm          = 6
	AsmE8ParserRULE_inst_reg_inm_variable = 7
	AsmE8ParserRULE_inst_ptr_reg          = 8
	AsmE8ParserRULE_inst_reg_ptr          = 9
	AsmE8ParserRULE_inst_reg_ptr_offset   = 10
	AsmE8ParserRULE_inst_ptr_offset_reg   = 11
	AsmE8ParserRULE_inst_single_reg       = 12
	AsmE8ParserRULE_inst_single_inm       = 13
	AsmE8ParserRULE_inst_single_tag       = 14
	AsmE8ParserRULE_inst_single           = 15
	AsmE8ParserRULE_mnemonic              = 16
	AsmE8ParserRULE_reg                   = 17
	AsmE8ParserRULE_ptr                   = 18
	AsmE8ParserRULE_ptr_offset            = 19
	AsmE8ParserRULE_variable              = 20
	AsmE8ParserRULE_directives            = 21
	AsmE8ParserRULE_inm_list              = 22
	AsmE8ParserRULE_inm                   = 23
	AsmE8ParserRULE_tag                   = 24
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
		p.SetState(50)
		p.Instruction()
	}
	{
		p.SetState(51)
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

	p.EnterOuterAlt(localctx, 1)
	p.SetState(54)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4458176053278) != 0) {
		{
			p.SetState(53)
			p.Line()
		}

		p.SetState(56)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(59)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == AsmE8ParserLINE_COMMENT {
		{
			p.SetState(58)
			p.Match(AsmE8ParserLINE_COMMENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
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

	p.SetState(67)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(61)
			p.Variable()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(62)
			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&14) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(63)
			p.Inst()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(64)
			p.Label()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(65)
			p.Directives()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(66)
			p.Match(AsmE8ParserT__3)
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
		p.SetState(69)
		p.Tag()
	}
	{
		p.SetState(70)
		p.Match(AsmE8ParserT__4)
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
	Inst_reg_inm() IInst_reg_inmContext
	Inst_reg_inm_variable() IInst_reg_inm_variableContext
	Inst_single_reg() IInst_single_regContext
	Inst_single_inm() IInst_single_inmContext
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

func (s *InstContext) Inst_reg_inm() IInst_reg_inmContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInst_reg_inmContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInst_reg_inmContext)
}

func (s *InstContext) Inst_reg_inm_variable() IInst_reg_inm_variableContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInst_reg_inm_variableContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInst_reg_inm_variableContext)
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

func (s *InstContext) Inst_single_inm() IInst_single_inmContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInst_single_inmContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInst_single_inmContext)
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
	p.SetState(83)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(72)
			p.Inst_reg_reg()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(73)
			p.Inst_reg_inm()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(74)
			p.Inst_reg_inm_variable()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(75)
			p.Inst_single_reg()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(76)
			p.Inst_single_inm()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(77)
			p.Inst_single_tag()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(78)
			p.Inst_single()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(79)
			p.Inst_ptr_reg()
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(80)
			p.Inst_reg_ptr()
		}

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(81)
			p.Inst_reg_ptr_offset()
		}

	case 11:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(82)
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
		p.SetState(85)
		p.Mnemonic()
	}
	{
		p.SetState(86)
		p.Match(AsmE8ParserT__5)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(87)
		p.Reg()
	}
	{
		p.SetState(88)
		p.Match(AsmE8ParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(89)
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

// IInst_reg_inmContext is an interface to support dynamic dispatch.
type IInst_reg_inmContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Mnemonic() IMnemonicContext
	Reg() IRegContext
	Inm() IInmContext

	// IsInst_reg_inmContext differentiates from other interfaces.
	IsInst_reg_inmContext()
}

type Inst_reg_inmContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInst_reg_inmContext() *Inst_reg_inmContext {
	var p = new(Inst_reg_inmContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsmE8ParserRULE_inst_reg_inm
	return p
}

func InitEmptyInst_reg_inmContext(p *Inst_reg_inmContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsmE8ParserRULE_inst_reg_inm
}

func (*Inst_reg_inmContext) IsInst_reg_inmContext() {}

func NewInst_reg_inmContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Inst_reg_inmContext {
	var p = new(Inst_reg_inmContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AsmE8ParserRULE_inst_reg_inm

	return p
}

func (s *Inst_reg_inmContext) GetParser() antlr.Parser { return s.parser }

func (s *Inst_reg_inmContext) Mnemonic() IMnemonicContext {
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

func (s *Inst_reg_inmContext) Reg() IRegContext {
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

func (s *Inst_reg_inmContext) Inm() IInmContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInmContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInmContext)
}

func (s *Inst_reg_inmContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Inst_reg_inmContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Inst_reg_inmContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsmE8Listener); ok {
		listenerT.EnterInst_reg_inm(s)
	}
}

func (s *Inst_reg_inmContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsmE8Listener); ok {
		listenerT.ExitInst_reg_inm(s)
	}
}

func (p *AsmE8Parser) Inst_reg_inm() (localctx IInst_reg_inmContext) {
	localctx = NewInst_reg_inmContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, AsmE8ParserRULE_inst_reg_inm)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(91)
		p.Mnemonic()
	}
	{
		p.SetState(92)
		p.Match(AsmE8ParserT__5)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(93)
		p.Reg()
	}
	{
		p.SetState(94)
		p.Match(AsmE8ParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(95)
		p.Inm()
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

// IInst_reg_inm_variableContext is an interface to support dynamic dispatch.
type IInst_reg_inm_variableContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Mnemonic() IMnemonicContext
	Reg() IRegContext
	Tag() ITagContext

	// IsInst_reg_inm_variableContext differentiates from other interfaces.
	IsInst_reg_inm_variableContext()
}

type Inst_reg_inm_variableContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInst_reg_inm_variableContext() *Inst_reg_inm_variableContext {
	var p = new(Inst_reg_inm_variableContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsmE8ParserRULE_inst_reg_inm_variable
	return p
}

func InitEmptyInst_reg_inm_variableContext(p *Inst_reg_inm_variableContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsmE8ParserRULE_inst_reg_inm_variable
}

func (*Inst_reg_inm_variableContext) IsInst_reg_inm_variableContext() {}

func NewInst_reg_inm_variableContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Inst_reg_inm_variableContext {
	var p = new(Inst_reg_inm_variableContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AsmE8ParserRULE_inst_reg_inm_variable

	return p
}

func (s *Inst_reg_inm_variableContext) GetParser() antlr.Parser { return s.parser }

func (s *Inst_reg_inm_variableContext) Mnemonic() IMnemonicContext {
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

func (s *Inst_reg_inm_variableContext) Reg() IRegContext {
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

func (s *Inst_reg_inm_variableContext) Tag() ITagContext {
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

func (s *Inst_reg_inm_variableContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Inst_reg_inm_variableContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Inst_reg_inm_variableContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsmE8Listener); ok {
		listenerT.EnterInst_reg_inm_variable(s)
	}
}

func (s *Inst_reg_inm_variableContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsmE8Listener); ok {
		listenerT.ExitInst_reg_inm_variable(s)
	}
}

func (p *AsmE8Parser) Inst_reg_inm_variable() (localctx IInst_reg_inm_variableContext) {
	localctx = NewInst_reg_inm_variableContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, AsmE8ParserRULE_inst_reg_inm_variable)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(97)
		p.Mnemonic()
	}
	{
		p.SetState(98)
		p.Match(AsmE8ParserT__5)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(99)
		p.Reg()
	}
	{
		p.SetState(100)
		p.Match(AsmE8ParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(101)
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
		p.SetState(103)
		p.Mnemonic()
	}
	{
		p.SetState(104)
		p.Match(AsmE8ParserT__5)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(105)
		p.Ptr()
	}
	{
		p.SetState(106)
		p.Match(AsmE8ParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(107)
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
		p.SetState(109)
		p.Mnemonic()
	}
	{
		p.SetState(110)
		p.Match(AsmE8ParserT__5)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(111)
		p.Reg()
	}
	{
		p.SetState(112)
		p.Match(AsmE8ParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(113)
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
		p.SetState(115)
		p.Mnemonic()
	}
	{
		p.SetState(116)
		p.Match(AsmE8ParserT__5)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(117)
		p.Reg()
	}
	{
		p.SetState(118)
		p.Match(AsmE8ParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(119)
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
		p.SetState(121)
		p.Mnemonic()
	}
	{
		p.SetState(122)
		p.Match(AsmE8ParserT__5)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(123)
		p.Ptr_offset()
	}
	{
		p.SetState(124)
		p.Match(AsmE8ParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(125)
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
		p.SetState(127)
		p.Mnemonic()
	}
	{
		p.SetState(128)
		p.Match(AsmE8ParserT__5)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(129)
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

// IInst_single_inmContext is an interface to support dynamic dispatch.
type IInst_single_inmContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Mnemonic() IMnemonicContext
	Inm() IInmContext

	// IsInst_single_inmContext differentiates from other interfaces.
	IsInst_single_inmContext()
}

type Inst_single_inmContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInst_single_inmContext() *Inst_single_inmContext {
	var p = new(Inst_single_inmContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsmE8ParserRULE_inst_single_inm
	return p
}

func InitEmptyInst_single_inmContext(p *Inst_single_inmContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsmE8ParserRULE_inst_single_inm
}

func (*Inst_single_inmContext) IsInst_single_inmContext() {}

func NewInst_single_inmContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Inst_single_inmContext {
	var p = new(Inst_single_inmContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AsmE8ParserRULE_inst_single_inm

	return p
}

func (s *Inst_single_inmContext) GetParser() antlr.Parser { return s.parser }

func (s *Inst_single_inmContext) Mnemonic() IMnemonicContext {
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

func (s *Inst_single_inmContext) Inm() IInmContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInmContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInmContext)
}

func (s *Inst_single_inmContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Inst_single_inmContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Inst_single_inmContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsmE8Listener); ok {
		listenerT.EnterInst_single_inm(s)
	}
}

func (s *Inst_single_inmContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsmE8Listener); ok {
		listenerT.ExitInst_single_inm(s)
	}
}

func (p *AsmE8Parser) Inst_single_inm() (localctx IInst_single_inmContext) {
	localctx = NewInst_single_inmContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, AsmE8ParserRULE_inst_single_inm)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(131)
		p.Mnemonic()
	}
	{
		p.SetState(132)
		p.Match(AsmE8ParserT__5)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(133)
		p.Inm()
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
		p.SetState(135)
		p.Mnemonic()
	}
	{
		p.SetState(136)
		p.Match(AsmE8ParserT__5)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(137)
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
		p.SetState(139)
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
		p.SetState(141)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&33554176) != 0) {
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
		p.SetState(143)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&503316480) != 0) {
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
	Inm() IInmContext
	AllWHITE_SPACE() []antlr.TerminalNode
	WHITE_SPACE(i int) antlr.TerminalNode
	Tag() ITagContext

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

func (s *PtrContext) Inm() IInmContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInmContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInmContext)
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

	p.SetState(165)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(145)
			p.Match(AsmE8ParserT__28)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(147)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == AsmE8ParserWHITE_SPACE {
			{
				p.SetState(146)
				p.Match(AsmE8ParserWHITE_SPACE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(149)
			p.Inm()
		}
		p.SetState(151)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == AsmE8ParserWHITE_SPACE {
			{
				p.SetState(150)
				p.Match(AsmE8ParserWHITE_SPACE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(153)
			p.Match(AsmE8ParserT__29)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(155)
			p.Match(AsmE8ParserT__28)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(157)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == AsmE8ParserWHITE_SPACE {
			{
				p.SetState(156)
				p.Match(AsmE8ParserWHITE_SPACE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(159)
			p.Tag()
		}
		p.SetState(161)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == AsmE8ParserWHITE_SPACE {
			{
				p.SetState(160)
				p.Match(AsmE8ParserWHITE_SPACE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(163)
			p.Match(AsmE8ParserT__29)
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

// IPtr_offsetContext is an interface to support dynamic dispatch.
type IPtr_offsetContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Inm() IInmContext
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

func (s *Ptr_offsetContext) Inm() IInmContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInmContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInmContext)
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
	p.EnterRule(localctx, 38, AsmE8ParserRULE_ptr_offset)
	var _la int

	p.SetState(203)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 17, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(167)
			p.Match(AsmE8ParserT__28)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(169)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == AsmE8ParserWHITE_SPACE {
			{
				p.SetState(168)
				p.Match(AsmE8ParserWHITE_SPACE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(171)
			p.Inm()
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
			p.Match(AsmE8ParserT__30)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
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
			p.Reg()
		}
		p.SetState(181)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == AsmE8ParserWHITE_SPACE {
			{
				p.SetState(180)
				p.Match(AsmE8ParserWHITE_SPACE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(183)
			p.Match(AsmE8ParserT__29)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(185)
			p.Match(AsmE8ParserT__28)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(187)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == AsmE8ParserWHITE_SPACE {
			{
				p.SetState(186)
				p.Match(AsmE8ParserWHITE_SPACE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(189)
			p.Tag()
		}
		p.SetState(191)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == AsmE8ParserWHITE_SPACE {
			{
				p.SetState(190)
				p.Match(AsmE8ParserWHITE_SPACE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(193)
			p.Match(AsmE8ParserT__30)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(195)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == AsmE8ParserWHITE_SPACE {
			{
				p.SetState(194)
				p.Match(AsmE8ParserWHITE_SPACE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(197)
			p.Reg()
		}
		p.SetState(199)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == AsmE8ParserWHITE_SPACE {
			{
				p.SetState(198)
				p.Match(AsmE8ParserWHITE_SPACE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(201)
			p.Match(AsmE8ParserT__29)
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
	Inm() IInmContext
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

func (s *VariableContext) Inm() IInmContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInmContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInmContext)
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
	p.EnterRule(localctx, 40, AsmE8ParserRULE_variable)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(205)
		p.Tag()
	}
	p.SetState(207)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == AsmE8ParserWHITE_SPACE {
		{
			p.SetState(206)
			p.Match(AsmE8ParserWHITE_SPACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(209)
		p.Match(AsmE8ParserT__31)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(211)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == AsmE8ParserWHITE_SPACE {
		{
			p.SetState(210)
			p.Match(AsmE8ParserWHITE_SPACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(213)
		p.Inm()
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
	Inm_list() IInm_listContext

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

func (s *DirectivesContext) Inm_list() IInm_listContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInm_listContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInm_listContext)
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
	p.EnterRule(localctx, 42, AsmE8ParserRULE_directives)
	p.SetState(221)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case AsmE8ParserT__32:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(215)
			p.Match(AsmE8ParserT__32)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(216)
			p.Inm_list()
		}

	case AsmE8ParserT__33:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(217)
			p.Match(AsmE8ParserT__33)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(218)
			p.Inm_list()
		}

	case AsmE8ParserT__34:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(219)
			p.Match(AsmE8ParserT__34)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(220)
			p.Inm_list()
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

// IInm_listContext is an interface to support dynamic dispatch.
type IInm_listContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllInm() []IInmContext
	Inm(i int) IInmContext
	AllTag() []ITagContext
	Tag(i int) ITagContext

	// IsInm_listContext differentiates from other interfaces.
	IsInm_listContext()
}

type Inm_listContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInm_listContext() *Inm_listContext {
	var p = new(Inm_listContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsmE8ParserRULE_inm_list
	return p
}

func InitEmptyInm_listContext(p *Inm_listContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsmE8ParserRULE_inm_list
}

func (*Inm_listContext) IsInm_listContext() {}

func NewInm_listContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Inm_listContext {
	var p = new(Inm_listContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AsmE8ParserRULE_inm_list

	return p
}

func (s *Inm_listContext) GetParser() antlr.Parser { return s.parser }

func (s *Inm_listContext) AllInm() []IInmContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IInmContext); ok {
			len++
		}
	}

	tst := make([]IInmContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IInmContext); ok {
			tst[i] = t.(IInmContext)
			i++
		}
	}

	return tst
}

func (s *Inm_listContext) Inm(i int) IInmContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInmContext); ok {
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

	return t.(IInmContext)
}

func (s *Inm_listContext) AllTag() []ITagContext {
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

func (s *Inm_listContext) Tag(i int) ITagContext {
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

func (s *Inm_listContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Inm_listContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Inm_listContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsmE8Listener); ok {
		listenerT.EnterInm_list(s)
	}
}

func (s *Inm_listContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsmE8Listener); ok {
		listenerT.ExitInm_list(s)
	}
}

func (p *AsmE8Parser) Inm_list() (localctx IInm_listContext) {
	localctx = NewInm_listContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, AsmE8ParserRULE_inm_list)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(223)
		p.Inm()
	}
	p.SetState(231)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == AsmE8ParserT__6 || _la == AsmE8ParserT__35 {
		{
			p.SetState(224)
			_la = p.GetTokenStream().LA(1)

			if !(_la == AsmE8ParserT__6 || _la == AsmE8ParserT__35) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		p.SetState(227)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case AsmE8ParserT__36, AsmE8ParserT__37, AsmE8ParserINT, AsmE8ParserHEX, AsmE8ParserBINARY:
			{
				p.SetState(225)
				p.Inm()
			}

		case AsmE8ParserSTR:
			{
				p.SetState(226)
				p.Tag()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(233)
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

// IInmContext is an interface to support dynamic dispatch.
type IInmContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	INT() antlr.TerminalNode
	HEX() antlr.TerminalNode
	BINARY() antlr.TerminalNode
	CHAR() antlr.TerminalNode

	// IsInmContext differentiates from other interfaces.
	IsInmContext()
}

type InmContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInmContext() *InmContext {
	var p = new(InmContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsmE8ParserRULE_inm
	return p
}

func InitEmptyInmContext(p *InmContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsmE8ParserRULE_inm
}

func (*InmContext) IsInmContext() {}

func NewInmContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InmContext {
	var p = new(InmContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AsmE8ParserRULE_inm

	return p
}

func (s *InmContext) GetParser() antlr.Parser { return s.parser }

func (s *InmContext) INT() antlr.TerminalNode {
	return s.GetToken(AsmE8ParserINT, 0)
}

func (s *InmContext) HEX() antlr.TerminalNode {
	return s.GetToken(AsmE8ParserHEX, 0)
}

func (s *InmContext) BINARY() antlr.TerminalNode {
	return s.GetToken(AsmE8ParserBINARY, 0)
}

func (s *InmContext) CHAR() antlr.TerminalNode {
	return s.GetToken(AsmE8ParserCHAR, 0)
}

func (s *InmContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InmContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InmContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsmE8Listener); ok {
		listenerT.EnterInm(s)
	}
}

func (s *InmContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsmE8Listener); ok {
		listenerT.ExitInm(s)
	}
}

func (p *AsmE8Parser) Inm() (localctx IInmContext) {
	localctx = NewInmContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, AsmE8ParserRULE_inm)
	p.SetState(241)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case AsmE8ParserINT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(234)
			p.Match(AsmE8ParserINT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case AsmE8ParserHEX:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(235)
			p.Match(AsmE8ParserHEX)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case AsmE8ParserBINARY:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(236)
			p.Match(AsmE8ParserBINARY)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case AsmE8ParserT__36:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(237)
			p.Match(AsmE8ParserT__36)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(238)
			p.Match(AsmE8ParserCHAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(239)
			p.Match(AsmE8ParserT__36)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case AsmE8ParserT__37:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(240)
			p.Match(AsmE8ParserT__37)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
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
	p.EnterRule(localctx, 48, AsmE8ParserRULE_tag)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(243)
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
