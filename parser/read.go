package parser

import (
	"isigo/ast"
	"isigo/context"
	"isigo/syntax"
	"isigo/tokens"
)

func (c *Parser) Read(ctx *context.Context, delta TokenDelta) (ast.Read, TokenDelta, error) {
	// -> leia
	if !delta.token.IsReservedWord() {
		return ast.Read{}, delta, unexpectedTokenTypeError(delta, tokens.ReservedWord)
	}

	if !delta.token.Is(syntax.Read) {
		return ast.Read{}, delta, unexpectedContentError(delta, syntax.Read)
	}

	// -> (
	var err error
	delta, err = c.nextToken()
	if err != nil {
		return ast.Read{}, delta, err
	}

	if !delta.token.IsOpenParenthesis() {
		return ast.Read{}, delta, unexpectedTokenTypeError(delta, tokens.OpenParenthesis)
	}

	// -> ID
	delta, err = c.nextToken()
	if err != nil {
		return ast.Read{}, delta, err
	}

	// Precisa ter um e apenas um identificador
	if !delta.token.IsIdentifier() {
		return ast.Read{}, delta, unexpectedTokenTypeError(delta, tokens.Identifier)
	}

	identifier := delta.token.Content()
	if !ctx.SymbolExists(identifier) {
		return ast.Read{}, delta, usedBeforeDeclaration(identifier)
	}

	symbol, err := ctx.RetrieveSymbol(identifier)
	if err != nil {
		return ast.Read{}, delta, err
	}

	// -> )
	delta, err = c.nextToken()
	if err != nil {
		return ast.Read{}, delta, err
	}

	if !delta.token.IsCloseParenthesis() {
		return ast.Read{}, delta, unexpectedTokenTypeError(delta, tokens.CloseParenthesis)
	}

	// -> .
	delta, err = c.nextToken()
	if err != nil {
		return ast.Read{}, delta, err
	}

	if !delta.token.IsStatementTerminator() {
		return ast.Read{}, delta, unexpectedTokenTypeError(delta, tokens.StatementTerminator)
	}

	// delta
	delta, err = c.nextToken()
	if err != nil {
		return ast.Read{}, delta, err
	}

	return ast.NewRead(ctx, symbol), delta, nil
}
