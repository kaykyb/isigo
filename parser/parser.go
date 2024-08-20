package parser

import (
	"isigo/common"
	"isigo/lexer"
	"isigo/tokens"
)

type Parser struct {
	lexicalAnalysis *lexer.LexicalAnalysis
}

type TokenDelta struct {
	token    tokens.Token
	position common.CodePosition
}

func New(lexicalAnalysis *lexer.LexicalAnalysis) Parser {
	return Parser{
		lexicalAnalysis: lexicalAnalysis,
	}
}

func (c *Parser) nextToken() (TokenDelta, error) {
	token, position, err := c.lexicalAnalysis.NextToken()
	return TokenDelta{
		token:    token,
		position: position,
	}, err
}

func (c *Parser) NextToken() (TokenDelta, error) {
	token, position, err := c.lexicalAnalysis.NextToken()
	return TokenDelta{
		token:    token,
		position: position,
	}, err
}
