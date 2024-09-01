package lang

import (
	"fmt"
	"isigo/context"
	"isigo/failure"
	"isigo/value_types"
)

type SumExpr struct {
	context *context.Context
	left    Expr
	term    Term
}

func NewSumExpr(ctx *context.Context, left Expr, term Term) (SumExpr, error) {
	leftType, err := left.ResultingType()
	if err != nil {
		return SumExpr{}, err
	}

	leftTypeSumable, ok := leftType.(value_types.SumableValueType)
	if !ok {
		return SumExpr{}, failure.TypeNotSumable(leftType.Name())
	}

	termType, err := term.ResultingType()
	if err != nil {
		return SumExpr{}, err
	}

	_, err = leftTypeSumable.ResultingSumType(termType)

	return SumExpr{
		context: ctx,
		left:    left,
		term:    term,
	}, err
}

func (p SumExpr) Output() (string, error) {
	leftContent, err := p.left.Output()
	if err != nil {
		return "", err
	}

	rightContent, err := p.term.Output()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s + %s", leftContent, rightContent), nil
}

func (p SumExpr) Eval(ctx *context.Context) (any, error) {
	leftEval, err := p.left.Eval(ctx)
	if err != nil {
		return nil, err
	}

	termEval, err := p.term.Eval(ctx)
	if err != nil {
		return nil, err
	}

	switch left := leftEval.(type) {
	case int:
		switch factor := termEval.(type) {
		case int:
			return left + factor, nil
		case float64:
			return float64(left) + factor, nil
		default:
			return nil, fmt.Errorf("não é possível somar: %T", termEval)
		}
	case float64:
		switch factor := termEval.(type) {
		case int:
			return left + float64(factor), nil
		case float64:
			return left + factor, nil
		default:
			return nil, fmt.Errorf("não é possível somar: %T", termEval)
		}
	default:
		return nil, fmt.Errorf("não é possível somar: %T", leftEval)
	}
}

func (n SumExpr) ResultingType() (value_types.ValueType, error) {
	leftType, err := n.left.ResultingType()
	if err != nil {
		return value_types.FloatValueTypeEntity, err
	}

	termType, err := n.term.ResultingType()
	if err != nil {
		return value_types.FloatValueTypeEntity, err
	}

	leftTypeSumable, ok := leftType.(value_types.SumableValueType)
	if !ok {
		return value_types.FloatValueTypeEntity, failure.TypeNotSumable(leftType.Name())
	}

	return leftTypeSumable.ResultingSumType(termType)
}

func (p SumExpr) IsExpr() bool {
	return true
}
