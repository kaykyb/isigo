package lang_test

import (
	"fmt"
	"testing"

	"isigo/context"
	"isigo/lang"
	"isigo/symbol"
	"isigo/value_types"

	"github.com/stretchr/testify/assert"
)

type MockExpr struct {
	resultType value_types.ValueType
	value      any
	output     string
	err        error
}

func (m *MockExpr) ResultingType() (value_types.ValueType, error) {
	return m.resultType, m.err
}

func (m *MockExpr) Output() (string, error) {
	return m.output, m.err
}

func (m *MockExpr) Eval(ctx *context.Context) (any, error) {
	return m.value, m.err
}

func (m *MockExpr) IsExpr() bool {
	return true
}

func TestAssignment_Output(t *testing.T) {
	ctx := context.New()
	sym := &symbol.Symbol{
		Identifier: "x",
		Type:       value_types.IntegerValueTypeEntity,
	}

	expr := &MockExpr{
		output: "42",
	}

	assignment := lang.NewAssignment(&ctx, sym, expr)

	output, err := assignment.Output()

	assert.NoError(t, err)
	assert.Equal(t, "x = 42", output)
}

func TestAssignment_Output_Error(t *testing.T) {
	ctx := context.New()
	sym := &symbol.Symbol{
		Identifier: "x",
		Type:       value_types.IntegerValueTypeEntity,
	}

	expr := &MockExpr{
		output: "",
		err:    fmt.Errorf("mock error"),
	}

	assignment := lang.NewAssignment(&ctx, sym, expr)

	_, err := assignment.Output()

	assert.Error(t, err)
	assert.EqualError(t, err, "mock error")
}

func TestAssignment_Eval(t *testing.T) {
	ctx := context.New()
	sym := &symbol.Symbol{
		Identifier: "x",
		Type:       value_types.IntegerValueTypeEntity,
	}

	expr := &MockExpr{
		value: 42,
	}

	assignment := lang.NewAssignment(&ctx, sym, expr)

	val, err := assignment.Eval(&ctx)

	assert.NoError(t, err)
	assert.Nil(t, val)
	assert.True(t, sym.Assigned)
	assert.Equal(t, 42, sym.RuntimeValue())
}

func TestAssignment_Eval_Error(t *testing.T) {
	ctx := context.New()
	sym := &symbol.Symbol{
		Identifier: "x",
		Type:       value_types.IntegerValueTypeEntity,
	}

	expr := &MockExpr{
		err: fmt.Errorf("mock error"),
	}

	assignment := lang.NewAssignment(&ctx, sym, expr)

	_, err := assignment.Eval(&ctx)

	assert.Error(t, err)
	assert.EqualError(t, err, "mock error")
	assert.False(t, sym.Assigned)
}
