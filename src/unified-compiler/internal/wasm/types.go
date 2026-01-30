package wasm

import (
	"unified-compiler/internal/bytecode"
)

// ValueType represents WASM value types
type ValueType byte

const (
	I32 ValueType = 0x7F
	I64 ValueType = 0x7E
	F32 ValueType = 0x7D
	F64 ValueType = 0x7C
)

// Value represents a runtime value in WASM
type Value struct {
	Type ValueType
	I32  int32
	I64  int64
	F32  float32
	F64  float64
}

// ConvertFromBytecode converts bytecode.Value to wasm.Value
func ConvertFromBytecode(bv bytecode.Value) Value {
	switch bv.Type {
	case bytecode.ValueTypeInt:
		return Value{Type: I64, I64: bv.Int}
	case bytecode.ValueTypeFloat:
		return Value{Type: F64, F64: bv.Float}
	case bytecode.ValueTypeBool:
		if bv.Bool {
			return Value{Type: I32, I32: 1}
		}
		return Value{Type: I32, I32: 0}
	default:
		return Value{Type: I32, I32: 0}
	}
}

// ConvertToBytecode converts wasm.Value to bytecode.Value
func ConvertToBytecode(wv Value) bytecode.Value {
	switch wv.Type {
	case I32:
		return bytecode.NewIntValue(int64(wv.I32))
	case I64:
		return bytecode.NewIntValue(wv.I64)
	case F32:
		return bytecode.NewFloatValue(float64(wv.F32))
	case F64:
		return bytecode.NewFloatValue(wv.F64)
	default:
		return bytecode.NewIntValue(0)
	}
}
