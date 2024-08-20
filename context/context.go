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

func (c *Context) CreateSymbol(identifier string) (*symbol.Symbol, error) {
	symbolEntity := &symbol.Symbol{
		Identifier: identifier,
		Type:       value_types.IntegerValueTypeEntity,
	}

	err := c.symbolTable.PutNew(identifier, symbolEntity)
	return symbolEntity, err
}
