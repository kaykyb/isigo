package lang

import (
	"isigo/context"
)

type ReplEnd struct {
	context *context.Context
}

func NewReplEnd(ctx *context.Context) ReplEnd {
	return ReplEnd{
		context: ctx,
	}
}

func (p ReplEnd) Output() (string, error) {
	return "", nil
}

func (p ReplEnd) Eval(ctx *context.Context) (any, error) {
	ctx.ReplacementContext = ctx
	return nil, nil
}
