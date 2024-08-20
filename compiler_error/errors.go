package compiler_error

import (
	"fmt"
)

func TypeNotSumable(ltype string) error {
	return fmt.Errorf("Não é possível somar em um '%s'.", ltype)
}

func TypeNotSubtractable(ltype string) error {
	return fmt.Errorf("Não é possível subtrair de um '%s'.", ltype)
}

func TypeNotMultipliable(ltype string) error {
	return fmt.Errorf("Não é possível multiplicar um '%s'.", ltype)
}

func TypeNotDivisible(ltype string) error {
	return fmt.Errorf("Não é possível dividir um '%s'.", ltype)
}

func CannotDivideTypes(left string, by string) error {
	return fmt.Errorf("Não é possível aplicar uma divisão de um '%s' por um '%s'.", left, by)
}
