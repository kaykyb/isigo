package symbol

import (
	"fmt"
	"isigo/failure"
)

type Table struct {
	symbols map[string]*Symbol
}

func NewTable() Table {
	return Table{
		symbols: make(map[string]*Symbol),
	}
}

func (s *Table) Exists(identifier string) bool {
	_, ok := s.symbols[identifier]
	return ok
}

func (s *Table) Retrieve(identifier string) *Symbol {
	return s.symbols[identifier]
}

func (s *Table) PutNew(identifier string, symbol *Symbol) error {
	if _, exists := s.symbols[identifier]; exists {
		return fmt.Errorf("Símbolo %s já existe no contexto atual", identifier)
	}

	s.symbols[identifier] = symbol
	return nil
}

func (s *Table) ValidateSymbolUsage() error {
	for _, entity := range s.symbols {
		if !entity.Assigned {
			return failure.NeverUsed(entity.Identifier)
		}
	}

	return nil
}
