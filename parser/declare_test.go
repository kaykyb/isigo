package parser

import (
	"isigo/lang"
	"isigo/value_types"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeclare(t *testing.T) {
	p, c := SetupLPC(t, "declare a inteiro.")
	delta := AssertNextToken(t, p)

	doc, delta, err := p.Declare(c, delta)

	expectedDeclare := lang.NewDeclare(c, []lang.Variable{
		lang.NewVariable(c, "a", value_types.IntegerValueTypeEntity),
	})

	assert.NoError(t, err)
	assert.Equal(t, expectedDeclare, doc)

	_, err = c.RetrieveSymbol("a")
	assert.NoError(t, err)
}

func TestDeclareMultipleVariables(t *testing.T) {
	p, c := SetupLPC(t, "declare a inteiro, b inteiro.")
	delta := AssertNextToken(t, p)

	doc, delta, err := p.Declare(c, delta)

	expectedDeclare := lang.NewDeclare(c, []lang.Variable{
		lang.NewVariable(c, "a", value_types.IntegerValueTypeEntity),
		lang.NewVariable(c, "b", value_types.IntegerValueTypeEntity),
	})

	assert.NoError(t, err)
	assert.Equal(t, expectedDeclare, doc)

	_, err = c.RetrieveSymbol("a")
	assert.NoError(t, err)

	_, err = c.RetrieveSymbol("b")
	assert.NoError(t, err)
}
