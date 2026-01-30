# Unified Programming Language

[![Status](https://img.shields.io/badge/status-phase%201%20complete-green)](docs/PROJECT_STATUS.md)
[![Phase](https://img.shields.io/badge/current-phase%202%20ready-blue)](docs/planning/AI_IMPLEMENTATION_PLAN.md)
[![Tests](https://img.shields.io/badge/tests-26%20passing-yellow)](TODO.md)
[![License](https://img.shields.io/badge/license-MIT-green)](LICENSE)

A modern systems programming language that combines memory safety, performance, and developer ergonomics.

## 📋 Quick Links

### For AI Agents & Developers
- **[🤖 AI Implementation Plan](docs/planning/AI_IMPLEMENTATION_PLAN.md)** - **START HERE** for implementation
- **[📊 Project Status](docs/PROJECT_STATUS.md)** - Current status and next steps
- **[📖 Implementation Guide](docs/IMPLEMENTATION_GUIDE.md)** - How to use the documentation
- **[🧪 Testing Guide](src/unified-compiler/TESTING.md)** - How to write and run tests

### Language Documentation
- **[📐 Language Spec](spec/UnifiedSpecifiation.md)** - Complete language specification
- **[🗺️ Phased Roadmap](docs/planning/PHASED_ROADMAP.md)** - High-level implementation plan
- **[🏛️ Architecture](docs/design/ARCHITECTURE.md)** - Compiler architecture
- **[⚡ WASM Backend](WASM_MIGRATION_SUMMARY.md)** - WebAssembly backend documentation

### Getting Started
- **[🤝 Contributing](CONTRIBUTING.md)** - How to contribute
- **[🎓 Learn Unified](LearnUnifiedinYMinutes.md)** - Quick language tour

## 🎯 Current Status

**Phase 1: WebAssembly Backend** - ✅ COMPLETE

The Unified compiler now targets WebAssembly with comprehensive feature support:
- ✅ WebAssembly code generation (replaces custom VM)
- ✅ Struct operations (heap allocation, field access)
- ✅ Enum operations (tagged unions, variant construction)
- ✅ Array operations (literals, indexing with bounds checking)
- ✅ String manipulation (literals in data section, deduplication)
- ✅ For loops (array iteration)
- ✅ Pattern matching (match expressions)
- ✅ Advanced control flow (break, continue)
- ✅ Memory management (bump allocator)
- ✅ 15+ unit tests (100% passing)
- ✅ Complete documentation

**Runtime:** wazero (pure Go WebAssembly runtime)

See [WASM Migration Summary](WASM_MIGRATION_SUMMARY.md) and [WASM Feature Expansion](WASM_FEATURE_EXPANSION_SUMMARY.md) for complete details.

**Test Results:** 26 of 121 tests passing (21.5%) with WASM backend. See [TODO.md](TODO.md) for detailed test results and implementation roadmap.

## 🚀 Overview

# Unified

I'd call it "Unified" - a language that takes the most compelling features from various languages while addressing common pain points.

## Core Design Principles

1. **Pragmatic simplicity** - Like Go and Lua, the language would have a small, consistent core with clear semantics
2. **Memory safety without GC overhead** - Borrowing Rust's ownership model but with more ergonomic syntax
3. **Type flexibility** - Gradual typing like TypeScript that allows both dynamic exploration and static verification
4. **Concurrency by design** - Built-in concurrency primitives from Go and Erlang/Elixir
5. **Expression-oriented** - Everything is an expression (like Rust and functional languages)

## Key Features

```unified
// Semicolons are optional - clean, readable syntax
fn fibonacci(n: Int) -> Int {
    if n <= 1 {
        return n
    }
    let mut a = 0
    let mut b = 1
    let mut result = 0
    for i in 2..=n {
        result = a + b
        a = b
        b = result
    }
    return result
}

// Pattern matching and algebraic data types (from Rust/ML family)
type Result<T, E> = Success(T) | Error(E)

// First-class functions with closures (like JavaScript/Rust)
fn map<T, U>(list: List<T>, transform: fn(T) -> U) -> List<U> {
    // Implementation
}

// Ownership system with ergonomic borrowing (inspired by Rust)
fn process(data: &mut String) {
    // Mutable borrow
}

// Erlang/Elixir-style concurrency with lightweight processes
actor Counter {
    var count = 0
    
    fn increment() {
        count += 1
        return count
    }
}

// Go-style error handling but with Result types
fn read_file(path: String) -> Result<String, IOError> {
    // Implementation
}

// String interpolation (from Swift/JavaScript)
let message = "Hello, ${name}!"
```


Let's see the possibility of creating a language that handles the memory safety/performance tradeoff better than existing options - something with Rust's safety guarantees but Go's approachability and development speed.

## A  Summary of Unified

Here's a summary of the Unified programming language as developed in our conversation:

# Unified Programming Language Overview

## Core Philosophy
Unified is a modern systems programming language that combines memory safety, performance, and developer ergonomics. It draws inspiration from Rust, Go, Swift, and Elixir, focusing on providing Rust's safety guarantees with Go's approachability.

## Key Features

### Memory Management
- Ownership system similar to Rust but with more intuitive syntax
- Explicit `move` keyword for ownership transfers
- Region-based memory management for efficient allocation
- Optional garbage collection with `@gc` annotation

### Type System
- Strong static typing with powerful type inference
- Algebraic data types (enums, structs)
- Generic programming capabilities
- Advanced type features: union types, phantom types, refinement types
- Expression-oriented syntax

### Concurrency
- Actor model for isolation and message-passing
- Async/await for asynchronous programming
- Channels for communication between processes
- Work-stealing scheduler for lightweight processes
- Low-level synchronization primitives (Mutex, Atomic)

### Error Handling
- `Result<T, E>` and `Option<T>` types
- Pattern matching and destructuring
- `?` operator for error propagation
- No exceptions

### Syntax Highlights
- `let` for immutable variables, `let mut` for mutable ones
- `fn` for functions, with expression-based returns
- `||` for closures
- `module` for code organization
- `interface` for shared behavior contracts
- `switch` for pattern matching
- `region` for memory management blocks

### Standard Library Organization
- Modules organized by functionality
- Comprehensive collection types
- I/O, networking, filesystem utilities

## Compilation
- Uses WebAssembly (WASM) as the compilation target
- Powered by wazero - a pure Go WebAssembly runtime
- Generates portable bytecode that runs anywhere
- Zero C dependencies for maximum portability

The language emphasizes memory safety without garbage collection overhead, expressive but pragmatic syntax, and powerful concurrency primitives, aiming to be suitable for systems programming while maintaining developer productivity.

This summary captures the key aspects of Unified we've designed, including its syntax, semantics, memory model, type system, concurrency features, and compilation strategy.

## 🏗️ Development

### Current Capabilities

**Working Features** (Phase 1 Complete):
```unified
fn factorial(n: Int) -> Int {
    let result = 1
    let i = n
    // Note: loops coming in Phase 2!
    return result * i
}

fn main() -> Int {
    return factorial(5)
}
```

**What Works Now**:
- ✅ Function declarations with typed parameters
- ✅ Function calls with arguments
- ✅ Local variables (`let` statements)
- ✅ Return values
- ✅ All arithmetic operations (+, -, *, /, %, unary -)
- ✅ All comparison operations (==, !=, <, <=, >, >=)
- ✅ All logical operations (&&, ||, !)
- ✅ Integer, float, boolean, string, and null literals
- ✅ Complex nested expressions

**Coming in Phase 2**:
- 🔄 If/else statements
- 🔄 While loops
- 🔄 For loops with ranges
- 🔄 Break/continue statements

See the [complete phase plan](docs/planning/AI_IMPLEMENTATION_PLAN.md) for all upcoming features.

## 🚀 Quick Start for Developers

### Prerequisites

Before you begin, ensure you have the following installed:

- **Go 1.21+** - [Download Go](https://golang.org/dl/)
- **Make** - Usually pre-installed on Linux/macOS, or install via package manager
- **ANTLR4** (optional, only needed if modifying grammar files) - [Install ANTLR4](https://www.antlr.org/download.html)

### Setup Development Environment

1. **Clone the repository**
   ```bash
   git clone https://github.com/kristofer/Unified.git
   cd Unified
   ```

2. **Navigate to the compiler directory**
   ```bash
   cd src/unified-compiler
   ```

3. **Build the compiler**
   ```bash
   make build
   # Or build with parser generation (if grammar was modified)
   make all
   ```

4. **Verify the build**
   ```bash
   ./bin/unified --help
   ```

### Running Tests Locally

The Unified compiler uses a WASM backend. To test all .uni files:

#### Run All .uni Test Files
```bash
# From repository root
./test_all_uni_files.sh
```

This tests all 121 .uni files with the WASM compiler and generates:
- `test_working.txt` - List of passing tests
- `test_failing.txt` - List of failing tests  
- `test_results.txt` - Detailed output for all tests

**Current Results:** 26 passing (21.5%), 95 failing (78.5%)

See [TODO.md](TODO.md) for detailed analysis and implementation roadmap.

#### Run Go Unit Tests
```bash
cd src/unified-compiler
make test
```

This runs all unit tests for the compiler components.

#### Run Tests for Specific Components
```bash
# Test the WASM generator
go test ./internal/wasm -v

# Test the AST builder
go test ./internal/ast -v

# Test integration (end-to-end compilation)
go test ./cmd/compiler -v

# Test all packages verbosely
go test -v ./...
```

#### Run Specific Test Functions
```bash
# Run a specific test by name
go test ./internal/wasm -run TestWASMGenerator -v

# Run tests matching a pattern
go test ./internal/vm -run TestVM.* -v
```

#### Test Coverage
```bash
# View coverage summary
go test ./... -cover

# Generate detailed coverage report
go test ./... -coverprofile=coverage.out
go tool cover -html=coverage.out  # Opens in browser
```

#### Watch Mode for Development
```bash
# Run tests on file changes (requires entr or similar tool)
find . -name "*.go" | entr -c go test ./...
```

### Run Example Programs

The compiler comes with example Unified programs you can compile and run:

```bash
# Compile and run a simple program
./bin/unified --input test/integration/simple_return.uni

# Compile a function call example
./bin/unified --input test/integration/function_call.uni

# Compile local variables example
./bin/unified --input test/integration/local_variables.uni
```

### Common Development Tasks

```bash
# Build the compiler
make build

# Run all tests
make test

# Clean build artifacts
make clean

# Regenerate parser (only after grammar changes)
make parser

# Build and run a specific test program
make build && ./bin/unified --input test/integration/simple_return.uni
```

### Troubleshooting

**Issue: Tests fail to run**
- Ensure you're in the `src/unified-compiler` directory
- Run `go mod download` to fetch dependencies
- Verify Go version with `go version` (must be 1.21+)

**Issue: Build fails**
- Check that you have Go installed: `go version`
- Run `make clean` then `make build`

**Issue: Parser generation fails**
- ANTLR4 is only needed if you're modifying grammar files
- For normal development, pre-generated parser files are included

For more detailed testing information, see [Testing Guide](src/unified-compiler/TESTING.md).

### For AI Agents

This project is optimized for development with GitHub Copilot and other AI coding assistants:

1. **Start here**: [TODO.md](TODO.md) - Implementation roadmap with all failing tests
2. **Current status**: WASM backend functional but needs feature completion
3. **Test results**: 26/121 tests passing (21.5%)
4. **Priority**: Fix structs, arrays, for loops, and strings (Priority 1 in TODO.md)
5. **Test continuously**: Run `./test_all_uni_files.sh` to verify progress

See [WASM Migration Summary](WASM_MIGRATION_SUMMARY.md) for architecture details.

### Phased Development

Unified's compiler architecture has evolved:

- **✅ Phase 1**: WASM Backend Implementation (COMPLETE - architecture in place)
- **🎯 Current**: Feature parity with all 121 test files
  - Priority 1: Structs, arrays, for loops, strings (30 tests)
  - Priority 2: Generics, try operator (22 tests)
  - Priority 3: Standard library (24 tests)
- **Future**: Advanced language features (ownership, concurrency, tooling)

The original 15-phase plan has been superseded by the WASM migration. The new implementation roadmap focuses on achieving feature completeness with the WASM backend.

See [TODO.md](TODO.md) for the complete implementation roadmap with specific tasks for each failing test.

### For AI Agents - Quick Start

**Ready to implement?** Here's your workflow:

1. Read [TODO.md](TODO.md) for the complete implementation roadmap
2. Focus on Priority 1 items first (structs, arrays, for loops, strings)
3. Test files show expected behavior - use them as specification
4. Run `./test_all_uni_files.sh` to verify each change
5. Use `report_progress` tool to commit changes
6. See [WASM Migration Summary](WASM_MIGRATION_SUMMARY.md) for architecture details

### Reference Project

This project follows best practices from [kristofer/smog](https://github.com/kristofer/smog), an excellent example of:
- Well-structured documentation
- Phased implementation
- Comprehensive testing
- Clear architecture

# Implementation Notes

26 Mar, Asked claude to create an ANTLR4 grammar file. `unified.g4`
Will try to use ANTLR to gen parser, and ask claude to help build the middle-end of the parser from `.uni` code to LLVM IR.

3 Apr 2025, how about making the backend to Unified be WASM instead of LLVM? 
Asking claude this was the reply:

Targeting WebAssembly (WASM) instead of LLVM as the backend for Unified is a very strategic pivot that makes a lot of sense given the circumstances. Here's why this could be a good direction:

### Benefits of Targeting WASM

1. **Better Go tooling integration**: Go has first-class support for WebAssembly compilation. You could use Go's established WASM tools rather than relying on unmaintained LLVM bindings.

2. **Platform independence**: WASM is designed to be portable and runs in browsers, standalone runtimes (Wasmtime, WAMR, Wasmer), and can be embedded in many applications.

3. **Growing ecosystem**: WASM is gaining tremendous momentum with WASI (WebAssembly System Interface) expanding its capabilities beyond the browser.

4. **Modern specification**: Unlike LLVM which is primarily a compiler infrastructure, WASM is designed as a compilation target from the ground up.

5. **Security model**: WASM's sandboxed execution model provides security guarantees that align well with Unified's safety goals.

6. **Simpler implementation path**: The WASM specification is more constrained than LLVM's full feature set, which could make implementation more straightforward.

### Implementation Approach

You could structure the Unified compiler to:

1. Parse Unified code into an AST
2. Perform semantic analysis, type checking, and apply Unified's ownership rules
3. Lower to a WASM-friendly intermediate representation
4. Generate WASM binary format or text format (WAT)

The Go ecosystem has good support for this via packages like `github.com/bytecodealliance/wasmtime-go` for runtime and testing.

### Considerations

- **Memory model**: You'd need to map Unified's ownership system and region-based memory to WASM's linear memory model
- **Concurrency**: WASM is single-threaded in its core spec, though thread proposals exist. You'd need to implement Unified's actor model on top of this.
- **Host interop**: Consider how Unified programs would interact with their host environment through WASI or custom interfaces

### Verdict

Switching from LLVM to WASM as a compilation target is an excellent strategic choice, especially given the Go ecosystem's better support for WASM compared to LLVM. It provides a more modern, portable runtime target while potentially simplifying your implementation path.

This approach would also make Unified programs more deployable across different environments - from servers to browsers to edge computing - which could significantly increase adoption potential.