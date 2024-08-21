package parser

import (
	"isigo/context"
	"isigo/lang"
	"isigo/syntax"
	"isigo/tokens"
	"isigo/value_types"
)

func (c *Parser) Declare(ctx *context.Context, delta TokenDelta) (lang.Declare, TokenDelta, error) {
	// -> declare
	if !delta.token.IsReservedWord() {
		return lang.Declare{}, delta, unexpectedTokenTypeError(delta, tokens.ReservedWord)
	}

	if !delta.token.Is(syntax.Declare) {
		return lang.Declare{}, delta, unexpectedContentError(delta, syntax.Declare)
	}

	var err error
	delta, err = c.nextToken()
	if err != nil {
		return lang.Declare{}, delta, err
	}

	// -> Ids
	variables, delta, err := c.declarations(ctx, delta, []lang.Variable{})
	if err != nil {
		return lang.Declare{}, delta, err
	}

	// -> statement terminator
	if !delta.token.IsStatementTerminator() {
		return lang.Declare{}, delta, unexpectedTokenTypeError(delta, tokens.StatementTerminator)
	}

	// -> delta
	delta, err = c.nextToken()
	if err != nil {
		return lang.Declare{}, delta, err
	}

	return lang.NewDeclare(ctx, variables), delta, nil
}

func (c *Parser) declarations(ctx *context.Context, delta TokenDelta, current []lang.Variable) ([]lang.Variable, TokenDelta, error) {
	// -> id
	if !delta.token.IsIdentifier() {
		return current, delta, unexpectedTokenTypeError(delta, tokens.Identifier)
	}

	identifier := delta.token.Content()

	if ctx.SymbolExists(identifier) {
		return current, delta, alreadyDeclared(identifier)
	}

	// Tipo
	var err error
	delta, err = c.nextToken()
	if err != nil {
		return current, delta, err
	}

	if !delta.token.IsTypeT() {
		return current, delta, unexpectedTokenTypeError(delta, tokens.TypeT)
	}

	typeEntity := value_types.TypeEntityForTypeT(delta.token.Content())
	if _, err := ctx.CreateSymbol(identifier, typeEntity); err != nil {
		return current, delta, err
	}

	current = append(current, lang.NewVariable(ctx, identifier, typeEntity))

	// Tenta puxar o próximo separator
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
