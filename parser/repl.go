package parser

import (
	"isigo/context"
	"isigo/lang"
)

func (c *Parser) REPL(ctx *context.Context, delta TokenDelta) (lang.EvaluableNode, TokenDelta, error) {
	// Parse dos statements dentro do programa
	progBlock, delta, err := c.DeclareOrCommand(ctx, delta)
	if err != nil {
		return lang.Block{}, delta, err
	}

	return progBlock, delta, nil
}
