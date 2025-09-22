# Unified Compiler Implementation Plan

## Current State Analysis

### ✅ **What's Working**
1. **Parser Generation**: ANTLR4 grammar files compile and generate Go code
2. **Basic Parsing**: Simple function declarations with basic types parse correctly 
3. **AST Building**: Visitor pattern converts parse trees to AST structures
4. **Project Structure**: Well-organized Go modules and build system

### ❌ **Major Blockers**

#### 1. **LLVM Integration Broken** 
- The LLVM Go bindings (`github.com/llvm/bindings/go/llvm`) don't exist
- Code generation is completely non-functional
- Need to find working LLVM bindings or alternative backends

#### 2. **Grammar Incomplete**
- Missing many essential language features from the `.uni` test files
- Type syntax issues (`Int` vs `i32`)
- Range operators (`..=`) not fully implemented
- Method call syntax incomplete
- For loop syntax issues

#### 3. **AST Visitor Incomplete**
- Many expression types cause panics 
- Binary/unary operator handling incomplete
- Complex expressions not properly parsed

#### 4. **No Runtime System**
- No way to execute compiled code
- Missing print functions, basic I/O
- No standard library bindings

## 🎯 **Roadmap to Working System**

### **Phase 1: Fix Core Parser (1-2 weeks)**
1. **Fix LLVM Bindings**
   - Replace with `tinygo-org/go-llvm` or similar working bindings
   - Or implement a simpler C-based backend initially

2. **Complete Essential Grammar**
   - Add missing operators and expressions 
   - Fix type syntax consistency
   - Add basic for/while loop support
   - Add simple method calls

3. **Fix AST Visitor**
   - Add null checks and error handling
   - Complete binary/unary expression parsing
   - Add missing statement types

### **Phase 2: Minimal Code Generation (1-2 weeks)**  
1. **Basic LLVM Integration**
   - Generate simple functions
   - Handle basic arithmetic
   - Support return statements
   - Add variable declarations/assignments

2. **Minimal Runtime**
   - Link with C runtime for basic I/O
   - Add simple print function
   - Handle program entry point

### **Phase 3: Essential Features (2-3 weeks)**
1. **Control Flow**
   - If/else statements
   - Basic loops (for, while)
   - Function calls

2. **Data Types**
   - Integers, floats, booleans
   - Basic strings
   - Simple arrays/lists

### **Phase 4: Target Programs (1 week)**
1. **Make Test Programs Work**
   - Fix fibonacci program
   - Fix fizzbuzz program
   - Add any missing language features they need

## 🚧 **Immediate Next Steps**

To get started right away:

1. **Fix LLVM bindings** - Replace with `github.com/llir/llvm` or `tinygo-org/go-llvm`
2. **Create minimal test programs** that work with current grammar
3. **Get basic code generation working** for simple arithmetic
4. **Add runtime integration** so programs can actually run

The project has a solid foundation but needs significant work on the code generation and missing language features. The good news is that the parser architecture is sound, so it's mainly about filling in the gaps systematically.

**Estimated timeline to working system: 4-6 weeks** with focused development effort.

## Test Programs to Target

### Simple Test (Currently Working)
```unified
fn main() -> i32 { 
    return 42; 
}
```

### Basic Arithmetic (Target)
```unified
fn add(a: i32, b: i32) -> i32 {
    return a + b;
}

fn main() -> i32 {
    let result = add(5, 7);
    return result;
}
```

### Fibonacci (End Goal)
```unified
fn fibonacci(n: i32) -> i32 {
    if n <= 1 {
        return n;
    }
    return fibonacci(n - 1) + fibonacci(n - 2);
}

fn main() -> i32 {
    return fibonacci(10);
}
```