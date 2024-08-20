package lang

import (
	"fmt"
	"isigo/compiler_error"
	"isigo/context"
	"isigo/value_types"
)

type Term interface {
	Node
	ResultingType() (value_types.ValueType, error)
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

	return fmt.Sprintf("%s", content), nil
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

	return fmt.Sprintf("%s / %s", leftContent, rightContent), nil
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
		return value_types.FloatValueTypeEntity, compiler_error.TypeNotDivisible(leftType.Name())
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
		return value_types.FloatValueTypeEntity, compiler_error.TypeNotMultipliable(leftType.Name())
	}

	return leftTypeMultipliable.ResultingMultiplicationType(factorType)
}
