package wasm

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"strconv"
	"unified-compiler/internal/ast"
)

// generateStatement generates WASM bytecode for a statement
func (g *Generator) generateStatement(body *bytes.Buffer, stmt ast.Statement) error {
	switch s := stmt.(type) {
	case *ast.ReturnStatement:
		return g.generateReturn(body, s)
	case *ast.LetStatement:
		return g.generateLet(body, s)
	case *ast.ExprStatement:
		if err := g.generateExpression(body, s.Expression); err != nil {
			return err
		}
		// Pop the expression result if it's not used
		body.WriteByte(0x1A) // drop
		return nil
	case *ast.IfStatement:
		return g.generateIf(body, s)
	case *ast.WhileStatement:
		return g.generateWhile(body, s)
	case *ast.AssignmentStatement:
		return g.generateAssignment(body, s)
	default:
		return fmt.Errorf("unsupported statement type: %T", stmt)
	}
}

// generateReturn generates WASM bytecode for a return statement
func (g *Generator) generateReturn(body *bytes.Buffer, ret *ast.ReturnStatement) error {
	if ret.Value != nil {
		if err := g.generateExpression(body, ret.Value); err != nil {
			return err
		}
	}
	body.WriteByte(0x0F) // return
	return nil
}

// generateLet generates WASM bytecode for a let statement
func (g *Generator) generateLet(body *bytes.Buffer, let *ast.LetStatement) error {
	// Generate the initial value
	if let.Value != nil {
		if err := g.generateExpression(body, let.Value); err != nil {
			return err
		}
	} else {
		// Default value
		g.emitDefaultValue(body, let.Type)
	}

	// Assign local variable index
	localIndex := g.localVarCount
	g.localVars[let.Name] = localIndex
	g.localVarTypes[let.Name] = let.Type
	g.localVarCount++

	// Store to local variable
	g.emitSetLocal(body, localIndex)

	return nil
}

// generateAssignment generates WASM bytecode for an assignment statement
func (g *Generator) generateAssignment(body *bytes.Buffer, assign *ast.AssignmentStatement) error {
	// Generate the value expression
	if err := g.generateExpression(body, assign.Value); err != nil {
		return err
	}

	// Look up the variable
	localIndex, ok := g.localVars[assign.Target]
	if !ok {
		return fmt.Errorf("undefined variable: %s", assign.Target)
	}

	// Store to local variable
	g.emitSetLocal(body, localIndex)

	return nil
}

// generateIf generates WASM bytecode for an if statement
func (g *Generator) generateIf(body *bytes.Buffer, ifStmt *ast.IfStatement) error {
	// Generate condition
	if err := g.generateExpression(body, ifStmt.Condition); err != nil {
		return err
	}

	// TODO: Type checking and conversion needed!
	// The if instruction requires i32 on the stack, but comparison operators
	// return i32 while other values might be i64. Need proper type conversion.
	// Currently assumes the condition expression produces i32.
	
	// if instruction
	body.WriteByte(0x04) // if
	body.WriteByte(0x40) // empty block type

	// Then block
	if ifStmt.ThenBlock != nil {
		for _, stmt := range ifStmt.ThenBlock.Statements {
			if err := g.generateStatement(body, stmt); err != nil {
				return err
			}
		}
	}

	// Else block
	if ifStmt.ElseBlock != nil && len(ifStmt.ElseBlock.Statements) > 0 {
		body.WriteByte(0x05) // else
		for _, stmt := range ifStmt.ElseBlock.Statements {
			if err := g.generateStatement(body, stmt); err != nil {
				return err
			}
		}
	}

	body.WriteByte(0x0B) // end

	return nil
}

// generateWhile generates WASM bytecode for a while loop
func (g *Generator) generateWhile(body *bytes.Buffer, whileStmt *ast.WhileStatement) error {
	// WASM while loop using block and loop
	// block (outer - for break)
	body.WriteByte(0x02) // block
	body.WriteByte(0x40) // empty block type

	// loop (inner - for continue)
	body.WriteByte(0x03) // loop
	body.WriteByte(0x40) // empty block type

	// TODO: Type checking needed for while loop condition
	// Branch instructions require i32, but generated condition may be i64
	// Generate condition
	if err := g.generateExpression(body, whileStmt.Condition); err != nil {
		return err
	}

	// Branch if false (exit loop)
	// NOTE: i32.eqz expects i32 input - this will fail if condition is i64
	body.WriteByte(0x45) // i32.eqz (negate condition)
	body.WriteByte(0x0D) // br_if
	body.WriteByte(0x01) // break to outer block

	// Loop body
	for _, stmt := range whileStmt.Body.Statements {
		if err := g.generateStatement(body, stmt); err != nil {
			return err
		}
	}

	// Branch back to loop start
	body.WriteByte(0x0C) // br
	body.WriteByte(0x00) // continue to loop

	body.WriteByte(0x0B) // end loop
	body.WriteByte(0x0B) // end block

	return nil
}

// generateExpression generates WASM bytecode for an expression
func (g *Generator) generateExpression(body *bytes.Buffer, expr ast.Expression) error {
	switch e := expr.(type) {
	case *ast.Literal:
		return g.generateLiteral(body, e)
	case *ast.BinaryExpr:
		return g.generateBinaryExpr(body, e)
	case *ast.UnaryExpr:
		return g.generateUnaryExpr(body, e)
	case *ast.Identifier:
		return g.generateIdentifier(body, e)
	case *ast.CallExpr:
		return g.generateCall(body, e)
	default:
		return fmt.Errorf("unsupported expression type: %T", expr)
	}
}

// generateLiteral generates WASM bytecode for a literal
func (g *Generator) generateLiteral(body *bytes.Buffer, lit *ast.Literal) error {
	switch lit.Kind {
	case ast.LiteralInt:
		// Parse the integer value
		val, err := strconv.ParseInt(lit.Value, 0, 64)
		if err != nil {
			return fmt.Errorf("invalid integer literal: %s", lit.Value)
		}
		body.WriteByte(0x42) // i64.const
		g.emitLEB128(body, val)
	case ast.LiteralFloat:
		// Parse the float value
		val, err := strconv.ParseFloat(lit.Value, 64)
		if err != nil {
			return fmt.Errorf("invalid float literal: %s", lit.Value)
		}
		body.WriteByte(0x44) // f64.const
		binary.Write(body, binary.LittleEndian, val)
	case ast.LiteralBool:
		// Boolean as i32
		body.WriteByte(0x41) // i32.const
		if lit.Value == "true" {
			body.WriteByte(0x01)
		} else {
			body.WriteByte(0x00)
		}
	default:
		return fmt.Errorf("unsupported literal kind: %v", lit.Kind)
	}
	return nil
}

// generateIdentifier generates WASM bytecode for an identifier (variable)
func (g *Generator) generateIdentifier(body *bytes.Buffer, id *ast.Identifier) error {
	localIndex, ok := g.localVars[id.Name]
	if !ok {
		return fmt.Errorf("undefined variable: %s", id.Name)
	}

	g.emitGetLocal(body, localIndex)
	return nil
}

// generateCall generates WASM bytecode for a function call
func (g *Generator) generateCall(body *bytes.Buffer, call *ast.CallExpr) error {
	// For now, we only support simple function calls where Function is an Identifier
	funcId, ok := call.Function.(*ast.Identifier)
	if !ok {
		return fmt.Errorf("complex function calls not yet supported")
	}

	// Generate arguments in order
	for _, arg := range call.Arguments {
		if err := g.generateExpression(body, arg); err != nil {
			return err
		}
	}

	// Find function index
	funcIndex := -1
	for i, fn := range g.module.Functions {
		if fn.Name == funcId.Name {
			funcIndex = i
			break
		}
	}

	if funcIndex == -1 {
		return fmt.Errorf("undefined function: %s", funcId.Name)
	}

	// Emit call instruction
	body.WriteByte(0x10) // call
	g.emitULEB128(body, uint64(funcIndex))

	return nil
}

// generateBinaryExpr generates WASM bytecode for a binary expression
func (g *Generator) generateBinaryExpr(body *bytes.Buffer, expr *ast.BinaryExpr) error {
	// Generate left operand
	if err := g.generateExpression(body, expr.Left); err != nil {
		return err
	}

	// Generate right operand
	if err := g.generateExpression(body, expr.Right); err != nil {
		return err
	}

	// Generate operator based on OperatorType
	switch expr.Operator {
	case ast.OperatorAdd:
		body.WriteByte(0x7C) // i64.add
	case ast.OperatorSub:
		body.WriteByte(0x7D) // i64.sub
	case ast.OperatorMul:
		body.WriteByte(0x7E) // i64.mul
	case ast.OperatorDiv:
		body.WriteByte(0x7F) // i64.div_s
	case ast.OperatorMod:
		body.WriteByte(0x81) // i64.rem_s
	case ast.OperatorEq:
		body.WriteByte(0x51) // i64.eq
	case ast.OperatorNe:
		body.WriteByte(0x52) // i64.ne
	case ast.OperatorLt:
		body.WriteByte(0x53) // i64.lt_s
	case ast.OperatorGt:
		body.WriteByte(0x55) // i64.gt_s
	case ast.OperatorLe:
		body.WriteByte(0x54) // i64.le_s
	case ast.OperatorGe:
		body.WriteByte(0x56) // i64.ge_s
	case ast.OperatorAnd:
		// TODO: Type mismatch - should use i64.and (0x83) for consistency
		// Currently uses i32.and which will cause validation errors with i64 operands
		body.WriteByte(0x71) // i32.and (FIXME: should match operand types)
	case ast.OperatorOr:
		// TODO: Type mismatch - should use i64.or (0x84) for consistency
		body.WriteByte(0x72) // i32.or (FIXME: should match operand types)
	case ast.OperatorBitAnd:
		body.WriteByte(0x83) // i64.and
	case ast.OperatorBitOr:
		body.WriteByte(0x84) // i64.or
	case ast.OperatorBitXor:
		body.WriteByte(0x85) // i64.xor
	case ast.OperatorLShift:
		body.WriteByte(0x86) // i64.shl
	case ast.OperatorRShift:
		body.WriteByte(0x88) // i64.shr_s
	default:
		return fmt.Errorf("unsupported binary operator: %v", expr.Operator)
	}

	return nil
}

// generateUnaryExpr generates WASM bytecode for a unary expression
func (g *Generator) generateUnaryExpr(body *bytes.Buffer, expr *ast.UnaryExpr) error {
	// Generate operand
	if err := g.generateExpression(body, expr.Operand); err != nil {
		return err
	}

	// Generate operator based on OperatorType
	switch expr.Operator {
	case ast.OperatorUnaryMinus:
		// Negate by multiplying by -1
		body.WriteByte(0x42) // i64.const
		g.emitLEB128(body, -1)
		body.WriteByte(0x7E) // i64.mul
	case ast.OperatorNot:
		// TODO: Type mismatch - i32.eqz expects i32 but operand may be i64
		// Use i64.eqz (0x50) for i64 operands or add type conversion
		body.WriteByte(0x45) // i32.eqz (FIXME: should match operand type)
	case ast.OperatorBitNot:
		// Bitwise not: XOR with -1
		body.WriteByte(0x42) // i64.const
		g.emitLEB128(body, -1)
		body.WriteByte(0x85) // i64.xor
	default:
		return fmt.Errorf("unsupported unary operator: %v", expr.Operator)
	}

	return nil
}

// Helper functions for WASM encoding

// emitGetLocal emits a local.get instruction
func (g *Generator) emitGetLocal(body *bytes.Buffer, index int) {
	body.WriteByte(0x20) // local.get
	g.emitULEB128(body, uint64(index))
}

// emitSetLocal emits a local.set instruction
func (g *Generator) emitSetLocal(body *bytes.Buffer, index int) {
	body.WriteByte(0x21) // local.set
	g.emitULEB128(body, uint64(index))
}

// emitLEB128 emits a signed LEB128 encoded integer
func (g *Generator) emitLEB128(body *bytes.Buffer, value int64) {
	for {
		b := byte(value & 0x7F)
		value >>= 7
		if (value == 0 && (b&0x40) == 0) || (value == -1 && (b&0x40) != 0) {
			body.WriteByte(b)
			break
		}
		body.WriteByte(b | 0x80)
	}
}

// emitULEB128 emits an unsigned LEB128 encoded integer
func (g *Generator) emitULEB128(body *bytes.Buffer, value uint64) {
	for {
		b := byte(value & 0x7F)
		value >>= 7
		if value == 0 {
			body.WriteByte(b)
			break
		}
		body.WriteByte(b | 0x80)
	}
}
