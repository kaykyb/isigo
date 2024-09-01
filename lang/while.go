package lang

import (
	"fmt"
	"isigo/context"
)

type WhileStatement struct {
	context *context.Context
	cond    Expr
	child   Block
}

func NewWhileStatement(ctx *context.Context, cond Expr, child Block) WhileStatement {
	return WhileStatement{
		context: ctx,
		cond:    cond,
		child:   child,
	}
}

func (p WhileStatement) Output() (string, error) {
	condContent, err := p.cond.Output()
	if err != nil {
		return "", err
	}

	childContent, err := p.child.Output()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("for %s {\n%s\n}", condContent, childContent), nil
}

func (p WhileStatement) Eval(ctx *context.Context) (any, error) {
	var condVal any
	var err error

	for {
		condVal, err = p.cond.Eval(ctx)
		if err != nil {
			return nil, err
		}

		if !condVal.(bool) {
			return nil, nil
		}

		_, err := p.child.Eval(ctx)
		if err != nil {
			return nil, err
		}
	}
}
