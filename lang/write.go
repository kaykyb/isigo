package lang

import (
	"fmt"
	"isigo/context"
	"isigo/symbol"
)

type Write struct {
	context *context.Context
	output  *symbol.Symbol
}

func NewWrite(ctx *context.Context, output *symbol.Symbol) Write {
	return Write{
		context: ctx,
		output:  output,
	}
}

func (p Write) Output() (string, error) {
	return fmt.Sprintf("print(%s)", p.output.Identifier), nil
}
