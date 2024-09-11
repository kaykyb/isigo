package lang

import (
	"fmt"
	"isigo/context"
	"isigo/failure"
	"isigo/symbol"
)

type Assignment struct {
	context *context.Context
	to      *symbol.Symbol
	expr    Expr
}

func NewAssignment(context *context.Context, to *symbol.Symbol, expr Expr) (Assignment, error) {
	exprResultingType, err := expr.ResultingType()
	if err != nil {
		return Assignment{}, err
	}

	if !to.Type.CanAssign(exprResultingType) {
		return Assignment{}, failure.SymbolTypeDiffers(to.Identifier, to.Type.Name(), exprResultingType.Name())
	}

	return Assignment{
		context: context,
		to:      to,
		expr:    expr,
	}, nil
}

func (p Assignment) Output() (string, error) {
	exprContent, err := p.expr.Output()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s = %s(%s)", p.to.Identifier, p.to.Type.Output(), exprContent), nil
}

func (p Assignment) Eval(ctx *context.Context) (any, error) {
	exprVal, err := p.expr.Eval(ctx)
	if err != nil {
		return "", err
	}

	p.to.AssignRuntimeValue(exprVal)
	return nil, nil
}
