package vm

import (
"strings"
"testing"
"unified-compiler/internal/bytecode"
)

// Test 1: Array creation
func TestVMArrayCreation(t *testing.T) {
bc := bytecode.NewBytecode()
bc.AddFunction("main", 0)

// Create an array with 3 elements: [1, 2, 3]
idx1 := bc.AddConstant(bytecode.NewIntValue(1))
idx2 := bc.AddConstant(bytecode.NewIntValue(2))
idx3 := bc.AddConstant(bytecode.NewIntValue(3))

bc.AddInstruction(bytecode.OpPush, int64(idx1))
bc.AddInstruction(bytecode.OpPush, int64(idx2))
bc.AddInstruction(bytecode.OpPush, int64(idx3))
bc.AddInstruction(bytecode.OpAllocArray, 3)
bc.AddInstruction(bytecode.OpHalt, 0)

vm := NewVM(bc)
result, err := vm.Run()
if err != nil {
t.Fatalf("VM execution failed: %v", err)
}

if result.Type != bytecode.ValueTypeArray {
t.Errorf("Expected array type, got %v", result.Type)
}
if result.Array.Length != 3 {
t.Errorf("Expected array length 3, got %d", result.Array.Length)
}
}

// Test 2: Array indexing (read)
func TestVMArrayIndexRead(t *testing.T) {
bc := bytecode.NewBytecode()
bc.AddFunction("main", 0)

idx1 := bc.AddConstant(bytecode.NewIntValue(10))
idx2 := bc.AddConstant(bytecode.NewIntValue(20))
idx3 := bc.AddConstant(bytecode.NewIntValue(30))

bc.AddInstruction(bytecode.OpPush, int64(idx1))
bc.AddInstruction(bytecode.OpPush, int64(idx2))
bc.AddInstruction(bytecode.OpPush, int64(idx3))
bc.AddInstruction(bytecode.OpAllocArray, 3)
bc.AddInstruction(bytecode.OpDup, 0)

idxPos := bc.AddConstant(bytecode.NewIntValue(1))
bc.AddInstruction(bytecode.OpPush, int64(idxPos))
bc.AddInstruction(bytecode.OpLoadArray, 0)
bc.AddInstruction(bytecode.OpHalt, 0)

vm := NewVM(bc)
result, err := vm.Run()
if err != nil {
t.Fatalf("VM execution failed: %v", err)
}

if result.Type != bytecode.ValueTypeInt {
t.Errorf("Expected int type, got %v", result.Type)
}
if result.Int != 20 {
t.Errorf("Expected value 20, got %d", result.Int)
}
}

// Test 3: Array bounds checking (index too large)
func TestVMArrayBoundsCheckTooLarge(t *testing.T) {
bc := bytecode.NewBytecode()
bc.AddFunction("main", 0)

idx1 := bc.AddConstant(bytecode.NewIntValue(1))
idx2 := bc.AddConstant(bytecode.NewIntValue(2))
idx3 := bc.AddConstant(bytecode.NewIntValue(3))

bc.AddInstruction(bytecode.OpPush, int64(idx1))
bc.AddInstruction(bytecode.OpPush, int64(idx2))
bc.AddInstruction(bytecode.OpPush, int64(idx3))
bc.AddInstruction(bytecode.OpAllocArray, 3)

idxPos := bc.AddConstant(bytecode.NewIntValue(5))
bc.AddInstruction(bytecode.OpPush, int64(idxPos))
bc.AddInstruction(bytecode.OpLoadArray, 0)
bc.AddInstruction(bytecode.OpHalt, 0)

vm := NewVM(bc)
_, err := vm.Run()
if err == nil {
t.Fatal("Expected bounds check error, but got none")
}
if !strings.Contains(err.Error(), "array index out of bounds") {
t.Errorf("Expected bounds check error, got: %v", err)
}
}

// Test 4: Array bounds checking (negative index)
func TestVMArrayBoundsCheckNegative(t *testing.T) {
bc := bytecode.NewBytecode()
bc.AddFunction("main", 0)

idx1 := bc.AddConstant(bytecode.NewIntValue(1))
idx2 := bc.AddConstant(bytecode.NewIntValue(2))
idx3 := bc.AddConstant(bytecode.NewIntValue(3))

bc.AddInstruction(bytecode.OpPush, int64(idx1))
bc.AddInstruction(bytecode.OpPush, int64(idx2))
bc.AddInstruction(bytecode.OpPush, int64(idx3))
bc.AddInstruction(bytecode.OpAllocArray, 3)

idxPos := bc.AddConstant(bytecode.NewIntValue(-1))
bc.AddInstruction(bytecode.OpPush, int64(idxPos))
bc.AddInstruction(bytecode.OpLoadArray, 0)
bc.AddInstruction(bytecode.OpHalt, 0)

vm := NewVM(bc)
_, err := vm.Run()
if err == nil {
t.Fatal("Expected bounds check error, but got none")
}
if !strings.Contains(err.Error(), "array index out of bounds") {
t.Errorf("Expected bounds check error, got: %v", err)
}
}

// Test 5: Array length operation
func TestVMArrayLength(t *testing.T) {
bc := bytecode.NewBytecode()
bc.AddFunction("main", 0)

for i := 1; i <= 5; i++ {
idx := bc.AddConstant(bytecode.NewIntValue(int64(i * 5)))
bc.AddInstruction(bytecode.OpPush, int64(idx))
}
bc.AddInstruction(bytecode.OpAllocArray, 5)
bc.AddInstruction(bytecode.OpArrayLen, 0)
bc.AddInstruction(bytecode.OpHalt, 0)

vm := NewVM(bc)
result, err := vm.Run()
if err != nil {
t.Fatalf("VM execution failed: %v", err)
}

if result.Type != bytecode.ValueTypeInt {
t.Errorf("Expected int type, got %v", result.Type)
}
if result.Int != 5 {
t.Errorf("Expected length 5, got %d", result.Int)
}
}

// Test 6: Empty array
func TestVMEmptyArray(t *testing.T) {
bc := bytecode.NewBytecode()
bc.AddFunction("main", 0)

bc.AddInstruction(bytecode.OpAllocArray, 0)
bc.AddInstruction(bytecode.OpArrayLen, 0)
bc.AddInstruction(bytecode.OpHalt, 0)

vm := NewVM(bc)
result, err := vm.Run()
if err != nil {
t.Fatalf("VM execution failed: %v", err)
}

if result.Type != bytecode.ValueTypeInt {
t.Errorf("Expected int type, got %v", result.Type)
}
if result.Int != 0 {
t.Errorf("Expected length 0, got %d", result.Int)
}
}

// Test 7: Array element assignment
func TestVMArrayIndexWrite(t *testing.T) {
bc := bytecode.NewBytecode()
bc.AddFunction("main", 0)

idx1 := bc.AddConstant(bytecode.NewIntValue(1))
idx2 := bc.AddConstant(bytecode.NewIntValue(2))
idx3 := bc.AddConstant(bytecode.NewIntValue(3))

bc.AddInstruction(bytecode.OpPush, int64(idx1))
bc.AddInstruction(bytecode.OpPush, int64(idx2))
bc.AddInstruction(bytecode.OpPush, int64(idx3))
bc.AddInstruction(bytecode.OpAllocArray, 3)
bc.AddInstruction(bytecode.OpStoreLocal, 0)

bc.AddInstruction(bytecode.OpLoadLocal, 0)
idxPos := bc.AddConstant(bytecode.NewIntValue(1))
bc.AddInstruction(bytecode.OpPush, int64(idxPos))
idx99 := bc.AddConstant(bytecode.NewIntValue(99))
bc.AddInstruction(bytecode.OpPush, int64(idx99))
bc.AddInstruction(bytecode.OpStoreArray, 0)

bc.AddInstruction(bytecode.OpLoadLocal, 0)
bc.AddInstruction(bytecode.OpPush, int64(idxPos))
bc.AddInstruction(bytecode.OpLoadArray, 0)
bc.AddInstruction(bytecode.OpHalt, 0)

vm := NewVM(bc)
result, err := vm.Run()
if err != nil {
t.Fatalf("VM execution failed: %v", err)
}

if result.Type != bytecode.ValueTypeInt {
t.Errorf("Expected int type, got %v", result.Type)
}
if result.Int != 99 {
t.Errorf("Expected value 99, got %d", result.Int)
}
}

// Test 8: Array boundary access
func TestVMArrayBoundaryAccess(t *testing.T) {
bc := bytecode.NewBytecode()
bc.AddFunction("main", 0)

idx1 := bc.AddConstant(bytecode.NewIntValue(100))
idx2 := bc.AddConstant(bytecode.NewIntValue(200))
idx3 := bc.AddConstant(bytecode.NewIntValue(300))

bc.AddInstruction(bytecode.OpPush, int64(idx1))
bc.AddInstruction(bytecode.OpPush, int64(idx2))
bc.AddInstruction(bytecode.OpPush, int64(idx3))
bc.AddInstruction(bytecode.OpAllocArray, 3)
bc.AddInstruction(bytecode.OpStoreLocal, 0)

bc.AddInstruction(bytecode.OpLoadLocal, 0)
idx0 := bc.AddConstant(bytecode.NewIntValue(0))
bc.AddInstruction(bytecode.OpPush, int64(idx0))
bc.AddInstruction(bytecode.OpLoadArray, 0)
bc.AddInstruction(bytecode.OpStoreLocal, 1)

bc.AddInstruction(bytecode.OpLoadLocal, 0)
idx2Pos := bc.AddConstant(bytecode.NewIntValue(2))
bc.AddInstruction(bytecode.OpPush, int64(idx2Pos))
bc.AddInstruction(bytecode.OpLoadArray, 0)
bc.AddInstruction(bytecode.OpStoreLocal, 2)

bc.AddInstruction(bytecode.OpLoadLocal, 1)
bc.AddInstruction(bytecode.OpLoadLocal, 2)
bc.AddInstruction(bytecode.OpAdd, 0)
bc.AddInstruction(bytecode.OpHalt, 0)

vm := NewVM(bc)
result, err := vm.Run()
if err != nil {
t.Fatalf("VM execution failed: %v", err)
}

if result.Type != bytecode.ValueTypeInt {
t.Errorf("Expected int type, got %v", result.Type)
}
if result.Int != 400 {
t.Errorf("Expected value 400, got %d", result.Int)
}
}

// Test 9: Array modification
func TestVMArrayModification(t *testing.T) {
bc := bytecode.NewBytecode()
bc.AddFunction("main", 0)

for i := 0; i < 3; i++ {
idx := bc.AddConstant(bytecode.NewIntValue(1))
bc.AddInstruction(bytecode.OpPush, int64(idx))
}
bc.AddInstruction(bytecode.OpAllocArray, 3)
bc.AddInstruction(bytecode.OpStoreLocal, 0)

for i := 0; i < 3; i++ {
bc.AddInstruction(bytecode.OpLoadLocal, 0)
idxPos := bc.AddConstant(bytecode.NewIntValue(int64(i)))
bc.AddInstruction(bytecode.OpPush, int64(idxPos))
valIdx := bc.AddConstant(bytecode.NewIntValue(int64((i + 1) * 2)))
bc.AddInstruction(bytecode.OpPush, int64(valIdx))
bc.AddInstruction(bytecode.OpStoreArray, 0)
}

bc.AddInstruction(bytecode.OpLoadLocal, 0)
idx0 := bc.AddConstant(bytecode.NewIntValue(0))
bc.AddInstruction(bytecode.OpPush, int64(idx0))
bc.AddInstruction(bytecode.OpLoadArray, 0)

bc.AddInstruction(bytecode.OpLoadLocal, 0)
idx1 := bc.AddConstant(bytecode.NewIntValue(1))
bc.AddInstruction(bytecode.OpPush, int64(idx1))
bc.AddInstruction(bytecode.OpLoadArray, 0)
bc.AddInstruction(bytecode.OpAdd, 0)

bc.AddInstruction(bytecode.OpLoadLocal, 0)
idx2 := bc.AddConstant(bytecode.NewIntValue(2))
bc.AddInstruction(bytecode.OpPush, int64(idx2))
bc.AddInstruction(bytecode.OpLoadArray, 0)
bc.AddInstruction(bytecode.OpAdd, 0)

bc.AddInstruction(bytecode.OpHalt, 0)

vm := NewVM(bc)
result, err := vm.Run()
if err != nil {
t.Fatalf("VM execution failed: %v", err)
}

if result.Type != bytecode.ValueTypeInt {
t.Errorf("Expected int type, got %v", result.Type)
}
if result.Int != 12 {
t.Errorf("Expected sum 12, got %d", result.Int)
}
}

// Test 10: Array operations
func TestVMArrayOperations(t *testing.T) {
bc := bytecode.NewBytecode()
bc.AddFunction("main", 0)

idx1 := bc.AddConstant(bytecode.NewIntValue(5))
idx2 := bc.AddConstant(bytecode.NewIntValue(10))
idx3 := bc.AddConstant(bytecode.NewIntValue(15))

bc.AddInstruction(bytecode.OpPush, int64(idx1))
bc.AddInstruction(bytecode.OpPush, int64(idx2))
bc.AddInstruction(bytecode.OpPush, int64(idx3))
bc.AddInstruction(bytecode.OpAllocArray, 3)
bc.AddInstruction(bytecode.OpStoreLocal, 0)

bc.AddInstruction(bytecode.OpLoadLocal, 0)
idx0 := bc.AddConstant(bytecode.NewIntValue(0))
bc.AddInstruction(bytecode.OpPush, int64(idx0))
bc.AddInstruction(bytecode.OpLoadArray, 0)
two := bc.AddConstant(bytecode.NewIntValue(2))
bc.AddInstruction(bytecode.OpPush, int64(two))
bc.AddInstruction(bytecode.OpMul, 0)

bc.AddInstruction(bytecode.OpLoadLocal, 0)
idx1Pos := bc.AddConstant(bytecode.NewIntValue(1))
bc.AddInstruction(bytecode.OpPush, int64(idx1Pos))
bc.AddInstruction(bytecode.OpLoadArray, 0)
bc.AddInstruction(bytecode.OpPush, int64(two))
bc.AddInstruction(bytecode.OpMul, 0)
bc.AddInstruction(bytecode.OpAdd, 0)

bc.AddInstruction(bytecode.OpLoadLocal, 0)
idx2Pos := bc.AddConstant(bytecode.NewIntValue(2))
bc.AddInstruction(bytecode.OpPush, int64(idx2Pos))
bc.AddInstruction(bytecode.OpLoadArray, 0)
bc.AddInstruction(bytecode.OpPush, int64(two))
bc.AddInstruction(bytecode.OpMul, 0)
bc.AddInstruction(bytecode.OpAdd, 0)

bc.AddInstruction(bytecode.OpHalt, 0)

vm := NewVM(bc)
result, err := vm.Run()
if err != nil {
t.Fatalf("VM execution failed: %v", err)
}

if result.Type != bytecode.ValueTypeInt {
t.Errorf("Expected int type, got %v", result.Type)
}
if result.Int != 60 {
t.Errorf("Expected sum 60, got %d", result.Int)
}
}
