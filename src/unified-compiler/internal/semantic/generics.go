package semantic

import (
	"fmt"
	"strings"
	"unified-compiler/internal/ast"
)

// TypeParameter represents a generic type parameter like T in fn identity<T>
type TypeParameter struct {
	Name       string
	Bounds     []ast.Type
	Variance   Variance
	ID         int // Unique identifier for this type parameter
}

// Variance describes how a type parameter varies
type Variance int

const (
	Invariant Variance = iota
	Covariant
	Contravariant
)

// TypeVariable represents a type variable used during type inference
type TypeVariable struct {
	ID       int
	Bounds   []ast.Type
	Resolved ast.Type // nil if not yet resolved
}

// GenericContext tracks type parameters and substitutions during compilation
type GenericContext struct {
	// Map from type parameter name to its definition
	TypeParams map[string]*TypeParameter
	
	// Current substitutions (type param name -> concrete type)
	Substitutions map[string]ast.Type
	
	// Type variables for inference
	TypeVars map[int]*TypeVariable
	
	// Counter for generating unique IDs
	nextVarID int
	nextParamID int
	
	// Parent context for nested generics
	parent *GenericContext
}

// NewGenericContext creates a new generic context
func NewGenericContext() *GenericContext {
	return &GenericContext{
		TypeParams:    make(map[string]*TypeParameter),
		Substitutions: make(map[string]ast.Type),
		TypeVars:      make(map[int]*TypeVariable),
		nextVarID:     1,
		nextParamID:   1,
	}
}

// NewChildContext creates a child context for nested generics
func (gc *GenericContext) NewChildContext() *GenericContext {
	child := NewGenericContext()
	child.parent = gc
	child.nextVarID = gc.nextVarID
	child.nextParamID = gc.nextParamID
	return child
}

// AddTypeParameter adds a type parameter to the context
func (gc *GenericContext) AddTypeParameter(param *ast.GenericParam) *TypeParameter {
	tp := &TypeParameter{
		Name:     param.Name,
		Bounds:   param.Constraints,
		Variance: Invariant,
		ID:       gc.nextParamID,
	}
	gc.nextParamID++
	gc.TypeParams[param.Name] = tp
	return tp
}

// LookupTypeParam looks up a type parameter by name (checks parent contexts)
func (gc *GenericContext) LookupTypeParam(name string) *TypeParameter {
	if tp, ok := gc.TypeParams[name]; ok {
		return tp
	}
	if gc.parent != nil {
		return gc.parent.LookupTypeParam(name)
	}
	return nil
}

// NewTypeVariable creates a new type variable for inference
func (gc *GenericContext) NewTypeVariable() *TypeVariable {
	tv := &TypeVariable{
		ID:       gc.nextVarID,
		Bounds:   []ast.Type{},
		Resolved: nil,
	}
	gc.nextVarID++
	gc.TypeVars[tv.ID] = tv
	return tv
}

// Substitute applies type substitutions to a type
func (gc *GenericContext) Substitute(t ast.Type) ast.Type {
	if t == nil {
		return nil
	}
	
	switch typ := t.(type) {
	case *ast.TypeReference:
		// Check if this is a type parameter
		if sub, ok := gc.Substitutions[typ.Name]; ok {
			return sub
		}
		
		// Recursively substitute type arguments
		if len(typ.TypeArgs) > 0 {
			newArgs := make([]ast.Type, len(typ.TypeArgs))
			for i, arg := range typ.TypeArgs {
				newArgs[i] = gc.Substitute(arg)
			}
			return &ast.TypeReference{
				Name:     typ.Name,
				TypeArgs: newArgs,
				Position: typ.Position,
			}
		}
		return typ
		
	case *ast.FunctionType:
		newParams := make([]ast.Type, len(typ.ParamTypes))
		for i, param := range typ.ParamTypes {
			newParams[i] = gc.Substitute(param)
		}
		return &ast.FunctionType{
			ParamTypes: newParams,
			ReturnType: gc.Substitute(typ.ReturnType),
			Position:   typ.Position,
		}
		
	case *ast.TupleType:
		newElems := make([]ast.Type, len(typ.ElementTypes))
		for i, elem := range typ.ElementTypes {
			newElems[i] = gc.Substitute(elem)
		}
		return &ast.TupleType{
			ElementTypes: newElems,
			Position:     typ.Position,
		}
		
	case *ast.ReferenceType:
		return &ast.ReferenceType{
			BaseType:  gc.Substitute(typ.BaseType),
			IsMutable: typ.IsMutable,
			Position:  typ.Position,
		}
		
	default:
		return typ
	}
}

// Unify attempts to unify two types, updating type variable bindings
func (gc *GenericContext) Unify(t1, t2 ast.Type) error {
	// Apply current substitutions
	t1 = gc.Substitute(t1)
	t2 = gc.Substitute(t2)
	
	// If either is nil, they can't be unified
	if t1 == nil || t2 == nil {
		return fmt.Errorf("cannot unify nil types")
	}
	
	// Same type references
	ref1, isRef1 := t1.(*ast.TypeReference)
	ref2, isRef2 := t2.(*ast.TypeReference)
	
	if isRef1 && isRef2 {
		// Check if one is a type parameter
		if gc.LookupTypeParam(ref1.Name) != nil {
			// t1 is a type parameter, bind it to t2
			gc.Substitutions[ref1.Name] = t2
			return nil
		}
		if gc.LookupTypeParam(ref2.Name) != nil {
			// t2 is a type parameter, bind it to t1
			gc.Substitutions[ref2.Name] = t1
			return nil
		}
		
		// Both are concrete types
		if ref1.Name != ref2.Name {
			return fmt.Errorf("cannot unify %s with %s", ref1.Name, ref2.Name)
		}
		
		// Unify type arguments
		if len(ref1.TypeArgs) != len(ref2.TypeArgs) {
			return fmt.Errorf("type argument count mismatch for %s", ref1.Name)
		}
		
		for i := range ref1.TypeArgs {
			if err := gc.Unify(ref1.TypeArgs[i], ref2.TypeArgs[i]); err != nil {
				return err
			}
		}
		return nil
	}
	
	// Function types
	fn1, isFn1 := t1.(*ast.FunctionType)
	fn2, isFn2 := t2.(*ast.FunctionType)
	if isFn1 && isFn2 {
		if len(fn1.ParamTypes) != len(fn2.ParamTypes) {
			return fmt.Errorf("parameter count mismatch")
		}
		for i := range fn1.ParamTypes {
			if err := gc.Unify(fn1.ParamTypes[i], fn2.ParamTypes[i]); err != nil {
				return err
			}
		}
		return gc.Unify(fn1.ReturnType, fn2.ReturnType)
	}
	
	// Tuple types
	tup1, isTup1 := t1.(*ast.TupleType)
	tup2, isTup2 := t2.(*ast.TupleType)
	if isTup1 && isTup2 {
		if len(tup1.ElementTypes) != len(tup2.ElementTypes) {
			return fmt.Errorf("tuple size mismatch")
		}
		for i := range tup1.ElementTypes {
			if err := gc.Unify(tup1.ElementTypes[i], tup2.ElementTypes[i]); err != nil {
				return err
			}
		}
		return nil
	}
	
	return fmt.Errorf("cannot unify types %T and %T", t1, t2)
}

// InferTypeArguments infers type arguments from function call arguments
func (gc *GenericContext) InferTypeArguments(
	genericParams []*ast.GenericParam,
	paramTypes []ast.Type,
	argTypes []ast.Type,
) (map[string]ast.Type, error) {
	if len(paramTypes) != len(argTypes) {
		return nil, fmt.Errorf("argument count mismatch")
	}
	
	// Create a child context for inference
	inferCtx := gc.NewChildContext()
	
	// Add type parameters
	for _, gp := range genericParams {
		inferCtx.AddTypeParameter(gp)
	}
	
	// Unify each parameter type with its argument type
	for i := range paramTypes {
		if err := inferCtx.Unify(paramTypes[i], argTypes[i]); err != nil {
			return nil, fmt.Errorf("type inference failed for parameter %d: %v", i, err)
		}
	}
	
	// Extract the inferred substitutions
	result := make(map[string]ast.Type)
	for _, gp := range genericParams {
		if sub, ok := inferCtx.Substitutions[gp.Name]; ok {
			result[gp.Name] = sub
		} else {
			return nil, fmt.Errorf("could not infer type for parameter %s", gp.Name)
		}
	}
	
	return result, nil
}

// MonomorphizationKey generates a unique key for a monomorphized function
type MonomorphizationKey struct {
	FunctionName string
	TypeArgs     []string // String representations of type arguments
}

func (mk MonomorphizationKey) String() string {
	if len(mk.TypeArgs) == 0 {
		return mk.FunctionName
	}
	return fmt.Sprintf("%s<%s>", mk.FunctionName, strings.Join(mk.TypeArgs, ","))
}

// TypeToString converts a type to a string representation for monomorphization
func TypeToString(t ast.Type) string {
	if t == nil {
		return "void"
	}
	
	switch typ := t.(type) {
	case *ast.TypeReference:
		if len(typ.TypeArgs) == 0 {
			return typ.Name
		}
		args := make([]string, len(typ.TypeArgs))
		for i, arg := range typ.TypeArgs {
			args[i] = TypeToString(arg)
		}
		return fmt.Sprintf("%s<%s>", typ.Name, strings.Join(args, ","))
		
	case *ast.FunctionType:
		params := make([]string, len(typ.ParamTypes))
		for i, p := range typ.ParamTypes {
			params[i] = TypeToString(p)
		}
		return fmt.Sprintf("fn(%s)->%s", strings.Join(params, ","), TypeToString(typ.ReturnType))
		
	case *ast.TupleType:
		elems := make([]string, len(typ.ElementTypes))
		for i, e := range typ.ElementTypes {
			elems[i] = TypeToString(e)
		}
		return fmt.Sprintf("(%s)", strings.Join(elems, ","))
		
	case *ast.ReferenceType:
		prefix := "&"
		if typ.IsMutable {
			prefix = "&mut "
		}
		return prefix + TypeToString(typ.BaseType)
		
	default:
		return fmt.Sprintf("%T", typ)
	}
}

// MangleName generates a mangled name for a monomorphized function
func MangleName(baseName string, typeArgs []ast.Type) string {
	if len(typeArgs) == 0 {
		return baseName
	}
	
	argStrs := make([]string, len(typeArgs))
	for i, arg := range typeArgs {
		argStrs[i] = TypeToString(arg)
	}
	
	// Create a mangled name that's valid as an identifier
	// Replace special characters with underscores
	mangled := baseName + "_" + strings.Join(argStrs, "_")
	mangled = strings.ReplaceAll(mangled, "<", "_")
	mangled = strings.ReplaceAll(mangled, ">", "_")
	mangled = strings.ReplaceAll(mangled, ",", "_")
	mangled = strings.ReplaceAll(mangled, " ", "")
	mangled = strings.ReplaceAll(mangled, "(", "_")
	mangled = strings.ReplaceAll(mangled, ")", "_")
	mangled = strings.ReplaceAll(mangled, "&", "ref_")
	mangled = strings.ReplaceAll(mangled, "mut", "mut_")
	
	return mangled
}

// CheckConstraints verifies that a type satisfies the given constraints
func CheckConstraints(t ast.Type, constraints []ast.Type) error {
	// For now, we don't have a trait system, so just return success
	// This will be expanded when we implement traits/interfaces
	return nil
}
