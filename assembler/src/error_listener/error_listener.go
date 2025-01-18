package error_listener

import (
	"errors"
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

func (l *CustomErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
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
	return errors.New(s)
}
