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

