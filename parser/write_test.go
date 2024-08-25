package parser

import (
	"isigo/lang"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWrite(t *testing.T) {
	p, c := SetupLPC(t, "escreva(1).")

	delta := AssertNextToken(t, p)

	read, delta, err := p.Write(c, delta)

	expectedWrite := lang.NewWrite(c, lang.NewTermExpr(c, lang.NewFactorTerm(c, lang.NewIntegerFactor(c, 1))))

	assert.NoError(t, err)
	assert.Equal(t, expectedWrite, read)
}
