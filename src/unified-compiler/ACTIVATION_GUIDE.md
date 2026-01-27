# New Keyword Implementation - ACTIVATION GUIDE

## Status: Ready for Parser Regeneration

All code has been implemented and is ready to be activated once the parser is regenerated with ANTLR.

## What Has Been Implemented

### ✅ Grammar Files Updated
- `grammar/UnifiedLexer.g4` - Added `NEW : 'new';` keyword
- `grammar/UnifiedParser.g4` - Added `constructorExpr` rule

### ✅ AST Node Created
- `internal/ast/ast.go` - `NewExpr` struct with full documentation

### ✅ Visitor Implementation (Ready to Activate)
- `internal/ast/visitor.go` - `VisitConstructorExpr()` method implemented
- Code is commented out but ready to be uncommented after parser regeneration
- See lines 814 and 845-882

### ✅ Bytecode Generator Implemented
- `internal/bytecode/generator.go` - `generateNewExpr()` method fully implemented
- Handles positional constructor arguments
- Maps arguments to struct fields in declaration order
- Validates argument count matches field count

### ✅ Test Files Created
- `test/new_keyword_basic.uni` - Basic test cases
- `test/integration/new_keyword.uni` - Comprehensive integration tests

### ✅ Documentation Updated
- All 62 occurrences of old syntax replaced with new syntax
- Implementation guides created

## Quick Activation Steps (After Parser Regeneration)

Once you have regenerated the parser using ANTLR, follow these steps to activate the new keyword:

### Step 1: Regenerate Parser

```bash
cd src/unified-compiler
make parser
```

This will regenerate the parser with the new `constructorExpr` rule and create:
- `parser.ConstructorExprContext` type
- `ctx.ConstructorExpr()` method on PrimaryContext
- `IConstructorExprContext` interface

### Step 2: Uncomment Visitor Code

In `internal/ast/visitor.go`:

**Line ~814** - Uncomment these lines:
```go
// Currently commented:
// if constructorExprCtx := ctx.ConstructorExpr(); constructorExprCtx != nil {
// 	return v.VisitConstructorExpr(constructorExprCtx.(*parser.ConstructorExprContext))
// }

// Uncomment to:
if constructorExprCtx := ctx.ConstructorExpr(); constructorExprCtx != nil {
	return v.VisitConstructorExpr(constructorExprCtx.(*parser.ConstructorExprContext))
}
```

**Line ~846** - Update the VisitConstructorExpr function:
```go
// Currently:
func (v *ASTBuilder) VisitConstructorExpr(ctx *parser.ConstructorExprContext) interface{} {
	// Placeholder implementation

// Change to:
func (v *ASTBuilder) VisitNewExpr(ctx *parser.NewExprContext) interface{} {
```

Then uncomment the implementation inside (lines ~854-879) and remove the placeholder `return nil`.

### Step 3: Build and Test

```bash
# Build the compiler
make build

# Run unit tests
go test ./internal/bytecode/...

# Test with new keyword syntax
./bin/unified --input test/new_keyword_basic.uni
./bin/unified --input test/integration/new_keyword.uni

# Run all integration tests
make run-all-tests
```

### Step 4: Verify

- [ ] Compiler builds without errors
- [ ] All unit tests pass
- [ ] `new Type()` syntax compiles successfully
- [ ] `new Type<T>()` generic syntax works
- [ ] `new Type(arg1, arg2)` with arguments works
- [ ] Old `Type::new()` syntax produces error (as expected)

## Implementation Details

### How It Works

The `new` keyword is syntactic sugar that creates struct instances using positional constructor arguments instead of named field initializers.

**Traditional struct instantiation:**
```unified
let point = Point{ x: 10, y: 20 }
```

**New keyword syntax:**
```unified
let point = new Point(10, 20)
```

Both compile to the same bytecode using `OpAllocStruct`.

### Constructor Argument Mapping

Arguments are mapped to struct fields in declaration order:

```unified
struct Point {
    x: Int  // First field -> first argument
    y: Int  // Second field -> second argument
}

let p = new Point(10, 20)  // 10 -> x, 20 -> y
```

### Validation

The bytecode generator validates:
- Struct type exists
- Argument count matches field count
- All arguments can be evaluated to expressions

### Error Messages

```
new UnknownType(): "undefined struct type: UnknownType"
new Point(): "new Point: constructor requires 2 arguments"
new Point(1, 2, 3): "new Point: expected 2 arguments, got 3"
```

## Files Modified

### Source Code
- `internal/ast/visitor.go` - Lines 812-882 (visitor implementation, commented)
- `internal/bytecode/generator.go` - Lines 357, 1019-1079 (fully implemented)

### Tests
- `test/new_keyword_basic.uni` - Basic test cases
- `test/integration/new_keyword.uni` - Comprehensive tests

## Known Limitations

1. **No default values**: All struct fields must be provided as arguments
2. **Positional only**: Arguments must be in field declaration order
3. **No named arguments**: Cannot use `new Point(y: 20, x: 10)`
4. **No varargs**: Constructor must match field count exactly

These limitations can be addressed in future enhancements.

## Troubleshooting

### Build Error: "ctx.NewExpr undefined"
**Cause**: Parser not regenerated yet  
**Solution**: Run `make parser` first

### Build Error: "NewExprContext is not a type"
**Cause**: Parser not regenerated yet  
**Solution**: Run `make parser` first

### Runtime Error: "undefined struct type"
**Cause**: Trying to instantiate undefined struct  
**Solution**: Declare the struct before using it with `new`

### Parser Regeneration Fails
**Cause**: ANTLR not installed or network issues  
**Solution**: See `PARSER_REGENERATION_NEEDED.md` for installation options

## Next Steps

1. Install ANTLR (see `PARSER_REGENERATION_NEEDED.md`)
2. Regenerate parser: `make parser`
3. Uncomment visitor code (this file, Step 2)
4. Build and test: `make build && make test`
5. Celebrate! 🎉

## Questions?

See the comprehensive guides:
- `NEW_KEYWORD_IMPLEMENTATION_GUIDE.md` - Full implementation details
- `PARSER_REGENERATION_NEEDED.md` - Parser regeneration options
- `NEW_KEYWORD_SUMMARY.md` - Complete change summary
