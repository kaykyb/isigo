package parser

import (
	"isigo/lang"
	"isigo/value_types"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExecutionContextDeclare(t *testing.T) {
	p, c := SetupLPC(t, "declare a inteiro. leia(a). }")
	delta := AssertNextToken(t, p)

	doc, delta, err := p.ExecutionContext(c, delta)

	assert.NoError(t, err)
	assert.IsType(t, lang.ExecutionContext{}, doc)

	assert.Len(t, doc.Children(), 1)
	assert.IsType(t, lang.VariableContext{}, doc.Children()[0])
}

func TestExecutionContextCommand(t *testing.T) {
	p, c := SetupLPC(t, "a := 1. }")
	c.CreateSymbol("a", value_types.IntegerValueTypeEntity)

	delta := AssertNextToken(t, p)

	doc, delta, err := p.ExecutionContext(c, delta)

	assert.NoError(t, err)
	assert.IsType(t, lang.ExecutionContext{}, doc)

	assert.Len(t, doc.Children(), 1)
	assert.IsType(t, lang.Assignment{}, doc.Children()[0])
}
