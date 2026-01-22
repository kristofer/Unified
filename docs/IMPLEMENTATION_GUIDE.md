# Unified Language Implementation Guide

**For AI Agents and Human Developers**

## Quick Start

👉 **Ready to implement a phase?** Start here: [AI Implementation Plan](planning/AI_IMPLEMENTATION_PLAN.md)

## Overview

This guide helps you navigate the Unified programming language implementation documentation and get started with development.

## Documentation Structure

### For AI Agents

1. **[AI Implementation Plan](planning/AI_IMPLEMENTATION_PLAN.md)** - **START HERE**
   - Detailed phase-by-phase implementation guide
   - Task-level breakdown for each phase
   - Test requirements and success criteria
   - Example programs for each phase
   - AI-specific implementation notes

2. **[Project Status](PROJECT_STATUS.md)**
   - Current implementation state
   - Feature completion status
   - Next priorities

3. **[Testing Guide](../src/unified-compiler/TESTING.md)**
   - How to write and run tests
   - Test categories and structure
   - Coverage requirements

### For Understanding the Language

1. **[Language Specification](../spec/UnifiedSpecifiation.md)**
   - Complete language specification (2100+ lines)
   - All language features defined
   - Grammar and semantics

2. **[Phased Roadmap](planning/PHASED_ROADMAP.md)**
   - High-level overview of all phases
   - Phase dependencies
   - Timeline and priorities

3. **[Architecture Guide](design/ARCHITECTURE.md)**
   - Compiler architecture
   - VM design
   - System overview

### For Day-to-Day Development

1. **[Contributing Guide](../CONTRIBUTING.md)**
   - Development workflow
   - Code standards
   - Pull request process

2. **[README](../README.md)**
   - Project overview
   - Quick links
   - Getting started

3. **[VM Documentation](../src/unified-compiler/VM_README.md)**
   - VM architecture
   - Bytecode instruction set
   - Execution model

## Current State (January 2026)

### ✅ What's Done

**Phase 1: VM-Based Minimal Compiler - COMPLETE**

- Stack-based virtual machine
- Bytecode generator
- Function calls and returns
- Arithmetic, comparison, logical operations
- Local variables
- 76 tests passing (100%)
- Full documentation

### 🎯 What's Next

**Phase 2: Control Flow** (Next priority)

- If/else statements
- While loops
- For loops with ranges
- Break/continue statements

See [Phase 2 detailed plan](planning/AI_IMPLEMENTATION_PLAN.md#phase-2-control-flow-and-parser-grammar-updates) for implementation guide.

## How to Use This Documentation

### If You're an AI Agent Implementing a Phase

1. Read the current [Project Status](PROJECT_STATUS.md)
2. Open the [AI Implementation Plan](planning/AI_IMPLEMENTATION_PLAN.md)
3. Find your phase (probably Phase 2)
4. Follow the tasks sequentially
5. Write tests as you go
6. Update documentation when done
7. Use `report_progress` tool to commit changes

### If You're a Human Developer

1. Read the [Contributing Guide](../CONTRIBUTING.md)
2. Check [Project Status](PROJECT_STATUS.md) for current state
3. Review [AI Implementation Plan](planning/AI_IMPLEMENTATION_PLAN.md) for your area
4. Set up development environment (see README)
5. Pick a task and start coding!

### If You're Learning About Unified

1. Read the [README](../README.md) for overview
2. Check [Language Specification](../spec/UnifiedSpecifiation.md)
3. Look at example programs in `src/unified-compiler/test/`
4. Read [Learn Unified in Y Minutes](../LearnUnifiedinYMinutes.md)

## Development Workflow

### Standard Development Cycle

```bash
# 1. Check current state
cd /home/runner/work/Unified/Unified/src/unified-compiler
git status
make test

# 2. Make changes
# ... edit files ...

# 3. Test continuously
make test                    # All tests
go test ./internal/bytecode  # Specific module

# 4. Generate parser if grammar changed
make parser

# 5. Build
make build

# 6. Commit (use report_progress tool for AI agents)
git add .
git commit -m "Implement feature X"
git push
```

### Testing Workflow

```bash
# Run all tests
make test

# Run specific module tests
go test ./internal/vm -v

# Run with coverage
go test ./... -cover

# Run single test
go test ./internal/vm -run TestVMSimpleArithmetic -v
```

## Phase Implementation Checklist

When implementing a phase, use this checklist:

- [ ] Read phase documentation completely
- [ ] Understand prerequisites and dependencies
- [ ] Review example programs for the phase
- [ ] Implement features following task order
- [ ] Write tests for each feature (TDD preferred)
- [ ] Run tests continuously
- [ ] Update grammar if needed (regenerate parser)
- [ ] Add integration tests
- [ ] Write documentation
- [ ] Create PHASE_N_SUMMARY.md
- [ ] Update PROJECT_STATUS.md
- [ ] Update README.md if needed
- [ ] Run final test suite (100% passing)
- [ ] Request code review
- [ ] Address review feedback
- [ ] Mark phase complete

## Key Files and Locations

### Source Code

```
src/unified-compiler/
├── cmd/compiler/main.go          # Compiler entry point
├── grammar/                      # ANTLR4 grammar files
│   ├── UnifiedLexer.g4
│   └── UnifiedParser.g4
├── internal/
│   ├── ast/                      # Abstract Syntax Tree
│   ├── bytecode/                 # Bytecode generator
│   │   ├── generator.go
│   │   └── instructions.go
│   ├── parser/                   # Generated parser (ANTLR)
│   └── vm/                       # Virtual machine
│       ├── vm.go
│       └── stack.go
└── test/                         # Test programs (.uni files)
```

### Documentation

```
docs/
├── IMPLEMENTATION_GUIDE.md       # This file
├── PROJECT_STATUS.md             # Current status
├── planning/
│   ├── AI_IMPLEMENTATION_PLAN.md # Detailed phase plans
│   └── PHASED_ROADMAP.md         # High-level roadmap
└── design/
    └── ARCHITECTURE.md           # Architecture docs
```

### Tests

```
src/unified-compiler/
├── internal/bytecode/
│   ├── generator_test.go         # Generator tests
│   └── instructions_test.go      # Instruction tests
├── internal/vm/
│   ├── vm_test.go                # VM execution tests
│   └── stack_test.go             # Stack tests
└── cmd/compiler/
    └── integration_test.go       # End-to-end tests
```

## Common Tasks

### Add a New Language Feature

1. Update grammar in `grammar/UnifiedParser.g4`
2. Regenerate parser: `make parser`
3. Add AST node in `internal/ast/ast.go`
4. Update visitor in `internal/ast/visitor.go`
5. Add bytecode generation in `internal/bytecode/generator.go`
6. Add VM support if needed in `internal/vm/vm.go`
7. Write tests at each level
8. Add integration test
9. Document in appropriate PHASE_N_SUMMARY.md

### Add a New Test

1. Find appropriate test file (or create new one)
2. Write test function:
   ```go
   func TestFeatureName(t *testing.T) {
       // Arrange
       // Act
       // Assert
   }
   ```
3. Run: `go test ./path/to/package -v`

### Debug a Failing Test

1. Run test verbosely: `go test ./... -v`
2. Add debug output in code
3. Check bytecode generation (add print statements)
4. Trace VM execution
5. Verify AST construction

### Update Documentation

1. Add examples to test/ directory
2. Update relevant PHASE_N_SUMMARY.md
3. Update PROJECT_STATUS.md completion status
4. Update README.md if major feature
5. Add to language spec if new syntax

## Resources

### Internal Documentation

- [Language Specification](../spec/UnifiedSpecifiation.md)
- [VM Architecture](../src/unified-compiler/VM_README.md)
- [Testing Guide](../src/unified-compiler/TESTING.md)
- [Phase 1 Summary](../src/unified-compiler/PHASE1_SUMMARY.md)

### Learning Materials

- [Learn Unified in Y Minutes](../LearnUnifiedinYMinutes.md)
- [Grammar Documentation](../GrammarForUnified.md)
- Example programs in `src/unified-compiler/test/`

### Reference Projects

- [kristofer/smog](https://github.com/kristofer/smog) - Well-structured example project

## Getting Help

### For AI Agents

- Consult [AI Implementation Plan](planning/AI_IMPLEMENTATION_PLAN.md)
- Check "AI Agent Notes" sections in phase documentation
- Review "Common Pitfalls" in implementation plan

### For Human Developers

- Check existing documentation first
- Review similar implementations in codebase
- Look at test files for examples
- Read phase summaries for context

## FAQ

**Q: Where do I start as an AI agent?**  
A: Read [AI Implementation Plan](planning/AI_IMPLEMENTATION_PLAN.md) and start with Phase 2.

**Q: How do I know what phase we're on?**  
A: Check [Project Status](PROJECT_STATUS.md) - currently ready for Phase 2.

**Q: What if I need to change the grammar?**  
A: Edit `.g4` files in `grammar/`, then run `make parser` to regenerate.

**Q: How much test coverage is required?**  
A: Aim for ≥85% coverage. All new features must have tests.

**Q: Can I skip documentation?**  
A: No. Documentation is part of phase completion criteria.

**Q: What if tests fail after my changes?**  
A: Fix regressions immediately. Never commit broken tests.

**Q: How do I add a new bytecode instruction?**  
A: Add to OpCode enum in `instructions.go`, implement in `vm.go`, test in `vm_test.go`.

**Q: Where should example programs go?**  
A: Put `.uni` files in `src/unified-compiler/test/` directory.

## Next Steps

1. **Check your status**: Read [Project Status](PROJECT_STATUS.md)
2. **Pick your phase**: Review [AI Implementation Plan](planning/AI_IMPLEMENTATION_PLAN.md)
3. **Start coding**: Follow the tasks for your phase
4. **Test thoroughly**: Write tests as you implement
5. **Document everything**: Update docs as you go
6. **Report progress**: Use `report_progress` tool regularly

Happy coding! 🚀
