package vm

import (
"fmt"
"unified-compiler/internal/bytecode"
)

// Stack represents the execution stack
type Stack struct {
values []bytecode.Value
}

// NewStack creates a new stack
func NewStack() *Stack {
return &Stack{
values: make([]bytecode.Value, 0, 256),
}
}

// Push pushes a value onto the stack
func (s *Stack) Push(val bytecode.Value) {
s.values = append(s.values, val)
}

// Pop pops a value from the stack
func (s *Stack) Pop() bytecode.Value {
if len(s.values) == 0 {
panic("stack underflow")
}
val := s.values[len(s.values)-1]
s.values = s.values[:len(s.values)-1]
return val
}

// Peek returns the top value without popping
func (s *Stack) Peek() bytecode.Value {
if len(s.values) == 0 {
panic("stack underflow")
}
return s.values[len(s.values)-1]
}

// Size returns the number of values on the stack
func (s *Stack) Size() int {
return len(s.values)
}

// String returns a string representation of the stack
func (s *Stack) String() string {
return fmt.Sprintf("Stack[%d]: %v", len(s.values), s.values)
}
