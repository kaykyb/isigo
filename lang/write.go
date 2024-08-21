package lang

import (
	"fmt"
	"isigo/context"
)

type Write struct {
	context *context.Context
	output  Expr
}

func NewWrite(ctx *context.Context, output Expr) Write {
	return Write{
		context: ctx,
		output:  output,
	}
}

func (p Write) Output() (string, error) {
	content, err := p.output.Output()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("fmt.Println(%s)", content), nil
}

func (p Write) Eval(ctx *context.Context) (any, error) {
	exprVal, err := p.output.Eval(ctx)
	if err != nil {
		return nil, err
	}

	fmt.Println(exprVal)

	return nil, nil
}
