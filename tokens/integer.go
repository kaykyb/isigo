package tokens

func NewInteger(content string) Token {
	return newToken(Integer, content)
}
