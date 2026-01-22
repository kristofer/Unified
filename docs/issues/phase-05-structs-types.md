# Phase 5: Structs and Basic Types

**Status**: Not Started  
**Duration**: 3-4 weeks  
**Priority**: MEDIUM  
**Dependencies**: Phase 4 Complete (Functions and Advanced Expressions)

## 🎯 Goals

Implement user-defined struct types with fields and methods. This is the foundation for object-oriented features and provides the basic building blocks for more complex data structures.

## 📋 Prerequisites

- [x] Phase 4 complete (functions and expressions working)
- [ ] Understanding of struct memory layout
- [ ] Familiarity with method dispatch concepts

## ✨ Language Features to Add

### 1. Struct Declarations
- Field definitions with types
- Struct instantiation with field initialization
- Field access using dot notation
- Field visibility (public/private preparation)

### 2. Methods
- Methods on structs
- `self` parameter (immutable self)
- `&mut self` parameter (mutable self)
- Method call syntax
- Method chaining

### 3. Struct Methods
- Associated functions (static methods)
- Instance methods
- Constructor patterns
- Builder pattern support

### 4. Memory Layout
- Struct field ordering
- Padding and alignment
- Size calculations

## 📝 Implementation Tasks

### Task 5.1: Struct Grammar (3-4 hours)
- [ ] Add struct declaration to grammar
- [ ] Add struct body with fields and methods
- [ ] Add field declaration syntax
- [ ] Add method declaration syntax
- [ ] Add struct instantiation syntax
- [ ] Add field access (dot notation)
- [ ] Regenerate parser
- [ ] Test parsing struct declarations

**Grammar to add:**
```antlr
structDecl
    : STRUCT IDENTIFIER genericParams? LBRACE structBody RBRACE
    ;

structBody
    : (fieldDecl | methodDecl)*
    ;

fieldDecl
    : (PUB)? IDENTIFIER COLON type SEMICOLON?
    ;

methodDecl
    : (PUB)? FN IDENTIFIER LPAREN selfParam? (COMMA paramList)? RPAREN (ARROW type)? block
    ;

selfParam
    : SELF
    | AMPERSAND MUT SELF
    ;

structLiteral
    : IDENTIFIER LBRACE fieldInit (COMMA fieldInit)* RBRACE
    ;

fieldInit
    : IDENTIFIER COLON expression
    ;

fieldAccess
    : primaryExpr DOT IDENTIFIER
    ;
```

### Task 5.2: Struct Type System (4-6 hours)
- [ ] Create StructType in type system
- [ ] Track field names and types in order
- [ ] Implement field lookup by name
- [ ] Calculate field offsets
- [ ] Add struct to type registry
- [ ] Implement type checking for field access
- [ ] Add memory layout calculations
- [ ] Handle field visibility rules
- [ ] Add unit tests for type system

**Type system structure:**
```go
type StructType struct {
    Name    string
    Fields  []StructField
    Methods map[string]*MethodInfo
}

type StructField struct {
    Name    string
    Type    Type
    Offset  int
    Public  bool
}

type MethodInfo struct {
    Name       string
    Receiver   ReceiverType
    Parameters []Parameter
    ReturnType Type
}
```

### Task 5.3: Struct Instantiation (4-5 hours)
- [ ] Implement struct literal parsing
- [ ] Validate all fields are initialized
- [ ] Validate field types match
- [ ] Generate bytecode for struct creation
- [ ] Allocate struct in VM memory/heap
- [ ] Initialize fields in correct order
- [ ] Add tests for struct creation
- [ ] Test error cases (missing fields, wrong types)

**Example bytecode:**
```
ALLOC_STRUCT <size>
STORE_FIELD <offset> <value>
...
```

### Task 5.4: Field Access (3-4 hours)
- [ ] Implement field access in AST
- [ ] Validate field exists on struct
- [ ] Calculate field offset at compile time
- [ ] Generate bytecode for field load
- [ ] Generate bytecode for field store
- [ ] Implement in VM (field load/store instructions)
- [ ] Add tests for field access
- [ ] Test nested field access

### Task 5.5: Method Declarations (4-5 hours)
- [ ] Parse method declarations in struct body
- [ ] Add methods to struct type
- [ ] Handle self parameter specially
- [ ] Distinguish &self vs &mut self
- [ ] Validate self type matches struct
- [ ] Add methods to symbol table
- [ ] Generate bytecode for method definitions
- [ ] Add unit tests for method parsing

### Task 5.6: Method Calls (6-8 hours)
- [ ] Implement method call syntax
- [ ] Resolve method from struct type
- [ ] Pass self as implicit first parameter
- [ ] Generate bytecode for method calls
- [ ] Implement method dispatch in VM
- [ ] Support method chaining
- [ ] Add comprehensive method call tests
- [ ] Test self vs &mut self behavior

### Task 5.7: Associated Functions (3-4 hours)
- [ ] Support functions without self parameter
- [ ] Implement struct::function() call syntax
- [ ] Common pattern: `new()` constructor
- [ ] Generate appropriate bytecode
- [ ] Test associated functions
- [ ] Document constructor patterns

### Task 5.8: Write Tests (5-6 hours)
- [ ] Parser tests for struct syntax
- [ ] Type system tests for structs
- [ ] Struct instantiation tests
- [ ] Field access tests
- [ ] Method call tests
- [ ] Associated function tests
- [ ] Integration tests with complex structs
- [ ] Error case tests
- [ ] Ensure code coverage ≥ 85%

### Task 5.9: Update Documentation (3-4 hours)
- [ ] Create PHASE5_SUMMARY.md
- [ ] Document structs in language guide
- [ ] Document methods and self
- [ ] Add struct examples
- [ ] Update README.md
- [ ] Create struct best practices guide

## ✅ Test Requirements

**Minimum Test Coverage**: 85% for new code

### Parser Tests
- [ ] Parse struct declaration with fields
- [ ] Parse struct with methods
- [ ] Parse struct literal
- [ ] Parse field access
- [ ] Parse method call
- [ ] Parse associated function call
- [ ] Parse nested structs
- [ ] Reject invalid syntax

### Type System Tests
- [ ] Create struct types
- [ ] Look up fields by name
- [ ] Calculate field offsets
- [ ] Validate field types
- [ ] Register methods on structs
- [ ] Type check field access
- [ ] Type check method calls

### Instantiation Tests
- [ ] Create struct with all fields
- [ ] Error on missing fields
- [ ] Error on extra fields
- [ ] Error on wrong field types
- [ ] Create nested structs
- [ ] Initialize with expressions

### Field Access Tests
- [ ] Access struct fields
- [ ] Assign to mutable struct fields
- [ ] Error on accessing non-existent field
- [ ] Error on assigning to immutable field
- [ ] Chain field accesses
- [ ] Access nested struct fields

### Method Tests
- [ ] Declare methods on structs
- [ ] Call methods with self
- [ ] Call methods with &mut self
- [ ] Error on calling mut method on immutable
- [ ] Method chaining works
- [ ] Associated functions work

### Integration Tests
- [ ] Point struct with distance method
- [ ] Rectangle struct with area/perimeter
- [ ] Builder pattern example
- [ ] Nested struct example
- [ ] Complex method interactions

## 📚 Documentation Updates Required

### Update Existing Docs
- [ ] README.md: Add structs to feature list
- [ ] TESTING.md: Add struct test categories
- [ ] VM_README.md: Document struct-related bytecode

### Create New Docs
- [ ] PHASE5_SUMMARY.md: Complete phase summary
- [ ] examples/STRUCTS.md: Guide to structs and methods
- [ ] docs/MEMORY_LAYOUT.md: Struct memory layout documentation
- [ ] examples/OOP_PATTERNS.md: Object-oriented patterns

### Add Example Programs
- [ ] `test/point.uni`: Point struct with methods
- [ ] `test/rectangle.uni`: Rectangle with area/perimeter
- [ ] `test/builder.uni`: Builder pattern example
- [ ] `test/nested_structs.uni`: Nested struct example

## 🎓 Example Programs

### Point Struct
```unified
struct Point {
    x: Float
    y: Float
    
    fn new(x: Float, y: Float) -> Point {
        return Point { x: x, y: y }
    }
    
    fn distance_squared(self) -> Float {
        return self.x * self.x + self.y * self.y
    }
    
    fn move_by(&mut self, dx: Float, dy: Float) {
        self.x += dx
        self.y += dy
    }
    
    fn scale(&mut self, factor: Float) {
        self.x *= factor
        self.y *= factor
    }
}

fn main() -> Int {
    let mut p = Point::new(3.0, 4.0)
    let dist_sq = p.distance_squared()  // 25.0
    p.move_by(1.0, 1.0)                  // p is now (4.0, 5.0)
    p.scale(2.0)                         // p is now (8.0, 10.0)
    return 0
}
```

### Rectangle Struct
```unified
struct Rectangle {
    width: Float
    height: Float
    
    fn new(width: Float, height: Float) -> Rectangle {
        return Rectangle { width: width, height: height }
    }
    
    fn area(self) -> Float {
        return self.width * self.height
    }
    
    fn perimeter(self) -> Float {
        return 2.0 * (self.width + self.height)
    }
    
    fn is_square(self) -> Bool {
        return self.width == self.height
    }
    
    fn resize(&mut self, new_width: Float, new_height: Float) {
        self.width = new_width
        self.height = new_height
    }
}

fn main() -> Int {
    let mut rect = Rectangle::new(5.0, 10.0)
    let area = rect.area()           // 50.0
    let perimeter = rect.perimeter() // 30.0
    let is_sq = rect.is_square()     // false
    
    rect.resize(7.0, 7.0)
    let is_sq_now = rect.is_square() // true
    
    return 0
}
```

### Person Struct (Nested Structs)
```unified
struct Address {
    street: String
    city: String
    zipcode: Int
}

struct Person {
    name: String
    age: Int
    address: Address
    
    fn new(name: String, age: Int) -> Person {
        return Person {
            name: name,
            age: age,
            address: Address {
                street: "",
                city: "",
                zipcode: 0
            }
        }
    }
    
    fn set_address(&mut self, street: String, city: String, zip: Int) {
        self.address.street = street
        self.address.city = city
        self.address.zipcode = zip
    }
    
    fn get_city(self) -> String {
        return self.address.city
    }
}

fn main() -> Int {
    let mut person = Person::new("Alice", 30)
    person.set_address("123 Main St", "Boston", 12345)
    let city = person.get_city()  // "Boston"
    return 0
}
```

### Counter Struct (Mutable Self)
```unified
struct Counter {
    value: Int
    
    fn new() -> Counter {
        return Counter { value: 0 }
    }
    
    fn increment(&mut self) {
        self.value += 1
    }
    
    fn decrement(&mut self) {
        self.value -= 1
    }
    
    fn get_value(self) -> Int {
        return self.value
    }
    
    fn reset(&mut self) {
        self.value = 0
    }
}

fn main() -> Int {
    let mut counter = Counter::new()
    counter.increment()
    counter.increment()
    counter.increment()
    // counter.value is now 3
    let val = counter.get_value()
    return val  // Returns 3
}
```

### Method Chaining Example
```unified
struct Builder {
    value: Int
    
    fn new() -> Builder {
        return Builder { value: 0 }
    }
    
    fn add(&mut self, n: Int) -> &mut Builder {
        self.value += n
        return self
    }
    
    fn multiply(&mut self, n: Int) -> &mut Builder {
        self.value *= n
        return self
    }
    
    fn build(self) -> Int {
        return self.value
    }
}

fn main() -> Int {
    let mut builder = Builder::new()
    // Method chaining
    let result = builder.add(5).multiply(2).add(10).build()
    // result = (5 * 2) + 10 = 20
    return result
}
```

## 🏆 Success Criteria

- [ ] All parser tests pass (minimum 15 tests)
- [ ] All type system tests pass (minimum 12 tests)
- [ ] All instantiation tests pass (minimum 10 tests)
- [ ] All field access tests pass (minimum 12 tests)
- [ ] All method tests pass (minimum 15 tests)
- [ ] All integration tests pass (minimum 8 tests)
- [ ] Struct types work end-to-end
- [ ] Methods can access and modify fields
- [ ] Type checking works for structs
- [ ] Associated functions work correctly
- [ ] Method chaining is possible
- [ ] No regressions in previous phase tests
- [ ] Code coverage ≥ 85%
- [ ] Documentation complete
- [ ] All example programs run correctly

## 💡 Implementation Notes

### Implementation Order
1. Struct grammar and parsing
2. Type system for structs
3. Struct instantiation
4. Field access
5. Method declarations
6. Method calls
7. Associated functions
8. Integration and testing

### Testing Strategy
- Test struct features incrementally
- Start with simple structs, progress to complex
- Test both immutable and mutable methods
- Verify type checking at each step
- Test error cases thoroughly
- Use integration tests for realistic scenarios

### Common Pitfalls
1. Forgetting to validate all fields are initialized
2. Field offset calculation errors
3. Self parameter type confusion (&self vs &mut self)
4. Method name resolution conflicts
5. Memory layout and alignment issues
6. Not handling nested structs correctly
7. Method chaining return type issues

### Debugging Tips
- Print struct type information during compilation
- Add debug output for field offset calculations
- Verify bytecode for struct operations
- Test memory layout manually
- Use simple examples before complex ones
- Check symbol table for method registration

### Memory Management
- Structs allocated on heap (or stack for now)
- Field offsets calculated at compile time
- Consider alignment and padding
- Plan for future garbage collection integration

### Future Considerations
This phase prepares for:
- Traits and interfaces (Phase 11)
- Generic structs (Phase 10)
- Ownership and borrowing on struct fields
- Advanced patterns (destructuring, etc.)
- Operator overloading on structs

---

**Labels**: `phase-5`, `enhancement`, `types`, `structs`, `oop`  
**Milestone**: Phase 5 - Structs and Basic Types  
**Assignee**: TBD
