package parser

import (
	"isigo/context"
	"isigo/lang"
	"isigo/syntax"
	"isigo/tokens"
)

func (c *Parser) Prog(ctx *context.Context, delta TokenDelta) (lang.Program, TokenDelta, error) {
	// -> programa
	delta, err := c.startProg(delta)
	if err != nil {
		return lang.Program{}, delta, err
	}

	// Parse dos statements dentro do programa
	progBlock, delta, err := c.Block(ctx, delta)
	if err != nil {
		return lang.Program{}, delta, err
	}

	// -> fimprog
	delta, err = c.endProg(delta)
	if err != nil {
		return lang.Program{}, delta, err
	}

	// Sucesso
	program := lang.NewProgram(ctx, progBlock)

	return program, delta, nil
}

func (c *Parser) startProg(delta TokenDelta) (TokenDelta, error) {
	if !delta.token.IsReservedWord() {
		return TokenDelta{}, unexpectedTokenTypeError(delta, tokens.ReservedWord)
	}

	if !delta.token.Is(syntax.Program) {
		return TokenDelta{}, unexpectedContentError(delta, syntax.Program)
	}

	return c.nextToken()
}

func (c *Parser) endProg(delta TokenDelta) (TokenDelta, error) {
	if !delta.token.IsReservedWord() {
		return delta, unexpectedTokenTypeError(delta, tokens.ReservedWord)
	}

	if !delta.token.Is(syntax.EndProgram) {
		return delta, unexpectedContentError(delta, syntax.EndProgram)
	}

	return c.nextToken()
}
