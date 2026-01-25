package bytecode

import "fmt"

// OpCode represents a bytecode instruction operation
type OpCode byte

// Bytecode instruction set for the Unified VM (Phase 1)
const (
	// Stack operations
	OpPush    OpCode = iota // Push constant onto stack
	OpPop                   // Pop value from stack
	OpDup                   // Duplicate top of stack

	// Arithmetic operations
	OpAdd // Add two values
	OpSub // Subtract two values
	OpMul // Multiply two values
	OpDiv // Divide two values
	OpMod // Modulo operation
	OpNeg // Negate value

	// Bitwise operations
	OpBitAnd   // Bitwise AND
	OpBitOr    // Bitwise OR
	OpBitXor   // Bitwise XOR
	OpBitNot   // Bitwise NOT (unary)
	OpLShift   // Left shift
	OpRShift   // Right shift

	// Comparison operations
	OpEq // Equal
	OpNe // Not equal
	OpLt // Less than
	OpLe // Less than or equal
	OpGt // Greater than
	OpGe // Greater than or equal

	// Logical operations
	OpAnd // Logical AND
	OpOr  // Logical OR
	OpNot // Logical NOT

	// Variable operations
	OpLoadLocal   // Load local variable
	OpStoreLocal  // Store local variable
	OpLoadGlobal  // Load global variable
	OpStoreGlobal // Store global variable

	// Control flow
	OpJump        // Unconditional jump
	OpJumpIfFalse // Jump if top of stack is false
	OpJumpIfTrue  // Jump if top of stack is true

	// Function operations
	OpCall        // Call function
	OpReturn      // Return from function
	OpReturnValue // Return value from function

	// Struct operations
	OpAllocStruct // Allocate struct instance
	OpLoadField   // Load field from struct
	OpStoreField  // Store field to struct

	// Enum operations
	OpAllocEnum       // Allocate enum variant
	OpMatchVariant    // Match enum variant by tag
	OpExtractVariant  // Extract data from enum variant

	// Error handling operations
	OpTryPropagate    // Try operator (?) for error propagation

	// Array operations
	OpAllocArray  // Allocate array with size
	OpLoadArray   // Load element from array (with bounds checking)
	OpStoreArray  // Store element to array (with bounds checking)
	OpArrayLen    // Get array length

	// Special operations
	OpHalt // Halt execution
	OpNop  // No operation
)

// Instruction represents a single bytecode instruction
type Instruction struct {
Op      OpCode
Operand int64 // Generic operand (index, offset, constant, etc.)
	ArgCount int  // For function calls: number of arguments
}

// String returns a human-readable representation of the instruction
func (i Instruction) String() string {
return fmt.Sprintf("%s %d", i.Op, i.Operand)
}

// String returns a human-readable name for the opcode
func (op OpCode) String() string {
switch op {
case OpPush:
return "PUSH"
case OpPop:
return "POP"
case OpDup:
return "DUP"
case OpAdd:
return "ADD"
case OpSub:
return "SUB"
case OpMul:
return "MUL"
case OpDiv:
return "DIV"
case OpMod:
return "MOD"
case OpNeg:
return "NEG"
case OpBitAnd:
return "BIT_AND"
case OpBitOr:
return "BIT_OR"
case OpBitXor:
return "BIT_XOR"
case OpBitNot:
return "BIT_NOT"
case OpLShift:
return "LSHIFT"
case OpRShift:
return "RSHIFT"
case OpEq:
return "EQ"
case OpNe:
return "NE"
case OpLt:
return "LT"
case OpLe:
return "LE"
case OpGt:
return "GT"
case OpGe:
return "GE"
case OpAnd:
return "AND"
case OpOr:
return "OR"
case OpNot:
return "NOT"
case OpLoadLocal:
return "LOAD_LOCAL"
case OpStoreLocal:
return "STORE_LOCAL"
case OpLoadGlobal:
return "LOAD_GLOBAL"
case OpStoreGlobal:
return "STORE_GLOBAL"
case OpJump:
return "JUMP"
case OpJumpIfFalse:
return "JUMP_IF_FALSE"
case OpJumpIfTrue:
return "JUMP_IF_TRUE"
case OpCall:
return "CALL"
case OpReturn:
return "RETURN"
case OpReturnValue:
return "RETURN_VALUE"
case OpAllocStruct:
return "ALLOC_STRUCT"
case OpLoadField:
return "LOAD_FIELD"
case OpStoreField:
return "STORE_FIELD"
case OpAllocEnum:
return "ALLOC_ENUM"
case OpMatchVariant:
return "MATCH_VARIANT"
case OpExtractVariant:
return "EXTRACT_VARIANT"
case OpTryPropagate:
return "TRY_PROPAGATE"
case OpAllocArray:
return "ALLOC_ARRAY"
case OpLoadArray:
return "LOAD_ARRAY"
case OpStoreArray:
return "STORE_ARRAY"
case OpArrayLen:
return "ARRAY_LEN"
case OpHalt:
return "HALT"
case OpNop:
return "NOP"
default:
return fmt.Sprintf("UNKNOWN(%d)", op)
}
}

// Bytecode represents a compiled program
type Bytecode struct {
Instructions []Instruction
Constants    []Value // Constant pool
Functions    map[string]int // Function name -> instruction index
}

// NewBytecode creates a new bytecode container
func NewBytecode() *Bytecode {
return &Bytecode{
Instructions: make([]Instruction, 0),
Constants:    make([]Value, 0),
Functions:    make(map[string]int),
}
}

// AddInstruction adds an instruction to the bytecode
func (b *Bytecode) AddInstruction(op OpCode, operand int64) {
b.Instructions = append(b.Instructions, Instruction{Op: op, Operand: operand})
}

// AddConstant adds a constant to the pool and returns its index
func (b *Bytecode) AddConstant(val Value) int {
b.Constants = append(b.Constants, val)
return len(b.Constants) - 1
}

// AddFunction registers a function entry point
func (b *Bytecode) AddFunction(name string, index int) {
b.Functions[name] = index
}

// CurrentPosition returns the current instruction position
func (b *Bytecode) CurrentPosition() int {
return len(b.Instructions)
}

// PatchJump patches a jump instruction at the given position
func (b *Bytecode) PatchJump(position int, target int) {
b.Instructions[position].Operand = int64(target)
}

// Value represents a value in the VM
type Value struct {
	Type   ValueType
	Int    int64
	Float  float64
	Bool   bool
	Str    string
	Struct *StructValue // For struct instances
	Enum   *EnumValue   // For enum instances
	Array  *ArrayValue  // For array instances
}

// StructValue represents a struct instance in the VM
type StructValue struct {
	TypeName string
	Fields   map[string]Value
}

// EnumValue represents an enum instance in the VM
type EnumValue struct {
	EnumName    string
	VariantName string
	VariantTag  int
	Data        []Value // Variant data
}

// ArrayValue represents an array instance in the VM
type ArrayValue struct {
	Elements []Value
	Length   int
}

// ValueType represents the type of a value
type ValueType byte

const (
	ValueTypeInt ValueType = iota
	ValueTypeFloat
	ValueTypeBool
	ValueTypeString
	ValueTypeNull
	ValueTypeStruct
	ValueTypeEnum
	ValueTypeArray
)

// NewIntValue creates an integer value
func NewIntValue(i int64) Value {
return Value{Type: ValueTypeInt, Int: i}
}

// NewFloatValue creates a float value
func NewFloatValue(f float64) Value {
return Value{Type: ValueTypeFloat, Float: f}
}

// NewBoolValue creates a boolean value
func NewBoolValue(b bool) Value {
return Value{Type: ValueTypeBool, Bool: b}
}

// NewStringValue creates a string value
func NewStringValue(s string) Value {
return Value{Type: ValueTypeString, Str: s}
}

// NewNullValue creates a null value
func NewNullValue() Value {
return Value{Type: ValueTypeNull}
}

// NewStructValue creates a struct value
func NewStructValue(typeName string, fields map[string]Value) Value {
return Value{
Type: ValueTypeStruct,
Struct: &StructValue{
TypeName: typeName,
Fields:   fields,
},
}
}

// NewEnumValue creates an enum value
func NewEnumValue(enumName, variantName string, variantTag int, data []Value) Value {
	return Value{
		Type: ValueTypeEnum,
		Enum: &EnumValue{
			EnumName:    enumName,
			VariantName: variantName,
			VariantTag:  variantTag,
			Data:        data,
		},
	}
}

// NewArrayValue creates an array value
func NewArrayValue(elements []Value) Value {
	return Value{
		Type: ValueTypeArray,
		Array: &ArrayValue{
			Elements: elements,
			Length:   len(elements),
		},
	}
}

// String returns a human-readable representation of the value
func (v Value) String() string {
switch v.Type {
case ValueTypeInt:
return fmt.Sprintf("%d", v.Int)
case ValueTypeFloat:
return fmt.Sprintf("%f", v.Float)
case ValueTypeBool:
return fmt.Sprintf("%t", v.Bool)
case ValueTypeString:
return fmt.Sprintf("%q", v.Str)
case ValueTypeNull:
return "null"
case ValueTypeStruct:
return fmt.Sprintf("struct:%s", v.Struct.TypeName)
case ValueTypeEnum:
return fmt.Sprintf("enum:%s::%s", v.Enum.EnumName, v.Enum.VariantName)
case ValueTypeArray:
return fmt.Sprintf("array[%d]", v.Array.Length)
default:
return "unknown"
}
}

// IsTruthy returns whether the value is truthy
func (v Value) IsTruthy() bool {
switch v.Type {
case ValueTypeBool:
return v.Bool
case ValueTypeInt:
return v.Int != 0
case ValueTypeFloat:
return v.Float != 0.0
case ValueTypeString:
return len(v.Str) > 0
case ValueTypeNull:
return false
case ValueTypeStruct:
return true // Structs are always truthy if they exist
case ValueTypeEnum:
return true // Enums are always truthy if they exist
case ValueTypeArray:
return true // Arrays are always truthy if they exist
default:
return false
}
}

// AddInstructionWithArgCount adds an instruction with argument count
func (b *Bytecode) AddInstructionWithArgCount(op OpCode, operand int64, argCount int) {
b.Instructions = append(b.Instructions, Instruction{Op: op, Operand: operand, ArgCount: argCount})
}
