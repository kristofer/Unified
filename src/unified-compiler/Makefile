.PHONY: all parser build test clean

all: parser build

parser:
	@echo "Generating parser..."
	@bash scripts/generate.sh

build:
	@echo "Building compiler..."
	@go build -o bin/unified-compiler ./cmd/compiler

test:
	@echo "Running tests..."
	@go test -v ./...

clean:
	@echo "Cleaning up..."
	@rm -rf bin/
	@find ./internal/parser -name "Unified*.go" -type f -delete
	