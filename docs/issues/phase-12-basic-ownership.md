# Phase 12: Basic Ownership and Move Semantics

**Status**: Not Started  
**Duration**: 3-4 weeks  
**Priority**: CRITICAL  
**Dependencies**: Phase 11 Complete (Modules and Imports)

## 🎯 Goals

Implement the foundation of Unified's memory safety system: ownership tracking, move semantics, and basic borrow checking. This phase establishes the core mechanisms that prevent memory safety bugs without garbage collection.

## 📋 Prerequisites

- [x] Phase 11 complete (modules and visibility working)
- [ ] Understanding of Rust's ownership model
- [ ] Understanding of linear types and affine types
- [ ] Knowledge of lifetime analysis basics
- [ ] Familiarity with static analysis techniques

## ✨ Language Features to Add

### 1. Ownership Tracking
- Every value has a single owner
- Ownership transfer on assignment (move semantics)
- Automatic cleanup when owner goes out of scope
- Track ownership in symbol table

### 2. Move Semantics
- Move-by-default for non-Copy types
- Invalidate moved-from variables
- Detect use-after-move errors at compile time
- Move in function calls and returns

### 3. Copy Types
- Copy trait for simple types (Int, Float, Bool)
- Copy semantics for Copy types
- Explicit Copy vs Move distinction
- Derive Copy for user-defined types

### 4. Basic Borrowing
- Immutable references (`&T`)
- Mutable references (`&mut T`)
- Reference creation and dereferencing
- Single mutable or multiple immutable rule

### 5. Ownership Analysis
- Static analysis pass after parsing
- Ownership transfer tracking
- Use-after-move detection
- Borrow checker (basic version)
- Lifetime tracking (simplified)

## 📝 Implementation Tasks

### Task 12.1: Add Ownership Grammar (2-3 hours)
- [ ] Open `grammar/UnifiedParser.g4`
- [ ] Add reference type syntax (`&T`, `&mut T`)
- [ ] Add borrow operators (`&`, `&mut`)
- [ ] Add dereference operator (`*`)
- [ ] Update type grammar for references
- [ ] Regenerate parser: `make parser`
- [ ] Test parsing ownership syntax

**Grammar to add:**
```antlr
type
    : AMPERSAND MUT? type  // &T or &mut T
    | basicType
    | IDENTIFIER
    ;

borrowExpression
    : AMPERSAND MUT? expression  // &x or &mut x
    ;

dereferenceExpression
    : STAR expression  // *x
    ;

primaryExpression
    : borrowExpression
    | dereferenceExpression
    | /* existing rules */
    ;
```

### Task 12.2: Implement Ownership AST Nodes (2-3 hours)
- [ ] Open `internal/ast/ast.go`
- [ ] Add ReferenceType struct
- [ ] Add BorrowExpression struct
- [ ] Add DereferenceExpression struct
- [ ] Add ownership metadata to types
- [ ] Update visitor pattern
- [ ] Write unit tests

**AST Structures:**
```go
type ReferenceType struct {
    Mutable    bool
    Referent   Type
}

type BorrowExpression struct {
    Mutable    bool
    Expression Expression
}

type DereferenceExpression struct {
    Expression Expression
}

type TypeInfo struct {
    IsCopy     bool
    IsMove     bool
    // ... existing fields
}
```

### Task 12.3: Implement Ownership Analyzer (8-10 hours)
- [ ] Create `internal/ownership/analyzer.go`
- [ ] Implement ownership tracking data structures
- [ ] Track variable ownership state
- [ ] Detect ownership transfers (moves)
- [ ] Detect use-after-move errors
- [ ] Track variable initialization
- [ ] Handle conditional moves
- [ ] Write comprehensive tests

**Key Components:**
```go
type OwnershipAnalyzer struct {
    symbolTable  *symbols.SymbolTable
    ownershipMap map[string]OwnershipState
}

type OwnershipState int
const (
    StateUninitialized OwnershipState = iota
    StateOwned
    StateMoved
    StateBorrowed
)

func (a *OwnershipAnalyzer) Analyze(ast *AST) []OwnershipError
func (a *OwnershipAnalyzer) CheckMove(variable string) error
func (a *OwnershipAnalyzer) TrackAssignment(lhs, rhs string) error
```

### Task 12.4: Implement Borrow Checker (10-12 hours)
- [ ] Create `internal/ownership/borrow_checker.go`
- [ ] Track active borrows
- [ ] Enforce single mutable or multiple immutable rule
- [ ] Detect conflicting borrows
- [ ] Track borrow lifetimes (simplified)
- [ ] Handle nested scopes
- [ ] Generate clear error messages
- [ ] Write comprehensive tests

**Borrow Checker:**
```go
type BorrowChecker struct {
    activeBorrows map[string][]Borrow
}

type Borrow struct {
    Variable string
    Mutable  bool
    Location SourceLocation
    Scope    *Scope
}

func (b *BorrowChecker) CheckBorrow(variable string, mutable bool) error
func (b *BorrowChecker) RegisterBorrow(variable string, mutable bool, location SourceLocation)
func (b *BorrowChecker) EndScope(scope *Scope)
```

### Task 12.5: Implement Copy Trait System (4-5 hours)
- [ ] Create `internal/types/traits.go`
- [ ] Define Copy trait
- [ ] Mark built-in types as Copy (Int, Float, Bool)
- [ ] Implement Copy derivation for structs
- [ ] Update type checker to use Copy information
- [ ] Handle Copy vs Move in assignments
- [ ] Write tests for Copy semantics

**Copy Trait:**
```go
type Trait interface {
    Name() string
}

type CopyTrait struct{}

func (t *Type) IsCopy() bool
func (t *Type) DeriveCopy() bool  // Can this type derive Copy?
func MarkAsCopy(t *Type)
```

### Task 12.6: Update Symbol Table for Ownership (3-4 hours)
- [ ] Update `internal/symbols/table.go`
- [ ] Add ownership state to symbols
- [ ] Track moved variables
- [ ] Track borrowed variables
- [ ] Add scope-aware borrow tracking
- [ ] Update lookup to check ownership state
- [ ] Write tests

### Task 12.7: Update Type Checker (4-5 hours)
- [ ] Update `internal/semantic/type_checker.go`
- [ ] Integrate ownership analyzer
- [ ] Check ownership before use
- [ ] Validate borrow operations
- [ ] Ensure reference types are used correctly
- [ ] Add ownership-related type errors
- [ ] Write tests

### Task 12.8: Update Bytecode Generator (5-6 hours)
- [ ] Update `internal/bytecode/generator.go`
- [ ] Generate code for borrow operations
- [ ] Generate code for dereference operations
- [ ] Track ownership in generated code
- [ ] Add runtime ownership checks (debug mode)
- [ ] Handle move semantics in codegen
- [ ] Write tests

### Task 12.9: Update VM for References (4-5 hours)
- [ ] Update `internal/vm/vm.go`
- [ ] Implement reference values
- [ ] Implement borrow operations
- [ ] Implement dereference operations
- [ ] Add runtime borrow checking (optional, for debugging)
- [ ] Handle Copy vs Move at runtime
- [ ] Write tests

### Task 12.10: Write Comprehensive Tests (8-10 hours)
- [ ] Parser tests (8+ tests)
- [ ] Ownership analyzer tests (15+ tests)
- [ ] Borrow checker tests (12+ tests)
- [ ] Copy trait tests (6+ tests)
- [ ] Type checker tests (10+ tests)
- [ ] Bytecode generation tests (8+ tests)
- [ ] VM execution tests (8+ tests)
- [ ] Integration tests (10+ tests)

### Task 12.11: Update Documentation (4-5 hours)
- [ ] Create PHASE12_SUMMARY.md
- [ ] Create OWNERSHIP_GUIDE.md
- [ ] Document ownership rules
- [ ] Document borrowing rules
- [ ] Document Copy trait
- [ ] Add ownership examples
- [ ] Update README.md

## ✅ Test Requirements

**Minimum Test Coverage**: 95% for ownership-critical code

### Parser Tests (Ownership Syntax)
- [ ] Parse immutable reference type
- [ ] Parse mutable reference type
- [ ] Parse borrow expression (immutable)
- [ ] Parse borrow expression (mutable)
- [ ] Parse dereference expression
- [ ] Parse nested references
- [ ] Parse function with reference parameters
- [ ] Parse function returning reference

### Ownership Analyzer Tests
- [ ] Detect use-after-move (simple)
- [ ] Detect use-after-move (in function call)
- [ ] Allow use of Copy types after move
- [ ] Track ownership through assignments
- [ ] Track ownership through function calls
- [ ] Track ownership in conditional branches
- [ ] Track ownership in loops
- [ ] Detect uninitialized variable use
- [ ] Track partial moves (struct fields)
- [ ] Handle early returns correctly
- [ ] Handle multiple moves in sequence
- [ ] Detect move in loop body
- [ ] Track ownership in nested scopes
- [ ] Handle ownership in closures (if implemented)
- [ ] Provide clear error messages

### Borrow Checker Tests
- [ ] Allow multiple immutable borrows
- [ ] Reject multiple mutable borrows
- [ ] Reject mutable and immutable borrows together
- [ ] Allow sequential mutable borrows
- [ ] Track borrow lifetimes in scopes
- [ ] Detect borrow outliving value
- [ ] Allow borrowing Copy types
- [ ] Detect use of borrowed value after owner moved
- [ ] Handle nested borrows
- [ ] Handle borrows in function calls
- [ ] Handle borrows in return values
- [ ] Provide clear conflict messages

### Copy Trait Tests
- [ ] Verify Int is Copy
- [ ] Verify Float is Copy
- [ ] Verify Bool is Copy
- [ ] Verify structs are not Copy by default
- [ ] Allow Copy for simple structs
- [ ] Prevent Copy for structs with references
- [ ] Test derived Copy implementation

### Type Checker Tests
- [ ] Type check borrow expressions
- [ ] Type check dereference expressions
- [ ] Ensure reference types match
- [ ] Detect type mismatches in borrows
- [ ] Validate function signatures with references
- [ ] Check reference parameter passing
- [ ] Check reference return types
- [ ] Integrate with ownership checks
- [ ] Generate helpful error messages
- [ ] Handle reference type conversions

### Integration Tests (End-to-End)
- [ ] Compile and run program with moves
- [ ] Compile and run program with borrows
- [ ] Reject program with use-after-move
- [ ] Reject program with invalid borrows
- [ ] Compile program with Copy types
- [ ] Compile program with references
- [ ] Compile program with mutable borrows
- [ ] Test ownership across function calls
- [ ] Test ownership with structs
- [ ] Test comprehensive ownership scenario

## 📚 Documentation Updates Required

### Update Existing Docs
- [ ] README.md: Add ownership to feature list
- [ ] TESTING.md: Document ownership testing
- [ ] spec/UnifiedSpecification.md: Verify ownership spec

### Create New Docs
- [ ] PHASE12_SUMMARY.md: Complete phase summary
- [ ] docs/OWNERSHIP_GUIDE.md: Comprehensive ownership guide
- [ ] docs/BORROWING_GUIDE.md: Borrowing rules and examples
- [ ] docs/MEMORY_SAFETY.md: Memory safety guarantees

### Add Example Programs
- [ ] `test/phase12/move_semantics.uni`: Basic moves
- [ ] `test/phase12/borrow_immutable.uni`: Immutable borrows
- [ ] `test/phase12/borrow_mutable.uni`: Mutable borrows
- [ ] `test/phase12/copy_types.uni`: Copy semantics
- [ ] `test/phase12/ownership_transfer.uni`: Ownership transfer
- [ ] `test/phase12/use_after_move_error.uni`: Error example
- [ ] `test/phase12/borrow_conflict_error.uni`: Error example
- [ ] `test/phase12/linked_list.uni`: Real-world example

## 🎓 Example Programs

### Move Semantics
```unified
fn main() -> Int {
    let s1 = String::from("hello")
    let s2 = s1  // s1 moved to s2
    // let x = s1  // ERROR: use after move
    
    let s3 = s2  // s2 moved to s3
    print(s3)
    return 0
}
```

### Immutable Borrows
```unified
fn calculate_length(s: &String) -> Int {
    return s.len()
}

fn main() -> Int {
    let s1 = String::from("hello")
    let len = calculate_length(&s1)
    print(s1)  // Still valid - s1 wasn't moved
    return len
}
```

### Mutable Borrows
```unified
fn append_world(s: &mut String) {
    s.push_str(" world")
}

fn main() -> Int {
    let mut s = String::from("hello")
    append_world(&mut s)
    print(s)  // Prints "hello world"
    return 0
}
```

### Copy Types
```unified
fn main() -> Int {
    let x = 42
    let y = x  // Copy, not move (Int is Copy)
    let z = x + y  // Both x and y are still valid
    return z
}
```

### Borrowing Rules
```unified
fn main() -> Int {
    let mut s = String::from("hello")
    
    // Multiple immutable borrows OK
    let r1 = &s
    let r2 = &s
    print(r1)
    print(r2)
    
    // Mutable borrow OK after immutable borrows out of scope
    let r3 = &mut s
    r3.push_str(" world")
    
    return 0
}
```

### Use-After-Move Error
```unified
fn take_ownership(s: String) {
    print(s)
}

fn main() -> Int {
    let s = String::from("hello")
    take_ownership(s)  // s moved into function
    // print(s)  // ERROR: use after move
    return 0
}
```

### Borrow Conflict Error
```unified
fn main() -> Int {
    let mut s = String::from("hello")
    
    let r1 = &s      // Immutable borrow
    // let r2 = &mut s  // ERROR: cannot borrow as mutable
                        // because it's borrowed as immutable
    
    print(r1)
    return 0
}
```

### Realistic Example: Linked List Node
```unified
struct Node {
    value: Int
    next: Option<Box<Node>>
}

impl Node {
    fn new(value: Int) -> Node {
        return Node { value: value, next: None }
    }
    
    fn set_next(&mut self, next: Node) {
        self.next = Some(new Box(next))
    }
    
    fn get_value(&self) -> Int {
        return self.value
    }
}

fn main() -> Int {
    let mut head = new Node(1)
    let second = new Node(2)
    head.set_next(second)  // second moved
    
    return head.get_value()
}
```

## 🏆 Success Criteria

- [ ] All parser tests pass (minimum 8 tests)
- [ ] All ownership analyzer tests pass (minimum 15 tests)
- [ ] All borrow checker tests pass (minimum 12 tests)
- [ ] All Copy trait tests pass (minimum 6 tests)
- [ ] All type checker tests pass (minimum 10 tests)
- [ ] All bytecode tests pass (minimum 8 tests)
- [ ] All VM tests pass (minimum 8 tests)
- [ ] All integration tests pass (minimum 10 tests)
- [ ] Use-after-move is detected at compile time
- [ ] Borrow conflicts are detected at compile time
- [ ] Copy types work correctly
- [ ] No regressions in previous phases
- [ ] Code coverage ≥ 95% for ownership code
- [ ] Documentation is comprehensive
- [ ] Example programs demonstrate all features

## 💡 Implementation Notes

### Implementation Order
1. Grammar and parser (syntax)
2. AST nodes (representation)
3. Ownership analyzer (move tracking)
4. Copy trait system (Copy vs Move)
5. Borrow checker (borrow validation)
6. Type checker integration
7. Bytecode generator updates
8. VM updates (runtime support)
9. Tests and documentation

### Testing Strategy
- Start with simple move cases, then complex scenarios
- Test each borrow rule independently
- Use intentionally failing programs to test error detection
- Create comprehensive integration tests
- Test edge cases: loops, conditionals, early returns
- Verify error messages are clear and helpful

### Common Pitfalls
1. **Forgetting Copy Types**: Remember Int, Float, Bool are Copy
2. **Nested Scopes**: Borrow lifetimes end at scope exit
3. **Partial Moves**: Moving struct fields is complex
4. **Conditional Moves**: Track moves in all branches
5. **Loop Moves**: Can't move in loop that runs multiple times
6. **Error Messages**: Make them actionable and clear

### Debugging Tips
- Print ownership state at each program point
- Visualize ownership transfer with diagrams
- Use small test cases to isolate issues
- Check borrow checker logic with hand-written examples
- Add verbose error messages during development

### Performance Considerations
- Ownership analysis is compile-time, no runtime cost
- Cache analysis results when possible
- Optimize common patterns (Copy types, simple moves)
- Consider using worklist algorithms for efficiency

### Simplifications (For This Phase)
- Simplified lifetime tracking (full lifetimes in later phase)
- No lifetime parameters yet
- No advanced borrow patterns
- No interior mutability
- No unsafe code support

### Future Extensions (Not This Phase)
- Full lifetime analysis
- Lifetime parameters
- Interior mutability (RefCell, etc.)
- Unsafe code blocks
- Non-lexical lifetimes (NLL)
- Async/await ownership integration

---

**Labels**: `phase-12`, `critical`, `ownership`, `borrow-checker`, `memory-safety`, `compiler`  
**Milestone**: Phase 12 - Basic Ownership  
**Assignee**: TBD
