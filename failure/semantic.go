package failure

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
