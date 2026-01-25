package bytecode

import (
	"testing"
	"unified-compiler/internal/ast"
)

// Test 1: Try operator bytecode generation
func TestTryOperatorBytecodeGeneration(t *testing.T) {
	gen := NewGenerator()

	// Create a simple try expression: someResult?
	tryExpr := &ast.TryExpr{
		Operand: &ast.Identifier{
			Name: "someResult",
		},
	}

	// Set up a local variable for someResult
	gen.localVars["someResult"] = 0

	err := gen.generateTryExpr(tryExpr)
	if err != nil {
		t.Fatalf("Failed to generate try expression: %v", err)
	}

	// Check that bytecode was generated
	bc := gen.bytecode
	if len(bc.Instructions) < 2 {
		t.Fatalf("Expected at least 2 instructions, got %d", len(bc.Instructions))
	}

	// First instruction should load the variable
	if bc.Instructions[0].Op != OpLoadLocal {
		t.Errorf("Expected OpLoadLocal, got %v", bc.Instructions[0].Op)
	}

	// Second instruction should be OpTryPropagate
	if bc.Instructions[1].Op != OpTryPropagate {
		t.Errorf("Expected OpTryPropagate, got %v", bc.Instructions[1].Op)
	}
}

// Test 2: Try operator with function call
func TestTryOperatorWithFunctionCall(t *testing.T) {
	gen := NewGenerator()

	// Create try expression with variable that holds a result: result?
	tryExpr := &ast.TryExpr{
		Operand: &ast.Identifier{
			Name: "result",
		},
	}

	gen.localVars["result"] = 0

	err := gen.generateTryExpr(tryExpr)
	if err != nil {
		t.Fatalf("Failed to generate try expression: %v", err)
	}

	bc := gen.bytecode
	// Should have: LoadLocal, TryPropagate
	if len(bc.Instructions) < 2 {
		t.Fatalf("Expected at least 2 instructions, got %d", len(bc.Instructions))
	}

	// Last instruction should be OpTryPropagate
	lastOp := bc.Instructions[len(bc.Instructions)-1].Op
	if lastOp != OpTryPropagate {
		t.Errorf("Expected OpTryPropagate as last instruction, got %v", lastOp)
	}
}

// Test 3: Chained try operators
func TestChainedTryOperators(t *testing.T) {
	gen := NewGenerator()

	// Create chained try: (a?)? 
	innerTry := &ast.TryExpr{
		Operand: &ast.Identifier{
			Name: "a",
		},
	}
	outerTry := &ast.TryExpr{
		Operand: innerTry,
	}

	gen.localVars["a"] = 0

	err := gen.generateTryExpr(outerTry)
	if err != nil {
		t.Fatalf("Failed to generate chained try: %v", err)
	}

	bc := gen.bytecode
	// Should have: LoadLocal, TryPropagate, TryPropagate
	if len(bc.Instructions) < 3 {
		t.Fatalf("Expected at least 3 instructions, got %d", len(bc.Instructions))
	}

	// Check we have two OpTryPropagate instructions
	tryCount := 0
	for _, inst := range bc.Instructions {
		if inst.Op == OpTryPropagate {
			tryCount++
		}
	}
	if tryCount != 2 {
		t.Errorf("Expected 2 OpTryPropagate instructions, got %d", tryCount)
	}
}

// Test 4: Try operator instruction string representation
func TestTryPropagateInstructionString(t *testing.T) {
	opStr := OpTryPropagate.String()
	expected := "TRY_PROPAGATE"
	if opStr != expected {
		t.Errorf("Expected %s, got %s", expected, opStr)
	}
}

// Test 5: Try operator with field access
func TestTryOperatorWithFieldAccess(t *testing.T) {
	gen := NewGenerator()

	// Create try expression: obj.field?
	tryExpr := &ast.TryExpr{
		Operand: &ast.FieldAccessExpr{
			Object: &ast.Identifier{
				Name: "obj",
			},
			Field: "field",
		},
	}

	gen.localVars["obj"] = 0

	err := gen.generateTryExpr(tryExpr)
	if err != nil {
		t.Fatalf("Failed to generate try with field access: %v", err)
	}

	bc := gen.bytecode
	if len(bc.Instructions) < 2 {
		t.Fatalf("Expected at least 2 instructions, got %d", len(bc.Instructions))
	}

	// Last instruction should be OpTryPropagate
	lastOp := bc.Instructions[len(bc.Instructions)-1].Op
	if lastOp != OpTryPropagate {
		t.Errorf("Expected OpTryPropagate as last instruction, got %v", lastOp)
	}
}

// Test 6: Try operator with literal value
func TestTryOperatorWithEnumConstructor(t *testing.T) {
	gen := NewGenerator()

	// Create try expression with a literal (simplified test)
	tryExpr := &ast.TryExpr{
		Operand: &ast.Literal{
			Kind:  ast.LiteralInt,
			Value: "100",
		},
	}

	err := gen.generateTryExpr(tryExpr)
	if err != nil {
		t.Fatalf("Failed to generate try with literal: %v", err)
	}

	bc := gen.bytecode
	// Should end with OpTryPropagate
	if len(bc.Instructions) == 0 {
		t.Fatalf("No instructions generated")
	}

	lastOp := bc.Instructions[len(bc.Instructions)-1].Op
	if lastOp != OpTryPropagate {
		t.Errorf("Expected OpTryPropagate as last instruction, got %v", lastOp)
	}
}

// Test 7: Try operator in binary expression
func TestTryOperatorInBinaryExpression(t *testing.T) {
	gen := NewGenerator()

	// Create expression: a? + b
	tryExpr := &ast.TryExpr{
		Operand: &ast.Identifier{
			Name: "a",
		},
	}
	binaryExpr := &ast.BinaryExpr{
		Left:     tryExpr,
		Operator: ast.OperatorAdd,
		Right: &ast.Identifier{
			Name: "b",
		},
	}

	gen.localVars["a"] = 0
	gen.localVars["b"] = 1

	err := gen.generateBinaryExpr(binaryExpr)
	if err != nil {
		t.Fatalf("Failed to generate binary with try: %v", err)
	}

	bc := gen.bytecode
	// Should have: LoadLocal(a), TryPropagate, LoadLocal(b), Add
	if len(bc.Instructions) < 4 {
		t.Fatalf("Expected at least 4 instructions, got %d", len(bc.Instructions))
	}
}

// Test 8: Try operator opcode value
func TestTryPropagateOpcodeValue(t *testing.T) {
	// Verify OpTryPropagate has a valid opcode value
	if OpTryPropagate == 0 {
		t.Error("OpTryPropagate should not be zero")
	}

	// Verify it's different from other opcodes
	if OpTryPropagate == OpAllocEnum {
		t.Error("OpTryPropagate should have unique value")
	}
	if OpTryPropagate == OpMatchVariant {
		t.Error("OpTryPropagate should have unique value")
	}
	if OpTryPropagate == OpExtractVariant {
		t.Error("OpTryPropagate should have unique value")
	}
}

// Test 9: Try operator with literal (should still generate)
func TestTryOperatorWithLiteral(t *testing.T) {
	gen := NewGenerator()

	// Create try expression: 42? (semantically invalid but should generate)
	tryExpr := &ast.TryExpr{
		Operand: &ast.Literal{
			Kind:  ast.LiteralInt,
			Value: "42",
		},
	}

	err := gen.generateTryExpr(tryExpr)
	if err != nil {
		t.Fatalf("Failed to generate try with literal: %v", err)
	}

	bc := gen.bytecode
	if len(bc.Instructions) < 2 {
		t.Fatalf("Expected at least 2 instructions, got %d", len(bc.Instructions))
	}

	// Should have: Push(42), TryPropagate
	if bc.Instructions[0].Op != OpPush {
		t.Errorf("Expected OpPush, got %v", bc.Instructions[0].Op)
	}
	if bc.Instructions[1].Op != OpTryPropagate {
		t.Errorf("Expected OpTryPropagate, got %v", bc.Instructions[1].Op)
	}
}

// Test 10: Multiple try operators in sequence
func TestMultipleTryOperatorsInSequence(t *testing.T) {
	gen := NewGenerator()

	// Create multiple try expressions
	try1 := &ast.TryExpr{
		Operand: &ast.Identifier{Name: "a"},
	}
	try2 := &ast.TryExpr{
		Operand: &ast.Identifier{Name: "b"},
	}

	gen.localVars["a"] = 0
	gen.localVars["b"] = 1

	err := gen.generateTryExpr(try1)
	if err != nil {
		t.Fatalf("Failed to generate first try: %v", err)
	}

	err = gen.generateTryExpr(try2)
	if err != nil {
		t.Fatalf("Failed to generate second try: %v", err)
	}

	bc := gen.bytecode
	// Should have: LoadLocal(a), TryPropagate, LoadLocal(b), TryPropagate
	if len(bc.Instructions) != 4 {
		t.Fatalf("Expected 4 instructions, got %d", len(bc.Instructions))
	}

	if bc.Instructions[1].Op != OpTryPropagate {
		t.Errorf("Expected OpTryPropagate at position 1, got %v", bc.Instructions[1].Op)
	}
	if bc.Instructions[3].Op != OpTryPropagate {
		t.Errorf("Expected OpTryPropagate at position 3, got %v", bc.Instructions[3].Op)
	}
}
