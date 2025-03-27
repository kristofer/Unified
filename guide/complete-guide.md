# Beginner's Guide to Unified

## 1. Introduction to Unified
- **What is Unified?** - A modern systems programming language combining memory safety, performance, and developer ergonomics
- **Philosophy** - Taking the best parts of Rust, Go, Swift, and other languages while making memory management more approachable
- **Use Cases** - Systems programming, application development, server-side applications, and embedded systems
- **Key Advantages** - Memory safety without garbage collection, intuitive ownership model, powerful concurrency primitives, expressive type system

## 2. Getting Started
- **Installation** - Setting up the Unified compiler and toolchain
- **Hello World** - Writing your first Unified program
- **Project Structure** - Understanding the organization of Unified projects
- **Building and Running** - Using the compiler to build and execute programs
- **Package Management** - Working with Unified's package ecosystem

## 3. Basic Syntax and Types
- **Variables and Constants** - Declaring with `let`, `let mut`, and `const`
- **Basic Types** - Int, Float, Bool, String, and more
- **Type Inference** - How Unified determines types automatically
- **Operators** - Arithmetic, comparison, logical, and other operators
- **String Interpolation** - Using `${...}` for string formatting
- **Comments** - Single-line `//` and multi-line `/* */`

## 4. Control Flow
- **If Expressions** - Conditional logic with expression syntax
- **Loops** - For loops, while loops, and loop control
- **Switch and Pattern Matching** - Powerful pattern matching with the `switch` keyword
- **Early Returns** - Using `return` to exit functions early
- **Error Handling Flow** - Using `?` for error propagation

## 5. Functions and Closures
- **Function Declaration** - Using the `fn` keyword
- **Parameters and Return Types** - Type annotations and multiple returns
- **Default Parameters** - Specifying default values
- **Expression Syntax** - Implicit returns for single-expression functions
- **Closures** - Using `||` syntax for anonymous functions
- **Higher-order Functions** - Functions that take or return other functions

## 6. Data Structures
- **Tuples** - Ordered groups of values with mixed types
- **Lists** - Dynamic arrays with common operations
- **Maps** - Key-value collections
- **Sets** - Collections of unique values
- **Structs** - Defining custom data structures
- **Enums** - Creating algebraic data types with variants

## 7. Ownership and Borrowing
- **Understanding Ownership** - The core memory management concept
- **The `move` Keyword** - Explicit ownership transfer
- **Borrowing** - Using `&` for references
- **Mutable Borrowing** - Using `&mut` for mutable references
- **Lifetimes** - Implicit lifetime management
- **Copy vs. Move Types** - Understanding which types are copied vs. moved

## 8. Error Handling
- **The Result Type** - Using `Result<T, E>` for operations that can fail
- **Option Type** - Using `Option<T>` for values that might be absent
- **Error Propagation** - Using the `?` operator
- **Pattern Matching Errors** - Using `match` for comprehensive error handling
- **Creating Custom Errors** - Defining your own error types

## 9. Modules and Imports
- **Module System** - Organizing code with the `module` keyword
- **Importing** - Using the `import` statement
- **Visibility** - Public (`pub`) and private items
- **Namespacing** - Working with qualified paths
- **Standard Library** - Overview of key modules and functions

## 10. Object-Oriented Programming in Unified
- **Methods** - Adding behavior to structs
- **Self Parameter** - Using `self`, `&self`, and `&mut self`
- **Static Methods** - Methods that don't take a self parameter
- **Interfaces** - Using the `impl` keyword for behavior contracts
- **Compositional Design** - Favoring composition over inheritance

## 11. Generic Programming
- **Generic Types** - Creating data structures that work with any type
- **Generic Functions** - Writing flexible, reusable algorithms
- **Constraints** - Specifying requirements with `where` clauses
- **Associated Types** - Using type placeholders in interfaces
- **Type Inference** - How the compiler resolves generic types

## 12. Concurrency
- **Actors** - Isolated, message-passing concurrent entities
- **Channels** - Thread-safe communication primitives
- **Async/Await** - Writing non-blocking asynchronous code
- **Futures** - Working with asynchronous computations
- **Mutexes and Atomics** - Low-level synchronization primitives

## 13. Advanced Types
- **Region-based Memory** - Using the `region` keyword for efficient memory allocation
- **Optional GC** - Opting into garbage collection with `@gc`
- **Type Aliases** - Creating readable type shortcuts
- **Union Types** - Representing values that could be one of several types
- **Phantom Types** - Using the type system for extra safety guarantees

## 14. Pattern Matching and Destructuring
- **Basic Patterns** - Matching literals, variables, and wildcards
- **Destructuring** - Breaking down complex data types
- **Guards** - Adding conditions to pattern matches
- **Exhaustiveness** - Ensuring all cases are handled
- **Custom Patterns** - Creating your own pattern matching rules

## 15. Metaprogramming
- **Compile-time Evaluation** - Computing values at compile time
- **Code Generation** - Creating code dynamically
- **Attributes** - Adding metadata to code elements
- **Reflection** - Inspecting types at runtime
- **Custom DSLs** - Building domain-specific languages

## 16. Testing and Documentation
- **Unit Tests** - Writing and running tests
- **Test Attributes** - Configuring test behavior
- **Documentation Comments** - Creating rich documentation
- **Examples as Tests** - Testing documentation examples
- **Benchmarking** - Measuring performance

## 17. Interoperability
- **FFI** - Interfacing with C and other languages
- **Platform-specific Code** - Conditional compilation
- **LLVM Integration** - Leveraging the LLVM ecosystem
- **Embedding Unified** - Using Unified as a scripting language
- **WebAssembly** - Compiling to the web

## 18. Performance Optimization
- **Compiler Optimizations** - Understanding the optimization pipeline
- **Profiling** - Finding performance bottlenecks
- **Memory Layout** - Controlling data representations
- **Inline Assembly** - Using platform-specific instructions
- **SIMD Support** - Vectorized operations

## 19. Advanced Memory Management
- **Custom Allocators** - Controlling memory allocation strategies
- **Reference Counting** - Using `Rc<T>` for shared ownership
- **Weak References** - Breaking reference cycles
- **Unsafe Code** - Escaping the ownership system when necessary
- **Memory Mapping** - Working with memory-mapped files

## 20. Best Practices and Idioms
- **Coding Style** - Unified's recommended practices
- **API Design** - Creating intuitive interfaces
- **Error Handling Strategies** - Choosing the right approach
- **Performance Considerations** - Writing efficient Unified code
- **Migrating from Other Languages** - Transitioning existing projects
# 3. Basic Syntax and Types

## Variables and Constants

In Unified, you declare variables using the `let` keyword. By default, variables are immutable, which helps prevent accidental changes and makes your code safer.

```unified
// Immutable variable declaration
let name = "Alice"
// name = "Bob"  // This would cause a compiler error
```

To create a mutable variable, use the `let mut` syntax:

```unified
// Mutable variable
let mut counter = 0
counter = counter + 1  // This is valid
```

For values that never change, use the `const` keyword. Constants must have an explicit type and are evaluated at compile time:

```unified
// Constant declaration
const MAX_USERS: Int = 1000
const PI: Float = 3.14159
```

## Basic Types

Unified provides several built-in primitive types:

```unified
// Integer types
let i: Int = 42       // Default integer, platform-dependent size
let i8: Int8 = 127    // 8-bit signed integer
let i16: Int16 = 32767 // 16-bit signed integer
let i32: Int32 = 2147483647 // 32-bit signed integer
let i64: Int64 = 9223372036854775807 // 64-bit signed integer

// Unsigned integers
let ui: UInt = 42     // Default unsigned integer
let ui8: UInt8 = 255  // 8-bit unsigned integer
// UInt16, UInt32, UInt64 also available

// Floating-point types
let f: Float = 3.14   // Default floating-point, platform-dependent
let f32: Float32 = 3.14 // 32-bit floating-point
let f64: Float64 = 3.14159265359 // 64-bit floating-point

// Boolean type
let flag: Bool = true

// Character type
let letter: Char = 'A'

// String type
let message: String = "Hello, Unified!"

// Unit type (similar to void in other languages)
let nothing: Unit = ()
```

## Type Inference

Unified has a powerful type inference system that can usually determine the type of a variable from its initialization, so explicit type annotations are often optional:

```unified
// The compiler infers these types automatically
let age = 30         // Inferred as Int
let name = "Bob"     // Inferred as String
let isActive = true  // Inferred as Bool
let pi = 3.14159     // Inferred as Float
```

However, type annotations can improve readability and are required in certain contexts, such as function parameters and return types.

## Operators

Unified provides familiar operators for various operations:

```unified
// Arithmetic operators
let sum = 5 + 3       // Addition: 8
let difference = 5 - 3 // Subtraction: 2
let product = 5 * 3    // Multiplication: 15
let quotient = 5 / 3   // Division: 1 (integer division)
let remainder = 5 % 3  // Modulo: 2

// For floating-point division:
let floatQuotient = 5.0 / 3.0  // 1.666...

// Compound assignment operators
let mut value = 10
value += 5  // value is now 15
value -= 3  // value is now 12
value *= 2  // value is now 24
value /= 3  // value is now 8

// Comparison operators
let isEqual = 5 == 5       // true
let isNotEqual = 5 != 3    // true
let isGreater = 5 > 3      // true
let isLess = 5 < 10        // true
let isGreaterOrEqual = 5 >= 5 // true
let isLessOrEqual = 5 <= 10   // true

// Logical operators
let and = true && false    // false
let or = true || false     // true
let not = !true            // false

// Range operators
let range = 1..5           // Range from 1 to 4 (exclusive)
let inclusiveRange = 1..=5 // Range from 1 to 5 (inclusive)
```

## String Interpolation

Unified makes it easy to embed expressions within strings using the `${}` syntax:

```unified
let name = "World"
let greeting = "Hello, ${name}!"  // "Hello, World!"

let a = 5
let b = 10
let sum = "${a} + ${b} = ${a + b}"  // "5 + 10 = 15"
```

## Comments

Unified supports both single-line and multi-line comments:

```unified
// This is a single-line comment

/*
   This is a multi-line comment
   that spans multiple lines
*/

let x = 5  // Inline comment after code
```

# 4. Control Flow

## If Expressions

In Unified, `if` is an expression, which means it can return a value. The condition doesn't need parentheses, but the body must be enclosed in braces:

```unified
let number = 42

if number > 0 {
    print("Positive")
} else if number < 0 {
    print("Negative")
} else {
    print("Zero")
}
```

Since `if` is an expression, you can use it to assign values:

```unified
let number = 42
let message = if number > 0 {
    "Positive"
} else {
    "Non-positive"
}
// message is "Positive"
```

## Loops

Unified provides several looping constructs:

### For Loops

The `for` loop iterates over ranges, collections, or iterables:

```unified
// Loop over a range
for i in 0..5 {
    print(i)  // Prints 0, 1, 2, 3, 4
}

// Loop over an inclusive range
for i in 0..=5 {
    print(i)  // Prints 0, 1, 2, 3, 4, 5
}

// Loop over a collection
let names = ["Alice", "Bob", "Charlie"]
for name in names {
    print(name)
}
```

### While Loops

The `while` loop runs as long as a condition is true:

```unified
let mut count = 0
while count < 5 {
    print(count)
    count += 1
}

// Loop with condition and continue/break
let mut sum = 0
let mut i = 0
while i < 100 {
    i += 1
    if i % 2 == 0 {
        continue  // Skip even numbers
    }
    sum += i
    if sum > 1000 {
        break  // Exit the loop early
    }
}
```

### Infinite Loops

Use the `loop` keyword for an infinite loop that you can exit with `break`:

```unified
let mut count = 0
loop {
    print(count)
    count += 1
    if count >= 5 {
        break
    }
}
```

### Loop Labels

You can label loops to specify which loop to break from or continue:

```unified
'outer: for i in 0..5 {
    'inner: for j in 0..5 {
        if i * j > 10 {
            break 'outer  // Exits the outer loop
        }
        print("${i}, ${j}")
    }
}
```

## Switch and Pattern Matching

Unified's `switch` statement is a powerful pattern matching tool:

```unified
let value = 42

switch value {
    case 0 -> print("Zero")
    case 1..5 -> print("Small")
    case 6..10 -> print("Medium")
    case let x if x > 10 && x < 100 -> print("Large: ${x}")
    case _ -> print("Huge")
}
```

Pattern matching can also be used with enums and other complex types:

```unified
enum Result<T, E> {
    Success(T)
    Error(E)
}

let result: Result<Int, String> = Result.Success(42)

switch result {
    case Success(let value) -> print("Success: ${value}")
    case Error(let message) -> print("Error: ${message}")
}
```

## Early Returns

You can exit a function early using the `return` keyword:

```unified
fn isPositive(number: Int) -> Bool {
    if number <= 0 {
        return false
    }
    
    return true
}
```

## Error Handling Flow

Unified uses the `?` operator for concise error propagation:

```unified
fn readConfig() -> Result<Config, Error> {
    let file = File.open("config.toml")?  // Returns early if Error
    let content = file.readToString()?    // Returns early if Error
    let config = parse(content)?          // Returns early if Error
    
    return Success(config)
}
```

The `?` operator unwraps a `Success` value or returns the `Error` early from the function. This allows for clean error handling without deeply nested if statements or verbose error checks.

# 5. Functions and Closures

## Function Declaration

Functions in Unified are declared using the `fn` keyword. A basic function looks like this:

```unified
fn greet() {
    print("Hello, world!")
}
```

To call a function, use its name followed by parentheses:

```unified
greet()  // Prints "Hello, world!"
```

## Parameters and Return Types

Functions can take parameters and return values. Parameter types must be explicitly declared, and the return type is specified after an arrow (`->`):

```unified
fn add(a: Int, b: Int) -> Int {
    return a + b
}

let sum = add(3, 5)  // sum = 8
```

If a function doesn't return a value, you can either omit the return type or use the `Unit` type:

```unified
fn logMessage(message: String) {
    print("LOG: ${message}")
}

// Equivalent to:
fn logMessage(message: String) -> Unit {
    print("LOG: ${message}")
}
```

## Multiple Return Values

Unified allows returning multiple values using tuples:

```unified
fn divideWithRemainder(a: Int, b: Int) -> (Int, Int) {
    let quotient = a / b
    let remainder = a % b
    return (quotient, remainder)
}

let (quot, rem) = divideWithRemainder(10, 3)  // quot = 3, rem = 1
```

You can also use named tuples for clarity:

```unified
fn findMinMax(values: List<Int>) -> (min: Int, max: Int) {
    let min = values.reduce(values[0], |acc, value| if value < acc { value } else { acc })
    let max = values.reduce(values[0], |acc, value| if value > acc { value } else { acc })
    return (min: min, max: max)
}

let result = findMinMax([3, 1, 5, 2, 4])
print("Min: ${result.min}, Max: ${result.max}")  // "Min: 1, Max: 5"
```

## Default Parameters

Unified supports default parameter values, making certain arguments optional:

```unified
fn greet(name: String, greeting: String = "Hello") -> String {
    return "${greeting}, ${name}!"
}

let msg1 = greet("Alice")             // "Hello, Alice!"
let msg2 = greet("Bob", "Welcome")    // "Welcome, Bob!"
```

## Expression Syntax

In Unified, functions are expressions, and the last expression in a function body automatically becomes the return value, making the `return` keyword optional in many cases:

```unified
// Return keyword can be omitted for single-expression functions
fn square(x: Int) -> Int {
    x * x  // Implicit return
}

// This is equivalent to:
fn square(x: Int) -> Int {
    return x * x
}
```

This expression-oriented syntax makes for cleaner code, especially for simple functions.

## Closures

Closures (or lambda expressions) are anonymous functions that can capture variables from their surrounding scope. They use the `||` syntax for parameters:

```unified
// Simple closure with no parameters
let sayHello = || {
    print("Hello!")
}
sayHello()  // Prints "Hello!"

// Closure with parameters
let add = |a: Int, b: Int| {
    a + b
}
let sum = add(5, 3)  // sum = 8

// Type inference for closure parameters
let multiply = |a, b| a * b  // Types inferred from context
let product = multiply(4, 7)  // product = 28
```

## Capturing Variables

Closures can capture variables from their surrounding scope:

```unified
let factor = 2
let double = |x| x * factor  // Captures 'factor' from surrounding scope

let result = double(5)  // result = 10
```

By default, closures capture variables by reference. Use the `move` keyword to take ownership of captured variables:

```unified
let name = "Alice"
let greeting = move || {
    print("Hello, ${name}!")
}  // 'name' is moved into the closure

// name is no longer valid here if String is not a Copy type
```

## Higher-order Functions

Unified treats functions as first-class citizens, meaning they can be passed as arguments to other functions or returned from functions:

```unified
// Function that takes a function as a parameter
fn applyTwice(x: Int, f: fn(Int) -> Int) -> Int {
    return f(f(x))
}

fn double(x: Int) -> Int {
    return x * 2
}

let result = applyTwice(3, double)  // result = 12 (double(double(3)))
```

Functions can also return other functions:

```unified
fn makeMultiplier(factor: Int) -> fn(Int) -> Int {
    return |x| x * factor
}

let triple = makeMultiplier(3)
let result = triple(4)  // result = 12
```

## Function Overloading

Unified supports function overloading based on parameter types:

```unified
fn process(value: Int) -> String {
    return "Processing integer: ${value}"
}

fn process(value: String) -> String {
    return "Processing string: ${value}"
}

let result1 = process(42)        // "Processing integer: 42"
let result2 = process("Hello")   // "Processing string: Hello"
```

# 6. Data Structures

## Tuples

Tuples group multiple values of different types into a single compound value:

```unified
// Creating a tuple
let person = ("Alice", 30, true)  // (String, Int, Bool)

// Accessing tuple elements by index
let name = person.0    // "Alice"
let age = person.1     // 30
let isActive = person.2  // true

// Destructuring tuples
let (userName, userAge, userStatus) = person
print(userName)  // "Alice"
```

Named tuples provide better readability by assigning names to elements:

```unified
// Named tuple
let user = (name: "Bob", age: 25, active: false)

// Access by name
print(user.name)    // "Bob"
print(user.age)     // 25
print(user.active)  // false
```

## Lists

Lists (similar to arrays or vectors in other languages) are ordered collections of elements with the same type:

```unified
// Creating a list
let numbers = [1, 2, 3, 4, 5]

// Accessing elements by index
let first = numbers[0]  // 1

// Getting the length
let count = numbers.len()  // 5

// Lists can be mutable
let mut scores = [75, 80, 85]
scores[1] = 90  // Now [75, 90, 85]

// Common operations
scores.push(95)      // Add to the end: [75, 90, 85, 95]
let last = scores.pop()  // Remove from the end: last = 95, scores = [75, 90, 85]
scores.insert(0, 70)  // Insert at index: [70, 75, 90, 85]
scores.remove(2)      // Remove at index: [70, 75, 85]
```

List comprehensions provide a concise way to create new lists:

```unified
// List comprehension
let squares = [x * x for x in 1..6]  // [1, 4, 9, 16, 25]

// With conditional filtering
let evenSquares = [x * x for x in 1..10 if x % 2 == 0]  // [4, 16, 36, 64]
```

## Maps

Maps (or dictionaries) store key-value pairs:

```unified
// Creating a map
let scores = {
    "Alice": 92,
    "Bob": 85,
    "Charlie": 78
}

// Accessing values
let aliceScore = scores["Alice"]  // 92

// Checking if a key exists
let hasDave = scores.hasKey("Dave")  // false

// Modifying a mutable map
let mut inventory = {
    "apple": 5,
    "banana": 8
}

inventory["apple"] = 3        // Update value
inventory["orange"] = 10      // Add new key-value pair
inventory.remove("banana")    // Remove entry
```

Maps are often implemented as hash maps for O(1) access time on average.

## Sets

Sets are collections of unique values:

```unified
// Creating a set
let colors = {"red", "green", "blue"}

// Checking membership
let hasRed = colors.contains("red")      // true
let hasYellow = colors.contains("yellow")  // false

// Set operations
let primary = {"red", "blue", "yellow"}
let secondary = {"green", "orange", "purple"}

let union = primary.union(secondary)
let intersection = primary.intersection(colors)
let difference = primary.difference(colors)
```

## Structs

Structs allow you to create custom data types by grouping related values:

```unified
// Struct definition
struct Point {
    x: Float
    y: Float
}

// Creating a struct instance
let origin = Point{x: 0.0, y: 0.0}
let point = Point{x: 3.5, y: 2.5}

// Accessing fields
let xCoord = point.x  // 3.5
```

Structs can have methods:

```unified
struct Rectangle {
    width: Float
    height: Float
    
    // Method with self parameter
    fn area(self) -> Float {
        return self.width * self.height
    }
    
    // Method that modifies the struct
    fn scale(self: &mut Self, factor: Float) {
        self.width *= factor
        self.height *= factor
    }
    
    // Static method (no self parameter)
    fn square(size: Float) -> Rectangle {
        return Rectangle{width: size, height: size}
    }
}

// Using methods
let rect = Rectangle{width: 10.0, height: 5.0}
let area = rect.area()  // 50.0

let mut mutableRect = Rectangle{width: 4.0, height: 3.0}
mutableRect.scale(2.0)  // Now 8.0 x 6.0

// Using static method
let square = Rectangle.square(5.0)  // 5.0 x 5.0
```

## Enums

Enums define a type that can have a fixed set of variants:

```unified
// Simple enum
enum Color {
    Red
    Green
    Blue
}

let selectedColor = Color.Red

// Pattern matching with enums
switch selectedColor {
    case Red -> print("Roses are red")
    case Green -> print("Grass is green")
    case Blue -> print("Sky is blue")
}
```

Enums can carry additional data:

```unified
// Enum with associated values
enum Shape {
    Circle(radius: Float)
    Rectangle(width: Float, height: Float)
    Triangle(base: Float, height: Float)
}

let shape = Shape.Circle(radius: 5.0)

// Pattern matching with data extraction
fn calculateArea(shape: Shape) -> Float {
    switch shape {
        case Circle(let radius) -> 3.14159 * radius * radius
        case Rectangle(let width, let height) -> width * height
        case Triangle(let base, let height) -> 0.5 * base * height
    }
}

let area = calculateArea(shape)  // ~78.54
```

## Generic Data Structures

You can create data structures that work with any type using generics:

```unified
// Generic struct
struct Pair<A, B> {
    first: A
    second: B
}

let pair = Pair{first: 42, second: "Answer"}

// Generic enum (like Result and Option)
enum Option<T> {
    Some(T)
    None
}

let foundValue: Option<Int> = Option.Some(42)
let emptyValue: Option<Int> = Option.None
```

Many of Unified's standard library collections are generic, allowing them to work with any type that meets their requirements.

# 7. Ownership and Borrowing

Unified's ownership system is one of its most distinctive features, providing memory safety without garbage collection. This system allows Unified to prevent memory errors like use-after-free, double-free, and data races at compile time.

## Understanding Ownership

Every value in Unified has a single owner at any given time. When the owner goes out of scope, the value is automatically deallocated.

```unified
{
    let name = String.from("Unified")  // 'name' owns this string
    // use name...
}  // 'name' goes out of scope here, the string is freed automatically
```

When you assign a value to another variable or pass it to a function, ownership may transfer depending on the type:

```unified
// For types that don't implement Copy:
let s1 = String.from("hello")
let s2 = s1                 // Ownership moves from s1 to s2
// print(s1)                // Error: s1 is no longer valid

// For types that implement Copy (e.g., primitives):
let n1 = 5
let n2 = n1                 // n1 is copied to n2
print(n1)                   // Still valid, because Int is Copy
```

## The `move` Keyword

Unified uses an explicit `move` keyword to make ownership transfers more obvious and intentional:

```unified
let s1 = String.from("hello")
let s2 = move s1           // Explicit transfer of ownership
// print(s1)               // Error: s1 is no longer valid

fn takeOwnership(s: String) {
    // Function takes ownership of the parameter
    print(s)
}

let message = String.from("hello")
takeOwnership(move message)  // Explicitly transfer ownership to the function
// print(message)            // Error: message is no longer valid
```

Without the `move` keyword, the compiler will infer what should happen based on the context, often choosing to borrow rather than transfer ownership.

## Borrowing

Instead of transferring ownership, you can borrow a value using references. A reference allows you to refer to a value without taking ownership of it.

```unified
fn calculate_length(s: &String) -> Int {
    s.len()  // s is a reference to a String
}

let message = String.from("hello")
let length = calculate_length(&message)  // Borrow message
print(message)  // Still valid because we only borrowed it
```

By default, references are immutable, meaning you can't modify the borrowed value.

## Mutable Borrowing

To modify a borrowed value, you need a mutable reference:

```unified
fn append_world(s: &mut String) {
    s.append(" world")
}

let mut message = String.from("hello")
append_world(&mut message)   // Mutable borrow
print(message)               // "hello world"
```

Unified enforces these borrowing rules at compile time:

1. You can have any number of immutable references (`&T`) to a value at a time, OR
2. You can have exactly one mutable reference (`&mut T`) at a time.

This prevents data races:

```unified
let mut value = String.from("hello")

let r1 = &value     // Immutable borrow
let r2 = &value     // Another immutable borrow - OK
// let rm = &mut value  // Error: Cannot borrow as mutable when already borrowed as immutable

print(r1)
print(r2)

// After last use of immutable borrows
let rm = &mut value  // Now OK to borrow as mutable
rm.append(" world")
```

## Lifetimes

Unified manages lifetimes mostly automatically, but there are cases where you need to be explicit about how long references should live:

```unified
// Function that returns a reference needs a lifetime annotation
fn longest<'a>(x: &'a String, y: &'a String) -> &'a String {
    if x.len() > y.len() {
        x
    } else {
        y
    }
}
```

The lifetime annotation `'a` tells the compiler that the returned reference will be valid as long as both input references are valid.

## Copy vs. Move Types

Types in Unified fall into two categories:

1. **Copy types**: Simple types stored on the stack that are inexpensive to copy. These include primitives like `Int`, `Float`, `Bool`, and tuples containing only Copy types.

2. **Move types**: Types that contain heap-allocated data or are otherwise expensive to copy. These include `String`, `List<T>`, custom structs, etc.

```unified
// Copy types are duplicated on assignment
let a = 5
let b = a  // 'a' is copied to 'b'
// Both 'a' and 'b' are valid and independent

// Move types transfer ownership on assignment, unless borrowed
let s1 = String.from("hello")
let s2 = move s1  // Ownership transferred from 's1' to 's2'
// 's1' is no longer valid

// To copy a Move type, use an explicit clone
let s1 = String.from("hello")
let s2 = s1.clone()  // Creates a deep copy
// Both 's1' and 's2' are valid and independent
```

You can make custom types Copy by implementing the Copy interface:

```unified
struct Point {
    x: Int
    y: Int
}

impl Copy for Point {}  // Now Point is a Copy type
```

## Region-based Memory Management

For more sophisticated memory management needs, Unified provides region-based allocation:

```unified
region temp {
    let large_data = readLargeFile()  // Allocated in the temporary region
    process(large_data)
}  // Entire region freed at once
```

This allows efficient allocation and deallocation of multiple related objects without the complexity of manual memory management.

# 8. Error Handling

Unified provides a robust system for handling errors through the `Result` and `Option` types, eliminating the need for exceptions while maintaining type safety.

## The Result Type

The `Result` type represents an operation that might fail:

```unified
enum Result<T, E> {
    Success(T)   // Contains the success value
    Error(E)     // Contains the error value
}
```

Functions that can fail return a `Result`:

```unified
fn divide(a: Int, b: Int) -> Result<Int, String> {
    if b == 0 {
        return Error("Division by zero")
    }
    return Success(a / b)
}
```

You can handle a `Result` using pattern matching:

```unified
let result = divide(10, 2)

switch result {
    case Success(let value) -> print("Result: ${value}")
    case Error(let message) -> print("Error: ${message}")
}
```

## Option Type

The `Option` type represents a value that might be absent:

```unified
enum Option<T> {
    Some(T)  // Contains a value
    None     // Represents no value
}
```

This eliminates the need for null/nil references:

```unified
fn findUser(id: UserId) -> Option<User> {
    // Look up user in database
    if userExists(id) {
        return Some(getUserById(id))
    } else {
        return None
    }
}
```

You can handle an `Option` using pattern matching:

```unified
let user = findUser(userId)

switch user {
    case Some(let foundUser) -> displayUserInfo(foundUser)
    case None -> displayUserNotFound()
}
```

## Error Propagation

The `?` operator provides a concise way to propagate errors up the call stack:

```unified
fn readConfig() -> Result<Config, Error> {
    let file = File.open("config.toml")?  // Returns early if Error
    let content = file.readToString()?    // Returns early if Error
    let config = parse(content)?          // Returns early if Error
    
    return Success(config)
}
```

When applied to a `Result`, the `?` operator:
1. If the `Result` is `Success(value)`, extracts the `value`
2. If the `Result` is `Error(e)`, returns early from the function with `Error(e)`

Similarly, for `Option`:

```unified
fn username(id: UserId) -> Option<String> {
    let user = findUser(id)?         // Returns None if user not found
    let profile = user.getProfile()? // Returns None if profile not found
    return Some(profile.username)
}
```

## Pattern Matching Errors

Unified's pattern matching system allows for detailed error handling:

```unified
enum NetworkError {
    ConnectionFailed(String)
    Timeout(Int)
    InvalidResponse(StatusCode)
}

let result = fetchData(url)

switch result {
    case Success(let data) -> processData(data)
    case Error(ConnectionFailed(let reason)) -> retryConnection(reason)
    case Error(Timeout(let seconds)) -> {
        if seconds < 30 {
            retryWithLongerTimeout()
        } else {
            abandonRequest()
        }
    }
    case Error(InvalidResponse(let status)) -> reportBadStatus(status)
}
```

## Creating Custom Errors

You can define your own error types to provide rich, domain-specific error information:

```unified
// Define a custom error type
enum FileError {
    NotFound(String)           // File not found (with path)
    PermissionDenied(String)   // Permission error (with path)
    Corrupt(String, String)    // Corrupt file (path and details)
}

// Function that uses the custom error type
fn readConfig(path: String) -> Result<Config, FileError> {
    if !fileExists(path) {
        return Error(NotFound(path))
    }
    
    if !hasPermission(path) {
        return Error(PermissionDenied(path))
    }
    
    let content = readFile(path)
    if !isValidConfig(content) {
        return Error(Corrupt(path, "Invalid config format"))
    }
    
    return Success(parseConfig(content))
}
```

## Combining Result and Option

Unified provides utilities to convert between `Result` and `Option`:

```unified
// Convert Option to Result with a default error
let opt: Option<Int> = None
let res: Result<Int, String> = opt.okOr("Value was None")

// Convert Result to Option (discarding the error)
let res: Result<Int, String> = Error("Failed")
let opt: Option<Int> = res.ok()
```

## The `try` Block

For multiple operations that might fail, you can use a `try` block:

```unified
fn processData() -> Result<Output, ProcessError> {
    try {
        let file = openFile("data.txt")?
        let content = file.readToString()?
        let data = parseJson(content)?
        let result = transformData(data)?
        return Success(result)
    }
}
```

This creates a scope where `?` can be used, and automatically wraps the return value in `Success` if all operations succeed.

## Default Values and Fallbacks

Unified provides methods to handle errors with default values:

```unified
// For Option
let value = findUser(id).unwrapOr(defaultUser)
let name = user.getName().unwrapOrElse(|| generateRandomName())

// For Result
let result = fetchData().unwrapOr(cachedData)
let config = parseConfig().unwrapOrElse(|err| {
    logError(err)
    return defaultConfig()
})
```

The error handling system in Unified is designed to make errors explicit and impossible to ignore, while still being ergonomic to use. This leads to more robust and maintainable code.
# 9. Modules and Imports

Unified's module system helps you organize code into logical units, manage namespaces, and control visibility. It provides a clean way to structure larger applications and share code between projects.

## Module System

A module is a collection of items such as functions, structs, traits, and other modules. The `module` keyword (or its shorthand `mod`) defines a module:

```unified
// Define a module
module math {
    // Public function, accessible outside the module
    pub fn add(a: Int, b: Int) -> Int {
        return a + b
    }
    
    // Private function, only accessible within the module
    fn is_positive(x: Int) -> Bool {
        return x > 0
    }
    
    // Nested module
    pub module advanced {
        pub fn square_root(x: Float) -> Float {
            // Implementation
            return x.sqrt()
        }
    }
}
```

## Importing

The `import` statement brings items from other modules into scope:

```unified
// Import specific items
import math.add
import math.advanced.square_root

// Use imported functions
let sum = add(5, 3)
let sqrt = square_root(16.0)
```

You can import multiple items at once:

```unified
// Import multiple items
import math.{add, subtract, multiply}

// Import all public items from a module
import math.*
```

You can also give imports aliases to avoid name conflicts:

```unified
// Import with alias
import math.add as math_add
import stats.add as stats_add

let sum1 = math_add(2, 3)
let sum2 = stats_add([1, 2, 3, 4])
```

## Visibility

By default, items in Unified are private to their containing module. Use the `pub` keyword to make them public:

```unified
module shapes {
    // Public struct
    pub struct Circle {
        // Public field
        pub radius: Float
        
        // Private field
        center_x: Float
        center_y: Float
        
        // Public method
        pub fn area(self) -> Float {
            return 3.14159 * self.radius * self.radius
        }
        
        // Private method
        fn diameter(self) -> Float {
            return 2.0 * self.radius
        }
    }
    
    // Private function
    fn calculate_circumference(circle: &Circle) -> Float {
        return 2.0 * 3.14159 * circle.radius
    }
}
```

Visibility can be more granular with module paths:

```unified
module network {
    // Only visible to the 'server' module
    pub(server) fn internal_connect() {
        // Implementation
    }
    
    // Public to direct parent module only
    pub(super) fn status_check() {
        // Implementation
    }
    
    // Public to the entire crate
    pub(crate) fn log_connection() {
        // Implementation
    }
}
```

## Namespacing

Modules provide namespaces to avoid name collisions:

```unified
module graphics {
    pub struct Point {
        pub x: Int
        pub y: Int
    }
}

module geography {
    pub struct Point {
        pub latitude: Float
        pub longitude: Float
    }
}

// Use fully qualified paths to disambiguate
let screen_point = graphics.Point{x: 10, y: 20}
let map_point = geography.Point{latitude: 40.7, longitude: -74.0}
```

## File Organization

In larger projects, modules often correspond to files and directories:

```
project/
├── main.un               # Main file of your project
├── math/                 # Directory for the math module
│   ├── mod.un            # Entry point for the math module
│   ├── basic.un          # Basic math operations
│   └── advanced.un       # Advanced math operations
└── utils/                # Directory for utilities
    ├── mod.un            # Entry point for the utils module
    └── logging.un        # Logging utilities
```

You can reference these modules in your code:

```unified
// In main.un
import math.{add, subtract}
import math.advanced.square_root
import utils.logging.{log, debug}

fn main() {
    log("Starting application...")
    
    let result = add(square_root(16.0), 5)
    debug("Result: ${result}")
}
```

## Standard Library

Unified comes with a rich standard library organized into modules:

```unified
// Common standard library imports
import std.io.{print, read_line}
import std.collections.{List, Map, Set}
import std.math.{sin, cos, tan}
import std.time.{now, sleep}
```

The standard library provides essential functionality for common tasks, organized logically into modules like:

- `std.io`: Input/output operations
- `std.collections`: Data structures like List, Map, Set
- `std.math`: Mathematical functions
- `std.string`: String manipulation
- `std.time`: Time-related functions
- `std.thread`: Threading and concurrency
- `std.net`: Networking operations
- `std.fs`: File system operations

# 10. Object-Oriented Programming in Unified

While Unified isn't a traditional object-oriented language, it provides powerful features for organizing code around data and behavior. Unified takes a composition-focused approach rather than relying on inheritance hierarchies.

## Methods

You can add methods to structs to create object-like behavior:

```unified
struct Rectangle {
    width: Float
    height: Float
    
    // Method with self parameter
    fn area(self) -> Float {
        return self.width * self.height
    }
    
    // Method that modifies the struct
    fn scale(self: &mut Self, factor: Float) {
        self.width *= factor
        self.height *= factor
    }
}

// Using methods
let rect = Rectangle{width: 10.0, height: 5.0}
let area = rect.area()  // 50.0

let mut mutable_rect = Rectangle{width: 4.0, height: 3.0}
mutable_rect.scale(2.0)  // Now 8.0 x 6.0
```

## Self Parameter

Methods can take various forms of the `self` parameter:

```unified
struct Counter {
    count: Int
    
    // Takes the struct by value (moves or copies it)
    fn consume(self) -> Int {
        return self.count
    }
    
    // Borrows the struct immutably
    fn get(self: &Self) -> Int {
        return self.count
    }
    
    // Borrows the struct mutably
    fn increment(self: &mut Self) {
        self.count += 1
    }
    
    // Static method (no self parameter)
    fn new() -> Self {
        return Self{count: 0}
    }
}
```

The `Self` type is an alias for the type the method is defined on, which makes code more maintainable and reusable.

## Static Methods

Static methods don't take a `self` parameter and are associated with the type rather than an instance:

```unified
struct Point {
    x: Float
    y: Float
    
    // Static method to create a new point
    fn new(x: Float, y: Float) -> Self {
        return Self{x: x, y: y}
    }
    
    // Static method to create the origin point
    fn origin() -> Self {
        return Self{x: 0.0, y: 0.0}
    }
    
    // Instance method
    fn distance(self, other: &Point) -> Float {
        let dx = self.x - other.x
        let dy = self.y - other.y
        return (dx * dx + dy * dy).sqrt()
    }
}

// Using static methods
let origin = Point.origin()
let point = Point.new(3.0, 4.0)
let dist = point.distance(&origin)  // 5.0
```

## Interfaces

Unified uses interfaces (defined with the `interface` keyword) to define shared behavior that different types can implement:

```unified
// Define an interface
interface Shape {
    // Method signature without implementation
    fn area(self) -> Float
    
    // Method signature without implementation
    fn perimeter(self) -> Float
    
    // Method with default implementation
    fn describe(self) -> String {
        return "A shape with area ${self.area()} and perimeter ${self.perimeter()}"
    }
}

// Implement the interface for a struct
struct Circle {
    radius: Float
    
    fn new(radius: Float) -> Self {
        return Self{radius: radius}
    }
}

impl Shape for Circle {
    fn area(self) -> Float {
        return 3.14159 * self.radius * self.radius
    }
    
    fn perimeter(self) -> Float {
        return 2.0 * 3.14159 * self.radius
    }
    
    // We can override default methods too
    fn describe(self) -> String {
        return "A circle with radius ${self.radius}"
    }
}

// Implement for another struct
struct Rectangle {
    width: Float
    height: Float
}

impl Shape for Rectangle {
    fn area(self) -> Float {
        return self.width * self.height
    }
    
    fn perimeter(self) -> Float {
        return 2.0 * (self.width + self.height)
    }
}

// Function that works with any Shape
fn print_shape_info(shape: &impl Shape) {
    print("Area: ${shape.area()}")
    print("Perimeter: ${shape.perimeter()}")
    print("Description: ${shape.describe()}")
}

// Usage
let circle = Circle.new(5.0)
let rectangle = Rectangle{width: 4.0, height: 3.0}

print_shape_info(&circle)
print_shape_info(&rectangle)
```

## Compositional Design

Unified encourages composition over inheritance:

```unified
// Composition example
struct Engine {
    horsepower: Int
    
    fn start(self) {
        print("Engine started with ${self.horsepower} HP")
    }
}

struct Wheels {
    count: Int
    size: Float
    
    fn rotate(self) {
        print("${self.count} wheels rotating")
    }
}

struct Car {
    engine: Engine
    wheels: Wheels
    brand: String
    
    fn drive(self) {
        self.engine.start()
        self.wheels.rotate()
        print("${self.brand} car is driving")
    }
}

// Creating a car through composition
let car = Car{
    engine: Engine{horsepower: 150},
    wheels: Wheels{count: 4, size: 17.0},
    brand: "Unified"
}

car.drive()
```

## Method Chaining

You can design fluent interfaces with method chaining:

```unified
struct Builder {
    value: Int
    
    fn new() -> Self {
        return Self{value: 0}
    }
    
    fn add(self: &mut Self, x: Int) -> &mut Self {
        self.value += x
        return self
    }
    
    fn multiply(self: &mut Self, x: Int) -> &mut Self {
        self.value *= x
        return self
    }
    
    fn build(self) -> Int {
        return self.value
    }
}

// Method chaining
let result = Builder.new()
    .add(5)
    .multiply(2)
    .add(10)
    .build()  // result = 20
```

## Encapsulation

Unified provides encapsulation through module privacy and private struct fields:

```unified
module banking {
    pub struct Account {
        // Private fields
        holder_name: String
        balance: Float
        
        // Public constructor
        pub fn new(name: String, initial_balance: Float) -> Self {
            return Self{
                holder_name: name,
                balance: initial_balance
            }
        }
        
        // Public methods for controlled access
        pub fn get_balance(self) -> Float {
            return self.balance
        }
        
        pub fn deposit(self: &mut Self, amount: Float) -> Result<Float, String> {
            if amount <= 0.0 {
                return Error("Deposit amount must be positive")
            }
            
            self.balance += amount
            return Success(self.balance)
        }
        
        pub fn withdraw(self: &mut Self, amount: Float) -> Result<Float, String> {
            if amount <= 0.0 {
                return Error("Withdrawal amount must be positive")
            }
            
            if amount > self.balance {
                return Error("Insufficient funds")
            }
            
            self.balance -= amount
            return Success(self.balance)
        }
    }
    
    // Private helper function
    fn validate_amount(amount: Float) -> Bool {
        return amount > 0.0
    }
}

// Using the encapsulated type
let mut account = banking.Account.new("Alice", 1000.0)
let balance = account.get_balance()  // 1000.0

// Direct access to private fields is prevented
// account.balance = 2000.0  // Error: 'balance' is private

// Must use public methods instead
account.deposit(500.0)
account.withdraw(200.0)
```

While Unified doesn't support traditional inheritance, its combination of interfaces, composition, and powerful type system allows for flexible, maintainable object-oriented designs. This approach encourages delegation over inheritance, and composition over extension, leading to more modular and reusable code.
# 11. Generic Programming

Generic programming allows you to write flexible, reusable code that works with different types while maintaining type safety. Unified's generic system is powerful yet intuitive, enabling you to create highly adaptable libraries and applications.

## Generic Types

You can define generic structs and enums that work with any type:

```unified
// Generic struct with one type parameter
struct Box<T> {
    value: T
    
    fn new(value: T) -> Self {
        return Self{value: value}
    }
    
    fn get(self) -> &T {
        return &self.value
    }
    
    fn update(self: &mut Self, new_value: T) {
        self.value = new_value
    }
}

// Usage with different types
let int_box = Box<Int>.new(42)
let string_box = Box<String>.new("Hello")

// Type inference often works too
let float_box = Box.new(3.14)  // Type inferred as Box<Float>
```

Generic types can have multiple type parameters:

```unified
// Generic struct with two type parameters
struct Pair<A, B> {
    first: A
    second: B
    
    fn swap(self) -> Pair<B, A> {
        return Pair{first: self.second, second: self.first}
    }
}

// Usage
let pair = Pair{first: 1, second: "one"}
let swapped = pair.swap()  // Pair{first: "one", second: 1}
```

Generic enums like `Option<T>` and `Result<T, E>` are fundamental to Unified's type system:

```unified
// Generic enum definition (already built into Unified)
enum Option<T> {
    Some(T)
    None
}

enum Result<T, E> {
    Success(T)
    Error(E)
}
```

## Generic Functions

Functions can also be generic, working with multiple types:

```unified
// Generic function
fn identity<T>(value: T) -> T {
    return value
}

// Usage
let x = identity(42)       // x is Int
let s = identity("hello")  // s is String

// Generic function with multiple type parameters
fn pair<A, B>(a: A, b: B) -> Pair<A, B> {
    return Pair{first: a, second: b}
}

// Usage
let p = pair(5, "five")  // p is Pair<Int, String>
```

## Constraints

Often, you need to require that generic types have certain capabilities. In Unified, you can add constraints using the `where` clause:

```unified
// Generic function with constraint
fn max<T>(a: T, b: T) -> T where T: Ord {
    if a > b {
        return a
    } else {
        return b
    }
}

// Usage
let m1 = max(5, 10)       // Works because Int implements Ord
let m2 = max("a", "b")    // Works because String implements Ord
```

You can specify multiple constraints:

```unified
fn process<T>(value: T) -> String 
    where T: Display,        // Must be displayable as string
          T: Serialize,      // Must be serializable
          T: Clone            // Must be cloneable
{
    let copy = value.clone()
    let serialized = value.serialize()
    return "Processed: ${value}"
}
```

## Associated Types

Interfaces can have associated types, which are placeholder types specified when the interface is implemented:

```unified
// Interface with associated type
interface Iterator {
    type Item;  // Associated type
    
    fn next(self: &mut Self) -> Option<Self::Item>;
    fn has_next(self) -> Bool;
}

// Implementation with specific associated type
struct Counter {
    current: Int
    max: Int
}

impl Iterator for Counter {
    type Item = Int;  // Specify the associated type
    
    fn next(self: &mut Self) -> Option<Int> {
        if self.current < self.max {
            let value = self.current
            self.current += 1
            return Some(value)
        } else {
            return None
        }
    }
    
    fn has_next(self) -> Bool {
        return self.current < self.max
    }
}
```

## Type Inference

Unified's type inference system often allows you to omit explicit type parameters:

```unified
// Without type inference
let values: List<Int> = List<Int>.new()
values.push(1)
values.push(2)

// With type inference
let values = List.new()  // Type inferred as List<Int> when first used
values.push(1)           // Now the compiler knows it's List<Int>
```

Functions can infer their generic parameters from arguments:

```unified
// Generic function
fn first<T>(list: &List<T>) -> Option<&T> {
    if list.isEmpty() {
        return None
    }
    return Some(&list[0])
}

// Call without explicit type parameter
let numbers = [1, 2, 3]
let first_num = first(&numbers)  // Type parameter T inferred as Int
```

## Generic Implementation

You can implement interfaces for generic types:

```unified
// Generic implementation
impl<T> Display for Box<T> where T: Display {
    fn to_string(self) -> String {
        return "Box(${self.value})"
    }
}

// Usage
let box = Box.new(42)
print(box)  // "Box(42)"
```

You can also implement generic interfaces for specific types:

```unified
// Convert any type to Option<T>
interface Into<T> {
    fn into(self) -> T;
}

// Implement for specific conversion
impl Into<Option<Int>> for Int {
    fn into(self) -> Option<Int> {
        return Some(self)
    }
}

impl Into<Option<String>> for String {
    fn into(self) -> Option<String> {
        if self.isEmpty() {
            return None
        }
        return Some(self)
    }
}
```

Generic programming in Unified provides powerful abstractions without sacrificing performance, as the compiler monomorphizes generic code (creates specialized versions for each type used) during compilation.

# 12. Concurrency

Unified's concurrency model combines ideas from several languages to create a safe, efficient system for parallel and asynchronous programming.

## Actors

Actors are the primary high-level abstraction for concurrency in Unified. An actor encapsulates state and behavior, processing messages one at a time:

```unified
// Define an actor
actor Counter {
    // Private state
    var count: Int = 0
    
    // Message handlers
    fn increment() -> Int {
        count += 1
        return count
    }
    
    fn decrement() -> Int {
        count -= 1
        return count
    }
    
    fn get() -> Int {
        return count
    }
    
    fn reset() {
        count = 0
    }
}

// Create an actor instance
let counter = spawn Counter{}

// Send messages to the actor (non-blocking)
counter.increment()  // Message sent asynchronously
counter.increment()
let future_count = counter.get()  // Returns a future

// Await the result
let count = await future_count  // Suspends execution until result is available
print("Count: ${count}")  // "Count: 2"
```

Actors provide automatic synchronization—all message handling is serialized, so you never need to worry about data races within an actor.

## Channels

Channels provide a way for concurrent processes to communicate by passing messages:

```unified
// Create a channel
let channel = Channel<String>.new()  // Unbounded channel
// or
let bounded_channel = Channel<String>.withCapacity(10)  // Bounded channel

// Sender can send messages
fn sender(channel: &Channel<String>) {
    channel.send("Hello")
    channel.send("World")
    channel.close()  // Signal no more messages
}

// Receiver can receive messages
fn receiver(channel: &Channel<String>) {
    while let Some(message) = channel.receive() {
        print("Got: ${message}")
    }
    // Channel is closed when receive() returns None
}

// Create two concurrent tasks
let send_task = async {
    sender(&channel)
}

let receive_task = async {
    receiver(&channel)
}

// Wait for both to complete
await all([send_task, receive_task])
```

Channels can be bounded (with a maximum capacity) or unbounded. Sending to a full bounded channel will block until space is available.

## Async/Await

The `async` and `await` keywords enable writing asynchronous code in a synchronous style:

```unified
// Asynchronous function
async fn fetch_url(url: String) -> Result<String, Error> {
    // This would be non-blocking I/O
    return Success("Content of ${url}")
}

// Using async/await
async fn process() {
    let future1 = fetch_url("https://example.com/data1")
    let future2 = fetch_url("https://example.com/data2")
    
    // Wait for both futures to complete
    let result1 = await future1?
    let result2 = await future2?
    
    print("Results: ${result1}, ${result2}")
}

// Execute the async function
await process()
```

The `async` keyword creates a function that returns a future. The `await` keyword suspends execution until the future completes, without blocking the thread.

## Futures

Futures represent asynchronous computations that will complete in the future:

```unified
// Create a future
let future = Future<Int>.new()

// Complete the future (possibly from another task)
future.complete(42)

// Register a callback
future.onComplete(|value| {
    print("Got value: ${value}")
})

// Chain operations
let transformed = future.map(|value| value * 2)
let filtered = future.filter(|value| value > 0)
let flat_mapped = future.flatMap(|value| anotherAsyncOperation(value))
```

Futures can be composed in various ways:

```unified
// Wait for the first future to complete
let first = Future.any([future1, future2, future3])

// Wait for all futures to complete
let all = Future.all([future1, future2, future3])

// Apply a timeout
let with_timeout = future.timeout(5.seconds)
```

## Select

The `select` construct allows waiting for multiple concurrent operations:

```unified
select {
    case message = channel1.receive() -> handleMessage(message)
    case channel2.send(value) -> print("Sent value")
    case result = future.completed() -> processResult(result)
    case after 1.seconds -> print("Timeout")
}
```

This is useful for implementing timeouts, handling multiple channels, or responding to the first of several events.

## Mutex and Atomic Types

For the rare cases when shared mutable state is needed, Unified provides low-level synchronization primitives:

```unified
// Create a mutex
let counter = Mutex<Int>.new(0)

// Critical section
fn increment(counter: &Mutex<Int>) {
    let mut guard = counter.lock()  // Acquires the lock
    *guard += 1
    // Lock is automatically released when guard goes out of scope
}

// Atomic types for simple operations
let atomic_counter = Atomic<Int>.new(0)
atomic_counter.increment()  // Atomic operation
let value = atomic_counter.get()  // Get current value
atomic_counter.compareAndSwap(value, value + 1)  // Compare and swap
```

## Parallel Iteration

Unified makes it easy to parallelize operations on collections:

```unified
// Sequential map
let squares = numbers.map(|x| x * x)

// Parallel map
let squares = numbers.par_map(|x| x * x)

// Parallel filter
let evens = numbers.par_filter(|x| x % 2 == 0)

// Parallel reduce
let sum = numbers.par_reduce(0, |acc, x| acc + x)
```

The parallel methods automatically distribute work across available CPU cores.

## Thread Management

While Unified's concurrency model is primarily based on lightweight tasks rather than OS threads, you can work directly with threads when needed:

```unified
// Spawn a thread
let handle = Thread.spawn(|| {
    // Thread code
    for i in 0..10 {
        print("Thread: ${i}")
        Thread.sleep(100.milliseconds)
    }
})

// Wait for thread to finish
handle.join()
```

## Work Stealing Scheduler

Under the hood, Unified uses a work-stealing scheduler to efficiently manage thousands of lightweight tasks across available CPU cores. The scheduler automatically balances the workload by allowing idle threads to "steal" tasks from busy threads' queues.

This approach ensures high CPU utilization without the overhead of context switching between many OS threads, similar to Go's goroutine scheduler but with Unified's memory safety guarantees.

The combination of actors, channels, async/await, and the work-stealing scheduler makes Unified ideal for building highly concurrent applications, from network servers to parallel data processing pipelines.
# 13. Advanced Types

Unified's type system goes beyond the basics to provide sophisticated features for creating safe, expressive, and efficient code. This section explores some of the more advanced type capabilities available in the language.

## Region-based Memory

One of Unified's distinctive features is its support for region-based memory management, which allows for efficient allocation and deallocation of related objects:

```unified
// Define a region
region temp {
    // Objects allocated within this region
    let large_data = load_large_file()
    let processed = process_data(large_data)
    compute_results(processed)
}  // All memory from region freed at once

// Multiple regions can exist simultaneously
region long_lived {
    let config = load_configuration()
    
    region request_scope {
        let request_data = fetch_data()
        let response = generate_response(request_data, config)
        send(response)
    }  // request_scope memory freed
    
    // config still accessible here
}  // long_lived memory freed
```

Regions are particularly useful for:
- Server request handling where all allocations for a request can be freed at once
- Large temporary computations that would otherwise require complex ownership tracking
- Performance-critical code sections where allocation patterns are predictable

## Optional GC

In some scenarios, automatic memory management is preferable to ownership tracking. Unified allows opting into garbage collection for specific types:

```unified
// Opt into garbage collection for a type
@gc class Graph {
    nodes: List<Node>
    edges: List<Edge>
    
    // Complex cyclic structure handled by GC
    fn add_edge(self: &mut Self, from: NodeId, to: NodeId) {
        // No need to worry about reference cycles
        self.edges.push(Edge{from, to})
        self.nodes[from].connections.push(to)
        self.nodes[to].connections.push(from)
    }
}

// GC and ownership can coexist
fn process() {
    let owned_data = String.from("Regular ownership")
    
    @gc let graph = Graph.new()
    graph.build_complex_structure()
    
    // owned_data freed by ownership rules
    // graph freed by garbage collector when no longer reachable
}
```

The `@gc` attribute signals to the compiler that objects of this type should be managed by the garbage collector rather than Unified's ownership system.

## Type Aliases

Type aliases create new names for existing types, improving readability and maintainability:

```unified
// Simple type alias
type UserId = Int

// Function that uses the alias
fn get_user(id: UserId) -> Option<User> {
    // Implementation
}

// Type alias with generics
type Result<T> = std.Result<T, Error>

// Using the type alias
fn process_data() -> Result<Data> {
    // Implementation
}

// More complex type alias
type UserMap = Map<UserId, User>
type Handler<T> = fn(Request) -> Result<T>
```

Type aliases are especially useful for:
- Making complex types more readable
- Creating domain-specific vocabulary in your code
- Reducing repetition in generic type signatures

## Union Types

Union types represent values that could be one of several types:

```unified
// Define a union type
type JsonValue = String | Int | Float | Bool | JsonObject | JsonArray | Null

// Using a union type
fn parse_json(input: String) -> JsonValue {
    // Implementation
}

// Handling union types with pattern matching
fn process_value(value: JsonValue) {
    switch value {
        case let s: String -> print("String: ${s}")
        case let i: Int -> print("Integer: ${i}")
        case let f: Float -> print("Float: ${f}")
        case let b: Bool -> print("Boolean: ${b}")
        case let o: JsonObject -> process_object(o)
        case let a: JsonArray -> process_array(a)
        case Null -> print("Null value")
    }
}
```

Union types provide more flexibility than enums when you need to combine existing types without creating wrapper types.

## Phantom Types

Phantom types use the type system to enforce constraints at compile time without any runtime overhead:

```unified
// Define marker types
struct Encrypted {}
struct Plaintext {}

// Generic type with phantom parameter
struct Password<State> {
    data: String
    // State is not used in the struct but affects the type
}

// Functions that operate on specific states
fn encrypt(password: Password<Plaintext>) -> Password<Encrypted> {
    // Implementation
    return Password<Encrypted>{data: encrypted_data}
}

fn verify(password: Password<Encrypted>, input: String) -> Bool {
    // Implementation
}

// Usage
let plain_pass = Password<Plaintext>{data: "secret123"}
let secure_pass = encrypt(plain_pass)

// Type safety prevents mistakes
verify(secure_pass, "attempt")  // OK
// verify(plain_pass, "attempt")  // Error: Expected Password<Encrypted>
```

Phantom types are useful for:
- Enforcing state transitions (like the above example)
- Ensuring protocol compliance
- Preventing logical errors through the type system

## Refinement Types

Refinement types allow you to constrain the values of a type:

```unified
// Define a refinement type
type NonEmptyString = String refine |s| !s.isEmpty()

// Function that uses refinement type
fn process(input: NonEmptyString) {
    // No need to check if empty, guaranteed by type
}

// Using refinement types
let valid: NonEmptyString = "Hello"  // OK
// let invalid: NonEmptyString = ""  // Compile-time error

// Refinement type for numbers
type PositiveInt = Int refine |n| n > 0

// Range-based refinement
type Percentage = Int refine |n| n >= 0 && n <= 100
```

The compiler ensures that values satisfy the refinement predicate at compile time when possible, and inserts runtime checks when necessary.

## Dependent Types

Unified supports limited dependent types, where the type depends on a value:

```unified
// Array with size in type
fn create_matrix<N: Int>(rows: Int, cols: Int) -> Array<N, Float> {
    // Create and return a 2D array with size determined by N
}

// Using dependent type
let matrix3x3 = create_matrix<9>(3, 3)
let matrix4x4 = create_matrix<16>(4, 4)

// The type system guarantees the size matches
fn process_3x3_matrix(m: Array<9, Float>) {
    // Can safely access all 9 elements
}

process_3x3_matrix(matrix3x3)  // OK
// process_3x3_matrix(matrix4x4)  // Error: Type mismatch
```

## Type-level Programming

Unified provides facilities for type-level programming, enabling you to perform computations at the type level:

```unified
// Type-level integers
type Zero {}
type Succ<N> {}

// Type-level operations
type Add<A, B> = if A == Zero { B } else { Succ<Add<Pred<A>, B>> }
type Pred<Succ<N>> = N

// Type-level conditionals
type If<Cond, Then, Else> = Cond match {
    True => Then,
    False => Else
}

// Type-level equality
type Equals<A, B> = // Implementation

// Using type-level programming
fn vector_add<N, A: Add, B: Add>(a: Vector<N, A>, b: Vector<N, B>) -> Vector<N, Add<A, B>> {
    // Implementation
}
```

Type-level programming enables advanced generic programming techniques that ensure correctness by encoding invariants in the type system.

# 14. Pattern Matching and Destructuring

Pattern matching is a powerful feature in Unified that allows you to check a value against patterns and extract components. When combined with destructuring, it enables concise, expressive code.

## Basic Patterns

The `switch` statement is the primary way to use pattern matching in Unified:

```unified
let value = 42

switch value {
    case 0 -> print("Zero")
    case 1 -> print("One")
    case 2 -> print("Two")
    case _ -> print("Something else")
}
```

Patterns can include ranges and conditions:

```unified
switch value {
    case 0 -> print("Zero")
    case 1..5 -> print("Small")
    case 6..10 -> print("Medium")
    case let x if x > 10 && x < 100 -> print("Large")
    case _ -> print("Huge")
}
```

## Destructuring

Destructuring allows you to extract components from compound values:

```unified
// Destructuring tuples
let point = (10, 20)
let (x, y) = point
print("X: ${x}, Y: ${y}")  // "X: 10, Y: 20"

// Destructuring in function parameters
fn process_point((x, y): (Int, Int)) {
    print("Processing point at ${x}, ${y}")
}

// Destructuring structs
struct Person {
    name: String
    age: Int
}

let person = Person{name: "Alice", age: 30}
let Person{name, age} = person
print("${name} is ${age} years old")  // "Alice is 30 years old"

// Partial destructuring
let Person{name, ..} = person  // Extract only name
```

## Guards

Pattern guards add additional conditions to patterns:

```unified
switch value {
    case let n if n % 2 == 0 -> print("Even: ${n}")
    case let n if n % 2 == 1 -> print("Odd: ${n}")
}

// Guards with destructuring
switch user {
    case User{name, age} if age >= 18 -> allowAccess(name)
    case User{name, ..} -> denyAccess(name)
}
```

## Exhaustiveness

The compiler checks that pattern matching is exhaustive, ensuring all possible cases are handled:

```unified
enum Color {
    Red
    Green
    Blue
}

// Missing cases will cause a compile error
fn describe(color: Color) -> String {
    switch color {
        case Red -> "The color of roses"
        case Green -> "The color of grass"
        // Error: Non-exhaustive pattern matching, missing case Blue
    }
}
```

This helps prevent bugs by forcing you to consider all possibilities.

## Pattern Matching with Enums

Pattern matching is particularly powerful with enums:

```unified
enum Shape {
    Circle(radius: Float)
    Rectangle(width: Float, height: Float)
    Triangle(base: Float, height: Float)
}

let shape = Shape.Circle(radius: 5.0)

switch shape {
    case Circle(let radius) -> {
        let area = 3.14159 * radius * radius
        print("Circle with area ${area}")
    }
    case Rectangle(let width, let height) -> {
        let area = width * height
        print("Rectangle with area ${area}")
    }
    case Triangle(let base, let height) -> {
        let area = 0.5 * base * height
        print("Triangle with area ${area}")
    }
}
```

The pattern extracts the associated values from the enum variant.

## Nested Pattern Matching

Patterns can be nested to match complex structures:

```unified
enum Tree<T> {
    Leaf(value: T)
    Node(left: Tree<T>, value: T, right: Tree<T>)
}

let tree = Tree.Node(
    left: Tree.Leaf(value: 1),
    value: 2,
    right: Tree.Node(
        left: Tree.Leaf(value: 3),
        value: 4,
        right: Tree.Leaf(value: 5)
    )
)

switch tree {
    case Leaf(let value) -> 
        print("Leaf with value ${value}")
    
    case Node(Leaf(let left_val), let value, Leaf(let right_val)) -> 
        print("Simple node: ${left_val} <- ${value} -> ${right_val}")
    
    case Node(let left, let value, let right) -> 
        print("Complex node with value ${value}")
}
```

## Pattern Matching with Options and Results

Pattern matching works seamlessly with `Option` and `Result` types:

```unified
let maybe_value: Option<Int> = Some(42)

switch maybe_value {
    case Some(let value) -> print("Got value: ${value}")
    case None -> print("No value")
}

let result: Result<String, Error> = Success("Data")

switch result {
    case Success(let data) -> process(data)
    case Error(let error) -> handle_error(error)
}
```

## If-let and While-let

For simple cases, `if let` and `while let` provide more concise syntax:

```unified
// If-let
if let Some(value) = get_optional() {
    print("Got value: ${value}")
}

// Equivalent to:
switch get_optional() {
    case Some(let value) -> print("Got value: ${value}")
    case None -> {}
}

// While-let
while let Some(item) = iterator.next() {
    process(item)
}

// Equivalent to:
loop {
    switch iterator.next() {
        case Some(let item) -> process(item)
        case None -> break
    }
}
```

## Custom Patterns

You can define custom patterns for your types by implementing the `Pattern` interface:

```unified
interface Pattern<T, P> {
    fn matches(value: &T) -> Bool;
    fn extract(value: &T) -> Option<P>;
}

struct EmailPattern {}

impl Pattern<String, (String, String)> for EmailPattern {
    fn matches(value: &String) -> Bool {
        return value.contains("@")
    }
    
    fn extract(value: &String) -> Option<(String, String)> {
        let parts = value.split("@")
        if parts.len() == 2 {
            return Some((parts[0], parts[1]))
        }
        return None
    }
}

// Using custom pattern
switch email {
    case EmailPattern(let username, let domain) -> {
        print("Username: ${username}")
        print("Domain: ${domain}")
    }
    case _ -> print("Invalid email")
}
```

Pattern matching and destructuring make Unified code more concise, readable, and less error-prone by eliminating boilerplate and ensuring that all cases are properly handled.
# 20. Best Practices and Idioms

As you become more familiar with Unified, adopting established best practices and idioms will help you write more maintainable, efficient, and idiomatic code. This section covers recommended approaches for various aspects of the language.

## Coding Style

Unified has developed an opinionated but practical style guide based on experience with the language.

### Naming Conventions

```unified
// Types use PascalCase
struct User {}
enum ConnectionState {}
interface Serializable {}

// Functions, methods, and variables use camelCase
fn connectToDatabase() {}
let userData = fetchUserData()

// Constants use UPPER_SNAKE_CASE
const MAX_CONNECTIONS = 100
const PI = 3.14159

// Module names use snake_case
module network_utils {
    // Module contents
}

// Generic type parameters use single uppercase letters or PascalCase
struct List<T> {}
struct KeyValuePair<K, V> {}
struct HttpResponse<ResponseData> {}
```

### Formatting

- Indentation: 4 spaces (not tabs)
- Line length: aim for 80-100 characters maximum
- Opening braces on the same line for control structures and function declarations
- Space after keywords and before braces
- Spaces around binary operators
- No space after unary operators

```unified
// Proper formatting
fn calculate(values: List<Int>) -> Int {
    let mut sum = 0
    
    for value in values {
        if value > 0 {
            sum += value
        } else if value < -10 {
            sum -= value
        }
    }
    
    return sum * 2
}
```

### Documentation Comments

Document public APIs with full sentences and examples:

```unified
/// Returns the sum of all positive numbers in the list.
///
/// Negative numbers are ignored during the calculation.
///
/// # Examples
///
/// ```
/// let values = [1, -2, 3, -4, 5]
/// let sum = sumPositives(values)  // Returns 9 (1 + 3 + 5)
/// ```
fn sumPositives(values: List<Int>) -> Int {
    return values.filter(|n| n > 0).reduce(0, |acc, n| acc + n)
}
```

## API Design

When designing APIs in Unified, favor these principles:

### Be Explicit About Ownership

Make it clear when ownership is transferred:

```unified
// Good: explicit about taking ownership
fn processAndStore(data: String) {
    // Takes ownership of data
}

// Good: explicit about borrowing
fn analyze(data: &String) -> Stats {
    // Only borrows data
}

// When using functions:
processAndStore(move myData)  // Explicitly transferring ownership
analyze(&myData)              // Explicitly borrowing
```

### Return Results for Fallible Operations

Functions that can fail should return `Result` or `Option`:

```unified
// Good
fn divide(a: Int, b: Int) -> Result<Float, String> {
    if b == 0 {
        return Error("Division by zero")
    }
    return Success(a as Float / b as Float)
}

// Also good for values that might not exist
fn findUser(id: UserId) -> Option<User> {
    // Implementation
}
```

### Favor Composition Over Inheritance

Build complex types and behaviors through composition:

```unified
// Good: compose functionality
struct Database {
    connection: Connection
    cache: Cache
    logger: Logger
    
    fn query(self, query: String) -> Result<QueryResult, DbError> {
        self.logger.log("Executing query: ${query}")
        
        if let Some(cached) = self.cache.get(query) {
            return Success(cached)
        }
        
        let result = self.connection.execute(query)?
        self.cache.store(query, result.clone())
        return Success(result)
    }
}
```

### Use Method Chaining Where Appropriate

For fluent interfaces, return `self` from methods that modify state:

```unified
struct QueryBuilder {
    table: String
    conditions: List<String>
    ordering: Option<String>
    
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
    
    fn build(self) -> String {
        // Build and return SQL query string
    }
}

// Usage
let query = QueryBuilder.new()
    .from("users")
    .where("age > 18")
    .orderBy("last_name")
    .build()
```

## Error Handling Strategies

### Use the Type System to Prevent Errors

Where possible, make illegal states unrepresentable:

```unified
// Instead of a boolean flag:
fn processPayment(amount: Float, isRefund: Bool) // Not ideal

// Better: Use an enum to make the intent explicit
enum PaymentType {
    Purchase
    Refund
}

fn processPayment(amount: Float, type: PaymentType) // Better!
```

### Propagate Errors with `?` Operator

Use the `?` operator for clean error handling:

```unified
fn processUserData() -> Result<UserProfile, Error> {
    let userId = getUserId()?
    let userData = fetchUserData(userId)?
    let profile = createProfile(userData)?
    return Success(profile)
}
```

### Use Pattern Matching for Comprehensive Error Handling

When you need to handle errors differently based on their type:

```unified
let result = fetchData()

switch result {
    case Success(let data) -> processData(data)
    case Error(NetworkError.Timeout) -> retryWithBackoff()
    case Error(NetworkError.Authentication) -> promptLogin()
    case Error(let err) -> reportError(err)
}
```

### Create Custom Error Types

For rich domain-specific error information:

```unified
enum OrderError {
    ProductNotFound(productId: String)
    InsufficientInventory(productId: String, requested: Int, available: Int)
    PaymentDeclined(reason: String)
}

fn placeOrder(order: Order) -> Result<OrderConfirmation, OrderError> {
    // Implementation with specific errors
}
```

## Performance Considerations

### Avoid Premature Optimization

Write clear, correct code first, then optimize if necessary:

```unified
// Start with the simple, readable version
fn processData(data: List<Int>) -> Int {
    return data.filter(|n| n > 0).map(|n| n * n).sum()
}

// Only if profiling shows this is a bottleneck, consider:
fn processDataOptimized(data: List<Int>) -> Int {
    let mut sum = 0
    for n in data {
        if n > 0 {
            sum += n * n
        }
    }
    return sum
}
```

### Use Region-Based Memory for Performance-Critical Code

```unified
fn processLargeDataset(dataset: &Dataset) -> Results {
    region temp {
        let intermediate1 = createLargeIntermediate(dataset)
        let intermediate2 = furtherProcess(intermediate1)
        return finalizeResults(intermediate2)
    }  // All intermediates freed at once
}
```

### Consider Memory Layout for Data-Heavy Applications

Group related data for better cache locality:

```unified
// Better cache locality for particle simulations
struct Particles {
    // Structure of Arrays (SoA) for better cache locality
    positions_x: Array<Float>
    positions_y: Array<Float>
    positions_z: Array<Float>
    velocities_x: Array<Float>
    velocities_y: Array<Float>
    velocities_z: Array<Float>
}

// Better than Array of Structures (AoS) for parallel processing
// struct Particle { x, y, z, vx, vy, vz }
// struct Particles { particles: Array<Particle> }
```

### Use Lazy Evaluation for On-Demand Computation

```unified
// Create a lazy sequence that only computes values when needed
let infiniteSequence = LazySequence.counting(from: 1)
    .map(|n| n * n)
    .filter(|n| n % 2 == 0)

// Only computes the first 5 values
let firstFive = infiniteSequence.take(5).collect()
```

## Migrating from Other Languages

### For Java/C# Developers

- Embrace immutability as the default
- Use composition over inheritance
- Think in terms of values rather than objects
- Use interfaces instead of abstract classes
- Leverage the ownership system instead of garbage collection

```unified
// Rather than Java's:
class UserService {
    private UserRepository repository;
    
    public User findUser(String id) {
        return repository.findById(id);
    }
}

// In Unified:
struct UserService {
    repository: UserRepository
    
    fn findUser(self, id: String) -> Option<User> {
        return self.repository.findById(id)
    }
}
```

### For JavaScript/TypeScript Developers

- Embrace static typing
- Use pattern matching instead of type checks
- Use actors instead of callbacks for asynchronous code
- Use Result/Option instead of undefined/null checks

```unified
// Rather than TypeScript:
function processValue(value: string | number | null) {
    if (value === null) {
        return "No value";
    } else if (typeof value === "string") {
        return value.toUpperCase();
    } else {
        return value * 2;
    }
}

// In Unified:
fn processValue(value: Option<String | Int>) -> String | Int {
    switch value {
        case None -> "No value"
        case Some(let s: String) -> s.toUpperCase()
        case Some(let n: Int) -> n * 2
    }
}
```

### For C/C++ Developers

- Rely on the ownership system instead of manual memory management
- Use Result/Option instead of error codes or exceptions
- Use algebraic data types instead of unions
- Use interfaces instead of virtual inheritance

```unified
// Rather than C++:
class Shape {
public:
    virtual double area() const = 0;
    virtual ~Shape() {}
};

class Circle : public Shape {
private:
    double radius;
public:
    Circle(double r) : radius(r) {}
    double area() const override { return 3.14159 * radius * radius; }
};

// In Unified:
interface Shape {
    fn area(self) -> Float;
}

struct Circle {
    radius: Float
}

impl Shape for Circle {
    fn area(self) -> Float {
        return 3.14159 * self.radius * self.radius
    }
}
```

### For Rust Developers

- Use explicit `move` for ownership transfers
- Leverage region-based memory where appropriate
- Use actors for concurrent code
- Use async/await for asynchronous operations

```unified
// Rather than Rust:
fn process_data(data: Vec<i32>) -> Result<i32, Error> {
    let sum: i32 = data.iter()
        .filter(|&x| *x > 0)
        .map(|&x| x * 2)
        .sum();
    
    Ok(sum)
}

// In Unified:
fn processData(data: List<Int>) -> Result<Int, Error> {
    let sum = data.filter(|x| x > 0)
        .map(|x| x * 2)
        .sum()
    
    return Success(sum)
}
```

By following these best practices and idioms, you'll write Unified code that is not only correct and efficient but also idiomatic and maintainable. Remember that clarity should be your primary goal—code is read far more often than it is written.
