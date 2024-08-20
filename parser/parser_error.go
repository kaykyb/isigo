package parser

import (
	"isigo/failure"
	"isigo/tokens"
)

func alreadyDeclared(identifier string) error {
	return failure.AlreadyDeclared(identifier)
}

func usedBeforeDeclaration(identifier string) error {
	return failure.UsedBeforeDeclaration(identifier)
}

func noMatchTypeError(delta TokenDelta) error {
	return failure.UnexpectedTokenError(delta.token.FriendlyString())
}

func unexpectedTokenTypeError(delta TokenDelta, expectedType tokens.Type) error {
	return failure.UnexpectedTokenTypeError(delta.token.FriendlyString(), delta.token.Content(), tokens.FriendlyString(expectedType))
}

func unexpectedContentError(delta TokenDelta, expectedContent string) error {
	return failure.UnexpectedTokenContentError(delta.token.Content(), expectedContent)
}

func expressionBlockExpected(delta TokenDelta) error {
	return failure.ExpressionBlockExpected(delta.token.FriendlyString(), delta.token.Content())
}
