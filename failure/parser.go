package failure

func AlreadyDeclared(id string) error {
	return SyntaxErrorf("O símbolo '%s' já foi declarado neste contexto.", id)
}

func UsedBeforeDeclaration(id string) error {
	return SyntaxErrorf("O símbolo '%s' foi usado antes de ser declarado.", id)
}

func UnexpectedTokenError(ttype string) error {
	return SyntaxErrorf("'%s' não é esperado neste contexto.", ttype)
}

func UnexpectedTokenTypeError(tkType, tkContent, expectedType string) error {
	return SyntaxErrorf("Esperava uma %s, encontrada %s '%s'.", expectedType, tkType, tkContent)
}

func UnexpectedTokenContentError(tkContent, expectedContent string) error {
	return SyntaxErrorf("Esperava encontrar '%s', mas foi encontrado '%s'.", expectedContent, tkContent)
}

func ExpressionBlockExpected(tkType, tkContent string) error {
	return SyntaxErrorf("Esperava um expressão ou bloco de comando, encontrado %s '%s'.", tkType, tkContent)
}
