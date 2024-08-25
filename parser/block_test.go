package parser

import (
	"isigo/common"
	"isigo/lang"
	"isigo/tokens"
	"isigo/value_types"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBlockExecutionContext(t *testing.T) {
	p, c := SetupLPC(t, "a := 1. }")

	c.CreateSymbol("a", value_types.IntegerValueTypeEntity)

	delta, err := p.nextToken()
	assert.NoError(t, err)

	block, delta, err := p.Block(c, delta)

	assert.NoError(t, err)

	assert.IsType(t, lang.Block{}, block)
	assert.IsType(t, lang.ExecutionContext{}, block.Child())

	assert.Equal(t, NewTokenDelta(tokens.NewCloseBrace("}"), common.NewCodePosition(8, 0, 8)), delta)
}

func TestBlockDeclareContext(t *testing.T) {
	p, c := SetupLPC(t, "declare a inteiro. a := 1. }")

	delta, err := p.nextToken()
	assert.NoError(t, err)

	block, delta, err := p.Block(c, delta)

	assert.NoError(t, err)

	assert.IsType(t, lang.Block{}, block)
	assert.IsType(t, lang.VariableContext{}, block.Child())

	assert.Equal(t, NewTokenDelta(tokens.NewCloseBrace("}"), common.NewCodePosition(27, 0, 27)), delta)
}

func TestBlockExecutionContextProgram(t *testing.T) {
	p, c := SetupLPC(t, "a := 1. fimprog.")

	c.CreateSymbol("a", value_types.IntegerValueTypeEntity)

	delta, err := p.nextToken()
	assert.NoError(t, err)

	block, delta, err := p.Block(c, delta)

	assert.NoError(t, err)

	assert.IsType(t, lang.Block{}, block)
	assert.IsType(t, lang.ExecutionContext{}, block.Child())

	assert.Equal(t, NewTokenDelta(tokens.NewReservedWord("fimprog"), common.NewCodePosition(8, 0, 8)), delta)
}
