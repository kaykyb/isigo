package failure

import "fmt"

type Syntax struct {
	err string
}

func (e Syntax) Error() string {
	return fmt.Sprintf("[Syntax Error]: %s", e.err)
}

func (e Syntax) Environment() Environment {
	return SyntaxEnvironment
}

func SyntaxErrorf(format string, args ...interface{}) error {
	return Syntax{
		err: fmt.Sprintf(format, args...),
	}
}

func AlreadyDeclared(id string) error {
	return SyntaxErrorf("O símbolo '%s' já foi declarado neste contexto.", id)
}

func UsedBeforeDeclaration(id string) error {
	return SyntaxErrorf("O símbolo '%s' foi usado antes de ser declarado.", id)
}

func UsedBeforeAssignment(id string) error {
	return SyntaxErrorf("O símbolo '%s' foi usado antes de ser inicializado.", id)
}

func NeverUsed(id string) error {
	return SyntaxErrorf("O símbolo '%s' foi declarado mas nunca foi usado.", id)
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
