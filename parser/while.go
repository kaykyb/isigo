package parser

import (
	"isigo/context"
	"isigo/lang"
	"isigo/syntax"
	"isigo/tokens"
)

func (c *Parser) WhileStatement(ctx *context.Context, delta TokenDelta) (lang.WhileStatement, TokenDelta, error) {
	// -> enquanto
	if !delta.token.IsReservedWord() {
		return lang.WhileStatement{}, delta, unexpectedTokenTypeError(delta, tokens.ReservedWord)
	}

	if !delta.token.Is(syntax.While) {
		return lang.WhileStatement{}, delta, unexpectedContentError(delta, syntax.While)
	}

	// -> (
	var err error
	delta, err = c.nextToken()
	if err != nil {
		return lang.WhileStatement{}, delta, err
	}

	if !delta.token.IsOpenParenthesis() {
		return lang.WhileStatement{}, delta, unexpectedTokenTypeError(delta, tokens.OpenParenthesis)
	}

	// -> Expr
	delta, err = c.nextToken()
	if err != nil {
		return lang.WhileStatement{}, delta, err
	}

	cond, delta, err := c.RelationalExpr(ctx, delta)
	if err != nil {
		return lang.WhileStatement{}, delta, err
	}

	// -> )
	if !delta.token.IsCloseParenthesis() {
		return lang.WhileStatement{}, delta, unexpectedTokenTypeError(delta, tokens.CloseParenthesis)
	}

	// -> {
	delta, err = c.nextToken()
	if err != nil {
		return lang.WhileStatement{}, delta, err
	}

	if !delta.token.IsOpenBrace() {
		return lang.WhileStatement{}, delta, unexpectedTokenTypeError(delta, tokens.OpenBrace)
	}

	// -> Block
	delta, err = c.nextToken()
	if err != nil {
		return lang.WhileStatement{}, delta, err
	}

	child, delta, err := c.Block(ctx, delta)
	if err != nil {
		return lang.WhileStatement{}, delta, err
	}

	// -> }
	if !delta.token.IsCloseBrace() {
		return lang.WhileStatement{}, delta, unexpectedTokenTypeError(delta, tokens.CloseBrace)
	}

	delta, err = c.nextToken()
	if err != nil {
		return lang.WhileStatement{}, delta, err
	}

	return lang.NewWhileStatement(ctx, cond, child), delta, nil
}
