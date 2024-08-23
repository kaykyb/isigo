package lang_test

import (
	"fmt"
	"testing"

	"isigo/context"
	"isigo/lang"

	"github.com/stretchr/testify/assert"
)

type MockNode struct {
	output string
	value  any
	err    error
}

func (m *MockNode) Output() (string, error) {
	return m.output, m.err
}

func (m *MockNode) Eval(ctx *context.Context) (any, error) {
	return m.value, m.err
}

func TestBlock_Output(t *testing.T) {
	ctx := context.New()
	mockNode := &MockNode{
		output: "child output",
	}

	block := lang.NewBlock(&ctx, mockNode)

	output, err := block.Output()

	assert.NoError(t, err)
	assert.Equal(t, "child output", output)
}

func TestBlock_Output_Error(t *testing.T) {
	ctx := context.New()
	mockNode := &MockNode{
		output: "",
		err:    fmt.Errorf("mock error"),
	}

	block := lang.NewBlock(&ctx, mockNode)

	_, err := block.Output()

	assert.Error(t, err)
	assert.EqualError(t, err, "mock error")
}

func TestBlock_Eval(t *testing.T) {
	ctx := context.New()
	mockNode := &MockNode{
		value: 42,
	}

	block := lang.NewBlock(&ctx, mockNode)

	val, err := block.Eval(&ctx)

	assert.NoError(t, err)
	assert.Equal(t, 42, val)
}

func TestBlock_Eval_Error(t *testing.T) {
	ctx := context.New()
	mockNode := &MockNode{
		value: nil,
		err:   fmt.Errorf("mock error"),
	}

	block := lang.NewBlock(&ctx, mockNode)

	_, err := block.Eval(&ctx)

	assert.Error(t, err)
	assert.EqualError(t, err, "mock error")
}
