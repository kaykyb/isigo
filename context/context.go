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
	if c.symbolTable.Exists(identifier) {
		return true
	}

	return c.parent != nil && c.parent.SymbolExists(identifier)
}

func (c *Context) RetrieveSymbol(identifier string) (*symbol.Symbol, error) {
	if sym := c.symbolTable.Retrieve(identifier); sym != nil {
		return sym, nil
	}

	if c.parent != nil {
		return c.parent.RetrieveSymbol(identifier)
	}

	return nil, fmt.Errorf("o símbolo %s não existe no contexto atual", identifier)
}

func (c *Context) CreateSymbol(identifier string, typeT value_types.ValueType) (*symbol.Symbol, error) {
	symbolEntity := &symbol.Symbol{
		Identifier: identifier,
		Type:       typeT,
	}

	return symbolEntity, c.symbolTable.PutNew(identifier, symbolEntity)
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
