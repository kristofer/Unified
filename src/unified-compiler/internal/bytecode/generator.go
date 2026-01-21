package bytecode

import (
"fmt"
"unified-compiler/internal/ast"
)

// Generator converts AST to bytecode
type Generator struct {
bytecode      *Bytecode
localVars     map[string]int // Variable name -> stack index
localVarCount int            // Number of local variables
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
default:
return fmt.Errorf("unsupported statement type: %T", stmt)
}
}

// generateLetStatement generates bytecode for let statement
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
default:
return fmt.Errorf("unsupported binary operator: %s", expr.Operator)
}

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
g.bytecode.AddInstruction(OpCall, int64(funcIdx))

return nil
}
