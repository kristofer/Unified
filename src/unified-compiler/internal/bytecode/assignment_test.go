package bytecode

import (
	"testing"
	"unified-compiler/internal/ast"
)

func TestGenerateAssignmentToMutableVariable(t *testing.T) {
	// fn main() { let mut x = 5; x = 10; return x }
	program := &ast.Program{
		Items: []ast.Item{
			&ast.FunctionDecl{
				Name: "main",
				Body: &ast.Block{
					Statements: []ast.Statement{
						&ast.LetStatement{
							Name: "x",
							Value: &ast.Literal{
								Kind:  ast.LiteralInt,
								Value: "5",
							},
							IsMutable: true,
						},
						&ast.AssignmentStatement{
							Target:   "x",
							Operator: ast.AssignNormal,
							Value: &ast.Literal{
								Kind:  ast.LiteralInt,
								Value: "10",
							},
						},
						&ast.ReturnStatement{
							Value: &ast.Identifier{
								Name: "x",
							},
						},
					},
				},
			},
		},
	}

	generator := NewGenerator()
	bc, err := generator.Generate(program)
	
	if err != nil {
		t.Fatalf("Generation failed: %v", err)
	}

	// Should have PUSH, STORE_LOCAL, PUSH, STORE_LOCAL, LOAD_LOCAL, RETURN_VALUE, HALT
	storeCount := 0
	loadCount := 0
	for _, inst := range bc.Instructions {
		if inst.Op == OpStoreLocal {
			storeCount++
		}
		if inst.Op == OpLoadLocal {
			loadCount++
		}
	}

	if storeCount != 2 {
		t.Errorf("Expected 2 STORE_LOCAL instructions (let + assignment), got %d", storeCount)
	}
	if loadCount != 1 {
		t.Errorf("Expected 1 LOAD_LOCAL instruction (return), got %d", loadCount)
	}
}

func TestGenerateCompoundAssignment(t *testing.T) {
	tests := []struct {
		name     string
		operator ast.AssignOp
		expected OpCode
	}{
		{"AddAssign", ast.AssignAdd, OpAdd},
		{"SubAssign", ast.AssignSub, OpSub},
		{"MulAssign", ast.AssignMul, OpMul},
		{"DivAssign", ast.AssignDiv, OpDiv},
		{"ModAssign", ast.AssignMod, OpMod},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// fn main() { let mut x = 10; x += 5; }
			program := &ast.Program{
				Items: []ast.Item{
					&ast.FunctionDecl{
						Name: "main",
						Body: &ast.Block{
							Statements: []ast.Statement{
								&ast.LetStatement{
									Name: "x",
									Value: &ast.Literal{
										Kind:  ast.LiteralInt,
										Value: "10",
									},
									IsMutable: true,
								},
								&ast.AssignmentStatement{
									Target:   "x",
									Operator: tt.operator,
									Value: &ast.Literal{
										Kind:  ast.LiteralInt,
										Value: "5",
									},
								},
							},
						},
					},
				},
			}

			generator := NewGenerator()
			bc, err := generator.Generate(program)
			
			if err != nil {
				t.Fatalf("Generation failed: %v", err)
			}

			// Find the operator instruction
			found := false
			for _, inst := range bc.Instructions {
				if inst.Op == tt.expected {
					found = true
					break
				}
			}

			if !found {
				t.Errorf("Expected %s instruction not found", tt.expected)
			}

			// Verify we load the variable before the operation for compound assignment
			hasLoadLocal := false
			for _, inst := range bc.Instructions {
				if inst.Op == OpLoadLocal {
					hasLoadLocal = true
					break
				}
			}

			if !hasLoadLocal {
				t.Error("Expected LOAD_LOCAL instruction for compound assignment")
			}
		})
	}
}

func TestGenerateCounterWithMutableVariable(t *testing.T) {
	// fn main() { let mut counter = 0; counter += 1; counter += 1; return counter }
	program := &ast.Program{
		Items: []ast.Item{
			&ast.FunctionDecl{
				Name: "main",
				Body: &ast.Block{
					Statements: []ast.Statement{
						&ast.LetStatement{
							Name: "counter",
							Value: &ast.Literal{
								Kind:  ast.LiteralInt,
								Value: "0",
							},
							IsMutable: true,
						},
						&ast.AssignmentStatement{
							Target:   "counter",
							Operator: ast.AssignAdd,
							Value: &ast.Literal{
								Kind:  ast.LiteralInt,
								Value: "1",
							},
						},
						&ast.AssignmentStatement{
							Target:   "counter",
							Operator: ast.AssignAdd,
							Value: &ast.Literal{
								Kind:  ast.LiteralInt,
								Value: "1",
							},
						},
						&ast.ReturnStatement{
							Value: &ast.Identifier{
								Name: "counter",
							},
						},
					},
				},
			},
		},
	}

	generator := NewGenerator()
	bc, err := generator.Generate(program)
	
	if err != nil {
		t.Fatalf("Generation failed: %v", err)
	}

	// Count ADD instructions (should be 2 for compound assignments)
	addCount := 0
	for _, inst := range bc.Instructions {
		if inst.Op == OpAdd {
			addCount++
		}
	}

	if addCount != 2 {
		t.Errorf("Expected 2 ADD instructions, got %d", addCount)
	}
}

func TestGenerateMultipleMutableVariables(t *testing.T) {
	// fn main() { let mut x = 1; let mut y = 2; x += 3; y += 4; return x + y }
	program := &ast.Program{
		Items: []ast.Item{
			&ast.FunctionDecl{
				Name: "main",
				Body: &ast.Block{
					Statements: []ast.Statement{
						&ast.LetStatement{
							Name: "x",
							Value: &ast.Literal{
								Kind:  ast.LiteralInt,
								Value: "1",
							},
							IsMutable: true,
						},
						&ast.LetStatement{
							Name: "y",
							Value: &ast.Literal{
								Kind:  ast.LiteralInt,
								Value: "2",
							},
							IsMutable: true,
						},
						&ast.AssignmentStatement{
							Target:   "x",
							Operator: ast.AssignAdd,
							Value: &ast.Literal{
								Kind:  ast.LiteralInt,
								Value: "3",
							},
						},
						&ast.AssignmentStatement{
							Target:   "y",
							Operator: ast.AssignAdd,
							Value: &ast.Literal{
								Kind:  ast.LiteralInt,
								Value: "4",
							},
						},
						&ast.ReturnStatement{
							Value: &ast.BinaryExpr{
								Left: &ast.Identifier{
									Name: "x",
								},
								Operator: ast.OperatorAdd,
								Right: &ast.Identifier{
									Name: "y",
								},
							},
						},
					},
				},
			},
		},
	}

	generator := NewGenerator()
	bc, err := generator.Generate(program)
	
	if err != nil {
		t.Fatalf("Generation failed: %v", err)
	}

	// Count STORE_LOCAL (2 for let, 2 for assignments = 4)
	storeCount := 0
	for _, inst := range bc.Instructions {
		if inst.Op == OpStoreLocal {
			storeCount++
		}
	}

	if storeCount != 4 {
		t.Errorf("Expected 4 STORE_LOCAL instructions, got %d", storeCount)
	}
}

func TestGenerateErrorUndefinedVariableInAssignment(t *testing.T) {
	// fn main() { x = 10; }  // x is undefined
	program := &ast.Program{
		Items: []ast.Item{
			&ast.FunctionDecl{
				Name: "main",
				Body: &ast.Block{
					Statements: []ast.Statement{
						&ast.AssignmentStatement{
							Target:   "undefined_var",
							Operator: ast.AssignNormal,
							Value: &ast.Literal{
								Kind:  ast.LiteralInt,
								Value: "10",
							},
						},
					},
				},
			},
		},
	}

	generator := NewGenerator()
	_, err := generator.Generate(program)
	
	if err == nil {
		t.Error("Expected error for assignment to undefined variable")
	}
}
