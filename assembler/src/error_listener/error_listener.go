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
	errorMessage := fmt.Sprintf("line %d:%d %s", line, column, msg)
	l.Errors = append(l.Errors, errorMessage)
}

func (l *CustomErrorListener) ReportAmbiguity(_ antlr.Parser, _ *antlr.DFA, _ int, _ int, _ bool, _ *antlr.BitSet, _ *antlr.ATNConfigSet) {
}
func (l *CustomErrorListener) ReportAttemptingFullContext(_ antlr.Parser, _ *antlr.DFA, _ int, _ int, _ *antlr.BitSet, _ *antlr.ATNConfigSet) {
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
