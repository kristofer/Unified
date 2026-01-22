# Examples Directory

This directory contains example programs written in the Unified programming language. Examples are organized by complexity and use case.

## Directory Structure

```
examples/
├── basic/          # Simple examples for beginners
├── advanced/       # Complex examples showcasing advanced features
└── tutorials/      # Step-by-step tutorial examples
```

## Current Status

The examples directory structure is in place but examples are not yet implemented. Examples will be added as language features are implemented in each phase.

## Planned Examples

### Basic Examples (Phase 1-3)

- `hello_world.uni` - Print "Hello, World!"
- `fibonacci.uni` - Calculate Fibonacci numbers
- `fizzbuzz.uni` - FizzBuzz implementation
- `factorial.uni` - Factorial calculation
- `primes.uni` - Prime number generation

### Advanced Examples (Phase 4-8)

- `binary_tree.uni` - Binary tree implementation
- `hash_map.uni` - Hash map implementation
- `json_parser.uni` - JSON parsing
- `web_server.uni` - Simple HTTP server
- `actor_system.uni` - Actor-based concurrency

### Tutorial Examples (Ongoing)

Step-by-step tutorials covering:
- Getting started with Unified
- Understanding ownership and borrowing
- Working with pattern matching
- Concurrent programming with actors
- Building a CLI application
- Building a web application

## Contributing Examples

When contributing examples:

1. Place in appropriate directory (basic/advanced/tutorials)
2. Add clear comments explaining the code
3. Include expected output in comments
4. Follow the code style guidelines
5. Test that the example compiles and runs correctly

## Testing Examples

Examples should be tested as part of the integration test suite:

```bash
cd src/unified-compiler
./bin/unified --input ../../examples/basic/hello_world.uni
```

## Documentation

Each example should include:
- Purpose of the example
- What language features it demonstrates
- Expected output
- Difficulty level
- Prerequisites

Example header:
```unified
// Example: Fibonacci Sequence
// Purpose: Demonstrates recursion and basic functions
// Features: Functions, recursion, control flow
// Level: Basic
// Expected output: First 10 Fibonacci numbers

// Implementation...
```

## Status

Examples will be added progressively as each phase is completed. Check the [PROJECT_STATUS.md](../docs/PROJECT_STATUS.md) to see which phases are complete.
