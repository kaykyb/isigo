package value_types

type BooleanValueType interface {
	ValueType
}

type booleanValueType struct{}

var BooleanValueTypeEntity = booleanValueType{}

func (booleanValueType) Name() string {
	return "booleano"
}

func (v booleanValueType) CanAssign(a ValueType) bool {
	return a == booleanValueType{}
}

func (v booleanValueType) Output() string {
	return "bool"
}
