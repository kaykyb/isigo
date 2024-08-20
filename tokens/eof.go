package tokens

func NewEOF(content string) Token {
	return newToken(EOF, content)
}
