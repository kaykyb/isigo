package parser

import (
	"isigo/context"
	"isigo/lang"
	"isigo/syntax"
	"isigo/tokens"
)

func (c *Parser) IfStatement(ctx *context.Context, delta TokenDelta) (lang.IfStatement, TokenDelta, error) {
	// -> se
	if !delta.token.IsReservedWord() {
		return lang.IfStatement{}, delta, unexpectedTokenTypeError(delta, tokens.ReservedWord)
	}

	if !delta.token.Is(syntax.If) {
		return lang.IfStatement{}, delta, unexpectedContentError(delta, syntax.If)
	}

	// -> (
	var err error
	delta, err = c.nextToken()
	if err != nil {
		return lang.IfStatement{}, delta, err
	}

	if !delta.token.IsOpenParenthesis() {
		return lang.IfStatement{}, delta, unexpectedTokenTypeError(delta, tokens.OpenParenthesis)
	}

	// -> Expr
	delta, err = c.nextToken()
	if err != nil {
		return lang.IfStatement{}, delta, err
	}

	cond, delta, err := c.RelationalExpr(ctx, delta)
	if err != nil {
		return lang.IfStatement{}, delta, err
	}

	// -> )
	if !delta.token.IsCloseParenthesis() {
		return lang.IfStatement{}, delta, unexpectedTokenTypeError(delta, tokens.CloseParenthesis)
	}

	// -> {
	delta, err = c.nextToken()
	if err != nil {
		return lang.IfStatement{}, delta, err
	}

	if !delta.token.IsOpenBrace() {
		return lang.IfStatement{}, delta, unexpectedTokenTypeError(delta, tokens.OpenBrace)
	}

	// -> Block
	delta, err = c.nextToken()
	if err != nil {
		return lang.IfStatement{}, delta, err
	}

	child, delta, err := c.Block(ctx, delta)
	if err != nil {
		return lang.IfStatement{}, delta, err
	}

	// -> }
	if !delta.token.IsCloseBrace() {
		return lang.IfStatement{}, delta, unexpectedTokenTypeError(delta, tokens.CloseBrace)
	}

	delta, err = c.nextToken()
	if err != nil {
		return lang.IfStatement{}, delta, err
	}

	if delta.token.IsReservedWord() && delta.token.Is(syntax.Else) {
		return c.IfWithElseStatement(ctx, cond, child, delta)
	}

	return lang.NewIfStatement(ctx, cond, child, nil), delta, nil
}

func (c *Parser) IfWithElseStatement(ctx *context.Context, cond lang.Expr, ifChild lang.Block, delta TokenDelta) (lang.IfStatement, TokenDelta, error) {
	if !delta.token.IsReservedWord() {
		return lang.IfStatement{}, delta, unexpectedTokenTypeError(delta, tokens.ReservedWord)
	}

	if !delta.token.Is(syntax.Else) {
		return lang.IfStatement{}, delta, unexpectedContentError(delta, syntax.Else)
	}

	// -> {
	delta, err := c.nextToken()
	if err != nil {
		return lang.IfStatement{}, delta, err
	}

	if !delta.token.IsOpenBrace() {
		return lang.IfStatement{}, delta, unexpectedTokenTypeError(delta, tokens.OpenBrace)
	}

	// -> Block
	delta, err = c.nextToken()
	if err != nil {
		return lang.IfStatement{}, delta, err
	}

	elseChild, delta, err := c.Block(ctx, delta)
	if err != nil {
		return lang.IfStatement{}, delta, err
	}

	// -> }
	if !delta.token.IsCloseBrace() {
		return lang.IfStatement{}, delta, unexpectedTokenTypeError(delta, tokens.CloseBrace)
	}

	delta, err = c.nextToken()
	if err != nil {
		return lang.IfStatement{}, delta, err
	}

	return lang.NewIfStatement(ctx, cond, ifChild, &elseChild), delta, nil
}
