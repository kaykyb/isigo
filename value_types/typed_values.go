package value_types

import (
	"isigo/syntax"
)

type ValueType interface {
	Name() string
	CanAssign(ValueType) bool
	Output() string
}

type SumableValueType interface {
	ValueType
	ResultingSumType(by ValueType) (ValueType, error)
}

type SubtractableValueType interface {
	ValueType
	ResultingSubtractType(by ValueType) (ValueType, error)
}

type MultipliableValueType interface {
	ValueType
	ResultingMultiplicationType(by ValueType) (ValueType, error)
}

type DivisibleValueType interface {
	ValueType
	ResultingDivisionType(by ValueType) (ValueType, error)
}

func TypeEntityForTypeT(typeT string) ValueType {
	switch typeT {
	case syntax.IntegerT:
		return IntegerValueTypeEntity
	case syntax.FloatT:
		return FloatValueTypeEntity
	case syntax.StringT:
		return StringValueTypeEntity
	default:
		return IntegerValueTypeEntity
	}
}
