package ast

import (
	"isigo/context"
)

type Declare struct {
	context   *context.Context
	variables []Variable
}

func NewDeclare(ctx *context.Context, variables []Variable) Declare {
	return Declare{
		context:   ctx,
		variables: variables,
	}
}
