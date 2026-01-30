package wasm

import (
	"testing"
	"unified-compiler/internal/ast"
)

// TestStructGeneration tests struct expression generation
func TestStructGeneration(t *testing.T) {
	g := NewGenerator()
	
	// Create a simple struct expression: Point { x: 10, y: 20 }
	structExpr := &ast.StructExpr{
		Name: "Point",
		FieldInits: []*ast.FieldInit{
			{
				Name: "x",
				Value: &ast.Literal{
					Kind:  ast.LiteralInt,
					Value: "10",
				},
			},
			{
				Name: "y",
				Value: &ast.Literal{
					Kind:  ast.LiteralInt,
					Value: "20",
				},
			},
		},
	}
	
	// Create a simple function to test
	fn := &ast.FunctionDecl{
		Name: "test",
		ReturnType: &ast.TypeReference{Name: "i32"},
		Body: &ast.Block{
			Statements: []ast.Statement{
				&ast.ReturnStatement{
					Value: structExpr,
				},
			},
		},
	}
	
	err := g.addFunction(fn)
	if err != nil {
		t.Fatalf("Failed to generate function: %v", err)
	}
	
	// Check that function was added
	if len(g.module.Functions) != 1 {
		t.Errorf("Expected 1 function, got %d", len(g.module.Functions))
	}
}

// TestListGeneration tests array/list literal generation
func TestListGeneration(t *testing.T) {
	g := NewGenerator()
	
	// Create a list: [1, 2, 3]
	listExpr := &ast.ListExpr{
		Elements: []ast.Expression{
			&ast.Literal{Kind: ast.LiteralInt, Value: "1"},
			&ast.Literal{Kind: ast.LiteralInt, Value: "2"},
			&ast.Literal{Kind: ast.LiteralInt, Value: "3"},
		},
	}
	
	fn := &ast.FunctionDecl{
		Name: "test",
		ReturnType: &ast.TypeReference{Name: "i32"},
		Body: &ast.Block{
			Statements: []ast.Statement{
				&ast.ReturnStatement{
					Value: listExpr,
				},
			},
		},
	}
	
	err := g.addFunction(fn)
	if err != nil {
		t.Fatalf("Failed to generate function: %v", err)
	}
	
	if len(g.module.Functions) != 1 {
		t.Errorf("Expected 1 function, got %d", len(g.module.Functions))
	}
}

// TestStringLiteral tests string literal generation
func TestStringLiteral(t *testing.T) {
	g := NewGenerator()
	
	// Create a string literal
	stringLit := &ast.Literal{
		Kind:  ast.LiteralString,
		Value: "hello",
	}
	
	fn := &ast.FunctionDecl{
		Name: "test",
		ReturnType: &ast.TypeReference{Name: "i32"},
		Body: &ast.Block{
			Statements: []ast.Statement{
				&ast.ReturnStatement{
					Value: stringLit,
				},
			},
		},
	}
	
	err := g.addFunction(fn)
	if err != nil {
		t.Fatalf("Failed to generate function: %v", err)
	}
	
	// Check that string was added to string table
	if len(g.stringTable) != 1 {
		t.Errorf("Expected 1 string in table, got %d", len(g.stringTable))
	}
	
	// Check that string is in data segments
	if len(g.memoryAllocator.GetDataSegments()) != 1 {
		t.Errorf("Expected 1 data segment, got %d", len(g.memoryAllocator.GetDataSegments()))
	}
}

// TestIndexExpr tests array indexing generation
func TestIndexExpr(t *testing.T) {
	g := NewGenerator()
	
	// Create array indexing: arr[0]
	indexExpr := &ast.IndexExpr{
		Object: &ast.Identifier{Name: "arr"},
		Index:  &ast.Literal{Kind: ast.LiteralInt, Value: "0"},
	}
	
	fn := &ast.FunctionDecl{
		Name: "test",
		Parameters: []*ast.Parameter{
			{Name: "arr", Type: &ast.TypeReference{Name: "i32"}},
		},
		ReturnType: &ast.TypeReference{Name: "i64"},
		Body: &ast.Block{
			Statements: []ast.Statement{
				&ast.ReturnStatement{
					Value: indexExpr,
				},
			},
		},
	}
	
	err := g.addFunction(fn)
	if err != nil {
		t.Fatalf("Failed to generate function: %v", err)
	}
}

// TestForLoop tests for loop generation
func TestForLoop(t *testing.T) {
	g := NewGenerator()
	
	// Create a for loop: for i in arr { }
	forStmt := &ast.ForStatement{
		Variable: "i",
		Iterable: &ast.Identifier{Name: "arr"},
		Body: &ast.Block{
			Statements: []ast.Statement{},
		},
	}
	
	fn := &ast.FunctionDecl{
		Name: "test",
		Parameters: []*ast.Parameter{
			{Name: "arr", Type: &ast.TypeReference{Name: "i32"}},
		},
		ReturnType: &ast.TypeReference{Name: "i64"},
		Body: &ast.Block{
			Statements: []ast.Statement{
				forStmt,
				&ast.ReturnStatement{
					Value: &ast.Literal{Kind: ast.LiteralInt, Value: "0"},
				},
			},
		},
	}
	
	err := g.addFunction(fn)
	if err != nil {
		t.Fatalf("Failed to generate function: %v", err)
	}
}

// TestBreakContinue tests break and continue statement generation
func TestBreakContinue(t *testing.T) {
	g := NewGenerator()
	
	// Create a while loop with break
	whileStmt := &ast.WhileStatement{
		Condition: &ast.Literal{Kind: ast.LiteralBool, Value: "true"},
		Body: &ast.Block{
			Statements: []ast.Statement{
				&ast.BreakStatement{},
			},
		},
	}
	
	fn := &ast.FunctionDecl{
		Name: "test",
		ReturnType: &ast.TypeReference{Name: "i64"},
		Body: &ast.Block{
			Statements: []ast.Statement{
				whileStmt,
				&ast.ReturnStatement{
					Value: &ast.Literal{Kind: ast.LiteralInt, Value: "0"},
				},
			},
		},
	}
	
	err := g.addFunction(fn)
	if err != nil {
		t.Fatalf("Failed to generate function: %v", err)
	}
}

// TestMemoryAllocator tests the memory allocator
func TestMemoryAllocator(t *testing.T) {
	ma := NewMemoryAllocator()
	
	// Test basic allocation
	offset1 := ma.Allocate(16)
	if offset1 != 1024 {
		t.Errorf("Expected first allocation at 1024, got %d", offset1)
	}
	
	// Test alignment
	offset2 := ma.Allocate(16)
	if offset2 < offset1+16 {
		t.Errorf("Expected second allocation after first, got offset %d", offset2)
	}
	
	// Test data allocation
	data := []byte("hello")
	offset3 := ma.AllocateWithData(data)
	
	segments := ma.GetDataSegments()
	if len(segments) != 1 {
		t.Errorf("Expected 1 data segment, got %d", len(segments))
	}
	
	if segments[0].Offset != offset3 {
		t.Errorf("Data segment offset mismatch")
	}
}
