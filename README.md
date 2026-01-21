# Unified Programming Language

[![Status](https://img.shields.io/badge/status-in%20development-yellow)](docs/PROJECT_STATUS.md)
[![Phase](https://img.shields.io/badge/phase-0%20planning-blue)](docs/planning/PHASED_ROADMAP.md)
[![License](https://img.shields.io/badge/license-MIT-green)](LICENSE)

A modern systems programming language that combines memory safety, performance, and developer ergonomics.

## 📋 Quick Links

- **[Project Status](docs/PROJECT_STATUS.md)** - Current implementation status
- **[Phased Roadmap](docs/planning/PHASED_ROADMAP.md)** - Complete implementation plan
- **[Contributing Guide](CONTRIBUTING.md)** - How to contribute
- **[Copilot Guide](docs/COPILOT_GUIDE.md)** - For AI-assisted development
- **[Architecture](docs/design/ARCHITECTURE.md)** - Compiler architecture
- **[Language Spec](spec/UnifiedSpecifiation.md)** - Complete language specification

## 🎯 Current Status

**Phase 0: Project Foundation and Planning** - 90% Complete

The project has comprehensive documentation and planning in place. Phase 1 (Minimal Compiler Pipeline) will begin once Phase 0 is complete. See [PROJECT_STATUS.md](docs/PROJECT_STATUS.md) for details.

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
- Uses LLVM as a backend for wide platform support
- Generates efficient native code with strong optimizations

The language emphasizes memory safety without garbage collection overhead, expressive but pragmatic syntax, and powerful concurrency primitives, aiming to be suitable for systems programming while maintaining developer productivity.

This summary captures the key aspects of Unified we've designed, including its syntax, semantics, memory model, type system, concurrency features, and compilation strategy.

## 🏗️ Development

### Getting Started

**Prerequisites:**
- Go 1.21+
- ANTLR4
- Make

**Build the compiler:**
```bash
cd src/unified-compiler
make all  # Generate parser and build compiler
make test # Run tests
```

**Current capabilities:**
- ✅ Parser infrastructure (ANTLR4 grammar)
- ✅ Basic AST structures
- ✅ Project organization and build system
- ❌ Code generation (blocked by LLVM bindings)
- ❌ Executable programs (Phase 1 goal)

See [CONTRIBUTING.md](CONTRIBUTING.md) for detailed development instructions.

### Phased Development

Unified is being built in 15 carefully planned phases:

- **Phase 0** (Current): Project foundation and planning
- **Phase 1**: Minimal compiler pipeline
- **Phase 2**: Basic expressions and arithmetic
- **Phase 3**: Control flow
- **Phases 4-15**: Functions, types, memory management, concurrency, tooling, and more

Each phase has specific goals, test requirements, and success criteria. See [PHASED_ROADMAP.md](docs/planning/PHASED_ROADMAP.md) for the complete plan.

### For AI Agents

This project is optimized for development with GitHub Copilot and other AI coding assistants:

- Read [COPILOT_GUIDE.md](docs/COPILOT_GUIDE.md) for guidance
- Follow the phased roadmap strictly
- Write tests before implementation
- Document as you code

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