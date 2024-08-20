package tokens

func NewDecimal(content string) Token {
	return newToken(Decimal, content)
}
