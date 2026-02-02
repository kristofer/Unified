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

The compiler currently supports (56 of 123 tests passing):
- Basic function declarations and calls
- Variable declarations and assignments
- Arithmetic, logical, and bitwise expressions
- Control flow (if statements, while loops, for loops)
- Mutable variables with assignment
- Literal values (integers, floats, booleans)
- Optional semicolons
- Basic enums and generics
- Structs with field access ✅
- Arrays (literals, indexing, iteration, bounds checking) ✅
- Strings (length, concat, trim, case, search, substring) ✅
- For loops with ranges ✅
- Nested loops with break/continue ✅

**In Progress** (see TODO.md and TEST_BASELINE_2026-02-02.md for details):
- Parser grammar improvements (Self, new, mut self, field syntax) - **HIGHEST PRIORITY**
- Advanced generics (monomorphization, type inference)
- Try operator (?) codegen completion
- Type system improvements (i32/i64 coercion)
- Edge cases for structs, arrays, strings

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

The repository contains 123 `.uni` test files across multiple directories:
- `test/integration/` - Integration tests for language features
- `test/generics/` - Generic function tests
- `test/stdlib/` - Standard library tests
- Root test files - Basic functionality tests

**Current Test Results:** 56 passing (45.5%), 67 failing (54.5%)

See `TODO.md` for detailed analysis of test results and `TEST_BASELINE_2026-02-02.md` for comprehensive baseline report.

## Implementation Status

**Current Phase:** WASM Backend Feature Completion

The compiler has migrated from a custom VM to WebAssembly. Core architecture is complete, but many language features need WASM code generation implementation.

**Test Status:**
- **56 tests passing (45.5%)** - Core features work: functions, variables, if/else, while, for loops, arrays, strings, structs, enums
- **67 tests failing (54.5%)** - Mainly parser grammar gaps (Self, new, mut self) and advanced features (generics monomorphization)

**Priority Tasks** (see TODO.md and TEST_BASELINE_2026-02-02.md for complete details):

1. **Add parser grammar features** - Self, new, mut self, field syntax (26 tests, 21% improvement) - **HIGHEST IMPACT**
2. **Fix type system issues** - i32/i64 coercion (7 tests)
3. **Complete try operator** - Codegen for all cases (7 tests)
4. **Improve generics** - Monomorphization and advanced type inference (13 tests)
5. **Fix edge cases** - Structs, arrays, strings (8 tests)

Once parser grammar is complete, ~82 tests should pass (67%).

## Important Notes

- The compiler uses WebAssembly as the compilation target (WASM 1.0 spec)
- Runtime is wazero - pure Go, zero C dependencies, highly portable
- Grammar files use ANTLR4 LL(*) parsing approach
- AST nodes implement visitor pattern for traversal
- WASM bytecode is stack-based following the standard
- Many language features work, but need WASM codegen implementation
- Project follows Go module structure with `go.mod` in compiler directory
- Test suite has 123 .uni files - run `./test_all_uni_files.sh` to verify changes
- **Current baseline: 56 passing (45.5%)** - see TEST_BASELINE_2026-02-02.md for details