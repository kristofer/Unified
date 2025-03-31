package codegen

import (
	"fmt"
	"strings"
	"testing"
	"unified-compiler/internal/ast"
)

// TestLiteralGeneration tests the generation of LLVM IR for literal values
func TestLiteralGeneration(t *testing.T) {
	// Create a generator instance for testing
	gen := NewCodeGenerator("test_module")
	defer gen.Dispose()

	tests := []struct {
		name     string
		literal  *ast.Literal
		expected string // partial string that should appear in the IR
	}{
		{
			name: "integer literal",
			literal: &ast.Literal{
				Kind:     ast.LiteralInt,
				Value:    "42",
				Position: ast.Position{Line: 1, Column: 1},
			},
			expected: "i32 42",
		},
		{
			name: "float literal",
			literal: &ast.Literal{
				Kind:     ast.LiteralFloat,
				Value:    "3.14",
				Position: ast.Position{Line: 1, Column: 1},
			},
			expected: "double 3.14",
		},
		{
			name: "boolean true",
			literal: &ast.Literal{
				Kind:     ast.LiteralBool,
				Value:    "true",
				Position: ast.Position{Line: 1, Column: 1},
			},
			expected: "i1 true",
		},
		{
			name: "string literal",
			literal: &ast.Literal{
				Kind:     ast.LiteralString,
				Value:    "hello",
				Position: ast.Position{Line: 1, Column: 1},
			},
			expected: "hello",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Generate LLVM Value for literal
			llvmValue := gen.GenerateLiteral(tt.literal)

			// Check if not nil
			if llvmValue.IsNil() {
				t.Fatalf("generated LLVM value is nil")
			}

			// Convert to string and check
			valueStr := llvmValue.String()
			if !strings.Contains(valueStr, tt.expected) {
				t.Errorf("expected IR to contain %q, got %q", tt.expected, valueStr)
			}
		})
	}
}

// TestSimpleFunctionGeneration tests generating a simple function
func TestSimpleFunctionGeneration(t *testing.T) {
	gen := NewCodeGenerator("test_module")
	defer gen.Dispose()

	// Create a simple function that returns an int literal
	funcDecl := &ast.FunctionDecl{
		Name:       "test_func",
		Parameters: []*ast.Parameter{},
		ReturnType: &ast.TypeReference{Name: "i32"},
		Body: &ast.Block{
			Statements: []ast.Statement{
				&ast.ReturnStatement{
					Value: &ast.Literal{
						Kind:  ast.LiteralInt,
						Value: "42",
					},
				},
			},
		},
	}

	// Generate the function
	_, err := gen.GenerateFunction(funcDecl)
	if err != nil {
		t.Fatalf("failed to generate function: %v", err)
	}

	// Get module IR
	ir := gen.module.String()

	// Check if function definition exists
	if !strings.Contains(ir, "define i32 @test_func()") {
		t.Errorf("function definition not found in IR: %s", ir)
	}

	// Check if return value exists
	if !strings.Contains(ir, "ret i32 42") {
		t.Errorf("return statement not found in IR: %s", ir)
	}
}

// TestIfStatement tests generating if statements
func TestIfStatement(t *testing.T) {
	gen := NewCodeGenerator("test_module")
	defer gen.Dispose()

	// Create a function with an if statement
	funcDecl := &ast.FunctionDecl{
		Name:       "test_if",
		Parameters: []*ast.Parameter{},
		ReturnType: &ast.TypeReference{Name: "i32"},
		Body: &ast.Block{
			Statements: []ast.Statement{
				&ast.IfStatement{
					Condition: &ast.Literal{
						Kind:  ast.LiteralBool,
						Value: "true",
					},
					ThenBlock: &ast.Block{
						Statements: []ast.Statement{
							&ast.ReturnStatement{
								Value: &ast.Literal{
									Kind:  ast.LiteralInt,
									Value: "1",
								},
							},
						},
					},
					ElseBlock: &ast.Block{
						Statements: []ast.Statement{
							&ast.ReturnStatement{
								Value: &ast.Literal{
									Kind:  ast.LiteralInt,
									Value: "0",
								},
							},
						},
					},
				},
			},
		},
	}

	// Generate the function
	_, err := gen.GenerateFunction(funcDecl)
	if err != nil {
		t.Fatalf("failed to generate function: %v", err)
	}

	// Get module IR
	ir := gen.module.String()

	// Check for conditional branch
	if !strings.Contains(ir, "br i1") {
		t.Errorf("conditional branch not found in IR: %s", ir)
	}

	// Check for then/else blocks
	expectedBlocks := []string{"then", "else", "ifcont"}
	for _, block := range expectedBlocks {
		if !strings.Contains(ir, block+":") {
			t.Errorf("expected block %q not found in IR: %s", block, ir)
		}
	}
}

// TestBinaryExpressions tests generating binary expressions
func TestBinaryExpressions(t *testing.T) {
	gen := NewCodeGenerator("test_module")
	defer gen.Dispose()

	// Create a function that evaluates a simple binary expression
	funcDecl := &ast.FunctionDecl{
		Name:       "test_bin_op",
		Parameters: []*ast.Parameter{},
		ReturnType: &ast.TypeReference{Name: "i32"},
		Body: &ast.Block{
			Statements: []ast.Statement{
				&ast.ReturnStatement{
					Value: &ast.BinaryExpr{
						Left: &ast.Literal{
							Kind:  ast.LiteralInt,
							Value: "5",
						},
						Operator: ast.OperatorAdd,
						Right: &ast.Literal{
							Kind:  ast.LiteralInt,
							Value: "3",
						},
					},
				},
			},
		},
	}

	// Generate the function
	_, err := gen.GenerateFunction(funcDecl)
	if err != nil {
		t.Fatalf("failed to generate function: %v", err)
	}

	// Get module IR
	ir := gen.module.String()

	// Check for addition operation
	if !strings.Contains(ir, "add") {
		t.Errorf("addition operation not found in IR: %s", ir)
	}

	// Check for return of the result
	if !strings.Contains(ir, "ret i32 %add") {
		t.Errorf("return of addition result not found in IR: %s", ir)
	}
}

// TestVariableDeclaration tests generating variable declarations
func TestVariableDeclaration(t *testing.T) {
	gen := NewCodeGenerator("test_module")
	defer gen.Dispose()

	// Create a function with variable declarations
	funcDecl := &ast.FunctionDecl{
		Name:       "test_vars",
		Parameters: []*ast.Parameter{},
		ReturnType: &ast.TypeReference{Name: "i32"},
		Body: &ast.Block{
			Statements: []ast.Statement{
				// Create a variable
				&ast.VarStatement{
					Name: "x",
					Type: &ast.TypeReference{Name: "i32"},
					Value: &ast.Literal{
						Kind:  ast.LiteralInt,
						Value: "10",
					},
				},
				// Return the variable
				&ast.ReturnStatement{
					Value: &ast.Identifier{
						Name: "x",
					},
				},
			},
		},
	}

	// Generate the function
	_, err := gen.GenerateFunction(funcDecl)
	if err != nil {
		t.Fatalf("failed to generate function: %v", err)
	}

	// Get module IR
	ir := gen.module.String()

	// Check for alloca instruction
	if !strings.Contains(ir, "alloca i32") {
		t.Errorf("alloca instruction not found in IR: %s", ir)
	}

	// Check for store instruction
	if !strings.Contains(ir, "store i32 10") {
		t.Errorf("store instruction not found in IR: %s", ir)
	}

	// Check for load instruction
	if !strings.Contains(ir, "load i32") {
		t.Errorf("load instruction not found in IR: %s", ir)
	}
}

// TestCompleteProgram tests generating a complete program
func TestCompleteProgram(t *testing.T) {
	gen := NewCodeGenerator("test_program")
	defer gen.Dispose()

	// Create a simple program with a function that computes factorial
	program := &ast.Program{
		Items: []ast.Item{
			&ast.FunctionDecl{
				Name: "factorial",
				Parameters: []*ast.Parameter{
					{
						Name: "n",
						Type: &ast.TypeReference{Name: "i32"},
					},
				},
				ReturnType: &ast.TypeReference{Name: "i32"},
				Body: &ast.Block{
					Statements: []ast.Statement{
						// if n <= 1 return 1
						&ast.IfStatement{
							Condition: &ast.BinaryExpr{
								Left:     &ast.Identifier{Name: "n"},
								Operator: ast.OperatorLe,
								Right:    &ast.Literal{Kind: ast.LiteralInt, Value: "1"},
							},
							ThenBlock: &ast.Block{
								Statements: []ast.Statement{
									&ast.ReturnStatement{
										Value: &ast.Literal{Kind: ast.LiteralInt, Value: "1"},
									},
								},
							},
						},
						// return n * factorial(n-1)
						&ast.ReturnStatement{
							Value: &ast.BinaryExpr{
								Left:     &ast.Identifier{Name: "n"},
								Operator: ast.OperatorMul,
								Right: &ast.BinaryExpr{
									Left:     &ast.Identifier{Name: "n"},
									Operator: ast.OperatorSub,
									Right:    &ast.Literal{Kind: ast.LiteralInt, Value: "1"},
								},
							},
						},
					},
				},
			},
		},
	}

	// Generate IR for the program
	ir, err := gen.Generate(program)
	if err != nil {
		t.Fatalf("failed to generate program: %v", err)
	}

	// Verify the important parts of the IR
	expectedParts := []string{
		"define i32 @factorial(i32 %n)",
		"alloca i32",
		"icmp sle",
		"br i1",
		"ret i32 1",
		"mul",
		"sub",
	}

	for _, part := range expectedParts {
		if !strings.Contains(ir, part) {
			t.Errorf("expected IR to contain %q, but it does not:\n%s", part, ir)
		}
	}
}

// TestGeneratorErrors tests error cases
func TestGeneratorErrors(t *testing.T) {
	gen := NewCodeGenerator("test_errors")
	defer gen.Dispose()

	// Test undefined variable
	funcWithUndefinedVar := &ast.FunctionDecl{
		Name:       "test_error",
		Parameters: []*ast.Parameter{},
		ReturnType: &ast.TypeReference{Name: "i32"},
		Body: &ast.Block{
			Statements: []ast.Statement{
				&ast.ReturnStatement{
					Value: &ast.Identifier{
						Name: "undefined_variable", // This variable doesn't exist
					},
				},
			},
		},
	}

	// This should produce an error
	_, err := gen.GenerateFunction(funcWithUndefinedVar)
	if err == nil {
		t.Errorf("expected error for undefined variable, but got none")
	}
}

// TestCodeGeneration shows how to directly test the generator
func TestCodeGeneration() {
	// Create a simple AST for testing
	program := &ast.Program{
		Items: []ast.Item{
			&ast.FunctionDecl{
				Name:       "test",
				Parameters: []*ast.Parameter{},
				Body: &ast.Block{
					Statements: []ast.Statement{
						&ast.ReturnStatement{
							Value: &ast.Literal{
								Kind:  ast.LiteralInt,
								Value: "42",
							},
						},
					},
				},
			},
		},
	}

	// Generate IR
	generator := NewCodeGenerator("test_module")
	ir, err := generator.Generate(program)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Println(ir)
	}
	generator.Dispose()
}
