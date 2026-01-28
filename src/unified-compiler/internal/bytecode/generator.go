package bytecode

import (
"fmt"
"unified-compiler/internal/ast"
"unified-compiler/internal/semantic"
)

// LoopContext tracks loop information for break/continue
type LoopContext struct {
startPos      int    // Position to jump to for continue
breakJumps    []int  // Positions of break jumps to patch
continueJumps []int  // Positions of continue jumps to patch
label         string // Optional loop label
}

// MonomorphizedFunction tracks a specialized version of a generic function
type MonomorphizedFunction struct {
BaseName   string
TypeArgs   []ast.Type
MangledName string
Generated  bool
}

// MonomorphizedStruct tracks a specialized version of a generic struct
type MonomorphizedStruct struct {
BaseName    string
TypeArgs    []ast.Type
MangledName string
Generated   bool
}

// MonomorphizedEnum tracks a specialized version of a generic enum
type MonomorphizedEnum struct {
BaseName    string
TypeArgs    []ast.Type
MangledName string
Generated   bool
}

// Generator converts AST to bytecode
type Generator struct {
bytecode      *Bytecode
localVars     map[string]int // Variable name -> stack index
localVarCount int            // Number of local variables
loopStack     []LoopContext  // Stack of nested loop contexts
structTypes   map[string]*StructTypeInfo // Struct type information
enumTypes     map[string]*EnumTypeInfo   // Enum type information
	
// Generic programming support
genericContext *semantic.GenericContext // Generic type parameter context for current function
	
// Monomorphization tracking
// Maps mangled name -> MonomorphizedFunction. Created during Generate() when generic
// functions are instantiated with concrete types. Used to avoid duplicate generation.
monomorphized map[string]*MonomorphizedFunction
	
// Generic function templates
// Maps function name -> FunctionDecl. Populated in first pass of Generate() to store
// generic function templates that will be monomorphized on-demand during call generation.
genericFunctions map[string]*ast.FunctionDecl

// Generic struct templates
// Maps struct name -> StructDecl. Populated in first pass to store generic struct
// templates that will be monomorphized on-demand during instantiation.
genericStructs map[string]*ast.StructDecl

// Monomorphized struct tracking
// Maps mangled name -> MonomorphizedStruct to avoid duplicate generation.
monomorphizedStructs map[string]*MonomorphizedStruct

// Generic enum templates
// Maps enum name -> EnumDecl. Populated in first pass to store generic enum
// templates that will be monomorphized on-demand during instantiation.
genericEnums map[string]*ast.EnumDecl

// Monomorphized enum tracking
// Maps mangled name -> MonomorphizedEnum to avoid duplicate generation.
monomorphizedEnums map[string]*MonomorphizedEnum
}

// StructTypeInfo holds metadata about a struct type
type StructTypeInfo struct {
Name    string
Fields  map[string]int // Field name -> index
Methods map[string]*ast.FunctionDecl
}

// EnumTypeInfo holds metadata about an enum type
type EnumTypeInfo struct {
Name     string
Variants map[string]*VariantInfo // Variant name -> info
}

// VariantInfo holds metadata about an enum variant
type VariantInfo struct {
Name  string
Tag   int
Arity int // Number of data fields
}

// NewGenerator creates a new bytecode generator
func NewGenerator() *Generator {
return &Generator{
bytecode:    NewBytecode(),
localVars:   make(map[string]int),
structTypes: make(map[string]*StructTypeInfo),
enumTypes:   make(map[string]*EnumTypeInfo),
genericContext: semantic.NewGenericContext(),
monomorphized: make(map[string]*MonomorphizedFunction),
genericFunctions: make(map[string]*ast.FunctionDecl),
genericStructs: make(map[string]*ast.StructDecl),
monomorphizedStructs: make(map[string]*MonomorphizedStruct),
genericEnums: make(map[string]*ast.EnumDecl),
monomorphizedEnums: make(map[string]*MonomorphizedEnum),
}
}

// Generate converts a program AST to bytecode
func (g *Generator) Generate(program *ast.Program) (*Bytecode, error) {
// First pass: Register struct and enum types
for _, item := range program.Items {
switch item := item.(type) {
case *ast.StructDecl:
// Store generic structs for monomorphization
if len(item.GenericParams) > 0 {
g.genericStructs[item.Name] = item
} else {
// Register non-generic structs immediately
if err := g.registerStructType(item); err != nil {
return nil, fmt.Errorf("error registering struct %s: %w", item.Name, err)
}
}
case *ast.EnumDecl:
// Store generic enums for monomorphization
if len(item.GenericParams) > 0 {
g.genericEnums[item.Name] = item
} else {
// Register non-generic enums immediately
if err := g.registerEnumType(item); err != nil {
return nil, fmt.Errorf("error registering enum %s: %w", item.Name, err)
}
}
case *ast.FunctionDecl:
// Store generic functions for monomorphization
if len(item.GenericParams) > 0 {
g.genericFunctions[item.Name] = item
}
}
}

// Second pass: Generate code for all top-level items
for _, item := range program.Items {
switch item := item.(type) {
case *ast.FunctionDecl:
// Skip generic functions - they will be monomorphized on demand
if len(item.GenericParams) > 0 {
continue
}
if err := g.generateFunction(item); err != nil {
return nil, fmt.Errorf("error generating function %s: %w", item.Name, err)
}
case *ast.StructDecl:
// Struct methods are handled separately
if err := g.generateStructMethods(item); err != nil {
return nil, fmt.Errorf("error generating methods for struct %s: %w", item.Name, err)
}
case *ast.EnumDecl:
// Enums don't have methods (yet), so nothing to generate
case *ast.ConstantDecl:
// Constants are handled at compile time
// For now, skip them in bytecode generation
default:
return nil, fmt.Errorf("unsupported top-level item: %T", item)
}
}

// Add final HALT instruction
g.bytecode.AddInstruction(OpHalt, 0)

return g.bytecode, nil
}

// generateFunction generates bytecode for a function
func (g *Generator) generateFunction(fn *ast.FunctionDecl) error {
// Record function entry point
entryPoint := g.bytecode.CurrentPosition()
g.bytecode.AddFunction(fn.Name, entryPoint)

// Reset local variables for this function
g.localVars = make(map[string]int)
g.localVarCount = 0

// Add parameters to local variables
for i, param := range fn.Parameters {
g.localVars[param.Name] = i
g.localVarCount++
}

// Generate function body
if fn.Body != nil {
if err := g.generateBlock(fn.Body); err != nil {
return err
}

// If function has trailing expression, generate it and return
if fn.Body.Expression != nil {
if err := g.generateExpression(fn.Body.Expression); err != nil {
return err
}
g.bytecode.AddInstruction(OpReturnValue, 0)
} else {
// Otherwise, add implicit return
g.bytecode.AddInstruction(OpReturn, 0)
}
} else {
// Empty function body
g.bytecode.AddInstruction(OpReturn, 0)
}

return nil
}

// generateBlock generates bytecode for a block statement
func (g *Generator) generateBlock(block *ast.Block) error {
for _, stmt := range block.Statements {
if err := g.generateStatement(stmt); err != nil {
return err
}
}
return nil
}

// generateBlockExpression generates bytecode for a block used as an expression
func (g *Generator) generateBlockExpression(block *ast.Block) error {
// Generate all statements in the block
for _, stmt := range block.Statements {
if err := g.generateStatement(stmt); err != nil {
return err
}
}

// If block has a trailing expression, generate it
// The value will remain on the stack as the block's result
if block.Expression != nil {
if err := g.generateExpression(block.Expression); err != nil {
return err
}
} else {
// If no expression, push null/zero as the block value
idx := g.bytecode.AddConstant(NewIntValue(0))
g.bytecode.AddInstruction(OpPush, int64(idx))
}

return nil
}

// generateStatement generates bytecode for a statement
func (g *Generator) generateStatement(stmt ast.Statement) error {
	switch stmt := stmt.(type) {
	case *ast.LetStatement:
		return g.generateLetStatement(stmt)
	case *ast.VarStatement:
		return g.generateVarStatement(stmt)
	case *ast.AssignmentStatement:
		return g.generateAssignmentStatement(stmt)
	case *ast.ReturnStatement:
		return g.generateReturnStatement(stmt)
	case *ast.ExprStatement:
		// Generate expression and pop result (expression statements don't use the value)
		if err := g.generateExpression(stmt.Expression); err != nil {
			return err
		}
		g.bytecode.AddInstruction(OpPop, 0)
		return nil
	case *ast.IfStatement:
		return g.generateIfStatement(stmt)
	case *ast.WhileStatement:
		return g.generateWhileStatement(stmt)
	case *ast.ForStatement:
		return g.generateForStatement(stmt)
	case *ast.LoopStatement:
		return g.generateLoopStatement(stmt)
	case *ast.BreakStatement:
		return g.generateBreakStatement(stmt)
	case *ast.ContinueStatement:
		return g.generateContinueStatement(stmt)
	default:
		return fmt.Errorf("unsupported statement type: %T", stmt)
	}
}

func (g *Generator) generateLetStatement(stmt *ast.LetStatement) error {
// Generate the initializer expression
if stmt.Value != nil {
if err := g.generateExpression(stmt.Value); err != nil {
return err
}
} else {
// Default initialize to 0
constIdx := g.bytecode.AddConstant(NewIntValue(0))
g.bytecode.AddInstruction(OpPush, int64(constIdx))
}

// Allocate local variable
varIdx := g.localVarCount
g.localVars[stmt.Name] = varIdx
g.localVarCount++

// Store the value
g.bytecode.AddInstruction(OpStoreLocal, int64(varIdx))

return nil
}

// generateVarStatement generates bytecode for var statement
func (g *Generator) generateVarStatement(stmt *ast.VarStatement) error {
// Similar to let statement for Phase 1
if stmt.Value != nil {
if err := g.generateExpression(stmt.Value); err != nil {
return err
}
} else {
constIdx := g.bytecode.AddConstant(NewIntValue(0))
g.bytecode.AddInstruction(OpPush, int64(constIdx))
}

varIdx := g.localVarCount
g.localVars[stmt.Name] = varIdx
g.localVarCount++

g.bytecode.AddInstruction(OpStoreLocal, int64(varIdx))

return nil
}

// generateReturnStatement generates bytecode for return statement
func (g *Generator) generateReturnStatement(stmt *ast.ReturnStatement) error {
if stmt.Value != nil {
if err := g.generateExpression(stmt.Value); err != nil {
return err
}
g.bytecode.AddInstruction(OpReturnValue, 0)
} else {
g.bytecode.AddInstruction(OpReturn, 0)
}
return nil
}

// generateIfStatement generates bytecode for if statement
func (g *Generator) generateIfStatement(stmt *ast.IfStatement) error {
// Generate condition
if err := g.generateExpression(stmt.Condition); err != nil {
return err
}

// Jump if false
jumpIfFalse := g.bytecode.CurrentPosition()
g.bytecode.AddInstruction(OpJumpIfFalse, 0) // Placeholder

// Generate then block
if err := g.generateBlock(stmt.ThenBlock); err != nil {
return err
}

if stmt.ElseBlock != nil {
// Jump over else block
jumpOverElse := g.bytecode.CurrentPosition()
g.bytecode.AddInstruction(OpJump, 0) // Placeholder

// Patch jump if false to point to else block
elseStart := g.bytecode.CurrentPosition()
g.bytecode.PatchJump(jumpIfFalse, elseStart)

// Generate else block
if err := g.generateBlock(stmt.ElseBlock); err != nil {
return err
}

// Patch jump over else to point past else block
afterElse := g.bytecode.CurrentPosition()
g.bytecode.PatchJump(jumpOverElse, afterElse)
} else {
// No else block, patch jump if false to point past then block
afterThen := g.bytecode.CurrentPosition()
g.bytecode.PatchJump(jumpIfFalse, afterThen)
}

return nil
}

// generateExpression generates bytecode for an expression
func (g *Generator) generateExpression(expr ast.Expression) error {
	switch expr := expr.(type) {
	case *ast.Literal:
		return g.generateLiteral(expr)
	case *ast.Identifier:
		return g.generateIdentifier(expr)
	case *ast.BinaryExpr:
		return g.generateBinaryExpr(expr)
	case *ast.UnaryExpr:
		return g.generateUnaryExpr(expr)
	case *ast.CallExpr:
		return g.generateCallExpr(expr)
	case *ast.StructExpr:
		return g.generateStructExpr(expr)
	case *ast.NewExpr:
		return g.generateNewExpr(expr)
	case *ast.FieldAccessExpr:
		return g.generateFieldAccessExpr(expr)
	case *ast.MethodCallExpr:
		return g.generateMethodCallExpr(expr)
	case *ast.EnumConstructorExpr:
		return g.generateEnumConstructorExpr(expr)
	case *ast.MatchExpr:
		return g.generateMatchExpr(expr)
	case *ast.TryExpr:
		return g.generateTryExpr(expr)
	case *ast.Block:
		return g.generateBlockExpression(expr)
	case *ast.ListExpr:
		return g.generateListExpr(expr)
	case *ast.IndexExpr:
		return g.generateIndexExpr(expr)
	default:
		return fmt.Errorf("unsupported expression type: %T", expr)
	}
}

// generateLiteral generates bytecode for a literal
func (g *Generator) generateLiteral(lit *ast.Literal) error {
var val Value
switch lit.Kind {
case ast.LiteralInt:
var i int64
fmt.Sscanf(lit.Value, "%d", &i)
val = NewIntValue(i)
case ast.LiteralFloat:
var f float64
fmt.Sscanf(lit.Value, "%f", &f)
val = NewFloatValue(f)
case ast.LiteralBool:
val = NewBoolValue(lit.Value == "true")
case ast.LiteralString:
// Strip quotes from string literal
strValue := lit.Value
if len(strValue) >= 2 && strValue[0] == '"' && strValue[len(strValue)-1] == '"' {
strValue = strValue[1 : len(strValue)-1]
}
// Handle escape sequences
strValue = g.unescapeString(strValue)
val = NewStringValue(strValue)
case ast.LiteralNull:
val = NewNullValue()
default:
return fmt.Errorf("unsupported literal kind: %d", lit.Kind)
}

constIdx := g.bytecode.AddConstant(val)
g.bytecode.AddInstruction(OpPush, int64(constIdx))
return nil
}

// generateIdentifier generates bytecode for an identifier
func (g *Generator) generateIdentifier(ident *ast.Identifier) error {
// Look up variable
if idx, ok := g.localVars[ident.Name]; ok {
g.bytecode.AddInstruction(OpLoadLocal, int64(idx))
return nil
}

return fmt.Errorf("undefined variable: %s", ident.Name)
}

// generateBinaryExpr generates bytecode for a binary expression
func (g *Generator) generateBinaryExpr(expr *ast.BinaryExpr) error {
// Handle assignment specially
if expr.Operator == ast.OperatorAssign {
return g.generateAssignment(expr)
}

// Range expressions are only valid in for loops, not as standalone expressions
if expr.Operator == ast.OperatorRange || expr.Operator == ast.OperatorRangeIncl {
return fmt.Errorf("range expressions can only be used in for loops")
}

// Generate left operand
if err := g.generateExpression(expr.Left); err != nil {
return err
}

// Generate right operand
if err := g.generateExpression(expr.Right); err != nil {
return err
}

// Generate operator
switch expr.Operator {
case ast.OperatorAdd:
g.bytecode.AddInstruction(OpAdd, 0)
case ast.OperatorSub:
g.bytecode.AddInstruction(OpSub, 0)
case ast.OperatorMul:
g.bytecode.AddInstruction(OpMul, 0)
case ast.OperatorDiv:
g.bytecode.AddInstruction(OpDiv, 0)
case ast.OperatorMod:
g.bytecode.AddInstruction(OpMod, 0)
case ast.OperatorEq:
g.bytecode.AddInstruction(OpEq, 0)
case ast.OperatorNe:
g.bytecode.AddInstruction(OpNe, 0)
case ast.OperatorLt:
g.bytecode.AddInstruction(OpLt, 0)
case ast.OperatorLe:
g.bytecode.AddInstruction(OpLe, 0)
case ast.OperatorGt:
g.bytecode.AddInstruction(OpGt, 0)
case ast.OperatorGe:
g.bytecode.AddInstruction(OpGe, 0)
case ast.OperatorAnd:
g.bytecode.AddInstruction(OpAnd, 0)
case ast.OperatorOr:
g.bytecode.AddInstruction(OpOr, 0)
case ast.OperatorBitAnd:
g.bytecode.AddInstruction(OpBitAnd, 0)
case ast.OperatorBitOr:
g.bytecode.AddInstruction(OpBitOr, 0)
case ast.OperatorBitXor:
g.bytecode.AddInstruction(OpBitXor, 0)
case ast.OperatorLShift:
g.bytecode.AddInstruction(OpLShift, 0)
case ast.OperatorRShift:
g.bytecode.AddInstruction(OpRShift, 0)
default:
return fmt.Errorf("unsupported binary operator: %s", expr.Operator)
}

return nil
}

// generateAssignment generates bytecode for variable or array element assignment
func (g *Generator) generateAssignment(expr *ast.BinaryExpr) error {
	// Check if assignment target is a variable
	if ident, ok := expr.Left.(*ast.Identifier); ok {
		// Generate the value being assigned
		if err := g.generateExpression(expr.Right); err != nil {
			return err
		}

		// Look up variable
		varIdx, ok := g.localVars[ident.Name]
		if !ok {
			return fmt.Errorf("undefined variable: %s", ident.Name)
		}

		// Duplicate the value on the stack (assignment is an expression that returns the assigned value)
		g.bytecode.AddInstruction(OpDup, 0)

		// Store to variable
		g.bytecode.AddInstruction(OpStoreLocal, int64(varIdx))

		return nil
	}

	// Check if assignment target is an array index
	if indexExpr, ok := expr.Left.(*ast.IndexExpr); ok {
		// Generate the array object
		if err := g.generateExpression(indexExpr.Object); err != nil {
			return err
		}

		// Generate the index
		if err := g.generateExpression(indexExpr.Index); err != nil {
			return err
		}

		// Generate the value being assigned
		if err := g.generateExpression(expr.Right); err != nil {
			return err
		}

		// Store to array
		g.bytecode.AddInstruction(OpStoreArray, 0)

		return nil
	}

	return fmt.Errorf("assignment target must be a variable or array index")
}

// generateUnaryExpr generates bytecode for a unary expression
func (g *Generator) generateUnaryExpr(expr *ast.UnaryExpr) error {
// Generate operand
if err := g.generateExpression(expr.Operand); err != nil {
return err
}

// Generate operator
switch expr.Operator {
case ast.OperatorUnaryMinus:
g.bytecode.AddInstruction(OpNeg, 0)
case ast.OperatorNot:
g.bytecode.AddInstruction(OpNot, 0)
case ast.OperatorBitNot:
g.bytecode.AddInstruction(OpBitNot, 0)
case ast.OperatorUnaryPlus:
// No-op
default:
return fmt.Errorf("unsupported unary operator: %s", expr.Operator)
}

return nil
}

// generateCallExpr generates bytecode for a function call
func (g *Generator) generateCallExpr(expr *ast.CallExpr) error {
// For Phase 1, only support direct function calls
funcIdent, ok := expr.Function.(*ast.Identifier)
if !ok {
return fmt.Errorf("only direct function calls are supported in Phase 1")
}

// Check if this is a call to a generic function
if genericFn, ok := g.genericFunctions[funcIdent.Name]; ok {
// This is a generic function call - we need to monomorphize it
var typeArgs []ast.Type
var err error
		
// Check if explicit type arguments are provided
if len(expr.TypeArgs) > 0 {
typeArgs = expr.TypeArgs
} else {
// Infer type arguments from call arguments
typeArgs, err = g.inferCallTypeArguments(genericFn, expr.Arguments)
if err != nil {
return fmt.Errorf("type inference failed for %s: %w", funcIdent.Name, err)
}
}
		
// Monomorphize the function
mangledName, err := g.monomorphizeFunction(funcIdent.Name, typeArgs)
if err != nil {
return fmt.Errorf("monomorphization failed for %s: %w", funcIdent.Name, err)
}
		
// Generate arguments (pushed in order)
for _, arg := range expr.Arguments {
if err := g.generateExpression(arg); err != nil {
return err
}
}
		
// Call the monomorphized version
funcIdx, ok := g.bytecode.Functions[mangledName]
if !ok {
return fmt.Errorf("monomorphized function not found: %s", mangledName)
}
		
g.bytecode.AddInstructionWithArgCount(OpCall, int64(funcIdx), len(expr.Arguments))
return nil
}

// Generate arguments (pushed in order)
for _, arg := range expr.Arguments {
if err := g.generateExpression(arg); err != nil {
return err
}
}

// Look up function
funcIdx, ok := g.bytecode.Functions[funcIdent.Name]
if !ok {
return fmt.Errorf("undefined function: %s", funcIdent.Name)
}

// Generate call with argument count
g.bytecode.AddInstructionWithArgCount(OpCall, int64(funcIdx), len(expr.Arguments))

return nil
}

// generateWhileStatement generates bytecode for while loop
func (g *Generator) generateWhileStatement(stmt *ast.WhileStatement) error {
// Mark loop start
loopStart := g.bytecode.CurrentPosition()

// Create loop context
loopCtx := LoopContext{
startPos:      loopStart,
breakJumps:    []int{},
continueJumps: []int{},
label:         stmt.Label,
}
g.loopStack = append(g.loopStack, loopCtx)

// Generate condition
if err := g.generateExpression(stmt.Condition); err != nil {
return err
}

// Jump to end if condition is false
jumpIfFalse := g.bytecode.CurrentPosition()
g.bytecode.AddInstruction(OpJumpIfFalse, 0) // Placeholder

// Generate loop body
if err := g.generateBlock(stmt.Body); err != nil {
return err
}

// Jump back to loop start
g.bytecode.AddInstruction(OpJump, int64(loopStart))

// Patch break jumps to point here (after loop)
afterLoop := g.bytecode.CurrentPosition()
g.bytecode.PatchJump(jumpIfFalse, afterLoop)

// Pop loop context and patch break/continue jumps
loopCtx = g.loopStack[len(g.loopStack)-1]
g.loopStack = g.loopStack[:len(g.loopStack)-1]

for _, breakPos := range loopCtx.breakJumps {
g.bytecode.PatchJump(breakPos, afterLoop)
}
for _, continuePos := range loopCtx.continueJumps {
g.bytecode.PatchJump(continuePos, loopStart)
}

return nil
}

// generateForStatement generates bytecode for for loop
func (g *Generator) generateForStatement(stmt *ast.ForStatement) error {
// For Phase 2, we'll support range expressions
// Expect iterable to be a binary expression with RANGE or RANGE_INCL operator

// Check if iterable is a binary expression with RANGE or RANGE_INCL operator
var rangeStart, rangeEnd ast.Expression
var isInclusive bool

if binaryExpr, ok := stmt.Iterable.(*ast.BinaryExpr); ok {
// Check for range operators
if binaryExpr.Operator == ast.OperatorRange {
rangeStart = binaryExpr.Left
rangeEnd = binaryExpr.Right
isInclusive = false
} else if binaryExpr.Operator == ast.OperatorRangeIncl {
rangeStart = binaryExpr.Left
rangeEnd = binaryExpr.Right
isInclusive = true
} else {
return fmt.Errorf("for loop iterable must be a range expression (.. or ..=)")
}
} else {
return fmt.Errorf("for loops currently only support range expressions")
}

// Allocate iterator variable
iterIdx := g.localVarCount
g.localVars[stmt.Variable] = iterIdx
g.localVarCount++

// Initialize iterator with range start
if err := g.generateExpression(rangeStart); err != nil {
return err
}
g.bytecode.AddInstruction(OpStoreLocal, int64(iterIdx))

// Store range end in a temporary variable
endIdx := g.localVarCount
g.localVarCount++
if err := g.generateExpression(rangeEnd); err != nil {
return err
}
g.bytecode.AddInstruction(OpStoreLocal, int64(endIdx))

// Loop start
loopStart := g.bytecode.CurrentPosition()

// Create loop context
loopCtx := LoopContext{
startPos:      loopStart,
breakJumps:    []int{},
continueJumps: []int{},
label:         stmt.Label,
}
g.loopStack = append(g.loopStack, loopCtx)

// Check condition: iterator < end (or <= for inclusive)
g.bytecode.AddInstruction(OpLoadLocal, int64(iterIdx))
g.bytecode.AddInstruction(OpLoadLocal, int64(endIdx))
if isInclusive {
g.bytecode.AddInstruction(OpLe, 0) // iterator <= end
} else {
g.bytecode.AddInstruction(OpLt, 0) // iterator < end
}

// Jump to end if condition is false
jumpIfFalse := g.bytecode.CurrentPosition()
g.bytecode.AddInstruction(OpJumpIfFalse, 0) // Placeholder

// Generate loop body
if err := g.generateBlock(stmt.Body); err != nil {
return err
}

// Increment iterator
incrementPos := g.bytecode.CurrentPosition()
g.bytecode.AddInstruction(OpLoadLocal, int64(iterIdx))
idx1 := g.bytecode.AddConstant(NewIntValue(1))
g.bytecode.AddInstruction(OpPush, int64(idx1))
g.bytecode.AddInstruction(OpAdd, 0)
g.bytecode.AddInstruction(OpStoreLocal, int64(iterIdx))

// Jump back to loop start
g.bytecode.AddInstruction(OpJump, int64(loopStart))

// Patch jumps
afterLoop := g.bytecode.CurrentPosition()
g.bytecode.PatchJump(jumpIfFalse, afterLoop)

// Pop loop context and patch break/continue jumps
loopCtx = g.loopStack[len(g.loopStack)-1]
g.loopStack = g.loopStack[:len(g.loopStack)-1]

for _, breakPos := range loopCtx.breakJumps {
g.bytecode.PatchJump(breakPos, afterLoop)
}
for _, continuePos := range loopCtx.continueJumps {
g.bytecode.PatchJump(continuePos, incrementPos)
}

return nil
}

// generateLoopStatement generates bytecode for infinite loop
func (g *Generator) generateLoopStatement(stmt *ast.LoopStatement) error {
// Mark loop start
loopStart := g.bytecode.CurrentPosition()

// Create loop context
loopCtx := LoopContext{
startPos:      loopStart,
breakJumps:    []int{},
continueJumps: []int{},
label:         stmt.Label,
}
g.loopStack = append(g.loopStack, loopCtx)

// Generate loop body
if err := g.generateBlock(stmt.Body); err != nil {
return err
}

// Jump back to loop start
g.bytecode.AddInstruction(OpJump, int64(loopStart))

// After loop (only reachable via break)
afterLoop := g.bytecode.CurrentPosition()

// Pop loop context and patch break/continue jumps
loopCtx = g.loopStack[len(g.loopStack)-1]
g.loopStack = g.loopStack[:len(g.loopStack)-1]

for _, breakPos := range loopCtx.breakJumps {
g.bytecode.PatchJump(breakPos, afterLoop)
}
for _, continuePos := range loopCtx.continueJumps {
g.bytecode.PatchJump(continuePos, loopStart)
}

return nil
}

// generateBreakStatement generates bytecode for break statement
func (g *Generator) generateBreakStatement(stmt *ast.BreakStatement) error {
if len(g.loopStack) == 0 {
return fmt.Errorf("break statement outside of loop")
}

// Find the target loop context
var targetIdx int
if stmt.Label != "" {
// Find labeled loop
found := false
for i := len(g.loopStack) - 1; i >= 0; i-- {
if g.loopStack[i].label == stmt.Label {
targetIdx = i
found = true
break
}
}
if !found {
return fmt.Errorf("break label not found: %s", stmt.Label)
}
} else {
// Break from innermost loop
targetIdx = len(g.loopStack) - 1
}

// Add a jump instruction (will be patched later)
jumpPos := g.bytecode.CurrentPosition()
g.bytecode.AddInstruction(OpJump, 0) // Placeholder

// Record this jump position in the loop context
g.loopStack[targetIdx].breakJumps = append(g.loopStack[targetIdx].breakJumps, jumpPos)

return nil
}

// generateContinueStatement generates bytecode for continue statement
func (g *Generator) generateContinueStatement(stmt *ast.ContinueStatement) error {
if len(g.loopStack) == 0 {
return fmt.Errorf("continue statement outside of loop")
}

// Find the target loop context
var targetIdx int
if stmt.Label != "" {
// Find labeled loop
found := false
for i := len(g.loopStack) - 1; i >= 0; i-- {
if g.loopStack[i].label == stmt.Label {
targetIdx = i
found = true
break
}
}
if !found {
return fmt.Errorf("continue label not found: %s", stmt.Label)
}
} else {
// Continue to innermost loop
targetIdx = len(g.loopStack) - 1
}

// Add a jump instruction (will be patched later)
jumpPos := g.bytecode.CurrentPosition()
g.bytecode.AddInstruction(OpJump, 0) // Placeholder

// Record this jump position in the loop context
g.loopStack[targetIdx].continueJumps = append(g.loopStack[targetIdx].continueJumps, jumpPos)

return nil
}

// registerStructType registers a struct type's metadata
func (g *Generator) registerStructType(structDecl *ast.StructDecl) error {
info := &StructTypeInfo{
Name:    structDecl.Name,
Fields:  make(map[string]int),
Methods: make(map[string]*ast.FunctionDecl),
}

// Register fields with their indices
fieldIndex := 0
for _, member := range structDecl.Members {
if !member.IsMethod {
info.Fields[member.Name] = fieldIndex
fieldIndex++
} else {
// Register method
info.Methods[member.Method.Name] = member.Method
}
}

g.structTypes[structDecl.Name] = info
return nil
}

// registerEnumType registers an enum type and its variants
func (g *Generator) registerEnumType(enumDecl *ast.EnumDecl) error {
info := &EnumTypeInfo{
Name:     enumDecl.Name,
Variants: make(map[string]*VariantInfo),
}

// Register variants with their tags and arity
for tag, variant := range enumDecl.Variants {
variantInfo := &VariantInfo{
Name:  variant.Name,
Tag:   tag,
Arity: len(variant.Parameters),
}
info.Variants[variant.Name] = variantInfo
}

g.enumTypes[enumDecl.Name] = info
return nil
}

// generateStructMethods generates bytecode for struct methods
func (g *Generator) generateStructMethods(structDecl *ast.StructDecl) error {
for _, member := range structDecl.Members {
if member.IsMethod {
// Generate method with struct type prefix
methodName := structDecl.Name + "::" + member.Method.Name
method := member.Method
method.Name = methodName
if err := g.generateFunction(method); err != nil {
return fmt.Errorf("error generating method %s: %w", methodName, err)
}
}
}
return nil
}

// generateStructExpr generates bytecode for struct instantiation
func (g *Generator) generateStructExpr(expr *ast.StructExpr) error {
	// Determine the actual struct name to use (might be monomorphized)
	structName := expr.Name
	
	// Check if this is a generic struct instantiation
	if len(expr.TypeArgs) > 0 {
		// Check if this struct is a generic template
		if _, ok := g.genericStructs[expr.Name]; ok {
			// Monomorphize the struct with the provided type arguments
			mangledName, err := g.monomorphizeStruct(expr.Name, expr.TypeArgs)
			if err != nil {
				return fmt.Errorf("error monomorphizing struct %s: %w", expr.Name, err)
			}
			structName = mangledName
		} else {
			return fmt.Errorf("struct %s is not generic but type arguments provided", expr.Name)
		}
	}
	
	// Get struct type info (either concrete or monomorphized)
	structInfo, ok := g.structTypes[structName]
	if !ok {
		return fmt.Errorf("undefined struct type: %s", structName)
	}
	
	// Create a map to track which fields are initialized
	initializedFields := make(map[string]bool)
	
	// Build ordered list of field names and values
	type fieldPair struct {
		name  string
		value ast.Expression
		index int
	}
	fieldPairs := make([]fieldPair, len(structInfo.Fields))
	
	for _, fieldInit := range expr.FieldInits {
		fieldIndex, ok := structInfo.Fields[fieldInit.Name]
		if !ok {
			return fmt.Errorf("unknown field %s in struct %s", fieldInit.Name, structName)
		}
		fieldPairs[fieldIndex] = fieldPair{
			name:  fieldInit.Name,
			value: fieldInit.Value,
			index: fieldIndex,
		}
		initializedFields[fieldInit.Name] = true
	}
	
	// Check that all fields are initialized
	for fieldName := range structInfo.Fields {
		if !initializedFields[fieldName] {
			return fmt.Errorf("missing field initialization: %s", fieldName)
		}
	}
	
	// Push field name and value pairs onto stack
	for _, pair := range fieldPairs {
		// Push field name
		fieldNameIdx := g.bytecode.AddConstant(NewStringValue(pair.name))
		g.bytecode.AddInstruction(OpPush, int64(fieldNameIdx))
		
		// Push field value
		if err := g.generateExpression(pair.value); err != nil {
			return err
		}
	}
	
	// Push struct type name as a constant (use mangled name for generic structs)
	typeNameIdx := g.bytecode.AddConstant(NewStringValue(structName))
	g.bytecode.AddInstruction(OpPush, int64(typeNameIdx))
	
	// Allocate struct with field count
	g.bytecode.AddInstruction(OpAllocStruct, int64(len(structInfo.Fields)))
	
	return nil
}

// generateNewExpr generates bytecode for new expressions
func (g *Generator) generateNewExpr(expr *ast.NewExpr) error {
	// For now, new Type() is semantically equivalent to Type{} (struct literal)
	// This is a simplified implementation that treats new as struct initialization
	// The main difference is that new uses constructor arguments instead of field initializers
	
	// Get struct type info
	structInfo, ok := g.structTypes[expr.TypeName]
	if !ok {
		return fmt.Errorf("undefined struct type: %s", expr.TypeName)
	}
	
	// If there are constructor arguments, we need to map them to struct fields
	// For now, we'll support positional arguments that map to fields in declaration order
	if len(expr.Args) > 0 {
		// Get field names in order
		fieldNames := make([]string, len(structInfo.Fields))
		for fieldName, fieldIndex := range structInfo.Fields {
			fieldNames[fieldIndex] = fieldName
		}
		
		// Check argument count matches field count
		if len(expr.Args) != len(structInfo.Fields) {
			return fmt.Errorf("new %s: expected %d arguments, got %d", 
				expr.TypeName, len(structInfo.Fields), len(expr.Args))
		}
		
		// Push field name and value pairs onto stack
		for i, arg := range expr.Args {
			// Push field name
			fieldNameIdx := g.bytecode.AddConstant(NewStringValue(fieldNames[i]))
			g.bytecode.AddInstruction(OpPush, int64(fieldNameIdx))
			
			// Push field value
			if err := g.generateExpression(arg); err != nil {
				return err
			}
		}
	} else {
		// No arguments - need to initialize all fields with default values
		// For now, this is an error - all fields must be initialized
		if len(structInfo.Fields) > 0 {
			return fmt.Errorf("new %s: constructor requires %d arguments", 
				expr.TypeName, len(structInfo.Fields))
		}
	}
	
	// Push struct type name as a constant
	typeNameIdx := g.bytecode.AddConstant(NewStringValue(expr.TypeName))
	g.bytecode.AddInstruction(OpPush, int64(typeNameIdx))
	
	// Allocate struct with field count
	g.bytecode.AddInstruction(OpAllocStruct, int64(len(structInfo.Fields)))
	
	return nil
}

// generateFieldAccessExpr generates bytecode for field access
func (g *Generator) generateFieldAccessExpr(expr *ast.FieldAccessExpr) error {
// Generate the object expression
if err := g.generateExpression(expr.Object); err != nil {
return err
}

// For now, we'll use a simple approach: push field name and use OpLoadField
// In a more sophisticated implementation, we'd resolve the field index at compile time
fieldNameIdx := g.bytecode.AddConstant(NewStringValue(expr.Field))
g.bytecode.AddInstruction(OpPush, int64(fieldNameIdx))
g.bytecode.AddInstruction(OpLoadField, 0)

return nil
}

// generateMethodCallExpr generates bytecode for method calls
func (g *Generator) generateMethodCallExpr(expr *ast.MethodCallExpr) error {
// Generate arguments first
for _, arg := range expr.Arguments {
if err := g.generateExpression(arg); err != nil {
return err
}
}

// Generate the object (self) last so it's on top of the arguments
if err := g.generateExpression(expr.Object); err != nil {
return err
}

// For methods, we need to determine the struct type to resolve the method
// For now, we'll use a simplified approach with runtime resolution
// In a real implementation, we'd do type inference here

// Push the method name onto the stack
methodNameIdx := g.bytecode.AddConstant(NewStringValue(expr.Method))
g.bytecode.AddInstruction(OpPush, int64(methodNameIdx))

// Call the method (will be resolved at runtime)
// The operand is the number of arguments + 1 (for self)
argCount := len(expr.Arguments) + 1
inst := Instruction{
Op:       OpCall,
Operand:  int64(methodNameIdx), // This will need to be resolved to function index
ArgCount: argCount,
}
g.bytecode.Instructions = append(g.bytecode.Instructions, inst)

return nil
}

// generateEnumConstructorExpr generates bytecode for enum variant construction
func (g *Generator) generateEnumConstructorExpr(expr *ast.EnumConstructorExpr) error {
	// Determine the actual enum name to use (might be monomorphized)
	enumName := expr.EnumName
	
	// Check if this is a generic enum instantiation
	if len(expr.TypeArgs) > 0 {
		// Check if this enum is a generic template
		if _, ok := g.genericEnums[expr.EnumName]; ok {
			// Monomorphize the enum with the provided type arguments
			mangledName, err := g.monomorphizeEnum(expr.EnumName, expr.TypeArgs)
			if err != nil {
				return fmt.Errorf("error monomorphizing enum %s: %w", expr.EnumName, err)
			}
			enumName = mangledName
		} else {
			return fmt.Errorf("enum %s is not generic but type arguments provided", expr.EnumName)
		}
	}
	
	// Get enum type info (either concrete or monomorphized)
	enumInfo, ok := g.enumTypes[enumName]
	if !ok {
		return fmt.Errorf("undefined enum type: %s", enumName)
	}
	
	// Get variant info
	variantInfo, ok := enumInfo.Variants[expr.Variant]
	if !ok {
		return fmt.Errorf("undefined variant %s in enum %s", expr.Variant, enumName)
	}
	
	// Check argument count
	if len(expr.Arguments) != variantInfo.Arity {
		return fmt.Errorf("variant %s::%s expects %d arguments, got %d",
			enumName, expr.Variant, variantInfo.Arity, len(expr.Arguments))
	}
	
	// Push variant data (arguments) onto stack
	for _, arg := range expr.Arguments {
		if err := g.generateExpression(arg); err != nil {
			return err
		}
	}
	
	// Push enum name (use mangled name for generic enums)
	enumNameIdx := g.bytecode.AddConstant(NewStringValue(enumName))
	g.bytecode.AddInstruction(OpPush, int64(enumNameIdx))
	
	// Push variant name
	variantNameIdx := g.bytecode.AddConstant(NewStringValue(expr.Variant))
	g.bytecode.AddInstruction(OpPush, int64(variantNameIdx))
	
	// Push variant tag
	tagIdx := g.bytecode.AddConstant(NewIntValue(int64(variantInfo.Tag)))
	g.bytecode.AddInstruction(OpPush, int64(tagIdx))
	
	// Allocate enum with data count
	g.bytecode.AddInstruction(OpAllocEnum, int64(variantInfo.Arity))
	
	return nil
}

// generateMatchExpr generates bytecode for match expressions
func (g *Generator) generateMatchExpr(expr *ast.MatchExpr) error {
// Generate the value to match against
if err := g.generateExpression(expr.Value); err != nil {
return err
}

// Track jump positions for each case
var caseEndJumps []int

// Generate code for each case
for i, caseExpr := range expr.Cases {
// Duplicate the value on stack for pattern matching
g.bytecode.AddInstruction(OpDup, 0)

// Generate pattern matching code
// This will consume one copy of the value and leave a boolean on stack
if err := g.generatePatternMatch(caseExpr.Pattern); err != nil {
return err
}

// Jump to next case if pattern doesnt match
nextCaseJump := g.bytecode.CurrentPosition()
g.bytecode.AddInstruction(OpJumpIfFalse, 0) // Placeholder

// Pop the duplicate value since we matched
g.bytecode.AddInstruction(OpPop, 0)

// Generate the case expression
if err := g.generateExpression(caseExpr.Expression); err != nil {
return err
}

// Jump to end of match expression
if i < len(expr.Cases)-1 {
endJump := g.bytecode.CurrentPosition()
g.bytecode.AddInstruction(OpJump, 0) // Placeholder
caseEndJumps = append(caseEndJumps, endJump)
}

// Patch the next case jump to point here
nextCasePos := g.bytecode.CurrentPosition()
g.bytecode.PatchJump(nextCaseJump, nextCasePos)
}

// Pop the remaining value if no cases matched (should not happen with exhaustive matching)
g.bytecode.AddInstruction(OpPop, 0)

// Patch all end jumps to point past the match expression
endPos := g.bytecode.CurrentPosition()
for _, jumpPos := range caseEndJumps {
g.bytecode.PatchJump(jumpPos, endPos)
}

return nil
}

// generatePatternMatch generates bytecode for pattern matching
// It consumes the value to match and pushes a boolean (true if match, false otherwise)
func (g *Generator) generatePatternMatch(pattern *ast.Pattern) error {
switch pattern.Kind {
case ast.PatternWildcard, ast.PatternVariable:
// Wildcard and variable patterns always match
// Pop the value to match (consume it)
g.bytecode.AddInstruction(OpPop, 0)
// Push true for the match result
trueIdx := g.bytecode.AddConstant(NewBoolValue(true))
g.bytecode.AddInstruction(OpPush, int64(trueIdx))

case ast.PatternLiteral:
// Match against literal value
// Push the literal
if err := g.generateLiteral(pattern.Literal); err != nil {
return err
}
// Compare for equality
g.bytecode.AddInstruction(OpEq, 0)

case ast.PatternConstructor:
// Match enum variant
// For now, assume pattern.Value is "Variant" or "Enum::Variant"
variantNameIdx := g.bytecode.AddConstant(NewStringValue(pattern.Value))
g.bytecode.AddInstruction(OpPush, int64(variantNameIdx))
g.bytecode.AddInstruction(OpMatchVariant, 0)

default:
return fmt.Errorf("unsupported pattern kind: %v", pattern.Kind)
}

return nil
}


// generateTryExpr generates bytecode for the ? operator
// The ? operator works with Result<T, E> enums:
// - If the value is Ok(value), unwraps and continues with value
// - If the value is Err(error), propagates the error by returning early
func (g *Generator) generateTryExpr(expr *ast.TryExpr) error {
// Generate the operand expression (should be a Result)
if err := g.generateExpression(expr.Operand); err != nil {
return err
}

// Generate the OpTryPropagate instruction
// This will check if the value is Ok or Err and handle accordingly
g.bytecode.AddInstruction(OpTryPropagate, 0)

return nil
}

// generateListExpr generates bytecode for a list/array literal
func (g *Generator) generateListExpr(expr *ast.ListExpr) error {
// For now, only support simple array literals, not comprehensions
if expr.Comprehension != nil {
return fmt.Errorf("list comprehensions not yet supported")
}

// Generate code for each element (push onto stack)
for _, elem := range expr.Elements {
if err := g.generateExpression(elem); err != nil {
return fmt.Errorf("error generating array element: %w", err)
}
}

// Allocate array with N elements
g.bytecode.AddInstruction(OpAllocArray, int64(len(expr.Elements)))
return nil
}

// generateIndexExpr generates bytecode for array indexing (arr[i])
func (g *Generator) generateIndexExpr(expr *ast.IndexExpr) error {
// Generate code for the array/object being indexed
if err := g.generateExpression(expr.Object); err != nil {
return fmt.Errorf("error generating indexed object: %w", err)
}

// Generate code for the index
if err := g.generateExpression(expr.Index); err != nil {
return fmt.Errorf("error generating index: %w", err)
}

// Load element from array
g.bytecode.AddInstruction(OpLoadArray, 0)
return nil
}

// unescapeString processes escape sequences in a string literal
func (g *Generator) unescapeString(s string) string {
result := make([]byte, 0, len(s))
i := 0
for i < len(s) {
if s[i] == '\\' && i+1 < len(s) {
switch s[i+1] {
case 'n':
result = append(result, '\n')
case 't':
result = append(result, '\t')
case 'r':
result = append(result, '\r')
case 'b':
result = append(result, '\b')
case 'f':
result = append(result, '\f')
case '"':
result = append(result, '"')
case '\'':
result = append(result, '\'')
case '\\':
result = append(result, '\\')
default:
// Unknown escape sequence, keep as is
result = append(result, s[i], s[i+1])
}
i += 2
} else {
result = append(result, s[i])
i++
}
}
return string(result)
}

// monomorphizeFunction generates a specialized version of a generic function
func (g *Generator) monomorphizeFunction(baseName string, typeArgs []ast.Type) (string, error) {
// Create mangled name for this specialization
mangledName := semantic.MangleName(baseName, typeArgs)
	
// Check if already monomorphized
if mono, ok := g.monomorphized[mangledName]; ok && mono.Generated {
return mangledName, nil
}
	
// Get the generic function template
template, ok := g.genericFunctions[baseName]
if !ok {
return "", fmt.Errorf("generic function %s not found", baseName)
}
	
// Verify type argument count matches
if len(template.GenericParams) != len(typeArgs) {
return "", fmt.Errorf("wrong number of type arguments for %s: expected %d, got %d",
baseName, len(template.GenericParams), len(typeArgs))
}
	
// Create a child generic context for this monomorphization
monoCtx := g.genericContext.NewChildContext()
	
// Add type parameter substitutions
for i, param := range template.GenericParams {
monoCtx.AddTypeParameter(param)
monoCtx.Substitutions[param.Name] = typeArgs[i]
}
	
// Create a specialized version of the function
specializedFn := &ast.FunctionDecl{
Name:       mangledName,
IsPublic:   template.IsPublic,
Parameters: g.substituteParameters(template.Parameters, monoCtx),
ReturnType: monoCtx.Substitute(template.ReturnType),
Body:       template.Body, // Body stays the same, types are handled at runtime
Position:   template.Position,
}
	
// Mark as being generated
g.monomorphized[mangledName] = &MonomorphizedFunction{
BaseName:   baseName,
TypeArgs:   typeArgs,
MangledName: mangledName,
Generated:  false,
}
	
// Save current generic context
savedCtx := g.genericContext
g.genericContext = monoCtx
	
// Generate the specialized function
if err := g.generateFunction(specializedFn); err != nil {
g.genericContext = savedCtx
return "", fmt.Errorf("error monomorphizing %s: %w", baseName, err)
}
	
// Restore generic context
g.genericContext = savedCtx
	
// Mark as generated
g.monomorphized[mangledName].Generated = true
	
return mangledName, nil
}

// substituteParameters applies type substitutions to function parameters
func (g *Generator) substituteParameters(params []*ast.Parameter, ctx *semantic.GenericContext) []*ast.Parameter {
result := make([]*ast.Parameter, len(params))
for i, param := range params {
result[i] = &ast.Parameter{
Name:        param.Name,
Type:        ctx.Substitute(param.Type),
IsSelf:      param.IsSelf,
IsReference: param.IsReference,
IsMutable:   param.IsMutable,
Position:    param.Position,
}
}
return result
}

// monomorphizeStruct generates a specialized version of a generic struct
func (g *Generator) monomorphizeStruct(baseName string, typeArgs []ast.Type) (string, error) {
// Create mangled name for this specialization
mangledName := semantic.MangleName(baseName, typeArgs)

// Check if already monomorphized
if mono, ok := g.monomorphizedStructs[mangledName]; ok && mono.Generated {
return mangledName, nil
}

// Get the generic struct template
template, ok := g.genericStructs[baseName]
if !ok {
return "", fmt.Errorf("generic struct %s not found", baseName)
}

// Verify type argument count matches
if len(template.GenericParams) != len(typeArgs) {
return "", fmt.Errorf("wrong number of type arguments for %s: expected %d, got %d",
baseName, len(template.GenericParams), len(typeArgs))
}

// Create a child generic context for this monomorphization
monoCtx := g.genericContext.NewChildContext()

// Add type parameter substitutions
for i, param := range template.GenericParams {
monoCtx.AddTypeParameter(param)
monoCtx.Substitutions[param.Name] = typeArgs[i]
}

// Create a specialized version of the struct
specializedStruct := &ast.StructDecl{
Name:       mangledName,
IsPublic:   template.IsPublic,
Members:    g.substituteStructMembers(template.Members, monoCtx),
Position:   template.Position,
// No GenericParams - this is a concrete type
}

// Mark as being generated
g.monomorphizedStructs[mangledName] = &MonomorphizedStruct{
BaseName:    baseName,
TypeArgs:    typeArgs,
MangledName: mangledName,
Generated:   false,
}

// Register the specialized struct type
if err := g.registerStructType(specializedStruct); err != nil {
return "", fmt.Errorf("error registering monomorphized struct %s: %w", mangledName, err)
}

// Mark as generated
g.monomorphizedStructs[mangledName].Generated = true

return mangledName, nil
}

// substituteStructMembers applies type substitutions to struct members
func (g *Generator) substituteStructMembers(members []*ast.StructMember, ctx *semantic.GenericContext) []*ast.StructMember {
result := make([]*ast.StructMember, len(members))
for i, member := range members {
result[i] = &ast.StructMember{
Name:     member.Name,
Type:     ctx.Substitute(member.Type),
IsMethod: member.IsMethod,
Method:   member.Method, // TODO: Substitute method types if needed
Position: member.Position,
}
}
return result
}

// monomorphizeEnum generates a specialized version of a generic enum
func (g *Generator) monomorphizeEnum(baseName string, typeArgs []ast.Type) (string, error) {
// Create mangled name for this specialization
mangledName := semantic.MangleName(baseName, typeArgs)

// Check if already monomorphized
if mono, ok := g.monomorphizedEnums[mangledName]; ok && mono.Generated {
return mangledName, nil
}

// Get the generic enum template
template, ok := g.genericEnums[baseName]
if !ok {
return "", fmt.Errorf("generic enum %s not found", baseName)
}

// Verify type argument count matches
if len(template.GenericParams) != len(typeArgs) {
return "", fmt.Errorf("wrong number of type arguments for %s: expected %d, got %d",
baseName, len(template.GenericParams), len(typeArgs))
}

// Create a child generic context for this monomorphization
monoCtx := g.genericContext.NewChildContext()

// Add type parameter substitutions
for i, param := range template.GenericParams {
monoCtx.AddTypeParameter(param)
monoCtx.Substitutions[param.Name] = typeArgs[i]
}

// Create a specialized version of the enum
specializedEnum := &ast.EnumDecl{
Name:     mangledName,
IsPublic: template.IsPublic,
Variants: g.substituteEnumVariants(template.Variants, monoCtx),
Position: template.Position,
// No GenericParams - this is a concrete type
}

// Mark as being generated
g.monomorphizedEnums[mangledName] = &MonomorphizedEnum{
BaseName:    baseName,
TypeArgs:    typeArgs,
MangledName: mangledName,
Generated:   false,
}

// Register the specialized enum type
if err := g.registerEnumType(specializedEnum); err != nil {
return "", fmt.Errorf("error registering monomorphized enum %s: %w", mangledName, err)
}

// Mark as generated
g.monomorphizedEnums[mangledName].Generated = true

return mangledName, nil
}

// substituteEnumVariants applies type substitutions to enum variants
func (g *Generator) substituteEnumVariants(variants []*ast.EnumVariant, ctx *semantic.GenericContext) []*ast.EnumVariant {
result := make([]*ast.EnumVariant, len(variants))
for i, variant := range variants {
// Substitute types in variant parameters
newParams := make([]*ast.EnumVariantParam, len(variant.Parameters))
for j, param := range variant.Parameters {
newParams[j] = &ast.EnumVariantParam{
Name:     param.Name,
Type:     ctx.Substitute(param.Type),
Position: param.Position,
}
}

result[i] = &ast.EnumVariant{
Name:       variant.Name,
Parameters: newParams,
Position:   variant.Position,
}
}
return result
}

// inferCallTypeArguments infers type arguments for a generic function call
func (g *Generator) inferCallTypeArguments(fn *ast.FunctionDecl, args []ast.Expression) ([]ast.Type, error) {
// Get parameter types
paramTypes := make([]ast.Type, len(fn.Parameters))
for i, param := range fn.Parameters {
paramTypes[i] = param.Type
}
	
// Infer argument types
argTypes := make([]ast.Type, len(args))
for i, arg := range args {
// For now, use a simple type inference based on literals
// In a full implementation, this would use the semantic analyzer
argType, err := g.inferExpressionType(arg)
if err != nil {
return nil, fmt.Errorf("cannot infer type for argument %d: %w", i, err)
}
argTypes[i] = argType
}
	
// Use the generic context to infer type arguments
return g.inferTypeArgsFromCall(fn.GenericParams, paramTypes, argTypes)
}

// inferExpressionType performs simple type inference on an expression
func (g *Generator) inferExpressionType(expr ast.Expression) (ast.Type, error) {
switch e := expr.(type) {
case *ast.Literal:
switch e.Kind {
case ast.LiteralInt:
return &ast.TypeReference{Name: "Int"}, nil
case ast.LiteralFloat:
return &ast.TypeReference{Name: "Float"}, nil
case ast.LiteralBool:
return &ast.TypeReference{Name: "Bool"}, nil
case ast.LiteralString:
return &ast.TypeReference{Name: "String"}, nil
case ast.LiteralChar:
return &ast.TypeReference{Name: "Char"}, nil
default:
return nil, fmt.Errorf("unknown literal kind: %v", e.Kind)
}
case *ast.Identifier:
// Look up in local variables to get type
// For now, return nil to indicate we need more context
return nil, fmt.Errorf("cannot infer type for identifier %s", e.Name)
default:
return nil, fmt.Errorf("cannot infer type for expression: %T", expr)
}
}

// inferTypeArgsFromCall infers type arguments from a function call
func (g *Generator) inferTypeArgsFromCall(genericParams []*ast.GenericParam, paramTypes []ast.Type, argTypes []ast.Type) ([]ast.Type, error) {
if len(paramTypes) != len(argTypes) {
return nil, fmt.Errorf("parameter count mismatch")
}
	
// Create a temporary context for inference
inferCtx := g.genericContext.NewChildContext()
	
// Add type parameters
for _, param := range genericParams {
inferCtx.AddTypeParameter(param)
}
	
// Unify each parameter with its argument
for i := range paramTypes {
if err := inferCtx.Unify(paramTypes[i], argTypes[i]); err != nil {
return nil, fmt.Errorf("type inference failed for parameter %d: %w", i, err)
}
}
	
// Extract inferred types
result := make([]ast.Type, len(genericParams))
for i, param := range genericParams {
if sub, ok := inferCtx.Substitutions[param.Name]; ok {
result[i] = sub
} else {
return nil, fmt.Errorf("could not infer type for parameter %s", param.Name)
}
}
	
return result, nil
}
