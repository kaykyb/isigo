package parser

import (
	"isigo/ast"
	"isigo/context"
	"isigo/syntax"
	"isigo/tokens"
)

func (c *Parser) Declare(ctx *context.Context, delta TokenDelta) (ast.Declare, TokenDelta, error) {
	// -> declare
	if !delta.token.IsReservedWord() {
		return ast.Declare{}, delta, unexpectedTokenTypeError(delta, tokens.ReservedWord)
	}

	if !delta.token.Is(syntax.Declare) {
		return ast.Declare{}, delta, unexpectedContentError(delta, syntax.Declare)
	}

	var err error
	delta, err = c.nextToken()
	if err != nil {
		return ast.Declare{}, delta, err
	}

	// -> Ids
	variables, delta, err := c.declarations(ctx, delta, []ast.Variable{})
	if err != nil {
		return ast.Declare{}, delta, err
	}

	// -> statement terminator
	if !delta.token.IsStatementTerminator() {
		return ast.Declare{}, delta, unexpectedTokenTypeError(delta, tokens.StatementTerminator)
	}

	// -> delta
	delta, err = c.nextToken()
	if err != nil {
		return ast.Declare{}, delta, err
	}

	return ast.NewDeclare(ctx, variables), delta, nil
}

func (c *Parser) declarations(ctx *context.Context, delta TokenDelta, current []ast.Variable) ([]ast.Variable, TokenDelta, error) {
	// -> id
	if !delta.token.IsIdentifier() {
		return current, delta, unexpectedTokenTypeError(delta, tokens.Identifier)
	}

	identifier := delta.token.Content()

	if ctx.SymbolExists(identifier) {
		print("HEEEE")
		return current, delta, alreadyDeclared(identifier)
	}

	if _, err := ctx.CreateSymbol(identifier); err != nil {
		return current, delta, err
	}

	current = append(current, ast.NewVariable(ctx, identifier))

	// Tenta puxar o próximo separator
	var err error
	delta, err = c.nextToken()
	if err != nil {
		return current, delta, err
	}

	// Existem mais IDs sendo declarados
	if delta.token.IsSeparator() {
		delta, err = c.nextToken()
		if err != nil {
			return current, delta, err
		}

		return c.declarations(ctx, delta, current)
	}

	return current, delta, nil
}
