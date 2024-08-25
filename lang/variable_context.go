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

func (p VariableContext) Eval(ctx *context.Context) (any, error) {
	newContext := context.NewWithParent(ctx)
	_, err := p.declare.Eval(&newContext)
	if err != nil {
		return nil, err
	}

	return p.child.Eval(&newContext)
}

func (p VariableContext) Declare() Node {
	return p.declare
}

func (p VariableContext) Child() Node {
	return p.child
}

func (p VariableContext) DeepestContext() *context.Context {
	executionContext, ok := p.child.(ExecutionContext)
	if ok && executionContext.DeepestContext() != nil {
		return executionContext.DeepestContext()
	}

	variableContext, ok := p.child.(VariableContext)
	if ok {
		return variableContext.DeepestContext()
	}

	return p.context
}
