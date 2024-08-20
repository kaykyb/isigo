package tokens

func NewOpenParenthesis(content string) Token {
	return newToken(OpenParenthesis, content)
}
