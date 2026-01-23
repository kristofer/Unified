package bytecode

import (
"fmt"
"unified-compiler/internal/ast"
)

// generateAssignmentStatement generates bytecode for assignment statement
func (g *Generator) generateAssignmentStatement(stmt *ast.AssignmentStatement) error {
// Look up the variable
varIdx, ok := g.localVars[stmt.Target]
if !ok {
return fmt.Errorf("undefined variable: %s", stmt.Target)
}

// For compound assignments (+=, -=, etc.), load current value first
if stmt.Operator != ast.AssignNormal {
g.bytecode.AddInstruction(OpLoadLocal, int64(varIdx))
}

// Generate the value expression
if err := g.generateExpression(stmt.Value); err != nil {
return err
}

// For compound assignments, apply the operation
switch stmt.Operator {
case ast.AssignAdd:
g.bytecode.AddInstruction(OpAdd, 0)
case ast.AssignSub:
g.bytecode.AddInstruction(OpSub, 0)
case ast.AssignMul:
g.bytecode.AddInstruction(OpMul, 0)
case ast.AssignDiv:
g.bytecode.AddInstruction(OpDiv, 0)
case ast.AssignMod:
g.bytecode.AddInstruction(OpMod, 0)
case ast.AssignNormal:
// No operation needed
default:
return fmt.Errorf("unsupported assignment operator: %v", stmt.Operator)
}

// Store the value
g.bytecode.AddInstruction(OpStoreLocal, int64(varIdx))

return nil
}
