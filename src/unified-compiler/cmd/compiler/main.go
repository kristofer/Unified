package main

import (
	"flag"
	"fmt"
	"os"

	// Import compiler packages
	"unified-compiler/internal/ast"
	"unified-compiler/internal/bytecode"
	"unified-compiler/internal/parser"
	"unified-compiler/internal/vm"

	"github.com/antlr4-go/antlr/v4"
)

func compile(inputFile *string) {
	if *inputFile == "" {
		fmt.Printf("Error: No input file specified\n")
		os.Exit(1)
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

	// 2. Parse the input and build the AST
	parseTree := p.Program()
	astBuilder := ast.NewASTBuilder(*inputFile)
	program := astBuilder.VisitProgram(parseTree.(*parser.ProgramContext)).(*ast.Program)

	fmt.Printf("AST built with %d top-level items\n", len(program.Items))

	// 3. Generate bytecode
	generator := bytecode.NewGenerator()
	bc, err := generator.Generate(program)
	if err != nil {
		fmt.Printf("Error generating bytecode: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Generated %d instructions\n", len(bc.Instructions))

	// 4. Execute with VM
	virtualMachine := vm.NewVM(bc)
	result, err := virtualMachine.Run()
	if err != nil {
		fmt.Printf("Runtime error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Program completed successfully\n")
	fmt.Printf("Return value: %v\n", result)

	// Exit with the return value if it's an integer
	if result.Type == bytecode.ValueTypeInt {
		os.Exit(int(result.Int))
	} else {
		os.Exit(0)
	}
}

func main() {
	// Parse command-line args
	inputFile := flag.String("input", "", "Input source file")
	flag.Parse()

	compile(inputFile)
}