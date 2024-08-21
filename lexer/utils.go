package lexer

import (
	"isigo/syntax"
)

func IsWhitespaceRune(c rune) bool {
	return c == rune(' ')
}

func IsNewLineRune(c rune) bool {
	return c == rune('\n')
}

func IsTabRune(c rune) bool {
	return c == rune('\t')
}

func IsCartridgeReturnRune(c rune) bool {
	return c == rune('\r')
}

func IsOperatorRune(c rune) bool {
	switch c {
	case '+', '-', '*', '/', '<', '>', '!', '=':
		return true
	default:
		return false
	}
}

func IsLetterRune(c rune) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z')
}

func IsDigitRune(c rune) bool {
	return c >= '0' && c <= '9'
}

func IsDecimalSeparator(c rune) bool {
	return c == '.'
}

func IsStatementTerminator(c rune) bool {
	return c == '.'
}

func IsColon(c rune) bool {
	return c == ':'
}

func IsEqualSign(c rune) bool {
	return c == '='
}

func IsOpenParenthesis(c rune) bool {
	return c == '('
}

func IsCloseParenthesis(c rune) bool {
	return c == ')'
}

func IsOpenBrace(c rune) bool {
	return c == '{'
}

func IsCloseBrace(c rune) bool {
	return c == '}'
}

func IsComma(c rune) bool {
	return c == ','
}

func IsStringDelimiter(c rune) bool {
	return c == '"'
}

func IsReservedWord(s string) bool {
	switch s {
	case syntax.Program, syntax.EndProgram, syntax.Declare, syntax.If, syntax.Then, syntax.Else, syntax.Read, syntax.Write:
		return true
	default:
		return false
	}
}

func IsTypeT(s string) bool {
	switch s {
	case syntax.IntegerT, syntax.FloatT, syntax.StringT:
		return true
	default:
		return false
	}
}
