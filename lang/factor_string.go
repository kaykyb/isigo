package lang

import (
	"fmt"
	"isigo/context"
	"isigo/value_types"
)

type StringFactor struct {
	context *context.Context
	value   string
}

func (n StringFactor) ResultingType() (value_types.ValueType, error) {
	return value_types.StringValueTypeEntity, nil
}

func NewStringFactor(ctx *context.Context, value string) StringFactor {
	return StringFactor{
		context: ctx,
		value:   value,
	}
}

func (p StringFactor) Output() (string, error) {
	return fmt.Sprintf("\"%s\"", p.value), nil
}

func (p StringFactor) Eval(ctx *context.Context) (any, error) {
	return p.value, nil
}

func (p StringFactor) IsFactor() bool {
	return true
}
