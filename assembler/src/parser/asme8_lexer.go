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
		"", "'\\t'", "'    '", "'  '", "'\\n'", "':'", "' '", "', '", "'mov'",
		"'add'", "'adc'", "'sub'", "'sbb'", "'cmp'", "'inc'", "'dec'", "'jmp'",
		"'jz'", "'jnz'", "'jsr'", "'rts'", "'brk'", "'nop'", "'push'", "'pop'",
		"'a'", "'b'", "'c'", "'d'", "'['", "']'", "'+'", "' = '",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "LINE_COMMENT",
		"INT", "STR", "HEX", "WHITE_SPACE",
	}
	staticData.RuleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
		"T__9", "T__10", "T__11", "T__12", "T__13", "T__14", "T__15", "T__16",
		"T__17", "T__18", "T__19", "T__20", "T__21", "T__22", "T__23", "T__24",
		"T__25", "T__26", "T__27", "T__28", "T__29", "T__30", "T__31", "LINE_COMMENT",
		"INT", "STR", "HEX", "WHITE_SPACE",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 37, 222, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36,
		7, 36, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 3,
		1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7,
		1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1,
		10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13,
		1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 1, 15, 1,
		16, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 1, 17, 1, 18, 1, 18, 1, 18, 1, 18,
		1, 19, 1, 19, 1, 19, 1, 19, 1, 20, 1, 20, 1, 20, 1, 20, 1, 21, 1, 21, 1,
		21, 1, 21, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 23, 1, 23, 1, 23, 1, 23,
		1, 24, 1, 24, 1, 25, 1, 25, 1, 26, 1, 26, 1, 27, 1, 27, 1, 28, 1, 28, 1,
		29, 1, 29, 1, 30, 1, 30, 1, 31, 1, 31, 1, 31, 1, 31, 1, 32, 5, 32, 182,
		8, 32, 10, 32, 12, 32, 185, 9, 32, 1, 32, 1, 32, 5, 32, 189, 8, 32, 10,
		32, 12, 32, 192, 9, 32, 1, 32, 1, 32, 1, 33, 4, 33, 197, 8, 33, 11, 33,
		12, 33, 198, 1, 34, 1, 34, 5, 34, 203, 8, 34, 10, 34, 12, 34, 206, 9, 34,
		1, 35, 1, 35, 1, 35, 1, 35, 4, 35, 212, 8, 35, 11, 35, 12, 35, 213, 1,
		36, 4, 36, 217, 8, 36, 11, 36, 12, 36, 218, 1, 36, 1, 36, 0, 0, 37, 1,
		1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11,
		23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17, 35, 18, 37, 19, 39, 20,
		41, 21, 43, 22, 45, 23, 47, 24, 49, 25, 51, 26, 53, 27, 55, 28, 57, 29,
		59, 30, 61, 31, 63, 32, 65, 33, 67, 34, 69, 35, 71, 36, 73, 37, 1, 0, 6,
		2, 0, 10, 10, 13, 13, 1, 0, 48, 57, 3, 0, 65, 90, 95, 95, 97, 122, 4, 0,
		48, 57, 65, 90, 95, 95, 97, 122, 3, 0, 48, 57, 65, 70, 97, 102, 1, 0, 32,
		32, 227, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1,
		0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15,
		1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0,
		23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0,
		0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0,
		0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0,
		0, 0, 0, 47, 1, 0, 0, 0, 0, 49, 1, 0, 0, 0, 0, 51, 1, 0, 0, 0, 0, 53, 1,
		0, 0, 0, 0, 55, 1, 0, 0, 0, 0, 57, 1, 0, 0, 0, 0, 59, 1, 0, 0, 0, 0, 61,
		1, 0, 0, 0, 0, 63, 1, 0, 0, 0, 0, 65, 1, 0, 0, 0, 0, 67, 1, 0, 0, 0, 0,
		69, 1, 0, 0, 0, 0, 71, 1, 0, 0, 0, 0, 73, 1, 0, 0, 0, 1, 75, 1, 0, 0, 0,
		3, 77, 1, 0, 0, 0, 5, 82, 1, 0, 0, 0, 7, 85, 1, 0, 0, 0, 9, 87, 1, 0, 0,
		0, 11, 89, 1, 0, 0, 0, 13, 91, 1, 0, 0, 0, 15, 94, 1, 0, 0, 0, 17, 98,
		1, 0, 0, 0, 19, 102, 1, 0, 0, 0, 21, 106, 1, 0, 0, 0, 23, 110, 1, 0, 0,
		0, 25, 114, 1, 0, 0, 0, 27, 118, 1, 0, 0, 0, 29, 122, 1, 0, 0, 0, 31, 126,
		1, 0, 0, 0, 33, 130, 1, 0, 0, 0, 35, 133, 1, 0, 0, 0, 37, 137, 1, 0, 0,
		0, 39, 141, 1, 0, 0, 0, 41, 145, 1, 0, 0, 0, 43, 149, 1, 0, 0, 0, 45, 153,
		1, 0, 0, 0, 47, 158, 1, 0, 0, 0, 49, 162, 1, 0, 0, 0, 51, 164, 1, 0, 0,
		0, 53, 166, 1, 0, 0, 0, 55, 168, 1, 0, 0, 0, 57, 170, 1, 0, 0, 0, 59, 172,
		1, 0, 0, 0, 61, 174, 1, 0, 0, 0, 63, 176, 1, 0, 0, 0, 65, 183, 1, 0, 0,
		0, 67, 196, 1, 0, 0, 0, 69, 200, 1, 0, 0, 0, 71, 207, 1, 0, 0, 0, 73, 216,
		1, 0, 0, 0, 75, 76, 5, 9, 0, 0, 76, 2, 1, 0, 0, 0, 77, 78, 5, 32, 0, 0,
		78, 79, 5, 32, 0, 0, 79, 80, 5, 32, 0, 0, 80, 81, 5, 32, 0, 0, 81, 4, 1,
		0, 0, 0, 82, 83, 5, 32, 0, 0, 83, 84, 5, 32, 0, 0, 84, 6, 1, 0, 0, 0, 85,
		86, 5, 10, 0, 0, 86, 8, 1, 0, 0, 0, 87, 88, 5, 58, 0, 0, 88, 10, 1, 0,
		0, 0, 89, 90, 5, 32, 0, 0, 90, 12, 1, 0, 0, 0, 91, 92, 5, 44, 0, 0, 92,
		93, 5, 32, 0, 0, 93, 14, 1, 0, 0, 0, 94, 95, 5, 109, 0, 0, 95, 96, 5, 111,
		0, 0, 96, 97, 5, 118, 0, 0, 97, 16, 1, 0, 0, 0, 98, 99, 5, 97, 0, 0, 99,
		100, 5, 100, 0, 0, 100, 101, 5, 100, 0, 0, 101, 18, 1, 0, 0, 0, 102, 103,
		5, 97, 0, 0, 103, 104, 5, 100, 0, 0, 104, 105, 5, 99, 0, 0, 105, 20, 1,
		0, 0, 0, 106, 107, 5, 115, 0, 0, 107, 108, 5, 117, 0, 0, 108, 109, 5, 98,
		0, 0, 109, 22, 1, 0, 0, 0, 110, 111, 5, 115, 0, 0, 111, 112, 5, 98, 0,
		0, 112, 113, 5, 98, 0, 0, 113, 24, 1, 0, 0, 0, 114, 115, 5, 99, 0, 0, 115,
		116, 5, 109, 0, 0, 116, 117, 5, 112, 0, 0, 117, 26, 1, 0, 0, 0, 118, 119,
		5, 105, 0, 0, 119, 120, 5, 110, 0, 0, 120, 121, 5, 99, 0, 0, 121, 28, 1,
		0, 0, 0, 122, 123, 5, 100, 0, 0, 123, 124, 5, 101, 0, 0, 124, 125, 5, 99,
		0, 0, 125, 30, 1, 0, 0, 0, 126, 127, 5, 106, 0, 0, 127, 128, 5, 109, 0,
		0, 128, 129, 5, 112, 0, 0, 129, 32, 1, 0, 0, 0, 130, 131, 5, 106, 0, 0,
		131, 132, 5, 122, 0, 0, 132, 34, 1, 0, 0, 0, 133, 134, 5, 106, 0, 0, 134,
		135, 5, 110, 0, 0, 135, 136, 5, 122, 0, 0, 136, 36, 1, 0, 0, 0, 137, 138,
		5, 106, 0, 0, 138, 139, 5, 115, 0, 0, 139, 140, 5, 114, 0, 0, 140, 38,
		1, 0, 0, 0, 141, 142, 5, 114, 0, 0, 142, 143, 5, 116, 0, 0, 143, 144, 5,
		115, 0, 0, 144, 40, 1, 0, 0, 0, 145, 146, 5, 98, 0, 0, 146, 147, 5, 114,
		0, 0, 147, 148, 5, 107, 0, 0, 148, 42, 1, 0, 0, 0, 149, 150, 5, 110, 0,
		0, 150, 151, 5, 111, 0, 0, 151, 152, 5, 112, 0, 0, 152, 44, 1, 0, 0, 0,
		153, 154, 5, 112, 0, 0, 154, 155, 5, 117, 0, 0, 155, 156, 5, 115, 0, 0,
		156, 157, 5, 104, 0, 0, 157, 46, 1, 0, 0, 0, 158, 159, 5, 112, 0, 0, 159,
		160, 5, 111, 0, 0, 160, 161, 5, 112, 0, 0, 161, 48, 1, 0, 0, 0, 162, 163,
		5, 97, 0, 0, 163, 50, 1, 0, 0, 0, 164, 165, 5, 98, 0, 0, 165, 52, 1, 0,
		0, 0, 166, 167, 5, 99, 0, 0, 167, 54, 1, 0, 0, 0, 168, 169, 5, 100, 0,
		0, 169, 56, 1, 0, 0, 0, 170, 171, 5, 91, 0, 0, 171, 58, 1, 0, 0, 0, 172,
		173, 5, 93, 0, 0, 173, 60, 1, 0, 0, 0, 174, 175, 5, 43, 0, 0, 175, 62,
		1, 0, 0, 0, 176, 177, 5, 32, 0, 0, 177, 178, 5, 61, 0, 0, 178, 179, 5,
		32, 0, 0, 179, 64, 1, 0, 0, 0, 180, 182, 3, 73, 36, 0, 181, 180, 1, 0,
		0, 0, 182, 185, 1, 0, 0, 0, 183, 181, 1, 0, 0, 0, 183, 184, 1, 0, 0, 0,
		184, 186, 1, 0, 0, 0, 185, 183, 1, 0, 0, 0, 186, 190, 5, 59, 0, 0, 187,
		189, 8, 0, 0, 0, 188, 187, 1, 0, 0, 0, 189, 192, 1, 0, 0, 0, 190, 188,
		1, 0, 0, 0, 190, 191, 1, 0, 0, 0, 191, 193, 1, 0, 0, 0, 192, 190, 1, 0,
		0, 0, 193, 194, 6, 32, 0, 0, 194, 66, 1, 0, 0, 0, 195, 197, 7, 1, 0, 0,
		196, 195, 1, 0, 0, 0, 197, 198, 1, 0, 0, 0, 198, 196, 1, 0, 0, 0, 198,
		199, 1, 0, 0, 0, 199, 68, 1, 0, 0, 0, 200, 204, 7, 2, 0, 0, 201, 203, 7,
		3, 0, 0, 202, 201, 1, 0, 0, 0, 203, 206, 1, 0, 0, 0, 204, 202, 1, 0, 0,
		0, 204, 205, 1, 0, 0, 0, 205, 70, 1, 0, 0, 0, 206, 204, 1, 0, 0, 0, 207,
		208, 5, 48, 0, 0, 208, 209, 5, 120, 0, 0, 209, 211, 1, 0, 0, 0, 210, 212,
		7, 4, 0, 0, 211, 210, 1, 0, 0, 0, 212, 213, 1, 0, 0, 0, 213, 211, 1, 0,
		0, 0, 213, 214, 1, 0, 0, 0, 214, 72, 1, 0, 0, 0, 215, 217, 7, 5, 0, 0,
		216, 215, 1, 0, 0, 0, 217, 218, 1, 0, 0, 0, 218, 216, 1, 0, 0, 0, 218,
		219, 1, 0, 0, 0, 219, 220, 1, 0, 0, 0, 220, 221, 6, 36, 0, 0, 221, 74,
		1, 0, 0, 0, 7, 0, 183, 190, 198, 204, 213, 218, 1, 6, 0, 0,
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
	AsmE8LexerT__0         = 1
	AsmE8LexerT__1         = 2
	AsmE8LexerT__2         = 3
	AsmE8LexerT__3         = 4
	AsmE8LexerT__4         = 5
	AsmE8LexerT__5         = 6
	AsmE8LexerT__6         = 7
	AsmE8LexerT__7         = 8
	AsmE8LexerT__8         = 9
	AsmE8LexerT__9         = 10
	AsmE8LexerT__10        = 11
	AsmE8LexerT__11        = 12
	AsmE8LexerT__12        = 13
	AsmE8LexerT__13        = 14
	AsmE8LexerT__14        = 15
	AsmE8LexerT__15        = 16
	AsmE8LexerT__16        = 17
	AsmE8LexerT__17        = 18
	AsmE8LexerT__18        = 19
	AsmE8LexerT__19        = 20
	AsmE8LexerT__20        = 21
	AsmE8LexerT__21        = 22
	AsmE8LexerT__22        = 23
	AsmE8LexerT__23        = 24
	AsmE8LexerT__24        = 25
	AsmE8LexerT__25        = 26
	AsmE8LexerT__26        = 27
	AsmE8LexerT__27        = 28
	AsmE8LexerT__28        = 29
	AsmE8LexerT__29        = 30
	AsmE8LexerT__30        = 31
	AsmE8LexerT__31        = 32
	AsmE8LexerLINE_COMMENT = 33
	AsmE8LexerINT          = 34
	AsmE8LexerSTR          = 35
	AsmE8LexerHEX          = 36
	AsmE8LexerWHITE_SPACE  = 37
)
