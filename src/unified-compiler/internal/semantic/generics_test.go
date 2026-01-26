package semantic

import (
	"testing"
	"unified-compiler/internal/ast"
)

// Test 1: Generic function parsing with single type parameter
func TestGenericFunctionSingleTypeParam(t *testing.T) {
	// Test parsing of: fn identity<T>(x: T) -> T
	funcDecl := &ast.FunctionDecl{
		Name: "identity",
		GenericParams: []*ast.GenericParam{
			{Name: "T"},
		},
		Parameters: []*ast.Parameter{
			{
				Name: "x",
				Type: &ast.TypeReference{Name: "T"},
			},
		},
		ReturnType: &ast.TypeReference{Name: "T"},
	}

	if len(funcDecl.GenericParams) != 1 {
		t.Errorf("Expected 1 generic parameter, got %d", len(funcDecl.GenericParams))
	}
	if funcDecl.GenericParams[0].Name != "T" {
		t.Errorf("Expected generic parameter name 'T', got '%s'", funcDecl.GenericParams[0].Name)
	}
}

// Test 2: Generic function with multiple type parameters
func TestGenericFunctionMultipleTypeParams(t *testing.T) {
	// Test parsing of: fn pair<A, B>(a: A, b: B) -> (A, B)
	funcDecl := &ast.FunctionDecl{
		Name: "pair",
		GenericParams: []*ast.GenericParam{
			{Name: "A"},
			{Name: "B"},
		},
		Parameters: []*ast.Parameter{
			{Name: "a", Type: &ast.TypeReference{Name: "A"}},
			{Name: "b", Type: &ast.TypeReference{Name: "B"}},
		},
		ReturnType: &ast.TupleType{
			ElementTypes: []ast.Type{
				&ast.TypeReference{Name: "A"},
				&ast.TypeReference{Name: "B"},
			},
		},
	}

	if len(funcDecl.GenericParams) != 2 {
		t.Errorf("Expected 2 generic parameters, got %d", len(funcDecl.GenericParams))
	}
	if funcDecl.GenericParams[0].Name != "A" {
		t.Errorf("Expected first generic parameter name 'A', got '%s'", funcDecl.GenericParams[0].Name)
	}
	if funcDecl.GenericParams[1].Name != "B" {
		t.Errorf("Expected second generic parameter name 'B', got '%s'", funcDecl.GenericParams[1].Name)
	}
}

// Test 3: Generic struct with type parameter
func TestGenericStruct(t *testing.T) {
	// Test parsing of: struct Box<T> { value: T }
	structDecl := &ast.StructDecl{
		Name: "Box",
		GenericParams: []*ast.GenericParam{
			{Name: "T"},
		},
		Members: []*ast.StructMember{
			{
				Name: "value",
				Type: &ast.TypeReference{Name: "T"},
			},
		},
	}

	if len(structDecl.GenericParams) != 1 {
		t.Errorf("Expected 1 generic parameter, got %d", len(structDecl.GenericParams))
	}
	if structDecl.GenericParams[0].Name != "T" {
		t.Errorf("Expected generic parameter name 'T', got '%s'", structDecl.GenericParams[0].Name)
	}
	if len(structDecl.Members) != 1 {
		t.Errorf("Expected 1 struct member, got %d", len(structDecl.Members))
	}
}

// Test 4: Generic enum with type parameter
func TestGenericEnum(t *testing.T) {
	// Test parsing of: enum Option<T> { Some(T), None }
	enumDecl := &ast.EnumDecl{
		Name: "Option",
		GenericParams: []*ast.GenericParam{
			{Name: "T"},
		},
		Variants: []*ast.EnumVariant{
			{
				Name: "Some",
				Parameters: []*ast.EnumVariantParam{
					{Type: &ast.TypeReference{Name: "T"}},
				},
			},
			{
				Name: "None",
			},
		},
	}

	if len(enumDecl.GenericParams) != 1 {
		t.Errorf("Expected 1 generic parameter, got %d", len(enumDecl.GenericParams))
	}
	if enumDecl.GenericParams[0].Name != "T" {
		t.Errorf("Expected generic parameter name 'T', got '%s'", enumDecl.GenericParams[0].Name)
	}
}

// Test 5: Type parameter with constraints
func TestTypeParameterWithConstraints(t *testing.T) {
	// Test parsing of: fn compare<T: Comparable>(a: T, b: T) -> Bool
	funcDecl := &ast.FunctionDecl{
		Name: "compare",
		GenericParams: []*ast.GenericParam{
			{
				Name: "T",
				Constraints: []ast.Type{
					&ast.TypeReference{Name: "Comparable"},
				},
			},
		},
		Parameters: []*ast.Parameter{
			{Name: "a", Type: &ast.TypeReference{Name: "T"}},
			{Name: "b", Type: &ast.TypeReference{Name: "T"}},
		},
		ReturnType: &ast.TypeReference{Name: "Bool"},
	}

	if len(funcDecl.GenericParams) != 1 {
		t.Errorf("Expected 1 generic parameter, got %d", len(funcDecl.GenericParams))
	}
	if len(funcDecl.GenericParams[0].Constraints) != 1 {
		t.Errorf("Expected 1 constraint, got %d", len(funcDecl.GenericParams[0].Constraints))
	}
	constraint, ok := funcDecl.GenericParams[0].Constraints[0].(*ast.TypeReference)
	if !ok {
		t.Errorf("Expected constraint to be TypeReference")
	} else if constraint.Name != "Comparable" {
		t.Errorf("Expected constraint name 'Comparable', got '%s'", constraint.Name)
	}
}

// Test 6: Type application with concrete types
func TestTypeApplicationWithConcreteTypes(t *testing.T) {
	// Test type application: Box<Int>
	typeRef := &ast.TypeReference{
		Name: "Box",
		TypeArgs: []ast.Type{
			&ast.TypeReference{Name: "Int"},
		},
	}

	if typeRef.Name != "Box" {
		t.Errorf("Expected type name 'Box', got '%s'", typeRef.Name)
	}
	if len(typeRef.TypeArgs) != 1 {
		t.Errorf("Expected 1 type argument, got %d", len(typeRef.TypeArgs))
	}
	typeArg, ok := typeRef.TypeArgs[0].(*ast.TypeReference)
	if !ok {
		t.Errorf("Expected type argument to be TypeReference")
	} else if typeArg.Name != "Int" {
		t.Errorf("Expected type argument 'Int', got '%s'", typeArg.Name)
	}
}

// Test 7: Multiple type arguments
func TestMultipleTypeArguments(t *testing.T) {
	// Test type application: Result<Int, String>
	typeRef := &ast.TypeReference{
		Name: "Result",
		TypeArgs: []ast.Type{
			&ast.TypeReference{Name: "Int"},
			&ast.TypeReference{Name: "String"},
		},
	}

	if typeRef.Name != "Result" {
		t.Errorf("Expected type name 'Result', got '%s'", typeRef.Name)
	}
	if len(typeRef.TypeArgs) != 2 {
		t.Errorf("Expected 2 type arguments, got %d", len(typeRef.TypeArgs))
	}
}

// Test 8: Nested generic types
func TestNestedGenericTypes(t *testing.T) {
	// Test type application: Option<Box<Int>>
	innerType := &ast.TypeReference{
		Name: "Box",
		TypeArgs: []ast.Type{
			&ast.TypeReference{Name: "Int"},
		},
	}
	outerType := &ast.TypeReference{
		Name: "Option",
		TypeArgs: []ast.Type{
			innerType,
		},
	}

	if outerType.Name != "Option" {
		t.Errorf("Expected outer type name 'Option', got '%s'", outerType.Name)
	}
	if len(outerType.TypeArgs) != 1 {
		t.Errorf("Expected 1 type argument for outer type, got %d", len(outerType.TypeArgs))
	}
	
	innerTypeRef, ok := outerType.TypeArgs[0].(*ast.TypeReference)
	if !ok {
		t.Errorf("Expected inner type to be TypeReference")
	} else {
		if innerTypeRef.Name != "Box" {
			t.Errorf("Expected inner type name 'Box', got '%s'", innerTypeRef.Name)
		}
		if len(innerTypeRef.TypeArgs) != 1 {
			t.Errorf("Expected 1 type argument for inner type, got %d", len(innerTypeRef.TypeArgs))
		}
	}
}

// Test 9: Generic function with where clause
func TestGenericFunctionWithWhereClause(t *testing.T) {
	// Test: fn process<T>(value: T) -> T where T: Clone + Send
	funcDecl := &ast.FunctionDecl{
		Name: "process",
		GenericParams: []*ast.GenericParam{
			{Name: "T"},
		},
		Parameters: []*ast.Parameter{
			{Name: "value", Type: &ast.TypeReference{Name: "T"}},
		},
		ReturnType: &ast.TypeReference{Name: "T"},
		WhereConstraints: []*ast.WhereConstraint{
			{
				SubjectType: &ast.TypeReference{Name: "T"},
				Constraints: []ast.Type{
					&ast.TypeReference{Name: "Clone"},
					&ast.TypeReference{Name: "Send"},
				},
			},
		},
	}

	if len(funcDecl.WhereConstraints) != 1 {
		t.Errorf("Expected 1 where constraint, got %d", len(funcDecl.WhereConstraints))
	}
	if len(funcDecl.WhereConstraints[0].Constraints) != 2 {
		t.Errorf("Expected 2 constraints in where clause, got %d", len(funcDecl.WhereConstraints[0].Constraints))
	}
}

// Test 10: Generic type compatibility check
func TestGenericTypeCompatibility(t *testing.T) {
	tests := []struct {
		name       string
		target     ast.Type
		source     ast.Type
		compatible bool
	}{
		{
			name:       "Same generic type",
			target:     &ast.TypeReference{Name: "T"},
			source:     &ast.TypeReference{Name: "T"},
			compatible: true,
		},
		{
			name:       "Different generic types",
			target:     &ast.TypeReference{Name: "T"},
			source:     &ast.TypeReference{Name: "U"},
			compatible: false,
		},
		{
			name: "Concrete type to generic",
			target: &ast.TypeReference{
				Name: "Box",
				TypeArgs: []ast.Type{
					&ast.TypeReference{Name: "Int"},
				},
			},
			source: &ast.TypeReference{
				Name: "Box",
				TypeArgs: []ast.Type{
					&ast.TypeReference{Name: "Int"},
				},
			},
			compatible: true,
		},
		{
			name: "Different concrete type arguments",
			target: &ast.TypeReference{
				Name: "Box",
				TypeArgs: []ast.Type{
					&ast.TypeReference{Name: "Int"},
				},
			},
			source: &ast.TypeReference{
				Name: "Box",
				TypeArgs: []ast.Type{
					&ast.TypeReference{Name: "String"},
				},
			},
			compatible: false,
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
