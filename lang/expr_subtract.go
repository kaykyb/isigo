package lang

import (
	"fmt"
	"isigo/context"
	"isigo/failure"
	"isigo/value_types"
)

type SubtractExpr struct {
	context *context.Context
	left    Expr
	term    Term
}

func NewSubtractExpr(ctx *context.Context, left Expr, term Term) (SubtractExpr, error) {
	leftType, err := left.ResultingType()
	if err != nil {
		return SubtractExpr{}, err
	}

	leftTypeSubtractable, ok := leftType.(value_types.SubtractableValueType)
	if !ok {
		return SubtractExpr{}, failure.TypeNotSubtractable(leftType.Name())
	}

	termType, err := term.ResultingType()
	if err != nil {
		return SubtractExpr{}, err
	}

	_, err = leftTypeSubtractable.ResultingSubtractType(termType)
	if err != nil {
		return SubtractExpr{}, err
	}

	return SubtractExpr{
		context: ctx,
		left:    left,
		term:    term,
	}, nil
}

func (p SubtractExpr) Output() (string, error) {
	leftContent, err := p.left.Output()
	if err != nil {
		return "", err
	}

	rightContent, err := p.term.Output()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s - %s", leftContent, rightContent), nil
}

func (p SubtractExpr) Eval(ctx *context.Context) (any, error) {
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
			return nil, fmt.Errorf("não é possível subtrair: %T", termEval)
		}
	case float64:
		switch factor := termEval.(type) {
		case int:
			return left + float64(factor), nil
		case float64:
			return left + factor, nil
		default:
			return nil, fmt.Errorf("não é possível subtrair: %T", termEval)
		}
	default:
		return nil, fmt.Errorf("não é possível subtrair: %T", leftEval)
	}
}

func (n SubtractExpr) ResultingType() (value_types.ValueType, error) {
	leftType, err := n.left.ResultingType()
	if err != nil {
		return value_types.FloatValueTypeEntity, err
	}

	termType, err := n.term.ResultingType()
	if err != nil {
		return value_types.FloatValueTypeEntity, err
	}

	leftTypeSubtractable, ok := leftType.(value_types.SubtractableValueType)
	if !ok {
		return value_types.FloatValueTypeEntity, failure.TypeNotSubtractable(leftType.Name())
	}

	return leftTypeSubtractable.ResultingSubtractType(termType)
}

func (p SubtractExpr) IsExpr() bool {
	return true
}
