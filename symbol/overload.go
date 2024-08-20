package symbol

import "isigo/value_types"

type Overload struct {
	params []value_types.ValueType
}

func NewOverload(params ...value_types.ValueType) *Overload {
	return &Overload{params: params}
}

func (o *Overload) ParamAt(i int) value_types.ValueType {
	return o.params[i]
}
