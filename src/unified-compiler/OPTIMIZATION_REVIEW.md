# Future Optimization and Review Tasks

This document tracks areas identified during testing that need optimization, review, or enhancement for future work.

## Critical Issues (Must Fix Before Production)

### 1. ANTLR Version Mismatch
**Status**: Blocked  
**Priority**: High  
**Issue**: Project uses ANTLR 4.13+ Go runtime, but system ANTLR is 4.9.2  
**Impact**: Cannot regenerate parser from grammar changes  
**Solution**: 
- Install ANTLR 4.13.1 from official source
- Alternative: Use Go-based ANTLR generator if available
- Document installation process in README

**Related Issues**:
- Cannot fix FloatLiteral lexer rule properly
- Cannot fix struct instantiation grammar ambiguity
- Forces workarounds (spaces in ranges, parentheses for structs)

### 2. Range Operator Lexer Ambiguity
**Status**: Workaround in place  
**Priority**: High  
**Issue**: `1..4` is lexed as `1.` + `.4` instead of `1` + `..` + `4`  
**Current Workaround**: Require spaces around range operator (`1 .. 4`)  
**Proper Fix**:
```antlr
# In UnifiedLexer.g4, line 113:
FloatLiteral: [0-9]+ '.' [0-9]+ ([eE] [+-]? [0-9]+)? | [0-9]+ [eE] [+-]? [0-9]+;
```
**Testing Required**:
- Verify float literals still parse correctly
- Test edge cases: `1.0`, `1.0e10`, `.5` (should fail)
- Test ranges: `1..10`, `1..=10`, `a..b`
- Add negative tests for invalid syntax

### 3. Struct Instantiation Grammar Ambiguity  
**Status**: Workaround in place  
**Priority**: Medium  
**Issue**: `Point { x: 10 }` doesn't parse, must use `(Point { x: 10 })`  
**Root Cause**: Parser commits to matching identifier before seeing `{`  
**Proper Fix Options**:
1. Reorder primary alternatives (risky, may break other parsing)
2. Use semantic predicate to look ahead for `{`
3. Merge identifier and structExpr into single production:
   ```antlr
   primary: identifier (LBRACE fieldInitList RBRACE)?
   ```
4. Make struct instantiation a statement-level construct

**Recommendation**: Option 3 is cleanest but changes AST structure  
**Testing Required**:
- All existing identifier uses still work
- Struct instantiation works without parens
- Nested struct expressions work
- Structs in function arguments work

## Performance Optimization Opportunities

### VM Instruction Dispatch
**Current**: Big switch statement (~55 cases)  
**Measured Performance**: Not yet benchmarked  
**Optimization Options**:
1. **Jump Table**: Let compiler generate (Go likely already does this)
2. **Direct Threading**: Requires C or assembly (future backend)
3. **Instruction Combining**: Fuse common sequences
   - Example: `PUSH_LOCAL n` instead of `PUSH n; LOAD_LOCAL`
   - Example: `CALL_DIRECT` for known functions

**Next Steps**:
1. Create comprehensive benchmarks
2. Profile real-world programs
3. Measure impact of different dispatch methods
4. Only optimize if proven bottleneck

### Call Frame Allocation
**Current**: Allocate 100-element array per function call  
**Impact**: ~800 bytes per call (assuming 8-byte values)  
**Issues**:
- Most functions use <10 locals, wasting 90% of allocation
- No frame reuse, causing GC pressure

**Optimization Options**:
1. **Static Analysis**: Determine max locals per function at compile time
2. **Frame Pooling**: Reuse frames from a sync.Pool
3. **Separate Stacks**: Use separate stack for locals (more complex)
4. **Hybrid**: Small frames from pool, large frames allocated fresh

**Implementation Steps**:
1. Add local count to bytecode function metadata
2. Allocate exact size needed per function
3. Measure memory improvement
4. Add frame pooling if still needed

### Stack Capacity Tuning
**Current**: Preallocate 256 value capacity  
**Measured Usage**: Not tracked  
**Next Steps**:
1. Add instrumentation to track max stack depth
2. Run test suite and record statistics
3. Adjust initial capacity based on 95th percentile
4. Consider per-function stack allocation hints

### Value Type Boxing
**Current**: All values boxed in 48-byte struct  
**Impact**: High memory usage, poor cache locality  
**Mitigation Options**:
1. **Unboxed Arrays**: Store primitive arrays contiguously
2. **Tagged Pointers**: Use NaN boxing for numbers (64-bit)
3. **Value Specialization**: Generate specialized code for types
4. **Escape Analysis**: Stack-allocate non-escaping values

**Complexity**: High - requires major refactoring  
**When**: After profiling shows it's a real bottleneck

### Constant Pool Access  
**Current**: Map lookup for function names  
**Optimization**: Array indexing with compile-time function IDs  
**Impact**: Moderate (only affects function calls)  
**Implementation**:
1. Assign numeric ID to each function during compilation
2. Store function entry points in array indexed by ID
3. Use OpCall with function ID instead of name lookup

## Code Quality Improvements

### Error Messages and Debugging

#### 1. Source Location in Errors
**Current**: Runtime errors show instruction pointer  
**Needed**: Line and column numbers from source file  
**Implementation**:
- Store source location map (IP → Position) in bytecode
- Include in error messages
- Add to stack traces

#### 2. Stack Traces on Panic
**Current**: Single error message  
**Needed**: Full call stack with function names and locations  
**Implementation**:
- Maintain function name in CallFrame
- Walk call stack on error
- Format as traditional stack trace

#### 3. Debug Mode
**Current**: No debug output  
**Needed**: Instruction tracing, stack visualization  
**Implementation**:
- Add --debug flag to compiler
- Print each instruction before execution
- Show stack and locals state
- Highlight changes

### Test Coverage Gaps

#### 1. Nested Struct Mutations
**Current**: Basic field access works  
**Missing**: Mutable field assignment in nested structs  
**Example**:
```unified
let mut person = Person { addr: Address { street: 0 } };
person.addr.street = 123;  // Not tested
```
**Action**: Add tests when OpStoreField for nested access is implemented

#### 2. Numeric Edge Cases
**Missing Tests**:
- Integer overflow/underflow
- Division by zero (tested for OpDiv, need OpMod)
- Float NaN and infinity handling
- Type conversion edge cases

**Action**: Create comprehensive numeric_edge_cases.uni test

#### 3. String Unicode Handling
**Current**: Basic ASCII string operations work  
**Missing**: Unicode-aware operations  
**Tests Needed**:
- Multi-byte character handling
- String length with emoji
- Substring with grapheme clusters
- Case conversion for non-ASCII

**Action**: Add unicode_strings.uni test when UTF-8 support added

#### 4. Error Recovery
**Current**: First error stops compilation  
**Missing**: Continue after errors, report multiple  
**Action**: Enhance error handling in parser and semantic analyzer

## Feature Completeness

### Phase 5 Completion (Structs)
**Completed**:
- ✅ Struct declarations
- ✅ Struct instantiation (with parentheses workaround)
- ✅ Field access
- ✅ Nested structs

**Blocked**:
- ❌ Methods on structs (need parser regeneration)
- ❌ Associated functions (need parser regeneration)

**Partially Complete**:
- ⏳ Mutable field assignment (OpStoreField exists, needs testing)

**Action Items**:
1. Fix ANTLR installation issue
2. Regenerate parser with method syntax
3. Test method calls and associated functions
4. Add comprehensive struct mutation tests

### Phase 6 Features (Advanced Types)
**Priority**: Medium  
**Dependencies**: Phase 5 completion  
**Features Needed**:
- Default parameter values
- Optional types (Option<T>)
- Result types for error handling
- Generic functions and structs
- Trait/interface system

### Phase 7 Features (Concurrency)
**Priority**: Low (foundational work needed)  
**Major Changes Required**:
- Actor model implementation
- Message passing primitives
- Async/await runtime
- Channel types
- Goroutine-style concurrency

## Documentation Gaps

### 1. Language Guide
**Missing**: Comprehensive language tutorial  
**Needed Sections**:
- Getting started
- Basic syntax
- Control flow
- Functions and closures
- Structs and methods
- Error handling
- Concurrency (when implemented)

### 2. VM Architecture Documentation
**Current**: VM_README.md covers basics  
**Missing**:
- Detailed instruction reference
- Bytecode format specification
- Calling convention details
- Memory layout diagrams
- Performance characteristics

### 3. Grammar Documentation
**Current**: ANTLR grammar files have minimal comments  
**Needed**:
- Operator precedence table
- Grammar rules explanation
- Common parsing pitfalls
- How to extend the grammar

### 4. Contributing Guide Enhancements
**Missing**:
- Code style guidelines
- Test writing guidelines
- Performance testing procedures
- Release process
- Commit message format

## Tooling Needs

### 1. Language Server Protocol (LSP)
**Priority**: High for editor support  
**Features Needed**:
- Syntax highlighting
- Auto-completion
- Go to definition
- Find references
- Inline error checking
- Code formatting

### 2. Debugger Integration
**Priority**: Medium  
**Options**:
- GDB/LLDB integration (if switching to native compilation)
- Custom debugger for VM
- VSCode debug adapter protocol

### 3. Package Manager
**Priority**: Medium  
**Needed For**: Standard library and third-party packages  
**Features**:
- Dependency resolution
- Version management
- Package repository
- Build system integration

### 4. Formatter
**Priority**: Low (can use external tools initially)  
**Behavior**: Enforce consistent style  
**Challenges**: Preserve semantic whitespace

## Testing Infrastructure

### 1. Continuous Integration
**Current**: Likely manual testing  
**Needed**:
- Automated test runs on commit
- Multiple platform testing
- Performance regression detection
- Coverage reporting

### 2. Fuzzing
**Priority**: Medium for robustness  
**Targets**:
- Parser (random valid/invalid input)
- Semantic analyzer
- Bytecode generator
- VM (random bytecode sequences)

### 3. Property-Based Testing
**Priority**: Low  
**Use Cases**:
- Verify compiler optimizations preserve semantics
- Test VM instruction commutativity
- Validate type system properties

### 4. Benchmark Suite
**Priority**: Medium  
**Benchmarks Needed**:
- Fibonacci (recursion)
- Array sorting (memory access patterns)
- String manipulation
- Struct allocation and access
- Function call overhead
- Real-world scenarios

## Security Considerations

### 1. Memory Safety
**Current Status**: Go provides safety at VM level  
**Potential Issues**:
- Array bounds checking (implemented)
- Integer overflow (not checked)
- Stack overflow (not checked)

**Actions**:
1. Add stack depth limit
2. Add configurable recursion limit
3. Consider checked arithmetic mode

### 2. Resource Limits
**Missing**:
- Maximum heap size
- Execution time limits
- Instruction count limits

**Use Case**: Running untrusted code safely  
**Implementation**: Add VM configuration options

### 3. Input Validation
**Parser**: Uses ANTLR error recovery  
**Semantic Analysis**: Validates types and scopes  
**VM**: Validates bytecode structure  
**Action**: Add fuzzing to find edge cases

## Maintenance Tasks

### Regular Reviews Needed
1. **Dependency Updates**: Keep ANTLR and Go runtime current
2. **Test Suite Maintenance**: Remove obsolete tests, add new ones
3. **Documentation Updates**: Keep in sync with code changes
4. **Performance Monitoring**: Track regression over time
5. **Issue Triage**: Regular review of bug reports and feature requests

### Technical Debt
1. **Grammar Cleanup**: Remove unused rules, clarify ambiguities
2. **Code Duplication**: Refactor similar test patterns
3. **Magic Numbers**: Replace with named constants (frame size, stack capacity)
4. **Error Messages**: Make more consistent and helpful
5. **Naming Conventions**: Ensure consistency across codebase

## Timeline Recommendations

### Immediate (Next Sprint)
1. Fix ANTLR installation
2. Regenerate parser with grammar fixes
3. Remove workarounds from test files
4. Complete Phase 5 struct features

### Short Term (1-2 Months)
1. Add comprehensive numeric edge case tests
2. Implement frame allocation optimization
3. Create language tutorial documentation
4. Set up CI/CD pipeline

### Medium Term (3-6 Months)
1. Begin Phase 6 implementation
2. Create LSP server
3. Add benchmark suite
4. Implement resource limits for VM

### Long Term (6-12 Months)
1. Complete Phase 6 and 7
2. Consider JIT compilation
3. Build standard library
4. Create ecosystem tooling

---

**Last Updated**: January 26, 2026  
**Next Review**: Before Phase 6 implementation begins
