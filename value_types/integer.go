package value_types

import "isigo/failure"

type IntegerValueType interface {
	ValueType
	SumableValueType
	SubtractableValueType
	MultipliableValueType
	DivisibleValueType
	OrdenableValueType
}

type integerValueType struct{}

var IntegerValueTypeEntity = integerValueType{}

func (integerValueType) Name() string {
	return "inteiro"
}

func (v integerValueType) ResultingSubtractType(by ValueType) (ValueType, error) {
	if by == IntegerValueTypeEntity {
		return IntegerValueTypeEntity, nil
	}

	if by == FloatValueTypeEntity {
		return FloatValueTypeEntity, nil
	}

	return FloatValueTypeEntity, failure.CannotDivideTypes(v.Name(), by.Name())
}

func (v integerValueType) ResultingSumType(by ValueType) (ValueType, error) {
	if by == IntegerValueTypeEntity {
		return IntegerValueTypeEntity, nil
	}

	if by == FloatValueTypeEntity {
		return FloatValueTypeEntity, nil
	}

	return FloatValueTypeEntity, failure.CannotDivideTypes(v.Name(), by.Name())
}

func (v integerValueType) ResultingMultiplicationType(by ValueType) (ValueType, error) {
	if by == IntegerValueTypeEntity || by == FloatValueTypeEntity {
		return FloatValueTypeEntity, nil
	}

	return FloatValueTypeEntity, failure.CannotDivideTypes(v.Name(), by.Name())
}

func (v integerValueType) ResultingDivisionType(by ValueType) (ValueType, error) {
	if by == IntegerValueTypeEntity || by == FloatValueTypeEntity {
		return FloatValueTypeEntity, nil
	}

	return FloatValueTypeEntity, failure.CannotDivideTypes(v.Name(), by.Name())
}

func (v integerValueType) CanAssign(a ValueType) bool {
	return a == integerValueType{}
}

func (v integerValueType) Output() string {
	return "int64"
}

func (v integerValueType) ToOrdenable(value any) (float64, error) {
	integerValue, ok := value.(int64)
	if !ok {
		panic("O valor atribuído não é um inteiro")
	}

	return float64(integerValue), nil
}
