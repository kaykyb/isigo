package lang

import (
	"fmt"
	"isigo/context"
	"isigo/value_types"
)

type ExpressionFactor struct {
	context    *context.Context
	expression Expr
}

func (n ExpressionFactor) ResultingType() (value_types.ValueType, error) {
	return n.expression.ResultingType()
}

func NewExpressionFactor(ctx *context.Context, expression Expr) ExpressionFactor {
	return ExpressionFactor{
		context:    ctx,
		expression: expression,
	}
}

func (p ExpressionFactor) Output() (string, error) {
	content, err := p.expression.Output()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("(%s)", content), nil
}

func (p ExpressionFactor) Eval(ctx *context.Context) (any, error) {
	return p.expression.Eval(ctx)
}

func (p ExpressionFactor) IsFactor() bool {
	return true
}
