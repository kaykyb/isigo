package value_types

import "isigo/failure"

type FloatValueType interface {
	ValueType
	SumableValueType
	SubtractableValueType
	MultipliableValueType
	DivisibleValueType
}

type floatValueType struct{}

var FloatValueTypeEntity = floatValueType{}

func (floatValueType) Name() string {
	return "decimal"
}

func (v floatValueType) ResultingSubtractType(by ValueType) (ValueType, error) {
	if by == IntegerValueTypeEntity || by == FloatValueTypeEntity {
		return FloatValueTypeEntity, nil
	}

	return FloatValueTypeEntity, failure.CannotDivideTypes(v.Name(), by.Name())
}

func (v floatValueType) ResultingSumType(by ValueType) (ValueType, error) {
	if by == IntegerValueTypeEntity || by == FloatValueTypeEntity {
		return FloatValueTypeEntity, nil
	}

	return FloatValueTypeEntity, failure.CannotDivideTypes(v.Name(), by.Name())
}

func (v floatValueType) ResultingMultiplicationType(by ValueType) (ValueType, error) {
	if by == IntegerValueTypeEntity || by == FloatValueTypeEntity {
		return FloatValueTypeEntity, nil
	}

	return FloatValueTypeEntity, failure.CannotDivideTypes(v.Name(), by.Name())
}

func (v floatValueType) ResultingDivisionType(by ValueType) (ValueType, error) {
	if by == IntegerValueTypeEntity || by == FloatValueTypeEntity {
		return FloatValueTypeEntity, nil
	}

	return FloatValueTypeEntity, failure.CannotDivideTypes(v.Name(), by.Name())
}

func (v floatValueType) CanAssign(a ValueType) bool {
	return a == floatValueType{} || a == integerValueType{}
}

func (v floatValueType) Output() string {
	return "float64"
}
