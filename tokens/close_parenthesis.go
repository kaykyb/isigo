package tokens

func NewCloseParenthesis(content string) Token {
	return newToken(CloseParenthesis, content)
}
