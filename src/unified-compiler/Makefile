.PHONY: all parser build test clean run-test run-all-tests vm-info

all: parser build

# Generate parser from ANTLR grammar
parser:
	@echo "Generating parser from ANTLR grammar..."
	@bash scripts/generate.sh

# Build the Unified compiler with VM backend
build:
	@echo "Building Unified compiler with VM backend..."
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
	for test_file in test/*.uni; do \
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
	if [ $$failed -eq 0 ]; then \
		echo "All tests passed!"; \
	else \
		echo "$$failed test(s) failed"; \
		exit 1; \
	fi

# Display VM information
vm-info:
	@echo "Unified VM-based Compiler"
	@echo "========================="
	@echo "Architecture: Stack-based virtual machine"
	@echo "Backend: Custom Go VM (no LLVM dependency)"
	@echo ""
	@echo "Build: make build"
	@echo "Run: ./bin/unified --input <file.uni>"
	@echo "See VM_README.md for more information"
