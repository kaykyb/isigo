package lang

import (
	"fmt"
	"isigo/context"
	"isigo/std"
	"isigo/symbol"
	"isigo/value_types"
)

type Read struct {
	context *context.Context
	output  *symbol.Symbol
}

func NewRead(ctx *context.Context, output *symbol.Symbol) Read {
	return Read{
		context: ctx,
		output:  output,
	}
}

func (p Read) Output() (string, error) {
	switch p.output.Type {
	case value_types.IntegerValueTypeEntity:
		return fmt.Sprintf("%s = std.Leia__int()", p.output.Identifier), nil
	case value_types.FloatValueTypeEntity:
		return fmt.Sprintf("%s = std.Leia__float()", p.output.Identifier), nil
	default:
		return fmt.Sprintf("%s = std.Leia__string()", p.output.Identifier), nil
	}
}

func (p Read) Eval(ctx *context.Context) (any, error) {
	var val any

	switch p.output.Type {
	case value_types.IntegerValueTypeEntity:
		val = std.Leia__int()
	case value_types.FloatValueTypeEntity:
		val = std.Leia__float()
	default:
		val = std.Leia__string()
	}

	p.output.AssignRuntimeValue(val)
	return val, nil
}
