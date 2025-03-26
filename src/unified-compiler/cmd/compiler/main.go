package main

import (
	"flag"
	"fmt"
	"os"

	"unified-compiler/internal/ast"    // Import your generated parser package
	"unified-compiler/internal/parser" // Import your generated parser package

	//"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/antlr4-go/antlr/v4"
)

func main() {
	// Parse command-line arguments
	inputFile := flag.String("input", "", "Input source file")
	outputFile := flag.String("output", "", "Output LLVM IR file")
	flag.Parse()

	if *inputFile == "" {
		fmt.Println("Error: Input file required")
		flag.Usage()
		os.Exit(1)
	}

	fmt.Printf("Compiling %s to %s\n", *inputFile, *outputFile)

	// TODO: Implement compilation pipeline
	// 1. Parse input file using ANTLR parser
	input, err := os.ReadFile(*inputFile)
	if err != nil {
		fmt.Printf("Error reading input file: %v\n", err)
		os.Exit(1)
	}

	// Create the input stream
	inputStream := antlr.NewInputStream(string(input))

	// Create the lexer
	lexer := parser.NewUnifiedLexer(inputStream)

	// Create the token stream
	tokenStream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the parser
	p := parser.NewUnifiedParser(tokenStream)

	// For better error handling
	p.RemoveErrorListeners()
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))

	// Start parsing from the program rule
	//programContext := p.Program()
	_ = p.Program()

	// Check for syntax errors
	// if p.GetNumberOfSyntaxErrors() > 0 {
	// 	fmt.Printf("Found %d syntax errors\n", p.GetNumberOfSyntaxErrors())
	// 	os.Exit(1)
	// }

	fmt.Println("Parsing successful")

	// 2. Build AST
	programContext := p.Program()

	// Check for syntax errors
	// if p.GetNumberOfSyntaxErrors() > 0 {
	// 	fmt.Printf("Found %d syntax errors\n", p.GetNumberOfSyntaxErrors())
	// 	os.Exit(1)
	// }

	fmt.Println("Parsing successful")

	// 2. Build AST
	astBuilder := ast.NewASTBuilder(*inputFile)
	astRoot := astBuilder.VisitProgram(programContext.(*parser.ProgramContext)).(*ast.Program)

	fmt.Printf("Built AST with %d top-level items\n", len(astRoot.Items))

	// 3. Perform semantic analysis
	// 4. Generate LLVM IR
	// 5. Output to file
}
