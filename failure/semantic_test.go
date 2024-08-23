package failure

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSemanticErrorf(t *testing.T) {
	err := SemanticErrorf("unexpected error: %s", "test")
	expectedErrorMessage := "[Semantic Error]: unexpected error: test"
	assert.Equal(t, expectedErrorMessage, err.Error(), "SemanticErrorf should format the error message correctly.")
}

func TestTypeNotSumable(t *testing.T) {
	err := TypeNotSumable("inteiro")
	expectedErrorMessage := "[Semantic Error]: Não é possível somar em um 'inteiro'."
	assert.Equal(t, expectedErrorMessage, err.Error(), "TypeNotSumable should return the correct error message.")
}

func TestTypeNotSubtractable(t *testing.T) {
	err := TypeNotSubtractable("decimal")
	expectedErrorMessage := "[Semantic Error]: Não é possível subtrair de um 'decimal'."
	assert.Equal(t, expectedErrorMessage, err.Error(), "TypeNotSubtractable should return the correct error message.")
}

func TestTypeNotMultipliable(t *testing.T) {
	err := TypeNotMultipliable("texto")
	expectedErrorMessage := "[Semantic Error]: Não é possível multiplicar um 'texto'."
	assert.Equal(t, expectedErrorMessage, err.Error(), "TypeNotMultipliable should return the correct error message.")
}

func TestTypeNotDivisible(t *testing.T) {
	err := TypeNotDivisible("bool")
	expectedErrorMessage := "[Semantic Error]: Não é possível dividir um 'bool'."
	assert.Equal(t, expectedErrorMessage, err.Error(), "TypeNotDivisible should return the correct error message.")
}

func TestCannotDivideTypes(t *testing.T) {
	err := CannotDivideTypes("inteiro", "texto")
	expectedErrorMessage := "[Semantic Error]: Não é possível aplicar uma divisão de um 'inteiro' por um 'texto'."
	assert.Equal(t, expectedErrorMessage, err.Error(), "CannotDivideTypes should return the correct error message.")
}

func TestSymbolTypeDiffers(t *testing.T) {
	err := SymbolTypeDiffers("x", "decimal", "inteiro")
	expectedErrorMessage := "[Semantic Error]: Não é possível atribuir um 'inteiro' ao símbolo 'x' (decimal)."
	assert.Equal(t, expectedErrorMessage, err.Error(), "SymbolTypeDiffers should return the correct error message.")
}

func TestSemanticEnvironment(t *testing.T) {
	semantic := Semantic{}
	assert.Equal(t, SemanticEnvironment, semantic.Environment(), "Environment should return the correct SemanticEnvironment.")
}
