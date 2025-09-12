# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is the implementation of the **Unified Programming Language** - a modern systems programming language designed to combine memory safety, performance, and developer ergonomics. The language draws inspiration from Rust, Go, Swift, and Elixir while establishing its own unique approach to programming challenges.

## Core Architecture

### Compiler Structure (`src/unified-compiler/`)

The Unified compiler is built in Go and uses ANTLR4 for parsing, with LLVM as the code generation backend:

- **`cmd/compiler/main.go`**: Main entry point, handles command-line arguments and orchestrates compilation pipeline
- **`grammar/`**: ANTLR4 grammar files (`UnifiedLexer.g4`, `UnifiedParser.g4`) defining language syntax
- **`internal/ast/`**: Abstract Syntax Tree definitions and visitor pattern implementation
- **`internal/parser/`**: ANTLR-generated parser code (generated from grammar files)
- **`internal/codegen/`**: LLVM IR code generation from AST
- **`internal/semantic/`**: Semantic analysis (not yet implemented)
- **`test/`**: Test cases including `.uni` source files like `fib.uni` and `fizz.uni`

### Language Features

Unified implements several advanced programming language features:
- Memory safety through ownership system (similar to Rust but more ergonomic)
- Actor model for concurrency with built-in async/await
- Pattern matching and algebraic data types
- Type inference and generic programming
- Region-based memory management as alternative to garbage collection
- Expression-oriented syntax where everything returns a value

## Development Commands

### Building the Compiler
```bash
cd src/unified-compiler
make parser  # Generate parser from ANTLR grammar
make build   # Build the compiler binary
```

### Testing
```bash
make test    # Run Go tests
./bin/unified-compiler --input test/fib.uni --output fib.ll  # Compile test file
```

### Clean Build
```bash
make clean   # Remove generated files and binaries
```

### Parser Generation
The parser must be regenerated whenever grammar files are modified:
```bash
bash scripts/generate.sh  # Requires ANTLR4 to be installed
```

## Code Generation Pipeline

1. **Lexical Analysis**: Source code → tokens (ANTLR lexer)
2. **Parsing**: Tokens → parse tree (ANTLR parser) 
3. **AST Building**: Parse tree → AST (`ast.go` visitor pattern)
4. **Code Generation**: AST → LLVM IR (`codegen/generator.go`)
5. **Output**: LLVM IR written to `.ll` file

The compiler is currently in Phase 1 implementation, supporting:
- Basic function declarations and calls
- Variable declarations and assignments
- Arithmetic and logical expressions
- Control flow (if statements, basic loops)
- Literal values (integers, floats, booleans, strings)

## Key Dependencies

- **ANTLR4**: Parser generator, must be installed separately
- **Go LLVM bindings**: `github.com/llvm/bindings/go/llvm` for code generation
- **ANTLR Go runtime**: `github.com/antlr4-go/antlr/v4` for generated parser

## Language Specification

The complete language specification is documented in `spec/UnifiedSpecifiation.md` (2000+ lines), covering:
- Memory management with ownership and borrowing
- Type system with generics and constraints
- Concurrency with actors and async/await
- Pattern matching and algebraic data types
- Standard library organization
- LLVM implementation guidelines

## Test Files

Example Unified programs in `test/`:
- `fib.uni`: Fibonacci implementations (recursive and iterative)
- `fizz.uni`: FizzBuzz-style example

## Implementation Status

**Current Phase 1** supports basic language constructs. Future phases will add:
- **Phase 2**: Memory management (ownership, borrowing, regions)
- **Phase 3**: Advanced type system (generics, interfaces, type inference)  
- **Phase 4**: Concurrency (actors, async/await, channels)
- **Phase 5**: Standard library and tooling

## Important Notes

- The compiler currently targets LLVM IR generation, not direct machine code
- Grammar files use ANTLR4 LL(*) parsing approach
- AST nodes implement visitor pattern for traversal
- LLVM bindings require proper LLVM installation
- Project follows Go module structure with `go.mod` in compiler directory