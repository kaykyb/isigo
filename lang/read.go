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
		return fmt.Sprintf("scanfPanicInt(&%s)", p.output.Identifier), nil
	case value_types.FloatValueTypeEntity:
		return fmt.Sprintf("scanfPanicFloat(&%s)", p.output.Identifier), nil
	default:
		return fmt.Sprintf("scanfPanicString(&%s)", p.output.Identifier), nil
	}
}
