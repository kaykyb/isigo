package lexer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsWhitespaceRune(t *testing.T) {
	assert.True(t, IsWhitespaceRune(' '))
	assert.False(t, IsWhitespaceRune('\r'))
}

func TestIsNewLineRune(t *testing.T) {
	assert.True(t, IsNewLineRune('\n'))
	assert.False(t, IsNewLineRune('\r'))
}

func TestIsTabRune(t *testing.T) {
	assert.True(t, IsTabRune('\t'))
	assert.False(t, IsTabRune('\r'))
}

func TestIsCartridgeReturn(t *testing.T) {
	assert.True(t, IsCartridgeReturnRune('\r'))
	assert.False(t, IsCartridgeReturnRune('\n'))
}

func TestIsOperatorRune(t *testing.T) {
	assert.True(t, IsOperatorRune('+'))
	assert.True(t, IsOperatorRune('-'))
	assert.True(t, IsOperatorRune('*'))
	assert.True(t, IsOperatorRune('/'))
	assert.True(t, IsOperatorRune('<'))
	assert.True(t, IsOperatorRune('>'))
	assert.True(t, IsOperatorRune('!'))
	assert.True(t, IsOperatorRune('='))

	assert.False(t, IsOperatorRune('a'))
	assert.False(t, IsOperatorRune('@'))
	assert.False(t, IsOperatorRune('`'))
	assert.False(t, IsOperatorRune('m'))
	assert.False(t, IsOperatorRune('n'))
	assert.False(t, IsOperatorRune('x'))
	assert.False(t, IsOperatorRune('y'))
	assert.False(t, IsOperatorRune(':'))
}

func TestIsLetterRune(t *testing.T) {
	assert.True(t, IsLetterRune('a'))
	assert.True(t, IsLetterRune('b'))
	assert.True(t, IsLetterRune('z'))

	assert.True(t, IsLetterRune('A'))
	assert.True(t, IsLetterRune('B'))
	assert.True(t, IsLetterRune('Z'))

	assert.False(t, IsLetterRune('1'))
	assert.False(t, IsLetterRune('2'))
	assert.False(t, IsLetterRune('0'))
}

func TestIsDigitRune(t *testing.T) {
	assert.True(t, IsDigitRune('0'))
	assert.True(t, IsDigitRune('1'))
	assert.True(t, IsDigitRune('2'))
	assert.True(t, IsDigitRune('3'))
	assert.True(t, IsDigitRune('4'))
	assert.True(t, IsDigitRune('5'))
	assert.True(t, IsDigitRune('6'))
	assert.True(t, IsDigitRune('7'))
	assert.True(t, IsDigitRune('8'))
	assert.True(t, IsDigitRune('9'))

	assert.False(t, IsDigitRune('a'))
	assert.False(t, IsDigitRune('/'))
	assert.False(t, IsDigitRune('z'))
	assert.False(t, IsDigitRune('A'))
	assert.False(t, IsDigitRune('*'))
	assert.False(t, IsDigitRune('Z'))
}

func TestIsDecimalSeparator(t *testing.T) {
	assert.True(t, IsDecimalSeparator('.'))
	assert.False(t, IsDecimalSeparator(','))
}

func TestIsStatementTerminator(t *testing.T) {
	assert.True(t, IsStatementTerminator('.'))
	assert.False(t, IsStatementTerminator(','))
	assert.False(t, IsStatementTerminator(';'))
}
