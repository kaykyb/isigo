package tokens

func NewStatementTerminator(content string) Token {
	return newToken(StatementTerminator, content)
}
