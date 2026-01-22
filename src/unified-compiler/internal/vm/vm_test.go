package vm

import (
	"testing"
	"unified-compiler/internal/bytecode"
)

func TestNewVM(t *testing.T) {
	bc := bytecode.NewBytecode()
	vm := NewVM(bc)
	
	if vm == nil {
		t.Fatal("NewVM() returned nil")
	}
	if vm.bytecode != bc {
		t.Error("VM bytecode not set correctly")
	}
	if vm.stack == nil {
		t.Error("VM stack not initialized")
	}
	if vm.ip != 0 {
		t.Errorf("Initial IP should be 0, got %d", vm.ip)
	}
}

func TestVMSimpleArithmetic(t *testing.T) {
	tests := []struct {
		name     string
		setup    func(*bytecode.Bytecode)
		expected int64
	}{
		{
			name: "Add two numbers",
			setup: func(bc *bytecode.Bytecode) {
				bc.AddFunction("main", 0)
				idx1 := bc.AddConstant(bytecode.NewIntValue(10))
				idx2 := bc.AddConstant(bytecode.NewIntValue(20))
				bc.AddInstruction(bytecode.OpPush, int64(idx1))
				bc.AddInstruction(bytecode.OpPush, int64(idx2))
				bc.AddInstruction(bytecode.OpAdd, 0)
				bc.AddInstruction(bytecode.OpReturnValue, 0)
				bc.AddInstruction(bytecode.OpHalt, 0)
			},
			expected: 30,
		},
		{
			name: "Subtract two numbers",
			setup: func(bc *bytecode.Bytecode) {
				bc.AddFunction("main", 0)
				idx1 := bc.AddConstant(bytecode.NewIntValue(50))
				idx2 := bc.AddConstant(bytecode.NewIntValue(20))
				bc.AddInstruction(bytecode.OpPush, int64(idx1))
				bc.AddInstruction(bytecode.OpPush, int64(idx2))
				bc.AddInstruction(bytecode.OpSub, 0)
				bc.AddInstruction(bytecode.OpReturnValue, 0)
				bc.AddInstruction(bytecode.OpHalt, 0)
			},
			expected: 30,
		},
		{
			name: "Multiply two numbers",
			setup: func(bc *bytecode.Bytecode) {
				bc.AddFunction("main", 0)
				idx1 := bc.AddConstant(bytecode.NewIntValue(6))
				idx2 := bc.AddConstant(bytecode.NewIntValue(7))
				bc.AddInstruction(bytecode.OpPush, int64(idx1))
				bc.AddInstruction(bytecode.OpPush, int64(idx2))
				bc.AddInstruction(bytecode.OpMul, 0)
				bc.AddInstruction(bytecode.OpReturnValue, 0)
				bc.AddInstruction(bytecode.OpHalt, 0)
			},
			expected: 42,
		},
		{
			name: "Divide two numbers",
			setup: func(bc *bytecode.Bytecode) {
				bc.AddFunction("main", 0)
				idx1 := bc.AddConstant(bytecode.NewIntValue(84))
				idx2 := bc.AddConstant(bytecode.NewIntValue(2))
				bc.AddInstruction(bytecode.OpPush, int64(idx1))
				bc.AddInstruction(bytecode.OpPush, int64(idx2))
				bc.AddInstruction(bytecode.OpDiv, 0)
				bc.AddInstruction(bytecode.OpReturnValue, 0)
				bc.AddInstruction(bytecode.OpHalt, 0)
			},
			expected: 42,
		},
		{
			name: "Modulo operation",
			setup: func(bc *bytecode.Bytecode) {
				bc.AddFunction("main", 0)
				idx1 := bc.AddConstant(bytecode.NewIntValue(17))
				idx2 := bc.AddConstant(bytecode.NewIntValue(5))
				bc.AddInstruction(bytecode.OpPush, int64(idx1))
				bc.AddInstruction(bytecode.OpPush, int64(idx2))
				bc.AddInstruction(bytecode.OpMod, 0)
				bc.AddInstruction(bytecode.OpReturnValue, 0)
				bc.AddInstruction(bytecode.OpHalt, 0)
			},
			expected: 2,
		},
		{
			name: "Negation",
			setup: func(bc *bytecode.Bytecode) {
				bc.AddFunction("main", 0)
				idx := bc.AddConstant(bytecode.NewIntValue(42))
				bc.AddInstruction(bytecode.OpPush, int64(idx))
				bc.AddInstruction(bytecode.OpNeg, 0)
				bc.AddInstruction(bytecode.OpReturnValue, 0)
				bc.AddInstruction(bytecode.OpHalt, 0)
			},
			expected: -42,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bc := bytecode.NewBytecode()
			tt.setup(bc)
			
			vm := NewVM(bc)
			result, err := vm.Run()
			
			if err != nil {
				t.Fatalf("VM execution failed: %v", err)
			}
			
			if result.Type != bytecode.ValueTypeInt {
				t.Fatalf("Expected int result, got %v", result.Type)
			}
			
			if result.Int != tt.expected {
				t.Errorf("Expected %d, got %d", tt.expected, result.Int)
			}
		})
	}
}

func TestVMComparison(t *testing.T) {
	tests := []struct {
		name     string
		setup    func(*bytecode.Bytecode)
		expected bool
	}{
		{
			name: "Equal - true",
			setup: func(bc *bytecode.Bytecode) {
				bc.AddFunction("main", 0)
				idx := bc.AddConstant(bytecode.NewIntValue(42))
				bc.AddInstruction(bytecode.OpPush, int64(idx))
				bc.AddInstruction(bytecode.OpPush, int64(idx))
				bc.AddInstruction(bytecode.OpEq, 0)
				bc.AddInstruction(bytecode.OpReturnValue, 0)
				bc.AddInstruction(bytecode.OpHalt, 0)
			},
			expected: true,
		},
		{
			name: "Equal - false",
			setup: func(bc *bytecode.Bytecode) {
				bc.AddFunction("main", 0)
				idx1 := bc.AddConstant(bytecode.NewIntValue(42))
				idx2 := bc.AddConstant(bytecode.NewIntValue(43))
				bc.AddInstruction(bytecode.OpPush, int64(idx1))
				bc.AddInstruction(bytecode.OpPush, int64(idx2))
				bc.AddInstruction(bytecode.OpEq, 0)
				bc.AddInstruction(bytecode.OpReturnValue, 0)
				bc.AddInstruction(bytecode.OpHalt, 0)
			},
			expected: false,
		},
		{
			name: "Not equal - true",
			setup: func(bc *bytecode.Bytecode) {
				bc.AddFunction("main", 0)
				idx1 := bc.AddConstant(bytecode.NewIntValue(42))
				idx2 := bc.AddConstant(bytecode.NewIntValue(43))
				bc.AddInstruction(bytecode.OpPush, int64(idx1))
				bc.AddInstruction(bytecode.OpPush, int64(idx2))
				bc.AddInstruction(bytecode.OpNe, 0)
				bc.AddInstruction(bytecode.OpReturnValue, 0)
				bc.AddInstruction(bytecode.OpHalt, 0)
			},
			expected: true,
		},
		{
			name: "Less than - true",
			setup: func(bc *bytecode.Bytecode) {
				bc.AddFunction("main", 0)
				idx1 := bc.AddConstant(bytecode.NewIntValue(10))
				idx2 := bc.AddConstant(bytecode.NewIntValue(20))
				bc.AddInstruction(bytecode.OpPush, int64(idx1))
				bc.AddInstruction(bytecode.OpPush, int64(idx2))
				bc.AddInstruction(bytecode.OpLt, 0)
				bc.AddInstruction(bytecode.OpReturnValue, 0)
				bc.AddInstruction(bytecode.OpHalt, 0)
			},
			expected: true,
		},
		{
			name: "Greater than - true",
			setup: func(bc *bytecode.Bytecode) {
				bc.AddFunction("main", 0)
				idx1 := bc.AddConstant(bytecode.NewIntValue(30))
				idx2 := bc.AddConstant(bytecode.NewIntValue(20))
				bc.AddInstruction(bytecode.OpPush, int64(idx1))
				bc.AddInstruction(bytecode.OpPush, int64(idx2))
				bc.AddInstruction(bytecode.OpGt, 0)
				bc.AddInstruction(bytecode.OpReturnValue, 0)
				bc.AddInstruction(bytecode.OpHalt, 0)
			},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bc := bytecode.NewBytecode()
			tt.setup(bc)
			
			vm := NewVM(bc)
			result, err := vm.Run()
			
			if err != nil {
				t.Fatalf("VM execution failed: %v", err)
			}
			
			if result.Type != bytecode.ValueTypeBool {
				t.Fatalf("Expected bool result, got %v", result.Type)
			}
			
			if result.Bool != tt.expected {
				t.Errorf("Expected %v, got %v", tt.expected, result.Bool)
			}
		})
	}
}

func TestVMLogical(t *testing.T) {
	tests := []struct {
		name     string
		setup    func(*bytecode.Bytecode)
		expected bool
	}{
		{
			name: "AND - true",
			setup: func(bc *bytecode.Bytecode) {
				bc.AddFunction("main", 0)
				idxTrue := bc.AddConstant(bytecode.NewBoolValue(true))
				bc.AddInstruction(bytecode.OpPush, int64(idxTrue))
				bc.AddInstruction(bytecode.OpPush, int64(idxTrue))
				bc.AddInstruction(bytecode.OpAnd, 0)
				bc.AddInstruction(bytecode.OpReturnValue, 0)
				bc.AddInstruction(bytecode.OpHalt, 0)
			},
			expected: true,
		},
		{
			name: "AND - false",
			setup: func(bc *bytecode.Bytecode) {
				bc.AddFunction("main", 0)
				idxTrue := bc.AddConstant(bytecode.NewBoolValue(true))
				idxFalse := bc.AddConstant(bytecode.NewBoolValue(false))
				bc.AddInstruction(bytecode.OpPush, int64(idxTrue))
				bc.AddInstruction(bytecode.OpPush, int64(idxFalse))
				bc.AddInstruction(bytecode.OpAnd, 0)
				bc.AddInstruction(bytecode.OpReturnValue, 0)
				bc.AddInstruction(bytecode.OpHalt, 0)
			},
			expected: false,
		},
		{
			name: "OR - true",
			setup: func(bc *bytecode.Bytecode) {
				bc.AddFunction("main", 0)
				idxTrue := bc.AddConstant(bytecode.NewBoolValue(true))
				idxFalse := bc.AddConstant(bytecode.NewBoolValue(false))
				bc.AddInstruction(bytecode.OpPush, int64(idxTrue))
				bc.AddInstruction(bytecode.OpPush, int64(idxFalse))
				bc.AddInstruction(bytecode.OpOr, 0)
				bc.AddInstruction(bytecode.OpReturnValue, 0)
				bc.AddInstruction(bytecode.OpHalt, 0)
			},
			expected: true,
		},
		{
			name: "NOT - true",
			setup: func(bc *bytecode.Bytecode) {
				bc.AddFunction("main", 0)
				idxFalse := bc.AddConstant(bytecode.NewBoolValue(false))
				bc.AddInstruction(bytecode.OpPush, int64(idxFalse))
				bc.AddInstruction(bytecode.OpNot, 0)
				bc.AddInstruction(bytecode.OpReturnValue, 0)
				bc.AddInstruction(bytecode.OpHalt, 0)
			},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bc := bytecode.NewBytecode()
			tt.setup(bc)
			
			vm := NewVM(bc)
			result, err := vm.Run()
			
			if err != nil {
				t.Fatalf("VM execution failed: %v", err)
			}
			
			if result.Type != bytecode.ValueTypeBool {
				t.Fatalf("Expected bool result, got %v", result.Type)
			}
			
			if result.Bool != tt.expected {
				t.Errorf("Expected %v, got %v", tt.expected, result.Bool)
			}
		})
	}
}

func TestVMLocalVariables(t *testing.T) {
	// Test storing and loading local variables
	bc := bytecode.NewBytecode()
	bc.AddFunction("main", 0)
	
	// Store 42 to local variable 0
	idx := bc.AddConstant(bytecode.NewIntValue(42))
	bc.AddInstruction(bytecode.OpPush, int64(idx))
	bc.AddInstruction(bytecode.OpStoreLocal, 0)
	
	// Load local variable 0
	bc.AddInstruction(bytecode.OpLoadLocal, 0)
	bc.AddInstruction(bytecode.OpReturnValue, 0)
	bc.AddInstruction(bytecode.OpHalt, 0)
	
	vm := NewVM(bc)
	result, err := vm.Run()
	
	if err != nil {
		t.Fatalf("VM execution failed: %v", err)
	}
	
	if result.Type != bytecode.ValueTypeInt || result.Int != 42 {
		t.Errorf("Expected 42, got %v", result)
	}
}

func TestVMJump(t *testing.T) {
	// Test unconditional jump
	bc := bytecode.NewBytecode()
	bc.AddFunction("main", 0)
	
	// Jump over the first value
	bc.AddInstruction(bytecode.OpJump, 3)
	
	// This value should be skipped
	idx1 := bc.AddConstant(bytecode.NewIntValue(99))
	bc.AddInstruction(bytecode.OpPush, int64(idx1))
	bc.AddInstruction(bytecode.OpReturnValue, 0)
	
	// This is where we jump to
	idx2 := bc.AddConstant(bytecode.NewIntValue(42))
	bc.AddInstruction(bytecode.OpPush, int64(idx2))
	bc.AddInstruction(bytecode.OpReturnValue, 0)
	bc.AddInstruction(bytecode.OpHalt, 0)
	
	vm := NewVM(bc)
	result, err := vm.Run()
	
	if err != nil {
		t.Fatalf("VM execution failed: %v", err)
	}
	
	if result.Type != bytecode.ValueTypeInt || result.Int != 42 {
		t.Errorf("Expected 42, got %v", result)
	}
}

func TestVMJumpIfFalse(t *testing.T) {
	tests := []struct {
		name      string
		condition bool
		expected  int64
	}{
		{"Jump when false", false, 99},
		{"Don't jump when true", true, 42},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bc := bytecode.NewBytecode()
			bc.AddFunction("main", 0)
			
			// Push condition
			idxCond := bc.AddConstant(bytecode.NewBoolValue(tt.condition))
			bc.AddInstruction(bytecode.OpPush, int64(idxCond))
			
			// Jump if false to else block
			jumpPos := bc.CurrentPosition()
			bc.AddInstruction(bytecode.OpJumpIfFalse, 0) // Will patch later
			
			// THEN block: executes if condition is true
			idx1 := bc.AddConstant(bytecode.NewIntValue(42))
			bc.AddInstruction(bytecode.OpPush, int64(idx1))
			
			// Jump over else block
			jumpOverElse := bc.CurrentPosition()
			bc.AddInstruction(bytecode.OpJump, 0) // Will patch later
			
			// ELSE block: executes if condition is false
			elseStart := bc.CurrentPosition()
			idx2 := bc.AddConstant(bytecode.NewIntValue(99))
			bc.AddInstruction(bytecode.OpPush, int64(idx2))
			
			// After if-else
			afterIf := bc.CurrentPosition()
			
			// Patch jumps
			bc.PatchJump(jumpPos, elseStart)
			bc.PatchJump(jumpOverElse, afterIf)
			
			// Return the result
			bc.AddInstruction(bytecode.OpReturnValue, 0)
			bc.AddInstruction(bytecode.OpHalt, 0)
			
			vm := NewVM(bc)
			result, err := vm.Run()
			
			if err != nil {
				t.Fatalf("VM execution failed: %v", err)
			}
			
			if result.Type != bytecode.ValueTypeInt || result.Int != tt.expected {
				t.Errorf("Expected %d, got %v", tt.expected, result)
			}
		})
	}
}

func TestVMFunctionCall(t *testing.T) {
	// Test function call with arguments
	bc := bytecode.NewBytecode()
	
	// Helper function at position 0: add(a, b) -> a + b
	bc.AddFunction("add", 0)
	bc.AddInstruction(bytecode.OpLoadLocal, 0) // Load param a
	bc.AddInstruction(bytecode.OpLoadLocal, 1) // Load param b
	bc.AddInstruction(bytecode.OpAdd, 0)
	bc.AddInstruction(bytecode.OpReturnValue, 0)
	
	// Main function
	mainStart := bc.CurrentPosition()
	bc.AddFunction("main", mainStart)
	
	// Push arguments for add(10, 20)
	idx1 := bc.AddConstant(bytecode.NewIntValue(10))
	idx2 := bc.AddConstant(bytecode.NewIntValue(20))
	bc.AddInstruction(bytecode.OpPush, int64(idx1))
	bc.AddInstruction(bytecode.OpPush, int64(idx2))
	
	// Call add function
	bc.AddInstructionWithArgCount(bytecode.OpCall, 0, 2)
	
	// Return the result
	bc.AddInstruction(bytecode.OpReturnValue, 0)
	bc.AddInstruction(bytecode.OpHalt, 0)
	
	vm := NewVM(bc)
	result, err := vm.Run()
	
	if err != nil {
		t.Fatalf("VM execution failed: %v", err)
	}
	
	if result.Type != bytecode.ValueTypeInt || result.Int != 30 {
		t.Errorf("Expected 30, got %v", result)
	}
}

func TestVMDivisionByZero(t *testing.T) {
	bc := bytecode.NewBytecode()
	bc.AddFunction("main", 0)
	
	idx1 := bc.AddConstant(bytecode.NewIntValue(42))
	idx2 := bc.AddConstant(bytecode.NewIntValue(0))
	bc.AddInstruction(bytecode.OpPush, int64(idx1))
	bc.AddInstruction(bytecode.OpPush, int64(idx2))
	bc.AddInstruction(bytecode.OpDiv, 0)
	bc.AddInstruction(bytecode.OpReturnValue, 0)
	bc.AddInstruction(bytecode.OpHalt, 0)
	
	vm := NewVM(bc)
	_, err := vm.Run()
	
	if err == nil {
		t.Error("Expected division by zero error")
	}
}

func TestVMNoMainFunction(t *testing.T) {
	bc := bytecode.NewBytecode()
	// Don't add a main function
	bc.AddInstruction(bytecode.OpHalt, 0)
	
	vm := NewVM(bc)
	_, err := vm.Run()
	
	if err == nil {
		t.Error("Expected error for missing main function")
	}
}
