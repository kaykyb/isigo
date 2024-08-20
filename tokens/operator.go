package tokens

func NewOperator(content string) Token {
	return newToken(Operator, content)
}
