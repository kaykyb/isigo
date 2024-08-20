package parser

import (
	"isigo/ast"
	"isigo/context"
	"isigo/tokens"
)

func (c *Parser) Assignment(ctx *context.Context, delta TokenDelta) (ast.Assignment, TokenDelta, error) {
	// -> ID
	if !delta.token.IsIdentifier() {
		return ast.Assignment{}, delta, unexpectedTokenTypeError(delta, tokens.Identifier)
	}

	identifier := delta.token.Content()

	if !ctx.SymbolExists(identifier) {
		return ast.Assignment{}, delta, usedBeforeDeclaration(identifier)
	}

	symbolTo, err := ctx.RetrieveSymbol(identifier)
	if err != nil {
		return ast.Assignment{}, delta, err
	}

	// -> :=
	delta, err = c.nextToken()
	if err != nil {
		return ast.Assignment{}, delta, err
	}

	if !delta.token.IsAssign() {
		return ast.Assignment{}, delta, unexpectedTokenTypeError(delta, tokens.Assign)
	}

	delta, err = c.nextToken()
	if err != nil {
		return ast.Assignment{}, delta, err
	}

	// -> Expr
	expr, delta, err := c.Expr(ctx, delta)
	if err != nil {
		return ast.Assignment{}, delta, err
	}

	assignment := ast.NewAssignment(ctx, symbolTo, expr)

	// -> .
	if !delta.token.IsStatementTerminator() {
		return ast.Assignment{}, delta, unexpectedTokenTypeError(delta, tokens.StatementTerminator)
	}

	// delta
	delta, err = c.nextToken()
	if err != nil {
		return ast.Assignment{}, delta, err
	}

	return assignment, delta, nil
}
