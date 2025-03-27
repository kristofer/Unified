package parser

import (
    "fmt"
    "github.com/antlr4-go/antlr/v4"
)

// DebugListener implements UnifiedParserListener for debugging purposes
type DebugListener struct {
    *BaseUnifiedParserListener
}

// NewDebugListener creates a new debug listener
func NewDebugListener() *DebugListener {
    return &DebugListener{
        BaseUnifiedParserListener: &BaseUnifiedParserListener{},
    }
}

// Print position helper function
func printPos(ruleName string, ctx antlr.ParserRuleContext, entering bool) {
    token := ctx.GetStart()
    direction := "→"
    if !entering {
        direction = "←"
    }
    fmt.Printf("%s %s [line %d:%d]\n", direction, ruleName, token.GetLine(), token.GetColumn())
}

// Program
func (l *DebugListener) EnterProgram(c *ProgramContext) {
    printPos("Program", c, true)
}

func (l *DebugListener) ExitProgram(c *ProgramContext) {
    printPos("Program", c, false)
}

// Item
func (l *DebugListener) EnterItem(c *ItemContext) {
    printPos("Item", c, true)
}

func (l *DebugListener) ExitItem(c *ItemContext) {
    printPos("Item", c, false)
}

// ModuleDecl
func (l *DebugListener) EnterModuleDecl(c *ModuleDeclContext) {
    printPos("ModuleDecl", c, true)
}

func (l *DebugListener) ExitModuleDecl(c *ModuleDeclContext) {
    printPos("ModuleDecl", c, false)
}

// ImportDecl
func (l *DebugListener) EnterImportDecl(c *ImportDeclContext) {
    printPos("ImportDecl", c, true)
}

func (l *DebugListener) ExitImportDecl(c *ImportDeclContext) {
    printPos("ImportDecl", c, false)
}

// ImportPath
func (l *DebugListener) EnterImportPath(c *ImportPathContext) {
    printPos("ImportPath", c, true)
}

func (l *DebugListener) ExitImportPath(c *ImportPathContext) {
    printPos("ImportPath", c, false)
}

// ImportList
func (l *DebugListener) EnterImportList(c *ImportListContext) {
    printPos("ImportList", c, true)
}

func (l *DebugListener) ExitImportList(c *ImportListContext) {
    printPos("ImportList", c, false)
}

// FunctionDecl
func (l *DebugListener) EnterFunctionDecl(c *FunctionDeclContext) {
    printPos("FunctionDecl", c, true)
}

func (l *DebugListener) ExitFunctionDecl(c *FunctionDeclContext) {
    printPos("FunctionDecl", c, false)
}

// ParamList
func (l *DebugListener) EnterParamList(c *ParamListContext) {
    printPos("ParamList", c, true)
}

func (l *DebugListener) ExitParamList(c *ParamListContext) {
    printPos("ParamList", c, false)
}

// Parameter
func (l *DebugListener) EnterParameter(c *ParameterContext) {
    printPos("Parameter", c, true)
}

func (l *DebugListener) ExitParameter(c *ParameterContext) {
    printPos("Parameter", c, false)
}

// GenericParams
func (l *DebugListener) EnterGenericParams(c *GenericParamsContext) {
    printPos("GenericParams", c, true)
}

func (l *DebugListener) ExitGenericParams(c *GenericParamsContext) {
    printPos("GenericParams", c, false)
}

// GenericParam
func (l *DebugListener) EnterGenericParam(c *GenericParamContext) {
    printPos("GenericParam", c, true)
}

func (l *DebugListener) ExitGenericParam(c *GenericParamContext) {
    printPos("GenericParam", c, false)
}

// TypeConstraint
func (l *DebugListener) EnterTypeConstraint(c *TypeConstraintContext) {
    printPos("TypeConstraint", c, true)
}

func (l *DebugListener) ExitTypeConstraint(c *TypeConstraintContext) {
    printPos("TypeConstraint", c, false)
}

// WhereClause
func (l *DebugListener) EnterWhereClause(c *WhereClauseContext) {
    printPos("WhereClause", c, true)
}

func (l *DebugListener) ExitWhereClause(c *WhereClauseContext) {
    printPos("WhereClause", c, false)
}

// WhereConstraint
func (l *DebugListener) EnterWhereConstraint(c *WhereConstraintContext) {
    printPos("WhereConstraint", c, true)
}

func (l *DebugListener) ExitWhereConstraint(c *WhereConstraintContext) {
    printPos("WhereConstraint", c, false)
}

// StructDecl
func (l *DebugListener) EnterStructDecl(c *StructDeclContext) {
    printPos("StructDecl", c, true)
}

func (l *DebugListener) ExitStructDecl(c *StructDeclContext) {
    printPos("StructDecl", c, false)
}

// StructMember
func (l *DebugListener) EnterStructMember(c *StructMemberContext) {
    printPos("StructMember", c, true)
}

func (l *DebugListener) ExitStructMember(c *StructMemberContext) {
    printPos("StructMember", c, false)
}

// EnumDecl
func (l *DebugListener) EnterEnumDecl(c *EnumDeclContext) {
    printPos("EnumDecl", c, true)
}

func (l *DebugListener) ExitEnumDecl(c *EnumDeclContext) {
    printPos("EnumDecl", c, false)
}

// EnumVariant
func (l *DebugListener) EnterEnumVariant(c *EnumVariantContext) {
    printPos("EnumVariant", c, true)
}

func (l *DebugListener) ExitEnumVariant(c *EnumVariantContext) {
    printPos("EnumVariant", c, false)
}

// EnumVariantParams
func (l *DebugListener) EnterEnumVariantParams(c *EnumVariantParamsContext) {
    printPos("EnumVariantParams", c, true)
}

func (l *DebugListener) ExitEnumVariantParams(c *EnumVariantParamsContext) {
    printPos("EnumVariantParams", c, false)
}

// EnumVariantParam
func (l *DebugListener) EnterEnumVariantParam(c *EnumVariantParamContext) {
    printPos("EnumVariantParam", c, true)
}

func (l *DebugListener) ExitEnumVariantParam(c *EnumVariantParamContext) {
    printPos("EnumVariantParam", c, false)
}

// InterfaceDecl
func (l *DebugListener) EnterInterfaceDecl(c *InterfaceDeclContext) {
    printPos("InterfaceDecl", c, true)
}

func (l *DebugListener) ExitInterfaceDecl(c *InterfaceDeclContext) {
    printPos("InterfaceDecl", c, false)
}

// InterfaceMember
func (l *DebugListener) EnterInterfaceMember(c *InterfaceMemberContext) {
    printPos("InterfaceMember", c, true)
}

func (l *DebugListener) ExitInterfaceMember(c *InterfaceMemberContext) {
    printPos("InterfaceMember", c, false)
}

// FunctionSig
func (l *DebugListener) EnterFunctionSig(c *FunctionSigContext) {
    printPos("FunctionSig", c, true)
}

func (l *DebugListener) ExitFunctionSig(c *FunctionSigContext) {
    printPos("FunctionSig", c, false)
}

// ImplDecl
func (l *DebugListener) EnterImplDecl(c *ImplDeclContext) {
    printPos("ImplDecl", c, true)
}

func (l *DebugListener) ExitImplDecl(c *ImplDeclContext) {
    printPos("ImplDecl", c, false)
}

// ImplMember
func (l *DebugListener) EnterImplMember(c *ImplMemberContext) {
    printPos("ImplMember", c, true)
}

func (l *DebugListener) ExitImplMember(c *ImplMemberContext) {
    printPos("ImplMember", c, false)
}

// ActorDecl
func (l *DebugListener) EnterActorDecl(c *ActorDeclContext) {
    printPos("ActorDecl", c, true)
}

func (l *DebugListener) ExitActorDecl(c *ActorDeclContext) {
    printPos("ActorDecl", c, false)
}

// ActorMember
func (l *DebugListener) EnterActorMember(c *ActorMemberContext) {
    printPos("ActorMember", c, true)
}

func (l *DebugListener) ExitActorMember(c *ActorMemberContext) {
    printPos("ActorMember", c, false)
}

// TypeAlias
func (l *DebugListener) EnterTypeAlias(c *TypeAliasContext) {
    printPos("TypeAlias", c, true)
}

func (l *DebugListener) ExitTypeAlias(c *TypeAliasContext) {
    printPos("TypeAlias", c, false)
}

// ConstantDecl
func (l *DebugListener) EnterConstantDecl(c *ConstantDeclContext) {
    printPos("ConstantDecl", c, true)
}

func (l *DebugListener) ExitConstantDecl(c *ConstantDeclContext) {
    printPos("ConstantDecl", c, false)
}

// Type
func (l *DebugListener) EnterType(c *TypeContext) {
    printPos("Type", c, true)
}

func (l *DebugListener) ExitType(c *TypeContext) {
    printPos("Type", c, false)
}

// TypeList
func (l *DebugListener) EnterTypeList(c *TypeListContext) {
    printPos("TypeList", c, true)
}

func (l *DebugListener) ExitTypeList(c *TypeListContext) {
    printPos("TypeList", c, false)
}

// Statement
func (l *DebugListener) EnterStatement(c *StatementContext) {
    printPos("Statement", c, true)
}

func (l *DebugListener) ExitStatement(c *StatementContext) {
    printPos("Statement", c, false)
}

// LetStatement
func (l *DebugListener) EnterLetStatement(c *LetStatementContext) {
    printPos("LetStatement", c, true)
}

func (l *DebugListener) ExitLetStatement(c *LetStatementContext) {
    printPos("LetStatement", c, false)
}

// VarStatement
func (l *DebugListener) EnterVarStatement(c *VarStatementContext) {
    printPos("VarStatement", c, true)
}

func (l *DebugListener) ExitVarStatement(c *VarStatementContext) {
    printPos("VarStatement", c, false)
}

// RegionStatement
func (l *DebugListener) EnterRegionStatement(c *RegionStatementContext) {
    printPos("RegionStatement", c, true)
}

func (l *DebugListener) ExitRegionStatement(c *RegionStatementContext) {
    printPos("RegionStatement", c, false)
}

// ExprStatement
func (l *DebugListener) EnterExprStatement(c *ExprStatementContext) {
    printPos("ExprStatement", c, true)
}

func (l *DebugListener) ExitExprStatement(c *ExprStatementContext) {
    printPos("ExprStatement", c, false)
}

// ReturnStatement
func (l *DebugListener) EnterReturnStatement(c *ReturnStatementContext) {
    printPos("ReturnStatement", c, true)
}

func (l *DebugListener) ExitReturnStatement(c *ReturnStatementContext) {
    printPos("ReturnStatement", c, false)
}

// IfStatement
func (l *DebugListener) EnterIfStatement(c *IfStatementContext) {
    printPos("IfStatement", c, true)
}

func (l *DebugListener) ExitIfStatement(c *IfStatementContext) {
    printPos("IfStatement", c, false)
}

// LoopStatement
func (l *DebugListener) EnterLoopStatement(c *LoopStatementContext) {
    printPos("LoopStatement", c, true)
}

func (l *DebugListener) ExitLoopStatement(c *LoopStatementContext) {
    printPos("LoopStatement", c, false)
}

// WhileStatement
func (l *DebugListener) EnterWhileStatement(c *WhileStatementContext) {
    printPos("WhileStatement", c, true)
}

func (l *DebugListener) ExitWhileStatement(c *WhileStatementContext) {
    printPos("WhileStatement", c, false)
}

// ForStatement
func (l *DebugListener) EnterForStatement(c *ForStatementContext) {
    printPos("ForStatement", c, true)
}

func (l *DebugListener) ExitForStatement(c *ForStatementContext) {
    printPos("ForStatement", c, false)
}

// SwitchStatement
func (l *DebugListener) EnterSwitchStatement(c *SwitchStatementContext) {
    printPos("SwitchStatement", c, true)
}

func (l *DebugListener) ExitSwitchStatement(c *SwitchStatementContext) {
    printPos("SwitchStatement", c, false)
}

// CaseClause
func (l *DebugListener) EnterCaseClause(c *CaseClauseContext) {
    printPos("CaseClause", c, true)
}

func (l *DebugListener) ExitCaseClause(c *CaseClauseContext) {
    printPos("CaseClause", c, false)
}

// BreakStatement
func (l *DebugListener) EnterBreakStatement(c *BreakStatementContext) {
    printPos("BreakStatement", c, true)
}

func (l *DebugListener) ExitBreakStatement(c *BreakStatementContext) {
    printPos("BreakStatement", c, false)
}

// ContinueStatement
func (l *DebugListener) EnterContinueStatement(c *ContinueStatementContext) {
    printPos("ContinueStatement", c, true)
}

func (l *DebugListener) ExitContinueStatement(c *ContinueStatementContext) {
    printPos("ContinueStatement", c, false)
}

// BlockStatement
func (l *DebugListener) EnterBlockStatement(c *BlockStatementContext) {
    printPos("BlockStatement", c, true)
}

func (l *DebugListener) ExitBlockStatement(c *BlockStatementContext) {
    printPos("BlockStatement", c, false)
}

// Pattern
func (l *DebugListener) EnterPattern(c *PatternContext) {
    printPos("Pattern", c, true)
}

func (l *DebugListener) ExitPattern(c *PatternContext) {
    printPos("Pattern", c, false)
}

// PatternList
func (l *DebugListener) EnterPatternList(c *PatternListContext) {
    printPos("PatternList", c, true)
}

func (l *DebugListener) ExitPatternList(c *PatternListContext) {
    printPos("PatternList", c, false)
}

// FieldPattern
func (l *DebugListener) EnterFieldPattern(c *FieldPatternContext) {
    printPos("FieldPattern", c, true)
}

func (l *DebugListener) ExitFieldPattern(c *FieldPatternContext) {
    printPos("FieldPattern", c, false)
}

// FieldPatternList
func (l *DebugListener) EnterFieldPatternList(c *FieldPatternListContext) {
    printPos("FieldPatternList", c, true)
}

func (l *DebugListener) ExitFieldPatternList(c *FieldPatternListContext) {
    printPos("FieldPatternList", c, false)
}

// Expr
func (l *DebugListener) EnterExpr(c *ExprContext) {
    printPos("Expr", c, true)
}

func (l *DebugListener) ExitExpr(c *ExprContext) {
    printPos("Expr", c, false)
}

// CaseExpr
func (l *DebugListener) EnterCaseExpr(c *CaseExprContext) {
    printPos("CaseExpr", c, true)
}

func (l *DebugListener) ExitCaseExpr(c *CaseExprContext) {
    printPos("CaseExpr", c, false)
}

// Primary
func (l *DebugListener) EnterPrimary(c *PrimaryContext) {
    printPos("Primary", c, true)
}

func (l *DebugListener) ExitPrimary(c *PrimaryContext) {
    printPos("Primary", c, false)
}

// LambdaExpr
func (l *DebugListener) EnterLambdaExpr(c *LambdaExprContext) {
    printPos("LambdaExpr", c, true)
}

func (l *DebugListener) ExitLambdaExpr(c *LambdaExprContext) {
    printPos("LambdaExpr", c, false)
}

// AsyncExpr
func (l *DebugListener) EnterAsyncExpr(c *AsyncExprContext) {
    printPos("AsyncExpr", c, true)
}

func (l *DebugListener) ExitAsyncExpr(c *AsyncExprContext) {
    printPos("AsyncExpr", c, false)
}

// StructExpr
func (l *DebugListener) EnterStructExpr(c *StructExprContext) {
    printPos("StructExpr", c, true)
}

func (l *DebugListener) ExitStructExpr(c *StructExprContext) {
    printPos("StructExpr", c, false)
}

// FieldInitList
func (l *DebugListener) EnterFieldInitList(c *FieldInitListContext) {
    printPos("FieldInitList", c, true)
}

func (l *DebugListener) ExitFieldInitList(c *FieldInitListContext) {
    printPos("FieldInitList", c, false)
}

// FieldInit
func (l *DebugListener) EnterFieldInit(c *FieldInitContext) {
    printPos("FieldInit", c, true)
}

func (l *DebugListener) ExitFieldInit(c *FieldInitContext) {
    printPos("FieldInit", c, false)
}

// ListExpr
func (l *DebugListener) EnterListExpr(c *ListExprContext) {
    printPos("ListExpr", c, true)
}

func (l *DebugListener) ExitListExpr(c *ListExprContext) {
    printPos("ListExpr", c, false)
}

// MapExpr
func (l *DebugListener) EnterMapExpr(c *MapExprContext) {
    printPos("MapExpr", c, true)
}

func (l *DebugListener) ExitMapExpr(c *MapExprContext) {
    printPos("MapExpr", c, false)
}

// KeyValue
func (l *DebugListener) EnterKeyValue(c *KeyValueContext) {
    printPos("KeyValue", c, true)
}

func (l *DebugListener) ExitKeyValue(c *KeyValueContext) {
    printPos("KeyValue", c, false)
}

// SetExpr
func (l *DebugListener) EnterSetExpr(c *SetExprContext) {
    printPos("SetExpr", c, true)
}

func (l *DebugListener) ExitSetExpr(c *SetExprContext) {
    printPos("SetExpr", c, false)
}

// TupleExpr
func (l *DebugListener) EnterTupleExpr(c *TupleExprContext) {
    printPos("TupleExpr", c, true)
}

func (l *DebugListener) ExitTupleExpr(c *TupleExprContext) {
    printPos("TupleExpr", c, false)
}

// Block
func (l *DebugListener) EnterBlock(c *BlockContext) {
    printPos("Block", c, true)
}

func (l *DebugListener) ExitBlock(c *BlockContext) {
    printPos("Block", c, false)
}

// ArgList
func (l *DebugListener) EnterArgList(c *ArgListContext) {
    printPos("ArgList", c, true)
}

func (l *DebugListener) ExitArgList(c *ArgListContext) {
    printPos("ArgList", c, false)
}

// Literal
func (l *DebugListener) EnterLiteral(c *LiteralContext) {
    printPos("Literal", c, true)
    if c.IntLiteral() != nil {
        fmt.Printf("  Int Literal: %s\n", c.IntLiteral().GetText())
    } else if c.FloatLiteral() != nil {
        fmt.Printf("  Float Literal: %s\n", c.FloatLiteral().GetText())
    } else if c.StringLiteral() != nil {
        fmt.Printf("  String Literal: %s\n", c.StringLiteral().GetText())
    } else if c.CharLiteral() != nil {
        fmt.Printf("  Char Literal: %s\n", c.CharLiteral().GetText())
    } else if c.BoolLiteral() != nil {
        fmt.Printf("  Bool Literal: %s\n", c.BoolLiteral().GetText())
    } else if c.NullLiteral() != nil {
        fmt.Printf("  Null Literal\n")
    }
}

func (l *DebugListener) ExitLiteral(c *LiteralContext) {
    printPos("Literal", c, false)
}

// Identifier
func (l *DebugListener) EnterIdentifier(c *IdentifierContext) {
    printPos("Identifier", c, true)
    fmt.Printf("  Name: %s\n", c.GetText())
}

func (l *DebugListener) ExitIdentifier(c *IdentifierContext) {
    printPos("Identifier", c, false)
}
