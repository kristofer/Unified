# GitHub Copilot Guide for Unified Language Implementation

This document provides guidance for GitHub Copilot and other AI coding assistants when working on the Unified Programming Language project.

## Project Overview

Unified is a modern systems programming language combining memory safety, performance, and developer ergonomics. It draws inspiration from Rust, Go, Swift, and Elixir.

### Key Features
- Memory safety through ownership system (similar to Rust but more ergonomic)
- Actor model for concurrency with built-in async/await
- Pattern matching and algebraic data types
- Type inference and generic programming
- Expression-oriented syntax

### Current Status
The project is in early development with a basic compiler infrastructure in place. We are following a phased implementation plan (see `docs/planning/PHASED_ROADMAP.md`).

## Architecture

### Compiler Pipeline

```
Source Code (.uni) → Lexer → Parser → AST → Code Generator → LLVM IR/WASM → Executable
```

### Directory Structure

```
Unified/
├── docs/                    # Documentation
│   ├── planning/           # Roadmaps and planning docs
│   ├── spec/               # Language specifications
│   └── design/             # Design documents
├── src/
│   └── unified-compiler/   # Main compiler implementation
│       ├── cmd/            # Executable entry points
│       ├── grammar/        # ANTLR4 grammar files
│       ├── internal/       # Internal packages
│       │   ├── ast/        # Abstract Syntax Tree
│       │   ├── parser/     # Generated parser (from ANTLR)
│       │   ├── codegen/    # Code generation
│       │   └── semantic/   # Semantic analysis (future)
│       └── test/           # Test programs (.uni files)
├── examples/               # Example programs
└── spec/                   # Language specification
```

## Working with Phases

### Current Phase

Always check the current phase before starting work:
- Read `docs/planning/PHASED_ROADMAP.md` to understand the current phase
- Focus only on features required for the current phase
- Don't implement features from future phases

### Phase Workflow

1. **Understand the Phase**
   - Read the phase goals, prerequisites, and deliverables
   - Review the test requirements
   - Understand the success criteria

2. **Test-Driven Development**
   - Write tests BEFORE implementation
   - Tests should cover all requirements from the phase
   - Run tests frequently during development

3. **Incremental Implementation**
   - Implement one feature at a time
   - Test after each feature
   - Commit working code frequently

4. **Documentation**
   - Update documentation as you implement features
   - Add code comments for complex logic
   - Update README if user-facing changes

## Language Specification

The complete language specification is in `spec/UnifiedSpecifiation.md`. Key sections:

- **Lexical Structure**: Keywords, operators, literals
- **Type System**: Primitives, structs, enums, generics
- **Memory Management**: Ownership, borrowing, regions
- **Concurrency**: Actors, async/await, channels
- **Pattern Matching**: Match expressions, patterns

## Compiler Components

### Lexer (ANTLR4)

Location: `src/unified-compiler/grammar/UnifiedLexer.g4`

The lexer converts source code into tokens. When adding new syntax:
1. Add token definitions to the lexer grammar
2. Regenerate parser with `make parser`
3. Test token recognition

### Parser (ANTLR4)

Location: `src/unified-compiler/grammar/UnifiedParser.g4`

The parser builds a parse tree from tokens. When adding new syntax:
1. Add grammar rules to the parser
2. Regenerate parser with `make parser`
3. Update AST builder to handle new parse tree nodes

### AST (Abstract Syntax Tree)

Location: `src/unified-compiler/internal/ast/`

The AST is a simplified representation of the program. When adding new features:
1. Define new node types in `ast.go`
2. Implement visitor methods in `visitor.go`
3. Update AST builder to convert parse trees to AST nodes

### Code Generator

Location: `src/unified-compiler/internal/codegen/`

Generates LLVM IR (or WASM) from AST. When adding new features:
1. Add code generation logic for new AST nodes
2. Test generated code compiles
3. Test generated code executes correctly

## Testing Strategy

### Test Organization

Tests are organized by phase:
```
src/unified-compiler/test/
├── phase1/          # Minimal compiler tests
├── phase2/          # Expression tests
├── phase3/          # Control flow tests
├── phase4/          # Function tests
└── ...
```

### Test Types

1. **Unit Tests**: Test individual components
   - Lexer tests: Token recognition
   - Parser tests: AST construction
   - Codegen tests: LLVM IR generation

2. **Integration Tests**: Test the full pipeline
   - Compile .uni files end-to-end
   - Execute and verify output
   - Check error messages

3. **Example Programs**: Real-world usage
   - Located in `examples/`
   - Should compile and run
   - Demonstrate language features

### Writing Tests

For Go unit tests:
```go
func TestFeature(t *testing.T) {
    // Arrange: Set up test data
    input := "fn main() -> i32 { return 42 }"
    
    // Act: Execute the code being tested
    result := compile(input)
    
    // Assert: Verify the result
    if result.ExitCode != 42 {
        t.Errorf("Expected 42, got %d", result.ExitCode)
    }
}
```

For .uni test files:
```unified
// test/phase2/arithmetic.uni
// Expected output: 12
fn main() -> i32 {
    return 5 + 7
}
```

## Common Tasks

### Adding a New Operator

1. Add token to `UnifiedLexer.g4`:
   ```antlr
   PLUS: '+';
   ```

2. Add grammar rule to `UnifiedParser.g4`:
   ```antlr
   expression
       : expression PLUS expression  # AddExpression
       | ...
       ;
   ```

3. Regenerate parser: `make parser`

4. Add AST node in `ast.go`:
   ```go
   type BinaryExpr struct {
       Left  Expr
       Op    string
       Right Expr
   }
   ```

5. Update visitor in `visitor.go` to build AST node

6. Add code generation in `generator.go`

7. Write tests

### Adding a New Statement Type

1. Add keywords to lexer if needed
2. Add grammar rule to parser
3. Regenerate parser
4. Add AST node type
5. Update visitor
6. Add code generation
7. Add semantic checks if needed
8. Write tests

### Adding a New Type

1. Add type to type system representation
2. Update parser to recognize type syntax
3. Add type checking logic
4. Add code generation for type
5. Write tests

## Code Style

### Go Code

- Follow standard Go formatting (`gofmt`)
- Use descriptive variable names
- Add comments for exported functions
- Keep functions focused and short
- Use early returns to reduce nesting

Example:
```go
// GenerateFunction generates LLVM IR for a function declaration
func (g *Generator) GenerateFunction(fn *ast.FunctionDecl) error {
    if fn == nil {
        return fmt.Errorf("function declaration is nil")
    }
    
    // Generate function signature
    sig := g.buildFunctionSignature(fn)
    
    // Generate function body
    return g.generateBody(fn.Body)
}
```

### ANTLR Grammar

- Use clear rule names (camelCase)
- Add comments for complex rules
- Group related rules together
- Use labels (#RuleName) for alternative productions

Example:
```antlr
// Expression with operator precedence
expression
    : primary                              # PrimaryExpression
    | expression ('*' | '/') expression    # MultiplicativeExpression
    | expression ('+' | '-') expression    # AdditiveExpression
    ;
```

### Unified Language Code

- Use 4-space indentation
- Follow Rust-like conventions
- Prefer `let` over `let mut` when possible
- Add comments for complex algorithms

Example:
```unified
// Calculate the nth Fibonacci number iteratively
fn fibonacci(n: i32) -> i32 {
    if n <= 1 {
        return n
    }
    
    let mut a = 0
    let mut b = 1
    
    for _ in 2..=n {
        let temp = a + b
        a = b
        b = temp
    }
    
    return b
}
```

## Build and Test Commands

### Building

```bash
cd src/unified-compiler

# Generate parser from grammar
make parser

# Build compiler
make build

# Do both
make all
```

### Testing

```bash
# Run all Go tests
make test

# Run specific test
go test -v ./internal/ast -run TestFunction

# Compile and run a .uni file
./bin/unified --input test/fib.uni
```

## Common Pitfalls

### 1. Parser Generation

**Problem**: Changes to grammar don't take effect

**Solution**: Always run `make parser` after modifying `.g4` files

### 2. Nil Pointer Errors

**Problem**: Visitor crashes on nil nodes

**Solution**: Always check for nil before accessing fields:
```go
if node == nil {
    return nil, fmt.Errorf("unexpected nil node")
}
```

### 3. Type Mismatches

**Problem**: LLVM IR generation fails due to type errors

**Solution**: Ensure type checking pass validates all types before code generation

### 4. Test Failures

**Problem**: Tests fail after changes

**Solution**: 
- Run tests before making changes to establish baseline
- Fix broken tests immediately
- Don't commit code with failing tests

## Debugging

### Debugging the Compiler

1. **Add Debug Logging**
   ```go
   log.Printf("DEBUG: Processing node: %+v", node)
   ```

2. **Use Debug Visitor**
   The parser includes a debug visitor in `internal/parser/debug_visitor.go`

3. **Inspect AST**
   Print the AST structure to understand the tree:
   ```go
   fmt.Printf("AST: %+v\n", astNode)
   ```

4. **Check Generated LLVM IR**
   Review the `.ll` output file to see what code was generated

### Debugging Unified Programs

Currently, debugging compiled Unified programs is limited. Future phases will add:
- Stack traces
- Debugger support
- Better error messages

## Resources

### Internal Documentation

- `docs/planning/PHASED_ROADMAP.md` - Complete implementation roadmap
- `spec/UnifiedSpecifiation.md` - Language specification
- `CLAUDE.md` - High-level project overview
- `README.md` - Getting started guide

### External Resources

- [ANTLR4 Documentation](https://github.com/antlr/antlr4/blob/master/doc/index.md)
- [LLVM Language Reference](https://llvm.org/docs/LangRef.html)
- [Go LLVM Bindings](https://pkg.go.dev/github.com/llir/llvm)
- [Smog Language](https://github.com/kristofer/smog) - Reference implementation

### Language Design Inspiration

- **Rust**: Ownership system, pattern matching, expression-oriented
- **Go**: Simplicity, concurrency primitives, fast compilation
- **Swift**: Clean syntax, type inference
- **Elixir**: Actor model, message passing

## Best Practices for AI Agents

1. **Always Read Context First**
   - Check current phase in PHASED_ROADMAP.md
   - Review existing code before modifying
   - Understand the test requirements

2. **Test-Driven Development**
   - Write tests before implementation
   - Run tests after each change
   - Fix failing tests immediately

3. **Incremental Changes**
   - Make small, focused changes
   - Test each change before moving on
   - Commit working code frequently

4. **Clear Communication**
   - Explain what you're doing and why
   - Document complex decisions
   - Ask for clarification when uncertain

5. **Follow the Plan**
   - Don't implement features from future phases
   - Focus on current phase goals
   - Complete one phase before moving to the next

6. **Quality Over Speed**
   - Write clean, maintainable code
   - Add comprehensive tests
   - Document as you go

## Getting Help

If you encounter issues:

1. Check existing documentation in `docs/`
2. Review the language specification
3. Look at the smog project for reference
4. Check ANTLR4 documentation for parser issues
5. Review LLVM documentation for code generation issues

## Version History

- **v1.0** (January 2026): Initial version for Phase 0-1 development

## Updates

This guide will be updated as the project evolves. Always check the latest version in the repository.
