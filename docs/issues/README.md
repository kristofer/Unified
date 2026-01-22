# GitHub Issues for Unified Compiler Phases

This directory contains comprehensive issue templates for each phase of the Unified compiler implementation (Phases 2-15).

## Overview

Each phase has a dedicated issue template that includes:
- **Goals and Prerequisites**: What the phase achieves and what must be completed first
- **Implementation Tasks**: Detailed, actionable tasks with time estimates (50-70 tasks per phase)
- **Test Requirements**: Comprehensive test coverage requirements (40+ test cases per phase)
- **Documentation Updates**: Required documentation changes
- **Example Programs**: 5-8 example programs demonstrating new features
- **Success Criteria**: Measurable completion criteria
- **Implementation Notes**: Common pitfalls, debugging tips, and best practices

## Phase Templates

| Phase | File | Duration | Focus Area |
|-------|------|----------|------------|
| Phase 2 | [phase-02-control-flow.md](phase-02-control-flow.md) | 1-2 weeks | Control flow statements (if/else, while, for, loop) |
| Phase 3 | [phase-03-variables-mutability.md](phase-03-variables-mutability.md) | 2 weeks | Mutability, assignment, symbol tables, type inference |
| Phase 4 | [phase-04-functions-advanced.md](phase-04-functions-advanced.md) | 2-3 weeks | Tuples, lambdas, default parameters, block expressions |
| Phase 5 | [phase-05-structs-types.md](phase-05-structs-types.md) | 3-4 weeks | Struct declarations, methods, memory layout |
| Phase 6 | [phase-06-enums-pattern-matching.md](phase-06-enums-pattern-matching.md) | 3-4 weeks | Enums, pattern matching, exhaustiveness checking |
| Phase 7 | [phase-07-error-handling.md](phase-07-error-handling.md) | 1-2 weeks | ? operator, error propagation |
| Phase 8 | [phase-08-arrays-slices.md](phase-08-arrays-slices.md) | 2-3 weeks | Fixed arrays, slices, bounds checking |
| Phase 9 | [phase-09-string-operations.md](phase-09-string-operations.md) | 2 weeks | String methods, interpolation, UTF-8 |
| Phase 10 | [phase-10-generics.md](phase-10-generics.md) | 3-4 weeks | Type parameters, monomorphization, type inference |
| Phase 11 | [phase-11-modules-imports.md](phase-11-modules-imports.md) | 2-3 weeks | Module system, imports/exports, visibility |
| Phase 12 | [phase-12-basic-ownership.md](phase-12-basic-ownership.md) | 3-4 weeks | Ownership tracking, move semantics, borrow checker |
| Phase 13 | [phase-13-stdlib-foundation.md](phase-13-stdlib-foundation.md) | 2-3 weeks | Core types, collections, I/O primitives |
| Phase 14 | [phase-14-concurrency-basics.md](phase-14-concurrency-basics.md) | 4-6 weeks | Async/await, actors, channels |
| Phase 15 | [phase-15-tooling-polish.md](phase-15-tooling-polish.md) | 3-4 weeks | Error messages, formatter, docs, package manager |

**Total Estimated Duration**: 35-48 weeks (8-11 months)

## How to Create GitHub Issues

### Important Note
These templates are ready to be converted into GitHub issues. Since automated issue creation requires specific GitHub permissions, you'll need to create the issues manually or use the GitHub CLI.

### Method 1: Manual Creation (Recommended)

For each phase:

1. **Navigate to GitHub Issues**:
   - Go to https://github.com/kristofer/Unified/issues/new

2. **Create New Issue**:
   - Click "New Issue"
   - Copy the entire content from the phase template file
   - Paste into the issue description

3. **Set Issue Metadata**:
   - **Title**: Use the phase name (e.g., "Phase 2: Control Flow and Parser Grammar Updates")
   - **Labels**: Add labels from the bottom of each template (e.g., `phase-2`, `enhancement`, `parser`, `vm`)
   - **Milestone**: Create/select the appropriate phase milestone
   - **Assignee**: Assign to yourself or leave unassigned

4. **Create Issue**:
   - Click "Submit new issue"

### Method 2: GitHub CLI (Automated)

If you have the GitHub CLI (`gh`) installed and configured:

```bash
# Navigate to the repository
cd /home/runner/work/Unified/Unified

# Create issues for all phases
./docs/issues/create-all-issues.sh
```

Or create individual issues:

```bash
# Phase 2 example
gh issue create \
  --title "Phase 2: Control Flow and Parser Grammar Updates" \
  --body-file docs/issues/phase-02-control-flow.md \
  --label "phase-2,enhancement,parser,vm,control-flow" \
  --milestone "Phase 2"

# Phase 3 example
gh issue create \
  --title "Phase 3: Variables, Mutability, and Assignment" \
  --body-file docs/issues/phase-03-variables-mutability.md \
  --label "phase-3,enhancement,parser,vm,semantic-analysis" \
  --milestone "Phase 3"
```

### Method 3: GitHub API (Programmatic)

Use the GitHub REST API with a personal access token:

```bash
# Set your GitHub token
export GITHUB_TOKEN="your_github_token_here"

# Create issues using the API
./docs/issues/create-issues-api.sh
```

## Issue Creation Script

A helper script `create-all-issues.sh` is provided to create all phase issues at once using the GitHub CLI.

### Prerequisites
- GitHub CLI installed: https://cli.github.com/
- Authenticated with GitHub: `gh auth login`

### Usage

```bash
cd /home/runner/work/Unified/Unified
chmod +x docs/issues/create-all-issues.sh
./docs/issues/create-all-issues.sh
```

This will create 14 issues (one for each phase 2-15) with proper titles, labels, and milestones.

## Tracking Progress

Once issues are created, you can track progress by:

1. **Checking off tasks**: Mark tasks complete in the issue as you work
2. **Adding comments**: Document decisions, blockers, and progress
3. **Updating labels**: Add `in-progress`, `blocked`, or `completed` labels
4. **Linking PRs**: Link pull requests that implement parts of a phase
5. **Closing issues**: Close when all success criteria are met

## Issue Labels

Each phase uses consistent labels:

- `phase-N`: Identifies the phase number (e.g., `phase-2`, `phase-3`)
- `enhancement`: Indicates a new feature
- Component labels:
  - `parser`: Parser/grammar changes
  - `vm`: Virtual machine changes
  - `bytecode`: Bytecode generation changes
  - `semantic-analysis`: Type checking and analysis
  - `stdlib`: Standard library additions
  - `tooling`: Developer tools
  - `documentation`: Documentation updates

Priority labels:
- `priority-high`: Critical for next phase
- `priority-medium`: Important but not blocking
- `priority-low`: Nice to have

Status labels (add as needed):
- `in-progress`: Currently being worked on
- `blocked`: Blocked by another issue or dependency
- `needs-review`: Implementation complete, needs review
- `completed`: All success criteria met

## Milestones

Create GitHub milestones for each phase:

1. **Phase 2 - Control Flow** (1-2 weeks)
2. **Phase 3 - Variables & Mutability** (2 weeks)
3. **Phase 4 - Functions Advanced** (2-3 weeks)
4. **Phase 5 - Structs & Types** (3-4 weeks)
5. **Phase 6 - Enums & Pattern Matching** (3-4 weeks)
6. **Phase 7 - Error Handling** (1-2 weeks)
7. **Phase 8 - Arrays & Slices** (2-3 weeks)
8. **Phase 9 - String Operations** (2 weeks)
9. **Phase 10 - Generics** (3-4 weeks)
10. **Phase 11 - Modules & Imports** (2-3 weeks)
11. **Phase 12 - Basic Ownership** (3-4 weeks)
12. **Phase 13 - Standard Library** (2-3 weeks)
13. **Phase 14 - Concurrency** (4-6 weeks)
14. **Phase 15 - Tooling & Polish** (3-4 weeks)

## Project Board

Consider creating a GitHub Project board to track all phases:

1. **Columns**:
   - Not Started
   - In Progress
   - In Review
   - Completed

2. **Views**:
   - By Phase (group by milestone)
   - By Priority (group by priority label)
   - Timeline (Gantt chart view)

## Related Documentation

- [Phased Roadmap](../planning/PHASED_ROADMAP.md) - High-level phase overview
- [AI Implementation Plan](../planning/AI_IMPLEMENTATION_PLAN.md) - Detailed implementation guide
- [Project Status](../PROJECT_STATUS.md) - Current implementation status
- [Contributing Guide](../../CONTRIBUTING.md) - How to contribute

## Questions?

If you have questions about any phase or need clarification on implementation details:

1. Check the [AI Implementation Plan](../planning/AI_IMPLEMENTATION_PLAN.md)
2. Review the [Language Specification](../../spec/UnifiedSpecifiation.md)
3. Look at reference implementation: [Smog Language](https://github.com/kristofer/smog)
4. Open a GitHub Discussion: https://github.com/kristofer/Unified/discussions

## Updates

This directory is maintained as part of the Unified compiler project. Templates are updated as implementation progresses and new insights are gained.

**Last Updated**: January 2026  
**Template Version**: 1.0  
**Total Templates**: 14 (Phases 2-15)
