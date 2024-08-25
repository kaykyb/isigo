package lang

import (
	"fmt"
	"isigo/context"
	"isigo/failure"
	"isigo/value_types"
)

type Term interface {
	Node
	ResultingType() (value_types.ValueType, error)
	IsTerm() bool
}

type FactorTerm struct {
	context *context.Context
	factor  Factor
}

func NewFactorTerm(ctx *context.Context, factor Factor) FactorTerm {
	return FactorTerm{context: ctx, factor: factor}
}

func (p FactorTerm) Output() (string, error) {
	content, err := p.factor.Output()
	if err != nil {
		return "", err
	}

	return content, nil
}

func (p FactorTerm) Eval(ctx *context.Context) (any, error) {
	return p.factor.Eval(ctx)
}

func (p FactorTerm) IsTerm() bool {
	return true
}

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
	case int:
		switch factor := factorEval.(type) {
		case int:
			return float64(left) / float64(factor), nil
		case float64:
			return float64(left) / factor, nil
		default:
			return nil, fmt.Errorf("não é possível dividir: %T", factorEval)
		}
	case float64:
		switch factor := factorEval.(type) {
		case int:
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

func (n FactorTerm) ResultingType() (value_types.ValueType, error) {
	return n.factor.ResultingType()
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
