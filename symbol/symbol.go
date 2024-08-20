package symbol

import "isigo/value_types"

type Symbol struct {
	Identifier string
	Type       value_types.ValueType
	Overloads  []Overload
}

func New(identifier string, t value_types.ValueType) Symbol {
	return Symbol{
		Identifier: identifier,
		Type:       t,
	}
}
