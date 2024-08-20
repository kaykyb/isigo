package value_types

import (
	"isigo/compiler_error"
	"reflect"
)

type ValueType interface {
	Name() string
}

func implements(structInstance interface{}, interfaceType interface{}) bool {
	structType := reflect.TypeOf(structInstance)
	interfaceTypeType := reflect.TypeOf(interfaceType).Elem()

	return structType.Implements(interfaceTypeType)
}

type SumableValueType interface {
	ValueType
	ResultingSumType(by ValueType) (ValueType, error)
}

func IsSumable(v ValueType) bool {
	return implements(v, (*SumableValueType)(nil))
}

type SubtractableValueType interface {
	ValueType
	ResultingSubtractType(by ValueType) (ValueType, error)
}

func IsSubtractable(v ValueType) bool {
	return implements(v, (*SubtractableValueType)(nil))
}

type MultipliableValueType interface {
	ValueType
	ResultingMultiplicationType(by ValueType) (ValueType, error)
}

func IsMultipliable(v ValueType) bool {
	return implements(v, (*MultipliableValueType)(nil))
}

type DivisibleValueType interface {
	ValueType
	ResultingDivisionType(by ValueType) (ValueType, error)
}

func IsDivisible(v ValueType) bool {
	return implements(v, (*DivisibleValueType)(nil))
}

type IntegerValueType interface {
	ValueType
	SumableValueType
	SubtractableValueType
	MultipliableValueType
	DivisibleValueType
}

type FloatValueType interface {
	ValueType
	SumableValueType
	SubtractableValueType
	MultipliableValueType
	DivisibleValueType
}

type integerValueType struct{}
type floatValueType struct{}

var IntegerValueTypeEntity = integerValueType{}
var FloatValueTypeEntity = floatValueType{}

// ------------- Integers
func (integerValueType) Name() string {
	return "int"
}

func (v integerValueType) ResultingSubtractType(by ValueType) (ValueType, error) {
	if by == IntegerValueTypeEntity {
		return IntegerValueTypeEntity, nil
	}

	if by == FloatValueTypeEntity {
		return FloatValueTypeEntity, nil
	}

	return FloatValueTypeEntity, compiler_error.CannotDivideTypes(v.Name(), by.Name())
}

func (v integerValueType) ResultingSumType(by ValueType) (ValueType, error) {
	if by == IntegerValueTypeEntity {
		return IntegerValueTypeEntity, nil
	}

	if by == FloatValueTypeEntity {
		return FloatValueTypeEntity, nil
	}

	return FloatValueTypeEntity, compiler_error.CannotDivideTypes(v.Name(), by.Name())
}

func (v integerValueType) ResultingMultiplicationType(by ValueType) (ValueType, error) {
	if by == IntegerValueTypeEntity || by == FloatValueTypeEntity {
		return FloatValueTypeEntity, nil
	}

	return FloatValueTypeEntity, compiler_error.CannotDivideTypes(v.Name(), by.Name())
}

func (v integerValueType) ResultingDivisionType(by ValueType) (ValueType, error) {
	if by == IntegerValueTypeEntity || by == FloatValueTypeEntity {
		return FloatValueTypeEntity, nil
	}

	return FloatValueTypeEntity, compiler_error.CannotDivideTypes(v.Name(), by.Name())
}

// ------------- Floats

func (floatValueType) Name() string {
	return "float"
}

func (v floatValueType) ResultingSubtractType(by ValueType) (ValueType, error) {
	if by == IntegerValueTypeEntity || by == FloatValueTypeEntity {
		return FloatValueTypeEntity, nil
	}

	return FloatValueTypeEntity, compiler_error.CannotDivideTypes(v.Name(), by.Name())
}

func (v floatValueType) ResultingSumType(by ValueType) (ValueType, error) {
	if by == IntegerValueTypeEntity || by == FloatValueTypeEntity {
		return FloatValueTypeEntity, nil
	}

	return FloatValueTypeEntity, compiler_error.CannotDivideTypes(v.Name(), by.Name())
}

func (v floatValueType) ResultingMultiplicationType(by ValueType) (ValueType, error) {
	if by == IntegerValueTypeEntity || by == FloatValueTypeEntity {
		return FloatValueTypeEntity, nil
	}

	return FloatValueTypeEntity, compiler_error.CannotDivideTypes(v.Name(), by.Name())
}

func (v floatValueType) ResultingDivisionType(by ValueType) (ValueType, error) {
	if by == IntegerValueTypeEntity || by == FloatValueTypeEntity {
		return FloatValueTypeEntity, nil
	}

	return FloatValueTypeEntity, compiler_error.CannotDivideTypes(v.Name(), by.Name())
}

// ----------- Strings

type StringType struct{}

func (StringType) Name() string {
	return "string"
}
