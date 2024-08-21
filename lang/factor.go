package lang

import (
	"fmt"
	"isigo/context"
	"isigo/symbol"
	"isigo/value_types"
)

type Factor interface {
	Node
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

func (p IntegerFactor) Output() (string, error) {
	return fmt.Sprintf("%d", p.value), nil
}

func (p IntegerFactor) Eval(ctx *context.Context) (any, error) {
	return p.value, nil
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

func (p FloatFactor) Output() (string, error) {
	return fmt.Sprintf("%f", p.value), nil
}

func (p FloatFactor) Eval(ctx *context.Context) (any, error) {
	return p.value, nil
}

// String ----------
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

func (p SymbolFactor) Output() (string, error) {
	return fmt.Sprintf("%s", p.symbol.Identifier), nil
}

func (p SymbolFactor) Eval(ctx *context.Context) (any, error) {
	return p.symbol.RuntimeValue(), nil
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

func (p ExpressionFactor) Output() (string, error) {
	content, err := p.expression.Output()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("(%s)", content), nil
}

func (p ExpressionFactor) Eval(ctx *context.Context) (any, error) {
	return p.expression.Eval(ctx)
}
