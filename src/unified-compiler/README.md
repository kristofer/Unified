# Unified Compiler

A compiler for the Unified language that generates LLVM IR.

## Project Structure

- `cmd/compiler`: Command-line interface
- `grammar`: ANTLR grammar definition
- `internal/ast`: Abstract Syntax Tree definitions
- `internal/parser`: ANTLR-generated parser and visitors
- `internal/semantic`: Semantic analysis
- `internal/codegen`: LLVM IR code generation
- `pkg/compiler`: Public compiler API
- `scripts`: Build and utility scripts
- `test`: Test cases and fixtures

## Building

1. Ensure you have Go 1.19+ and ANTLR4 installed
2. Run `make parser` to generate the parser
3. Run `make build` to build the compiler

## Usage
```
./bin/unified-compiler --input program.uni --output program.ll
```


# Make generate script executable
chmod +x unified-compiler/scripts/generate.sh

echo "Project structure created successfully in ./unified-compiler/"
echo "Next steps:"
echo "1. cd unified-compiler"
echo "2. Install ANTLR4 if not already installed"
echo "3. Run 'make parser' to generate the parser"
echo "4. Run 'make build' to build the compiler"

