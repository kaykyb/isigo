package parser

import (
	"isigo/context"
	"isigo/lang"
	"isigo/syntax"
)

func (c *Parser) VariableContext(ctx *context.Context, delta TokenDelta) (lang.Node, TokenDelta, error) {
	// -> cmd
	if !delta.token.IsReservedWord() || !delta.token.Is(syntax.Declare) {
		return c.ExecutionContext(ctx, delta)
	}

	// -> declare
	newContext := context.NewWithParent(ctx)

	declare, delta, err := c.Declare(&newContext, delta)
	if err != nil {
		return lang.VariableContext{}, delta, err
	}

	// -> fim bloco
	if isBlockTerminator(delta.token) {
		if !c.isRepl {
			err = newContext.ValidateSymbolUsage()
			if err != nil {
				return lang.VariableContext{}, delta, err
			}
		}

		return declare, delta, nil
	}

	// -> declare ou expr
	insideVariableContext, delta, err := c.VariableContext(&newContext, delta)
	if err != nil {
		return lang.VariableContext{}, delta, err
	}

	if !c.isRepl {
		err = newContext.ValidateSymbolUsage()
		if err != nil {
			return lang.VariableContext{}, delta, err
		}
	}

	return lang.NewVariableContext(ctx, declare, insideVariableContext), delta, err
}
