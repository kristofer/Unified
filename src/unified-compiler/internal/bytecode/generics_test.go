package bytecode

import (
	"testing"
	"unified-compiler/internal/ast"
)

// Test 31: Monomorphize simple generic identity function
func TestMonomorphizeIdentityFunction(t *testing.T) {
	gen := NewGenerator()
	
	// Create generic function: fn identity<T>(x: T) -> T { return x }
	identityFn := &ast.FunctionDecl{
		Name: "identity",
		GenericParams: []*ast.GenericParam{
			{Name: "T"},
		},
		Parameters: []*ast.Parameter{
			{Name: "x", Type: &ast.TypeReference{Name: "T"}},
		},
		ReturnType: &ast.TypeReference{Name: "T"},
		Body: &ast.Block{
			Statements: []ast.Statement{},
			Expression: &ast.Identifier{Name: "x"},
		},
	}
	
	// Store generic function
	gen.genericFunctions["identity"] = identityFn
	
	// Monomorphize with Int
	typeArgs := []ast.Type{
		&ast.TypeReference{Name: "Int"},
	}
	
	mangledName, err := gen.monomorphizeFunction("identity", typeArgs)
	if err != nil {
		t.Fatalf("Monomorphization failed: %v", err)
	}
	
	expectedName := "identity_Int"
	if mangledName != expectedName {
		t.Errorf("Expected mangled name %s, got %s", expectedName, mangledName)
	}
	
	// Check that the function was registered
	if _, ok := gen.bytecode.Functions[mangledName]; !ok {
		t.Errorf("Monomorphized function %s not registered", mangledName)
	}
}

// Test 32: Monomorphize generic function with multiple type parameters
func TestMonomorphizePairFunction(t *testing.T) {
	gen := NewGenerator()
	
	// Create generic function: fn pair<A, B>(a: A, b: B) -> A
	// Simplified to avoid TupleExpr which isn't implemented yet
	pairFn := &ast.FunctionDecl{
		Name: "pair",
		GenericParams: []*ast.GenericParam{
			{Name: "A"},
			{Name: "B"},
		},
		Parameters: []*ast.Parameter{
			{Name: "a", Type: &ast.TypeReference{Name: "A"}},
			{Name: "b", Type: &ast.TypeReference{Name: "B"}},
		},
		ReturnType: &ast.TypeReference{Name: "A"},
		Body: &ast.Block{
			Statements: []ast.Statement{},
			Expression: &ast.Identifier{Name: "a"},
		},
	}
	
	gen.genericFunctions["pair"] = pairFn
	
	// Monomorphize with Int and String
	typeArgs := []ast.Type{
		&ast.TypeReference{Name: "Int"},
		&ast.TypeReference{Name: "String"},
	}
	
	mangledName, err := gen.monomorphizeFunction("pair", typeArgs)
	if err != nil {
		t.Fatalf("Monomorphization failed: %v", err)
	}
	
	expectedName := "pair_Int_String"
	if mangledName != expectedName {
		t.Errorf("Expected mangled name %s, got %s", expectedName, mangledName)
	}
}

// Test 33: Generate code for generic function call with type inference
func TestGenerateGenericCallWithInference(t *testing.T) {
	gen := NewGenerator()
	
	// Create generic function
	identityFn := &ast.FunctionDecl{
		Name: "identity",
		GenericParams: []*ast.GenericParam{
			{Name: "T"},
		},
		Parameters: []*ast.Parameter{
			{Name: "x", Type: &ast.TypeReference{Name: "T"}},
		},
		ReturnType: &ast.TypeReference{Name: "T"},
		Body: &ast.Block{
			Statements: []ast.Statement{},
			Expression: &ast.Identifier{Name: "x"},
		},
	}
	
	gen.genericFunctions["identity"] = identityFn
	
	// Create a call expression: identity(42)
	callExpr := &ast.CallExpr{
		Function: &ast.Identifier{Name: "identity"},
		Arguments: []ast.Expression{
			&ast.Literal{
				Kind:  ast.LiteralInt,
				Value: "42",
			},
		},
	}
	
	// Generate the call - should infer T = Int
	err := gen.generateCallExpr(callExpr)
	if err != nil {
		t.Fatalf("Failed to generate generic call: %v", err)
	}
	
	// Verify monomorphized function was created
	mangledName := "identity_Int"
	if _, ok := gen.bytecode.Functions[mangledName]; !ok {
		t.Errorf("Monomorphized function %s not created", mangledName)
	}
}

// Test 34: Type inference from literal arguments
func TestInferExpressionTypeLiterals(t *testing.T) {
	gen := NewGenerator()
	
	tests := []struct {
		name     string
		expr     ast.Expression
		expected string
	}{
		{
			name:     "Int literal",
			expr:     &ast.Literal{Kind: ast.LiteralInt, Value: "42"},
			expected: "Int",
		},
		{
			name:     "Float literal",
			expr:     &ast.Literal{Kind: ast.LiteralFloat, Value: "3.14"},
			expected: "Float",
		},
		{
			name:     "Bool literal",
			expr:     &ast.Literal{Kind: ast.LiteralBool, Value: "true"},
			expected: "Bool",
		},
		{
			name:     "String literal",
			expr:     &ast.Literal{Kind: ast.LiteralString, Value: "hello"},
			expected: "String",
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inferredType, err := gen.inferExpressionType(tt.expr)
			if err != nil {
				t.Fatalf("Type inference failed: %v", err)
			}
			
			typeRef, ok := inferredType.(*ast.TypeReference)
			if !ok {
				t.Fatalf("Expected TypeReference, got %T", inferredType)
			}
			
			if typeRef.Name != tt.expected {
				t.Errorf("Expected type %s, got %s", tt.expected, typeRef.Name)
			}
		})
	}
}

// Test 35: Multiple monomorphizations of same function
func TestMultipleMonomorphizations(t *testing.T) {
	gen := NewGenerator()
	
	// Create generic function
	identityFn := &ast.FunctionDecl{
		Name: "identity",
		GenericParams: []*ast.GenericParam{
			{Name: "T"},
		},
		Parameters: []*ast.Parameter{
			{Name: "x", Type: &ast.TypeReference{Name: "T"}},
		},
		ReturnType: &ast.TypeReference{Name: "T"},
		Body: &ast.Block{
			Statements: []ast.Statement{},
			Expression: &ast.Identifier{Name: "x"},
		},
	}
	
	gen.genericFunctions["identity"] = identityFn
	
	// Monomorphize with Int
	name1, err := gen.monomorphizeFunction("identity", []ast.Type{
		&ast.TypeReference{Name: "Int"},
	})
	if err != nil {
		t.Fatalf("First monomorphization failed: %v", err)
	}
	
	// Monomorphize with String
	name2, err := gen.monomorphizeFunction("identity", []ast.Type{
		&ast.TypeReference{Name: "String"},
	})
	if err != nil {
		t.Fatalf("Second monomorphization failed: %v", err)
	}
	
	// Verify different names
	if name1 == name2 {
		t.Errorf("Expected different mangled names, both are %s", name1)
	}
	
	// Verify both functions are registered
	if _, ok := gen.bytecode.Functions[name1]; !ok {
		t.Errorf("First monomorphization %s not registered", name1)
	}
	if _, ok := gen.bytecode.Functions[name2]; !ok {
		t.Errorf("Second monomorphization %s not registered", name2)
	}
}

// Test 36: Idempotent monomorphization (calling twice with same types)
func TestIdempotentMonomorphization(t *testing.T) {
	gen := NewGenerator()
	
	identityFn := &ast.FunctionDecl{
		Name: "identity",
		GenericParams: []*ast.GenericParam{
			{Name: "T"},
		},
		Parameters: []*ast.Parameter{
			{Name: "x", Type: &ast.TypeReference{Name: "T"}},
		},
		ReturnType: &ast.TypeReference{Name: "T"},
		Body: &ast.Block{
			Statements: []ast.Statement{},
			Expression: &ast.Identifier{Name: "x"},
		},
	}
	
	gen.genericFunctions["identity"] = identityFn
	
	typeArgs := []ast.Type{&ast.TypeReference{Name: "Int"}}
	
	// First call
	name1, err := gen.monomorphizeFunction("identity", typeArgs)
	if err != nil {
		t.Fatalf("First call failed: %v", err)
	}
	
	// Second call with same type args
	name2, err := gen.monomorphizeFunction("identity", typeArgs)
	if err != nil {
		t.Fatalf("Second call failed: %v", err)
	}
	
	// Should return same name
	if name1 != name2 {
		t.Errorf("Expected same name from both calls: %s vs %s", name1, name2)
	}
}

// Test 37: Type argument count mismatch error
func TestTypeArgumentCountMismatch(t *testing.T) {
	gen := NewGenerator()
	
	identityFn := &ast.FunctionDecl{
		Name: "identity",
		GenericParams: []*ast.GenericParam{
			{Name: "T"},
		},
		Parameters: []*ast.Parameter{
			{Name: "x", Type: &ast.TypeReference{Name: "T"}},
		},
		ReturnType: &ast.TypeReference{Name: "T"},
		Body: &ast.Block{},
	}
	
	gen.genericFunctions["identity"] = identityFn
	
	// Try to monomorphize with wrong number of type args
	_, err := gen.monomorphizeFunction("identity", []ast.Type{
		&ast.TypeReference{Name: "Int"},
		&ast.TypeReference{Name: "String"},
	})
	
	if err == nil {
		t.Error("Expected error for type argument count mismatch")
	}
}

// Test 38: Parameter type substitution
func TestParameterTypeSubstitution(t *testing.T) {
	gen := NewGenerator()
	
	// Create context with T -> Int substitution
	ctx := gen.genericContext
	ctx.AddTypeParameter(&ast.GenericParam{Name: "T"})
	ctx.Substitutions["T"] = &ast.TypeReference{Name: "Int"}
	
	// Create parameter with type T
	params := []*ast.Parameter{
		{
			Name: "x",
			Type: &ast.TypeReference{Name: "T"},
		},
	}
	
	// Substitute
	result := gen.substituteParameters(params, ctx)
	
	if len(result) != 1 {
		t.Fatalf("Expected 1 parameter, got %d", len(result))
	}
	
	paramType, ok := result[0].Type.(*ast.TypeReference)
	if !ok {
		t.Fatalf("Expected TypeReference, got %T", result[0].Type)
	}
	
	if paramType.Name != "Int" {
		t.Errorf("Expected parameter type Int, got %s", paramType.Name)
	}
}

// Test 39: Generic function with nested type
func TestMonomorphizeNestedGenericType(t *testing.T) {
	gen := NewGenerator()
	
	// fn process<T>(x: Box<T>) -> Box<T>
	processFn := &ast.FunctionDecl{
		Name: "process",
		GenericParams: []*ast.GenericParam{
			{Name: "T"},
		},
		Parameters: []*ast.Parameter{
			{
				Name: "x",
				Type: &ast.TypeReference{
					Name: "Box",
					TypeArgs: []ast.Type{
						&ast.TypeReference{Name: "T"},
					},
				},
			},
		},
		ReturnType: &ast.TypeReference{
			Name: "Box",
			TypeArgs: []ast.Type{
				&ast.TypeReference{Name: "T"},
			},
		},
		Body: &ast.Block{
			Expression: &ast.Identifier{Name: "x"},
		},
	}
	
	gen.genericFunctions["process"] = processFn
	
	// Monomorphize with Int
	mangledName, err := gen.monomorphizeFunction("process", []ast.Type{
		&ast.TypeReference{Name: "Int"},
	})
	if err != nil {
		t.Fatalf("Monomorphization failed: %v", err)
	}
	
	if mangledName != "process_Int" {
		t.Errorf("Expected mangled name process_Int, got %s", mangledName)
	}
}

// Test 40: Infer type arguments from multiple parameters
func TestInferTypeArgsFromMultipleParams(t *testing.T) {
	gen := NewGenerator()
	
	genericParams := []*ast.GenericParam{
		{Name: "T"},
	}
	
	paramTypes := []ast.Type{
		&ast.TypeReference{Name: "T"},
		&ast.TypeReference{Name: "T"},
	}
	
	argTypes := []ast.Type{
		&ast.TypeReference{Name: "Int"},
		&ast.TypeReference{Name: "Int"},
	}
	
	inferred, err := gen.inferTypeArgsFromCall(genericParams, paramTypes, argTypes)
	if err != nil {
		t.Fatalf("Type inference failed: %v", err)
	}
	
	if len(inferred) != 1 {
		t.Fatalf("Expected 1 inferred type, got %d", len(inferred))
	}
	
	inferredType, ok := inferred[0].(*ast.TypeReference)
	if !ok {
		t.Fatalf("Expected TypeReference, got %T", inferred[0])
	}
	
	if inferredType.Name != "Int" {
		t.Errorf("Expected inferred type Int, got %s", inferredType.Name)
	}
}
