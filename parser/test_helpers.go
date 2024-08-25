package parser

import (
	"bufio"
	"isigo/context"
	"isigo/lexer"
	"isigo/sources"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func SetupLPC(t *testing.T, str string) (*Parser, *context.Context) {
	reader := bufio.NewReader(strings.NewReader(str))
	sourceStream := sources.NewBuildReader(reader)

	l := lexer.New(sourceStream)
	p := New(&l)
	c := context.New()

	return &p, &c
}

func AssertNextToken(t *testing.T, p *Parser) TokenDelta {
	delta, err := p.nextToken()
	assert.NoError(t, err)

	return delta
}
