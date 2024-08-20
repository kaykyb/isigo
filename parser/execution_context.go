package parser

import (
	"isigo/ast"
	"isigo/context"
)

func (c *Parser) ExecutionContext(ctx *context.Context, delta TokenDelta) (ast.ExecutionContext, TokenDelta, error) {
	var children []ast.Node
	var child ast.Node
	var err error

	for {
		child, delta, err = c.DeclareOrCommand(ctx, delta)
		if err != nil {
			return ast.ExecutionContext{}, delta, err
		}

		children = append(children, child)

		if isBlockTerminator(delta.token) {
			return ast.NewExecutionContext(ctx, children), delta, nil
		}
	}
}
