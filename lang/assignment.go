package lang

import (
	"fmt"
	"isigo/context"
	"isigo/symbol"
)

type Assignment struct {
	context *context.Context
	to      *symbol.Symbol
	expr    Expr
}

func NewAssignment(context *context.Context, to *symbol.Symbol, expr Expr) Assignment {
	return Assignment{
		context: context,
		to:      to,
		expr:    expr,
	}
}

func (p Assignment) Output() (string, error) {
	exprContent, err := p.expr.Output()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s = %s", p.to.Identifier, exprContent), nil
}
