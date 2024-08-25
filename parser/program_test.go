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

	doc, delta, err := p.Prog(c)

	assert.NoError(t, err)
	assert.IsType(t, lang.Program{}, doc)
	assert.IsType(t, lang.Block{}, doc.Child())

	assert.Equal(t, NewTokenDelta(tokens.NewStatementTerminator("."), common.NewCodePosition(28, 0, 28)), delta)
}
