package vm

import (
	"testing"
	"unified-compiler/internal/bytecode"
)

// Test 1: Try operator unwraps Ok variant
func TestTryOperatorUnwrapsOk(t *testing.T) {
	bc := bytecode.NewBytecode()
	bc.AddFunction("main", 0)

	// Create Result::Ok(42)
	okValueIdx := bc.AddConstant(bytecode.NewIntValue(42))
	enumNameIdx := bc.AddConstant(bytecode.NewStringValue("Result"))
	variantNameIdx := bc.AddConstant(bytecode.NewStringValue("Ok"))
	tagIdx := bc.AddConstant(bytecode.NewIntValue(0))

	bc.AddInstruction(bytecode.OpPush, int64(okValueIdx))       // Push value
	bc.AddInstruction(bytecode.OpPush, int64(enumNameIdx))      // Push enum name
	bc.AddInstruction(bytecode.OpPush, int64(variantNameIdx))   // Push variant name
	bc.AddInstruction(bytecode.OpPush, int64(tagIdx))           // Push tag
	bc.AddInstruction(bytecode.OpAllocEnum, 1)                  // Create Ok(42)
	bc.AddInstruction(bytecode.OpTryPropagate, 0)               // Try operator
	bc.AddInstruction(bytecode.OpReturnValue, 0)                // Return the unwrapped value

	vm := NewVM(bc)
	result, err := vm.Run()
	if err != nil {
		t.Fatalf("VM error: %v", err)
	}

	// Result should have the unwrapped value (42)
	if result.Type != bytecode.ValueTypeInt {
		t.Errorf("Expected int type, got %v", result.Type)
	}
	if result.Int != 42 {
		t.Errorf("Expected 42, got %d", result.Int)
	}
}

// Test 2: Try operator propagates Err variant
func TestTryOperatorPropagatesErr(t *testing.T) {
	bc := bytecode.NewBytecode()
	bc.AddFunction("main", 0)

	// Create Result::Err(100)
	errValueIdx := bc.AddConstant(bytecode.NewIntValue(100))
	enumNameIdx := bc.AddConstant(bytecode.NewStringValue("Result"))
	variantNameIdx := bc.AddConstant(bytecode.NewStringValue("Err"))
	tagIdx := bc.AddConstant(bytecode.NewIntValue(1))

	bc.AddInstruction(bytecode.OpPush, int64(errValueIdx))      // Push error value
	bc.AddInstruction(bytecode.OpPush, int64(enumNameIdx))      // Push enum name
	bc.AddInstruction(bytecode.OpPush, int64(variantNameIdx))   // Push variant name
	bc.AddInstruction(bytecode.OpPush, int64(tagIdx))           // Push tag
	bc.AddInstruction(bytecode.OpAllocEnum, 1)                  // Create Err(100)
	bc.AddInstruction(bytecode.OpTryPropagate, 0)               // Try operator
	bc.AddInstruction(bytecode.OpReturnValue, 0)

	vm := NewVM(bc)
	result, err := vm.Run()
	if err != nil {
		t.Fatalf("VM error: %v", err)
	}

	// Result should have the Err variant (early return)
	if result.Type != bytecode.ValueTypeEnum {
		t.Errorf("Expected enum type, got %v", result.Type)
	}
	if result.Enum.VariantName != "Err" {
		t.Errorf("Expected Err variant, got %s", result.Enum.VariantName)
	}
	if len(result.Enum.Data) == 0 || result.Enum.Data[0].Int != 100 {
		t.Errorf("Expected error value 100")
	}
}

// Test 3: Multiple try operators with Ok values
func TestMultipleTryOperatorsWithOk(t *testing.T) {
	bc := bytecode.NewBytecode()
	bc.AddFunction("main", 0)

	// Create Ok(10) and Ok(20), then try both
	enumNameIdx := bc.AddConstant(bytecode.NewStringValue("Result"))
	okVariantIdx := bc.AddConstant(bytecode.NewStringValue("Ok"))
	tagIdx := bc.AddConstant(bytecode.NewIntValue(0))

	// First Ok(10)
	val1Idx := bc.AddConstant(bytecode.NewIntValue(10))
	bc.AddInstruction(bytecode.OpPush, int64(val1Idx))
	bc.AddInstruction(bytecode.OpPush, int64(enumNameIdx))
	bc.AddInstruction(bytecode.OpPush, int64(okVariantIdx))
	bc.AddInstruction(bytecode.OpPush, int64(tagIdx))
	bc.AddInstruction(bytecode.OpAllocEnum, 1)
	bc.AddInstruction(bytecode.OpTryPropagate, 0)  // Unwrap to 10

	// Second Ok(20)
	val2Idx := bc.AddConstant(bytecode.NewIntValue(20))
	bc.AddInstruction(bytecode.OpPush, int64(val2Idx))
	bc.AddInstruction(bytecode.OpPush, int64(enumNameIdx))
	bc.AddInstruction(bytecode.OpPush, int64(okVariantIdx))
	bc.AddInstruction(bytecode.OpPush, int64(tagIdx))
	bc.AddInstruction(bytecode.OpAllocEnum, 1)
	bc.AddInstruction(bytecode.OpTryPropagate, 0)  // Unwrap to 20

	// Add the two values
	bc.AddInstruction(bytecode.OpAdd, 0)
	bc.AddInstruction(bytecode.OpReturnValue, 0)

	vm := NewVM(bc)
	result, err := vm.Run()
	if err != nil {
		t.Fatalf("VM error: %v", err)
	}

	// Should have 30 as result
	if result.Type != bytecode.ValueTypeInt || result.Int != 30 {
		t.Errorf("Expected 30, got %v", result)
	}
}

// Test 4: Try operator short-circuits on first Err
func TestTryOperatorShortCircuits(t *testing.T) {
	bc := bytecode.NewBytecode()
	bc.AddFunction("main", 0)

	enumNameIdx := bc.AddConstant(bytecode.NewStringValue("Result"))
	okVariantIdx := bc.AddConstant(bytecode.NewStringValue("Ok"))
	errVariantIdx := bc.AddConstant(bytecode.NewStringValue("Err"))
	okTagIdx := bc.AddConstant(bytecode.NewIntValue(0))
	errTagIdx := bc.AddConstant(bytecode.NewIntValue(1))

	// First Ok(10)
	val1Idx := bc.AddConstant(bytecode.NewIntValue(10))
	bc.AddInstruction(bytecode.OpPush, int64(val1Idx))
	bc.AddInstruction(bytecode.OpPush, int64(enumNameIdx))
	bc.AddInstruction(bytecode.OpPush, int64(okVariantIdx))
	bc.AddInstruction(bytecode.OpPush, int64(okTagIdx))
	bc.AddInstruction(bytecode.OpAllocEnum, 1)
	bc.AddInstruction(bytecode.OpTryPropagate, 0)  // Unwrap to 10

	// Second Err(99)
	errValIdx := bc.AddConstant(bytecode.NewIntValue(99))
	bc.AddInstruction(bytecode.OpPush, int64(errValIdx))
	bc.AddInstruction(bytecode.OpPush, int64(enumNameIdx))
	bc.AddInstruction(bytecode.OpPush, int64(errVariantIdx))
	bc.AddInstruction(bytecode.OpPush, int64(errTagIdx))
	bc.AddInstruction(bytecode.OpAllocEnum, 1)
	bc.AddInstruction(bytecode.OpTryPropagate, 0)  // Should propagate error

	// This should not execute
	bc.AddInstruction(bytecode.OpAdd, 0)
	bc.AddInstruction(bytecode.OpReturnValue, 0)

	vm := NewVM(bc)
	result, err := vm.Run()
	if err != nil {
		t.Fatalf("VM error: %v", err)
	}

	// Should have Err as result, not the sum
	if result.Type != bytecode.ValueTypeEnum {
		t.Errorf("Expected enum type, got %v", result.Type)
	}
	if result.Enum.VariantName != "Err" {
		t.Errorf("Expected Err variant, got %s", result.Enum.VariantName)
	}
}

// Test 5: Try operator with wrong type should error
func TestTryOperatorWithWrongType(t *testing.T) {
	bc := bytecode.NewBytecode()
	bc.AddFunction("main", 0)

	// Try to use ? on a plain integer
	valIdx := bc.AddConstant(bytecode.NewIntValue(42))
	bc.AddInstruction(bytecode.OpPush, int64(valIdx))
	bc.AddInstruction(bytecode.OpTryPropagate, 0)
	bc.AddInstruction(bytecode.OpReturnValue, 0)

	vm := NewVM(bc)
	_, err := vm.Run()
	if err == nil {
		t.Fatal("Expected error when using ? on non-Result type")
	}
}

// Test 6: Try operator with Option::Some
func TestTryOperatorWithOptionSome(t *testing.T) {
	bc := bytecode.NewBytecode()
	bc.AddFunction("main", 0)

	// Create Option::Some(42) - but try operator expects Result
	someValueIdx := bc.AddConstant(bytecode.NewIntValue(42))
	enumNameIdx := bc.AddConstant(bytecode.NewStringValue("Option"))
	variantNameIdx := bc.AddConstant(bytecode.NewStringValue("Some"))
	tagIdx := bc.AddConstant(bytecode.NewIntValue(0))

	bc.AddInstruction(bytecode.OpPush, int64(someValueIdx))
	bc.AddInstruction(bytecode.OpPush, int64(enumNameIdx))
	bc.AddInstruction(bytecode.OpPush, int64(variantNameIdx))
	bc.AddInstruction(bytecode.OpPush, int64(tagIdx))
	bc.AddInstruction(bytecode.OpAllocEnum, 1)
	bc.AddInstruction(bytecode.OpTryPropagate, 0)
	bc.AddInstruction(bytecode.OpReturnValue, 0)

	vm := NewVM(bc)
	_, err := vm.Run()
	if err == nil {
		t.Fatal("Expected error when using ? on Option instead of Result")
	}
}

// Test 7: Try operator with Ok but no data should error
func TestTryOperatorWithEmptyOk(t *testing.T) {
	bc := bytecode.NewBytecode()
	bc.AddFunction("main", 0)

	// Create Result::Ok() with no data
	enumNameIdx := bc.AddConstant(bytecode.NewStringValue("Result"))
	variantNameIdx := bc.AddConstant(bytecode.NewStringValue("Ok"))
	tagIdx := bc.AddConstant(bytecode.NewIntValue(0))

	bc.AddInstruction(bytecode.OpPush, int64(enumNameIdx))
	bc.AddInstruction(bytecode.OpPush, int64(variantNameIdx))
	bc.AddInstruction(bytecode.OpPush, int64(tagIdx))
	bc.AddInstruction(bytecode.OpAllocEnum, 0)  // No data
	bc.AddInstruction(bytecode.OpTryPropagate, 0)
	bc.AddInstruction(bytecode.OpReturnValue, 0)

	vm := NewVM(bc)
	_, err := vm.Run()
	if err == nil {
		t.Fatal("Expected error when Ok has no data")
	}
}

// Test 8: Try operator with string value in Ok
func TestTryOperatorWithStringOk(t *testing.T) {
	bc := bytecode.NewBytecode()
	bc.AddFunction("main", 0)

	// Create Result::Ok("hello")
	strValueIdx := bc.AddConstant(bytecode.NewStringValue("hello"))
	enumNameIdx := bc.AddConstant(bytecode.NewStringValue("Result"))
	variantNameIdx := bc.AddConstant(bytecode.NewStringValue("Ok"))
	tagIdx := bc.AddConstant(bytecode.NewIntValue(0))

	bc.AddInstruction(bytecode.OpPush, int64(strValueIdx))
	bc.AddInstruction(bytecode.OpPush, int64(enumNameIdx))
	bc.AddInstruction(bytecode.OpPush, int64(variantNameIdx))
	bc.AddInstruction(bytecode.OpPush, int64(tagIdx))
	bc.AddInstruction(bytecode.OpAllocEnum, 1)
	bc.AddInstruction(bytecode.OpTryPropagate, 0)
	bc.AddInstruction(bytecode.OpReturnValue, 0)

	vm := NewVM(bc)
	result, err := vm.Run()
	if err != nil {
		t.Fatalf("VM error: %v", err)
	}

	if result.Type != bytecode.ValueTypeString {
		t.Errorf("Expected string type, got %v", result.Type)
	}
	if result.Str != "hello" {
		t.Errorf("Expected 'hello', got '%s'", result.Str)
	}
}

// Test 9: Try operator with boolean value in Ok
func TestTryOperatorWithBoolOk(t *testing.T) {
	bc := bytecode.NewBytecode()
	bc.AddFunction("main", 0)

	// Create Result::Ok(true)
	boolValueIdx := bc.AddConstant(bytecode.NewBoolValue(true))
	enumNameIdx := bc.AddConstant(bytecode.NewStringValue("Result"))
	variantNameIdx := bc.AddConstant(bytecode.NewStringValue("Ok"))
	tagIdx := bc.AddConstant(bytecode.NewIntValue(0))

	bc.AddInstruction(bytecode.OpPush, int64(boolValueIdx))
	bc.AddInstruction(bytecode.OpPush, int64(enumNameIdx))
	bc.AddInstruction(bytecode.OpPush, int64(variantNameIdx))
	bc.AddInstruction(bytecode.OpPush, int64(tagIdx))
	bc.AddInstruction(bytecode.OpAllocEnum, 1)
	bc.AddInstruction(bytecode.OpTryPropagate, 0)
	bc.AddInstruction(bytecode.OpReturnValue, 0)

	vm := NewVM(bc)
	result, err := vm.Run()
	if err != nil {
		t.Fatalf("VM error: %v", err)
	}

	if result.Type != bytecode.ValueTypeBool {
		t.Errorf("Expected bool type, got %v", result.Type)
	}
	if result.Bool != true {
		t.Errorf("Expected true, got %v", result.Bool)
	}
}

// Test 10: Try operator with nested Result
func TestTryOperatorWithNestedResult(t *testing.T) {
	bc := bytecode.NewBytecode()
	bc.AddFunction("main", 0)

	enumNameIdx := bc.AddConstant(bytecode.NewStringValue("Result"))
	okVariantIdx := bc.AddConstant(bytecode.NewStringValue("Ok"))
	tagIdx := bc.AddConstant(bytecode.NewIntValue(0))

	// Create inner Result::Ok(5)
	innerValueIdx := bc.AddConstant(bytecode.NewIntValue(5))
	bc.AddInstruction(bytecode.OpPush, int64(innerValueIdx))
	bc.AddInstruction(bytecode.OpPush, int64(enumNameIdx))
	bc.AddInstruction(bytecode.OpPush, int64(okVariantIdx))
	bc.AddInstruction(bytecode.OpPush, int64(tagIdx))
	bc.AddInstruction(bytecode.OpAllocEnum, 1)

	// Wrap in outer Result::Ok(Result::Ok(5))
	bc.AddInstruction(bytecode.OpPush, int64(enumNameIdx))
	bc.AddInstruction(bytecode.OpPush, int64(okVariantIdx))
	bc.AddInstruction(bytecode.OpPush, int64(tagIdx))
	bc.AddInstruction(bytecode.OpAllocEnum, 1)

	// First try unwraps outer Result
	bc.AddInstruction(bytecode.OpTryPropagate, 0)

	// Second try unwraps inner Result
	bc.AddInstruction(bytecode.OpTryPropagate, 0)

	bc.AddInstruction(bytecode.OpReturnValue, 0)

	vm := NewVM(bc)
	result, err := vm.Run()
	if err != nil {
		t.Fatalf("VM error: %v", err)
	}

	if result.Type != bytecode.ValueTypeInt {
		t.Errorf("Expected int type, got %v", result.Type)
	}
	if result.Int != 5 {
		t.Errorf("Expected 5, got %d", result.Int)
	}
}
