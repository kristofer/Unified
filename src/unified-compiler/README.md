# Unified Compiler

A compiler for the Unified language with a WebAssembly backend written in Go.

## Overview

The Unified compiler uses ANTLR4 for parsing and generates WebAssembly bytecode for execution. The WASM runtime is provided by wazero, a pure Go implementation with zero C dependencies.

## Architecture

```
Source Code (.uni) → Lexer → Parser → AST → WASM Generator → WASM Binary → wazero Runtime → Execution
```

See [WASM_MIGRATION_SUMMARY.md](../../WASM_MIGRATION_SUMMARY.md) for details on the WASM backend architecture.

## Project Structure

- `cmd/compiler`: Command-line interface and integration tests
- `grammar`: ANTLR grammar definition
- `internal/ast`: Abstract Syntax Tree definitions and visitor
- `internal/parser`: ANTLR-generated parser
- `internal/wasm`: WebAssembly code generator and encoder
- `internal/bytecode`: Type definitions (shared with WASM)
- `internal/semantic`: Semantic analysis (symbol table, type inference, checker)
- `internal/vm`: Legacy VM code (deprecated, replaced by WASM)
- `scripts`: Build and utility scripts
- `test`: Test cases and integration test files

## Building

1. Ensure you have Go 1.22+ and ANTLR4 installed
2. Run `make parser` to generate the parser from grammar
3. Run `make build` to build the compiler

```bash
cd src/unified-compiler
make parser
make build
```

## Usage

```bash
./bin/unified --input program.uni
# The exit code will be the return value of main()
echo $?
```

## Testing

The compiler is tested with 121 .uni test files across the repository.

**Current Status: 26 tests passing (21.5%)**

```bash
# Run all .uni test files (from repository root)
cd ../..
./test_all_uni_files.sh

# Run Go unit tests (from compiler directory)
go test ./...

# Run with verbose output
go test ./... -v

# Run with coverage
go test ./... -cover

# Run specific package tests
go test ./internal/wasm -v
go test ./internal/ast -v
```

See [../../TODO.md](../../TODO.md) for detailed test results and implementation roadmap.

## Current Features

### Working Features (26 tests passing) ✅

**Phase 1 - Core Language Features:**
- ✅ Function declarations and calls
- ✅ Function parameters and return values
- ✅ Arithmetic operations (+, -, *, /, %)
- ✅ Comparison operations (==, !=, <, <=, >, >=)
- ✅ Logical operations (&&, ||, !)
- ✅ Bitwise operations (&, |, ^, ~, <<, >>)
- ✅ Local variables (let statements)
- ✅ Mutable variables (let mut)
- ✅ Variable assignment
- ✅ Compound assignment (+=, -=, *=, /=, %=)
- ✅ If/else statements
- ✅ While loops
- ✅ Return statements
- ✅ Integer, float, boolean literals
- ✅ Basic expressions and precedence
- ✅ Optional semicolons
- ✅ Type inference for basic types
- ✅ Simple enums
- ✅ Basic generics

### Features In Progress (95 tests failing)

See [../../TODO.md](../../TODO.md) for the complete list of features being implemented:

**Priority 1 (Critical):**
- 🔴 Struct support (heap allocation, field access)
- 🔴 Array operations (literals, indexing, iteration)
- 🔴 For loops with ranges
- 🔴 String operations (length, concat, etc.)

**Priority 2 (Important):**
- 🟡 Advanced generics and type inference
- 🟡 Try operator (?) for error handling
- 🟡 Nested loop control flow

**Priority 3 (Nice to have):**
- 🟢 Standard library collections
- 🟢 Block expressions
- 🟢 Variable shadowing
- ✅ Symbol table with scope management
- ✅ Semantic analysis (mutability, undefined variables, type checking)
- ✅ Clear error messages for violations

### Phase 4 - Advanced Expressions ✅
- ✅ Bitwise operations (&, |, ^, ~, <<, >>)
- ✅ Operator precedence (correct evaluation order)
- ✅ Block expressions (blocks as values)
- ⏳ Tuple support (infrastructure ready)
- ⏳ Lambda expressions (infrastructure ready)
- ⏳ Default parameters (not started)

### Phase 5 - Structs and Basic Types 🚧
- ✅ Struct declarations with fields
- ✅ Struct instantiation (requires parentheses: `(Point { x: 10 })` )
- ✅ Field access (dot notation)
- ✅ Nested structs
- ❌ Methods on structs (blocked: parser regeneration needed)
- ❌ Associated functions (blocked: parser regeneration needed)
- ⏳ Mutable field assignment (OpStoreField exists, syntax needed)

## Known Issues

### Syntax Workarounds

1. **Range Operator Spacing**: Range operators must have spaces around them
   - ✅ Correct: `for i in 1 .. 4` or `for i in 1 ..= 10`
   - ❌ Incorrect: `for i in 1..4`
   - **Reason**: Lexer ambiguity between `1.` (float) and `1` + `..` (range)
   - **Fix**: Requires ANTLR 4.13+ to regenerate parser

2. **Struct Instantiation**: Must wrap in parentheses
   - ✅ Correct: `let p = (Point { x: 10, y: 20 });`
   - ❌ Incorrect: `let p = Point { x: 10, y: 20 };`
   - **Reason**: Parser precedence ambiguity
   - **Fix**: Requires grammar restructuring and parser regeneration

See [OPTIMIZATION_REVIEW.md](OPTIMIZATION_REVIEW.md) for detailed analysis and future fixes.

## Documentation

- [VM_README.md](VM_README.md) - Virtual machine architecture and bytecode instruction set
- [TESTING.md](TESTING.md) - Comprehensive testing guide  
- [TEST_RESULTS_2026-01-26.md](TEST_RESULTS_2026-01-26.md) - Detailed test results and analysis
- [OPTIMIZATION_REVIEW.md](OPTIMIZATION_REVIEW.md) - Future optimization and review tasks
- [ARCHITECTURE_CHANGE_SUMMARY.md](../../ARCHITECTURE_CHANGE_SUMMARY.md) - Migration from LLVM to VM

## Development

### Running Examples

```bash
# Simple return
./bin/unified --input test/integration/simple_return.uni

# Function calls
./bin/unified --input test/integration/function_call.uni

# Local variables
./bin/unified --input test/integration/local_variables.uni
```

### Make Targets

```bash
make parser    # Generate parser from ANTLR grammar
make build     # Build the compiler
make test      # Run all tests
make clean     # Clean build artifacts
make vm-info   # Display VM architecture information
```
