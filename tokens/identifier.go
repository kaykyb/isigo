package tokens

func NewIdentifier(content string) Token {
	return newToken(Identifier, content)
}
