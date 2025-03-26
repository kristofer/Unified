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

