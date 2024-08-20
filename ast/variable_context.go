package ast

import (
	"isigo/context"
)

type VariableContext struct {
	context *context.Context
	declare Declare
	child   Node
}

func NewVariableContext(ctx *context.Context, declare Declare, child Node) VariableContext {
	return VariableContext{
		context: ctx,
		declare: declare,
		child:   child,
	}
}
