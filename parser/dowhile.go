package parser

import (
	"isigo/context"
	"isigo/lang"
	"isigo/syntax"
	"isigo/tokens"
)

func (c *Parser) DoWhileStatement(ctx *context.Context, delta TokenDelta) (lang.DoWhileStatement, TokenDelta, error) {
	var err error

	// -> faca
	if !delta.token.IsReservedWord() {
		return lang.DoWhileStatement{}, delta, unexpectedTokenTypeError(delta, tokens.ReservedWord)
	}

	if !delta.token.Is(syntax.Do) {
		return lang.DoWhileStatement{}, delta, unexpectedContentError(delta, syntax.Do)
	}

	// -> {
	delta, err = c.nextToken()
	if err != nil {
		return lang.DoWhileStatement{}, delta, err
	}

	if !delta.token.IsOpenBrace() {
		return lang.DoWhileStatement{}, delta, unexpectedTokenTypeError(delta, tokens.OpenBrace)
	}

	// -> Block
	delta, err = c.nextToken()
	if err != nil {
		return lang.DoWhileStatement{}, delta, err
	}

	child, delta, err := c.Block(ctx, delta)
	if err != nil {
		return lang.DoWhileStatement{}, delta, err
	}

	// -> }
	if !delta.token.IsCloseBrace() {
		return lang.DoWhileStatement{}, delta, unexpectedTokenTypeError(delta, tokens.CloseBrace)
	}

	delta, err = c.nextToken()
	if err != nil {
		return lang.DoWhileStatement{}, delta, err
	}

	// -> enquanto
	if !delta.token.IsReservedWord() {
		return lang.DoWhileStatement{}, delta, unexpectedTokenTypeError(delta, tokens.ReservedWord)
	}

	if !delta.token.Is(syntax.While) {
		return lang.DoWhileStatement{}, delta, unexpectedContentError(delta, syntax.Do)
	}

	// -> (
	delta, err = c.nextToken()
	if err != nil {
		return lang.DoWhileStatement{}, delta, err
	}

	if !delta.token.IsOpenParenthesis() {
		return lang.DoWhileStatement{}, delta, unexpectedTokenTypeError(delta, tokens.OpenParenthesis)
	}

	// -> Expr
	delta, err = c.nextToken()
	if err != nil {
		return lang.DoWhileStatement{}, delta, err
	}

	cond, delta, err := c.RelationalExpr(ctx, delta)
	if err != nil {
		return lang.DoWhileStatement{}, delta, err
	}

	// -> )
	if !delta.token.IsCloseParenthesis() {
		return lang.DoWhileStatement{}, delta, unexpectedTokenTypeError(delta, tokens.CloseParenthesis)
	}

	delta, err = c.nextToken()
	if err != nil {
		return lang.DoWhileStatement{}, delta, err
	}

	return lang.NewDoWhileStatement(ctx, cond, child), delta, nil
}
