package lexer

import (
	"isigo/common"
	"isigo/tokens"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTokenizeWhitespacesString(t *testing.T) {
	l := NewStrLexer(t, " \n\t ")
	tk, pos, err := l.NextToken()

	assert.NoError(t, err)
	assert.Equal(t, tokens.NewEOF(""), tk)
	assert.Equal(t, common.NewCodePosition(4, 1, 5), pos)
}

func TestTokenizeNewlinesString(t *testing.T) {
	l := NewStrLexer(t, "\n\n\n")
	tk, pos, err := l.NextToken()

	assert.NoError(t, err)
	assert.Equal(t, tokens.NewEOF(""), tk)
	assert.Equal(t, common.NewCodePosition(3, 3, 0), pos)
}

func TestTokenizeTabsString(t *testing.T) {
	l := NewStrLexer(t, "\t\t\t")
	tk, pos, err := l.NextToken()

	assert.NoError(t, err)
	assert.Equal(t, tokens.NewEOF(""), tk)
	assert.Equal(t, common.NewCodePosition(3, 0, 12), pos)
}

func TestTokenizeCarriageReturnsString(t *testing.T) {
	l := NewStrLexer(t, "\r\r\r")
	tk, pos, err := l.NextToken()

	assert.NoError(t, err)
	assert.Equal(t, tokens.NewEOF(""), tk)
	assert.Equal(t, common.NewCodePosition(3, 0, 0), pos)
}

func TestTokenizeOperatorString(t *testing.T) {
	l := NewStrLexer(t, "+ - * / > < >= <= != ==")
	tk, pos, err := l.NextToken()

	assert.NoError(t, err)
	assert.Equal(t, tokens.NewOperator("+"), tk)
	assert.Equal(t, common.NewCodePosition(0, 0, 0), pos)

	tk, pos, err = l.NextToken()
	assert.NoError(t, err)
	assert.Equal(t, tokens.NewOperator("-"), tk)
	assert.Equal(t, common.NewCodePosition(2, 0, 2), pos)

	tk, pos, err = l.NextToken()
	assert.NoError(t, err)
	assert.Equal(t, tokens.NewOperator("*"), tk)
	assert.Equal(t, common.NewCodePosition(4, 0, 4), pos)

	tk, pos, err = l.NextToken()
	assert.NoError(t, err)
	assert.Equal(t, tokens.NewOperator("/"), tk)
	assert.Equal(t, common.NewCodePosition(6, 0, 6), pos)

	tk, pos, err = l.NextToken()
	assert.NoError(t, err)
	assert.Equal(t, tokens.NewOperator(">"), tk)
	assert.Equal(t, common.NewCodePosition(8, 0, 8), pos)

	tk, pos, err = l.NextToken()
	assert.NoError(t, err)
	assert.Equal(t, tokens.NewOperator("<"), tk)
	assert.Equal(t, common.NewCodePosition(10, 0, 10), pos)

	tk, pos, err = l.NextToken()
	assert.NoError(t, err)
	assert.Equal(t, tokens.NewOperator(">="), tk)
	assert.Equal(t, common.NewCodePosition(12, 0, 12), pos)

	tk, pos, err = l.NextToken()
	assert.NoError(t, err)
	assert.Equal(t, tokens.NewOperator("<="), tk)
	assert.Equal(t, common.NewCodePosition(15, 0, 15), pos)

	tk, pos, err = l.NextToken()
	assert.NoError(t, err)
	assert.Equal(t, tokens.NewOperator("!="), tk)
	assert.Equal(t, common.NewCodePosition(18, 0, 18), pos)

	tk, pos, err = l.NextToken()
	assert.NoError(t, err)
	assert.Equal(t, tokens.NewOperator("=="), tk)
	assert.Equal(t, common.NewCodePosition(21, 0, 21), pos)

	tk, pos, err = l.NextToken()
	assert.NoError(t, err)
	assert.Equal(t, tokens.NewEOF(""), tk)
	assert.Equal(t, common.NewCodePosition(23, 0, 23), pos)
}

func TestTokenizeIdentifierString(t *testing.T) {
	l := NewStrLexer(t, "hello world")
	tk, pos, err := l.NextToken()

	assert.NoError(t, err)
	assert.Equal(t, tokens.NewIdentifier("hello"), tk)
	assert.Equal(t, common.NewCodePosition(0, 0, 0), pos)

	tk, pos, err = l.NextToken()
	assert.NoError(t, err)
	assert.Equal(t, tokens.NewIdentifier("world"), tk)
	assert.Equal(t, common.NewCodePosition(6, 0, 6), pos)

	tk, pos, err = l.NextToken()
	assert.NoError(t, err)
	assert.Equal(t, tokens.NewEOF(""), tk)
	assert.Equal(t, common.NewCodePosition(11, 0, 11), pos)
}

func TestTokenizeIntegerString(t *testing.T) {
	l := NewStrLexer(t, "123 456 789")
	tk, pos, err := l.NextToken()

	assert.NoError(t, err)
	assert.Equal(t, tokens.NewInteger("123"), tk)
	assert.Equal(t, common.NewCodePosition(0, 0, 0), pos)

	tk, pos, err = l.NextToken()
	assert.NoError(t, err)
	assert.Equal(t, tokens.NewInteger("456"), tk)
	assert.Equal(t, common.NewCodePosition(4, 0, 4), pos)

	tk, pos, err = l.NextToken()
	assert.NoError(t, err)
	assert.Equal(t, tokens.NewInteger("789"), tk)
	assert.Equal(t, common.NewCodePosition(8, 0, 8), pos)

	tk, pos, err = l.NextToken()
	assert.NoError(t, err)
	assert.Equal(t, tokens.NewEOF(""), tk)
	assert.Equal(t, common.NewCodePosition(11, 0, 11), pos)
}

func TestTokenizeFloatString(t *testing.T) {
	l := NewStrLexer(t, "123.45 678.90 987.65")
	tk, pos, err := l.NextToken()

	assert.NoError(t, err)
	assert.Equal(t, tokens.NewDecimal("123.45"), tk)
	assert.Equal(t, common.NewCodePosition(0, 0, 0), pos)

	tk, pos, err = l.NextToken()
	assert.NoError(t, err)
	assert.Equal(t, tokens.NewDecimal("678.90"), tk)
	assert.Equal(t, common.NewCodePosition(7, 0, 7), pos)

	tk, pos, err = l.NextToken()
	assert.NoError(t, err)
	assert.Equal(t, tokens.NewDecimal("987.65"), tk)
	assert.Equal(t, common.NewCodePosition(14, 0, 14), pos)
}

func TestTokenizeStatementTerminatorString(t *testing.T) {
	l := NewStrLexer(t, ".")
	tk, pos, err := l.NextToken()

	assert.NoError(t, err)
	assert.Equal(t, tokens.NewStatementTerminator("."), tk)
	assert.Equal(t, common.NewCodePosition(0, 0, 0), pos)
}

func TestNoConflictStatementTerminatorDecimalString(t *testing.T) {
	l := NewStrLexer(t, "2.56.")
	tk, pos, err := l.NextToken()

	assert.NoError(t, err)
	assert.Equal(t, tokens.NewDecimal("2.56"), tk)
	assert.Equal(t, common.NewCodePosition(0, 0, 0), pos)

	tk, pos, err = l.NextToken()
	assert.NoError(t, err)
	assert.Equal(t, tokens.NewStatementTerminator("."), tk)
	assert.Equal(t, common.NewCodePosition(4, 0, 4), pos)
}

func TestTokenizeAssignmentOperatorString(t *testing.T) {
	l := NewStrLexer(t, ":=")
	tk, pos, err := l.NextToken()

	assert.NoError(t, err)
	assert.Equal(t, tokens.NewAssign(":="), tk)
	assert.Equal(t, common.NewCodePosition(0, 0, 0), pos)
}

func TestTokenizeAssignmentOperatorStringFail(t *testing.T) {
	l := NewStrLexer(t, ":")
	_, pos, err := l.NextToken()

	assert.Error(t, err)
	assert.Equal(t, common.NewCodePosition(2, 0, 2), pos)
}

func TestTokenizeParenthesesString(t *testing.T) {
	l := NewStrLexer(t, "()")
	tk, pos, err := l.NextToken()

	assert.NoError(t, err)
	assert.Equal(t, tokens.NewOpenParenthesis("("), tk)
	assert.Equal(t, common.NewCodePosition(0, 0, 0), pos)

	tk, pos, err = l.NextToken()
	assert.NoError(t, err)
	assert.Equal(t, tokens.NewCloseParenthesis(")"), tk)
	assert.Equal(t, common.NewCodePosition(1, 0, 1), pos)
}

func TestTokenizeBracesString(t *testing.T) {
	l := NewStrLexer(t, "{}")
	tk, pos, err := l.NextToken()

	assert.NoError(t, err)
	assert.Equal(t, tokens.NewOpenBrace("{"), tk)
	assert.Equal(t, common.NewCodePosition(0, 0, 0), pos)

	tk, pos, err = l.NextToken()
	assert.NoError(t, err)
	assert.Equal(t, tokens.NewCloseBrace("}"), tk)
	assert.Equal(t, common.NewCodePosition(1, 0, 1), pos)
}

func TestTokenizeCommaString(t *testing.T) {
	l := NewStrLexer(t, ",")
	tk, pos, err := l.NextToken()

	assert.NoError(t, err)
	assert.Equal(t, tokens.NewSeparator(","), tk)
	assert.Equal(t, common.NewCodePosition(0, 0, 0), pos)
}

func TestTokenizeStringString(t *testing.T) {
	l := NewStrLexer(t, `"Hello, world!" "This is a test."`)
	tk, pos, err := l.NextToken()

	assert.NoError(t, err)
	assert.Equal(t, tokens.NewString("Hello, world!"), tk)
	assert.Equal(t, common.NewCodePosition(0, 0, 0), pos)

	tk, pos, err = l.NextToken()
	assert.NoError(t, err)
	assert.Equal(t, tokens.NewString("This is a test."), tk)
	assert.Equal(t, common.NewCodePosition(16, 0, 16), pos)
}

func TestTokenizeStringStringFail(t *testing.T) {
	l := NewStrLexer(t, `"Hello, world!`)
	_, pos, err := l.NextToken()

	assert.Error(t, err)
	assert.Equal(t, common.NewCodePosition(14, 0, 14), pos)
}

func TestTokenizeFail(t *testing.T) {
	l := NewStrLexer(t, `&`)
	_, pos, err := l.NextToken()

	assert.Error(t, err)
	assert.Equal(t, common.NewCodePosition(0, 0, 0), pos)
}
