package bytecode

import (
	"testing"
)

func TestOpCodeString(t *testing.T) {
	tests := []struct {
		name     string
		op       OpCode
		expected string
	}{
		{"PUSH", OpPush, "PUSH"},
		{"POP", OpPop, "POP"},
		{"DUP", OpDup, "DUP"},
		{"ADD", OpAdd, "ADD"},
		{"SUB", OpSub, "SUB"},
		{"MUL", OpMul, "MUL"},
		{"DIV", OpDiv, "DIV"},
		{"MOD", OpMod, "MOD"},
		{"NEG", OpNeg, "NEG"},
		{"EQ", OpEq, "EQ"},
		{"NE", OpNe, "NE"},
		{"LT", OpLt, "LT"},
		{"LE", OpLe, "LE"},
		{"GT", OpGt, "GT"},
		{"GE", OpGe, "GE"},
		{"AND", OpAnd, "AND"},
		{"OR", OpOr, "OR"},
		{"NOT", OpNot, "NOT"},
		{"LOAD_LOCAL", OpLoadLocal, "LOAD_LOCAL"},
		{"STORE_LOCAL", OpStoreLocal, "STORE_LOCAL"},
		{"JUMP", OpJump, "JUMP"},
		{"JUMP_IF_FALSE", OpJumpIfFalse, "JUMP_IF_FALSE"},
		{"JUMP_IF_TRUE", OpJumpIfTrue, "JUMP_IF_TRUE"},
		{"CALL", OpCall, "CALL"},
		{"RETURN", OpReturn, "RETURN"},
		{"RETURN_VALUE", OpReturnValue, "RETURN_VALUE"},
		{"HALT", OpHalt, "HALT"},
		{"NOP", OpNop, "NOP"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.op.String()
			if result != tt.expected {
				t.Errorf("OpCode.String() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestInstructionString(t *testing.T) {
	tests := []struct {
		name     string
		inst     Instruction
		expected string
	}{
		{
			name:     "PUSH with operand",
			inst:     Instruction{Op: OpPush, Operand: 5},
			expected: "PUSH 5",
		},
		{
			name:     "ADD with no operand",
			inst:     Instruction{Op: OpAdd, Operand: 0},
			expected: "ADD 0",
		},
		{
			name:     "JUMP with target",
			inst:     Instruction{Op: OpJump, Operand: 100},
			expected: "JUMP 100",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.inst.String()
			if result != tt.expected {
				t.Errorf("Instruction.String() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestNewBytecode(t *testing.T) {
	bc := NewBytecode()
	
	if bc == nil {
		t.Fatal("NewBytecode() returned nil")
	}
	if bc.Instructions == nil {
		t.Error("Instructions not initialized")
	}
	if bc.Constants == nil {
		t.Error("Constants not initialized")
	}
	if bc.Functions == nil {
		t.Error("Functions not initialized")
	}
	if len(bc.Instructions) != 0 {
		t.Errorf("Expected 0 instructions, got %d", len(bc.Instructions))
	}
	if len(bc.Constants) != 0 {
		t.Errorf("Expected 0 constants, got %d", len(bc.Constants))
	}
	if len(bc.Functions) != 0 {
		t.Errorf("Expected 0 functions, got %d", len(bc.Functions))
	}
}

func TestAddInstruction(t *testing.T) {
	bc := NewBytecode()
	bc.AddInstruction(OpPush, 5)
	bc.AddInstruction(OpAdd, 0)

	if len(bc.Instructions) != 2 {
		t.Fatalf("Expected 2 instructions, got %d", len(bc.Instructions))
	}

	if bc.Instructions[0].Op != OpPush || bc.Instructions[0].Operand != 5 {
		t.Errorf("First instruction incorrect: %v", bc.Instructions[0])
	}

	if bc.Instructions[1].Op != OpAdd || bc.Instructions[1].Operand != 0 {
		t.Errorf("Second instruction incorrect: %v", bc.Instructions[1])
	}
}

func TestAddConstant(t *testing.T) {
	bc := NewBytecode()
	
	idx1 := bc.AddConstant(NewIntValue(42))
	idx2 := bc.AddConstant(NewFloatValue(3.14))
	idx3 := bc.AddConstant(NewBoolValue(true))

	if idx1 != 0 {
		t.Errorf("First constant index should be 0, got %d", idx1)
	}
	if idx2 != 1 {
		t.Errorf("Second constant index should be 1, got %d", idx2)
	}
	if idx3 != 2 {
		t.Errorf("Third constant index should be 2, got %d", idx3)
	}

	if len(bc.Constants) != 3 {
		t.Errorf("Expected 3 constants, got %d", len(bc.Constants))
	}
}

func TestAddFunction(t *testing.T) {
	bc := NewBytecode()
	bc.AddFunction("main", 0)
	bc.AddFunction("helper", 10)

	if len(bc.Functions) != 2 {
		t.Errorf("Expected 2 functions, got %d", len(bc.Functions))
	}

	if bc.Functions["main"] != 0 {
		t.Errorf("main function should be at index 0, got %d", bc.Functions["main"])
	}

	if bc.Functions["helper"] != 10 {
		t.Errorf("helper function should be at index 10, got %d", bc.Functions["helper"])
	}
}

func TestCurrentPosition(t *testing.T) {
	bc := NewBytecode()
	
	if bc.CurrentPosition() != 0 {
		t.Errorf("Initial position should be 0, got %d", bc.CurrentPosition())
	}

	bc.AddInstruction(OpPush, 1)
	if bc.CurrentPosition() != 1 {
		t.Errorf("Position after one instruction should be 1, got %d", bc.CurrentPosition())
	}

	bc.AddInstruction(OpAdd, 0)
	bc.AddInstruction(OpPop, 0)
	if bc.CurrentPosition() != 3 {
		t.Errorf("Position after three instructions should be 3, got %d", bc.CurrentPosition())
	}
}

func TestPatchJump(t *testing.T) {
	bc := NewBytecode()
	bc.AddInstruction(OpJump, 0) // Placeholder jump
	bc.AddInstruction(OpPush, 1)
	bc.AddInstruction(OpPush, 2)
	
	target := bc.CurrentPosition()
	bc.PatchJump(0, target)

	if bc.Instructions[0].Operand != int64(target) {
		t.Errorf("Jump not patched correctly, expected %d, got %d", target, bc.Instructions[0].Operand)
	}
}

func TestNewIntValue(t *testing.T) {
	val := NewIntValue(42)
	
	if val.Type != ValueTypeInt {
		t.Errorf("Expected ValueTypeInt, got %v", val.Type)
	}
	if val.Int != 42 {
		t.Errorf("Expected Int value 42, got %d", val.Int)
	}
}

func TestNewFloatValue(t *testing.T) {
	val := NewFloatValue(3.14)
	
	if val.Type != ValueTypeFloat {
		t.Errorf("Expected ValueTypeFloat, got %v", val.Type)
	}
	if val.Float != 3.14 {
		t.Errorf("Expected Float value 3.14, got %f", val.Float)
	}
}

func TestNewBoolValue(t *testing.T) {
	valTrue := NewBoolValue(true)
	valFalse := NewBoolValue(false)
	
	if valTrue.Type != ValueTypeBool {
		t.Errorf("Expected ValueTypeBool, got %v", valTrue.Type)
	}
	if !valTrue.Bool {
		t.Error("Expected true value")
	}
	if valFalse.Bool {
		t.Error("Expected false value")
	}
}

func TestNewStringValue(t *testing.T) {
	val := NewStringValue("hello")
	
	if val.Type != ValueTypeString {
		t.Errorf("Expected ValueTypeString, got %v", val.Type)
	}
	if val.Str != "hello" {
		t.Errorf("Expected String value 'hello', got '%s'", val.Str)
	}
}

func TestNewNullValue(t *testing.T) {
	val := NewNullValue()
	
	if val.Type != ValueTypeNull {
		t.Errorf("Expected ValueTypeNull, got %v", val.Type)
	}
}

func TestValueString(t *testing.T) {
	tests := []struct {
		name     string
		value    Value
		expected string
	}{
		{"Int", NewIntValue(42), "42"},
		{"Float", NewFloatValue(3.14), "3.140000"},
		{"Bool true", NewBoolValue(true), "true"},
		{"Bool false", NewBoolValue(false), "false"},
		{"String", NewStringValue("hello"), `"hello"`},
		{"Null", NewNullValue(), "null"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.value.String()
			if result != tt.expected {
				t.Errorf("Value.String() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestValueIsTruthy(t *testing.T) {
	tests := []struct {
		name     string
		value    Value
		expected bool
	}{
		{"Bool true", NewBoolValue(true), true},
		{"Bool false", NewBoolValue(false), false},
		{"Int non-zero", NewIntValue(42), true},
		{"Int zero", NewIntValue(0), false},
		{"Float non-zero", NewFloatValue(3.14), true},
		{"Float zero", NewFloatValue(0.0), false},
		{"String non-empty", NewStringValue("hello"), true},
		{"String empty", NewStringValue(""), false},
		{"Null", NewNullValue(), false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.value.IsTruthy()
			if result != tt.expected {
				t.Errorf("Value.IsTruthy() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestAddInstructionWithArgCount(t *testing.T) {
	bc := NewBytecode()
	bc.AddInstructionWithArgCount(OpCall, 10, 3)

	if len(bc.Instructions) != 1 {
		t.Fatalf("Expected 1 instruction, got %d", len(bc.Instructions))
	}

	inst := bc.Instructions[0]
	if inst.Op != OpCall {
		t.Errorf("Expected OpCall, got %v", inst.Op)
	}
	if inst.Operand != 10 {
		t.Errorf("Expected operand 10, got %d", inst.Operand)
	}
	if inst.ArgCount != 3 {
		t.Errorf("Expected ArgCount 3, got %d", inst.ArgCount)
	}
}
