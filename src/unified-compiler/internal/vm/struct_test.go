package vm

import (
	"testing"
	"unified-compiler/internal/bytecode"
)

func TestVMStructAllocation(t *testing.T) {
	bc := bytecode.NewBytecode()
	bc.AddFunction("main", 0)
	
	// Create a simple struct: Point { x: 10, y: 20 }
	// Push field name "x"
	xNameIdx := bc.AddConstant(bytecode.NewStringValue("x"))
	bc.AddInstruction(bytecode.OpPush, int64(xNameIdx))
	// Push field value 10
	xValIdx := bc.AddConstant(bytecode.NewIntValue(10))
	bc.AddInstruction(bytecode.OpPush, int64(xValIdx))
	
	// Push field name "y"
	yNameIdx := bc.AddConstant(bytecode.NewStringValue("y"))
	bc.AddInstruction(bytecode.OpPush, int64(yNameIdx))
	// Push field value 20
	yValIdx := bc.AddConstant(bytecode.NewIntValue(20))
	bc.AddInstruction(bytecode.OpPush, int64(yValIdx))
	
	// Push struct type name
	typeNameIdx := bc.AddConstant(bytecode.NewStringValue("Point"))
	bc.AddInstruction(bytecode.OpPush, int64(typeNameIdx))
	
	// Allocate struct with 2 fields
	bc.AddInstruction(bytecode.OpAllocStruct, 2)
	
	bc.AddInstruction(bytecode.OpReturnValue, 0)
	bc.AddInstruction(bytecode.OpHalt, 0)
	
	vm := NewVM(bc)
	result, err := vm.Run()
	
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	
	if result.Type != bytecode.ValueTypeStruct {
		t.Errorf("Expected struct type, got %v", result.Type)
	}
	
	if result.Struct.TypeName != "Point" {
		t.Errorf("Expected type name 'Point', got '%s'", result.Struct.TypeName)
	}
	
	xVal, ok := result.Struct.Fields["x"]
	if !ok {
		t.Fatal("Field 'x' not found")
	}
	if xVal.Int != 10 {
		t.Errorf("Expected x=10, got %d", xVal.Int)
	}
	
	yVal, ok := result.Struct.Fields["y"]
	if !ok {
		t.Fatal("Field 'y' not found")
	}
	if yVal.Int != 20 {
		t.Errorf("Expected y=20, got %d", yVal.Int)
	}
}

func TestVMFieldAccess(t *testing.T) {
	bc := bytecode.NewBytecode()
	bc.AddFunction("main", 0)
	
	// Create a struct: Point { x: 42, y: 99 }
	xNameIdx := bc.AddConstant(bytecode.NewStringValue("x"))
	bc.AddInstruction(bytecode.OpPush, int64(xNameIdx))
	xValIdx := bc.AddConstant(bytecode.NewIntValue(42))
	bc.AddInstruction(bytecode.OpPush, int64(xValIdx))
	
	yNameIdx := bc.AddConstant(bytecode.NewStringValue("y"))
	bc.AddInstruction(bytecode.OpPush, int64(yNameIdx))
	yValIdx := bc.AddConstant(bytecode.NewIntValue(99))
	bc.AddInstruction(bytecode.OpPush, int64(yValIdx))
	
	typeNameIdx := bc.AddConstant(bytecode.NewStringValue("Point"))
	bc.AddInstruction(bytecode.OpPush, int64(typeNameIdx))
	bc.AddInstruction(bytecode.OpAllocStruct, 2)
	
	// Access field x
	fieldNameIdx := bc.AddConstant(bytecode.NewStringValue("x"))
	bc.AddInstruction(bytecode.OpPush, int64(fieldNameIdx))
	bc.AddInstruction(bytecode.OpLoadField, 0)
	
	bc.AddInstruction(bytecode.OpReturnValue, 0)
	bc.AddInstruction(bytecode.OpHalt, 0)
	
	vm := NewVM(bc)
	result, err := vm.Run()
	
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	
	if result.Type != bytecode.ValueTypeInt {
		t.Errorf("Expected int type, got %v", result.Type)
	}
	
	if result.Int != 42 {
		t.Errorf("Expected 42, got %d", result.Int)
	}
}

func TestVMNestedStructs(t *testing.T) {
	t.Skip("Nested struct test requires more complex bytecode generation - tested via integration tests instead")
}
