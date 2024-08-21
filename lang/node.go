package lang

import "isigo/context"

type EvaluableNode interface {
	Eval(ctx *context.Context) (any, error)
}
type Node interface {
	EvaluableNode
	Output() (string, error)
}
