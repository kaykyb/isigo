package tokens

func NewReservedWord(content string) Token {
	return newToken(ReservedWord, content)
}
