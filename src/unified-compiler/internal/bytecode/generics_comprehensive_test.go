package bytecode

import (
	"testing"
	"unified-compiler/internal/ast"
	"unified-compiler/internal/semantic"
)

// TestGenericStructMonomorphization tests monomorphizing generic structs
func TestGenericStructMonomorphization(t *testing.T) {
	g := NewGenerator()
	
	// Create a generic Box<T> struct
	boxStruct := &ast.StructDecl{
		Name:     "Box",
		IsPublic: true,
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
	
	// Store the generic struct template
	g.genericStructs["Box"] = boxStruct
	
	// Monomorphize Box<Int>
	intType := &ast.TypeReference{Name: "Int"}
	mangledName, err := g.monomorphizeStruct("Box", []ast.Type{intType})
	if err != nil {
		t.Fatalf("Failed to monomorphize Box<Int>: %v", err)
	}
	
	// Check the mangled name
	expectedName := semantic.MangleName("Box", []ast.Type{intType})
	if mangledName != expectedName {
		t.Errorf("Expected mangled name %s, got %s", expectedName, mangledName)
	}
	
	// Verify the struct was registered
	if _, ok := g.structTypes[mangledName]; !ok {
		t.Errorf("Monomorphized struct %s was not registered", mangledName)
	}
	
	// Verify it's marked as generated
	if mono, ok := g.monomorphizedStructs[mangledName]; !ok || !mono.Generated {
		t.Error("Monomorphized struct not marked as generated")
	}
	
	// Test duplicate monomorphization returns same name
	mangledName2, err := g.monomorphizeStruct("Box", []ast.Type{intType})
	if err != nil {
		t.Fatalf("Failed on duplicate monomorphization: %v", err)
	}
	if mangledName != mangledName2 {
		t.Errorf("Duplicate monomorphization returned different name: %s vs %s", mangledName, mangledName2)
	}
}

// TestGenericStructMultipleSpecializations tests creating multiple specializations
func TestGenericStructMultipleSpecializations(t *testing.T) {
	g := NewGenerator()
	
	// Create a generic Box<T> struct
	boxStruct := &ast.StructDecl{
		Name:     "Box",
		IsPublic: true,
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
	
	g.genericStructs["Box"] = boxStruct
	
	// Monomorphize Box<Int>
	intType := &ast.TypeReference{Name: "Int"}
	intName, err := g.monomorphizeStruct("Box", []ast.Type{intType})
	if err != nil {
		t.Fatalf("Failed to monomorphize Box<Int>: %v", err)
	}
	
	// Monomorphize Box<String>
	stringType := &ast.TypeReference{Name: "String"}
	stringName, err := g.monomorphizeStruct("Box", []ast.Type{stringType})
	if err != nil {
		t.Fatalf("Failed to monomorphize Box<String>: %v", err)
	}
	
	// Names should be different
	if intName == stringName {
		t.Error("Different type arguments produced same mangled name")
	}
	
	// Both should be registered
	if _, ok := g.structTypes[intName]; !ok {
		t.Error("Box<Int> not registered")
	}
	if _, ok := g.structTypes[stringName]; !ok {
		t.Error("Box<String> not registered")
	}
}

// TestGenericEnumMonomorphization tests monomorphizing generic enums
func TestGenericEnumMonomorphization(t *testing.T) {
	g := NewGenerator()
	
	// Create a generic Option<T> enum
	optionEnum := &ast.EnumDecl{
		Name:     "Option",
		IsPublic: true,
		GenericParams: []*ast.GenericParam{
			{Name: "T"},
		},
		Variants: []*ast.EnumVariant{
			{
				Name: "Some",
				Parameters: []*ast.EnumVariantParam{
					{Name: "value", Type: &ast.TypeReference{Name: "T"}},
				},
			},
			{
				Name:       "None",
				Parameters: []*ast.EnumVariantParam{},
			},
		},
	}
	
	g.genericEnums["Option"] = optionEnum
	
	// Monomorphize Option<Int>
	intType := &ast.TypeReference{Name: "Int"}
	mangledName, err := g.monomorphizeEnum("Option", []ast.Type{intType})
	if err != nil {
		t.Fatalf("Failed to monomorphize Option<Int>: %v", err)
	}
	
	// Check the mangled name
	expectedName := semantic.MangleName("Option", []ast.Type{intType})
	if mangledName != expectedName {
		t.Errorf("Expected mangled name %s, got %s", expectedName, mangledName)
	}
	
	// Verify the enum was registered
	if _, ok := g.enumTypes[mangledName]; !ok {
		t.Errorf("Monomorphized enum %s was not registered", mangledName)
	}
	
	// Verify it's marked as generated
	if mono, ok := g.monomorphizedEnums[mangledName]; !ok || !mono.Generated {
		t.Error("Monomorphized enum not marked as generated")
	}
}

// TestGenericEnumMultipleSpecializations tests creating multiple enum specializations
func TestGenericEnumMultipleSpecializations(t *testing.T) {
	g := NewGenerator()
	
	// Create a generic Result<T, E> enum
	resultEnum := &ast.EnumDecl{
		Name:     "Result",
		IsPublic: true,
		GenericParams: []*ast.GenericParam{
			{Name: "T"},
			{Name: "E"},
		},
		Variants: []*ast.EnumVariant{
			{
				Name: "Success",
				Parameters: []*ast.EnumVariantParam{
					{Name: "value", Type: &ast.TypeReference{Name: "T"}},
				},
			},
			{
				Name: "Error",
				Parameters: []*ast.EnumVariantParam{
					{Name: "error", Type: &ast.TypeReference{Name: "E"}},
				},
			},
		},
	}
	
	g.genericEnums["Result"] = resultEnum
	
	// Monomorphize Result<Int, String>
	intType := &ast.TypeReference{Name: "Int"}
	stringType := &ast.TypeReference{Name: "String"}
	name1, err := g.monomorphizeEnum("Result", []ast.Type{intType, stringType})
	if err != nil {
		t.Fatalf("Failed to monomorphize Result<Int, String>: %v", err)
	}
	
	// Monomorphize Result<String, Int>
	name2, err := g.monomorphizeEnum("Result", []ast.Type{stringType, intType})
	if err != nil {
		t.Fatalf("Failed to monomorphize Result<String, Int>: %v", err)
	}
	
	// Names should be different (different type order)
	if name1 == name2 {
		t.Error("Different type argument order produced same mangled name")
	}
	
	// Both should be registered
	if _, ok := g.enumTypes[name1]; !ok {
		t.Error("Result<Int, String> not registered")
	}
	if _, ok := g.enumTypes[name2]; !ok {
		t.Error("Result<String, Int> not registered")
	}
}

// TestGenericStructWrongArity tests error handling for wrong number of type arguments
func TestGenericStructWrongArity(t *testing.T) {
	g := NewGenerator()
	
	boxStruct := &ast.StructDecl{
		Name:     "Box",
		GenericParams: []*ast.GenericParam{
			{Name: "T"},
		},
		Members: []*ast.StructMember{},
	}
	
	g.genericStructs["Box"] = boxStruct
	
	// Try to monomorphize with wrong number of type arguments
	intType := &ast.TypeReference{Name: "Int"}
	stringType := &ast.TypeReference{Name: "String"}
	_, err := g.monomorphizeStruct("Box", []ast.Type{intType, stringType})
	
	if err == nil {
		t.Error("Expected error for wrong type argument count, got nil")
	}
}

// TestGenericEnumWrongArity tests error handling for wrong number of type arguments
func TestGenericEnumWrongArity(t *testing.T) {
	g := NewGenerator()
	
	optionEnum := &ast.EnumDecl{
		Name:     "Option",
		GenericParams: []*ast.GenericParam{
			{Name: "T"},
		},
		Variants: []*ast.EnumVariant{},
	}
	
	g.genericEnums["Option"] = optionEnum
	
	// Try to monomorphize with no type arguments
	_, err := g.monomorphizeEnum("Option", []ast.Type{})
	
	if err == nil {
		t.Error("Expected error for wrong type argument count, got nil")
	}
}

// TestNonGenericStructWithTypeArgs tests error when providing type args to non-generic struct
func TestNonGenericStructWithTypeArgs(t *testing.T) {
	g := NewGenerator()
	
	// Create a non-generic struct
	simpleStruct := &ast.StructDecl{
		Name:    "Simple",
		Members: []*ast.StructMember{},
	}
	
	// Register it normally (not in genericStructs)
	g.registerStructType(simpleStruct)
	
	// Try to use type arguments with it in a struct expression
	structExpr := &ast.StructExpr{
		Name:     "Simple",
		TypeArgs: []ast.Type{&ast.TypeReference{Name: "Int"}},
	}
	
	err := g.generateStructExpr(structExpr)
	if err == nil {
		t.Error("Expected error for non-generic struct with type arguments")
	}
}

// TestNonGenericEnumWithTypeArgs tests error when providing type args to non-generic enum
func TestNonGenericEnumWithTypeArgs(t *testing.T) {
	g := NewGenerator()
	
	// Create a non-generic enum
	simpleEnum := &ast.EnumDecl{
		Name: "Simple",
		Variants: []*ast.EnumVariant{
			{Name: "A"},
		},
	}
	
	// Register it normally (not in genericEnums)
	g.registerEnumType(simpleEnum)
	
	// Try to use type arguments with it in an enum constructor
	enumExpr := &ast.EnumConstructorExpr{
		EnumName: "Simple",
		Variant:  "A",
		TypeArgs: []ast.Type{&ast.TypeReference{Name: "Int"}},
	}
	
	err := g.generateEnumConstructorExpr(enumExpr)
	if err == nil {
		t.Error("Expected error for non-generic enum with type arguments")
	}
}

// TestNestedGenericTypes tests monomorphization with nested generic types
func TestNestedGenericTypes(t *testing.T) {
	g := NewGenerator()
	
	// Create Box<T>
	boxStruct := &ast.StructDecl{
		Name:     "Box",
		GenericParams: []*ast.GenericParam{
			{Name: "T"},
		},
		Members: []*ast.StructMember{
			{Name: "value", Type: &ast.TypeReference{Name: "T"}},
		},
	}
	g.genericStructs["Box"] = boxStruct
	
	// Create Option<T>
	optionEnum := &ast.EnumDecl{
		Name:     "Option",
		GenericParams: []*ast.GenericParam{
			{Name: "T"},
		},
		Variants: []*ast.EnumVariant{
			{
				Name: "Some",
				Parameters: []*ast.EnumVariantParam{
					{Name: "value", Type: &ast.TypeReference{Name: "T"}},
				},
			},
		},
	}
	g.genericEnums["Option"] = optionEnum
	
	// First monomorphize Box<Int>
	intType := &ast.TypeReference{Name: "Int"}
	boxIntName, err := g.monomorphizeStruct("Box", []ast.Type{intType})
	if err != nil {
		t.Fatalf("Failed to monomorphize Box<Int>: %v", err)
	}
	
	// Then monomorphize Option<Box<Int>> - use the mangled name
	boxIntType := &ast.TypeReference{Name: boxIntName}
	optionBoxIntName, err := g.monomorphizeEnum("Option", []ast.Type{boxIntType})
	if err != nil {
		t.Fatalf("Failed to monomorphize Option<Box<Int>>: %v", err)
	}
	
	// Verify both are registered
	if _, ok := g.structTypes[boxIntName]; !ok {
		t.Error("Box<Int> not registered")
	}
	if _, ok := g.enumTypes[optionBoxIntName]; !ok {
		t.Error("Option<Box<Int>> not registered")
	}
}

// TestGenericStructFieldSubstitution tests that field types are properly substituted
func TestGenericStructFieldSubstitution(t *testing.T) {
	g := NewGenerator()
	
	// Create a generic Pair<A, B> struct
	pairStruct := &ast.StructDecl{
		Name:     "Pair",
		GenericParams: []*ast.GenericParam{
			{Name: "A"},
			{Name: "B"},
		},
		Members: []*ast.StructMember{
			{Name: "first", Type: &ast.TypeReference{Name: "A"}},
			{Name: "second", Type: &ast.TypeReference{Name: "B"}},
		},
	}
	g.genericStructs["Pair"] = pairStruct
	
	// Monomorphize Pair<Int, String>
	intType := &ast.TypeReference{Name: "Int"}
	stringType := &ast.TypeReference{Name: "String"}
	mangledName, err := g.monomorphizeStruct("Pair", []ast.Type{intType, stringType})
	if err != nil {
		t.Fatalf("Failed to monomorphize Pair<Int, String>: %v", err)
	}
	
	// Verify the struct is registered with correct fields
	structInfo, ok := g.structTypes[mangledName]
	if !ok {
		t.Fatal("Monomorphized struct not registered")
	}
	
	// Check that both fields exist
	if _, ok := structInfo.Fields["first"]; !ok {
		t.Error("Field 'first' not found in monomorphized struct")
	}
	if _, ok := structInfo.Fields["second"]; !ok {
		t.Error("Field 'second' not found in monomorphized struct")
	}
}

// TestEmptyGenericStruct tests monomorphizing a struct with no fields
func TestEmptyGenericStruct(t *testing.T) {
	g := NewGenerator()
	
	emptyStruct := &ast.StructDecl{
		Name:     "Empty",
		GenericParams: []*ast.GenericParam{
			{Name: "T"},
		},
		Members: []*ast.StructMember{},
	}
	g.genericStructs["Empty"] = emptyStruct
	
	intType := &ast.TypeReference{Name: "Int"}
	mangledName, err := g.monomorphizeStruct("Empty", []ast.Type{intType})
	if err != nil {
		t.Fatalf("Failed to monomorphize Empty<Int>: %v", err)
	}
	
	// Should still be registered
	if _, ok := g.structTypes[mangledName]; !ok {
		t.Error("Empty monomorphized struct not registered")
	}
}

// TestSingleVariantGenericEnum tests monomorphizing enum with one variant
func TestSingleVariantGenericEnum(t *testing.T) {
	g := NewGenerator()
	
	singleEnum := &ast.EnumDecl{
		Name:     "Single",
		GenericParams: []*ast.GenericParam{
			{Name: "T"},
		},
		Variants: []*ast.EnumVariant{
			{
				Name: "Only",
				Parameters: []*ast.EnumVariantParam{
					{Name: "value", Type: &ast.TypeReference{Name: "T"}},
				},
			},
		},
	}
	g.genericEnums["Single"] = singleEnum
	
	intType := &ast.TypeReference{Name: "Int"}
	mangledName, err := g.monomorphizeEnum("Single", []ast.Type{intType})
	if err != nil {
		t.Fatalf("Failed to monomorphize Single<Int>: %v", err)
	}
	
	// Verify it's registered
	enumInfo, ok := g.enumTypes[mangledName]
	if !ok {
		t.Fatal("Monomorphized enum not registered")
	}
	
	// Check variant exists
	if _, ok := enumInfo.Variants["Only"]; !ok {
		t.Error("Variant 'Only' not found in monomorphized enum")
	}
}
