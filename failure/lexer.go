package failure

import "fmt"

type Lexer struct {
	err string
}

func (e Lexer) Error() string {
	return fmt.Sprintf("[Lexical Error]: %s", e.err)
}

func (e Lexer) Environment() Environment {
	return LexerEnvironment
}

func LexerErrorf(format string, args ...interface{}) error {
	return Lexer{
		err: fmt.Sprintf(format, args...),
	}
}

func ExpectedEndQuote() error {
	return LexerErrorf("Final de string '\"' esperado, mas não foi encontrado.")
}

func MalformedAssignmentOperator(c rune) error {
	return LexerErrorf("Operador de atribuição mal-formado. Esperado '=', encontrado %c", c)
}

func UnexpectedCharacter(c rune) error {
	return LexerErrorf("'%c' não é um caractere esperado neste contexto.", c)
}
