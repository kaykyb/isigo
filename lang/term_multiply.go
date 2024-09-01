package lang

import (
	"fmt"
	"isigo/context"
	"isigo/failure"
	"isigo/value_types"
)

type MultiplyTerm struct {
	context *context.Context
	left    Term
	factor  Factor
}

func NewMultiplyTerm(ctx *context.Context, left Term, factor Factor) MultiplyTerm {
	return MultiplyTerm{
		context: ctx,
		left:    left,
		factor:  factor,
	}
}

func (p MultiplyTerm) Output() (string, error) {
	leftContent, err := p.left.Output()
	if err != nil {
		return "", err
	}

	rightContent, err := p.factor.Output()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s * %s", leftContent, rightContent), nil
}

func (p MultiplyTerm) Eval(ctx *context.Context) (any, error) {
	leftEval, err := p.left.Eval(ctx)
	if err != nil {
		return nil, err
	}

	factorEval, err := p.factor.Eval(ctx)
	if err != nil {
		return nil, err
	}

	switch left := leftEval.(type) {
	case int:
		switch factor := factorEval.(type) {
		case int:
			return left * factor, nil
		case float64:
			return float64(left) * factor, nil
		default:
			return nil, fmt.Errorf("não é possível multiplicar: %T", factorEval)
		}
	case float64:
		switch factor := factorEval.(type) {
		case int:
			return left * float64(factor), nil
		case float64:
			return left * factor, nil
		default:
			return nil, fmt.Errorf("não é possível multiplicar: %T", factorEval)
		}
	default:
		return nil, fmt.Errorf("não é possível multiplicar: %T", leftEval)
	}
}

func (p MultiplyTerm) IsTerm() bool {
	return true
}

func (n MultiplyTerm) ResultingType() (value_types.ValueType, error) {
	leftType, err := n.left.ResultingType()
	if err != nil {
		return value_types.FloatValueTypeEntity, err
	}

	factorType, err := n.factor.ResultingType()
	if err != nil {
		return value_types.FloatValueTypeEntity, err
	}

	leftTypeMultipliable, ok := leftType.(value_types.MultipliableValueType)
	if !ok {
		return value_types.FloatValueTypeEntity, failure.TypeNotMultipliable(leftType.Name())
	}

	return leftTypeMultipliable.ResultingMultiplicationType(factorType)
}
