package symbol

import "isigo/value_types"

type Symbol struct {
	Identifier string
	Type       value_types.ValueType
	Overloads  []Overload
	Assigned   bool
}
