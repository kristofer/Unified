package parser

import (
	"fmt"
	"strings"

	"github.com/antlr4-go/antlr/v4"
)

// CustomErrorListener suppresses "missing ';'" errors since semicolons are optional in Unified
type CustomErrorListener struct {
	*antlr.DefaultErrorListener
}

func NewCustomErrorListener() *CustomErrorListener {
	return &CustomErrorListener{
		DefaultErrorListener: antlr.NewDefaultErrorListener(),
	}
}

func (c *CustomErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	// Suppress "missing ';'" errors since semicolons are optional
	if strings.Contains(msg, "missing ';'") {
		return
	}
	
	// Report all other errors
	fmt.Printf("line %d:%d %s\n", line, column, msg)
}
