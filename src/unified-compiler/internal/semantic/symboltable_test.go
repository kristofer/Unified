package semantic

import (
	"testing"
	"unified-compiler/internal/ast"
)

func TestSymbolTableDefine(t *testing.T) {
	st := NewSymbolTable()
	
	// Define a variable
	err := st.Define("x", &ast.TypeReference{Name: "Int"}, false)
	if err != nil {
		t.Errorf("Failed to define variable: %v", err)
	}
	
	// Look it up
	symbol, err := st.Lookup("x")
	if err != nil {
		t.Errorf("Failed to lookup variable: %v", err)
	}
	if symbol.Name != "x" {
		t.Errorf("Expected name 'x', got '%s'", symbol.Name)
	}
	if symbol.Mutable {
		t.Error("Expected immutable variable")
	}
}

func TestSymbolTableMutable(t *testing.T) {
	st := NewSymbolTable()
	
	// Define a mutable variable
	err := st.Define("counter", &ast.TypeReference{Name: "Int"}, true)
	if err != nil {
		t.Errorf("Failed to define variable: %v", err)
	}
	
	// Look it up and check mutability
	symbol, err := st.Lookup("counter")
	if err != nil {
		t.Errorf("Failed to lookup variable: %v", err)
	}
	if !symbol.Mutable {
		t.Error("Expected mutable variable")
	}
}

func TestSymbolTableDuplicateError(t *testing.T) {
	st := NewSymbolTable()
	
	// Define a variable
	st.Define("x", &ast.TypeReference{Name: "Int"}, false)
	
	// Try to define it again in the same scope
	err := st.Define("x", &ast.TypeReference{Name: "Int"}, false)
	if err == nil {
		t.Error("Expected error for duplicate variable definition")
	}
}

func TestSymbolTableUndefinedError(t *testing.T) {
	st := NewSymbolTable()
	
	// Look up undefined variable
	_, err := st.Lookup("undefined")
	if err == nil {
		t.Error("Expected error for undefined variable")
	}
}

func TestSymbolTableScopes(t *testing.T) {
	st := NewSymbolTable()
	
	// Define in global scope
	st.Define("x", &ast.TypeReference{Name: "Int"}, false)
	
	// Enter a new scope
	st.EnterScope()
	if st.CurrentScope() != 1 {
		t.Errorf("Expected scope 1, got %d", st.CurrentScope())
	}
	
	// Should still be able to see parent scope variable
	_, err := st.Lookup("x")
	if err != nil {
		t.Error("Should be able to lookup parent scope variable")
	}
	
	// Define a new variable in nested scope
	st.Define("y", &ast.TypeReference{Name: "String"}, false)
	
	// Exit scope
	st.ExitScope()
	if st.CurrentScope() != 0 {
		t.Errorf("Expected scope 0, got %d", st.CurrentScope())
	}
	
	// y should not be visible anymore
	_, err = st.Lookup("y")
	if err == nil {
		t.Error("Should not be able to lookup child scope variable")
	}
}

func TestSymbolTableShadowing(t *testing.T) {
	st := NewSymbolTable()
	
	// Define x in global scope
	st.Define("x", &ast.TypeReference{Name: "Int"}, false)
	
	// Enter nested scope
	st.EnterScope()
	
	// Shadow x with a different type
	err := st.Define("x", &ast.TypeReference{Name: "String"}, true)
	if err != nil {
		t.Errorf("Should be able to shadow variable: %v", err)
	}
	
	// Lookup should find the shadowed version
	symbol, _ := st.Lookup("x")
	if symbol.Type.(*ast.TypeReference).Name != "String" {
		t.Error("Should find shadowed variable")
	}
	if !symbol.Mutable {
		t.Error("Shadowed variable should be mutable")
	}
	
	// Exit scope
	st.ExitScope()
	
	// Should see original x again
	symbol, _ = st.Lookup("x")
	if symbol.Type.(*ast.TypeReference).Name != "Int" {
		t.Error("Should find original variable after exiting scope")
	}
	if symbol.Mutable {
		t.Error("Original variable should be immutable")
	}
}

func TestSymbolTableMultipleScopes(t *testing.T) {
	st := NewSymbolTable()
	
	// Global scope
	st.Define("a", &ast.TypeReference{Name: "Int"}, false)
	
	// Level 1
	st.EnterScope()
	st.Define("b", &ast.TypeReference{Name: "Int"}, false)
	
	// Level 2
	st.EnterScope()
	st.Define("c", &ast.TypeReference{Name: "Int"}, false)
	
	// Should see all three variables
	if _, err := st.Lookup("a"); err != nil {
		t.Error("Should see variable a")
	}
	if _, err := st.Lookup("b"); err != nil {
		t.Error("Should see variable b")
	}
	if _, err := st.Lookup("c"); err != nil {
		t.Error("Should see variable c")
	}
	
	// Exit to level 1
	st.ExitScope()
	if _, err := st.Lookup("c"); err == nil {
		t.Error("Should not see variable c")
	}
	
	// Exit to global
	st.ExitScope()
	if _, err := st.Lookup("b"); err == nil {
		t.Error("Should not see variable b")
	}
}

func TestSymbolTableExitGlobalScopeError(t *testing.T) {
	st := NewSymbolTable()
	
	// Try to exit global scope
	err := st.ExitScope()
	if err == nil {
		t.Error("Should not be able to exit global scope")
	}
}
