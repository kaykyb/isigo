package lang

import (
	"fmt"
	"isigo/context"
)

type VariableContext struct {
	context *context.Context
	declare Declare
	child   Node
}

func NewVariableContext(ctx *context.Context, declare Declare, child Node) VariableContext {
	return VariableContext{
		context: ctx,
		declare: declare,
		child:   child,
	}
}

func (p VariableContext) Output() (string, error) {
	declareContent, err := p.declare.Output()
	if err != nil {
		return "", err
	}

	childContent, err := p.child.Output()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s\n%s", declareContent, childContent), nil
}
