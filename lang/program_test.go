package lang_test

import (
	"fmt"
	"testing"

	"isigo/context"
	"isigo/lang"

	"github.com/stretchr/testify/assert"
)

func TestProgram_Output(t *testing.T) {
	ctx := context.New()
	mockNode := &MockNode{
		output: "fmt.Println(\"Hello, World!\")",
	}

	block := lang.NewBlock(&ctx, mockNode)
	program := lang.NewProgram(&ctx, block)

	output, err := program.Output()

	expectedOutput := "package main\n\nimport \"isigoprogram/std\"\n\nfunc main() {\n\tfmt.Println(\"Hello, World!\")\n}\n"

	assert.NoError(t, err)
	assert.Equal(t, expectedOutput, output)
}

func TestProgram_Output_Error(t *testing.T) {
	ctx := context.New()
	mockNode := &MockNode{
		err: fmt.Errorf("mock error"),
	}

	block := lang.NewBlock(&ctx, mockNode)
	program := lang.NewProgram(&ctx, block)

	output, err := program.Output()

	assert.Error(t, err)
	assert.Equal(t, "", output)
	assert.EqualError(t, err, "mock error")
}

func TestProgram_Eval(t *testing.T) {
	ctx := context.New()
	mockNode := &MockNode{
		value: "Execution result",
	}

	block := lang.NewBlock(&ctx, mockNode)
	program := lang.NewProgram(&ctx, block)

	val, err := program.Eval(&ctx)

	assert.NoError(t, err)
	assert.Equal(t, "Execution result", val)
}

func TestProgram_Eval_Error(t *testing.T) {
	ctx := context.New()
	mockNode := &MockNode{
		err: fmt.Errorf("mock error"),
	}

	block := lang.NewBlock(&ctx, mockNode)
	program := lang.NewProgram(&ctx, block)

	val, err := program.Eval(&ctx)

	assert.Error(t, err)
	assert.Nil(t, val)
	assert.EqualError(t, err, "mock error")
}
