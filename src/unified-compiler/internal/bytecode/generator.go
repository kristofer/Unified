package bytecode

import (
"fmt"
"unified-compiler/internal/ast"
)

// LoopContext tracks loop information for break/continue
type LoopContext struct {
startPos      int    // Position to jump to for continue
breakJumps    []int  // Positions of break jumps to patch
continueJumps []int  // Positions of continue jumps to patch
label         string // Optional loop label
}

// Generator converts AST to bytecode
type Generator struct {
bytecode      *Bytecode
localVars     map[string]int // Variable name -> stack index
localVarCount int            // Number of local variables
loopStack     []LoopContext  // Stack of nested loop contexts
}

// NewGenerator creates a new bytecode generator
func NewGenerator() *Generator {
return &Generator{
bytecode:  NewBytecode(),
localVars: make(map[string]int),
}
}

// Generate converts a program AST to bytecode
func (g *Generator) Generate(program *ast.Program) (*Bytecode, error) {
// Generate code for all top-level items
for _, item := range program.Items {
switch item := item.(type) {
case *ast.FunctionDecl:
if err := g.generateFunction(item); err != nil {
return nil, fmt.Errorf("error generating function %s: %w", item.Name, err)
}
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
val = NewStringValue(lit.Value)
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

// generateAssignment generates bytecode for variable assignment
func (g *Generator) generateAssignment(expr *ast.BinaryExpr) error {
// Left side must be an identifier
ident, ok := expr.Left.(*ast.Identifier)
if !ok {
return fmt.Errorf("assignment target must be a variable")
}

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
