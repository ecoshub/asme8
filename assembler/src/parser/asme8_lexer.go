// Code generated from AsmE8.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type AsmE8Lexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var AsmE8LexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func asme8lexerLexerInit() {
	staticData := &AsmE8LexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "'\\t'", "'    '", "'  '", "' '", "'\\n'", "':'", "', '", "'mov'",
		"'add'", "'adc'", "'sub'", "'sbb'", "'cmp'", "'inc'", "'dec'", "'xor'",
		"'and'", "'or'", "'not'", "'shl'", "'shr'", "'rol'", "'ror'", "'jmp'",
		"'jz'", "'jnz'", "'jc'", "'js'", "'jns'", "'jsr'", "'rts'", "'brk'",
		"'nop'", "'push'", "'pop'", "'clc'", "'a'", "'b'", "'c'", "'d'", "'sp'",
		"'['", "']'", "'+'", "'-'", "'='", "'.byte '", "'.word '", "'.resb '",
		"'.org '", "'.segment '", "'\"'", "'global '", "'extern '",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "LINE_COMMENT", "SUFFIX_WHITESPACE", "WHITE_SPACE",
		"BINARY", "INT", "CHAR", "HEX", "STR",
	}
	staticData.RuleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
		"T__9", "T__10", "T__11", "T__12", "T__13", "T__14", "T__15", "T__16",
		"T__17", "T__18", "T__19", "T__20", "T__21", "T__22", "T__23", "T__24",
		"T__25", "T__26", "T__27", "T__28", "T__29", "T__30", "T__31", "T__32",
		"T__33", "T__34", "T__35", "T__36", "T__37", "T__38", "T__39", "T__40",
		"T__41", "T__42", "T__43", "T__44", "T__45", "T__46", "T__47", "T__48",
		"T__49", "T__50", "T__51", "T__52", "T__53", "LINE_COMMENT", "SUFFIX_WHITESPACE",
		"WHITE_SPACE", "BINARY", "INT", "CHAR", "HEX", "STR",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 62, 405, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36,
		7, 36, 2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7,
		41, 2, 42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 2, 46, 7, 46,
		2, 47, 7, 47, 2, 48, 7, 48, 2, 49, 7, 49, 2, 50, 7, 50, 2, 51, 7, 51, 2,
		52, 7, 52, 2, 53, 7, 53, 2, 54, 7, 54, 2, 55, 7, 55, 2, 56, 7, 56, 2, 57,
		7, 57, 2, 58, 7, 58, 2, 59, 7, 59, 2, 60, 7, 60, 2, 61, 7, 61, 1, 0, 1,
		0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1,
		4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1,
		8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 11, 1,
		11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 1, 13,
		1, 14, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 1, 15, 1, 16, 1, 16, 1,
		16, 1, 16, 1, 17, 1, 17, 1, 17, 1, 18, 1, 18, 1, 18, 1, 18, 1, 19, 1, 19,
		1, 19, 1, 19, 1, 20, 1, 20, 1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 1, 21, 1,
		22, 1, 22, 1, 22, 1, 22, 1, 23, 1, 23, 1, 23, 1, 23, 1, 24, 1, 24, 1, 24,
		1, 25, 1, 25, 1, 25, 1, 25, 1, 26, 1, 26, 1, 26, 1, 27, 1, 27, 1, 27, 1,
		28, 1, 28, 1, 28, 1, 28, 1, 29, 1, 29, 1, 29, 1, 29, 1, 30, 1, 30, 1, 30,
		1, 30, 1, 31, 1, 31, 1, 31, 1, 31, 1, 32, 1, 32, 1, 32, 1, 32, 1, 33, 1,
		33, 1, 33, 1, 33, 1, 33, 1, 34, 1, 34, 1, 34, 1, 34, 1, 35, 1, 35, 1, 35,
		1, 35, 1, 36, 1, 36, 1, 37, 1, 37, 1, 38, 1, 38, 1, 39, 1, 39, 1, 40, 1,
		40, 1, 40, 1, 41, 1, 41, 1, 42, 1, 42, 1, 43, 1, 43, 1, 44, 1, 44, 1, 45,
		1, 45, 1, 46, 1, 46, 1, 46, 1, 46, 1, 46, 1, 46, 1, 46, 1, 47, 1, 47, 1,
		47, 1, 47, 1, 47, 1, 47, 1, 47, 1, 48, 1, 48, 1, 48, 1, 48, 1, 48, 1, 48,
		1, 48, 1, 49, 1, 49, 1, 49, 1, 49, 1, 49, 1, 49, 1, 50, 1, 50, 1, 50, 1,
		50, 1, 50, 1, 50, 1, 50, 1, 50, 1, 50, 1, 50, 1, 51, 1, 51, 1, 52, 1, 52,
		1, 52, 1, 52, 1, 52, 1, 52, 1, 52, 1, 52, 1, 53, 1, 53, 1, 53, 1, 53, 1,
		53, 1, 53, 1, 53, 1, 53, 1, 54, 5, 54, 335, 8, 54, 10, 54, 12, 54, 338,
		9, 54, 1, 54, 1, 54, 5, 54, 342, 8, 54, 10, 54, 12, 54, 345, 9, 54, 1,
		54, 1, 54, 1, 55, 5, 55, 350, 8, 55, 10, 55, 12, 55, 353, 9, 55, 1, 55,
		1, 55, 1, 55, 1, 55, 1, 56, 4, 56, 360, 8, 56, 11, 56, 12, 56, 361, 1,
		57, 1, 57, 1, 57, 1, 57, 4, 57, 368, 8, 57, 11, 57, 12, 57, 369, 1, 58,
		4, 58, 373, 8, 58, 11, 58, 12, 58, 374, 1, 59, 1, 59, 1, 59, 1, 59, 1,
		59, 1, 59, 1, 59, 1, 59, 1, 59, 1, 59, 1, 59, 1, 59, 3, 59, 389, 8, 59,
		1, 60, 1, 60, 1, 60, 1, 60, 4, 60, 395, 8, 60, 11, 60, 12, 60, 396, 1,
		61, 1, 61, 5, 61, 401, 8, 61, 10, 61, 12, 61, 404, 9, 61, 0, 0, 62, 1,
		1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11,
		23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17, 35, 18, 37, 19, 39, 20,
		41, 21, 43, 22, 45, 23, 47, 24, 49, 25, 51, 26, 53, 27, 55, 28, 57, 29,
		59, 30, 61, 31, 63, 32, 65, 33, 67, 34, 69, 35, 71, 36, 73, 37, 75, 38,
		77, 39, 79, 40, 81, 41, 83, 42, 85, 43, 87, 44, 89, 45, 91, 46, 93, 47,
		95, 48, 97, 49, 99, 50, 101, 51, 103, 52, 105, 53, 107, 54, 109, 55, 111,
		56, 113, 57, 115, 58, 117, 59, 119, 60, 121, 61, 123, 62, 1, 0, 8, 2, 0,
		10, 10, 13, 13, 1, 0, 32, 32, 1, 0, 48, 49, 1, 0, 48, 57, 3, 0, 10, 10,
		13, 13, 39, 39, 3, 0, 48, 57, 65, 70, 97, 102, 3, 0, 65, 90, 95, 95, 97,
		122, 4, 0, 48, 57, 65, 90, 95, 95, 97, 122, 415, 0, 1, 1, 0, 0, 0, 0, 3,
		1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11,
		1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0,
		19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0,
		0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0,
		0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0,
		0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0, 0, 0, 49, 1,
		0, 0, 0, 0, 51, 1, 0, 0, 0, 0, 53, 1, 0, 0, 0, 0, 55, 1, 0, 0, 0, 0, 57,
		1, 0, 0, 0, 0, 59, 1, 0, 0, 0, 0, 61, 1, 0, 0, 0, 0, 63, 1, 0, 0, 0, 0,
		65, 1, 0, 0, 0, 0, 67, 1, 0, 0, 0, 0, 69, 1, 0, 0, 0, 0, 71, 1, 0, 0, 0,
		0, 73, 1, 0, 0, 0, 0, 75, 1, 0, 0, 0, 0, 77, 1, 0, 0, 0, 0, 79, 1, 0, 0,
		0, 0, 81, 1, 0, 0, 0, 0, 83, 1, 0, 0, 0, 0, 85, 1, 0, 0, 0, 0, 87, 1, 0,
		0, 0, 0, 89, 1, 0, 0, 0, 0, 91, 1, 0, 0, 0, 0, 93, 1, 0, 0, 0, 0, 95, 1,
		0, 0, 0, 0, 97, 1, 0, 0, 0, 0, 99, 1, 0, 0, 0, 0, 101, 1, 0, 0, 0, 0, 103,
		1, 0, 0, 0, 0, 105, 1, 0, 0, 0, 0, 107, 1, 0, 0, 0, 0, 109, 1, 0, 0, 0,
		0, 111, 1, 0, 0, 0, 0, 113, 1, 0, 0, 0, 0, 115, 1, 0, 0, 0, 0, 117, 1,
		0, 0, 0, 0, 119, 1, 0, 0, 0, 0, 121, 1, 0, 0, 0, 0, 123, 1, 0, 0, 0, 1,
		125, 1, 0, 0, 0, 3, 127, 1, 0, 0, 0, 5, 132, 1, 0, 0, 0, 7, 135, 1, 0,
		0, 0, 9, 137, 1, 0, 0, 0, 11, 139, 1, 0, 0, 0, 13, 141, 1, 0, 0, 0, 15,
		144, 1, 0, 0, 0, 17, 148, 1, 0, 0, 0, 19, 152, 1, 0, 0, 0, 21, 156, 1,
		0, 0, 0, 23, 160, 1, 0, 0, 0, 25, 164, 1, 0, 0, 0, 27, 168, 1, 0, 0, 0,
		29, 172, 1, 0, 0, 0, 31, 176, 1, 0, 0, 0, 33, 180, 1, 0, 0, 0, 35, 184,
		1, 0, 0, 0, 37, 187, 1, 0, 0, 0, 39, 191, 1, 0, 0, 0, 41, 195, 1, 0, 0,
		0, 43, 199, 1, 0, 0, 0, 45, 203, 1, 0, 0, 0, 47, 207, 1, 0, 0, 0, 49, 211,
		1, 0, 0, 0, 51, 214, 1, 0, 0, 0, 53, 218, 1, 0, 0, 0, 55, 221, 1, 0, 0,
		0, 57, 224, 1, 0, 0, 0, 59, 228, 1, 0, 0, 0, 61, 232, 1, 0, 0, 0, 63, 236,
		1, 0, 0, 0, 65, 240, 1, 0, 0, 0, 67, 244, 1, 0, 0, 0, 69, 249, 1, 0, 0,
		0, 71, 253, 1, 0, 0, 0, 73, 257, 1, 0, 0, 0, 75, 259, 1, 0, 0, 0, 77, 261,
		1, 0, 0, 0, 79, 263, 1, 0, 0, 0, 81, 265, 1, 0, 0, 0, 83, 268, 1, 0, 0,
		0, 85, 270, 1, 0, 0, 0, 87, 272, 1, 0, 0, 0, 89, 274, 1, 0, 0, 0, 91, 276,
		1, 0, 0, 0, 93, 278, 1, 0, 0, 0, 95, 285, 1, 0, 0, 0, 97, 292, 1, 0, 0,
		0, 99, 299, 1, 0, 0, 0, 101, 305, 1, 0, 0, 0, 103, 315, 1, 0, 0, 0, 105,
		317, 1, 0, 0, 0, 107, 325, 1, 0, 0, 0, 109, 336, 1, 0, 0, 0, 111, 351,
		1, 0, 0, 0, 113, 359, 1, 0, 0, 0, 115, 363, 1, 0, 0, 0, 117, 372, 1, 0,
		0, 0, 119, 388, 1, 0, 0, 0, 121, 390, 1, 0, 0, 0, 123, 398, 1, 0, 0, 0,
		125, 126, 5, 9, 0, 0, 126, 2, 1, 0, 0, 0, 127, 128, 5, 32, 0, 0, 128, 129,
		5, 32, 0, 0, 129, 130, 5, 32, 0, 0, 130, 131, 5, 32, 0, 0, 131, 4, 1, 0,
		0, 0, 132, 133, 5, 32, 0, 0, 133, 134, 5, 32, 0, 0, 134, 6, 1, 0, 0, 0,
		135, 136, 5, 32, 0, 0, 136, 8, 1, 0, 0, 0, 137, 138, 5, 10, 0, 0, 138,
		10, 1, 0, 0, 0, 139, 140, 5, 58, 0, 0, 140, 12, 1, 0, 0, 0, 141, 142, 5,
		44, 0, 0, 142, 143, 5, 32, 0, 0, 143, 14, 1, 0, 0, 0, 144, 145, 5, 109,
		0, 0, 145, 146, 5, 111, 0, 0, 146, 147, 5, 118, 0, 0, 147, 16, 1, 0, 0,
		0, 148, 149, 5, 97, 0, 0, 149, 150, 5, 100, 0, 0, 150, 151, 5, 100, 0,
		0, 151, 18, 1, 0, 0, 0, 152, 153, 5, 97, 0, 0, 153, 154, 5, 100, 0, 0,
		154, 155, 5, 99, 0, 0, 155, 20, 1, 0, 0, 0, 156, 157, 5, 115, 0, 0, 157,
		158, 5, 117, 0, 0, 158, 159, 5, 98, 0, 0, 159, 22, 1, 0, 0, 0, 160, 161,
		5, 115, 0, 0, 161, 162, 5, 98, 0, 0, 162, 163, 5, 98, 0, 0, 163, 24, 1,
		0, 0, 0, 164, 165, 5, 99, 0, 0, 165, 166, 5, 109, 0, 0, 166, 167, 5, 112,
		0, 0, 167, 26, 1, 0, 0, 0, 168, 169, 5, 105, 0, 0, 169, 170, 5, 110, 0,
		0, 170, 171, 5, 99, 0, 0, 171, 28, 1, 0, 0, 0, 172, 173, 5, 100, 0, 0,
		173, 174, 5, 101, 0, 0, 174, 175, 5, 99, 0, 0, 175, 30, 1, 0, 0, 0, 176,
		177, 5, 120, 0, 0, 177, 178, 5, 111, 0, 0, 178, 179, 5, 114, 0, 0, 179,
		32, 1, 0, 0, 0, 180, 181, 5, 97, 0, 0, 181, 182, 5, 110, 0, 0, 182, 183,
		5, 100, 0, 0, 183, 34, 1, 0, 0, 0, 184, 185, 5, 111, 0, 0, 185, 186, 5,
		114, 0, 0, 186, 36, 1, 0, 0, 0, 187, 188, 5, 110, 0, 0, 188, 189, 5, 111,
		0, 0, 189, 190, 5, 116, 0, 0, 190, 38, 1, 0, 0, 0, 191, 192, 5, 115, 0,
		0, 192, 193, 5, 104, 0, 0, 193, 194, 5, 108, 0, 0, 194, 40, 1, 0, 0, 0,
		195, 196, 5, 115, 0, 0, 196, 197, 5, 104, 0, 0, 197, 198, 5, 114, 0, 0,
		198, 42, 1, 0, 0, 0, 199, 200, 5, 114, 0, 0, 200, 201, 5, 111, 0, 0, 201,
		202, 5, 108, 0, 0, 202, 44, 1, 0, 0, 0, 203, 204, 5, 114, 0, 0, 204, 205,
		5, 111, 0, 0, 205, 206, 5, 114, 0, 0, 206, 46, 1, 0, 0, 0, 207, 208, 5,
		106, 0, 0, 208, 209, 5, 109, 0, 0, 209, 210, 5, 112, 0, 0, 210, 48, 1,
		0, 0, 0, 211, 212, 5, 106, 0, 0, 212, 213, 5, 122, 0, 0, 213, 50, 1, 0,
		0, 0, 214, 215, 5, 106, 0, 0, 215, 216, 5, 110, 0, 0, 216, 217, 5, 122,
		0, 0, 217, 52, 1, 0, 0, 0, 218, 219, 5, 106, 0, 0, 219, 220, 5, 99, 0,
		0, 220, 54, 1, 0, 0, 0, 221, 222, 5, 106, 0, 0, 222, 223, 5, 115, 0, 0,
		223, 56, 1, 0, 0, 0, 224, 225, 5, 106, 0, 0, 225, 226, 5, 110, 0, 0, 226,
		227, 5, 115, 0, 0, 227, 58, 1, 0, 0, 0, 228, 229, 5, 106, 0, 0, 229, 230,
		5, 115, 0, 0, 230, 231, 5, 114, 0, 0, 231, 60, 1, 0, 0, 0, 232, 233, 5,
		114, 0, 0, 233, 234, 5, 116, 0, 0, 234, 235, 5, 115, 0, 0, 235, 62, 1,
		0, 0, 0, 236, 237, 5, 98, 0, 0, 237, 238, 5, 114, 0, 0, 238, 239, 5, 107,
		0, 0, 239, 64, 1, 0, 0, 0, 240, 241, 5, 110, 0, 0, 241, 242, 5, 111, 0,
		0, 242, 243, 5, 112, 0, 0, 243, 66, 1, 0, 0, 0, 244, 245, 5, 112, 0, 0,
		245, 246, 5, 117, 0, 0, 246, 247, 5, 115, 0, 0, 247, 248, 5, 104, 0, 0,
		248, 68, 1, 0, 0, 0, 249, 250, 5, 112, 0, 0, 250, 251, 5, 111, 0, 0, 251,
		252, 5, 112, 0, 0, 252, 70, 1, 0, 0, 0, 253, 254, 5, 99, 0, 0, 254, 255,
		5, 108, 0, 0, 255, 256, 5, 99, 0, 0, 256, 72, 1, 0, 0, 0, 257, 258, 5,
		97, 0, 0, 258, 74, 1, 0, 0, 0, 259, 260, 5, 98, 0, 0, 260, 76, 1, 0, 0,
		0, 261, 262, 5, 99, 0, 0, 262, 78, 1, 0, 0, 0, 263, 264, 5, 100, 0, 0,
		264, 80, 1, 0, 0, 0, 265, 266, 5, 115, 0, 0, 266, 267, 5, 112, 0, 0, 267,
		82, 1, 0, 0, 0, 268, 269, 5, 91, 0, 0, 269, 84, 1, 0, 0, 0, 270, 271, 5,
		93, 0, 0, 271, 86, 1, 0, 0, 0, 272, 273, 5, 43, 0, 0, 273, 88, 1, 0, 0,
		0, 274, 275, 5, 45, 0, 0, 275, 90, 1, 0, 0, 0, 276, 277, 5, 61, 0, 0, 277,
		92, 1, 0, 0, 0, 278, 279, 5, 46, 0, 0, 279, 280, 5, 98, 0, 0, 280, 281,
		5, 121, 0, 0, 281, 282, 5, 116, 0, 0, 282, 283, 5, 101, 0, 0, 283, 284,
		5, 32, 0, 0, 284, 94, 1, 0, 0, 0, 285, 286, 5, 46, 0, 0, 286, 287, 5, 119,
		0, 0, 287, 288, 5, 111, 0, 0, 288, 289, 5, 114, 0, 0, 289, 290, 5, 100,
		0, 0, 290, 291, 5, 32, 0, 0, 291, 96, 1, 0, 0, 0, 292, 293, 5, 46, 0, 0,
		293, 294, 5, 114, 0, 0, 294, 295, 5, 101, 0, 0, 295, 296, 5, 115, 0, 0,
		296, 297, 5, 98, 0, 0, 297, 298, 5, 32, 0, 0, 298, 98, 1, 0, 0, 0, 299,
		300, 5, 46, 0, 0, 300, 301, 5, 111, 0, 0, 301, 302, 5, 114, 0, 0, 302,
		303, 5, 103, 0, 0, 303, 304, 5, 32, 0, 0, 304, 100, 1, 0, 0, 0, 305, 306,
		5, 46, 0, 0, 306, 307, 5, 115, 0, 0, 307, 308, 5, 101, 0, 0, 308, 309,
		5, 103, 0, 0, 309, 310, 5, 109, 0, 0, 310, 311, 5, 101, 0, 0, 311, 312,
		5, 110, 0, 0, 312, 313, 5, 116, 0, 0, 313, 314, 5, 32, 0, 0, 314, 102,
		1, 0, 0, 0, 315, 316, 5, 34, 0, 0, 316, 104, 1, 0, 0, 0, 317, 318, 5, 103,
		0, 0, 318, 319, 5, 108, 0, 0, 319, 320, 5, 111, 0, 0, 320, 321, 5, 98,
		0, 0, 321, 322, 5, 97, 0, 0, 322, 323, 5, 108, 0, 0, 323, 324, 5, 32, 0,
		0, 324, 106, 1, 0, 0, 0, 325, 326, 5, 101, 0, 0, 326, 327, 5, 120, 0, 0,
		327, 328, 5, 116, 0, 0, 328, 329, 5, 101, 0, 0, 329, 330, 5, 114, 0, 0,
		330, 331, 5, 110, 0, 0, 331, 332, 5, 32, 0, 0, 332, 108, 1, 0, 0, 0, 333,
		335, 3, 113, 56, 0, 334, 333, 1, 0, 0, 0, 335, 338, 1, 0, 0, 0, 336, 334,
		1, 0, 0, 0, 336, 337, 1, 0, 0, 0, 337, 339, 1, 0, 0, 0, 338, 336, 1, 0,
		0, 0, 339, 343, 5, 59, 0, 0, 340, 342, 8, 0, 0, 0, 341, 340, 1, 0, 0, 0,
		342, 345, 1, 0, 0, 0, 343, 341, 1, 0, 0, 0, 343, 344, 1, 0, 0, 0, 344,
		346, 1, 0, 0, 0, 345, 343, 1, 0, 0, 0, 346, 347, 6, 54, 0, 0, 347, 110,
		1, 0, 0, 0, 348, 350, 7, 1, 0, 0, 349, 348, 1, 0, 0, 0, 350, 353, 1, 0,
		0, 0, 351, 349, 1, 0, 0, 0, 351, 352, 1, 0, 0, 0, 352, 354, 1, 0, 0, 0,
		353, 351, 1, 0, 0, 0, 354, 355, 7, 0, 0, 0, 355, 356, 1, 0, 0, 0, 356,
		357, 6, 55, 0, 0, 357, 112, 1, 0, 0, 0, 358, 360, 7, 1, 0, 0, 359, 358,
		1, 0, 0, 0, 360, 361, 1, 0, 0, 0, 361, 359, 1, 0, 0, 0, 361, 362, 1, 0,
		0, 0, 362, 114, 1, 0, 0, 0, 363, 364, 5, 48, 0, 0, 364, 365, 5, 98, 0,
		0, 365, 367, 1, 0, 0, 0, 366, 368, 7, 2, 0, 0, 367, 366, 1, 0, 0, 0, 368,
		369, 1, 0, 0, 0, 369, 367, 1, 0, 0, 0, 369, 370, 1, 0, 0, 0, 370, 116,
		1, 0, 0, 0, 371, 373, 7, 3, 0, 0, 372, 371, 1, 0, 0, 0, 373, 374, 1, 0,
		0, 0, 374, 372, 1, 0, 0, 0, 374, 375, 1, 0, 0, 0, 375, 118, 1, 0, 0, 0,
		376, 377, 5, 39, 0, 0, 377, 378, 5, 32, 0, 0, 378, 389, 5, 39, 0, 0, 379,
		380, 5, 34, 0, 0, 380, 381, 5, 39, 0, 0, 381, 389, 5, 34, 0, 0, 382, 383,
		5, 39, 0, 0, 383, 384, 5, 44, 0, 0, 384, 389, 5, 39, 0, 0, 385, 386, 5,
		39, 0, 0, 386, 387, 8, 4, 0, 0, 387, 389, 5, 39, 0, 0, 388, 376, 1, 0,
		0, 0, 388, 379, 1, 0, 0, 0, 388, 382, 1, 0, 0, 0, 388, 385, 1, 0, 0, 0,
		389, 120, 1, 0, 0, 0, 390, 391, 5, 48, 0, 0, 391, 392, 5, 120, 0, 0, 392,
		394, 1, 0, 0, 0, 393, 395, 7, 5, 0, 0, 394, 393, 1, 0, 0, 0, 395, 396,
		1, 0, 0, 0, 396, 394, 1, 0, 0, 0, 396, 397, 1, 0, 0, 0, 397, 122, 1, 0,
		0, 0, 398, 402, 7, 6, 0, 0, 399, 401, 7, 7, 0, 0, 400, 399, 1, 0, 0, 0,
		401, 404, 1, 0, 0, 0, 402, 400, 1, 0, 0, 0, 402, 403, 1, 0, 0, 0, 403,
		124, 1, 0, 0, 0, 404, 402, 1, 0, 0, 0, 10, 0, 336, 343, 351, 361, 369,
		374, 388, 396, 402, 1, 6, 0, 0,
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

// AsmE8LexerInit initializes any static state used to implement AsmE8Lexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewAsmE8Lexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func AsmE8LexerInit() {
	staticData := &AsmE8LexerLexerStaticData
	staticData.once.Do(asme8lexerLexerInit)
}

// NewAsmE8Lexer produces a new lexer instance for the optional input antlr.CharStream.
func NewAsmE8Lexer(input antlr.CharStream) *AsmE8Lexer {
	AsmE8LexerInit()
	l := new(AsmE8Lexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &AsmE8LexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "AsmE8.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// AsmE8Lexer tokens.
const (
	AsmE8LexerT__0              = 1
	AsmE8LexerT__1              = 2
	AsmE8LexerT__2              = 3
	AsmE8LexerT__3              = 4
	AsmE8LexerT__4              = 5
	AsmE8LexerT__5              = 6
	AsmE8LexerT__6              = 7
	AsmE8LexerT__7              = 8
	AsmE8LexerT__8              = 9
	AsmE8LexerT__9              = 10
	AsmE8LexerT__10             = 11
	AsmE8LexerT__11             = 12
	AsmE8LexerT__12             = 13
	AsmE8LexerT__13             = 14
	AsmE8LexerT__14             = 15
	AsmE8LexerT__15             = 16
	AsmE8LexerT__16             = 17
	AsmE8LexerT__17             = 18
	AsmE8LexerT__18             = 19
	AsmE8LexerT__19             = 20
	AsmE8LexerT__20             = 21
	AsmE8LexerT__21             = 22
	AsmE8LexerT__22             = 23
	AsmE8LexerT__23             = 24
	AsmE8LexerT__24             = 25
	AsmE8LexerT__25             = 26
	AsmE8LexerT__26             = 27
	AsmE8LexerT__27             = 28
	AsmE8LexerT__28             = 29
	AsmE8LexerT__29             = 30
	AsmE8LexerT__30             = 31
	AsmE8LexerT__31             = 32
	AsmE8LexerT__32             = 33
	AsmE8LexerT__33             = 34
	AsmE8LexerT__34             = 35
	AsmE8LexerT__35             = 36
	AsmE8LexerT__36             = 37
	AsmE8LexerT__37             = 38
	AsmE8LexerT__38             = 39
	AsmE8LexerT__39             = 40
	AsmE8LexerT__40             = 41
	AsmE8LexerT__41             = 42
	AsmE8LexerT__42             = 43
	AsmE8LexerT__43             = 44
	AsmE8LexerT__44             = 45
	AsmE8LexerT__45             = 46
	AsmE8LexerT__46             = 47
	AsmE8LexerT__47             = 48
	AsmE8LexerT__48             = 49
	AsmE8LexerT__49             = 50
	AsmE8LexerT__50             = 51
	AsmE8LexerT__51             = 52
	AsmE8LexerT__52             = 53
	AsmE8LexerT__53             = 54
	AsmE8LexerLINE_COMMENT      = 55
	AsmE8LexerSUFFIX_WHITESPACE = 56
	AsmE8LexerWHITE_SPACE       = 57
	AsmE8LexerBINARY            = 58
	AsmE8LexerINT               = 59
	AsmE8LexerCHAR              = 60
	AsmE8LexerHEX               = 61
	AsmE8LexerSTR               = 62
)
