package lang

import (
	"isigo/context"
	"isigo/value_types"
)

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

func (n FactorTerm) ResultingType() (value_types.ValueType, error) {
	return n.factor.ResultingType()
}
