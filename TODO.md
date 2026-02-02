# TODO: WASM Backend Implementation Tasks

## Summary

This document tracks the implementation tasks needed to make all 123 test files pass with the WASM backend.

**Current Status:** 56 tests passing (45.5%), 67 tests failing (54.5%)

**Last Updated:** February 2, 2026 - After comprehensive test audit (see TEST_BASELINE_2026-02-02.md)

## Test Results Overview

- **Total Tests:** 123
- **✅ Passing:** 56 tests (45.5%)
- **❌ Failing:** 67 tests (54.5%)

**See TEST_BASELINE_2026-02-02.md for comprehensive analysis**

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

## Working Features (56 tests passing) ✅

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
- ✅ For loops with ranges ✅ **WORKING**
- ✅ Break/continue in loops ✅ **WORKING**
- ✅ Type inference for basic types
- ✅ Operator precedence
- ✅ Optional semicolons
- ✅ Basic enum variants
- ✅ Generic functions (basic cases)
- ✅ Structs with field access ✅ **WORKING**
- ✅ Arrays (literals, indexing, iteration) ✅ **WORKING**
- ✅ Strings (length, concat, trim, case, search, substring) ✅ **WORKING**

### Passing Test Files (56 tests)
```
./src/unified-compiler/minimal_test.uni
./src/unified-compiler/simple_comparison_test.uni
./src/unified-compiler/simple_test.uni
./src/unified-compiler/test/bitwise.uni
./src/unified-compiler/test/compound_assign.uni
./src/unified-compiler/test/compound_assign_no_semi.uni
./src/unified-compiler/test/counter_mut.uni
./src/unified-compiler/test/enum_option.uni
./src/unified-compiler/test/enum_result.uni
./src/unified-compiler/test/enum_simple.uni
./src/unified-compiler/test/generics/13_generic_with_control_flow.uni
./src/unified-compiler/test/generics/14_multiple_same_type_calls.uni
./src/unified-compiler/test/generics/16_nested_calls.uni
./src/unified-compiler/test/generics/19_arithmetic_ops.uni
./src/unified-compiler/test/generics/20_complex_inference.uni
./src/unified-compiler/test/integration/array_basics.uni
./src/unified-compiler/test/integration/array_boundary.uni
./src/unified-compiler/test/integration/array_empty.uni
./src/unified-compiler/test/integration/array_find_max.uni
./src/unified-compiler/test/integration/array_length.uni
./src/unified-compiler/test/integration/array_search.uni
./src/unified-compiler/test/integration/array_sum.uni
./src/unified-compiler/test/integration/fizzbuzz_complete.uni
./src/unified-compiler/test/integration/for_sum.uni
./src/unified-compiler/test/integration/function_call.uni
./src/unified-compiler/test/integration/if_else.uni
./src/unified-compiler/test/integration/local_variables.uni
./src/unified-compiler/test/integration/nested_loops.uni
./src/unified-compiler/test/integration/range_test.uni
./src/unified-compiler/test/integration/simple_for_loop.uni
./src/unified-compiler/test/integration/simple_for_test.uni
./src/unified-compiler/test/integration/simple_return.uni
./src/unified-compiler/test/integration/string_array.uni
./src/unified-compiler/test/integration/string_case.uni
./src/unified-compiler/test/integration/string_compare.uni
./src/unified-compiler/test/integration/string_concat.uni
./src/unified-compiler/test/integration/string_demo_all.uni
./src/unified-compiler/test/integration/string_length.uni
./src/unified-compiler/test/integration/string_search.uni
./src/unified-compiler/test/integration/string_substring.uni
./src/unified-compiler/test/integration/string_trim.uni
./src/unified-compiler/test/integration/while_factorial.uni
./src/unified-compiler/test/no_semicolons_multi.uni
./src/unified-compiler/test/no_semicolons_simple.uni
./src/unified-compiler/test/point.uni
./src/unified-compiler/test/precedence.uni
./src/unified-compiler/test/rectangle.uni
./src/unified-compiler/test/semicolons_all.uni
./src/unified-compiler/test/semicolons_mixed.uni
./src/unified-compiler/test/simple_comparison.uni
./src/unified-compiler/test/simple_precedence.uni
./src/unified-compiler/test/simple_test.uni
./src/unified-compiler/test/try_operator_multiple_errors.uni
./src/unified-compiler/test/try_operator_propagate_err.uni
./src/unified-compiler/test/try_operator_short_circuit.uni
./src/unified-compiler/test/type_inference.uni
```

## Required Implementations (67 tests failing)

The following features need to be implemented to make the failing tests pass:

**See TEST_BASELINE_2026-02-02.md for detailed error analysis and categorization**

---

## Priority 0: Test Infrastructure ✅ COMPLETE

### Establish Accurate Testing Baseline ✅

**Status:** Complete - see TEST_BASELINE_2026-02-02.md

**Findings:**
- ✅ Test infrastructure works correctly on Linux
- ✅ 123 total tests (not 121 as previously documented)
- ✅ 56 passing (45.5%, not 21.5%!)
- ⚠️ macOS compatibility issue: `timeout` command doesn't exist
- ✅ Comprehensive baseline report created

**macOS Compatibility:**
- Issue: `timeout` command in `test_all_uni_files.sh` doesn't exist on macOS
- Solutions available:
  - Option A: Install coreutils (`brew install coreutils`) and use `gtimeout`
  - Option B: Create cross-platform timeout wrapper
  - Option C: Remove timeout (not recommended)
- Recommended: Document Option A for macOS users

---

## Priority 1: Critical Language Features (Updated Based on Baseline)

**MAJOR UPDATE:** Based on TEST_BASELINE_2026-02-02.md analysis, priority order has changed significantly.

### NEW Priority 1.1: Parser Grammar Improvements (26 tests) 🔴 HIGHEST PRIORITY

**Status:** Not started - **BIGGEST IMPACT OPPORTUNITY**

**Error Category:** 38.8% of all failures are parser grammar issues

**Failing Tests:** All 26 standard library tests fail due to missing parser features
- `lib/List.uni`, `lib/Tree.uni`
- `stdlib/BinaryTree.uni`, `stdlib/HashMap.uni`, `stdlib/List.uni`, `stdlib/Queue.uni`, `stdlib/Set.uni`, `stdlib/Stack.uni`
- All `test/stdlib/*.uni` tests (18 tests)

**Implementation Required:**
- Add `Self` keyword to parser (refers to current type in methods)
- Add `new` method syntax support
- Add `mut self` parameter syntax
- Add struct field shorthand syntax (fields without explicit types in body)
- Support method definitions inside structs

**Impact:** Unlocking these 26 tests = +21% pass rate improvement!

**Related Files:**
- `grammar/UnifiedParser.g4` - Add grammar rules
- `grammar/UnifiedLexer.g4` - Add keywords if needed
- `internal/ast/ast.go` - Add AST nodes for new syntax
- `internal/semantic/` - Method resolution

---

### 1.1 Struct Support (2 tests) 🟢 MOSTLY WORKING

**Status:** ✅ Structs work! Only 2 edge case failures remaining

**Progress Made:**
- ✅ Basic struct allocation works
- ✅ Field access works (point.uni, rectangle.uni pass!)
- ✅ WASM global section encoding fixed
- ✅ Heap allocation works
- ❌ Nested struct field access has type mismatch (1 test)
- ❌ new keyword with structs (1 test)

**Passing Tests:**
- `test/point.uni` ✅
- `test/rectangle.uni` ✅

**Failing Tests:**
- `test/nested_structs.uni` - Nested field access type mismatch (i64 vs i32)
- `test/new_keyword_basic.uni` - new keyword support

**Next Steps:**
1. Fix nested struct field access type issue
2. Add new keyword support (likely parser grammar issue)

---

### 1.2 Array Support (3 tests) 🟢 MOSTLY WORKING

**Status:** ✅ Arrays work! Only 3 edge case failures remaining

**Progress Made:**
- ✅ Array literals work
- ✅ Array indexing works
- ✅ Array iteration works
- ✅ Array length access works
- ✅ Bounds checking works
- ❌ Array element assignment (1 test)
- ❌ Array operations with doubles (1 test)
- ❌ Array reversal (1 test)

**Passing Tests (8 tests):**
- `test/integration/array_basics.uni` ✅
- `test/integration/array_boundary.uni` ✅
- `test/integration/array_empty.uni` ✅
- `test/integration/array_find_max.uni` ✅
- `test/integration/array_length.uni` ✅
- `test/integration/array_search.uni` ✅
- `test/integration/array_sum.uni` ✅

**Failing Tests (3 tests):**
- `test/integration/array_assignment.uni` - Element assignment
- `test/integration/array_double.uni` - Double type operations
- `test/integration/array_reverse.uni` - Reversal algorithm
- `test/integration/new_keyword.uni` - new keyword with arrays

**Next Steps:**
1. Implement array element assignment
2. Fix double type handling
3. Debug array reversal issue

---

### 1.3 For Loop & Range Support ✅ WORKING

**Status:** ✅ For loops and ranges work!

**All 3 tests passing:**
- `test/integration/for_sum.uni` ✅
- `test/integration/simple_for_loop.uni` ✅
- `test/integration/simple_for_test.uni` ✅
- `test/integration/range_test.uni` ✅

**No action needed** - Feature complete!

---

### 1.4 String Operations ✅ MOSTLY WORKING

**Status:** ✅ Strings work! Only 2 edge case failures

**Progress Made:**
- ✅ String literals work
- ✅ String length works
- ✅ String comparison works
- ✅ String concatenation works
- ✅ String substring works
- ✅ String search works
- ✅ String trim works
- ✅ String case conversion works
- ❌ Complex string operations (1 test)
- ❌ String functions (1 test)

**Passing Tests (11 tests):**
- `test/integration/string_array.uni` ✅
- `test/integration/string_case.uni` ✅
- `test/integration/string_compare.uni` ✅
- `test/integration/string_concat.uni` ✅
- `test/integration/string_demo_all.uni` ✅
- `test/integration/string_length.uni` ✅
- `test/integration/string_search.uni` ✅
- `test/integration/string_substring.uni` ✅
- `test/integration/string_trim.uni` ✅

**Failing Tests (2 tests):**
- `test/integration/string_comprehensive.uni` - Complex operations
- `test/integration/string_function.uni` - Function handling

**Next Steps:**
1. Debug comprehensive string test
2. Fix string function test

---

### NEW Priority 1.5: Type System Issues (7 tests) 🟡 MEDIUM PRIORITY

**Status:** Type mismatch errors in simple test files

**Error Pattern:** "type mismatch: expected i32, but was i64" on function calls

**Failing Tests:**
- `src/unified-compiler/add_test.uni`
- `src/unified-compiler/basic_test.uni`
- `src/unified-compiler/call_test.uni`
- `src/unified-compiler/comparison_test.uni`
- `src/unified-compiler/simple_arith.uni`
- `src/unified-compiler/simple_call_test.uni`
- `src/unified-compiler/simple_comparison_test2.uni`

**Implementation Required:**
- Fix i32/i64 type coercion in function parameters
- Ensure consistent type usage across WASM generator
- May be related to parameter passing convention

**Impact:** Quick wins, relatively simple fixes

**Related Files:**
- `internal/wasm/codegen.go` - Function call type handling

---

## Priority 2: Advanced Language Features (Updated Based on Baseline)

### 2.1 Generic Functions (13 tests) 🟡 MEDIUM PRIORITY

**Status:** Basic generics work (5 tests pass), advanced inference fails (13 tests fail)

**Passing Tests (5 tests):**
- `test/generics/13_generic_with_control_flow.uni` ✅
- `test/generics/14_multiple_same_type_calls.uni` ✅
- `test/generics/16_nested_calls.uni` ✅
- `test/generics/19_arithmetic_ops.uni` ✅
- `test/generics/20_complex_inference.uni` ✅

**Failing Tests (13 tests):**
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

### 2.2 Try Operator (?) (7 tests) 🟡 MEDIUM PRIORITY

**Status:** Parser recognizes `?` and enum constructors, 3 tests pass, 7 tests fail

**Passing Tests (3 tests):**
- `test/try_operator_multiple_errors.uni` ✅
- `test/try_operator_propagate_err.uni` ✅
- `test/try_operator_short_circuit.uni` ✅

**Failing Tests (7 tests):**
- `test/try_operator_basic_ok.uni`
- `test/try_operator_bool.uni`
- `test/try_operator_chained.uni`
- `test/try_operator_in_conditional.uni`
- `test/try_operator_in_expression.uni`
- `test/try_operator_nested_calls.uni`
- `test/try_operator_string.uni`

**Progress Made:**
- ✅ Parser recognizes `?` operator
- ✅ Enum constructor syntax `Type::Variant(args)` works
- ✅ Some basic try operator cases work
- ❌ Complex try operator expressions fail

**Implementation Required:**
- Complete try operator WASM codegen for all cases
- Handle try operator in complex expressions
- Support chained try operators
- Implement proper error unwrapping

**Related Files:**
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

## Priority 3: Additional Features (Updated Based on Baseline)

### 3.1 Block Expressions (1 test) 🟢 LOW PRIORITY

**Failing Test:**
- `test/block_expr.uni`

**Error:** "unsupported expression type: *ast.Block"

**Implementation Required:**
- Support blocks as expressions that return values
- Last expression in block becomes the block's value
- No explicit return needed for block expressions

---

### 3.2 Variable Shadowing (1 test) 🟢 LOW PRIORITY

**Failing Test:**
- `test/shadowing.uni`

**Implementation Required:**
- Support variable shadowing in nested scopes
- Ensure proper scope tracking in WASM local variables

---

### 3.3 FizzBuzz and Other Programs (2 tests) 🟢 WORKING/LOW PRIORITY

**Status:** FizzBuzz complete works! ✅

**Passing:**
- `test/integration/fizzbuzz_complete.uni` ✅

**Failing Tests:**
- `test/fib.uni` - Fibonacci implementation (needs investigation)
- `test/fizz.uni` - FizzBuzz implementation (needs investigation)

**Implementation Required:**
- Investigate why simpler fib/fizz tests fail when fizzbuzz_complete works
- May use deprecated syntax or have other issues

---

## Priority 4: Critical Bug Fixes

### 4.1 Infinite Loop/Timeout ✅ RESOLVED

**Status:** ✅ counter_mut.uni PASSES - no timeout!

**Resolution:**
- Previous documentation claimed this test timed out
- Comprehensive testing shows it passes successfully
- No infinite loop issue exists

**This issue is resolved and can be removed from priority list.**

---

## Implementation Roadmap (Updated February 2, 2026)

**See TEST_BASELINE_2026-02-02.md for comprehensive analysis**

### Current Status: 56/123 tests passing (45.5%)

### Phase 1 (Immediate - Highest Impact) - Parser Grammar

**Focus:** Add missing parser features (26 tests, 21% improvement)

1. Add `Self` keyword support
2. Add `new` method syntax
3. Add `mut self` parameter syntax
4. Add struct field shorthand syntax

**Expected Result:** 82 tests passing (67%)

### Phase 2 (Short-term - Edge Cases)

**Focus:** Fix remaining edge cases and type issues (10 tests)

1. Fix type system i32/i64 issues (7 tests)
2. Fix nested struct field access (1 test)
3. Fix array edge cases (3 tests)
4. Fix string edge cases (2 tests)

**Expected Result:** 92 tests passing (75%)

### Phase 3 (Medium-term - Advanced Features)

**Focus:** Generics and try operator (20 tests)

1. Implement generic monomorphization (13 tests)
2. Complete try operator codegen (7 tests)

**Expected Result:** 112 tests passing (91%)

### Phase 4 (Long-term - Polish)

**Focus:** Remaining features (11 tests)

1. Block expressions (1 test)
2. Variable shadowing (1 test)
3. Investigate simple test failures (2 tests)

**Expected Result:** 120+ tests passing (97%+)

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

## Notes (Updated February 2, 2026)

- The WASM backend is **much more functional** than previously documented
- **56 tests passing (45.5%)**, not 26 (21.5%) as documented
- **Major features work**: functions, variables, control flow, arrays, strings, for loops, structs, enums
- Most failures (38.8%) are due to **parser grammar gaps** (Self, new, mut self, field syntax)
- **Biggest opportunity**: Adding parser support would unlock 26 tests (+21% improvement)
- The migration from VM to WASM is sound; many features already implemented
- Once parser grammar is fixed, ~82 tests will pass (67%)

**Key Insight:** Focus on parser grammar first (Priority 1.1) for maximum impact!

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
