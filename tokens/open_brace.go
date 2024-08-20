package tokens

func NewOpenBrace(content string) Token {
	return newToken(OpenBrace, content)
}
