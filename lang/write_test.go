package lang_test

import (
	"fmt"
	"testing"

	"isigo/context"
	"isigo/lang"

	"github.com/stretchr/testify/assert"
)

// Mocking std.Escreva function
var escrevaOutput any

func MockEscreva(val any) {
	escrevaOutput = val
}

func TestWrite_Output(t *testing.T) {
	ctx := context.New()

	mockExpr := &MockExpr{
		output: "\"Hello, World!\"",
	}

	write := lang.NewWrite(&ctx, mockExpr)

	output, err := write.Output()

	expectedOutput := "std.Escreva(\"Hello, World!\")"
	assert.NoError(t, err)
	assert.Equal(t, expectedOutput, output)
}

func TestWrite_Output_Error(t *testing.T) {
	ctx := context.New()

	mockExpr := &MockExpr{
		err: fmt.Errorf("mock error"),
	}

	write := lang.NewWrite(&ctx, mockExpr)

	output, err := write.Output()

	assert.Error(t, err)
	assert.Equal(t, "", output)
	assert.EqualError(t, err, "mock error")
}

// func TestWrite_Eval(t *testing.T) {
// 	ctx := context.New()

// 	mockExpr := &MockExpr{
// 		value: "Hello, World!",
// 	}

// 	write := lang.NewWrite(&ctx, mockExpr)

// 	originalEscreva := std.Escreva
// 	defer func() { std.Escreva = originalEscreva }()
// 	std.Escreva = MockEscreva

// 	_, err := write.Eval(&ctx)

// 	assert.NoError(t, err)
// 	assert.Equal(t, "Hello, World!", escrevaOutput)
// }

func TestWrite_Eval_Error(t *testing.T) {
	ctx := context.New()

	mockExpr := &MockExpr{
		err: fmt.Errorf("mock error"),
	}

	write := lang.NewWrite(&ctx, mockExpr)

	val, err := write.Eval(&ctx)

	assert.Error(t, err)
	assert.Nil(t, val)
	assert.EqualError(t, err, "mock error")
}
