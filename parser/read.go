package parser

import (
	"isigo/context"
	"isigo/lang"
	"isigo/syntax"
	"isigo/tokens"
)

func (c *Parser) Read(ctx *context.Context, delta TokenDelta) (lang.Read, TokenDelta, error) {
	// -> leia
	if !delta.token.IsReservedWord() {
		return lang.Read{}, delta, unexpectedTokenTypeError(delta, tokens.ReservedWord)
	}

	if !delta.token.Is(syntax.Read) {
		return lang.Read{}, delta, unexpectedContentError(delta, syntax.Read)
	}

	// -> (
	var err error
	delta, err = c.nextToken()
	if err != nil {
		return lang.Read{}, delta, err
	}

	if !delta.token.IsOpenParenthesis() {
		return lang.Read{}, delta, unexpectedTokenTypeError(delta, tokens.OpenParenthesis)
	}

	// -> ID
	delta, err = c.nextToken()
	if err != nil {
		return lang.Read{}, delta, err
	}

	// Precisa ter um e apenas um identificador
	if !delta.token.IsIdentifier() {
		return lang.Read{}, delta, unexpectedTokenTypeError(delta, tokens.Identifier)
	}

	identifier := delta.token.Content()
	if !ctx.SymbolExists(identifier) {
		return lang.Read{}, delta, usedBeforeDeclaration(identifier)
	}

	symbol, err := ctx.RetrieveSymbol(identifier)
	if err != nil {
		return lang.Read{}, delta, err
	}

	// -> )
	delta, err = c.nextToken()
	if err != nil {
		return lang.Read{}, delta, err
	}

	if !delta.token.IsCloseParenthesis() {
		return lang.Read{}, delta, unexpectedTokenTypeError(delta, tokens.CloseParenthesis)
	}

	// -> .
	delta, err = c.nextToken()
	if err != nil {
		return lang.Read{}, delta, err
	}

	if !delta.token.IsStatementTerminator() {
		return lang.Read{}, delta, unexpectedTokenTypeError(delta, tokens.StatementTerminator)
	}

	err = ctx.AssignSymbol(identifier)
	if err != nil {
		return lang.Read{}, delta, err
	}

	// delta
	delta, err = c.nextToken()
	if err != nil {
		return lang.Read{}, delta, err
	}

	return lang.NewRead(ctx, symbol), delta, nil
}
