package context

import (
	"fmt"
	"isigo/symbol"
	"isigo/value_types"
)

type Context struct {
	parent      *Context
	symbolTable symbol.Table
}

func New() Context {
	return Context{
		symbolTable: symbol.NewTable(),
	}
}

func NewWithParent(parent *Context) Context {
	return Context{
		parent:      parent,
		symbolTable: symbol.NewTable(),
	}
}

func (c *Context) SymbolExists(identifier string) bool {
	symbolExistsInsideThisContext := c.symbolTable.Exists(identifier)

	if symbolExistsInsideThisContext {
		return true
	}

	if c.parent != nil {
		return c.parent.SymbolExists(identifier)
	}

	return false
}

func (c *Context) RetrieveSymbol(identifier string) (*symbol.Symbol, error) {
	symbolInsideThisContext := c.symbolTable.Retrieve(identifier)

	if symbolInsideThisContext != nil {
		return symbolInsideThisContext, nil
	}

	// Tenta procurar nos contextos superiores
	if c.parent != nil {
		return c.parent.RetrieveSymbol(identifier)
	}

	return nil, fmt.Errorf("O símbolo %s não existe no contexto atual", identifier)
}

func (c *Context) CreateSymbol(identifier string, typeT value_types.ValueType) (*symbol.Symbol, error) {
	symbolEntity := &symbol.Symbol{
		Identifier: identifier,
		Type:       typeT,
	}

	err := c.symbolTable.PutNew(identifier, symbolEntity)
	return symbolEntity, err
}

func (c *Context) AssignSymbol(identifier string) error {
	symbolEntity, err := c.RetrieveSymbol(identifier)
	if err != nil {
		return err
	}

	symbolEntity.Assigned = true
	return nil
}

func (c *Context) ValidateSymbolUsage() error {
	return c.symbolTable.ValidateSymbolUsage()
}
