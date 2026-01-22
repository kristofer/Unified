# Phase 9: String Operations

**Status**: Not Started  
**Duration**: 2 weeks  
**Priority**: MEDIUM  
**Dependencies**: Phase 8 Complete (Arrays and Slices)

## 🎯 Goals

Implement comprehensive string operations and string interpolation. This brings strings from basic literals to a fully-featured text manipulation system with modern conveniences.

## 📋 Prerequisites

- [x] Phase 8 complete (arrays and slices working)
- [ ] Understanding of UTF-8 encoding
- [ ] Familiarity with string representation options
- [ ] Knowledge of string interpolation techniques

## ✨ Language Features to Add

### 1. String Methods
- `len()` - Get string length
- `is_empty()` - Check if empty
- `substring(start, end)` - Extract substring
- `split(delimiter)` - Split into array
- `to_upper()` - Convert to uppercase
- `to_lower()` - Convert to lowercase
- `trim()` - Remove whitespace
- `replace(old, new)` - Replace substrings
- `contains(substring)` - Check containment
- `starts_with(prefix)` - Check prefix
- `ends_with(suffix)` - Check suffix

### 2. String Interpolation
- Template strings: `"Hello, ${name}!"`
- Expression interpolation: `"Result: ${x + y}"`
- Format specifiers (basic): `"${value:2}"`
- Escaping: `"Use \${} for literal"`

### 3. String Concatenation
- `+` operator for string concatenation
- Efficient string building
- Automatic string conversion for interpolation

### 4. String Representation
- UTF-8 internal representation
- String as (pointer, length, capacity)
- Immutable strings (for now)
- String slicing compatibility with [u8] slices

## 📝 Implementation Tasks

### Task 9.1: String Representation (3-4 hours)
- [ ] Define String structure in VM
- [ ] Implement as (pointer, length, capacity)
- [ ] UTF-8 byte storage
- [ ] Handle string allocation
- [ ] Implement string copying
- [ ] Add memory management for strings
- [ ] Add tests for string creation

**String structure:**
```go
type String struct {
    Data     []byte  // UTF-8 bytes
    Length   int     // Number of bytes
    Capacity int     // Allocated capacity
}
```

### Task 9.2: String Methods Grammar (2-3 hours)
- [ ] Add method call syntax for strings (already in Phase 5)
- [ ] Ensure string type supports methods
- [ ] Add built-in string methods to prelude
- [ ] Test parsing string method calls

### Task 9.3: Implement Core String Methods (6-8 hours)
- [ ] Implement `len()` method
- [ ] Implement `is_empty()` method
- [ ] Implement `substring(start, end)` method
- [ ] Implement `contains(substring)` method
- [ ] Implement `starts_with(prefix)` method
- [ ] Implement `ends_with(suffix)` method
- [ ] Add bytecode instructions for string operations
- [ ] Implement in VM
- [ ] Add comprehensive tests

### Task 9.4: Implement String Manipulation Methods (6-8 hours)
- [ ] Implement `to_upper()` method
- [ ] Implement `to_lower()` method
- [ ] Implement `trim()` method
- [ ] Implement `replace(old, new)` method
- [ ] Implement `split(delimiter)` method
- [ ] Handle UTF-8 correctly
- [ ] Add tests for all methods

### Task 9.5: String Interpolation Grammar (3-4 hours)
- [ ] Add interpolated string syntax to grammar
- [ ] Parse `${}` expressions within strings
- [ ] Support nested expressions
- [ ] Handle escaping
- [ ] Test parsing interpolated strings

**Grammar to add:**
```antlr
stringLiteral
    : PLAIN_STRING
    | INTERPOLATED_STRING
    ;

interpolatedString
    : QUOTE interpolatedPart* QUOTE
    ;

interpolatedPart
    : STRING_TEXT
    | DOLLAR LBRACE expression RBRACE
    ;
```

### Task 9.6: Implement String Interpolation (4-6 hours)
- [ ] Desugar interpolation to concatenation
- [ ] Convert expressions to strings
- [ ] Generate efficient bytecode
- [ ] Handle format specifiers (basic)
- [ ] Add comprehensive tests

**Desugaring example:**
```unified
// Source:
"Hello, ${name}!"

// Desugars to:
"Hello, " + name.to_string() + "!"
```

### Task 9.7: String Concatenation (2-3 hours)
- [ ] Implement `+` operator for strings
- [ ] Efficient concatenation in VM
- [ ] Handle string + string
- [ ] Test concatenation

### Task 9.8: String Conversion (2-3 hours)
- [ ] Implement `to_string()` method for primitives
- [ ] Int to string
- [ ] Float to string
- [ ] Bool to string
- [ ] Add to_string trait/interface foundation
- [ ] Test conversions

### Task 9.9: Write Tests (4-5 hours)
- [ ] Parser tests for string methods
- [ ] Parser tests for interpolation
- [ ] String method execution tests
- [ ] String manipulation tests
- [ ] Interpolation tests
- [ ] Concatenation tests
- [ ] UTF-8 handling tests
- [ ] Integration tests
- [ ] Ensure code coverage ≥ 85%

### Task 9.10: Update Documentation (3-4 hours)
- [ ] Create PHASE9_SUMMARY.md
- [ ] Document string methods
- [ ] Document string interpolation
- [ ] Add string examples
- [ ] Update README.md
- [ ] Create string manipulation guide

## ✅ Test Requirements

**Minimum Test Coverage**: 85% for new code

### Parser Tests
- [ ] Parse string method calls
- [ ] Parse interpolated strings
- [ ] Parse nested expressions in interpolation
- [ ] Parse escaped interpolation
- [ ] Parse string concatenation
- [ ] Reject invalid syntax

### String Method Tests
- [ ] len() returns correct length
- [ ] is_empty() works correctly
- [ ] substring() extracts correctly
- [ ] contains() finds substrings
- [ ] starts_with() checks prefix
- [ ] ends_with() checks suffix
- [ ] All methods handle edge cases

### String Manipulation Tests
- [ ] to_upper() converts correctly
- [ ] to_lower() converts correctly
- [ ] trim() removes whitespace
- [ ] replace() substitutes correctly
- [ ] split() creates array
- [ ] Handle empty strings
- [ ] Handle special characters

### Interpolation Tests
- [ ] Simple variable interpolation
- [ ] Expression interpolation
- [ ] Multiple interpolations in one string
- [ ] Nested expressions
- [ ] Escaped ${}
- [ ] Format specifiers (if implemented)

### Concatenation Tests
- [ ] String + string works
- [ ] Multiple concatenations
- [ ] Empty string concatenation
- [ ] Performance is reasonable

### UTF-8 Tests
- [ ] ASCII strings work
- [ ] Multi-byte characters work
- [ ] len() counts bytes correctly
- [ ] substring() respects UTF-8 boundaries
- [ ] Methods handle UTF-8 correctly

### Integration Tests
- [ ] Build sentence from parts
- [ ] Parse and process text
- [ ] Format output with interpolation
- [ ] String searching and replacing

## 📚 Documentation Updates Required

### Update Existing Docs
- [ ] README.md: Add string features
- [ ] TESTING.md: Add string tests
- [ ] VM_README.md: Document string bytecode

### Create New Docs
- [ ] PHASE9_SUMMARY.md: Complete phase summary
- [ ] examples/STRINGS.md: String operations guide
- [ ] examples/INTERPOLATION.md: String interpolation guide
- [ ] docs/UTF8.md: UTF-8 handling documentation

### Add Example Programs
- [ ] `test/string_methods.uni`: String method examples
- [ ] `test/interpolation.uni`: Interpolation examples
- [ ] `test/text_processing.uni`: Text processing example
- [ ] `test/string_builder.uni`: String building example

## 🎓 Example Programs

### String Methods
```unified
fn string_operations() -> Int {
    let text = "Hello, World!"
    
    let length = text.len()              // 13
    let is_empty = text.is_empty()       // false
    let contains_world = text.contains("World")  // true
    let starts = text.starts_with("Hello")       // true
    let ends = text.ends_with("!")               // true
    let sub = text.substring(0, 5)       // "Hello"
    
    return length  // Returns 13
}

fn main() -> Int {
    return string_operations()
}
```

### String Manipulation
```unified
fn process_text(text: String) -> String {
    let trimmed = text.trim()
    let upper = trimmed.to_upper()
    let replaced = upper.replace("OLD", "NEW")
    return replaced
}

fn main() -> Int {
    let input = "  hello old world  "
    let result = process_text(input)
    // result = "HELLO NEW WORLD"
    
    return result.len()  // Returns 15
}
```

### String Interpolation
```unified
fn greet(name: String, age: Int) -> String {
    return "Hello, ${name}! You are ${age} years old."
}

fn calculate_message(x: Int, y: Int) -> String {
    let sum = x + y
    return "The sum of ${x} and ${y} is ${sum}"
}

fn main() -> Int {
    let greeting = greet("Alice", 30)
    // greeting = "Hello, Alice! You are 30 years old."
    
    let message = calculate_message(10, 20)
    // message = "The sum of 10 and 20 is 30"
    
    return 0
}
```

### String Splitting
```unified
fn count_words(text: String) -> Int {
    let words = text.split(" ")
    return words.len()
}

fn parse_csv(line: String) -> [String] {
    return line.split(",")
}

fn main() -> Int {
    let sentence = "The quick brown fox"
    let word_count = count_words(sentence)  // 4
    
    let csv = "apple,banana,cherry"
    let fruits = parse_csv(csv)
    // fruits = ["apple", "banana", "cherry"]
    
    return word_count  // Returns 4
}
```

### String Concatenation
```unified
fn build_sentence(subject: String, verb: String, object: String) -> String {
    return subject + " " + verb + " " + object + "."
}

fn main() -> Int {
    let sentence = build_sentence("The cat", "chased", "the mouse")
    // sentence = "The cat chased the mouse."
    
    return sentence.len()  // Returns 29
}
```

### Text Formatting
```unified
fn format_price(item: String, price: Int) -> String {
    return "Item: ${item}, Price: $${price}"
}

fn format_coordinates(x: Float, y: Float) -> String {
    return "(${x}, ${y})"
}

fn main() -> Int {
    let price_tag = format_price("Apple", 2)
    // price_tag = "Item: Apple, Price: $2"
    
    let coords = format_coordinates(3.14, 2.71)
    // coords = "(3.14, 2.71)"
    
    return 0
}
```

### String Search and Replace
```unified
fn censor_text(text: String) -> String {
    let step1 = text.replace("bad", "***")
    let step2 = step1.replace("worse", "***")
    let step3 = step2.replace("worst", "***")
    return step3
}

fn normalize_whitespace(text: String) -> String {
    let trimmed = text.trim()
    let no_double = trimmed.replace("  ", " ")
    return no_double
}

fn main() -> Int {
    let original = "This bad text is worse"
    let censored = censor_text(original)
    // censored = "This *** text is ***"
    
    return censored.len()
}
```

### Case Conversion
```unified
fn title_case(text: String) -> String {
    // Simplified title case (just first letter)
    if text.is_empty() {
        return text
    }
    
    let first = text.substring(0, 1).to_upper()
    let rest = text.substring(1, text.len()).to_lower()
    return first + rest
}

fn main() -> Int {
    let name = "aLiCe"
    let proper = title_case(name)
    // proper = "Alice"
    
    return proper.len()  // Returns 5
}
```

## 🏆 Success Criteria

- [ ] All parser tests pass (minimum 12 tests)
- [ ] All string method tests pass (minimum 20 tests)
- [ ] All manipulation tests pass (minimum 15 tests)
- [ ] All interpolation tests pass (minimum 12 tests)
- [ ] All concatenation tests pass (minimum 8 tests)
- [ ] All UTF-8 tests pass (minimum 10 tests)
- [ ] All integration tests pass (minimum 6 tests)
- [ ] String operations work correctly
- [ ] Interpolation is convenient and works
- [ ] UTF-8 handling is correct
- [ ] String methods are efficient
- [ ] No regressions in previous phase tests
- [ ] Code coverage ≥ 85%
- [ ] Documentation complete
- [ ] All example programs run correctly

## 💡 Implementation Notes

### Implementation Order
1. String representation in VM
2. Basic string methods (len, is_empty)
3. String inspection methods (contains, starts_with, ends_with)
4. String manipulation (substring, trim, case conversion)
5. String splitting
6. String concatenation
7. String interpolation grammar
8. String interpolation implementation
9. String conversion (to_string)
10. Integration and testing

### Testing Strategy
- Test ASCII strings first, then UTF-8
- Verify each method independently
- Test edge cases (empty strings, boundaries)
- Test interpolation with various expressions
- Performance test concatenation
- Test memory management

### Common Pitfalls
1. Incorrect UTF-8 handling (byte vs character confusion)
2. Off-by-one errors in substring
3. Memory leaks in string operations
4. Inefficient concatenation (repeated copying)
5. Not handling empty strings
6. Interpolation parsing complexity
7. Case conversion with non-ASCII

### Debugging Tips
- Print string bytes during development
- Verify UTF-8 encoding manually
- Test with simple ASCII first
- Check memory allocation/deallocation
- Use small examples before complex ones
- Verify interpolation desugaring

### String Representation Details
- Store as UTF-8 bytes (standard)
- Length = byte count, not character count
- Consider rope data structure for large strings (future)
- String interning for literals (future optimization)

### UTF-8 Considerations
- 1-4 bytes per character
- Validate UTF-8 sequences
- Substring must not split multi-byte characters
- Length counts bytes, not code points
- Consider separate char_len() method

### Performance Considerations
- String concatenation creates new string
- Multiple concatenations are expensive
- Consider StringBuilder pattern (future)
- String methods should avoid copying when possible
- Implement copy-on-write (future)

### Future Enhancements
This phase prepares for:
- StringBuilder for efficient building
- Regular expressions
- String formatting with printf-style
- Unicode normalization
- String interning
- Raw strings and byte strings

---

**Labels**: `phase-9`, `enhancement`, `strings`, `text-processing`  
**Milestone**: Phase 9 - String Operations  
**Assignee**: TBD
