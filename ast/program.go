package ast

import (
	"isigo/context"
)

type Program struct {
	context *context.Context
	child   Block
}

func NewProgram(ctx *context.Context, child Block) Program {
	return Program{
		context: ctx,
		child:   child,
	}
}
