# Quick Reference - Unified Development

A quick reference card for developers working on the Unified Programming Language.

## 📚 Essential Documents

| Document | Purpose | When to Read |
|----------|---------|--------------|
| [PHASED_ROADMAP.md](planning/PHASED_ROADMAP.md) | Complete 15-phase plan | Before starting any work |
| [COPILOT_GUIDE.md](COPILOT_GUIDE.md) | AI development guide | When using AI assistants |
| [ARCHITECTURE.md](design/ARCHITECTURE.md) | Compiler architecture | When working on compiler |
| [PROJECT_STATUS.md](PROJECT_STATUS.md) | Current status | To see what's done |
| [CONTRIBUTING.md](../CONTRIBUTING.md) | How to contribute | Before first contribution |

## 🎯 Current Phase

**Phase 0**: Project Foundation and Planning (95% complete)

**Next Phase**: Phase 1 - Minimal Compiler Pipeline

## 🏗️ Quick Start

```bash
# Clone and setup
git clone https://github.com/kristofer/Unified.git
cd Unified/src/unified-compiler

# Build
make all        # Generate parser and build compiler
make parser     # Just regenerate parser
make build      # Just build compiler
make test       # Run tests
make clean      # Clean up

# Test compilation
./bin/unified-compiler --input test/simple.uni --output simple.ll
```

## 📋 Common Tasks

### Adding a New Language Feature

1. ✅ Check current phase allows this feature
2. ✅ Write tests first
3. Update grammar files (`.g4`)
4. Run `make parser` to regenerate
5. Update AST structures (`ast.go`)
6. Update visitor (`visitor.go`)
7. Add code generation (`generator.go`)
8. Run tests
9. Update documentation

### Running Tests

```bash
# All tests
make test

# Specific package
go test -v ./internal/ast
go test -v ./internal/codegen

# With coverage
go test -cover ./...
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

### Debugging

```bash
# Enable debug output
export DEBUG=1

# Run with verbose output
go run ./cmd/compiler --input test/simple.uni --output simple.ll --verbose

# View generated AST
# (add debug prints in visitor.go)

# Check generated LLVM IR
cat simple.ll
```

## 📝 Code Style

### Go
- Use `gofmt` and `golint`
- Clear variable names
- Add comments for exported functions
- Early returns to reduce nesting

### ANTLR Grammar
- camelCase for rule names
- Use labels: `# RuleName`
- Comments for complex rules

### Unified Language
- 4-space indentation
- Follow Rust-like conventions
- Prefer `let` over `let mut`

## 🧪 Test Requirements

Every feature needs:
- ✅ Unit tests (Go)
- ✅ Integration tests (.uni files)
- ✅ Documentation
- ✅ Example programs

## 🚫 Common Pitfalls

1. **Forgot to regenerate parser** → Run `make parser`
2. **Nil pointer errors** → Always check for nil
3. **Tests failing** → Fix immediately, don't commit
4. **Working on wrong phase** → Check PHASED_ROADMAP.md

## 📊 Project Structure

```
Unified/
├── docs/              # All documentation
├── examples/          # Example programs
├── spec/              # Language specification
├── src/
│   └── unified-compiler/
│       ├── cmd/           # Executables
│       ├── grammar/       # ANTLR4 grammars
│       ├── internal/      # Internal packages
│       │   ├── ast/       # Abstract Syntax Tree
│       │   ├── parser/    # Generated parser
│       │   └── codegen/   # Code generation
│       └── test/          # Test programs
└── CONTRIBUTING.md    # How to contribute
```

## 🔧 Tools

### Required
- Go 1.21+
- ANTLR4
- Make

### Recommended
- golint
- staticcheck
- VS Code with Go extension

## 📞 Getting Help

1. Check documentation in `docs/`
2. Search GitHub Issues
3. Open a Discussion
4. Ask in PR comments

## ⚡ Quick Commands

```bash
# Generate parser
make parser

# Build compiler
make build

# Run tests
make test

# Clean build artifacts
make clean

# Full rebuild
make clean all

# Compile a .uni file
./bin/unified-compiler --input file.uni --output file.ll

# Run compiled program (future)
lli file.ll
```

## 📈 Phase Checklist

Before claiming a phase is done:

- [ ] All tests pass
- [ ] Documentation updated
- [ ] Example programs work
- [ ] Code reviewed
- [ ] Success criteria met
- [ ] PR approved and merged

## 🎨 AI Development

When using GitHub Copilot:

1. Read [COPILOT_GUIDE.md](COPILOT_GUIDE.md)
2. Focus on current phase only
3. Write tests first
4. Follow code style
5. Document as you go
6. Check success criteria

## 🔗 External Resources

- [ANTLR4 Docs](https://github.com/antlr/antlr4/blob/master/doc/index.md)
- [LLVM Docs](https://llvm.org/docs/)
- [Go Docs](https://golang.org/doc/)
- [smog Project](https://github.com/kristofer/smog)

## 📅 Timeline

- **Q1 2026**: Phases 0-2 (Foundation, minimal compiler, expressions)
- **Q2 2026**: Phases 3-5 (Control flow, functions, types)
- **Q3 2026**: Phases 6-8 (Memory, patterns, generics)
- **Q4 2026**: Phases 9-10 (Stdlib, async)
- **Q1 2027**: Phases 11-12 (Actors, advanced types)
- **Q2 2027**: Phases 13-14 (Optimization, tooling)
- **Ongoing**: Phase 15 (Stdlib expansion)

## ✅ Success Metrics

Each phase tracks:
- Test coverage %
- Test pass rate %
- Documentation coverage %
- Example programs working
- Compilation time
- Runtime performance

---

**Remember**: Quality over speed. Write tests, document well, and build incrementally.

For detailed information, see the full documentation in `docs/`.
