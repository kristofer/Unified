package bytecode

import (
	"testing"
	"unified-compiler/internal/ast"
)

func TestGenerateWhileLoop(t *testing.T) {
	// while i < 10 { i = i + 1 }
	program := &ast.Program{
		Items: []ast.Item{
			&ast.FunctionDecl{
				Name: "main",
				Body: &ast.Block{
					Statements: []ast.Statement{
						// let mut i = 0
						&ast.LetStatement{
							Name:      "i",
							IsMutable: true,
							Value: &ast.Literal{
								Kind:  ast.LiteralInt,
								Value: "0",
							},
						},
						// while i < 10 { i = i + 1 }
						&ast.WhileStatement{
							Condition: &ast.BinaryExpr{
								Left: &ast.Identifier{
									Name: "i",
								},
								Operator: ast.OperatorLt,
								Right: &ast.Literal{
									Kind:  ast.LiteralInt,
									Value: "10",
								},
							},
							Body: &ast.Block{
								Statements: []ast.Statement{
									// For simplicity, just increment placeholder
									&ast.ExprStatement{
										Expression: &ast.Literal{
											Kind:  ast.LiteralInt,
											Value: "1",
										},
									},
								},
							},
						},
						// return i
						&ast.ReturnStatement{
							Value: &ast.Identifier{
								Name: "i",
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

	// Check for jump instructions (while loop should have jumps)
	hasJump := false
	hasJumpIfFalse := false
	for _, inst := range bc.Instructions {
		if inst.Op == OpJump {
			hasJump = true
		}
		if inst.Op == OpJumpIfFalse {
			hasJumpIfFalse = true
		}
	}

	if !hasJump {
		t.Error("While loop should generate OpJump instruction")
	}
	if !hasJumpIfFalse {
		t.Error("While loop should generate OpJumpIfFalse instruction")
	}
}

func TestGenerateForLoop(t *testing.T) {
	// for i in 0..10 { ... }
	program := &ast.Program{
		Items: []ast.Item{
			&ast.FunctionDecl{
				Name: "main",
				Body: &ast.Block{
					Statements: []ast.Statement{
						// for i in 0..10 { ... }
						&ast.ForStatement{
							Variable: "i",
							Iterable: &ast.BinaryExpr{
								Left: &ast.Literal{
									Kind:  ast.LiteralInt,
									Value: "0",
								},
								Operator: ast.OperatorRange,
								Right: &ast.Literal{
									Kind:  ast.LiteralInt,
									Value: "10",
								},
							},
							Body: &ast.Block{
								Statements: []ast.Statement{
									// Empty body for now
								},
							},
						},
						// return 0
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
	}

	generator := NewGenerator()
	bc, err := generator.Generate(program)

	if err != nil {
		t.Fatalf("Generation failed: %v", err)
	}

	// Check for jump instructions and comparison
	hasJump := false
	hasLt := false
	for _, inst := range bc.Instructions {
		if inst.Op == OpJump {
			hasJump = true
		}
		if inst.Op == OpLt {
			hasLt = true
		}
	}

	if !hasJump {
		t.Error("For loop should generate OpJump instruction")
	}
	if !hasLt {
		t.Error("For loop should generate OpLt for range comparison")
	}
}

func TestGenerateForLoopInclusive(t *testing.T) {
	// for i in 0..=10 { ... }
	program := &ast.Program{
		Items: []ast.Item{
			&ast.FunctionDecl{
				Name: "main",
				Body: &ast.Block{
					Statements: []ast.Statement{
						// for i in 0..=10 { ... }
						&ast.ForStatement{
							Variable: "i",
							Iterable: &ast.BinaryExpr{
								Left: &ast.Literal{
									Kind:  ast.LiteralInt,
									Value: "0",
								},
								Operator: ast.OperatorRangeIncl,
								Right: &ast.Literal{
									Kind:  ast.LiteralInt,
									Value: "10",
								},
							},
							Body: &ast.Block{
								Statements: []ast.Statement{},
							},
						},
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
	}

	generator := NewGenerator()
	bc, err := generator.Generate(program)

	if err != nil {
		t.Fatalf("Generation failed: %v", err)
	}

	// Check for OpLe (less than or equal) for inclusive range
	hasLe := false
	for _, inst := range bc.Instructions {
		if inst.Op == OpLe {
			hasLe = true
		}
	}

	if !hasLe {
		t.Error("Inclusive for loop should generate OpLe instruction")
	}
}

func TestGenerateLoopStatement(t *testing.T) {
	// loop { break; }
	program := &ast.Program{
		Items: []ast.Item{
			&ast.FunctionDecl{
				Name: "main",
				Body: &ast.Block{
					Statements: []ast.Statement{
						&ast.LoopStatement{
							Body: &ast.Block{
								Statements: []ast.Statement{
									&ast.BreakStatement{},
								},
							},
						},
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
	}

	generator := NewGenerator()
	bc, err := generator.Generate(program)

	if err != nil {
		t.Fatalf("Generation failed: %v", err)
	}

	// Check for jump instructions
	hasJump := false
	for _, inst := range bc.Instructions {
		if inst.Op == OpJump {
			hasJump = true
		}
	}

	if !hasJump {
		t.Error("Loop statement should generate OpJump instruction")
	}
}

func TestGenerateBreakStatement(t *testing.T) {
	// while true { break; }
	program := &ast.Program{
		Items: []ast.Item{
			&ast.FunctionDecl{
				Name: "main",
				Body: &ast.Block{
					Statements: []ast.Statement{
						&ast.WhileStatement{
							Condition: &ast.Literal{
								Kind:  ast.LiteralBool,
								Value: "true",
							},
							Body: &ast.Block{
								Statements: []ast.Statement{
									&ast.BreakStatement{},
								},
							},
						},
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
	}

	generator := NewGenerator()
	bc, err := generator.Generate(program)

	if err != nil {
		t.Fatalf("Generation failed: %v", err)
	}

	// Break should generate a jump
	jumpCount := 0
	for _, inst := range bc.Instructions {
		if inst.Op == OpJump {
			jumpCount++
		}
	}

	if jumpCount < 2 {
		t.Errorf("Break statement should generate at least 2 jumps (break + loop back), got %d", jumpCount)
	}
}

func TestGenerateContinueStatement(t *testing.T) {
	// while true { continue; }
	program := &ast.Program{
		Items: []ast.Item{
			&ast.FunctionDecl{
				Name: "main",
				Body: &ast.Block{
					Statements: []ast.Statement{
						&ast.LoopStatement{
							Body: &ast.Block{
								Statements: []ast.Statement{
									&ast.ContinueStatement{},
									&ast.BreakStatement{}, // Add break to make it terminate
								},
							},
						},
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
	}

	generator := NewGenerator()
	bc, err := generator.Generate(program)

	if err != nil {
		t.Fatalf("Generation failed: %v", err)
	}

	// Continue should generate jumps
	jumpCount := 0
	for _, inst := range bc.Instructions {
		if inst.Op == OpJump {
			jumpCount++
		}
	}

	if jumpCount < 2 {
		t.Errorf("Continue statement should generate jumps, got %d", jumpCount)
	}
}

func TestBreakOutsideLoop(t *testing.T) {
	// break; (without loop)
	program := &ast.Program{
		Items: []ast.Item{
			&ast.FunctionDecl{
				Name: "main",
				Body: &ast.Block{
					Statements: []ast.Statement{
						&ast.BreakStatement{},
					},
				},
			},
		},
	}

	generator := NewGenerator()
	_, err := generator.Generate(program)

	if err == nil {
		t.Error("Expected error for break outside of loop")
	}
}

func TestContinueOutsideLoop(t *testing.T) {
	// continue; (without loop)
	program := &ast.Program{
		Items: []ast.Item{
			&ast.FunctionDecl{
				Name: "main",
				Body: &ast.Block{
					Statements: []ast.Statement{
						&ast.ContinueStatement{},
					},
				},
			},
		},
	}

	generator := NewGenerator()
	_, err := generator.Generate(program)

	if err == nil {
		t.Error("Expected error for continue outside of loop")
	}
}

func TestGenerateNestedLoops(t *testing.T) {
	// while i < 10 { while j < 10 { j = j + 1 } }
	program := &ast.Program{
		Items: []ast.Item{
			&ast.FunctionDecl{
				Name: "main",
				Body: &ast.Block{
					Statements: []ast.Statement{
						&ast.LetStatement{
							Name:      "i",
							IsMutable: true,
							Value: &ast.Literal{
								Kind:  ast.LiteralInt,
								Value: "0",
							},
						},
						&ast.WhileStatement{
							Condition: &ast.BinaryExpr{
								Left: &ast.Identifier{
									Name: "i",
								},
								Operator: ast.OperatorLt,
								Right: &ast.Literal{
									Kind:  ast.LiteralInt,
									Value: "10",
								},
							},
							Body: &ast.Block{
								Statements: []ast.Statement{
									&ast.LetStatement{
										Name:      "j",
										IsMutable: true,
										Value: &ast.Literal{
											Kind:  ast.LiteralInt,
											Value: "0",
										},
									},
									&ast.WhileStatement{
										Condition: &ast.BinaryExpr{
											Left: &ast.Identifier{
												Name: "j",
											},
											Operator: ast.OperatorLt,
											Right: &ast.Literal{
												Kind:  ast.LiteralInt,
												Value: "10",
											},
										},
										Body: &ast.Block{
											Statements: []ast.Statement{},
										},
									},
								},
							},
						},
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
	}

	generator := NewGenerator()
	bc, err := generator.Generate(program)

	if err != nil {
		t.Fatalf("Generation failed for nested loops: %v", err)
	}

	// Should have multiple jumps for nested loops
	jumpCount := 0
	for _, inst := range bc.Instructions {
		if inst.Op == OpJump || inst.Op == OpJumpIfFalse {
			jumpCount++
		}
	}

	if jumpCount < 4 {
		t.Errorf("Nested loops should generate at least 4 jump instructions, got %d", jumpCount)
	}
}
