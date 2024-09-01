package lang

import (
	"isigo/value_types"
)

type Expr interface {
	Node
	ResultingType() (value_types.ValueType, error)
	IsExpr() bool
}
