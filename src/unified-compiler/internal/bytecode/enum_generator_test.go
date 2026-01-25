package bytecode

import (
	"testing"
	"unified-compiler/internal/ast"
)

// TestEnumTypeRegistration tests that enum types are registered correctly
func TestEnumTypeRegistration(t *testing.T) {
	gen := NewGenerator()

	// Create a simple enum
	enum := &ast.EnumDecl{
		Name: "Color",
		Variants: []*ast.EnumVariant{
			{Name: "Red", Parameters: []*ast.EnumVariantParam{}},
			{Name: "Green", Parameters: []*ast.EnumVariantParam{}},
			{Name: "Blue", Parameters: []*ast.EnumVariantParam{}},
		},
	}

	err := gen.registerEnumType(enum)
	if err != nil {
		t.Fatalf("Failed to register enum type: %v", err)
	}

	// Check that enum was registered
	info, ok := gen.enumTypes["Color"]
	if !ok {
		t.Fatal("Enum type 'Color' was not registered")
	}

	if info.Name != "Color" {
		t.Errorf("Expected enum name 'Color', got '%s'", info.Name)
	}

	if len(info.Variants) != 3 {
		t.Errorf("Expected 3 variants, got %d", len(info.Variants))
	}

	// Check variants
	variants := []string{"Red", "Green", "Blue"}
	for i, name := range variants {
		variant, ok := info.Variants[name]
		if !ok {
			t.Errorf("Variant '%s' not found", name)
			continue
		}
		if variant.Name != name {
			t.Errorf("Expected variant name '%s', got '%s'", name, variant.Name)
		}
		if variant.Tag != i {
			t.Errorf("Expected tag %d for variant '%s', got %d", i, name, variant.Tag)
		}
		if variant.Arity != 0 {
			t.Errorf("Expected arity 0 for variant '%s', got %d", name, variant.Arity)
		}
	}
}

// TestEnumWithDataRegistration tests enum variants with data
func TestEnumWithDataRegistration(t *testing.T) {
	gen := NewGenerator()

	// Create Option enum with data
	enum := &ast.EnumDecl{
		Name: "Option",
		Variants: []*ast.EnumVariant{
			{Name: "None", Parameters: []*ast.EnumVariantParam{}},
			{
				Name: "Some",
				Parameters: []*ast.EnumVariantParam{
					{Type: &ast.TypeReference{Name: "Int"}},
				},
			},
		},
	}

	err := gen.registerEnumType(enum)
	if err != nil {
		t.Fatalf("Failed to register enum type: %v", err)
	}

	info, ok := gen.enumTypes["Option"]
	if !ok {
		t.Fatal("Enum type 'Option' was not registered")
	}

	// Check None variant
	none, ok := info.Variants["None"]
	if !ok {
		t.Fatal("Variant 'None' not found")
	}
	if none.Arity != 0 {
		t.Errorf("Expected arity 0 for 'None', got %d", none.Arity)
	}

	// Check Some variant
	some, ok := info.Variants["Some"]
	if !ok {
		t.Fatal("Variant 'Some' not found")
	}
	if some.Arity != 1 {
		t.Errorf("Expected arity 1 for 'Some', got %d", some.Arity)
	}
}

// TestEnumConstructorGeneration tests bytecode generation for enum constructors
func TestEnumConstructorGeneration(t *testing.T) {
	gen := NewGenerator()

	// Register Color enum
	colorEnum := &ast.EnumDecl{
		Name: "Color",
		Variants: []*ast.EnumVariant{
			{Name: "Red", Parameters: []*ast.EnumVariantParam{}},
		},
	}
	gen.registerEnumType(colorEnum)

	// Generate constructor for Color::Red
	constructor := &ast.EnumConstructorExpr{
		EnumName: "Color",
		Variant:  "Red",
		Arguments: []ast.Expression{},
	}

	err := gen.generateEnumConstructorExpr(constructor)
	if err != nil {
		t.Fatalf("Failed to generate enum constructor: %v", err)
	}

	// Check that bytecode was generated
	bc := gen.bytecode
	if len(bc.Instructions) == 0 {
		t.Fatal("No bytecode instructions generated")
	}

	// Last instruction should be OpAllocEnum
	lastInst := bc.Instructions[len(bc.Instructions)-1]
	if lastInst.Op != OpAllocEnum {
		t.Errorf("Expected OpAllocEnum, got %v", lastInst.Op)
	}
	if lastInst.Operand != 0 {
		t.Errorf("Expected operand 0 (no data), got %d", lastInst.Operand)
	}
}

// TestEnumConstructorWithData tests bytecode generation for enum with data
func TestEnumConstructorWithData(t *testing.T) {
	gen := NewGenerator()

	// Register Option enum
	optionEnum := &ast.EnumDecl{
		Name: "Option",
		Variants: []*ast.EnumVariant{
			{
				Name: "Some",
				Parameters: []*ast.EnumVariantParam{
					{Type: &ast.TypeReference{Name: "Int"}},
				},
			},
		},
	}
	gen.registerEnumType(optionEnum)

	// Generate constructor for Option::Some(42)
	constructor := &ast.EnumConstructorExpr{
		EnumName: "Option",
		Variant:  "Some",
		Arguments: []ast.Expression{
			&ast.Literal{Kind: ast.LiteralInt, Value: "42"},
		},
	}

	err := gen.generateEnumConstructorExpr(constructor)
	if err != nil {
		t.Fatalf("Failed to generate enum constructor: %v", err)
	}

	// Check that bytecode was generated
	bc := gen.bytecode
	if len(bc.Instructions) == 0 {
		t.Fatal("No bytecode instructions generated")
	}

	// Last instruction should be OpAllocEnum with operand 1 (one data field)
	lastInst := bc.Instructions[len(bc.Instructions)-1]
	if lastInst.Op != OpAllocEnum {
		t.Errorf("Expected OpAllocEnum, got %v", lastInst.Op)
	}
	if lastInst.Operand != 1 {
		t.Errorf("Expected operand 1 (one data field), got %d", lastInst.Operand)
	}
}

// TestEnumConstructorValidation tests error handling
func TestEnumConstructorValidation(t *testing.T) {
	tests := []struct {
		name        string
		setup       func(*Generator)
		constructor *ast.EnumConstructorExpr
		wantErr     bool
	}{
		{
			name: "Undefined enum type",
			setup: func(g *Generator) {},
			constructor: &ast.EnumConstructorExpr{
				EnumName: "UndefinedEnum",
				Variant:  "Variant",
			},
			wantErr: true,
		},
		{
			name: "Undefined variant",
			setup: func(g *Generator) {
				g.registerEnumType(&ast.EnumDecl{
					Name: "Color",
					Variants: []*ast.EnumVariant{
						{Name: "Red"},
					},
				})
			},
			constructor: &ast.EnumConstructorExpr{
				EnumName: "Color",
				Variant:  "UndefinedVariant",
			},
			wantErr: true,
		},
		{
			name: "Wrong argument count",
			setup: func(g *Generator) {
				g.registerEnumType(&ast.EnumDecl{
					Name: "Option",
					Variants: []*ast.EnumVariant{
						{
							Name: "Some",
							Parameters: []*ast.EnumVariantParam{
								{Type: &ast.TypeReference{Name: "Int"}},
							},
						},
					},
				})
			},
			constructor: &ast.EnumConstructorExpr{
				EnumName:  "Option",
				Variant:   "Some",
				Arguments: []ast.Expression{}, // Should have 1 argument
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gen := NewGenerator()
			tt.setup(gen)

			err := gen.generateEnumConstructorExpr(tt.constructor)
			if (err != nil) != tt.wantErr {
				t.Errorf("Expected error: %v, got: %v", tt.wantErr, err)
			}
		})
	}
}
