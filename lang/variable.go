package lang

import (
	"isigo/context"
	"isigo/value_types"
)

type Variable struct {
	context      *context.Context
	identifier   string
	variableType value_types.ValueType
}

func NewVariable(ctx *context.Context, identifier string, variableType value_types.ValueType) Variable {
	return Variable{
		context:      ctx,
		identifier:   identifier,
		variableType: variableType,
	}
}
