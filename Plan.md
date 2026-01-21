# Unified Compiler Implementation Plan

## Current State Analysis

### ✅ **What's Working**
1. **Parser Generation**: ANTLR4 grammar files compile and generate Go code
2. **Basic Parsing**: Simple function declarations with basic types parse correctly 
3. **AST Building**: Visitor pattern converts parse trees to AST structures
4. **Project Structure**: Well-organized Go modules and build system

### 🔄 **Architecture Change**

**Decision**: Switch from LLVM-based code generation to a custom Go-based Virtual Machine (VM)

**Rationale**:
- **Portability**: VM works anywhere Go runs (no LLVM dependency)
- **Simplicity**: Easier to implement and maintain than LLVM bindings
- **Faster Development**: No need to deal with LLVM binding issues
- **Better for Iteration**: Simpler debugging and testing
- **Still Good Performance**: VM performance is competitive with interpreted languages

### ❌ **Previous Blockers (Now Addressed)**

#### 1. **LLVM Integration Issues** (RESOLVED)
- ~~The LLVM Go bindings don't exist or are poorly maintained~~
- **Solution**: Use custom VM instead of LLVM
- No external dependencies needed

## 🎯 **Roadmap to Working System**

### **Phase 1: VM-Based Minimal Compiler (2-3 weeks)**
1. **Design VM Architecture**
   - Define bytecode instruction set
   - Design stack-based VM structure
   - Plan runtime data structures

2. **Implement Bytecode Generator**
   - Convert AST to bytecode instructions
   - Handle basic expressions and statements
   - Support function definitions and calls

3. **Implement Virtual Machine**
   - Create stack-based execution engine
   - Implement instruction execution
   - Add runtime support (I/O, etc.)

4. **Test and Validate**
   - Ensure minimal programs work
   - Test arithmetic and control flow
   - Validate against test suite

#### 2. **Grammar Needs Completion**
- Missing many essential language features from the `.uni` test files
- Type syntax issues (`Int` vs `i32`)
- Range operators (`..=`) not fully implemented
- Method call syntax incomplete
- For loop syntax needs work

#### 3. **AST Visitor Needs Improvement**
- Some expression types cause issues
- Binary/unary operator handling needs completion
- Complex expressions need proper parsing

## 🚧 **Immediate Next Steps**

To get started right away:

1. **Design VM bytecode instruction set** - Define simple, stack-based instructions
2. **Implement bytecode generator** - Convert AST to bytecode
3. **Create virtual machine** - Stack-based execution engine
4. **Test with minimal program** - Get `fn main() -> i32 { return 42 }` working

The project has a solid foundation with the parser working. Now we need to implement the VM backend to make programs executable. The VM approach is much simpler than LLVM and will let us iterate quickly.

**Estimated timeline to working system: 2-3 weeks** with focused development effort on the VM.

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