package error_listener

import (
	"fmt"

	"github.com/antlr4-go/antlr/v4"
)

type CustomErrorListener struct {
	*antlr.DefaultErrorListener
	Errors []string
}

func NewCustomErrorListener() *CustomErrorListener {
	return &CustomErrorListener{
		Errors: make([]string, 0, 2),
	}
}

func (l *CustomErrorListener) SyntaxError(_ antlr.Recognizer, _ interface{}, line, column int, msg string, _ antlr.RecognitionException) {
	fmt.Println("--")
	errorMessage := fmt.Sprintf("line %d:%d %s", line, column, msg)
	l.Errors = append(l.Errors, errorMessage)
}

// ReportAmbiguity prints readable info about parser ambiguities
func (l *CustomErrorListener) ReportAmbiguity(prs antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, exact bool, ambigAlts *antlr.BitSet, configs *antlr.ATNConfigSet) {
	// // Get the specific token where ambiguity starts
	// token := prs.GetTokenStream().Get(startIndex)

	// fmt.Printf("Ambiguity at line %d, position %d. Token: %q\n",
	// 	token.GetLine(),
	// 	token.GetColumn(),
	// 	token.GetText(),
	// )
}

func (l *CustomErrorListener) ReportAttemptingFullContext(prs antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, conflictingAlts *antlr.BitSet, configs *antlr.ATNConfigSet) {
	// // Get the token where the parser decided it needs to "think harder"
	// token := prs.GetTokenStream().Get(startIndex)

	// fmt.Printf("Attempting Full Context at line %d:%d. Input: %q\n",
	// 	token.GetLine(),
	// 	token.GetColumn(),
	// 	token.GetText(),
	// )
}

func (l *CustomErrorListener) ReportContextSensitivity(_ antlr.Parser, _ *antlr.DFA, _ int, _ int, _ int, _ *antlr.ATNConfigSet) {
}

func (l *CustomErrorListener) NewSimpleError(msg string, line, column int) {
	errorMessage := fmt.Sprintf("line %d:%d %s", line, column, msg)
	l.Errors = append(l.Errors, errorMessage)
}

func (l *CustomErrorListener) GetError() error {
	if len(l.Errors) == 0 {
		return nil
	}
	s := "Assemble failed.\r\n"
	s += "Syntax Error(s):\r\n"
	for _, e := range l.Errors {
		s += "\t" + e + "\r\n"
	}
	return WrapErrorMessages(l.Errors...)
}
