package parser

import (
	"isigo/context"
	"isigo/lang"
	"isigo/tokens"
)

func (c *Parser) Assignment(ctx *context.Context, delta TokenDelta) (lang.Assignment, TokenDelta, error) {
	// -> ID
	if !delta.token.IsIdentifier() {
		return lang.Assignment{}, delta, unexpectedTokenTypeError(delta, tokens.Identifier)
	}

	identifier := delta.token.Content()

	if !ctx.SymbolExists(identifier) {
		return lang.Assignment{}, delta, usedBeforeDeclaration(identifier)
	}

	symbolTo, err := ctx.RetrieveSymbol(identifier)
	if err != nil {
		return lang.Assignment{}, delta, err
	}

	// -> :=
	delta, err = c.nextToken()
	if err != nil {
		return lang.Assignment{}, delta, err
	}

	if !delta.token.IsAssign() {
		return lang.Assignment{}, delta, unexpectedTokenTypeError(delta, tokens.Assign)
	}

	delta, err = c.nextToken()
	if err != nil {
		return lang.Assignment{}, delta, err
	}

	// -> Expr
	expr, delta, err := c.Expr(ctx, delta)
	if err != nil {
		return lang.Assignment{}, delta, err
	}

	assignment := lang.NewAssignment(ctx, symbolTo, expr)

	// -> .
	if !delta.token.IsStatementTerminator() {
		return lang.Assignment{}, delta, unexpectedTokenTypeError(delta, tokens.StatementTerminator)
	}

	// delta
	delta, err = c.nextToken()
	if err != nil {
		return lang.Assignment{}, delta, err
	}

	return assignment, delta, nil
}
