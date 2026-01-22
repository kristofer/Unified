# Testing Guide for Unified Compiler

This guide explains how to write and run tests for the Unified compiler.

## Test Structure

The Unified compiler has a comprehensive test suite organized into several categories:

### 1. Unit Tests

#### Bytecode Module (`internal/bytecode`)
Tests for bytecode instructions, constants, and values.

**Files:**
- `instructions_test.go` - Tests for OpCode, Instructions, Bytecode, and Values
- `generator_test.go` - Tests for AST to bytecode generation

**Example:**
```go
func TestOpCodeString(t *testing.T) {
    op := bytecode.OpPush
    if op.String() != "PUSH" {
        t.Errorf("Expected 'PUSH', got '%s'", op.String())
    }
}
```

#### VM Module (`internal/vm`)
Tests for the virtual machine execution engine.

**Files:**
- `stack_test.go` - Tests for the execution stack
- `vm_test.go` - Tests for VM execution logic

**Example:**
```go
func TestVMSimpleArithmetic(t *testing.T) {
    bc := bytecode.NewBytecode()
    bc.AddFunction("main", 0)
    // ... setup bytecode ...
    
    vm := vm.NewVM(bc)
    result, err := vm.Run()
    // ... assertions ...
}
```

### 2. Integration Tests

End-to-end tests that compile and run complete Unified programs.

**Location:** `cmd/compiler/integration_test.go`

**Test Files:** `test/integration/*.uni`

**Example:**
```go
func TestIntegrationSimpleReturn(t *testing.T) {
    testCompileAndRun(t, "../../test/integration/simple_return.uni", 42)
}
```

## Running Tests

### Run All Tests
```bash
cd src/unified-compiler
go test ./...
```

### Run Tests with Verbose Output
```bash
go test ./... -v
```

### Run Tests for Specific Module
```bash
go test ./internal/bytecode -v
go test ./internal/vm -v
go test ./cmd/compiler -v
```

### Run Specific Test
```bash
go test ./internal/vm -run TestVMSimpleArithmetic -v
```

### Run with Coverage
```bash
go test ./... -cover
go test ./... -coverprofile=coverage.out
go tool cover -html=coverage.out
```

### Run with Race Detection
```bash
go test ./... -race
```

## Writing Tests

### Unit Test Guidelines

1. **Test One Thing**: Each test should verify a single behavior
2. **Use Table-Driven Tests**: For testing multiple inputs
3. **Clear Names**: Test names should describe what they test
4. **Helper Functions**: Use `t.Helper()` for test utilities

**Example Table-Driven Test:**
```go
func TestValueIsTruthy(t *testing.T) {
    tests := []struct {
        name     string
        value    Value
        expected bool
    }{
        {"Bool true", NewBoolValue(true), true},
        {"Int zero", NewIntValue(0), false},
        {"String empty", NewStringValue(""), false},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := tt.value.IsTruthy()
            if result != tt.expected {
                t.Errorf("Expected %v, got %v", tt.expected, result)
            }
        })
    }
}
```

### Integration Test Guidelines

1. **Create Test File**: Add a `.uni` file to `test/integration/`
2. **Add Test Function**: Add corresponding test in `cmd/compiler/integration_test.go`
3. **Use Helper**: Use `testCompileAndRun()` helper function

**Example:**
```go
func TestIntegrationNewFeature(t *testing.T) {
    testCompileAndRun(t, "../../test/integration/new_feature.uni", expectedExitCode)
}
```

## Test Coverage

Current test coverage by module:

- **Bytecode Instructions**: 28 test cases
  - OpCode operations
  - Instruction manipulation
  - Constant pool
  - Value types

- **VM Stack**: 9 test cases
  - Push/Pop operations
  - Peek operations
  - Size management

- **VM Execution**: 19 test cases
  - Arithmetic operations
  - Comparisons
  - Logical operations
  - Control flow
  - Function calls

- **Bytecode Generator**: 16 test cases
  - Expression generation
  - Statement generation
  - Function generation
  - Error handling

- **Integration**: 3 test cases
  - Simple programs
  - Function calls
  - Local variables

**Total: 76 test cases**

## Continuous Integration

Tests are automatically run on every commit and pull request. All tests must pass before merging.

## Adding New Tests

When adding new features:

1. **Write tests first** (TDD approach recommended)
2. **Add unit tests** for the new components
3. **Add integration tests** for complete features
4. **Update this guide** if adding new test categories
5. **Ensure all tests pass** before committing

## Debugging Tests

### View Detailed Output
```bash
go test ./internal/vm -v -run TestVMSimpleArithmetic
```

### Print Debug Information
```go
func TestExample(t *testing.T) {
    // Use t.Logf for debug output
    t.Logf("Debug info: %v", someValue)
    
    // Or use t.Error for failures with context
    if result != expected {
        t.Errorf("Expected %v, got %v. Details: %+v", expected, result, context)
    }
}
```

### Skip Tests Temporarily
```go
func TestExample(t *testing.T) {
    t.Skip("Skipping until feature X is implemented")
    // test code...
}
```

## Best Practices

1. **Keep tests simple and focused**
2. **Use meaningful test names**
3. **Test both success and failure cases**
4. **Clean up resources** (files, connections, etc.)
5. **Don't rely on test execution order**
6. **Use test fixtures** for complex setup
7. **Mock external dependencies** when necessary
8. **Keep tests fast** - unit tests should run in milliseconds
9. **Test edge cases** (empty inputs, null values, boundaries)
10. **Document complex test scenarios**

## Common Patterns

### Testing Errors
```go
func TestErrorCase(t *testing.T) {
    _, err := SomeFunction()
    if err == nil {
        t.Error("Expected error, got nil")
    }
}
```

### Setup and Teardown
```go
func TestMain(m *testing.M) {
    // Setup
    setup()
    
    // Run tests
    code := m.Run()
    
    // Teardown
    teardown()
    
    os.Exit(code)
}
```

### Parallel Tests
```go
func TestParallel(t *testing.T) {
    t.Parallel()
    // test code...
}
```

## Resources

- [Go Testing Package](https://pkg.go.dev/testing)
- [Table Driven Tests](https://github.com/golang/go/wiki/TableDrivenTests)
- [Advanced Testing with Go](https://www.youtube.com/watch?v=8hQG7QlcLBk)
