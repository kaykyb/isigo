package parser

import (
	"isigo/ast"
	"isigo/context"
	"isigo/syntax"
)

func (c *Parser) VariableContext(ctx *context.Context, delta TokenDelta) (ast.Node, TokenDelta, error) {
	// -> cmd
	if !delta.token.IsReservedWord() || !delta.token.Is(syntax.Declare) {
		return c.ExecutionContext(ctx, delta)
	}

	// -> declare
	newContext := context.NewWithParent(ctx)

	declare, delta, err := c.Declare(&newContext, delta)
	if err != nil {
		return ast.VariableContext{}, delta, err
	}

	// -> fim bloco
	if isBlockTerminator(delta.token) {
		return declare, delta, nil
	}

	// -> declare ou expr
	insideVariableContext, delta, err := c.VariableContext(&newContext, delta)
	if err != nil {
		return ast.VariableContext{}, delta, err
	}

	return ast.NewVariableContext(ctx, declare, insideVariableContext), delta, err
}
