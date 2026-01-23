package semantic

import (
	"testing"
	"unified-compiler/internal/ast"
)

func TestInferLiteralTypes(t *testing.T) {
	tests := []struct {
		name         string
		literal      *ast.Literal
		expectedType string
	}{
		{
			name:         "Int literal",
			literal:      &ast.Literal{Kind: ast.LiteralInt, Value: "42"},
			expectedType: "Int",
		},
		{
			name:         "Float literal",
			literal:      &ast.Literal{Kind: ast.LiteralFloat, Value: "3.14"},
			expectedType: "Float",
		},
		{
			name:         "Bool literal",
			literal:      &ast.Literal{Kind: ast.LiteralBool, Value: "true"},
			expectedType: "Bool",
		},
		{
			name:         "String literal",
			literal:      &ast.Literal{Kind: ast.LiteralString, Value: "hello"},
			expectedType: "String",
		},
		{
			name:         "Char literal",
			literal:      &ast.Literal{Kind: ast.LiteralChar, Value: "a"},
			expectedType: "Char",
		},
		{
			name:         "Null literal",
			literal:      &ast.Literal{Kind: ast.LiteralNull, Value: "null"},
			expectedType: "Null",
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inferredType, err := InferType(tt.literal)
			if err != nil {
				t.Fatalf("Failed to infer type: %v", err)
			}
			
			typeRef, ok := inferredType.(*ast.TypeReference)
			if !ok {
				t.Fatalf("Expected TypeReference, got %T", inferredType)
			}
			
			if typeRef.Name != tt.expectedType {
				t.Errorf("Expected type %s, got %s", tt.expectedType, typeRef.Name)
			}
		})
	}
}

func TestInferArithmeticExpressionType(t *testing.T) {
	tests := []struct {
		name         string
		operator     ast.OperatorType
		expectedType string
	}{
		{"Add", ast.OperatorAdd, "Int"},
		{"Sub", ast.OperatorSub, "Int"},
		{"Mul", ast.OperatorMul, "Int"},
		{"Div", ast.OperatorDiv, "Int"},
		{"Mod", ast.OperatorMod, "Int"},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			expr := &ast.BinaryExpr{
				Left:     &ast.Literal{Kind: ast.LiteralInt, Value: "10"},
				Operator: tt.operator,
				Right:    &ast.Literal{Kind: ast.LiteralInt, Value: "20"},
			}
			
			inferredType, err := InferType(expr)
			if err != nil {
				t.Fatalf("Failed to infer type: %v", err)
			}
			
			typeRef, ok := inferredType.(*ast.TypeReference)
			if !ok {
				t.Fatalf("Expected TypeReference, got %T", inferredType)
			}
			
			if typeRef.Name != tt.expectedType {
				t.Errorf("Expected type %s, got %s", tt.expectedType, typeRef.Name)
			}
		})
	}
}

func TestInferComparisonExpressionType(t *testing.T) {
	tests := []struct {
		name     string
		operator ast.OperatorType
	}{
		{"Eq", ast.OperatorEq},
		{"Ne", ast.OperatorNe},
		{"Lt", ast.OperatorLt},
		{"Le", ast.OperatorLe},
		{"Gt", ast.OperatorGt},
		{"Ge", ast.OperatorGe},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			expr := &ast.BinaryExpr{
				Left:     &ast.Literal{Kind: ast.LiteralInt, Value: "10"},
				Operator: tt.operator,
				Right:    &ast.Literal{Kind: ast.LiteralInt, Value: "20"},
			}
			
			inferredType, err := InferType(expr)
			if err != nil {
				t.Fatalf("Failed to infer type: %v", err)
			}
			
			typeRef, ok := inferredType.(*ast.TypeReference)
			if !ok {
				t.Fatalf("Expected TypeReference, got %T", inferredType)
			}
			
			if typeRef.Name != "Bool" {
				t.Errorf("Expected type Bool, got %s", typeRef.Name)
			}
		})
	}
}

func TestInferLogicalExpressionType(t *testing.T) {
	tests := []struct {
		name     string
		operator ast.OperatorType
	}{
		{"And", ast.OperatorAnd},
		{"Or", ast.OperatorOr},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			expr := &ast.BinaryExpr{
				Left:     &ast.Literal{Kind: ast.LiteralBool, Value: "true"},
				Operator: tt.operator,
				Right:    &ast.Literal{Kind: ast.LiteralBool, Value: "false"},
			}
			
			inferredType, err := InferType(expr)
			if err != nil {
				t.Fatalf("Failed to infer type: %v", err)
			}
			
			typeRef, ok := inferredType.(*ast.TypeReference)
			if !ok {
				t.Fatalf("Expected TypeReference, got %T", inferredType)
			}
			
			if typeRef.Name != "Bool" {
				t.Errorf("Expected type Bool, got %s", typeRef.Name)
			}
		})
	}
}

func TestInferUnaryExpressionType(t *testing.T) {
	tests := []struct {
		name         string
		operator     ast.OperatorType
		operand      *ast.Literal
		expectedType string
	}{
		{
			name:         "UnaryMinus",
			operator:     ast.OperatorUnaryMinus,
			operand:      &ast.Literal{Kind: ast.LiteralInt, Value: "42"},
			expectedType: "Int",
		},
		{
			name:         "UnaryPlus",
			operator:     ast.OperatorUnaryPlus,
			operand:      &ast.Literal{Kind: ast.LiteralInt, Value: "42"},
			expectedType: "Int",
		},
		{
			name:         "Not",
			operator:     ast.OperatorNot,
			operand:      &ast.Literal{Kind: ast.LiteralBool, Value: "true"},
			expectedType: "Bool",
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			expr := &ast.UnaryExpr{
				Operator: tt.operator,
				Operand:  tt.operand,
			}
			
			inferredType, err := InferType(expr)
			if err != nil {
				t.Fatalf("Failed to infer type: %v", err)
			}
			
			typeRef, ok := inferredType.(*ast.TypeReference)
			if !ok {
				t.Fatalf("Expected TypeReference, got %T", inferredType)
			}
			
			if typeRef.Name != tt.expectedType {
				t.Errorf("Expected type %s, got %s", tt.expectedType, typeRef.Name)
			}
		})
	}
}

func TestTypesCompatible(t *testing.T) {
	tests := []struct {
		name       string
		target     ast.Type
		source     ast.Type
		compatible bool
	}{
		{
			name:       "Int to Int",
			target:     &ast.TypeReference{Name: "Int"},
			source:     &ast.TypeReference{Name: "Int"},
			compatible: true,
		},
		{
			name:       "Int to Float",
			target:     &ast.TypeReference{Name: "Int"},
			source:     &ast.TypeReference{Name: "Float"},
			compatible: false,
		},
		{
			name:       "Nil source",
			target:     &ast.TypeReference{Name: "Int"},
			source:     nil,
			compatible: true,
		},
		{
			name:       "Nil target",
			target:     nil,
			source:     &ast.TypeReference{Name: "Int"},
			compatible: true,
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := TypesCompatible(tt.target, tt.source)
			if result != tt.compatible {
				t.Errorf("Expected compatibility %v, got %v", tt.compatible, result)
			}
		})
	}
}
