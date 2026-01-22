.PHONY: all parser build test clean run-test vm-info

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
