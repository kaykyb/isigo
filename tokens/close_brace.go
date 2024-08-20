package tokens

func NewCloseBrace(content string) Token {
	return newToken(CloseBrace, content)
}
