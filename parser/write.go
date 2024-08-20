package parser

import (
	"isigo/ast"
	"isigo/context"
	"isigo/syntax"
	"isigo/tokens"
)

func (c *Parser) Write(ctx *context.Context, delta TokenDelta) (ast.Write, TokenDelta, error) {
	// -> escreva
	if !delta.token.IsReservedWord() {
		return ast.Write{}, delta, unexpectedTokenTypeError(delta, tokens.ReservedWord)
	}

	if !delta.token.Is(syntax.Write) {
		return ast.Write{}, delta, unexpectedContentError(delta, syntax.Write)
	}

	// -> (
	var err error
	delta, err = c.nextToken()
	if err != nil {
		return ast.Write{}, delta, err
	}

	if !delta.token.IsOpenParenthesis() {
		return ast.Write{}, delta, unexpectedTokenTypeError(delta, tokens.OpenParenthesis)
	}

	// -> ID
	delta, err = c.nextToken()
	if err != nil {
		return ast.Write{}, delta, err
	}

	// Precisa ter um e apenas um identificador
	if !delta.token.IsIdentifier() {
		return ast.Write{}, delta, unexpectedTokenTypeError(delta, tokens.Identifier)
	}

	identifier := delta.token.Content()
	if !ctx.SymbolExists(identifier) {
		return ast.Write{}, delta, usedBeforeDeclaration(identifier)
	}

	symbol, err := ctx.RetrieveSymbol(identifier)
	if err != nil {
		return ast.Write{}, delta, err
	}

	// -> )
	delta, err = c.nextToken()
	if err != nil {
		return ast.Write{}, delta, err
	}

	if !delta.token.IsCloseParenthesis() {
		return ast.Write{}, delta, unexpectedTokenTypeError(delta, tokens.CloseParenthesis)
	}

	// -> .
	delta, err = c.nextToken()
	if err != nil {
		return ast.Write{}, delta, err
	}

	if !delta.token.IsStatementTerminator() {
		return ast.Write{}, delta, unexpectedTokenTypeError(delta, tokens.StatementTerminator)
	}

	// delta
	delta, err = c.nextToken()
	if err != nil {
		return ast.Write{}, delta, err
	}

	return ast.NewWrite(ctx, symbol), delta, nil
}
