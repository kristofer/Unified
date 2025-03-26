# Unified
Claude and I designed a language!
I'd love to design a programming language! This is a fascinating thought experiment.

If I were to create a new programming language that unifies concepts across imperative languages while maintaining their strengths, here's what I'd envision:

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

## Implementation Notes

26 Mar, Asked claude to create an ANTLR4 grammar file. `unified.g4`
Will try to use ANTLR to gen parser, and ask claude to help build the middle-end of the parser from `.uni` code to LLVM IR.
