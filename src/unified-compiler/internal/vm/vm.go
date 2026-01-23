package vm

import (
"fmt"
"unified-compiler/internal/bytecode"
)

// VM represents the virtual machine
type VM struct {
bytecode *bytecode.Bytecode
stack    *Stack
ip       int // Instruction pointer
callStack []CallFrame
globals  map[string]bytecode.Value
}

// CallFrame represents a function call frame
type CallFrame struct {
returnIP int // Return instruction pointer
locals   []bytecode.Value
basePtr  int // Base pointer for local variables
}

// NewVM creates a new virtual machine
func NewVM(bc *bytecode.Bytecode) *VM {
return &VM{
bytecode:  bc,
stack:     NewStack(),
ip:        0,
callStack: make([]CallFrame, 0),
globals:   make(map[string]bytecode.Value),
}
}

// Run executes the bytecode
func (vm *VM) Run() (bytecode.Value, error) {
// Find main function
mainIdx, ok := vm.bytecode.Functions["main"]
if !ok {
return bytecode.NewNullValue(), fmt.Errorf("no main function found")
}

// Jump to main
vm.ip = mainIdx

// Execute instructions
for vm.ip < len(vm.bytecode.Instructions) {
instruction := vm.bytecode.Instructions[vm.ip]

if err := vm.executeInstruction(instruction); err != nil {
return bytecode.NewNullValue(), fmt.Errorf("runtime error at instruction %d: %w", vm.ip, err)
}

// Check if we should halt
if instruction.Op == bytecode.OpHalt {
break
}
}

// Return value from stack if any
if vm.stack.Size() > 0 {
return vm.stack.Pop(), nil
}

return bytecode.NewIntValue(0), nil
}

// executeInstruction executes a single bytecode instruction
func (vm *VM) executeInstruction(inst bytecode.Instruction) error {
switch inst.Op {
case bytecode.OpPush:
// Push constant onto stack
if int(inst.Operand) >= len(vm.bytecode.Constants) {
return fmt.Errorf("constant index out of bounds: %d", inst.Operand)
}
val := vm.bytecode.Constants[inst.Operand]
vm.stack.Push(val)
vm.ip++

case bytecode.OpPop:
vm.stack.Pop()
vm.ip++

case bytecode.OpDup:
val := vm.stack.Peek()
vm.stack.Push(val)
vm.ip++

case bytecode.OpAdd:
right := vm.stack.Pop()
left := vm.stack.Pop()
result, err := vm.add(left, right)
if err != nil {
return err
}
vm.stack.Push(result)
vm.ip++

case bytecode.OpSub:
right := vm.stack.Pop()
left := vm.stack.Pop()
result, err := vm.sub(left, right)
if err != nil {
return err
}
vm.stack.Push(result)
vm.ip++

case bytecode.OpMul:
right := vm.stack.Pop()
left := vm.stack.Pop()
result, err := vm.mul(left, right)
if err != nil {
return err
}
vm.stack.Push(result)
vm.ip++

case bytecode.OpDiv:
right := vm.stack.Pop()
left := vm.stack.Pop()
result, err := vm.div(left, right)
if err != nil {
return err
}
vm.stack.Push(result)
vm.ip++

case bytecode.OpMod:
right := vm.stack.Pop()
left := vm.stack.Pop()
result, err := vm.mod(left, right)
if err != nil {
return err
}
vm.stack.Push(result)
vm.ip++

case bytecode.OpNeg:
val := vm.stack.Pop()
result, err := vm.neg(val)
if err != nil {
return err
}
vm.stack.Push(result)
vm.ip++

case bytecode.OpEq:
right := vm.stack.Pop()
left := vm.stack.Pop()
result := vm.eq(left, right)
vm.stack.Push(result)
vm.ip++

case bytecode.OpNe:
right := vm.stack.Pop()
left := vm.stack.Pop()
result := vm.ne(left, right)
vm.stack.Push(result)
vm.ip++

case bytecode.OpLt:
right := vm.stack.Pop()
left := vm.stack.Pop()
result, err := vm.lt(left, right)
if err != nil {
return err
}
vm.stack.Push(result)
vm.ip++

case bytecode.OpLe:
right := vm.stack.Pop()
left := vm.stack.Pop()
result, err := vm.le(left, right)
if err != nil {
return err
}
vm.stack.Push(result)
vm.ip++

case bytecode.OpGt:
right := vm.stack.Pop()
left := vm.stack.Pop()
result, err := vm.gt(left, right)
if err != nil {
return err
}
vm.stack.Push(result)
vm.ip++

case bytecode.OpGe:
right := vm.stack.Pop()
left := vm.stack.Pop()
result, err := vm.ge(left, right)
if err != nil {
return err
}
vm.stack.Push(result)
vm.ip++

case bytecode.OpAnd:
right := vm.stack.Pop()
left := vm.stack.Pop()
result := bytecode.NewBoolValue(left.IsTruthy() && right.IsTruthy())
vm.stack.Push(result)
vm.ip++

case bytecode.OpOr:
right := vm.stack.Pop()
left := vm.stack.Pop()
result := bytecode.NewBoolValue(left.IsTruthy() || right.IsTruthy())
vm.stack.Push(result)
vm.ip++

case bytecode.OpNot:
val := vm.stack.Pop()
result := bytecode.NewBoolValue(!val.IsTruthy())
vm.stack.Push(result)
vm.ip++

case bytecode.OpBitAnd:
right := vm.stack.Pop()
left := vm.stack.Pop()
if left.Type != bytecode.ValueTypeInt || right.Type != bytecode.ValueTypeInt {
return fmt.Errorf("bitwise AND requires integer operands")
}
result := bytecode.NewIntValue(left.Int & right.Int)
vm.stack.Push(result)
vm.ip++

case bytecode.OpBitOr:
right := vm.stack.Pop()
left := vm.stack.Pop()
if left.Type != bytecode.ValueTypeInt || right.Type != bytecode.ValueTypeInt {
return fmt.Errorf("bitwise OR requires integer operands")
}
result := bytecode.NewIntValue(left.Int | right.Int)
vm.stack.Push(result)
vm.ip++

case bytecode.OpBitXor:
right := vm.stack.Pop()
left := vm.stack.Pop()
if left.Type != bytecode.ValueTypeInt || right.Type != bytecode.ValueTypeInt {
return fmt.Errorf("bitwise XOR requires integer operands")
}
result := bytecode.NewIntValue(left.Int ^ right.Int)
vm.stack.Push(result)
vm.ip++

case bytecode.OpBitNot:
val := vm.stack.Pop()
if val.Type != bytecode.ValueTypeInt {
return fmt.Errorf("bitwise NOT requires integer operand")
}
result := bytecode.NewIntValue(^val.Int)
vm.stack.Push(result)
vm.ip++

case bytecode.OpLShift:
right := vm.stack.Pop()
left := vm.stack.Pop()
if left.Type != bytecode.ValueTypeInt || right.Type != bytecode.ValueTypeInt {
return fmt.Errorf("left shift requires integer operands")
}
result := bytecode.NewIntValue(left.Int << uint(right.Int))
vm.stack.Push(result)
vm.ip++

case bytecode.OpRShift:
right := vm.stack.Pop()
left := vm.stack.Pop()
if left.Type != bytecode.ValueTypeInt || right.Type != bytecode.ValueTypeInt {
return fmt.Errorf("right shift requires integer operands")
}
result := bytecode.NewIntValue(left.Int >> uint(right.Int))
vm.stack.Push(result)
vm.ip++

case bytecode.OpLoadLocal:
// Load local variable from current frame
if len(vm.callStack) == 0 {
return fmt.Errorf("no call frame for local variable access")
}
frame := vm.callStack[len(vm.callStack)-1]
if int(inst.Operand) >= len(frame.locals) {
return fmt.Errorf("local variable index out of bounds: %d", inst.Operand)
}
val := frame.locals[inst.Operand]
vm.stack.Push(val)
vm.ip++

case bytecode.OpStoreLocal:
// Store to local variable in current frame
val := vm.stack.Pop()
if len(vm.callStack) == 0 {
// Create initial frame for main function
vm.callStack = append(vm.callStack, CallFrame{
returnIP: -1,
locals:   make([]bytecode.Value, 100), // Preallocate space
})
}
frameIdx := len(vm.callStack) - 1
if int(inst.Operand) >= len(vm.callStack[frameIdx].locals) {
// Expand locals if needed
newLocals := make([]bytecode.Value, inst.Operand+1)
copy(newLocals, vm.callStack[frameIdx].locals)
vm.callStack[frameIdx].locals = newLocals
}
vm.callStack[frameIdx].locals[inst.Operand] = val
vm.ip++

case bytecode.OpJump:
vm.ip = int(inst.Operand)

case bytecode.OpJumpIfFalse:
val := vm.stack.Pop()
if !val.IsTruthy() {
vm.ip = int(inst.Operand)
} else {
vm.ip++
}

case bytecode.OpJumpIfTrue:
val := vm.stack.Pop()
if val.IsTruthy() {
vm.ip = int(inst.Operand)
} else {
vm.ip++
}

	case bytecode.OpCall:
		// Function call with arguments
		// inst.Operand is the function entry point
		// inst.ArgCount is the number of arguments
		// Arguments are on the stack (in order)
		
		// Pop arguments from stack (in reverse order since stack is LIFO)
		args := make([]bytecode.Value, inst.ArgCount)
		for i := inst.ArgCount - 1; i >= 0; i-- {
			args[i] = vm.stack.Pop()
		}
		
		// Create new call frame
		frame := CallFrame{
			returnIP: vm.ip + 1,
			locals:   make([]bytecode.Value, 100),
		}
		
		// Set arguments as local variables
		for i, arg := range args {
			frame.locals[i] = arg
		}
		
		vm.callStack = append(vm.callStack, frame)
		vm.ip = int(inst.Operand)

case bytecode.OpReturn:
// Return without value
if len(vm.callStack) > 0 {
frame := vm.callStack[len(vm.callStack)-1]
vm.callStack = vm.callStack[:len(vm.callStack)-1]
if frame.returnIP >= 0 {
vm.ip = frame.returnIP
} else {
// Returning from main
vm.ip = len(vm.bytecode.Instructions)
}
} else {
vm.ip++
}

case bytecode.OpReturnValue:
// Return with value on stack
returnValue := vm.stack.Pop()
if len(vm.callStack) > 0 {
frame := vm.callStack[len(vm.callStack)-1]
vm.callStack = vm.callStack[:len(vm.callStack)-1]
vm.stack.Push(returnValue)
if frame.returnIP >= 0 {
vm.ip = frame.returnIP
} else {
// Returning from main
vm.ip = len(vm.bytecode.Instructions)
}
} else {
vm.stack.Push(returnValue)
vm.ip++
}

case bytecode.OpHalt:
return nil

case bytecode.OpNop:
vm.ip++

case bytecode.OpAllocStruct:
// Allocate struct: expects field count as operand
// Stack: [field_name_0, field_value_0, field_name_1, field_value_1, ...] [type_name]
fieldCount := int(inst.Operand)
typeName := vm.stack.Pop()
if typeName.Type != bytecode.ValueTypeString {
return fmt.Errorf("expected string for struct type name")
}

// Pop field name/value pairs from stack (in reverse order)
fields := make(map[string]bytecode.Value)
for i := 0; i < fieldCount; i++ {
// Pop value then name (since we pushed name first, value second)
fieldValue := vm.stack.Pop()
fieldName := vm.stack.Pop()

if fieldName.Type != bytecode.ValueTypeString {
return fmt.Errorf("expected string for field name")
}

fields[fieldName.Str] = fieldValue
}

// Create struct value and push it
structVal := bytecode.NewStructValue(typeName.Str, fields)
vm.stack.Push(structVal)
vm.ip++

case bytecode.OpLoadField:
// Load field from struct
// Stack: [struct] [field_name]
fieldName := vm.stack.Pop()
structVal := vm.stack.Pop()

if structVal.Type != bytecode.ValueTypeStruct {
return fmt.Errorf("expected struct for field access")
}
if fieldName.Type != bytecode.ValueTypeString {
return fmt.Errorf("expected string for field name")
}

// Get field value
fieldValue, ok := structVal.Struct.Fields[fieldName.Str]
if !ok {
return fmt.Errorf("field not found: %s", fieldName.Str)
}

vm.stack.Push(fieldValue)
vm.ip++

case bytecode.OpStoreField:
// Store field to struct
// Stack: [struct] [field_name] [value]
value := vm.stack.Pop()
fieldName := vm.stack.Pop()
structVal := vm.stack.Pop()

if structVal.Type != bytecode.ValueTypeStruct {
return fmt.Errorf("expected struct for field assignment")
}
if fieldName.Type != bytecode.ValueTypeString {
return fmt.Errorf("expected string for field name")
}

// Store field value
structVal.Struct.Fields[fieldName.Str] = value
vm.stack.Push(structVal)
vm.ip++

default:
return fmt.Errorf("unknown opcode: %d", inst.Op)
}

return nil
}

// Arithmetic operations

func (vm *VM) add(left, right bytecode.Value) (bytecode.Value, error) {
if left.Type == bytecode.ValueTypeInt && right.Type == bytecode.ValueTypeInt {
return bytecode.NewIntValue(left.Int + right.Int), nil
}
if left.Type == bytecode.ValueTypeFloat || right.Type == bytecode.ValueTypeFloat {
leftFloat := vm.toFloat(left)
rightFloat := vm.toFloat(right)
return bytecode.NewFloatValue(leftFloat + rightFloat), nil
}
return bytecode.NewNullValue(), fmt.Errorf("cannot add %v and %v", left.Type, right.Type)
}

func (vm *VM) sub(left, right bytecode.Value) (bytecode.Value, error) {
if left.Type == bytecode.ValueTypeInt && right.Type == bytecode.ValueTypeInt {
return bytecode.NewIntValue(left.Int - right.Int), nil
}
if left.Type == bytecode.ValueTypeFloat || right.Type == bytecode.ValueTypeFloat {
leftFloat := vm.toFloat(left)
rightFloat := vm.toFloat(right)
return bytecode.NewFloatValue(leftFloat - rightFloat), nil
}
return bytecode.NewNullValue(), fmt.Errorf("cannot subtract %v and %v", left.Type, right.Type)
}

func (vm *VM) mul(left, right bytecode.Value) (bytecode.Value, error) {
if left.Type == bytecode.ValueTypeInt && right.Type == bytecode.ValueTypeInt {
return bytecode.NewIntValue(left.Int * right.Int), nil
}
if left.Type == bytecode.ValueTypeFloat || right.Type == bytecode.ValueTypeFloat {
leftFloat := vm.toFloat(left)
rightFloat := vm.toFloat(right)
return bytecode.NewFloatValue(leftFloat * rightFloat), nil
}
return bytecode.NewNullValue(), fmt.Errorf("cannot multiply %v and %v", left.Type, right.Type)
}

func (vm *VM) div(left, right bytecode.Value) (bytecode.Value, error) {
if left.Type == bytecode.ValueTypeInt && right.Type == bytecode.ValueTypeInt {
if right.Int == 0 {
return bytecode.NewNullValue(), fmt.Errorf("division by zero")
}
return bytecode.NewIntValue(left.Int / right.Int), nil
}
if left.Type == bytecode.ValueTypeFloat || right.Type == bytecode.ValueTypeFloat {
leftFloat := vm.toFloat(left)
rightFloat := vm.toFloat(right)
if rightFloat == 0.0 {
return bytecode.NewNullValue(), fmt.Errorf("division by zero")
}
return bytecode.NewFloatValue(leftFloat / rightFloat), nil
}
return bytecode.NewNullValue(), fmt.Errorf("cannot divide %v and %v", left.Type, right.Type)
}

func (vm *VM) mod(left, right bytecode.Value) (bytecode.Value, error) {
if left.Type == bytecode.ValueTypeInt && right.Type == bytecode.ValueTypeInt {
if right.Int == 0 {
return bytecode.NewNullValue(), fmt.Errorf("modulo by zero")
}
return bytecode.NewIntValue(left.Int % right.Int), nil
}
return bytecode.NewNullValue(), fmt.Errorf("cannot modulo %v and %v", left.Type, right.Type)
}

func (vm *VM) neg(val bytecode.Value) (bytecode.Value, error) {
if val.Type == bytecode.ValueTypeInt {
return bytecode.NewIntValue(-val.Int), nil
}
if val.Type == bytecode.ValueTypeFloat {
return bytecode.NewFloatValue(-val.Float), nil
}
return bytecode.NewNullValue(), fmt.Errorf("cannot negate %v", val.Type)
}

// Comparison operations

func (vm *VM) eq(left, right bytecode.Value) bytecode.Value {
if left.Type != right.Type {
return bytecode.NewBoolValue(false)
}
switch left.Type {
case bytecode.ValueTypeInt:
return bytecode.NewBoolValue(left.Int == right.Int)
case bytecode.ValueTypeFloat:
return bytecode.NewBoolValue(left.Float == right.Float)
case bytecode.ValueTypeBool:
return bytecode.NewBoolValue(left.Bool == right.Bool)
case bytecode.ValueTypeString:
return bytecode.NewBoolValue(left.Str == right.Str)
case bytecode.ValueTypeNull:
return bytecode.NewBoolValue(true)
default:
return bytecode.NewBoolValue(false)
}
}

func (vm *VM) ne(left, right bytecode.Value) bytecode.Value {
eq := vm.eq(left, right)
return bytecode.NewBoolValue(!eq.Bool)
}

func (vm *VM) lt(left, right bytecode.Value) (bytecode.Value, error) {
if left.Type == bytecode.ValueTypeInt && right.Type == bytecode.ValueTypeInt {
return bytecode.NewBoolValue(left.Int < right.Int), nil
}
if left.Type == bytecode.ValueTypeFloat || right.Type == bytecode.ValueTypeFloat {
leftFloat := vm.toFloat(left)
rightFloat := vm.toFloat(right)
return bytecode.NewBoolValue(leftFloat < rightFloat), nil
}
return bytecode.NewBoolValue(false), fmt.Errorf("cannot compare %v and %v", left.Type, right.Type)
}

func (vm *VM) le(left, right bytecode.Value) (bytecode.Value, error) {
if left.Type == bytecode.ValueTypeInt && right.Type == bytecode.ValueTypeInt {
return bytecode.NewBoolValue(left.Int <= right.Int), nil
}
if left.Type == bytecode.ValueTypeFloat || right.Type == bytecode.ValueTypeFloat {
leftFloat := vm.toFloat(left)
rightFloat := vm.toFloat(right)
return bytecode.NewBoolValue(leftFloat <= rightFloat), nil
}
return bytecode.NewBoolValue(false), fmt.Errorf("cannot compare %v and %v", left.Type, right.Type)
}

func (vm *VM) gt(left, right bytecode.Value) (bytecode.Value, error) {
if left.Type == bytecode.ValueTypeInt && right.Type == bytecode.ValueTypeInt {
return bytecode.NewBoolValue(left.Int > right.Int), nil
}
if left.Type == bytecode.ValueTypeFloat || right.Type == bytecode.ValueTypeFloat {
leftFloat := vm.toFloat(left)
rightFloat := vm.toFloat(right)
return bytecode.NewBoolValue(leftFloat > rightFloat), nil
}
return bytecode.NewBoolValue(false), fmt.Errorf("cannot compare %v and %v", left.Type, right.Type)
}

func (vm *VM) ge(left, right bytecode.Value) (bytecode.Value, error) {
if left.Type == bytecode.ValueTypeInt && right.Type == bytecode.ValueTypeInt {
return bytecode.NewBoolValue(left.Int >= right.Int), nil
}
if left.Type == bytecode.ValueTypeFloat || right.Type == bytecode.ValueTypeFloat {
leftFloat := vm.toFloat(left)
rightFloat := vm.toFloat(right)
return bytecode.NewBoolValue(leftFloat >= rightFloat), nil
}
return bytecode.NewBoolValue(false), fmt.Errorf("cannot compare %v and %v", left.Type, right.Type)
}

// Helper functions

func (vm *VM) toFloat(val bytecode.Value) float64 {
if val.Type == bytecode.ValueTypeFloat {
return val.Float
}
if val.Type == bytecode.ValueTypeInt {
return float64(val.Int)
}
return 0.0
}
