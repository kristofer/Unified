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
