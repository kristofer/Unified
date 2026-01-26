# Phase 10: Generics Basics - Test Suite

This directory contains comprehensive tests for Phase 10 generic functionality.

## Test Structure

### Go Unit Tests (10 tests in `internal/semantic/generics_test.go`)

These tests validate the AST structures and semantic analysis for generics:

1. **TestGenericFunctionSingleTypeParam** - Tests parsing of `fn identity<T>(x: T) -> T`
2. **TestGenericFunctionMultipleTypeParams** - Tests parsing of `fn pair<A, B>(a: A, b: B)`
3. **TestGenericStruct** - Tests parsing of `struct Box<T> { value: T }`
4. **TestGenericEnum** - Tests parsing of `enum Option<T> { Some(T), None }`
5. **TestTypeParameterWithConstraints** - Tests type parameter bounds like `T: Comparable`
6. **TestTypeApplicationWithConcreteTypes** - Tests instantiation like `Box<Int>`
7. **TestMultipleTypeArguments** - Tests multi-param instantiation like `Result<Int, String>`
8. **TestNestedGenericTypes** - Tests nested generics like `Option<Box<Int>>`
9. **TestGenericFunctionWithWhereClause** - Tests where clause constraints
10. **TestGenericTypeCompatibility** - Tests type compatibility checking for generics

### Unified Language Tests (10 files in `test/generics/*.uni`)

These are complete Unified programs that demonstrate generic features:

1. **01_identity_function.uni** - Basic generic identity function with type inference
2. **02_box_struct.uni** - Generic struct instantiation and field access
3. **03_option_enum.uni** - Pattern matching with generic Option enum
4. **04_result_enum.uni** - Error handling with generic Result enum
5. **05_pair_function.uni** - Multiple type parameters and tuple handling
6. **06_swap_function.uni** - Generic references and mutation
7. **07_array_operations.uni** - Generics with array/slice types
8. **08_container_methods.uni** - Methods on generic structs with impl blocks
9. **09_comparison_function.uni** - Constrained generics with trait bounds
10. **10_linked_list.uni** - Recursive generic structures

## Running the Tests

### Go Unit Tests

```bash
cd /home/runner/work/Unified/Unified/src/unified-compiler
go test ./internal/semantic -run TestGeneric -v
```

### Unified Integration Tests

These tests will be integrated into the compiler test suite once generics are fully implemented:

```bash
./bin/unified --input test/generics/01_identity_function.uni
./bin/unified --input test/generics/02_box_struct.uni
# ... etc
```

## Test Coverage

These tests cover the following aspects of generics:

### Core Features
- ✅ Generic function declarations
- ✅ Generic struct declarations
- ✅ Generic enum declarations
- ✅ Type parameter lists
- ✅ Type argument lists
- ✅ Multiple type parameters

### Type System
- ✅ Type inference for generic functions
- ✅ Type constraints/bounds
- ✅ Where clauses
- ✅ Type compatibility checking
- ✅ Nested generic types

### Advanced Features
- ✅ Methods on generic types (impl blocks)
- ✅ Generic references and mutation
- ✅ Pattern matching with generic enums
- ✅ Recursive generic structures
- ✅ Generic array/slice operations

## Expected Behavior

### Current State (Pre-Implementation)
The Go tests validate AST structure and will pass if the parser correctly builds generic AST nodes (which is already mostly in place in the grammar and AST).

The Unified tests will fail to compile until the full generics implementation is complete, including:
- Type parameter resolution
- Type inference
- Monomorphization (generating specialized versions)
- Generic method dispatch

### Post-Implementation
All tests should:
1. Parse correctly
2. Type-check correctly
3. Generate correct bytecode
4. Execute with expected results (return code 1 for success, 0 for failure)

## Notes

- Tests are designed to be minimal and focused
- Each test validates a specific aspect of generics
- Tests build on each other in complexity
- Tests use realistic patterns from other languages (Rust, Swift, etc.)
- Tests integrate with existing language features (enums, structs, pattern matching)
