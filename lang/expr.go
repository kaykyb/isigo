package lang

import (
	"fmt"
	"isigo/context"
	"isigo/failure"
	"isigo/value_types"
)

type Expr interface {
	Node
	ResultingType() (value_types.ValueType, error)
}

// ---- TERM ----

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

func (p TermExpr) Output() (string, error) {
	return p.term.Output()
}

func (n TermExpr) ResultingType() (value_types.ValueType, error) {
	return n.term.ResultingType()
}

// ---- SUM EXPR ----

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

func (p SumExpr) Output() (string, error) {
	leftContent, err := p.left.Output()
	if err != nil {
		return "", err
	}

	rightContent, err := p.term.Output()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s + %s", leftContent, rightContent), nil
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
		return value_types.FloatValueTypeEntity, failure.TypeNotSumable(leftType.Name())
	}

	return leftTypeSumable.ResultingSumType(termType)
}

// ---- SUBTRACT EXPR ----

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

func (p SubtractExpr) Output() (string, error) {
	leftContent, err := p.left.Output()
	if err != nil {
		return "", err
	}

	rightContent, err := p.term.Output()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s - %s", leftContent, rightContent), nil
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
		return value_types.FloatValueTypeEntity, failure.TypeNotSubtractable(leftType.Name())
	}

	return leftTypeSubtractable.ResultingSubtractType(termType)
}
