# Phase 8: Arrays and Slices

**Status**: Not Started  
**Duration**: 2-3 weeks  
**Priority**: HIGH  
**Dependencies**: Phase 7 Complete (Error Handling)

## 🎯 Goals

Implement fixed-size arrays and dynamic slices with bounds checking. This provides essential data structures for collections and enables safe memory access patterns.

## 📋 Prerequisites

- [x] Phase 7 complete (error handling with ?)
- [ ] Understanding of array memory layout
- [ ] Familiarity with bounds checking techniques
- [ ] Knowledge of slice implementations (pointer + length)

## ✨ Language Features to Add

### 1. Array Types
- Fixed-size arrays: `[Int; 10]`
- Array literals: `[1, 2, 3, 4, 5]`
- Array indexing with bounds checking
- Array type checking and inference
- Multidimensional arrays (if time permits)

### 2. Slice Types
- Dynamic slices: `[Int]`
- Slice operations: `arr[1..5]`, `arr[..3]`, `arr[2..]`
- Slice as (pointer, length) representation
- Slice methods (len, is_empty)
- Slice bounds checking

### 3. For-Each Loops
- Iterate over arrays and slices
- Iterator protocol foundation
- Index-based and value-based iteration
- Mutable iteration

## 📝 Implementation Tasks

### Task 8.1: Array Grammar and Types (4-6 hours)
- [ ] Add array type syntax to grammar
- [ ] Add array literal syntax
- [ ] Add array indexing syntax
- [ ] Add slice syntax (range indexing)
- [ ] Regenerate parser
- [ ] Test parsing arrays and slices

**Grammar to add:**
```antlr
arrayType
    : LBRACKET type SEMICOLON INTEGER RBRACKET   // Fixed size: [Int; 10]
    | LBRACKET type RBRACKET                      // Slice: [Int]
    ;

arrayLiteral
    : LBRACKET (expression (COMMA expression)*)? RBRACKET
    ;

indexExpr
    : postfixExpr LBRACKET expression RBRACKET
    ;

sliceExpr
    : postfixExpr LBRACKET sliceRange RBRACKET
    ;

sliceRange
    : expression? DOTDOT expression?        // start..end
    | expression? DOTDOT_EQ expression?     // start..=end
    ;
```

### Task 8.2: Array Type System (4-6 hours)
- [ ] Create ArrayType in type system
- [ ] Store element type and size
- [ ] Create SliceType in type system
- [ ] Type checking for array literals
- [ ] Type checking for array indexing
- [ ] Type checking for slicing operations
- [ ] Add unit tests for array types

**Type system structure:**
```go
type ArrayType struct {
    ElementType Type
    Size        int
}

type SliceType struct {
    ElementType Type
}
```

### Task 8.3: Array Memory Layout (4-5 hours)
- [ ] Define array memory layout (contiguous)
- [ ] Calculate array size at compile time
- [ ] Implement array allocation in VM
- [ ] Implement element offset calculation
- [ ] Handle array initialization
- [ ] Zero-initialize arrays
- [ ] Add tests for memory layout

### Task 8.4: Array Indexing (4-5 hours)
- [ ] Implement index expression in AST
- [ ] Generate bounds checking code
- [ ] Generate load bytecode for arr[i]
- [ ] Generate store bytecode for arr[i] = value
- [ ] Implement bounds checking in VM
- [ ] Runtime error on out-of-bounds access
- [ ] Add comprehensive indexing tests

**Bytecode:**
```
LOAD_ARRAY <base> <index>    // Checks bounds, loads element
STORE_ARRAY <base> <index> <value>  // Checks bounds, stores element
```

### Task 8.5: Slice Implementation (6-8 hours)
- [ ] Implement slice as (pointer, length) struct
- [ ] Generate slice creation bytecode
- [ ] Implement slicing operation (arr[start..end])
- [ ] Handle partial slices (arr[..end], arr[start..])
- [ ] Bounds checking for slice creation
- [ ] Implement slice indexing
- [ ] Add tests for slice operations

### Task 8.6: Array/Slice Iteration (4-6 hours)
- [ ] Update for loop to support array iteration
- [ ] Implement index-based iteration
- [ ] Implement value-based iteration (for item in array)
- [ ] Support mutable iteration
- [ ] Generate efficient iteration bytecode
- [ ] Add iterator protocol foundation
- [ ] Test all iteration patterns

**Example iteration patterns:**
```unified
for i in 0..arr.len() { }           // Index iteration
for item in arr { }                  // Value iteration
for (index, item) in arr.enumerate() { }  // Both (future)
```

### Task 8.7: Array Methods (3-4 hours)
- [ ] Implement len() method for arrays
- [ ] Implement is_empty() method
- [ ] Plan for future methods (push, pop, etc. for Vec)
- [ ] Add tests for methods

### Task 8.8: Write Tests (5-6 hours)
- [ ] Parser tests for array syntax
- [ ] Parser tests for slice syntax
- [ ] Type system tests
- [ ] Array creation tests
- [ ] Array indexing tests (success and bounds errors)
- [ ] Slice creation tests
- [ ] Slice indexing tests
- [ ] Iteration tests
- [ ] Integration tests
- [ ] Ensure code coverage ≥ 85%

### Task 8.9: Update Documentation (3-4 hours)
- [ ] Create PHASE8_SUMMARY.md
- [ ] Document arrays in language guide
- [ ] Document slices in language guide
- [ ] Add array/slice examples
- [ ] Update README.md
- [ ] Create collections guide

## ✅ Test Requirements

**Minimum Test Coverage**: 85% for new code

### Parser Tests
- [ ] Parse array type `[Int; 10]`
- [ ] Parse slice type `[Int]`
- [ ] Parse array literal `[1, 2, 3]`
- [ ] Parse array indexing `arr[5]`
- [ ] Parse slice range `arr[1..5]`
- [ ] Parse partial slices `arr[..5]`, `arr[2..]`
- [ ] Reject invalid syntax

### Type System Tests
- [ ] Array type with size
- [ ] Slice type without size
- [ ] Type check array literals
- [ ] Infer array size from literal
- [ ] Type check element access
- [ ] Type check slice operations
- [ ] Detect type mismatches

### Array Creation Tests
- [ ] Create fixed-size arrays
- [ ] Initialize from literal
- [ ] Zero-initialize arrays
- [ ] Nested arrays
- [ ] Error on size mismatch

### Indexing Tests
- [ ] Access array elements
- [ ] Assign to array elements
- [ ] Bounds checking catches out-of-bounds
- [ ] Error messages are helpful
- [ ] Negative index handling
- [ ] Index at boundaries (0, len-1)

### Slice Tests
- [ ] Create slices from arrays
- [ ] Slice with range `arr[1..5]`
- [ ] Partial slices work
- [ ] Slice of slice
- [ ] Bounds checking for slices
- [ ] Empty slices

### Iteration Tests
- [ ] Iterate over array with index
- [ ] Iterate over array values
- [ ] Iterate over slice
- [ ] Mutable iteration
- [ ] Break/continue in array loops
- [ ] Nested iteration

### Integration Tests
- [ ] Sum array elements
- [ ] Find element in array
- [ ] Reverse array
- [ ] Sort array (simple bubble sort)
- [ ] Matrix operations (2D arrays)
- [ ] Slice manipulation

## 📚 Documentation Updates Required

### Update Existing Docs
- [ ] README.md: Add arrays and slices to feature list
- [ ] TESTING.md: Add collection tests
- [ ] VM_README.md: Document array bytecode

### Create New Docs
- [ ] PHASE8_SUMMARY.md: Complete phase summary
- [ ] examples/ARRAYS.md: Guide to arrays
- [ ] examples/SLICES.md: Guide to slices
- [ ] docs/COLLECTIONS.md: Collections overview

### Add Example Programs
- [ ] `test/array_basics.uni`: Basic array operations
- [ ] `test/slicing.uni`: Slice operations
- [ ] `test/array_iteration.uni`: Iteration examples
- [ ] `test/matrix.uni`: 2D array example

## 🎓 Example Programs

### Array Basics
```unified
fn sum_array(arr: [Int; 5]) -> Int {
    let mut sum = 0
    for i in 0..5 {
        sum += arr[i]
    }
    return sum
}

fn main() -> Int {
    let numbers = [1, 2, 3, 4, 5]
    let total = sum_array(numbers)
    return total  // Returns 15
}
```

### Array Iteration
```unified
fn find_max(arr: [Int; 10]) -> Int {
    let mut max = arr[0]
    for item in arr {
        if item > max {
            max = item
        }
    }
    return max
}

fn main() -> Int {
    let numbers = [5, 2, 8, 1, 9, 3, 7, 4, 6, 10]
    return find_max(numbers)  // Returns 10
}
```

### Slice Operations
```unified
fn sum_slice(slice: [Int]) -> Int {
    let mut sum = 0
    for item in slice {
        sum += item
    }
    return sum
}

fn main() -> Int {
    let numbers = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10]
    
    let first_half = numbers[..5]      // [1, 2, 3, 4, 5]
    let second_half = numbers[5..]     // [6, 7, 8, 9, 10]
    let middle = numbers[3..7]         // [4, 5, 6, 7]
    
    let sum1 = sum_slice(first_half)   // 15
    let sum2 = sum_slice(second_half)  // 40
    
    return sum1 + sum2  // Returns 55
}
```

### Mutable Array Operations
```unified
fn double_elements(arr: &mut [Int; 5]) {
    for i in 0..5 {
        arr[i] = arr[i] * 2
    }
}

fn main() -> Int {
    let mut numbers = [1, 2, 3, 4, 5]
    double_elements(&mut numbers)
    // numbers is now [2, 4, 6, 8, 10]
    return numbers[0]  // Returns 2
}
```

### 2D Arrays (Matrix)
```unified
fn matrix_sum(matrix: [[Int; 3]; 3]) -> Int {
    let mut sum = 0
    for i in 0..3 {
        for j in 0..3 {
            sum += matrix[i][j]
        }
    }
    return sum
}

fn main() -> Int {
    let matrix = [
        [1, 2, 3],
        [4, 5, 6],
        [7, 8, 9]
    ]
    return matrix_sum(matrix)  // Returns 45
}
```

### Array Search
```unified
fn linear_search(arr: [Int; 10], target: Int) -> Option<Int> {
    for i in 0..10 {
        if arr[i] == target {
            return Option::Some(i)
        }
    }
    return Option::None
}

fn main() -> Int {
    let numbers = [10, 20, 30, 40, 50, 60, 70, 80, 90, 100]
    let result = linear_search(numbers, 50)
    
    match result {
        Option::Some(index) -> return index  // Returns 4
        Option::None -> return -1
    }
}
```

### Bubble Sort
```unified
fn bubble_sort(arr: &mut [Int; 8]) {
    for i in 0..8 {
        for j in 0..(8 - i - 1) {
            if arr[j] > arr[j + 1] {
                // Swap
                let temp = arr[j]
                arr[j] = arr[j + 1]
                arr[j + 1] = temp
            }
        }
    }
}

fn main() -> Int {
    let mut numbers = [64, 34, 25, 12, 22, 11, 90, 88]
    bubble_sort(&mut numbers)
    // numbers is now [11, 12, 22, 25, 34, 64, 88, 90]
    return numbers[0]  // Returns 11
}
```

### Array Reversal
```unified
fn reverse(arr: &mut [Int; 6]) {
    let len = 6
    for i in 0..(len / 2) {
        let temp = arr[i]
        arr[i] = arr[len - 1 - i]
        arr[len - 1 - i] = temp
    }
}

fn main() -> Int {
    let mut numbers = [1, 2, 3, 4, 5, 6]
    reverse(&mut numbers)
    // numbers is now [6, 5, 4, 3, 2, 1]
    return numbers[0]  // Returns 6
}
```

## 🏆 Success Criteria

- [ ] All parser tests pass (minimum 15 tests)
- [ ] All type system tests pass (minimum 12 tests)
- [ ] All array creation tests pass (minimum 10 tests)
- [ ] All indexing tests pass (minimum 15 tests)
- [ ] All slice tests pass (minimum 12 tests)
- [ ] All iteration tests pass (minimum 10 tests)
- [ ] All integration tests pass (minimum 8 tests)
- [ ] Arrays and slices work correctly
- [ ] Bounds checking prevents errors
- [ ] Can iterate over collections
- [ ] Array literals type check correctly
- [ ] Slicing operations work
- [ ] No regressions in previous phase tests
- [ ] Code coverage ≥ 85%
- [ ] Documentation complete
- [ ] All example programs run correctly

## 💡 Implementation Notes

### Implementation Order
1. Array types and parsing
2. Array memory layout
3. Array literal creation
4. Array indexing (read)
5. Array indexing (write)
6. Bounds checking
7. Slice types and representation
8. Slice creation from arrays
9. Slice indexing
10. Array/slice iteration
11. Integration and testing

### Testing Strategy
- Test bounds checking thoroughly (security critical!)
- Test all slice range combinations
- Verify memory layout is correct
- Test iteration patterns
- Use integration tests for realistic scenarios
- Verify error messages are helpful

### Common Pitfalls
1. Off-by-one errors in bounds checking
2. Incorrect slice length calculations
3. Memory layout misalignment
4. Forgetting to check bounds on slices
5. Iterator invalidation issues
6. Negative index handling
7. Empty array/slice edge cases

### Debugging Tips
- Print array memory layout during development
- Verify bounds checking logic carefully
- Test boundary conditions (0, len-1, len)
- Check slice pointer arithmetic
- Use simple examples before complex ones
- Visualize memory layout

### Memory Layout
**Array:**
```
[element_0][element_1][element_2]...[element_n-1]
```
- Contiguous memory
- Fixed size known at compile time
- Elements of same type

**Slice:**
```
Slice {
    ptr: *T,      // Pointer to first element
    len: usize    // Number of elements
}
```

### Bounds Checking
Always check:
- `0 <= index < length`
- Compile-time checking where possible
- Runtime checking otherwise
- Clear error messages with index and length

### Performance Considerations
- Bounds checking has small overhead
- Consider removing checks in unsafe blocks (future)
- Iteration can be optimized to avoid repeated checks
- Array access should be as fast as C for simple cases

### Future Enhancements
This phase prepares for:
- Dynamic arrays (Vec<T>)
- Iterator trait
- Array methods (map, filter, etc.)
- Unsafe unchecked indexing
- SIMD operations on arrays
- String as [u8] slice

---

**Labels**: `phase-8`, `enhancement`, `collections`, `arrays`, `safety`  
**Milestone**: Phase 8 - Arrays and Slices  
**Assignee**: TBD
