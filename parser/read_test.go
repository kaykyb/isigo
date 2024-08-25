package parser

import (
	"isigo/lang"
	"isigo/value_types"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRead(t *testing.T) {
	p, c := SetupLPC(t, "leia(a).")

	s, _ := c.CreateSymbol("a", value_types.IntegerValueTypeEntity)

	delta := AssertNextToken(t, p)

	read, delta, err := p.Read(c, delta)

	assert.NoError(t, err)
	assert.Equal(t, lang.NewRead(c, s), read)
}
