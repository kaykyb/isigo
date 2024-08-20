package compiler_error

import (
	"fmt"
	"isigo/common"
)

type PositionedError struct {
	err      error
	position common.CodePosition
}

func (e *PositionedError) Error() string {
	return fmt.Sprintf("[Linha %d, Coluna %d]: %s", e.position.Line+1, e.position.Column+1, e.err.Error())
}
