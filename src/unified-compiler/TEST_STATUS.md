# Test Status Report

Generated: 2026-01-27

## Summary
- **Total Tests**: 103
- **Passing**: 49 (47.6%)
- **Failing**: 54 (52.4%)

## Test Results by Category

### ✓ Passing Tests (49)

#### Basic Language Features (15 tests)
- `test/bitwise.uni` ✓
- `test/block_expr.uni` ✓
- `test/compound_assign.uni` ✓
- `test/compound_assign_no_semi.uni` ✓
- `test/counter_mut.uni` ✓
- `test/enum_option.uni` ✓ (Note: doesn't actually test enum functionality)
- `test/enum_result.uni` ✓ (Note: doesn't actually test enum functionality)
- `test/enum_simple.uni` ✓ (Note: doesn't actually test enum functionality)
- `test/nested_structs.uni` ✓
- `test/no_semicolons_multi.uni` ✓
- `test/no_semicolons_simple.uni` ✓
- `test/point.uni` ✓
- `test/precedence.uni` ✓
- `test/rectangle.uni` ✓
- `test/semicolons_all.uni` ✓
- `test/semicolons_mixed.uni` ✓
- `test/shadowing.uni` ✓
- `test/simple_precedence.uni` ✓
- `test/type_inference.uni` ✓

#### Integration Tests (23/30 passing)
- `test/integration/array_basics.uni` ✓
- `test/integration/array_boundary.uni` ✓
- `test/integration/array_empty.uni` ✓
- `test/integration/array_find_max.uni` ✓
- `test/integration/array_length.uni` ✓
- `test/integration/array_search.uni` ✓
- `test/integration/array_sum.uni` ✓
- `test/integration/fizzbuzz_complete.uni` ✓
- `test/integration/for_sum.uni` ✓
- `test/integration/function_call.uni` ✓
- `test/integration/if_else.uni` ✓
- `test/integration/local_variables.uni` ✓
- `test/integration/nested_loops.uni` ✓
- `test/integration/simple_for_loop.uni` ✓
- `test/integration/simple_for_test.uni` ✓
- `test/integration/simple_return.uni` ✓
- `test/integration/string_array.uni` ✓
- `test/integration/string_case.uni` ✓
- `test/integration/string_compare.uni` ✓
- `test/integration/string_comprehensive.uni` ✓
- `test/integration/string_concat.uni` ✓
- `test/integration/string_demo_all.uni` ✓
- `test/integration/string_function.uni` ✓
- `test/integration/string_length.uni` ✓
- `test/integration/string_search.uni` ✓
- `test/integration/string_substring.uni` ✓
- `test/integration/string_trim.uni` ✓
- `test/integration/while_factorial.uni` ✓

#### Generics Tests (2/20 passing)
- `test/generics/06_swap_function.uni` ✓
- `test/generics/09_comparison_function.uni` ✓

### ✗ Failing Tests (54)

#### Category 1: Missing Feature - Method Call Syntax (24 tests)
**Issue**: `Type.method()` and `obj.method()` syntax not implemented (Phase 1 limitation)
**Error**: "only direct function calls are supported in Phase 1"

Affected tests:
- `test/fib.uni` - Uses `List<Int>.new()`, `sequence.push()`, `seq.join()`
- `test/fizz.uni` - Uses `List<Int>.new()`, `numbers.push()`, `nums.join()`
- `test/integration/array_assignment.uni` - Uses array methods
- `test/integration/array_double.uni` - Uses array methods
- `test/integration/array_reverse.uni` - Uses array methods
- All 18 stdlib tests (require method syntax):
  - `test/stdlib/binarytree_int.uni`
  - `test/stdlib/binarytree_string.uni`
  - `test/stdlib/hashmap_int.uni`
  - `test/stdlib/hashmap_string.uni`
  - `test/stdlib/list_float.uni`
  - `test/stdlib/list_int.uni`
  - `test/stdlib/list_string.uni`
  - `test/stdlib/list_struct.uni`
  - `test/stdlib/queue_float.uni`
  - `test/stdlib/queue_int.uni`
  - `test/stdlib/queue_string.uni`
  - `test/stdlib/queue_struct.uni`
  - `test/stdlib/set_int.uni`
  - `test/stdlib/set_string.uni`
  - `test/stdlib/stack_float.uni`
  - `test/stdlib/stack_int.uni`
  - `test/stdlib/stack_string.uni`
  - `test/stdlib/stack_struct.uni`

#### Category 2: Missing Feature - Enum Variant Syntax (10 tests)
**Issue**: `EnumType::Variant()` syntax not supported in expressions
**Error**: "extraneous input '::' expecting {...}"

Affected tests:
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

#### Category 3: Generic Function Issues (18 tests)
**Issue**: Generic functions either have runtime issues or missing features

Sub-issues:
- Runtime error "no call frame for local variable access" (several tests)
- Generic type instantiation not working properly
- Generic method calls not supported

Affected tests:
- `test/generics/01_identity_function.uni` - Runtime: no call frame
- `test/generics/02_box_struct.uni`
- `test/generics/03_option_enum.uni`
- `test/generics/04_result_enum.uni`
- `test/generics/05_pair_function.uni`
- `test/generics/07_array_operations.uni`
- `test/generics/08_container_methods.uni`
- `test/generics/10_linked_list.uni`
- `test/generics/11_explicit_type_args.uni`
- `test/generics/12_multiple_type_params.uni`
- `test/generics/13_generic_with_control_flow.uni`
- `test/generics/14_multiple_same_type_calls.uni`
- `test/generics/15_different_type_calls.uni`
- `test/generics/16_nested_calls.uni`
- `test/generics/17_local_variables.uni`
- `test/generics/18_same_type_params.uni`
- `test/generics/19_arithmetic_ops.uni`
- `test/generics/20_complex_inference.uni`

#### Category 4: New Keyword with Generics (2 tests)
**Issue**: `new Type<T>()` syntax parsing issues

Affected tests:
- `test/new_keyword_basic.uni`
- `test/integration/new_keyword.uni`

#### Category 5: Range Expression Outside For Loop (1 test)
**Issue**: Range expressions (`..`, `..=`) only work in for loops
**Error**: "range expressions can only be used in for loops"

Affected tests:
- `test/integration/range_test.uni`

## Implementation Status

### Fully Implemented Features
- ✅ Basic arithmetic and logical operations
- ✅ Variable declarations (let/var, mut)
- ✅ Function declarations and direct calls
- ✅ Control flow (if/else, loops, for with ranges)
- ✅ Arrays (basic operations without methods)
- ✅ Strings (basic operations)
- ✅ Structs (definition and instantiation)
- ✅ Basic enums (definition only, not variant construction)
- ✅ Compound assignments (+=, -=, etc.)
- ✅ Bitwise operations
- ✅ Type inference (basic)
- ✅ Semicolon optional syntax

### Partially Implemented Features
- ⚠️ Enums: Can define but cannot construct variants with `::`
- ⚠️ Ranges: Work in for loops but not as standalone expressions
- ⚠️ Generics: Partially working, has runtime issues

### Not Implemented (Phase 1 Limitations)
- ❌ Method call syntax (`obj.method()`, `Type.method()`)
- ❌ Enum variant construction syntax (`Enum::Variant()`)
- ❌ Try operator (`?`)
- ❌ String interpolation (`"${var}"`)
- ❌ Generic type instantiation (partially works)
- ❌ Standard library (requires method calls)

## Recommendations

### Quick Wins (Can Fix Easily)
1. **Underscore in for loops**: Already works but shows parser warning
   - Fix: Update grammar to accept `_` as valid identifier in for loop pattern
   - Impact: Fixes parsing warnings in 2+ tests

2. **Range test**: Test expects feature not in Phase 1
   - Fix: Update test to use range only in for loop, or mark as Phase 2
   - Impact: 1 test fixed

3. **Enum tests**: Tests claim to pass but don't test actual enum usage
   - Fix: Either mark as placeholder tests or add TODO comments
   - Impact: Better test accuracy

### Medium Effort
4. **Generic function runtime**: "no call frame" error
   - Investigation needed in VM call frame handling for generic functions
   - Impact: Could fix ~18 generic tests if root cause found

### High Effort (Phase 2+ Features)
5. **Method call syntax**: Fundamental Phase 1 limitation
   - Impact: Would fix 24 tests but requires significant implementation

6. **Enum variant syntax (::)**: Requires parser and bytecode changes
   - Impact: Would fix 10 try_operator tests

## Next Steps
1. Fix underscore in for loops (grammar update)
2. Update range_test.uni to align with Phase 1
3. Document enum tests as placeholders
4. Investigate generic function call frame issue
5. Create documentation for Phase 2 features needed
