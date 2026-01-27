# Standard Library Extended - HashMap, Set, BinaryTree

## Overview

Extended the Unified standard library with three additional generic data structures: HashMap, Set, and BinaryTree. These additions complement the existing List, Stack, and Queue implementations.

## New Data Structures

### 1. HashMap<K, V> - Hash Table

**File**: [stdlib/HashMap.uni](stdlib/HashMap.uni)  
**Lines of Code**: 180  
**Algorithm**: Open addressing with linear probing

**Features**:
- Generic key-value storage
- Dynamic resizing at 75% load factor
- Linear probing for collision resolution
- Type-safe operations for both keys and values

**Methods** (11 total):
- `new()`, `withCapacity(cap)` - Constructors
- `len()`, `isEmpty()` - Query operations
- `insert(key, value)` - Add/update entry
- `get(key)` - Retrieve value by key
- `containsKey(key)` - Check for key existence
- `remove(key)` - Delete entry
- `clear()` - Remove all entries
- `hash(key)`, `findSlot(key)`, `resize()` - Internal operations

**Complexity**:
- Insert: O(1) average, O(n) worst case
- Get: O(1) average, O(n) worst case
- Remove: O(1) average, O(n) worst case

### 2. Set<T> - Unique Element Collection

**File**: [stdlib/Set.uni](stdlib/Set.uni)  
**Lines of Code**: 195  
**Algorithm**: Hash-based set with linear probing

**Features**:
- Stores unique elements only
- Automatic duplicate prevention
- Set operations (union, intersection)
- Dynamic resizing

**Methods** (13 total):
- `new()`, `withCapacity(cap)` - Constructors
- `len()`, `isEmpty()` - Query operations
- `insert(item)` - Add element (returns bool)
- `contains(item)` - Membership test
- `remove(item)` - Delete element (returns bool)
- `clear()` - Remove all elements
- `toArray()` - Convert to array for iteration
- `union(other)` - Combine two sets
- `intersection(other)` - Common elements
- `hash(item)`, `findSlot(item)`, `resize()` - Internal operations

**Complexity**:
- Insert: O(1) average
- Contains: O(1) average
- Remove: O(1) average
- Union: O(m + n)
- Intersection: O(min(m, n))

### 3. BinaryTree<T> - Binary Search Tree

**File**: [stdlib/BinaryTree.uni](stdlib/BinaryTree.uni)  
**Lines of Code**: 262  
**Algorithm**: Standard BST with array-based node storage

**Features**:
- Maintains sorted order
- Efficient search operations
- Min/max finding
- Height calculation
- Array-based node storage (no pointers)

**Methods** (15 total):
- `new()`, `withCapacity(cap)` - Constructors
- `len()`, `isEmpty()` - Query operations
- `insert(value)` - Add element (maintains BST property)
- `contains(value)` - Search for element
- `findMin()`, `findMax()` - Find extremes
- `height()` - Tree height calculation
- `clear()` - Remove all nodes
- Helper methods: `insertHelper()`, `containsHelper()`, `findMinHelper()`, `findMaxHelper()`, `heightHelper()`, `allocateNode()`, `resize()`

**Complexity**:
- Insert: O(log n) average, O(n) worst case
- Search: O(log n) average, O(n) worst case
- FindMin/Max: O(log n) average, O(n) worst case
- Height: O(n)

## Test Suite

**Total New Tests**: 6 test files

### HashMap Tests (2 files)
- [hashmap_int.uni](test/stdlib/hashmap_int.uni) - Int keys/values
- [hashmap_string.uni](test/stdlib/hashmap_string.uni) - String keys/values

**Coverage**: Insert, get, update, containsKey, remove, clear

### Set Tests (2 files)
- [set_int.uni](test/stdlib/set_int.uni) - Integer elements
- [set_string.uni](test/stdlib/set_string.uni) - String elements

**Coverage**: Insert (with duplicate detection), contains, remove, clear

### BinaryTree Tests (2 files)
- [binarytree_int.uni](test/stdlib/binarytree_int.uni) - Integer values
- [binarytree_string.uni](test/stdlib/binarytree_string.uni) - String values (alphabetical)

**Coverage**: Insert, contains, findMin, findMax, height, clear

## Implementation Highlights

### HashMap & Set Design
Both use similar hash-based storage:
- **Linear Probing**: Simple and cache-friendly collision resolution
- **Load Factor**: Resize at 75% capacity to maintain performance
- **Array-based**: Efficient memory layout with two parallel arrays (Set) or structured entries (HashMap)

### BinaryTree Design
- **Array-based Nodes**: Uses indices instead of pointers for memory efficiency
- **Free List**: Tracks available node slots for reuse
- **Recursive Operations**: Clean, idiomatic tree traversals
- **Not Self-Balancing**: Standard BST (AVL/Red-Black trees planned for future)

### Common Patterns
All three structures follow the stdlib conventions:
- Default capacity of 16 elements
- Doubling growth strategy
- Panic on invalid operations
- Clear separation of public API and internal helpers

## Files Created

### Standard Library (3 files)
1. `/src/unified-compiler/stdlib/HashMap.uni` (180 lines)
2. `/src/unified-compiler/stdlib/Set.uni` (195 lines)
3. `/src/unified-compiler/stdlib/BinaryTree.uni` (262 lines)

### Tests (6 files)
4-5. HashMap tests: `hashmap_{int,string}.uni`  
6-7. Set tests: `set_{int,string}.uni`  
8-9. BinaryTree tests: `binarytree_{int,string}.uni`

### Documentation (2 files updated)
- [stdlib/README.md](stdlib/README.md) - Added API docs for all three structures
- [Makefile](Makefile) - Updated `stdlib-info` target

## Complete Standard Library Summary

**Total Data Structures**: 6
- List<T>
- Stack<T>
- Queue<T>
- HashMap<K, V>
- Set<T>
- BinaryTree<T>

**Total Test Files**: 18
- 4 List tests (Int, Float, String, Struct)
- 4 Stack tests (Int, Float, String, Struct)
- 4 Queue tests (Int, Float, String, Struct)
- 2 HashMap tests (Int, String)
- 2 Set tests (Int, String)
- 2 BinaryTree tests (Int, String)

**Total Lines of Code**: ~1,400+ new lines (637 stdlib + tests)

## Testing

All tests can be run with:
```bash
make run-lib-tests
```

Expected output: 18 tests, all passing

## Performance Characteristics

| Structure | Insert | Get/Contains | Remove | Space |
|-----------|--------|--------------|--------|-------|
| List<T> | O(1)* | O(1) | O(1)* | O(n) |
| Stack<T> | O(1) | O(1) | O(1) | O(n) |
| Queue<T> | O(1) | O(1) | O(1) | O(n) |
| HashMap<K,V> | O(1)† | O(1)† | O(1)† | O(n) |
| Set<T> | O(1)† | O(1)† | O(1)† | O(n) |
| BinaryTree<T> | O(log n)‡ | O(log n)‡ | - | O(n) |

\* Amortized constant time (due to resizing)  
† Average case (worst case O(n) with many collisions)  
‡ Average case (worst case O(n) with unbalanced tree)

## Future Enhancements

Potential improvements:
- **HashMap**: Better hash functions, separate chaining option
- **Set**: Implement difference and symmetric difference operations
- **BinaryTree**: Add removal operation, implement AVL or Red-Black balancing
- **All**: Iterator support for foreach loops
- **New**: PriorityQueue (heap-based), Graph, Trie

## Success Criteria

✅ HashMap<K, V> implementation working  
✅ Set<T> implementation working  
✅ BinaryTree<T> implementation working  
✅ Tests for Int type  
✅ Tests for String type  
✅ Documentation updated  
✅ Type-safe generic operations  
✅ Efficient algorithms with known complexity  

## Notes

These implementations demonstrate:
- Advanced use of generics (including two type parameters for HashMap)
- Array-based implementations avoiding pointer complexity
- Classic computer science data structures in Unified
- Consistent API design across the standard library

The complete standard library now provides a comprehensive foundation for building complex applications in Unified.
