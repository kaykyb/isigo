package lang

import (
	"fmt"
	"isigo/context"
	"isigo/value_types"
)

type IntegerFactor struct {
	context *context.Context
	value   int64
}

func (n IntegerFactor) ResultingType() (value_types.ValueType, error) {
	return value_types.IntegerValueTypeEntity, nil
}

func NewIntegerFactor(ctx *context.Context, value int64) IntegerFactor {
	return IntegerFactor{
		context: ctx,
		value:   value,
	}
}

func (p IntegerFactor) Output() (string, error) {
	return fmt.Sprintf("%d", p.value), nil
}

func (p IntegerFactor) Eval(ctx *context.Context) (any, error) {
	return p.value, nil
}

func (p IntegerFactor) IsFactor() bool {
	return true
}
