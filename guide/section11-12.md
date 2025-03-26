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
