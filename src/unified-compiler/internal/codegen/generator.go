package codegen

import (
	"fmt"
	"log"
	"unified-compiler/internal/ast"

	"tinygo.org/x/go-llvm"
)

// CodeGenerator handles LLVM IR generation
type CodeGenerator struct {
	context        llvm.Context
	module         llvm.Module
	builder        llvm.Builder
	symbolTable    map[string]llvm.Value // Track variables in scope
	typeTable      map[string]llvm.Type  // Map Unified types to LLVM types
	currentFunc    llvm.Value            // Current function being generated
	stringLiterals map[string]llvm.Value // String literal cache
}

// NewCodeGenerator creates a new LLVM IR generator
func NewCodeGenerator(moduleName string) *CodeGenerator {
	// Initialize LLVM
	llvm.InitializeNativeTarget()
	llvm.InitializeNativeAsmPrinter()

	context := llvm.NewContext()
	module := context.NewModule(moduleName)
	builder := context.NewBuilder()

	cg := &CodeGenerator{
		context:        context,
		module:         module,
		builder:        builder,
		symbolTable:    make(map[string]llvm.Value),
		typeTable:      make(map[string]llvm.Type),
		stringLiterals: make(map[string]llvm.Value),
	}

	// Initialize basic types mapping
	cg.initializeTypeTable()

	return cg
}

// Initialize type mapping between Unified types and LLVM types
func (g *CodeGenerator) initializeTypeTable() {
	// Basic types
	g.typeTable["i8"] = g.context.Int8Type()
	g.typeTable["i16"] = g.context.Int16Type()
	g.typeTable["i32"] = g.context.Int32Type()
	g.typeTable["i64"] = g.context.Int64Type()
	g.typeTable["u8"] = g.context.Int8Type()
	g.typeTable["u16"] = g.context.Int16Type()
	g.typeTable["u32"] = g.context.Int32Type()
	g.typeTable["u64"] = g.context.Int64Type()
	g.typeTable["f32"] = g.context.FloatType()
	g.typeTable["f64"] = g.context.DoubleType()
	g.typeTable["bool"] = g.context.Int1Type()
	g.typeTable["char"] = g.context.Int8Type()
	g.typeTable["void"] = g.context.VoidType()
}

// GetLLVMType converts a Unified type to an LLVM type
func (g *CodeGenerator) GetLLVMType(t ast.Type) (llvm.Type, error) {
	// Handle different type structures
	switch ty := t.(type) {
	case *ast.TypeReference:
		if llvmType, exists := g.typeTable[ty.Name]; exists {
			return llvmType, nil
		}
		return llvm.Type{}, fmt.Errorf("unknown type: %s", ty.Name)

	case *ast.ReferenceType:
		// Handle reference types
		baseType, err := g.GetLLVMType(ty.BaseType)
		if err != nil {
			return llvm.Type{}, err
		}
		return llvm.PointerType(baseType, 0), nil
	}

	// Default to i32 if type can't be resolved
	log.Printf("Warning: Using default i32 type for unrecognized type")
	return g.context.Int32Type(), nil
}

// Generate produces LLVM IR for the entire program
func (g *CodeGenerator) Generate(program *ast.Program) (string, error) {
	// Phase 1: Only handle global variables and functions
	for _, item := range program.Items {
		switch item := item.(type) {
		case *ast.FunctionDecl:
			if _, err := g.GenerateFunction(item); err != nil {
				return "", fmt.Errorf("error generating function: %w", err)
			}

		case *ast.ConstantDecl:
			if _, err := g.GenerateGlobalConst(item); err != nil {
				return "", fmt.Errorf("error generating constant: %w", err)
			}

		default:
			// Skip other top-level declarations in Phase 1
			log.Printf("Skipping unsupported top-level declaration: %T", item)
		}
	}

	// Verify the generated module
	if err := llvm.VerifyModule(g.module, llvm.PrintMessageAction); err != nil {
		return "", fmt.Errorf("module verification failed: %v", err)
	}

	// Return the module as LLVM IR text
	return g.module.String(), nil
}

// GenerateFunction creates an LLVM function (Phase 1)
func (g *CodeGenerator) GenerateFunction(funcDecl *ast.FunctionDecl) (llvm.Value, error) {
	// Create function type
	var paramTypes []llvm.Type
	for _, param := range funcDecl.Parameters {
		paramType, err := g.GetLLVMType(param.Type)
		if err != nil {
			paramType = g.context.Int32Type()
			log.Printf("Using default type for parameter %s", param.Name)
		}
		paramTypes = append(paramTypes, paramType)
	}

	// Determine return type
	var returnType llvm.Type = g.context.VoidType()
	if funcDecl.ReturnType != nil {
		var err error
		returnType, err = g.GetLLVMType(funcDecl.ReturnType)
		if err != nil {
			log.Printf("Using void return type for function %s due to: %v",
				funcDecl.Name, err)
		}
	}

	// Create function type
	funcType := llvm.FunctionType(returnType, paramTypes, false)

	// Create function
	function := llvm.AddFunction(g.module, funcDecl.Name, funcType)

	// Save the current context
	prevFunc := g.currentFunc
	prevSymbolTable := g.symbolTable

	// Set new function context
	g.currentFunc = function
	g.symbolTable = make(map[string]llvm.Value)

	// Only generate body if it exists
	if funcDecl.Body != nil {
		// Create entry block
		entryBlock := llvm.AddBasicBlock(function, "entry")
		g.builder.SetInsertPoint(entryBlock, entryBlock.FirstInstruction())

		// Add parameters to symbol table
		for i, param := range funcDecl.Parameters {
			paramValue := function.Param(i)
			paramValue.SetName(param.Name)

			// For now, store parameters directly (not as allocas)  
			// TODO: Later we might need allocas for mutable parameters
			g.symbolTable[param.Name] = paramValue
		}

		// Generate statements in function body
		for _, stmt := range funcDecl.Body.Statements {
			if _, err := g.GenerateStatement(stmt); err != nil {
				g.currentFunc = prevFunc
				g.symbolTable = prevSymbolTable
				return llvm.Value{}, fmt.Errorf("error in function body: %w", err)
			}
		}

		// Handle trailing expression if present
		if funcDecl.Body.Expression != nil && returnType.C != g.context.VoidType().C {
			exprValue, err := g.GenerateExpr(funcDecl.Body.Expression)
			if err != nil {
				g.currentFunc = prevFunc
				g.symbolTable = prevSymbolTable
				return llvm.Value{}, fmt.Errorf("error in trailing expression: %w", err)
			}
			g.builder.CreateRet(exprValue)
		} else {
			// Only add implicit return if we haven't already returned
			// For now, skip implicit returns when we have explicit returns
		}
	}

	// Restore previous context
	g.currentFunc = prevFunc
	g.symbolTable = prevSymbolTable

	return function, nil
}

// GenerateGlobalConst creates a global constant (Phase 1)
func (g *CodeGenerator) GenerateGlobalConst(constDecl *ast.ConstantDecl) (llvm.Value, error) {
	// Only literals are supported in Phase 1
	lit, ok := constDecl.Value.(*ast.Literal)
	if !ok {
		return llvm.Value{}, fmt.Errorf("only literal constants supported in Phase 1")
	}

	// Generate the constant value
	constValue := g.GenerateLiteral(lit)

	// Create global constant
	global := llvm.AddGlobal(g.module, constValue.Type(), constDecl.Name)
	global.SetInitializer(constValue)
	global.SetGlobalConstant(true)

	// Add to symbol table
	g.symbolTable[constDecl.Name] = global

	return global, nil
}

// GenerateStatement generates LLVM IR for a statement (Phase 1)
func (g *CodeGenerator) GenerateStatement(stmt ast.Statement) (llvm.Value, error) {
	switch stmt := stmt.(type) {
	case *ast.LetStatement:
		return g.GenerateLetStatement(stmt)

	case *ast.VarStatement:
		return g.GenerateVarDecl(stmt)

	case *ast.ExprStatement:
		return g.GenerateExpr(stmt.Expression)

	case *ast.ReturnStatement:
		return g.GenerateReturnStatement(stmt)

	case *ast.IfStatement:
		return g.GenerateIfStatement(stmt)

	default:
		log.Printf("Warning: Skipping unsupported statement type in Phase 1: %T", stmt)
		return llvm.Value{}, nil
	}
}

// GenerateLetStatement generates a let statement (Phase 1)
func (g *CodeGenerator) GenerateLetStatement(letStmt *ast.LetStatement) (llvm.Value, error) {
	// Determine variable type - use i32 as default for literals without explicit type
	var varType llvm.Type = g.context.Int32Type() // default type
	if letStmt.Type != nil {
		var err error
		varType, err = g.GetLLVMType(letStmt.Type)
		if err != nil {
			log.Printf("Using default i32 type for let variable %s", letStmt.Name)
		}
	}

	// Create stack allocation
	log.Printf("Creating alloca for %s with type: %s", letStmt.Name, varType.String())
	alloca := g.builder.CreateAlloca(varType, letStmt.Name)
	log.Printf("Alloca created successfully, type: %s", alloca.Type().String())

	// Initialize with the value
	if letStmt.Value != nil {
		initValue, err := g.GenerateExpr(letStmt.Value)
		if err != nil {
			return llvm.Value{}, fmt.Errorf("error in initializer: %w", err)
		}
		g.builder.CreateStore(initValue, alloca)
	}

	// Add to symbol table
	g.symbolTable[letStmt.Name] = alloca

	return alloca, nil
}

// GenerateVarDecl generates a local variable declaration (Phase 1)
func (g *CodeGenerator) GenerateVarDecl(varStmt *ast.VarStatement) (llvm.Value, error) {
	// Determine variable type
	varType, err := g.GetLLVMType(varStmt.Type)
	if err != nil {
		varType = g.context.Int32Type()
		log.Printf("Using default i32 type for variable %s", varStmt.Name)
	}

	// Create stack allocation
	alloca := g.builder.CreateAlloca(varType, varStmt.Name)

	// Initialize if there's a value
	if varStmt.Value != nil {
		initValue, err := g.GenerateExpr(varStmt.Value)
		if err != nil {
			return llvm.Value{}, fmt.Errorf("error in initializer: %w", err)
		}
		g.builder.CreateStore(initValue, alloca)
	}

	// Add to symbol table
	g.symbolTable[varStmt.Name] = alloca

	return alloca, nil
}

// GenerateReturnStatement handles return statements (Phase 1)
func (g *CodeGenerator) GenerateReturnStatement(stmt *ast.ReturnStatement) (llvm.Value, error) {
	if stmt.Value == nil {
		return g.builder.CreateRetVoid(), nil
	}

	retValue, err := g.GenerateExpr(stmt.Value)
	if err != nil {
		return llvm.Value{}, err
	}
	return g.builder.CreateRet(retValue), nil
}

// GenerateIfStatement handles if statements (Phase 1)
func (g *CodeGenerator) GenerateIfStatement(ifStmt *ast.IfStatement) (llvm.Value, error) {
	// Create basic blocks
	thenBlock := llvm.AddBasicBlock(g.currentFunc, "then")
	var elseBlock llvm.BasicBlock
	if ifStmt.ElseBlock != nil {
		elseBlock = llvm.AddBasicBlock(g.currentFunc, "else")
	}
	mergeBlock := llvm.AddBasicBlock(g.currentFunc, "ifcont")

	// Generate condition
	condValue, err := g.GenerateExpr(ifStmt.Condition)
	if err != nil {
		return llvm.Value{}, fmt.Errorf("error in condition: %w", err)
	}

	// Convert condition to boolean if needed
	if condValue.Type().C != g.context.Int1Type().C {
		condValue = g.builder.CreateICmp(llvm.IntNE, condValue,
			llvm.ConstInt(condValue.Type(), 0, false), "cond")
	}

	// Create branch instruction
	if ifStmt.ElseBlock != nil {
		g.builder.CreateCondBr(condValue, thenBlock, elseBlock)
	} else {
		g.builder.CreateCondBr(condValue, thenBlock, mergeBlock)
	}

	// Generate then block
	g.builder.SetInsertPointAtEnd(thenBlock)
	for _, stmt := range ifStmt.ThenBlock.Statements {
		if _, err := g.GenerateStatement(stmt); err != nil {
			return llvm.Value{}, fmt.Errorf("error in then block: %w", err)
		}
	}

	// Add branch to merge block if needed
	// Simplified terminator check
	if true { // TODO: implement proper terminator check
		g.builder.CreateBr(mergeBlock)
	}

	// Generate else block if present
	if ifStmt.ElseBlock != nil {
		g.builder.SetInsertPointAtEnd(elseBlock)
		for _, stmt := range ifStmt.ElseBlock.Statements {
			if _, err := g.GenerateStatement(stmt); err != nil {
				return llvm.Value{}, fmt.Errorf("error in else block: %w", err)
			}
		}

		// Add branch to merge block
		// Simplified terminator check
	if true { // TODO: implement proper terminator check
			g.builder.CreateBr(mergeBlock)
		}
	}

	// Continue with merge block
	g.builder.SetInsertPointAtEnd(mergeBlock)

	return llvm.Value{}, nil
}

// GenerateExpr handles expression generation (Phase 1)
func (g *CodeGenerator) GenerateExpr(expr ast.Expression) (llvm.Value, error) {
	switch expr := expr.(type) {
	case *ast.Literal:
		return g.GenerateLiteral(expr), nil

	case *ast.Identifier:
		// Look up variable in the symbol table
		if val, exists := g.symbolTable[expr.Name]; exists {
			// If it's a global constant, return it directly
			if !val.IsAGlobalVariable().IsNil() && val.IsGlobalConstant() {
				return val, nil
			}
			
			// Check if it's an alloca (local variable) that needs loading
			if val.Type().TypeKind() == llvm.PointerTypeKind {
				loadType := val.Type().ElementType()
				// Ensure we have a valid element type
				if !loadType.IsNil() {
					return g.builder.CreateLoad(loadType, val, expr.Name), nil
				} else {
					log.Printf("Warning: Invalid element type for variable %s, returning alloca directly", expr.Name)
				}
			}
			
			// Otherwise return the value directly (e.g., function parameters)
			return val, nil
		}
		return llvm.Value{}, fmt.Errorf("undefined variable: %s", expr.Name)

	case *ast.BinaryExpr:
		return g.GenerateBinaryExpr(expr)

	case *ast.UnaryExpr:
		return g.GenerateUnaryExpr(expr)

	case *ast.CallExpr:
		return g.GenerateCallExpr(expr)

	default:
		return llvm.Value{}, fmt.Errorf("unsupported expression in Phase 1: %T", expr)
	}
}

// GenerateLiteral creates LLVM values for literals (Phase 1)
func (g *CodeGenerator) GenerateLiteral(lit *ast.Literal) llvm.Value {
	switch lit.Kind {
	case ast.LiteralInt:
		// Create 32-bit integer
		var val int64
		fmt.Sscanf(lit.Value, "%d", &val)
		return llvm.ConstInt(g.context.Int32Type(), uint64(val), true)

	case ast.LiteralFloat:
		// Create 64-bit float
		var val float64
		fmt.Sscanf(lit.Value, "%f", &val)
		return llvm.ConstFloat(g.context.DoubleType(), val)

	case ast.LiteralBool:
		// Create boolean
		var val uint64
		if lit.Value == "true" {
			val = 1
		}
		return llvm.ConstInt(g.context.Int1Type(), val, false)

	case ast.LiteralString:
		// Check if we've already created this string literal
		if val, exists := g.stringLiterals[lit.Value]; exists {
			return val
		}

		// Create a global string constant
		strVal := g.context.ConstString(lit.Value, false)
		strPtr := llvm.AddGlobal(g.module, strVal.Type(), "")
		strPtr.SetLinkage(llvm.PrivateLinkage)
		strPtr.SetGlobalConstant(true)
		strPtr.SetInitializer(strVal)
		strPtr.SetUnnamedAddr(true)

		// Convert to i8* pointer
		zero := llvm.ConstInt(g.context.Int32Type(), 0, false)
		indices := []llvm.Value{zero, zero}
		strPtrCast := llvm.ConstGEP(strVal.Type(), strPtr, indices)

		g.stringLiterals[lit.Value] = strPtrCast
		return strPtrCast

	case ast.LiteralNull:
		// Create a null pointer
		return llvm.ConstNull(llvm.PointerType(g.context.Int8Type(), 0))

	default:
		log.Printf("Warning: Unsupported literal kind: %d", lit.Kind)
		return llvm.ConstNull(g.context.Int32Type())
	}
}

// GenerateBinaryExpr handles binary expressions (Phase 1)
func (g *CodeGenerator) GenerateBinaryExpr(expr *ast.BinaryExpr) (llvm.Value, error) {
	// Generate operands
	left, err := g.GenerateExpr(expr.Left)
	if err != nil {
		return llvm.Value{}, err
	}

	right, err := g.GenerateExpr(expr.Right)
	if err != nil {
		return llvm.Value{}, err
	}

	// Handle arithmetic, comparison and logical operations
	isFloatingPoint := left.Type().TypeKind() == llvm.FloatTypeKind ||
		left.Type().TypeKind() == llvm.DoubleTypeKind

	if isFloatingPoint {
		// Floating-point operations
		switch expr.Operator {
		case ast.OperatorAdd:
			return g.builder.CreateFAdd(left, right, "fadd"), nil
		case ast.OperatorSub:
			return g.builder.CreateFSub(left, right, "fsub"), nil
		case ast.OperatorMul:
			return g.builder.CreateFMul(left, right, "fmul"), nil
		case ast.OperatorDiv:
			return g.builder.CreateFDiv(left, right, "fdiv"), nil
		case ast.OperatorMod:
			return g.builder.CreateFRem(left, right, "frem"), nil
		case ast.OperatorEq:
			return g.builder.CreateFCmp(llvm.FloatOEQ, left, right, "fcmp_eq"), nil
		case ast.OperatorNe:
			return g.builder.CreateFCmp(llvm.FloatONE, left, right, "fcmp_ne"), nil
		case ast.OperatorLt:
			return g.builder.CreateFCmp(llvm.FloatOLT, left, right, "fcmp_lt"), nil
		case ast.OperatorLe:
			return g.builder.CreateFCmp(llvm.FloatOLE, left, right, "fcmp_le"), nil
		case ast.OperatorGt:
			return g.builder.CreateFCmp(llvm.FloatOGT, left, right, "fcmp_gt"), nil
		case ast.OperatorGe:
			return g.builder.CreateFCmp(llvm.FloatOGE, left, right, "fcmp_ge"), nil
		default:
			return llvm.Value{}, fmt.Errorf("unsupported floating-point operator: %s", expr.Operator)
		}
	} else {
		// Integer operations
		switch expr.Operator {
		case ast.OperatorAdd:
			return g.builder.CreateAdd(left, right, "add"), nil
		case ast.OperatorSub:
			return g.builder.CreateSub(left, right, "sub"), nil
		case ast.OperatorMul:
			return g.builder.CreateMul(left, right, "mul"), nil
		case ast.OperatorDiv:
			return g.builder.CreateSDiv(left, right, "div"), nil
		case ast.OperatorMod:
			return g.builder.CreateSRem(left, right, "rem"), nil
		case ast.OperatorEq:
			return g.builder.CreateICmp(llvm.IntEQ, left, right, "icmp_eq"), nil
		case ast.OperatorNe:
			return g.builder.CreateICmp(llvm.IntNE, left, right, "icmp_ne"), nil
		case ast.OperatorLt:
			return g.builder.CreateICmp(llvm.IntSLT, left, right, "icmp_lt"), nil
		case ast.OperatorLe:
			return g.builder.CreateICmp(llvm.IntSLE, left, right, "icmp_le"), nil
		case ast.OperatorGt:
			return g.builder.CreateICmp(llvm.IntSGT, left, right, "icmp_gt"), nil
		case ast.OperatorGe:
			return g.builder.CreateICmp(llvm.IntSGE, left, right, "icmp_ge"), nil
		case ast.OperatorAnd:
			return g.createLogicalAnd(left, right), nil
		case ast.OperatorOr:
			return g.createLogicalOr(left, right), nil
		default:
			return llvm.Value{}, fmt.Errorf("unsupported integer operator: %s", expr.Operator)
		}
	}
}

// GenerateUnaryExpr handles unary expressions (Phase 1)
func (g *CodeGenerator) GenerateUnaryExpr(expr *ast.UnaryExpr) (llvm.Value, error) {
	// Generate operand
	operand, err := g.GenerateExpr(expr.Operand)
	if err != nil {
		return llvm.Value{}, err
	}

	// Handle different unary operators
	switch expr.Operator {
	case ast.OperatorUnaryPlus:
		// Unary plus is a no-op
		return operand, nil

	case ast.OperatorUnaryMinus:
		// Negation depends on type
		if operand.Type().TypeKind() == llvm.FloatTypeKind ||
			operand.Type().TypeKind() == llvm.DoubleTypeKind {
			return g.builder.CreateFNeg(operand, "fneg"), nil
		}
		return g.builder.CreateNeg(operand, "neg"), nil

	case ast.OperatorNot:
		// Logical not - compare with zero or use not for boolean
		if operand.Type().C == g.context.Int1Type().C {
			return g.builder.CreateNot(operand, "not"), nil
		}
		return g.builder.CreateICmp(llvm.IntEQ, operand,
			llvm.ConstInt(operand.Type(), 0, false), "lnot"), nil

	default:
		return llvm.Value{}, fmt.Errorf("unsupported unary operator: %s", expr.Operator)
	}
}

// GenerateCallExpr handles function calls (Phase 1)
func (g *CodeGenerator) GenerateCallExpr(expr *ast.CallExpr) (llvm.Value, error) {
	// Get the function name - for Phase 1, we expect identifiers only
	funcIdent, ok := expr.Function.(*ast.Identifier)
	if !ok {
		return llvm.Value{}, fmt.Errorf("function calls must be identifiers in Phase 1")
	}

	// Look up the function in the module
	function := g.module.NamedFunction(funcIdent.Name)
	if function.IsNil() {
		return llvm.Value{}, fmt.Errorf("undefined function: %s", funcIdent.Name)
	}
	
	log.Printf("Found function %s in module", funcIdent.Name)

	// Generate arguments with validation
	var args []llvm.Value
	for i, argExpr := range expr.Arguments {
		argValue, err := g.GenerateExpr(argExpr)
		if err != nil {
			return llvm.Value{}, fmt.Errorf("error generating function argument %d: %w", i, err)
		}
		if argValue.IsNil() {
			return llvm.Value{}, fmt.Errorf("argument %d is nil", i)
		}
		log.Printf("Argument %d generated successfully", i)
		args = append(args, argValue)
	}

	// In older LLVM versions, we need to be more careful about function calls
	// Let's try using the GlobalValue approach and get function type differently
	
	// Get the function type by using LLVM's type system
	functionPtrType := function.Type()
	
	// Most functions in LLVM are pointer types to function types  
	var funcType llvm.Type
	if functionPtrType.TypeKind() == llvm.PointerTypeKind {
		funcType = functionPtrType.ElementType()
	} else {
		funcType = functionPtrType
	}
	
	// Validate we have a function type
	if funcType.IsNil() {
		return llvm.Value{}, fmt.Errorf("cannot determine function type for %s", funcIdent.Name)
	}
	
	log.Printf("About to call function %s", funcIdent.Name)
	
	// Use CreateCall - this should work with the tinygo LLVM bindings
	result := g.builder.CreateCall(funcType, function, args, "")
	
	log.Printf("Function call completed successfully")
	return result, nil
}

// createLogicalAnd implements short-circuit evaluation for logical AND (Phase 1)
func (g *CodeGenerator) createLogicalAnd(left, right llvm.Value) llvm.Value {
	// Create basic blocks
	rightBlock := llvm.AddBasicBlock(g.currentFunc, "and.right")
	mergeBlock := llvm.AddBasicBlock(g.currentFunc, "and.end")

	// Convert left to boolean if needed
	if left.Type().C != g.context.Int1Type().C {
		left = g.builder.CreateICmp(llvm.IntNE, left,
			llvm.ConstInt(left.Type(), 0, false), "left.bool")
	}

	// Create an alloca for the result
	result := g.builder.CreateAlloca(g.context.Int1Type(), "and.result")

	// If left is false, short-circuit with false
	g.builder.CreateStore(left, result)
	g.builder.CreateCondBr(left, rightBlock, mergeBlock)

	// Evaluate right side only if left is true
	g.builder.SetInsertPointAtEnd(rightBlock)
	if right.Type().C != g.context.Int1Type().C {
		right = g.builder.CreateICmp(llvm.IntNE, right,
			llvm.ConstInt(right.Type(), 0, false), "right.bool")
	}
	g.builder.CreateStore(right, result)
	g.builder.CreateBr(mergeBlock)

	// Continue with merge block
	g.builder.SetInsertPointAtEnd(mergeBlock)
	return g.builder.CreateLoad(result.Type().ElementType(), result, "and.result.load")
}

// createLogicalOr implements short-circuit evaluation for logical OR (Phase 1)
func (g *CodeGenerator) createLogicalOr(left, right llvm.Value) llvm.Value {
	// Create basic blocks
	rightBlock := llvm.AddBasicBlock(g.currentFunc, "or.right")
	mergeBlock := llvm.AddBasicBlock(g.currentFunc, "or.end")

	// Convert left to boolean if needed
	if left.Type().C != g.context.Int1Type().C {
		left = g.builder.CreateICmp(llvm.IntNE, left,
			llvm.ConstInt(left.Type(), 0, false), "left.bool")
	}

	// Create an alloca for the result
	result := g.builder.CreateAlloca(g.context.Int1Type(), "or.result")

	// If left is true, short-circuit with true
	g.builder.CreateStore(left, result)
	g.builder.CreateCondBr(left, mergeBlock, rightBlock)

	// Evaluate right side only if left is false
	g.builder.SetInsertPointAtEnd(rightBlock)
	if right.Type().C != g.context.Int1Type().C {
		right = g.builder.CreateICmp(llvm.IntNE, right,
			llvm.ConstInt(right.Type(), 0, false), "right.bool")
	}
	g.builder.CreateStore(right, result)
	g.builder.CreateBr(mergeBlock)

	// Continue with merge block
	g.builder.SetInsertPointAtEnd(mergeBlock)
	return g.builder.CreateLoad(result.Type().ElementType(), result, "or.result.load")
}

// Dispose releases LLVM resources
func (g *CodeGenerator) Dispose() {
	g.builder.Dispose()
	g.module.Dispose()
	g.context.Dispose()
}
