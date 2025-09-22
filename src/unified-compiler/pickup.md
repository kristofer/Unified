# Unified Compiler Development Session State

## Project Overview
This is the implementation of the **Unified Programming Language** - a modern systems programming language. The compiler is built in Go using ANTLR4 for parsing and LLVM for code generation.

## Current Project Structure
- **Main directory**: `/Users/kristofer/LocalProjects/Unified/src/unified-compiler`
- **Source files**: 
  - `cmd/compiler/main.go` - Main compiler entry point
  - `internal/ast/` - AST definitions and visitor implementation
  - `internal/codegen/` - LLVM IR code generation
  - `internal/parser/` - ANTLR-generated parser code
  - `grammar/` - ANTLR4 grammar files (UnifiedLexer.g4, UnifiedParser.g4)
  - `test/` - Test programs in .uni format

## Build System
- Uses `Makefile` with targets: `parser`, `build`, `test`, `clean`
- `make parser` - Regenerates parser from grammar files using ANTLR4
- `make build` - Compiles the Go binary to `bin/unified-compiler`
- Uses Go modules with `go.mod`

## Current Dependencies
- `github.com/antlr4-go/antlr/v4 v4.13.1` - ANTLR Go runtime
- `tinygo.org/x/go-llvm v0.0.0-20250422114502-b8f170971e74` - LLVM bindings

## LLVM Integration Status
**FIXED**: LLVM bindings issue was resolved by switching from broken `github.com/llvm/bindings/go/llvm` to working `tinygo.org/x/go-llvm`. All LLVM API calls updated to use context-based methods:
- `llvm.Int32Type()` → `g.context.Int32Type()`  
- `llvm.NewBuilder()` → `g.context.NewBuilder()`
- Type equality: `type1.Equal(type2)` → `type1.C == type2.C`
- Load operations require explicit types: `CreateLoad(type, ptr, name)`

## Working Features
The compiler can successfully compile these patterns:

### Simple Functions
```unified
fn main() -> i32 {
    return 42;
}
```

### Functions with Parameters and Arithmetic
```unified
fn add(a: i32, b: i32) -> i32 {
    return a + b;
}
```

**Generated LLVM IR Example**:
```llvm
define i32 @add(i32 %a, i32 %b) {
entry:
  %add = add i32 %a, %b
  ret i32 %add
}
```

## Known Limitations and Issues

### Parser/AST Issues
1. **Function calls not supported** - `multiply(6, 7)` causes panic in AST visitor
2. **Let statements crash** - segfault in LLVM CreateLoad due to alloca type issues
3. **Limited expression parsing** - many expression types return nil, causing panics

### Grammar Limitations  
1. Missing function call syntax in expression grammar
2. Type syntax inconsistencies (`Int` vs `i32`)
3. Range operators (`..=`) not fully implemented
4. Method call syntax incomplete (`List<Int>.new()`)
5. For loop syntax issues

### Codegen Limitations
1. Let statements create allocas but type handling fails
2. Only basic arithmetic operators implemented  
3. No control flow (if/while/for) in codegen
4. Function parameters stored directly (not as allocas) for simplicity

## Test Files Available
- `minimal_test.uni` - Simple return 42 (WORKS)
- `add_test.uni` - Function with parameters (WORKS)  
- `basic_test.uni` - Uses let statements (CRASHES)
- `test/fib.uni` - Complex fibonacci with advanced features (FAILS)
- `test/fizz.uni` - FizzBuzz example (FAILS)

## Compilation Process
1. `./bin/unified-compiler --input file.uni --output file.ll`
2. Parses .uni file using ANTLR4 parser
3. Builds AST using visitor pattern
4. Generates LLVM IR using codegen
5. Writes .ll file with LLVM IR
6. Verifies module before output

## Error Patterns Encountered
- **Segmentation faults** in CreateLoad calls when types are null
- **Panics** in AST visitor when expressions are nil  
- **Module verification failures** when multiple terminators in basic blocks
- **Undefined variable errors** when symbol table lookup fails

## Recent Working Session
Was implementing Phase 1 of development plan focused on fixing core parser and LLVM integration. Successfully fixed LLVM bindings and got basic arithmetic working. Current todo list was focused on expanding language features:

Current active todos (before session end):
1. Document current working features and limitations (in progress)
2. Add function call support to grammar and visitor
3. Fix let statement alloca handling  
4. Add missing binary operators to lexer grammar
5. Add missing comparison operators to parser grammar
6. Fix range operator syntax (..=) in grammar
7. Add method call syntax to expression grammar
8. Fix for loop syntax in grammar
9. Add null checks to AST visitor methods
10. Complete binary expression parsing in visitor

## Key Files Modified Recently
- `internal/codegen/generator.go` - Updated all LLVM API calls for TinyGo compatibility
- `cmd/compiler/main.go` - Toggles between testparse and compile functions
- `go.mod` - Updated to use tinygo.org/x/go-llvm

## Build Commands Used
- `make parser && make build` - Full rebuild
- `./bin/unified-compiler --input test.uni --output test.ll` - Compile test