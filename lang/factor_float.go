package lang

import (
	"fmt"
	"isigo/context"
	"isigo/value_types"
)

type FloatFactor struct {
	context *context.Context
	value   float64
}

func (n FloatFactor) ResultingType() (value_types.ValueType, error) {
	return value_types.IntegerValueTypeEntity, nil
}

func NewFloatFactor(ctx *context.Context, value float64) FloatFactor {
	return FloatFactor{
		context: ctx,
		value:   value,
	}
}

func (p FloatFactor) Output() (string, error) {
	return fmt.Sprintf("%f", p.value), nil
}

func (p FloatFactor) Eval(ctx *context.Context) (any, error) {
	return p.value, nil
}

func (p FloatFactor) IsFactor() bool {
	return true
}
