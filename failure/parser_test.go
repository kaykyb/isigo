package failure

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSyntaxErrorf(t *testing.T) {
	err := SyntaxErrorf("unexpected error: %s", "test")
	expectedErrorMessage := "[Syntax Error]: unexpected error: test"
	assert.Equal(t, expectedErrorMessage, err.Error(), "SyntaxErrorf should format the error message correctly.")
}

func TestAlreadyDeclared(t *testing.T) {
	err := AlreadyDeclared("x")
	expectedErrorMessage := "[Syntax Error]: O símbolo 'x' já foi declarado neste contexto."
	assert.Equal(t, expectedErrorMessage, err.Error(), "AlreadyDeclared should return the correct error message.")
}

func TestUsedBeforeDeclaration(t *testing.T) {
	err := UsedBeforeDeclaration("y")
	expectedErrorMessage := "[Syntax Error]: O símbolo 'y' foi usado antes de ser declarado."
	assert.Equal(t, expectedErrorMessage, err.Error(), "UsedBeforeDeclaration should return the correct error message.")
}

func TestUsedBeforeAssignment(t *testing.T) {
	err := UsedBeforeAssignment("z")
	expectedErrorMessage := "[Syntax Error]: O símbolo 'z' foi usado antes de ser inicializado."
	assert.Equal(t, expectedErrorMessage, err.Error(), "UsedBeforeAssignment should return the correct error message.")
}

func TestNeverUsed(t *testing.T) {
	err := NeverUsed("a")
	expectedErrorMessage := "[Syntax Error]: O símbolo 'a' foi declarado mas nunca foi usado."
	assert.Equal(t, expectedErrorMessage, err.Error(), "NeverUsed should return the correct error message.")
}

func TestUnexpectedTokenError(t *testing.T) {
	err := UnexpectedTokenError("TOKEN_TYPE")
	expectedErrorMessage := "[Syntax Error]: 'TOKEN_TYPE' não é esperado neste contexto."
	assert.Equal(t, expectedErrorMessage, err.Error(), "UnexpectedTokenError should return the correct error message.")
}

func TestUnexpectedTokenTypeError(t *testing.T) {
	err := UnexpectedTokenTypeError("int", "5", "float")
	expectedErrorMessage := "[Syntax Error]: Esperava uma float, encontrada int '5'."
	assert.Equal(t, expectedErrorMessage, err.Error(), "UnexpectedTokenTypeError should return the correct error message.")
}

func TestUnexpectedTokenContentError(t *testing.T) {
	err := UnexpectedTokenContentError("else", "if")
	expectedErrorMessage := "[Syntax Error]: Esperava encontrar 'if', mas foi encontrado 'else'."
	assert.Equal(t, expectedErrorMessage, err.Error(), "UnexpectedTokenContentError should return the correct error message.")
}

func TestExpressionBlockExpected(t *testing.T) {
	err := ExpressionBlockExpected("identifier", "x")
	expectedErrorMessage := "[Syntax Error]: Esperava um expressão ou bloco de comando, encontrado identifier 'x'."
	assert.Equal(t, expectedErrorMessage, err.Error(), "ExpressionBlockExpected should return the correct error message.")
}

func TestSyntaxEnvironment(t *testing.T) {
	syntax := Syntax{}
	assert.Equal(t, SyntaxEnvironment, syntax.Environment(), "Environment should return the correct SyntaxEnvironment.")
}
