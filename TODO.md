# TODO: WASM Backend Implementation Tasks

## Summary

This document tracks the implementation tasks needed to make all 121 test files pass with the WASM backend.

**Current Status:** 21 tests passing (17.4%), 100 tests failing (82.6%)

**Last Updated:** January 30, 2026 - After Priority 2 parser work

## Test Results Overview

- **Total Tests:** 121
- **✅ Passing:** 21 tests (17.4%)
- **❌ Failing:** 100 tests (82.6%)

## Priority 1 & 2 Status Summary

### Priority 1: Critical Language Features (29 tests) ⚠️ **PARTIAL**
- **Status:** Infrastructure complete, blocked on WASM code generation bugs
- **Completed:** WASM global section, heap allocator, struct registry
- **Blocked:** Field access type mismatch, array allocation issues

### Priority 2: Advanced Language Features (32 tests) ⚠️ **15% COMPLETE**
- **2.1 Generic Functions (15 tests):** ❌ Not started - needs monomorphization
- **2.2 Try Operator (10 tests):** 🟡 50% complete - parser done, codegen blocked
- **2.3 Standard Library (24 tests):** ❌ Not started - needs Self keyword, methods

**See PRIORITY2_WORK_SUMMARY.md for detailed analysis**

## Recent Changes (January 30, 2026)

### Priority 2 Work Session - Parser Infrastructure

#### ✅ Completed Infrastructure Improvements
1. **ANTLR 4.13.2 Installation**: Downloaded from Maven Central to avoid keyword conflicts
2. **Parser Regeneration**: Successfully regenerated all parser files
3. **Enum Constructor Support**: Added `Type::Variant(args)` syntax to grammar
4. **Visitor Implementation**: Implemented VisitEnumConstructorExpr in AST builder

#### 🟡 Partial Completions
1. **Enum Constructor Parsing**: ✅ Works - `Result::Ok(42)` now parses
2. **Try Operator Parsing**: ✅ Works - `expr?` recognized
3. **Enum Constructor Codegen**: ❌ Type mismatch errors
4. **Try Operator Codegen**: ❌ Not implemented

#### ❌ Outstanding Issues (From Priority 1)
1. **CRITICAL BLOCKER**: Field access type mismatch ("expected i32, but was i64")
2. **Regression**: 5 tests that were passing now fail (needs investigation)

## Working Features (21 tests passing) ✅

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

### 1.1 Struct Support (4 tests) 🔴 HIGH PRIORITY - **BLOCKED**

**Status:** Infrastructure complete, but blocked on field access type mismatch error

**Progress Made:**
- ✅ Added WASM global section encoding (was missing, causing "invalid index for global.get")
- ✅ Fixed heap pointer ULEB128 encoding (1024 = 0x80, 0x08, not 0x00, 0x04, 0x00)
- ✅ Created struct registry to track field names, types, and offsets
- ✅ Implemented first-pass struct declaration collection
- ✅ Fixed all WASM load/store memarg encoding (alignment and offset now use ULEB128)
- ✅ Implemented field offset lookup based on struct type
- ✅ Updated type inference to return actual field types
- ✅ Struct allocation works correctly

**Current Blocker:**
- ❌ Field access fails with "type mismatch: expected i32, but was i64"
- Works: `let p = Point { x: 42 }`
- Fails: `return p.x` (where x is type Int/i64)
- Root cause unclear despite extensive debugging

**Failing Tests:**
- `test/point.uni` - Basic struct with field access
- `test/rectangle.uni` - Struct with multiple fields
- `test/nested_structs.uni` - Nested struct access
- `test/new_keyword_basic.uni` - Struct initialization with `new`

**Error Pattern:**
```
Runtime error: failed to compile module: invalid function[0] export["main"]: 
type mismatch: expected i32, but was i64
```

**Next Steps:**
1. Debug WASM bytecode sequence for field access
2. Try minimal reproduction with WASM tools (wasm-tools, wabt)
3. Review type conversions in field access code path
4. Consider alternative implementations

**Related Files:**
- `internal/wasm/encoder.go` - Global section encoding
- `internal/wasm/generator.go` - Struct registry  
- `internal/wasm/codegen.go` - Field access codegen
- `internal/wasm/memory.go` - Heap allocation

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

**Next Review Date:** After Priority 1.1 struct field access blocker is resolved

---

## How to Proceed with Priority 2, 3, and 4

### ⚠️ CRITICAL: Priority 1 Must Be Completed First

**DO NOT** proceed with Priority 2, 3, or 4 until Priority 1 is **100% complete**. The struct field access blocker and array allocation issues must be resolved first.

### Priority 2: Advanced Language Features (32 tests) - **15% COMPLETE**

**Status**: Parser infrastructure complete, WASM codegen needs work

#### Current State After January 30 Work Session:
- ✅ **Parser Infrastructure**: ANTLR 4.13.2 installed, parser regenerated
- ✅ **Enum Constructor Syntax**: `Type::Variant(args)` parsing works
- ✅ **Try Operator Syntax**: `expr?` parsing works
- ❌ **Enum Constructor Codegen**: Type mismatch bugs (i32 vs i64)
- ❌ **Try Operator Codegen**: Not implemented
- ❌ **Generic Monomorphization**: Not started
- ❌ **Self Keyword**: Not added to grammar
- ❌ **Method Syntax**: Not added to grammar

#### Recommended Approach (After Priority 1 Complete):

**Step 1: Fix Enum Constructor WASM Codegen** (2-4 hours, 0-2 tests)
1. Debug type mismatch in `internal/wasm/codegen.go::generateEnumConstructor`
2. Ensure enum memory layout is correct
3. Fix i32/i64 pointer vs value type issues
4. Test with: `test/enum_simple.uni`, enum constructor tests

**Step 2: Implement Try Operator Codegen** (4-8 hours, 10 tests)
1. Implement `generateTryExpr` in `internal/wasm/codegen.go`
2. Add Result/Option enum variant detection
3. Generate early return pattern on Error/None variants
4. Implement value unwrapping for Ok/Some variants
5. Test with: `test/try_operator_*.uni` (10 tests)

**Step 3: Implement Generic Function Monomorphization** (2-3 days, 15 tests)
1. Track type arguments at function call sites in semantic analysis
2. Generate specialized function instances per type combination
3. Update type inference for generic return types
4. Handle generic constraints and bounds
5. Update function name mangling for monomorphized versions
6. Test with: `test/generics/*.uni` (15 tests)

**Step 4: Add Self Keyword and Method Syntax** (3-5 days, 24 tests)
1. Add `Self` keyword to lexer and parser (already in grammar as SELF)
2. Update type resolution to handle Self references
3. Add method syntax to grammar (`structDecl` member functions)
4. Implement method resolution in semantic analysis
5. Generate WASM code for method calls (implicit self parameter)
6. Test with: `test/stdlib/*.uni`, `lib/*.uni` (24 tests)

**Expected Result After Priority 2**: 70-72 tests passing (58-60%)

---

### Priority 3: Additional Features (15 tests)

**Status:** Not started - low priority

**Dependency**: Can start alongside Priority 2, but should focus on Priority 2 first

#### Recommended Approach:

**Block Expressions (3.1)**: Simple, 1 test
- Implement blocks as expressions (last expression is value)
- Low complexity, can be done anytime
- Test with: `test/block_expr.uni`

**Nested Loops (3.2)**: Medium, 1 test
- Ensure break/continue work in nested contexts
- May need label tracking
- Test with: `test/integration/nested_loops.uni`

**FizzBuzz (3.3)**: Should work automatically, 3 tests
- Likely will pass once for loops and modulo work
- Good integration test
- Test with: `test/fib.uni`, `test/fizz.uni`, `test/integration/fizzbuzz_complete.uni`

**Variable Shadowing (3.4)**: Medium, 1 test
- Track scopes properly in WASM locals
- Test with: `test/shadowing.uni`

**Simple Tests (3.5)**: Investigation needed, 10 tests
- May use deprecated syntax
- Review individually after priorities 1-2
- Files: `add_test.uni`, `basic_test.uni`, etc.

**Expected Result After Priority 3**: 85-87 tests passing (70-72%)

---

### Priority 4: Critical Bug Fixes (1 test)

**Status:** Not started - needs investigation

#### The Timeout Bug:
- `test/counter_mut.uni` enters infinite loop
- Could be compiler hang or runtime infinite loop
- High priority once found, but low impact (1 test)

#### Recommended Approach:
1. Add timeout protection to compiler itself
2. Debug with simple mutable counter test case
3. May be related to mutable variable handling

**Dependencies:** Can be investigated in parallel with other work

**Expected Result After Priority 4**: 86-88 tests passing (71-73%)

---

## Priority Completion Checklist

Use this checklist to track overall progress:

### Priority 1: Critical Language Features (29 tests)
- [ ] 1.1 Struct Support (4 tests) - **BLOCKED** on field access type mismatch
- [ ] 1.2 Array Support (11 tests) - Needs heap allocation fixes
- [ ] 1.3 For Loop & Range (4 tests) - Needs range operator in codegen
- [ ] 1.4 String Operations (10 tests) - Type mismatches, missing runtime functions

### Priority 2: Advanced Language Features (32 tests)
- [ ] 2.1 Generic Functions (15 tests) - Needs monomorphization
- [x] 2.2 Try Operator - Parser (50% complete)
- [ ] 2.2 Try Operator - Codegen (0% complete)
- [ ] 2.3 Standard Library (24 tests) - Needs Self, methods

### Priority 3: Additional Features (15 tests)
- [ ] 3.1 Block Expressions (1 test)
- [ ] 3.2 Nested Loops (1 test)
- [ ] 3.3 FizzBuzz (3 tests)
- [ ] 3.4 Variable Shadowing (1 test)
- [ ] 3.5 Simple Tests (10 tests)

### Priority 4: Critical Bug Fixes (1 test)
- [ ] 4.1 Infinite Loop/Timeout (1 test)

---

## Documentation Updates Needed

After completing each priority:

### After Priority 1 Complete:
- [ ] Update CLAUDE.md with struct/array implementation details
- [ ] Update test status in this TODO.md
- [ ] Document WASM memory layout for structs/arrays

### After Priority 2 Complete:
- [ ] Update CLAUDE.md with generic monomorphization approach
- [ ] Document try operator code generation pattern
- [ ] Update spec/UnifiedSpecification.md if needed

### After Priority 3 Complete:
- [ ] Update test status to reflect ~85 tests passing
- [ ] Review and update all example code

### After Priority 4 Complete:
- [ ] Final test suite run and documentation
- [ ] Update README.md with current capabilities
- [ ] Create migration guide for any breaking changes

---

**Maintainer Notes:**
- **DO NOT skip Priority 1** - it's foundational for all other work
- Test suite must be run after each major change
- Regressions must be addressed immediately  
- Update this TODO.md after each work session
- Refer to PRIORITY2_WORK_SUMMARY.md for detailed Priority 2 analysis

**Next Review Date:** After Priority 1 struct field access blocker is resolved
