package lang

import (
	"fmt"
	"isigo/context"
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
		return fmt.Sprintf("%s = scanfPanicInt()", p.output.Identifier), nil
	case value_types.FloatValueTypeEntity:
		return fmt.Sprintf("%s = scanfPanicFloat()", p.output.Identifier), nil
	default:
		return fmt.Sprintf("%s = scanfPanicString()", p.output.Identifier), nil
	}
}
