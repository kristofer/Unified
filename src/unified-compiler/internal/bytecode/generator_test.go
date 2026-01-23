package bytecode

import (
	"testing"
	"unified-compiler/internal/ast"
)

func TestGenerateSimpleFunction(t *testing.T) {
	// fn main() -> Int { return 42 }
	program := &ast.Program{
		Items: []ast.Item{
			&ast.FunctionDecl{
				Name: "main",
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

	generator := NewGenerator()
	bc, err := generator.Generate(program)
	
	if err != nil {
		t.Fatalf("Generation failed: %v", err)
	}

	if len(bc.Functions) != 1 {
		t.Errorf("Expected 1 function, got %d", len(bc.Functions))
	}

	if _, ok := bc.Functions["main"]; !ok {
		t.Error("main function not found")
	}

	// Should have: PUSH 42, RETURN_VALUE, HALT
	if len(bc.Instructions) < 3 {
		t.Errorf("Expected at least 3 instructions, got %d", len(bc.Instructions))
	}
}

func TestGenerateBinaryExpression(t *testing.T) {
	tests := []struct {
		name     string
		operator ast.OperatorType
		expected OpCode
	}{
		{"Add", ast.OperatorAdd, OpAdd},
		{"Sub", ast.OperatorSub, OpSub},
		{"Mul", ast.OperatorMul, OpMul},
		{"Div", ast.OperatorDiv, OpDiv},
		{"Mod", ast.OperatorMod, OpMod},
		{"Eq", ast.OperatorEq, OpEq},
		{"Ne", ast.OperatorNe, OpNe},
		{"Lt", ast.OperatorLt, OpLt},
		{"Le", ast.OperatorLe, OpLe},
		{"Gt", ast.OperatorGt, OpGt},
		{"Ge", ast.OperatorGe, OpGe},
		{"And", ast.OperatorAnd, OpAnd},
		{"Or", ast.OperatorOr, OpOr},
		{"BitAnd", ast.OperatorBitAnd, OpBitAnd},
		{"BitOr", ast.OperatorBitOr, OpBitOr},
		{"BitXor", ast.OperatorBitXor, OpBitXor},
		{"LShift", ast.OperatorLShift, OpLShift},
		{"RShift", ast.OperatorRShift, OpRShift},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			program := &ast.Program{
				Items: []ast.Item{
					&ast.FunctionDecl{
						Name: "main",
						Body: &ast.Block{
							Statements: []ast.Statement{
								&ast.ReturnStatement{
									Value: &ast.BinaryExpr{
										Left: &ast.Literal{
											Kind:  ast.LiteralInt,
											Value: "10",
										},
										Operator: tt.operator,
										Right: &ast.Literal{
											Kind:  ast.LiteralInt,
											Value: "20",
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
		})
	}
}

func TestGenerateUnaryExpression(t *testing.T) {
	tests := []struct {
		name     string
		operator ast.OperatorType
		expected OpCode
	}{
		{"Negate", ast.OperatorUnaryMinus, OpNeg},
		{"Not", ast.OperatorNot, OpNot},
		{"BitNot", ast.OperatorBitNot, OpBitNot},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			program := &ast.Program{
				Items: []ast.Item{
					&ast.FunctionDecl{
						Name: "main",
						Body: &ast.Block{
							Statements: []ast.Statement{
								&ast.ReturnStatement{
									Value: &ast.UnaryExpr{
										Operator: tt.operator,
										Operand: &ast.Literal{
											Kind:  ast.LiteralInt,
											Value: "42",
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
		})
	}
}

func TestGenerateLetStatement(t *testing.T) {
	// fn main() { let x = 42; return x }
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
								Value: "42",
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

	// Should have PUSH, STORE_LOCAL, LOAD_LOCAL, RETURN_VALUE, HALT
	hasStoreLocal := false
	hasLoadLocal := false
	for _, inst := range bc.Instructions {
		if inst.Op == OpStoreLocal {
			hasStoreLocal = true
		}
		if inst.Op == OpLoadLocal {
			hasLoadLocal = true
		}
	}

	if !hasStoreLocal {
		t.Error("STORE_LOCAL instruction not found")
	}
	if !hasLoadLocal {
		t.Error("LOAD_LOCAL instruction not found")
	}
}

func TestGenerateIfStatement(t *testing.T) {
	// fn main() { if true { return 1 } else { return 2 } }
	program := &ast.Program{
		Items: []ast.Item{
			&ast.FunctionDecl{
				Name: "main",
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
											Value: "2",
										},
									},
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

	// Should have JUMP_IF_FALSE and JUMP instructions for if-else
	hasJumpIfFalse := false
	hasJump := false
	for _, inst := range bc.Instructions {
		if inst.Op == OpJumpIfFalse {
			hasJumpIfFalse = true
		}
		if inst.Op == OpJump {
			hasJump = true
		}
	}

	if !hasJumpIfFalse {
		t.Error("JUMP_IF_FALSE instruction not found")
	}
	if !hasJump {
		t.Error("JUMP instruction not found (for skipping else block)")
	}
}

func TestGenerateFunctionCall(t *testing.T) {
	// fn add(a: Int, b: Int) -> Int { return a + b }
	// fn main() -> Int { return add(1, 2) }
	program := &ast.Program{
		Items: []ast.Item{
			&ast.FunctionDecl{
				Name: "add",
				Parameters: []*ast.Parameter{
					{Name: "a"},
					{Name: "b"},
				},
				Body: &ast.Block{
					Statements: []ast.Statement{
						&ast.ReturnStatement{
							Value: &ast.BinaryExpr{
								Left: &ast.Identifier{
									Name: "a",
								},
								Operator: ast.OperatorAdd,
								Right: &ast.Identifier{
									Name: "b",
								},
							},
						},
					},
				},
			},
			&ast.FunctionDecl{
				Name: "main",
				Body: &ast.Block{
					Statements: []ast.Statement{
						&ast.ReturnStatement{
							Value: &ast.CallExpr{
								Function: &ast.Identifier{
									Name: "add",
								},
								Arguments: []ast.Expression{
									&ast.Literal{
										Kind:  ast.LiteralInt,
										Value: "1",
									},
									&ast.Literal{
										Kind:  ast.LiteralInt,
										Value: "2",
									},
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

	if len(bc.Functions) != 2 {
		t.Errorf("Expected 2 functions, got %d", len(bc.Functions))
	}

	// Should have CALL instruction
	hasCall := false
	for _, inst := range bc.Instructions {
		if inst.Op == OpCall {
			hasCall = true
			if inst.ArgCount != 2 {
				t.Errorf("Expected ArgCount 2, got %d", inst.ArgCount)
			}
		}
	}

	if !hasCall {
		t.Error("CALL instruction not found")
	}
}

func TestGenerateLiterals(t *testing.T) {
	tests := []struct {
		name         string
		literal      *ast.Literal
		expectedType ValueType
	}{
		{
			name: "Integer",
			literal: &ast.Literal{
				Kind:  ast.LiteralInt,
				Value: "42",
			},
			expectedType: ValueTypeInt,
		},
		{
			name: "Float",
			literal: &ast.Literal{
				Kind:  ast.LiteralFloat,
				Value: "3.14",
			},
			expectedType: ValueTypeFloat,
		},
		{
			name: "Bool True",
			literal: &ast.Literal{
				Kind:  ast.LiteralBool,
				Value: "true",
			},
			expectedType: ValueTypeBool,
		},
		{
			name: "Bool False",
			literal: &ast.Literal{
				Kind:  ast.LiteralBool,
				Value: "false",
			},
			expectedType: ValueTypeBool,
		},
		{
			name: "String",
			literal: &ast.Literal{
				Kind:  ast.LiteralString,
				Value: "hello",
			},
			expectedType: ValueTypeString,
		},
		{
			name: "Null",
			literal: &ast.Literal{
				Kind:  ast.LiteralNull,
				Value: "null",
			},
			expectedType: ValueTypeNull,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			program := &ast.Program{
				Items: []ast.Item{
					&ast.FunctionDecl{
						Name: "main",
						Body: &ast.Block{
							Statements: []ast.Statement{
								&ast.ReturnStatement{
									Value: tt.literal,
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

			if len(bc.Constants) < 1 {
				t.Fatal("No constants generated")
			}

			if bc.Constants[0].Type != tt.expectedType {
				t.Errorf("Expected constant type %v, got %v", tt.expectedType, bc.Constants[0].Type)
			}
		})
	}
}

func TestGenerateMultipleLocalVariables(t *testing.T) {
	// fn main() { let x = 1; let y = 2; let z = x + y; return z }
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
						},
						&ast.LetStatement{
							Name: "y",
							Value: &ast.Literal{
								Kind:  ast.LiteralInt,
								Value: "2",
							},
						},
						&ast.LetStatement{
							Name: "z",
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
						&ast.ReturnStatement{
							Value: &ast.Identifier{
								Name: "z",
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

	// Count STORE_LOCAL instructions (should be 3)
	storeCount := 0
	for _, inst := range bc.Instructions {
		if inst.Op == OpStoreLocal {
			storeCount++
		}
	}

	if storeCount != 3 {
		t.Errorf("Expected 3 STORE_LOCAL instructions, got %d", storeCount)
	}
}

func TestGeneratorErrorUndefinedVariable(t *testing.T) {
	// fn main() { return x }  // x is undefined
	program := &ast.Program{
		Items: []ast.Item{
			&ast.FunctionDecl{
				Name: "main",
				Body: &ast.Block{
					Statements: []ast.Statement{
						&ast.ReturnStatement{
							Value: &ast.Identifier{
								Name: "undefined_var",
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
		t.Error("Expected error for undefined variable")
	}
}

func TestGeneratorErrorUndefinedFunction(t *testing.T) {
	// fn main() { return foo() }  // foo is undefined
	program := &ast.Program{
		Items: []ast.Item{
			&ast.FunctionDecl{
				Name: "main",
				Body: &ast.Block{
					Statements: []ast.Statement{
						&ast.ReturnStatement{
							Value: &ast.CallExpr{
								Function: &ast.Identifier{
									Name: "undefined_func",
								},
								Arguments: []ast.Expression{},
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
		t.Error("Expected error for undefined function")
	}
}

func TestGenerateFunctionParameters(t *testing.T) {
	// fn add(a: Int, b: Int) -> Int { return a + b }
	program := &ast.Program{
		Items: []ast.Item{
			&ast.FunctionDecl{
				Name: "add",
				Parameters: []*ast.Parameter{
					{Name: "a"},
					{Name: "b"},
				},
				Body: &ast.Block{
					Statements: []ast.Statement{
						&ast.ReturnStatement{
							Value: &ast.BinaryExpr{
								Left: &ast.Identifier{
									Name: "a",
								},
								Operator: ast.OperatorAdd,
								Right: &ast.Identifier{
									Name: "b",
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

	// Parameters should be loadable as local variables
	hasLoadLocal := false
	for _, inst := range bc.Instructions {
		if inst.Op == OpLoadLocal {
			hasLoadLocal = true
		}
	}

	if !hasLoadLocal {
		t.Error("Parameters not accessible as local variables")
	}
}
