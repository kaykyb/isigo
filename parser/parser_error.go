package parser

import (
	"fmt"
	"isigo/tokens"
)

type ParserError struct {
	err string
}

func (e *ParserError) Error() string {
	return fmt.Sprintf("[Syntax Error]: %s", e.err)
}

func alreadyDeclared(identifier string) *ParserError {
	return errorf("O símbolo '%s' já foi declarado neste contexto.", identifier)
}

func usedBeforeDeclaration(identifier string) *ParserError {
	return errorf("O símbolo '%s' foi usado antes de ser declarado.", identifier)
}

func noMatchTypeError(delta TokenDelta) *ParserError {
	return errorf("%s não é esperado neste contexto.", delta.token.FriendlyString())
}

func unexpectedTokenTypeError(delta TokenDelta, expectedType tokens.Type) *ParserError {
	return errorf("Esperava uma %s, encontrada %s '%s'.", tokens.FriendlyString(expectedType), delta.token.FriendlyString(), delta.token.Content())
}

func unexpectedContentError(delta TokenDelta, expectedContent string) *ParserError {
	return errorf("Esperava encontrar '%s', mas foi encontrado '%s'.", expectedContent, delta.token.Content())
}

func errorf(content string, a ...any) *ParserError {
	return &ParserError{
		err: fmt.Sprintf(content, a...),
	}
}
