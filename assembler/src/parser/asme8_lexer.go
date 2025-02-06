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
		"'jz'", "'jnz'", "'jc'", "'js'", "'jns'", "'call'", "'ret'", "'brk'",
		"'nop'", "'push'", "'pop'", "'clc'", "'a'", "'b'", "'c'", "'d'", "'ipl'",
		"'iph'", "'sp'", "'ip'", "'['", "']'", "'+'", "'-'", "'='", "'.asciiz '",
		"'.ascii '", "'.byte '", "'.word '", "'.resb '", "'.org '", "'.segment '",
		"'\"'", "'global '", "'extern '",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "LINE_COMMENT", "SUFFIX_WHITESPACE",
		"WHITE_SPACE", "BINARY", "INT", "CHAR", "ANY", "HEX", "STR",
	}
	staticData.RuleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
		"T__9", "T__10", "T__11", "T__12", "T__13", "T__14", "T__15", "T__16",
		"T__17", "T__18", "T__19", "T__20", "T__21", "T__22", "T__23", "T__24",
		"T__25", "T__26", "T__27", "T__28", "T__29", "T__30", "T__31", "T__32",
		"T__33", "T__34", "T__35", "T__36", "T__37", "T__38", "T__39", "T__40",
		"T__41", "T__42", "T__43", "T__44", "T__45", "T__46", "T__47", "T__48",
		"T__49", "T__50", "T__51", "T__52", "T__53", "T__54", "T__55", "T__56",
		"T__57", "T__58", "LINE_COMMENT", "SUFFIX_WHITESPACE", "WHITE_SPACE",
		"BINARY", "INT", "CHAR", "ANY", "HEX", "STR",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 68, 455, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
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
		7, 57, 2, 58, 7, 58, 2, 59, 7, 59, 2, 60, 7, 60, 2, 61, 7, 61, 2, 62, 7,
		62, 2, 63, 7, 63, 2, 64, 7, 64, 2, 65, 7, 65, 2, 66, 7, 66, 2, 67, 7, 67,
		1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3,
		1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8,
		1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1,
		11, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13,
		1, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 1, 15, 1, 16, 1,
		16, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 1, 18, 1, 18, 1, 18, 1, 18, 1, 19,
		1, 19, 1, 19, 1, 19, 1, 20, 1, 20, 1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 1,
		21, 1, 22, 1, 22, 1, 22, 1, 22, 1, 23, 1, 23, 1, 23, 1, 23, 1, 24, 1, 24,
		1, 24, 1, 25, 1, 25, 1, 25, 1, 25, 1, 26, 1, 26, 1, 26, 1, 27, 1, 27, 1,
		27, 1, 28, 1, 28, 1, 28, 1, 28, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 30,
		1, 30, 1, 30, 1, 30, 1, 31, 1, 31, 1, 31, 1, 31, 1, 32, 1, 32, 1, 32, 1,
		32, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 34, 1, 34, 1, 34, 1, 34, 1, 35,
		1, 35, 1, 35, 1, 35, 1, 36, 1, 36, 1, 37, 1, 37, 1, 38, 1, 38, 1, 39, 1,
		39, 1, 40, 1, 40, 1, 40, 1, 40, 1, 41, 1, 41, 1, 41, 1, 41, 1, 42, 1, 42,
		1, 42, 1, 43, 1, 43, 1, 43, 1, 44, 1, 44, 1, 45, 1, 45, 1, 46, 1, 46, 1,
		47, 1, 47, 1, 48, 1, 48, 1, 49, 1, 49, 1, 49, 1, 49, 1, 49, 1, 49, 1, 49,
		1, 49, 1, 49, 1, 50, 1, 50, 1, 50, 1, 50, 1, 50, 1, 50, 1, 50, 1, 50, 1,
		51, 1, 51, 1, 51, 1, 51, 1, 51, 1, 51, 1, 51, 1, 52, 1, 52, 1, 52, 1, 52,
		1, 52, 1, 52, 1, 52, 1, 53, 1, 53, 1, 53, 1, 53, 1, 53, 1, 53, 1, 53, 1,
		54, 1, 54, 1, 54, 1, 54, 1, 54, 1, 54, 1, 55, 1, 55, 1, 55, 1, 55, 1, 55,
		1, 55, 1, 55, 1, 55, 1, 55, 1, 55, 1, 56, 1, 56, 1, 57, 1, 57, 1, 57, 1,
		57, 1, 57, 1, 57, 1, 57, 1, 57, 1, 58, 1, 58, 1, 58, 1, 58, 1, 58, 1, 58,
		1, 58, 1, 58, 1, 59, 5, 59, 376, 8, 59, 10, 59, 12, 59, 379, 9, 59, 1,
		59, 1, 59, 5, 59, 383, 8, 59, 10, 59, 12, 59, 386, 9, 59, 1, 59, 1, 59,
		1, 60, 5, 60, 391, 8, 60, 10, 60, 12, 60, 394, 9, 60, 1, 60, 1, 60, 1,
		60, 1, 60, 1, 61, 4, 61, 401, 8, 61, 11, 61, 12, 61, 402, 1, 62, 1, 62,
		1, 62, 1, 62, 4, 62, 409, 8, 62, 11, 62, 12, 62, 410, 1, 63, 4, 63, 414,
		8, 63, 11, 63, 12, 63, 415, 1, 64, 1, 64, 1, 64, 1, 64, 1, 64, 1, 64, 1,
		64, 1, 64, 1, 64, 1, 64, 1, 64, 1, 64, 3, 64, 430, 8, 64, 1, 65, 1, 65,
		5, 65, 434, 8, 65, 10, 65, 12, 65, 437, 9, 65, 1, 65, 1, 65, 1, 66, 1,
		66, 1, 66, 1, 66, 4, 66, 445, 8, 66, 11, 66, 12, 66, 446, 1, 67, 1, 67,
		5, 67, 451, 8, 67, 10, 67, 12, 67, 454, 9, 67, 0, 0, 68, 1, 1, 3, 2, 5,
		3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25,
		13, 27, 14, 29, 15, 31, 16, 33, 17, 35, 18, 37, 19, 39, 20, 41, 21, 43,
		22, 45, 23, 47, 24, 49, 25, 51, 26, 53, 27, 55, 28, 57, 29, 59, 30, 61,
		31, 63, 32, 65, 33, 67, 34, 69, 35, 71, 36, 73, 37, 75, 38, 77, 39, 79,
		40, 81, 41, 83, 42, 85, 43, 87, 44, 89, 45, 91, 46, 93, 47, 95, 48, 97,
		49, 99, 50, 101, 51, 103, 52, 105, 53, 107, 54, 109, 55, 111, 56, 113,
		57, 115, 58, 117, 59, 119, 60, 121, 61, 123, 62, 125, 63, 127, 64, 129,
		65, 131, 66, 133, 67, 135, 68, 1, 0, 9, 2, 0, 10, 10, 13, 13, 1, 0, 32,
		32, 1, 0, 48, 49, 1, 0, 48, 57, 3, 0, 10, 10, 13, 13, 39, 39, 3, 0, 10,
		10, 13, 13, 34, 34, 3, 0, 48, 57, 65, 70, 97, 102, 3, 0, 65, 90, 95, 95,
		97, 122, 4, 0, 48, 57, 65, 90, 95, 95, 97, 122, 466, 0, 1, 1, 0, 0, 0,
		0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0,
		0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0,
		0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0,
		0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1,
		0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41,
		1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0, 0, 0,
		49, 1, 0, 0, 0, 0, 51, 1, 0, 0, 0, 0, 53, 1, 0, 0, 0, 0, 55, 1, 0, 0, 0,
		0, 57, 1, 0, 0, 0, 0, 59, 1, 0, 0, 0, 0, 61, 1, 0, 0, 0, 0, 63, 1, 0, 0,
		0, 0, 65, 1, 0, 0, 0, 0, 67, 1, 0, 0, 0, 0, 69, 1, 0, 0, 0, 0, 71, 1, 0,
		0, 0, 0, 73, 1, 0, 0, 0, 0, 75, 1, 0, 0, 0, 0, 77, 1, 0, 0, 0, 0, 79, 1,
		0, 0, 0, 0, 81, 1, 0, 0, 0, 0, 83, 1, 0, 0, 0, 0, 85, 1, 0, 0, 0, 0, 87,
		1, 0, 0, 0, 0, 89, 1, 0, 0, 0, 0, 91, 1, 0, 0, 0, 0, 93, 1, 0, 0, 0, 0,
		95, 1, 0, 0, 0, 0, 97, 1, 0, 0, 0, 0, 99, 1, 0, 0, 0, 0, 101, 1, 0, 0,
		0, 0, 103, 1, 0, 0, 0, 0, 105, 1, 0, 0, 0, 0, 107, 1, 0, 0, 0, 0, 109,
		1, 0, 0, 0, 0, 111, 1, 0, 0, 0, 0, 113, 1, 0, 0, 0, 0, 115, 1, 0, 0, 0,
		0, 117, 1, 0, 0, 0, 0, 119, 1, 0, 0, 0, 0, 121, 1, 0, 0, 0, 0, 123, 1,
		0, 0, 0, 0, 125, 1, 0, 0, 0, 0, 127, 1, 0, 0, 0, 0, 129, 1, 0, 0, 0, 0,
		131, 1, 0, 0, 0, 0, 133, 1, 0, 0, 0, 0, 135, 1, 0, 0, 0, 1, 137, 1, 0,
		0, 0, 3, 139, 1, 0, 0, 0, 5, 144, 1, 0, 0, 0, 7, 147, 1, 0, 0, 0, 9, 149,
		1, 0, 0, 0, 11, 151, 1, 0, 0, 0, 13, 153, 1, 0, 0, 0, 15, 156, 1, 0, 0,
		0, 17, 160, 1, 0, 0, 0, 19, 164, 1, 0, 0, 0, 21, 168, 1, 0, 0, 0, 23, 172,
		1, 0, 0, 0, 25, 176, 1, 0, 0, 0, 27, 180, 1, 0, 0, 0, 29, 184, 1, 0, 0,
		0, 31, 188, 1, 0, 0, 0, 33, 192, 1, 0, 0, 0, 35, 196, 1, 0, 0, 0, 37, 199,
		1, 0, 0, 0, 39, 203, 1, 0, 0, 0, 41, 207, 1, 0, 0, 0, 43, 211, 1, 0, 0,
		0, 45, 215, 1, 0, 0, 0, 47, 219, 1, 0, 0, 0, 49, 223, 1, 0, 0, 0, 51, 226,
		1, 0, 0, 0, 53, 230, 1, 0, 0, 0, 55, 233, 1, 0, 0, 0, 57, 236, 1, 0, 0,
		0, 59, 240, 1, 0, 0, 0, 61, 245, 1, 0, 0, 0, 63, 249, 1, 0, 0, 0, 65, 253,
		1, 0, 0, 0, 67, 257, 1, 0, 0, 0, 69, 262, 1, 0, 0, 0, 71, 266, 1, 0, 0,
		0, 73, 270, 1, 0, 0, 0, 75, 272, 1, 0, 0, 0, 77, 274, 1, 0, 0, 0, 79, 276,
		1, 0, 0, 0, 81, 278, 1, 0, 0, 0, 83, 282, 1, 0, 0, 0, 85, 286, 1, 0, 0,
		0, 87, 289, 1, 0, 0, 0, 89, 292, 1, 0, 0, 0, 91, 294, 1, 0, 0, 0, 93, 296,
		1, 0, 0, 0, 95, 298, 1, 0, 0, 0, 97, 300, 1, 0, 0, 0, 99, 302, 1, 0, 0,
		0, 101, 311, 1, 0, 0, 0, 103, 319, 1, 0, 0, 0, 105, 326, 1, 0, 0, 0, 107,
		333, 1, 0, 0, 0, 109, 340, 1, 0, 0, 0, 111, 346, 1, 0, 0, 0, 113, 356,
		1, 0, 0, 0, 115, 358, 1, 0, 0, 0, 117, 366, 1, 0, 0, 0, 119, 377, 1, 0,
		0, 0, 121, 392, 1, 0, 0, 0, 123, 400, 1, 0, 0, 0, 125, 404, 1, 0, 0, 0,
		127, 413, 1, 0, 0, 0, 129, 429, 1, 0, 0, 0, 131, 431, 1, 0, 0, 0, 133,
		440, 1, 0, 0, 0, 135, 448, 1, 0, 0, 0, 137, 138, 5, 9, 0, 0, 138, 2, 1,
		0, 0, 0, 139, 140, 5, 32, 0, 0, 140, 141, 5, 32, 0, 0, 141, 142, 5, 32,
		0, 0, 142, 143, 5, 32, 0, 0, 143, 4, 1, 0, 0, 0, 144, 145, 5, 32, 0, 0,
		145, 146, 5, 32, 0, 0, 146, 6, 1, 0, 0, 0, 147, 148, 5, 32, 0, 0, 148,
		8, 1, 0, 0, 0, 149, 150, 5, 10, 0, 0, 150, 10, 1, 0, 0, 0, 151, 152, 5,
		58, 0, 0, 152, 12, 1, 0, 0, 0, 153, 154, 5, 44, 0, 0, 154, 155, 5, 32,
		0, 0, 155, 14, 1, 0, 0, 0, 156, 157, 5, 109, 0, 0, 157, 158, 5, 111, 0,
		0, 158, 159, 5, 118, 0, 0, 159, 16, 1, 0, 0, 0, 160, 161, 5, 97, 0, 0,
		161, 162, 5, 100, 0, 0, 162, 163, 5, 100, 0, 0, 163, 18, 1, 0, 0, 0, 164,
		165, 5, 97, 0, 0, 165, 166, 5, 100, 0, 0, 166, 167, 5, 99, 0, 0, 167, 20,
		1, 0, 0, 0, 168, 169, 5, 115, 0, 0, 169, 170, 5, 117, 0, 0, 170, 171, 5,
		98, 0, 0, 171, 22, 1, 0, 0, 0, 172, 173, 5, 115, 0, 0, 173, 174, 5, 98,
		0, 0, 174, 175, 5, 98, 0, 0, 175, 24, 1, 0, 0, 0, 176, 177, 5, 99, 0, 0,
		177, 178, 5, 109, 0, 0, 178, 179, 5, 112, 0, 0, 179, 26, 1, 0, 0, 0, 180,
		181, 5, 105, 0, 0, 181, 182, 5, 110, 0, 0, 182, 183, 5, 99, 0, 0, 183,
		28, 1, 0, 0, 0, 184, 185, 5, 100, 0, 0, 185, 186, 5, 101, 0, 0, 186, 187,
		5, 99, 0, 0, 187, 30, 1, 0, 0, 0, 188, 189, 5, 120, 0, 0, 189, 190, 5,
		111, 0, 0, 190, 191, 5, 114, 0, 0, 191, 32, 1, 0, 0, 0, 192, 193, 5, 97,
		0, 0, 193, 194, 5, 110, 0, 0, 194, 195, 5, 100, 0, 0, 195, 34, 1, 0, 0,
		0, 196, 197, 5, 111, 0, 0, 197, 198, 5, 114, 0, 0, 198, 36, 1, 0, 0, 0,
		199, 200, 5, 110, 0, 0, 200, 201, 5, 111, 0, 0, 201, 202, 5, 116, 0, 0,
		202, 38, 1, 0, 0, 0, 203, 204, 5, 115, 0, 0, 204, 205, 5, 104, 0, 0, 205,
		206, 5, 108, 0, 0, 206, 40, 1, 0, 0, 0, 207, 208, 5, 115, 0, 0, 208, 209,
		5, 104, 0, 0, 209, 210, 5, 114, 0, 0, 210, 42, 1, 0, 0, 0, 211, 212, 5,
		114, 0, 0, 212, 213, 5, 111, 0, 0, 213, 214, 5, 108, 0, 0, 214, 44, 1,
		0, 0, 0, 215, 216, 5, 114, 0, 0, 216, 217, 5, 111, 0, 0, 217, 218, 5, 114,
		0, 0, 218, 46, 1, 0, 0, 0, 219, 220, 5, 106, 0, 0, 220, 221, 5, 109, 0,
		0, 221, 222, 5, 112, 0, 0, 222, 48, 1, 0, 0, 0, 223, 224, 5, 106, 0, 0,
		224, 225, 5, 122, 0, 0, 225, 50, 1, 0, 0, 0, 226, 227, 5, 106, 0, 0, 227,
		228, 5, 110, 0, 0, 228, 229, 5, 122, 0, 0, 229, 52, 1, 0, 0, 0, 230, 231,
		5, 106, 0, 0, 231, 232, 5, 99, 0, 0, 232, 54, 1, 0, 0, 0, 233, 234, 5,
		106, 0, 0, 234, 235, 5, 115, 0, 0, 235, 56, 1, 0, 0, 0, 236, 237, 5, 106,
		0, 0, 237, 238, 5, 110, 0, 0, 238, 239, 5, 115, 0, 0, 239, 58, 1, 0, 0,
		0, 240, 241, 5, 99, 0, 0, 241, 242, 5, 97, 0, 0, 242, 243, 5, 108, 0, 0,
		243, 244, 5, 108, 0, 0, 244, 60, 1, 0, 0, 0, 245, 246, 5, 114, 0, 0, 246,
		247, 5, 101, 0, 0, 247, 248, 5, 116, 0, 0, 248, 62, 1, 0, 0, 0, 249, 250,
		5, 98, 0, 0, 250, 251, 5, 114, 0, 0, 251, 252, 5, 107, 0, 0, 252, 64, 1,
		0, 0, 0, 253, 254, 5, 110, 0, 0, 254, 255, 5, 111, 0, 0, 255, 256, 5, 112,
		0, 0, 256, 66, 1, 0, 0, 0, 257, 258, 5, 112, 0, 0, 258, 259, 5, 117, 0,
		0, 259, 260, 5, 115, 0, 0, 260, 261, 5, 104, 0, 0, 261, 68, 1, 0, 0, 0,
		262, 263, 5, 112, 0, 0, 263, 264, 5, 111, 0, 0, 264, 265, 5, 112, 0, 0,
		265, 70, 1, 0, 0, 0, 266, 267, 5, 99, 0, 0, 267, 268, 5, 108, 0, 0, 268,
		269, 5, 99, 0, 0, 269, 72, 1, 0, 0, 0, 270, 271, 5, 97, 0, 0, 271, 74,
		1, 0, 0, 0, 272, 273, 5, 98, 0, 0, 273, 76, 1, 0, 0, 0, 274, 275, 5, 99,
		0, 0, 275, 78, 1, 0, 0, 0, 276, 277, 5, 100, 0, 0, 277, 80, 1, 0, 0, 0,
		278, 279, 5, 105, 0, 0, 279, 280, 5, 112, 0, 0, 280, 281, 5, 108, 0, 0,
		281, 82, 1, 0, 0, 0, 282, 283, 5, 105, 0, 0, 283, 284, 5, 112, 0, 0, 284,
		285, 5, 104, 0, 0, 285, 84, 1, 0, 0, 0, 286, 287, 5, 115, 0, 0, 287, 288,
		5, 112, 0, 0, 288, 86, 1, 0, 0, 0, 289, 290, 5, 105, 0, 0, 290, 291, 5,
		112, 0, 0, 291, 88, 1, 0, 0, 0, 292, 293, 5, 91, 0, 0, 293, 90, 1, 0, 0,
		0, 294, 295, 5, 93, 0, 0, 295, 92, 1, 0, 0, 0, 296, 297, 5, 43, 0, 0, 297,
		94, 1, 0, 0, 0, 298, 299, 5, 45, 0, 0, 299, 96, 1, 0, 0, 0, 300, 301, 5,
		61, 0, 0, 301, 98, 1, 0, 0, 0, 302, 303, 5, 46, 0, 0, 303, 304, 5, 97,
		0, 0, 304, 305, 5, 115, 0, 0, 305, 306, 5, 99, 0, 0, 306, 307, 5, 105,
		0, 0, 307, 308, 5, 105, 0, 0, 308, 309, 5, 122, 0, 0, 309, 310, 5, 32,
		0, 0, 310, 100, 1, 0, 0, 0, 311, 312, 5, 46, 0, 0, 312, 313, 5, 97, 0,
		0, 313, 314, 5, 115, 0, 0, 314, 315, 5, 99, 0, 0, 315, 316, 5, 105, 0,
		0, 316, 317, 5, 105, 0, 0, 317, 318, 5, 32, 0, 0, 318, 102, 1, 0, 0, 0,
		319, 320, 5, 46, 0, 0, 320, 321, 5, 98, 0, 0, 321, 322, 5, 121, 0, 0, 322,
		323, 5, 116, 0, 0, 323, 324, 5, 101, 0, 0, 324, 325, 5, 32, 0, 0, 325,
		104, 1, 0, 0, 0, 326, 327, 5, 46, 0, 0, 327, 328, 5, 119, 0, 0, 328, 329,
		5, 111, 0, 0, 329, 330, 5, 114, 0, 0, 330, 331, 5, 100, 0, 0, 331, 332,
		5, 32, 0, 0, 332, 106, 1, 0, 0, 0, 333, 334, 5, 46, 0, 0, 334, 335, 5,
		114, 0, 0, 335, 336, 5, 101, 0, 0, 336, 337, 5, 115, 0, 0, 337, 338, 5,
		98, 0, 0, 338, 339, 5, 32, 0, 0, 339, 108, 1, 0, 0, 0, 340, 341, 5, 46,
		0, 0, 341, 342, 5, 111, 0, 0, 342, 343, 5, 114, 0, 0, 343, 344, 5, 103,
		0, 0, 344, 345, 5, 32, 0, 0, 345, 110, 1, 0, 0, 0, 346, 347, 5, 46, 0,
		0, 347, 348, 5, 115, 0, 0, 348, 349, 5, 101, 0, 0, 349, 350, 5, 103, 0,
		0, 350, 351, 5, 109, 0, 0, 351, 352, 5, 101, 0, 0, 352, 353, 5, 110, 0,
		0, 353, 354, 5, 116, 0, 0, 354, 355, 5, 32, 0, 0, 355, 112, 1, 0, 0, 0,
		356, 357, 5, 34, 0, 0, 357, 114, 1, 0, 0, 0, 358, 359, 5, 103, 0, 0, 359,
		360, 5, 108, 0, 0, 360, 361, 5, 111, 0, 0, 361, 362, 5, 98, 0, 0, 362,
		363, 5, 97, 0, 0, 363, 364, 5, 108, 0, 0, 364, 365, 5, 32, 0, 0, 365, 116,
		1, 0, 0, 0, 366, 367, 5, 101, 0, 0, 367, 368, 5, 120, 0, 0, 368, 369, 5,
		116, 0, 0, 369, 370, 5, 101, 0, 0, 370, 371, 5, 114, 0, 0, 371, 372, 5,
		110, 0, 0, 372, 373, 5, 32, 0, 0, 373, 118, 1, 0, 0, 0, 374, 376, 3, 123,
		61, 0, 375, 374, 1, 0, 0, 0, 376, 379, 1, 0, 0, 0, 377, 375, 1, 0, 0, 0,
		377, 378, 1, 0, 0, 0, 378, 380, 1, 0, 0, 0, 379, 377, 1, 0, 0, 0, 380,
		384, 5, 59, 0, 0, 381, 383, 8, 0, 0, 0, 382, 381, 1, 0, 0, 0, 383, 386,
		1, 0, 0, 0, 384, 382, 1, 0, 0, 0, 384, 385, 1, 0, 0, 0, 385, 387, 1, 0,
		0, 0, 386, 384, 1, 0, 0, 0, 387, 388, 6, 59, 0, 0, 388, 120, 1, 0, 0, 0,
		389, 391, 7, 1, 0, 0, 390, 389, 1, 0, 0, 0, 391, 394, 1, 0, 0, 0, 392,
		390, 1, 0, 0, 0, 392, 393, 1, 0, 0, 0, 393, 395, 1, 0, 0, 0, 394, 392,
		1, 0, 0, 0, 395, 396, 7, 0, 0, 0, 396, 397, 1, 0, 0, 0, 397, 398, 6, 60,
		0, 0, 398, 122, 1, 0, 0, 0, 399, 401, 7, 1, 0, 0, 400, 399, 1, 0, 0, 0,
		401, 402, 1, 0, 0, 0, 402, 400, 1, 0, 0, 0, 402, 403, 1, 0, 0, 0, 403,
		124, 1, 0, 0, 0, 404, 405, 5, 48, 0, 0, 405, 406, 5, 98, 0, 0, 406, 408,
		1, 0, 0, 0, 407, 409, 7, 2, 0, 0, 408, 407, 1, 0, 0, 0, 409, 410, 1, 0,
		0, 0, 410, 408, 1, 0, 0, 0, 410, 411, 1, 0, 0, 0, 411, 126, 1, 0, 0, 0,
		412, 414, 7, 3, 0, 0, 413, 412, 1, 0, 0, 0, 414, 415, 1, 0, 0, 0, 415,
		413, 1, 0, 0, 0, 415, 416, 1, 0, 0, 0, 416, 128, 1, 0, 0, 0, 417, 418,
		5, 39, 0, 0, 418, 419, 5, 32, 0, 0, 419, 430, 5, 39, 0, 0, 420, 421, 5,
		34, 0, 0, 421, 422, 5, 39, 0, 0, 422, 430, 5, 34, 0, 0, 423, 424, 5, 39,
		0, 0, 424, 425, 5, 44, 0, 0, 425, 430, 5, 39, 0, 0, 426, 427, 5, 39, 0,
		0, 427, 428, 8, 4, 0, 0, 428, 430, 5, 39, 0, 0, 429, 417, 1, 0, 0, 0, 429,
		420, 1, 0, 0, 0, 429, 423, 1, 0, 0, 0, 429, 426, 1, 0, 0, 0, 430, 130,
		1, 0, 0, 0, 431, 435, 5, 34, 0, 0, 432, 434, 8, 5, 0, 0, 433, 432, 1, 0,
		0, 0, 434, 437, 1, 0, 0, 0, 435, 433, 1, 0, 0, 0, 435, 436, 1, 0, 0, 0,
		436, 438, 1, 0, 0, 0, 437, 435, 1, 0, 0, 0, 438, 439, 5, 34, 0, 0, 439,
		132, 1, 0, 0, 0, 440, 441, 5, 48, 0, 0, 441, 442, 5, 120, 0, 0, 442, 444,
		1, 0, 0, 0, 443, 445, 7, 6, 0, 0, 444, 443, 1, 0, 0, 0, 445, 446, 1, 0,
		0, 0, 446, 444, 1, 0, 0, 0, 446, 447, 1, 0, 0, 0, 447, 134, 1, 0, 0, 0,
		448, 452, 7, 7, 0, 0, 449, 451, 7, 8, 0, 0, 450, 449, 1, 0, 0, 0, 451,
		454, 1, 0, 0, 0, 452, 450, 1, 0, 0, 0, 452, 453, 1, 0, 0, 0, 453, 136,
		1, 0, 0, 0, 454, 452, 1, 0, 0, 0, 11, 0, 377, 384, 392, 402, 410, 415,
		429, 435, 446, 452, 1, 6, 0, 0,
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
	AsmE8LexerT__54             = 55
	AsmE8LexerT__55             = 56
	AsmE8LexerT__56             = 57
	AsmE8LexerT__57             = 58
	AsmE8LexerT__58             = 59
	AsmE8LexerLINE_COMMENT      = 60
	AsmE8LexerSUFFIX_WHITESPACE = 61
	AsmE8LexerWHITE_SPACE       = 62
	AsmE8LexerBINARY            = 63
	AsmE8LexerINT               = 64
	AsmE8LexerCHAR              = 65
	AsmE8LexerANY               = 66
	AsmE8LexerHEX               = 67
	AsmE8LexerSTR               = 68
)
