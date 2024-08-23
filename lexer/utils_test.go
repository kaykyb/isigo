package lexer_test

import (
	"testing"

	"isigo/lexer"
	"isigo/syntax"

	"github.com/stretchr/testify/assert"
)

func TestIsWhitespaceRune(t *testing.T) {
	assert.True(t, lexer.IsWhitespaceRune(' '))
	assert.False(t, lexer.IsWhitespaceRune('a'))
}

func TestIsNewLineRune(t *testing.T) {
	assert.True(t, lexer.IsNewLineRune('\n'))
	assert.False(t, lexer.IsNewLineRune('a'))
}

func TestIsTabRune(t *testing.T) {
	assert.True(t, lexer.IsTabRune('\t'))
	assert.False(t, lexer.IsTabRune('a'))
}

func TestIsCartridgeReturnRune(t *testing.T) {
	assert.True(t, lexer.IsCartridgeReturnRune('\r'))
	assert.False(t, lexer.IsCartridgeReturnRune('a'))
}

func TestIsOperatorRune(t *testing.T) {
	assert.True(t, lexer.IsOperatorRune('+'))
	assert.True(t, lexer.IsOperatorRune('='))
	assert.False(t, lexer.IsOperatorRune('&'))
}

func TestIsLetterRune(t *testing.T) {
	assert.True(t, lexer.IsLetterRune('a'))
	assert.True(t, lexer.IsLetterRune('Z'))
	assert.False(t, lexer.IsLetterRune('1'))
	assert.False(t, lexer.IsLetterRune('@'))
}

func TestIsDigitRune(t *testing.T) {
	assert.True(t, lexer.IsDigitRune('0'))
	assert.True(t, lexer.IsDigitRune('9'))
	assert.False(t, lexer.IsDigitRune('a'))
	assert.False(t, lexer.IsDigitRune('-'))
}

func TestIsDecimalSeparator(t *testing.T) {
	assert.True(t, lexer.IsDecimalSeparator('.'))
	assert.False(t, lexer.IsDecimalSeparator(','))
}

func TestIsStatementTerminator(t *testing.T) {
	assert.True(t, lexer.IsStatementTerminator('.'))
	assert.False(t, lexer.IsStatementTerminator(';'))
}

func TestIsColon(t *testing.T) {
	assert.True(t, lexer.IsColon(':'))
	assert.False(t, lexer.IsColon(';'))
}

func TestIsOpenParenthesis(t *testing.T) {
	assert.True(t, lexer.IsOpenParenthesis('('))
	assert.False(t, lexer.IsOpenParenthesis(')'))
}

func TestIsCloseParenthesis(t *testing.T) {
	assert.True(t, lexer.IsCloseParenthesis(')'))
	assert.False(t, lexer.IsCloseParenthesis('('))
}

func TestIsOpenBrace(t *testing.T) {
	assert.True(t, lexer.IsOpenBrace('{'))
	assert.False(t, lexer.IsOpenBrace('}'))
}

func TestIsCloseBrace(t *testing.T) {
	assert.True(t, lexer.IsCloseBrace('}'))
	assert.False(t, lexer.IsCloseBrace('{'))
}

func TestIsComma(t *testing.T) {
	assert.True(t, lexer.IsComma(','))
	assert.False(t, lexer.IsComma('.'))
}

func TestIsStringDelimiter(t *testing.T) {
	assert.True(t, lexer.IsStringDelimiter('"'))
	assert.False(t, lexer.IsStringDelimiter('\''))
}

func TestIsReservedWord(t *testing.T) {
	assert.True(t, lexer.IsReservedWord(syntax.Program))
	assert.True(t, lexer.IsReservedWord(syntax.Read))
	assert.False(t, lexer.IsReservedWord("customWord"))
}

func TestIsTypeT(t *testing.T) {
	assert.True(t, lexer.IsTypeT(syntax.IntegerT))
	assert.True(t, lexer.IsTypeT(syntax.FloatT))
	assert.False(t, lexer.IsTypeT("customType"))
}
