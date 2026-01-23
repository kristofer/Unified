package semantic

import (
	"fmt"
	"unified-compiler/internal/ast"
)

// SemanticError represents a semantic analysis error
type SemanticError struct {
	Message  string
	Position ast.Position
}

func (e *SemanticError) Error() string {
	return fmt.Sprintf("%s:%d:%d: %s", e.Position.File, e.Position.Line, e.Position.Column, e.Message)
}

// Checker performs semantic analysis on the AST
type Checker struct {
	symbolTable *SymbolTable
	errors      []*SemanticError
}

// NewChecker creates a new semantic checker
func NewChecker() *Checker {
	return &Checker{
		symbolTable: NewSymbolTable(),
		errors:      []*SemanticError{},
	}
}

// Check performs semantic analysis on a program
func (c *Checker) Check(program *ast.Program) error {
	// Check all items
	for _, item := range program.Items {
		if err := c.checkItem(item); err != nil {
			return err
		}
	}
	
	// Return errors if any
	if len(c.errors) > 0 {
		return c.errors[0]
	}
	
	return nil
}

// checkItem checks a top-level item
func (c *Checker) checkItem(item ast.Item) error {
	switch item := item.(type) {
	case *ast.FunctionDecl:
		return c.checkFunction(item)
	case *ast.ConstantDecl:
		// Check constant declaration
		return nil
	default:
		return nil
	}
}

// checkFunction checks a function declaration
func (c *Checker) checkFunction(fn *ast.FunctionDecl) error {
	// Enter function scope
	c.symbolTable.EnterScope()
	defer c.symbolTable.ExitScope()
	
	// Add parameters to symbol table
	for _, param := range fn.Parameters {
		if param.Name != "" {
			c.symbolTable.Define(param.Name, param.Type, false)
		}
	}
	
	// Check function body
	if fn.Body != nil {
		return c.checkBlock(fn.Body)
	}
	
	return nil
}

// checkBlock checks a block statement
func (c *Checker) checkBlock(block *ast.Block) error {
	// Check all statements
	for _, stmt := range block.Statements {
		if err := c.checkStatement(stmt); err != nil {
			return err
		}
	}
	
	// Check trailing expression
	if block.Expression != nil {
		if err := c.checkExpression(block.Expression); err != nil {
			return err
		}
	}
	
	return nil
}

// checkStatement checks a statement
func (c *Checker) checkStatement(stmt ast.Statement) error {
	switch stmt := stmt.(type) {
	case *ast.LetStatement:
		return c.checkLetStatement(stmt)
	case *ast.AssignmentStatement:
		return c.checkAssignmentStatement(stmt)
	case *ast.ExprStatement:
		return c.checkExpression(stmt.Expression)
	case *ast.ReturnStatement:
		if stmt.Value != nil {
			return c.checkExpression(stmt.Value)
		}
		return nil
	case *ast.IfStatement:
		return c.checkIfStatement(stmt)
	case *ast.WhileStatement:
		return c.checkWhileStatement(stmt)
	case *ast.BlockStatement:
		c.symbolTable.EnterScope()
		defer c.symbolTable.ExitScope()
		return c.checkBlock(stmt.Block)
	default:
		return nil
	}
}

// checkLetStatement checks a let statement
func (c *Checker) checkLetStatement(stmt *ast.LetStatement) error {
	// Check the value expression
	if err := c.checkExpression(stmt.Value); err != nil {
		return err
	}
	
	// Infer type if not specified
	var varType ast.Type
	if stmt.Type != nil {
		varType = stmt.Type
	} else {
		// Infer type from value
		inferredType, err := InferType(stmt.Value)
		if err != nil {
			return &SemanticError{
				Message:  fmt.Sprintf("cannot infer type for variable '%s': %v", stmt.Name, err),
				Position: stmt.Position,
			}
		}
		varType = inferredType
	}
	
	// Add to symbol table
	if err := c.symbolTable.Define(stmt.Name, varType, stmt.IsMutable); err != nil {
		return &SemanticError{
			Message:  err.Error(),
			Position: stmt.Position,
		}
	}
	
	return nil
}

// checkAssignmentStatement checks an assignment statement
func (c *Checker) checkAssignmentStatement(stmt *ast.AssignmentStatement) error {
	// Check that the target variable exists
	symbol, err := c.symbolTable.Lookup(stmt.Target)
	if err != nil {
		return &SemanticError{
			Message:  fmt.Sprintf("cannot assign to undefined variable '%s'", stmt.Target),
			Position: stmt.Position,
		}
	}
	
	// Check that the variable is mutable
	if !symbol.Mutable {
		return &SemanticError{
			Message:  fmt.Sprintf("cannot assign to immutable variable '%s'", stmt.Target),
			Position: stmt.Position,
		}
	}
	
	// Check the value expression
	if err := c.checkExpression(stmt.Value); err != nil {
		return err
	}
	
	// Check type compatibility
	valueType, err := InferType(stmt.Value)
	if err == nil && valueType != nil {
		if !TypesCompatible(symbol.Type, valueType) {
			return &SemanticError{
				Message:  fmt.Sprintf("type mismatch in assignment to '%s'", stmt.Target),
				Position: stmt.Position,
			}
		}
	}
	
	return nil
}

// checkIfStatement checks an if statement
func (c *Checker) checkIfStatement(stmt *ast.IfStatement) error {
	// Check condition
	if err := c.checkExpression(stmt.Condition); err != nil {
		return err
	}
	
	// Check then block
	c.symbolTable.EnterScope()
	err := c.checkBlock(stmt.ThenBlock)
	c.symbolTable.ExitScope()
	if err != nil {
		return err
	}
	
	// Check else-if branches
	for _, branch := range stmt.ElseIfBranches {
		if err := c.checkExpression(branch.Condition); err != nil {
			return err
		}
		c.symbolTable.EnterScope()
		err := c.checkBlock(branch.Body)
		c.symbolTable.ExitScope()
		if err != nil {
			return err
		}
	}
	
	// Check else block
	if stmt.ElseBlock != nil {
		c.symbolTable.EnterScope()
		err := c.checkBlock(stmt.ElseBlock)
		c.symbolTable.ExitScope()
		if err != nil {
			return err
		}
	}
	
	return nil
}

// checkWhileStatement checks a while statement
func (c *Checker) checkWhileStatement(stmt *ast.WhileStatement) error {
	// Check condition
	if err := c.checkExpression(stmt.Condition); err != nil {
		return err
	}
	
	// Check body
	c.symbolTable.EnterScope()
	defer c.symbolTable.ExitScope()
	return c.checkBlock(stmt.Body)
}

// checkExpression checks an expression
func (c *Checker) checkExpression(expr ast.Expression) error {
	switch expr := expr.(type) {
	case *ast.Identifier:
		// Check that variable is defined
		_, err := c.symbolTable.Lookup(expr.Name)
		if err != nil {
			return &SemanticError{
				Message:  fmt.Sprintf("undefined variable '%s'", expr.Name),
				Position: expr.Position,
			}
		}
		return nil
		
	case *ast.BinaryExpr:
		// Check both operands
		if err := c.checkExpression(expr.Left); err != nil {
			return err
		}
		if err := c.checkExpression(expr.Right); err != nil {
			return err
		}
		
		// Handle assignment operators
		if expr.Operator == ast.OperatorAssign {
			// This is an assignment expression, check mutability
			if ident, ok := expr.Left.(*ast.Identifier); ok {
				symbol, err := c.symbolTable.Lookup(ident.Name)
				if err != nil {
					return &SemanticError{
						Message:  fmt.Sprintf("cannot assign to undefined variable '%s'", ident.Name),
						Position: expr.Position,
					}
				}
				if !symbol.Mutable {
					return &SemanticError{
						Message:  fmt.Sprintf("cannot assign to immutable variable '%s'", ident.Name),
						Position: expr.Position,
					}
				}
			}
		}
		return nil
		
	case *ast.UnaryExpr:
		return c.checkExpression(expr.Operand)
		
	case *ast.CallExpr:
		// Check function arguments
		for _, arg := range expr.Arguments {
			if err := c.checkExpression(arg); err != nil {
				return err
			}
		}
		return nil
		
	case *ast.Literal:
		return nil
		
	default:
		return nil
	}
}

// GetSymbolTable returns the symbol table (for testing purposes)
func (c *Checker) GetSymbolTable() *SymbolTable {
	return c.symbolTable
}
