# Unified Programming Language - Implementation Roadmap 2026

**Created**: February 2, 2026
**Status**: Active Development - Phase 1 Complete (Core Features)
**Current Priority**: Bug Fixes & Priority 1 Features
**Target**: 70-80% test pass rate by Q1 2026 end

---

## Overview

This roadmap outlines the implementation plan for bringing the Unified Programming Language from its current state (24% tests passing) to a feature-complete MVP (70-80% tests passing) and beyond.

**Current Baseline**:
- ~30 of 123 tests passing (24%)
- Core features working (functions, variables, control flow, structs, enums)
- WASM backend complete and functional
- Critical bugs blocking ~30 tests
- Advanced features (generics, pattern matching) not yet implemented

**6-Month Goal**: Feature-complete language with 90%+ tests passing

---

## 🚨 Priority 0: Infrastructure (URGENT) - Week 1

### Goal: Establish accurate testing baseline

These issues are **blocking visibility** into actual progress and must be fixed before any feature work begins.

### Tasks

#### 1. Fix Test Infrastructure (15 minutes)

**Problem**: `test_all_uni_files.sh` uses `timeout` command which doesn't exist on macOS

**Solution Options**:

```bash
# Option A: Install coreutils (provides gtimeout)
brew install coreutils

# Then update test_all_uni_files.sh:
# Change: timeout 5 ./bin/unified ...
# To: gtimeout 5 ./bin/unified ...
```

```bash
# Option B: Rewrite without timeout command
# Use Go-based test runner or remove timeout entirely
```

```bash
# Option C: Use Go test harness exclusively
cd src/unified-compiler
go test ./... -v
```

**Effort**: 15 minutes
**Assignee**: Infrastructure
**Blocker**: YES - blocks all accurate test measurement

#### 2. Run Comprehensive Test Audit (30 minutes)

**Tasks**:
1. Fix test infrastructure (above)
2. Run full test suite: `./test_all_uni_files.sh`
3. Run Go unit tests: `cd src/unified-compiler && go test ./... -v`
4. Document exact pass/fail counts
5. Categorize failures by root cause
6. Create baseline report

**Output**: `TEST_BASELINE_2026-02-02.md` with exact numbers

**Effort**: 30 minutes
**Dependencies**: Task 1 (test infrastructure)
**Blocker**: YES - needed for accurate planning

#### 3. Update Documentation (30 minutes)

**Files to update**:
- [TODO.md](TODO.md) - Update test counts and priorities
- [README.md](README.md) - Update badges and status
- [docs/PROJECT_STATUS.md](docs/PROJECT_STATUS.md) - Archive and reference new status
- [CLAUDE.md](CLAUDE.md) - Update current status section

**Effort**: 30 minutes
**Dependencies**: Task 2 (test audit)
**Blocker**: NO - but important for team coordination

---

## 🔥 Priority 1: Critical Bug Fixes - Weeks 1-3

### Goal: Fix blocking bugs preventing ~30 tests from passing

These are bugs in **already-implemented features** that just need debugging.

### 1.1 Fix For Loop Implementation (2-4 hours)

**Problem**: For loops compile and run but produce wrong values

**Evidence**:
```
TestIntegrationForSum:
  Expected: 55 (sum 0..10)
  Got: 66 (20% too high)

TestIntegrationNestedLoops:
  Expected: 18
  Got: 40 (122% too high)
```

**Root Cause**: Likely off-by-one error or accumulator bug in loop codegen

**Investigation Steps**:
1. Read [src/unified-compiler/internal/wasm/codegen.go:visitForStatement](src/unified-compiler/internal/wasm/codegen.go)
2. Add debug logging to track loop counter values
3. Compare against working while loop implementation
4. Check range operator handling (inclusive vs exclusive)
5. Verify loop variable initialization and increment

**Files to modify**:
- `src/unified-compiler/internal/wasm/codegen.go` - Loop codegen
- `src/unified-compiler/internal/ast/visitor.go` - Loop AST building (if needed)

**Testing**:
```bash
cd src/unified-compiler
go test -v -run TestIntegrationForSum
go test -v -run TestIntegrationNestedLoops
```

**Effort**: 2-4 hours
**Impact**: 2-4 tests
**Difficulty**: Medium (logic bug in existing code)

### 1.2 Resolve Type System Issues (1-3 days)

**Problem**: Persistent i32 vs i64 type mismatches throughout codebase

**Affected Areas**:
- Nested struct field access
- String operations
- Array indexing
- Generic function calls

**Example Errors**:
```
type mismatch: expected i32, but was i64
WASM validation error: global.get index mismatch
```

**Root Cause Analysis Needed**:
1. Are we using i32 for offsets but i64 for values?
2. Is wazero expecting specific integer types?
3. Are we inconsistent in type inference?
4. Do we need explicit type conversions?

**Investigation Steps**:
1. **Audit all integer type decisions**:
   - Search codebase for `i32` and `i64` usage
   - Document where each is used and why
   - Identify inconsistencies

2. **Check WASM memory model assumptions**:
   - What does wazero expect for memory addresses? (i32)
   - What does wazero expect for stack values? (i32 or i64?)
   - What does WASM spec require?

3. **Review struct field access**:
   - How are offsets calculated? (i32)
   - How are values loaded/stored? (i32/i64?)
   - Are we mixing the two?

4. **Fix systematically**:
   - Define clear rules: "Memory addresses are i32, values are i64" (or vice versa)
   - Update codegen consistently
   - Add type conversion functions where needed
   - Update semantic analyzer to enforce rules

**Files likely needing changes**:
- `src/unified-compiler/internal/wasm/codegen.go` - Type generation
- `src/unified-compiler/internal/wasm/generator.go` - WASM module structure
- `src/unified-compiler/internal/semantic/typechecker.go` - Type inference
- `src/unified-compiler/internal/ast/types.go` - Type definitions

**Testing**:
```bash
# After fixes, these should pass:
go test -v -run TestIntegrationStructNested
go test -v -run TestStringLength
go test -v -run TestArrayIndex
```

**Effort**: 1-3 days (depends on scope of issue)
**Impact**: 15-20 tests
**Difficulty**: High (systemic issue, requires careful analysis)
**Blocker**: YES - blocks multiple feature areas

### 1.3 Fix Array Support (1-2 days)

**Problem**: Array operations mostly broken

**Current Status**:
- ✅ Array literals parse correctly
- ❌ Heap allocation not working
- ❌ Indexing broken (type mismatch)
- ❌ Iteration not implemented

**Implementation Steps**:

**Step 1: Fix heap allocation for arrays**
- Review `allocateOnHeap` function in codegen
- Ensure array size calculation correct: `element_size * element_count`
- Add bounds metadata (length) before array data
- Test with simple array creation

**Step 2: Fix array indexing**
- Calculate offset: `base_address + (index * element_size)`
- Add bounds checking (compare index < length)
- Generate WASM load instruction for element type
- Handle type conversions (i32 index vs i64 values)

**Step 3: Implement array iteration**
- For loop over array: `for i in 0..array.length`
- Array.length property access
- Element access in loop body

**Step 4: Add array operations**
- Array assignment: `arr[i] = value`
- Array methods: `.length`, `.push()`, `.pop()` (if in spec)

**Test Files**:
- `test/arrays/array_literal.uni`
- `test/arrays/array_index.uni`
- `test/arrays/array_iteration.uni`

**Effort**: 1-2 days
**Impact**: 11 tests
**Difficulty**: Medium-High
**Dependencies**: Type system fixes (1.2)

### 1.4 Fix String Operations (1-2 days)

**Problem**: String literals work, but operations fail

**Current Status**:
- ✅ String literals parse and store in data section
- ✅ String deduplication works
- ❌ Type mismatches in operations
- ❌ Missing runtime functions

**Implementation Steps**:

**Step 1: Define string representation**
- Decide format: pointer to (length, data) tuple? Or null-terminated?
- Document in ARCHITECTURE.md
- Ensure consistent with current literal implementation

**Step 2: Implement string operations**
- `string.length` - Return i32 length
- `string[index]` - Return i32 character code
- `string + string` - Concatenation (allocate new string)
- `string.substring(start, end)` - Slice operation
- `string == string` - Equality comparison

**Step 3: Add runtime functions**
```go
// In runtime.go
func stringLength(ptr int32) int32
func stringConcat(ptr1, ptr2 int32) int32
func stringSubstring(ptr, start, end int32) int32
func stringEquals(ptr1, ptr2 int32) int32
```

**Step 4: Generate WASM calls to runtime**
- String operations generate `call $runtime.stringLength` etc.
- Import runtime functions in WASM module
- Handle return values correctly

**Test Files**:
- `test/strings/string_length.uni`
- `test/strings/string_concat.uni`
- `test/strings/string_index.uni`
- `test/strings/string_compare.uni`

**Effort**: 1-2 days
**Impact**: 11 tests
**Difficulty**: Medium
**Dependencies**: Type system fixes (1.2)

---

## 📦 Priority 2: High-Value Features - Weeks 4-6

### Goal: Implement features with parser support but missing codegen

These features have **parser work already complete**, just need code generation.

### 2.1 Complete Try Operator (6-12 hours)

**Status**:
- ✅ Parser support (`expr?` syntax)
- ✅ AST nodes defined
- ❌ WASM codegen not implemented

**Implementation Steps**:

**Step 1: Design error handling convention**
- How are errors represented? `Result<T, E>` enum?
- How does `?` propagate errors? Early return?
- What does WASM code look like?

**Step 2: Implement codegen for try operator**
```go
// In codegen.go
func (g *Generator) visitTryExpr(expr *ast.TryExpr) error {
    // 1. Generate code for inner expression
    // 2. Check if result is Error variant
    // 3. If Error: return immediately
    // 4. If Success: unwrap value and continue
}
```

**Step 3: Enum variant checking**
- Use enum tag to check Success vs Error
- Branch on tag value
- Extract inner value from variant

**Step 4: Early return mechanism**
- Generate WASM `return` instruction for Error case
- Ensure proper stack cleanup
- Test with nested try operators

**Test Files**:
- `test/errors/try_operator.uni`
- `test/errors/try_nested.uni`
- `test/errors/try_propagation.uni`

**Effort**: 6-12 hours
**Impact**: 10 tests
**Difficulty**: Medium
**Value**: HIGH (essential for ergonomic error handling)

### 2.2 Improve Struct Operations (6-8 hours)

**Goal**: Complete remaining struct features

**Tasks**:

1. **Fix nested field access** (covered by type system fixes)
2. **Implement struct methods** (if in spec)
3. **Add struct initialization syntax** (if not already working)
4. **Test complex struct patterns**

**Effort**: 6-8 hours (assuming type system fixed)
**Impact**: 3-4 additional tests
**Difficulty**: Low-Medium

---

## 🎯 Priority 3: Advanced Features - Weeks 7-12

### Goal: Implement sophisticated type system features

### 3.1 Generic Monomorphization (2-3 weeks)

**Problem**: Generic syntax parses but no type instantiation

**Current Status**:
- ✅ Generic function syntax parses: `fn foo<T>(x: T) -> T`
- ✅ Generic type parameters recognized
- ❌ No monomorphization pass
- ❌ No type specialization

**Background**: Monomorphization is generating specialized versions of generic functions for each concrete type used.

**Example**:
```unified
fn identity<T>(x: T) -> T { return x }

let a = identity(42)      // Generates identity_Int
let b = identity(3.14)    // Generates identity_Float
let c = identity("hello") // Generates identity_String
```

**Implementation Plan**:

**Phase 1: Collect generic function usage (1-2 days)**
- Add pass before codegen to scan for generic function calls
- Track concrete type arguments for each call site
- Build map: `generic_function_name -> Set<ConcreteTypes>`

**Phase 2: Generate specialized functions (3-4 days)**
- For each generic function + concrete types:
  - Clone AST of generic function
  - Substitute type parameters with concrete types
  - Rename function: `identity<T>` → `identity_Int`
  - Add specialized function to module

**Phase 3: Rewrite call sites (2-3 days)**
- Replace generic calls with specialized function names
- Update semantic analysis to resolve generic types
- Add type inference for generic arguments (if not explicit)

**Phase 4: Handle generic structs and enums (3-4 days)**
- Similar approach: `List<T>` → `List_Int`, `List_String`
- More complex: need to handle field types, method calls

**Files to modify**:
- `src/unified-compiler/internal/semantic/generics.go` (new file)
- `src/unified-compiler/internal/wasm/generator.go` - Add monomorphization pass
- `src/unified-compiler/internal/ast/visitor.go` - Clone and substitute types

**Test Files**: 15 generic tests in `test/generics/`

**Effort**: 2-3 weeks
**Impact**: 15 tests
**Difficulty**: HIGH (complex type system feature)
**Value**: HIGH (essential for reusable code)

### 3.2 Methods and Self Keyword (1-2 weeks)

**Problem**: Standard library needs method syntax

**Current Status**:
- ❌ No `Self` keyword
- ❌ No method syntax
- ❌ No method dispatch

**Implementation Plan**:

**Phase 1: Add method syntax to parser (2-3 days)**
```unified
struct Point {
    x: Int,
    y: Int
}

impl Point {
    fn new(x: Int, y: Int) -> Self {
        return Point { x: x, y: y }
    }

    fn distance(&self) -> Float {
        return sqrt(self.x * self.x + self.y * self.y)
    }
}
```

- Update grammar with `impl` blocks
- Add `Self` keyword
- Parse method declarations

**Phase 2: Semantic analysis (3-4 days)**
- Resolve `Self` to impl type
- Track methods in struct registry
- Type check method calls
- Handle `self`, `&self`, `&mut self` parameters

**Phase 3: Code generation (3-4 days)**
- Translate methods to regular functions with extra parameter
- `point.distance()` → `Point_distance(point)`
- Implement method dispatch

**Test Files**: 24 stdlib tests need methods

**Effort**: 1-2 weeks
**Impact**: 24 tests
**Difficulty**: HIGH
**Value**: HIGH (enables OOP patterns, stdlib)

### 3.3 Pattern Matching (2-3 weeks)

**Problem**: Match expressions not implemented

**Current Status**:
- ❌ No match expression parsing
- ❌ No pattern syntax
- ❌ No exhaustiveness checking

**Implementation Plan**:

**Phase 1: Parser (3-4 days)**
```unified
match result {
    Result::Ok(value) => {
        println("Success: {value}")
    }
    Result::Error(msg) => {
        println("Error: {msg}")
    }
}
```

- Add match expression grammar
- Parse pattern syntax
- Handle destructuring

**Phase 2: Semantic analysis (4-5 days)**
- Type check patterns against scrutinee
- Ensure exhaustiveness (all cases covered)
- Handle wildcard patterns `_`
- Implement pattern variable binding

**Phase 3: Code generation (4-5 days)**
- Generate dispatch table for enum variants
- Branch on enum tag
- Extract variant data
- Jump to appropriate arm code

**Test Files**: Multiple test categories use pattern matching

**Effort**: 2-3 weeks
**Impact**: 10-15 tests initially, enables many more
**Difficulty**: HIGH
**Value**: CRITICAL (core language feature)

---

## 🚀 Priority 4: Specification Features - Weeks 13-26 (Q2 2026)

### Goal: Implement remaining specification features

### 4.1 Ownership and Borrowing (4-6 weeks)

**Scope**: Implement Rust-like ownership system

**Features**:
- Move semantics
- Borrow checking
- Lifetime analysis
- Mutable/immutable borrows

**Complexity**: VERY HIGH
**Value**: CRITICAL (core safety feature)

### 4.2 Actor-Based Concurrency (4-6 weeks)

**Scope**: Erlang/Elixir-style actors

**Features**:
- Actor definitions
- Message passing
- Mailboxes
- Process spawning
- Supervision trees

**Complexity**: VERY HIGH
**Value**: HIGH (differentiating feature)

### 4.3 Async/Await (3-4 weeks)

**Scope**: Asynchronous programming support

**Features**:
- `async` functions
- `await` keyword
- Futures/Promises
- Event loop integration

**Complexity**: HIGH
**Value**: HIGH (modern concurrency)

### 4.4 Standard Library Expansion (6-8 weeks)

**Scope**: Complete standard library

**Modules**:
- Collections (List, Map, Set, Queue, Stack)
- I/O (File, Network, Streams)
- String utilities
- Math functions
- Time and date
- OS integration (via WASI)
- Error types

**Complexity**: MEDIUM (lots of code, not complex)
**Value**: CRITICAL (needed for real programs)

---

## 🛠️ Priority 5: Tooling & Polish - Weeks 27-40 (Q3 2026)

### Goal: Production-ready tooling

### 5.1 Language Server Protocol (4-6 weeks)

**Features**:
- Syntax highlighting
- Autocomplete
- Go to definition
- Find references
- Inline errors
- Refactoring support

**Value**: HIGH (developer experience)

### 5.2 Package Manager (3-4 weeks)

**Features**:
- Dependency management
- Package registry
- Version resolution
- Build integration

**Value**: HIGH (ecosystem enabler)

### 5.3 Debugger Integration (2-3 weeks)

**Features**:
- Breakpoints
- Step through code
- Inspect variables
- Call stack

**Value**: MEDIUM (nice to have)

### 5.4 Documentation Generator (1-2 weeks)

**Features**:
- Generate docs from comments
- API reference
- Examples
- Searchable website

**Value**: MEDIUM (important for adoption)

---

## 📊 Progress Milestones

### Milestone 1: "Core Language Complete" (Week 3)

**Target Date**: February 23, 2026
**Test Pass Rate**: 50-60% (60-75 tests)

**Criteria**:
- ✅ Test infrastructure working
- ✅ For loops working correctly
- ✅ Type system consistent (no i32/i64 mismatches)
- ✅ Arrays fully functional
- ✅ Strings fully functional
- ✅ Try operator working

**Deliverables**:
- All Priority 0 and Priority 1 tasks complete
- Updated documentation
- Comprehensive test report

### Milestone 2: "Advanced Features Complete" (Week 12)

**Target Date**: April 30, 2026
**Test Pass Rate**: 70-80% (85-100 tests)

**Criteria**:
- ✅ Generic monomorphization working
- ✅ Methods and Self keyword working
- ✅ Pattern matching working
- ✅ All Priority 2 and Priority 3 tasks complete

**Deliverables**:
- Feature-complete core language
- Updated specification
- Tutorial and examples

### Milestone 3: "Specification Complete" (Week 26)

**Target Date**: August 1, 2026
**Test Pass Rate**: 90%+ (110+ tests)

**Criteria**:
- ✅ Ownership and borrowing working
- ✅ Actors and concurrency working
- ✅ Async/await working
- ✅ Standard library 80% complete

**Deliverables**:
- Full specification implementation
- Comprehensive standard library
- Real-world example programs

### Milestone 4: "Production Ready" (Week 40)

**Target Date**: November 1, 2026
**Test Pass Rate**: 95%+ (115+ tests)

**Criteria**:
- ✅ Language Server Protocol working
- ✅ Package manager working
- ✅ Debugger integration working
- ✅ Documentation generator working
- ✅ Standard library complete
- ✅ CI/CD pipeline operational

**Deliverables**:
- Production-ready compiler
- Complete tooling
- Official 1.0 release

---

## 📅 Timeline Summary

| Phase | Weeks | Dates | Goals | Tests |
|-------|-------|-------|-------|-------|
| **Priority 0** | 1 | Feb 2-9 | Fix infrastructure | Baseline |
| **Priority 1** | 2-3 | Feb 9-23 | Fix critical bugs | 50-60% |
| **Priority 2** | 4-6 | Feb 23 - Apr 5 | High-value features | 60-70% |
| **Priority 3** | 7-12 | Apr 5 - Apr 30 | Advanced features | 70-80% |
| **Priority 4** | 13-26 | May 1 - Aug 1 | Spec features | 90%+ |
| **Priority 5** | 27-40 | Aug 1 - Nov 1 | Tooling & polish | 95%+ |

**Total Duration**: 40 weeks (~9 months)
**Target Release**: November 2026

---

## 🎯 Resource Requirements

### Development Time

**Per Week Estimates** (for reference):
- Priority 0: ~2 hours
- Priority 1: ~15-20 hours/week (3 weeks) = 45-60 hours
- Priority 2: ~15-20 hours/week (3 weeks) = 45-60 hours
- Priority 3: ~20-25 hours/week (6 weeks) = 120-150 hours
- Priority 4: ~20-25 hours/week (14 weeks) = 280-350 hours
- Priority 5: ~15-20 hours/week (14 weeks) = 210-280 hours

**Total Estimated Hours**: 700-900 hours

### Skillsets Needed

1. **Compiler Engineer** (Priority 1-4)
   - Go programming
   - ANTLR/parser generators
   - WebAssembly
   - Type systems
   - Code generation

2. **Language Designer** (Priority 3-4)
   - Type theory
   - Language semantics
   - API design
   - Documentation

3. **Tooling Engineer** (Priority 5)
   - LSP protocol
   - IDE integration
   - Build systems
   - Package management

4. **QA/Testing** (All priorities)
   - Test design
   - Integration testing
   - Performance testing
   - Documentation testing

### External Dependencies

- **ANTLR 4.13.2**: Parser generation
- **wazero**: WASM runtime
- **Go 1.21+**: Compiler implementation
- **coreutils** (macOS): For testing
- **GitHub Actions**: CI/CD (Priority 5)

---

## 🔄 Adjustment Process

This roadmap should be reviewed and updated:

**Weekly** during Priority 0-1:
- Track test pass rate
- Adjust estimates based on actual progress
- Reprioritize if blockers discovered

**Bi-weekly** during Priority 2-3:
- Review milestone progress
- Adjust scope if needed
- Update documentation

**Monthly** during Priority 4-5:
- Major feature assessment
- Community feedback integration
- Release planning

### Success Metrics

**Technical**:
- Test pass rate increasing weekly
- No regressions in passing tests
- Code quality maintained (no technical debt)
- Documentation kept current

**Process**:
- Hitting milestone dates (±1 week acceptable)
- Issues triaged within 48 hours
- PRs reviewed within 3 days
- Weekly progress updates

**Quality**:
- All new code has tests
- Test coverage >80%
- No known critical bugs
- Performance benchmarks stable

---

## 📝 Notes

### About This Roadmap

This roadmap was created based on:
- Current codebase analysis (40,490 lines Go code)
- Test suite analysis (123 .uni test files)
- Specification review (2,121 lines)
- Industry best practices for compiler development
- Realistic effort estimates for each feature

### Assumptions

1. **Single developer full-time**: Timeline assumes ~20 hours/week
2. **No major blockers**: Assumes predictable progress
3. **Specification stable**: No major design changes
4. **Infrastructure works**: Test harness reliable after Priority 0

### Risks

**High Risk**:
- Type system issues deeper than expected (could add 1-2 weeks)
- Generic monomorphization more complex than estimated (+1 week)
- Ownership system very complex (+2-4 weeks)

**Medium Risk**:
- Test infrastructure changes needed
- WASM runtime limitations discovered
- Performance issues requiring optimization

**Low Risk**:
- String operations take longer than expected
- Standard library larger than planned

### Mitigation Strategies

1. **Start with hardest problems first** (type system, generics)
2. **Build comprehensive tests before features**
3. **Regular progress reviews** (weekly standups)
4. **Keep scope flexible** (can defer Priority 5)
5. **Focus on milestones** (not individual tasks)

---

## 📚 Related Documents

- [PROJECT_STATUS_2026-02-02.md](PROJECT_STATUS_2026-02-02.md) - Current detailed status
- [TODO.md](TODO.md) - Detailed task breakdown (needs update)
- [spec/UnifiedSpecifiation.md](spec/UnifiedSpecifiation.md) - Language specification
- [CLAUDE.md](CLAUDE.md) - AI agent instructions
- [CONTRIBUTING.md](CONTRIBUTING.md) - How to contribute

---

*Created: February 2, 2026*
*Next Review: After Priority 0 completion*
*Status: Active - awaiting Priority 0 start*
