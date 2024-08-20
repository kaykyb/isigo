package ast

import (
	"isigo/context"
)

type Block struct {
	context *context.Context
	child   Node
}

func NewBlock(ctx *context.Context, child Node) Block {
	return Block{
		context: ctx,
		child:   child,
	}
}
