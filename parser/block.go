package parser

import (
	"isigo/ast"
	"isigo/context"
	"isigo/syntax"
	"isigo/tokens"
)

func (c *Parser) Block(ctx *context.Context, delta TokenDelta) (ast.Block, TokenDelta, error) {
	newContext := context.NewWithParent(ctx)
	child, delta, err := c.VariableContext(&newContext, delta)

	if err != nil {
		return ast.Block{}, delta, err
	}

	return ast.NewBlock(ctx, child), delta, err
}

func isBlockTerminator(token tokens.Token) bool {
	return (token.IsReservedWord() && token.Is(syntax.EndProgram)) || token.IsCloseBrace()
}
