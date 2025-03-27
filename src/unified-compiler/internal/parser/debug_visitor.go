package parser

import (
	"fmt"
	"strings"

	"github.com/antlr4-go/antlr/v4"
)

// DebugVisitor implements UnifiedParserVisitor for debugging
type DebugVisitor struct {
	BaseUnifiedParserVisitor
	indent int
}

// NewDebugVisitor creates a new debug visitor
func NewDebugVisitor() *DebugVisitor {
	return &DebugVisitor{indent: 0}
}

// Print helper function
func (v *DebugVisitor) printNode(name string, ctx antlr.ParserRuleContext) {
	indentStr := strings.Repeat("  ", v.indent)
	startToken := ctx.GetStart()
	stopToken := ctx.GetStop()

	if stopToken == nil {
		fmt.Printf("%s%s [L%d:C%d]\n", indentStr, name,
			startToken.GetLine(), startToken.GetColumn())
	} else {
		fmt.Printf("%s%s [L%d:C%d to L%d:C%d]\n", indentStr, name,
			startToken.GetLine(), startToken.GetColumn(),
			stopToken.GetLine(), stopToken.GetColumn())
	}
}

// VisitProgram visits a program node
func (v *DebugVisitor) VisitProgram(ctx *ProgramContext) interface{} {
	v.printNode("Program", ctx)
	v.indent++

	// Visit all items
	for _, item := range ctx.AllItem() {
		item.Accept(v)
	}

	v.indent--
	return nil
}

// VisitItem visits an item node
func (v *DebugVisitor) VisitItem(ctx *ItemContext) interface{} {
	v.printNode("Item", ctx)
	v.indent++

	// Visit the appropriate child based on the item type
	if ctx.ModuleDecl() != nil {
		ctx.ModuleDecl().Accept(v)
	} else if ctx.FunctionDecl() != nil {
		ctx.FunctionDecl().Accept(v)
	} else if ctx.StructDecl() != nil {
		ctx.StructDecl().Accept(v)
	} else if ctx.EnumDecl() != nil {
		ctx.EnumDecl().Accept(v)
	} else if ctx.InterfaceDecl() != nil {
		ctx.InterfaceDecl().Accept(v)
	} else if ctx.ImplDecl() != nil {
		ctx.ImplDecl().Accept(v)
	} else if ctx.ActorDecl() != nil {
		ctx.ActorDecl().Accept(v)
	} else if ctx.TypeAlias() != nil {
		ctx.TypeAlias().Accept(v)
	} else if ctx.ConstantDecl() != nil {
		ctx.ConstantDecl().Accept(v)
	} else if ctx.ImportDecl() != nil {
		ctx.ImportDecl().Accept(v)
	}

	v.indent--
	return nil
}

// VisitModuleDecl visits a module declaration
func (v *DebugVisitor) VisitModuleDecl(ctx *ModuleDeclContext) interface{} {
	v.printNode("Module", ctx)
	v.indent++

	// Print module name
	if ctx.Identifier() != nil {
		fmt.Printf("%sName: %s\n", strings.Repeat("  ", v.indent), ctx.Identifier().GetText())
	}

	// Visit all items in the module
	for _, item := range ctx.AllItem() {
		item.Accept(v)
	}

	v.indent--
	return nil
}

// VisitImportDecl visits an import declaration
func (v *DebugVisitor) VisitImportDecl(ctx *ImportDeclContext) interface{} {
	v.printNode("Import", ctx)
	v.indent++

	if ctx.ImportPath() != nil {
		ctx.ImportPath().Accept(v)
	}

	if ctx.ImportList() != nil {
		ctx.ImportList().Accept(v)
	}

	v.indent--
	return nil
}

// VisitImportPath visits an import path
func (v *DebugVisitor) VisitImportPath(ctx *ImportPathContext) interface{} {
	v.printNode("ImportPath", ctx)
	v.indent++

	path := ctx.GetText()
	fmt.Printf("%sPath: %s\n", strings.Repeat("  ", v.indent), path)

	v.indent--
	return nil
}

// VisitImportList visits an import list
func (v *DebugVisitor) VisitImportList(ctx *ImportListContext) interface{} {
	v.printNode("ImportList", ctx)
	v.indent++

	// Visit all identifiers
	for _, id := range ctx.AllIdentifier() {
		id.Accept(v)
	}

	v.indent--
	return nil
}

// VisitFunctionDecl visits a function declaration
func (v *DebugVisitor) VisitFunctionDecl(ctx *FunctionDeclContext) interface{} {
	v.printNode("Function", ctx)
	v.indent++

	// Print visibility
	if ctx.PUB() != nil {
		fmt.Printf("%sVisibility: public\n", strings.Repeat("  ", v.indent))
	} else {
		fmt.Printf("%sVisibility: private\n", strings.Repeat("  ", v.indent))
	}

	// Print function name
	if ctx.Identifier() != nil {
		fmt.Printf("%sName: %s\n", strings.Repeat("  ", v.indent), ctx.Identifier().GetText())
	}

	// Visit generic parameters if present
	if ctx.GenericParams() != nil {
		ctx.GenericParams().Accept(v)
	}

	// Visit parameters
	if ctx.ParamList() != nil {
		ctx.ParamList().Accept(v)
	}

	// Visit return type if present
	if ctx.Type_() != nil {
		fmt.Printf("%sReturn Type:\n", strings.Repeat("  ", v.indent))
		v.indent++
		ctx.Type_().Accept(v)
		v.indent--
	}

	// Visit where clause if present
	if ctx.WhereClause() != nil {
		ctx.WhereClause().Accept(v)
	}

	// Visit function body
	if ctx.Block() != nil {
		ctx.Block().Accept(v)
	}

	v.indent--
	return nil
}

// VisitParamList visits a parameter list
func (v *DebugVisitor) VisitParamList(ctx *ParamListContext) interface{} {
	v.printNode("Parameters", ctx)
	v.indent++

	// Visit all parameters
	for _, param := range ctx.AllParameter() {
		param.Accept(v)
	}

	v.indent--
	return nil
}

// VisitParameter visits a parameter
func (v *DebugVisitor) VisitParameter(ctx *ParameterContext) interface{} {
	v.printNode("Parameter", ctx)
	v.indent++

	// Check if this is a self parameter
	if ctx.SELF(0) != nil {
		fmt.Printf("%sSelf parameter", strings.Repeat("  ", v.indent))
		if ctx.MUT() != nil {
			fmt.Printf(" (mutable)")
		}
		if ctx.BIT_AND() != nil {
			fmt.Printf(" (reference)")
		}
		fmt.Println()
	} else {
		// Regular parameter
		if ctx.Identifier() != nil {
			fmt.Printf("%sName: %s\n", strings.Repeat("  ", v.indent), ctx.Identifier().GetText())
		}

		// Visit parameter type
		if ctx.Type_() != nil {
			fmt.Printf("%sType:\n", strings.Repeat("  ", v.indent))
			v.indent++
			ctx.Type_().Accept(v)
			v.indent--
		}
	}

	v.indent--
	return nil
}

// VisitGenericParams visits generic parameters
func (v *DebugVisitor) VisitGenericParams(ctx *GenericParamsContext) interface{} {
	v.printNode("GenericParams", ctx)
	v.indent++

	// Visit all generic parameters
	for _, param := range ctx.AllGenericParam() {
		param.Accept(v)
	}

	v.indent--
	return nil
}

// VisitGenericParam visits a generic parameter
func (v *DebugVisitor) VisitGenericParam(ctx *GenericParamContext) interface{} {
	v.printNode("GenericParam", ctx)
	v.indent++

	if ctx.Identifier() != nil {
		fmt.Printf("%sName: %s\n", strings.Repeat("  ", v.indent), ctx.Identifier().GetText())
	}

	if ctx.TypeConstraint() != nil {
		ctx.TypeConstraint().Accept(v)
	}

	v.indent--
	return nil
}

// VisitTypeConstraint visits a type constraint
func (v *DebugVisitor) VisitTypeConstraint(ctx *TypeConstraintContext) interface{} {
	v.printNode("TypeConstraint", ctx)
	v.indent++

	// Visit all constraint types
	for _, typeCtx := range ctx.AllType_() {
		typeCtx.Accept(v)
	}

	v.indent--
	return nil
}

// VisitWhereClause visits a where clause
func (v *DebugVisitor) VisitWhereClause(ctx *WhereClauseContext) interface{} {
	v.printNode("WhereClause", ctx)
	v.indent++

	// Visit all where constraints
	for _, constraint := range ctx.AllWhereConstraint() {
		constraint.Accept(v)
	}

	v.indent--
	return nil
}

// VisitWhereConstraint visits a where constraint
func (v *DebugVisitor) VisitWhereConstraint(ctx *WhereConstraintContext) interface{} {
	v.printNode("WhereConstraint", ctx)
	v.indent++

	// Visit the subject type
	if ctx.Type_() != nil {
		fmt.Printf("%sSubject:\n", strings.Repeat("  ", v.indent))
		v.indent++
		ctx.Type_().Accept(v)
		v.indent--
	}

	// Visit the constraint
	if ctx.TypeConstraint() != nil {
		fmt.Printf("%sConstraint:\n", strings.Repeat("  ", v.indent))
		ctx.TypeConstraint().Accept(v)
	}

	v.indent--
	return nil
}

// VisitStructDecl visits a struct declaration
func (v *DebugVisitor) VisitStructDecl(ctx *StructDeclContext) interface{} {
	v.printNode("Struct", ctx)
	v.indent++

	// Print visibility
	if ctx.PUB() != nil {
		fmt.Printf("%sVisibility: public\n", strings.Repeat("  ", v.indent))
	} else {
		fmt.Printf("%sVisibility: private\n", strings.Repeat("  ", v.indent))
	}

	// Print struct name
	if ctx.Identifier() != nil {
		fmt.Printf("%sName: %s\n", strings.Repeat("  ", v.indent), ctx.Identifier().GetText())
	}

	// Visit generic parameters if present
	if ctx.GenericParams() != nil {
		ctx.GenericParams().Accept(v)
	}

	// Visit struct members
	fmt.Printf("%sMembers:\n", strings.Repeat("  ", v.indent))
	v.indent++
	for _, member := range ctx.AllStructMember() {
		member.Accept(v)
	}
	v.indent--

	v.indent--
	return nil
}

// VisitStructMember visits a struct member
func (v *DebugVisitor) VisitStructMember(ctx *StructMemberContext) interface{} {
	v.printNode("StructMember", ctx)
	v.indent++

	// Check if this is a field or a method
	if ctx.Identifier() != nil {
		// It's a field
		if ctx.PUB() != nil {
			fmt.Printf("%sVisibility: public\n", strings.Repeat("  ", v.indent))
		} else {
			fmt.Printf("%sVisibility: private\n", strings.Repeat("  ", v.indent))
		}

		fmt.Printf("%sField: %s\n", strings.Repeat("  ", v.indent), ctx.Identifier().GetText())

		// Visit field type
		if ctx.Type_() != nil {
			fmt.Printf("%sType:\n", strings.Repeat("  ", v.indent))
			v.indent++
			ctx.Type_().Accept(v)
			v.indent--
		}
	} else if ctx.FunctionDecl() != nil {
		// It's a method
		ctx.FunctionDecl().Accept(v)
	}

	v.indent--
	return nil
}

// VisitType_ visits a type
func (v *DebugVisitor) VisitType_(ctx *TypeContext) interface{} {
	v.printNode("Type", ctx)
	v.indent++

	if ctx.Identifier(0) != nil {
		fmt.Printf("%sName: %s\n", strings.Repeat("  ", v.indent), ctx.Identifier(0).GetText())

		// Visit type arguments if present
		if ctx.TypeList() != nil {
			fmt.Printf("%sType Arguments:\n", strings.Repeat("  ", v.indent))
			v.indent++
			ctx.TypeList().Accept(v)
			v.indent--
		}
	} else if ctx.FN() != nil {
		fmt.Printf("%sFunction Type\n", strings.Repeat("  ", v.indent))

		// Visit parameter types
		if ctx.TypeList() != nil {
			fmt.Printf("%sParameter Types:\n", strings.Repeat("  ", v.indent))
			v.indent++
			ctx.TypeList().Accept(v)
			v.indent--
		}

		// Visit return type
		if len(ctx.AllType_()) > 0 {
			fmt.Printf("%sReturn Type:\n", strings.Repeat("  ", v.indent))
			v.indent++
			ctx.Type_(0).Accept(v)
			v.indent--
		}
	}
	// Handle other type forms...

	v.indent--
	return nil
}

// VisitTypeList visits a type list
func (v *DebugVisitor) VisitTypeList(ctx *TypeListContext) interface{} {
	v.printNode("TypeList", ctx)
	v.indent++

	// Visit all types
	for _, typeCtx := range ctx.AllType_() {
		typeCtx.Accept(v)
	}

	v.indent--
	return nil
}

// VisitStatement visits a statement
func (v *DebugVisitor) VisitStatement(ctx *StatementContext) interface{} {
	v.printNode("Statement", ctx)
	v.indent++

	// Visit the appropriate statement type
	if ctx.LetStatement() != nil {
		ctx.LetStatement().Accept(v)
	} else if ctx.VarStatement() != nil {
		ctx.VarStatement().Accept(v)
	} else if ctx.RegionStatement() != nil {
		ctx.RegionStatement().Accept(v)
	} else if ctx.ExprStatement() != nil {
		ctx.ExprStatement().Accept(v)
	} else if ctx.ReturnStatement() != nil {
		ctx.ReturnStatement().Accept(v)
	} else if ctx.IfStatement() != nil {
		ctx.IfStatement().Accept(v)
	} else if ctx.LoopStatement() != nil {
		ctx.LoopStatement().Accept(v)
	} else if ctx.WhileStatement() != nil {
		ctx.WhileStatement().Accept(v)
	} else if ctx.ForStatement() != nil {
		ctx.ForStatement().Accept(v)
	} else if ctx.SwitchStatement() != nil {
		ctx.SwitchStatement().Accept(v)
	} else if ctx.BreakStatement() != nil {
		ctx.BreakStatement().Accept(v)
	} else if ctx.ContinueStatement() != nil {
		ctx.ContinueStatement().Accept(v)
	} else if ctx.BlockStatement() != nil {
		ctx.BlockStatement().Accept(v)
	}

	v.indent--
	return nil
}

// VisitBlock visits a block
func (v *DebugVisitor) VisitBlock(ctx *BlockContext) interface{} {
	v.printNode("Block", ctx)
	v.indent++

	// Visit all statements
	for _, stmt := range ctx.AllStatement() {
		stmt.Accept(v)
	}

	// Visit trailing expression if present
	if ctx.Expr() != nil {
		fmt.Printf("%sTrailing Expression:\n", strings.Repeat("  ", v.indent))
		ctx.Expr().Accept(v)
	}

	v.indent--
	return nil
}

// VisitExpr visits an expression
func (v *DebugVisitor) VisitExpr(ctx *ExprContext) interface{} {
	v.printNode("Expression", ctx)
	v.indent++

	// Check expression type based on context
	if ctx.Primary() != nil {
		ctx.Primary().Accept(v)
	} else if len(ctx.AllExpr()) == 2 {
		// Binary operation
		operator := "unknown"
		if ctx.PLUS() != nil {
			operator = "+"
		} else if ctx.MINUS() != nil {
			operator = "-"
		} else if ctx.STAR() != nil {
			operator = "*"
		} else if ctx.DIV() != nil {
			operator = "/"
		} else if ctx.MOD() != nil {
			operator = "%"
		} else if ctx.EQ() != nil {
			operator = "=="
		} else if ctx.NE() != nil {
			operator = "!="
		} else if ctx.LT() != nil {
			operator = "<"
		} else if ctx.LE() != nil {
			operator = "<="
		} else if ctx.GT() != nil {
			operator = ">"
		} else if ctx.GE() != nil {
			operator = ">="
		} else if ctx.AND() != nil {
			operator = "&&"
		} else if ctx.OR() != nil {
			operator = "||"
		} else if ctx.BIT_AND() != nil {
			operator = "&"
		} else if ctx.BIT_OR() != nil {
			operator = "|"
		} else if ctx.BIT_XOR() != nil {
			operator = "^"
		}

		fmt.Printf("%sBinary Op: %s\n", strings.Repeat("  ", v.indent), operator)
		fmt.Printf("%sLeft:\n", strings.Repeat("  ", v.indent))
		v.indent++
		ctx.Expr(0).Accept(v)
		v.indent--

		fmt.Printf("%sRight:\n", strings.Repeat("  ", v.indent))
		v.indent++
		ctx.Expr(1).Accept(v)
		v.indent--
	} else if len(ctx.AllExpr()) == 1 {
		// Unary operation or member access
		if ctx.PLUS() != nil {
			fmt.Printf("%sUnary +\n", strings.Repeat("  ", v.indent))
		} else if ctx.MINUS() != nil {
			fmt.Printf("%sUnary -\n", strings.Repeat("  ", v.indent))
		} else if ctx.NOT() != nil {
			fmt.Printf("%sUnary !\n", strings.Repeat("  ", v.indent))
		} else if ctx.DOT() != nil && ctx.Identifier() != nil {
			fmt.Printf("%sField Access: %s\n", strings.Repeat("  ", v.indent), ctx.Identifier().GetText())
		}

		ctx.Expr(0).Accept(v)
	}

	v.indent--
	return nil
}

// VisitPrimary visits a primary expression
func (v *DebugVisitor) VisitPrimary(ctx *PrimaryContext) interface{} {
	v.printNode("Primary", ctx)
	v.indent++

	if ctx.Literal() != nil {
		ctx.Literal().Accept(v)
	} else if ctx.Identifier() != nil {
		ctx.Identifier().Accept(v)
	} else if ctx.LPAREN() != nil && ctx.Expr() != nil {
		fmt.Printf("%sParenthesized Expression:\n", strings.Repeat("  ", v.indent))
		ctx.Expr().Accept(v)
	}
	// Handle other primary expressions...

	v.indent--
	return nil
}

// VisitLiteral visits a literal
func (v *DebugVisitor) VisitLiteral(ctx *LiteralContext) interface{} {
	v.printNode("Literal", ctx)
	v.indent++

	if ctx.IntLiteral() != nil {
		fmt.Printf("%sInteger: %s\n", strings.Repeat("  ", v.indent), ctx.IntLiteral().GetText())
	} else if ctx.FloatLiteral() != nil {
		fmt.Printf("%sFloat: %s\n", strings.Repeat("  ", v.indent), ctx.FloatLiteral().GetText())
	} else if ctx.StringLiteral() != nil {
		fmt.Printf("%sString: %s\n", strings.Repeat("  ", v.indent), ctx.StringLiteral().GetText())
	} else if ctx.CharLiteral() != nil {
		fmt.Printf("%sChar: %s\n", strings.Repeat("  ", v.indent), ctx.CharLiteral().GetText())
	} else if ctx.BoolLiteral() != nil {
		fmt.Printf("%sBool: %s\n", strings.Repeat("  ", v.indent), ctx.BoolLiteral().GetText())
	} else if ctx.NullLiteral() != nil {
		fmt.Printf("%sNull\n", strings.Repeat("  ", v.indent))
	}

	v.indent--
	return nil
}

// VisitIdentifier visits an identifier
func (v *DebugVisitor) VisitIdentifier(ctx *IdentifierContext) interface{} {
	v.printNode("Identifier", ctx)
	v.indent++

	fmt.Printf("%sName: %s\n", strings.Repeat("  ", v.indent), ctx.GetText())

	v.indent--
	return nil
}

// Helper method to implement remaining visitor methods
func (v *DebugVisitor) visitDefault(name string, ctx antlr.ParserRuleContext) interface{} {
	v.printNode(name, ctx)
	return nil
}

// Remaining method implementations - these are simplified but follow the pattern
func (v *DebugVisitor) VisitEnumDecl(ctx *EnumDeclContext) interface{} {
	return v.visitDefault("EnumDecl", ctx)
}
func (v *DebugVisitor) VisitEnumVariant(ctx *EnumVariantContext) interface{} {
	return v.visitDefault("EnumVariant", ctx)
}
func (v *DebugVisitor) VisitEnumVariantParams(ctx *EnumVariantParamsContext) interface{} {
	return v.visitDefault("EnumVariantParams", ctx)
}
func (v *DebugVisitor) VisitEnumVariantParam(ctx *EnumVariantParamContext) interface{} {
	return v.visitDefault("EnumVariantParam", ctx)
}
func (v *DebugVisitor) VisitInterfaceDecl(ctx *InterfaceDeclContext) interface{} {
	return v.visitDefault("InterfaceDecl", ctx)
}
func (v *DebugVisitor) VisitInterfaceMember(ctx *InterfaceMemberContext) interface{} {
	return v.visitDefault("InterfaceMember", ctx)
}
func (v *DebugVisitor) VisitFunctionSig(ctx *FunctionSigContext) interface{} {
	return v.visitDefault("FunctionSig", ctx)
}
func (v *DebugVisitor) VisitImplDecl(ctx *ImplDeclContext) interface{} {
	return v.visitDefault("ImplDecl", ctx)
}
func (v *DebugVisitor) VisitImplMember(ctx *ImplMemberContext) interface{} {
	return v.visitDefault("ImplMember", ctx)
}
func (v *DebugVisitor) VisitActorDecl(ctx *ActorDeclContext) interface{} {
	return v.visitDefault("ActorDecl", ctx)
}
func (v *DebugVisitor) VisitActorMember(ctx *ActorMemberContext) interface{} {
	return v.visitDefault("ActorMember", ctx)
}
func (v *DebugVisitor) VisitTypeAlias(ctx *TypeAliasContext) interface{} {
	return v.visitDefault("TypeAlias", ctx)
}
func (v *DebugVisitor) VisitConstantDecl(ctx *ConstantDeclContext) interface{} {
	return v.visitDefault("ConstantDecl", ctx)
}
func (v *DebugVisitor) VisitLetStatement(ctx *LetStatementContext) interface{} {
	return v.visitDefault("LetStatement", ctx)
}
func (v *DebugVisitor) VisitVarStatement(ctx *VarStatementContext) interface{} {
	return v.visitDefault("VarStatement", ctx)
}
func (v *DebugVisitor) VisitRegionStatement(ctx *RegionStatementContext) interface{} {
	return v.visitDefault("RegionStatement", ctx)
}
func (v *DebugVisitor) VisitExprStatement(ctx *ExprStatementContext) interface{} {
	return v.visitDefault("ExprStatement", ctx)
}
func (v *DebugVisitor) VisitReturnStatement(ctx *ReturnStatementContext) interface{} {
	return v.visitDefault("ReturnStatement", ctx)
}
func (v *DebugVisitor) VisitIfStatement(ctx *IfStatementContext) interface{} {
	return v.visitDefault("IfStatement", ctx)
}
func (v *DebugVisitor) VisitLoopStatement(ctx *LoopStatementContext) interface{} {
	return v.visitDefault("LoopStatement", ctx)
}
func (v *DebugVisitor) VisitWhileStatement(ctx *WhileStatementContext) interface{} {
	return v.visitDefault("WhileStatement", ctx)
}
func (v *DebugVisitor) VisitForStatement(ctx *ForStatementContext) interface{} {
	return v.visitDefault("ForStatement", ctx)
}
func (v *DebugVisitor) VisitSwitchStatement(ctx *SwitchStatementContext) interface{} {
	return v.visitDefault("SwitchStatement", ctx)
}
func (v *DebugVisitor) VisitCaseClause(ctx *CaseClauseContext) interface{} {
	return v.visitDefault("CaseClause", ctx)
}
func (v *DebugVisitor) VisitBreakStatement(ctx *BreakStatementContext) interface{} {
	return v.visitDefault("BreakStatement", ctx)
}
func (v *DebugVisitor) VisitContinueStatement(ctx *ContinueStatementContext) interface{} {
	return v.visitDefault("ContinueStatement", ctx)
}
func (v *DebugVisitor) VisitBlockStatement(ctx *BlockStatementContext) interface{} {
	return v.visitDefault("BlockStatement", ctx)
}

func (v *DebugVisitor) VisitPattern(ctx *PatternContext) interface{} {
	return v.visitDefault("Pattern", ctx)
}
func (v *DebugVisitor) VisitPatternList(ctx *PatternListContext) interface{} {
	return v.visitDefault("PatternList", ctx)
}
func (v *DebugVisitor) VisitFieldPattern(ctx *FieldPatternContext) interface{} {
	return v.visitDefault("FieldPattern", ctx)
}
func (v *DebugVisitor) VisitFieldPatternList(ctx *FieldPatternListContext) interface{} {
	return v.visitDefault("FieldPatternList", ctx)
}
func (v *DebugVisitor) VisitCaseExpr(ctx *CaseExprContext) interface{} {
	return v.visitDefault("CaseExpr", ctx)
}
func (v *DebugVisitor) VisitLambdaExpr(ctx *LambdaExprContext) interface{} {
	return v.visitDefault("LambdaExpr", ctx)
}
func (v *DebugVisitor) VisitAsyncExpr(ctx *AsyncExprContext) interface{} {
	return v.visitDefault("AsyncExpr", ctx)
}
func (v *DebugVisitor) VisitStructExpr(ctx *StructExprContext) interface{} {
	return v.visitDefault("StructExpr", ctx)
}
func (v *DebugVisitor) VisitFieldInitList(ctx *FieldInitListContext) interface{} {
	return v.visitDefault("FieldInitList", ctx)
}
func (v *DebugVisitor) VisitFieldInit(ctx *FieldInitContext) interface{} {
	return v.visitDefault("FieldInit", ctx)
}
func (v *DebugVisitor) VisitListExpr(ctx *ListExprContext) interface{} {
	return v.visitDefault("ListExpr", ctx)
}
func (v *DebugVisitor) VisitMapExpr(ctx *MapExprContext) interface{} {
	return v.visitDefault("MapExpr", ctx)
}
func (v *DebugVisitor) VisitKeyValue(ctx *KeyValueContext) interface{} {
	return v.visitDefault("KeyValue", ctx)
}
func (v *DebugVisitor) VisitSetExpr(ctx *SetExprContext) interface{} {
	return v.visitDefault("SetExpr", ctx)
}
func (v *DebugVisitor) VisitTupleExpr(ctx *TupleExprContext) interface{} {
	return v.visitDefault("TupleExpr", ctx)
}
func (v *DebugVisitor) VisitArgList(ctx *ArgListContext) interface{} {
	return v.visitDefault("ArgList", ctx)
}
