# Semicolons in Unified

## Summary

Semicolons are **optional** in the Unified programming language. The language follows a philosophy similar to Go, where semicolons are not required at the end of statements within blocks.

## Implementation

The Unified compiler uses ANTLR's error recovery mechanism combined with a custom error listener to make semicolons optional:

1. The grammar defines semicolons in statement rules (for ANTLR compatibility)
2. A custom error listener (`CustomErrorListener`) suppresses "missing ';'" error messages
3. ANTLR's automatic error recovery inserts semicolons where needed during parsing

This approach ensures backward compatibility (code with semicolons still works) while allowing cleaner code without them.

## Examples

### Without Semicolons (Recommended)

```unified
fn fibonacci(n: Int) -> Int {
    if n <= 1 {
        return n
    }
    
    let mut a = 0
    let mut b = 1
    let mut result = 0
    
    for i in 2..=n {
        result = a + b
        a = b
        b = result
    }
    
    return result
}

fn main() -> Int {
    let fib10 = fibonacci(10)
    return fib10
}
```

### With Semicolons (Also Valid)

```unified
fn fibonacci(n: Int) -> Int {
    if n <= 1 {
        return n;
    }
    
    let mut a = 0;
    let mut b = 1;
    let mut result = 0;
    
    for i in 2..=n {
        result = a + b;
        a = b;
        b = result;
    }
    
    return result;
}

fn main() -> Int {
    let fib10 = fibonacci(10);
    return fib10;
}
```

### Mixed (Also Valid)

```unified
fn calculate() -> Int {
    let a = 5;      // with semicolon
    let b = 10      // without semicolon
    let c = 15;     // with semicolon  
    let d = 20      // without semicolon
    return a + b + c + d
}
```

## When Semicolons Are Not Needed

Semicolons are not required after:

- Variable declarations (`let`, `var`)
- Assignments and compound assignments (`=`, `+=`, `-=`, etc.)
- Return statements
- Expression statements
- Break and continue statements
- Top-level declarations (imports, constants, type aliases)

## When Semicolons May Be Used

You can still use semicolons if you prefer, or to write multiple statements on one line:

```unified
fn inline() -> Int {
    let x = 1; let y = 2; return x + y
}
```

## Best Practices

1. **Prefer no semicolons**: Write clean, readable code without semicolons
2. **Be consistent**: If you choose to use semicolons, use them consistently throughout your codebase
3. **Multiple statements per line**: Use semicolons when putting multiple statements on one line

## Testing

The optional semicolon feature is thoroughly tested with:

- **20+ Go unit tests** in `internal/parser/semicolon_test.go` covering all statement types
- **5+ .uni integration tests** demonstrating real-world usage
- All existing test files work without modification (they don't use semicolons)

## Technical Details

The implementation is in:

- `internal/parser/error_listener.go` - Custom error listener that suppresses semicolon warnings
- `cmd/compiler/main.go` - Compiler setup to use the custom error listener
- `grammar/UnifiedParser.g4` - Grammar file with semicolon rules marked as optional (SEMI?)

The ANTLR parser automatically recovers from "missing semicolon" errors by inserting them, which allows the code to parse and compile successfully even without explicit semicolons in the source code.
