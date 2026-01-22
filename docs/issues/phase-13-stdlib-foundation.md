# Phase 13: Standard Library Foundation

**Status**: Not Started  
**Duration**: 2-3 weeks  
**Priority**: HIGH  
**Dependencies**: Phase 12 Complete (Basic Ownership)

## 🎯 Goals

Implement the foundational standard library providing essential data structures, I/O primitives, and string utilities. This establishes the core building blocks for all Unified programs.

## 📋 Prerequisites

- [x] Phase 12 complete (ownership and borrowing working)
- [ ] Understanding of standard library design patterns
- [ ] Knowledge of collection implementations
- [ ] Familiarity with I/O abstractions
- [ ] Understanding of FFI for calling C libraries (for I/O)

## ✨ Language Features to Add

### 1. Collections Module
- **List<T>**: Dynamic array with ownership
- **Map<K, V>**: Hash map for key-value storage
- **Set<T>**: Hash set for unique values
- Generic implementations
- Iterator support

### 2. I/O Primitives
- **print()**: Output to stdout
- **println()**: Output with newline
- **read_line()**: Read from stdin
- **File**: Basic file operations (read, write)
- **Error handling** for I/O operations

### 3. String Utilities
- **String**: Owned, growable string type
- **str**: String slice (borrowed string)
- Common string operations (len, push, concat, split, trim)
- UTF-8 support
- String formatting

### 4. Core Types
- **Option<T>**: Nullable value abstraction
- **Result<T, E>**: Error handling
- **Box<T>**: Heap allocation
- Iterator trait

## 📝 Implementation Tasks

### Task 13.1: Design Standard Library Structure (2-3 hours)
- [ ] Create `stdlib/` directory structure
- [ ] Plan module organization
- [ ] Design public API surface
- [ ] Document design decisions
- [ ] Create stdlib build integration

**Directory Structure:**
```
stdlib/
├── core/
│   ├── option.uni
│   ├── result.uni
│   └── iterator.uni
├── collections/
│   ├── list.uni
│   ├── map.uni
│   └── set.uni
├── io/
│   ├── print.uni
│   └── file.uni
├── string/
│   ├── string.uni
│   └── str.uni
└── prelude.uni  // Auto-imported items
```

### Task 13.2: Implement Option<T> (3-4 hours)
- [ ] Create `stdlib/core/option.uni`
- [ ] Define Option enum (Some, None)
- [ ] Implement common methods (is_some, is_none, unwrap, map, etc.)
- [ ] Add pattern matching support
- [ ] Write unit tests
- [ ] Document with examples

**Implementation:**
```unified
pub enum Option<T> {
    Some(T),
    None,
}

impl<T> Option<T> {
    pub fn is_some(&self) -> Bool
    pub fn is_none(&self) -> Bool
    pub fn unwrap(self) -> T
    pub fn unwrap_or(self, default: T) -> T
    pub fn map<U>(self, f: fn(T) -> U) -> Option<U>
}
```

### Task 13.3: Implement Result<T, E> (3-4 hours)
- [ ] Create `stdlib/core/result.uni`
- [ ] Define Result enum (Ok, Err)
- [ ] Implement common methods
- [ ] Add error propagation support
- [ ] Write unit tests
- [ ] Document with examples

**Implementation:**
```unified
pub enum Result<T, E> {
    Ok(T),
    Err(E),
}

impl<T, E> Result<T, E> {
    pub fn is_ok(&self) -> Bool
    pub fn is_err(&self) -> Bool
    pub fn unwrap(self) -> T
    pub fn expect(self, msg: &str) -> T
    pub fn map<U>(self, f: fn(T) -> U) -> Result<U, E>
}
```

### Task 13.4: Implement String and str (6-8 hours)
- [ ] Create `stdlib/string/string.uni`
- [ ] Implement String (owned)
- [ ] Implement str slice (borrowed)
- [ ] Add common string methods
- [ ] Implement UTF-8 handling
- [ ] Add string formatting basics
- [ ] Write comprehensive tests
- [ ] Document API

**Key Methods:**
```unified
pub struct String {
    // Internal implementation
}

impl String {
    pub fn new() -> String
    pub fn from(s: &str) -> String
    pub fn len(&self) -> Int
    pub fn push_str(&mut self, s: &str)
    pub fn as_str(&self) -> &str
    pub fn split(&self, delim: &str) -> List<&str>
    pub fn trim(&self) -> &str
}
```

### Task 13.5: Implement List<T> (6-8 hours)
- [ ] Create `stdlib/collections/list.uni`
- [ ] Implement dynamic array
- [ ] Add capacity management
- [ ] Implement common operations (push, pop, get, set, len)
- [ ] Add iterator support
- [ ] Handle ownership correctly
- [ ] Write comprehensive tests
- [ ] Document with examples

**Implementation:**
```unified
pub struct List<T> {
    // Internal: pointer, length, capacity
}

impl<T> List<T> {
    pub fn new() -> List<T>
    pub fn with_capacity(capacity: Int) -> List<T>
    pub fn push(&mut self, value: T)
    pub fn pop(&mut self) -> Option<T>
    pub fn get(&self, index: Int) -> Option<&T>
    pub fn len(&self) -> Int
    pub fn is_empty(&self) -> Bool
    pub fn clear(&mut self)
}
```

### Task 13.6: Implement Map<K, V> (8-10 hours)
- [ ] Create `stdlib/collections/map.uni`
- [ ] Implement hash map (simple hash table)
- [ ] Implement hashing for basic types
- [ ] Add insert, get, remove, contains_key
- [ ] Handle collisions
- [ ] Implement iterator over entries
- [ ] Write comprehensive tests
- [ ] Document API

**Implementation:**
```unified
pub struct Map<K, V> {
    // Internal: buckets, size, capacity
}

impl<K, V> Map<K, V> {
    pub fn new() -> Map<K, V>
    pub fn insert(&mut self, key: K, value: V) -> Option<V>
    pub fn get(&self, key: &K) -> Option<&V>
    pub fn remove(&mut self, key: &K) -> Option<V>
    pub fn contains_key(&self, key: &K) -> Bool
    pub fn len(&self) -> Int
}
```

### Task 13.7: Implement Set<T> (4-5 hours)
- [ ] Create `stdlib/collections/set.uni`
- [ ] Implement hash set (wrapper around Map)
- [ ] Add insert, remove, contains
- [ ] Implement set operations (union, intersection, difference)
- [ ] Write tests
- [ ] Document API

### Task 13.8: Implement I/O Primitives (6-8 hours)
- [ ] Create `stdlib/io/print.uni`
- [ ] Implement print() via FFI to C
- [ ] Implement println()
- [ ] Implement read_line()
- [ ] Add basic file operations
- [ ] Handle I/O errors with Result
- [ ] Write tests
- [ ] Document I/O API

**Implementation:**
```unified
pub fn print(s: &str)
pub fn println(s: &str)
pub fn read_line() -> Result<String, IoError>

pub struct File {
    // Internal file handle
}

impl File {
    pub fn open(path: &str) -> Result<File, IoError>
    pub fn read_to_string(&mut self) -> Result<String, IoError>
    pub fn write(&mut self, s: &str) -> Result<(), IoError>
}
```

### Task 13.9: Implement Box<T> (4-5 hours)
- [ ] Create `stdlib/core/box.uni`
- [ ] Implement heap allocation wrapper
- [ ] Add deref support
- [ ] Ensure proper cleanup
- [ ] Write tests
- [ ] Document usage

**Implementation:**
```unified
pub struct Box<T> {
    // Internal: pointer to heap
}

impl<T> Box<T> {
    pub fn new(value: T) -> Box<T>
    // Deref implemented via language support
}
```

### Task 13.10: Add FFI Support for C Interop (5-6 hours)
- [ ] Create `internal/ffi/` package
- [ ] Implement C function calling
- [ ] Add C type marshalling
- [ ] Support for linking C libraries
- [ ] Use for print/I/O implementation
- [ ] Write tests
- [ ] Document FFI usage

### Task 13.11: Implement Prelude (2-3 hours)
- [ ] Create `stdlib/prelude.uni`
- [ ] Re-export commonly used types
- [ ] Auto-import prelude in all modules
- [ ] Document prelude contents

**Prelude Contents:**
```unified
pub use core::option::Option::{self, Some, None}
pub use core::result::Result::{self, Ok, Err}
pub use collections::list::List
pub use collections::map::Map
pub use string::String
pub use io::{print, println}
```

### Task 13.12: Update Compiler for stdlib (4-5 hours)
- [ ] Update compiler to find stdlib
- [ ] Auto-import prelude
- [ ] Link stdlib during compilation
- [ ] Add stdlib path configuration
- [ ] Test stdlib integration

### Task 13.13: Write Comprehensive Tests (8-10 hours)
- [ ] Unit tests for each stdlib module (30+ tests)
- [ ] Integration tests using stdlib (10+ tests)
- [ ] Property-based tests for collections
- [ ] I/O tests with temp files
- [ ] Memory safety tests
- [ ] Performance benchmarks

### Task 13.14: Update Documentation (4-5 hours)
- [ ] Create PHASE13_SUMMARY.md
- [ ] Create STDLIB_GUIDE.md
- [ ] Document each module with examples
- [ ] Create stdlib API reference
- [ ] Update README.md
- [ ] Add cookbook-style examples

## ✅ Test Requirements

**Minimum Test Coverage**: 90% for stdlib code

### Option<T> Tests
- [ ] Create Some variant
- [ ] Create None variant
- [ ] Test is_some()
- [ ] Test is_none()
- [ ] Test unwrap() on Some
- [ ] Test unwrap() panics on None
- [ ] Test unwrap_or()
- [ ] Test map()
- [ ] Test pattern matching

### Result<T, E> Tests
- [ ] Create Ok variant
- [ ] Create Err variant
- [ ] Test is_ok()
- [ ] Test is_err()
- [ ] Test unwrap() on Ok
- [ ] Test unwrap() panics on Err
- [ ] Test expect()
- [ ] Test map()

### String Tests
- [ ] Create empty string
- [ ] Create from string literal
- [ ] Test len()
- [ ] Test push_str()
- [ ] Test concatenation
- [ ] Test split()
- [ ] Test trim()
- [ ] Test UTF-8 handling
- [ ] Test ownership and borrowing

### List<T> Tests
- [ ] Create empty list
- [ ] Push elements
- [ ] Pop elements
- [ ] Get by index
- [ ] Test bounds checking
- [ ] Test capacity growth
- [ ] Test clear()
- [ ] Test with different types
- [ ] Test ownership of elements

### Map<K, V> Tests
- [ ] Create empty map
- [ ] Insert key-value pairs
- [ ] Get values by key
- [ ] Test key not found
- [ ] Update existing key
- [ ] Remove keys
- [ ] Test collision handling
- [ ] Test with different types
- [ ] Test iteration

### Set<T> Tests
- [ ] Create empty set
- [ ] Insert elements
- [ ] Test contains()
- [ ] Test uniqueness
- [ ] Test set operations
- [ ] Test with different types

### I/O Tests
- [ ] Test print()
- [ ] Test println()
- [ ] Test read_line() (with input simulation)
- [ ] Test file open
- [ ] Test file read
- [ ] Test file write
- [ ] Test I/O error handling
- [ ] Test file cleanup

### Integration Tests
- [ ] Build program using List
- [ ] Build program using Map
- [ ] Build program using I/O
- [ ] Build program with error handling
- [ ] Build program combining multiple stdlib features
- [ ] Test prelude auto-import
- [ ] Test stdlib linking
- [ ] Verify no memory leaks
- [ ] Test cross-module usage
- [ ] Build realistic example programs

## 📚 Documentation Updates Required

### Update Existing Docs
- [ ] README.md: Add stdlib to feature list
- [ ] TESTING.md: Document stdlib testing approach

### Create New Docs
- [ ] PHASE13_SUMMARY.md: Complete phase summary
- [ ] docs/STDLIB_GUIDE.md: Comprehensive stdlib guide
- [ ] docs/STDLIB_API.md: API reference
- [ ] docs/COLLECTIONS_GUIDE.md: Collections usage
- [ ] docs/IO_GUIDE.md: I/O operations guide
- [ ] docs/ERROR_HANDLING.md: Option and Result guide

### Add Example Programs
- [ ] `examples/collections/list_demo.uni`
- [ ] `examples/collections/map_demo.uni`
- [ ] `examples/collections/set_demo.uni`
- [ ] `examples/io/file_operations.uni`
- [ ] `examples/io/user_input.uni`
- [ ] `examples/strings/string_processing.uni`
- [ ] `examples/error_handling/result_demo.uni`
- [ ] `examples/cookbook/word_count.uni`

## 🎓 Example Programs

### Using List<T>
```unified
use collections::List

fn main() -> Int {
    let mut numbers = List::new()
    numbers.push(1)
    numbers.push(2)
    numbers.push(3)
    
    for num in numbers {
        println(num)
    }
    
    return 0
}
```

### Using Map<K, V>
```unified
use collections::Map

fn main() -> Int {
    let mut scores = Map::new()
    scores.insert("Alice", 100)
    scores.insert("Bob", 95)
    
    match scores.get("Alice") {
        Some(score) => println(score),
        None => println("Not found"),
    }
    
    return 0
}
```

### Error Handling with Result
```unified
use io::File

fn read_config() -> Result<String, IoError> {
    let mut file = File::open("config.txt")?
    return file.read_to_string()
}

fn main() -> Int {
    match read_config() {
        Ok(content) => println(content),
        Err(e) => println("Error reading config"),
    }
    return 0
}
```

### String Processing
```unified
fn main() -> Int {
    let text = String::from("  hello world  ")
    let trimmed = text.trim()
    let words = trimmed.split(" ")
    
    for word in words {
        println(word)
    }
    
    return 0
}
```

### Word Count Program
```unified
use collections::Map
use io::File

fn word_count(filename: &str) -> Result<Map<String, Int>, IoError> {
    let mut file = File::open(filename)?
    let content = file.read_to_string()?
    
    let mut counts = Map::new()
    let words = content.split(" ")
    
    for word in words {
        let count = counts.get(word).unwrap_or(0)
        counts.insert(word.to_string(), count + 1)
    }
    
    return Ok(counts)
}

fn main() -> Int {
    match word_count("input.txt") {
        Ok(counts) => {
            for (word, count) in counts {
                println(word + ": " + count.to_string())
            }
        },
        Err(_) => println("Error reading file"),
    }
    return 0
}
```

### Stack-Based Calculator
```unified
use collections::List

fn evaluate_rpn(expression: &str) -> Option<Int> {
    let mut stack = List::new()
    let tokens = expression.split(" ")
    
    for token in tokens {
        match token {
            "+" => {
                let b = stack.pop()?
                let a = stack.pop()?
                stack.push(a + b)
            },
            "-" => {
                let b = stack.pop()?
                let a = stack.pop()?
                stack.push(a - b)
            },
            _ => {
                let num = token.parse_int()?
                stack.push(num)
            },
        }
    }
    
    return stack.pop()
}

fn main() -> Int {
    match evaluate_rpn("3 4 + 2 *") {
        Some(result) => {
            println(result)  // Prints 14
            return 0
        },
        None => return 1,
    }
}
```

## 🏆 Success Criteria

- [ ] All Option tests pass (minimum 9 tests)
- [ ] All Result tests pass (minimum 8 tests)
- [ ] All String tests pass (minimum 9 tests)
- [ ] All List tests pass (minimum 9 tests)
- [ ] All Map tests pass (minimum 9 tests)
- [ ] All Set tests pass (minimum 6 tests)
- [ ] All I/O tests pass (minimum 8 tests)
- [ ] All integration tests pass (minimum 10 tests)
- [ ] Can compile programs using stdlib
- [ ] Prelude is auto-imported
- [ ] No memory leaks in stdlib
- [ ] No regressions in previous phases
- [ ] Code coverage ≥ 90%
- [ ] Documentation is comprehensive
- [ ] Example programs run correctly

## 💡 Implementation Notes

### Implementation Order
1. Core types (Option, Result, Box)
2. String and str
3. Collections (List, Map, Set)
4. I/O primitives
5. FFI support
6. Prelude
7. Compiler integration
8. Tests and documentation

### Testing Strategy
- Test each module independently first
- Write property-based tests for collections
- Test memory safety thoroughly
- Use real-world example programs
- Benchmark performance of collections
- Test edge cases (empty, single element, large size)

### Common Pitfalls
1. **Memory Leaks**: Ensure proper cleanup in collections
2. **Hash Collisions**: Handle collisions correctly in Map
3. **Capacity Management**: Grow arrays efficiently
4. **UTF-8**: Handle multi-byte characters correctly
5. **Error Handling**: Don't panic, use Result
6. **Ownership**: Collections must manage element ownership

### Debugging Tips
- Add debug printing to collection operations
- Visualize hash table buckets
- Test with small data first
- Use valgrind or similar to detect memory leaks
- Add assertions for invariants

### Performance Considerations
- Pre-allocate capacity when size is known
- Use power-of-2 growth for arrays
- Optimize hash function for common types
- Consider cache locality in data structures
- Profile and optimize hot paths

### FFI Considerations
- Keep FFI interface minimal
- Validate all data crossing FFI boundary
- Handle C errors properly
- Document safety requirements
- Test FFI thoroughly

### Future Extensions (Not This Phase)
- More collection types (HashMap, BTreeMap, LinkedList)
- Async I/O
- Network I/O
- Serialization/deserialization
- Regular expressions
- More string utilities
- Math library
- Random number generation

---

**Labels**: `phase-13`, `enhancement`, `stdlib`, `collections`, `io`, `high-priority`  
**Milestone**: Phase 13 - Standard Library Foundation  
**Assignee**: TBD
