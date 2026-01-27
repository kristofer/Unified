// Code generated from UnifiedParser.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // UnifiedParser
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by UnifiedParser.
type UnifiedParserVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by UnifiedParser#program.
	VisitProgram(ctx *ProgramContext) interface{}

	// Visit a parse tree produced by UnifiedParser#item.
	VisitItem(ctx *ItemContext) interface{}

	// Visit a parse tree produced by UnifiedParser#moduleDecl.
	VisitModuleDecl(ctx *ModuleDeclContext) interface{}

	// Visit a parse tree produced by UnifiedParser#importDecl.
	VisitImportDecl(ctx *ImportDeclContext) interface{}

	// Visit a parse tree produced by UnifiedParser#importPath.
	VisitImportPath(ctx *ImportPathContext) interface{}

	// Visit a parse tree produced by UnifiedParser#importList.
	VisitImportList(ctx *ImportListContext) interface{}

	// Visit a parse tree produced by UnifiedParser#functionDecl.
	VisitFunctionDecl(ctx *FunctionDeclContext) interface{}

	// Visit a parse tree produced by UnifiedParser#paramList.
	VisitParamList(ctx *ParamListContext) interface{}

	// Visit a parse tree produced by UnifiedParser#parameter.
	VisitParameter(ctx *ParameterContext) interface{}

	// Visit a parse tree produced by UnifiedParser#genericParams.
	VisitGenericParams(ctx *GenericParamsContext) interface{}

	// Visit a parse tree produced by UnifiedParser#genericParam.
	VisitGenericParam(ctx *GenericParamContext) interface{}

	// Visit a parse tree produced by UnifiedParser#typeConstraint.
	VisitTypeConstraint(ctx *TypeConstraintContext) interface{}

	// Visit a parse tree produced by UnifiedParser#whereClause.
	VisitWhereClause(ctx *WhereClauseContext) interface{}

	// Visit a parse tree produced by UnifiedParser#whereConstraint.
	VisitWhereConstraint(ctx *WhereConstraintContext) interface{}

	// Visit a parse tree produced by UnifiedParser#structDecl.
	VisitStructDecl(ctx *StructDeclContext) interface{}

	// Visit a parse tree produced by UnifiedParser#structMember.
	VisitStructMember(ctx *StructMemberContext) interface{}

	// Visit a parse tree produced by UnifiedParser#enumDecl.
	VisitEnumDecl(ctx *EnumDeclContext) interface{}

	// Visit a parse tree produced by UnifiedParser#enumVariant.
	VisitEnumVariant(ctx *EnumVariantContext) interface{}

	// Visit a parse tree produced by UnifiedParser#enumVariantParams.
	VisitEnumVariantParams(ctx *EnumVariantParamsContext) interface{}

	// Visit a parse tree produced by UnifiedParser#enumVariantParam.
	VisitEnumVariantParam(ctx *EnumVariantParamContext) interface{}

	// Visit a parse tree produced by UnifiedParser#interfaceDecl.
	VisitInterfaceDecl(ctx *InterfaceDeclContext) interface{}

	// Visit a parse tree produced by UnifiedParser#interfaceMember.
	VisitInterfaceMember(ctx *InterfaceMemberContext) interface{}

	// Visit a parse tree produced by UnifiedParser#functionSig.
	VisitFunctionSig(ctx *FunctionSigContext) interface{}

	// Visit a parse tree produced by UnifiedParser#implDecl.
	VisitImplDecl(ctx *ImplDeclContext) interface{}

	// Visit a parse tree produced by UnifiedParser#implMember.
	VisitImplMember(ctx *ImplMemberContext) interface{}

	// Visit a parse tree produced by UnifiedParser#actorDecl.
	VisitActorDecl(ctx *ActorDeclContext) interface{}

	// Visit a parse tree produced by UnifiedParser#actorMember.
	VisitActorMember(ctx *ActorMemberContext) interface{}

	// Visit a parse tree produced by UnifiedParser#typeAlias.
	VisitTypeAlias(ctx *TypeAliasContext) interface{}

	// Visit a parse tree produced by UnifiedParser#constantDecl.
	VisitConstantDecl(ctx *ConstantDeclContext) interface{}

	// Visit a parse tree produced by UnifiedParser#type.
	VisitType(ctx *TypeContext) interface{}

	// Visit a parse tree produced by UnifiedParser#typeList.
	VisitTypeList(ctx *TypeListContext) interface{}

	// Visit a parse tree produced by UnifiedParser#statement.
	VisitStatement(ctx *StatementContext) interface{}

	// Visit a parse tree produced by UnifiedParser#letStatement.
	VisitLetStatement(ctx *LetStatementContext) interface{}

	// Visit a parse tree produced by UnifiedParser#varStatement.
	VisitVarStatement(ctx *VarStatementContext) interface{}

	// Visit a parse tree produced by UnifiedParser#assignmentStatement.
	VisitAssignmentStatement(ctx *AssignmentStatementContext) interface{}

	// Visit a parse tree produced by UnifiedParser#assignmentOp.
	VisitAssignmentOp(ctx *AssignmentOpContext) interface{}

	// Visit a parse tree produced by UnifiedParser#regionStatement.
	VisitRegionStatement(ctx *RegionStatementContext) interface{}

	// Visit a parse tree produced by UnifiedParser#exprStatement.
	VisitExprStatement(ctx *ExprStatementContext) interface{}

	// Visit a parse tree produced by UnifiedParser#returnStatement.
	VisitReturnStatement(ctx *ReturnStatementContext) interface{}

	// Visit a parse tree produced by UnifiedParser#ifStatement.
	VisitIfStatement(ctx *IfStatementContext) interface{}

	// Visit a parse tree produced by UnifiedParser#loopStatement.
	VisitLoopStatement(ctx *LoopStatementContext) interface{}

	// Visit a parse tree produced by UnifiedParser#whileStatement.
	VisitWhileStatement(ctx *WhileStatementContext) interface{}

	// Visit a parse tree produced by UnifiedParser#forStatement.
	VisitForStatement(ctx *ForStatementContext) interface{}

	// Visit a parse tree produced by UnifiedParser#switchStatement.
	VisitSwitchStatement(ctx *SwitchStatementContext) interface{}

	// Visit a parse tree produced by UnifiedParser#caseClause.
	VisitCaseClause(ctx *CaseClauseContext) interface{}

	// Visit a parse tree produced by UnifiedParser#breakStatement.
	VisitBreakStatement(ctx *BreakStatementContext) interface{}

	// Visit a parse tree produced by UnifiedParser#continueStatement.
	VisitContinueStatement(ctx *ContinueStatementContext) interface{}

	// Visit a parse tree produced by UnifiedParser#blockStatement.
	VisitBlockStatement(ctx *BlockStatementContext) interface{}

	// Visit a parse tree produced by UnifiedParser#pattern.
	VisitPattern(ctx *PatternContext) interface{}

	// Visit a parse tree produced by UnifiedParser#patternList.
	VisitPatternList(ctx *PatternListContext) interface{}

	// Visit a parse tree produced by UnifiedParser#fieldPattern.
	VisitFieldPattern(ctx *FieldPatternContext) interface{}

	// Visit a parse tree produced by UnifiedParser#fieldPatternList.
	VisitFieldPatternList(ctx *FieldPatternListContext) interface{}

	// Visit a parse tree produced by UnifiedParser#expr.
	VisitExpr(ctx *ExprContext) interface{}

	// Visit a parse tree produced by UnifiedParser#caseExpr.
	VisitCaseExpr(ctx *CaseExprContext) interface{}

	// Visit a parse tree produced by UnifiedParser#primary.
	VisitPrimary(ctx *PrimaryContext) interface{}

	// Visit a parse tree produced by UnifiedParser#lambdaExpr.
	VisitLambdaExpr(ctx *LambdaExprContext) interface{}

	// Visit a parse tree produced by UnifiedParser#asyncExpr.
	VisitAsyncExpr(ctx *AsyncExprContext) interface{}

	// Visit a parse tree produced by UnifiedParser#constructorExpr.
	VisitConstructorExpr(ctx *ConstructorExprContext) interface{}

	// Visit a parse tree produced by UnifiedParser#structExpr.
	VisitStructExpr(ctx *StructExprContext) interface{}

	// Visit a parse tree produced by UnifiedParser#fieldInitList.
	VisitFieldInitList(ctx *FieldInitListContext) interface{}

	// Visit a parse tree produced by UnifiedParser#fieldInit.
	VisitFieldInit(ctx *FieldInitContext) interface{}

	// Visit a parse tree produced by UnifiedParser#listExpr.
	VisitListExpr(ctx *ListExprContext) interface{}

	// Visit a parse tree produced by UnifiedParser#mapExpr.
	VisitMapExpr(ctx *MapExprContext) interface{}

	// Visit a parse tree produced by UnifiedParser#keyValue.
	VisitKeyValue(ctx *KeyValueContext) interface{}

	// Visit a parse tree produced by UnifiedParser#setExpr.
	VisitSetExpr(ctx *SetExprContext) interface{}

	// Visit a parse tree produced by UnifiedParser#tupleExpr.
	VisitTupleExpr(ctx *TupleExprContext) interface{}

	// Visit a parse tree produced by UnifiedParser#namedTupleField.
	VisitNamedTupleField(ctx *NamedTupleFieldContext) interface{}

	// Visit a parse tree produced by UnifiedParser#block.
	VisitBlock(ctx *BlockContext) interface{}

	// Visit a parse tree produced by UnifiedParser#argList.
	VisitArgList(ctx *ArgListContext) interface{}

	// Visit a parse tree produced by UnifiedParser#literal.
	VisitLiteral(ctx *LiteralContext) interface{}

	// Visit a parse tree produced by UnifiedParser#identifier.
	VisitIdentifier(ctx *IdentifierContext) interface{}
}
