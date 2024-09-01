package lang

import (
	"isigo/value_types"
)

type Term interface {
	Node
	ResultingType() (value_types.ValueType, error)
	IsTerm() bool
}
