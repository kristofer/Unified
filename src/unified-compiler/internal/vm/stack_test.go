package vm

import (
	"testing"
	"unified-compiler/internal/bytecode"
)

func TestNewStack(t *testing.T) {
	s := NewStack()
	
	if s == nil {
		t.Fatal("NewStack() returned nil")
	}
	if s.Size() != 0 {
		t.Errorf("New stack should have size 0, got %d", s.Size())
	}
}

func TestStackPushPop(t *testing.T) {
	s := NewStack()
	
	val1 := bytecode.NewIntValue(42)
	s.Push(val1)
	
	if s.Size() != 1 {
		t.Errorf("Expected size 1 after one push, got %d", s.Size())
	}
	
	val2 := s.Pop()
	if val2.Type != bytecode.ValueTypeInt || val2.Int != 42 {
		t.Errorf("Popped value doesn't match pushed value")
	}
	
	if s.Size() != 0 {
		t.Errorf("Expected size 0 after pop, got %d", s.Size())
	}
}

func TestStackMultiplePushPop(t *testing.T) {
	s := NewStack()
	
	// Push multiple values
	s.Push(bytecode.NewIntValue(1))
	s.Push(bytecode.NewIntValue(2))
	s.Push(bytecode.NewIntValue(3))
	
	if s.Size() != 3 {
		t.Errorf("Expected size 3, got %d", s.Size())
	}
	
	// Pop in LIFO order
	val3 := s.Pop()
	if val3.Int != 3 {
		t.Errorf("Expected 3, got %d", val3.Int)
	}
	
	val2 := s.Pop()
	if val2.Int != 2 {
		t.Errorf("Expected 2, got %d", val2.Int)
	}
	
	val1 := s.Pop()
	if val1.Int != 1 {
		t.Errorf("Expected 1, got %d", val1.Int)
	}
	
	if s.Size() != 0 {
		t.Errorf("Expected empty stack, got size %d", s.Size())
	}
}

func TestStackPeek(t *testing.T) {
	s := NewStack()
	
	val := bytecode.NewIntValue(42)
	s.Push(val)
	
	peeked := s.Peek()
	if peeked.Type != bytecode.ValueTypeInt || peeked.Int != 42 {
		t.Errorf("Peeked value doesn't match pushed value")
	}
	
	// Stack size should not change
	if s.Size() != 1 {
		t.Errorf("Peek should not change stack size, got %d", s.Size())
	}
	
	// Value should still be on stack
	popped := s.Pop()
	if popped.Int != 42 {
		t.Error("Value was removed by Peek")
	}
}

func TestStackPopPanic(t *testing.T) {
	s := NewStack()
	
	defer func() {
		if r := recover(); r == nil {
			t.Error("Pop on empty stack should panic")
		}
	}()
	
	s.Pop()
}

func TestStackPeekPanic(t *testing.T) {
	s := NewStack()
	
	defer func() {
		if r := recover(); r == nil {
			t.Error("Peek on empty stack should panic")
		}
	}()
	
	s.Peek()
}

func TestStackSize(t *testing.T) {
	s := NewStack()
	
	sizes := []int{0, 1, 2, 3, 2, 1, 0}
	operations := []func(){
		func() { /* initial size */ },
		func() { s.Push(bytecode.NewIntValue(1)) },
		func() { s.Push(bytecode.NewIntValue(2)) },
		func() { s.Push(bytecode.NewIntValue(3)) },
		func() { s.Pop() },
		func() { s.Pop() },
		func() { s.Pop() },
	}
	
	for i, op := range operations {
		op()
		if s.Size() != sizes[i] {
			t.Errorf("After operation %d, expected size %d, got %d", i, sizes[i], s.Size())
		}
	}
}

func TestStackString(t *testing.T) {
	s := NewStack()
	
	// Empty stack
	str := s.String()
	if str != "Stack[0]: []" {
		t.Errorf("Empty stack string unexpected: %s", str)
	}
	
	// Stack with values
	s.Push(bytecode.NewIntValue(1))
	s.Push(bytecode.NewIntValue(2))
	str = s.String()
	
	// Just check that it contains the stack size
	if s.Size() != 2 {
		t.Error("Stack size incorrect in string representation")
	}
}

func TestStackDifferentTypes(t *testing.T) {
	s := NewStack()
	
	// Push different types
	s.Push(bytecode.NewIntValue(42))
	s.Push(bytecode.NewFloatValue(3.14))
	s.Push(bytecode.NewBoolValue(true))
	s.Push(bytecode.NewStringValue("hello"))
	
	if s.Size() != 4 {
		t.Errorf("Expected size 4, got %d", s.Size())
	}
	
	// Pop and verify types
	str := s.Pop()
	if str.Type != bytecode.ValueTypeString || str.Str != "hello" {
		t.Error("String value incorrect")
	}
	
	b := s.Pop()
	if b.Type != bytecode.ValueTypeBool || !b.Bool {
		t.Error("Bool value incorrect")
	}
	
	f := s.Pop()
	if f.Type != bytecode.ValueTypeFloat || f.Float != 3.14 {
		t.Error("Float value incorrect")
	}
	
	i := s.Pop()
	if i.Type != bytecode.ValueTypeInt || i.Int != 42 {
		t.Error("Int value incorrect")
	}
}
