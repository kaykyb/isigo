package lang_test

import (
	"fmt"
	"testing"

	"isigo/context"
	"isigo/lang"
	"isigo/value_types"

	"github.com/stretchr/testify/assert"
)

type MockVariable struct {
	identifier   string
	variableType value_types.ValueType
}

func (m *MockVariable) Output() string {
	return m.identifier + " " + m.variableType.Output()
}

func TestVariableContext_Output(t *testing.T) {
	ctx := context.New()

	variables := []lang.Variable{
		lang.NewVariable(&ctx, "x", value_types.IntegerValueTypeEntity),
		lang.NewVariable(&ctx, "y", value_types.FloatValueTypeEntity),
	}

	declare := lang.NewDeclare(&ctx, variables)

	mockNode := &MockNode{
		output: "x = 42",
	}

	variableContext := lang.NewVariableContext(&ctx, declare, mockNode)

	output, err := variableContext.Output()

	expectedOutput := "var x int64\nvar y float64\nx = 42"
	assert.NoError(t, err)
	assert.Equal(t, expectedOutput, output)
}

func TestVariableContext_Output_Error_Child(t *testing.T) {
	ctx := context.New()

	variables := []lang.Variable{
		lang.NewVariable(&ctx, "x", value_types.IntegerValueTypeEntity),
	}

	declare := lang.NewDeclare(&ctx, variables)

	mockNode := &MockNode{
		err: fmt.Errorf("child error"),
	}

	variableContext := lang.NewVariableContext(&ctx, declare, mockNode)

	output, err := variableContext.Output()

	assert.Error(t, err)
	assert.Equal(t, "", output)
	assert.EqualError(t, err, "child error")
}

func TestVariableContext_Eval(t *testing.T) {
	ctx := context.New()

	variables := []lang.Variable{
		lang.NewVariable(&ctx, "x", value_types.IntegerValueTypeEntity),
	}

	declare := lang.NewDeclare(&ctx, variables)

	mockNode := &MockNode{
		value: "Execution result",
	}

	variableContext := lang.NewVariableContext(&ctx, declare, mockNode)

	val, err := variableContext.Eval(&ctx)

	assert.NoError(t, err)
	assert.Equal(t, "Execution result", val)
}

func TestVariableContext_Eval_Error_Child(t *testing.T) {
	ctx := context.New()

	variables := []lang.Variable{
		lang.NewVariable(&ctx, "x", value_types.IntegerValueTypeEntity),
	}

	declare := lang.NewDeclare(&ctx, variables)

	mockNode := &MockNode{
		err: fmt.Errorf("child error"),
	}

	variableContext := lang.NewVariableContext(&ctx, declare, mockNode)

	val, err := variableContext.Eval(&ctx)

	assert.Error(t, err)
	assert.Nil(t, val)
	assert.EqualError(t, err, "child error")
}
