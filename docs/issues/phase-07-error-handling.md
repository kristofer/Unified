# Phase 7: Error Handling with ? Operator

**Status**: Not Started  
**Duration**: 1-2 weeks  
**Priority**: MEDIUM  
**Dependencies**: Phase 6 Complete (Enums and Pattern Matching)

## 🎯 Goals

Implement the `?` operator for ergonomic error propagation with Result types. This operator provides Rust-style error handling, making error propagation concise and readable while maintaining type safety.

## 📋 Prerequisites

- [x] Phase 6 complete (enums and pattern matching)
- [x] Result<T, E> type implemented
- [ ] Understanding of error propagation patterns
- [ ] Familiarity with desugaring techniques

## ✨ Language Features to Add

### 1. ? Operator
- Automatic error propagation from Result types
- Early return on Error variant
- Unwrap Success value for continued use
- Type checking for compatible Result types

### 2. Error Context
- Error type compatibility checking
- Error type conversion (if time permits)
- Clear error messages for type mismatches

### 3. Syntactic Sugar
- Desugar `?` to match expression
- Maintain clean error handling semantics
- Optimize common cases

## 📝 Implementation Tasks

### Task 7.1: Add ? to Grammar (1 hour)
- [ ] Open `grammar/UnifiedParser.g4`
- [ ] Add QUESTION token to lexer if not present
- [ ] Add postfix ? operator to expression grammar
- [ ] Update operator precedence
- [ ] Regenerate parser: `make parser`
- [ ] Verify parser generation succeeds
- [ ] Test parsing expressions with ?

**Grammar to add:**
```antlr
postfixExpr
    : primaryExpr postfixOp*
    ;

postfixOp
    : QUESTION                              // ? operator
    | DOT IDENTIFIER                        // field access
    | LPAREN argumentList? RPAREN           // function call
    | LBRACKET expression RBRACKET          // array index
    ;
```

### Task 7.2: AST for ? Operator (1-2 hours)
- [ ] Create TryExpression AST node
- [ ] Add to expression visitor
- [ ] Store inner expression
- [ ] Add unit tests for AST construction

**AST structure:**
```go
type TryExpression struct {
    Expression Expression
    Position   Position
}
```

### Task 7.3: Type Checking for ? (3-4 hours)
- [ ] Verify expression is Result<T, E> type
- [ ] Extract inner success type T
- [ ] Extract error type E
- [ ] Verify function returns Result<U, E> (same error type)
- [ ] Report clear errors for type mismatches
- [ ] Add comprehensive type checking tests

**Type checking rules:**
```
expr: Result<T, E>
expr?: T (within function returning Result<U, E>)
```

### Task 7.4: Implement Desugaring (4-6 hours)
- [ ] Transform `expr?` to match expression
- [ ] Generate early return on Error
- [ ] Unwrap Success value
- [ ] Preserve error value and type
- [ ] Optimize bytecode generation
- [ ] Add tests for desugaring

**Desugaring transformation:**
```unified
// Source code:
let value = risky_operation()?

// Desugars to:
let value = match risky_operation() {
    Result::Success(v) -> v
    Result::Error(e) -> return Result::Error(e)
}
```

### Task 7.5: Bytecode Generation (2-3 hours)
- [ ] Generate bytecode for try expression
- [ ] Efficient pattern matching on Result
- [ ] Early return implementation
- [ ] Value unwrapping
- [ ] Test bytecode generation

### Task 7.6: Error Type Compatibility (2-3 hours)
- [ ] Check error types match in function
- [ ] Allow same error types only (for now)
- [ ] Plan for error type conversion (future)
- [ ] Add helpful error messages
- [ ] Test error type checking

### Task 7.7: Write Tests (3-4 hours)
- [ ] Parser tests for ? operator
- [ ] Type checking tests (success and failure)
- [ ] Desugaring tests
- [ ] Bytecode generation tests
- [ ] Error propagation tests (integration)
- [ ] Chained ? operator tests
- [ ] Edge case tests
- [ ] Ensure code coverage ≥ 85%

### Task 7.8: Update Documentation (2-3 hours)
- [ ] Create PHASE7_SUMMARY.md
- [ ] Document ? operator in language guide
- [ ] Add error handling best practices
- [ ] Create examples showing ? operator
- [ ] Update README.md
- [ ] Add migration guide from explicit matching

## ✅ Test Requirements

**Minimum Test Coverage**: 85% for new code

### Parser Tests
- [ ] Parse simple ? expression
- [ ] Parse chained ? expressions
- [ ] Parse ? in various contexts (let, return, function arg)
- [ ] Correct operator precedence
- [ ] Reject invalid ? usage

### Type Checking Tests
- [ ] ? on Result type succeeds
- [ ] ? on non-Result type fails
- [ ] Error type compatibility checked
- [ ] Success type extracted correctly
- [ ] Function return type validated
- [ ] Clear error messages for type mismatches

### Desugaring Tests
- [ ] Single ? desugars correctly
- [ ] Chained ? desugars correctly
- [ ] ? in complex expressions
- [ ] Preserves error value

### Execution Tests
- [ ] ? unwraps Success correctly
- [ ] ? returns Error early
- [ ] Chained ? stops at first Error
- [ ] Success values flow through
- [ ] Type safety maintained at runtime

### Integration Tests
- [ ] File operations with ?
- [ ] Parsing with error handling
- [ ] Chained operations with ?
- [ ] Real-world error handling scenarios

## 📚 Documentation Updates Required

### Update Existing Docs
- [ ] README.md: Add ? operator to feature list
- [ ] TESTING.md: Add error handling tests
- [ ] examples/OPTION_RESULT.md: Add ? operator examples

### Create New Docs
- [ ] PHASE7_SUMMARY.md: Complete phase summary
- [ ] examples/ERROR_HANDLING.md: Comprehensive error handling guide
- [ ] docs/TRY_OPERATOR.md: Detailed ? operator documentation

### Add Example Programs
- [ ] `test/try_operator.uni`: Basic ? operator usage
- [ ] `test/chained_errors.uni`: Chained error handling
- [ ] `test/file_operations.uni`: Realistic error handling
- [ ] `test/error_propagation.uni`: Multiple error sources

## 🎓 Example Programs

### Basic ? Operator Usage
```unified
enum Result<T, E> {
    Success(T),
    Error(E)
}

fn read_file(path: String) -> Result<String, String> {
    // Simulated file reading
    if path == "" {
        return Result::Error("Empty path")
    }
    return Result::Success("file contents")
}

fn parse_int(s: String) -> Result<Int, String> {
    // Simulated parsing
    if s == "file contents" {
        return Result::Error("Not a number")
    }
    return Result::Success(42)
}

fn read_and_parse(path: String) -> Result<Int, String> {
    let contents = read_file(path)?     // Propagates error if reading fails
    let number = parse_int(contents)?   // Propagates error if parsing fails
    return Result::Success(number)
}

fn main() -> Int {
    let result = read_and_parse("data.txt")
    match result {
        Result::Success(n) -> return n
        Result::Error(msg) -> return -1
    }
}
```

### Chained Error Handling
```unified
fn step1() -> Result<Int, String> {
    return Result::Success(10)
}

fn step2(x: Int) -> Result<Int, String> {
    if x < 0 {
        return Result::Error("Negative value")
    }
    return Result::Success(x * 2)
}

fn step3(x: Int) -> Result<Int, String> {
    if x > 100 {
        return Result::Error("Too large")
    }
    return Result::Success(x + 5)
}

fn process() -> Result<Int, String> {
    let a = step1()?
    let b = step2(a)?
    let c = step3(b)?
    return Result::Success(c)
}

fn main() -> Int {
    match process() {
        Result::Success(value) -> return value  // Returns 25
        Result::Error(msg) -> return -1
    }
}
```

### File Processing Pipeline
```unified
fn open_file(path: String) -> Result<File, String> {
    // Simulated
    return Result::Success(File { handle: 1 })
}

fn read_line(file: File) -> Result<String, String> {
    // Simulated
    return Result::Success("42")
}

fn parse_number(s: String) -> Result<Int, String> {
    // Simulated
    return Result::Success(42)
}

fn close_file(file: File) -> Result<Bool, String> {
    // Simulated
    return Result::Success(true)
}

fn process_file(path: String) -> Result<Int, String> {
    let file = open_file(path)?
    let line = read_line(file)?
    let number = parse_number(line)?
    close_file(file)?
    
    return Result::Success(number)
}
```

### Error Recovery
```unified
fn divide(a: Int, b: Int) -> Result<Int, String> {
    if b == 0 {
        return Result::Error("Division by zero")
    }
    return Result::Success(a / b)
}

fn safe_divide_chain(a: Int, b: Int, c: Int) -> Result<Int, String> {
    let step1 = divide(a, b)?
    let step2 = divide(step1, c)?
    return Result::Success(step2)
}

fn main() -> Int {
    // This will fail if b or c is zero
    let result = safe_divide_chain(100, 10, 2)
    
    match result {
        Result::Success(value) -> return value  // Returns 5
        Result::Error(msg) -> {
            // Error recovery logic
            return 0
        }
    }
}
```

### Multiple Error Checks
```unified
fn validate_age(age: Int) -> Result<Int, String> {
    if age < 0 {
        return Result::Error("Age cannot be negative")
    }
    if age > 150 {
        return Result::Error("Age is unrealistic")
    }
    return Result::Success(age)
}

fn validate_name(name: String) -> Result<String, String> {
    if name == "" {
        return Result::Error("Name cannot be empty")
    }
    return Result::Success(name)
}

fn create_user(name: String, age: Int) -> Result<User, String> {
    let valid_name = validate_name(name)?
    let valid_age = validate_age(age)?
    
    // Both validations passed
    return Result::Success(User {
        name: valid_name,
        age: valid_age
    })
}
```

## 🏆 Success Criteria

- [ ] All parser tests pass (minimum 8 tests)
- [ ] All type checking tests pass (minimum 10 tests)
- [ ] All desugaring tests pass (minimum 8 tests)
- [ ] All execution tests pass (minimum 10 tests)
- [ ] All integration tests pass (minimum 6 tests)
- [ ] ? operator unwraps Success correctly
- [ ] ? operator returns Error early
- [ ] Can chain multiple ? operators
- [ ] Type checking prevents invalid usage
- [ ] Error handling is ergonomic
- [ ] No regressions in previous phase tests
- [ ] Code coverage ≥ 85%
- [ ] Documentation complete
- [ ] All example programs run correctly

## 💡 Implementation Notes

### Implementation Order
1. Grammar update (? operator)
2. AST representation
3. Type checking (critical!)
4. Desugaring implementation
5. Bytecode generation
6. Testing and documentation

### Testing Strategy
- Test type checking thoroughly (most critical)
- Verify desugaring is correct
- Test both success and error paths
- Test chaining behavior
- Ensure error messages are helpful

### Common Pitfalls
1. Not checking error type compatibility
2. Incorrect desugaring for chained ? operators
3. Lost error context during propagation
4. Type inference issues with generic Result types
5. Operator precedence conflicts
6. Not handling all Result variants

### Debugging Tips
- Print desugared code during development
- Verify type checker logic with simple examples
- Test error paths explicitly
- Check bytecode for early return implementation
- Use integration tests to verify real behavior

### Desugaring Details
The ? operator is pure syntactic sugar:
- Does NOT add runtime overhead
- Compiles to same code as manual match
- Makes error handling more readable
- Encourages proper error propagation

### Type System Integration
- Must work with generic Result<T, E>
- Error type E must be compatible
- Success type T is extracted for use
- Function must return Result<U, E>

### Future Enhancements
This phase prepares for:
- Error type conversion/mapping
- Try blocks (try { ... })
- Custom error types with conversions
- Error trait/interface
- Multiple error types in one function

### Performance Considerations
- Desugaring should generate efficient code
- No runtime overhead vs manual matching
- Consider inlining for simple cases
- Jump optimization for error paths

---

**Labels**: `phase-7`, `enhancement`, `error-handling`, `syntax-sugar`  
**Milestone**: Phase 7 - Error Handling  
**Assignee**: TBD
