package lang_test

import (
	"testing"

	"isigo/context"
	"isigo/lang"
	"isigo/symbol"
	"isigo/value_types"

	"github.com/stretchr/testify/assert"
)

func TestRead_Output_Int(t *testing.T) {
	ctx := context.New()
	sym := &symbol.Symbol{
		Identifier: "x",
		Type:       value_types.IntegerValueTypeEntity,
	}

	read := lang.NewRead(&ctx, sym)

	output, err := read.Output()

	assert.NoError(t, err)
	assert.Equal(t, "x = std.Leia__int()", output)
}

func TestRead_Output_Float(t *testing.T) {
	ctx := context.New()
	sym := &symbol.Symbol{
		Identifier: "y",
		Type:       value_types.FloatValueTypeEntity,
	}

	read := lang.NewRead(&ctx, sym)

	output, err := read.Output()

	assert.NoError(t, err)
	assert.Equal(t, "y = std.Leia__float()", output)
}

func TestRead_Output_String(t *testing.T) {
	ctx := context.New()
	sym := &symbol.Symbol{
		Identifier: "z",
		Type:       value_types.StringValueTypeEntity,
	}

	read := lang.NewRead(&ctx, sym)

	output, err := read.Output()

	assert.NoError(t, err)
	assert.Equal(t, "z = std.Leia__string()", output)
}
