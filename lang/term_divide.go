package lang

import (
	"fmt"
	"isigo/context"
	"isigo/failure"
	"isigo/value_types"
)

type DivideTerm struct {
	context *context.Context
	left    Term
	factor  Factor
}

func NewDivideTerm(ctx *context.Context, left Term, factor Factor) DivideTerm {
	return DivideTerm{
		context: ctx,
		left:    left,
		factor:  factor,
	}
}

func (p DivideTerm) Output() (string, error) {
	leftContent, err := p.left.Output()
	if err != nil {
		return "", err
	}

	rightContent, err := p.factor.Output()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("float64(%s) / float64(%s)", leftContent, rightContent), nil
}

func (p DivideTerm) Eval(ctx *context.Context) (any, error) {
	leftEval, err := p.left.Eval(ctx)
	if err != nil {
		return nil, err
	}

	factorEval, err := p.factor.Eval(ctx)
	if err != nil {
		return nil, err
	}

	switch left := leftEval.(type) {
	case int64:
		switch factor := factorEval.(type) {
		case int64:
			return float64(left) / float64(factor), nil
		case float64:
			return float64(left) / factor, nil
		default:
			return nil, fmt.Errorf("não é possível dividir: %T", factorEval)
		}
	case float64:
		switch factor := factorEval.(type) {
		case int64:
			return left / float64(factor), nil
		case float64:
			return left / factor, nil
		default:
			return nil, fmt.Errorf("não é possível dividir: %T", factorEval)
		}
	default:
		return nil, fmt.Errorf("não é possível dividir: %T", leftEval)
	}
}

func (p DivideTerm) IsTerm() bool {
	return true
}

func (n DivideTerm) ResultingType() (value_types.ValueType, error) {
	leftType, err := n.left.ResultingType()
	if err != nil {
		return value_types.FloatValueTypeEntity, err
	}

	factorType, err := n.factor.ResultingType()
	if err != nil {
		return value_types.FloatValueTypeEntity, err
	}

	leftTypeDivisible, ok := leftType.(value_types.DivisibleValueType)
	if !ok {
		return value_types.FloatValueTypeEntity, failure.TypeNotDivisible(leftType.Name())
	}

	return leftTypeDivisible.ResultingDivisionType(factorType)
}
