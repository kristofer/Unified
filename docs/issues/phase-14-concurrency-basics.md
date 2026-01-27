# Phase 14: Concurrency Basics - Async/Await and Actors

**Status**: Not Started  
**Duration**: 4-6 weeks  
**Priority**: HIGH  
**Dependencies**: Phase 13 Complete (Standard Library Foundation)

## 🎯 Goals

Implement foundational concurrency features including async/await for cooperative multitasking and a basic actor model for safe concurrent programming. This phase establishes Unified's approach to fearless concurrency.

## 📋 Prerequisites

- [x] Phase 13 complete (stdlib foundation available)
- [ ] Understanding of async/await patterns
- [ ] Understanding of actor model concurrency
- [ ] Knowledge of event loops and task scheduling
- [ ] Familiarity with message passing systems
- [ ] Understanding of concurrent data structures

## ✨ Language Features to Add

### 1. Async/Await Syntax
- `async fn` keyword for async functions
- `await` keyword for awaiting futures
- Future trait and implementation
- Task spawning
- Async blocks

### 2. Runtime Support
- Event loop / executor
- Task scheduler
- Future polling mechanism
- Wake mechanism for pending futures
- Runtime initialization

### 3. Actor Model
- Actor definitions (`actor` keyword)
- Message passing
- Actor spawning (`spawn` keyword)
- Mailbox/message queue
- Actor lifecycle (creation, execution, termination)

### 4. Channels
- Message channels for actor communication
- Sender and Receiver types
- Bounded and unbounded channels
- Channel operations (send, receive)

### 5. Synchronization Primitives
- Mutex for shared state
- Channel-based communication
- Atomic operations (basic)

## 📝 Implementation Tasks

### Task 14.1: Add Async/Await Grammar (3-4 hours)
- [ ] Open `grammar/UnifiedParser.g4`
- [ ] Add `async` keyword to function declarations
- [ ] Add `await` keyword as expression
- [ ] Add actor declaration grammar
- [ ] Add spawn expression grammar
- [ ] Regenerate parser: `make parser`
- [ ] Test parsing async syntax

**Grammar to add:**
```antlr
asyncFunctionDeclaration
    : ASYNC FN IDENTIFIER typeParameters? parameterList (ARROW type)? block
    ;

awaitExpression
    : AWAIT expression
    ;

actorDeclaration
    : ACTOR IDENTIFIER typeParameters? LBRACE actorMember* RBRACE
    ;

actorMember
    : fieldDeclaration
    | functionDeclaration
    ;

spawnExpression
    : SPAWN expression
    ;
```

### Task 14.2: Implement Async AST Nodes (2-3 hours)
- [ ] Open `internal/ast/ast.go`
- [ ] Add AsyncFunctionDeclaration struct
- [ ] Add AwaitExpression struct
- [ ] Add ActorDeclaration struct
- [ ] Add SpawnExpression struct
- [ ] Update visitor pattern
- [ ] Write unit tests

**AST Structures:**
```go
type AsyncFunctionDeclaration struct {
    Name       string
    Params     []Parameter
    ReturnType Type
    Body       *Block
}

type AwaitExpression struct {
    Expression Expression
}

type ActorDeclaration struct {
    Name    string
    Fields  []FieldDeclaration
    Methods []FunctionDeclaration
}

type SpawnExpression struct {
    ActorConstructor Expression
}
```

### Task 14.3: Implement Future Trait and Runtime (8-10 hours)
- [ ] Create `stdlib/async/future.uni`
- [ ] Define Future trait
- [ ] Implement basic future types
- [ ] Create runtime executor
- [ ] Implement task spawning
- [ ] Implement waker mechanism
- [ ] Write comprehensive tests

**Future Trait:**
```unified
pub trait Future {
    type Output
    fn poll(&mut self, waker: &Waker) -> Poll<Self::Output>
}

pub enum Poll<T> {
    Ready(T),
    Pending,
}
```

**Runtime:**
```go
type Runtime struct {
    tasks      []*Task
    readyQueue chan *Task
    running    bool
}

type Task struct {
    future Future
    waker  Waker
    state  TaskState
}

func (r *Runtime) Spawn(future Future) TaskHandle
func (r *Runtime) Run()
func (r *Runtime) Block(future Future) result
```

### Task 14.4: Implement Async Code Generation (6-8 hours)
- [ ] Update `internal/bytecode/generator.go`
- [ ] Transform async functions to state machines
- [ ] Generate await points
- [ ] Generate future implementations
- [ ] Handle async function calls
- [ ] Write comprehensive tests

**State Machine Transformation:**
```
async fn foo() -> Int {
    let x = await bar()
    let y = await baz()
    return x + y
}

// Transforms to state machine:
enum FooState {
    Start,
    AwaitBar,
    AwaitBaz,
    Done,
}
```

### Task 14.5: Implement VM Async Support (6-8 hours)
- [ ] Update `internal/vm/vm.go`
- [ ] Add async execution support
- [ ] Integrate runtime executor
- [ ] Handle future polling
- [ ] Implement cooperative scheduling
- [ ] Support multiple concurrent tasks
- [ ] Write tests

### Task 14.6: Implement Actor System (10-12 hours)
- [ ] Create `stdlib/actor/actor.uni`
- [ ] Implement actor runtime
- [ ] Create actor mailbox/message queue
- [ ] Implement message dispatch
- [ ] Handle actor lifecycle
- [ ] Implement spawn functionality
- [ ] Support actor isolation
- [ ] Write comprehensive tests

**Actor System:**
```unified
pub struct ActorHandle<T> {
    // Internal actor reference
}

pub trait Actor {
    type Message
    fn receive(&mut self, msg: Self::Message)
}

pub fn spawn<A: Actor>(actor: A) -> ActorHandle<A>
```

**Implementation:**
```go
type ActorRuntime struct {
    actors map[ActorID]*ActorInstance
    router *MessageRouter
}

type ActorInstance struct {
    id       ActorID
    mailbox  *Mailbox
    state    interface{}
    running  bool
}

type Mailbox struct {
    queue    chan Message
    capacity int
}

func (r *ActorRuntime) Spawn(actorType ActorType) ActorHandle
func (r *ActorRuntime) Send(handle ActorHandle, msg Message)
func (a *ActorInstance) ProcessMessages()
```

### Task 14.7: Implement Channels (6-8 hours)
- [ ] Create `stdlib/async/channel.uni`
- [ ] Implement channel creation
- [ ] Implement Sender type
- [ ] Implement Receiver type
- [ ] Support bounded and unbounded channels
- [ ] Add async send/receive
- [ ] Handle channel closure
- [ ] Write comprehensive tests

**Channel API:**
```unified
pub fn channel<T>() -> (Sender<T>, Receiver<T>)
pub fn bounded_channel<T>(capacity: Int) -> (Sender<T>, Receiver<T>)

pub struct Sender<T> {
    // Internal
}

impl<T> Sender<T> {
    pub async fn send(&self, value: T) -> Result<(), SendError>
    pub fn try_send(&self, value: T) -> Result<(), TrySendError>
}

pub struct Receiver<T> {
    // Internal
}

impl<T> Receiver<T> {
    pub async fn recv(&self) -> Option<T>
    pub fn try_recv(&self) -> Result<T, TryRecvError>
}
```

### Task 14.8: Implement Basic Mutex (4-5 hours)
- [ ] Create `stdlib/sync/mutex.uni`
- [ ] Implement Mutex<T>
- [ ] Add lock/unlock operations
- [ ] Implement MutexGuard for RAII
- [ ] Integrate with ownership system
- [ ] Write tests

**Mutex API:**
```unified
pub struct Mutex<T> {
    // Internal
}

impl<T> Mutex<T> {
    pub fn new(value: T) -> Mutex<T>
    pub fn lock(&self) -> MutexGuard<T>
    pub fn try_lock(&self) -> Option<MutexGuard<T>>
}

pub struct MutexGuard<'a, T> {
    // Automatically unlocks on drop
}
```

### Task 14.9: Add Concurrency to Compiler (5-6 hours)
- [ ] Update type checker for async functions
- [ ] Validate await usage
- [ ] Check actor message types
- [ ] Ensure Send trait for concurrent types
- [ ] Generate async runtime initialization code
- [ ] Update VM to start runtime

### Task 14.10: Write Comprehensive Tests (10-12 hours)
- [ ] Parser tests for async/actor syntax (10+ tests)
- [ ] Future implementation tests (8+ tests)
- [ ] Runtime executor tests (10+ tests)
- [ ] Actor system tests (12+ tests)
- [ ] Channel tests (10+ tests)
- [ ] Mutex tests (6+ tests)
- [ ] Integration tests (12+ tests)
- [ ] Concurrent stress tests
- [ ] Deadlock detection tests

### Task 14.11: Update Documentation (5-6 hours)
- [ ] Create PHASE14_SUMMARY.md
- [ ] Create ASYNC_GUIDE.md
- [ ] Create ACTOR_GUIDE.md
- [ ] Document concurrency model
- [ ] Add async/await examples
- [ ] Add actor examples
- [ ] Update README.md

## ✅ Test Requirements

**Minimum Test Coverage**: 90% for concurrency code

### Parser Tests (Async/Actor Syntax)
- [ ] Parse async function declaration
- [ ] Parse await expression
- [ ] Parse actor declaration
- [ ] Parse spawn expression
- [ ] Parse async block
- [ ] Parse actor with fields and methods
- [ ] Parse complex async/await chains
- [ ] Parse nested actors
- [ ] Parse channel operations
- [ ] Parse mutex operations

### Future/Runtime Tests
- [ ] Create simple future
- [ ] Poll future to completion
- [ ] Test waker mechanism
- [ ] Spawn task
- [ ] Execute multiple tasks
- [ ] Test task scheduling
- [ ] Test cooperative yielding
- [ ] Test future chaining
- [ ] Test error propagation in futures
- [ ] Test runtime shutdown

### Actor System Tests
- [ ] Create actor instance
- [ ] Spawn actor
- [ ] Send message to actor
- [ ] Receive message in actor
- [ ] Test message ordering
- [ ] Test actor isolation
- [ ] Test multiple actors
- [ ] Test actor-to-actor communication
- [ ] Test actor termination
- [ ] Test mailbox overflow handling
- [ ] Test actor state management
- [ ] Test concurrent actor execution

### Channel Tests
- [ ] Create unbounded channel
- [ ] Create bounded channel
- [ ] Send and receive messages
- [ ] Test async send/receive
- [ ] Test try_send/try_recv
- [ ] Test channel closure
- [ ] Test multiple senders
- [ ] Test multiple receivers
- [ ] Test channel backpressure
- [ ] Test channel iteration

### Mutex Tests
- [ ] Lock and unlock mutex
- [ ] Test mutual exclusion
- [ ] Test try_lock
- [ ] Test MutexGuard drop
- [ ] Test concurrent access
- [ ] Test deadlock prevention

### Integration Tests (End-to-End)
- [ ] Compile async function
- [ ] Execute async function
- [ ] Compile and run actor program
- [ ] Test multi-actor communication
- [ ] Test channel-based communication
- [ ] Test async with channels
- [ ] Test complex concurrent program
- [ ] Test error handling in async code
- [ ] Test actor supervision (basic)
- [ ] Test concurrent data processing
- [ ] Performance test: many actors
- [ ] Stress test: high message volume

## 📚 Documentation Updates Required

### Update Existing Docs
- [ ] README.md: Add concurrency features
- [ ] TESTING.md: Document concurrency testing

### Create New Docs
- [ ] PHASE14_SUMMARY.md: Complete phase summary
- [ ] docs/ASYNC_GUIDE.md: Async/await comprehensive guide
- [ ] docs/ACTOR_GUIDE.md: Actor model guide
- [ ] docs/CONCURRENCY_MODEL.md: Overall concurrency design
- [ ] docs/CHANNELS_GUIDE.md: Channel usage patterns
- [ ] docs/SYNC_PRIMITIVES.md: Synchronization guide

### Add Example Programs
- [ ] `examples/async/simple_async.uni`
- [ ] `examples/async/async_file_io.uni`
- [ ] `examples/async/parallel_tasks.uni`
- [ ] `examples/actors/counter_actor.uni`
- [ ] `examples/actors/ping_pong.uni`
- [ ] `examples/actors/chat_system.uni`
- [ ] `examples/channels/producer_consumer.uni`
- [ ] `examples/concurrency/pipeline.uni`
- [ ] `examples/concurrency/web_crawler.uni`

## 🎓 Example Programs

### Simple Async Function
```unified
async fn fetch_data() -> String {
    let data = await read_file("data.txt")
    return data
}

async fn main() -> Int {
    let result = await fetch_data()
    println(result)
    return 0
}
```

### Parallel Async Tasks
```unified
async fn process_item(item: Int) -> Int {
    // Simulate async work
    await sleep(100)
    return item * 2
}

async fn main() -> Int {
    let tasks = new List()
    
    for i in 1..=10 {
        tasks.push(spawn(process_item(i)))
    }
    
    let mut total = 0
    for task in tasks {
        total += await task
    }
    
    println(total)
    return 0
}
```

### Counter Actor
```unified
actor Counter {
    var count: Int = 0
    
    fn increment(&mut self) {
        self.count += 1
    }
    
    fn get(&self) -> Int {
        return self.count
    }
}

fn main() -> Int {
    let counter = spawn new Counter()
    
    counter.send(Counter::Increment)
    counter.send(Counter::Increment)
    
    let value = await counter.send(Counter::Get)
    println(value)  // Prints 2
    
    return 0
}
```

### Ping Pong Actors
```unified
actor Ping {
    fn ping(&self, pong: ActorHandle<Pong>) {
        println("Ping!")
        pong.send(Pong::Pong(self.handle()))
    }
}

actor Pong {
    fn pong(&self, ping: ActorHandle<Ping>) {
        println("Pong!")
        ping.send(Ping::Ping(self.handle()))
    }
}

fn main() -> Int {
    let ping = spawn new Ping()
    let pong = spawn new Pong()
    
    ping.send(Ping::Ping(pong))
    
    // Let them play
    await sleep(1000)
    
    return 0
}
```

### Producer-Consumer with Channels
```unified
async fn producer(tx: Sender<Int>) {
    for i in 1..=10 {
        await tx.send(i)
        println("Produced: " + i.to_string())
    }
}

async fn consumer(rx: Receiver<Int>) {
    loop {
        match await rx.recv() {
            Some(value) => {
                println("Consumed: " + value.to_string())
            },
            None => break,
        }
    }
}

async fn main() -> Int {
    let (tx, rx) = channel()
    
    spawn(producer(tx))
    spawn(consumer(rx))
    
    // Wait for completion
    await sleep(1000)
    
    return 0
}
```

### Parallel Data Processing Pipeline
```unified
async fn stage1(input: Receiver<Int>, output: Sender<Int>) {
    loop {
        match await input.recv() {
            Some(value) => await output.send(value * 2),
            None => break,
        }
    }
}

async fn stage2(input: Receiver<Int>, output: Sender<Int>) {
    loop {
        match await input.recv() {
            Some(value) => await output.send(value + 1),
            None => break,
        }
    }
}

async fn main() -> Int {
    let (tx1, rx1) = channel()
    let (tx2, rx2) = channel()
    let (tx3, rx3) = channel()
    
    spawn(stage1(rx1, tx2))
    spawn(stage2(rx2, tx3))
    
    // Feed data into pipeline
    for i in 1..=5 {
        await tx1.send(i)
    }
    
    // Collect results
    for _ in 1..=5 {
        let result = await rx3.recv()
        println(result)
    }
    
    return 0
}
```

### Actor-Based Chat System
```unified
actor ChatRoom {
    var clients: List<ActorHandle<Client>> = new List()
    
    fn join(&mut self, client: ActorHandle<Client>) {
        self.clients.push(client)
    }
    
    fn broadcast(&self, message: String, sender: ActorHandle<Client>) {
        for client in self.clients {
            if client != sender {
                client.send(Client::Message(message))
            }
        }
    }
}

actor Client {
    var name: String
    var room: ActorHandle<ChatRoom>
    
    fn send_message(&self, message: String) {
        self.room.send(ChatRoom::Broadcast(message, self.handle()))
    }
    
    fn receive_message(&self, message: String) {
        println(self.name + " received: " + message)
    }
}

fn main() -> Int {
    let room = spawn new ChatRoom()
    
    let alice = spawn new Client("Alice", room)
    let bob = spawn new Client("Bob", room)
    
    room.send(ChatRoom::Join(alice))
    room.send(ChatRoom::Join(bob))
    
    alice.send(Client::SendMessage("Hello!"))
    bob.send(Client::SendMessage("Hi there!"))
    
    return 0
}
```

## 🏆 Success Criteria

- [ ] All parser tests pass (minimum 10 tests)
- [ ] All future/runtime tests pass (minimum 10 tests)
- [ ] All actor system tests pass (minimum 12 tests)
- [ ] All channel tests pass (minimum 10 tests)
- [ ] All mutex tests pass (minimum 6 tests)
- [ ] All integration tests pass (minimum 12 tests)
- [ ] Can compile and run async functions
- [ ] Can spawn and communicate with actors
- [ ] Channels work correctly
- [ ] No data races in concurrent code
- [ ] No deadlocks in test programs
- [ ] No regressions in previous phases
- [ ] Code coverage ≥ 90%
- [ ] Documentation is comprehensive
- [ ] Example programs demonstrate features

## 💡 Implementation Notes

### Implementation Order
1. Grammar and parser (syntax support)
2. AST nodes (representation)
3. Future trait and runtime (async foundation)
4. Async code generation (state machines)
5. VM async support (execution)
6. Actor system (concurrency model)
7. Channels (communication)
8. Mutex (synchronization)
9. Compiler integration
10. Tests and documentation

### Testing Strategy
- Test async and actors separately first
- Use deterministic scheduling in tests where possible
- Test concurrent edge cases thoroughly
- Use stress tests to find race conditions
- Test both success and error paths
- Verify no memory leaks in async code
- Test cancellation and cleanup

### Common Pitfalls
1. **Race Conditions**: Test concurrent access thoroughly
2. **Deadlocks**: Design careful communication patterns
3. **Memory Leaks**: Ensure futures are properly dropped
4. **Actor Isolation**: Don't share mutable state
5. **Channel Buffering**: Consider backpressure
6. **Task Starvation**: Ensure fair scheduling
7. **Error Handling**: Propagate errors correctly in async code

### Debugging Tips
- Add extensive logging to runtime scheduler
- Visualize task execution timeline
- Use deterministic seeds for testing
- Add deadlock detection in debug mode
- Monitor mailbox sizes
- Track actor lifecycle events
- Use async stack traces

### Performance Considerations
- Minimize allocations in hot paths
- Use efficient task scheduling
- Consider work-stealing for multi-core
- Batch message processing where possible
- Use lock-free data structures for channels
- Profile async overhead
- Optimize state machine generation

### Simplifications (For This Phase)
- Single-threaded async runtime (multi-threaded later)
- Basic actor supervision (full OTP-style supervision later)
- Simple message routing (optimizations later)
- Cooperative scheduling only (no preemption)
- Basic error handling (improve later)

### Future Extensions (Not This Phase)
- Multi-threaded async executor
- Work-stealing scheduler
- Actor supervision trees
- Hot code reloading for actors
- Distributed actors
- Async I/O integration
- select! macro for channels
- Async iterators and streams
- Backpressure mechanisms
- Actor persistence

---

**Labels**: `phase-14`, `enhancement`, `async`, `actors`, `concurrency`, `high-priority`  
**Milestone**: Phase 14 - Concurrency Basics  
**Assignee**: TBD
