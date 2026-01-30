# WASM Feature Testing Guide

This document provides comprehensive testing information for the newly implemented WASM features.

## Overview

With the WASM feature expansion, the compiler now includes **15 comprehensive unit tests** that validate all major language features:

- Struct operations
- Enum operations  
- Array operations
- String manipulation
- For loops
- Pattern matching
- Advanced control flow

## Running Tests

### Run All WASM Tests

```bash
cd src/unified-compiler
go test ./internal/wasm/... -v
```

### Run Specific Test

```bash
cd src/unified-compiler
go test ./internal/wasm/... -v -run TestStructGeneration
```

### Run with Coverage

```bash
cd src/unified-compiler
go test ./internal/wasm/... -cover
```

## Test Descriptions

### 1. TestStructGeneration
**File:** `codegen_test.go`  
**Purpose:** Tests struct literal expression generation  
**Validates:**
- Creating struct expressions with multiple fields
- Heap allocation for structs
- Type ID generation

**Example AST:**
```go
Point { x: 10, y: 20 }
```

### 2. TestListGeneration
**File:** `codegen_test.go`  
**Purpose:** Tests array/list literal generation  
**Validates:**
- Creating array literals
- Memory layout with length prefix
- Element storage

**Example AST:**
```go
[1, 2, 3]
```

### 3. TestStringLiteral
**File:** `codegen_test.go`  
**Purpose:** Tests string literal handling  
**Validates:**
- String storage in data section
- String table management
- Data segment creation

**Example AST:**
```go
"hello"
```

### 4. TestIndexExpr
**File:** `codegen_test.go`  
**Purpose:** Tests array indexing  
**Validates:**
- Index expression generation
- Bounds checking code
- Memory load instructions

**Example AST:**
```go
arr[0]
```

### 5. TestForLoop
**File:** `codegen_test.go`  
**Purpose:** Tests for loop generation  
**Validates:**
- Loop iteration over arrays
- Loop variable binding
- Index management

**Example AST:**
```go
for i in arr { }
```

### 6. TestBreakContinue
**File:** `codegen_test.go`  
**Purpose:** Tests break and continue statements  
**Validates:**
- Break instruction generation
- Continue instruction generation
- Proper branching

**Example AST:**
```go
while true { break }
```

### 7. TestMemoryAllocator
**File:** `codegen_test.go`  
**Purpose:** Tests memory allocation system  
**Validates:**
- Bump allocator functionality
- 8-byte alignment
- Data segment allocation
- Memory offset tracking

### 8. TestStructFieldAccess
**File:** `codegen_test.go`  
**Purpose:** Tests field access on structs  
**Validates:**
- Field offset calculation
- Memory load from struct
- Field identifier resolution

**Example AST:**
```go
point.x
```

### 9. TestEnumConstruction
**File:** `codegen_test.go`  
**Purpose:** Tests enum variant construction  
**Validates:**
- Tagged union creation
- Variant data storage
- Tag assignment

**Example AST:**
```go
Some(42)
```

### 10. TestMatchExpression
**File:** `codegen_test.go`  
**Purpose:** Tests pattern matching  
**Validates:**
- Match expression generation
- If-else chain creation
- Pattern handling

**Example AST:**
```go
match x {
    1 => 1,
    _ => 0
}
```

### 11. TestMultipleStringLiterals
**File:** `codegen_test.go`  
**Purpose:** Tests string deduplication  
**Validates:**
- String table lookup
- Single data segment for duplicates
- Memory efficiency

**Example AST:**
```go
let s1 = "hello"
let s2 = "hello"  // Should reuse same memory
```

### 12. TestNestedStructs
**File:** `codegen_test.go`  
**Purpose:** Tests structs containing other structs  
**Validates:**
- Nested struct allocation
- Multiple heap allocations
- Complex struct expressions

**Example AST:**
```go
Rectangle { 
    top_left: Point{x:0, y:0}, 
    bottom_right: Point{x:10, y:10} 
}
```

### 13. TestArrayWithStructs
**File:** `codegen_test.go`  
**Purpose:** Tests arrays containing structs  
**Validates:**
- Array element allocation
- Struct pointers in arrays
- Mixed data structure handling

**Example AST:**
```go
[Point{x:1, y:2}, Point{x:3, y:4}]
```

### 14. TestForLoopWithBreak
**File:** `codegen_test.go`  
**Purpose:** Tests for loop with break statement  
**Validates:**
- Break in for loop
- Conditional break (if)
- Nested control structures

**Example AST:**
```go
for i in arr {
    if i > 5 { break }
}
```

### 15. TestContinueInLoop
**File:** `codegen_test.go`  
**Purpose:** Tests continue in while loop  
**Validates:**
- Continue statement generation
- Loop control flow
- Conditional continue

**Example AST:**
```go
while x < 10 {
    if x % 2 { continue }
}
```

## Test Coverage Summary

| Feature Category | Tests | Status |
|-----------------|-------|---------|
| Struct Operations | 3 | ✅ All Pass |
| Array Operations | 3 | ✅ All Pass |
| String Operations | 2 | ✅ All Pass |
| For Loops | 2 | ✅ All Pass |
| Control Flow | 2 | ✅ All Pass |
| Enums & Matching | 2 | ✅ All Pass |
| Memory Management | 1 | ✅ All Pass |
| **Total** | **15** | **✅ 100%** |

## Integration Testing

### Testing with .uni Files

To test with actual Unified source files:

```bash
cd src/unified-compiler
make build
./bin/unified --input test/point.uni
./bin/unified --input test/integration/array_basics.uni
```

### Expected Working Files

See [WASM_TEST_COMPATIBILITY.md](../WASM_TEST_COMPATIBILITY.md) for a complete list of test files that should work with the new WASM backend.

## Debugging Failed Tests

### Common Issues

1. **Memory allocation errors**
   - Check heap pointer initialization
   - Verify 8-byte alignment
   - Ensure offset calculations are correct

2. **Type mismatches**
   - Review type inference in `getExpressionType()`
   - Check i32/i64 conversions
   - Verify function signatures

3. **Bounds checking failures**
   - Ensure array length is stored at offset 0
   - Check index calculation (array + 4 + index*8)
   - Verify trap instruction placement

### Debug Mode

Run tests with verbose output:

```bash
go test ./internal/wasm/... -v -count=1
```

### Profiling

Generate test coverage report:

```bash
go test ./internal/wasm/... -coverprofile=coverage.out
go tool cover -html=coverage.out
```

## Adding New Tests

### Test Template

```go
func TestNewFeature(t *testing.T) {
    g := NewGenerator()
    
    // Create AST nodes for your test
    expr := &ast.SomeExpr{
        // ... configure expression
    }
    
    fn := &ast.FunctionDecl{
        Name: "test",
        ReturnType: &ast.TypeReference{Name: "i64"},
        Body: &ast.Block{
            Statements: []ast.Statement{
                &ast.ReturnStatement{Value: expr},
            },
        },
    }
    
    err := g.addFunction(fn)
    if err != nil {
        t.Fatalf("Failed to generate: %v", err)
    }
    
    // Add your assertions here
}
```

### Best Practices

1. **Focus on one feature** - Each test should validate a specific capability
2. **Use descriptive names** - Test name should indicate what's being tested
3. **Check multiple aspects** - Verify both success and expected state
4. **Include error cases** - Test edge cases and error conditions
5. **Document expectations** - Add comments explaining expected behavior

## Performance Testing

### Benchmark Template

```go
func BenchmarkStructGeneration(b *testing.B) {
    for i := 0; i < b.N; i++ {
        g := NewGenerator()
        // ... generate code
    }
}
```

### Running Benchmarks

```bash
go test ./internal/wasm/... -bench=. -benchmem
```

## Continuous Integration

All tests are automatically run on:
- Every commit to feature branches
- Pull request creation
- Merge to main branch

Test results are reported in CI/CD pipeline.

## See Also

- [WASM_FEATURE_EXPANSION_SUMMARY.md](../WASM_FEATURE_EXPANSION_SUMMARY.md) - Implementation details
- [WASM_TEST_COMPATIBILITY.md](../WASM_TEST_COMPATIBILITY.md) - Compatible .uni test files
- [WASM_MIGRATION_SUMMARY.md](../WASM_MIGRATION_SUMMARY.md) - Migration overview
