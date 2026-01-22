# Phase 10: Generics Basics

**Status**: Not Started  
**Duration**: 3-4 weeks  
**Priority**: MEDIUM  
**Dependencies**: Phase 9 Complete (String Operations)

## 🎯 Goals

Implement basic generic types and functions with type parameters. This enables writing reusable code that works with multiple types while maintaining type safety through monomorphization.

## 📋 Prerequisites

- [x] Phase 9 complete (string operations working)
- [ ] Understanding of generic programming concepts
- [ ] Familiarity with type parameter systems
- [ ] Knowledge of monomorphization vs type erasure

## ✨ Language Features to Add

### 1. Generic Functions
- Type parameters: `fn identity<T>(x: T) -> T`
- Multiple type parameters: `fn pair<A, B>(a: A, b: B)`
- Type inference for generics
- Monomorphization (generate specialized versions)
- Generic function constraints (basic)

### 2. Generic Structs
- `struct Box<T> { value: T }`
- Generic struct instantiation: `Box<Int> { value: 42 }`
- Methods on generic structs
- Multiple type parameters

### 3. Generic Enums
- `enum Option<T> { Some(T), None }` (already defined, now generic)
- `enum Result<T, E> { Success(T), Error(E) }` (already defined, now generic)
- Pattern matching with generic enums
- Generic enum methods

### 4. Type Inference
- Infer type parameters from arguments
- Minimize explicit type annotations
- Type parameter constraints
- Error messages for inference failures

## 📝 Implementation Tasks

### Task 10.1: Generic Syntax (3-4 hours)
- [ ] Add type parameter syntax to grammar
- [ ] Add angle brackets for type parameters
- [ ] Support type parameter lists
- [ ] Support type parameter bounds (basic)
- [ ] Regenerate parser
- [ ] Test parsing generic declarations

**Grammar to add:**
```antlr
typeParams
    : LT typeParam (COMMA typeParam)* GT
    ;

typeParam
    : IDENTIFIER (COLON typeBound)?
    ;

typeBound
    : type (PLUS type)*
    ;

genericFunctionDecl
    : FN IDENTIFIER typeParams? LPAREN paramList? RPAREN (ARROW type)? block
    ;

genericStructDecl
    : STRUCT IDENTIFIER typeParams? LBRACE structBody RBRACE
    ;

genericEnumDecl
    : ENUM IDENTIFIER typeParams? LBRACE enumVariants RBRACE
    ;

typeApplication
    : IDENTIFIER (LT typeList GT)?
    ;
```

### Task 10.2: Type Parameter System (6-8 hours)
- [ ] Create TypeParameter in type system
- [ ] Create TypeVariable for inference
- [ ] Implement type substitution
- [ ] Track type parameter bindings
- [ ] Implement type parameter scope
- [ ] Add type parameter to symbol table
- [ ] Add unit tests for type parameters

**Type system structure:**
```go
type TypeParameter struct {
    Name       string
    Bounds     []Type
    Variance   Variance
}

type TypeVariable struct {
    ID         int
    Bounds     []Type
    Resolved   Type // nil if not yet resolved
}

type GenericType struct {
    Base       Type
    TypeParams []TypeParameter
}
```

### Task 10.3: Generic Function Implementation (8-10 hours)
- [ ] Parse generic function declarations
- [ ] Store type parameters in function signature
- [ ] Implement type parameter binding at call sites
- [ ] Type check generic function bodies
- [ ] Generate specialized versions (monomorphization)
- [ ] Mangle names for specialized versions
- [ ] Add tests for generic functions

### Task 10.4: Type Inference for Generics (6-8 hours)
- [ ] Implement type inference algorithm
- [ ] Infer type parameters from arguments
- [ ] Unification algorithm for type variables
- [ ] Constraint solving for type parameters
- [ ] Handle inference failures gracefully
- [ ] Provide helpful error messages
- [ ] Add comprehensive inference tests

**Example inference:**
```unified
fn identity<T>(x: T) -> T { return x }

let a = identity(42)         // Infers T = Int
let b = identity("hello")    // Infers T = String
```

### Task 10.5: Generic Structs (6-8 hours)
- [ ] Parse generic struct declarations
- [ ] Allow type parameters in field types
- [ ] Implement generic struct instantiation
- [ ] Type check generic struct usage
- [ ] Generate specialized struct types
- [ ] Methods on generic structs
- [ ] Add tests for generic structs

### Task 10.6: Generic Enums (4-6 hours)
- [ ] Update Option<T> to use type parameters
- [ ] Update Result<T, E> to use type parameters
- [ ] Pattern matching with generic enums
- [ ] Type inference for enum variants
- [ ] Add tests for generic enums

### Task 10.7: Monomorphization (8-10 hours)
- [ ] Implement monomorphization pass
- [ ] Generate specialized versions for each type combination
- [ ] Track which specializations are needed
- [ ] Avoid duplicate specializations
- [ ] Name mangling for specialized functions
- [ ] Code generation for specialized versions
- [ ] Optimize monomorphization process
- [ ] Add tests for monomorphization

**Monomorphization example:**
```unified
// Source:
fn max<T>(a: T, b: T) -> T { ... }
let x = max(5, 10)
let y = max(3.14, 2.71)

// Generated:
fn max_Int(a: Int, b: Int) -> Int { ... }
fn max_Float(a: Float, b: Float) -> Float { ... }
```

### Task 10.8: Type Parameter Constraints (4-6 hours)
- [ ] Implement basic trait bounds (foundation)
- [ ] Check constraints are satisfied
- [ ] Error on constraint violations
- [ ] Add tests for constraints

**Example constraints:**
```unified
fn add<T: Numeric>(a: T, b: T) -> T {
    return a + b
}
```

### Task 10.9: Write Tests (5-6 hours)
- [ ] Parser tests for generic syntax
- [ ] Type parameter tests
- [ ] Generic function tests
- [ ] Type inference tests
- [ ] Generic struct tests
- [ ] Generic enum tests
- [ ] Monomorphization tests
- [ ] Constraint tests
- [ ] Integration tests
- [ ] Ensure code coverage ≥ 85%

### Task 10.10: Update Documentation (3-4 hours)
- [ ] Create PHASE10_SUMMARY.md
- [ ] Document generic functions
- [ ] Document generic structs and enums
- [ ] Document type inference
- [ ] Add generic examples
- [ ] Update README.md
- [ ] Create generics guide

## ✅ Test Requirements

**Minimum Test Coverage**: 85% for new code

### Parser Tests
- [ ] Parse generic function declaration
- [ ] Parse generic struct declaration
- [ ] Parse generic enum declaration
- [ ] Parse type parameters
- [ ] Parse type bounds
- [ ] Parse type application
- [ ] Reject invalid syntax

### Type Parameter Tests
- [ ] Create type parameters
- [ ] Bind type parameters
- [ ] Substitute type parameters
- [ ] Scope type parameters correctly
- [ ] Multiple type parameters
- [ ] Type parameter bounds

### Generic Function Tests
- [ ] Declare generic functions
- [ ] Call generic functions with explicit types
- [ ] Call generic functions with inferred types
- [ ] Multiple type parameters
- [ ] Generic functions calling generic functions
- [ ] Return generic types

### Type Inference Tests
- [ ] Infer single type parameter
- [ ] Infer multiple type parameters
- [ ] Infer from complex expressions
- [ ] Error on ambiguous inference
- [ ] Error on conflicting inferences
- [ ] Helpful error messages

### Generic Struct Tests
- [ ] Declare generic structs
- [ ] Instantiate generic structs
- [ ] Access fields of generic structs
- [ ] Methods on generic structs
- [ ] Nested generic types
- [ ] Multiple type parameters

### Generic Enum Tests
- [ ] Option<T> works with different types
- [ ] Result<T, E> works with different types
- [ ] Pattern matching on generic enums
- [ ] Generic enum methods
- [ ] Type inference with enums

### Monomorphization Tests
- [ ] Generate specialized functions
- [ ] No duplicate specializations
- [ ] Correct name mangling
- [ ] Specialized code works correctly
- [ ] Complex generic combinations
- [ ] Recursive generics

### Integration Tests
- [ ] Generic stack implementation
- [ ] Generic linked list (if time)
- [ ] Map/filter with generics
- [ ] Option and Result with various types
- [ ] Complex generic scenarios

## 📚 Documentation Updates Required

### Update Existing Docs
- [ ] README.md: Add generics to feature list
- [ ] TESTING.md: Add generic tests
- [ ] VM_README.md: Document monomorphization

### Create New Docs
- [ ] PHASE10_SUMMARY.md: Complete phase summary
- [ ] examples/GENERICS.md: Guide to generics
- [ ] docs/TYPE_INFERENCE.md: Type inference documentation
- [ ] docs/MONOMORPHIZATION.md: Monomorphization details

### Add Example Programs
- [ ] `test/generic_functions.uni`: Generic function examples
- [ ] `test/generic_structs.uni`: Generic struct examples
- [ ] `test/generic_stack.uni`: Generic stack implementation
- [ ] `test/type_inference.uni`: Type inference examples

## 🎓 Example Programs

### Generic Identity Function
```unified
fn identity<T>(x: T) -> T {
    return x
}

fn main() -> Int {
    let num = identity(42)          // T = Int
    let text = identity("hello")    // T = String
    let flag = identity(true)       // T = Bool
    
    return num  // Returns 42
}
```

### Generic Pair Type
```unified
struct Pair<A, B> {
    first: A
    second: B
    
    fn new(first: A, second: B) -> Pair<A, B> {
        return Pair { first: first, second: second }
    }
    
    fn get_first(self) -> A {
        return self.first
    }
    
    fn get_second(self) -> B {
        return self.second
    }
}

fn main() -> Int {
    let p1 = Pair::new(42, "answer")     // Pair<Int, String>
    let p2 = Pair::new(3.14, true)       // Pair<Float, Bool>
    
    let num = p1.get_first()             // 42
    let pi = p2.get_first()              // 3.14
    
    return num  // Returns 42
}
```

### Generic Box Type
```unified
struct Box<T> {
    value: T
    
    fn new(value: T) -> Box<T> {
        return Box { value: value }
    }
    
    fn get(self) -> T {
        return self.value
    }
    
    fn set(&mut self, value: T) {
        self.value = value
    }
}

fn main() -> Int {
    let mut int_box = Box::new(42)
    let val = int_box.get()          // 42
    int_box.set(100)
    
    let string_box = Box::new("hello")
    let text = string_box.get()      // "hello"
    
    return int_box.get()  // Returns 100
}
```

### Generic Option (Already Defined)
```unified
enum Option<T> {
    Some(T),
    None
}

fn find<T>(arr: [T; 5], predicate: fn(T) -> Bool) -> Option<T> {
    for i in 0..5 {
        if predicate(arr[i]) {
            return Option::Some(arr[i])
        }
    }
    return Option::None
}

fn main() -> Int {
    let numbers = [1, 2, 3, 4, 5]
    let result = find(numbers, |x| x > 3)
    
    match result {
        Option::Some(value) -> return value  // Returns 4
        Option::None -> return 0
    }
}
```

### Generic Result (Already Defined)
```unified
enum Result<T, E> {
    Success(T),
    Error(E)
}

fn parse_positive<T>(value: Int) -> Result<Int, String> {
    if value < 0 {
        return Result::Error("Value must be positive")
    }
    return Result::Success(value)
}

fn main() -> Int {
    let r1 = parse_positive(42)
    let r2 = parse_positive(-5)
    
    match r1 {
        Result::Success(v) -> return v  // Returns 42
        Result::Error(e) -> return 0
    }
}
```

### Generic Stack
```unified
struct Stack<T> {
    items: [T; 100]
    size: Int
    
    fn new() -> Stack<T> {
        return Stack {
            items: [/* zero initialized */],
            size: 0
        }
    }
    
    fn push(&mut self, item: T) {
        if self.size < 100 {
            self.items[self.size] = item
            self.size += 1
        }
    }
    
    fn pop(&mut self) -> Option<T> {
        if self.size > 0 {
            self.size -= 1
            return Option::Some(self.items[self.size])
        }
        return Option::None
    }
    
    fn is_empty(self) -> Bool {
        return self.size == 0
    }
}

fn main() -> Int {
    let mut int_stack = Stack::new()
    int_stack.push(10)
    int_stack.push(20)
    int_stack.push(30)
    
    let top = int_stack.pop()
    match top {
        Option::Some(value) -> return value  // Returns 30
        Option::None -> return 0
    }
}
```

### Generic Swap Function
```unified
fn swap<T>(a: &mut T, b: &mut T) {
    let temp = *a
    *a = *b
    *b = temp
}

fn main() -> Int {
    let mut x = 5
    let mut y = 10
    swap(&mut x, &mut y)
    // x is now 10, y is now 5
    
    let mut s1 = "hello"
    let mut s2 = "world"
    swap(&mut s1, &mut s2)
    // s1 is now "world", s2 is now "hello"
    
    return x  // Returns 10
}
```

### Generic Maximum Function
```unified
fn max<T>(a: T, b: T) -> T {
    // Assuming > operator works (requires constraint)
    if a > b {
        return a
    }
    return b
}

fn main() -> Int {
    let max_int = max(5, 10)         // 10
    let max_float = max(3.14, 2.71)  // 3.14
    
    return max_int  // Returns 10
}
```

## 🏆 Success Criteria

- [ ] All parser tests pass (minimum 15 tests)
- [ ] All type parameter tests pass (minimum 12 tests)
- [ ] All generic function tests pass (minimum 20 tests)
- [ ] All type inference tests pass (minimum 15 tests)
- [ ] All generic struct tests pass (minimum 15 tests)
- [ ] All generic enum tests pass (minimum 12 tests)
- [ ] All monomorphization tests pass (minimum 10 tests)
- [ ] All integration tests pass (minimum 8 tests)
- [ ] Generics work for common cases
- [ ] Type inference minimizes annotations
- [ ] Monomorphization generates correct code
- [ ] Option and Result work as generic types
- [ ] No regressions in previous phase tests
- [ ] Code coverage ≥ 85%
- [ ] Documentation complete
- [ ] All example programs run correctly

## 💡 Implementation Notes

### Implementation Order
1. Generic syntax and parsing
2. Type parameter system
3. Type variables and substitution
4. Generic functions (simplest case)
5. Type inference for functions
6. Generic structs
7. Generic enums (update existing Option/Result)
8. Monomorphization implementation
9. Type parameter constraints (basic)
10. Integration and testing

### Testing Strategy
- Start with simple generic functions
- Add complexity gradually
- Test type inference thoroughly
- Verify monomorphization produces correct code
- Test with existing generic types (Option, Result)
- Use integration tests for realistic scenarios

### Common Pitfalls
1. Type inference too complex or ambiguous
2. Monomorphization explosion (too many specializations)
3. Circular dependencies in generic types
4. Type parameter scope issues
5. Constraint checking errors
6. Name mangling conflicts
7. Infinite generic recursion

### Debugging Tips
- Print type parameter bindings during inference
- Verify monomorphized code is correct
- Test simple cases first
- Check type substitution manually
- Use clear error messages
- Visualize type inference steps

### Monomorphization Strategy
- Lazy monomorphization (only what's used)
- Track instantiations to avoid duplicates
- Name mangling: `function_Type1_Type2`
- Consider code size implications
- Optimize common cases

### Type Inference Algorithm
1. Collect type constraints from expressions
2. Unify type variables with constraints
3. Solve constraint system
4. Substitute resolved types
5. Report errors for unsolvable constraints

### Performance Considerations
- Monomorphization increases code size
- Better runtime performance than type erasure
- Compile time increases with generic usage
- Consider compilation caching

### Future Enhancements
This phase prepares for:
- Advanced trait constraints
- Associated types
- Higher-kinded types
- Generic type aliases
- Const generics
- Specialization

---

**Labels**: `phase-10`, `enhancement`, `generics`, `type-system`  
**Milestone**: Phase 10 - Generics Basics  
**Assignee**: TBD
