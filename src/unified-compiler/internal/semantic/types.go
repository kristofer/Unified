package semantic

import (
	"fmt"
	"unified-compiler/internal/ast"
)

// InferType infers the type of an expression
func InferType(expr ast.Expression) (ast.Type, error) {
	switch e := expr.(type) {
	case *ast.Literal:
		return inferLiteralType(e)
	case *ast.BinaryExpr:
		return inferBinaryExprType(e)
	case *ast.UnaryExpr:
		return inferUnaryExprType(e)
	case *ast.Identifier:
		// For identifiers, we need the symbol table context
		// This will be handled by the checker
		return nil, nil
	default:
		return nil, fmt.Errorf("cannot infer type for expression: %T", expr)
	}
}

// inferLiteralType infers the type of a literal value
func inferLiteralType(lit *ast.Literal) (ast.Type, error) {
	switch lit.Kind {
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
	case ast.LiteralNull:
		return &ast.TypeReference{Name: "Null"}, nil
	default:
		return nil, fmt.Errorf("unknown literal kind: %v", lit.Kind)
	}
}

// inferBinaryExprType infers the type of a binary expression
func inferBinaryExprType(expr *ast.BinaryExpr) (ast.Type, error) {
	leftType, err := InferType(expr.Left)
	if err != nil {
		return nil, err
	}
	rightType, err := InferType(expr.Right)
	if err != nil {
		return nil, err
	}
	
	// For arithmetic operations, result type matches operands
	switch expr.Operator {
	case ast.OperatorAdd, ast.OperatorSub, ast.OperatorMul, ast.OperatorDiv, ast.OperatorMod:
		// Both operands should have the same type
		if leftType != nil && rightType != nil {
			// Simplified: assume compatible types
			return leftType, nil
		}
		return &ast.TypeReference{Name: "Int"}, nil
		
	case ast.OperatorEq, ast.OperatorNe, ast.OperatorLt, ast.OperatorLe, ast.OperatorGt, ast.OperatorGe:
		// Comparison operators return Bool
		return &ast.TypeReference{Name: "Bool"}, nil
		
	case ast.OperatorAnd, ast.OperatorOr:
		// Logical operators return Bool
		return &ast.TypeReference{Name: "Bool"}, nil
		
	default:
		return nil, fmt.Errorf("cannot infer type for operator: %v", expr.Operator)
	}
}

// inferUnaryExprType infers the type of a unary expression
func inferUnaryExprType(expr *ast.UnaryExpr) (ast.Type, error) {
	operandType, err := InferType(expr.Operand)
	if err != nil {
		return nil, err
	}
	
	switch expr.Operator {
	case ast.OperatorUnaryMinus, ast.OperatorUnaryPlus:
		// Unary +/- preserve the operand type
		if operandType != nil {
			return operandType, nil
		}
		return &ast.TypeReference{Name: "Int"}, nil
		
	case ast.OperatorNot:
		// Not operator returns Bool
		return &ast.TypeReference{Name: "Bool"}, nil
		
	default:
		return nil, fmt.Errorf("cannot infer type for unary operator: %v", expr.Operator)
	}
}

// TypesCompatible checks if two types are compatible for assignment
func TypesCompatible(target, source ast.Type) bool {
	return typesCompatibleWithDepth(target, source, 0, 100)
}

// typesCompatibleWithDepth checks type compatibility with a depth limit to prevent infinite recursion
func typesCompatibleWithDepth(target, source ast.Type, depth, maxDepth int) bool {
	// Prevent infinite recursion with circular type references
	if depth > maxDepth {
		return false
	}
	
	// Simplified type compatibility check
	// In a real implementation, this would be much more sophisticated
	
	if target == nil || source == nil {
		return true // Allow nil types for now
	}
	
	targetRef, targetOk := target.(*ast.TypeReference)
	sourceRef, sourceOk := source.(*ast.TypeReference)
	
	if targetOk && sourceOk {
		// Check base type name
		if targetRef.Name != sourceRef.Name {
			return false
		}
		
		// Check type arguments if present
		if len(targetRef.TypeArgs) != len(sourceRef.TypeArgs) {
			return false
		}
		
		// Recursively check each type argument
		for i := range targetRef.TypeArgs {
			if !typesCompatibleWithDepth(targetRef.TypeArgs[i], sourceRef.TypeArgs[i], depth+1, maxDepth) {
				return false
			}
		}
		
		return true
	}
	
	// For other types, do a simple equality check
	return true
}
