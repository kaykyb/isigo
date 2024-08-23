package lang_test

import (
	"testing"

	"isigo/context"
	"isigo/lang"
	"isigo/value_types"

	"github.com/stretchr/testify/assert"
)

func TestDeclare_Output(t *testing.T) {
	ctx := context.New()

	variables := []lang.Variable{
		lang.NewVariable(&ctx, "x", value_types.IntegerValueTypeEntity),
		lang.NewVariable(&ctx, "y", value_types.FloatValueTypeEntity),
	}

	declare := lang.NewDeclare(&ctx, variables)

	output, err := declare.Output()

	assert.NoError(t, err)
	expectedOutput := "var x int64\nvar y float64"
	assert.Equal(t, expectedOutput, output)
}

func TestDeclare_Eval(t *testing.T) {
	ctx := context.New()

	variables := []lang.Variable{
		lang.NewVariable(&ctx, "x", value_types.IntegerValueTypeEntity),
		lang.NewVariable(&ctx, "y", value_types.FloatValueTypeEntity),
	}

	declare := lang.NewDeclare(&ctx, variables)

	val, err := declare.Eval(&ctx)

	assert.NoError(t, err)
	assert.Nil(t, val)

	// Verifica que as variáveis foram criadas no contexto
	symbolX, err := ctx.RetrieveSymbol("x")
	assert.NoError(t, err)
	assert.NotNil(t, symbolX)
	assert.Equal(t, value_types.IntegerValueTypeEntity, symbolX.Type)

	symbolY, err := ctx.RetrieveSymbol("y")
	assert.NoError(t, err)
	assert.NotNil(t, symbolY)
	assert.Equal(t, value_types.FloatValueTypeEntity, symbolY.Type)
}

func TestDeclare_Eval_Error(t *testing.T) {
	ctx := context.New()

	variables := []lang.Variable{
		lang.NewVariable(&ctx, "x", value_types.IntegerValueTypeEntity),
	}

	// Simula um cenário onde a variável já existe
	_, err := ctx.CreateSymbol("x", value_types.IntegerValueTypeEntity)
	assert.NoError(t, err)

	declare := lang.NewDeclare(&ctx, variables)

	_, err = declare.Eval(&ctx)

	assert.Error(t, err)
	assert.EqualError(t, err, "Símbolo x já existe no contexto atual")
}
