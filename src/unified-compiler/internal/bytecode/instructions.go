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
OpAdd                   // Add two values
OpSub                   // Subtract two values
OpMul                   // Multiply two values
OpDiv                   // Divide two values
OpMod                   // Modulo operation
OpNeg                   // Negate value

// Comparison operations
OpEq                    // Equal
OpNe                    // Not equal
OpLt                    // Less than
OpLe                    // Less than or equal
OpGt                    // Greater than
OpGe                    // Greater than or equal

// Logical operations
OpAnd                   // Logical AND
OpOr                    // Logical OR
OpNot                   // Logical NOT

// Variable operations
OpLoadLocal             // Load local variable
OpStoreLocal            // Store local variable
OpLoadGlobal            // Load global variable
OpStoreGlobal           // Store global variable

// Control flow
OpJump                  // Unconditional jump
OpJumpIfFalse           // Jump if top of stack is false
OpJumpIfTrue            // Jump if top of stack is true

// Function operations
OpCall                  // Call function
OpReturn                // Return from function
OpReturnValue           // Return value from function

// Special operations
OpHalt                  // Halt execution
OpNop                   // No operation
)

// Instruction represents a single bytecode instruction
type Instruction struct {
Op      OpCode
Operand int64 // Generic operand (index, offset, constant, etc.)
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
Type  ValueType
Int   int64
Float float64
Bool  bool
Str   string
}

// ValueType represents the type of a value
type ValueType byte

const (
ValueTypeInt ValueType = iota
ValueTypeFloat
ValueTypeBool
ValueTypeString
ValueTypeNull
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
default:
return false
}
}
