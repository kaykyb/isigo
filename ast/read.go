package ast

import (
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
