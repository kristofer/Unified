package semantic

import (
	"fmt"
	"unified-compiler/internal/ast"
)

// Symbol represents a variable or parameter in the symbol table
type Symbol struct {
	Name    string
	Type    ast.Type
	Mutable bool
	Defined bool
}

// SymbolTable manages variable scopes and symbols
type SymbolTable struct {
	scopes  []map[string]*Symbol
	current int
}

// NewSymbolTable creates a new symbol table
func NewSymbolTable() *SymbolTable {
	return &SymbolTable{
		scopes:  []map[string]*Symbol{make(map[string]*Symbol)},
		current: 0,
	}
}

// EnterScope creates a new nested scope
func (st *SymbolTable) EnterScope() {
	st.scopes = append(st.scopes, make(map[string]*Symbol))
	st.current++
}

// ExitScope exits the current scope
func (st *SymbolTable) ExitScope() error {
	if st.current == 0 {
		return fmt.Errorf("cannot exit global scope")
	}
	st.scopes = st.scopes[:len(st.scopes)-1]
	st.current--
	return nil
}

// Define adds a new symbol to the current scope
func (st *SymbolTable) Define(name string, typ ast.Type, mutable bool) error {
	// Check if symbol already exists in current scope
	if _, exists := st.scopes[st.current][name]; exists {
		return fmt.Errorf("variable '%s' already defined in this scope", name)
	}
	
	st.scopes[st.current][name] = &Symbol{
		Name:    name,
		Type:    typ,
		Mutable: mutable,
		Defined: true,
	}
	return nil
}

// Lookup searches for a symbol in the current and parent scopes
func (st *SymbolTable) Lookup(name string) (*Symbol, error) {
	// Search from current scope back to global scope
	for i := st.current; i >= 0; i-- {
		if symbol, exists := st.scopes[i][name]; exists {
			return symbol, nil
		}
	}
	return nil, fmt.Errorf("undefined variable '%s'", name)
}

// CurrentScope returns the current scope index
func (st *SymbolTable) CurrentScope() int {
	return st.current
}

// ScopeCount returns the total number of scopes
func (st *SymbolTable) ScopeCount() int {
	return len(st.scopes)
}
