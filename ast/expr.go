package ast

import (
	"isigo/compiler_error"
	"isigo/context"
	"isigo/value_types"
)

type Expr interface {
	ResultingType() (value_types.ValueType, error)
}

type TermExpr struct {
	context *context.Context
	term    Term
}

func NewTermExpr(ctx *context.Context, term Term) TermExpr {
	return TermExpr{
		context: ctx,
		term:    term,
	}
}

type SumExpr struct {
	context *context.Context
	left    Expr
	term    Term
}

func NewSumExpr(ctx *context.Context, left Expr, term Term) SumExpr {
	return SumExpr{
		context: ctx,
		left:    left,
		term:    term,
	}
}

type SubtractExpr struct {
	context *context.Context
	left    Expr
	term    Term
}

func NewSubtractExpr(ctx *context.Context, left Expr, term Term) SubtractExpr {
	return SubtractExpr{
		context: ctx,
		left:    left,
		term:    term,
	}
}

func (n TermExpr) ResultingType() (value_types.ValueType, error) {
	return n.term.ResultingType()
}

func (n SumExpr) ResultingType() (value_types.ValueType, error) {
	leftType, err := n.left.ResultingType()
	if err != nil {
		return value_types.FloatValueTypeEntity, err
	}

	termType, err := n.term.ResultingType()
	if err != nil {
		return value_types.FloatValueTypeEntity, err
	}

	leftTypeSumable, ok := leftType.(value_types.SumableValueType)
	if !ok {
		return value_types.FloatValueTypeEntity, compiler_error.TypeNotSumable(leftType.Name())
	}

	return leftTypeSumable.ResultingSumType(termType)
}

func (n SubtractExpr) ResultingType() (value_types.ValueType, error) {
	leftType, err := n.left.ResultingType()
	if err != nil {
		return value_types.FloatValueTypeEntity, err
	}

	termType, err := n.term.ResultingType()
	if err != nil {
		return value_types.FloatValueTypeEntity, err
	}

	leftTypeSubtractable, ok := leftType.(value_types.SubtractableValueType)
	if !ok {
		return value_types.FloatValueTypeEntity, compiler_error.TypeNotSubtractable(leftType.Name())
	}

	return leftTypeSubtractable.ResultingSubtractType(termType)
}
