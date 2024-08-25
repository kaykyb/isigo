package parser

import (
	"isigo/lang"
	"isigo/value_types"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeclareOrComand(t *testing.T) {
	p, c := SetupLPC(t, "declare a inteiro. leia(a). }")
	delta := AssertNextToken(t, p)

	doc, delta, err := p.DeclareOrCommand(c, delta)

	assert.NoError(t, err)
	assert.IsType(t, lang.VariableContext{}, doc)
}

func TestDeclareOrComandRead(t *testing.T) {
	p, c := SetupLPC(t, "leia(a).")
	c.CreateSymbol("a", value_types.IntegerValueTypeEntity)
	delta := AssertNextToken(t, p)

	doc, delta, err := p.DeclareOrCommand(c, delta)

	assert.NoError(t, err)
	assert.IsType(t, lang.Read{}, doc)
}

func TestDeclareOrComandWrite(t *testing.T) {
	p, c := SetupLPC(t, "escreva(1).")
	delta := AssertNextToken(t, p)

	doc, delta, err := p.DeclareOrCommand(c, delta)

	assert.NoError(t, err)
	assert.IsType(t, lang.Write{}, doc)
}

func TestDeclareOrComandAssignment(t *testing.T) {
	p, c := SetupLPC(t, "a := 1.")
	c.CreateSymbol("a", value_types.IntegerValueTypeEntity)
	delta := AssertNextToken(t, p)

	doc, delta, err := p.DeclareOrCommand(c, delta)

	assert.NoError(t, err)
	assert.IsType(t, lang.Assignment{}, doc)
}
