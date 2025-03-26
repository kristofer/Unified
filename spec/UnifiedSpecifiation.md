# Unified Programming Language Specification

Version 1.0.0

## 1. Introduction

Unified is a modern systems programming language designed to combine memory safety, performance, and developer ergonomics. It draws inspiration from languages like Rust, Go, Swift, and Elixir while establishing its own unique approach to solving common programming challenges.

This specification is intended for compiler developers implementing the Unified language with LLVM as a backend.

## 2. Design Goals

Unified aims to achieve the following design goals:

1. **Memory Safety without Garbage Collection**: Provide strong memory safety guarantees through compile-time checks without requiring a garbage collector for most code.
2. **Ergonomic Ownership**: Implement an ownership system with explicit moves to make memory management more intuitive than in languages like Rust.
3. **Concurrency by Design**: Incorporate built-in concurrency primitives based on the actor model and async/await.
4. **Expressive Type System**: Support powerful generic programming, algebraic data types, and pattern matching.
5. **Pragmatic Developer Experience**: Prioritize clear syntax, helpful error messages, and rapid development cycles.

## 3. Lexical Structure

### 3.1 Keywords

```
actor    async    await    break    case     const
continue else     enum     false    fn       for
if       impl     import   in       interface let
loop     match    mod      module   move     mut
pub      region   return   self     spawn    struct
switch   true     try      type     var      where
while
```

### 3.2 Operators and Punctuation

```
+    -    *    /    %    &    |    ^    !    ~
+=   -=   *=   /=   %=   &=   |=   ^=   <<   >>
<<=  >>=  ==   !=   <    >    <=   >=   &&   ||
.    ..   ...  ,    ;    :    ::   ->   =>   ?
??   ?.   (    )    [    ]    {    }    @
```

### 3.3 Comments

```unified
// Single-line comment

/* 
   Multi-line comment
   that spans multiple lines
*/
```

### 3.4 Literals

- **Integer literals**: `42`, `0xFF`, `0b1010`
- **Float literals**: `3.14`, `1e-10`
- **String literals**: `"Hello, world!"`
- **Character literals**: `'A'`
- **Boolean literals**: `true`, `false`
- **Null literal**: `null`

### 3.5 Identifiers

Identifiers start with a letter or underscore, followed by zero or more letters, digits, or underscores.

## 4. Types

### 4.1 Basic Types

- **Integer types**: `Int`, `Int8`, `Int16`, `Int32`, `Int64`, `UInt`, `UInt8`, `UInt16`, `UInt32`, `UInt64`
- **Floating-point types**: `Float`, `Float32`, `Float64`
- **Boolean type**: `Bool`
- **Character type**: `Char`
- **String type**: `String`
- **Unit type**: `Unit` (represented as `()`)

### 4.2 Compound Types

- **Tuples**: `(T1, T2, ...)` 
- **Lists**: `List<T>`
- **Maps**: `Map<K, V>`
- **Sets**: `Set<T>`
- **References**: `&T` (immutable), `&mut T` (mutable)

### 4.3 User-Defined Types

- **Structs**: Named collections of fields
- **Enums**: Tagged unions with variants
- **Type aliases**: Named references to other types
- **Interfaces**: Abstract type specifications

### 4.4 Generic Types

Types can be parameterized by one or more type parameters:

```unified
struct Pair<A, B> {
    first: A
    second: B
}
```

### 4.5 Advanced Types

- **Union types**: `T1 | T2 | ...`
- **Function types**: `fn(T1, T2, ...) -> R`
- **Refinement types**: `type PositiveInt = Int refine |n| n > 0`

## 5. Variables and Bindings

### 5.1 Variable Declarations

```unified
// Immutable variable (default)
let name = "Alice"

// Mutable variable
let mut counter = 0

// Explicit type annotation
let id: Int64 = 1000000000

// Constants
const MAX_USERS = 100
```

### 5.2 Ownership and Borrowing

```unified
// Ownership transfer with explicit move
let s1 = "Hello"
let s2 = move s1  // s1 is no longer valid after this point

// Borrowing (non-mutable)
fn analyze(data: &String) -> Int {
    return data.length  // Can read but not modify
}

// Mutable borrowing
fn modify(data: &mut String) {
    data.append(" modified")  // Can modify
}
```

### 5.3 Region-based Memory Management

```unified
region temp {
    let large_data = load_large_file()
    process(large_data)
}  // large_data is freed when region ends
```

## 6. Expressions

### 6.1 Literals and Basic Expressions

```unified
// Literals
42              // Integer
3.14            // Float
"Hello"         // String
'A'             // Character
true, false     // Boolean
null            // Null

// Variable access
let x = 5
let y = x       // Access variable x

// Arithmetic operations
let sum = a + b
let product = a * b
```

### 6.2 Function Calls

```unified
let result = calculate(x, y)
```

### 6.3 Method Calls

```unified
let length = name.length()
```

### 6.4 Operator Expressions

```unified
// Arithmetic operators: +, -, *, /, %
let sum = a + b

// Comparison operators: ==, !=, <, >, <=, >=
let isEqual = a == b

// Logical operators: &&, ||, !
let combined = a && b

// Bitwise operators: &, |, ^, ~, <<, >>
let bits = a & b
```

### 6.5 Control Flow Expressions

```unified
// If expressions
let status = if active { "running" } else { "stopped" }

// Match expressions
let description = switch value {
    case 0 -> "Zero"
    case 1..5 -> "Small"
    case _ -> "Large"
}
```

### 6.6 Lambda Expressions

```unified
let add = |a, b| a + b
let multiplier = |x| { return x * factor }
```

### 6.7 Blocks

```unified
let result = {
    let temp = x * 2
    temp + y
}
```

## 7. Statements

### 7.1 Variable Declarations

```unified
let x = 5
let mut y = 10
```

### 7.2 Assignment Statements

```unified
counter = 0
counter += 1
```

### 7.3 Expression Statements

```unified
print("Hello")
calculate()
```

### 7.4 Control Flow Statements

```unified
// If statement
if condition {
    doSomething()
} else {
    doSomethingElse()
}

// While loop
while condition {
    process()
}

// For loop
for item in items {
    process(item)
}

// Loop with break/continue
loop {
    if done {
        break
    }
    if skip {
        continue
    }
    process()
}
```

### 7.5 Pattern Matching

```unified
switch value {
    case 0 -> print("Zero")
    case 1..5 -> print("Small")
    case let x if x > 100 -> print("Large: ${x}")
    case _ -> print("Something else")
}
```

### 7.6 Return Statement

```unified
fn calculate() -> Int {
    if condition {
        return 0
    }
    return result
}
```

## 8. Functions

### 8.1 Function Declarations

```unified
fn add(a: Int, b: Int) -> Int {
    return a + b
}

// Function with default parameter value
fn greet(name: String, greeting: String = "Hello") -> String {
    return "${greeting}, ${name}!"
}

// Function with multiple return values (via tuple)
fn divideWithRemainder(a: Int, b: Int) -> (Int, Int) {
    return (a / b, a % b)
}
```

### 8.2 Methods

```unified
struct Rectangle {
    width: Float
    height: Float
    
    fn area(self) -> Float {
        return self.width * self.height
    }
    
    fn scale(self: &mut Self, factor: Float) {
        self.width *= factor
        self.height *= factor
    }
}
```

### 8.3 Generic Functions

```unified
fn identity<T>(value: T) -> T {
    return value
}

fn findMax<T>(values: List<T>) -> Option<T> where T: Ord {
    // Implementation
}
```

### 8.4 Higher-order Functions

```unified
fn applyTwice(x: Int, f: fn(Int) -> Int) -> Int {
    return f(f(x))
}
```

## 9. Structs and Enums

### 9.1 Struct Declarations

```unified
struct Point {
    x: Float
    y: Float
}

// Struct with methods
struct Circle {
    center: Point
    radius: Float
    
    fn area(self) -> Float {
        return 3.14159 * self.radius * self.radius
    }
}
```

### 9.2 Struct Instantiation

```unified
let origin = Point{x: 0.0, y: 0.0}
let circle = Circle{center: origin, radius: 5.0}
```

### 9.3 Enum Declarations

```unified
enum Color {
    Red
    Green
    Blue
}

// Enum with associated values
enum Shape {
    Circle(radius: Float)
    Rectangle(width: Float, height: Float)
    Triangle(base: Float, height: Float)
}
```

### 9.4 Enum Usage

```unified
let color = Color.Red

let shape = Shape.Circle(radius: 5.0)

// Pattern matching with enums
switch shape {
    case Circle(let radius) -> 3.14159 * radius * radius
    case Rectangle(let width, let height) -> width * height
    case Triangle(let base, let height) -> 0.5 * base * height
}
```

### 9.5 Generic Structs and Enums

```unified
struct Box<T> {
    value: T
}

enum Result<T, E> {
    Success(T)
    Error(E)
}
```

## 10. Interfaces and Implementations

### 10.1 Interface Declarations

```unified
interface Serializable {
    fn serialize(self) -> String
    fn deserialize(data: String) -> Self
}

// Interface with associated type
interface Iterator {
    type Item
    
    fn next(self: &mut Self) -> Option<Self::Item>
    fn hasNext(self) -> Bool
}
```

### 10.2 Implementing Interfaces

```unified
impl Serializable for User {
    fn serialize(self) -> String {
        // Implementation
    }
    
    fn deserialize(data: String) -> Self {
        // Implementation
    }
}

// Implementation with associated type
impl Iterator for Counter {
    type Item = Int
    
    fn next(self: &mut Self) -> Option<Int> {
        // Implementation
    }
    
    fn hasNext(self) -> Bool {
        // Implementation
    }
}
```

### 10.3 Generic Implementations

```unified
impl<T> Serializable for List<T> where T: Serializable {
    // Implementation
}
```

## 11. Modules and Imports

### 11.1 Module Declarations

```unified
module math {
    pub fn add(a: Int, b: Int) -> Int {
        return a + b
    }
    
    fn helper() {
        // Private function
    }
}
```

### 11.2 Importing Items

```unified
// Import specific items
import math.add
import math.subtract

// Import multiple items
import math.{add, subtract}

// Import with alias
import math.add as mathAdd

// Import all public items
import math.*
```

## 12. Error Handling

### 12.1 Result Type

```unified
enum Result<T, E> {
    Success(T)
    Error(E)
}

fn divide(a: Int, b: Int) -> Result<Int, String> {
    if b == 0 {
        return Error("Division by zero")
    }
    return Success(a / b)
}
```

### 12.2 Option Type

```unified
enum Option<T> {
    Some(T)
    None
}

fn findUser(id: String) -> Option<User> {
    // Implementation
}
```

### 12.3 Error Propagation

```unified
fn processData() -> Result<Data, Error> {
    let file = openFile("data.txt")?  // Returns early if Error
    let content = file.readToString()?
    let data = parseData(content)?
    return Success(data)
}
```

### 12.4 Pattern Matching for Error Handling

```unified
let result = divide(10, 0)

switch result {
    case Success(let value) -> print("Result: ${value}")
    case Error(let message) -> print("Error: ${message}")
}
```

## 13. Concurrency

### 13.1 Actors

```unified
actor Counter {
    var count = 0
    
    fn increment() -> Int {
        count += 1
        return count
    }
    
    fn reset() {
        count = 0
    }
}

// Creating and using an actor
let counter = spawn Counter{}
counter.increment()  // Message sent asynchronously
```

### 13.2 Async/Await

```unified
async fn fetchData(url: String) -> Result<String, Error> {
    // Asynchronous implementation
}

async fn processUrls(urls: List<String>) -> List<String> {
    let tasks = urls.map(url -> async fetchData(url))
    let results = await all(tasks)
    // Process results
}
```

### 13.3 Channels

```unified
// Create a channel
let channel = Channel<String>.new()

// Send and receive
channel.send("Hello")
let message = channel.receive()

// Select over multiple channels
select {
    case msg = channel1.receive() -> handleMessage(msg)
    case channel2.send(value) -> print("Sent value")
    case after 1.seconds -> print("Timeout")
}
```

## 14. Memory Management

### 14.1 Ownership System

```unified
// Ownership model
let s1 = "Hello"       // s1 owns the string
let s2 = move s1       // Ownership transferred to s2

// s1 is no longer valid here
```

### 14.2 Borrowing

```unified
// Immutable borrow
fn analyze(data: &String) {
    // Can read but not modify data
}

// Mutable borrow
fn modify(data: &mut String) {
    // Can modify data
}
```

### 14.3 Region-based Memory Management

```unified
region temp {
    // Memory allocated in this region
    let large_array = Array<Int>.new(1000000)
    
    // Use large_array
    
}  // All memory from region freed at once
```

### 14.4 Optional Garbage Collection

```unified
// Opt into garbage collection for specific types
@gc class Graph {
    nodes: List<Node>
    edges: List<Edge>
    
    // Complex cyclic structure handled by GC
    fn addEdge(self: &mut Self, from: NodeId, to: NodeId) {
        // No need to worry about reference cycles
        self.edges.push(Edge{from, to})
        self.nodes[from].connections.push(to)
        self.nodes[to].connections.push(from)
    }
}
```

### 14.5 Lifetime Annotations

```unified
// Function that returns a reference needs a lifetime annotation
fn longest<'a>(x: &'a String, y: &'a String) -> &'a String {
    if x.len() > y.len() {
        return x
    } else {
        return y
    }
}
```

## 15. Pattern Matching and Destructuring

### 15.1 Basic Pattern Matching

```unified
switch value {
    case 0 -> print("Zero")
    case 1..5 -> print("Small")
    case let x if x > 100 -> print("Large: ${x}")
    case _ -> print("Something else")
}
```

### 15.2 Destructuring

```unified
// Destructuring tuples
let point = (10, 20)
let (x, y) = point

// Destructuring structs
let Person{name, age} = person

// Destructuring in function parameters
fn processPoint((x, y): (Int, Int)) {
    // Implementation
}
```

### 15.3 Pattern Guards

```unified
switch user {
    case User{name, age} if age >= 18 -> allowAccess(name)
    case User{name, ..} -> denyAccess(name)
}
```

## 16. Generics and Type Constraints

### 16.1 Generic Functions and Types

```unified
fn identity<T>(value: T) -> T {
    return value
}

struct Pair<A, B> {
    first: A
    second: B
}
```

### 16.2 Type Constraints

```unified
fn max<T>(a: T, b: T) -> T where T: Ord {
    if a > b {
        return a
    } else {
        return b
    }
}

fn process<T>(value: T) -> String 
    where T: Display,        // Must be displayable as string
          T: Serialize       // Must be serializable
{
    // Implementation
}
```

### 16.3 Associated Types

```unified
interface Collection {
    type Item
    
    fn add(self: &mut Self, item: Self::Item)
    fn contains(self, item: &Self::Item) -> Bool
}

impl Collection for List<Int> {
    type Item = Int
    
    fn add(self: &mut Self, item: Int) {
        self.push(item)
    }
    
    fn contains(self, item: &Int) -> Bool {
        return self.indexOf(item).isSome()
    }
}
```

## 17. Advanced Features

### 17.1 String Interpolation

```unified
let name = "Alice"
let greeting = "Hello, ${name}!"

let a = 5
let b = 10
let sum = "${a} + ${b} = ${a + b}"
```

### 17.2 List Comprehensions

```unified
// Basic list comprehension
let squares = [x * x for x in 1..10]

// With condition
let evenSquares = [x * x for x in 1..10 if x % 2 == 0]
```

### 17.3 Optional Chaining

```unified
// Optional chaining
let name = user?.profile?.name

// Null coalescing
let displayName = user?.name ?? "Anonymous"
```

### 17.4 Extension Methods

```unified
// Extend existing types with new methods
extension String {
    fn capitalize(self) -> String {
        if self.isEmpty() {
            return self
        }
        return self[0].toUpperCase() + self[1..]
    }
}

// Usage
let capitalized = "hello".capitalize()  // "Hello"
```

## 18. Compiler Implementation Guidelines

### 18.1 LLVM Integration

Unified uses LLVM as its backend for code generation. The compiler should translate Unified's high-level constructs to LLVM IR before generating machine code.

#### 18.1.1 Type Mapping

| Unified Type | LLVM Type |
|--------------|-----------|
| Bool | i1 |
| Int8 | i8 |
| Int16 | i16 |
| Int32 | i32 |
| Int64 | i64 |
| UInt8 | i8 |
| UInt16 | i16 |
| UInt32 | i32 |
| UInt64 | i64 |
| Float32 | float |
| Float64 | double |
| Char | i8 or i32 depending on Unicode support |
| String | %String struct (with length and pointer to data) |
| &T | Pointer to T |
| (T1, T2, ...) | Struct with fields of corresponding types |

#### 18.1.2 Ownership Implementation

The ownership system should be implemented primarily at compile time through the following approaches:

1. **Ownership Tracking**: The compiler tracks ownership of values through a flow-sensitive analysis.
2. **Move Semantics**: When a value is moved, mark the original variable as invalid.
3. **Borrow Checking**: Ensure that borrows do not outlive the referenced value and that mutable borrows are exclusive.

```llvm
; Example of how a move might be represented in LLVM IR
%old_owner = alloca %String
%new_owner = alloca %String
; Move operation
%temp = load %String, %String* %old_owner
store %String %temp, %String* %new_owner
; Insert invalidation of old_owner
call void @invalidate_owner(%String* %old_owner)
```

#### 18.1.3 Region-based Memory Management

Regions should be implemented using custom allocators:

```llvm
; Region creation
%region = call i8* @region_create()

; Allocation within region
%ptr = call i8* @region_alloc(%region, i64 %size)

; Region destruction
call void @region_destroy(%region)
```

#### 18.1.4 Actor Implementation

Actors should be compiled to state machines with message queues:

```llvm
; Actor creation
%actor = call %Actor* @actor_create(%TypeDescriptor* @ActorType)

; Message dispatch
define void @actor_dispatch(%Actor* %self, %Message* %msg) {
  %msg_type = load i32, i32* getelementptr %msg, i32 0
  switch i32 %msg_type, label %default [
    i32 1, label %handle_increment
    i32 2, label %handle_reset
  ]
  
  handle_increment:
    ; Actor method implementation
    ret void
  
  handle_reset:
    ; Actor method implementation
    ret void
}
```

### 18.2 Compilation Pipeline

The Unified compiler should implement the following pipeline:

1. **Lexical Analysis**: Convert source code to tokens
2. **Parsing**: Build AST from tokens
3. **Name Resolution**: Resolve symbols and scopes
4. **Type Checking**: Verify type correctness
5. **Ownership Checking**: Verify ownership rules
6. **Monomorphization**: Generate specialized versions of generic code
7. **IR Generation**: Convert AST to LLVM IR
8. **Optimization**: Apply LLVM optimization passes
9. **Code Generation**: Generate machine code

### 18.3 Standard Library Implementation

The standard library should be implemented in a combination of Unified itself and lower-level languages (C/C++ or Rust) for platform-specific functionality.

Core data structures like List, Map, and Set should be implemented in Unified to serve as examples of idiomatic code.

## 19. Grammar (EBNF)

```ebnf
program        ::= item*

item           ::= moduleDecl | functionDecl | structDecl 
                 | enumDecl | interfaceDecl | implDecl 
                 | actorDecl | typeAlias | importDecl

moduleDecl     ::= 'module' identifier '{' item* '}'

functionDecl   ::= 'fn' identifier genericParams? 
                   '(' paramList? ')' ('->' type)? block

structDecl     ::= 'struct' identifier genericParams? 
                   '{' (field | functionDecl)* '}'

enumDecl       ::= 'enum' identifier genericParams? 
                   '{' enumVariant (',' enumVariant)* '}'

interfaceDecl  ::= 'interface' identifier genericParams? 
                   '{' (functionSig | typeAssoc)* '}'

implDecl       ::= 'impl' genericParams? type 'for' type 
                   '{' functionDecl* '}'

actorDecl      ::= 'actor' identifier genericParams? 
                   '{' (field | functionDecl)* '}'

typeAlias      ::= 'type' identifier genericParams? '=' type ';'

importDecl     ::= 'import' importPath ';'

paramList      ::= param (',' param)*

param          ::= identifier ':' type

genericParams  ::= '<' genericParam (',' genericParam)* '>'

type           ::= identifier ('<' typeList '>')?
                 | '&' type
                 | '&' 'mut' type
                 | '(' typeList ')'
                 | 'fn' '(' typeList? ')' '->' type

block          ::= '{' statement* expr? '}'

statement      ::= letStatement
                 | exprStatement
                 | returnStatement
                 | ifStatement
                 | forStatement
                 | whileStatement
                 | switchStatement

expr           ::= primary
                 | expr operator expr
                 | expr '(' exprList? ')'
                 | expr '.' identifier
                 | 'if' expr block ('else' block)?
                 | 'switch' expr '{' caseClause* '}'
                 | lambdaExpr
```

## 20. Compatibility and Interoperability

### 20.1 C FFI (Foreign Function Interface)

Unified provides a Foreign Function Interface for interoperating with C code:

```unified
// Importing C functions
@extern("C")
fn printf(format: &Char, ...) -> Int

// Exporting Unified functions to C
@export("C")
fn unified_function() -> Int {
    return 42
}
```

### 20.2 Platform-Specific Code

Conditional compilation allows for platform-specific implementations:

```unified
@if(os = "windows")
fn getPathSeparator() -> Char {
    return '\\'
}

@if(os = "linux" || os = "macos")
fn getPathSeparator() -> Char {
    return '/'
}
```

### 20.3 Inline Assembly

For low-level system programming, inline assembly is supported:

```unified
fn fast_add(a: Int, b: Int) -> Int {
    let result: Int
    
    @asm {
        "add %[a], %[b]"
        : [result] "=r" (result)
        : [a] "r" (a), [b] "r" (b)
    }
    
    return result
}
```

## 21. Versioning and Evolution

The Unified language uses semantic versioning (MAJOR.MINOR.PATCH) for its releases.

### 21.1 Language Stability

1. **Major Versions**: May contain breaking changes to the language syntax or semantics.
2. **Minor Versions**: Add new features in a backward-compatible manner.
3. **Patch Versions**: Only contain bug fixes and performance improvements.

### 21.2 Deprecation Policy

Features slated for removal will be:
1. Marked as deprecated in one minor version
2. Generate warnings in the next minor version
3. Removed in the next major version

### 21.3 Edition Mechanism

Similar to Rust, Unified uses an "edition" mechanism to handle larger language evolution while maintaining backward compatibility:

```unified
// Specify which language edition to use
@edition("2025")
module my_module {
    // Code using 2025 edition features
}
```

## Appendix A: LLVM IR Examples

### A.1 Function Declaration

```llvm
; Unified function: fn add(a: Int, b: Int) -> Int { return a + b }
define i32 @add(i32 %a, i32 %b) {
  %result = add i32 %a, %b
  ret i32 %result
}
```

### A.2 Struct Definition

```llvm
; Unified struct: struct Point { x: Float, y: Float }
%Point = type { float, float }

; Constructor
define %Point @Point_new(float %x, float %y) {
  %point = alloca %Point
  %x_ptr = getelementptr %Point, %Point* %point, i32 0, i32 0
  %y_ptr = getelementptr %Point, %Point* %point, i32 0, i32 1
  store float %x, float* %x_ptr
  store float %y, float* %y_ptr
  %result = load %Point, %Point* %point
  ret %Point %result
}
```

### A.3 Region-Based Memory

```llvm
; Region implementation
%Region = type { i8*, i64, i64 }

; Create a new region
define %Region* @region_create() {
  %size = mul i64 1024, 1024  ; 1MB initial size
  %mem = call i8* @malloc(i64 %size)
  %region = call %Region* @malloc(i64 ptrtoint (%Region* getelementptr (%Region, %Region* null, i32 1) to i64))
  %mem_ptr = getelementptr %Region, %Region* %region, i32 0, i32 0
  store i8* %mem, i8** %mem_ptr
  %size_ptr = getelementptr %Region, %Region* %region, i32 0, i32 1
  store i64 %size, i64* %size_ptr
  %used_ptr = getelementptr %Region, %Region* %region, i32 0, i32 2
  store i64 0, i64* %used_ptr
  ret %Region* %region
}

; Allocate memory from region
define i8* @region_alloc(%Region* %region, i64 %size) {
  ; Implementation
}

; Free entire region
define void @region_destroy(%Region* %region) {
  %mem_ptr = getelementptr %Region, %Region* %region, i32 0, i32 0
  %mem = load i8*, i8** %mem_ptr
  call void @free(i8* %mem)
  call void @free(i8* %region)
  ret void
}
```

## Appendix B: Runtime Library Interface

The Unified runtime library provides core functionality required by the language, including:

### B.1 Memory Management

```c
// Core memory management functions
void* unified_allocate(size_t size);
void unified_free(void* ptr);
void* unified_realloc(void* ptr, size_t new_size);

// Region-based memory management
Region* unified_region_create(void);
void* unified_region_alloc(Region* region, size_t size);
void unified_region_destroy(Region* region);

// Reference counting for shared ownership
void unified_rc_increment(void* ptr);
bool unified_rc_decrement(void* ptr);

// Optional garbage collection
void unified_gc_register(void* ptr, size

## Appendix B: Runtime Library Interface (continued)

### B.1 Memory Management (continued)

```c
// Optional garbage collection
void unified_gc_register(void* ptr, size_t size, TypeDescriptor* type);
void unified_gc_unregister(void* ptr);
void unified_gc_collect(void);
void unified_gc_disable(void);
void unified_gc_enable(void);
```

### B.2 Concurrency Runtime

```c
// Actor system
Actor* unified_actor_create(TypeDescriptor* type);
void unified_actor_send(Actor* actor, Message* message);
void unified_actor_destroy(Actor* actor);

// Scheduler
void unified_scheduler_init(int num_threads);
void unified_scheduler_shutdown(void);
TaskId unified_scheduler_spawn(Task* task);
void unified_scheduler_yield(void);

// Channels
Channel* unified_channel_create(size_t capacity, TypeDescriptor* elem_type);
bool unified_channel_send(Channel* channel, void* data, bool blocking);
bool unified_channel_receive(Channel* channel, void** data, bool blocking);
void unified_channel_close(Channel* channel);

// Async/await
Future* unified_future_create(TypeDescriptor* result_type);
void unified_future_complete(Future* future, void* result);
void unified_future_complete_error(Future* future, Error* error);
void unified_await(Future* future, void** result);
```

### B.3 Type System

```c
// Type descriptors for runtime type information
TypeDescriptor* unified_type_get(const char* name);
bool unified_type_is_subtype(TypeDescriptor* type, TypeDescriptor* potential_supertype);
size_t unified_type_size(TypeDescriptor* type);
void unified_type_destroy(void* instance, TypeDescriptor* type);
void unified_type_copy(void* dest, void* src, TypeDescriptor* type);
char* unified_type_to_string(void* instance, TypeDescriptor* type);

// Interface method dispatch
void* unified_interface_method_lookup(void* instance, TypeDescriptor* type, const char* method_name);
```

### B.4 Error Handling

```c
// Error creation and handling
Error* unified_error_create(ErrorType type, const char* message);
void unified_error_destroy(Error* error);
const char* unified_error_message(Error* error);
ErrorType unified_error_type(Error* error);

// Panic handling
void unified_panic(const char* message, const char* file, int line);
void unified_set_panic_handler(PanicHandler handler);
```

### B.5 I/O and System Integration

```c
// File operations
File* unified_file_open(const char* path, FileMode mode);
size_t unified_file_read(File* file, void* buffer, size_t size);
size_t unified_file_write(File* file, const void* buffer, size_t size);
void unified_file_close(File* file);

// Console I/O
void unified_print(const char* format, ...);
char* unified_read_line(void);

// Network operations
Socket* unified_socket_create(SocketType type);
bool unified_socket_connect(Socket* socket, const char* host, int port);
bool unified_socket_bind(Socket* socket, int port);
void unified_socket_close(Socket* socket);
```

## Appendix C: Best Practices for LLVM Integration

### C.1 Memory Management Implementation

When implementing Unified's memory management system using LLVM, consider these approaches:

#### C.1.1 Ownership Tracking

The ownership system should be represented in the LLVM IR through a combination of:

1. **Static Analysis**: Most ownership checks should happen at compile time before generating LLVM IR.
2. **Dynamic Checks**: When necessary, insert runtime checks for borrowing rules that can't be verified statically.

```llvm
; Example of tracking ownership through metadata
%struct.String = type { i8*, i64 }

; Define metadata for ownership
!0 = !{!"owned"}
!1 = !{!"borrowed"}
!2 = !{!"moved"}

; Function that takes ownership
define void @take_ownership(%struct.String* %str) {
  ; Mark the parameter as moved
  %str_md = metadata !{%str, !2}
  ; Use the string
  ; ...
  ret void
}
```

#### C.1.2 Region-Based Memory

Implement regions using bump-allocation for efficiency:

```llvm
; Define region structure
%struct.Region = type { i8*, i64, i64 }  ; memory block, size, used

; Create a region
define %struct.Region* @region_create(i64 %size) {
  %region_ptr = call i8* @malloc(i64 ptrtoint (%struct.Region* getelementptr (%struct.Region, %struct.Region* null, i32 1) to i64))
  %region = bitcast i8* %region_ptr to %struct.Region*
  
  ; Allocate memory block
  %block_ptr = call i8* @malloc(i64 %size)
  
  ; Initialize region
  %block_field = getelementptr %struct.Region, %struct.Region* %region, i32 0, i32 0
  store i8* %block_ptr, i8** %block_field
  
  %size_field = getelementptr %struct.Region, %struct.Region* %region, i32 0, i32 1
  store i64 %size, i64* %size_field
  
  %used_field = getelementptr %struct.Region, %struct.Region* %region, i32 0, i32 2
  store i64 0, i64* %used_field
  
  ret %struct.Region* %region
}

; Allocate from region (bump allocator)
define i8* @region_alloc(%struct.Region* %region, i64 %size) {
  ; Align size to proper boundary
  %aligned_size = add i64 %size, 7
  %aligned_size2 = and i64 %aligned_size, -8  ; Align to 8 bytes
  
  ; Get current position
  %used_field = getelementptr %struct.Region, %struct.Region* %region, i32 0, i32 2
  %used = load i64, i64* %used_field
  
  ; Calculate new position
  %new_used = add i64 %used, %aligned_size2
  
  ; Check if enough space
  %size_field = getelementptr %struct.Region, %struct.Region* %region, i32 0, i32 1
  %size = load i64, i64* %size_field
  %has_space = icmp ule i64 %new_used, %size
  br i1 %has_space, label %allocate, label %out_of_memory
  
allocate:
  ; Update used space
  store i64 %new_used, i64* %used_field
  
  ; Get memory block
  %block_field = getelementptr %struct.Region, %struct.Region* %region, i32 0, i32 0
  %block = load i8*, i8** %block_field
  
  ; Calculate pointer to allocated memory
  %result = getelementptr i8, i8* %block, i64 %used
  
  ret i8* %result
  
out_of_memory:
  call void @panic(i8* getelementptr inbounds ([21 x i8], [21 x i8]* @.str.out_of_memory, i32 0, i32 0))
  ret i8* null
}
```

### C.2 Actor Implementation

Actors should be implemented as state machines with message queues:

```llvm
; Actor structure
%struct.Actor = type { 
  i8* (i8*, %struct.Message*)*, ; message handler function
  i8*,                          ; actor state
  %struct.MessageQueue*         ; message queue
}

; Create an actor
define %struct.Actor* @actor_create(i8* (i8*, %struct.Message*)* %handler, i8* %initial_state) {
  %actor = call i8* @malloc(i64 ptrtoint (%struct.Actor* getelementptr (%struct.Actor, %struct.Actor* null, i32 1) to i64))
  %actor_ptr = bitcast i8* %actor to %struct.Actor*
  
  ; Set handler
  %handler_field = getelementptr %struct.Actor, %struct.Actor* %actor_ptr, i32 0, i32 0
  store i8* (i8*, %struct.Message*)* %handler, i8* (i8*, %struct.Message*)** %handler_field
  
  ; Set state
  %state_field = getelementptr %struct.Actor, %struct.Actor* %actor_ptr, i32 0, i32 1
  store i8* %initial_state, i8** %state_field
  
  ; Create message queue
  %queue = call %struct.MessageQueue* @message_queue_create()
  %queue_field = getelementptr %struct.Actor, %struct.Actor* %actor_ptr, i32 0, i32 2
  store %struct.MessageQueue* %queue, %struct.MessageQueue** %queue_field
  
  ; Register actor with scheduler
  call void @scheduler_register_actor(%struct.Actor* %actor_ptr)
  
  ret %struct.Actor* %actor_ptr
}

; Send message to actor (non-blocking)
define void @actor_send(%struct.Actor* %actor, %struct.Message* %message) {
  %queue_field = getelementptr %struct.Actor, %struct.Actor* %actor, i32 0, i32 2
  %queue = load %struct.MessageQueue*, %struct.MessageQueue** %queue_field
  
  call void @message_queue_push(%struct.MessageQueue* %queue, %struct.Message* %message)
  
  ; Notify scheduler that actor has work
  call void @scheduler_notify_actor_has_work(%struct.Actor* %actor)
  
  ret void
}

; Process messages for actor (called by scheduler)
define void @actor_process_messages(%struct.Actor* %actor) {
  %queue_field = getelementptr %struct.Actor, %struct.Actor* %actor, i32 0, i32 2
  %queue = load %struct.MessageQueue*, %struct.MessageQueue** %queue_field
  
  ; Process up to 10 messages at a time
  %i = alloca i32
  store i32 0, i32* %i
  
  br label %loop_header
  
loop_header:
  %i_val = load i32, i32* %i
  %continue = icmp slt i32 %i_val, 10
  br i1 %continue, label %loop_body, label %loop_exit
  
loop_body:
  ; Try to get a message
  %message = call %struct.Message* @message_queue_try_pop(%struct.MessageQueue* %queue)
  %has_message = icmp ne %struct.Message* %message, null
  br i1 %has_message, label %process_message, label %loop_exit
  
process_message:
  ; Get handler and state
  %handler_field = getelementptr %struct.Actor, %struct.Actor* %actor, i32 0, i32 0
  %handler = load i8* (i8*, %struct.Message*)*, i8* (i8*, %struct.Message*)** %handler_field
  
  %state_field = getelementptr %struct.Actor, %struct.Actor* %actor, i32 0, i32 1
  %state = load i8*, i8** %state_field
  
  ; Call handler
  %new_state = call i8* %handler(i8* %state, %struct.Message* %message)
  
  ; Update state
  store i8* %new_state, i8** %state_field
  
  ; Free message
  call void @message_free(%struct.Message* %message)
  
  ; Increment counter
  %i_next = add i32 %i_val, 1
  store i32 %i_next, i32* %i
  
  br label %loop_header
  
loop_exit:
  ret void
}
```

### C.3 Generic Code Implementation

Generic code should be implemented using monomorphization - generating specialized versions for each combination of type parameters:

```llvm
; Example of monomorphized function for List<Int>
define i32 @List_Int_sum(%struct.List_Int* %list) {
  ; Implementation for List<Int>
  ; ...
}

; Example of monomorphized function for List<Float>
define float @List_Float_sum(%struct.List_Float* %list) {
  ; Implementation for List<Float>
  ; ...
}
```

### C.4 Interface Implementation

Interfaces should be implemented using virtual method tables:

```llvm
; Interface vtable for Serializable
%struct.Serializable_vtable = type {
  %struct.String* (i8*)*, ; serialize method
  i8* (%struct.String*)*, ; deserialize method
}

; Implementation for a specific type
@User_Serializable_vtable = global %struct.Serializable_vtable {
  %struct.String* (i8*)* @User_serialize,
  i8* (%struct.String*)* @User_deserialize
}

; Interface dispatch
define %struct.String* @serialize(i8* %obj, %struct.Serializable_vtable* %vtable) {
  %serialize_fn = getelementptr %struct.Serializable_vtable, %struct.Serializable_vtable* %vtable, i32 0, i32 0
  %fn = load %struct.String* (i8*)*, %struct.String* (i8*)** %serialize_fn
  %result = call %struct.String* %fn(i8* %obj)
  ret %struct.String* %result
}
```

## Appendix D: Optimization Guidelines

### D.1 LLVM Optimization Passes

Recommended LLVM optimization passes for Unified code:

1. **Basic Optimizations**:
   - Instruction combining (`-instcombine`)
   - Common subexpression elimination (`-gvn`)
   - Dead code elimination (`-dce`)

2. **Intermediate Optimizations**:
   - Function inlining (`-inline`)
   - Loop unrolling (`-loop-unroll`)
   - Loop vectorization (`-loop-vectorize`)
   - Memory-to-register promotion (`-mem2reg`)

3. **Advanced Optimizations**:
   - Global value numbering (`-gvn`)
   - Scalar replacement of aggregates (`-sroa`)
   - Jump threading (`-jump-threading`)
   - Tail call elimination (`-tailcallelim`)

### D.2 Unified-Specific Optimizations

Specialized optimization passes for Unified features:

1. **Region Allocation Optimization**:
   - Detect when heap allocations can be converted to region allocations
   - Eliminate redundant region creation/destruction

2. **Ownership Optimization**:
   - Eliminate unnecessary ownership checks
   - Convert moves to copies when safe
   - Optimize borrow checking

3. **Actor Optimizations**:
   - Inline small actor methods
   - Batch message sends
   - Optimize actor scheduling

## Appendix E: Standard Library Overview

The Unified Standard Library consists of the following core modules:

### E.1 Core Primitives
- `std.core.Int`, `std.core.Float`, `std.core.Bool`, etc.
- `std.core.Option<T>` and `std.core.Result<T, E>`
- `std.core.Tuple` types

### E.2 Collections
- `std.collections.List<T>`: Dynamic array implementation
- `std.collections.Map<K, V>`: Hash-based key-value store
- `std.collections.Set<T>`: Hash-based set implementation
- `std.collections.Queue<T>`: FIFO queue implementation
- `std.collections.Stack<T>`: LIFO stack implementation
- `std.collections.LinkedList<T>`: Doubly-linked list
- `std.collections.Tree<T>`: Generic tree structure
- `std.collections.PriorityQueue<T>`: Priority queue based on heap

### E.3 Concurrency
- `std.concurrency.Actor`: Base actor functionality
- `std.concurrency.Channel<T>`: Communication channels
- `std.concurrency.Future<T>`: Asynchronous values
- `std.concurrency.Mutex<T>`: Mutual exclusion for shared state
- `std.concurrency.Atomic<T>`: Atomic operations
- `std.concurrency.Thread`: Thread management
- `std.concurrency.Scheduler`: Task scheduler configuration

### E.4 I/O
- `std.io.File`: File system operations
- `std.io.Console`: Terminal input/output
- `std.io.Stream`: Abstract data streams
- `std.io.Reader` and `std.io.Writer`: I/O abstractions
- `std.io.Path`: Platform-independent path handling
- `std.io.Serialization`: Generic serialization framework

### E.5 Network
- `std.net.Socket`: TCP/UDP socket operations
- `std.net.Http`: HTTP client and server
- `std.net.Url`: URL parsing and manipulation
- `std.net.Dns`: Domain name resolution
- `std.net.TLS`: Secure connections

### E.6 Text Processing
- `std.text.String`: String manipulation
- `std.text.Format`: String formatting
- `std.text.Regex`: Regular expressions
- `std.text.Unicode`: Unicode handling
- `std.text.Parsing`: Text parsing utilities

### E.7 Time and Date
- `std.time.Duration`: Time span representation
- `std.time.Instant`: Point in time
- `std.time.DateTime`: Calendar date and time
- `std.time.Timer`: Time measurement
- `std.time.Scheduler`: Time-based scheduling

### E.8 Math
- `std.math.Basic`: Common math functions
- `std.math.Complex`: Complex number operations
- `std.math.Vector`: Vector operations
- `std.math.Matrix`: Matrix operations
- `std.math.Random`: Random number generation
- `std.math.Statistics`: Statistical functions

### E.9 System
- `std.system.Process`: Process management
- `std.system.Env`: Environment variables
- `std.system.OS`: Operating system utilities
- `std.system.Memory`: Memory management
- `std.system.Signal`: Signal handling

### E.10 Testing
- `std.test.Assert`: Test assertions
- `std.test.Benchmark`: Performance measurement
- `std.test.Mock`: Object mocking
- `std.test.Property`: Property-based testing

## Appendix F: Common Design Patterns

This section outlines idiomatic patterns for solving common problems in Unified.

### F.1 RAII (Resource Acquisition Is Initialization)

Unified's ownership system naturally supports RAII:

```unified
struct FileWrapper {
    file: File
    
    fn new(path: String) -> Result<Self, IOError> {
        let file = File.open(path)?
        return Success(Self{file: file})
    }
    
    // File is automatically closed when FileWrapper is dropped
    fn destroy(self) {
        self.file.close()
    }
}

fn process_file() -> Result<(), IOError> {
    // File is automatically closed at the end of this scope
    let wrapper = FileWrapper.new("data.txt")?
    
    // Use wrapper.file
    let data = wrapper.file.readToString()?
    process_data(data)
    
    return Success(())
}
```

### F.2 Builder Pattern

For constructing complex objects:

```unified
struct QueryBuilder {
    table: String
    conditions: List<String>
    ordering: Option<String>
    limit: Option<Int>
    
    fn new() -> Self {
        return Self{
            table: "",
            conditions: List<String>.new(),
            ordering: None,
            limit: None
        }
    }
    
    fn from(self: &mut Self, table: String) -> &mut Self {
        self.table = table
        return self
    }
    
    fn where(self: &mut Self, condition: String) -> &mut Self {
        self.conditions.push(condition)
        return self
    }
    
    fn orderBy(self: &mut Self, field: String, ascending: Bool = true) -> &mut Self {
        let direction = if ascending { "ASC" } else { "DESC" }
        self.ordering = Some("${field} ${direction}")
        return self
    }
    
    fn limit(self: &mut Self, count: Int) -> &mut Self {
        self.limit = Some(count)
        return self
    }
    
    fn build(self) -> Result<Query, QueryError> {
        if self.table.isEmpty() {
            return Error(QueryError.MissingTable)
        }
        
        // Construct query
        let query = Query{
            table: self.table,
            conditions: self.conditions,
            ordering: self.ordering,
            limit: self.limit
        }
        
        return Success(query)
    }
}

// Usage
let query = QueryBuilder.new()
    .from("users")
    .where("age > 18")
    .orderBy("last_name")
    .limit(10)
    .build()?
```

### F.3 Strategy Pattern

Using interfaces to define interchangeable algorithms:

```unified
interface SortStrategy<T> {
    fn sort(self, items: &mut List<T>);
}

struct QuickSort<T> {}
impl<T> SortStrategy<T> for QuickSort<T> where T: Ord {
    fn sort(self, items: &mut List<T>) {
        // QuickSort implementation
    }
}

struct MergeSort<T> {}
impl<T> SortStrategy<T> for MergeSort<T> where T: Ord {
    fn sort(self, items: &mut List<T>) {
        // MergeSort implementation
    }
}

struct Sorter<T> {
    strategy: impl SortStrategy<T>
    
    fn new(strategy: impl SortStrategy<T>) -> Self {
        return Self{strategy: strategy}
    }
    
    fn sort(self, items: &mut List<T>) {
        self.strategy.sort(items)
    }
}

// Usage
let quickSorter = Sorter.new(QuickSort<Int>{})
quickSorter.sort(&mut numbers)

let mergeSorter = Sorter.new(MergeSort<Int>{})
mergeSorter.sort(&mut numbers)
```

### F.4 Observer Pattern

Using channels and actors for event notification:

```unified
actor EventPublisher<T> {
    subscribers: List<Channel<T>>
    
    fn subscribe(channel: Channel<T>) {
        subscribers.push(channel)
    }
    
    fn unsubscribe(channel: Channel<T>) {
        subscribers = subscribers.filter(|c| c != channel)
    }
    
    fn publish(event: T) {
        for channel in subscribers {
            channel.send(event)
        }
    }
}

// Usage
actor UserInterface {
    events: Channel<Event>
    
    fn new() -> Self {
        let channel = Channel<Event>.new()
        EventPublisher.subscribe(channel)
        return Self{events: channel}
    }
    
    fn run() {
        while true {
            switch events.receive() {
                case Some(event) -> handleEvent(event)
                case None -> break
            }
        }
    }
    
    fn handleEvent(event: Event) {
        // Process event
    }
}
```

### F.5 Dependency Injection

Using interface-based design for loose coupling:

```unified
interface Database {
    fn query(self, sql: String) -> Result<QueryResult, DBError>;
    fn execute(self, sql: String) -> Result<Int, DBError>;
}

struct PostgresDB {
    connection: PostgresConnection
    
    fn new(connString: String) -> Result<Self, DBError> {
        let connection = PostgresConnection.connect(connString)?
        return Success(Self{connection: connection})
    }
}

impl Database for PostgresDB {
    fn query(self, sql: String) -> Result<QueryResult, DBError> {
        return self.connection.query(sql)
    }
    
    fn execute(self, sql: String) -> Result<Int, DBError> {
        return self.connection.execute(sql)
    }
}

struct UserRepository {
    db: impl Database
    
    fn new(db: impl Database) -> Self {
        return Self{db: db}
    }
    
    fn findById(self, id: String) -> Result<User, RepositoryError> {
        let sql = "SELECT * FROM users WHERE id = '${id}'"
        let result = self.db.query(sql).mapError(|e| RepositoryError.Database(e))?
        
        if result.isEmpty() {
            return Error(RepositoryError.NotFound)
        }
        
        return Success(User.fromRow(result.rows[0]))
    }
}

// Usage
let db = PostgresDB.new("connection_string")?
let userRepo = UserRepository.new(db)
let user = userRepo.findById("123")?
```

## Appendix G: Compiler Validation Suite

The following test cases should be used to validate a Unified compiler implementation:

### G.1 Ownership and Borrowing Tests

1. Basic ownership transfer
2. Borrowing and lifetimes
3. Mutable borrowing constraints
4. Function parameter ownership
5. Region-based allocation
6. Nested borrows
7. Move semantics with closures

### G.2 Type System Tests

1. Generic type specialization
2. Interface implementations
3. Type constraints
4. Associated types
5. Union types
6. Complex type inference
7. Refinement types

### G.3 Concurrency Tests

1. Actor message passing
2. Channel usage patterns
3. Async/await functionality
4. Future composition
5. Thread synchronization
6. Scheduler behavior
7. Deadlock detection

### G.4 Memory Management Tests

1. Ownership model correctness
2. Region allocation/deallocation
3. Optional GC behavior
4. Memory leak detection
5. RAII patterns
6. Cyclic references

### G.5 Performance Benchmarks

1. Large data structure processing
2. Concurrent request handling
3. Memory allocation patterns
4. Interface method dispatch
5. Generic code performance
6. Actor system throughput

## Appendix H: Future Language Evolution

The following areas are identified for potential future expansion of the Unified language:

### H.1 Advanced Type System Features

1. **Dependent Types**: Expand support for values as part of types
2. **Effect Systems**: Track and constrain side effects through the type system
3. **Linear Types**: More precise control over resource usage
4. **Gradual Typing**: Allow dynamic typing in specific contexts
5. **Typestate Programming**: Track object state through the type system

### H.2 Compilation and Performance

1. **Whole-Program Optimization**: Cross-module optimizations
2. **Profile-Guided Optimization**: Optimize based on runtime behavior
3. **Heterogeneous Computing**: GPU and specialized hardware support
4. **Incremental Compilation**: Faster development cycles
5. **WASM Compilation**: First-class WebAssembly support

### H.3 Development Experience

1. **Language Server Protocol**: Rich IDE integration
2. **Interactive Programming**: REPL and notebook-style development
3. **Improved Error Messages**: More helpful diagnostics
4. **Code Generation**: Macros and compile-time metaprogramming
5. **Hot Code Reloading**: Dynamic code updates without restarting

## Appendix I: LLVM Implementation Timeline

A recommended phased approach for implementing the Unified compiler with LLVM:

### Phase 1: Core Language Features (3-6 months)
1. Lexer and parser
2. Basic type checking
3. Simple code generation for:
   - Basic expressions
   - Control flow
   - Functions
   - Structs
4. Integration with LLVM toolchain

### Phase 2: Memory Management (2-3 months)
1. Ownership system implementation
2. Borrow checking
3. Region-based memory management
4. LLVM IR generation for memory operations

### Phase 3: Type System (2-3 months)
1. Generic type implementation
2. Interface system
3. Type inference
4. Advanced type features

### Phase 4: Concurrency (3-4 months)
1. Actor model implementation
2. Async/await
3. Channels and messaging
4. Work-stealing scheduler

### Phase 5: Standard Library and Tooling (3-6 months)
1. Core standard library
2. Build system
3. Package manager
4. Development tools

### Phase 6: Optimization and Refinement (Ongoing)
1. Performance optimizations
2. Error message improvements
3. Extended standard library
4. Language server implementation