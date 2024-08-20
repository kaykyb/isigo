package symbol

type Value int

const (
	IntegerType Value = iota
	FloatType
	StringType
	InternalFunctionType
)
