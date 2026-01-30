# Contributing to Unified Programming Language

Thank you for your interest in contributing to the Unified Programming Language! This document provides guidelines and information for contributors.

## Table of Contents

1. [Getting Started](#getting-started)
2. [Development Process](#development-process)
3. [Phased Implementation](#phased-implementation)
4. [Code Style](#code-style)
5. [Testing](#testing)
6. [Documentation](#documentation)
7. [Pull Requests](#pull-requests)
8. [Community](#community)

## Getting Started

### Prerequisites

- **Go 1.21+**: Install from [golang.org](https://golang.org)
- **ANTLR4**: Install ANTLR4 for parser generation
  ```bash
  # On macOS with Homebrew
  brew install antlr
  
  # Or download from https://www.antlr.org/download.html
  ```
- **Git**: For version control
- **Make**: For build automation

### Setting Up Development Environment

1. **Clone the repository**
   ```bash
   git clone https://github.com/kristofer/Unified.git
   cd Unified
   ```

2. **Build the compiler**
   ```bash
   cd src/unified-compiler
   make all  # Generates parser and builds compiler
   ```

3. **Run tests**
   ```bash
   make test
   ```

4. **Read the documentation**
   - `docs/planning/PHASED_ROADMAP.md` - Implementation roadmap
   - `docs/COPILOT_GUIDE.md` - Guide for AI-assisted development
   - `docs/design/ARCHITECTURE.md` - Compiler architecture
   - `spec/UnifiedSpecifiation.md` - Language specification

## Development Process

### 1. Choose a Task

**For New Contributors:**
- Start with Phase 0 or Phase 1 tasks
- Look for issues labeled "good first issue"
- Focus on documentation improvements
- Write example programs

**For Experienced Contributors:**
- Check the current phase in `docs/planning/PHASED_ROADMAP.md`
- Pick tasks from the current phase
- Avoid implementing features from future phases

### 2. Create a Branch

```bash
git checkout -b feature/your-feature-name
# or
git checkout -b fix/bug-description
```

Branch naming conventions:
- `feature/` - New features
- `fix/` - Bug fixes
- `docs/` - Documentation changes
- `test/` - Test additions/improvements
- `refactor/` - Code refactoring

### 3. Make Changes

Follow these principles:
- **Test-Driven Development**: Write tests before implementation
- **Small Commits**: Make small, focused commits
- **Clear Messages**: Write clear commit messages
- **Documentation**: Update docs as you code

### 4. Test Your Changes

Always test your changes before submitting a pull request.

#### Run All Tests
```bash
cd src/unified-compiler
make test
```

#### Run Tests for Specific Components
```bash
# Test a specific package
go test ./internal/wasm -v
go test ./internal/ast -v
go test ./cmd/compiler -v

# Test all packages with verbose output
go test -v ./...
```

#### Run Specific Tests
```bash
# Run a single test function
go test ./internal/wasm -run TestWASMGenerator -v

# Run tests matching a pattern
go test ./internal/wasm -run TestWASM.* -v
```

#### Check Test Coverage
```bash
# View coverage summary
go test ./... -cover

# Generate detailed HTML coverage report
go test ./... -coverprofile=coverage.out
go tool cover -html=coverage.out
```

#### Run Integration Tests
```bash
# Build the compiler first
make build

# Run all .uni test files (from repository root)
cd ../..
./test_all_uni_files.sh

# Run specific example programs (from compiler directory)
cd src/unified-compiler
./bin/unified --input test/integration/simple_return.uni
./bin/unified --input test/integration/function_call.uni
./bin/unified --input test/integration/local_variables.uni
```

**Important:** All tests must pass before submitting a PR. The CI system will also run these tests.

**Current Test Status:** 26 of 121 .uni files passing (21.5%). See [TODO.md](../../TODO.md) for details.

### 5. Submit Pull Request

See [Pull Requests](#pull-requests) section below.

## Phased Implementation

The Unified compiler has migrated to a WebAssembly backend. **Current focus is on achieving feature parity.**

### Current Focus

See [TODO.md](TODO.md) for the complete implementation roadmap.

**Priority Tasks:**
1. Fix struct support (WASM global.get index bugs)
2. Fix array support (memory allocation and indexing)
3. Implement for loops with range operator
4. Fix string operations (type mismatches and runtime functions)

### Development Workflow

1. **Pick a Task from TODO.md**
   - Start with Priority 1 items
   - Each task shows affected test files
   - Tasks include expected error patterns

2. **Write/Update Tests**
   - Test files already exist in the repository
   - Run `./test_all_uni_files.sh` to verify changes
   - Focus on making failing tests pass

3. **Implement Features**
   - Update WASM code generator (`internal/wasm/codegen.go`)
   - Add runtime functions if needed (`internal/wasm/runtime.go`)
   - Update parser/AST if syntax support is missing

4. **Verify Success**
   - Run test suite again
   - Check that test count improves
   - Update TODO.md with progress

5. **Move to Next Phase**
   - Only after current phase is complete
   - Get review and approval

## Code Style

### Go Code

Follow the [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments):

- Use `gofmt` for formatting
- Use `golint` for linting
- Follow standard Go idioms
- Add comments for exported types and functions

**Example:**
```go
// ParseFunction parses a function declaration from the parse tree.
// It returns a FunctionDecl AST node or an error if parsing fails.
func (v *Visitor) ParseFunction(ctx *parser.FunctionDeclContext) (*ast.FunctionDecl, error) {
    if ctx == nil {
        return nil, fmt.Errorf("function context is nil")
    }
    
    name := ctx.IDENTIFIER().GetText()
    // ... rest of implementation
}
```

### ANTLR Grammar

- Use clear, descriptive rule names
- Add comments for complex rules
- Use labels for alternatives (`# RuleName`)
- Group related rules together

**Example:**
```antlr
// Expression with proper operator precedence
expression
    : primary                              # PrimaryExpression
    | expression ('*' | '/' | '%') expression  # MultiplicativeExpression
    | expression ('+' | '-') expression    # AdditiveExpression
    | expression ('==' | '!=') expression  # EqualityExpression
    ;
```

### Unified Language Code

- 4-space indentation (no tabs)
- Follow Rust-like conventions
- Clear variable names
- Comments for complex logic

**Example:**
```unified
// Calculate factorial iteratively
fn factorial(n: i32) -> i32 {
    let mut result = 1
    let mut i = 2
    
    while i <= n {
        result = result * i
        i = i + 1
    }
    
    return result
}
```

## Testing

### Running Tests Locally

Before writing new tests, familiarize yourself with the existing test suite:

```bash
# Navigate to the compiler directory
cd src/unified-compiler

# Run all tests
make test

# Run tests with verbose output to see individual test results
go test -v ./...

# Run tests for a specific package
go test ./internal/vm -v
go test ./internal/bytecode -v
go test ./cmd/compiler -v

# Run a specific test function
go test ./internal/vm -run TestVMSimpleArithmetic -v

# Check test coverage
go test ./... -cover

# Generate coverage report (opens in browser)
go test ./... -coverprofile=coverage.out
go tool cover -html=coverage.out
```

See the [Testing Guide](src/unified-compiler/TESTING.md) for comprehensive testing documentation.

### Test Organization

Tests are organized by phase and type:

```
src/unified-compiler/
├── test/                    # Integration tests (.uni files)
│   ├── phase1/             # Phase 1 test programs
│   ├── phase2/             # Phase 2 test programs
│   └── ...
└── internal/
    ├── ast/
    │   └── ast_test.go     # Unit tests for AST
    ├── codegen/
    │   └── genrator_test.go  # Unit tests for code gen
    └── ...
```

### Writing Tests

#### Unit Tests (Go)

```go
func TestFunctionParsing(t *testing.T) {
    tests := []struct {
        name     string
        input    string
        expected *ast.FunctionDecl
    }{
        {
            name:  "simple function",
            input: "fn main() -> i32 { return 42 }",
            expected: &ast.FunctionDecl{
                Name: "main",
                // ... expected structure
            },
        },
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := parseFun(tt.input)
            if !reflect.DeepEqual(result, tt.expected) {
                t.Errorf("got %+v, want %+v", result, tt.expected)
            }
        })
    }
}
```

#### Integration Tests (.uni files)

Create test files with expected behavior documented:

```unified
// test/phase2/arithmetic.uni
// Tests basic arithmetic operations
// Expected output: 12

fn main() -> i32 {
    let a = 5
    let b = 7
    return a + b
}
```

### Test Coverage

Aim for:
- **80%+ code coverage** for core components
- **100% coverage** for critical paths (type checking, memory safety)
- **All phases must have comprehensive tests**

Check coverage:
```bash
go test -cover ./...
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

## Documentation

### What to Document

1. **Code Comments**
   - All exported functions, types, constants
   - Complex algorithms
   - Non-obvious design decisions

2. **Documentation Files**
   - Update relevant docs when adding features
   - Add examples for new features
   - Update PHASED_ROADMAP.md when completing phases

3. **Test Documentation**
   - Document test intent
   - Explain expected behavior
   - Add examples of usage

### Documentation Style

- Use Markdown for documentation files
- Use clear, concise language
- Include code examples
- Add diagrams where helpful

## Pull Requests

### Before Submitting

- [ ] All tests pass (`make test`)
- [ ] Code is formatted (`gofmt`, `go fmt ./...`)
- [ ] Lint checks pass (`golint ./...`)
- [ ] Documentation is updated
- [ ] Commit messages are clear
- [ ] Branch is up to date with main

### PR Description Template

```markdown
## Description
Brief description of changes

## Related Issue
Closes #123

## Type of Change
- [ ] Bug fix
- [ ] New feature
- [ ] Documentation update
- [ ] Refactoring

## Phase
Which phase does this relate to? (e.g., Phase 2)

## Testing
- [ ] Unit tests added/updated
- [ ] Integration tests added/updated
- [ ] All tests pass

## Checklist
- [ ] Code follows style guidelines
- [ ] Self-review completed
- [ ] Comments added for complex code
- [ ] Documentation updated
- [ ] No new warnings generated
```

### Review Process

1. Automated checks run (CI/CD)
2. Code review by maintainers
3. Address feedback
4. Approval and merge

## Community

### Communication

- **GitHub Issues**: Bug reports, feature requests
- **GitHub Discussions**: Questions, ideas, general discussion
- **Pull Requests**: Code contributions

### Code of Conduct

- Be respectful and professional
- Welcome newcomers
- Provide constructive feedback
- Focus on the code, not the person
- Follow the [Contributor Covenant](https://www.contributor-covenant.org/)

### Getting Help

- Read the documentation first
- Search existing issues
- Ask in GitHub Discussions
- Provide context and examples when asking questions

## Recognition

Contributors will be:
- Listed in CONTRIBUTORS.md
- Credited in release notes
- Acknowledged for significant contributions

## License

By contributing, you agree that your contributions will be licensed under the same license as the project (see LICENSE file).

## Questions?

If you have questions about contributing:
1. Check the documentation in `docs/`
2. Search GitHub Issues and Discussions
3. Open a new Discussion
4. Contact the maintainers

Thank you for contributing to Unified! 🚀
