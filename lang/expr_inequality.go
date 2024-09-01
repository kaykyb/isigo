package lang

import (
	"fmt"
	"isigo/context"
	"isigo/failure"
	"isigo/value_types"
)

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
