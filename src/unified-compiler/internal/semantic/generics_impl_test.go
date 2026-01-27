package semantic

import (
	"testing"
	"unified-compiler/internal/ast"
)

// Test 11: Generic context creation and type parameter addition
func TestGenericContextCreation(t *testing.T) {
	gc := NewGenericContext()
	
	param := &ast.GenericParam{Name: "T"}
	tp := gc.AddTypeParameter(param)
	
	if tp.Name != "T" {
		t.Errorf("Expected type parameter name 'T', got '%s'", tp.Name)
	}
	
	// Lookup should find it
	found := gc.LookupTypeParam("T")
	if found == nil {
		t.Error("Expected to find type parameter T")
	}
	if found.Name != "T" {
		t.Errorf("Expected found parameter name 'T', got '%s'", found.Name)
	}
}

// Test 12: Type substitution basic case
func TestTypeSubstitutionBasic(t *testing.T) {
	gc := NewGenericContext()
	gc.AddTypeParameter(&ast.GenericParam{Name: "T"})
	
	// Substitute T -> Int
	gc.Substitutions["T"] = &ast.TypeReference{Name: "Int"}
	
	// Apply substitution
	input := &ast.TypeReference{Name: "T"}
	result := gc.Substitute(input)
	
	resultRef, ok := result.(*ast.TypeReference)
	if !ok {
		t.Fatal("Expected TypeReference result")
	}
	if resultRef.Name != "Int" {
		t.Errorf("Expected substitution to Int, got %s", resultRef.Name)
	}
}

// Test 13: Type substitution with generic types
func TestTypeSubstitutionGeneric(t *testing.T) {
	gc := NewGenericContext()
	gc.AddTypeParameter(&ast.GenericParam{Name: "T"})
	
	// Substitute T -> String
	gc.Substitutions["T"] = &ast.TypeReference{Name: "String"}
	
	// Apply to Box<T>
	input := &ast.TypeReference{
		Name: "Box",
		TypeArgs: []ast.Type{
			&ast.TypeReference{Name: "T"},
		},
	}
	
	result := gc.Substitute(input)
	resultRef, ok := result.(*ast.TypeReference)
	if !ok {
		t.Fatal("Expected TypeReference result")
	}
	
	if resultRef.Name != "Box" {
		t.Errorf("Expected Box, got %s", resultRef.Name)
	}
	if len(resultRef.TypeArgs) != 1 {
		t.Fatalf("Expected 1 type argument, got %d", len(resultRef.TypeArgs))
	}
	
	argRef, ok := resultRef.TypeArgs[0].(*ast.TypeReference)
	if !ok {
		t.Fatal("Expected TypeReference type argument")
	}
	if argRef.Name != "String" {
		t.Errorf("Expected String type argument, got %s", argRef.Name)
	}
}

// Test 14: Type unification with same types
func TestUnificationSameTypes(t *testing.T) {
	gc := NewGenericContext()
	
	t1 := &ast.TypeReference{Name: "Int"}
	t2 := &ast.TypeReference{Name: "Int"}
	
	err := gc.Unify(t1, t2)
	if err != nil {
		t.Errorf("Expected unification to succeed, got error: %v", err)
	}
}

// Test 15: Type unification with type parameter
func TestUnificationTypeParameter(t *testing.T) {
	gc := NewGenericContext()
	gc.AddTypeParameter(&ast.GenericParam{Name: "T"})
	
	t1 := &ast.TypeReference{Name: "T"}
	t2 := &ast.TypeReference{Name: "Int"}
	
	err := gc.Unify(t1, t2)
	if err != nil {
		t.Errorf("Expected unification to succeed, got error: %v", err)
	}
	
	// Check that T is now bound to Int
	if sub, ok := gc.Substitutions["T"]; ok {
		subRef, ok := sub.(*ast.TypeReference)
		if !ok {
			t.Fatal("Expected TypeReference substitution")
		}
		if subRef.Name != "Int" {
			t.Errorf("Expected T to be bound to Int, got %s", subRef.Name)
		}
	} else {
		t.Error("Expected T to have a substitution")
	}
}

// Test 16: Type unification failure with incompatible types
func TestUnificationIncompatibleTypes(t *testing.T) {
	gc := NewGenericContext()
	
	t1 := &ast.TypeReference{Name: "Int"}
	t2 := &ast.TypeReference{Name: "String"}
	
	err := gc.Unify(t1, t2)
	if err == nil {
		t.Error("Expected unification to fail for incompatible types")
	}
}

// Test 17: Type inference for identity function
func TestTypeInferenceIdentity(t *testing.T) {
	gc := NewGenericContext()
	
	// fn identity<T>(x: T) -> T
	genericParams := []*ast.GenericParam{
		{Name: "T"},
	}
	
	paramTypes := []ast.Type{
		&ast.TypeReference{Name: "T"},
	}
	
	// Call with Int
	argTypes := []ast.Type{
		&ast.TypeReference{Name: "Int"},
	}
	
	inferred, err := gc.InferTypeArguments(genericParams, paramTypes, argTypes)
	if err != nil {
		t.Fatalf("Type inference failed: %v", err)
	}
	
	if len(inferred) != 1 {
		t.Fatalf("Expected 1 inferred type, got %d", len(inferred))
	}
	
	inferredT, ok := inferred["T"]
	if !ok {
		t.Fatal("Expected to infer type for T")
	}
	
	inferredRef, ok := inferredT.(*ast.TypeReference)
	if !ok {
		t.Fatal("Expected TypeReference for inferred type")
	}
	if inferredRef.Name != "Int" {
		t.Errorf("Expected inferred type Int, got %s", inferredRef.Name)
	}
}

// Test 18: Type inference for pair function
func TestTypeInferencePair(t *testing.T) {
	gc := NewGenericContext()
	
	// fn pair<A, B>(a: A, b: B) -> (A, B)
	genericParams := []*ast.GenericParam{
		{Name: "A"},
		{Name: "B"},
	}
	
	paramTypes := []ast.Type{
		&ast.TypeReference{Name: "A"},
		&ast.TypeReference{Name: "B"},
	}
	
	// Call with Int and String
	argTypes := []ast.Type{
		&ast.TypeReference{Name: "Int"},
		&ast.TypeReference{Name: "String"},
	}
	
	inferred, err := gc.InferTypeArguments(genericParams, paramTypes, argTypes)
	if err != nil {
		t.Fatalf("Type inference failed: %v", err)
	}
	
	if len(inferred) != 2 {
		t.Fatalf("Expected 2 inferred types, got %d", len(inferred))
	}
	
	inferredA, ok := inferred["A"].(*ast.TypeReference)
	if !ok || inferredA.Name != "Int" {
		t.Errorf("Expected inferred type A=Int")
	}
	
	inferredB, ok := inferred["B"].(*ast.TypeReference)
	if !ok || inferredB.Name != "String" {
		t.Errorf("Expected inferred type B=String")
	}
}

// Test 19: Name mangling for monomorphization
func TestNameMangling(t *testing.T) {
	tests := []struct {
		name     string
		baseName string
		typeArgs []ast.Type
		expected string
	}{
		{
			name:     "No type arguments",
			baseName: "identity",
			typeArgs: []ast.Type{},
			expected: "identity",
		},
		{
			name:     "Single type argument",
			baseName: "identity",
			typeArgs: []ast.Type{
				&ast.TypeReference{Name: "Int"},
			},
			expected: "identity_Int",
		},
		{
			name:     "Multiple type arguments",
			baseName: "pair",
			typeArgs: []ast.Type{
				&ast.TypeReference{Name: "Int"},
				&ast.TypeReference{Name: "String"},
			},
			expected: "pair_Int_String",
		},
		{
			name:     "Generic type argument",
			baseName: "process",
			typeArgs: []ast.Type{
				&ast.TypeReference{
					Name: "Box",
					TypeArgs: []ast.Type{
						&ast.TypeReference{Name: "Int"},
					},
				},
			},
			expected: "process_Box_Int_",
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MangleName(tt.baseName, tt.typeArgs)
			if result != tt.expected {
				t.Errorf("Expected mangled name '%s', got '%s'", tt.expected, result)
			}
		})
	}
}

// Test 20: TypeToString conversion
func TestTypeToString(t *testing.T) {
	tests := []struct {
		name     string
		typ      ast.Type
		expected string
	}{
		{
			name:     "Simple type",
			typ:      &ast.TypeReference{Name: "Int"},
			expected: "Int",
		},
		{
			name: "Generic type",
			typ: &ast.TypeReference{
				Name: "Box",
				TypeArgs: []ast.Type{
					&ast.TypeReference{Name: "String"},
				},
			},
			expected: "Box<String>",
		},
		{
			name: "Nested generic type",
			typ: &ast.TypeReference{
				Name: "Option",
				TypeArgs: []ast.Type{
					&ast.TypeReference{
						Name: "Box",
						TypeArgs: []ast.Type{
							&ast.TypeReference{Name: "Int"},
						},
					},
				},
			},
			expected: "Option<Box<Int>>",
		},
		{
			name: "Tuple type",
			typ: &ast.TupleType{
				ElementTypes: []ast.Type{
					&ast.TypeReference{Name: "Int"},
					&ast.TypeReference{Name: "String"},
				},
			},
			expected: "(Int,String)",
		},
		{
			name: "Reference type",
			typ: &ast.ReferenceType{
				BaseType: &ast.TypeReference{Name: "Int"},
			},
			expected: "&Int",
		},
		{
			name: "Mutable reference type",
			typ: &ast.ReferenceType{
				BaseType:  &ast.TypeReference{Name: "String"},
				IsMutable: true,
			},
			expected: "&mut String",
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := TypeToString(tt.typ)
			if result != tt.expected {
				t.Errorf("Expected type string '%s', got '%s'", tt.expected, result)
			}
		})
	}
}

// Test 21: Child context inheritance
func TestChildContextInheritance(t *testing.T) {
	parent := NewGenericContext()
	parent.AddTypeParameter(&ast.GenericParam{Name: "T"})
	
	child := parent.NewChildContext()
	child.AddTypeParameter(&ast.GenericParam{Name: "U"})
	
	// Child should find both T and U
	if child.LookupTypeParam("U") == nil {
		t.Error("Child should find its own type parameter U")
	}
	if child.LookupTypeParam("T") == nil {
		t.Error("Child should find parent's type parameter T")
	}
	
	// Parent should only find T
	if parent.LookupTypeParam("T") == nil {
		t.Error("Parent should find its own type parameter T")
	}
	if parent.LookupTypeParam("U") != nil {
		t.Error("Parent should not find child's type parameter U")
	}
}

// Test 22: Type variable creation
func TestTypeVariableCreation(t *testing.T) {
	gc := NewGenericContext()
	
	tv1 := gc.NewTypeVariable()
	tv2 := gc.NewTypeVariable()
	
	if tv1.ID == tv2.ID {
		t.Error("Type variables should have unique IDs")
	}
	
	if tv1.Resolved != nil {
		t.Error("New type variable should not be resolved")
	}
}

// Test 23: Monomorphization key generation
func TestMonomorphizationKey(t *testing.T) {
	key1 := MonomorphizationKey{
		FunctionName: "identity",
		TypeArgs:     []string{"Int"},
	}
	
	key2 := MonomorphizationKey{
		FunctionName: "identity",
		TypeArgs:     []string{"String"},
	}
	
	if key1.String() == key2.String() {
		t.Error("Different monomorphizations should have different keys")
	}
	
	expected1 := "identity<Int>"
	if key1.String() != expected1 {
		t.Errorf("Expected key '%s', got '%s'", expected1, key1.String())
	}
}

// Test 24: Complex type substitution
func TestComplexTypeSubstitution(t *testing.T) {
	gc := NewGenericContext()
	gc.AddTypeParameter(&ast.GenericParam{Name: "T"})
	gc.AddTypeParameter(&ast.GenericParam{Name: "U"})
	
	// Substitute T -> Int, U -> String
	gc.Substitutions["T"] = &ast.TypeReference{Name: "Int"}
	gc.Substitutions["U"] = &ast.TypeReference{Name: "String"}
	
	// Apply to Result<T, U>
	input := &ast.TypeReference{
		Name: "Result",
		TypeArgs: []ast.Type{
			&ast.TypeReference{Name: "T"},
			&ast.TypeReference{Name: "U"},
		},
	}
	
	result := gc.Substitute(input)
	resultRef := result.(*ast.TypeReference)
	
	if resultRef.Name != "Result" {
		t.Errorf("Expected Result, got %s", resultRef.Name)
	}
	if len(resultRef.TypeArgs) != 2 {
		t.Fatalf("Expected 2 type arguments, got %d", len(resultRef.TypeArgs))
	}
	
	arg0 := resultRef.TypeArgs[0].(*ast.TypeReference)
	arg1 := resultRef.TypeArgs[1].(*ast.TypeReference)
	
	if arg0.Name != "Int" {
		t.Errorf("Expected first type arg Int, got %s", arg0.Name)
	}
	if arg1.Name != "String" {
		t.Errorf("Expected second type arg String, got %s", arg1.Name)
	}
}

// Test 25: Function type substitution
func TestFunctionTypeSubstitution(t *testing.T) {
	gc := NewGenericContext()
	gc.AddTypeParameter(&ast.GenericParam{Name: "T"})
	gc.Substitutions["T"] = &ast.TypeReference{Name: "Int"}
	
	// fn(T) -> T
	input := &ast.FunctionType{
		ParamTypes: []ast.Type{
			&ast.TypeReference{Name: "T"},
		},
		ReturnType: &ast.TypeReference{Name: "T"},
	}
	
	result := gc.Substitute(input)
	fnType := result.(*ast.FunctionType)
	
	paramType := fnType.ParamTypes[0].(*ast.TypeReference)
	returnType := fnType.ReturnType.(*ast.TypeReference)
	
	if paramType.Name != "Int" {
		t.Errorf("Expected param type Int, got %s", paramType.Name)
	}
	if returnType.Name != "Int" {
		t.Errorf("Expected return type Int, got %s", returnType.Name)
	}
}

// Test 26: Tuple type unification
func TestTupleTypeUnification(t *testing.T) {
	gc := NewGenericContext()
	
	t1 := &ast.TupleType{
		ElementTypes: []ast.Type{
			&ast.TypeReference{Name: "Int"},
			&ast.TypeReference{Name: "String"},
		},
	}
	
	t2 := &ast.TupleType{
		ElementTypes: []ast.Type{
			&ast.TypeReference{Name: "Int"},
			&ast.TypeReference{Name: "String"},
		},
	}
	
	err := gc.Unify(t1, t2)
	if err != nil {
		t.Errorf("Expected tuple unification to succeed, got: %v", err)
	}
}

// Test 27: Generic type argument unification
func TestGenericTypeArgumentUnification(t *testing.T) {
	gc := NewGenericContext()
	gc.AddTypeParameter(&ast.GenericParam{Name: "T"})
	
	// Box<T> with Box<Int>
	t1 := &ast.TypeReference{
		Name: "Box",
		TypeArgs: []ast.Type{
			&ast.TypeReference{Name: "T"},
		},
	}
	
	t2 := &ast.TypeReference{
		Name: "Box",
		TypeArgs: []ast.Type{
			&ast.TypeReference{Name: "Int"},
		},
	}
	
	err := gc.Unify(t1, t2)
	if err != nil {
		t.Errorf("Expected unification to succeed, got: %v", err)
	}
	
	// Check that T is bound to Int
	if sub, ok := gc.Substitutions["T"]; ok {
		subRef := sub.(*ast.TypeReference)
		if subRef.Name != "Int" {
			t.Errorf("Expected T bound to Int, got %s", subRef.Name)
		}
	} else {
		t.Error("Expected T to be bound")
	}
}

// Test 28: Inference failure with wrong argument count
func TestInferenceArgumentCountMismatch(t *testing.T) {
	gc := NewGenericContext()
	
	genericParams := []*ast.GenericParam{
		{Name: "T"},
	}
	
	paramTypes := []ast.Type{
		&ast.TypeReference{Name: "T"},
	}
	
	// Wrong number of arguments
	argTypes := []ast.Type{
		&ast.TypeReference{Name: "Int"},
		&ast.TypeReference{Name: "String"},
	}
	
	_, err := gc.InferTypeArguments(genericParams, paramTypes, argTypes)
	if err == nil {
		t.Error("Expected inference to fail with argument count mismatch")
	}
}

// Test 29: Constraints checking (basic)
func TestConstraintsChecking(t *testing.T) {
	// For now, constraints checking just returns success
	// This will be expanded when traits are implemented
	typ := &ast.TypeReference{Name: "Int"}
	constraints := []ast.Type{
		&ast.TypeReference{Name: "Comparable"},
	}
	
	err := CheckConstraints(typ, constraints)
	if err != nil {
		t.Errorf("Expected constraints check to succeed, got: %v", err)
	}
}

// Test 30: Multiple substitutions in sequence
func TestMultipleSubstitutions(t *testing.T) {
	gc := NewGenericContext()
	gc.AddTypeParameter(&ast.GenericParam{Name: "T"})
	gc.AddTypeParameter(&ast.GenericParam{Name: "U"})
	
	// First substitution: T -> U
	gc.Substitutions["T"] = &ast.TypeReference{Name: "U"}
	
	input1 := &ast.TypeReference{Name: "T"}
	result1 := gc.Substitute(input1)
	result1Ref := result1.(*ast.TypeReference)
	if result1Ref.Name != "U" {
		t.Errorf("Expected T to substitute to U, got %s", result1Ref.Name)
	}
	
	// Second substitution: U -> Int
	gc.Substitutions["U"] = &ast.TypeReference{Name: "Int"}
	
	input2 := &ast.TypeReference{Name: "U"}
	result2 := gc.Substitute(input2)
	result2Ref := result2.(*ast.TypeReference)
	if result2Ref.Name != "Int" {
		t.Errorf("Expected U to substitute to Int, got %s", result2Ref.Name)
	}
}
