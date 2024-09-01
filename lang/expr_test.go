package lang_test

import (
	"testing"

	"isigo/context"
	"isigo/lang"
	"isigo/value_types"

	"github.com/stretchr/testify/assert"
)

type MockTerm struct {
	output        string
	evalValue     any
	resultingType value_types.ValueType
	err           error
}

func (m *MockTerm) Output() (string, error) {
	return m.output, m.err
}

func (m *MockTerm) Eval(ctx *context.Context) (any, error) {
	return m.evalValue, m.err
}

func (m *MockTerm) ResultingType() (value_types.ValueType, error) {
	return m.resultingType, m.err
}

func (m *MockTerm) IsTerm() bool {
	return true
}

func TestTermExpr_Output(t *testing.T) {
	ctx := context.New()
	term := &MockTerm{output: "42"}

	termExpr := lang.NewTermExpr(&ctx, term)

	output, err := termExpr.Output()

	assert.NoError(t, err)
	assert.Equal(t, "42", output)
}

func TestTermExpr_Eval(t *testing.T) {
	ctx := context.New()
	term := &MockTerm{evalValue: 42}

	termExpr := lang.NewTermExpr(&ctx, term)

	val, err := termExpr.Eval(&ctx)

	assert.NoError(t, err)
	assert.Equal(t, 42, val)
}

func TestTermExpr_ResultingType(t *testing.T) {
	ctx := context.New()
	term := &MockTerm{resultingType: value_types.IntegerValueTypeEntity}

	termExpr := lang.NewTermExpr(&ctx, term)

	valType, err := termExpr.ResultingType()

	assert.NoError(t, err)
	assert.Equal(t, value_types.IntegerValueTypeEntity, valType)
}

func TestSumExpr_Output(t *testing.T) {
	ctx := context.New()

	left := &MockExpr{
		output:     "10",
		resultType: value_types.IntegerValueTypeEntity,
	}

	term := &MockTerm{
		output:        "32",
		resultingType: value_types.IntegerValueTypeEntity,
	}

	sumExpr, err := lang.NewSumExpr(&ctx, left, term)
	assert.NoError(t, err)

	output, err := sumExpr.Output()

	assert.NoError(t, err)
	assert.Equal(t, "10 + 32", output)
}

func TestSumExpr_Eval(t *testing.T) {
	ctx := context.New()

	left := &MockExpr{
		value:      10,
		resultType: value_types.IntegerValueTypeEntity,
	}
	term := &MockTerm{
		evalValue:     32,
		resultingType: value_types.IntegerValueTypeEntity,
	}

	sumExpr, err := lang.NewSumExpr(&ctx, left, term)
	assert.NoError(t, err)

	val, err := sumExpr.Eval(&ctx)

	assert.NoError(t, err)
	assert.Equal(t, 42, val)
}

func TestSumExpr_Eval_Float(t *testing.T) {
	ctx := context.New()

	left := &MockExpr{value: 10.5, resultType: value_types.FloatValueTypeEntity}
	term := &MockTerm{evalValue: 1.5, resultingType: value_types.FloatValueTypeEntity}

	sumExpr, err := lang.NewSumExpr(&ctx, left, term)
	assert.NoError(t, err)

	val, err := sumExpr.Eval(&ctx)

	assert.NoError(t, err)
	assert.Equal(t, 12.0, val)
}

func TestSumExpr_ResultingType(t *testing.T) {
	ctx := context.New()

	left := &MockExpr{resultType: value_types.IntegerValueTypeEntity}
	term := &MockTerm{resultingType: value_types.FloatValueTypeEntity}

	sumExpr, err := lang.NewSumExpr(&ctx, left, term)
	assert.NoError(t, err)

	valType, err := sumExpr.ResultingType()

	assert.NoError(t, err)
	assert.Equal(t, value_types.FloatValueTypeEntity, valType)
}

func TestSubtractExpr_Output(t *testing.T) {
	ctx := context.New()

	left := &MockExpr{output: "10", resultType: value_types.IntegerValueTypeEntity}
	term := &MockTerm{output: "32", resultingType: value_types.IntegerValueTypeEntity}

	subExpr, err := lang.NewSubtractExpr(&ctx, left, term)
	assert.NoError(t, err)

	output, err := subExpr.Output()

	assert.NoError(t, err)
	assert.Equal(t, "10 - 32", output)
}

func TestSubtractExpr_Eval(t *testing.T) {
	ctx := context.New()

	left := &MockExpr{value: 50, resultType: value_types.IntegerValueTypeEntity}
	term := &MockTerm{evalValue: 8, resultingType: value_types.IntegerValueTypeEntity}

	subExpr, err := lang.NewSubtractExpr(&ctx, left, term)
	assert.NoError(t, err)

	val, err := subExpr.Eval(&ctx)

	assert.NoError(t, err)
	assert.Equal(t, 58, val)
}

func TestSubtractExpr_Eval_Float(t *testing.T) {
	ctx := context.New()

	left := &MockExpr{value: 10.5, resultType: value_types.IntegerValueTypeEntity}
	term := &MockTerm{evalValue: 1.5, resultingType: value_types.IntegerValueTypeEntity}

	subExpr, err := lang.NewSubtractExpr(&ctx, left, term)
	assert.NoError(t, err)

	val, err := subExpr.Eval(&ctx)

	assert.NoError(t, err)
	assert.Equal(t, 12.0, val)
}

func TestSubtractExpr_ResultingType(t *testing.T) {
	ctx := context.New()

	left := &MockExpr{resultType: value_types.IntegerValueTypeEntity}
	term := &MockTerm{resultingType: value_types.FloatValueTypeEntity}

	subExpr, err := lang.NewSubtractExpr(&ctx, left, term)
	assert.NoError(t, err)

	valType, err := subExpr.ResultingType()

	assert.NoError(t, err)
	assert.Equal(t, value_types.FloatValueTypeEntity, valType)
}
