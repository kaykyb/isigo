package parser

import (
	"isigo/context"
	"isigo/lang"
)

func (c *Parser) ExecutionContext(ctx *context.Context, delta TokenDelta) (lang.ExecutionContext, TokenDelta, error) {
	var children []lang.Node
	var child lang.Node
	var err error

	for {
		child, delta, err = c.DeclareOrCommand(ctx, delta)
		if err != nil {
			return lang.ExecutionContext{}, delta, err
		}

		children = append(children, child)

		if isBlockTerminator(delta.token) {
			return lang.NewExecutionContext(ctx, children), delta, nil
		}
	}
}
