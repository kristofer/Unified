package ast

import (
	"unified-compiler/internal/parser"

	"github.com/antlr4-go/antlr/v4"
)

// ASTBuilder implements the ANTLR visitor interface to build our AST
type ASTBuilder struct {
	parser.UnifiedParserVisitor     // Embed the base visitor interface
	parser.BaseUnifiedParserVisitor // Base implementation

	filename string
}

// NewASTBuilder creates a new AST builder
func NewASTBuilder(filename string) *ASTBuilder {
	return &ASTBuilder{
		filename: filename,
	}
}

// Convert ANTLR context position to our Position
func (v *ASTBuilder) getPosition(ctx antlr.ParserRuleContext) Position {
	token := ctx.GetStart()
	return Position{
		Line:   token.GetLine(),
		Column: token.GetColumn(),
		File:   v.filename,
	}
}

// VisitProgram builds the root Program node
func (v *ASTBuilder) VisitProgram(ctx *parser.ProgramContext) interface{} {
	items := []Item{}

	// Visit all item children
	for _, itemCtx := range ctx.AllItem() {
		item := v.VisitItem(itemCtx.(*parser.ItemContext))
		if item != nil {
			items = append(items, item.(Item))
		}
	}

	return &Program{
		Items:    items,
		Position: v.getPosition(ctx),
	}
}

// VisitItem handles different item types
func (v *ASTBuilder) VisitItem(ctx *parser.ItemContext) interface{} {
	// Check which type of item we have
	if moduleCtx := ctx.ModuleDecl(); moduleCtx != nil {
		return v.VisitModuleDecl(moduleCtx.(*parser.ModuleDeclContext))
	}
	// Add handlers for other item types...

	return nil
}

// VisitModuleDecl builds a ModuleDecl node
func (v *ASTBuilder) VisitModuleDecl(ctx *parser.ModuleDeclContext) interface{} {
	// Extract module name
	identifier := ctx.Identifier().GetText()

	// Visit nested items
	items := []Item{}
	for _, itemCtx := range ctx.AllItem() {
		item := v.VisitItem(itemCtx.(*parser.ItemContext))
		if item != nil {
			items = append(items, item.(Item))
		}
	}

	return &ModuleDecl{
		Name:     identifier,
		Items:    items,
		Position: v.getPosition(ctx),
	}
}

// Add more visitor methods for other rule types...
