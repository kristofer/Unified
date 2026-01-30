# TODO: WASM Backend Implementation Tasks

## Summary

This document tracks the implementation tasks needed to make all 121 test files pass with the WASM backend.

**Current Status:** 26 tests passing (21.5%), 95 tests failing (78.5%)

**Last Updated:** January 30, 2026

## Test Results Overview

- **Total Tests:** 121
- **✅ Passing:** 26 tests (21.5%)
- **❌ Failing:** 94 tests (77.7%)
- **⏱️ Timeout:** 1 test (0.8%)

## Working Features (26 tests passing) ✅

The following language features are fully functional with the WASM backend:

### Core Language Features
- ✅ Basic function declarations and calls
- ✅ Integer literals and arithmetic
- ✅ Boolean literals and logic operations
- ✅ Comparison operations
- ✅ Bitwise operations
- ✅ Local variables with `let`
- ✅ Mutable variables with `let mut`
- ✅ Variable assignment
- ✅ Compound assignment operators (`+=`, `-=`, etc.)
- ✅ Return statements
- ✅ If/else statements
- ✅ While loops
- ✅ Type inference for basic types
- ✅ Operator precedence
- ✅ Optional semicolons
- ✅ Basic enum variants (simple cases)
- ✅ Generic functions (basic cases without complex type inference)

### Passing Test Files
```
./src/unified-compiler/minimal_test.uni
./src/unified-compiler/simple_comparison_test.uni
./src/unified-compiler/simple_test.uni
./src/unified-compiler/test/bitwise.uni
./src/unified-compiler/test/compound_assign.uni
./src/unified-compiler/test/compound_assign_no_semi.uni
./src/unified-compiler/test/enum_option.uni
./src/unified-compiler/test/enum_result.uni
./src/unified-compiler/test/enum_simple.uni
./src/unified-compiler/test/generics/13_generic_with_control_flow.uni
./src/unified-compiler/test/generics/14_multiple_same_type_calls.uni
./src/unified-compiler/test/generics/16_nested_calls.uni
./src/unified-compiler/test/generics/19_arithmetic_ops.uni
./src/unified-compiler/test/generics/20_complex_inference.uni
./src/unified-compiler/test/integration/function_call.uni
./src/unified-compiler/test/integration/if_else.uni
./src/unified-compiler/test/integration/local_variables.uni
./src/unified-compiler/test/integration/simple_return.uni
./src/unified-compiler/test/integration/while_factorial.uni
./src/unified-compiler/test/no_semicolons_multi.uni
./src/unified-compiler/test/no_semicolons_simple.uni
./src/unified-compiler/test/precedence.uni
./src/unified-compiler/test/semicolons_all.uni
./src/unified-compiler/test/semicolons_mixed.uni
./src/unified-compiler/test/simple_precedence.uni
./src/unified-compiler/test/type_inference.uni
```

## Required Implementations (95 tests failing)

The following features need to be implemented to make the failing tests pass:

---

## Priority 1: Critical Language Features (29 tests)

### 1.1 Struct Support (4 tests) 🔴 HIGH PRIORITY

**Status:** Parser recognizes structs, but WASM codegen has bugs with global.get indices

**Failing Tests:**
- `test/point.uni` - Basic struct with field access
- `test/rectangle.uni` - Struct with multiple fields
- `test/nested_structs.uni` - Nested struct access
- `test/new_keyword_basic.uni` - Struct initialization with `new`

**Error Pattern:**
```
Runtime error: failed to compile module: invalid function[0] export["main"]: 
invalid index for global.get
```

**Implementation Required:**
- Fix WASM global variable indexing for struct field access
- Implement struct field offset calculations in WASM linear memory
- Support dot notation for field access (e.g., `point.x`)
- Implement `new` keyword for heap allocation
- Ensure proper memory alignment for struct fields

**Related Files:**
- `internal/wasm/codegen.go` - Need to fix global.get index calculation
- `internal/wasm/generator.go` - Need proper struct layout in memory

---

### 1.2 Array Support (11 tests) 🔴 HIGH PRIORITY

**Status:** Parser supports array literals, but WASM codegen fails with global.get errors

**Failing Tests:**
- `test/integration/array_basics.uni` - Array literals and indexing
- `test/integration/array_assignment.uni` - Array element assignment
- `test/integration/array_boundary.uni` - Array bounds checking
- `test/integration/array_double.uni` - Array element operations
- `test/integration/array_empty.uni` - Empty array handling
- `test/integration/array_find_max.uni` - Array iteration and max
- `test/integration/array_length.uni` - Array length property
- `test/integration/array_reverse.uni` - Array reversal
- `test/integration/array_search.uni` - Array searching
- `test/integration/array_sum.uni` - Array sum calculation
- `test/integration/new_keyword.uni` - Array with new keyword

**Error Pattern:**
```
Runtime error: failed to compile module: invalid function[0] export["main"]: 
invalid index for global.get
```

**Implementation Required:**
- Fix array memory allocation in WASM heap
- Implement array indexing with bounds checking
- Support array literals `[1, 2, 3]`
- Implement array length access
- Support array element assignment
- Proper memory layout: [length, capacity, ...elements]

**Related Files:**
- `internal/wasm/codegen.go` - Array operations
- `internal/wasm/generator.go` - Memory allocator for arrays

---

### 1.3 For Loop & Range Support (4 tests) 🔴 HIGH PRIORITY

**Status:** Parser doesn't support range operator `..` in binary expressions

**Failing Tests:**
- `test/integration/for_sum.uni` - For loop with range
- `test/integration/simple_for_loop.uni` - Basic for loop
- `test/integration/simple_for_test.uni` - For loop variations
- `test/integration/range_test.uni` - Range expressions

**Error Pattern:**
```
Error generating WASM: error adding function sum_range: 
unsupported binary operator: Range
```

**Implementation Required:**
- Add Range operator to AST binary expression handling
- Implement for-in loop code generation for WASM
- Support range syntax: `for i in 0..10 {}`
- Support inclusive ranges: `for i in 0..=10 {}`
- Properly handle loop variable scoping

**Related Files:**
- `internal/wasm/codegen.go` - Add range operator handling
- `internal/ast/ast.go` - Ensure Range is in BinaryOp enum

---

### 1.4 String Operations (10 tests) 🟡 MEDIUM PRIORITY

**Status:** String literals work, but operations have type mismatches and missing runtime functions

**Failing Tests:**
- `test/integration/string_length.uni` - String length
- `test/integration/string_compare.uni` - String comparison
- `test/integration/string_concat.uni` - String concatenation
- `test/integration/string_substring.uni` - String slicing
- `test/integration/string_search.uni` - String searching
- `test/integration/string_trim.uni` - String trimming
- `test/integration/string_case.uni` - Case conversion
- `test/integration/string_array.uni` - String arrays
- `test/integration/string_function.uni` - String functions
- `test/integration/string_comprehensive.uni` - Complex string ops
- `test/integration/string_demo_all.uni` - String demo

**Error Pattern:**
```
Runtime error: failed to compile module: invalid function[0] export["main"]: 
type mismatch: expected i64, but was i32
```

**Implementation Required:**
- Fix type mismatch between i32 (string length) and i64 (expected return)
- Implement string length runtime function
- Implement string comparison runtime function
- Implement string concatenation runtime function
- Implement string substring/slice function
- Implement string search functions
- Store strings in data section with length prefix
- Add string manipulation helper functions in WASM

**Related Files:**
- `internal/wasm/codegen.go` - String operations
- `internal/wasm/runtime.go` - Add string helper functions

---

## Priority 2: Advanced Language Features (32 tests)

### 2.1 Generic Functions (9 tests) 🟡 MEDIUM PRIORITY

**Status:** Basic generics work, but complex type inference and monomorphization fail

**Failing Tests:**
- `test/generics/01_identity_function.uni`
- `test/generics/02_box_struct.uni`
- `test/generics/03_option_enum.uni`
- `test/generics/04_result_enum.uni`
- `test/generics/05_pair_function.uni`
- `test/generics/06_swap_function.uni`
- `test/generics/07_array_operations.uni`
- `test/generics/08_container_methods.uni`
- `test/generics/09_comparison_function.uni`
- `test/generics/10_linked_list.uni`
- `test/generics/11_explicit_type_args.uni`
- `test/generics/12_multiple_type_params.uni`
- `test/generics/15_different_type_calls.uni`
- `test/generics/17_local_variables.uni`
- `test/generics/18_same_type_params.uni`

**Implementation Required:**
- Implement full generic type parameter inference
- Support monomorphization (generating separate function for each type)
- Handle generic structs and enums
- Support explicit type arguments `func::<T>()`
- Implement constraint checking for generic bounds
- Handle generic return types

**Related Files:**
- `internal/semantic/` - Type inference and monomorphization
- `internal/wasm/generator.go` - Generic function instantiation

---

### 2.2 Try Operator (?) (10 tests) 🟡 MEDIUM PRIORITY

**Status:** Parser doesn't recognize `::` for enum variants, try operator not implemented

**Failing Tests:**
- `test/try_operator_basic_ok.uni`
- `test/try_operator_bool.uni`
- `test/try_operator_chained.uni`
- `test/try_operator_in_conditional.uni`
- `test/try_operator_in_expression.uni`
- `test/try_operator_multiple_errors.uni`
- `test/try_operator_nested_calls.uni`
- `test/try_operator_propagate_err.uni`
- `test/try_operator_short_circuit.uni`
- `test/try_operator_string.uni`

**Error Pattern:**
```
extraneous input '::' expecting ...
Error generating WASM: error adding function: undefined variable: Result
```

**Implementation Required:**
- Add `::` operator to grammar for enum variant access (e.g., `Result::Ok`)
- Implement try operator `?` in parser and AST
- Generate WASM code for early return on Error variant
- Support Result<T, E> and Option<T> pattern matching
- Implement automatic error propagation

**Related Files:**
- `grammar/UnifiedParser.g4` - Add `::` operator
- `internal/ast/ast.go` - Add TryExpression node
- `internal/wasm/codegen.go` - Try operator code generation

---

### 2.3 Standard Library Collections (24 tests) 🟢 LOW PRIORITY

**Status:** Library code uses unsupported syntax (struct fields, methods, Self)

**Failing Tests:**

**List (5 tests):**
- `lib/List.uni`
- `src/unified-compiler/stdlib/List.uni`
- `test/stdlib/list_int.uni`
- `test/stdlib/list_float.uni`
- `test/stdlib/list_string.uni`
- `test/stdlib/list_struct.uni`

**BinaryTree (3 tests):**
- `lib/Tree.uni`
- `src/unified-compiler/stdlib/BinaryTree.uni`
- `test/stdlib/binarytree_int.uni`
- `test/stdlib/binarytree_string.uni`

**HashMap (3 tests):**
- `src/unified-compiler/stdlib/HashMap.uni`
- `test/stdlib/hashmap_int.uni`
- `test/stdlib/hashmap_string.uni`

**Queue (5 tests):**
- `src/unified-compiler/stdlib/Queue.uni`
- `test/stdlib/queue_int.uni`
- `test/stdlib/queue_float.uni`
- `test/stdlib/queue_string.uni`
- `test/stdlib/queue_struct.uni`

**Set (2 tests):**
- `src/unified-compiler/stdlib/Set.uni`
- `test/stdlib/set_int.uni`
- `test/stdlib/set_string.uni`

**Stack (5 tests):**
- `src/unified-compiler/stdlib/Stack.uni`
- `test/stdlib/stack_int.uni`
- `test/stdlib/stack_float.uni`
- `test/stdlib/stack_string.uni`
- `test/stdlib/stack_struct.uni`

**Error Pattern:**
```
extraneous input 'data' expecting ':'
mismatched input 'new' expecting Identifier
mismatched input 'Self' expecting ...
extraneous input 'mut' expecting {')', SELF, Identifier}
```

**Implementation Required:**
- Support struct field syntax without explicit type in struct body
- Implement `Self` keyword for referring to current type
- Support `new` as a constructor method name
- Implement method syntax (struct methods)
- Support `mut self` parameters
- These require substantial parser and semantic analysis work

**Related Files:**
- `grammar/UnifiedParser.g4` - Struct field syntax, Self keyword
- `internal/semantic/` - Method resolution
- `internal/wasm/` - Method calling convention

---

## Priority 3: Additional Features (15 tests)

### 3.1 Block Expressions (1 test) 🟢 LOW PRIORITY

**Failing Test:**
- `test/block_expr.uni`

**Implementation Required:**
- Support blocks as expressions that return values
- Last expression in block becomes the block's value
- No explicit return needed for block expressions

---

### 3.2 Nested Loops with Break/Continue (1 test) 🟡 MEDIUM PRIORITY

**Failing Test:**
- `test/integration/nested_loops.uni`

**Error:** Unknown - needs investigation

**Implementation Required:**
- Ensure break/continue work in nested loops
- Support labeled loops for multi-level break/continue

---

### 3.3 FizzBuzz Implementation (2 tests) 🟢 LOW PRIORITY

**Failing Tests:**
- `test/fib.uni` - Fibonacci implementation
- `test/fizz.uni` - FizzBuzz implementation
- `test/integration/fizzbuzz_complete.uni`

**Implementation Required:**
- These likely fail due to missing features above (for loops, modulo, etc.)
- Should work once Priority 1 & 2 features are complete

---

### 3.4 Variable Shadowing (1 test) 🟢 LOW PRIORITY

**Failing Test:**
- `test/shadowing.uni`

**Implementation Required:**
- Support variable shadowing in nested scopes
- Ensure proper scope tracking in WASM local variables

---

### 3.5 Simple Tests with Unknown Issues (10 tests) 🟢 LOW PRIORITY

**Failing Tests:**
- `add_test.uni`
- `basic_test.uni`
- `call_test.uni`
- `comparison_test.uni`
- `simple_arith.uni`
- `simple_call_test.uni`
- `simple_comparison_test2.uni`

These are older test files that may use deprecated syntax or have other issues.
Need individual investigation.

---

## Priority 4: Critical Bug Fixes

### 4.1 Infinite Loop/Timeout (1 test) ⏱️ CRITICAL

**Failing Test:**
- `test/counter_mut.uni` (TIMEOUT)

**Issue:** Compiler or runtime enters infinite loop when processing this file

**Implementation Required:**
- Debug why compilation or execution hangs
- Add timeout protection to prevent infinite loops
- Investigate mutable counter implementation

---

## Implementation Roadmap

### Phase 1 (Immediate - Critical Features)
1. Fix struct global.get index bug (4 tests)
2. Fix array global.get index bug (11 tests)
3. Implement for loop and range operator (4 tests)
4. Fix string type mismatches (11 tests)

**Expected Result:** 56 tests passing (46%)

### Phase 2 (Short-term - Advanced Features)
1. Implement try operator and `::` syntax (10 tests)
2. Improve generic type inference (15 tests)
3. Fix nested loops and break/continue (1 test)

**Expected Result:** 82 tests passing (68%)

### Phase 3 (Medium-term - Library Support)
1. Add Self keyword support
2. Add method syntax
3. Fix struct field syntax
4. Implement standard library collections (24 tests)

**Expected Result:** 106 tests passing (88%)

### Phase 4 (Long-term - Polish)
1. Block expressions (1 test)
2. Variable shadowing (1 test)
3. Fix simple test files (10 tests)
4. Debug timeout issue (1 test)

**Expected Result:** 118+ tests passing (97%+)

---

## Technical Debt & Known Issues

### Type System Issues
- Type mismatch between i32 and i64 in various contexts
- String operations return i32 but functions expect i64
- Need systematic type coercion/conversion

### Memory Management
- Global variable indexing is broken for complex types
- Memory allocator needs improvement
- Heap allocation for structs/arrays needs proper implementation

### Parser Limitations
- `::` operator not supported (needed for enum variants)
- Range operator `..` not fully integrated
- Self keyword not implemented
- Method syntax not supported

### WASM Code Generation
- Many features generate incorrect indices
- Missing runtime helper functions for strings
- Need better error messages for unsupported features

---

## Testing Infrastructure

### Test Execution
A comprehensive test runner `test_all_uni_files.sh` has been created that:
- Tests all 121 .uni files
- Has 5-second timeout per test
- Categorizes results as PASS, FAIL, or TIMEOUT
- Generates detailed reports

### Test Reports
- `test_results.txt` - Full output of all tests
- `test_working.txt` - List of passing tests
- `test_failing.txt` - List of failing tests

---

## Notes

- The WASM backend is functional for basic language features (26 tests passing)
- Most failures are due to missing advanced features, not fundamental architecture issues
- The migration from VM to WASM is sound; implementation just needs to catch up
- Once structs and arrays are fixed (Priority 1.1 & 1.2), many other tests will likely pass

---

## Maintenance

This TODO should be updated as features are implemented:
1. Move completed items to a "Completed" section
2. Update test counts
3. Re-run test suite periodically
4. Document any new issues discovered

**Next Review Date:** After Priority 1 items are complete
