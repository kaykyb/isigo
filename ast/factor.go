package ast

import (
	"isigo/context"
	"isigo/symbol"
	"isigo/value_types"
)

type Factor interface {
	ResultingType() (value_types.ValueType, error)
}

// Integer --------
type IntegerFactor struct {
	context *context.Context
	value   int
}

func (n IntegerFactor) ResultingType() (value_types.ValueType, error) {
	return value_types.IntegerValueTypeEntity, nil
}

func NewIntegerFactor(ctx *context.Context, value int) IntegerFactor {
	return IntegerFactor{
		context: ctx,
		value:   value,
	}
}

// Float ----------
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

// Symbol ----------
type SymbolFactor struct {
	context *context.Context
	symbol  *symbol.Symbol
}

func (n SymbolFactor) ResultingType() (value_types.ValueType, error) {
	return n.symbol.Type, nil
}

func NewSymbolFactor(ctx *context.Context, symbol *symbol.Symbol) SymbolFactor {
	return SymbolFactor{
		context: ctx,
		symbol:  symbol,
	}
}

// Expr ------------
type ExpressionFactor struct {
	context    *context.Context
	expression Expr
}

func (n ExpressionFactor) ResultingType() (value_types.ValueType, error) {
	return n.expression.ResultingType()
}

func NewExpressionFactor(ctx *context.Context, expression Expr) ExpressionFactor {
	return ExpressionFactor{
		context:    ctx,
		expression: expression,
	}
}
