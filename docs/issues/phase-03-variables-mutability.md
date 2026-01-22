# Phase 3: Variables, Mutability, and Assignment

**Status**: Not Started  
**Duration**: 2 weeks  
**Priority**: HIGH  
**Dependencies**: Phase 2 Complete (Control Flow)

## 🎯 Goals

Implement proper variable declaration, mutability tracking, and assignment operations. This phase enhances the variable system beyond basic `let` statements and establishes the foundation for ownership concepts in future phases.

## 📋 Prerequisites

- [x] Phase 2 complete (control flow working)
- [ ] Understanding of ownership concepts (for future phases)
- [ ] Symbol table design concepts

## ✨ Language Features to Add

### 1. Mutable Variables
- `let mut` declarations
- Mutability checking at compile time
- Error on assignment to immutable variables
- Clear error messages for mutability violations

### 2. Assignment Statements
- Simple assignment: `x = value`
- Compound assignment: `+=`, `-=`, `*=`, `/=`, `%=`
- Validation of assignment targets
- Type checking for assignments

### 3. Variable Shadowing
- Allow redeclaration in nested scopes
- Track scope correctly
- Preserve type safety across shadows

### 4. Type Inference
- Basic type inference for let statements
- Type checking for assignments
- Infer types from literal values
- Infer types from expressions

## 📝 Implementation Tasks

### Task 3.1: Add Mutability to Grammar (1-2 hours)
- [ ] Open `grammar/UnifiedParser.g4`
- [ ] Update let statement grammar to support `mut` keyword
- [ ] Add assignment statement grammar rule
- [ ] Add assignment operators (=, +=, -=, *=, /=, %=)
- [ ] Regenerate parser: `make parser`
- [ ] Verify parser generation succeeds
- [ ] Test parsing mutable declarations

**Grammar to add:**
```antlr
letStatement
    : LET MUT? IDENTIFIER (COLON type)? ASSIGN expression SEMICOLON?
    ;

assignmentStatement
    : IDENTIFIER assignmentOp expression SEMICOLON?
    ;

assignmentOp
    : ASSIGN | PLUS_ASSIGN | MINUS_ASSIGN | STAR_ASSIGN 
    | SLASH_ASSIGN | PERCENT_ASSIGN
    ;
```

### Task 3.2: Update AST for Mutability (2 hours)
- [ ] Open `internal/ast/ast.go`
- [ ] Add mutable flag to LetStatement struct
- [ ] Create AssignmentStatement struct
- [ ] Add AssignOp enum for assignment operators
- [ ] Update visitor pattern to handle new nodes
- [ ] Add unit tests for AST construction

**AST structures:**
```go
type LetStatement struct {
    Name      string
    Type      Type // may be nil for inference
    Value     Expression
    Mutable   bool
}

type AssignmentStatement struct {
    Target   string
    Operator AssignOp // =, +=, -=, etc.
    Value    Expression
}

type AssignOp int
const (
    AssignNormal AssignOp = iota
    AssignAdd
    AssignSub
    AssignMul
    AssignDiv
    AssignMod
)
```

### Task 3.3: Implement Symbol Table (4-6 hours)
- [ ] Create `internal/semantic/symboltable.go`
- [ ] Define Symbol struct with name, type, mutability
- [ ] Define SymbolTable with scope management
- [ ] Implement EnterScope() method
- [ ] Implement ExitScope() method
- [ ] Implement Define(name, type, mutable) method
- [ ] Implement Lookup(name) method
- [ ] Add symbol table to bytecode generator
- [ ] Track variable mutability across scopes
- [ ] Add unit tests for symbol table

**Symbol table structure:**
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

### Task 3.4: Add Semantic Checking (3-4 hours)
- [ ] Create `internal/semantic/checker.go`
- [ ] Implement check for assignment to immutable variables
- [ ] Implement check for undefined variable errors
- [ ] Add error reporting with line numbers
- [ ] Create comprehensive error messages
- [ ] Add context to error messages (show variable name)
- [ ] Integrate checker into compilation pipeline
- [ ] Add tests for all error cases

### Task 3.5: Implement Type Inference (4-6 hours)
- [ ] Create `internal/semantic/types.go`
- [ ] Implement type inference for literal values
  - Int literals → Int type
  - Float literals → Float type
  - Bool literals → Bool type
  - String literals → String type
- [ ] Implement type inference for expressions
  - Binary ops → result type
  - Function calls → return type
- [ ] Store inferred types in symbol table
- [ ] Validate type compatibility in assignments
- [ ] Add type coercion rules where appropriate
- [ ] Add comprehensive type inference tests

### Task 3.6: Update VM for Assignment (2-3 hours)
- [ ] Review existing STORE_LOCAL implementation
- [ ] Implement STORE_LOCAL for variable reassignment
- [ ] Verify mutable variables can be updated
- [ ] Add runtime checks if needed
- [ ] Test that immutable variables cannot be updated
- [ ] Add VM tests for assignment operations

### Task 3.7: Write Tests (4-5 hours)
- [ ] Write parser tests for mut keyword
- [ ] Write parser tests for all assignment operators
- [ ] Write symbol table tests
- [ ] Write semantic checking tests (error cases)
- [ ] Write type inference tests
- [ ] Write integration tests with mutable variables
- [ ] Ensure code coverage ≥ 85%

### Task 3.8: Update Documentation (2 hours)
- [ ] Document mutability in language guide
- [ ] Create PHASE3_SUMMARY.md
- [ ] Add examples to test directory
- [ ] Update error message documentation
- [ ] Update README.md with new features
- [ ] Create migration guide from Phase 2

## ✅ Test Requirements

**Minimum Test Coverage**: 85% for new code

### Parser Tests
- [ ] Parse `let mut x = 5`
- [ ] Parse `let x = 5` (immutable)
- [ ] Parse assignment `x = 10`
- [ ] Parse compound assignment `x += 5`
- [ ] Parse all assignment operators (+=, -=, *=, /=, %=)
- [ ] Parse type annotations with mut
- [ ] Parse shadowed variables
- [ ] Reject invalid syntax

### Symbol Table Tests
- [ ] Define variable in scope
- [ ] Lookup variable in current scope
- [ ] Lookup variable in parent scope
- [ ] Shadow variable in nested scope
- [ ] Scope isolation (can't access child scope vars)
- [ ] Handle multiple scope levels
- [ ] Clear scopes correctly on exit

### Semantic Analysis Tests
- [ ] Error on assignment to immutable variable
- [ ] Error on assignment to undefined variable
- [ ] Error on type mismatch in assignment
- [ ] Allow assignment to mutable variable
- [ ] Allow variable shadowing
- [ ] Provide clear error messages
- [ ] Report correct line numbers for errors

### Type Inference Tests
- [ ] Infer Int from integer literal
- [ ] Infer Float from float literal
- [ ] Infer Bool from boolean literal
- [ ] Infer String from string literal
- [ ] Infer type from arithmetic expression
- [ ] Infer type from function call
- [ ] Detect type mismatch in assignment
- [ ] Handle explicit type annotations

### Integration Tests
- [ ] Counter with mutable variable
- [ ] Compound assignment in loop
- [ ] Variable shadowing in nested scopes
- [ ] Mixed mutable and immutable variables
- [ ] Type inference in complex expressions

## 📚 Documentation Updates Required

### Update Existing Docs
- [ ] README.md: Add mutability and type inference to feature list
- [ ] TESTING.md: Add semantic analysis test categories
- [ ] VM_README.md: Document any new instructions

### Create New Docs
- [ ] PHASE3_SUMMARY.md: Complete phase summary
- [ ] examples/VARIABLES.md: Guide to variables and mutability
- [ ] docs/TYPE_INFERENCE.md: Type inference documentation

### Add Example Programs
- [ ] `test/counter_mut.uni`: Counter with mutable variable
- [ ] `test/compound_assign.uni`: Compound assignment examples
- [ ] `test/shadowing.uni`: Variable shadowing demo
- [ ] `test/type_inference.uni`: Type inference examples

## 🎓 Example Programs

### Counter with Mutable Variable
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

### Type Inference
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

### Variable Shadowing
```unified
fn shadowing_example() -> Int {
    let x = 5
    {
        let x = 10       // Shadows outer x
        let y = x * 2    // Uses inner x (10)
        // y = 20
    }
    // x is still 5 here
    return x
}
```

### Compound Assignment
```unified
fn compound_ops() -> Int {
    let mut total = 100
    total += 50      // 150
    total -= 30      // 120
    total *= 2       // 240
    total /= 4       // 60
    total %= 25      // 10
    return total
}
```

## 🏆 Success Criteria

- [ ] All parser tests pass (minimum 15 tests)
- [ ] All symbol table tests pass (minimum 12 tests)
- [ ] All semantic analysis tests pass (minimum 15 tests)
- [ ] All type inference tests pass (minimum 12 tests)
- [ ] All integration tests pass (minimum 8 tests)
- [ ] Mutability tracking works correctly
- [ ] Assignments to immutable variables are caught at compile time
- [ ] Type inference works for basic cases
- [ ] Symbol table correctly handles scoping
- [ ] No regressions in Phase 1 and Phase 2 tests
- [ ] Code coverage ≥ 85%
- [ ] Documentation complete
- [ ] Clear error messages for all error cases

## 💡 Implementation Notes

### Implementation Order
Follow tasks in sequence. Grammar first, then AST, then symbol table (critical!), then semantic checking, then type inference, then VM updates.

### Critical Design Decision
**Symbol table is the foundation**: All future type checking, ownership analysis, and borrow checking will build on the symbol table. Design it carefully to support future extensions.

### Testing Strategy
- Write tests immediately after implementing each feature
- Focus on error cases as much as success cases
- Good error messages are critical for usability
- Test edge cases (empty scopes, deeply nested scopes, shadowing chains)

### Common Pitfalls
1. Forgetting to check mutability before allowing assignment
2. Not properly tracking scope entry/exit
3. Poor error messages that don't help users
4. Type inference being too permissive or too strict
5. Not handling shadowing correctly
6. Forgetting to test compound assignment operators

### Debugging Tips
- Use `make test` frequently to catch regressions
- Add debug output to symbol table to see scope changes
- Print inferred types during development
- Test error messages manually to ensure quality
- Use integration tests to verify end-to-end behavior

### Future Considerations
This phase lays groundwork for:
- Ownership and borrowing (Phase 11)
- Lifetime analysis
- Advanced type checking
- Generic type inference

---

**Labels**: `phase-3`, `enhancement`, `semantic-analysis`, `type-system`  
**Milestone**: Phase 3 - Variables and Mutability  
**Assignee**: TBD
