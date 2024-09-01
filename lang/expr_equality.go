package lang

import (
	"fmt"
	"isigo/context"
	"isigo/failure"
	"isigo/value_types"
)

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
