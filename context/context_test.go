package context

import (
	"testing"

	"isigo/value_types"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	ctx := New()

	assert.NotNil(t, ctx.symbolTable, "New context should initialize a symbol table.")
	assert.Nil(t, ctx.parent, "New context should not have a parent context.")
}

func TestNewWithParent(t *testing.T) {
	parentCtx := New()
	childCtx := NewWithParent(&parentCtx)

	assert.NotNil(t, childCtx.symbolTable, "New context with parent should initialize a symbol table.")
	assert.Equal(t, &parentCtx, childCtx.parent, "New context with parent should correctly set the parent context.")
}

func TestSymbolExists(t *testing.T) {
	ctx := New()
	ctx.CreateSymbol("x", value_types.IntegerValueTypeEntity)

	assert.True(t, ctx.SymbolExists("x"), "SymbolExists should return true for existing symbols.")
	assert.False(t, ctx.SymbolExists("y"), "SymbolExists should return false for non-existing symbols.")
}

func TestSymbolExistsWithParent(t *testing.T) {
	parentCtx := New()
	_, err := parentCtx.CreateSymbol("x", value_types.IntegerValueTypeEntity)
	assert.NoError(t, err)

	childCtx := NewWithParent(&parentCtx)
	assert.True(t, childCtx.SymbolExists("x"), "SymbolExists should return true for symbols existing in the parent context.")
}

func TestRetrieveSymbol(t *testing.T) {
	ctx := New()
	_, err := ctx.CreateSymbol("x", value_types.IntegerValueTypeEntity)
	assert.NoError(t, err)

	sym, err := ctx.RetrieveSymbol("x")
	assert.NoError(t, err)
	assert.NotNil(t, sym, "RetrieveSymbol should return the symbol if it exists.")
	assert.Equal(t, "x", sym.Identifier, "RetrieveSymbol should return the correct symbol.")

	_, err = ctx.RetrieveSymbol("y")
	assert.Error(t, err, "RetrieveSymbol should return an error if the symbol does not exist.")
}

func TestRetrieveSymbolWithParent(t *testing.T) {
	parentCtx := New()
	_, err := parentCtx.CreateSymbol("x", value_types.IntegerValueTypeEntity)
	assert.NoError(t, err)

	childCtx := NewWithParent(&parentCtx)
	sym, err := childCtx.RetrieveSymbol("x")
	assert.NoError(t, err)
	assert.NotNil(t, sym, "RetrieveSymbol should retrieve the symbol from the parent context.")
	assert.Equal(t, "x", sym.Identifier, "RetrieveSymbol should retrieve the correct symbol from the parent context.")
}

func TestCreateSymbol(t *testing.T) {
	ctx := New()
	sym, err := ctx.CreateSymbol("x", value_types.IntegerValueTypeEntity)
	assert.NoError(t, err)
	assert.NotNil(t, sym, "CreateSymbol should return a new symbol.")
	assert.Equal(t, "x", sym.Identifier, "CreateSymbol should correctly set the symbol identifier.")
}

func TestAssignSymbol(t *testing.T) {
	ctx := New()
	_, err := ctx.CreateSymbol("x", value_types.IntegerValueTypeEntity)
	assert.NoError(t, err)

	err = ctx.AssignSymbol("x")
	assert.NoError(t, err)

	sym, err := ctx.RetrieveSymbol("x")
	assert.NoError(t, err)
	assert.True(t, sym.Assigned, "AssignSymbol should mark the symbol as assigned.")
}

func TestAssignSymbolNonExistent(t *testing.T) {
	ctx := New()
	err := ctx.AssignSymbol("y")
	assert.Error(t, err, "AssignSymbol should return an error when assigning a non-existent symbol.")
}

func TestValidateSymbolUsage(t *testing.T) {
	ctx := New()
	_, err := ctx.CreateSymbol("x", value_types.IntegerValueTypeEntity)
	assert.NoError(t, err)

	err = ctx.ValidateSymbolUsage()
	assert.Error(t, err, "ValidateSymbolUsage should return an error if there are unassigned symbols.")

	err = ctx.AssignSymbol("x")
	assert.NoError(t, err)

	err = ctx.ValidateSymbolUsage()
	assert.NoError(t, err, "ValidateSymbolUsage should not return an error if all symbols are assigned.")
}
