package tokens

func NewAssign(content string) Token {
	return newToken(Assign, content)
}
