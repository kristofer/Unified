.PHONY: all parser build test clean run-test run-all-tests run-lib-tests wasm-info stdlib-info

all: parser build

# Generate parser from ANTLR grammar
parser:
	@echo "Generating parser from ANTLR grammar..."
	@bash scripts/generate.sh

# Build the Unified compiler with WASM backend
build:
	@echo "Building Unified compiler with WASM backend..."
	@go build -o bin/unified ./cmd/compiler
	@echo "Build complete! Binary: bin/unified"

# Run tests
test:
	@echo "Running tests..."
	@go test -v ./...

# Clean build artifacts
clean:
	@echo "Cleaning up..."
	@rm -rf bin/
	@find ./internal/parser -name "Unified*.go" -type f -delete

# Run a test program
run-test: build
	@echo "Running minimal test..."
	./bin/unified --input minimal_test.uni
	@echo "Exit code: $$?"

# Run all Unified test files
run-all-tests: build
	@echo "Running all Unified test files..."
	@failed=0; \
	total=0; \
	for test_file in $$(find test -name "*.uni" -type f | sort); do \
		total=$$((total + 1)); \
		echo ""; \
		echo "Testing $$test_file..."; \
		output=$$(./bin/unified --input "$$test_file" 2>&1); \
		exit_code=$$?; \
		echo "$$output"; \
		if echo "$$output" | grep -q "panic:\|Error generating\|error:"; then \
			echo "✗ $$test_file FAILED (compilation/runtime error)"; \
			failed=$$((failed + 1)); \
		elif echo "$$output" | grep -q "Program completed successfully"; then \
			echo "✓ $$test_file passed (exit code: $$exit_code)"; \
		else \
			echo "✗ $$test_file FAILED (unexpected output)"; \
			failed=$$((failed + 1)); \
		fi; \
	done; \
	echo ""; \
	echo "Ran $$total tests total"; \
	if [ $$failed -eq 0 ]; then \
		echo "All tests passed!"; \
	else \
		echo "$$failed test(s) failed"; \
		exit 1; \
	fi

# Run standard library tests
run-lib-tests: build
	@echo "Running standard library tests..."
	@failed=0; \
	total=0; \
	for test_file in $$(find test/stdlib -name "*.uni" -type f | sort); do \
		total=$$((total + 1)); \
		echo ""; \
		echo "Testing $$test_file..."; \
		output=$$(./bin/unified --input "$$test_file" 2>&1); \
		exit_code=$$?; \
		echo "$$output"; \
		if echo "$$output" | grep -q "panic:\|Error generating\|error:"; then \
			echo "✗ $$test_file FAILED (compilation/runtime error)"; \
			failed=$$((failed + 1)); \
		elif echo "$$output" | grep -q "Program completed successfully"; then \
			echo "✓ $$test_file passed (exit code: $$exit_code)"; \
		else \
			echo "✗ $$test_file FAILED (unexpected output)"; \
			failed=$$((failed + 1)); \
		fi; \
	done; \
	echo ""; \
	echo "Ran $$total standard library tests total"; \
	if [ $$failed -eq 0 ]; then \
		echo "All standard library tests passed!"; \
	else \
		echo "$$failed test(s) failed"; \
		exit 1; \
	fi

# Display WASM information
wasm-info:
	@echo "Unified WASM-based Compiler"
	@echo "==========================="
	@echo "Architecture: WebAssembly target"
	@echo "Backend: WASM generation with wazero runtime"
	@echo "Runtime: Embedded wazero WebAssembly runtime"
	@echo ""
	@echo "Build: make build"
	@echo "Run: ./bin/unified --input <file.uni>"
	@echo "See WASM documentation for more information"

# Display standard library information
stdlib-info:
	@echo "Unified Standard Library"
	@echo "======================="
	@echo ""
	@echo "Generic Data Structures:"
	@echo "  • List<T>        - Dynamic array"
	@echo "  • Stack<T>       - LIFO data structure"
	@echo "  • Queue<T>       - FIFO data structure"
	@echo "  • HashMap<K, V>  - Hash table"
	@echo "  • Set<T>         - Unique element collection"
	@echo "  • BinaryTree<T>  - Binary search tree"
	@echo ""
	@echo "Files:"
	@echo "  stdlib/List.uni"
	@echo "  stdlib/Stack.uni"
	@echo "  stdlib/Queue.uni"
	@echo "  stdlib/HashMap.uni"
	@echo "  stdlib/Set.uni"
	@echo "  stdlib/BinaryTree.uni"
	@echo ""
	@echo "Testing:"
	@echo "  make run-lib-tests  - Run all standard library tests (18 tests)"
	@echo ""
	@echo "Documentation:"
	@echo "  stdlib/README.md - Full API reference"
	@echo "  STDLIB_IMPLEMENTATION.md - Implementation summary"
