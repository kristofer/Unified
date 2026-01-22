# Phase 2: Control Flow and Parser Grammar Updates

**Status**: Not Started  
**Duration**: 1-2 weeks  
**Priority**: HIGH  
**Dependencies**: Phase 1 Complete

## 🎯 Goals

Complete the parser grammar to support control flow statements and prepare the foundation for more complex language features.

## 📋 Prerequisites

- [x] Phase 1 complete (VM and bytecode generator working)
- [ ] ANTLR4 installed and working
- [ ] Understanding of ANTLR grammar syntax

## ✨ Language Features to Add

### 1. If/Else Statements (Complete Implementation)
- Parser grammar rules for if/else
- Full AST support
- Integration with existing VM support

### 2. While Loops
- `while condition { body }`
- `break` statement
- `continue` statement

### 3. For Loops
- `for item in range { body }`
- Range expressions (`start..end`, `start..=end`)
- Iterator protocol foundation

### 4. Loop Statement
- Infinite loop with break
- Labeled breaks/continues

## 📝 Implementation Tasks

### Task 2.1: Update Grammar for If/Else (2-3 hours)
- [ ] Open `grammar/UnifiedParser.g4`
- [ ] Add/verify if statement grammar rule
- [ ] Update statement rule to include ifStatement
- [ ] Regenerate parser: `make parser`
- [ ] Verify parser generation succeeds
- [ ] Test parsing simple if statement

**Grammar to add:**
```antlr
ifStatement
    : IF expression block (ELSE block)?
    ;
```

### Task 2.2: Update AST for If/Else (1-2 hours)
- [ ] Open `internal/ast/ast.go`
- [ ] Add IfStatement struct if not present
- [ ] Update visitor to handle if statements
- [ ] Add unit tests for AST construction

**AST Structure:**
```go
type IfStatement struct {
    Condition  Expression
    ThenBlock  *Block
    ElseBlock  *Block // nullable
}
```

### Task 2.3: Add While Loop Grammar (2-3 hours)
- [ ] Add while statement to grammar
- [ ] Add break and continue statements
- [ ] Update statement rule
- [ ] Regenerate parser
- [ ] Test parsing

**Grammar to add:**
```antlr
whileStatement
    : WHILE expression block
    ;

breakStatement
    : BREAK SEMICOLON?
    ;

continueStatement
    : CONTINUE SEMICOLON?
    ;
```

### Task 2.4: Add While Loop Support to AST/VM (3-4 hours)
- [ ] Add AST nodes for while, break, continue
- [ ] Add bytecode instructions if needed
- [ ] Update bytecode generator for loops
- [ ] Handle break/continue with loop context tracking

### Task 2.5: Add For Loop and Range Support (4-6 hours)
- [ ] Add grammar for range expressions
- [ ] Add for loop grammar
- [ ] Implement range as iterator
- [ ] Update VM to support iteration

**Grammar to add:**
```antlr
rangeExpression
    : expression DOTDOT expression          // exclusive (1..5)
    | expression DOTDOT_EQ expression       // inclusive (1..=5)
    ;

forStatement
    : FOR IDENTIFIER IN expression block
    ;
```

### Task 2.6: Add Loop Statement (2-3 hours)
- [ ] Add loop statement grammar
- [ ] Implement as infinite loop
- [ ] Require break to exit
- [ ] Test with labeled breaks if time permits

**Grammar:**
```antlr
loopStatement
    : LOOP block
    ;
```

### Task 2.7: Write Comprehensive Tests (3-4 hours)
- [ ] Parser tests (9+ tests)
- [ ] Generator tests (8+ tests)
- [ ] VM tests (7+ tests)
- [ ] Integration tests (4+ tests)

### Task 2.8: Update Documentation (2-3 hours)
- [ ] Update PHASE1_SUMMARY.md → rename to IMPLEMENTATION_HISTORY.md
- [ ] Create PHASE2_SUMMARY.md
- [ ] Update README.md with new capabilities
- [ ] Add example programs to test directory

## ✅ Test Requirements

**Minimum Test Coverage**: 90% for new code

### Parser Tests (Grammar/AST)
- [ ] Parse if statement without else
- [ ] Parse if statement with else
- [ ] Parse nested if statements
- [ ] Parse while loop
- [ ] Parse for loop with range
- [ ] Parse loop statement
- [ ] Parse break statement
- [ ] Parse continue statement
- [ ] Parse labeled break/continue (if implemented)

### Generator Tests (Bytecode Generation)
- [ ] Generate correct bytecode for if without else
- [ ] Generate correct bytecode for if with else
- [ ] Generate correct jumps for while loop
- [ ] Generate correct iteration for for loop
- [ ] Generate correct infinite loop
- [ ] Generate correct break jumps
- [ ] Generate correct continue jumps
- [ ] Handle nested loops correctly

### VM Tests (Execution)
- [ ] Execute if statement (both branches)
- [ ] Execute while loop (multiple iterations)
- [ ] Execute for loop over range
- [ ] Execute infinite loop with break
- [ ] Execute break in while loop
- [ ] Execute continue in while loop
- [ ] Execute nested loops correctly

### Integration Tests (End-to-End)
- [ ] Compile and run factorial program (while loop)
- [ ] Compile and run sum program (for loop)
- [ ] Compile and run search program (early break)
- [ ] Compile and run FizzBuzz (if statements in loop)

## 📚 Documentation Updates Required

### Update Existing Docs
- [ ] README.md: Add control flow to feature list
- [ ] TESTING.md: Add new test categories
- [ ] VM_README.md: Document any new bytecode instructions

### Create New Docs
- [ ] PHASE2_SUMMARY.md: Complete phase summary
- [ ] examples/CONTROL_FLOW.md: Guide to control flow features

### Add Example Programs
- [ ] `test/while_factorial.uni`: Factorial using while
- [ ] `test/for_sum.uni`: Sum using for loop
- [ ] `test/nested_loops.uni`: Nested loop example
- [ ] `test/fizzbuzz_complete.uni`: Complete FizzBuzz

## 🎓 Example Programs

### Factorial with While Loop
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

### Sum with For Loop
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

### FizzBuzz
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

## 🏆 Success Criteria

- [ ] All new parser tests pass (minimum 15 tests)
- [ ] All new generator tests pass (minimum 12 tests)
- [ ] All new VM tests pass (minimum 10 tests)
- [ ] All integration tests pass (minimum 4 tests)
- [ ] Can compile and run factorial program
- [ ] Can compile and run FizzBuzz program
- [ ] No regressions in Phase 1 tests
- [ ] Code coverage ≥ 85%
- [ ] Documentation updated
- [ ] Example programs run correctly

## 💡 Implementation Notes

### Implementation Order
Follow tasks in sequence. Grammar first, then AST, then generator, then VM updates.

### Testing Strategy
- Write tests immediately after implementing each feature
- Use table-driven tests for multiple cases
- Test edge cases (empty loops, single iteration, etc.)

### Common Pitfalls
1. Forgetting to track loop labels for break/continue
2. Not handling empty else blocks correctly
3. Off-by-one errors in range iteration
4. Not testing nested loop scenarios

### Debugging Tips
- Use `make test` frequently to catch regressions
- Add debug output to bytecode generator to see generated code
- Use integration tests to verify end-to-end behavior

---

**Labels**: `phase-2`, `enhancement`, `parser`, `vm`, `control-flow`  
**Milestone**: Phase 2 - Control Flow  
**Assignee**: TBD
