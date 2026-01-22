# Phase 4: Functions and Advanced Expressions

**Status**: Not Started  
**Duration**: 2-3 weeks  
**Priority**: MEDIUM  
**Dependencies**: Phase 3 Complete (Variables and Mutability)

## 🎯 Goals

Enhance function support with more features and implement advanced expressions including blocks as expressions, function expressions (lambdas), and advanced operators. This phase brings functional programming capabilities to Unified.

## 📋 Prerequisites

- [x] Phase 3 complete (variables and mutability working)
- [ ] Understanding of closure concepts
- [ ] Familiarity with functional programming patterns

## ✨ Language Features to Add

### 1. Function Features
- Multiple return values (via tuples)
- Default parameter values
- Named parameters
- Variadic parameters (basic implementation)

### 2. Block Expressions
- Blocks that return values
- Last expression as return value
- Scoped blocks with return values
- Void blocks vs value-returning blocks

### 3. Lambda Expressions
- Anonymous functions: `|x, y| x + y`
- Closures (basic, no capture initially)
- First-class functions
- Higher-order functions (functions as parameters)

### 4. Advanced Operators
- Bitwise operators (&, |, ^, ~, <<, >>)
- Operator precedence handling
- Unary operators (!, -, +)
- Logical operators (&&, ||)

## 📝 Implementation Tasks

### Task 4.1: Multiple Return Values (4-6 hours)
- [ ] Add tuple type to type system
- [ ] Define TupleType struct in AST
- [ ] Add tuple literal syntax to grammar
- [ ] Implement tuple literals in parser
- [ ] Support tuple destructuring in let statements
- [ ] Generate bytecode for tuple creation
- [ ] Implement in VM (return multiple values on stack)
- [ ] Add tests for tuple operations

**Grammar to add:**
```antlr
tupleType
    : LPAREN type (COMMA type)+ RPAREN
    ;

tupleLiteral
    : LPAREN expression (COMMA expression)+ RPAREN
    ;

letStatement
    : LET MUT? (IDENTIFIER | tuplePattern) (COLON type)? ASSIGN expression SEMICOLON?
    ;

tuplePattern
    : LPAREN IDENTIFIER (COMMA IDENTIFIER)+ RPAREN
    ;
```

**Example:**
```unified
let (quotient, remainder) = divide(10, 3)
```

### Task 4.2: Default Parameters (3-4 hours)
- [ ] Add default value to parameter syntax in grammar
- [ ] Update Parameter AST node with optional default value
- [ ] Validate default parameters at parse time
- [ ] Generate bytecode to handle missing arguments
- [ ] Validate call sites (arguments vs parameters)
- [ ] Add tests for default parameters

**Grammar to add:**
```antlr
parameter
    : IDENTIFIER COLON type (ASSIGN expression)?
    ;
```

**Example:**
```unified
fn greet(name: String, greeting: String = "Hello") -> String {
    return greeting + ", " + name
}
```

### Task 4.3: Block Expressions (3-4 hours)
- [ ] Update block grammar to support expression blocks
- [ ] Distinguish statement blocks from expression blocks
- [ ] Return last expression value from block
- [ ] Handle void blocks (no return value)
- [ ] Update type checker to infer block types
- [ ] Generate appropriate bytecode
- [ ] Test blocks in various contexts (if, while, function bodies)

**Example:**
```unified
let x = {
    let a = 5
    let b = 10
    a + b  // Last expression returned
}  // x = 15
```

### Task 4.4: Lambda Expressions (6-8 hours)
- [ ] Add lambda syntax to grammar
- [ ] Create LambdaExpression AST node
- [ ] Support both expression and block bodies
- [ ] Create function objects in VM
- [ ] Implement function pointers/references
- [ ] Support passing functions as parameters
- [ ] Add function type syntax: `fn(Int, Int) -> Int`
- [ ] Implement closure environment (basic, no capture yet)
- [ ] Test higher-order functions
- [ ] Add comprehensive lambda tests

**Grammar to add:**
```antlr
lambdaExpr
    : PIPE paramList? PIPE (ARROW expression | block)
    ;

paramList
    : IDENTIFIER (COMMA IDENTIFIER)*
    ;

functionType
    : FN LPAREN typeList? RPAREN (ARROW type)?
    ;
```

**Examples:**
```unified
let add = |a, b| a + b
let double = |x| x * 2
let complex = |x| {
    let y = x * 2
    y + 10
}
```

### Task 4.5: Bitwise Operators (2-3 hours)
- [ ] Add bitwise operators to lexer (if not present)
- [ ] Add bitwise operators to grammar
- [ ] Add bytecode instructions for bitwise ops
- [ ] Implement in VM
- [ ] Add operator precedence rules
- [ ] Test all bitwise operators
- [ ] Test operator precedence

**Operators to add:**
- `&` - Bitwise AND
- `|` - Bitwise OR
- `^` - Bitwise XOR
- `~` - Bitwise NOT
- `<<` - Left shift
- `>>` - Right shift

### Task 4.6: Operator Precedence (2-3 hours)
- [ ] Define precedence levels in grammar
- [ ] Update expression parsing
- [ ] Test precedence with multiple operators
- [ ] Add comprehensive precedence tests

### Task 4.7: Write Tests (4-5 hours)
- [ ] Tuple tests (creation, destructuring, type checking)
- [ ] Default parameter tests (all combinations)
- [ ] Block expression tests (various contexts)
- [ ] Lambda expression tests (all forms)
- [ ] Higher-order function tests
- [ ] Bitwise operator tests
- [ ] Operator precedence tests
- [ ] Integration tests combining features
- [ ] Ensure code coverage ≥ 85%

### Task 4.8: Update Documentation (2-3 hours)
- [ ] Create PHASE4_SUMMARY.md
- [ ] Document tuples in language guide
- [ ] Document lambdas in language guide
- [ ] Add functional programming examples
- [ ] Update README.md with new features
- [ ] Create advanced examples directory

## ✅ Test Requirements

**Minimum Test Coverage**: 85% for new code

### Tuple Tests
- [ ] Create tuple literals
- [ ] Destructure tuples in let statements
- [ ] Return tuples from functions
- [ ] Type check tuples correctly
- [ ] Handle nested tuples
- [ ] Error on tuple size mismatch

### Default Parameter Tests
- [ ] Call function with all arguments
- [ ] Call function omitting optional arguments
- [ ] Multiple default parameters
- [ ] Default parameters with expressions
- [ ] Type checking with defaults
- [ ] Error on invalid default values

### Block Expression Tests
- [ ] Block as expression returns last value
- [ ] Block in if condition
- [ ] Block in let statement
- [ ] Nested blocks
- [ ] Void blocks
- [ ] Type inference for blocks

### Lambda Tests
- [ ] Create lambda expressions
- [ ] Call lambda expressions
- [ ] Pass lambdas to functions
- [ ] Return lambdas from functions
- [ ] Lambda with multiple parameters
- [ ] Lambda with no parameters
- [ ] Lambda with block body
- [ ] Type checking for lambdas

### Higher-Order Function Tests
- [ ] Functions accepting function parameters
- [ ] Functions returning functions
- [ ] Chaining higher-order functions
- [ ] Type checking function parameters

### Bitwise Operator Tests
- [ ] AND operation works correctly
- [ ] OR operation works correctly
- [ ] XOR operation works correctly
- [ ] NOT operation works correctly
- [ ] Left shift works correctly
- [ ] Right shift works correctly
- [ ] Operator precedence is correct

### Integration Tests
- [ ] Map function with lambda
- [ ] Filter function with lambda
- [ ] Reduce/fold function with lambda
- [ ] Complex expression with all operators
- [ ] Function returning tuple with lambda

## 📚 Documentation Updates Required

### Update Existing Docs
- [ ] README.md: Add advanced features to feature list
- [ ] TESTING.md: Add new test categories
- [ ] VM_README.md: Document new bytecode instructions

### Create New Docs
- [ ] PHASE4_SUMMARY.md: Complete phase summary
- [ ] examples/FUNCTIONAL.md: Guide to functional programming features
- [ ] examples/LAMBDAS.md: Lambda expression examples
- [ ] docs/TUPLES.md: Tuple documentation

### Add Example Programs
- [ ] `test/higher_order.uni`: Higher-order function examples
- [ ] `test/map_filter.uni`: Map/filter implementation
- [ ] `test/tuples.uni`: Tuple usage examples
- [ ] `test/bitwise.uni`: Bitwise operation examples

## 🎓 Example Programs

### Higher-Order Functions
```unified
fn apply_twice(x: Int, f: fn(Int) -> Int) -> Int {
    return f(f(x))
}

fn double(x: Int) -> Int {
    return x * 2
}

fn main() -> Int {
    let result = apply_twice(5, double)  // 5 * 2 * 2 = 20
    return result
}
```

### Lambda Expressions
```unified
fn main() -> Int {
    let add = |a, b| a + b
    let multiply = |a, b| a * b
    
    let sum = add(10, 20)       // 30
    let product = multiply(5, 6) // 30
    
    return sum + product        // 60
}
```

### Map Function
```unified
fn map(arr: [Int; 5], f: fn(Int) -> Int) -> [Int; 5] {
    let mut result = [0, 0, 0, 0, 0]
    for i in 0..5 {
        result[i] = f(arr[i])
    }
    return result
}

fn main() -> Int {
    let numbers = [1, 2, 3, 4, 5]
    let doubled = map(numbers, |x| x * 2)
    // doubled = [2, 4, 6, 8, 10]
    return doubled[0]  // Returns 2
}
```

### Tuple Return Values
```unified
fn divide_with_remainder(a: Int, b: Int) -> (Int, Int) {
    let quotient = a / b
    let remainder = a % b
    return (quotient, remainder)
}

fn main() -> Int {
    let (q, r) = divide_with_remainder(17, 5)
    // q = 3, r = 2
    return q + r  // Returns 5
}
```

### Default Parameters
```unified
fn power(base: Int, exponent: Int = 2) -> Int {
    let mut result = 1
    for i in 0..exponent {
        result *= base
    }
    return result
}

fn main() -> Int {
    let squared = power(5)      // 25 (uses default exponent=2)
    let cubed = power(5, 3)     // 125
    return squared + cubed      // 150
}
```

### Block Expressions
```unified
fn compute() -> Int {
    let x = {
        let temp = 10
        let multiplier = 2
        temp * multiplier
    }  // x = 20
    
    let y = {
        let a = 5
        let b = 3
        a + b
    }  // y = 8
    
    return x + y  // 28
}
```

### Bitwise Operations
```unified
fn bitwise_example() -> Int {
    let a = 12      // 1100 in binary
    let b = 10      // 1010 in binary
    
    let and_result = a & b   // 1000 = 8
    let or_result = a | b    // 1110 = 14
    let xor_result = a ^ b   // 0110 = 6
    let shifted = a << 2     // 110000 = 48
    
    return and_result + or_result  // 22
}
```

## 🏆 Success Criteria

- [ ] All tuple tests pass (minimum 10 tests)
- [ ] All default parameter tests pass (minimum 8 tests)
- [ ] All block expression tests pass (minimum 10 tests)
- [ ] All lambda tests pass (minimum 15 tests)
- [ ] All higher-order function tests pass (minimum 8 tests)
- [ ] All bitwise operator tests pass (minimum 10 tests)
- [ ] All integration tests pass (minimum 6 tests)
- [ ] Can implement map/filter using lambdas
- [ ] Can create and use tuples
- [ ] Can use default parameters
- [ ] Operator precedence works correctly
- [ ] No regressions in previous phase tests
- [ ] Code coverage ≥ 85%
- [ ] Documentation complete
- [ ] All example programs run correctly

## 💡 Implementation Notes

### Implementation Order
1. Tuples (foundation for multiple returns)
2. Block expressions (simpler than lambdas)
3. Bitwise operators (independent feature)
4. Default parameters (enhances existing functions)
5. Lambda expressions (most complex)
6. Integration and testing

### Testing Strategy
- Test each feature independently first
- Then test combinations of features
- Focus on type checking for lambdas and tuples
- Verify higher-order functions work correctly
- Test edge cases (empty tuples, single-element tuples, etc.)

### Common Pitfalls
1. Type inference for lambdas can be tricky
2. Tuple destructuring pattern matching complexity
3. Block expression vs statement block confusion
4. Default parameter evaluation timing
5. Function type compatibility checking
6. Closure variable capture (not implemented in this phase)
7. Operator precedence conflicts

### Debugging Tips
- Use `make test` frequently to catch regressions
- Print AST for complex expressions
- Add debug output to type inference
- Test lambda type checking carefully
- Verify bytecode generation for new constructs
- Use integration tests to verify end-to-end behavior

### Future Considerations
This phase prepares for:
- Full closure support with capture (Phase 11)
- Advanced pattern matching on tuples
- Trait/interface system for function types
- Generic functions (Phase 10)

### Performance Considerations
- Lambda creation should be lightweight
- Tuple operations should be efficient
- Consider inline optimization for simple lambdas

---

**Labels**: `phase-4`, `enhancement`, `functional`, `expressions`  
**Milestone**: Phase 4 - Functions and Advanced Expressions  
**Assignee**: TBD
