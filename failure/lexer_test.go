package failure

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLexerErrorf(t *testing.T) {
	err := LexerErrorf("unexpected error: %s", "test")
	expectedErrorMessage := "[Lexical Error]: unexpected error: test"
	assert.Equal(t, expectedErrorMessage, err.Error(), "LexerErrorf should format the error message correctly.")
}

func TestExpectedEndQuote(t *testing.T) {
	err := ExpectedEndQuote()
	expectedErrorMessage := "[Lexical Error]: Final de string '\"' esperado, mas não foi encontrado."
	assert.Equal(t, expectedErrorMessage, err.Error(), "ExpectedEndQuote should return the correct error message.")
}

func TestMalformedAssignmentOperator(t *testing.T) {
	err := MalformedAssignmentOperator('=')
	expectedErrorMessage := "[Lexical Error]: Operador de atribuição mal-formado. Esperado '=', encontrado ="
	assert.Equal(t, expectedErrorMessage, err.Error(), "MalformedAssignmentOperator should return the correct error message.")
}

func TestUnexpectedCharacter(t *testing.T) {
	err := UnexpectedCharacter('x')
	expectedErrorMessage := "[Lexical Error]: 'x' não é um caractere esperado neste contexto."
	assert.Equal(t, expectedErrorMessage, err.Error(), "UnexpectedCharacter should return the correct error message.")
}

func TestLexerEnvironment(t *testing.T) {
	lexer := Lexer{}
	assert.Equal(t, LexerEnvironment, lexer.Environment(), "Environment should return the correct LexerEnvironment.")
}
