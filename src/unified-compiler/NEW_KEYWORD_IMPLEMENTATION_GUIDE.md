# New Keyword Implementation Guide

## Overview

This guide provides step-by-step instructions for completing the `new` keyword implementation after the parser has been regenerated from the updated grammar files.

## Prerequisites

✅ **Completed Steps:**
1. Grammar files updated (UnifiedLexer.g4, UnifiedParser.g4)
2. AST node created (`NewExpr` in ast.go)
3. All documentation updated to use `new Type()` syntax
4. Test file created (test/new_keyword_basic.uni)

⚠️ **Pending Steps:**
1. Parser regeneration (see PARSER_REGENERATION_NEEDED.md)
2. Visitor implementation
3. Bytecode generation
4. Testing and validation

## Step 1: Regenerate the Parser

Follow instructions in `PARSER_REGENERATION_NEEDED.md` to regenerate the parser from grammar files.

After regeneration, the following new methods will be available:
- `VisitNewExpr(ctx *NewExprContext) interface{}` in the visitor interface
- `NewExpr()` method on `PrimaryContext`
- `INewExprContext` interface

## Step 2: Update AST Visitor

File: `internal/ast/visitor.go`

Add the following method after `VisitAsyncExpr`:

```go
// VisitNewExpr builds a NewExpr node
func (v *ASTBuilder) VisitNewExpr(ctx *parser.NewExprContext) interface{} {
	// Get type name
	typeName := ctx.Identifier().GetText()
	
	// Process generic type arguments if present
	var typeArgs []Type
	if ctx.LT() != nil && ctx.GT() != nil {
		if typeListCtx := ctx.TypeList(); typeListCtx != nil {
			typeArgs = v.processTypeList(typeListCtx.(*parser.TypeListContext))
		}
	}
	
	// Process constructor arguments if present
	var args []Expression
	if argListCtx := ctx.ArgList(); argListCtx != nil {
		args = v.processArgList(argListCtx.(*parser.ArgListContext))
	}
	
	return &NewExpr{
		TypeName: typeName,
		TypeArgs: typeArgs,
		Args:     args,
		Position: v.getPosition(ctx),
	}
}
```

Update `VisitPrimary` to handle new expressions (add after line ~810):

```go
if newExprCtx := ctx.NewExpr(); newExprCtx != nil {
	return v.VisitNewExpr(newExprCtx.(*parser.NewExprContext))
}
```

## Step 3: Update Bytecode Generator

File: `internal/bytecode/generator.go`

Find the `generateExpression` function and add a case for `*ast.NewExpr`:

```go
func (g *Generator) generateExpression(expr ast.Expression) error {
	switch e := expr.(type) {
	// ... existing cases ...
	
	case *ast.NewExpr:
		return g.generateNewExpr(e)
	
	// ... rest of cases ...
	}
	return fmt.Errorf("unsupported expression type: %T", expr)
}
```

Add the new expression generator method:

```go
// generateNewExpr generates bytecode for new expressions
func (g *Generator) generateNewExpr(expr *ast.NewExpr) error {
	// For now, new Type() is semantically equivalent to Type{} (struct literal)
	// This is a simplified implementation that treats new as struct initialization
	// Future enhancements can add constructor support
	
	// Generate arguments
	for _, arg := range expr.Args {
		if err := g.generateExpression(arg); err != nil {
			return err
		}
	}
	
	// Generate CALL instruction for constructor
	// The actual implementation depends on how constructors are represented
	// For basic structs without explicit constructors, this would be a STRUCT_NEW instruction
	
	g.emit(OpStructNew)
	g.emitString(expr.TypeName)
	g.emitInt(len(expr.Args))
	
	return nil
}
```

## Step 4: Add VM Support (if needed)

File: `internal/vm/vm.go`

If a new `OpStructNew` instruction is needed, add it to the instruction set:

File: `internal/bytecode/instructions.go`

```go
const (
	// ... existing opcodes ...
	OpStructNew    Opcode = 0x?? // Replace ?? with next available opcode number
)
```

Update the VM execution loop to handle the new instruction:

File: `internal/vm/vm.go`

```go
case OpStructNew:
	// Pop constructor arguments from stack
	// Create new struct instance
	// Push struct onto stack
	// Implementation depends on VM structure
```

## Step 5: Create Comprehensive Tests

File: `internal/bytecode/new_expr_test.go`

```go
package bytecode

import (
	"testing"
	"unified-compiler/internal/ast"
)

func TestNewExprBasic(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:  "simple new without args",
			input: "new Stack<Int>()",
			expected: "basic stack instantiation",
		},
		{
			name:  "new with args",
			input: "new Point(10, 20)",
			expected: "point with coordinates",
		},
		{
			name:  "new with generic type",
			input: "new List<String>()",
			expected: "generic list instantiation",
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Test implementation
		})
	}
}
```

File: `test/integration/new_keyword.uni`

```unified
// Comprehensive test for new keyword
fn test_basic_new() {
	let stack = new Stack<Int>()
	return 0
}

fn test_new_with_args() {
	let point = new Point(10, 20)
	return point.x + point.y
}

fn test_generic_new() {
	let list = new List<String>()
	return list.len()
}

struct Stack<T> {
	items: List<T>
}

struct Point {
	x: Int
	y: Int
}

struct List<T> {
	data: Array<T>
}
```

## Step 6: Build and Test

```bash
cd src/unified-compiler

# Build the compiler
make build

# Run unit tests
make test

# Run integration tests
./bin/unified --input test/integration/new_keyword.uni
./bin/unified --input test/new_keyword_basic.uni

# Run all tests
make run-all-tests
```

## Step 7: Verification Checklist

- [ ] Parser regenerated successfully
- [ ] No compilation errors in Go code
- [ ] `make build` succeeds
- [ ] Unit tests pass
- [ ] Integration tests pass
- [ ] Example programs using `new` keyword compile and run
- [ ] Documentation is accurate
- [ ] No old `Type::new()` or `Type.new()` syntax remains in codebase

## Expected Behavior

After implementation, the following should work:

```unified
// Before (old syntax - NO LONGER SUPPORTED)
// let stack = Stack<Int>::new()

// After (new syntax - CURRENT)
let stack = new Stack<Int>()

// With arguments
let point = new Point(10, 20)

// With generic types
let list = new List<String>()
```

## Troubleshooting

### Parser regeneration fails
- Ensure ANTLR 4.13.x is installed
- Check Java version (Java 11+ required)
- Verify grammar files have no syntax errors

### Compilation errors after parser regeneration
- Clear old generated files: `make clean`
- Regenerate: `make parser`
- Rebuild: `make build`

### Runtime errors with new keyword
- Check bytecode generation for NewExpr
- Verify VM instruction implementation
- Add debug output to trace execution

## Notes

- The `new` keyword is now a reserved keyword and cannot be used as an identifier
- Old syntax `Type::new()` and `Type.new()` is deprecated and will not compile
- Generic type parameters must be specified at instantiation time
- Constructor arguments are passed in parentheses after the type name
