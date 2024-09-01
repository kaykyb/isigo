package parser

import (
	"isigo/context"
	"isigo/lang"
)

func (c *Parser) ExecutionContext(ctx *context.Context, delta TokenDelta) (lang.ExecutionContext, TokenDelta, error) {
	if c.isRepl && delta.token.IsEOF() {
		return lang.NewExecutionContext(ctx, []lang.Node{lang.NewReplEnd(ctx)}), delta, nil
	}

	var children []lang.Node
	var child lang.Node
	var err error

	for {
		child, delta, err = c.DeclareOrCommand(ctx, delta)
		if err != nil {
			return lang.ExecutionContext{}, delta, err
		}

		children = append(children, child)

		if isBlockTerminator(delta.token) {
			if c.isRepl {
				children = append(children, lang.NewReplEnd(ctx))
			}

			return lang.NewExecutionContext(ctx, children), delta, nil
		}
	}
}
