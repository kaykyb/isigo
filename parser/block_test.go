package parser

import (
	"isigo/common"
	"isigo/context"
	"isigo/lang"
	"isigo/lexer"
	"isigo/tokens"
	"isigo/value_types"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBlockExecutionContext(t *testing.T) {
	l := lexer.New("a := 1. }")
	p := New(&l)

	c := context.New()
	c.CreateSymbol("a", value_types.IntegerValueTypeEntity)

	delta, err := p.nextToken()
	assert.NoError(t, err)

	block, delta, err := p.Block(&c, delta)

	assert.NoError(t, err)

	assert.IsType(t, lang.Block{}, block)
	assert.IsType(t, lang.ExecutionContext{}, block.Child())

	assert.Equal(t, NewTokenDelta(tokens.NewCloseBrace("}"), common.NewCodePosition(8, 0, 8)), delta)
}

func TestBlockDeclareContext(t *testing.T) {
	l := lexer.New("declare a inteiro. a := 1. }")
	p := New(&l)

	c := context.New()

	delta, err := p.nextToken()
	assert.NoError(t, err)

	block, delta, err := p.Block(&c, delta)

	assert.NoError(t, err)

	assert.IsType(t, lang.Block{}, block)
	assert.IsType(t, lang.VariableContext{}, block.Child())

	assert.Equal(t, NewTokenDelta(tokens.NewCloseBrace("}"), common.NewCodePosition(27, 0, 27)), delta)
}

func TestBlockExecutionContextProgram(t *testing.T) {
	l := lexer.New("a := 1. fimprog.")
	p := New(&l)

	c := context.New()
	c.CreateSymbol("a", value_types.IntegerValueTypeEntity)

	delta, err := p.nextToken()
	assert.NoError(t, err)

	block, delta, err := p.Block(&c, delta)

	assert.NoError(t, err)

	assert.IsType(t, lang.Block{}, block)
	assert.IsType(t, lang.ExecutionContext{}, block.Child())

	assert.Equal(t, NewTokenDelta(tokens.NewReservedWord("fimprog"), common.NewCodePosition(8, 0, 8)), delta)
}
