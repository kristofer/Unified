# `new` Keyword Implementation - Summary

## Overview

This PR implements support for the `new` keyword in the Unified programming language, replacing the old `Type::new()` and `Type.new()` syntax with the more familiar `new Type()` syntax, similar to Java, C++, and JavaScript.

## Motivation

The old syntax `Stack<Int>::new()` required:
- Implementation blocks (impl)
- Type resolution complexity
- Less intuitive syntax for newcomers

The new syntax `new Stack<Int>()` provides:
- Clearer intent
- Simpler parsing
- More familiar syntax for developers from other languages

## Changes Made

### 1. Grammar Updates ✅

**File: `src/unified-compiler/grammar/UnifiedLexer.g4`**
- Added `NEW : 'new';` keyword (line 105)

**File: `src/unified-compiler/grammar/UnifiedParser.g4`**
- Added `newExpr` rule: `NEW identifier (LT typeList GT)? LPAREN argList? RPAREN`
- Updated `primary` rule to include `newExpr` as first option

### 2. AST Updates ✅

**File: `src/unified-compiler/internal/ast/ast.go`**
- Added `NewExpr` struct (lines 500-517):
  ```go
  type NewExpr struct {
      TypeName string
      TypeArgs []Type       // Type arguments for generic types
      Args     []Expression // Constructor arguments
      Position Position
  }
  ```

### 3. Documentation Updates ✅

All occurrences of the old constructor syntax have been updated to use the new `new` keyword:

**Specification:**
- `spec/UnifiedSpecifiation.md` - 9 replacements, keywords updated

**Documentation:**
- `docs/issues/phase-05-structs-types.md` - 5 replacements
- `docs/issues/phase-10-generics.md` - 5 replacements
- `docs/issues/phase-11-modules-imports.md` - 2 replacements
- `docs/issues/phase-12-basic-ownership.md` - 3 replacements
- `docs/issues/phase-13-stdlib-foundation.md` - 4 replacements
- `docs/issues/phase-14-concurrency-basics.md` - 8 replacements
- `docs/issues/phase-15-tooling-polish.md` - 1 replacement

**Library Code:**
- `lib/List.uni` - 3 replacements
- `lib/Tree.uni` - 12 replacements

**Examples:**
- `ClassicExercises.md` - 9 replacements
- `PHASES_2_9_STATUS.md` - 1 replacement

**Total: 62 replacements across 15 files**

### 4. Test Files ✅

**File: `src/unified-compiler/test/new_keyword_basic.uni`**
- Created test file demonstrating new keyword usage with:
  - Basic new expression
  - New with arguments
  - New with generic type parameters

### 5. Implementation Guides ✅

**File: `src/unified-compiler/PARSER_REGENERATION_NEEDED.md`**
- Detailed instructions for regenerating the parser with ANTLR
- Multiple installation options (native, Python, Docker)
- Step-by-step guide for different platforms

**File: `src/unified-compiler/NEW_KEYWORD_IMPLEMENTATION_GUIDE.md`**
- Complete implementation guide for remaining steps
- Code examples for visitor implementation
- Bytecode generation guidelines
- VM instruction updates
- Test creation templates
- Troubleshooting section

## Syntax Examples

### Before (Old Syntax - Deprecated)
```unified
let channel = Channel<String>.new()
let stack = Stack<Int>::new()
let point = Point.new(10, 20)
```

### After (New Syntax - Current)
```unified
let channel = new Channel<String>()
let stack = new Stack<Int>()
let point = new Point(10, 20)
```

## Current Status

### Completed ✅
1. Grammar definition updated
2. AST node created
3. All documentation updated
4. Test files created
5. Implementation guides written
6. Compiler still builds successfully
7. Existing tests still pass

### Pending ⚠️
1. **Parser Regeneration** - Requires ANTLR installation
   - Network restrictions prevented automatic installation
   - Manual regeneration required (see PARSER_REGENERATION_NEEDED.md)

2. **Visitor Implementation** - After parser regeneration
   - Add `VisitNewExpr()` method to visitor
   - Update `VisitPrimary()` to handle new expressions

3. **Bytecode Generation** - After visitor implementation
   - Implement `generateNewExpr()` method
   - Add `OpStructNew` instruction if needed
   - Update VM to handle new instruction

4. **Testing** - After implementation
   - Create comprehensive unit tests
   - Create integration tests
   - Verify backward compatibility

## Impact Assessment

### Breaking Changes
- ❌ Old syntax `Type::new()` will no longer be supported
- ❌ Old syntax `Type.new()` will no longer be supported

### Migration Required
- All user code using the old syntax must be updated
- Automated migration tool could be created to convert old syntax

### Backward Compatibility
- None - this is a breaking change by design
- All existing library code has been updated
- All documentation has been updated

## Next Steps

1. **Immediate**: Regenerate parser using ANTLR
   ```bash
   cd src/unified-compiler
   make parser
   ```

2. **Implementation**: Follow NEW_KEYWORD_IMPLEMENTATION_GUIDE.md
   - Update visitor
   - Update bytecode generator
   - Update VM (if needed)

3. **Testing**: Comprehensive test suite
   - Unit tests for AST building
   - Bytecode generation tests
   - Integration tests
   - Example programs

4. **Validation**: Ensure all tests pass
   ```bash
   make test
   make run-all-tests
   ```

5. **Review**: Code review and security scan
   - Code review for changes
   - CodeQL security scan

## Files Modified

### Grammar
- `src/unified-compiler/grammar/UnifiedLexer.g4`
- `src/unified-compiler/grammar/UnifiedParser.g4`

### Source Code
- `src/unified-compiler/internal/ast/ast.go`

### Documentation
- `spec/UnifiedSpecifiation.md`
- `docs/issues/phase-05-structs-types.md`
- `docs/issues/phase-10-generics.md`
- `docs/issues/phase-11-modules-imports.md`
- `docs/issues/phase-12-basic-ownership.md`
- `docs/issues/phase-13-stdlib-foundation.md`
- `docs/issues/phase-14-concurrency-basics.md`
- `docs/issues/phase-15-tooling-polish.md`
- `ClassicExercises.md`
- `PHASES_2_9_STATUS.md`

### Library
- `lib/List.uni`
- `lib/Tree.uni`

### Tests
- `src/unified-compiler/test/new_keyword_basic.uni` (new)

### Guides
- `src/unified-compiler/PARSER_REGENERATION_NEEDED.md` (new)
- `src/unified-compiler/NEW_KEYWORD_IMPLEMENTATION_GUIDE.md` (new)

## Build Status

✅ Compiler builds successfully
✅ Existing tests still pass
⚠️ Parser regeneration pending

## Security Considerations

- No security vulnerabilities introduced by grammar changes
- Parser regeneration should be done from trusted ANTLR source
- CodeQL scan pending for final implementation

## Notes

This is a **BIG** change as indicated in the issue. The implementation has been carefully planned and documented, with all user-facing documentation already updated to reflect the new syntax. The remaining work is primarily mechanical (parser regeneration and visitor/bytecode implementation) and well-documented in the implementation guide.
