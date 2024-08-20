package tokens

func NewSeparator(content string) Token {
	return newToken(Separator, content)
}
