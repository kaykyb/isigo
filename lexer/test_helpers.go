package lexer

import (
	"bufio"
	"isigo/sources"
	"strings"
	"testing"
)

func NewStrLexer(t *testing.T, str string) Lexer {
	reader := bufio.NewReader(strings.NewReader(str))
	sourceStream := sources.NewBuildReader(reader)

	return New(sourceStream)
}
