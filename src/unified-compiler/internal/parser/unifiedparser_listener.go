// Code generated from UnifiedParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // UnifiedParser
import "github.com/antlr4-go/antlr/v4"

// UnifiedParserListener is a complete listener for a parse tree produced by UnifiedParser.
type UnifiedParserListener interface {
	antlr.ParseTreeListener

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterItem is called when entering the item production.
	EnterItem(c *ItemContext)

	// EnterModuleDecl is called when entering the moduleDecl production.
	EnterModuleDecl(c *ModuleDeclContext)

	// EnterImportDecl is called when entering the importDecl production.
	EnterImportDecl(c *ImportDeclContext)

	// EnterImportPath is called when entering the importPath production.
	EnterImportPath(c *ImportPathContext)

	// EnterImportList is called when entering the importList production.
	EnterImportList(c *ImportListContext)

	// EnterFunctionDecl is called when entering the functionDecl production.
	EnterFunctionDecl(c *FunctionDeclContext)

	// EnterParamList is called when entering the paramList production.
	EnterParamList(c *ParamListContext)

	// EnterParameter is called when entering the parameter production.
	EnterParameter(c *ParameterContext)

	// EnterGenericParams is called when entering the genericParams production.
	EnterGenericParams(c *GenericParamsContext)

	// EnterGenericParam is called when entering the genericParam production.
	EnterGenericParam(c *GenericParamContext)

	// EnterTypeConstraint is called when entering the typeConstraint production.
	EnterTypeConstraint(c *TypeConstraintContext)

	// EnterWhereClause is called when entering the whereClause production.
	EnterWhereClause(c *WhereClauseContext)

	// EnterWhereConstraint is called when entering the whereConstraint production.
	EnterWhereConstraint(c *WhereConstraintContext)

	// EnterStructDecl is called when entering the structDecl production.
	EnterStructDecl(c *StructDeclContext)

	// EnterStructMember is called when entering the structMember production.
	EnterStructMember(c *StructMemberContext)

	// EnterEnumDecl is called when entering the enumDecl production.
	EnterEnumDecl(c *EnumDeclContext)

	// EnterEnumVariant is called when entering the enumVariant production.
	EnterEnumVariant(c *EnumVariantContext)

	// EnterEnumVariantParams is called when entering the enumVariantParams production.
	EnterEnumVariantParams(c *EnumVariantParamsContext)

	// EnterEnumVariantParam is called when entering the enumVariantParam production.
	EnterEnumVariantParam(c *EnumVariantParamContext)

	// EnterInterfaceDecl is called when entering the interfaceDecl production.
	EnterInterfaceDecl(c *InterfaceDeclContext)

	// EnterInterfaceMember is called when entering the interfaceMember production.
	EnterInterfaceMember(c *InterfaceMemberContext)

	// EnterFunctionSig is called when entering the functionSig production.
	EnterFunctionSig(c *FunctionSigContext)

	// EnterImplDecl is called when entering the implDecl production.
	EnterImplDecl(c *ImplDeclContext)

	// EnterImplMember is called when entering the implMember production.
	EnterImplMember(c *ImplMemberContext)

	// EnterActorDecl is called when entering the actorDecl production.
	EnterActorDecl(c *ActorDeclContext)

	// EnterActorMember is called when entering the actorMember production.
	EnterActorMember(c *ActorMemberContext)

	// EnterTypeAlias is called when entering the typeAlias production.
	EnterTypeAlias(c *TypeAliasContext)

	// EnterConstantDecl is called when entering the constantDecl production.
	EnterConstantDecl(c *ConstantDeclContext)

	// EnterType is called when entering the type production.
	EnterType(c *TypeContext)

	// EnterTypeList is called when entering the typeList production.
	EnterTypeList(c *TypeListContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterLetStatement is called when entering the letStatement production.
	EnterLetStatement(c *LetStatementContext)

	// EnterVarStatement is called when entering the varStatement production.
	EnterVarStatement(c *VarStatementContext)

	// EnterAssignmentStatement is called when entering the assignmentStatement production.
	EnterAssignmentStatement(c *AssignmentStatementContext)

	// EnterAssignmentOp is called when entering the assignmentOp production.
	EnterAssignmentOp(c *AssignmentOpContext)

	// EnterRegionStatement is called when entering the regionStatement production.
	EnterRegionStatement(c *RegionStatementContext)

	// EnterExprStatement is called when entering the exprStatement production.
	EnterExprStatement(c *ExprStatementContext)

	// EnterReturnStatement is called when entering the returnStatement production.
	EnterReturnStatement(c *ReturnStatementContext)

	// EnterIfStatement is called when entering the ifStatement production.
	EnterIfStatement(c *IfStatementContext)

	// EnterLoopStatement is called when entering the loopStatement production.
	EnterLoopStatement(c *LoopStatementContext)

	// EnterWhileStatement is called when entering the whileStatement production.
	EnterWhileStatement(c *WhileStatementContext)

	// EnterForStatement is called when entering the forStatement production.
	EnterForStatement(c *ForStatementContext)

	// EnterSwitchStatement is called when entering the switchStatement production.
	EnterSwitchStatement(c *SwitchStatementContext)

	// EnterCaseClause is called when entering the caseClause production.
	EnterCaseClause(c *CaseClauseContext)

	// EnterBreakStatement is called when entering the breakStatement production.
	EnterBreakStatement(c *BreakStatementContext)

	// EnterContinueStatement is called when entering the continueStatement production.
	EnterContinueStatement(c *ContinueStatementContext)

	// EnterBlockStatement is called when entering the blockStatement production.
	EnterBlockStatement(c *BlockStatementContext)

	// EnterPattern is called when entering the pattern production.
	EnterPattern(c *PatternContext)

	// EnterPatternList is called when entering the patternList production.
	EnterPatternList(c *PatternListContext)

	// EnterFieldPattern is called when entering the fieldPattern production.
	EnterFieldPattern(c *FieldPatternContext)

	// EnterFieldPatternList is called when entering the fieldPatternList production.
	EnterFieldPatternList(c *FieldPatternListContext)

	// EnterExpr is called when entering the expr production.
	EnterExpr(c *ExprContext)

	// EnterCaseExpr is called when entering the caseExpr production.
	EnterCaseExpr(c *CaseExprContext)

	// EnterPrimary is called when entering the primary production.
	EnterPrimary(c *PrimaryContext)

	// EnterLambdaExpr is called when entering the lambdaExpr production.
	EnterLambdaExpr(c *LambdaExprContext)

	// EnterAsyncExpr is called when entering the asyncExpr production.
	EnterAsyncExpr(c *AsyncExprContext)

	// EnterConstructorExpr is called when entering the constructorExpr production.
	EnterConstructorExpr(c *ConstructorExprContext)

	// EnterEnumConstructorExpr is called when entering the enumConstructorExpr production.
	EnterEnumConstructorExpr(c *EnumConstructorExprContext)

	// EnterStructExpr is called when entering the structExpr production.
	EnterStructExpr(c *StructExprContext)

	// EnterFieldInitList is called when entering the fieldInitList production.
	EnterFieldInitList(c *FieldInitListContext)

	// EnterFieldInit is called when entering the fieldInit production.
	EnterFieldInit(c *FieldInitContext)

	// EnterListExpr is called when entering the listExpr production.
	EnterListExpr(c *ListExprContext)

	// EnterMapExpr is called when entering the mapExpr production.
	EnterMapExpr(c *MapExprContext)

	// EnterKeyValue is called when entering the keyValue production.
	EnterKeyValue(c *KeyValueContext)

	// EnterSetExpr is called when entering the setExpr production.
	EnterSetExpr(c *SetExprContext)

	// EnterTupleExpr is called when entering the tupleExpr production.
	EnterTupleExpr(c *TupleExprContext)

	// EnterNamedTupleField is called when entering the namedTupleField production.
	EnterNamedTupleField(c *NamedTupleFieldContext)

	// EnterBlock is called when entering the block production.
	EnterBlock(c *BlockContext)

	// EnterArgList is called when entering the argList production.
	EnterArgList(c *ArgListContext)

	// EnterLiteral is called when entering the literal production.
	EnterLiteral(c *LiteralContext)

	// EnterIdentifier is called when entering the identifier production.
	EnterIdentifier(c *IdentifierContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitItem is called when exiting the item production.
	ExitItem(c *ItemContext)

	// ExitModuleDecl is called when exiting the moduleDecl production.
	ExitModuleDecl(c *ModuleDeclContext)

	// ExitImportDecl is called when exiting the importDecl production.
	ExitImportDecl(c *ImportDeclContext)

	// ExitImportPath is called when exiting the importPath production.
	ExitImportPath(c *ImportPathContext)

	// ExitImportList is called when exiting the importList production.
	ExitImportList(c *ImportListContext)

	// ExitFunctionDecl is called when exiting the functionDecl production.
	ExitFunctionDecl(c *FunctionDeclContext)

	// ExitParamList is called when exiting the paramList production.
	ExitParamList(c *ParamListContext)

	// ExitParameter is called when exiting the parameter production.
	ExitParameter(c *ParameterContext)

	// ExitGenericParams is called when exiting the genericParams production.
	ExitGenericParams(c *GenericParamsContext)

	// ExitGenericParam is called when exiting the genericParam production.
	ExitGenericParam(c *GenericParamContext)

	// ExitTypeConstraint is called when exiting the typeConstraint production.
	ExitTypeConstraint(c *TypeConstraintContext)

	// ExitWhereClause is called when exiting the whereClause production.
	ExitWhereClause(c *WhereClauseContext)

	// ExitWhereConstraint is called when exiting the whereConstraint production.
	ExitWhereConstraint(c *WhereConstraintContext)

	// ExitStructDecl is called when exiting the structDecl production.
	ExitStructDecl(c *StructDeclContext)

	// ExitStructMember is called when exiting the structMember production.
	ExitStructMember(c *StructMemberContext)

	// ExitEnumDecl is called when exiting the enumDecl production.
	ExitEnumDecl(c *EnumDeclContext)

	// ExitEnumVariant is called when exiting the enumVariant production.
	ExitEnumVariant(c *EnumVariantContext)

	// ExitEnumVariantParams is called when exiting the enumVariantParams production.
	ExitEnumVariantParams(c *EnumVariantParamsContext)

	// ExitEnumVariantParam is called when exiting the enumVariantParam production.
	ExitEnumVariantParam(c *EnumVariantParamContext)

	// ExitInterfaceDecl is called when exiting the interfaceDecl production.
	ExitInterfaceDecl(c *InterfaceDeclContext)

	// ExitInterfaceMember is called when exiting the interfaceMember production.
	ExitInterfaceMember(c *InterfaceMemberContext)

	// ExitFunctionSig is called when exiting the functionSig production.
	ExitFunctionSig(c *FunctionSigContext)

	// ExitImplDecl is called when exiting the implDecl production.
	ExitImplDecl(c *ImplDeclContext)

	// ExitImplMember is called when exiting the implMember production.
	ExitImplMember(c *ImplMemberContext)

	// ExitActorDecl is called when exiting the actorDecl production.
	ExitActorDecl(c *ActorDeclContext)

	// ExitActorMember is called when exiting the actorMember production.
	ExitActorMember(c *ActorMemberContext)

	// ExitTypeAlias is called when exiting the typeAlias production.
	ExitTypeAlias(c *TypeAliasContext)

	// ExitConstantDecl is called when exiting the constantDecl production.
	ExitConstantDecl(c *ConstantDeclContext)

	// ExitType is called when exiting the type production.
	ExitType(c *TypeContext)

	// ExitTypeList is called when exiting the typeList production.
	ExitTypeList(c *TypeListContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitLetStatement is called when exiting the letStatement production.
	ExitLetStatement(c *LetStatementContext)

	// ExitVarStatement is called when exiting the varStatement production.
	ExitVarStatement(c *VarStatementContext)

	// ExitAssignmentStatement is called when exiting the assignmentStatement production.
	ExitAssignmentStatement(c *AssignmentStatementContext)

	// ExitAssignmentOp is called when exiting the assignmentOp production.
	ExitAssignmentOp(c *AssignmentOpContext)

	// ExitRegionStatement is called when exiting the regionStatement production.
	ExitRegionStatement(c *RegionStatementContext)

	// ExitExprStatement is called when exiting the exprStatement production.
	ExitExprStatement(c *ExprStatementContext)

	// ExitReturnStatement is called when exiting the returnStatement production.
	ExitReturnStatement(c *ReturnStatementContext)

	// ExitIfStatement is called when exiting the ifStatement production.
	ExitIfStatement(c *IfStatementContext)

	// ExitLoopStatement is called when exiting the loopStatement production.
	ExitLoopStatement(c *LoopStatementContext)

	// ExitWhileStatement is called when exiting the whileStatement production.
	ExitWhileStatement(c *WhileStatementContext)

	// ExitForStatement is called when exiting the forStatement production.
	ExitForStatement(c *ForStatementContext)

	// ExitSwitchStatement is called when exiting the switchStatement production.
	ExitSwitchStatement(c *SwitchStatementContext)

	// ExitCaseClause is called when exiting the caseClause production.
	ExitCaseClause(c *CaseClauseContext)

	// ExitBreakStatement is called when exiting the breakStatement production.
	ExitBreakStatement(c *BreakStatementContext)

	// ExitContinueStatement is called when exiting the continueStatement production.
	ExitContinueStatement(c *ContinueStatementContext)

	// ExitBlockStatement is called when exiting the blockStatement production.
	ExitBlockStatement(c *BlockStatementContext)

	// ExitPattern is called when exiting the pattern production.
	ExitPattern(c *PatternContext)

	// ExitPatternList is called when exiting the patternList production.
	ExitPatternList(c *PatternListContext)

	// ExitFieldPattern is called when exiting the fieldPattern production.
	ExitFieldPattern(c *FieldPatternContext)

	// ExitFieldPatternList is called when exiting the fieldPatternList production.
	ExitFieldPatternList(c *FieldPatternListContext)

	// ExitExpr is called when exiting the expr production.
	ExitExpr(c *ExprContext)

	// ExitCaseExpr is called when exiting the caseExpr production.
	ExitCaseExpr(c *CaseExprContext)

	// ExitPrimary is called when exiting the primary production.
	ExitPrimary(c *PrimaryContext)

	// ExitLambdaExpr is called when exiting the lambdaExpr production.
	ExitLambdaExpr(c *LambdaExprContext)

	// ExitAsyncExpr is called when exiting the asyncExpr production.
	ExitAsyncExpr(c *AsyncExprContext)

	// ExitConstructorExpr is called when exiting the constructorExpr production.
	ExitConstructorExpr(c *ConstructorExprContext)

	// ExitEnumConstructorExpr is called when exiting the enumConstructorExpr production.
	ExitEnumConstructorExpr(c *EnumConstructorExprContext)

	// ExitStructExpr is called when exiting the structExpr production.
	ExitStructExpr(c *StructExprContext)

	// ExitFieldInitList is called when exiting the fieldInitList production.
	ExitFieldInitList(c *FieldInitListContext)

	// ExitFieldInit is called when exiting the fieldInit production.
	ExitFieldInit(c *FieldInitContext)

	// ExitListExpr is called when exiting the listExpr production.
	ExitListExpr(c *ListExprContext)

	// ExitMapExpr is called when exiting the mapExpr production.
	ExitMapExpr(c *MapExprContext)

	// ExitKeyValue is called when exiting the keyValue production.
	ExitKeyValue(c *KeyValueContext)

	// ExitSetExpr is called when exiting the setExpr production.
	ExitSetExpr(c *SetExprContext)

	// ExitTupleExpr is called when exiting the tupleExpr production.
	ExitTupleExpr(c *TupleExprContext)

	// ExitNamedTupleField is called when exiting the namedTupleField production.
	ExitNamedTupleField(c *NamedTupleFieldContext)

	// ExitBlock is called when exiting the block production.
	ExitBlock(c *BlockContext)

	// ExitArgList is called when exiting the argList production.
	ExitArgList(c *ArgListContext)

	// ExitLiteral is called when exiting the literal production.
	ExitLiteral(c *LiteralContext)

	// ExitIdentifier is called when exiting the identifier production.
	ExitIdentifier(c *IdentifierContext)
}
