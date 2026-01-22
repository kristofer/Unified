# Phase 15: Tooling and Polish

**Status**: Not Started  
**Duration**: 3-4 weeks  
**Priority**: MEDIUM  
**Dependencies**: Phase 14 Complete (Concurrency Basics)

## 🎯 Goals

Enhance developer experience through improved tooling, better error messages, code formatting, documentation generation, and basic package management. This phase makes Unified a productive and pleasant language to work with.

## 📋 Prerequisites

- [x] Phase 14 complete (basic language features working)
- [ ] Understanding of compiler error message design
- [ ] Knowledge of code formatting approaches
- [ ] Familiarity with documentation generation tools
- [ ] Understanding of package management systems

## ✨ Language Features to Add

### 1. Enhanced Error Messages
- Source code snippets in errors
- Colorized terminal output
- Helpful suggestions and hints
- Error codes and documentation links
- Multi-line error displays
- Warning system

### 2. Code Formatter
- Automatic code formatting (`unified fmt`)
- Configurable style options
- IDE integration support
- Format-on-save capability
- Check mode (verify without modifying)

### 3. Documentation Generator
- Doc comments (`///` style)
- Generate HTML documentation
- Markdown support in docs
- Code examples in documentation
- Cross-references and links
- Search functionality

### 4. Package Manager Basics
- Package manifest (`Package.uni` file)
- Dependency declaration
- Package installation
- Version resolution (basic)
- Local package development
- Package building

### 5. Build System Improvements
- Incremental compilation
- Build caching
- Parallel compilation
- Build profiles (debug/release)
- Custom build scripts

### 6. Additional Developer Tools
- Syntax highlighting definitions
- Language server protocol (LSP) foundation
- REPL improvements
- Profiler integration
- Debugger integration

## 📝 Implementation Tasks

### Task 15.1: Enhanced Error Message System (8-10 hours)
- [ ] Create `internal/diagnostics/` package
- [ ] Implement error severity levels
- [ ] Add source location tracking
- [ ] Create error message formatter
- [ ] Add color support for terminals
- [ ] Implement code snippet extraction
- [ ] Add suggestion system
- [ ] Create error catalog with codes
- [ ] Write tests for error formatting

**Error System:**
```go
type Diagnostic struct {
    Severity    Severity  // Error, Warning, Info
    Code        string    // E0001, W0001, etc.
    Message     string
    Location    SourceLocation
    Suggestions []string
    Notes       []string
}

type DiagnosticRenderer struct {
    colorize bool
    verbose  bool
}

func (r *DiagnosticRenderer) Render(diag Diagnostic) string
func (r *DiagnosticRenderer) RenderWithSnippet(diag Diagnostic, source string) string
```

**Example Error Output:**
```
error[E0382]: use of moved value: `s1`
  --> main.uni:4:14
   |
 3 |     let s2 = s1
   |              -- value moved here
 4 |     print(s1)
   |           ^^ value used after move
   |
   = note: move occurs because `s1` has type `String`, which does not implement the `Copy` trait
   = help: consider cloning the value: `s1.clone()`
```

### Task 15.2: Implement Code Formatter (10-12 hours)
- [ ] Create `cmd/unified-fmt/` command
- [ ] Implement AST pretty-printer
- [ ] Define formatting rules
- [ ] Handle comments preservation
- [ ] Support configuration file
- [ ] Add check mode (--check)
- [ ] Add diff mode (--diff)
- [ ] Integrate with build system
- [ ] Write comprehensive tests

**Formatter:**
```go
type Formatter struct {
    config *FormatConfig
}

type FormatConfig struct {
    IndentSize    int
    MaxLineLength int
    TabsOrSpaces  string
    // ... more options
}

func (f *Formatter) FormatFile(path string) error
func (f *Formatter) FormatString(source string) (string, error)
func (f *Formatter) FormatAST(ast *AST) string
```

**Configuration File (.unified-fmt.toml):**
```toml
indent_size = 4
max_line_length = 100
use_tabs = false
trailing_comma = true
```

### Task 15.3: Implement Documentation Generator (10-12 hours)
- [ ] Create `cmd/unified-doc/` command
- [ ] Parse doc comments (`///`)
- [ ] Extract documentation metadata
- [ ] Generate HTML output
- [ ] Create default templates
- [ ] Support Markdown in comments
- [ ] Generate module documentation
- [ ] Create search index
- [ ] Add cross-linking
- [ ] Write tests

**Doc Comments:**
```unified
/// Calculates the factorial of a number.
///
/// # Examples
///
/// ```
/// let result = factorial(5)
/// assert_eq!(result, 120)
/// ```
///
/// # Panics
///
/// Panics if `n` is negative.
pub fn factorial(n: Int) -> Int {
    // implementation
}
```

**Documentation Generator:**
```go
type DocGenerator struct {
    outputDir string
    template  *Template
}

type DocItem struct {
    Name        string
    Kind        string  // function, struct, enum, etc.
    Visibility  string
    Signature   string
    Description string
    Examples    []string
    SeeAlso     []string
}

func (g *DocGenerator) Generate(module *Module) error
func (g *DocGenerator) GenerateHTML(items []DocItem) error
```

### Task 15.4: Implement Package Manager (12-15 hours)
- [ ] Create `cmd/unified-pkg/` command
- [ ] Design Package.uni manifest format
- [ ] Implement package metadata parsing
- [ ] Create dependency resolver (basic)
- [ ] Implement package installation
- [ ] Support local and remote packages
- [ ] Create package registry client (basic)
- [ ] Implement version constraints
- [ ] Add lock file support
- [ ] Write comprehensive tests

**Package.uni Manifest:**
```unified
package {
    name = "my-project"
    version = "0.1.0"
    authors = ["John Doe <john@example.com>"]
    license = "MIT"
    description = "A sample Unified package"
}

dependencies {
    http = "1.0.0"
    json = "^2.1"
}

dev-dependencies {
    test-framework = "0.5"
}
```

**Package Manager:**
```go
type PackageManager struct {
    registry *Registry
    cache    *PackageCache
}

type Package struct {
    Name         string
    Version      string
    Dependencies []Dependency
}

type Dependency struct {
    Name       string
    Constraint string
}

func (pm *PackageManager) Install(pkgName, version string) error
func (pm *PackageManager) Resolve(deps []Dependency) ([]Package, error)
func (pm *PackageManager) Build(pkg *Package) error
```

### Task 15.5: Implement Incremental Compilation (8-10 hours)
- [ ] Update `internal/compiler/` for incremental builds
- [ ] Implement dependency tracking
- [ ] Create compilation cache
- [ ] Detect changed files
- [ ] Recompile only changed modules
- [ ] Implement cache invalidation
- [ ] Add cache serialization
- [ ] Write tests

**Incremental Compiler:**
```go
type IncrementalCompiler struct {
    cache     *CompilationCache
    depGraph  *DependencyGraph
}

type CompilationCache struct {
    entries map[string]*CacheEntry
}

type CacheEntry struct {
    SourceHash  string
    Bytecode    []byte
    Dependencies []string
    Timestamp   time.Time
}

func (c *IncrementalCompiler) Compile(files []string) error
func (c *IncrementalCompiler) GetChangedFiles(files []string) []string
```

### Task 15.6: Implement Warning System (4-5 hours)
- [ ] Create warning categories
- [ ] Add warning detection passes
- [ ] Implement -W flags (enable/disable warnings)
- [ ] Add warning suppression annotations
- [ ] Integrate with error reporting
- [ ] Write tests

**Warnings:**
```go
type WarningKind int

const (
    WarningUnusedVariable WarningKind = iota
    WarningUnusedImport
    WarningDeadCode
    WarningDeprecated
    // ... more warnings
)

type WarningConfig struct {
    enabled map[WarningKind]bool
    asError bool
}

func (c *Compiler) DetectWarnings(ast *AST) []Warning
```

### Task 15.7: Create Build System Improvements (6-8 hours)
- [ ] Add build profiles (debug/release)
- [ ] Implement parallel compilation
- [ ] Add build scripts support
- [ ] Create build configuration file
- [ ] Add optimization levels
- [ ] Implement clean builds
- [ ] Write tests

**Build Configuration (build.uni):**
```unified
build {
    target = "native"
    optimization = "release"
    parallel = true
}

profile.debug {
    optimization = "none"
    debug_info = true
}

profile.release {
    optimization = "aggressive"
    debug_info = false
}
```

### Task 15.8: Create Syntax Highlighting Definitions (4-5 hours)
- [ ] Create TextMate grammar
- [ ] Create VSCode extension skeleton
- [ ] Create Vim syntax file
- [ ] Create Emacs mode
- [ ] Create Sublime Text package
- [ ] Document installation

### Task 15.9: LSP Foundation (8-10 hours)
- [ ] Create `cmd/unified-lsp/` server
- [ ] Implement basic LSP protocol
- [ ] Add hover information
- [ ] Add go-to-definition
- [ ] Add autocompletion basics
- [ ] Add syntax error reporting
- [ ] Write tests
- [ ] Document LSP features

**LSP Server:**
```go
type LanguageServer struct {
    workspace *Workspace
    analyzer  *Analyzer
}

func (s *LanguageServer) HandleHover(params HoverParams) (Hover, error)
func (s *LanguageServer) HandleDefinition(params DefinitionParams) ([]Location, error)
func (s *LanguageServer) HandleCompletion(params CompletionParams) ([]CompletionItem, error)
```

### Task 15.10: Improve REPL (4-5 hours)
- [ ] Update `cmd/unified-repl/`
- [ ] Add line editing (readline)
- [ ] Add history support
- [ ] Add tab completion
- [ ] Add multi-line input
- [ ] Add special commands (:help, :type, etc.)
- [ ] Write tests

### Task 15.11: Write Comprehensive Tests (8-10 hours)
- [ ] Error message formatting tests (10+ tests)
- [ ] Formatter tests (15+ tests)
- [ ] Doc generator tests (8+ tests)
- [ ] Package manager tests (12+ tests)
- [ ] Incremental compilation tests (8+ tests)
- [ ] Warning system tests (10+ tests)
- [ ] LSP tests (10+ tests)
- [ ] Integration tests (8+ tests)

### Task 15.12: Update Documentation (5-6 hours)
- [ ] Create PHASE15_SUMMARY.md
- [ ] Create TOOLING_GUIDE.md
- [ ] Document error codes
- [ ] Document formatter options
- [ ] Document package manager usage
- [ ] Create package authoring guide
- [ ] Update README.md
- [ ] Create IDE setup guides

## ✅ Test Requirements

**Minimum Test Coverage**: 85% for tooling code

### Error Message Tests
- [ ] Format basic error
- [ ] Format error with snippet
- [ ] Colorize error output
- [ ] Format multi-line error
- [ ] Include suggestions
- [ ] Format warnings
- [ ] Test error codes
- [ ] Test severity levels
- [ ] Test note and help messages
- [ ] Verify readable output

### Formatter Tests
- [ ] Format simple function
- [ ] Format complex expressions
- [ ] Preserve comments
- [ ] Format structs and enums
- [ ] Format nested blocks
- [ ] Handle long lines
- [ ] Test indentation
- [ ] Test check mode
- [ ] Test diff mode
- [ ] Verify idempotency (format twice = same result)
- [ ] Test configuration options
- [ ] Format entire file
- [ ] Format entire project
- [ ] Handle syntax errors gracefully
- [ ] Preserve intentional formatting where appropriate

### Documentation Generator Tests
- [ ] Parse doc comments
- [ ] Extract function documentation
- [ ] Extract struct documentation
- [ ] Generate HTML output
- [ ] Render Markdown
- [ ] Generate module index
- [ ] Test cross-references
- [ ] Generate search index

### Package Manager Tests
- [ ] Parse Package.uni manifest
- [ ] Resolve simple dependency
- [ ] Resolve transitive dependencies
- [ ] Handle version constraints
- [ ] Install package
- [ ] Detect dependency conflicts
- [ ] Create lock file
- [ ] Build package
- [ ] Test local packages
- [ ] Test remote packages
- [ ] Handle missing packages
- [ ] Update dependencies

### Incremental Compilation Tests
- [ ] Detect unchanged files
- [ ] Detect changed files
- [ ] Cache compilation results
- [ ] Invalidate cache on change
- [ ] Track dependencies
- [ ] Recompile dependent modules
- [ ] Handle removed files
- [ ] Verify correctness

### Warning System Tests
- [ ] Detect unused variable
- [ ] Detect unused import
- [ ] Detect dead code
- [ ] Enable/disable warnings
- [ ] Warnings as errors
- [ ] Suppress specific warnings
- [ ] Test all warning types
- [ ] Verify warning messages
- [ ] Test warning in different contexts
- [ ] Batch warning reporting

### LSP Tests
- [ ] Start LSP server
- [ ] Handle hover request
- [ ] Handle definition request
- [ ] Handle completion request
- [ ] Report syntax errors
- [ ] Test with multiple files
- [ ] Test incremental updates
- [ ] Test protocol compliance
- [ ] Test error handling
- [ ] Verify performance

### Integration Tests
- [ ] Format entire codebase
- [ ] Generate docs for stdlib
- [ ] Install and use package
- [ ] Incremental build large project
- [ ] LSP in realistic scenario
- [ ] End-to-end developer workflow
- [ ] Test tooling interoperability
- [ ] Performance benchmarks

## 📚 Documentation Updates Required

### Update Existing Docs
- [ ] README.md: Add tooling features
- [ ] CONTRIBUTING.md: Add formatting guidelines

### Create New Docs
- [ ] PHASE15_SUMMARY.md: Complete phase summary
- [ ] docs/TOOLING_GUIDE.md: Comprehensive tooling guide
- [ ] docs/ERROR_CODES.md: Error code reference
- [ ] docs/FORMATTER_GUIDE.md: Formatter usage
- [ ] docs/PACKAGE_GUIDE.md: Package authoring guide
- [ ] docs/IDE_SETUP.md: IDE integration guides
- [ ] docs/LSP_SPEC.md: LSP capabilities

### Add Example Configurations
- [ ] `.unified-fmt.toml`: Formatter config
- [ ] `Package.uni`: Package manifest example
- [ ] `build.uni`: Build configuration
- [ ] `.unified-lint.toml`: Linter config (future)

## 🎓 Example Usage

### Enhanced Error Messages
```bash
$ unified build main.uni
error[E0382]: use of moved value: `s1`
  --> main.uni:4:14
   |
 3 |     let s2 = s1
   |              -- value moved here
 4 |     print(s1)
   |           ^^ value used after move
   |
   = note: move occurs because `s1` has type `String`
   = help: consider cloning: `s1.clone()`

warning[W0001]: unused variable: `x`
  --> main.uni:2:9
   |
 2 |     let x = 42
   |         ^ help: prefix with underscore: `_x`

error: aborting due to 1 previous error; 1 warning emitted

For more information about an error, try `unified explain E0382`
```

### Code Formatter
```bash
# Format single file
$ unified fmt main.uni

# Format entire project
$ unified fmt .

# Check without modifying
$ unified fmt --check main.uni

# Show diff
$ unified fmt --diff main.uni

# Configure via file
$ cat .unified-fmt.toml
indent_size = 4
max_line_length = 100
use_tabs = false
```

**Before Formatting:**
```unified
fn   factorial(  n:Int  )->Int{
let mut result=1
for i in 1..=n{result=result*i}
return result}
```

**After Formatting:**
```unified
fn factorial(n: Int) -> Int {
    let mut result = 1
    for i in 1..=n {
        result = result * i
    }
    return result
}
```

### Documentation Generator
```unified
/// A collection type that stores values in a contiguous memory block.
///
/// `List<T>` is a dynamic array that can grow or shrink in size.
///
/// # Examples
///
/// ```
/// let mut list = List::new()
/// list.push(1)
/// list.push(2)
/// assert_eq!(list.len(), 2)
/// ```
///
/// # Performance
///
/// - Push: O(1) amortized
/// - Pop: O(1)
/// - Index access: O(1)
pub struct List<T> {
    // ...
}
```

```bash
# Generate documentation
$ unified doc --output ./docs

# Generate docs for specific module
$ unified doc stdlib/collections

# Generate with examples
$ unified doc --examples

# Serve docs locally
$ unified doc --serve
```

### Package Manager
```bash
# Create new package
$ unified pkg new my-project
Created package `my-project`

# Add dependency
$ unified pkg add http@1.0

# Install dependencies
$ unified pkg install

# Update dependencies
$ unified pkg update

# Build package
$ unified pkg build

# Publish package (future)
$ unified pkg publish
```

**Package.uni:**
```unified
package {
    name = "web-server"
    version = "0.1.0"
    authors = ["Alice <alice@example.com>"]
    license = "MIT"
}

dependencies {
    http = "1.0.0"
    json = "^2.1.0"
    async = ">=0.5.0, <1.0"
}
```

### Incremental Compilation
```bash
# First build (compiles everything)
$ unified build
   Compiling module1.uni
   Compiling module2.uni
   Compiling main.uni
    Finished in 2.3s

# Modify one file
$ echo "// comment" >> module1.uni

# Rebuild (only recompiles changed)
$ unified build
   Compiling module1.uni (changed)
    Finished in 0.3s (incremental)
```

### Language Server Protocol
```bash
# Start LSP server
$ unified-lsp

# Use with VSCode
# Install "Unified Language Support" extension

# Use with Vim/Neovim
# Configure LSP client to use unified-lsp
```

**LSP Features:**
- Hover for type information
- Go to definition
- Find references
- Auto-completion
- Syntax error highlighting
- Code actions (quick fixes)

### Warning Configuration
```bash
# Allow unused variables
$ unified build -W unused-variables

# Warnings as errors
$ unified build --warnings-as-errors

# Disable all warnings
$ unified build -W none
```

**In Code:**
```unified
// Suppress specific warning
#[allow(unused_variables)]
fn example() {
    let x = 42  // No warning
}
```

## 🏆 Success Criteria

- [ ] Error messages are clear and helpful
- [ ] Error messages include code snippets
- [ ] Error messages suggest fixes
- [ ] Formatter produces consistent output
- [ ] Formatter is idempotent
- [ ] Documentation is generated correctly
- [ ] Docs include examples and cross-links
- [ ] Package manager can install dependencies
- [ ] Package manager resolves versions correctly
- [ ] Incremental compilation works
- [ ] Incremental builds are faster
- [ ] Warning system detects common issues
- [ ] LSP server works with major editors
- [ ] All tests pass (minimum 81 tests)
- [ ] No regressions in previous phases
- [ ] Code coverage ≥ 85%
- [ ] Documentation is comprehensive
- [ ] Developer experience is significantly improved

## 💡 Implementation Notes

### Implementation Order
1. Enhanced error messages (foundation)
2. Warning system
3. Code formatter
4. Documentation generator
5. Incremental compilation
6. Package manager
7. LSP foundation
8. Build system improvements
9. Syntax highlighting
10. Tests and documentation

### Testing Strategy
- Test error formatting with many examples
- Verify formatter idempotency
- Test package manager with realistic scenarios
- Benchmark incremental compilation speedup
- Test LSP with real editors
- User testing for developer experience
- Test all tools together in workflows

### Common Pitfalls
1. **Error Messages**: Too technical or not helpful enough
2. **Formatter**: Breaking valid code or losing comments
3. **Doc Generator**: Not handling edge cases in Markdown
4. **Package Manager**: Version resolution conflicts
5. **Incremental Compilation**: Cache invalidation bugs
6. **LSP**: Performance issues with large files
7. **Cross-Platform**: Path handling differences

### Debugging Tips
- Test error messages with actual compiler errors
- Compare formatter output manually
- Validate generated HTML documentation
- Test package manager with local registry
- Profile incremental compilation cache
- Use LSP test clients
- Monitor tooling performance

### Performance Considerations
- Optimize error message rendering
- Make formatter fast enough for format-on-save
- Cache documentation generation results
- Optimize package dependency resolution
- Minimize incremental compilation overhead
- Keep LSP responses under 100ms
- Profile all tools regularly

### User Experience Focus
- Error messages should be actionable
- Formatter should be transparent
- Documentation should be browsable
- Package manager should be intuitive
- Build times should be fast
- LSP should be responsive
- Tools should integrate seamlessly

### Future Extensions (Not This Phase)
- Advanced linting rules
- Code refactoring tools
- Integrated testing framework UI
- Performance profiler UI
- Memory profiler
- Continuous integration helpers
- Package publishing and registry
- Advanced LSP features (refactoring, etc.)
- Debugger GUI
- Project templates
- Migration tools

---

**Labels**: `phase-15`, `enhancement`, `tooling`, `developer-experience`, `polish`  
**Milestone**: Phase 15 - Tooling and Polish  
**Assignee**: TBD
