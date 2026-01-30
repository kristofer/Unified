// Code generated from UnifiedParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // UnifiedParser
import "github.com/antlr4-go/antlr/v4"

type BaseUnifiedParserVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseUnifiedParserVisitor) VisitProgram(ctx *ProgramContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseUnifiedParserVisitor) VisitItem(ctx *ItemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseUnifiedParserVisitor) VisitModuleDecl(ctx *ModuleDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseUnifiedParserVisitor) VisitImportDecl(ctx *ImportDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseUnifiedParserVisitor) VisitImportPath(ctx *ImportPathContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseUnifiedParserVisitor) VisitImportList(ctx *ImportListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseUnifiedParserVisitor) VisitFunctionDecl(ctx *FunctionDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseUnifiedParserVisitor) VisitParamList(ctx *ParamListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseUnifiedParserVisitor) VisitParameter(ctx *ParameterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseUnifiedParserVisitor) VisitGenericParams(ctx *GenericParamsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseUnifiedParserVisitor) VisitGenericParam(ctx *GenericParamContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseUnifiedParserVisitor) VisitTypeConstraint(ctx *TypeConstraintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseUnifiedParserVisitor) VisitWhereClause(ctx *WhereClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseUnifiedParserVisitor) VisitWhereConstraint(ctx *WhereConstraintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseUnifiedParserVisitor) VisitStructDecl(ctx *StructDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseUnifiedParserVisitor) VisitStructMember(ctx *StructMemberContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseUnifiedParserVisitor) VisitEnumDecl(ctx *EnumDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseUnifiedParserVisitor) VisitEnumVariant(ctx *EnumVariantContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseUnifiedParserVisitor) VisitEnumVariantParams(ctx *EnumVariantParamsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseUnifiedParserVisitor) VisitEnumVariantParam(ctx *EnumVariantParamContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseUnifiedParserVisitor) VisitInterfaceDecl(ctx *InterfaceDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseUnifiedParserVisitor) VisitInterfaceMember(ctx *InterfaceMemberContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseUnifiedParserVisitor) VisitFunctionSig(ctx *FunctionSigContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseUnifiedParserVisitor) VisitImplDecl(ctx *ImplDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseUnifiedParserVisitor) VisitImplMember(ctx *ImplMemberContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseUnifiedParserVisitor) VisitActorDecl(ctx *ActorDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseUnifiedParserVisitor) VisitActorMember(ctx *ActorMemberContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseUnifiedParserVisitor) VisitTypeAlias(ctx *TypeAliasContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseUnifiedParserVisitor) VisitConstantDecl(ctx *ConstantDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseUnifiedParserVisitor) VisitType(ctx *TypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseUnifiedParserVisitor) VisitTypeList(ctx *TypeListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseUnifiedParserVisitor) VisitStatement(ctx *StatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseUnifiedParserVisitor) VisitLetStatement(ctx *LetStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseUnifiedParserVisitor) VisitVarStatement(ctx *VarStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseUnifiedParserVisitor) VisitAssignmentStatement(ctx *AssignmentStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseUnifiedParserVisitor) VisitAssignmentOp(ctx *AssignmentOpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseUnifiedParserVisitor) VisitRegionStatement(ctx *RegionStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseUnifiedParserVisitor) VisitExprStatement(ctx *ExprStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseUnifiedParserVisitor) VisitReturnStatement(ctx *ReturnStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseUnifiedParserVisitor) VisitIfStatement(ctx *IfStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseUnifiedParserVisitor) VisitLoopStatement(ctx *LoopStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseUnifiedParserVisitor) VisitWhileStatement(ctx *WhileStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseUnifiedParserVisitor) VisitForStatement(ctx *ForStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseUnifiedParserVisitor) VisitSwitchStatement(ctx *SwitchStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseUnifiedParserVisitor) VisitCaseClause(ctx *CaseClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseUnifiedParserVisitor) VisitBreakStatement(ctx *BreakStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseUnifiedParserVisitor) VisitContinueStatement(ctx *ContinueStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseUnifiedParserVisitor) VisitBlockStatement(ctx *BlockStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseUnifiedParserVisitor) VisitPattern(ctx *PatternContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseUnifiedParserVisitor) VisitPatternList(ctx *PatternListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseUnifiedParserVisitor) VisitFieldPattern(ctx *FieldPatternContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseUnifiedParserVisitor) VisitFieldPatternList(ctx *FieldPatternListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseUnifiedParserVisitor) VisitExpr(ctx *ExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseUnifiedParserVisitor) VisitCaseExpr(ctx *CaseExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseUnifiedParserVisitor) VisitPrimary(ctx *PrimaryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseUnifiedParserVisitor) VisitLambdaExpr(ctx *LambdaExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseUnifiedParserVisitor) VisitAsyncExpr(ctx *AsyncExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseUnifiedParserVisitor) VisitConstructorExpr(ctx *ConstructorExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseUnifiedParserVisitor) VisitEnumConstructorExpr(ctx *EnumConstructorExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseUnifiedParserVisitor) VisitStructExpr(ctx *StructExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseUnifiedParserVisitor) VisitFieldInitList(ctx *FieldInitListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseUnifiedParserVisitor) VisitFieldInit(ctx *FieldInitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseUnifiedParserVisitor) VisitListExpr(ctx *ListExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseUnifiedParserVisitor) VisitMapExpr(ctx *MapExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseUnifiedParserVisitor) VisitKeyValue(ctx *KeyValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseUnifiedParserVisitor) VisitSetExpr(ctx *SetExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseUnifiedParserVisitor) VisitTupleExpr(ctx *TupleExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseUnifiedParserVisitor) VisitNamedTupleField(ctx *NamedTupleFieldContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseUnifiedParserVisitor) VisitBlock(ctx *BlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseUnifiedParserVisitor) VisitArgList(ctx *ArgListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseUnifiedParserVisitor) VisitLiteral(ctx *LiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseUnifiedParserVisitor) VisitIdentifier(ctx *IdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}
