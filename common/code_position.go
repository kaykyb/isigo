package common

type CodePosition struct {
	BufferPosition int
	Line           int
	Column         int
}

func NewCodePosition(bufferPosition, line, column int) CodePosition {
	return CodePosition{
		BufferPosition: bufferPosition,
		Line:           line,
		Column:         column,
	}
}
