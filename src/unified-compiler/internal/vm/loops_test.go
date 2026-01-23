package vm

import (
	"testing"
	"unified-compiler/internal/bytecode"
)

func TestVMWhileLoop(t *testing.T) {
	// while loop that counts from 0 to 5
	bc := bytecode.NewBytecode()
	bc.AddFunction("main", 0)

	// let i = 0 (local 0)
	idx0 := bc.AddConstant(bytecode.NewIntValue(0))
	bc.AddInstruction(bytecode.OpPush, int64(idx0))
	bc.AddInstruction(bytecode.OpStoreLocal, 0)

	// Loop start
	loopStart := bc.CurrentPosition()

	// Condition: i < 5
	bc.AddInstruction(bytecode.OpLoadLocal, 0)
	idx5 := bc.AddConstant(bytecode.NewIntValue(5))
	bc.AddInstruction(bytecode.OpPush, int64(idx5))
	bc.AddInstruction(bytecode.OpLt, 0)

	// Jump if false
	jumpIfFalse := bc.CurrentPosition()
	bc.AddInstruction(bytecode.OpJumpIfFalse, 0) // Placeholder

	// Body: i = i + 1
	bc.AddInstruction(bytecode.OpLoadLocal, 0)
	idx1 := bc.AddConstant(bytecode.NewIntValue(1))
	bc.AddInstruction(bytecode.OpPush, int64(idx1))
	bc.AddInstruction(bytecode.OpAdd, 0)
	bc.AddInstruction(bytecode.OpStoreLocal, 0)

	// Jump back to loop start
	bc.AddInstruction(bytecode.OpJump, int64(loopStart))

	// After loop
	afterLoop := bc.CurrentPosition()
	bc.PatchJump(jumpIfFalse, afterLoop)

	// Return i
	bc.AddInstruction(bytecode.OpLoadLocal, 0)
	bc.AddInstruction(bytecode.OpReturnValue, 0)
	bc.AddInstruction(bytecode.OpHalt, 0)

	vm := NewVM(bc)
	result, err := vm.Run()

	if err != nil {
		t.Fatalf("VM execution failed: %v", err)
	}

	if result.Type != bytecode.ValueTypeInt || result.Int != 5 {
		t.Errorf("Expected 5, got %v", result)
	}
}

func TestVMForLoop(t *testing.T) {
	// for i in 0..5 (sum all values)
	bc := bytecode.NewBytecode()
	bc.AddFunction("main", 0)

	// let sum = 0 (local 0)
	idx0 := bc.AddConstant(bytecode.NewIntValue(0))
	bc.AddInstruction(bytecode.OpPush, int64(idx0))
	bc.AddInstruction(bytecode.OpStoreLocal, 0)

	// let i = 0 (local 1 - iterator)
	bc.AddInstruction(bytecode.OpPush, int64(idx0))
	bc.AddInstruction(bytecode.OpStoreLocal, 1)

	// let end = 5 (local 2)
	idx5 := bc.AddConstant(bytecode.NewIntValue(5))
	bc.AddInstruction(bytecode.OpPush, int64(idx5))
	bc.AddInstruction(bytecode.OpStoreLocal, 2)

	// Loop start
	loopStart := bc.CurrentPosition()

	// Condition: i < end
	bc.AddInstruction(bytecode.OpLoadLocal, 1)
	bc.AddInstruction(bytecode.OpLoadLocal, 2)
	bc.AddInstruction(bytecode.OpLt, 0)

	// Jump if false
	jumpIfFalse := bc.CurrentPosition()
	bc.AddInstruction(bytecode.OpJumpIfFalse, 0) // Placeholder

	// Body: sum = sum + i
	bc.AddInstruction(bytecode.OpLoadLocal, 0)
	bc.AddInstruction(bytecode.OpLoadLocal, 1)
	bc.AddInstruction(bytecode.OpAdd, 0)
	bc.AddInstruction(bytecode.OpStoreLocal, 0)

	// Increment: i = i + 1
	bc.AddInstruction(bytecode.OpLoadLocal, 1)
	idx1 := bc.AddConstant(bytecode.NewIntValue(1))
	bc.AddInstruction(bytecode.OpPush, int64(idx1))
	bc.AddInstruction(bytecode.OpAdd, 0)
	bc.AddInstruction(bytecode.OpStoreLocal, 1)

	// Jump back to loop start
	bc.AddInstruction(bytecode.OpJump, int64(loopStart))

	// After loop
	afterLoop := bc.CurrentPosition()
	bc.PatchJump(jumpIfFalse, afterLoop)

	// Return sum
	bc.AddInstruction(bytecode.OpLoadLocal, 0)
	bc.AddInstruction(bytecode.OpReturnValue, 0)
	bc.AddInstruction(bytecode.OpHalt, 0)

	vm := NewVM(bc)
	result, err := vm.Run()

	if err != nil {
		t.Fatalf("VM execution failed: %v", err)
	}

	// Sum of 0..4 = 0 + 1 + 2 + 3 + 4 = 10
	if result.Type != bytecode.ValueTypeInt || result.Int != 10 {
		t.Errorf("Expected 10, got %v", result)
	}
}

func TestVMInfiniteLoopWithBreak(t *testing.T) {
	// loop { i = i + 1; if i >= 5 { break } }
	bc := bytecode.NewBytecode()
	bc.AddFunction("main", 0)

	// let i = 0 (local 0)
	idx0 := bc.AddConstant(bytecode.NewIntValue(0))
	bc.AddInstruction(bytecode.OpPush, int64(idx0))
	bc.AddInstruction(bytecode.OpStoreLocal, 0)

	// Loop start
	loopStart := bc.CurrentPosition()

	// i = i + 1
	bc.AddInstruction(bytecode.OpLoadLocal, 0)
	idx1 := bc.AddConstant(bytecode.NewIntValue(1))
	bc.AddInstruction(bytecode.OpPush, int64(idx1))
	bc.AddInstruction(bytecode.OpAdd, 0)
	bc.AddInstruction(bytecode.OpStoreLocal, 0)

	// Condition: i >= 5
	bc.AddInstruction(bytecode.OpLoadLocal, 0)
	idx5 := bc.AddConstant(bytecode.NewIntValue(5))
	bc.AddInstruction(bytecode.OpPush, int64(idx5))
	bc.AddInstruction(bytecode.OpGe, 0)

	// If true, break (jump to after loop)
	breakJump := bc.CurrentPosition()
	bc.AddInstruction(bytecode.OpJumpIfTrue, 0) // Placeholder

	// Jump back to loop start
	bc.AddInstruction(bytecode.OpJump, int64(loopStart))

	// After loop
	afterLoop := bc.CurrentPosition()
	bc.PatchJump(breakJump, afterLoop)

	// Return i
	bc.AddInstruction(bytecode.OpLoadLocal, 0)
	bc.AddInstruction(bytecode.OpReturnValue, 0)
	bc.AddInstruction(bytecode.OpHalt, 0)

	vm := NewVM(bc)
	result, err := vm.Run()

	if err != nil {
		t.Fatalf("VM execution failed: %v", err)
	}

	if result.Type != bytecode.ValueTypeInt || result.Int != 5 {
		t.Errorf("Expected 5, got %v", result)
	}
}

func TestVMWhileLoopWithContinue(t *testing.T) {
	// sum = 0; i = 0
	// while i < 10 { i = i + 1; if i == 5 { continue }; sum = sum + i }
	// This should sum 1,2,3,4,6,7,8,9,10 (skipping 5) = 55 - 5 = 50
	bc := bytecode.NewBytecode()
	bc.AddFunction("main", 0)

	// let sum = 0 (local 0)
	idx0 := bc.AddConstant(bytecode.NewIntValue(0))
	bc.AddInstruction(bytecode.OpPush, int64(idx0))
	bc.AddInstruction(bytecode.OpStoreLocal, 0)

	// let i = 0 (local 1)
	bc.AddInstruction(bytecode.OpPush, int64(idx0))
	bc.AddInstruction(bytecode.OpStoreLocal, 1)

	// Loop start
	loopStart := bc.CurrentPosition()

	// Condition: i < 10
	bc.AddInstruction(bytecode.OpLoadLocal, 1)
	idx10 := bc.AddConstant(bytecode.NewIntValue(10))
	bc.AddInstruction(bytecode.OpPush, int64(idx10))
	bc.AddInstruction(bytecode.OpLt, 0)

	// Jump if false (exit loop)
	jumpIfFalse := bc.CurrentPosition()
	bc.AddInstruction(bytecode.OpJumpIfFalse, 0) // Placeholder

	// i = i + 1
	bc.AddInstruction(bytecode.OpLoadLocal, 1)
	idx1 := bc.AddConstant(bytecode.NewIntValue(1))
	bc.AddInstruction(bytecode.OpPush, int64(idx1))
	bc.AddInstruction(bytecode.OpAdd, 0)
	bc.AddInstruction(bytecode.OpStoreLocal, 1)

	// if i == 5 { continue }
	bc.AddInstruction(bytecode.OpLoadLocal, 1)
	idx5 := bc.AddConstant(bytecode.NewIntValue(5))
	bc.AddInstruction(bytecode.OpPush, int64(idx5))
	bc.AddInstruction(bytecode.OpEq, 0)

	// If true, continue (jump to loop start)
	bc.AddInstruction(bytecode.OpJumpIfTrue, int64(loopStart))

	// sum = sum + i
	bc.AddInstruction(bytecode.OpLoadLocal, 0)
	bc.AddInstruction(bytecode.OpLoadLocal, 1)
	bc.AddInstruction(bytecode.OpAdd, 0)
	bc.AddInstruction(bytecode.OpStoreLocal, 0)

	// Jump back to loop start
	bc.AddInstruction(bytecode.OpJump, int64(loopStart))

	// After loop
	afterLoop := bc.CurrentPosition()
	bc.PatchJump(jumpIfFalse, afterLoop)

	// Return sum
	bc.AddInstruction(bytecode.OpLoadLocal, 0)
	bc.AddInstruction(bytecode.OpReturnValue, 0)
	bc.AddInstruction(bytecode.OpHalt, 0)

	vm := NewVM(bc)
	result, err := vm.Run()

	if err != nil {
		t.Fatalf("VM execution failed: %v", err)
	}

	// Sum of 1..10 except 5 = 55 - 5 = 50
	if result.Type != bytecode.ValueTypeInt || result.Int != 50 {
		t.Errorf("Expected 50, got %v", result)
	}
}

func TestVMNestedLoops(t *testing.T) {
	// Nested loop: for i in 0..3 { for j in 0..3 { sum = sum + 1 } }
	// Should execute inner loop 3*3 = 9 times
	bc := bytecode.NewBytecode()
	bc.AddFunction("main", 0)

	// let sum = 0 (local 0)
	idx0 := bc.AddConstant(bytecode.NewIntValue(0))
	bc.AddInstruction(bytecode.OpPush, int64(idx0))
	bc.AddInstruction(bytecode.OpStoreLocal, 0)

	// Outer loop: i (local 1)
	bc.AddInstruction(bytecode.OpPush, int64(idx0))
	bc.AddInstruction(bytecode.OpStoreLocal, 1)

	// end_i = 3 (local 2)
	idx3 := bc.AddConstant(bytecode.NewIntValue(3))
	bc.AddInstruction(bytecode.OpPush, int64(idx3))
	bc.AddInstruction(bytecode.OpStoreLocal, 2)

	// Outer loop start
	outerLoopStart := bc.CurrentPosition()

	// Condition: i < 3
	bc.AddInstruction(bytecode.OpLoadLocal, 1)
	bc.AddInstruction(bytecode.OpLoadLocal, 2)
	bc.AddInstruction(bytecode.OpLt, 0)

	// Jump if false
	outerJumpIfFalse := bc.CurrentPosition()
	bc.AddInstruction(bytecode.OpJumpIfFalse, 0) // Placeholder

	// Inner loop: j (local 3)
	bc.AddInstruction(bytecode.OpPush, int64(idx0))
	bc.AddInstruction(bytecode.OpStoreLocal, 3)

	// end_j = 3 (local 4)
	bc.AddInstruction(bytecode.OpPush, int64(idx3))
	bc.AddInstruction(bytecode.OpStoreLocal, 4)

	// Inner loop start
	innerLoopStart := bc.CurrentPosition()

	// Condition: j < 3
	bc.AddInstruction(bytecode.OpLoadLocal, 3)
	bc.AddInstruction(bytecode.OpLoadLocal, 4)
	bc.AddInstruction(bytecode.OpLt, 0)

	// Jump if false
	innerJumpIfFalse := bc.CurrentPosition()
	bc.AddInstruction(bytecode.OpJumpIfFalse, 0) // Placeholder

	// sum = sum + 1
	bc.AddInstruction(bytecode.OpLoadLocal, 0)
	idx1 := bc.AddConstant(bytecode.NewIntValue(1))
	bc.AddInstruction(bytecode.OpPush, int64(idx1))
	bc.AddInstruction(bytecode.OpAdd, 0)
	bc.AddInstruction(bytecode.OpStoreLocal, 0)

	// j = j + 1
	bc.AddInstruction(bytecode.OpLoadLocal, 3)
	bc.AddInstruction(bytecode.OpPush, int64(idx1))
	bc.AddInstruction(bytecode.OpAdd, 0)
	bc.AddInstruction(bytecode.OpStoreLocal, 3)

	// Jump back to inner loop start
	bc.AddInstruction(bytecode.OpJump, int64(innerLoopStart))

	// After inner loop
	afterInnerLoop := bc.CurrentPosition()
	bc.PatchJump(innerJumpIfFalse, afterInnerLoop)

	// i = i + 1
	bc.AddInstruction(bytecode.OpLoadLocal, 1)
	bc.AddInstruction(bytecode.OpPush, int64(idx1))
	bc.AddInstruction(bytecode.OpAdd, 0)
	bc.AddInstruction(bytecode.OpStoreLocal, 1)

	// Jump back to outer loop start
	bc.AddInstruction(bytecode.OpJump, int64(outerLoopStart))

	// After outer loop
	afterOuterLoop := bc.CurrentPosition()
	bc.PatchJump(outerJumpIfFalse, afterOuterLoop)

	// Return sum
	bc.AddInstruction(bytecode.OpLoadLocal, 0)
	bc.AddInstruction(bytecode.OpReturnValue, 0)
	bc.AddInstruction(bytecode.OpHalt, 0)

	vm := NewVM(bc)
	result, err := vm.Run()

	if err != nil {
		t.Fatalf("VM execution failed: %v", err)
	}

	if result.Type != bytecode.ValueTypeInt || result.Int != 9 {
		t.Errorf("Expected 9 (3x3 iterations), got %v", result)
	}
}

func TestVMForLoopInclusive(t *testing.T) {
	// for i in 0..=4 (sum all values including 4)
	// Sum = 0 + 1 + 2 + 3 + 4 = 10
	bc := bytecode.NewBytecode()
	bc.AddFunction("main", 0)

	// let sum = 0 (local 0)
	idx0 := bc.AddConstant(bytecode.NewIntValue(0))
	bc.AddInstruction(bytecode.OpPush, int64(idx0))
	bc.AddInstruction(bytecode.OpStoreLocal, 0)

	// let i = 0 (local 1 - iterator)
	bc.AddInstruction(bytecode.OpPush, int64(idx0))
	bc.AddInstruction(bytecode.OpStoreLocal, 1)

	// let end = 4 (local 2)
	idx4 := bc.AddConstant(bytecode.NewIntValue(4))
	bc.AddInstruction(bytecode.OpPush, int64(idx4))
	bc.AddInstruction(bytecode.OpStoreLocal, 2)

	// Loop start
	loopStart := bc.CurrentPosition()

	// Condition: i <= end (inclusive)
	bc.AddInstruction(bytecode.OpLoadLocal, 1)
	bc.AddInstruction(bytecode.OpLoadLocal, 2)
	bc.AddInstruction(bytecode.OpLe, 0)

	// Jump if false
	jumpIfFalse := bc.CurrentPosition()
	bc.AddInstruction(bytecode.OpJumpIfFalse, 0) // Placeholder

	// Body: sum = sum + i
	bc.AddInstruction(bytecode.OpLoadLocal, 0)
	bc.AddInstruction(bytecode.OpLoadLocal, 1)
	bc.AddInstruction(bytecode.OpAdd, 0)
	bc.AddInstruction(bytecode.OpStoreLocal, 0)

	// Increment: i = i + 1
	bc.AddInstruction(bytecode.OpLoadLocal, 1)
	idx1 := bc.AddConstant(bytecode.NewIntValue(1))
	bc.AddInstruction(bytecode.OpPush, int64(idx1))
	bc.AddInstruction(bytecode.OpAdd, 0)
	bc.AddInstruction(bytecode.OpStoreLocal, 1)

	// Jump back to loop start
	bc.AddInstruction(bytecode.OpJump, int64(loopStart))

	// After loop
	afterLoop := bc.CurrentPosition()
	bc.PatchJump(jumpIfFalse, afterLoop)

	// Return sum
	bc.AddInstruction(bytecode.OpLoadLocal, 0)
	bc.AddInstruction(bytecode.OpReturnValue, 0)
	bc.AddInstruction(bytecode.OpHalt, 0)

	vm := NewVM(bc)
	result, err := vm.Run()

	if err != nil {
		t.Fatalf("VM execution failed: %v", err)
	}

	// Sum of 0..=4 = 0 + 1 + 2 + 3 + 4 = 10
	if result.Type != bytecode.ValueTypeInt || result.Int != 10 {
		t.Errorf("Expected 10, got %v", result)
	}
}
