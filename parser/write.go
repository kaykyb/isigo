package parser

import (
	"isigo/context"
	"isigo/lang"
	"isigo/syntax"
	"isigo/tokens"
)

func (c *Parser) Write(ctx *context.Context, delta TokenDelta) (lang.Write, TokenDelta, error) {
	// -> escreva
	if !delta.token.IsReservedWord() {
		return lang.Write{}, delta, unexpectedTokenTypeError(delta, tokens.ReservedWord)
	}

	if !delta.token.Is(syntax.Write) {
		return lang.Write{}, delta, unexpectedContentError(delta, syntax.Write)
	}

	// -> (
	var err error
	delta, err = c.nextToken()
	if err != nil {
		return lang.Write{}, delta, err
	}

	if !delta.token.IsOpenParenthesis() {
		return lang.Write{}, delta, unexpectedTokenTypeError(delta, tokens.OpenParenthesis)
	}

	// -> Expr
	delta, err = c.nextToken()
	if err != nil {
		return lang.Write{}, delta, err
	}

	expr, delta, err := c.Expr(ctx, delta)
	if err != nil {
		return lang.Write{}, delta, err
	}

	// -> )
	if !delta.token.IsCloseParenthesis() {
		return lang.Write{}, delta, unexpectedTokenTypeError(delta, tokens.CloseParenthesis)
	}

	// -> .
	delta, err = c.nextToken()
	if err != nil {
		return lang.Write{}, delta, err
	}

	if !delta.token.IsStatementTerminator() {
		return lang.Write{}, delta, unexpectedTokenTypeError(delta, tokens.StatementTerminator)
	}

	// delta
	delta, err = c.nextToken()
	if err != nil {
		return lang.Write{}, delta, err
	}

	return lang.NewWrite(ctx, expr), delta, nil
}
