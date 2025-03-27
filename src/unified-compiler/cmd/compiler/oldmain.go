package main

import (
	"flag"
	"fmt"
	"os"

	// Import your generated parser package
	"unified-compiler/internal/parser" // Import your generated parser package
	//"github.com/kristofer/Unified/src/unified-compiler/internal/parser"
	// OR with the correct module path from go.mod
	//"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/antlr4-go/antlr/v4"
)

func oldmain() {
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

	fmt.Printf("Parsing & Listening %s\n", *inputFile)
	// Create the input stream
	sourceCode, err := os.ReadFile(*inputFile)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		os.Exit(1)
	}
	//fmt.Printf("SourceCode: %v\n", string(sourceCode))
	inputStream := antlr.NewInputStream(string(sourceCode))

	// Create the lexer
	lexer := parser.NewUnifiedLexer(inputStream)

	// Create the token stream
	tokenStream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// fmt.Printf("Len of TokenStream: %v\n", len(tokenStream.GetAllText()))
	// tokenstr := tokenStream.GetAllTokens()
	// for _, token := range tokenstr {
	// 	fmt.Printf("Token: %v\n", token)
	// }
	// Create the parser
	p := parser.NewUnifiedParser(tokenStream)

	// Add our debug listener
	debugListener := parser.NewDebugListener()
	p.AddParseListener(debugListener)

	// For better error handling
	//p.RemoveErrorListeners()
	//p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))

	// Start parsing from the program rule
	// programContext := p.Program()
	pctxt := p.Program()

	fmt.Printf("Context: %+v \nLen: %v\n", pctxt, len(pctxt.GetChildren()))
	fmt.Printf("Parsing successful\n")
	fmt.Printf("Debug visitor: %v\n", debugListener)

	// Parse source code
	//inputStream = antlr.NewInputStream(*inputFile)
	sourceCode, err = os.ReadFile(*inputFile)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		os.Exit(1)
	}
	//fmt.Printf("SourceCode: %v\n", string(sourceCode))
	inputStream = antlr.NewInputStream(string(sourceCode))

	lexer = parser.NewUnifiedLexer(inputStream)
	tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p = parser.NewUnifiedParser(tokens)
	// Create the debug visitor
	debugVisitor := parser.NewDebugVisitor()
	//debugVisitor := &parser.DebugVisitor{} // Use direct struct initialization

	// Start parsing from the root rule and visit the tree
	tree := p.Program()
	fmt.Printf("Tree: %+v \nLen: %v\n", tree, len(tree.GetChildren()))
	debugVisitor.Visit(tree)

	tree.Accept(debugVisitor)
	fmt.Printf("Debug visitor: %+v\n", debugVisitor)

	fmt.Printf("Visiting successful\n")

	// // Start parsing from the program rule
	// //programContext := p.Program()

	// // Check for syntax errors
	// // if p.GetNumberOfSyntaxErrors() > 0 {
	// // 	fmt.Printf("Found %d syntax errors\n", p.GetNumberOfSyntaxErrors())
	// // 	os.Exit(1)
	// // }

	// // 2. Build AST
	// programContext := p.Program()

	// // Check for syntax errors
	// // if p.GetNumberOfSyntaxErrors() > 0 {
	// // 	fmt.Printf("Found %d syntax errors\n", p.GetNumberOfSyntaxErrors())
	// // 	os.Exit(1)
	// // }

	// fmt.Println("Parsing successful")

	// // 2. Build AST
	// astBuilder := ast.NewASTBuilder(*inputFile)
	// astRoot := astBuilder.VisitProgram(programContext.(*parser.ProgramContext)).(*ast.Program)

	// fmt.Printf("Built AST with %d top-level items\n", len(astRoot.Items))

	// // 3. Perform semantic analysis
	// // 4. Generate LLVM IR
	// // 5. Output to file
}
