package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	// Import your generated parser package
	"unified-compiler/internal/ast"
	"unified-compiler/internal/codegen"
	"unified-compiler/internal/parser" // Import your generated parser package

	//"github.com/kristofer/Unified/src/unified-compiler/internal/parser"
	// OR with the correct module path from go.mod
	//"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/antlr4-go/antlr/v4"
)

func compile(inputFile *string, outputFile *string) {
	// Parse command-line args
	// inputFile := flag.String("input", "", "Input source file")
	// outputFile := flag.String("output", "", "Output LLVM IR file")
	flag.Parse()

	if *inputFile == "" {
		fmt.Printf("Error: No input file specified\n")
		os.Exit(1)
	}

	if *outputFile == "" {
		*outputFile = strings.TrimSuffix(*inputFile, filepath.Ext(*inputFile)) + ".ll"
	}

	// 1. Set up the ANTLR parser
	sourceCode, err := os.ReadFile(*inputFile)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		os.Exit(1)
	}

	inputStream := antlr.NewInputStream(string(sourceCode))
	lexer := parser.NewUnifiedLexer(inputStream)
	tokenStream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewUnifiedParser(tokenStream)

	// Optional: Add error handling
	// p.RemoveErrorListeners()
	// p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))

	// 2. Parse the input and build the AST
	parseTree := p.Program()
	astBuilder := ast.NewASTBuilder(*inputFile)
	program := astBuilder.VisitProgram(parseTree.(*parser.ProgramContext)).(*ast.Program)

	fmt.Printf("AST built with %d top-level items\n", len(program.Items))

	// 3. Generate LLVM IR
	generator := codegen.NewCodeGenerator(filepath.Base(*inputFile))
	llvmIR, err := generator.Generate(program)
	if err != nil {
		fmt.Printf("Error generating LLVM IR: %v\n", err)
		os.Exit(1)
	}

	// 4. Write the LLVM IR to a file
	err = os.WriteFile(*outputFile, []byte(llvmIR), 0644)
	if err != nil {
		fmt.Printf("Error writing output file: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Successfully compiled %s to %s\n", *inputFile, *outputFile)

	// Clean up LLVM resources
	generator.Dispose()
}

func testparse(inputFile *string, outputFile *string) {

	if *inputFile == "" {
		fmt.Printf("Error: No input file specified\n")
		os.Exit(1)
	}

	if *outputFile == "" {
		*outputFile = strings.TrimSuffix(*inputFile, filepath.Ext(*inputFile)) + ".ll"
	}

	// 1. Set up the ANTLR parser
	sourceCode, err := os.ReadFile(*inputFile)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		os.Exit(1)
	}

	inputStream := antlr.NewInputStream(string(sourceCode))
	lexer := parser.NewUnifiedLexer(inputStream)
	tokenStream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewUnifiedParser(tokenStream)

	// Optional: Add error handling
	// p.RemoveErrorListeners()
	// p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))

	// 2. Parse the input and build the AST
	parseTree := p.Program()
	astBuilder := ast.NewASTBuilder(*inputFile)
	program := astBuilder.VisitProgram(parseTree.(*parser.ProgramContext)).(*ast.Program)

	fmt.Printf("AST built with %d top-level items\n", len(program.Items))

	// 3. Generate LLVM IR
	generator := codegen.NewCodeGenerator(filepath.Base(*inputFile))
	llvmIR, err := generator.Generate(program)
	if err != nil {
		fmt.Printf("Error generating LLVM IR: %v\n", err)
		os.Exit(1)
	}

	// 4. Write the LLVM IR to a file
	err = os.WriteFile(*outputFile, []byte(llvmIR), 0644)
	if err != nil {
		fmt.Printf("Error writing output file: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Successfully compiled %s to %s\n", *inputFile, *outputFile)

	// Clean up LLVM resources
	generator.Dispose()
}

func main() {
	// Parse command-line args
	inputFile := flag.String("input", "", "Input source file")
	outputFile := flag.String("output", "", "Output LLVM IR file")
	flag.Parse()

	testparse(inputFile, outputFile)
	//compile(inputFile, outputFile)
}