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

The compiler currently supports (verified via Go unit tests - ~24% of 123 tests passing):

**✅ Working (100% functional)**:
- Function declarations, calls, parameters, and returns
- Variable declarations (let) and mutable variables (let mut)
- Arithmetic, logical, comparison, and bitwise expressions
- Control flow (if/else statements, while loops)
- Assignment and compound assignment operators (+=, -=, etc.)
- Literal values (integers, floats, booleans)
- Struct declarations, instantiation, and field access
- Enum declarations with data and variant construction
- Optional semicolons
- Complex programs (FizzBuzz works correctly)

**🟡 Partially Working (needs bug fixes)**:
- For loops (infrastructure exists but produces wrong values - off-by-one bugs)
- Nested struct field access (type mismatch: i32 vs i64 issues)

**❌ Not Yet Implemented**:
- Array operations (literals parse, but allocation/indexing broken)
- String operations (literals work, but length/concat/indexing broken)
- Generic monomorphization (syntax parses, no type instantiation)
- Try operator (?) - parser complete, codegen missing
- Pattern matching (match expressions)
- Methods and Self keyword
- Ownership and borrowing
- Actors and concurrency

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

**Current Test Results:** ~30 passing (24%), ~93 failing (76%)

**⚠️ IMPORTANT**: The `test_all_uni_files.sh` script is currently broken on macOS because it uses the `timeout` command which doesn't exist. Install coreutils (`brew install coreutils`) or use Go unit tests instead:

```bash
cd src/unified-compiler
go test ./... -v  # Most reliable way to test
```

See [PROJECT_STATUS_2026-02-02.md](PROJECT_STATUS_2026-02-02.md) for comprehensive current status and [ROADMAP_2026.md](ROADMAP_2026.md) for implementation plan.

## Implementation Status

**Current Phase:** Bug Fixes & Priority 1 Features
**Last Reviewed:** February 2, 2026
**Status:** Active Development

The compiler has successfully migrated from a custom VM to WebAssembly. The WASM backend is complete and functional. Core language features work well (70% of integration tests passing), but several critical bugs are blocking ~30 tests.

**Immediate Priorities** (see [ROADMAP_2026.md](ROADMAP_2026.md) for complete plan):

### Priority 0: URGENT (Week 1)
1. **Fix test infrastructure** - `timeout` command missing on macOS (15 min)
2. **Run comprehensive test audit** - Get accurate baseline (30 min)
3. **Update documentation** - Reflect actual current state (30 min)

### Priority 1: Critical Bugs (Weeks 1-3)
1. **Fix for loop bugs** - Produces wrong values (off-by-one errors) - 2-4 hours
2. **Resolve type system issues** - i32 vs i64 mismatches throughout - 1-3 days
3. **Fix array support** - Allocation and indexing broken - 1-2 days
4. **Fix string operations** - Type mismatches, missing runtime functions - 1-2 days

**Expected Result after Priority 1**: 50-60 tests passing (50% pass rate)

### Priority 2: High-Value Features (Weeks 4-6)
1. **Complete try operator** - Parser done, just need codegen - 6-12 hours
2. **Improve struct operations** - After type system fixed - 6-8 hours

**Expected Result after Priority 2**: 60-70 tests passing (60% pass rate)

### Priority 3: Advanced Features (Weeks 7-12)
1. **Generic monomorphization** - Type instantiation and specialization - 2-3 weeks
2. **Methods and Self keyword** - Enable OOP patterns and stdlib - 1-2 weeks
3. **Pattern matching** - Match expressions with exhaustiveness checking - 2-3 weeks

**Expected Result after Priority 3**: 80-90 tests passing (70-75% pass rate)

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