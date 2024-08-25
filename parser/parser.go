package parser

import (
	"isigo/common"
	"isigo/context"
	"isigo/lang"
	"isigo/lexer"
	"isigo/tokens"
)

type Parser struct {
	lexicalAnalysis *lexer.Lexer
}

type TokenDelta struct {
	token    tokens.Token
	position common.CodePosition
}

func NewTokenDelta(t tokens.Token, p common.CodePosition) TokenDelta {
	return TokenDelta{
		token:    t,
		position: p,
	}
}

func New(lexicalAnalysis *lexer.Lexer) Parser {
	return Parser{
		lexicalAnalysis: lexicalAnalysis,
	}
}

func (d TokenDelta) Position() common.CodePosition {
	return d.position
}

func (c *Parser) nextToken() (TokenDelta, error) {
	token, position, err := c.lexicalAnalysis.NextToken()
	return TokenDelta{
		token:    token,
		position: position,
	}, err
}

func (c *Parser) ParseProgram(ctx *context.Context) (lang.Program, TokenDelta, error) {
	delta, err := c.nextToken()
	if err != nil {
		return lang.Program{}, delta, err
	}

	return c.Prog(ctx, delta)
}
