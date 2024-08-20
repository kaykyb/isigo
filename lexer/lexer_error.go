package lexer

import "fmt"

type LexerError struct {
	err string
}

func (e *LexerError) Error() string {
	return fmt.Sprintf("[Analisador Léxico] %s", e.err)
}
