# Test Baseline Report - February 2, 2026

**Date**: February 2, 2026  
**Compiler Version**: WASM Backend (wazero runtime)  
**Purpose**: Establish accurate testing baseline for progress tracking

---

## Executive Summary

### Test Suite Results

**`.uni` Integration Tests (test_all_uni_files.sh):**
- **Total Tests**: 123
- **Passing**: 56 (45.5%)
- **Failing**: 67 (54.5%)
- **Status**: ✅ Much better than previously documented (21.5%)

**Go Unit Tests (go test ./...):**
- **Total Packages**: 6
- **Passing Packages**: 5 (bytecode, parser, semantic, vm, wasm)
- **Failing Packages**: 1 (cmd/compiler - 3 integration tests fail)
- **Status**: ✅ Mostly passing, 3 known integration test failures

### Key Findings

1. **Actual test count is 123**, not 121 as documented in TODO.md and CLAUDE.md
2. **Actual pass rate is 45.5%**, significantly better than the 21.5% documented
3. **Test infrastructure is working correctly** - no timeout issues on Linux
4. **Documentation is severely outdated** - needs comprehensive update

---

## Detailed Test Results

### Integration Test Results (test_all_uni_files.sh)

#### Passing Tests (56 tests - 45.5%)

**Core Language Features (27 tests):**
- ✅ `minimal_test.uni` - Minimal program
- ✅ `simple_comparison_test.uni` - Comparisons
- ✅ `simple_test.uni` - Simple functions
- ✅ `bitwise.uni` - Bitwise operations
- ✅ `compound_assign.uni` - Compound assignments (+=, -=, etc.)
- ✅ `compound_assign_no_semi.uni` - Compound assignments without semicolons
- ✅ `counter_mut.uni` - Mutable counter (NOT timing out as documented!)
- ✅ `enum_option.uni` - Option enum
- ✅ `enum_result.uni` - Result enum
- ✅ `enum_simple.uni` - Simple enums
- ✅ `no_semicolons_multi.uni` - Multiple statements without semicolons
- ✅ `no_semicolons_simple.uni` - Simple no semicolons
- ✅ `point.uni` - Point struct
- ✅ `precedence.uni` - Operator precedence
- ✅ `rectangle.uni` - Rectangle struct
- ✅ `semicolons_all.uni` - All semicolons
- ✅ `semicolons_mixed.uni` - Mixed semicolons
- ✅ `simple_comparison.uni` - Simple comparisons
- ✅ `simple_precedence.uni` - Simple precedence
- ✅ `simple_test.uni` - Simple test (duplicate entry?)
- ✅ `type_inference.uni` - Type inference
- ✅ `function_call.uni` - Function calls
- ✅ `if_else.uni` - If/else statements
- ✅ `local_variables.uni` - Local variables
- ✅ `simple_return.uni` - Simple return
- ✅ `while_factorial.uni` - While loop factorial
- ✅ `range_test.uni` - Range operator

**Generics (5 tests):**
- ✅ `generics/13_generic_with_control_flow.uni`
- ✅ `generics/14_multiple_same_type_calls.uni`
- ✅ `generics/16_nested_calls.uni`
- ✅ `generics/19_arithmetic_ops.uni`
- ✅ `generics/20_complex_inference.uni`

**Arrays (8 tests):**
- ✅ `array_basics.uni` - Array literals and indexing
- ✅ `array_boundary.uni` - Array bounds checking
- ✅ `array_empty.uni` - Empty arrays
- ✅ `array_find_max.uni` - Array max finding
- ✅ `array_length.uni` - Array length
- ✅ `array_search.uni` - Array searching
- ✅ `array_sum.uni` - Array sum

**Strings (11 tests):**
- ✅ `string_array.uni` - String arrays
- ✅ `string_case.uni` - String case conversion
- ✅ `string_compare.uni` - String comparison
- ✅ `string_concat.uni` - String concatenation
- ✅ `string_demo_all.uni` - String demo
- ✅ `string_length.uni` - String length
- ✅ `string_search.uni` - String search
- ✅ `string_substring.uni` - String substring
- ✅ `string_trim.uni` - String trim

**For Loops (3 tests):**
- ✅ `for_sum.uni` - For loop sum
- ✅ `simple_for_loop.uni` - Simple for loop
- ✅ `simple_for_test.uni` - For loop test

**Advanced Features (2 tests):**
- ✅ `fizzbuzz_complete.uni` - Complete FizzBuzz
- ✅ `nested_loops.uni` - Nested loops with break/continue

**Try Operator (3 tests):**
- ✅ `try_operator_multiple_errors.uni`
- ✅ `try_operator_propagate_err.uni`
- ✅ `try_operator_short_circuit.uni`

#### Failing Tests (67 tests - 54.5%)

**Category 1: Parser Grammar Issues (24 tests)**

*Standard Library - Missing `Self`, `new`, `mut self`, struct field syntax:*
- ❌ `lib/List.uni` - Self, struct fields, methods
- ❌ `lib/Tree.uni` - Self, struct fields
- ❌ `stdlib/BinaryTree.uni` - Self, new, methods
- ❌ `stdlib/HashMap.uni` - Self, new, methods
- ❌ `stdlib/List.uni` - Self, new, methods
- ❌ `stdlib/Queue.uni` - Self, new, methods
- ❌ `stdlib/Set.uni` - Self, new, methods
- ❌ `stdlib/Stack.uni` - Self, new, methods
- ❌ `stdlib/binarytree_int.uni`
- ❌ `stdlib/binarytree_string.uni`
- ❌ `stdlib/hashmap_int.uni`
- ❌ `stdlib/hashmap_string.uni`
- ❌ `stdlib/list_float.uni`
- ❌ `stdlib/list_int.uni`
- ❌ `stdlib/list_string.uni`
- ❌ `stdlib/list_struct.uni`
- ❌ `stdlib/queue_float.uni`
- ❌ `stdlib/queue_int.uni`
- ❌ `stdlib/queue_string.uni`
- ❌ `stdlib/queue_struct.uni`
- ❌ `stdlib/set_int.uni`
- ❌ `stdlib/set_string.uni`
- ❌ `stdlib/stack_float.uni`
- ❌ `stdlib/stack_int.uni`
- ❌ `stdlib/stack_string.uni`
- ❌ `stdlib/stack_struct.uni`

**Category 2: Generics - Type Inference Issues (13 tests)**
- ❌ `generics/01_identity_function.uni`
- ❌ `generics/02_box_struct.uni`
- ❌ `generics/03_option_enum.uni`
- ❌ `generics/04_result_enum.uni`
- ❌ `generics/05_pair_function.uni`
- ❌ `generics/06_swap_function.uni`
- ❌ `generics/07_array_operations.uni`
- ❌ `generics/08_container_methods.uni`
- ❌ `generics/09_comparison_function.uni`
- ❌ `generics/10_linked_list.uni`
- ❌ `generics/11_explicit_type_args.uni`
- ❌ `generics/12_multiple_type_params.uni`
- ❌ `generics/15_different_type_calls.uni`
- ❌ `generics/17_local_variables.uni`
- ❌ `generics/18_same_type_params.uni`

**Category 3: Type Mismatch i32/i64 Issues (7 tests)**
- ❌ `add_test.uni` - Type mismatch i32/i64
- ❌ `basic_test.uni` - Type mismatch
- ❌ `call_test.uni` - Type mismatch
- ❌ `comparison_test.uni` - Type mismatch
- ❌ `simple_arith.uni` - Type mismatch
- ❌ `simple_call_test.uni` - Type mismatch
- ❌ `simple_comparison_test2.uni` - Type mismatch

**Category 4: Try Operator - Missing Implementation (6 tests)**
- ❌ `try_operator_basic_ok.uni`
- ❌ `try_operator_bool.uni`
- ❌ `try_operator_chained.uni`
- ❌ `try_operator_in_conditional.uni`
- ❌ `try_operator_in_expression.uni`
- ❌ `try_operator_nested_calls.uni`
- ❌ `try_operator_string.uni`

**Category 5: Advanced Features - Not Implemented (8 tests)**
- ❌ `block_expr.uni` - Block expressions not supported
- ❌ `fib.uni` - Complex recursion or other issues
- ❌ `fizz.uni` - FizzBuzz implementation issues
- ❌ `shadowing.uni` - Variable shadowing not supported
- ❌ `nested_structs.uni` - Nested struct field access issues
- ❌ `new_keyword_basic.uni` - new keyword issues
- ❌ `new_keyword.uni` - new keyword issues

**Category 6: Array & String Issues (3 tests)**
- ❌ `array_assignment.uni` - Array element assignment
- ❌ `array_double.uni` - Array double operations
- ❌ `array_reverse.uni` - Array reversal
- ❌ `string_comprehensive.uni` - Complex string operations
- ❌ `string_function.uni` - String function issues

---

## Go Unit Test Results

### Passing Packages (5/6)

1. **internal/bytecode** ✅
   - All bytecode instruction tests pass
   - Type definitions working correctly

2. **internal/parser** ✅
   - ANTLR parser generation works
   - AST construction successful

3. **internal/semantic** ✅
   - Symbol table tests pass
   - Type checking works for basic types

4. **internal/vm** ✅
   - All VM tests pass (legacy, but still maintained)
   - Stack operations, arithmetic, control flow work

5. **internal/wasm** ✅
   - WASM generator unit tests pass
   - Struct, array, string, for loop tests work
   - Enum, match expression tests work

### Failing Package (1/6)

**cmd/compiler** ❌ - 3 integration test failures:

1. **TestIntegrationForSum** ❌
   - Expected: exit code 55
   - Got: exit code 66
   - Issue: For loop range calculation incorrect

2. **TestIntegrationNestedLoops** ❌
   - Expected: exit code 18
   - Got: exit code 40
   - Issue: Nested loop calculation incorrect

3. **TestIntegrationStructNested** ❌
   - Expected: exit code 57
   - Got: exit code 1
   - Error: "type mismatch: expected i64, but was i32"
   - Issue: Nested struct field access type mismatch

---

## Error Category Analysis

### Error Type Distribution

| Error Type | Count | Percentage |
|------------|-------|------------|
| Parser grammar issues (Self, new, mut, fields) | 26 | 38.8% |
| Generic type inference | 13 | 19.4% |
| Type mismatch (i32/i64) | 7 | 10.4% |
| Try operator not implemented | 7 | 10.4% |
| Advanced features missing | 8 | 11.9% |
| Array/string operations | 6 | 9.0% |

### Root Cause Summary

1. **Parser Grammar Gaps** (26 tests, 38.8%)
   - Missing `Self` keyword support
   - Missing `new` method support
   - Missing `mut self` parameter syntax
   - Missing struct field shorthand syntax
   - **Impact**: Blocks all standard library tests

2. **Generic Type Inference** (13 tests, 19.4%)
   - Monomorphization not implemented
   - Type parameter inference incomplete
   - **Impact**: Advanced generic tests fail

3. **Type System Issues** (7 tests, 10.4%)
   - i32/i64 type mismatch in function calls
   - Type coercion missing
   - **Impact**: Simple test files fail

4. **Try Operator** (7 tests, 10.4%)
   - Parser recognizes `?` but codegen not complete
   - Missing Result/Option unwrapping logic
   - **Impact**: Error handling tests fail

5. **Missing Features** (14 tests, 20.9%)
   - Block expressions
   - Variable shadowing
   - Complex array operations
   - Complex string operations
   - Nested struct field access

---

## Test Infrastructure Status

### ✅ Working Infrastructure

1. **test_all_uni_files.sh**
   - ✅ Works on Linux (timeout command available)
   - ✅ 5-second timeout per test
   - ✅ Proper categorization (PASS/FAIL)
   - ✅ Generates useful reports
   - ⚠️ **Needs macOS compatibility** - timeout command doesn't exist on macOS

2. **Go Test Suite**
   - ✅ Unit tests comprehensive
   - ✅ Integration tests cover key features
   - ✅ Table-driven test approach
   - ✅ Excellent coverage of WASM generator

3. **Build System**
   - ✅ Makefile works correctly
   - ✅ Parser generation functional
   - ✅ Build automation solid

### ⚠️ Infrastructure Issues to Address

1. **macOS Compatibility**
   - Problem: `timeout` command doesn't exist on macOS
   - Solutions:
     - Option A: Install coreutils (`brew install coreutils`) and use `gtimeout`
     - Option B: Implement timeout in pure bash/Go
     - Option C: Remove timeout entirely (not recommended)
   - **Recommended**: Option A or create cross-platform wrapper

2. **Test Discovery**
   - Issue: Test count differs from documentation (123 vs 121)
   - Solution: Automated test counting in reports

3. **Documentation Sync**
   - Issue: Documentation severely out of date
   - Solution: Automated documentation generation from test results

---

## Comparison with Documentation

### Current Documentation Claims (TODO.md, CLAUDE.md, README.md)

- **Claimed Tests**: 121
- **Claimed Passing**: 26 (21.5%)
- **Claimed Failing**: 95 (78.5%)

### Actual Results (This Baseline)

- **Actual Tests**: 123 (+2 tests)
- **Actual Passing**: 56 (45.5%) (+30 tests, +24 percentage points!)
- **Actual Failing**: 67 (-28 tests)

### Documentation Accuracy Issues

1. ❌ Test count wrong (121 vs 123)
2. ❌ Pass rate severely understated (21.5% vs 45.5%)
3. ❌ Specific test status incorrect (many claimed failing tests now pass)
4. ❌ Feature status outdated (structs, arrays, strings, for loops work!)
5. ❌ counter_mut.uni claimed to timeout - actually passes!

---

## Recommendations

### Immediate Actions (Priority 0)

1. ✅ **Test Infrastructure Fixed** - Works on Linux
2. ⚠️ **Add macOS Support** - Implement cross-platform timeout
3. ✅ **Baseline Established** - This document
4. ⏳ **Update Documentation** - In progress

### Short-term Priorities (Priority 1)

Based on error analysis, prioritize:

1. **Parser Grammar Improvements** (26 tests, 38.8%)
   - Add `Self` keyword support
   - Add `new` method syntax
   - Add `mut self` parameter syntax
   - Add struct field shorthand syntax
   - **Impact**: Unlock all 26 standard library tests

2. **Fix Type System Issues** (7 tests, 10.4%)
   - Fix i32/i64 type mismatches
   - Add proper type coercion
   - **Impact**: Quick wins, simple fixes

3. **Complete Try Operator** (7 tests, 10.4%)
   - Implement codegen for `?` operator
   - Add Result/Option unwrapping
   - **Impact**: Error handling completeness

### Medium-term Priorities (Priority 2)

1. **Generic Monomorphization** (13 tests, 19.4%)
   - Implement type parameter instantiation
   - Generate specialized functions
   - **Impact**: Advanced type system features

2. **Advanced Features** (14 tests, 20.9%)
   - Block expressions
   - Variable shadowing
   - Complex array/string operations

---

## Test Files Summary

### Test Files Location

- **Integration Tests**: `src/unified-compiler/test/integration/` (27 files)
- **Generic Tests**: `src/unified-compiler/test/generics/` (20 files)
- **Standard Library Tests**: `src/unified-compiler/test/stdlib/` (26 files)
- **Root Tests**: `src/unified-compiler/test/` (remaining files)
- **Legacy Tests**: `src/unified-compiler/*.uni` (9 files)
- **Library Source**: `lib/*.uni` (2 files)
- **Stdlib Source**: `src/unified-compiler/stdlib/*.uni` (6 files)

### Total File Count: 123 .uni files

---

## Next Steps

### Documentation Updates Needed

1. **TODO.md**
   - Update test counts: 123 total, 56 passing (45.5%)
   - Correct feature status (many features work now!)
   - Update priority analysis based on error categories
   - Remove incorrect timeout claim for counter_mut.uni

2. **README.md**
   - Update badges: tests-56 passing-yellow
   - Update test results summary
   - Correct success rate to 45.5%

3. **CLAUDE.md**
   - Update test status
   - Correct pass rate
   - Update implementation status

4. **docs/PROJECT_STATUS.md**
   - Archive old status
   - Add reference to this baseline
   - Update current capabilities

---

## Conclusion

The Unified compiler is in **much better shape** than the documentation suggests:

- ✅ **45.5% of tests pass** (not 21.5% as documented)
- ✅ **Core features work**: functions, variables, control flow, arrays, strings, for loops, structs
- ✅ **Test infrastructure solid** (with minor macOS compatibility issue)
- ✅ **Clear path forward**: Focus on parser grammar (38.8% of failures)

**Biggest win**: Adding parser support for `Self`, `new`, `mut self`, and struct field syntax would immediately enable **26 more tests** (21% improvement) and unlock the entire standard library.

---

**Report Generated**: February 2, 2026  
**Generated By**: Copilot Code Agent  
**Next Review**: After Priority 1 grammar improvements
