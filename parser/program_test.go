package parser

import (
	"isigo/common"
	"isigo/lang"
	"isigo/tokens"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProgram(t *testing.T) {
	p, c := SetupLPC(t, "programa escreva(1). fimprog.")

	delta := AssertNextToken(t, p)

	doc, delta, err := p.Prog(c, delta)

	assert.NoError(t, err)
	assert.IsType(t, lang.Program{}, doc)
	assert.IsType(t, lang.Block{}, doc.Child())

	assert.Equal(t, NewTokenDelta(tokens.NewStatementTerminator("."), common.NewCodePosition(28, 0, 28)), delta)
}
