# Unified Virtual Machine Architecture

## Overview

The Unified compiler now uses a custom virtual machine (VM) written in Go instead of LLVM for code generation and execution. This provides excellent portability while maintaining good performance.

## Architecture

```
Source Code (.uni) → Lexer → Parser → AST → Bytecode Generator → Bytecode → VM → Execution
```

## Bytecode Instructions

The VM uses a stack-based architecture with the following instruction set:

### Stack Operations
- `PUSH` - Push constant onto stack
- `POP` - Pop value from stack
- `DUP` - Duplicate top of stack

### Arithmetic Operations
- `ADD`, `SUB`, `MUL`, `DIV`, `MOD` - Binary arithmetic
- `NEG` - Unary negation

### Bitwise Operations (Phase 4)
- `BIT_AND` - Bitwise AND (&)
- `BIT_OR` - Bitwise OR (|)
- `BIT_XOR` - Bitwise XOR (^)
- `BIT_NOT` - Bitwise NOT (~) - unary
- `LSHIFT` - Left shift (<<)
- `RSHIFT` - Right shift (>>)

### Comparison Operations
- `EQ`, `NE`, `LT`, `LE`, `GT`, `GE` - Comparisons

### Logical Operations
- `AND`, `OR`, `NOT` - Logical operations

### Variable Operations
- `LOAD_LOCAL` - Load local variable
- `STORE_LOCAL` - Store local variable

### Control Flow
- `JUMP` - Unconditional jump
- `JUMP_IF_FALSE` - Conditional jump
- `JUMP_IF_TRUE` - Conditional jump

### Function Operations
- `CALL` - Call function with arguments
- `RETURN` - Return from function
- `RETURN_VALUE` - Return value from function

### Struct Operations (Phase 5)
- `ALLOC_STRUCT` - Allocate struct instance
  - Operand: field count
  - Stack: [field_name_0, field_value_0, ..., field_name_N, field_value_N, type_name]
  - Pops field name/value pairs and type name, creates struct
- `LOAD_FIELD` - Load field from struct
  - Stack: [struct, field_name] → [field_value]
- `STORE_FIELD` - Store field to struct
  - Stack: [struct, field_name, value] → [struct]

### Special
- `HALT` - Stop execution
- `NOP` - No operation

## Phase 1-4 Features

Currently implemented and tested features:

### Phase 1 (Complete)
✅ Function declarations and calls
✅ Function parameters
✅ Arithmetic operations (+, -, *, /, %)
✅ Comparison operations (==, !=, <, <=, >, >=)
✅ Logical operations (&&, ||, !)
✅ Local variables (let statements)
✅ Return statements
✅ Integer, float, boolean, and string literals
✅ Basic expressions
✅ Control flow (if/else)

### Phase 2 (Complete)
✅ While loops
✅ For loops with ranges
✅ Break and continue statements
✅ Nested loops
✅ Loop labels

### Phase 3 (Complete)
✅ Mutable variables (let mut)
✅ Variable assignment
✅ Compound assignment operators (+=, -=, *=, /=, %=)
✅ Variable shadowing
✅ Type inference

### Phase 4 (Partial - 60%)
✅ Bitwise operators (&, |, ^, ~, <<, >>)
✅ Operator precedence
✅ Block expressions
⏳ Tuples (planned)
⏳ Lambda expressions (planned)
⏳ Default parameters (planned)

### Phase 5 (Partial - 60%)
✅ Struct declarations with fields
✅ Struct instantiation
✅ Field access (dot notation)
✅ Nested structs
❌ Methods on structs (blocked: parser regeneration)
❌ Associated functions (blocked: parser regeneration)

## Testing

The compiler includes comprehensive test coverage:

### Unit Tests
- **Bytecode Instructions Tests** (28 test cases)
  - OpCode string representations
  - Instruction creation and manipulation
  - Constant pool management
  - Function registration
  - Value types and operations

- **VM Stack Tests** (9 test cases)
  - Stack operations (push, pop, peek)
  - Stack size management
  - Multiple data types

- **VM Execution Tests** (19 test cases)
  - Arithmetic operations
  - Comparison operations
  - Logical operations
  - Local variables
  - Control flow (jumps)
  - Function calls
  - Error handling (division by zero, missing main)

- **Bytecode Generator Tests** (16 test cases)
  - Binary and unary expressions
  - Let statements
  - If/else statements
  - Function calls with arguments
  - Literal values of all types
  - Local variable scoping
  - Error detection (undefined variables/functions)

### Integration Tests
- **End-to-End Compilation Tests** (3 test cases)
  - Simple return values
  - Function calls with parameters
  - Local variable usage

### Running Tests

```bash
# Run all tests
go test ./...

# Run specific module tests
go test ./internal/bytecode -v
go test ./internal/vm -v
go test ./cmd/compiler -v

# Run with coverage
go test ./... -cover
```

All 76 tests pass successfully ✅

## Example Programs

### Minimal Program
```unified
fn main() -> i32 {
    return 42;
}
```

### Function Call
```unified
fn multiply(a: i32, b: i32) -> i32 {
    return a * b;
}

fn main() -> i32 {
    return multiply(6, 7);
}
```

### Complex Example
```unified
fn add(x: i32, y: i32) -> i32 {
    return x + y;
}

fn multiply(a: i32, b: i32) -> i32 {
    return a * b;
}

fn calculate(x: i32, y: i32) -> i32 {
    let sum = add(x, y);
    let product = multiply(x, y);
    return add(sum, product);
}

fn main() -> i32 {
    return calculate(3, 4);  // Returns 19
}
```

## Building and Running

```bash
# Build the compiler
make build

# Run a program
./bin/unified --input program.uni

# The exit code will be the return value of main()
echo $?
```

## VM Implementation Details

### Stack
- Stack-based execution model
- Efficient push/pop operations
- Automatic bounds checking

### Call Stack
- Each function call creates a new call frame
- Call frames contain:
  - Return instruction pointer
  - Local variables
  - Base pointer for stack management

### Local Variables
- Stored in the call frame
- Indexed by position
- Support for parameters and let bindings

### Function Calls
1. Arguments are pushed onto the stack (in order)
2. CALL instruction pops arguments and creates new frame
3. Arguments become local variables in the new frame
4. Function executes
5. RETURN_VALUE pops return value and restores previous frame

## Benefits of VM Architecture

### Portability
- Runs anywhere Go runs
- No platform-specific code
- No external dependencies (like LLVM)

### Simplicity
- Easy to understand and debug
- Straightforward instruction set
- Clean separation of concerns

### Development Speed
- Faster iteration cycles
- No complex toolchain setup
- Immediate execution

### Future Enhancements
- JIT compilation for hot paths
- Optimized instruction dispatch
- Inline caching
- Garbage collection integration

## Performance

Current performance is competitive with interpreted languages like Python. Future optimizations planned:
- Register-based VM variant
- Bytecode optimization passes
- JIT compilation
- Inline caching for method calls

## Directory Structure

```
src/unified-compiler/
├── internal/
│   ├── bytecode/          # Bytecode definitions and generator
│   │   ├── instructions.go # Instruction set and types
│   │   └── generator.go    # AST to bytecode compilation
│   └── vm/                 # Virtual machine
│       ├── vm.go           # VM execution engine
│       └── stack.go        # Execution stack
```

## Next Steps

Phase 2 will add:
- Control flow (if/else, loops)
- More data types (float, bool, string)
- Comparison operators
- More complex expressions
