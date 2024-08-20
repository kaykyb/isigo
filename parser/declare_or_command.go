package parser

import (
	"isigo/ast"
	"isigo/context"
	"isigo/syntax"
)

func (c *Parser) DeclareOrCommand(ctx *context.Context, delta TokenDelta) (ast.Node, TokenDelta, error) {
	if delta.token.IsReservedWord() {
		// -> declare
		if delta.token.Is(syntax.Declare) {
			return c.VariableContext(ctx, delta)
		}

		// -> if
		if delta.token.Is(syntax.If) {
		}

		// -> leia
		if delta.token.Is(syntax.Read) {
			return c.Read(ctx, delta)
		}

		// -> escreva
		if delta.token.Is(syntax.Write) {

		}
	}

	// -> assignment
	if delta.token.IsIdentifier() {
		return c.Assignment(ctx, delta)
	}

	return ast.Block{}, delta, errorf(
		"Esperava uma expressão ou bloco de comando, encontrada %s '%s'.",
		delta.token.FriendlyString(),
		delta.token.Content(),
	)
}
