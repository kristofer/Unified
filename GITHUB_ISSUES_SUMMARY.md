# GitHub Issue Templates for Unified Compiler Phases

## Overview

This document describes the comprehensive GitHub issue templates created for the Unified programming language compiler implementation. These templates provide detailed, actionable specifications for each development phase.

## What Was Created

### 📋 Issue Templates (14 files)

Comprehensive issue templates for Phases 2-15 of the compiler implementation:

| Phase | File | Size | Duration | Focus |
|-------|------|------|----------|-------|
| 2 | phase-02-control-flow.md | 7.6 KB | 1-2 weeks | Control flow (if/else, while, for, loop) |
| 3 | phase-03-variables-mutability.md | 11 KB | 2 weeks | Variables, mutability, symbol tables |
| 4 | phase-04-functions-advanced.md | 13 KB | 2-3 weeks | Tuples, lambdas, default parameters |
| 5 | phase-05-structs-types.md | 14 KB | 3-4 weeks | Struct declarations and methods |
| 6 | phase-06-enums-pattern-matching.md | 14 KB | 3-4 weeks | Enums and pattern matching |
| 7 | phase-07-error-handling.md | 12 KB | 1-2 weeks | Error handling with ? operator |
| 8 | phase-08-arrays-slices.md | 13 KB | 2-3 weeks | Arrays and slices |
| 9 | phase-09-string-operations.md | 14 KB | 2 weeks | String manipulation |
| 10 | phase-10-generics.md | 16 KB | 3-4 weeks | Generic programming |
| 11 | phase-11-modules-imports.md | 13 KB | 2-3 weeks | Module system |
| 12 | phase-12-basic-ownership.md | 16 KB | 3-4 weeks | Ownership and borrowing |
| 13 | phase-13-stdlib-foundation.md | 16 KB | 2-3 weeks | Standard library core |
| 14 | phase-14-concurrency-basics.md | 18 KB | 4-6 weeks | Async/await and actors |
| 15 | phase-15-tooling-polish.md | 20 KB | 3-4 weeks | Tooling and developer experience |

**Total**: ~8,000 lines of comprehensive phase specifications

### 📚 Documentation (3 files)

1. **README.md** (8.4 KB) - Complete guide to the templates and phase overview
2. **HOWTO.md** (7.0 KB) - Step-by-step instructions for creating GitHub issues
3. **create-all-issues.sh** (3.6 KB) - Automated script for issue creation

### 📊 Statistics

- **14 phase templates** covering Phases 2-15
- **~8,000 total lines** of detailed specifications
- **700+ implementation tasks** across all phases (50-70 per phase)
- **560+ test requirements** (40+ per phase)
- **80+ example programs** demonstrating features (5-8 per phase)
- **Estimated timeline**: 35-48 weeks (8-11 months)

## What Each Template Contains

Every phase template includes:

### 1. **Metadata**
- Status (Not Started / In Progress / Complete)
- Estimated duration
- Priority level
- Dependencies on previous phases

### 2. **Goals & Prerequisites**
- Clear statement of what the phase achieves
- List of prerequisites that must be complete
- Overview of language features being added

### 3. **Implementation Tasks** (50-70 per phase)
Each task includes:
- Clear, actionable description
- Time estimate
- Checkbox for tracking completion
- Code examples where applicable
- References to relevant files

Example:
```markdown
#### Task 2.1: Update Grammar for If/Else (2-3 hours)
- [ ] Open `grammar/UnifiedParser.g4`
- [ ] Add/verify if statement grammar rule
- [ ] Update statement rule to include ifStatement
- [ ] Regenerate parser: `make parser`
- [ ] Verify parser generation succeeds
- [ ] Test parsing simple if statement
```

### 4. **Test Requirements** (40+ per phase)
Organized by category:
- Parser tests (grammar/AST)
- Generator tests (bytecode generation)
- VM tests (execution)
- Integration tests (end-to-end)
- Semantic analysis tests
- Error handling tests

### 5. **Documentation Updates**
- Files to create or update
- Example programs to add
- Guides to write
- Specific documentation sections

### 6. **Example Programs** (5-8 per phase)
Working code examples demonstrating:
- New language features
- Common use cases
- Edge cases
- Integration with existing features

### 7. **Success Criteria**
Measurable completion criteria:
- Minimum test counts
- Code coverage targets
- Working example programs
- No regressions
- Documentation completeness

### 8. **Implementation Notes**
- Recommended implementation order
- Testing strategies
- Common pitfalls to avoid
- Debugging tips
- Performance considerations
- Future compatibility notes

### 9. **Labels & Metadata**
Suggested GitHub labels:
- Phase identifier (`phase-2`, `phase-3`, etc.)
- Component tags (`parser`, `vm`, `stdlib`, etc.)
- Priority level
- Feature area

## How to Use These Templates

### Important Note

**These are templates, not actual GitHub issues.** You need to create the issues from these templates using one of the following methods:

### Method 1: Automated (Recommended)

Use the provided script with GitHub CLI:

```bash
cd /home/runner/work/Unified/Unified
./docs/issues/create-all-issues.sh
```

**Requirements**: GitHub CLI installed and authenticated

### Method 2: Manual

1. Go to https://github.com/kristofer/Unified/issues/new
2. Copy content from a template file
3. Paste into issue description
4. Set title and labels
5. Create issue

### Method 3: GitHub API

Use the GitHub REST API with a personal access token (see HOWTO.md for details).

## Detailed Instructions

See the following files in `docs/issues/`:

- **[README.md](docs/issues/README.md)** - Overview and phase listing
- **[HOWTO.md](docs/issues/HOWTO.md)** - Step-by-step creation guide
- **create-all-issues.sh** - Automated creation script

## File Locations

All templates and documentation are in:
```
docs/issues/
├── README.md                          # Overview and guide
├── HOWTO.md                           # How to create issues
├── create-all-issues.sh              # Automation script
├── phase-02-control-flow.md          # Phase 2 template
├── phase-03-variables-mutability.md  # Phase 3 template
├── phase-04-functions-advanced.md    # Phase 4 template
├── phase-05-structs-types.md         # Phase 5 template
├── phase-06-enums-pattern-matching.md # Phase 6 template
├── phase-07-error-handling.md        # Phase 7 template
├── phase-08-arrays-slices.md         # Phase 8 template
├── phase-09-string-operations.md     # Phase 9 template
├── phase-10-generics.md              # Phase 10 template
├── phase-11-modules-imports.md       # Phase 11 template
├── phase-12-basic-ownership.md       # Phase 12 template
├── phase-13-stdlib-foundation.md     # Phase 13 template
├── phase-14-concurrency-basics.md    # Phase 14 template
└── phase-15-tooling-polish.md        # Phase 15 template
```

## Phase Dependencies

Phases must be completed in order due to dependencies:

```
Phase 1 (Complete) ✅
    └─> Phase 2: Control Flow
           └─> Phase 3: Variables & Mutability
                  └─> Phase 4: Functions Advanced
                         └─> Phase 5: Structs & Types
                                └─> Phase 6: Enums & Pattern Matching
                                       └─> Phase 7: Error Handling
                                              └─> Phase 8: Arrays & Slices
                                                     └─> Phase 9: String Operations
                                                            └─> Phase 10: Generics
                                                                   ├─> Phase 11: Modules & Imports
                                                                   └─> Phase 12: Basic Ownership
                                                                          └─> Phase 13: Stdlib Foundation
                                                                                 └─> Phase 14: Concurrency
                                                                                        └─> Phase 15: Tooling & Polish
```

## Quick Start

To begin using these templates:

1. **Review**: Read `docs/issues/README.md` for overview
2. **Prepare**: Install GitHub CLI if using automated method
3. **Create**: Run `create-all-issues.sh` or create manually
4. **Organize**: Create milestones for each phase
5. **Track**: Set up a project board for progress tracking
6. **Start**: Begin with Phase 2 implementation

## Related Documentation

- [Phased Roadmap](docs/planning/PHASED_ROADMAP.md) - High-level overview
- [AI Implementation Plan](docs/planning/AI_IMPLEMENTATION_PLAN.md) - Detailed guide
- [Project Status](docs/PROJECT_STATUS.md) - Current state
- [Contributing Guide](CONTRIBUTING.md) - How to contribute

## Benefits of These Templates

✅ **Comprehensive**: Every detail needed for implementation  
✅ **Actionable**: Clear tasks with time estimates  
✅ **Testable**: Specific test requirements for each feature  
✅ **Documented**: Documentation updates as first-class tasks  
✅ **Example-Driven**: Real code examples for each feature  
✅ **Trackable**: Checkboxes for progress tracking  
✅ **Consistent**: Same structure across all phases  
✅ **AI-Friendly**: Detailed enough for AI-assisted development  

## Contributing

To improve these templates:

1. Implement a phase and note what could be clearer
2. Update the template with improvements
3. Submit a PR with your enhancements
4. Help future implementers

## Questions?

- Check [HOWTO.md](docs/issues/HOWTO.md) for usage instructions
- Review [AI Implementation Plan](docs/planning/AI_IMPLEMENTATION_PLAN.md) for details
- Open a [GitHub Discussion](https://github.com/kristofer/Unified/discussions) for questions

---

**Created**: January 2026  
**Version**: 1.0  
**Total Templates**: 14 (Phases 2-15)  
**Total Lines**: ~8,000  
**Estimated Total Duration**: 35-48 weeks
