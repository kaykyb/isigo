package ast

import (
	"isigo/context"
)

type ExecutionContext struct {
	context  *context.Context
	children []Node
}

func NewExecutionContext(ctx *context.Context, children []Node) ExecutionContext {
	return ExecutionContext{
		context:  ctx,
		children: children,
	}
}
