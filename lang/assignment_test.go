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
	sym, _ := ctx.CreateSymbol("x", value_types.IntegerValueTypeEntity)

	expr := &MockExpr{
		output:     "42",
		resultType: value_types.IntegerValueTypeEntity,
	}

	assignment, err := lang.NewAssignment(&ctx, sym, expr)
	assert.NoError(t, err)

	output, err := assignment.Output()

	assert.NoError(t, err)
	assert.Equal(t, "x = int64(42)", output)
}

func TestAssignment_Output_Error(t *testing.T) {
	ctx := context.New()
	sym, _ := ctx.CreateSymbol("x", value_types.IntegerValueTypeEntity)

	expr := &MockExpr{
		output:     "",
		resultType: value_types.IntegerValueTypeEntity,
		err:        fmt.Errorf("mock error"),
	}

	_, err := lang.NewAssignment(&ctx, sym, expr)
	assert.Error(t, err)
	assert.EqualError(t, err, "mock error")
}

func TestAssignment_Eval(t *testing.T) {
	ctx := context.New()
	sym, _ := ctx.CreateSymbol("x", value_types.IntegerValueTypeEntity)

	expr := &MockExpr{
		value:      42,
		resultType: value_types.IntegerValueTypeEntity,
	}

	assignment, err := lang.NewAssignment(&ctx, sym, expr)
	assert.NoError(t, err)

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
		err:        fmt.Errorf("mock error"),
		resultType: value_types.IntegerValueTypeEntity,
	}

	_, err := lang.NewAssignment(&ctx, sym, expr)
	assert.Error(t, err)
	assert.EqualError(t, err, "mock error")
}
