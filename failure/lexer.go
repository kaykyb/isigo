package failure

func ExpectedEndQuote() error {
	return LexerErrorf("Final de string '\"' esperado, mas não foi encontrado.")
}

func MalformedAssignmentOperator(c rune) error {
	return LexerErrorf("Operador de atribuição mal-formado. Esperado '=', encontrado %c", c)
}

func UnexpectedCharacter(c rune) error {
	return LexerErrorf("'%c' não é um caractere esperado neste contexto.", c)
}
