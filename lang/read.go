package lang

import (
	"fmt"
	"isigo/context"
	"isigo/symbol"
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
	return fmt.Sprintf("scanfPanicInt(&%s)", p.output.Identifier), nil
}
