package lang

import (
	"isigo/context"
	"isigo/value_types"
)

type TermExpr struct {
	context *context.Context
	term    Term
}

func NewTermExpr(ctx *context.Context, term Term) TermExpr {
	return TermExpr{
		context: ctx,
		term:    term,
	}
}

func (p TermExpr) Output() (string, error) {
	return p.term.Output()
}

func (p TermExpr) Eval(ctx *context.Context) (any, error) {
	return p.term.Eval(ctx)
}

func (n TermExpr) ResultingType() (value_types.ValueType, error) {
	return n.term.ResultingType()
}

func (p TermExpr) IsExpr() bool {
	return true
}
