package parser

import (
	"isigo/context"
	"isigo/lexer"
	"testing"

	"github.com/stretchr/testify/assert"
)

func SetupLPC(t *testing.T, str string) (*Parser, *context.Context) {
	l := lexer.New(str)
	p := New(&l)
	c := context.New()

	return &p, &c
}

func AssertNextToken(t *testing.T, p *Parser) TokenDelta {
	delta, err := p.nextToken()
	assert.NoError(t, err)

	return delta
}
