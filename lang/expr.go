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
	IsExpr() bool
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

func (p TermExpr) Eval(ctx *context.Context) (any, error) {
	return p.term.Eval(ctx)
}

func (n TermExpr) ResultingType() (value_types.ValueType, error) {
	return n.term.ResultingType()
}

func (p TermExpr) IsExpr() bool {
	return true
}

// ---- SUM EXPR ----

type SumExpr struct {
	context *context.Context
	left    Expr
	term    Term
}

func NewSumExpr(ctx *context.Context, left Expr, term Term) (SumExpr, error) {
	leftType, err := left.ResultingType()
	if err != nil {
		return SumExpr{}, err
	}

	leftTypeSumable, ok := leftType.(value_types.SumableValueType)
	if !ok {
		return SumExpr{}, failure.TypeNotSumable(leftType.Name())
	}

	termType, err := term.ResultingType()
	if err != nil {
		return SumExpr{}, err
	}

	_, err = leftTypeSumable.ResultingSumType(termType)

	return SumExpr{
		context: ctx,
		left:    left,
		term:    term,
	}, err
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

func (p SumExpr) Eval(ctx *context.Context) (any, error) {
	leftEval, err := p.left.Eval(ctx)
	if err != nil {
		return nil, err
	}

	termEval, err := p.term.Eval(ctx)
	if err != nil {
		return nil, err
	}

	switch left := leftEval.(type) {
	case int:
		switch factor := termEval.(type) {
		case int:
			return left + factor, nil
		case float64:
			return float64(left) + factor, nil
		default:
			return nil, fmt.Errorf("não é possível somar: %T", termEval)
		}
	case float64:
		switch factor := termEval.(type) {
		case int:
			return left + float64(factor), nil
		case float64:
			return left + factor, nil
		default:
			return nil, fmt.Errorf("não é possível somar: %T", termEval)
		}
	default:
		return nil, fmt.Errorf("não é possível somar: %T", leftEval)
	}
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

func (p SumExpr) IsExpr() bool {
	return true
}

// ---- SUBTRACT EXPR ----

type SubtractExpr struct {
	context *context.Context
	left    Expr
	term    Term
}

func NewSubtractExpr(ctx *context.Context, left Expr, term Term) (SubtractExpr, error) {
	leftType, err := left.ResultingType()
	if err != nil {
		return SubtractExpr{}, err
	}

	leftTypeSubtractable, ok := leftType.(value_types.SubtractableValueType)
	if !ok {
		return SubtractExpr{}, failure.TypeNotSubtractable(leftType.Name())
	}

	termType, err := term.ResultingType()
	if err != nil {
		return SubtractExpr{}, err
	}

	_, err = leftTypeSubtractable.ResultingSubtractType(termType)
	if err != nil {
		return SubtractExpr{}, err
	}

	return SubtractExpr{
		context: ctx,
		left:    left,
		term:    term,
	}, nil
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

func (p SubtractExpr) Eval(ctx *context.Context) (any, error) {
	leftEval, err := p.left.Eval(ctx)
	if err != nil {
		return nil, err
	}

	termEval, err := p.term.Eval(ctx)
	if err != nil {
		return nil, err
	}

	switch left := leftEval.(type) {
	case int:
		switch factor := termEval.(type) {
		case int:
			return left + factor, nil
		case float64:
			return float64(left) + factor, nil
		default:
			return nil, fmt.Errorf("não é possível subtrair: %T", termEval)
		}
	case float64:
		switch factor := termEval.(type) {
		case int:
			return left + float64(factor), nil
		case float64:
			return left + factor, nil
		default:
			return nil, fmt.Errorf("não é possível subtrair: %T", termEval)
		}
	default:
		return nil, fmt.Errorf("não é possível subtrair: %T", leftEval)
	}
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

func (p SubtractExpr) IsExpr() bool {
	return true
}

// ---- Equality EXPR ----
type EqualityExpr struct {
	context       *context.Context
	shouldBeEqual bool
	left          Expr
	right         Expr
}

func NewEqualityExpr(ctx *context.Context, shouldBeEqual bool, left Expr, right Expr) (EqualityExpr, error) {
	leftType, err := left.ResultingType()
	if err != nil {
		return EqualityExpr{}, err
	}

	rightType, err := right.ResultingType()
	if err != nil {
		return EqualityExpr{}, err
	}

	if leftType != rightType {
		return EqualityExpr{}, failure.CannotCompareTypes(leftType.Name(), rightType.Name())
	}

	return EqualityExpr{
		context:       ctx,
		shouldBeEqual: shouldBeEqual,
		left:          left,
		right:         right,
	}, nil
}

func (p EqualityExpr) Output() (string, error) {
	leftContent, err := p.left.Output()
	if err != nil {
		return "", err
	}

	rightContent, err := p.right.Output()
	if err != nil {
		return "", err
	}

	if p.shouldBeEqual {
		return fmt.Sprintf("%s == %s", leftContent, rightContent), nil
	} else {
		return fmt.Sprintf("%s != %s", leftContent, rightContent), nil
	}
}

func (p EqualityExpr) Eval(ctx *context.Context) (any, error) {
	leftEval, err := p.left.Eval(ctx)
	if err != nil {
		return nil, err
	}

	rightEval, err := p.right.Eval(ctx)
	if err != nil {
		return nil, err
	}

	if p.shouldBeEqual {
		return leftEval == rightEval, nil
	} else {
		return leftEval != rightEval, nil
	}
}

func (n EqualityExpr) ResultingType() (value_types.ValueType, error) {
	return value_types.BooleanValueTypeEntity, nil
}

func (p EqualityExpr) IsExpr() bool {
	return true
}

// ---- Inequality EXPR ----
type InequalityExpr struct {
	context  *context.Context
	operator string
	left     Expr
	right    Expr
}

func NewInequalityExpr(ctx *context.Context, operator string, left Expr, right Expr) (InequalityExpr, error) {
	leftType, err := left.ResultingType()
	if err != nil {
		return InequalityExpr{}, err
	}

	rightType, err := right.ResultingType()
	if err != nil {
		return InequalityExpr{}, err
	}

	_, ok := leftType.(value_types.OrdenableValueType)
	if !ok {
		return InequalityExpr{}, failure.CannotCompareTypes(leftType.Name(), rightType.Name())
	}

	_, ok = rightType.(value_types.OrdenableValueType)
	if !ok {
		return InequalityExpr{}, failure.CannotCompareTypes(rightType.Name(), rightType.Name())
	}

	return InequalityExpr{
		context:  ctx,
		operator: operator,
		left:     left,
		right:    right,
	}, nil
}

func (p InequalityExpr) Output() (string, error) {
	leftContent, err := p.left.Output()
	if err != nil {
		return "", err
	}

	rightContent, err := p.right.Output()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s %s %s", leftContent, p.operator, rightContent), nil

}

func (p InequalityExpr) Eval(ctx *context.Context) (any, error) {
	leftEval, err := p.left.Eval(ctx)
	if err != nil {
		return nil, err
	}

	rightEval, err := p.right.Eval(ctx)
	if err != nil {
		return nil, err
	}

	leftType, err := p.left.ResultingType()
	if err != nil {
		return InequalityExpr{}, err
	}

	rightType, err := p.right.ResultingType()
	if err != nil {
		return InequalityExpr{}, err
	}

	ordernableLeftType, ok := leftType.(value_types.OrdenableValueType)
	if !ok {
		return InequalityExpr{}, failure.CannotCompareTypes(leftType.Name(), rightType.Name())
	}

	ordernableRightType, ok := rightType.(value_types.OrdenableValueType)
	if !ok {
		return InequalityExpr{}, failure.CannotCompareTypes(rightType.Name(), rightType.Name())
	}

	leftOrdenableValue, err := ordernableLeftType.ToOrdenable(leftEval)
	rightOrdenableValue, err := ordernableRightType.ToOrdenable(rightEval)

	switch p.operator {
	case ">":
		return leftOrdenableValue > rightOrdenableValue, nil
	case ">=":
		return leftOrdenableValue >= rightOrdenableValue, nil
	case "<=":
		return leftOrdenableValue <= rightOrdenableValue, nil
	case "<":
		return leftOrdenableValue < rightOrdenableValue, nil
	default:
		return nil, fmt.Errorf("operador inválido: %s", p.operator)
	}
}

func (n InequalityExpr) ResultingType() (value_types.ValueType, error) {
	return value_types.BooleanValueTypeEntity, nil
}

func (p InequalityExpr) IsExpr() bool {
	return true
}
