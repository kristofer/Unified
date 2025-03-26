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
