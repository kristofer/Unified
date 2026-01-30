package wasm

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"unified-compiler/internal/ast"
)

// Generator converts AST to WASM module
type Generator struct {
	module         *Module
	localVars      map[string]int
	localVarTypes  map[string]ast.Type
	localVarCount  int
	functionIndex  int
	localTypeOrder []ValueType          // Track the order of local variable types
	currentFuncReturnType ast.Type     // Return type of current function being generated
}

// Module represents a WASM module
type Module struct {
	Functions     []Function
	FunctionTypes []FunctionType
	Exports       []Export
	Memory        *Memory
	Data          []DataSegment
	Globals       []Global
}

// Function represents a WASM function
type Function struct {
	Name       string
	TypeIndex  int
	Locals     []LocalVar
	Body       []byte
	ParamCount int
}

// FunctionType represents a function signature
type FunctionType struct {
	Params  []ValueType
	Returns []ValueType
}

// Export represents an exported function or memory
type Export struct {
	Name  string
	Kind  byte // 0=func, 1=table, 2=mem, 3=global
	Index int
}

// Memory represents linear memory configuration
type Memory struct {
	Min uint32 // Minimum pages (64KB each)
	Max uint32 // Maximum pages
}

// DataSegment represents initialized data in memory
type DataSegment struct {
	Offset uint32
	Data   []byte
}

// Global represents a global variable
type Global struct {
	Type    ValueType
	Mutable bool
	Init    []byte
}

// LocalVar represents a local variable in a function
type LocalVar struct {
	Count int
	Type  ValueType
}

// NewGenerator creates a new WASM generator
func NewGenerator() *Generator {
	return &Generator{
		module: &Module{
			Functions:     make([]Function, 0),
			FunctionTypes: make([]FunctionType, 0),
			Exports:       make([]Export, 0),
			Memory:        &Memory{Min: 1, Max: 10}, // Start with 1 page (64KB)
			Data:          make([]DataSegment, 0),
			Globals:       make([]Global, 0),
		},
		localVars:      make(map[string]int),
		localVarTypes:  make(map[string]ast.Type),
		localTypeOrder: make([]ValueType, 0),
		functionIndex:  0,
	}
}

// Generate converts an AST program to a WASM module
func (g *Generator) Generate(program *ast.Program) (*Module, error) {
	// First pass: collect function declarations
	for _, item := range program.Items {
		if fn, ok := item.(*ast.FunctionDecl); ok {
			if err := g.addFunction(fn); err != nil {
				return nil, fmt.Errorf("error adding function %s: %w", fn.Name, err)
			}
		}
	}

	// Export main function if it exists
	for i, fn := range g.module.Functions {
		if fn.Name == "main" {
			g.module.Exports = append(g.module.Exports, Export{
				Name:  "main",
				Kind:  0, // function export
				Index: i,
			})
			break
		}
	}

	return g.module, nil
}

// addFunction adds a function to the module
func (g *Generator) addFunction(fn *ast.FunctionDecl) error {
	// Reset local variable tracking for this function
	g.localVars = make(map[string]int)
	g.localVarTypes = make(map[string]ast.Type)
	g.localTypeOrder = make([]ValueType, 0)
	g.localVarCount = 0

	// Create function type
	params := make([]ValueType, len(fn.Parameters))
	for i, param := range fn.Parameters {
		params[i] = g.convertType(param.Type)
		// Add parameter as local variable
		g.localVars[param.Name] = g.localVarCount
		g.localVarTypes[param.Name] = param.Type
		// Track parameter types in order
		g.localTypeOrder = append(g.localTypeOrder, params[i])
		g.localVarCount++
	}

	returns := make([]ValueType, 0)
	if fn.ReturnType != nil {
		returns = append(returns, g.convertType(fn.ReturnType))
	}

	fnType := FunctionType{
		Params:  params,
		Returns: returns,
	}

	// Check if this function type already exists
	typeIndex := g.findOrAddFunctionType(fnType)

	// Set current function return type for type conversions
	g.currentFuncReturnType = fn.ReturnType

	// Generate function body
	body, locals, err := g.generateFunctionBody(fn)
	if err != nil {
		return err
	}

	wasmFunc := Function{
		Name:       fn.Name,
		TypeIndex:  typeIndex,
		Locals:     locals,
		Body:       body,
		ParamCount: len(fn.Parameters),
	}

	g.module.Functions = append(g.module.Functions, wasmFunc)
	g.functionIndex++

	return nil
}

// findOrAddFunctionType finds or adds a function type and returns its index
func (g *Generator) findOrAddFunctionType(ft FunctionType) int {
	// Check if type already exists
	for i, existing := range g.module.FunctionTypes {
		if g.functionTypesEqual(existing, ft) {
			return i
		}
	}

	// Add new type
	g.module.FunctionTypes = append(g.module.FunctionTypes, ft)
	return len(g.module.FunctionTypes) - 1
}

// functionTypesEqual checks if two function types are equal
func (g *Generator) functionTypesEqual(a, b FunctionType) bool {
	if len(a.Params) != len(b.Params) || len(a.Returns) != len(b.Returns) {
		return false
	}

	for i := range a.Params {
		if a.Params[i] != b.Params[i] {
			return false
		}
	}

	for i := range a.Returns {
		if a.Returns[i] != b.Returns[i] {
			return false
		}
	}

	return true
}

// generateFunctionBody generates the WASM bytecode for a function body
func (g *Generator) generateFunctionBody(fn *ast.FunctionDecl) ([]byte, []LocalVar, error) {
	var body bytes.Buffer
	locals := make([]LocalVar, 0)

	// Track local variables added during body generation
	initialLocalCount := g.localVarCount

	// Generate code for each statement in the body block
	for _, stmt := range fn.Body.Statements {
		if err := g.generateStatement(&body, stmt); err != nil {
			return nil, nil, err
		}
	}

	// Add implicit return if function doesn't end with return
	if len(fn.Body.Statements) == 0 || !g.isReturnStatement(fn.Body.Statements[len(fn.Body.Statements)-1]) {
		if fn.ReturnType != nil {
			// Push default value for return type
			g.emitDefaultValue(&body, fn.ReturnType)
		}
	}

	// Declare local variables grouped by type
	// WASM requires locals to be declared in consecutive groups of the same type
	// However, the indices remain in the order they were declared
	if g.localVarCount > initialLocalCount {
		// Group consecutive locals of the same type
		for i := initialLocalCount; i < g.localVarCount; {
			currentType := g.localTypeOrder[i]
			count := 1
			
			// Count consecutive locals of the same type
			for i+count < g.localVarCount && g.localTypeOrder[i+count] == currentType {
				count++
			}
			
			locals = append(locals, LocalVar{
				Count: count,
				Type:  currentType,
			})
			
			i += count
		}
	}

	// Add end instruction
	body.WriteByte(0x0B)

	return body.Bytes(), locals, nil
}

// convertType converts AST type to WASM ValueType
func (g *Generator) convertType(t ast.Type) ValueType {
	switch t := t.(type) {
	case *ast.TypeReference:
		switch t.Name {
		case "i32", "Int32", "bool", "Bool":
			return I32
		case "i64", "Int64", "Int":
			return I64
		case "f32", "Float32":
			return F32
		case "f64", "Float64", "Float":
			return F64
		default:
			return I64
		}
	default:
		return I64
	}
}

// isReturnStatement checks if a statement is a return statement
func (g *Generator) isReturnStatement(stmt ast.Statement) bool {
	_, ok := stmt.(*ast.ReturnStatement)
	return ok
}

// emitDefaultValue emits a default value for a type
func (g *Generator) emitDefaultValue(body *bytes.Buffer, t ast.Type) {
	vt := g.convertType(t)
	switch vt {
	case I32:
		body.WriteByte(0x41) // i32.const
		body.WriteByte(0x00) // 0
	case I64:
		body.WriteByte(0x42) // i64.const
		body.WriteByte(0x00) // 0
	case F32:
		body.WriteByte(0x43) // f32.const
		binary.Write(body, binary.LittleEndian, float32(0.0))
	case F64:
		body.WriteByte(0x44) // f64.const
		binary.Write(body, binary.LittleEndian, float64(0.0))
	}
}
