# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is the implementation of the **Unified Programming Language** - a modern systems programming language designed to combine memory safety, performance, and developer ergonomics. The language draws inspiration from Rust, Go, Swift, and Elixir while establishing its own unique approach to programming challenges.

The compiler uses a **WebAssembly (WASM) backend** written in Go, providing excellent portability across all platforms. The WASM runtime is wazero, a pure Go implementation with zero C dependencies.

## Core Architecture

### Compiler Structure (`src/unified-compiler/`)

The Unified compiler is built in Go and uses ANTLR4 for parsing, with a WebAssembly code generator:

- **`cmd/compiler/main.go`**: Main entry point, handles command-line arguments and orchestrates compilation pipeline
- **`grammar/`**: ANTLR4 grammar files (`UnifiedLexer.g4`, `UnifiedParser.g4`) defining language syntax
- **`internal/ast/`**: Abstract Syntax Tree definitions and visitor pattern implementation
- **`internal/parser/`**: ANTLR-generated parser code (generated from grammar files)
- **`internal/wasm/`**: WebAssembly code generator that converts AST to WASM bytecode
- **`internal/bytecode/`**: Type definitions (shared between WASM and legacy code)
- **`internal/vm/`**: Legacy VM code (deprecated, replaced by WASM)
- **`internal/semantic/`**: Semantic analysis (symbol table, type checking, type inference)
- **`test/`**: Test cases including 121 `.uni` source files

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
# Run all .uni test files (from repository root)
./test_all_uni_files.sh

# Run Go unit tests
cd src/unified-compiler
make test

# Compile and run a specific test file
./bin/unified --input test/integration/simple_return.uni
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
4. **WASM Generation**: AST → WASM bytecode (`wasm/generator.go`, `wasm/codegen.go`)
5. **WASM Encoding**: WASM module → binary format (`wasm/encoder.go`)
6. **Runtime Execution**: WASM binary → wazero runtime → program execution

The compiler currently supports (26 of 121 tests passing):
- Basic function declarations and calls
- Variable declarations and assignments
- Arithmetic, logical, and bitwise expressions
- Control flow (if statements, while loops)
- Mutable variables with assignment
- Literal values (integers, floats, booleans)
- Optional semicolons
- Basic enums and generics

**In Progress** (see TODO.md for details):
- Struct operations (heap allocation, field access)
- Array operations (literals, indexing, iteration)
- For loops with ranges
- String operations (length, concat, etc.)
- Advanced generics
- Try operator (?) for error handling

## Key Dependencies

- **ANTLR4**: Parser generator, must be installed separately
- **ANTLR Go runtime**: `github.com/antlr4-go/antlr/v4` for generated parser
- **wazero**: `github.com/tetratelabs/wazero` - Pure Go WebAssembly runtime (zero C dependencies)

## Language Specification

The complete language specification is documented in `spec/UnifiedSpecifiation.md` (2000+ lines), covering:
- Memory management with ownership and borrowing
- Type system with generics and constraints
- Concurrency with actors and async/await
- Pattern matching and algebraic data types
- Standard library organization
- WebAssembly compilation target

## Test Files

The repository contains 121 `.uni` test files across multiple directories:
- `test/integration/` - Integration tests for language features
- `test/generics/` - Generic function tests
- `test/stdlib/` - Standard library tests
- Root test files - Basic functionality tests

**Current Test Results:** 26 passing (21.5%), 95 failing (78.5%)

See `TODO.md` for detailed analysis of test results and implementation roadmap.

## Implementation Status

**Current Phase:** WASM Backend Feature Completion

The compiler has migrated from a custom VM to WebAssembly. Core architecture is complete, but many language features need WASM code generation implementation.

**Test Status:**
- **26 tests passing (21.5%)** - Basic features work (functions, variables, if/else, while, arithmetic, logic)
- **95 tests failing (78.5%)** - Advanced features need implementation (structs, arrays, for loops, strings, etc.)

**Priority Tasks** (see TODO.md for complete details):

1. **Fix struct support** - WASM global.get index bugs (4 tests)
2. **Fix array support** - Memory allocation and indexing (11 tests)
3. **Implement for loops** - Range operator support (4 tests)
4. **Fix string operations** - Type mismatches and runtime functions (11 tests)
5. **Improve generics** - Advanced type inference (15 tests)
6. **Add try operator** - Error handling with ? operator (10 tests)

Once these are complete, ~82 tests should pass (68%).

## Important Notes

- The compiler uses WebAssembly as the compilation target (WASM 1.0 spec)
- Runtime is wazero - pure Go, zero C dependencies, highly portable
- Grammar files use ANTLR4 LL(*) parsing approach
- AST nodes implement visitor pattern for traversal
- WASM bytecode is stack-based following the standard
- Many language features work, but need WASM codegen implementation
- Project follows Go module structure with `go.mod` in compiler directory
- Test suite has 121 .uni files - run `./test_all_uni_files.sh` to verify changes