package parser

import (
	"isigo/context"
	"isigo/lang"
	"isigo/syntax"
	"isigo/tokens"
)

func (c *Parser) RelationalExpr(ctx *context.Context, delta TokenDelta) (lang.Expr, TokenDelta, error) {
	leftExpr, delta, err := c.Expr(ctx, delta)
	if err != nil {
		return lang.EqualityExpr{}, delta, err
	}

	if !delta.token.IsOperator() {
		return lang.EqualityExpr{}, delta, unexpectedTokenTypeError(delta, tokens.Operator)
	}

	if delta.token.Is(syntax.IsEqual) || delta.token.Is(syntax.Different) {
		return c.equalityExpr(ctx, leftExpr, delta)
	}

	if delta.token.Is(syntax.Less) || delta.token.Is(syntax.Greater) ||
		delta.token.Is(syntax.LEQ) || delta.token.Is(syntax.GEQ) {
		return c.inequalityExpr(ctx, leftExpr, delta)
	}

	return leftExpr, delta, nil
}

func (c *Parser) equalityExpr(ctx *context.Context, left lang.Expr, delta TokenDelta) (lang.EqualityExpr, TokenDelta, error) {
	if !delta.token.Is(syntax.IsEqual) && !delta.token.Is(syntax.Different) {
		return lang.EqualityExpr{}, delta, unexpectedContentError(delta, "== ou !=")
	}

	shouldBeEqual := delta.token.Is(syntax.IsEqual)

	delta, err := c.nextToken()
	if err != nil {
		return lang.EqualityExpr{}, delta, err
	}

	rightExpr, delta, err := c.Expr(ctx, delta)
	if err != nil {
		return lang.EqualityExpr{}, delta, err
	}

	equalityExpr, err := lang.NewEqualityExpr(ctx, shouldBeEqual, left, rightExpr)
	return equalityExpr, delta, err
}

func (c *Parser) inequalityExpr(ctx *context.Context, left lang.Expr, delta TokenDelta) (lang.InequalityExpr, TokenDelta, error) {
	if !delta.token.Is(syntax.Greater) && !delta.token.Is(syntax.Less) && !delta.token.Is(syntax.LEQ) && !delta.token.Is(syntax.GEQ) {
		return lang.InequalityExpr{}, delta, unexpectedContentError(delta, "<, <=, > ou >=")
	}

	operator := delta.token.Content()

	delta, err := c.nextToken()
	if err != nil {
		return lang.InequalityExpr{}, delta, err
	}

	rightExpr, delta, err := c.Expr(ctx, delta)
	if err != nil {
		return lang.InequalityExpr{}, delta, err
	}

	inequalityExpr, err := lang.NewInequalityExpr(ctx, operator, left, rightExpr)
	return inequalityExpr, delta, err
}
