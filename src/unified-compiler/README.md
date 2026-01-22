# Unified Compiler

A compiler for the Unified language with a custom VM-based backend written in Go.

## Overview

The Unified compiler uses ANTLR4 for parsing and implements a custom virtual machine for code execution. This provides excellent portability while maintaining good performance.

## Architecture

```
Source Code (.uni) → Lexer → Parser → AST → Bytecode Generator → Bytecode → VM → Execution
```

See [VM_README.md](VM_README.md) for details on the VM architecture.

## Project Structure

- `cmd/compiler`: Command-line interface and integration tests
- `grammar`: ANTLR grammar definition
- `internal/ast`: Abstract Syntax Tree definitions and visitor
- `internal/parser`: ANTLR-generated parser
- `internal/bytecode`: Bytecode generator and instruction set
- `internal/vm`: Virtual machine execution engine
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

The compiler has comprehensive test coverage with 76+ test cases.

```bash
# Run all tests
go test ./...

# Run with verbose output
go test ./... -v

# Run with coverage
go test ./... -cover
```

See [TESTING.md](TESTING.md) for the complete testing guide.

## Phase 1 Features

Currently implemented:
- ✅ Function declarations and calls
- ✅ Function parameters
- ✅ Arithmetic operations (+, -, *, /, %)
- ✅ Comparison operations (==, !=, <, <=, >, >=)
- ✅ Logical operations (&&, ||, !)
- ✅ Local variables (let statements)
- ✅ Return statements
- ✅ Integer, float, boolean, and string literals
- ✅ Basic expressions
- ✅ Control flow (if/else - VM support complete)

## Documentation

- [VM_README.md](VM_README.md) - Virtual machine architecture and bytecode instruction set
- [TESTING.md](TESTING.md) - Comprehensive testing guide
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
