package ast

import (
	"isigo/context"
)

type DeclareOrCommand struct {
	context *context.Context
	child   Node
}

func NewDeclareOrCommand(ctx *context.Context, child Node) DeclareOrCommand {
	return DeclareOrCommand{
		context: ctx,
		child:   child,
	}
}
