# Phase 6: Enums and Pattern Matching

**Status**: Not Started  
**Duration**: 3-4 weeks  
**Priority**: MEDIUM  
**Dependencies**: Phase 5 Complete (Structs and Basic Types)

## 🎯 Goals

Implement algebraic data types (enums) and pattern matching, enabling Result and Option types. This brings powerful functional programming capabilities and enables safe error handling without exceptions.

## 📋 Prerequisites

- [x] Phase 5 complete (structs working)
- [ ] Understanding of tagged unions
- [ ] Familiarity with pattern matching concepts
- [ ] Knowledge of exhaustiveness checking algorithms

## ✨ Language Features to Add

### 1. Enum Declarations
- Simple enums (unit variants): `enum Color { Red, Green, Blue }`
- Enums with data (tuple variants): `enum Result { Success(Int), Error(String) }`
- Enums with named fields (struct variants): `enum Shape { Circle { radius: Float } }`
- Enum variant constructors

### 2. Pattern Matching
- Match expressions with exhaustiveness checking
- Patterns for literals, variables, structs, enums
- Pattern guards (if conditions)
- Wildcard patterns (_)
- Variable binding in patterns
- Nested patterns

### 3. Standard Enums
- Option<T>: Some(T), None
- Result<T, E>: Success(T), Error(E)
- Prelude integration

## 📝 Implementation Tasks

### Task 6.1: Enum Grammar (3-4 hours)
- [ ] Add enum declaration to grammar
- [ ] Add unit variants (no data)
- [ ] Add tuple variants (with data)
- [ ] Add struct variants (with named fields)
- [ ] Add enum variant construction syntax
- [ ] Regenerate parser
- [ ] Test parsing enum declarations

**Grammar to add:**
```antlr
enumDecl
    : ENUM IDENTIFIER genericParams? LBRACE enumVariants RBRACE
    ;

enumVariants
    : enumVariant (COMMA enumVariant)* COMMA?
    ;

enumVariant
    : IDENTIFIER                                    // Unit variant
    | IDENTIFIER LPAREN typeList RPAREN            // Tuple variant
    | IDENTIFIER LBRACE structBody RBRACE          // Struct variant
    ;

enumConstructor
    : IDENTIFIER DOUBLE_COLON IDENTIFIER (LPAREN exprList RPAREN | LBRACE fieldInits RBRACE)?
    ;
```

### Task 6.2: Enum Type System (6-8 hours)
- [ ] Create EnumType in type system
- [ ] Implement tag-union representation
- [ ] Track variant names and types
- [ ] Calculate variant data sizes
- [ ] Implement variant constructors as functions
- [ ] Add enum to type registry
- [ ] Type checking for enum construction
- [ ] Add unit tests for enum types

**Type system structure:**
```go
type EnumType struct {
    Name     string
    Variants []EnumVariant
    TypeParams []TypeParam // For generic enums
}

type EnumVariant struct {
    Name   string
    Kind   VariantKind
    Fields []Field // For tuple or struct variants
}

type VariantKind int
const (
    UnitVariant VariantKind = iota
    TupleVariant
    StructVariant
)
```

### Task 6.3: Match Expression Grammar (3-4 hours)
- [ ] Add match expression to grammar
- [ ] Add pattern syntax
- [ ] Add match arms
- [ ] Add pattern guards
- [ ] Support multiple patterns per arm
- [ ] Regenerate parser
- [ ] Test parsing match expressions

**Grammar to add:**
```antlr
matchExpr
    : MATCH expression LBRACE matchArm+ RBRACE
    ;

matchArm
    : pattern (IF expression)? ARROW expression (COMMA | SEMICOLON)?
    ;

pattern
    : literalPattern
    | wildcardPattern
    | identifierPattern
    | tuplePattern
    | structPattern
    | enumPattern
    ;

literalPattern
    : INTEGER | FLOAT | STRING | TRUE | FALSE
    ;

wildcardPattern
    : UNDERSCORE
    ;

identifierPattern
    : IDENTIFIER
    ;

enumPattern
    : IDENTIFIER DOUBLE_COLON IDENTIFIER (LPAREN patternList RPAREN | LBRACE fieldPatterns RBRACE)?
    ;
```

### Task 6.4: Pattern Compilation (8-10 hours)
- [ ] Implement pattern matching algorithm
- [ ] Compile patterns to decision tree
- [ ] Generate efficient bytecode for matching
- [ ] Extract variables from patterns
- [ ] Handle nested patterns
- [ ] Optimize common cases
- [ ] Add tests for pattern compilation

### Task 6.5: Exhaustiveness Checking (6-8 hours)
- [ ] Implement exhaustiveness checker
- [ ] Detect missing patterns
- [ ] Detect unreachable patterns
- [ ] Report helpful error messages
- [ ] Handle complex enum combinations
- [ ] Add comprehensive tests

**Algorithm:**
- Build decision tree of all patterns
- Check if all enum variants are covered
- Check if wildcards make patterns unreachable
- Report specific missing cases

### Task 6.6: Match Bytecode Generation (4-6 hours)
- [ ] Generate VM instructions for matching
- [ ] Jump table for efficient dispatch
- [ ] Variable binding in patterns
- [ ] Guard evaluation
- [ ] Optimize simple cases
- [ ] Test bytecode generation

### Task 6.7: Option and Result Types (4-6 hours)
- [ ] Define Option<T> in prelude
- [ ] Define Result<T, E> in prelude
- [ ] Add constructors (Some, None, Success, Error)
- [ ] Implement methods (map, and_then, unwrap, etc.)
- [ ] Add tests for Option and Result
- [ ] Create examples using Option and Result

**Prelude definitions:**
```unified
enum Option<T> {
    Some(T),
    None
}

enum Result<T, E> {
    Success(T),
    Error(E)
}
```

### Task 6.8: Write Tests (5-6 hours)
- [ ] Parser tests for enums
- [ ] Parser tests for match expressions
- [ ] Type system tests for enums
- [ ] Pattern compilation tests
- [ ] Exhaustiveness checking tests
- [ ] Match execution tests
- [ ] Option tests
- [ ] Result tests
- [ ] Integration tests
- [ ] Ensure code coverage ≥ 85%

### Task 6.9: Update Documentation (3-4 hours)
- [ ] Create PHASE6_SUMMARY.md
- [ ] Document enums in language guide
- [ ] Document pattern matching
- [ ] Document Option and Result
- [ ] Add enum examples
- [ ] Update README.md

## ✅ Test Requirements

**Minimum Test Coverage**: 85% for new code

### Parser Tests
- [ ] Parse unit variant enums
- [ ] Parse tuple variant enums
- [ ] Parse struct variant enums
- [ ] Parse mixed variant enums
- [ ] Parse match expressions
- [ ] Parse all pattern types
- [ ] Parse pattern guards
- [ ] Reject invalid syntax

### Type System Tests
- [ ] Create enum types
- [ ] Register variants
- [ ] Type check enum construction
- [ ] Type check match expressions
- [ ] Validate variant data types
- [ ] Generic enum types (Option, Result)

### Pattern Matching Tests
- [ ] Match on unit variants
- [ ] Match on tuple variants
- [ ] Match on struct variants
- [ ] Extract data from patterns
- [ ] Bind variables in patterns
- [ ] Nested pattern matching
- [ ] Pattern guards work correctly
- [ ] Wildcard patterns

### Exhaustiveness Tests
- [ ] Detect missing enum variants
- [ ] Detect unreachable patterns
- [ ] Allow wildcard to catch all
- [ ] Error messages are helpful
- [ ] Complex enum combinations
- [ ] Nested enum matching

### Option Tests
- [ ] Create Some and None values
- [ ] Match on Option
- [ ] Use Option methods
- [ ] Type safety with Option

### Result Tests
- [ ] Create Success and Error values
- [ ] Match on Result
- [ ] Use Result methods
- [ ] Error propagation with Result

### Integration Tests
- [ ] Divide function returning Result
- [ ] Find function returning Option
- [ ] Complex pattern matching scenarios
- [ ] Nested enum matching
- [ ] Real-world use cases

## 📚 Documentation Updates Required

### Update Existing Docs
- [ ] README.md: Add enums and pattern matching
- [ ] TESTING.md: Add pattern matching tests
- [ ] VM_README.md: Document match bytecode

### Create New Docs
- [ ] PHASE6_SUMMARY.md: Complete phase summary
- [ ] examples/ENUMS.md: Guide to enums
- [ ] examples/PATTERN_MATCHING.md: Pattern matching guide
- [ ] docs/OPTION_RESULT.md: Option and Result documentation

### Add Example Programs
- [ ] `test/enums.uni`: Basic enum examples
- [ ] `test/pattern_matching.uni`: Pattern matching examples
- [ ] `test/option.uni`: Option type usage
- [ ] `test/result.uni`: Result type usage

## 🎓 Example Programs

### Basic Enum with Pattern Matching
```unified
enum Color {
    Red,
    Green,
    Blue,
    RGB(Int, Int, Int)
}

fn color_to_string(color: Color) -> String {
    match color {
        Color::Red -> "red"
        Color::Green -> "green"
        Color::Blue -> "blue"
        Color::RGB(r, g, b) -> "custom"
    }
}

fn main() -> Int {
    let c1 = Color::Red
    let c2 = Color::RGB(128, 0, 255)
    return 0
}
```

### Result Type Usage
```unified
enum Result<T, E> {
    Success(T),
    Error(E)
}

fn divide(a: Int, b: Int) -> Result<Int, String> {
    if b == 0 {
        return Result::Error("Division by zero")
    }
    return Result::Success(a / b)
}

fn main() -> Int {
    let result = divide(10, 2)
    match result {
        Result::Success(value) -> return value      // Returns 5
        Result::Error(msg) -> return -1
    }
}
```

### Option Type Usage
```unified
enum Option<T> {
    Some(T),
    None
}

fn find_first_positive(numbers: [Int; 5]) -> Option<Int> {
    for i in 0..5 {
        if numbers[i] > 0 {
            return Option::Some(numbers[i])
        }
    }
    return Option::None
}

fn main() -> Int {
    let nums = [-1, -5, 3, 7, -2]
    let result = find_first_positive(nums)
    
    match result {
        Option::Some(value) -> return value     // Returns 3
        Option::None -> return 0
    }
}
```

### Complex Pattern Matching
```unified
enum Message {
    Quit,
    Move { x: Int, y: Int },
    Write(String),
    ChangeColor(Int, Int, Int)
}

fn process_message(msg: Message) -> Int {
    match msg {
        Message::Quit -> {
            // Cleanup and exit
            return 0
        }
        Message::Move { x, y } -> {
            // Move to position
            return x + y
        }
        Message::Write(text) -> {
            // Write message
            return 1
        }
        Message::ChangeColor(r, g, b) if r > 0 -> {
            // Change color (only if r > 0)
            return r
        }
        Message::ChangeColor(_, _, _) -> {
            // Ignore invalid color
            return -1
        }
    }
}

fn main() -> Int {
    let msg1 = Message::Quit
    let msg2 = Message::Move { x: 10, y: 20 }
    let msg3 = Message::Write("Hello")
    let msg4 = Message::ChangeColor(255, 0, 0)
    
    return process_message(msg2)  // Returns 30
}
```

### Pattern Guards
```unified
enum Number {
    Int(Int),
    Float(Float)
}

fn classify(num: Number) -> String {
    match num {
        Number::Int(n) if n < 0 -> "negative integer"
        Number::Int(n) if n == 0 -> "zero"
        Number::Int(n) -> "positive integer"
        Number::Float(f) if f < 0.0 -> "negative float"
        Number::Float(f) -> "non-negative float"
    }
}
```

### Nested Pattern Matching
```unified
enum Option<T> {
    Some(T),
    None
}

enum Result<T, E> {
    Success(T),
    Error(E)
}

fn process(result: Result<Option<Int>, String>) -> Int {
    match result {
        Result::Success(Option::Some(value)) -> value
        Result::Success(Option::None) -> 0
        Result::Error(msg) -> -1
    }
}

fn main() -> Int {
    let r1 = Result::Success(Option::Some(42))
    let r2 = Result::Success(Option::None)
    let r3 = Result::Error("Failed")
    
    return process(r1)  // Returns 42
}
```

## 🏆 Success Criteria

- [ ] All parser tests pass (minimum 20 tests)
- [ ] All type system tests pass (minimum 15 tests)
- [ ] All pattern matching tests pass (minimum 25 tests)
- [ ] All exhaustiveness tests pass (minimum 10 tests)
- [ ] All Option tests pass (minimum 10 tests)
- [ ] All Result tests pass (minimum 10 tests)
- [ ] All integration tests pass (minimum 8 tests)
- [ ] Enums work end-to-end
- [ ] Pattern matching compiles correctly
- [ ] Exhaustiveness checking prevents bugs
- [ ] Can implement error handling with Result
- [ ] Can use Option for nullable values
- [ ] No regressions in previous phase tests
- [ ] Code coverage ≥ 85%
- [ ] Documentation complete
- [ ] All example programs run correctly

## 💡 Implementation Notes

### Implementation Order
1. Basic enum declarations (unit variants)
2. Enum type system
3. Tuple and struct variants
4. Match expression grammar
5. Simple pattern matching (literals, wildcards)
6. Enum pattern matching
7. Exhaustiveness checking
8. Pattern guards
9. Option and Result types
10. Integration and testing

### Testing Strategy
- Start with simple enums, add complexity gradually
- Test pattern matching incrementally
- Verify exhaustiveness checking thoroughly
- Test error messages for quality
- Use Option and Result in realistic scenarios

### Common Pitfalls
1. Incorrect tag-union memory layout
2. Exhaustiveness checking false positives/negatives
3. Pattern binding variable scope issues
4. Inefficient pattern matching codegen
5. Guard evaluation order
6. Nested pattern complexity
7. Generic enum type inference

### Debugging Tips
- Print pattern compilation decision tree
- Verify enum tag values are correct
- Check exhaustiveness algorithm carefully
- Test pattern matching bytecode manually
- Use simple examples before complex ones
- Verify Option and Result work correctly first

### Pattern Matching Algorithm
Consider using:
- Decision tree compilation for efficiency
- Backtracking for nested patterns
- Jump tables for simple enum switches
- Guard compilation to conditional jumps

### Memory Layout for Enums
```
[tag: 1 byte][data: max(variant sizes)]
```
- Tag identifies which variant
- Data area holds variant payload
- Size is max of all variant sizes + tag

### Future Considerations
This phase prepares for:
- ? operator for Result (Phase 7)
- Advanced pattern matching features
- Pattern matching on other types
- Const pattern matching
- Refutable vs irrefutable patterns

---

**Labels**: `phase-6`, `enhancement`, `types`, `enums`, `pattern-matching`  
**Milestone**: Phase 6 - Enums and Pattern Matching  
**Assignee**: TBD
