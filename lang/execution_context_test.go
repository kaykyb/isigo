package lang_test

import (
	"fmt"
	"testing"

	"isigo/context"
	"isigo/lang"

	"github.com/stretchr/testify/assert"
)

func TestExecutionContext_Output(t *testing.T) {
	ctx := context.New()

	children := []lang.Node{
		&MockNode{output: "child 1"},
		&MockNode{output: "child 2"},
		&MockNode{output: "child 3"},
	}

	execContext := lang.NewExecutionContext(&ctx, children)

	output, err := execContext.Output()

	assert.NoError(t, err)
	expectedOutput := "child 1\nchild 2\nchild 3"
	assert.Equal(t, expectedOutput, output)
}

func TestExecutionContext_Output_Error(t *testing.T) {
	ctx := context.New()

	children := []lang.Node{
		&MockNode{output: "child 1"},
		&MockNode{err: fmt.Errorf("mock error")},
		&MockNode{output: "child 3"},
	}

	execContext := lang.NewExecutionContext(&ctx, children)

	_, err := execContext.Output()

	assert.Error(t, err)
	assert.EqualError(t, err, "mock error")
}

func TestExecutionContext_Eval(t *testing.T) {
	ctx := context.New()

	children := []lang.Node{
		&MockNode{value: 1},
		&MockNode{value: 2},
		&MockNode{value: 3},
	}

	execContext := lang.NewExecutionContext(&ctx, children)

	lastReturnVal, err := execContext.Eval(&ctx)

	assert.NoError(t, err)
	assert.Equal(t, 3, lastReturnVal)
}

func TestExecutionContext_Eval_Error(t *testing.T) {
	ctx := context.New()

	children := []lang.Node{
		&MockNode{value: 1},
		&MockNode{err: fmt.Errorf("mock error")},
		&MockNode{value: 3},
	}

	execContext := lang.NewExecutionContext(&ctx, children)

	_, err := execContext.Eval(&ctx)

	assert.Error(t, err)
	assert.EqualError(t, err, "mock error")
}
