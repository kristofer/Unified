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
