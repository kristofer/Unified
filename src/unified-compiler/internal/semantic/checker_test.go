package semantic

import (
	"testing"
	"unified-compiler/internal/ast"
)

func TestCheckLetStatement(t *testing.T) {
	// fn main() { let x = 42; }
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
							IsMutable: false,
						},
					},
				},
			},
		},
	}
	
	checker := NewChecker()
	err := checker.Check(program)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}

func TestCheckMutableVariable(t *testing.T) {
	// fn main() { let mut x = 42; }
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
							IsMutable: true,
						},
					},
				},
			},
		},
	}
	
	checker := NewChecker()
	err := checker.Check(program)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}

func TestCheckAssignmentToMutableVariable(t *testing.T) {
	// fn main() { let mut x = 42; x = 10; }
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
					},
				},
			},
		},
	}
	
	checker := NewChecker()
	err := checker.Check(program)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}

func TestCheckAssignmentToImmutableVariable(t *testing.T) {
	// fn main() { let x = 42; x = 10; }  // Should error
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
							IsMutable: false,
						},
						&ast.AssignmentStatement{
							Target:   "x",
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
	
	checker := NewChecker()
	err := checker.Check(program)
	if err == nil {
		t.Error("Expected error for assignment to immutable variable")
	}
	
	semanticErr, ok := err.(*SemanticError)
	if !ok {
		t.Errorf("Expected SemanticError, got %T", err)
	} else if semanticErr.Message != "cannot assign to immutable variable 'x'" {
		t.Errorf("Unexpected error message: %s", semanticErr.Message)
	}
}

func TestCheckAssignmentToUndefinedVariable(t *testing.T) {
	// fn main() { x = 10; }  // x is undefined
	program := &ast.Program{
		Items: []ast.Item{
			&ast.FunctionDecl{
				Name: "main",
				Body: &ast.Block{
					Statements: []ast.Statement{
						&ast.AssignmentStatement{
							Target:   "x",
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
	
	checker := NewChecker()
	err := checker.Check(program)
	if err == nil {
		t.Error("Expected error for assignment to undefined variable")
	}
}

func TestCheckUndefinedVariableReference(t *testing.T) {
	// fn main() { return x; }  // x is undefined
	program := &ast.Program{
		Items: []ast.Item{
			&ast.FunctionDecl{
				Name: "main",
				Body: &ast.Block{
					Statements: []ast.Statement{
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
	
	checker := NewChecker()
	err := checker.Check(program)
	if err == nil {
		t.Error("Expected error for undefined variable reference")
	}
	
	semanticErr, ok := err.(*SemanticError)
	if !ok {
		t.Errorf("Expected SemanticError, got %T", err)
	} else if semanticErr.Message != "undefined variable 'x'" {
		t.Errorf("Unexpected error message: %s", semanticErr.Message)
	}
}

func TestCheckVariableShadowing(t *testing.T) {
	// fn main() { let x = 5; { let x = 10; } }
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
							IsMutable: false,
						},
						&ast.BlockStatement{
							Block: &ast.Block{
								Statements: []ast.Statement{
									&ast.LetStatement{
										Name: "x",
										Value: &ast.Literal{
											Kind:  ast.LiteralInt,
											Value: "10",
										},
										IsMutable: false,
									},
								},
							},
						},
					},
				},
			},
		},
	}
	
	checker := NewChecker()
	err := checker.Check(program)
	if err != nil {
		t.Errorf("Variable shadowing should be allowed: %v", err)
	}
}

func TestCheckDuplicateVariable(t *testing.T) {
	// fn main() { let x = 5; let x = 10; }  // Should error
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
							IsMutable: false,
						},
						&ast.LetStatement{
							Name: "x",
							Value: &ast.Literal{
								Kind:  ast.LiteralInt,
								Value: "10",
							},
							IsMutable: false,
						},
					},
				},
			},
		},
	}
	
	checker := NewChecker()
	err := checker.Check(program)
	if err == nil {
		t.Error("Expected error for duplicate variable in same scope")
	}
}

func TestCheckTypeInference(t *testing.T) {
	// fn main() { let x = 42; let y = x + 10; }
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
							IsMutable: false,
						},
						&ast.LetStatement{
							Name: "y",
							Value: &ast.BinaryExpr{
								Left: &ast.Identifier{
									Name: "x",
								},
								Operator: ast.OperatorAdd,
								Right: &ast.Literal{
									Kind:  ast.LiteralInt,
									Value: "10",
								},
							},
							IsMutable: false,
						},
					},
				},
			},
		},
	}
	
	checker := NewChecker()
	err := checker.Check(program)
	if err != nil {
		t.Errorf("Type inference should work: %v", err)
	}
}
