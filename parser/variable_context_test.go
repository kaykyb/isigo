package parser

import (
	"isigo/lang"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVariableContextDeclare(t *testing.T) {
	p, c := SetupLPC(t, "declare a inteiro. leia(a). }")
	delta := AssertNextToken(t, p)

	doc, delta, err := p.VariableContext(c, delta)

	assert.NoError(t, err)
	assert.IsType(t, lang.VariableContext{}, doc)
}
