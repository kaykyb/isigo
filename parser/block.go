package parser

import (
	"isigo/context"
	"isigo/lang"
	"isigo/syntax"
	"isigo/tokens"
)

func (c *Parser) Block(ctx *context.Context, delta TokenDelta) (lang.Block, TokenDelta, error) {
	child, delta, err := c.VariableContext(ctx, delta)

	if err != nil {
		return lang.Block{}, delta, err
	}

	return lang.NewBlock(ctx, child), delta, err
}

func isBlockTerminator(token tokens.Token) bool {
	return (token.IsReservedWord() && token.Is(syntax.EndProgram)) || token.IsCloseBrace()
}
