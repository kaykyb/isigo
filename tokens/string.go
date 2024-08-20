package tokens

func NewString(content string) Token {
	return newToken(String, content)
}
