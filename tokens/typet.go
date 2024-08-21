package tokens

func NewTypeT(content string) Token {
	return newToken(TypeT, content)
}
