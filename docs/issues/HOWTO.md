# How to Create GitHub Issues from Templates

This guide explains how to create GitHub issues from the phase templates since automated issue creation requires specific GitHub API access.

## Important Context

This implementation creates **issue templates** rather than directly creating GitHub issues because:

1. **GitHub API Limitations**: Direct issue creation requires GitHub API access with a personal access token
2. **Permission Requirements**: Creating issues programmatically requires specific repository permissions
3. **Better Control**: Manual or CLI-based creation gives you control over when and how issues are created

## What Has Been Created

✅ **14 Comprehensive Issue Templates** (Phases 2-15)
- Each template is 400-600 lines of detailed specifications
- Includes 50-70 implementation tasks per phase
- Contains 40+ test requirements per phase
- Features 5-8 example programs each
- Total: ~6,500 lines of documentation

✅ **Helper Scripts and Documentation**
- `README.md` - Complete guide to the templates
- `create-all-issues.sh` - Automated issue creation script using GitHub CLI
- Templates are ready to copy/paste into GitHub

## Three Ways to Create Issues

### Option 1: Manual Creation (Simplest)

**Best for**: First-time setup, or creating a few issues

1. **Open GitHub Issues**: Navigate to https://github.com/kristofer/Unified/issues/new

2. **For each phase template**:
   - Open the template file (e.g., `docs/issues/phase-02-control-flow.md`)
   - Copy the entire file contents
   - Paste into the GitHub issue description field
   - Set the title (e.g., "Phase 2: Control Flow and Parser Grammar Updates")
   - Add labels from the bottom of the template
   - Create the issue

3. **Repeat for all 14 phases** (2-15)

**Time required**: ~15-20 minutes for all phases

### Option 2: GitHub CLI (Recommended)

**Best for**: Automated creation of all issues at once

**Prerequisites**:
1. Install GitHub CLI: https://cli.github.com/
2. Authenticate: `gh auth login`

**Steps**:

```bash
# Navigate to repository
cd /home/runner/work/Unified/Unified

# Run the helper script
./docs/issues/create-all-issues.sh
```

This will:
- Create all 14 issues (Phases 2-15)
- Set proper titles
- Apply correct labels
- Use template content as issue body

**Time required**: ~1 minute (automated)

### Option 3: Individual CLI Commands

**Best for**: Creating specific phases or having fine control

```bash
# Phase 2
gh issue create \
  --title "Phase 2: Control Flow and Parser Grammar Updates" \
  --body-file docs/issues/phase-02-control-flow.md \
  --label "phase-2,enhancement,parser,vm,control-flow"

# Phase 3
gh issue create \
  --title "Phase 3: Variables, Mutability, and Assignment" \
  --body-file docs/issues/phase-03-variables-mutability.md \
  --label "phase-3,enhancement,parser,vm,semantic-analysis"

# ... repeat for other phases
```

## After Creating Issues

### 1. Create Milestones

Create a milestone for each phase to track progress:

1. Go to https://github.com/kristofer/Unified/milestones
2. Click "New milestone"
3. Create milestones:
   - "Phase 2 - Control Flow" (due: ~2 weeks from start)
   - "Phase 3 - Variables & Mutability" (due: ~2 weeks from Phase 2)
   - etc.

### 2. Assign Milestones to Issues

For each issue:
1. Click the issue
2. Click the gear icon next to "Milestone"
3. Select the appropriate phase milestone

### 3. Set Up Project Board (Optional)

Create a project board to visualize progress:

1. Go to https://github.com/kristofer/Unified/projects
2. Create "Unified Compiler Development" board
3. Add columns:
   - **Not Started**: Phases not yet begun
   - **In Progress**: Currently being worked on
   - **In Review**: Implementation complete, needs review
   - **Completed**: All success criteria met
4. Add all phase issues to the board

### 4. Track Progress

As you work on each phase:

1. **Update checkboxes**: Mark tasks complete using `- [x]` syntax
2. **Add comments**: Document decisions, blockers, and progress notes
3. **Link PRs**: Reference pull requests that implement features
4. **Update labels**: Add `in-progress`, `blocked`, or `completed` labels
5. **Close when done**: Close issue when all success criteria are met

## Template Structure

Each template follows this structure:

```markdown
# Phase N: Title

**Status**: Not Started
**Duration**: X weeks
**Priority**: HIGH/MEDIUM/LOW
**Dependencies**: Previous phases

## Goals
What this phase achieves

## Prerequisites
What must be complete before starting

## Language Features to Add
List of new features

## Implementation Tasks
- [ ] Task 1 (time estimate)
- [ ] Task 2 (time estimate)
... 50-70 tasks total

## Test Requirements
- [ ] Test category 1
- [ ] Test category 2
... 40+ test cases

## Documentation Updates Required
- [ ] Doc update 1
- [ ] Doc update 2

## Example Programs
```code examples```

## Success Criteria
- [ ] Criterion 1
- [ ] Criterion 2

## Implementation Notes
Tips, pitfalls, debugging advice

---
Labels: phase-N, enhancement, component-tags
```

## Issue Labels Reference

Apply these labels to issues:

**Phase Labels** (required):
- `phase-2` through `phase-15`

**Component Labels**:
- `parser` - Grammar/parsing changes
- `vm` - Virtual machine changes
- `bytecode` - Bytecode generation
- `semantic-analysis` - Type checking
- `stdlib` - Standard library
- `tooling` - Developer tools
- `documentation` - Docs only

**Priority Labels**:
- `priority-high` - Critical
- `priority-medium` - Important
- `priority-low` - Nice to have

**Status Labels** (add as work progresses):
- `in-progress` - Currently working
- `blocked` - Waiting on dependency
- `needs-review` - Ready for review
- `completed` - Done

## Troubleshooting

### "gh: command not found"

**Solution**: Install GitHub CLI from https://cli.github.com/

### "gh: authentication required"

**Solution**: Run `gh auth login` and follow the prompts

### "Permission denied"

**Solution**: 
- Ensure you have write access to the repository
- Check that your GitHub token has the `repo` scope
- Try refreshing authentication: `gh auth refresh`

### Script fails with "file not found"

**Solution**: 
- Ensure you're running the script from the repository root
- Check that all template files exist in `docs/issues/`
- Verify file naming matches the pattern: `phase-XX-*.md`

## Next Steps

After creating all issues:

1. ✅ Review each issue to ensure content is accurate
2. ✅ Create and assign milestones
3. ✅ Set up project board for tracking
4. ✅ Prioritize which phase to start with (likely Phase 2)
5. ✅ Begin implementation following the task checklist
6. ✅ Update issue with progress as you work

## Questions?

- Check the [AI Implementation Plan](../planning/AI_IMPLEMENTATION_PLAN.md) for detailed guidance
- Review [Project Status](../PROJECT_STATUS.md) for current state
- See [Contributing Guide](../../CONTRIBUTING.md) for workflow
- Open a [GitHub Discussion](https://github.com/kristofer/Unified/discussions) for questions

---

**Remember**: These templates are living documents. Update them as you learn and discover better approaches during implementation!
