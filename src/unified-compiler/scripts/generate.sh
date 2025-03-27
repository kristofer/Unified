#!/bin/bash

# This script generates Go code from the ANTLR grammar

# Ensure ANTLR is installed
if ! command -v antlr4 &> /dev/null; then
    echo "ANTLR4 not found. Please install it first."
    echo "Visit https://www.antlr.org/ for installation instructions."
    exit 1
fi

# Generate parser
echo "Generating parser from grammar..."
cd $(dirname "$0")/../grammar
antlr4 -Dlanguage=Go -package parser -visitor -listener -o ../internal/parser UnifiedParser.g4 UnifiedLexer.g4

echo "Done generating parser code."
