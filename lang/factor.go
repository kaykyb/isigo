package lang

import (
	"isigo/value_types"
)

type Factor interface {
	Node
	ResultingType() (value_types.ValueType, error)
	IsFactor() bool
}
