package failure

import "fmt"

type Semantic struct {
	err string
}

func (e Semantic) Error() string {
	return fmt.Sprintf("[Semantic Error]: %s", e.err)
}

func (e Semantic) Environment() Environment {
	return SemanticEnvironment
}

func SemanticErrorf(format string, args ...interface{}) error {
	return Semantic{
		err: fmt.Sprintf(format, args...),
	}
}

func TypeNotSumable(ltype string) error {
	return SemanticErrorf("Não é possível somar em um '%s'.", ltype)
}

func TypeNotSubtractable(ltype string) error {
	return SemanticErrorf("Não é possível subtrair de um '%s'.", ltype)
}

func TypeNotMultipliable(ltype string) error {
	return SemanticErrorf("Não é possível multiplicar um '%s'.", ltype)
}

func TypeNotDivisible(ltype string) error {
	return SemanticErrorf("Não é possível dividir um '%s'.", ltype)
}

func CannotDivideTypes(ltype string, rtype string) error {
	return SemanticErrorf("Não é possível aplicar uma divisão de um '%s' por um '%s'.", ltype, rtype)
}

func CannotCompareTypes(ltype string, rtype string) error {
	return SemanticErrorf("Não é possível comparar '%s' com um '%s'.", ltype, rtype)
}

func SymbolTypeDiffers(identifier, identifierType, ltype string) error {
	return SemanticErrorf("Não é possível atribuir um '%s' ao símbolo '%s' (%s).", ltype, identifier, identifierType)
}
