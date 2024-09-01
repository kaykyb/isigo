package lang

import (
	"fmt"
	"isigo/context"
)

type DoWhileStatement struct {
	context *context.Context
	cond    Expr
	child   Block
}

func NewDoWhileStatement(ctx *context.Context, cond Expr, child Block) DoWhileStatement {
	return DoWhileStatement{
		context: ctx,
		cond:    cond,
		child:   child,
	}
}

func (p DoWhileStatement) Output() (string, error) {
	condContent, err := p.cond.Output()
	if err != nil {
		return "", err
	}

	childContent, err := p.child.Output()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("for ISI__ok := true; ISI__ok; ISI__ok = %s {\n%s\n}", condContent, childContent), nil
}

func (p DoWhileStatement) Eval(ctx *context.Context) (any, error) {
	var condVal any
	var err error

	for {
		_, err = p.child.Eval(ctx)
		if err != nil {
			return nil, err
		}

		condVal, err = p.cond.Eval(ctx)
		if err != nil {
			return nil, err
		}

		if !condVal.(bool) {
			return nil, nil
		}
	}
}
