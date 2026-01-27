# Unified Standard Library

This directory contains the standard library implementation for the Unified Programming Language. The standard library provides generic data structures and utilities that work with any type.

## Generic Data Structures

### List<T>
A dynamic array implementation that automatically grows as needed.

**Methods:**
- `new() -> List<T>` - Create a new empty list with default capacity
- `withCapacity(cap: Int) -> List<T>` - Create a list with specified capacity
- `len(self) -> Int` - Get the number of elements
- `isEmpty(self) -> Bool` - Check if the list is empty
- `push(mut self, item: T)` - Add an element to the end
- `pop(mut self) -> T` - Remove and return the last element
- `get(self, index: Int) -> T` - Get element at index
- `set(mut self, index: Int, item: T)` - Set element at index
- `clear(mut self)` - Remove all elements
- `contains(self, item: T) -> Bool` - Check if element exists
- `indexOf(self, item: T) -> Int` - Find index of element (-1 if not found)

**Example:**
```unified
let list = List<Int>.new()
list.push(10)
list.push(20)
list.push(30)
let val = list.pop()  // Returns 30
```

### Stack<T>
A Last-In-First-Out (LIFO) data structure.

**Methods:**
- `new() -> Stack<T>` - Create a new empty stack
- `withCapacity(cap: Int) -> Stack<T>` - Create with specified capacity
- `len(self) -> Int` - Get the number of elements
- `isEmpty(self) -> Bool` - Check if the stack is empty
- `push(mut self, item: T)` - Push an element onto the stack
- `pop(mut self) -> T` - Pop and return the top element
- `peek(self) -> T` - View the top element without removing it
- `clear(mut self)` - Remove all elements
- `contains(self, item: T) -> Bool` - Check if element exists
- `getCapacity(self) -> Int` - Get current capacity

**Example:**
```unified
let stack = Stack<String>.new()
stack.push("first")
stack.push("second")
stack.push("third")
let top = stack.pop()  // Returns "third" (LIFO)
```

### Queue<T>
A First-In-First-Out (FIFO) data structure.

**Methods:**
- `new() -> Queue<T>` - Create a new empty queue
- `withCapacity(cap: Int) -> Queue<T>` - Create with specified capacity
- `len(self) -> Int` - Get the number of elements
- `isEmpty(self) -> Bool` - Check if the queue is empty
- `enqueue(mut self, item: T)` - Add an element to the back
- `dequeue(mut self) -> T` - Remove and return the front element
- `peek(self) -> T` - View the front element without removing it
- `clear(mut self)` - Remove all elements
- `contains(self, item: T) -> Bool` - Check if element exists
- `getCapacity(self) -> Int` - Get current capacity

**Example:**
```unified
let queue = Queue<Int>.new()
queue.enqueue(1)
queue.enqueue(2)
queue.enqueue(3)
let first = queue.dequeue()  // Returns 1 (FIFO)
```

### HashMap<K, V>
A hash table that maps keys to values using hashing for fast lookups.

**Methods:**
- `new() -> HashMap<K, V>` - Create a new empty hash map
- `withCapacity(cap: Int) -> HashMap<K, V>` - Create with specified capacity
- `len(self) -> Int` - Get the number of key-value pairs
- `isEmpty(self) -> Bool` - Check if the map is empty
- `insert(mut self, key: K, value: V)` - Insert or update a key-value pair
- `get(self, key: K) -> V` - Get value by key (panics if not found)
- `containsKey(self, key: K) -> Bool` - Check if key exists
- `remove(mut self, key: K)` - Remove a key-value pair
- `clear(mut self)` - Remove all entries

**Example:**
```unified
let map = HashMap<String, Int>.new()
map.insert("age", 25)
map.insert("score", 100)
let age = map.get("age")  // Returns 25
```

### Set<T>
A collection of unique elements with fast membership testing.

**Methods:**
- `new() -> Set<T>` - Create a new empty set
- `withCapacity(cap: Int) -> Set<T>` - Create with specified capacity
- `len(self) -> Int` - Get the number of elements
- `isEmpty(self) -> Bool` - Check if the set is empty
- `insert(mut self, item: T) -> Bool` - Add element (returns false if already exists)
- `contains(self, item: T) -> Bool` - Check if element exists
- `remove(mut self, item: T) -> Bool` - Remove element (returns true if removed)
- `clear(mut self)` - Remove all elements
- `toArray(self) -> Array<T>` - Get all elements as array
- `union(self, other: Set<T>) -> Set<T>` - Union with another set
- `intersection(self, other: Set<T>) -> Set<T>` - Intersection with another set

**Example:**
```unified
let set = Set<Int>.new()
set.insert(1)
set.insert(2)
set.insert(1)  // Returns false, already exists
let hasTwo = set.contains(2)  // Returns true
```

### BinaryTree<T>
A binary search tree that maintains elements in sorted order.

**Methods:**
- `new() -> BinaryTree<T>` - Create a new empty binary tree
- `withCapacity(cap: Int) -> BinaryTree<T>` - Create with specified capacity
- `len(self) -> Int` - Get the number of nodes
- `isEmpty(self) -> Bool` - Check if the tree is empty
- `insert(mut self, value: T)` - Insert a value (maintains BST property)
- `contains(self, value: T) -> Bool` - Check if value exists
- `findMin(self) -> T` - Find minimum value
- `findMax(self) -> T` - Find maximum value
- `height(self) -> Int` - Get the height of the tree
- `clear(mut self)` - Remove all nodes

**Example:**
```unified
let tree = BinaryTree<Int>.new()
tree.insert(50)
tree.insert(30)
tree.insert(70)
let min = tree.findMin()  // Returns 30
let max = tree.findMax()  // Returns 70
```

## Testing

All standard library components are thoroughly tested with multiple data types:
- **Int** - Integer values
- **Float** - Floating point numbers
- **String** - Text strings
- **Struct** - Custom data structures

To run the standard library tests:
```bash
make run-lib-tests
```

This will compile and run all 18 test files:
- List: `list_int.uni`, `list_float.uni`, `list_string.uni`, `list_struct.uni`
- Stack: `stack_int.uni`, `stack_float.uni`, `stack_string.uni`, `stack_struct.uni`
- Queue: `queue_int.uni`, `queue_float.uni`, `queue_string.uni`, `queue_struct.uni`
- HashMap: `hashmap_int.uni`, `hashmap_string.uni`
- Set: `set_int.uni`, `set_string.uni`
- BinaryTree: `binarytree_int.uni`, `binarytree_string.uni`

## Implementation Details

All data structures are implemented using:
- **Generics**: Type parameters allow working with any type
- **Array-based storage**: Efficient memory layout
- **Automatic growth**: Capacity doubles when needed
- **Hash-based indexing** (HashMap, Set): Linear probing for collision resolution
- **Node-based tree structure** (BinaryTree): Array-based node storage for efficiency
- **Circular buffer** (Queue only): Efficient FIFO operations

## Algorithms

- **HashMap & Set**: Open addressing with linear probing, 75% load factor before resize
- **BinaryTree**: Standard BST insertion with recursive traversal
- **All structures**: Automatic capacity doubling when growth threshold reached

## Future Additions

Planned additions to the standard library:
- PriorityQueue<T> - Priority-based queue with heap implementation
- LinkedList<T> - Doubly-linked list
- Iterator<T> - Iteration support for all collections
- Graph<T> - Graph data structure with adjacency list
- Trie<T> - Prefix tree for string operations
- AVLTree<T> - Self-balancing binary search tree
