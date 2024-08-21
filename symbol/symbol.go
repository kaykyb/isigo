package symbol

import "isigo/value_types"

type Symbol struct {
	Identifier   string
	Type         value_types.ValueType
	Overloads    []Overload
	Assigned     bool
	runtimeValue any
}

func (s *Symbol) AssignRuntimeValue(val any) {
	s.Assigned = true
	s.runtimeValue = val
}

func (s *Symbol) RuntimeValue() any {
	return s.runtimeValue
}
