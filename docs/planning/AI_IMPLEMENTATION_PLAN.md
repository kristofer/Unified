# Unified Language - AI Agent Implementation Plan

**Version**: 2.0  
**Created**: January 2026  
**Purpose**: Detailed phase-by-phase implementation guide for AI agents building the Unified programming language

## Document Purpose

This document provides AI agents with a comprehensive, step-by-step plan to implement the Unified programming language. Each phase is designed to be:
- **Independently implementable** - Can be completed by an AI agent in a focused session
- **Fully testable** - Clear test requirements and success criteria
- **Incremental** - Builds on previous phases without breaking existing functionality
- **Documented** - Includes documentation updates as part of completion

## Current State (January 2026)

### ✅ Phase 1 Complete: VM-Based Minimal Compiler

**Status**: COMPLETE (76 tests passing)

**Implemented Features**:
- Stack-based virtual machine with 30+ bytecode instructions
- Bytecode generator (AST → bytecode)
- Function declarations with parameters
- Arithmetic operations (+, -, *, /, %, unary -)
- Comparison operations (==, !=, <, <=, >, >=)
- Logical operations (&&, ||, !)
- Local variables (let statements)
- Return statements
- Function calls with arguments
- Integer, float, boolean, string, and null literals
- Basic expressions
- If/else statement support (in VM/generator, parser needs update)

**Test Coverage**: 76 tests total
- 28 bytecode instruction tests
- 9 VM stack tests
- 19 VM execution tests
- 16 bytecode generator tests
- 3 integration tests

**Documentation**:
- VM_README.md (VM architecture)
- TESTING.md (test guide)
- PHASE1_SUMMARY.md (completion summary)

### 🎯 Next Phase: Phase 2

Phase 2 will extend the parser grammar and add control flow constructs.

---

## Phase Structure Template

Each phase follows this structure:

### Phase N: [Name]

**Goals**: Brief description of what this phase achieves

**Duration**: Estimated time (1-2 weeks typical)

**Prerequisites**: 
- List of required completed phases
- Required knowledge/tools

**Language Features Added**:
- Bullet list of new language features

**Implementation Tasks**:
1. Numbered list of specific tasks
2. Each task should be completable in 1-4 hours
3. Tasks should be ordered for sequential completion

**Test Requirements**:
- [ ] Specific test cases that must pass
- [ ] Coverage requirements
- [ ] Integration test requirements

**Documentation Requirements**:
- Files to create/update
- Examples to add
- Guides to write

**Success Criteria**:
- ✅ Measurable completion criteria
- ✅ Quality gates that must pass

**Example Programs**:
```unified
// Example code showing new features
```

**AI Agent Notes**:
- Specific guidance for AI agents
- Common pitfalls to avoid
- Hints for implementation

---

## Phase 2: Control Flow and Parser Grammar Updates

**Status**: NOT STARTED  
**Duration**: 1-2 weeks  
**Priority**: HIGH

### Goals

Complete the parser grammar to support control flow statements and prepare the foundation for more complex language features.

### Prerequisites

- Phase 1 complete (VM and bytecode generator working)
- ANTLR4 installed and working
- Understanding of ANTLR grammar syntax

### Language Features Added

1. **If/Else Statements** (complete implementation)
   - Parser grammar rules for if/else
   - Full AST support
   - Integration with existing VM support

2. **While Loops**
   - while condition { body }
   - break statement
   - continue statement

3. **For Loops**
   - for item in range { body }
   - Range expressions (start..end, start..=end)
   - Iterator protocol foundation

4. **Loop Statement**
   - Infinite loop with break
   - labeled breaks/continues

### Implementation Tasks

#### Task 2.1: Update Grammar for If/Else (2-3 hours)
1. Open `grammar/UnifiedParser.g4`
2. Add/verify if statement grammar rule:
   ```antlr
   ifStatement
       : IF expression block (ELSE block)?
       ;
   ```
3. Update statement rule to include ifStatement
4. Regenerate parser: `make parser`
5. Verify parser generation succeeds
6. Test parsing simple if statement

#### Task 2.2: Update AST for If/Else (1-2 hours)
1. Open `internal/ast/ast.go`
2. Add IfStatement struct if not present:
   ```go
   type IfStatement struct {
       Condition  Expression
       ThenBlock  *Block
       ElseBlock  *Block // nullable
   }
   ```
3. Update visitor to handle if statements
4. Add unit tests for AST construction

#### Task 2.3: Add While Loop Grammar (2-3 hours)
1. Add to `grammar/UnifiedParser.g4`:
   ```antlr
   whileStatement
       : WHILE expression block
       ;
   ```
2. Add break and continue statements:
   ```antlr
   breakStatement
       : BREAK SEMICOLON?
       ;
   
   continueStatement
       : CONTINUE SEMICOLON?
       ;
   ```
3. Update statement rule
4. Regenerate parser
5. Test parsing

#### Task 2.4: Add While Loop Support to AST/VM (3-4 hours)
1. Add AST nodes for while, break, continue
2. Add bytecode instructions if needed:
   - May reuse existing JUMP instructions
   - Track loop start/end labels
3. Update bytecode generator:
   - Generate loop entry label
   - Generate condition check
   - Generate JUMP_IF_FALSE to loop end
   - Generate loop body
   - Generate JUMP to loop start
   - Generate loop end label
4. Handle break/continue:
   - Track loop context stack
   - Generate jumps to appropriate labels

#### Task 2.5: Add For Loop and Range Support (4-6 hours)
1. Add grammar for range expressions:
   ```antlr
   rangeExpression
       : expression DOTDOT expression          // exclusive
       | expression DOTDOT_EQ expression       // inclusive
       ;
   ```
2. Add for loop grammar:
   ```antlr
   forStatement
       : FOR IDENTIFIER IN expression block
       ;
   ```
3. Implement range as iterator:
   - Create range object with start, end, current
   - Implement next() method concept
   - Generate bytecode for iteration
4. Update VM to support iteration:
   - May need new instructions for iterator protocol
   - Or desugar to while loop in generator

#### Task 2.6: Add Loop Statement (2-3 hours)
1. Add grammar:
   ```antlr
   loopStatement
       : LOOP block
       ;
   ```
2. Implement as infinite loop
3. Require break to exit
4. Test with labeled breaks if time permits

#### Task 2.7: Write Comprehensive Tests (3-4 hours)
1. Parser tests:
   - Test if/else parsing
   - Test while loop parsing
   - Test for loop parsing
   - Test loop statement parsing
   - Test break/continue parsing
2. Generator tests:
   - Test bytecode for each construct
   - Test nested loops
   - Test break/continue in nested loops
3. VM tests:
   - Test loop execution
   - Test early exit with break
   - Test continue skip
4. Integration tests:
   - Factorial with while loop
   - Sum with for loop
   - Search with early break

#### Task 2.8: Update Documentation (2-3 hours)
1. Update PHASE1_SUMMARY.md → rename to IMPLEMENTATION_HISTORY.md
2. Create PHASE2_SUMMARY.md documenting:
   - Features added
   - Test coverage
   - Example programs
3. Update README.md with new capabilities
4. Add example programs to `test/` directory:
   - `while_loop.uni`
   - `for_loop.uni`
   - `nested_loops.uni`
   - `factorial.uni`

### Test Requirements

**Minimum Test Coverage**: 90% for new code

**Parser Tests** (grammar/AST):
- [ ] Parse if statement without else
- [ ] Parse if statement with else
- [ ] Parse nested if statements
- [ ] Parse while loop
- [ ] Parse for loop with range
- [ ] Parse loop statement
- [ ] Parse break statement
- [ ] Parse continue statement
- [ ] Parse labeled break/continue (if implemented)

**Generator Tests** (bytecode generation):
- [ ] Generate correct bytecode for if without else
- [ ] Generate correct bytecode for if with else
- [ ] Generate correct jumps for while loop
- [ ] Generate correct iteration for for loop
- [ ] Generate correct infinite loop
- [ ] Generate correct break jumps
- [ ] Generate correct continue jumps
- [ ] Handle nested loops correctly

**VM Tests** (execution):
- [ ] Execute if statement (both branches)
- [ ] Execute while loop (multiple iterations)
- [ ] Execute for loop over range
- [ ] Execute infinite loop with break
- [ ] Execute break in while loop
- [ ] Execute continue in while loop
- [ ] Execute nested loops correctly

**Integration Tests** (end-to-end):
- [ ] Compile and run factorial program (while loop)
- [ ] Compile and run sum program (for loop)
- [ ] Compile and run search program (early break)
- [ ] Compile and run FizzBuzz (if statements in loop)

### Documentation Requirements

1. **Update existing docs**:
   - README.md: Add control flow to feature list
   - TESTING.md: Add new test categories
   - VM_README.md: Document any new bytecode instructions

2. **Create new docs**:
   - PHASE2_SUMMARY.md: Complete phase summary
   - examples/CONTROL_FLOW.md: Guide to control flow features

3. **Add example programs**:
   - `test/while_factorial.uni`: Factorial using while
   - `test/for_sum.uni`: Sum using for loop
   - `test/nested_loops.uni`: Nested loop example
   - `test/fizzbuzz_complete.uni`: Complete FizzBuzz

### Success Criteria

- ✅ All new parser tests pass (minimum 15 tests)
- ✅ All new generator tests pass (minimum 12 tests)
- ✅ All new VM tests pass (minimum 10 tests)
- ✅ All integration tests pass (minimum 4 tests)
- ✅ Can compile and run factorial program
- ✅ Can compile and run FizzBuzz program
- ✅ No regressions in Phase 1 tests
- ✅ Code coverage ≥ 85%
- ✅ Documentation updated
- ✅ Example programs run correctly

### Example Programs

**Factorial with While Loop**:
```unified
fn factorial(n: Int) -> Int {
    let mut result = 1
    let mut i = n
    while i > 1 {
        result = result * i
        i = i - 1
    }
    return result
}

fn main() -> Int {
    return factorial(5)  // Returns 120
}
```

**Sum with For Loop**:
```unified
fn sum_range(start: Int, end: Int) -> Int {
    let mut total = 0
    for i in start..end {
        total = total + i
    }
    return total
}

fn main() -> Int {
    return sum_range(1, 11)  // Returns 55 (sum of 1-10)
}
```

**FizzBuzz**:
```unified
fn fizzbuzz(n: Int) -> Int {
    for i in 1..=n {
        if i % 15 == 0 {
            // print("FizzBuzz")
            let x = 0  // Placeholder until print works
        } else if i % 3 == 0 {
            // print("Fizz")
            let x = 0
        } else if i % 5 == 0 {
            // print("Buzz")
            let x = 0
        } else {
            // print(i)
            let x = 0
        }
    }
    return n
}

fn main() -> Int {
    return fizzbuzz(15)
}
```

### AI Agent Notes

**Implementation Order**: Follow tasks in sequence. Grammar first, then AST, then generator, then VM updates.

**Testing Strategy**: 
- Write tests immediately after implementing each feature
- Use table-driven tests for multiple cases
- Test edge cases (empty loops, single iteration, etc.)

**Common Pitfalls**:
1. Forgetting to track loop labels for break/continue
2. Not handling empty else blocks correctly
3. Off-by-one errors in range iteration
4. Not testing nested loop scenarios

**Debugging Tips**:
- Use `make test` frequently to catch regressions
- Add debug output to bytecode generator to see generated code
- Use integration tests to verify end-to-end behavior

---

## Phase 3: Variables, Mutability, and Assignment

**Status**: NOT STARTED  
**Duration**: 2 weeks  
**Priority**: HIGH

### Goals

Implement proper variable declaration, mutability tracking, and assignment operations. This phase enhances the variable system beyond basic `let` statements.

### Prerequisites

- Phase 2 complete (control flow working)
- Understanding of ownership concepts (for future phases)

### Language Features Added

1. **Mutable Variables**
   - `let mut` declarations
   - Mutability checking at compile time
   - Error on assignment to immutable variables

2. **Assignment Statements**
   - Simple assignment: `x = value`
   - Compound assignment: `+=`, `-=`, `*=`, `/=`, `%=`
   - Validation of assignment targets

3. **Variable Shadowing**
   - Allow redeclaration in nested scopes
   - Track scope correctly

4. **Type Inference**
   - Basic type inference for let statements
   - Type checking for assignments

### Implementation Tasks

#### Task 3.1: Add Mutability to Grammar (1-2 hours)
1. Update let statement grammar:
   ```antlr
   letStatement
       : LET MUT? IDENTIFIER (COLON type)? ASSIGN expression SEMICOLON?
       ;
   ```
2. Add assignment statement:
   ```antlr
   assignmentStatement
       : IDENTIFIER assignmentOp expression SEMICOLON?
       ;
   
   assignmentOp
       : ASSIGN | PLUS_ASSIGN | MINUS_ASSIGN | STAR_ASSIGN 
       | SLASH_ASSIGN | PERCENT_ASSIGN
       ;
   ```
3. Regenerate parser

#### Task 3.2: Update AST for Mutability (2 hours)
1. Add mutable flag to LetStatement:
   ```go
   type LetStatement struct {
       Name      string
       Type      Type // may be nil for inference
       Value     Expression
       Mutable   bool
   }
   ```
2. Add AssignmentStatement:
   ```go
   type AssignmentStatement struct {
       Target   string
       Operator AssignOp // =, +=, -=, etc.
       Value    Expression
   }
   ```

#### Task 3.3: Implement Symbol Table (4-6 hours)
1. Create symbol table structure:
   ```go
   type Symbol struct {
       Name    string
       Type    Type
       Mutable bool
       Defined bool
   }
   
   type SymbolTable struct {
       scopes  []map[string]*Symbol
       current int
   }
   ```
2. Implement scope management:
   - EnterScope()
   - ExitScope()
   - Define(name, type, mutable)
   - Lookup(name) -> Symbol
3. Add to bytecode generator
4. Track variable mutability

#### Task 3.4: Add Semantic Checking (3-4 hours)
1. Check assignment to immutable variables:
   - Lookup variable in symbol table
   - Error if not mutable
2. Check undefined variable errors
3. Add error reporting with line numbers
4. Create comprehensive error messages

#### Task 3.5: Implement Type Inference (4-6 hours)
1. Infer types from literal values:
   - Int literals → Int type
   - Float literals → Float type
   - Bool literals → Bool type
   - String literals → String type
2. Infer types from expressions:
   - Binary ops → result type
   - Function calls → return type
3. Store inferred types in symbol table
4. Validate type compatibility in assignments

#### Task 3.6: Update VM for Assignment (2-3 hours)
1. Implement STORE_LOCAL correctly for reassignment
2. Verify mutable variables can be updated
3. Test that immutable variables cannot be updated

#### Task 3.7: Write Tests (4-5 hours)
1. Parser tests for mut keyword and assignments
2. Symbol table tests
3. Semantic checking tests (error cases)
4. Type inference tests
5. Integration tests with mutable variables

#### Task 3.8: Update Documentation (2 hours)
1. Document mutability in language guide
2. Create PHASE3_SUMMARY.md
3. Add examples to test directory
4. Update error message documentation

### Test Requirements

**Parser Tests**:
- [ ] Parse `let mut x = 5`
- [ ] Parse assignment `x = 10`
- [ ] Parse compound assignment `x += 5`
- [ ] Parse all assignment operators

**Symbol Table Tests**:
- [ ] Define variable in scope
- [ ] Lookup variable in current scope
- [ ] Lookup variable in parent scope
- [ ] Shadow variable in nested scope
- [ ] Scope isolation (can't access child scope vars)

**Semantic Analysis Tests**:
- [ ] Error on assignment to immutable variable
- [ ] Error on assignment to undefined variable
- [ ] Error on type mismatch in assignment
- [ ] Allow assignment to mutable variable
- [ ] Allow variable shadowing

**Type Inference Tests**:
- [ ] Infer Int from integer literal
- [ ] Infer Float from float literal
- [ ] Infer Bool from boolean literal
- [ ] Infer type from arithmetic expression
- [ ] Detect type mismatch

**Integration Tests**:
- [ ] Counter with mutable variable
- [ ] Compound assignment in loop
- [ ] Variable shadowing in nested scopes

### Success Criteria

- ✅ Mutability tracking works correctly
- ✅ Assignments to immutable variables are caught at compile time
- ✅ Type inference works for basic cases
- ✅ Symbol table correctly handles scoping
- ✅ All tests pass (minimum 30 new tests)
- ✅ No regressions
- ✅ Documentation complete

### Example Programs

**Counter with Mutable Variable**:
```unified
fn count_to(n: Int) -> Int {
    let mut counter = 0
    while counter < n {
        counter += 1
    }
    return counter
}

fn main() -> Int {
    return count_to(10)  // Returns 10
}
```

**Type Inference**:
```unified
fn demonstrate_inference() -> Int {
    let x = 42          // Inferred as Int
    let y = 3.14        // Inferred as Float
    let z = true        // Inferred as Bool
    let s = "hello"     // Inferred as String
    
    let mut counter = 0
    counter += x        // Type checked
    
    return counter
}
```

### AI Agent Notes

**Critical**: Implement symbol table before semantic checking. The symbol table is the foundation for all future type checking and ownership analysis.

**Testing**: Focus on error cases as much as success cases. Good error messages are critical for usability.

---

## Phase 4: Functions and Advanced Expressions

**Status**: NOT STARTED  
**Duration**: 2-3 weeks  
**Priority**: MEDIUM

### Goals

Enhance function support with more features and implement advanced expressions including blocks as expressions, function expressions, and operator precedence.

### Prerequisites

- Phase 3 complete (variables and mutability working)

### Language Features Added

1. **Function Features**
   - Multiple return values (via tuples)
   - Default parameter values
   - Named parameters
   - Variadic parameters (basic)

2. **Block Expressions**
   - Blocks that return values
   - Last expression as return value
   - Scoped blocks

3. **Lambda Expressions**
   - Anonymous functions
   - Closures (basic, no capture initially)
   - First-class functions

4. **Advanced Operators**
   - Bitwise operators (&, |, ^, ~, <<, >>)
   - Operator precedence handling
   - Unary operators (!, -, +)

### Implementation Tasks

#### Task 4.1: Multiple Return Values (4-6 hours)
1. Add tuple type to type system
2. Support tuple literals in parser
3. Support tuple destructuring in let statements:
   ```unified
   let (a, b) = divide(10, 3)  // Returns (quotient, remainder)
   ```
4. Implement in VM (return multiple values on stack)

#### Task 4.2: Default Parameters (3-4 hours)
1. Add default value to parameter AST
2. Generate bytecode to handle missing arguments
3. Validate call sites

#### Task 4.3: Block Expressions (3-4 hours)
1. Update block to be an expression
2. Return last expression value
3. Handle void blocks
4. Test in various contexts

#### Task 4.4: Lambda Expressions (6-8 hours)
1. Add lambda syntax to grammar:
   ```antlr
   lambdaExpr
       : PIPE paramList? PIPE (ARROW expression | block)
       ;
   ```
2. Create function objects in VM
3. Support passing functions as parameters
4. Test higher-order functions

#### Task 4.5: Bitwise Operators (2-3 hours)
1. Add operators to grammar
2. Add bytecode instructions
3. Implement in VM
4. Test all operators

#### Task 4.6: Write Tests (4-5 hours)
Comprehensive tests for all new features

#### Task 4.7: Documentation (2-3 hours)
Document all new features with examples

### Test Requirements

- [ ] Tuple return values work
- [ ] Default parameters work
- [ ] Blocks as expressions work
- [ ] Lambda expressions work
- [ ] Lambdas can be passed to functions
- [ ] Bitwise operators work correctly
- [ ] Operator precedence is correct

### Success Criteria

- ✅ All advanced features work
- ✅ Can implement map/filter using lambdas
- ✅ Test coverage ≥ 85%
- ✅ Documentation complete

### Example Programs

**Higher-Order Functions**:
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

**Lambda Expression**:
```unified
fn main() -> Int {
    let add = |a, b| a + b
    return add(10, 20)  // Returns 30
}
```

---

## Phase 5: Structs and Basic Types

**Status**: NOT STARTED  
**Duration**: 3-4 weeks  
**Priority**: MEDIUM

### Goals

Implement user-defined struct types with fields and methods. This is the foundation for object-oriented features.

### Prerequisites

- Phase 4 complete (functions and expressions working)

### Language Features Added

1. **Struct Declarations**
   - Field definitions with types
   - Struct instantiation
   - Field access (dot notation)

2. **Methods**
   - Methods on structs
   - self parameter
   - Mutable self (&mut self)

3. **Struct Methods**
   - Associated functions (static methods)
   - Instance methods
   - Method call syntax

### Implementation Tasks

#### Task 5.1: Struct Grammar (3-4 hours)
1. Add struct declaration:
   ```antlr
   structDecl
       : STRUCT IDENTIFIER genericParams? LBRACE structBody RBRACE
       ;
   
   structBody
       : (fieldDecl | methodDecl)*
       ;
   
   fieldDecl
       : IDENTIFIER COLON type SEMICOLON?
       ;
   ```

#### Task 5.2: Struct Type System (4-6 hours)
1. Add StructType to type system
2. Track field names and types
3. Type checking for field access
4. Memory layout for structs

#### Task 5.3: Struct Instantiation (4-5 hours)
1. Add struct literal syntax:
   ```unified
   Point { x: 10, y: 20 }
   ```
2. Generate bytecode for struct creation
3. Allocate struct in VM memory

#### Task 5.4: Field Access (3-4 hours)
1. Add dot notation to grammar
2. Generate field offset calculations
3. Implement field load/store in VM

#### Task 5.5: Methods (6-8 hours)
1. Add method declarations
2. Implicit self parameter
3. Method calls with self
4. Mutable vs immutable self

#### Task 5.6: Tests and Documentation (4-5 hours)

### Test Requirements

- [ ] Declare struct types
- [ ] Instantiate structs
- [ ] Access fields
- [ ] Call methods
- [ ] Mutable methods modify state

### Success Criteria

- ✅ Struct types work end-to-end
- ✅ Methods can access and modify fields
- ✅ Type checking for structs

### Example Programs

**Point Struct**:
```unified
struct Point {
    x: Float
    y: Float
    
    fn distance(self) -> Float {
        return (self.x * self.x + self.y * self.y)  // Simplified, no sqrt
    }
    
    fn move_by(self: &mut Self, dx: Float, dy: Float) {
        self.x += dx
        self.y += dy
    }
}

fn main() -> Int {
    let mut p = Point { x: 3.0, y: 4.0 }
    let dist = p.distance()  // 25.0 (without sqrt)
    p.move_by(1.0, 1.0)
    return 0
}
```

---

## Phase 6: Enums and Pattern Matching

**Status**: NOT STARTED  
**Duration**: 3-4 weeks  
**Priority**: MEDIUM

### Goals

Implement algebraic data types (enums) and pattern matching, enabling Result and Option types.

### Language Features Added

1. **Enum Declarations**
   - Simple enums (unit variants)
   - Enums with data (tuple variants)
   - Enums with named fields (struct variants)

2. **Pattern Matching**
   - Match expressions
   - Patterns for literals, variables, structs, enums
   - Pattern guards
   - Exhaustiveness checking

3. **Standard Enums**
   - Option<T>
   - Result<T, E>

### Implementation Tasks

#### Task 6.1: Enum Grammar (3-4 hours)
```antlr
enumDecl
    : ENUM IDENTIFIER genericParams? LBRACE enumVariants RBRACE
    ;

enumVariant
    : IDENTIFIER (tupleFields | structFields)?
    ;
```

#### Task 6.2: Enum Type System (6-8 hours)
1. Tag-union representation
2. Variant constructors
3. Type checking

#### Task 6.3: Match Expression (8-10 hours)
1. Match syntax in grammar
2. Pattern compilation
3. Exhaustiveness checking
4. Jump table generation

#### Task 6.4: Option and Result (4-6 hours)
1. Define in prelude
2. Methods (map, and_then, unwrap, etc.)
3. Integration with ? operator (Phase 7)

#### Task 6.5: Tests and Documentation (4-5 hours)

### Test Requirements

- [ ] Define enum types
- [ ] Create enum values
- [ ] Match on enums
- [ ] Extract enum data in patterns
- [ ] Exhaustiveness checking works
- [ ] Option and Result work

### Success Criteria

- ✅ Enums work end-to-end
- ✅ Pattern matching compiles correctly
- ✅ Can implement error handling with Result

### Example Programs

**Result Type Usage**:
```unified
enum Result<T, E> {
    Success(T)
    Error(E)
}

fn divide(a: Int, b: Int) -> Result<Int, String> {
    if b == 0 {
        return Result.Error("Division by zero")
    }
    return Result.Success(a / b)
}

fn main() -> Int {
    let result = divide(10, 2)
    match result {
        Success(value) -> return value
        Error(msg) -> return -1
    }
}
```

---

## Phase 7: Error Handling with ? Operator

**Status**: NOT STARTED  
**Duration**: 1-2 weeks  
**Priority**: MEDIUM

### Goals

Implement the ? operator for ergonomic error propagation with Result types.

### Prerequisites

- Phase 6 complete (enums and pattern matching)

### Language Features Added

1. **? Operator**
   - Automatic error propagation
   - Early return on Error
   - Unwrap Success value

### Implementation Tasks

#### Task 7.1: Add ? to Grammar (1 hour)
```antlr
postfixExpr
    : primaryExpr (QUESTION | DOT IDENTIFIER | callExpr)*
    ;
```

#### Task 7.2: Implement Desugaring (4-6 hours)
Desugar `expr?` to:
```unified
match expr {
    Success(value) -> value
    Error(e) -> return Error(e)
}
```

#### Task 7.3: Type Checking (3-4 hours)
1. Verify expr is Result type
2. Verify function returns compatible Result
3. Extract inner type for use

#### Task 7.4: Tests and Documentation (2-3 hours)

### Test Requirements

- [ ] ? operator unwraps Success
- [ ] ? operator returns Error early
- [ ] Chain multiple ? operators
- [ ] Type checking works

### Success Criteria

- ✅ ? operator works correctly
- ✅ Error handling is ergonomic

### Example Programs

**Chained Error Handling**:
```unified
fn process_data() -> Result<Int, String> {
    let file = open_file("data.txt")?
    let content = read_file(file)?
    let value = parse_int(content)?
    return Success(value * 2)
}
```

---

## Phase 8: Arrays and Slices

**Status**: NOT STARTED  
**Duration**: 2-3 weeks  
**Priority**: HIGH

### Goals

Implement fixed-size arrays and dynamic slices with bounds checking.

### Language Features Added

1. **Array Types**
   - Fixed-size arrays: `[Int; 10]`
   - Array literals: `[1, 2, 3, 4, 5]`
   - Array indexing with bounds checking

2. **Slice Types**
   - Dynamic slices: `[Int]`
   - Slice operations: `arr[1..5]`
   - Slice methods (len, is_empty)

3. **For-Each Loops**
   - Iterate over arrays/slices
   - Iterator protocol

### Implementation Tasks

#### Task 8.1: Array Grammar and Types (4-6 hours)
```antlr
arrayType
    : LBRACKET type SEMICOLON INTEGER RBRACKET   // Fixed size
    | LBRACKET type RBRACKET                      // Slice
    ;

arrayLiteral
    : LBRACKET (expression (COMMA expression)*)? RBRACKET
    ;
```

#### Task 8.2: Array Memory Layout (4-5 hours)
1. Contiguous memory allocation
2. Bounds checking
3. Index calculation

#### Task 8.3: Slice Implementation (6-8 hours)
1. Slice as (pointer, length)
2. Subslicing
3. Bounds checking

#### Task 8.4: Array/Slice Iteration (4-6 hours)
1. Update for loops to support arrays
2. Index-based iteration
3. Value iteration

#### Task 8.5: Tests and Documentation (4-5 hours)

### Test Requirements

- [ ] Create fixed-size arrays
- [ ] Array literal syntax works
- [ ] Array indexing works
- [ ] Bounds checking catches errors
- [ ] Slice creation works
- [ ] Iteration over arrays works

### Success Criteria

- ✅ Arrays and slices work correctly
- ✅ Bounds checking prevents errors
- ✅ Can iterate over collections

### Example Programs

**Array Processing**:
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
    return sum_array(numbers)  // Returns 15
}
```

---

## Phase 9: String Operations

**Status**: NOT STARTED  
**Duration**: 2 weeks  
**Priority**: MEDIUM

### Goals

Implement comprehensive string operations and string interpolation.

### Language Features Added

1. **String Methods**
   - len(), is_empty()
   - substring(), split()
   - to_upper(), to_lower()
   - trim(), replace()

2. **String Interpolation**
   - `"Hello, ${name}!"`
   - Expression interpolation
   - Format specifiers

3. **String Concatenation**
   - + operator for strings
   - Efficient string building

### Implementation Tasks

#### Task 9.1: String Representation (3-4 hours)
1. String as (pointer, length, capacity)
2. UTF-8 handling
3. Memory management

#### Task 9.2: String Methods (6-8 hours)
Implement core string methods as built-ins

#### Task 9.3: String Interpolation (4-6 hours)
1. Parse ${} syntax
2. Desugar to concatenation
3. Format expressions

#### Task 9.4: Tests and Documentation (3-4 hours)

### Test Requirements

- [ ] String methods work
- [ ] Interpolation works
- [ ] Concatenation works
- [ ] UTF-8 handling correct

### Success Criteria

- ✅ String operations work
- ✅ Interpolation is convenient

---

## Phase 10: Generics Basics

**Status**: NOT STARTED  
**Duration**: 3-4 weeks  
**Priority**: MEDIUM

### Goals

Implement basic generic types and functions with type parameters.

### Language Features Added

1. **Generic Functions**
   - Type parameters: `fn identity<T>(x: T) -> T`
   - Type inference for generics
   - Monomorphization

2. **Generic Structs**
   - `struct Box<T> { value: T }`
   - Generic instantiation

3. **Generic Enums**
   - `enum Option<T> { Some(T), None }`
   - Already used Result and Option

### Implementation Tasks

#### Task 10.1: Generic Syntax (3-4 hours)
Add type parameters to grammar

#### Task 10.2: Type Parameter System (6-8 hours)
1. Type variables
2. Type substitution
3. Type inference for generics

#### Task 10.3: Monomorphization (8-10 hours)
1. Generate specialized versions
2. Type erasure considerations
3. Code generation

#### Task 10.4: Tests and Documentation (4-5 hours)

### Test Requirements

- [ ] Generic functions work
- [ ] Generic structs work
- [ ] Type inference works
- [ ] Monomorphization generates correct code

### Success Criteria

- ✅ Generics work for common cases
- ✅ Type inference minimizes annotations

---

## Phases 11-15: Advanced Features (Summary)

### Phase 11: Modules and Imports (2-3 weeks)
- Module system
- Import/export
- Visibility (pub)

### Phase 12: Basic Ownership (3-4 weeks)
- Ownership tracking
- Move semantics
- Borrow checker (basic)

### Phase 13: Standard Library Foundation (2-3 weeks)
- Collections (List, Map, Set)
- I/O primitives
- String utilities

### Phase 14: Concurrency Basics (4-6 weeks)
- Async/await
- Actors (basic)
- Channels

### Phase 15: Tooling and Polish (3-4 weeks)
- Better error messages
- Code formatter
- Documentation generator
- Package manager basics

---

## Implementation Guidelines for AI Agents

### Before Starting Any Phase

1. **Read Phase Requirements Completely**
   - Understand all goals
   - Review prerequisites
   - Check example programs

2. **Check Current State**
   ```bash
   cd /home/runner/work/Unified/Unified/src/unified-compiler
   git status
   make test
   ```

3. **Create Feature Branch** (if not already on one)
   ```bash
   git checkout -b phase-N-feature-name
   ```

### During Implementation

1. **Follow Task Order**
   - Complete tasks sequentially
   - Don't skip steps
   - Test after each major task

2. **Write Tests First** (TDD)
   - Write failing test
   - Implement feature
   - Verify test passes
   - Refactor if needed

3. **Commit Frequently**
   - Commit after each task
   - Use descriptive commit messages
   - Use `report_progress` tool

4. **Run Tests Continuously**
   ```bash
   make test                    # All tests
   go test ./internal/bytecode  # Specific module
   go test -v ./...             # Verbose
   ```

### After Completing Phase

1. **Verify All Tests Pass**
   ```bash
   make test
   # Should show 100% passing
   ```

2. **Update Documentation**
   - Create PHASE_N_SUMMARY.md
   - Update README.md
   - Add examples
   - Update TESTING.md

3. **Create Summary**
   Document:
   - Features implemented
   - Tests added (count)
   - Example programs
   - Challenges encountered
   - Next phase prerequisites

4. **Code Review** (use code_review tool)
   - Request review
   - Address feedback
   - Get approval

### Testing Best Practices

1. **Test Categories**
   - Unit tests (per module)
   - Integration tests (end-to-end)
   - Error case tests
   - Edge case tests

2. **Test Naming**
   ```go
   func TestFeatureName(t *testing.T)
   func TestFeatureName_EdgeCase(t *testing.T)
   func TestFeatureName_ErrorCase(t *testing.T)
   ```

3. **Table-Driven Tests**
   ```go
   tests := []struct {
       name     string
       input    string
       expected int
   }{
       {"case1", "input1", 42},
       {"case2", "input2", 100},
   }
   for _, tt := range tests {
       t.Run(tt.name, func(t *testing.T) {
           // test implementation
       })
   }
   ```

### Common Pitfalls to Avoid

1. **Don't Skip Grammar Updates**
   - Always regenerate parser after grammar changes
   - Test parser separately before VM changes

2. **Don't Ignore Type System**
   - Add proper type checking
   - Handle type errors gracefully

3. **Don't Forget Error Cases**
   - Test invalid input
   - Test boundary conditions
   - Verify error messages

4. **Don't Break Existing Tests**
   - Run full test suite before committing
   - Fix regressions immediately

5. **Don't Skip Documentation**
   - Document as you code
   - Add examples for new features
   - Update guides and references

### Code Quality Standards

1. **Code Style**
   - Follow Go conventions
   - Use `gofmt`
   - Run `go vet`

2. **Comments**
   - Document public APIs
   - Explain complex algorithms
   - Add TODOs for future work

3. **Error Handling**
   - Return errors, don't panic
   - Provide context in errors
   - Use descriptive error messages

4. **Test Coverage**
   - Aim for ≥85% coverage
   - Cover error paths
   - Test edge cases

### Performance Considerations

1. **Compiler Performance**
   - Compilation should be fast (<1s for small programs)
   - Use efficient data structures
   - Avoid unnecessary allocations

2. **Runtime Performance**
   - VM should be reasonably fast
   - Optimize hot paths
   - Profile before optimizing

3. **Memory Usage**
   - Minimize allocations in VM
   - Reuse objects when possible
   - Clean up resources properly

---

## Phase Dependencies Graph

```
Phase 1: VM Foundation
    ↓
Phase 2: Control Flow ← (you are here next)
    ↓
Phase 3: Variables & Mutability
    ↓
Phase 4: Functions & Expressions
    ↓
Phase 5: Structs ← Phase 6: Enums
    ↓              ↓
    ↓          Phase 7: ? Operator
    ↓              ↓
Phase 8: Arrays ← ┘
    ↓
Phase 9: Strings
    ↓
Phase 10: Generics
    ↓
Phase 11: Modules
    ↓
Phase 12: Ownership
    ↓
Phase 13: Std Library
    ↓
Phase 14: Concurrency
    ↓
Phase 15: Tooling
```

---

## Quick Reference

### File Locations

- **Grammar**: `/src/unified-compiler/grammar/*.g4`
- **AST**: `/src/unified-compiler/internal/ast/*.go`
- **Parser**: `/src/unified-compiler/internal/parser/*.go` (generated)
- **Bytecode**: `/src/unified-compiler/internal/bytecode/*.go`
- **VM**: `/src/unified-compiler/internal/vm/*.go`
- **Tests**: `/src/unified-compiler/*/`*_test.go`
- **Integration Tests**: `/src/unified-compiler/cmd/compiler/integration_test.go`
- **Test Programs**: `/src/unified-compiler/test/*.uni`

### Common Commands

```bash
# Build
cd src/unified-compiler
make parser    # Regenerate parser from grammar
make build     # Build compiler binary
make all       # parser + build

# Test
make test      # Run all tests
go test -v ./internal/vm     # Test specific module
go test -cover ./...         # With coverage

# Run
./bin/unified-compiler --input test/fib.uni --output output.bc

# Clean
make clean     # Remove generated files
```

### Key Concepts

**AST**: Abstract Syntax Tree - represents program structure  
**Bytecode**: Low-level instructions for the VM  
**VM**: Virtual Machine - executes bytecode  
**Symbol Table**: Tracks variables and their properties  
**Type Checking**: Verifies type correctness  
**Monomorphization**: Generate specialized generic code  

---

## Conclusion

This implementation plan provides a clear path from the current Phase 1 completion to a full-featured programming language. Each phase is designed to be independently implementable by an AI agent while building incrementally on previous work.

**Key Success Factors**:
1. Follow phases in order
2. Write tests first
3. Document as you go
4. Commit frequently
5. Don't break existing functionality

**Current Priority**: Begin Phase 2 (Control Flow) immediately.

Good luck implementing Unified! 🚀
