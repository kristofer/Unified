package ast

// Node is the interface implemented by all AST nodes
type Node interface {
    // Position returns the source position information for this node
    Position() Position
}

// Position represents a position in source code
type Position struct {
    Line   int
    Column int
}

// Program is the root of the AST
type Program struct {
    Items []Item
    Pos   Position
}

func (p *Program) Position() Position {
    return p.Pos
}

// Item represents a top-level declaration
type Item interface {
    Node
    isItem()
}
