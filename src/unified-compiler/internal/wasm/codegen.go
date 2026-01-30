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
	case *ast.ForStatement:
		return g.generateFor(body, s)
	case *ast.AssignmentStatement:
		return g.generateAssignment(body, s)
	case *ast.BreakStatement:
		return g.generateBreak(body, s)
	case *ast.ContinueStatement:
		return g.generateContinue(body, s)
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
		
		// Add type conversion if the expression type doesn't match return type
		if g.currentFuncReturnType != nil {
			exprType := g.getExpressionType(ret.Value)
			returnType := g.convertType(g.currentFuncReturnType)
			
			if exprType != returnType {
				// Convert between types
				if exprType == I32 && returnType == I64 {
					g.emitI32ToI64Conversion(body)
				} else if exprType == I64 && returnType == I32 {
					g.emitI64ToI32Conversion(body)
				}
				// Note: Float conversions would be added here in a complete implementation
			}
		}
	}
	body.WriteByte(0x0F) // return
	return nil
}

// generateLet generates WASM bytecode for a let statement
func (g *Generator) generateLet(body *bytes.Buffer, let *ast.LetStatement) error {
	// Determine the WASM type for this variable
	wasmType := g.convertType(let.Type)
	
	// Generate the initial value
	if let.Value != nil {
		if err := g.generateExpression(body, let.Value); err != nil {
			return err
		}
		
		// Add type conversion if needed
		exprType := g.getExpressionType(let.Value)
		if exprType != wasmType {
			// Convert between types
			if exprType == I32 && wasmType == I64 {
				g.emitI32ToI64Conversion(body)
			} else if exprType == I64 && wasmType == I32 {
				g.emitI64ToI32Conversion(body)
			}
			// Note: Float conversions would be added here in a complete implementation
		}
	} else {
		// Default value
		g.emitDefaultValue(body, let.Type)
	}

	// Assign local variable index
	localIndex := g.localVarCount
	g.localVars[let.Name] = localIndex
	g.localVarTypes[let.Name] = let.Type
	
	// Track the type order for this local
	g.localTypeOrder = append(g.localTypeOrder, wasmType)
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
	// Generate condition expression
	// The condition must produce an i32 value on the stack
	// Boolean literals produce i32
	// Comparison operators (i64.lt_s, i64.eq, etc.) return i32
	// Logical operators (i32.and, i32.or) work with i32
	if err := g.generateExpression(body, ifStmt.Condition); err != nil {
		return err
	}
	
	// if instruction
	body.WriteByte(0x04) // if
	body.WriteByte(0x40) // empty block type (void)

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

	// Generate condition expression (must produce i32)
	// Comparison operators return i32, boolean literals are i32
	if err := g.generateExpression(body, whileStmt.Condition); err != nil {
		return err
	}

	// Branch if false (exit loop)
	// i32.eqz expects i32 input and returns i32 (0 becomes 1, non-zero becomes 0)
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
	case *ast.StructExpr:
		return g.generateStructExpr(body, e)
	case *ast.FieldAccessExpr:
		return g.generateFieldAccess(body, e)
	case *ast.EnumConstructorExpr:
		return g.generateEnumConstructor(body, e)
	case *ast.ListExpr:
		return g.generateListExpr(body, e)
	case *ast.IndexExpr:
		return g.generateIndexExpr(body, e)
	case *ast.MatchExpr:
		return g.generateMatchExpr(body, e)
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
	case ast.LiteralString:
		// String literals need to be stored in data section and referenced
		// Check if we've already allocated this string
		if offset, ok := g.stringTable[lit.Value]; ok {
			// String already allocated, just return its pointer
			body.WriteByte(0x41) // i32.const
			g.emitULEB128(body, uint64(offset))
		} else {
			// Allocate new string in memory
			// Layout: [length:i32][bytes...]
			stringBytes := []byte(lit.Value)
			
			// Create data with length prefix
			data := make([]byte, 4+len(stringBytes))
			binary.LittleEndian.PutUint32(data[0:4], uint32(len(stringBytes)))
			copy(data[4:], stringBytes)
			
			// Allocate in memory
			offset := g.memoryAllocator.AllocateWithData(data)
			g.stringTable[lit.Value] = offset
			
			// Return pointer to string
			body.WriteByte(0x41) // i32.const
			g.emitULEB128(body, uint64(offset))
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
	// Check for string concatenation (if both operands are strings and operator is Add)
	// For now, we'll detect this by checking if operands are string literals
	// TODO: Implement proper type inference
	
	// Determine types of operands
	leftType := g.getExpressionType(expr.Left)
	rightType := g.getExpressionType(expr.Right)
	
	// Generate left operand
	if err := g.generateExpression(body, expr.Left); err != nil {
		return err
	}
	
	// Convert left operand if needed for comparison/arithmetic operations
	// For most operations, we want both operands to be the same type
	// Promote i32 to i64 if one operand is i64
	targetType := leftType
	if leftType == I32 && rightType == I64 {
		g.emitI32ToI64Conversion(body)
		targetType = I64
	}
	
	// Generate right operand
	if err := g.generateExpression(body, expr.Right); err != nil {
		return err
	}
	
	// Convert right operand if needed
	if rightType == I32 && targetType == I64 {
		g.emitI32ToI64Conversion(body)
	} else if rightType == I64 && leftType == I32 && targetType == I32 {
		g.emitI64ToI32Conversion(body)
	}

	// Generate operator based on OperatorType
	// Use the target type to determine which instruction variant to use
	switch expr.Operator {
	case ast.OperatorAdd:
		if targetType == I32 {
			body.WriteByte(0x6A) // i32.add
		} else {
			body.WriteByte(0x7C) // i64.add
		}
	case ast.OperatorSub:
		if targetType == I32 {
			body.WriteByte(0x6B) // i32.sub
		} else {
			body.WriteByte(0x7D) // i64.sub
		}
	case ast.OperatorMul:
		if targetType == I32 {
			body.WriteByte(0x6C) // i32.mul
		} else {
			body.WriteByte(0x7E) // i64.mul
		}
	case ast.OperatorDiv:
		if targetType == I32 {
			body.WriteByte(0x6D) // i32.div_s
		} else {
			body.WriteByte(0x7F) // i64.div_s
		}
	case ast.OperatorMod:
		if targetType == I32 {
			body.WriteByte(0x6F) // i32.rem_s
		} else {
			body.WriteByte(0x81) // i64.rem_s
		}
	case ast.OperatorEq:
		if targetType == I32 {
			body.WriteByte(0x46) // i32.eq
		} else {
			body.WriteByte(0x51) // i64.eq
		}
	case ast.OperatorNe:
		if targetType == I32 {
			body.WriteByte(0x47) // i32.ne
		} else {
			body.WriteByte(0x52) // i64.ne
		}
	case ast.OperatorLt:
		if targetType == I32 {
			body.WriteByte(0x48) // i32.lt_s
		} else {
			body.WriteByte(0x53) // i64.lt_s
		}
	case ast.OperatorGt:
		if targetType == I32 {
			body.WriteByte(0x4A) // i32.gt_s
		} else {
			body.WriteByte(0x55) // i64.gt_s
		}
	case ast.OperatorLe:
		if targetType == I32 {
			body.WriteByte(0x4C) // i32.le_s
		} else {
			body.WriteByte(0x54) // i64.le_s
		}
	case ast.OperatorGe:
		if targetType == I32 {
			body.WriteByte(0x4E) // i32.ge_s
		} else {
			body.WriteByte(0x56) // i64.ge_s
		}
	case ast.OperatorAnd:
		// Logical AND for boolean values (i32)
		// Both operands are i32 (from boolean literals or comparisons)
		body.WriteByte(0x71) // i32.and
	case ast.OperatorOr:
		// Logical OR for boolean values (i32)
		// Both operands are i32 (from boolean literals or comparisons)
		body.WriteByte(0x72) // i32.or
	case ast.OperatorBitAnd:
		if targetType == I32 {
			body.WriteByte(0x71) // i32.and
		} else {
			body.WriteByte(0x83) // i64.and
		}
	case ast.OperatorBitOr:
		if targetType == I32 {
			body.WriteByte(0x72) // i32.or
		} else {
			body.WriteByte(0x84) // i64.or
		}
	case ast.OperatorBitXor:
		if targetType == I32 {
			body.WriteByte(0x73) // i32.xor
		} else {
			body.WriteByte(0x85) // i64.xor
		}
	case ast.OperatorLShift:
		if targetType == I32 {
			body.WriteByte(0x74) // i32.shl
		} else {
			body.WriteByte(0x86) // i64.shl
		}
	case ast.OperatorRShift:
		if targetType == I32 {
			body.WriteByte(0x76) // i32.shr_s
		} else {
			body.WriteByte(0x88) // i64.shr_s
		}
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
		// Logical NOT for boolean values (i32)
		// Operand is i32 (from boolean literal or comparison result)
		body.WriteByte(0x45) // i32.eqz
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

// emitI32ToI64Conversion emits instructions to convert i32 to i64
func (g *Generator) emitI32ToI64Conversion(body *bytes.Buffer) {
	// i64.extend_i32_s - sign-extend i32 to i64
	body.WriteByte(0xAC)
}

// emitI64ToI32Conversion emits instructions to convert i64 to i32
func (g *Generator) emitI64ToI32Conversion(body *bytes.Buffer) {
	// i32.wrap_i64 - wrap i64 to i32 (truncate)
	body.WriteByte(0xA7)
}

// getExpressionType returns the WASM type that an expression produces
// This is a simplified implementation - a full type inference system would be more complex
func (g *Generator) getExpressionType(expr ast.Expression) ValueType {
	switch e := expr.(type) {
	case *ast.Literal:
		switch e.Kind {
		case ast.LiteralInt:
			return I64 // Integer literals are i64
		case ast.LiteralFloat:
			return F64
		case ast.LiteralBool:
			return I32 // Booleans are i32
		default:
			return I64
		}
	case *ast.BinaryExpr:
		switch e.Operator {
		case ast.OperatorEq, ast.OperatorNe, ast.OperatorLt, ast.OperatorLe, ast.OperatorGt, ast.OperatorGe:
			return I32 // Comparison operators return i32
		case ast.OperatorAnd, ast.OperatorOr:
			return I32 // Logical operators return i32
		default:
			// Arithmetic operators return the promoted type of their operands
			// If either operand is i64, the result is i64
			leftType := g.getExpressionType(e.Left)
			rightType := g.getExpressionType(e.Right)
			if leftType == I64 || rightType == I64 {
				return I64
			} else if leftType == F64 || rightType == F64 {
				return F64
			} else if leftType == F32 || rightType == F32 {
				return F32
			}
			return I32
		}
	case *ast.UnaryExpr:
		if e.Operator == ast.OperatorNot {
			return I32 // Logical NOT returns i32
		}
		return g.getExpressionType(e.Operand)
	case *ast.Identifier:
		// Look up the variable type
		if varType, ok := g.localVarTypes[e.Name]; ok {
			return g.convertType(varType)
		}
		return I64 // Default to i64
	case *ast.CallExpr:
		// Would need to look up function return type
		// For now, default to i64
		return I64
	case *ast.StructExpr, *ast.EnumConstructorExpr, *ast.ListExpr:
		// Complex types are represented as pointers (i32 in WASM)
		return I32
	case *ast.FieldAccessExpr, *ast.IndexExpr:
		// Field access and indexing return the field/element type
		// For simplicity, default to i64
		return I64
	default:
		return I64 // Default to i64
	}
}

// generateFor generates WASM bytecode for a for loop
func (g *Generator) generateFor(body *bytes.Buffer, forStmt *ast.ForStatement) error {
	// For loops iterate over a collection
	// We need to:
	// 1. Evaluate the iterable expression (array/list)
	// 2. Get its length
	// 3. Loop from 0 to length-1
	// 4. For each iteration, load the element and bind to the loop variable
	
	// Generate the iterable expression
	if err := g.generateExpression(body, forStmt.Iterable); err != nil {
		return err
	}
	
	// Store the iterable in a temporary local
	iterableLocal := g.localVarCount
	g.localTypeOrder = append(g.localTypeOrder, I32) // Pointer to array
	g.localVarCount++
	g.emitSetLocal(body, iterableLocal)
	
	// Create index variable (starts at 0)
	indexLocal := g.localVarCount
	g.localTypeOrder = append(g.localTypeOrder, I32)
	g.localVarCount++
	body.WriteByte(0x41) // i32.const
	body.WriteByte(0x00) // 0
	g.emitSetLocal(body, indexLocal)
	
	// Get length of iterable
	// For now, we'll assume the iterable has its length at offset 0
	g.emitGetLocal(body, iterableLocal)
	body.WriteByte(0x28) // i32.load
	body.WriteByte(0x02) // alignment
	body.WriteByte(0x00) // offset
	
	// Store length in a local
	lengthLocal := g.localVarCount
	g.localTypeOrder = append(g.localTypeOrder, I32)
	g.localVarCount++
	g.emitSetLocal(body, lengthLocal)
	
	// Create loop variable
	loopVarLocal := g.localVarCount
	g.localVars[forStmt.Variable] = loopVarLocal
	g.localTypeOrder = append(g.localTypeOrder, I64) // Element type
	g.localVarCount++
	
	// block (outer - for break)
	body.WriteByte(0x02) // block
	body.WriteByte(0x40) // empty block type
	
	// loop (inner - for continue)
	body.WriteByte(0x03) // loop
	body.WriteByte(0x40) // empty block type
	
	// Check if index < length
	g.emitGetLocal(body, indexLocal)
	g.emitGetLocal(body, lengthLocal)
	body.WriteByte(0x4E) // i32.ge_s (index >= length?)
	body.WriteByte(0x0D) // br_if
	body.WriteByte(0x01) // break to outer block if true
	
	// Load current element into loop variable
	// Calculate address: iterable + 4 + (index * 8)
	g.emitGetLocal(body, iterableLocal)
	body.WriteByte(0x41) // i32.const
	body.WriteByte(0x04) // 4 (skip length)
	body.WriteByte(0x6A) // i32.add
	g.emitGetLocal(body, indexLocal)
	body.WriteByte(0x41) // i32.const
	body.WriteByte(0x08) // 8 (element size)
	body.WriteByte(0x6C) // i32.mul
	body.WriteByte(0x6A) // i32.add
	body.WriteByte(0x29) // i64.load
	body.WriteByte(0x03) // alignment
	body.WriteByte(0x00) // offset
	g.emitSetLocal(body, loopVarLocal)
	
	// Generate loop body
	for _, stmt := range forStmt.Body.Statements {
		if err := g.generateStatement(body, stmt); err != nil {
			return err
		}
	}
	
	// Increment index
	g.emitGetLocal(body, indexLocal)
	body.WriteByte(0x41) // i32.const
	body.WriteByte(0x01) // 1
	body.WriteByte(0x6A) // i32.add
	g.emitSetLocal(body, indexLocal)
	
	// Branch back to loop start
	body.WriteByte(0x0C) // br
	body.WriteByte(0x00) // continue to loop
	
	body.WriteByte(0x0B) // end loop
	body.WriteByte(0x0B) // end block
	
	return nil
}

// generateBreak generates WASM bytecode for a break statement
func (g *Generator) generateBreak(body *bytes.Buffer, breakStmt *ast.BreakStatement) error {
	// TODO: Implement proper label tracking for nested loops
	// For now, break to depth 1 (outer block)
	body.WriteByte(0x0C) // br
	body.WriteByte(0x01) // break to outer block
	return nil
}

// generateContinue generates WASM bytecode for a continue statement
func (g *Generator) generateContinue(body *bytes.Buffer, continueStmt *ast.ContinueStatement) error {
	// TODO: Implement proper label tracking for nested loops
	// For now, continue to depth 0 (loop start)
	body.WriteByte(0x0C) // br
	body.WriteByte(0x00) // continue to loop
	return nil
}

// generateStructExpr generates WASM bytecode for struct instantiation
func (g *Generator) generateStructExpr(body *bytes.Buffer, structExpr *ast.StructExpr) error {
	// Allocate memory for the struct
	// For simplicity, we'll use a fixed layout:
	// - First 4 bytes: type ID (hash of struct name)
	// - Following bytes: fields in order (8 bytes each for i64)
	
	// Calculate struct size (simplified: 4 bytes for type + 8 bytes per field)
	structSize := uint32(4 + (len(structExpr.FieldInits) * 8))
	
	// Allocate memory on heap
	g.emitHeapAlloc(body, structSize)
	
	// Duplicate pointer for storing fields
	body.WriteByte(0x22) // local.tee
	tempLocal := g.localVarCount
	g.emitULEB128(body, uint64(tempLocal))
	
	// Store type ID (simple hash of struct name)
	typeID := uint32(0)
	for _, ch := range structExpr.Name {
		typeID = typeID*31 + uint32(ch)
	}
	
	g.emitGetLocal(body, tempLocal)
	body.WriteByte(0x41) // i32.const (type ID)
	g.emitULEB128(body, uint64(typeID))
	body.WriteByte(0x36) // i32.store
	body.WriteByte(0x02) // alignment
	body.WriteByte(0x00) // offset
	
	// Store each field
	for i, field := range structExpr.FieldInits {
		// Generate field value
		if err := g.generateExpression(body, field.Value); err != nil {
			return err
		}
		
		// Store at offset (4 + i*8)
		g.emitGetLocal(body, tempLocal)
		body.WriteByte(0x37) // i64.store
		body.WriteByte(0x03) // alignment
		g.emitULEB128(body, uint64(4+i*8)) // offset
	}
	
	// Push the struct pointer as result
	g.emitGetLocal(body, tempLocal)
	
	return nil
}

// generateFieldAccess generates WASM bytecode for field access
func (g *Generator) generateFieldAccess(body *bytes.Buffer, fieldAccess *ast.FieldAccessExpr) error {
	// Generate the object expression (should produce a pointer)
	if err := g.generateExpression(body, fieldAccess.Object); err != nil {
		return err
	}
	
	// TODO: Look up field offset based on struct type
	// For now, assume fields are at fixed offsets (4 bytes for type + 8*index)
	fieldOffset := 4 // Placeholder
	
	// Load field value
	body.WriteByte(0x29) // i64.load
	body.WriteByte(0x03) // alignment
	g.emitULEB128(body, uint64(fieldOffset))
	
	return nil
}

// generateEnumConstructor generates WASM bytecode for enum construction
func (g *Generator) generateEnumConstructor(body *bytes.Buffer, enumExpr *ast.EnumConstructorExpr) error {
	// Allocate memory for the enum
	// Layout: [tag:i32][data1:i64][data2:i64]...
	
	enumSize := uint32(4 + (len(enumExpr.Arguments) * 8))
	
	// Allocate memory on heap
	g.emitHeapAlloc(body, enumSize)
	
	// Duplicate pointer for storing data
	body.WriteByte(0x22) // local.tee
	tempLocal := g.localVarCount
	g.emitULEB128(body, uint64(tempLocal))
	
	// Store tag (variant index)
	// TODO: Look up variant tag from enum definition
	variantTag := uint32(0) // Placeholder
	
	g.emitGetLocal(body, tempLocal)
	body.WriteByte(0x41) // i32.const (tag value)
	g.emitULEB128(body, uint64(variantTag))
	body.WriteByte(0x36) // i32.store
	body.WriteByte(0x02) // alignment
	body.WriteByte(0x00) // offset
	
	// Store each argument
	for i, arg := range enumExpr.Arguments {
		if err := g.generateExpression(body, arg); err != nil {
			return err
		}
		
		g.emitGetLocal(body, tempLocal)
		body.WriteByte(0x37) // i64.store
		body.WriteByte(0x03) // alignment
		g.emitULEB128(body, uint64(4+i*8))
	}
	
	// Push enum pointer
	g.emitGetLocal(body, tempLocal)
	
	return nil
}

// generateListExpr generates WASM bytecode for array/list literal
func (g *Generator) generateListExpr(body *bytes.Buffer, listExpr *ast.ListExpr) error {
	// Allocate memory for the array
	// Layout: [length:i32][elem0:i64][elem1:i64]...
	
	arraySize := uint32(4 + (len(listExpr.Elements) * 8))
	
	// Allocate memory on heap
	g.emitHeapAlloc(body, arraySize)
	
	// Duplicate pointer for storing elements
	body.WriteByte(0x22) // local.tee
	tempLocal := g.localVarCount
	g.emitULEB128(body, uint64(tempLocal))
	
	// Store length
	g.emitGetLocal(body, tempLocal)
	body.WriteByte(0x41) // i32.const (length)
	g.emitULEB128(body, uint64(len(listExpr.Elements)))
	body.WriteByte(0x36) // i32.store
	body.WriteByte(0x02) // alignment
	body.WriteByte(0x00) // offset
	
	// Store each element
	for i, elem := range listExpr.Elements {
		if err := g.generateExpression(body, elem); err != nil {
			return err
		}
		
		g.emitGetLocal(body, tempLocal)
		body.WriteByte(0x37) // i64.store
		body.WriteByte(0x03) // alignment
		g.emitULEB128(body, uint64(4+i*8))
	}
	
	// Push array pointer
	g.emitGetLocal(body, tempLocal)
	
	return nil
}

// generateIndexExpr generates WASM bytecode for array indexing
func (g *Generator) generateIndexExpr(body *bytes.Buffer, indexExpr *ast.IndexExpr) error {
	// Generate the array expression (should produce a pointer)
	if err := g.generateExpression(body, indexExpr.Object); err != nil {
		return err
	}
	
	// Store array pointer in a temp local
	tempLocal := g.localVarCount
	g.localTypeOrder = append(g.localTypeOrder, I32)
	g.localVarCount++
	g.emitSetLocal(body, tempLocal)
	
	// Generate the index expression
	if err := g.generateExpression(body, indexExpr.Index); err != nil {
		return err
	}
	
	// Convert i64 index to i32 if needed
	// TODO: Add proper type checking
	body.WriteByte(0xA7) // i32.wrap_i64
	
	// Store index in a temp local
	indexLocal := g.localVarCount
	g.localTypeOrder = append(g.localTypeOrder, I32)
	g.localVarCount++
	g.emitSetLocal(body, indexLocal)
	
	// Bounds checking: load array length and compare
	g.emitGetLocal(body, tempLocal)
	body.WriteByte(0x28) // i32.load (load length)
	body.WriteByte(0x02) // alignment
	body.WriteByte(0x00) // offset
	
	// Check if index >= length
	g.emitGetLocal(body, indexLocal)
	body.WriteByte(0x4D) // i32.le_s (length <= index?)
	
	// If true, trap (out of bounds)
	body.WriteByte(0x04) // if
	body.WriteByte(0x40) // void
	body.WriteByte(0x00) // unreachable (trap)
	body.WriteByte(0x0B) // end
	
	// Calculate element address: array + 4 + (index * 8)
	g.emitGetLocal(body, tempLocal)
	body.WriteByte(0x41) // i32.const
	body.WriteByte(0x04) // 4 (skip length)
	body.WriteByte(0x6A) // i32.add
	g.emitGetLocal(body, indexLocal)
	body.WriteByte(0x41) // i32.const
	body.WriteByte(0x08) // 8
	body.WriteByte(0x6C) // i32.mul
	body.WriteByte(0x6A) // i32.add
	
	// Load element
	body.WriteByte(0x29) // i64.load
	body.WriteByte(0x03) // alignment
	body.WriteByte(0x00) // offset
	
	return nil
}

// generateMatchExpr generates WASM bytecode for pattern matching
func (g *Generator) generateMatchExpr(body *bytes.Buffer, matchExpr *ast.MatchExpr) error {
	// Generate the value to match
	if err := g.generateExpression(body, matchExpr.Value); err != nil {
		return err
	}
	
	// Store in temporary local
	tempLocal := g.localVarCount
	g.localTypeOrder = append(g.localTypeOrder, I32) // Assuming pointer
	g.localVarCount++
	g.emitSetLocal(body, tempLocal)
	
	// Generate a series of if-else blocks for each case
	for i, caseExpr := range matchExpr.Cases {
		// TODO: Implement pattern matching logic
		// For now, only support simple patterns
		
		if i < len(matchExpr.Cases)-1 {
			// Not the last case, generate if
			// TODO: Generate pattern test
			body.WriteByte(0x41) // i32.const (placeholder condition)
			body.WriteByte(0x01)
			
			body.WriteByte(0x04) // if
			body.WriteByte(0x40) // void
		}
		
		// Generate case expression
		if err := g.generateExpression(body, caseExpr.Expression); err != nil {
			return err
		}
		
		if i < len(matchExpr.Cases)-1 {
			body.WriteByte(0x05) // else
		}
	}
	
	// Close all if blocks
	for i := 0; i < len(matchExpr.Cases)-1; i++ {
		body.WriteByte(0x0B) // end
	}
	
	return nil
}
