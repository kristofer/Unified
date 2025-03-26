# Learn Unified in Y Minutes

Unified is a modern systems programming language that combines memory safety, expressive typing, and ergonomic concurrency. It draws inspiration from Rust, Go, Swift, and Elixir while maintaining its own unique approach.

```unified
// Single-line comments use double slashes

/* 
   Multi-line comments use
   slash-asterisk pairs
*/

// ----- VARIABLES AND BASIC TYPES -----

// Type inference with 'let'
let message = "Hello, world!"  // String
let count = 42                 // Int
let pi = 3.14159               // Float
let active = true              // Bool

// Explicit typing
let id: Int64 = 1000000000     // 64-bit integer
let name: String = "Unified"
let values: List<Int> = [1, 2, 3, 4, 5]

// Constants
const MAX_USERS = 100          // Immutable by default

// Immutability is default, mutability is explicit
let mut counter = 0           // Mutable variable
counter += 1                  // Modification allowed

// String interpolation
let greeting = "Welcome, ${name}!"

// ----- CONTROL FLOW -----

// If expressions (no parentheses needed)
if count > 40 {
    print("Over 40")
} else if count > 20 {
    print("Over 20")
} else {
    print("20 or less")
}

// If as an expression
let status = if active { "running" } else { "stopped" }

// Pattern matching with switch
switch someValue {
    case 0 -> print("Zero")
    case 1..5 -> print("Small")
    case let x if x > 100 -> print("Large: ${x}")
    case _ -> print("Something else")
}

// For loops
for i in 0..10 {              // Range 0 to 9
    print(i)
}

for item in values {
    print(item)
}

// While loops
while counter < 10 {
    counter += 1
}

// ----- FUNCTIONS -----

// Basic function
fn add(a: Int, b: Int) -> Int {
    return a + b
}

// Functions are expressions; return keyword optional for single expression
fn multiply(a: Int, b: Int) -> Int {
    a * b  // Implicit return
}

// Default parameters
fn greet(name: String, greeting: String = "Hello") -> String {
    "${greeting}, ${name}!"
}

// Multiple return values using tuples
fn divideWithRemainder(a: Int, b: Int) -> (Int, Int) {
    (a / b, a % b)
}

// First-class functions
let operation = add
let result = operation(5, 3)  // result = 8

// Higher-order functions
fn apply(value: Int, transform: fn(Int) -> Int) -> Int {
    transform(value)
}

// ----- DATA STRUCTURES -----

// Tuples
let point = (10, 20)
let (x, y) = point        // Destructuring
print(point.0)            // Access by index

// Lists (arrays)
let numbers = [1, 2, 3, 4]
print(numbers[2])         // Access by index
let first = numbers.first // Safe access returns Option<Int>

// Maps (dictionaries)
let scores = {
    "Alice": 42,
    "Bob": 37,
    "Charlie": 91
}
let aliceScore = scores["Alice"]  // Returns Option<Int>

// ----- CUSTOM TYPES -----

// Struct definition
struct Point {
    x: Float
    y: Float
    
    // Methods
    fn distance(self, other: Point) -> Float {
        let dx = self.x - other.x
        let dy = self.y - other.y
        return (dx * dx + dy * dy).sqrt()
    }
    
    // Static method
    fn origin() -> Point {
        Point{x: 0.0, y: 0.0}
    }
}

// Creating a struct instance
let p1 = Point{x: 10.0, y: 20.0}
let p2 = Point{x: 30.0, y: 40.0}
let dist = p1.distance(p2)

// Algebraic data types (enums)
enum Result<T, E> {
    Success(T)
    Error(E)
}

// Using enum variants
let success: Result<Int, String> = Result.Success(42)
let error: Result<Int, String> = Result.Error("Invalid input")

// ----- PATTERN MATCHING -----

// Matching on enums
fn handleResult(result: Result<Int, String>) {
    match result {
        Success(value) -> print("Success: ${value}")
        Error(msg) -> print("Error: ${msg}")
    }
}

// ----- OWNERSHIP AND BORROWING -----

// Ownership
fn process(mut data: String) {
    data.append(" processed")
    // data is dropped when function exits
}

// Borrowing (non-mutable)
fn analyze(data: &String) -> Int {
    data.length  // Can read but not modify
}

// Mutable borrowing
fn modify(data: &mut String) {
    data.append(" modified")  // Can modify
}

// Usage
let mut text = "Hello"
analyze(&text)     // Immutable borrow
modify(&mut text)  // Mutable borrow
// text moved here: process(text)
// Cannot use text after moving

// ----- ERROR HANDLING -----

// Using Result type
fn divideIfPossible(a: Int, b: Int) -> Result<Int, String> {
    if b == 0 {
        return Error("Division by zero")
    }
    return Success(a / b)
}

// Using the ? operator for error propagation
fn calculate(a: Int, b: Int) -> Result<Int, String> {
    let div = divideIfPossible(a, b)?  // Returns early on Error
    return Success(div * 2)
}

// Pattern matching with Result
let calculation = calculate(10, 2)
match calculation {
    Success(value) -> print("Result: ${value}")
    Error(msg) -> print("Failed: ${msg}")
}

// ----- CONCURRENCY -----

// Actors define isolated state
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
let newCount = counter.increment()  // Asynchronous message send

// Async/await
async fn fetchData(url: String) -> Result<String, Error> {
    // Asynchronous implementation
}

async fn processUrl(url: String) -> Result<Int, Error> {
    let data = await fetchData(url)?
    return Success(data.length)
}

// ----- MODULES AND IMPORTS -----

// Importing from the standard library
import std.collections.List
import std.io.{File, ReadError}

// Creating a module
module math {
    pub fn square(x: Int) -> Int {
        x * x
    }
    
    pub fn cube(x: Int) -> Int {
        x * x * x
    }
    
    // Private function
    fn helper() {
        // Implementation
    }
}

// Using module functions
let squared = math.square(5)
```

