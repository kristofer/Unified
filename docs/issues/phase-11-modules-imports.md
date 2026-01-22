# Phase 11: Modules and Imports

**Status**: Not Started  
**Duration**: 2-3 weeks  
**Priority**: HIGH  
**Dependencies**: Phase 10 Complete (Functions, Structs, and Methods)

## 🎯 Goals

Implement a comprehensive module system with imports, exports, and visibility control to enable code organization and reusability across multiple files.

## 📋 Prerequisites

- [x] Phase 10 complete (structs and methods working)
- [ ] Understanding of module resolution strategies
- [ ] Understanding of visibility and access control
- [ ] Knowledge of package management concepts

## ✨ Language Features to Add

### 1. Module System
- Module declarations (`mod` keyword)
- File-based modules (one module per file)
- Directory-based modules (mod.uni files)
- Module paths and namespacing

### 2. Import/Export
- Import statements (`use` keyword)
- Import paths (absolute and relative)
- Selective imports (`use module::{item1, item2}`)
- Wildcard imports (`use module::*`)
- Aliasing (`use module::Item as Alias`)
- Re-exporting (`pub use`)

### 3. Visibility Control
- Public visibility (`pub` keyword)
- Private-by-default semantics
- Module-level visibility
- Struct field visibility
- Function visibility

### 4. Module Resolution
- Path resolution algorithm
- Search paths and module roots
- Circular dependency detection
- Module caching

## 📝 Implementation Tasks

### Task 11.1: Add Module Grammar (3-4 hours)
- [ ] Open `grammar/UnifiedParser.g4`
- [ ] Add module declaration grammar
- [ ] Add import/use statement grammar
- [ ] Add visibility modifier (pub) to declarations
- [ ] Update compilation unit to support modules
- [ ] Regenerate parser: `make parser`
- [ ] Test parsing module declarations

**Grammar to add:**
```antlr
moduleDeclaration
    : MOD IDENTIFIER SEMICOLON
    ;

useStatement
    : USE usePath (AS IDENTIFIER)? SEMICOLON
    ;

usePath
    : IDENTIFIER (DOUBLECOLON usePathSegment)*
    ;

usePathSegment
    : IDENTIFIER
    | LBRACE usePathList RBRACE
    | STAR
    ;

usePathList
    : usePath (COMMA usePath)*
    ;

visibilityModifier
    : PUB
    ;

// Update function, struct, etc. to include optional visibility
functionDeclaration
    : visibilityModifier? FN IDENTIFIER ...
    ;
```

### Task 11.2: Implement Module AST Nodes (2-3 hours)
- [ ] Open `internal/ast/ast.go`
- [ ] Add ModuleDeclaration struct
- [ ] Add UseStatement struct
- [ ] Add Visibility enum/type
- [ ] Update existing declaration nodes to include visibility
- [ ] Update visitor pattern to handle modules
- [ ] Write unit tests for AST nodes

**AST Structures:**
```go
type ModuleDeclaration struct {
    Name string
}

type UseStatement struct {
    Path       []string
    Items      []string  // for selective imports
    Wildcard   bool
    Alias      string
}

type Visibility int
const (
    VisibilityPrivate Visibility = iota
    VisibilityPublic
)

type FunctionDeclaration struct {
    Visibility Visibility
    Name       string
    // ... existing fields
}
```

### Task 11.3: Implement Module Resolver (6-8 hours)
- [ ] Create `internal/modules/resolver.go`
- [ ] Implement module path resolution
- [ ] Implement file system module discovery
- [ ] Handle relative and absolute imports
- [ ] Build module dependency graph
- [ ] Detect circular dependencies
- [ ] Implement module caching
- [ ] Write tests for resolution logic

**Key Components:**
```go
type ModuleResolver struct {
    rootPath    string
    moduleCache map[string]*Module
}

func (r *ModuleResolver) Resolve(path string) (*Module, error)
func (r *ModuleResolver) LoadModule(filepath string) (*Module, error)
func (r *ModuleResolver) DetectCircularDeps() error
```

### Task 11.4: Implement Symbol Table with Scoping (5-6 hours)
- [ ] Create `internal/symbols/table.go`
- [ ] Implement hierarchical symbol tables
- [ ] Track module-level symbols
- [ ] Track imported symbols
- [ ] Implement name resolution across modules
- [ ] Handle visibility checking
- [ ] Write comprehensive tests

**Symbol Table:**
```go
type SymbolTable struct {
    parent     *SymbolTable
    symbols    map[string]*Symbol
    module     *Module
}

type Symbol struct {
    Name       string
    Type       Type
    Visibility Visibility
    Module     string
}

func (s *SymbolTable) Lookup(name string) (*Symbol, error)
func (s *SymbolTable) LookupInModule(module, name string) (*Symbol, error)
```

### Task 11.5: Update Bytecode Generator for Modules (4-5 hours)
- [ ] Update `internal/bytecode/generator.go`
- [ ] Generate module metadata in bytecode
- [ ] Handle cross-module function calls
- [ ] Implement qualified name resolution
- [ ] Generate import tables
- [ ] Update tests for multi-module programs

### Task 11.6: Update VM for Module Support (3-4 hours)
- [ ] Update `internal/vm/vm.go`
- [ ] Load multiple bytecode modules
- [ ] Implement module linking at runtime
- [ ] Handle cross-module calls and data access
- [ ] Maintain separate module namespaces
- [ ] Test multi-module execution

### Task 11.7: Add Multi-File Compilation (4-5 hours)
- [ ] Update `cmd/compiler/main.go`
- [ ] Accept multiple input files
- [ ] Discover modules in directories
- [ ] Orchestrate multi-file compilation
- [ ] Generate combined bytecode output
- [ ] Add command-line flags for module paths
- [ ] Test multi-file projects

### Task 11.8: Write Comprehensive Tests (6-8 hours)
- [ ] Parser tests for module syntax (10+ tests)
- [ ] Module resolver tests (8+ tests)
- [ ] Symbol table tests (10+ tests)
- [ ] Bytecode generation tests (6+ tests)
- [ ] VM module loading tests (5+ tests)
- [ ] Integration tests (6+ tests)
- [ ] Test circular dependency detection
- [ ] Test visibility enforcement

### Task 11.9: Update Documentation (3-4 hours)
- [ ] Create PHASE11_SUMMARY.md
- [ ] Update README.md with module system
- [ ] Document module resolution algorithm
- [ ] Create module system guide
- [ ] Add multi-file examples
- [ ] Document visibility rules

## ✅ Test Requirements

**Minimum Test Coverage**: 90% for new code

### Parser Tests (Module Syntax)
- [ ] Parse module declaration
- [ ] Parse simple use statement
- [ ] Parse qualified use path
- [ ] Parse selective imports
- [ ] Parse wildcard imports
- [ ] Parse aliased imports
- [ ] Parse pub visibility modifier
- [ ] Parse pub use re-export
- [ ] Parse nested module paths
- [ ] Parse multiple imports

### Module Resolution Tests
- [ ] Resolve simple module import
- [ ] Resolve relative path import
- [ ] Resolve absolute path import
- [ ] Resolve module from file
- [ ] Resolve module from directory
- [ ] Detect circular dependencies
- [ ] Cache loaded modules
- [ ] Handle missing modules

### Symbol Table Tests
- [ ] Lookup local symbols
- [ ] Lookup imported symbols
- [ ] Lookup qualified symbols
- [ ] Respect visibility modifiers
- [ ] Block access to private symbols
- [ ] Allow access to public symbols
- [ ] Handle symbol shadowing
- [ ] Handle module namespaces
- [ ] Handle re-exports
- [ ] Handle aliased imports

### Integration Tests (End-to-End)
- [ ] Compile two-file program
- [ ] Compile multi-module project
- [ ] Import and use functions from other module
- [ ] Import and use structs from other module
- [ ] Enforce visibility across modules
- [ ] Detect and report circular dependencies

## 📚 Documentation Updates Required

### Update Existing Docs
- [ ] README.md: Add module system to feature list
- [ ] TESTING.md: Document module testing approach
- [ ] VM_README.md: Document module bytecode format

### Create New Docs
- [ ] PHASE11_SUMMARY.md: Complete phase summary
- [ ] docs/MODULE_SYSTEM.md: Comprehensive module guide
- [ ] docs/VISIBILITY.md: Visibility and access control guide

### Add Example Programs
- [ ] `test/phase11/simple_module.uni`: Basic module example
- [ ] `test/phase11/math/mod.uni`: Math module
- [ ] `test/phase11/main.uni`: Uses math module
- [ ] `test/phase11/visibility.uni`: Visibility examples
- [ ] `test/phase11/circular/`: Circular dependency test
- [ ] `test/phase11/reexport.uni`: Re-export example

## 🎓 Example Programs

### Basic Module (math.uni)
```unified
// File: math.uni
mod math

pub fn add(a: Int, b: Int) -> Int {
    return a + b
}

pub fn multiply(a: Int, b: Int) -> Int {
    return a * b
}

fn internal_helper() -> Int {
    return 42  // Private function
}
```

### Using a Module (main.uni)
```unified
// File: main.uni
use math::{add, multiply}

fn main() -> Int {
    let sum = add(10, 20)
    let product = multiply(5, 6)
    return sum + product  // Returns 60
}
```

### Module with Structs (geometry.uni)
```unified
// File: geometry.uni
mod geometry

pub struct Point {
    pub x: Float
    pub y: Float
}

impl Point {
    pub fn new(x: Float, y: Float) -> Point {
        return Point { x: x, y: y }
    }
    
    pub fn distance(&self, other: &Point) -> Float {
        let dx = self.x - other.x
        let dy = self.y - other.y
        return sqrt(dx * dx + dy * dy)
    }
}
```

### Multi-File Project
```unified
// File: shapes/circle.uni
mod shapes::circle

use geometry::Point

pub struct Circle {
    pub center: Point
    pub radius: Float
}

impl Circle {
    pub fn new(center: Point, radius: Float) -> Circle {
        return Circle { center: center, radius: radius }
    }
    
    pub fn area(&self) -> Float {
        return 3.14159 * self.radius * self.radius
    }
}
```

```unified
// File: main.uni
use geometry::Point
use shapes::circle::Circle

fn main() -> Int {
    let center = Point::new(0.0, 0.0)
    let circle = Circle::new(center, 5.0)
    let area = circle.area()
    return 0
}
```

### Visibility and Privacy
```unified
// File: bank.uni
mod bank

pub struct Account {
    pub id: Int
    balance: Float  // Private field
}

impl Account {
    pub fn new(id: Int) -> Account {
        return Account { id: id, balance: 0.0 }
    }
    
    pub fn deposit(&mut self, amount: Float) -> Bool {
        if amount > 0.0 {
            self.balance = self.balance + amount
            return true
        }
        return false
    }
    
    pub fn get_balance(&self) -> Float {
        return self.balance
    }
}
```

## 🏆 Success Criteria

- [ ] All parser tests pass (minimum 10 tests)
- [ ] All module resolver tests pass (minimum 8 tests)
- [ ] All symbol table tests pass (minimum 10 tests)
- [ ] All bytecode generation tests pass (minimum 6 tests)
- [ ] All VM tests pass (minimum 5 tests)
- [ ] All integration tests pass (minimum 6 tests)
- [ ] Can compile multi-file projects
- [ ] Module visibility is enforced
- [ ] Circular dependencies are detected
- [ ] No regressions in previous phase tests
- [ ] Code coverage ≥ 90%
- [ ] Documentation is complete
- [ ] Example programs compile and run

## 💡 Implementation Notes

### Implementation Order
1. Grammar and parser (syntax support)
2. AST nodes (representation)
3. Module resolver (discovery and loading)
4. Symbol table (name resolution)
5. Bytecode generator updates (code generation)
6. VM updates (execution)
7. Multi-file compilation (tooling)
8. Tests and documentation

### Testing Strategy
- Test each component in isolation first
- Use mock modules for testing resolution
- Test visibility enforcement thoroughly
- Create realistic multi-module examples
- Test error cases (missing modules, circular deps, visibility violations)

### Common Pitfalls
1. **Circular Dependencies**: Implement detection early to avoid debugging nightmares
2. **Path Resolution**: Handle both relative and absolute paths consistently
3. **Symbol Conflicts**: Ensure proper scoping to avoid name collisions
4. **Visibility**: Remember private-by-default, unlike some languages
5. **Module Caching**: Cache parsed modules to avoid re-parsing
6. **Cross-Platform Paths**: Use Go's filepath package for portability

### Debugging Tips
- Add verbose logging to module resolver for troubleshooting
- Print symbol table contents when debugging name resolution
- Use graphviz to visualize module dependency graph
- Test with simple two-file examples before complex projects
- Validate AST structure after parsing modules

### Performance Considerations
- Cache parsed modules aggressively
- Lazy-load modules when possible
- Build dependency graph once, reuse for multiple compilations
- Consider parallel module parsing for large projects

### Future Extensions (Not in This Phase)
- Package versioning
- Package registry/repository
- Workspace support for multi-package projects
- Incremental compilation
- Module precompilation/caching

---

**Labels**: `phase-11`, `enhancement`, `modules`, `parser`, `compiler`, `high-priority`  
**Milestone**: Phase 11 - Modules and Imports  
**Assignee**: TBD
