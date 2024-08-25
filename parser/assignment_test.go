package parser

import (
	"isigo/common"
	"isigo/lang"
	"isigo/tokens"
	"isigo/value_types"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAssignmentCanAssignToDeclaredVariable(t *testing.T) {
	p, c := SetupLPC(t, "a := 1.")

	s, _ := c.CreateSymbol("a", value_types.IntegerValueTypeEntity)

	delta, err := p.nextToken()

	assert.NoError(t, err)

	assignment, delta, err := p.Assignment(c, delta)

	expectedExpr := lang.NewTermExpr(c, lang.NewFactorTerm(c, lang.NewIntegerFactor(c, 1)))

	assert.NoError(t, err)
	assert.Equal(t, lang.NewAssignment(c, s, expectedExpr), assignment)
	assert.Equal(t, NewTokenDelta(tokens.NewEOF(""), common.NewCodePosition(7, 0, 7)), delta)
}

func TestAssignmentCannotAssignToNotDeclaredVariable(t *testing.T) {
	p, c := SetupLPC(t, "a := 1.")

	delta, err := p.nextToken()

	assert.NoError(t, err)

	_, _, err = p.Assignment(c, delta)
	assert.Error(t, err)
}

func TestMalformedAssignmentNoStatementTerminator(t *testing.T) {
	p, c := SetupLPC(t, "a := 1")

	delta, err := p.nextToken()

	assert.NoError(t, err)

	_, _, err = p.Assignment(c, delta)
	assert.Error(t, err)
}
