// Code generated from UnifiedParser.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // UnifiedParser
import "github.com/antlr4-go/antlr/v4"

// BaseUnifiedParserListener is a complete listener for a parse tree produced by UnifiedParser.
type BaseUnifiedParserListener struct{}

var _ UnifiedParserListener = &BaseUnifiedParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseUnifiedParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseUnifiedParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseUnifiedParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseUnifiedParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProgram is called when production program is entered.
func (s *BaseUnifiedParserListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BaseUnifiedParserListener) ExitProgram(ctx *ProgramContext) {}

// EnterItem is called when production item is entered.
func (s *BaseUnifiedParserListener) EnterItem(ctx *ItemContext) {}

// ExitItem is called when production item is exited.
func (s *BaseUnifiedParserListener) ExitItem(ctx *ItemContext) {}

// EnterModuleDecl is called when production moduleDecl is entered.
func (s *BaseUnifiedParserListener) EnterModuleDecl(ctx *ModuleDeclContext) {}

// ExitModuleDecl is called when production moduleDecl is exited.
func (s *BaseUnifiedParserListener) ExitModuleDecl(ctx *ModuleDeclContext) {}

// EnterImportDecl is called when production importDecl is entered.
func (s *BaseUnifiedParserListener) EnterImportDecl(ctx *ImportDeclContext) {}

// ExitImportDecl is called when production importDecl is exited.
func (s *BaseUnifiedParserListener) ExitImportDecl(ctx *ImportDeclContext) {}

// EnterImportPath is called when production importPath is entered.
func (s *BaseUnifiedParserListener) EnterImportPath(ctx *ImportPathContext) {}

// ExitImportPath is called when production importPath is exited.
func (s *BaseUnifiedParserListener) ExitImportPath(ctx *ImportPathContext) {}

// EnterImportList is called when production importList is entered.
func (s *BaseUnifiedParserListener) EnterImportList(ctx *ImportListContext) {}

// ExitImportList is called when production importList is exited.
func (s *BaseUnifiedParserListener) ExitImportList(ctx *ImportListContext) {}

// EnterFunctionDecl is called when production functionDecl is entered.
func (s *BaseUnifiedParserListener) EnterFunctionDecl(ctx *FunctionDeclContext) {}

// ExitFunctionDecl is called when production functionDecl is exited.
func (s *BaseUnifiedParserListener) ExitFunctionDecl(ctx *FunctionDeclContext) {}

// EnterParamList is called when production paramList is entered.
func (s *BaseUnifiedParserListener) EnterParamList(ctx *ParamListContext) {}

// ExitParamList is called when production paramList is exited.
func (s *BaseUnifiedParserListener) ExitParamList(ctx *ParamListContext) {}

// EnterParameter is called when production parameter is entered.
func (s *BaseUnifiedParserListener) EnterParameter(ctx *ParameterContext) {}

// ExitParameter is called when production parameter is exited.
func (s *BaseUnifiedParserListener) ExitParameter(ctx *ParameterContext) {}

// EnterGenericParams is called when production genericParams is entered.
func (s *BaseUnifiedParserListener) EnterGenericParams(ctx *GenericParamsContext) {}

// ExitGenericParams is called when production genericParams is exited.
func (s *BaseUnifiedParserListener) ExitGenericParams(ctx *GenericParamsContext) {}

// EnterGenericParam is called when production genericParam is entered.
func (s *BaseUnifiedParserListener) EnterGenericParam(ctx *GenericParamContext) {}

// ExitGenericParam is called when production genericParam is exited.
func (s *BaseUnifiedParserListener) ExitGenericParam(ctx *GenericParamContext) {}

// EnterTypeConstraint is called when production typeConstraint is entered.
func (s *BaseUnifiedParserListener) EnterTypeConstraint(ctx *TypeConstraintContext) {}

// ExitTypeConstraint is called when production typeConstraint is exited.
func (s *BaseUnifiedParserListener) ExitTypeConstraint(ctx *TypeConstraintContext) {}

// EnterWhereClause is called when production whereClause is entered.
func (s *BaseUnifiedParserListener) EnterWhereClause(ctx *WhereClauseContext) {}

// ExitWhereClause is called when production whereClause is exited.
func (s *BaseUnifiedParserListener) ExitWhereClause(ctx *WhereClauseContext) {}

// EnterWhereConstraint is called when production whereConstraint is entered.
func (s *BaseUnifiedParserListener) EnterWhereConstraint(ctx *WhereConstraintContext) {}

// ExitWhereConstraint is called when production whereConstraint is exited.
func (s *BaseUnifiedParserListener) ExitWhereConstraint(ctx *WhereConstraintContext) {}

// EnterStructDecl is called when production structDecl is entered.
func (s *BaseUnifiedParserListener) EnterStructDecl(ctx *StructDeclContext) {}

// ExitStructDecl is called when production structDecl is exited.
func (s *BaseUnifiedParserListener) ExitStructDecl(ctx *StructDeclContext) {}

// EnterStructMember is called when production structMember is entered.
func (s *BaseUnifiedParserListener) EnterStructMember(ctx *StructMemberContext) {}

// ExitStructMember is called when production structMember is exited.
func (s *BaseUnifiedParserListener) ExitStructMember(ctx *StructMemberContext) {}

// EnterEnumDecl is called when production enumDecl is entered.
func (s *BaseUnifiedParserListener) EnterEnumDecl(ctx *EnumDeclContext) {}

// ExitEnumDecl is called when production enumDecl is exited.
func (s *BaseUnifiedParserListener) ExitEnumDecl(ctx *EnumDeclContext) {}

// EnterEnumVariant is called when production enumVariant is entered.
func (s *BaseUnifiedParserListener) EnterEnumVariant(ctx *EnumVariantContext) {}

// ExitEnumVariant is called when production enumVariant is exited.
func (s *BaseUnifiedParserListener) ExitEnumVariant(ctx *EnumVariantContext) {}

// EnterEnumVariantParams is called when production enumVariantParams is entered.
func (s *BaseUnifiedParserListener) EnterEnumVariantParams(ctx *EnumVariantParamsContext) {}

// ExitEnumVariantParams is called when production enumVariantParams is exited.
func (s *BaseUnifiedParserListener) ExitEnumVariantParams(ctx *EnumVariantParamsContext) {}

// EnterEnumVariantParam is called when production enumVariantParam is entered.
func (s *BaseUnifiedParserListener) EnterEnumVariantParam(ctx *EnumVariantParamContext) {}

// ExitEnumVariantParam is called when production enumVariantParam is exited.
func (s *BaseUnifiedParserListener) ExitEnumVariantParam(ctx *EnumVariantParamContext) {}

// EnterInterfaceDecl is called when production interfaceDecl is entered.
func (s *BaseUnifiedParserListener) EnterInterfaceDecl(ctx *InterfaceDeclContext) {}

// ExitInterfaceDecl is called when production interfaceDecl is exited.
func (s *BaseUnifiedParserListener) ExitInterfaceDecl(ctx *InterfaceDeclContext) {}

// EnterInterfaceMember is called when production interfaceMember is entered.
func (s *BaseUnifiedParserListener) EnterInterfaceMember(ctx *InterfaceMemberContext) {}

// ExitInterfaceMember is called when production interfaceMember is exited.
func (s *BaseUnifiedParserListener) ExitInterfaceMember(ctx *InterfaceMemberContext) {}

// EnterFunctionSig is called when production functionSig is entered.
func (s *BaseUnifiedParserListener) EnterFunctionSig(ctx *FunctionSigContext) {}

// ExitFunctionSig is called when production functionSig is exited.
func (s *BaseUnifiedParserListener) ExitFunctionSig(ctx *FunctionSigContext) {}

// EnterImplDecl is called when production implDecl is entered.
func (s *BaseUnifiedParserListener) EnterImplDecl(ctx *ImplDeclContext) {}

// ExitImplDecl is called when production implDecl is exited.
func (s *BaseUnifiedParserListener) ExitImplDecl(ctx *ImplDeclContext) {}

// EnterImplMember is called when production implMember is entered.
func (s *BaseUnifiedParserListener) EnterImplMember(ctx *ImplMemberContext) {}

// ExitImplMember is called when production implMember is exited.
func (s *BaseUnifiedParserListener) ExitImplMember(ctx *ImplMemberContext) {}

// EnterActorDecl is called when production actorDecl is entered.
func (s *BaseUnifiedParserListener) EnterActorDecl(ctx *ActorDeclContext) {}

// ExitActorDecl is called when production actorDecl is exited.
func (s *BaseUnifiedParserListener) ExitActorDecl(ctx *ActorDeclContext) {}

// EnterActorMember is called when production actorMember is entered.
func (s *BaseUnifiedParserListener) EnterActorMember(ctx *ActorMemberContext) {}

// ExitActorMember is called when production actorMember is exited.
func (s *BaseUnifiedParserListener) ExitActorMember(ctx *ActorMemberContext) {}

// EnterTypeAlias is called when production typeAlias is entered.
func (s *BaseUnifiedParserListener) EnterTypeAlias(ctx *TypeAliasContext) {}

// ExitTypeAlias is called when production typeAlias is exited.
func (s *BaseUnifiedParserListener) ExitTypeAlias(ctx *TypeAliasContext) {}

// EnterConstantDecl is called when production constantDecl is entered.
func (s *BaseUnifiedParserListener) EnterConstantDecl(ctx *ConstantDeclContext) {}

// ExitConstantDecl is called when production constantDecl is exited.
func (s *BaseUnifiedParserListener) ExitConstantDecl(ctx *ConstantDeclContext) {}

// EnterType is called when production type is entered.
func (s *BaseUnifiedParserListener) EnterType(ctx *TypeContext) {}

// ExitType is called when production type is exited.
func (s *BaseUnifiedParserListener) ExitType(ctx *TypeContext) {}

// EnterTypeList is called when production typeList is entered.
func (s *BaseUnifiedParserListener) EnterTypeList(ctx *TypeListContext) {}

// ExitTypeList is called when production typeList is exited.
func (s *BaseUnifiedParserListener) ExitTypeList(ctx *TypeListContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaseUnifiedParserListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseUnifiedParserListener) ExitStatement(ctx *StatementContext) {}

// EnterLetStatement is called when production letStatement is entered.
func (s *BaseUnifiedParserListener) EnterLetStatement(ctx *LetStatementContext) {}

// ExitLetStatement is called when production letStatement is exited.
func (s *BaseUnifiedParserListener) ExitLetStatement(ctx *LetStatementContext) {}

// EnterVarStatement is called when production varStatement is entered.
func (s *BaseUnifiedParserListener) EnterVarStatement(ctx *VarStatementContext) {}

// ExitVarStatement is called when production varStatement is exited.
func (s *BaseUnifiedParserListener) ExitVarStatement(ctx *VarStatementContext) {}

// EnterAssignmentStatement is called when production assignmentStatement is entered.
func (s *BaseUnifiedParserListener) EnterAssignmentStatement(ctx *AssignmentStatementContext) {}

// ExitAssignmentStatement is called when production assignmentStatement is exited.
func (s *BaseUnifiedParserListener) ExitAssignmentStatement(ctx *AssignmentStatementContext) {}

// EnterAssignmentOp is called when production assignmentOp is entered.
func (s *BaseUnifiedParserListener) EnterAssignmentOp(ctx *AssignmentOpContext) {}

// ExitAssignmentOp is called when production assignmentOp is exited.
func (s *BaseUnifiedParserListener) ExitAssignmentOp(ctx *AssignmentOpContext) {}

// EnterRegionStatement is called when production regionStatement is entered.
func (s *BaseUnifiedParserListener) EnterRegionStatement(ctx *RegionStatementContext) {}

// ExitRegionStatement is called when production regionStatement is exited.
func (s *BaseUnifiedParserListener) ExitRegionStatement(ctx *RegionStatementContext) {}

// EnterExprStatement is called when production exprStatement is entered.
func (s *BaseUnifiedParserListener) EnterExprStatement(ctx *ExprStatementContext) {}

// ExitExprStatement is called when production exprStatement is exited.
func (s *BaseUnifiedParserListener) ExitExprStatement(ctx *ExprStatementContext) {}

// EnterReturnStatement is called when production returnStatement is entered.
func (s *BaseUnifiedParserListener) EnterReturnStatement(ctx *ReturnStatementContext) {}

// ExitReturnStatement is called when production returnStatement is exited.
func (s *BaseUnifiedParserListener) ExitReturnStatement(ctx *ReturnStatementContext) {}

// EnterIfStatement is called when production ifStatement is entered.
func (s *BaseUnifiedParserListener) EnterIfStatement(ctx *IfStatementContext) {}

// ExitIfStatement is called when production ifStatement is exited.
func (s *BaseUnifiedParserListener) ExitIfStatement(ctx *IfStatementContext) {}

// EnterLoopStatement is called when production loopStatement is entered.
func (s *BaseUnifiedParserListener) EnterLoopStatement(ctx *LoopStatementContext) {}

// ExitLoopStatement is called when production loopStatement is exited.
func (s *BaseUnifiedParserListener) ExitLoopStatement(ctx *LoopStatementContext) {}

// EnterWhileStatement is called when production whileStatement is entered.
func (s *BaseUnifiedParserListener) EnterWhileStatement(ctx *WhileStatementContext) {}

// ExitWhileStatement is called when production whileStatement is exited.
func (s *BaseUnifiedParserListener) ExitWhileStatement(ctx *WhileStatementContext) {}

// EnterForStatement is called when production forStatement is entered.
func (s *BaseUnifiedParserListener) EnterForStatement(ctx *ForStatementContext) {}

// ExitForStatement is called when production forStatement is exited.
func (s *BaseUnifiedParserListener) ExitForStatement(ctx *ForStatementContext) {}

// EnterSwitchStatement is called when production switchStatement is entered.
func (s *BaseUnifiedParserListener) EnterSwitchStatement(ctx *SwitchStatementContext) {}

// ExitSwitchStatement is called when production switchStatement is exited.
func (s *BaseUnifiedParserListener) ExitSwitchStatement(ctx *SwitchStatementContext) {}

// EnterCaseClause is called when production caseClause is entered.
func (s *BaseUnifiedParserListener) EnterCaseClause(ctx *CaseClauseContext) {}

// ExitCaseClause is called when production caseClause is exited.
func (s *BaseUnifiedParserListener) ExitCaseClause(ctx *CaseClauseContext) {}

// EnterBreakStatement is called when production breakStatement is entered.
func (s *BaseUnifiedParserListener) EnterBreakStatement(ctx *BreakStatementContext) {}

// ExitBreakStatement is called when production breakStatement is exited.
func (s *BaseUnifiedParserListener) ExitBreakStatement(ctx *BreakStatementContext) {}

// EnterContinueStatement is called when production continueStatement is entered.
func (s *BaseUnifiedParserListener) EnterContinueStatement(ctx *ContinueStatementContext) {}

// ExitContinueStatement is called when production continueStatement is exited.
func (s *BaseUnifiedParserListener) ExitContinueStatement(ctx *ContinueStatementContext) {}

// EnterBlockStatement is called when production blockStatement is entered.
func (s *BaseUnifiedParserListener) EnterBlockStatement(ctx *BlockStatementContext) {}

// ExitBlockStatement is called when production blockStatement is exited.
func (s *BaseUnifiedParserListener) ExitBlockStatement(ctx *BlockStatementContext) {}

// EnterPattern is called when production pattern is entered.
func (s *BaseUnifiedParserListener) EnterPattern(ctx *PatternContext) {}

// ExitPattern is called when production pattern is exited.
func (s *BaseUnifiedParserListener) ExitPattern(ctx *PatternContext) {}

// EnterPatternList is called when production patternList is entered.
func (s *BaseUnifiedParserListener) EnterPatternList(ctx *PatternListContext) {}

// ExitPatternList is called when production patternList is exited.
func (s *BaseUnifiedParserListener) ExitPatternList(ctx *PatternListContext) {}

// EnterFieldPattern is called when production fieldPattern is entered.
func (s *BaseUnifiedParserListener) EnterFieldPattern(ctx *FieldPatternContext) {}

// ExitFieldPattern is called when production fieldPattern is exited.
func (s *BaseUnifiedParserListener) ExitFieldPattern(ctx *FieldPatternContext) {}

// EnterFieldPatternList is called when production fieldPatternList is entered.
func (s *BaseUnifiedParserListener) EnterFieldPatternList(ctx *FieldPatternListContext) {}

// ExitFieldPatternList is called when production fieldPatternList is exited.
func (s *BaseUnifiedParserListener) ExitFieldPatternList(ctx *FieldPatternListContext) {}

// EnterExpr is called when production expr is entered.
func (s *BaseUnifiedParserListener) EnterExpr(ctx *ExprContext) {}

// ExitExpr is called when production expr is exited.
func (s *BaseUnifiedParserListener) ExitExpr(ctx *ExprContext) {}

// EnterCaseExpr is called when production caseExpr is entered.
func (s *BaseUnifiedParserListener) EnterCaseExpr(ctx *CaseExprContext) {}

// ExitCaseExpr is called when production caseExpr is exited.
func (s *BaseUnifiedParserListener) ExitCaseExpr(ctx *CaseExprContext) {}

// EnterPrimary is called when production primary is entered.
func (s *BaseUnifiedParserListener) EnterPrimary(ctx *PrimaryContext) {}

// ExitPrimary is called when production primary is exited.
func (s *BaseUnifiedParserListener) ExitPrimary(ctx *PrimaryContext) {}

// EnterLambdaExpr is called when production lambdaExpr is entered.
func (s *BaseUnifiedParserListener) EnterLambdaExpr(ctx *LambdaExprContext) {}

// ExitLambdaExpr is called when production lambdaExpr is exited.
func (s *BaseUnifiedParserListener) ExitLambdaExpr(ctx *LambdaExprContext) {}

// EnterAsyncExpr is called when production asyncExpr is entered.
func (s *BaseUnifiedParserListener) EnterAsyncExpr(ctx *AsyncExprContext) {}

// ExitAsyncExpr is called when production asyncExpr is exited.
func (s *BaseUnifiedParserListener) ExitAsyncExpr(ctx *AsyncExprContext) {}

// EnterConstructorExpr is called when production constructorExpr is entered.
func (s *BaseUnifiedParserListener) EnterConstructorExpr(ctx *ConstructorExprContext) {}

// ExitConstructorExpr is called when production constructorExpr is exited.
func (s *BaseUnifiedParserListener) ExitConstructorExpr(ctx *ConstructorExprContext) {}

// EnterStructExpr is called when production structExpr is entered.
func (s *BaseUnifiedParserListener) EnterStructExpr(ctx *StructExprContext) {}

// ExitStructExpr is called when production structExpr is exited.
func (s *BaseUnifiedParserListener) ExitStructExpr(ctx *StructExprContext) {}

// EnterFieldInitList is called when production fieldInitList is entered.
func (s *BaseUnifiedParserListener) EnterFieldInitList(ctx *FieldInitListContext) {}

// ExitFieldInitList is called when production fieldInitList is exited.
func (s *BaseUnifiedParserListener) ExitFieldInitList(ctx *FieldInitListContext) {}

// EnterFieldInit is called when production fieldInit is entered.
func (s *BaseUnifiedParserListener) EnterFieldInit(ctx *FieldInitContext) {}

// ExitFieldInit is called when production fieldInit is exited.
func (s *BaseUnifiedParserListener) ExitFieldInit(ctx *FieldInitContext) {}

// EnterListExpr is called when production listExpr is entered.
func (s *BaseUnifiedParserListener) EnterListExpr(ctx *ListExprContext) {}

// ExitListExpr is called when production listExpr is exited.
func (s *BaseUnifiedParserListener) ExitListExpr(ctx *ListExprContext) {}

// EnterMapExpr is called when production mapExpr is entered.
func (s *BaseUnifiedParserListener) EnterMapExpr(ctx *MapExprContext) {}

// ExitMapExpr is called when production mapExpr is exited.
func (s *BaseUnifiedParserListener) ExitMapExpr(ctx *MapExprContext) {}

// EnterKeyValue is called when production keyValue is entered.
func (s *BaseUnifiedParserListener) EnterKeyValue(ctx *KeyValueContext) {}

// ExitKeyValue is called when production keyValue is exited.
func (s *BaseUnifiedParserListener) ExitKeyValue(ctx *KeyValueContext) {}

// EnterSetExpr is called when production setExpr is entered.
func (s *BaseUnifiedParserListener) EnterSetExpr(ctx *SetExprContext) {}

// ExitSetExpr is called when production setExpr is exited.
func (s *BaseUnifiedParserListener) ExitSetExpr(ctx *SetExprContext) {}

// EnterTupleExpr is called when production tupleExpr is entered.
func (s *BaseUnifiedParserListener) EnterTupleExpr(ctx *TupleExprContext) {}

// ExitTupleExpr is called when production tupleExpr is exited.
func (s *BaseUnifiedParserListener) ExitTupleExpr(ctx *TupleExprContext) {}

// EnterNamedTupleField is called when production namedTupleField is entered.
func (s *BaseUnifiedParserListener) EnterNamedTupleField(ctx *NamedTupleFieldContext) {}

// ExitNamedTupleField is called when production namedTupleField is exited.
func (s *BaseUnifiedParserListener) ExitNamedTupleField(ctx *NamedTupleFieldContext) {}

// EnterBlock is called when production block is entered.
func (s *BaseUnifiedParserListener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *BaseUnifiedParserListener) ExitBlock(ctx *BlockContext) {}

// EnterArgList is called when production argList is entered.
func (s *BaseUnifiedParserListener) EnterArgList(ctx *ArgListContext) {}

// ExitArgList is called when production argList is exited.
func (s *BaseUnifiedParserListener) ExitArgList(ctx *ArgListContext) {}

// EnterLiteral is called when production literal is entered.
func (s *BaseUnifiedParserListener) EnterLiteral(ctx *LiteralContext) {}

// ExitLiteral is called when production literal is exited.
func (s *BaseUnifiedParserListener) ExitLiteral(ctx *LiteralContext) {}

// EnterIdentifier is called when production identifier is entered.
func (s *BaseUnifiedParserListener) EnterIdentifier(ctx *IdentifierContext) {}

// ExitIdentifier is called when production identifier is exited.
func (s *BaseUnifiedParserListener) ExitIdentifier(ctx *IdentifierContext) {}
