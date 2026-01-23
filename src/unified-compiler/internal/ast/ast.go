package ast

// Position represents a source location
type Position struct {
	Line   int
	Column int
	File   string
}

// Node is the interface implemented by all AST nodes
type Node interface {
	Pos() Position
}

// Program is the root node of our AST
type Program struct {
	Items    []Item
	Position Position
}

func (p *Program) Pos() Position { return p.Position }

// Define OperatorType and its possible values
type OperatorType string

const (
	OperatorUnaryPlus  OperatorType = "UnaryPlus"
	OperatorUnaryMinus OperatorType = "UnaryMinus"
	OperatorNot        OperatorType = "Not"
	OperatorAdd        OperatorType = "Add"
	OperatorSub        OperatorType = "Sub"
	OperatorMul        OperatorType = "Mul"
	OperatorDiv        OperatorType = "Div"
	OperatorMod        OperatorType = "Mod"
	OperatorAnd        OperatorType = "And"
	OperatorOr         OperatorType = "Or"
	OperatorEq         OperatorType = "Eq"
	OperatorNe         OperatorType = "Ne"
	OperatorLt         OperatorType = "Lt"
	OperatorLe         OperatorType = "Le"
	OperatorGt         OperatorType = "Gt"
	OperatorGe         OperatorType = "Ge"
	OperatorRange      OperatorType = "Range"      // .. (exclusive)
	OperatorRangeIncl  OperatorType = "RangeIncl"  // ..= (inclusive)
)

//
// Items (Top-level declarations)
//

// Item represents a top-level declaration
type Item interface {
	Node
	itemNode()
}

// ModuleDecl represents a module declaration
type ModuleDecl struct {
	Name     string
	Items    []Item
	Position Position
}

func (m *ModuleDecl) Pos() Position { return m.Position }
func (m *ModuleDecl) itemNode()     {}

// FunctionDecl represents a function declaration
type FunctionDecl struct {
	Name             string
	IsPublic         bool
	Parameters       []*Parameter
	ReturnType       Type
	GenericParams    []*GenericParam
	WhereConstraints []*WhereConstraint
	Body             *Block
	Position         Position
}

func (f *FunctionDecl) Pos() Position { return f.Position }
func (f *FunctionDecl) itemNode()     {}

// Parameter represents a function parameter
type Parameter struct {
	Name        string
	Type        Type
	IsSelf      bool
	IsReference bool
	IsMutable   bool
	Position    Position
}

func (p *Parameter) Pos() Position { return p.Position }

// StructDecl represents a struct declaration
type StructDecl struct {
	Name          string
	IsPublic      bool
	GenericParams []*GenericParam
	Members       []*StructMember
	Position      Position
}

func (s *StructDecl) Pos() Position { return s.Position }
func (s *StructDecl) itemNode()     {}

// StructMember represents a field or method in a struct
type StructMember struct {
	Name     string
	IsPublic bool
	Type     Type // Only for fields
	IsMethod bool
	Method   *FunctionDecl // Only for methods
	Position Position
}

func (s *StructMember) Pos() Position { return s.Position }

// EnumDecl represents an enum declaration
type EnumDecl struct {
	Name          string
	IsPublic      bool
	GenericParams []*GenericParam
	Variants      []*EnumVariant
	Position      Position
}

func (e *EnumDecl) Pos() Position { return e.Position }
func (e *EnumDecl) itemNode()     {}

// EnumVariant represents a variant in an enum
type EnumVariant struct {
	Name       string
	Parameters []*EnumVariantParam
	Position   Position
}

func (e *EnumVariant) Pos() Position { return e.Position }

// EnumVariantParam represents a parameter in an enum variant
type EnumVariantParam struct {
	Name     string
	Type     Type
	Position Position
}

func (e *EnumVariantParam) Pos() Position { return e.Position }

// InterfaceDecl represents an interface declaration
type InterfaceDecl struct {
	Name          string
	IsPublic      bool
	GenericParams []*GenericParam
	Members       []*InterfaceMember
	Position      Position
}

func (i *InterfaceDecl) Pos() Position { return i.Position }
func (i *InterfaceDecl) itemNode()     {}

// InterfaceMember represents a function signature or type requirement in an interface
type InterfaceMember struct {
	IsFunction bool
	Function   *FunctionSig
	TypeName   string
	Position   Position
}

func (i *InterfaceMember) Pos() Position { return i.Position }

// FunctionSig represents a function signature in an interface
type FunctionSig struct {
	Name          string
	Parameters    []*Parameter
	ReturnType    Type
	GenericParams []*GenericParam
	HasBody       bool
	Body          *Block
	Position      Position
}

func (f *FunctionSig) Pos() Position { return f.Position }

// ImplDecl represents an implementation declaration
type ImplDecl struct {
	Type             Type
	ForType          Type // Nil if not a trait impl
	GenericParams    []*GenericParam
	WhereConstraints []*WhereConstraint
	Members          []*ImplMember
	Position         Position
}

func (i *ImplDecl) Pos() Position { return i.Position }
func (i *ImplDecl) itemNode()     {}

// ImplMember represents a function or associated type in an impl block
type ImplMember struct {
	IsFunction bool
	Function   *FunctionDecl
	TypeAlias  *TypeAlias
	Position   Position
}

func (i *ImplMember) Pos() Position { return i.Position }

// ActorDecl represents an actor declaration
type ActorDecl struct {
	Name          string
	IsPublic      bool
	GenericParams []*GenericParam
	Members       []*ActorMember
	Position      Position
}

func (a *ActorDecl) Pos() Position { return a.Position }
func (a *ActorDecl) itemNode()     {}

// ActorMember represents a field or method in an actor
type ActorMember struct {
	IsVar    bool
	Variable *VarField
	Method   *FunctionDecl
	Position Position
}

func (a *ActorMember) Pos() Position { return a.Position }

// VarField represents a variable field in an actor
type VarField struct {
	Name     string
	IsPublic bool
	Type     Type
	Value    Expression
	Position Position
}

func (v *VarField) Pos() Position { return v.Position }

// TypeAlias represents a type alias declaration
type TypeAlias struct {
	Name          string
	IsPublic      bool
	GenericParams []*GenericParam
	Type          Type
	Refinement    *TypeRefinement
	Position      Position
}

func (t *TypeAlias) Pos() Position { return t.Position }
func (t *TypeAlias) itemNode()     {}

// TypeRefinement represents a refinement constraint on a type alias
type TypeRefinement struct {
	Name      string
	Predicate Expression
	Position  Position
}

func (t *TypeRefinement) Pos() Position { return t.Position }

// ConstantDecl represents a constant declaration
type ConstantDecl struct {
	Name     string
	IsPublic bool
	Type     Type
	Value    Expression
	Position Position
}

func (c *ConstantDecl) Pos() Position { return c.Position }
func (c *ConstantDecl) itemNode()     {}

//
// Types
//

// Type is the interface for all type references
type Type interface {
	Node
	typeNode()
}

// TypeReference represents a named type possibly with generic arguments
type TypeReference struct {
	Name     string
	TypeArgs []Type
	Position Position
}

func (t *TypeReference) Pos() Position { return t.Position }
func (t *TypeReference) typeNode()     {}

// FunctionType represents a function type signature
type FunctionType struct {
	ParamTypes []Type
	ReturnType Type
	Position   Position
}

func (f *FunctionType) Pos() Position { return f.Position }
func (f *FunctionType) typeNode()     {}

// ReferenceType represents a reference type
type ReferenceType struct {
	BaseType  Type
	IsMutable bool
	Position  Position
}

func (r *ReferenceType) Pos() Position { return r.Position }
func (r *ReferenceType) typeNode()     {}

// TupleType represents a tuple type
type TupleType struct {
	ElementTypes []Type
	Position     Position
}

func (t *TupleType) Pos() Position { return t.Position }
func (t *TupleType) typeNode()     {}

// UnionType represents a union type
type UnionType struct {
	Types    []Type
	Position Position
}

func (u *UnionType) Pos() Position { return u.Position }
func (u *UnionType) typeNode()     {}

// ImplType represents an impl type
type ImplType struct {
	Interface string
	Position  Position
}

func (i *ImplType) Pos() Position { return i.Position }
func (i *ImplType) typeNode()     {}

// QualifiedType represents a qualified type
type QualifiedType struct {
	BaseType   Type
	Qualifiers []string
	Position   Position
}

func (q *QualifiedType) Pos() Position { return q.Position }
func (q *QualifiedType) typeNode()     {}

// GenericParam represents a generic type parameter
type GenericParam struct {
	Name        string
	Constraints []Type
	Position    Position
}

func (g *GenericParam) Pos() Position { return g.Position }

// WhereConstraint represents a where clause constraint
type WhereConstraint struct {
	SubjectType Type
	Constraints []Type
	Position    Position
}

func (w *WhereConstraint) Pos() Position { return w.Position }

//
// Expressions
//

// Expression is the interface for all expressions
type Expression interface {
	Node
	expressionNode()
}

// Identifier represents a variable reference
type Identifier struct {
	Name     string
	Position Position
}

func (i *Identifier) Pos() Position   { return i.Position }
func (i *Identifier) expressionNode() {}

// LiteralKind represents the type of a literal value
type LiteralKind int

const (
	LiteralInt LiteralKind = iota
	LiteralFloat
	LiteralString
	LiteralChar
	LiteralBool
	LiteralNull
)

// Literal represents a literal value in the source code
type Literal struct {
	Kind     LiteralKind
	Value    string
	Position Position
}

func (l *Literal) Pos() Position   { return l.Position }
func (l *Literal) expressionNode() {}

// BinaryExpr represents a binary operation
type BinaryExpr struct {
	Left     Expression
	Operator OperatorType
	Right    Expression
	Position Position
}

func (b *BinaryExpr) Pos() Position   { return b.Position }
func (b *BinaryExpr) expressionNode() {}

// UnaryExpr represents a unary operation
type UnaryExpr struct {
	Operator OperatorType
	Operand  Expression
	Position Position
}

func (u *UnaryExpr) Pos() Position   { return u.Position }
func (u *UnaryExpr) expressionNode() {}

// FieldAccessExpr represents a field access
type FieldAccessExpr struct {
	Object   Expression
	Field    string
	Position Position
}

func (f *FieldAccessExpr) Pos() Position   { return f.Position }
func (f *FieldAccessExpr) expressionNode() {}

// MethodCallExpr represents a method call
type MethodCallExpr struct {
	Object    Expression
	Method    string
	Arguments []Expression
	Position  Position
}

func (m *MethodCallExpr) Pos() Position   { return m.Position }
func (m *MethodCallExpr) expressionNode() {}

// IndexExpr represents an index access
type IndexExpr struct {
	Object   Expression
	Index    Expression
	Position Position
}

func (i *IndexExpr) Pos() Position   { return i.Position }
func (i *IndexExpr) expressionNode() {}

// CallExpr represents a function call
type CallExpr struct {
	Function  Expression
	Arguments []Expression
	Position  Position
}

func (c *CallExpr) Pos() Position   { return c.Position }
func (c *CallExpr) expressionNode() {}

// LambdaExpr represents a lambda expression
type LambdaExpr struct {
	Parameters []*Parameter
	ReturnType Type
	Body       interface{} // Block or Expression
	IsMoved    bool
	Position   Position
}

func (l *LambdaExpr) Pos() Position   { return l.Position }
func (l *LambdaExpr) expressionNode() {}

// AsyncExpr represents an async block
type AsyncExpr struct {
	Body     *Block
	Position Position
}

func (a *AsyncExpr) Pos() Position   { return a.Position }
func (a *AsyncExpr) expressionNode() {}

// StructExpr represents a struct instantiation
type StructExpr struct {
	Name       string
	FieldInits []*FieldInit
	Position   Position
}

func (s *StructExpr) Pos() Position   { return s.Position }
func (s *StructExpr) expressionNode() {}

// FieldInit represents a field initialization in a struct expression
type FieldInit struct {
	Name     string
	Value    Expression
	Position Position
}

func (f *FieldInit) Pos() Position { return f.Position }

// ListExpr represents a list literal
type ListExpr struct {
	Elements      []Expression
	Comprehension *Comprehension
	Position      Position
}

func (l *ListExpr) Pos() Position   { return l.Position }
func (l *ListExpr) expressionNode() {}

// Comprehension represents a list comprehension
type Comprehension struct {
	Value     Expression
	Variable  string
	Iterable  Expression
	Condition Expression
	Position  Position
}

func (c *Comprehension) Pos() Position { return c.Position }

// MapExpr represents a map literal
type MapExpr struct {
	KeyValues []*KeyValue
	Position  Position
}

func (m *MapExpr) Pos() Position   { return m.Position }
func (m *MapExpr) expressionNode() {}

// KeyValue represents a key-value pair in a map or set
type KeyValue struct {
	Key      Expression
	Value    Expression
	Position Position
}

func (k *KeyValue) Pos() Position { return k.Position }

// SetExpr represents a set literal
type SetExpr struct {
	Elements []Expression
	Position Position
}

func (s *SetExpr) Pos() Position   { return s.Position }
func (s *SetExpr) expressionNode() {}

// TupleExpr represents a tuple expression
type TupleExpr struct {
	Elements []Expression
	Position Position
}

func (t *TupleExpr) Pos() Position   { return t.Position }
func (t *TupleExpr) expressionNode() {}

// Block represents a code block
type Block struct {
	Statements []Statement
	Expression Expression
	Position   Position
}

func (b *Block) Pos() Position   { return b.Position }
func (b *Block) expressionNode() {}

//
// Statements
//

// Statement is the interface for all statements
type Statement interface {
	Node
	statementNode()
}

// LetStatement is defined in visitor.go

// VarStatement represents a variable statement
type VarStatement struct {
	Name     string
	Type     Type
	Value    Expression
	Position Position
}

func (s *VarStatement) Pos() Position  { return s.Position }
func (s *VarStatement) statementNode() {}

// RegionStatement represents a region statement
type RegionStatement struct {
	Name     string
	Body     *Block
	Position Position
}

func (s *RegionStatement) Pos() Position  { return s.Position }
func (s *RegionStatement) statementNode() {}

// ExprStatement represents an expression statement
type ExprStatement struct {
	Expression Expression
	Position   Position
}

func (s *ExprStatement) Pos() Position  { return s.Position }
func (s *ExprStatement) statementNode() {}

// ReturnStatement represents a return statement
type ReturnStatement struct {
	Value    Expression
	Position Position
}

func (s *ReturnStatement) Pos() Position  { return s.Position }
func (s *ReturnStatement) statementNode() {}

// IfBranch represents a branch in an if statement
type IfBranch struct {
	Condition Expression
	Body      *Block
}

// IfStatement represents an if statement
type IfStatement struct {
	Condition      Expression
	ThenBlock      *Block
	ElseIfBranches []*IfBranch
	ElseBlock      *Block
	Position       Position
}

func (s *IfStatement) Pos() Position  { return s.Position }
func (s *IfStatement) statementNode() {}

// LoopStatement represents a loop statement
type LoopStatement struct {
	Label    string
	Body     *Block
	Position Position
}

func (s *LoopStatement) Pos() Position  { return s.Position }
func (s *LoopStatement) statementNode() {}

// WhileStatement represents a while statement
type WhileStatement struct {
	Label     string
	Condition Expression
	Body      *Block
	Position  Position
}

func (s *WhileStatement) Pos() Position  { return s.Position }
func (s *WhileStatement) statementNode() {}

// ForStatement represents a for statement
type ForStatement struct {
	Label    string
	Variable string
	Iterable Expression
	Body     *Block
	Position Position
}

func (s *ForStatement) Pos() Position  { return s.Position }
func (s *ForStatement) statementNode() {}

// SwitchStatement represents a switch statement
type SwitchStatement struct {
	Value    Expression
	Cases    []*CaseClause
	Position Position
}

func (s *SwitchStatement) Pos() Position  { return s.Position }
func (s *SwitchStatement) statementNode() {}

// CaseClause represents a case clause in a switch statement
type CaseClause struct {
	Pattern    *Pattern
	Statements []Statement
	Position   Position
}

func (c *CaseClause) Pos() Position { return c.Position }

// PatternKind represents the type of pattern
type PatternKind int

const (
	PatternVariable PatternKind = iota
	PatternWildcard
	PatternLiteral
	PatternRange
	PatternConstructor
	PatternBinding
	PatternStruct
	PatternTuple
	PatternUnknown
)

// Pattern represents a pattern in a case clause
type Pattern struct {
	Kind          PatternKind
	Value         string          // For variable patterns
	Literal       *Literal        // For literal patterns
	Range         [2]*Pattern     // For range patterns
	FieldPatterns []*FieldPattern // For struct patterns
	SubPatterns   []*Pattern      // For tuple/constructor patterns
	Position      Position
}

func (p *Pattern) Pos() Position { return p.Position }

// FieldPattern represents a field pattern in a struct pattern
type FieldPattern struct {
	Name     string
	Pattern  *Pattern
	Position Position
}

func (f *FieldPattern) Pos() Position { return f.Position }

// BreakStatement represents a break statement
type BreakStatement struct {
	Label    string
	Position Position
}

func (s *BreakStatement) Pos() Position  { return s.Position }
func (s *BreakStatement) statementNode() {}

// ContinueStatement represents a continue statement
type ContinueStatement struct {
	Label    string
	Position Position
}

func (s *ContinueStatement) Pos() Position  { return s.Position }
func (s *ContinueStatement) statementNode() {}

// BlockStatement represents a block statement
type BlockStatement struct {
	Block    *Block
	Position Position
}

func (s *BlockStatement) Pos() Position  { return s.Position }
func (s *BlockStatement) statementNode() {}
