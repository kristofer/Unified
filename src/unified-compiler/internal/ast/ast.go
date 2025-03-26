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

// Add more node types for functions, structs, etc.