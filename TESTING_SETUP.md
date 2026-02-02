# Testing Setup Guide

This guide helps you set up and run the Unified compiler test suite on different platforms.

## Quick Start

### Linux

The test infrastructure works out of the box on Linux:

```bash
# Build the compiler
cd src/unified-compiler
make build

# Run all .uni test files
cd ../..
./test_all_uni_files.sh

# Run Go unit tests
cd src/unified-compiler
go test ./... -v
```

### macOS

The `test_all_uni_files.sh` script uses the `timeout` command which is not available on macOS by default.

**Option 1: Install coreutils (Recommended)**

```bash
# Install coreutils which provides gtimeout
brew install coreutils

# The script will automatically detect and use gtimeout
./test_all_uni_files.sh
```

**Option 2: Create an alias**

```bash
# Install coreutils
brew install coreutils

# Create an alias in your shell profile
echo "alias timeout=gtimeout" >> ~/.zshrc  # or ~/.bash_profile
source ~/.zshrc

# Now the script will work
./test_all_uni_files.sh
```

**Option 3: Run without timeout**

The script will automatically detect if timeout is not available and run without it:

```bash
# The script will warn but continue without timeout protection
./test_all_uni_files.sh
```

**Note**: Running without timeout is safe but may hang on infinite loops.

### Windows

Use WSL (Windows Subsystem for Linux) and follow the Linux instructions above.

---

## Test Infrastructure

### Integration Tests (`test_all_uni_files.sh`)

This script:
- Finds all `.uni` files in the repository (123 total)
- Compiles and runs each file with a 5-second timeout
- Categorizes results as PASS, FAIL, or TIMEOUT
- Generates detailed reports

**Output Files:**
- `test_results.txt` - Full output of all tests
- `test_working.txt` - List of passing tests  
- `test_failing.txt` - List of failing tests

**Current Results** (February 2, 2026):
- Total: 123 tests
- Passing: 56 (45.5%)
- Failing: 67 (54.5%)

### Go Unit Tests

The compiler has comprehensive Go unit tests:

```bash
cd src/unified-compiler

# Run all tests
go test ./...

# Run with verbose output
go test ./... -v

# Run specific package
go test ./internal/wasm -v

# Run with coverage
go test ./... -cover
```

**Current Results**:
- 5/6 packages passing
- 1 package with 3 known integration test failures (expected value differences)

---

## Understanding Test Results

### Integration Test Status

Tests are categorized by whether they compile and run successfully:

- **✅ PASS** - Compiles and completes successfully
- **❌ FAIL** - Compilation error or runtime error
- **⏱️ TIMEOUT** - Takes longer than 5 seconds (indicates infinite loop)

### Common Failure Patterns

From `TEST_BASELINE_2026-02-02.md`:

1. **Parser Grammar Issues (38.8%)** - Missing `Self`, `new`, `mut self` keywords
2. **Generic Type Inference (19.4%)** - Advanced generics not fully implemented
3. **Type Mismatch (10.4%)** - i32/i64 coercion issues
4. **Try Operator (10.4%)** - Incomplete codegen
5. **Edge Cases (20.9%)** - Various advanced features

---

## Troubleshooting

### "timeout: command not found" on macOS

**Solution**: Install coreutils as described above.

### Tests hang indefinitely

**Cause**: Infinite loop in compiler or generated code.

**Solution**: 
- On Linux: The timeout command will catch this
- On macOS without timeout: Press Ctrl+C to stop
- Report the issue with the specific test file

### Build fails with "go: module not found"

**Solution**:
```bash
cd src/unified-compiler
go mod download
go mod tidy
```

### "make: command not found"

**Solution**: Install make:
- macOS: `xcode-select --install`
- Linux: `sudo apt install build-essential` (Debian/Ubuntu)

### Parser generation fails

**Note**: Parser is pre-generated. You only need ANTLR4 if modifying grammar files.

**Solution**: Most users can skip parser generation. If needed:
```bash
# Install ANTLR4 (version 4.13.2 recommended)
# Then regenerate parser
cd src/unified-compiler
make parser
```

---

## Continuous Integration

The test suite runs automatically on:
- Every push to pull requests
- Merges to main branch
- Nightly builds

CI uses the same test scripts, ensuring consistency across environments.

---

## Adding New Tests

### Create a new .uni test file

1. Add your test file to the appropriate directory:
   - `src/unified-compiler/test/` - General tests
   - `src/unified-compiler/test/integration/` - Integration tests
   - `src/unified-compiler/test/generics/` - Generic tests
   - `src/unified-compiler/test/stdlib/` - Standard library tests

2. Follow the pattern of existing tests:
```unified
fn main() -> Int {
    // Your test code
    return 42  // Expected exit code
}
```

3. Run the test suite to verify:
```bash
./test_all_uni_files.sh
```

### Create a new Go unit test

1. Add test to appropriate package in `src/unified-compiler/internal/`

2. Follow Go testing conventions:
```go
func TestMyFeature(t *testing.T) {
    // Test code
}
```

3. Run the test:
```bash
cd src/unified-compiler
go test ./internal/yourpackage -v
```

---

## Performance Testing

For performance benchmarks:

```bash
cd src/unified-compiler

# Run benchmarks
go test -bench=. ./...

# Run with memory profiling
go test -bench=. -benchmem ./...
```

---

## Additional Resources

- **[TODO.md](TODO.md)** - Implementation roadmap and test status
- **[TEST_BASELINE_2026-02-02.md](TEST_BASELINE_2026-02-02.md)** - Comprehensive test analysis
- **[CONTRIBUTING.md](CONTRIBUTING.md)** - Contribution guidelines
- **[src/unified-compiler/TESTING.md](src/unified-compiler/TESTING.md)** - Detailed testing guide

---

## Questions?

- GitHub Issues: https://github.com/kristofer/Unified/issues
- Discussions: https://github.com/kristofer/Unified/discussions

---

**Last Updated**: February 2, 2026
