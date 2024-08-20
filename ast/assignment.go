package ast

import (
	"isigo/context"
	"isigo/symbol"
)

type Assignment struct {
	context *context.Context
	to      *symbol.Symbol
	expr    Expr
}

func NewAssignment(context *context.Context, to *symbol.Symbol, expr Expr) Assignment {
	return Assignment{context: context, to: to, expr: expr}
}
