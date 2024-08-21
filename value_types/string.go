package value_types

type StringValueType interface {
	ValueType
}

type stringValueType struct{}

var StringValueTypeEntity = stringValueType{}

func (stringValueType) Name() string {
	return "texto"
}

func (v stringValueType) CanAssign(a ValueType) bool {
	return a == stringValueType{}
}

func (v stringValueType) Output() string {
	return "string"
}
