package lang_test

import (
	"testing"

	"isigo/context"
	"isigo/lang"
	"isigo/value_types"

	"github.com/stretchr/testify/assert"
)

type MockFactor struct {
	output        string
	evalValue     any
	resultingType value_types.ValueType
	err           error
}

func (m *MockFactor) Output() (string, error) {
	return m.output, m.err
}

func (m *MockFactor) Eval(ctx *context.Context) (any, error) {
	return m.evalValue, m.err
}

func (m *MockFactor) ResultingType() (value_types.ValueType, error) {
	return m.resultingType, m.err
}

func (m *MockFactor) IsFactor() bool {
	return true
}

func TestFactorTerm_Output(t *testing.T) {
	ctx := context.New()
	factor := &MockFactor{output: "42"}

	term := lang.NewFactorTerm(&ctx, factor)

	output, err := term.Output()

	assert.NoError(t, err)
	assert.Equal(t, "42", output)
}

func TestFactorTerm_Eval(t *testing.T) {
	ctx := context.New()
	factor := &MockFactor{evalValue: 42}

	term := lang.NewFactorTerm(&ctx, factor)

	val, err := term.Eval(&ctx)

	assert.NoError(t, err)
	assert.Equal(t, 42, val)
}

func TestFactorTerm_ResultingType(t *testing.T) {
	ctx := context.New()
	factor := &MockFactor{resultingType: value_types.IntegerValueTypeEntity}

	term := lang.NewFactorTerm(&ctx, factor)

	valType, err := term.ResultingType()

	assert.NoError(t, err)
	assert.Equal(t, value_types.IntegerValueTypeEntity, valType)
}

func TestMultiplyTerm_Output(t *testing.T) {
	ctx := context.New()
	left := &MockFactor{output: "6"}
	factor := &MockFactor{output: "7"}

	term := lang.NewMultiplyTerm(&ctx, lang.NewFactorTerm(&ctx, left), factor)

	output, err := term.Output()

	assert.NoError(t, err)
	assert.Equal(t, "6 * 7", output)
}

func TestMultiplyTerm_Eval(t *testing.T) {
	ctx := context.New()
	left := &MockFactor{evalValue: int64(6)}
	factor := &MockFactor{evalValue: int64(7)}

	term := lang.NewMultiplyTerm(&ctx, lang.NewFactorTerm(&ctx, left), factor)

	val, err := term.Eval(&ctx)

	assert.NoError(t, err)
	assert.Equal(t, int64(42), val)
}

func TestMultiplyTerm_ResultingType(t *testing.T) {
	ctx := context.New()
	left := &MockFactor{resultingType: value_types.IntegerValueTypeEntity}
	factor := &MockFactor{resultingType: value_types.IntegerValueTypeEntity}

	term := lang.NewMultiplyTerm(&ctx, lang.NewFactorTerm(&ctx, left), factor)

	valType, err := term.ResultingType()

	assert.NoError(t, err)
	assert.Equal(t, value_types.IntegerValueTypeEntity, valType)
}

func TestDivideTerm_Output(t *testing.T) {
	ctx := context.New()
	left := &MockFactor{output: "42"}
	factor := &MockFactor{output: "7"}

	term := lang.NewDivideTerm(&ctx, lang.NewFactorTerm(&ctx, left), factor)

	output, err := term.Output()

	assert.NoError(t, err)
	assert.Equal(t, "float64(42) / float64(7)", output)
}

func TestDivideTerm_Eval(t *testing.T) {
	ctx := context.New()
	left := &MockFactor{evalValue: int64(42)}
	factor := &MockFactor{evalValue: int64(7)}

	term := lang.NewDivideTerm(&ctx, lang.NewFactorTerm(&ctx, left), factor)

	val, err := term.Eval(&ctx)

	assert.NoError(t, err)
	assert.Equal(t, 6.0, val)
}

func TestDivideTerm_EvalFloat(t *testing.T) {
	ctx := context.New()
	left := &MockFactor{evalValue: int64(42)}
	factor := &MockFactor{evalValue: int64(5)}

	term := lang.NewDivideTerm(&ctx, lang.NewFactorTerm(&ctx, left), factor)

	val, err := term.Eval(&ctx)

	assert.NoError(t, err)
	assert.Equal(t, 8.4, val)
}

func TestDivideTerm_ResultingType(t *testing.T) {
	ctx := context.New()
	left := &MockFactor{resultingType: value_types.IntegerValueTypeEntity}
	factor := &MockFactor{resultingType: value_types.IntegerValueTypeEntity}

	term := lang.NewDivideTerm(&ctx, lang.NewFactorTerm(&ctx, left), factor)

	valType, err := term.ResultingType()

	assert.NoError(t, err)
	assert.Equal(t, value_types.FloatValueTypeEntity, valType)
}
