package lang

import (
	"isigo/context"
	"isigo/symbol"
	"isigo/value_types"
)

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
	return p.symbol.Identifier, nil
}

func (p SymbolFactor) Eval(ctx *context.Context) (any, error) {
	return p.symbol.RuntimeValue(), nil
}

func (p SymbolFactor) IsFactor() bool {
	return true
}
