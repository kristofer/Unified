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
    Items []Item
    Position   Position
}

func (p *Program) Pos() Position { return p.Position }

// Item represents a top-level declaration
type Item interface {
    Node
    itemNode()
}

// More specific node types (examples)
type ModuleDecl struct {
    Name     string
    Items    []Item
    Position Position
}

func (m *ModuleDecl) Pos() Position { return m.Position }
func (m *ModuleDecl) itemNode()     {}

// FunctionDecl represents a function declaration
type FunctionDecl struct {
    Name            string
    IsPublic        bool
    Parameters      []*Parameter
    ReturnType      Type
    GenericParams   []*GenericParam
    WhereConstraints []*WhereConstraint
    Body            *Block
    Position        Position
}

func (f *FunctionDecl) Pos() Position { return f.Position }
func (f *FunctionDecl) itemNode()     {}

// Parameter represents a function parameter
type Parameter struct {
    Name         string
    Type         Type
    IsSelf       bool
    IsReference  bool
    IsMutable    bool
    Position     Position
}

func (p *Parameter) Pos() Position { return p.Position }

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

// Type is the interface for all type references
type Type interface {
    Node
    typeNode()
}

// TypeReference represents a named type possibly with generic arguments
type TypeReference struct {
    Name      string
    TypeArgs  []Type
    Position  Position
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
    Name       string
    IsPublic   bool
    Type       Type  // Only for fields
    IsMethod   bool
    Method     *FunctionDecl  // Only for methods
    Position   Position
}

func (s *StructMember) Pos() Position { return s.Position }

// Block represents a code block
type Block struct {
    Statements []Statement
    Expression Expression
    Position   Position
}

func (b *Block) Pos() Position { return b.Position }

// Statement is the interface for all statements
type Statement interface {
    Node
    statementNode()
}

// Expression is the interface for all expressions
type Expression interface {
    Node
    expressionNode()
}

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

func (l *Literal) Pos() Position { return l.Position }
func (l *Literal) expressionNode() {}

// Add more node types for functions, structs, etc.