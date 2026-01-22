# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is the implementation of the **Unified Programming Language** - a modern systems programming language designed to combine memory safety, performance, and developer ergonomics. The language draws inspiration from Rust, Go, Swift, and Elixir while establishing its own unique approach to programming challenges.

The compiler uses a **virtual machine (VM) architecture** written in Go, providing excellent portability across all platforms where Go runs, while maintaining good performance.

## Core Architecture

### Compiler Structure (`src/unified-compiler/`)

The Unified compiler is built in Go and uses ANTLR4 for parsing, with a custom bytecode generator and virtual machine:

- **`cmd/compiler/main.go`**: Main entry point, handles command-line arguments and orchestrates compilation pipeline
- **`grammar/`**: ANTLR4 grammar files (`UnifiedLexer.g4`, `UnifiedParser.g4`) defining language syntax
- **`internal/ast/`**: Abstract Syntax Tree definitions and visitor pattern implementation
- **`internal/parser/`**: ANTLR-generated parser code (generated from grammar files)
- **`internal/bytecode/`**: Bytecode generator that converts AST to VM bytecode instructions
- **`internal/vm/`**: Virtual machine that executes bytecode (stack-based)
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
4. **Bytecode Generation**: AST → bytecode instructions (`bytecode/generator.go`)
5. **VM Execution**: Bytecode → program execution (`vm/vm.go`)

The compiler is currently in Phase 1 implementation, supporting:
- Basic function declarations and calls
- Variable declarations and assignments
- Arithmetic and logical expressions
- Control flow (if statements, basic loops)
- Literal values (integers, floats, booleans, strings)

## Key Dependencies

- **ANTLR4**: Parser generator, must be installed separately
- **ANTLR Go runtime**: `github.com/antlr4-go/antlr/v4` for generated parser

Note: Previous LLVM dependencies have been removed in favor of the custom VM implementation.

## Language Specification

The complete language specification is documented in `spec/UnifiedSpecifiation.md` (2000+ lines), covering:
- Memory management with ownership and borrowing
- Type system with generics and constraints
- Concurrency with actors and async/await
- Pattern matching and algebraic data types
- Standard library organization
- VM bytecode implementation guidelines

## Test Files

Example Unified programs in `test/`:
- `fib.uni`: Fibonacci implementations (recursive and iterative)
- `fizz.uni`: FizzBuzz-style example

## Implementation Status

**Current Phase 1** supports basic language constructs with VM execution. Future phases will add:
- **Phase 2**: Advanced expressions and control flow
- **Phase 3**: Memory management (ownership, borrowing, regions)
- **Phase 4**: Advanced type system (generics, interfaces, type inference)  
- **Phase 5**: Concurrency (actors, async/await, channels)
- **Phase 6**: Standard library and tooling

## Important Notes

- The compiler uses a custom Go-based virtual machine for execution
- Grammar files use ANTLR4 LL(*) parsing approach
- AST nodes implement visitor pattern for traversal
- Bytecode is stack-based for simplicity
- VM is highly portable, running anywhere Go runs
- Project follows Go module structure with `go.mod` in compiler directory