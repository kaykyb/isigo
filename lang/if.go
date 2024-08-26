package lang

import (
	"fmt"
	"isigo/context"
)

type IfStatement struct {
	context   *context.Context
	cond      Expr
	child     Block
	elseChild *Block
}

func NewIfStatement(ctx *context.Context, cond Expr, child Block, elseChild *Block) IfStatement {
	return IfStatement{
		context:   ctx,
		cond:      cond,
		child:     child,
		elseChild: elseChild,
	}
}

func (p IfStatement) Output() (string, error) {
	if p.elseChild == nil {
		return p.outputWithoutElse()
	} else {
		return p.outputWithElse()
	}
}

func (p IfStatement) outputWithoutElse() (string, error) {
	condContent, err := p.cond.Output()
	if err != nil {
		return "", err
	}

	childContent, err := p.child.Output()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("if %s {\n%s\n}", condContent, childContent), nil
}

func (p IfStatement) outputWithElse() (string, error) {
	condContent, err := p.cond.Output()
	if err != nil {
		return "", err
	}

	childContent, err := p.child.Output()
	if err != nil {
		return "", err
	}

	elseChild := *p.elseChild
	elseChildContent, err := elseChild.Output()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("if %s {\n%s\n} else {\n%s\n}", condContent, childContent, elseChildContent), nil
}

func (p IfStatement) Eval(ctx *context.Context) (any, error) {
	if p.elseChild == nil {
		return p.evalWithoutElse(ctx)
	} else {
		return p.evalWithElse(ctx)
	}
}

func (p IfStatement) evalWithoutElse(ctx *context.Context) (any, error) {
	condVal, err := p.cond.Eval(ctx)
	if err != nil {
		return nil, err
	}

	if condVal.(bool) {
		return p.child.Eval(ctx)
	}

	return nil, nil
}

func (p IfStatement) evalWithElse(ctx *context.Context) (any, error) {
	condVal, err := p.cond.Eval(ctx)
	if err != nil {
		return nil, err
	}

	if condVal.(bool) {
		return p.child.Eval(ctx)
	} else {
		return p.elseChild.Eval(ctx)
	}
}
