# Unified Compiler Architecture

Version 2.0.0

## Overview

This document describes the architecture of the Unified Programming Language compiler. The compiler is built in Go and uses ANTLR4 for parsing, with a custom virtual machine (VM) written in Go as the execution backend. This VM-based approach provides excellent portability while maintaining good performance.

## Architecture Diagram

```
┌─────────────────┐
│  Source Code    │
│    (.uni)       │
└────────┬────────┘
         │
         ▼
┌─────────────────┐
│     Lexer       │  (ANTLR4)
│  UnifiedLexer   │
└────────┬────────┘
         │ tokens
         ▼
┌─────────────────┐
│     Parser      │  (ANTLR4)
│  UnifiedParser  │
└────────┬────────┘
         │ parse tree
         ▼
┌─────────────────┐
│  AST Builder    │
│   (Visitor)     │
└────────┬────────┘
         │ AST
         ▼
┌─────────────────┐
│   Semantic      │
│   Analysis      │  (Future)
└────────┬────────┘
         │ validated AST
         ▼
┌─────────────────┐
│   Bytecode      │
│   Generator     │
└────────┬────────┘
         │ bytecode
         ▼
┌─────────────────┐
│  Virtual        │
│  Machine (VM)   │
└────────┬────────┘
         │ execution
         ▼
┌─────────────────┐
│  Program        │
│  Output         │
└─────────────────┘
```

## Components

### 1. Lexer (Tokenization)

**Location**: `src/unified-compiler/grammar/UnifiedLexer.g4`

**Purpose**: Convert source code into a stream of tokens

**Input**: Source code string (`.uni` file)

**Output**: Token stream

**Technology**: ANTLR4 lexer grammar

**Key Responsibilities**:
- Recognize keywords (fn, let, return, etc.)
- Recognize identifiers
- Recognize literals (numbers, strings, booleans)
- Recognize operators and punctuation
- Track line and column numbers for error reporting
- Filter whitespace and comments

**Example**:
```
Input:  "fn main() -> i32 { return 42 }"
Output: [FN, IDENTIFIER("main"), LPAREN, RPAREN, ARROW, TYPE("i32"), LBRACE, RETURN, NUMBER(42), RBRACE]
```

### 2. Parser (Syntax Analysis)

**Location**: `src/unified-compiler/grammar/UnifiedParser.g4`

**Purpose**: Build a parse tree from tokens according to grammar rules

**Input**: Token stream

**Output**: Parse tree (ANTLR-generated tree structure)

**Technology**: ANTLR4 parser grammar

**Key Responsibilities**:
- Validate syntax according to grammar
- Build hierarchical parse tree
- Report syntax errors with locations
- Handle operator precedence
- Manage context-sensitive parsing

**Example**:
```
Tokens:  [FN, IDENTIFIER("main"), LPAREN, RPAREN, ARROW, TYPE("i32"), ...]
Output:  FunctionDeclaration
           ├── name: "main"
           ├── parameters: []
           ├── returnType: "i32"
           └── body: BlockStatement
                 └── ReturnStatement
                       └── IntLiteral(42)
```

### 3. AST Builder (Visitor Pattern)

**Location**: `src/unified-compiler/internal/ast/`

**Purpose**: Convert parse tree to Abstract Syntax Tree (AST)

**Input**: Parse tree (from ANTLR parser)

**Output**: AST (custom Go structs)

**Technology**: Visitor pattern over ANTLR parse tree

**Key Responsibilities**:
- Traverse parse tree using visitor pattern
- Convert parse tree nodes to AST nodes
- Simplify tree structure
- Resolve syntactic sugar
- Prepare tree for semantic analysis

**Files**:
- `ast.go` - AST node definitions
- `visitor.go` - Parse tree to AST conversion

**Example AST Nodes**:
```go
type FunctionDecl struct {
    Name       string
    Parameters []Parameter
    ReturnType Type
    Body       *BlockStmt
}

type BinaryExpr struct {
    Left  Expr
    Op    string
    Right Expr
}
```

### 4. Semantic Analysis (Future)

**Location**: `src/unified-compiler/internal/semantic/` (not yet implemented)

**Purpose**: Validate program semantics and add type information

**Input**: AST

**Output**: Validated and annotated AST

**Key Responsibilities**:
- Type checking
- Variable scope resolution
- Function signature validation
- Ownership and borrow checking
- Dead code detection
- Unreachable code detection
- Constant evaluation

**Phases** (planned):
1. **Symbol Table Construction**: Build table of all symbols
2. **Name Resolution**: Resolve all identifiers to definitions
3. **Type Checking**: Validate all types and add type annotations
4. **Borrow Checking**: Validate ownership and borrowing rules
5. **Control Flow Analysis**: Validate control flow (returns, loops, etc.)

### 5. Bytecode Generator

**Location**: `src/unified-compiler/internal/bytecode/`

**Purpose**: Generate bytecode instructions from AST

**Input**: Validated AST

**Output**: Bytecode instruction stream

**Technology**: Custom Go implementation

**Key Responsibilities**:
- Translate AST nodes to bytecode instructions
- Manage bytecode instruction stream
- Handle constant pool for literals
- Generate stack-based operations
- Emit function definitions and calls
- Handle control flow (branches, loops)

**Files**:
- `generator.go` - Main bytecode generator
- `instructions.go` - Bytecode instruction definitions
- `generator_test.go` - Code generation tests

**Bytecode Features**:
- Stack-based virtual machine architecture
- Simple instruction set for Phase 1
- Support for basic arithmetic and control flow
- Function call/return mechanisms

### 6. Virtual Machine (VM)

**Location**: `src/unified-compiler/internal/vm/`

**Purpose**: Execute bytecode instructions

**Input**: Bytecode instruction stream

**Output**: Program execution and output

**Technology**: Custom Go-based stack VM

**Key Responsibilities**:
- Execute bytecode instructions
- Manage execution stack
- Handle function calls and returns
- Manage local variables
- Provide runtime library functions
- Handle program I/O

**Files**:
- `vm.go` - Virtual machine implementation
- `stack.go` - Execution stack management
- `runtime.go` - Runtime library functions
- `vm_test.go` - VM execution tests

**VM Architecture**:
- Stack-based execution model
- Register-less design for simplicity
- Built-in runtime functions (print, etc.)
- Garbage collection (Go's GC)
- Portable across all Go-supported platforms

## Data Flow

### Compilation and Execution Pipeline

```
1. Read source file
2. Lexer: source → tokens
3. Parser: tokens → parse tree
4. AST Builder: parse tree → AST
5. Semantic Analysis: AST → validated AST
6. Bytecode Generator: validated AST → bytecode
7. Virtual Machine: bytecode → execution and output
```

### Error Handling

Errors can occur at each stage:

1. **Lexer Errors**: Invalid characters, malformed literals
2. **Parser Errors**: Syntax errors, unexpected tokens
3. **Semantic Errors**: Type errors, undefined variables, borrow check failures
4. **Bytecode Generation Errors**: Invalid instruction generation
5. **VM Runtime Errors**: Division by zero, stack overflow, type errors

Each error should include:
- File name
- Line and column number
- Error message
- Suggested fix (when possible)

## Module Organization

```
src/unified-compiler/
├── cmd/
│   └── compiler/
│       └── main.go              # Entry point, CLI handling
├── grammar/
│   ├── UnifiedLexer.g4          # Lexer grammar (ANTLR4)
│   └── UnifiedParser.g4         # Parser grammar (ANTLR4)
├── internal/
│   ├── ast/
│   │   ├── ast.go              # AST node definitions
│   │   └── visitor.go          # Parse tree → AST visitor
│   ├── parser/
│   │   ├── unified_lexer.go    # Generated lexer (from ANTLR)
│   │   ├── unifiedparser_*.go  # Generated parser (from ANTLR)
│   │   ├── debug_visitor.go    # Debug utilities
│   │   └── debug_listener.go   # Debug utilities
│   ├── bytecode/
│   │   ├── generator.go        # Bytecode generator
│   │   ├── instructions.go     # Instruction definitions
│   │   └── generator_test.go   # Tests
│   ├── vm/
│   │   ├── vm.go              # Virtual machine
│   │   ├── stack.go           # Execution stack
│   │   ├── runtime.go         # Runtime library
│   │   └── vm_test.go         # Tests
│   └── semantic/               # Future: semantic analysis
├── test/                       # Test programs (.uni files)
│   ├── fib.uni
│   └── fizz.uni
├── scripts/
│   └── generate.sh             # Parser generation script
├── Makefile                    # Build automation
└── go.mod                      # Go module definition
```

## Design Decisions

### Why ANTLR4?

**Pros**:
- Powerful LL(*) parsing algorithm
- Excellent error recovery
- Great tooling and IDE support
- Generates readable code
- Well-documented and mature

**Cons**:
- Additional dependency
- Generated code can be large
- Learning curve for grammar syntax

**Decision**: Use ANTLR4 for its power and tooling support

### Why Go?

**Pros**:
- Fast compilation
- Good concurrency support
- Simple, readable syntax
- Great standard library
- Cross-platform support
- No need for external toolchains
- Easy to embed a VM

**Cons**:
- Less mature compiler ecosystem than C++

**Decision**: Use Go for development velocity, simplicity, and VM portability

### Why a Virtual Machine?

**Pros**:
- Highly portable - works anywhere Go runs
- Simpler implementation than native code generation
- Easier debugging and testing
- No external dependencies (like LLVM)
- Faster compilation (no LLVM optimization passes)
- Better for iterative development
- Can add JIT compilation later if needed

**Cons**:
- Slower execution than native code (but still quite fast)
- Additional runtime overhead

**Decision**: Use a custom Go-based VM for portability and simplicity

### AST vs Parse Tree

**Parse Tree**: 
- Direct representation of grammar rules
- Contains all syntactic details
- Can be verbose

**AST**:
- Simplified, semantic representation
- Removes syntactic noise
- Easier to work with for analysis and code generation

**Decision**: Use both - parse tree from ANTLR, convert to AST for processing

## Extensibility

### Adding New Language Features

To add a new language feature:

1. **Syntax**: Update grammar files (.g4)
2. **AST**: Add new node types to ast.go
3. **Visitor**: Update visitor.go to build new nodes
4. **Semantic**: Add semantic checks (when implemented)
5. **Bytecode Gen**: Add bytecode generation for new nodes
6. **VM**: Add VM support for new bytecode instructions
7. **Tests**: Add tests for all stages

### Plugin Architecture (Future)

Planned plugin points:
- Custom optimizations
- Custom backends
- Custom linters
- Language extensions

## Performance Considerations

### Compilation Speed

Target: < 1 second for 10,000 lines of code

Optimizations:
- Incremental compilation (future)
- Parallel parsing (future)
- Efficient data structures
- Minimize memory allocations

### Runtime Performance

Target: Competitive with interpreted languages like Python

Strategies:
- Efficient bytecode design
- Stack-based VM for simplicity
- Minimal instruction overhead
- Future: Add JIT compilation for hot paths
- Future: Inline caching for method calls

## Testing Strategy

### Unit Tests

Test each component in isolation:
- Lexer tests: Token recognition
- Parser tests: Parse tree construction
- AST tests: AST node creation
- Bytecode gen tests: Bytecode generation
- VM tests: Instruction execution

### Integration Tests

Test the full pipeline:
- Compile .uni files end-to-end
- Execute and verify output
- Test error messages

### Benchmarks

Performance benchmarks:
- Compilation time
- Runtime performance
- Memory usage

## Future Enhancements

### Phase 1
- Implement basic bytecode generator
- Create stack-based virtual machine
- Support minimal working compiler

### Phase 2-5
- Complete language features
- Add semantic analysis
- Implement type system
- Add memory management

### Phase 6+
- Advanced features (generics, traits, etc.)
- JIT compilation for performance
- Developer tooling (REPL, debugger)
- Standard library

See `docs/planning/PHASED_ROADMAP.md` for complete plan.

## References

- [ANTLR4 Documentation](https://github.com/antlr/antlr4/blob/master/doc/index.md)
- [Virtual Machine Design](https://craftinginterpreters.com/a-virtual-machine.html)
- [Unified Specification](../../spec/UnifiedSpecifiation.md)
- [Phased Roadmap](../planning/PHASED_ROADMAP.md)

## Version History

- **v2.0** (January 2026): Changed to VM-based architecture
- **v1.0** (January 2026): Initial architecture document (LLVM-based)
