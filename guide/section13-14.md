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
