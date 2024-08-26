package parser

import (
	"isigo/context"
	"isigo/lang"
	"isigo/syntax"
)

func (c *Parser) ReplDeclareOrCommand(ctx *context.Context) (lang.Node, TokenDelta, error) {
	delta, err := c.nextToken()
	if err != nil {
		return lang.Block{}, delta, err
	}

	return c.DeclareOrCommand(ctx, delta)
}

func (c *Parser) DeclareOrCommand(ctx *context.Context, delta TokenDelta) (lang.Node, TokenDelta, error) {
	if delta.token.IsReservedWord() {
		// -> declare
		if delta.token.Is(syntax.Declare) {
			return c.VariableContext(ctx, delta)
		}

		// -> if
		if delta.token.Is(syntax.If) {
			return c.IfStatement(ctx, delta)
		}

		// -> leia
		if delta.token.Is(syntax.Read) {
			return c.Read(ctx, delta)
		}

		// -> escreva
		if delta.token.Is(syntax.Write) {
			return c.Write(ctx, delta)
		}
	}

	// -> assignment
	if delta.token.IsIdentifier() {
		return c.Assignment(ctx, delta)
	}

	return lang.Block{}, delta, expressionBlockExpected(delta)
}
