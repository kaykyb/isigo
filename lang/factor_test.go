package lang_test

import (
	"testing"

	"isigo/context"
	"isigo/lang"
	"isigo/symbol"
	"isigo/value_types"

	"github.com/stretchr/testify/assert"
)

func TestIntegerFactor_Output(t *testing.T) {
	ctx := context.New()
	intFactor := lang.NewIntegerFactor(&ctx, 42)

	output, err := intFactor.Output()

	assert.NoError(t, err)
	assert.Equal(t, "42", output)
}

func TestIntegerFactor_Eval(t *testing.T) {
	ctx := context.New()
	intFactor := lang.NewIntegerFactor(&ctx, 42)

	val, err := intFactor.Eval(&ctx)

	assert.NoError(t, err)
	assert.Equal(t, 42, val)
}

func TestIntegerFactor_ResultingType(t *testing.T) {
	ctx := context.New()
	intFactor := lang.NewIntegerFactor(&ctx, 42)

	valType, err := intFactor.ResultingType()

	assert.NoError(t, err)
	assert.Equal(t, value_types.IntegerValueTypeEntity, valType)
}

func TestFloatFactor_Output(t *testing.T) {
	ctx := context.New()
	floatFactor := lang.NewFloatFactor(&ctx, 42.5)

	output, err := floatFactor.Output()

	assert.NoError(t, err)
	assert.Equal(t, "42.500000", output)
}

func TestFloatFactor_Eval(t *testing.T) {
	ctx := context.New()
	floatFactor := lang.NewFloatFactor(&ctx, 42.5)

	val, err := floatFactor.Eval(&ctx)

	assert.NoError(t, err)
	assert.Equal(t, 42.5, val)
}

func TestFloatFactor_ResultingType(t *testing.T) {
	ctx := context.New()
	floatFactor := lang.NewFloatFactor(&ctx, 42.5)

	valType, err := floatFactor.ResultingType()

	assert.NoError(t, err)
	assert.Equal(t, value_types.IntegerValueTypeEntity, valType)
}

func TestStringFactor_Output(t *testing.T) {
	ctx := context.New()
	stringFactor := lang.NewStringFactor(&ctx, "Helena Deland")

	output, err := stringFactor.Output()

	assert.NoError(t, err)
	assert.Equal(t, "\"Helena Deland\"", output)
}

func TestStringFactor_Eval(t *testing.T) {
	ctx := context.New()
	stringFactor := lang.NewStringFactor(&ctx, "Helena Deland")

	val, err := stringFactor.Eval(&ctx)

	assert.NoError(t, err)
	assert.Equal(t, "Helena Deland", val)
}

func TestStringFactor_ResultingType(t *testing.T) {
	ctx := context.New()
	stringFactor := lang.NewStringFactor(&ctx, "Thom Yorke")

	valType, err := stringFactor.ResultingType()

	assert.NoError(t, err)
	assert.Equal(t, value_types.StringValueTypeEntity, valType)
}

func TestSymbolFactor_Output(t *testing.T) {
	ctx := context.New()
	symbolEntity := &symbol.Symbol{
		Identifier: "x",
		Type:       value_types.IntegerValueTypeEntity,
	}
	symbolFactor := lang.NewSymbolFactor(&ctx, symbolEntity)

	output, err := symbolFactor.Output()

	assert.NoError(t, err)
	assert.Equal(t, "x", output)
}

func TestSymbolFactor_Eval(t *testing.T) {
	ctx := context.New()
	symbolEntity := &symbol.Symbol{
		Identifier: "x",
		Type:       value_types.IntegerValueTypeEntity,
		Assigned:   true,
	}
	symbolEntity.AssignRuntimeValue(42)
	symbolFactor := lang.NewSymbolFactor(&ctx, symbolEntity)

	val, err := symbolFactor.Eval(&ctx)

	assert.NoError(t, err)
	assert.Equal(t, 42, val)
}

func TestSymbolFactor_ResultingType(t *testing.T) {
	ctx := context.New()
	symbolEntity := &symbol.Symbol{
		Identifier: "x",
		Type:       value_types.IntegerValueTypeEntity,
	}
	symbolFactor := lang.NewSymbolFactor(&ctx, symbolEntity)

	valType, err := symbolFactor.ResultingType()

	assert.NoError(t, err)
	assert.Equal(t, value_types.IntegerValueTypeEntity, valType)
}

func TestExpressionFactor_Output(t *testing.T) {
	ctx := context.New()
	mockExpr := &MockExpr{output: "x + y"}

	exprFactor := lang.NewExpressionFactor(&ctx, mockExpr)

	output, err := exprFactor.Output()

	assert.NoError(t, err)
	assert.Equal(t, "(x + y)", output)
}

func TestExpressionFactor_Eval(t *testing.T) {
	ctx := context.New()
	mockExpr := &MockExpr{value: 42}

	exprFactor := lang.NewExpressionFactor(&ctx, mockExpr)

	val, err := exprFactor.Eval(&ctx)

	assert.NoError(t, err)
	assert.Equal(t, 42, val)
}

func TestExpressionFactor_ResultingType(t *testing.T) {
	ctx := context.New()
	mockExpr := &MockExpr{resultType: value_types.IntegerValueTypeEntity}

	exprFactor := lang.NewExpressionFactor(&ctx, mockExpr)

	valType, err := exprFactor.ResultingType()

	assert.NoError(t, err)
	assert.Equal(t, value_types.IntegerValueTypeEntity, valType)
}
