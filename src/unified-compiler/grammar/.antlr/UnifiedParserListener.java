// Generated from /Users/kristofer/LocalProjects/Unified/src/unified-compiler/grammar/UnifiedParser.g4 by ANTLR 4.13.1
import org.antlr.v4.runtime.tree.ParseTreeListener;

/**
 * This interface defines a complete listener for a parse tree produced by
 * {@link UnifiedParser}.
 */
public interface UnifiedParserListener extends ParseTreeListener {
	/**
	 * Enter a parse tree produced by {@link UnifiedParser#program}.
	 * @param ctx the parse tree
	 */
	void enterProgram(UnifiedParser.ProgramContext ctx);
	/**
	 * Exit a parse tree produced by {@link UnifiedParser#program}.
	 * @param ctx the parse tree
	 */
	void exitProgram(UnifiedParser.ProgramContext ctx);
	/**
	 * Enter a parse tree produced by {@link UnifiedParser#item}.
	 * @param ctx the parse tree
	 */
	void enterItem(UnifiedParser.ItemContext ctx);
	/**
	 * Exit a parse tree produced by {@link UnifiedParser#item}.
	 * @param ctx the parse tree
	 */
	void exitItem(UnifiedParser.ItemContext ctx);
	/**
	 * Enter a parse tree produced by {@link UnifiedParser#moduleDecl}.
	 * @param ctx the parse tree
	 */
	void enterModuleDecl(UnifiedParser.ModuleDeclContext ctx);
	/**
	 * Exit a parse tree produced by {@link UnifiedParser#moduleDecl}.
	 * @param ctx the parse tree
	 */
	void exitModuleDecl(UnifiedParser.ModuleDeclContext ctx);
	/**
	 * Enter a parse tree produced by {@link UnifiedParser#importDecl}.
	 * @param ctx the parse tree
	 */
	void enterImportDecl(UnifiedParser.ImportDeclContext ctx);
	/**
	 * Exit a parse tree produced by {@link UnifiedParser#importDecl}.
	 * @param ctx the parse tree
	 */
	void exitImportDecl(UnifiedParser.ImportDeclContext ctx);
	/**
	 * Enter a parse tree produced by {@link UnifiedParser#importPath}.
	 * @param ctx the parse tree
	 */
	void enterImportPath(UnifiedParser.ImportPathContext ctx);
	/**
	 * Exit a parse tree produced by {@link UnifiedParser#importPath}.
	 * @param ctx the parse tree
	 */
	void exitImportPath(UnifiedParser.ImportPathContext ctx);
	/**
	 * Enter a parse tree produced by {@link UnifiedParser#importList}.
	 * @param ctx the parse tree
	 */
	void enterImportList(UnifiedParser.ImportListContext ctx);
	/**
	 * Exit a parse tree produced by {@link UnifiedParser#importList}.
	 * @param ctx the parse tree
	 */
	void exitImportList(UnifiedParser.ImportListContext ctx);
	/**
	 * Enter a parse tree produced by {@link UnifiedParser#functionDecl}.
	 * @param ctx the parse tree
	 */
	void enterFunctionDecl(UnifiedParser.FunctionDeclContext ctx);
	/**
	 * Exit a parse tree produced by {@link UnifiedParser#functionDecl}.
	 * @param ctx the parse tree
	 */
	void exitFunctionDecl(UnifiedParser.FunctionDeclContext ctx);
	/**
	 * Enter a parse tree produced by {@link UnifiedParser#paramList}.
	 * @param ctx the parse tree
	 */
	void enterParamList(UnifiedParser.ParamListContext ctx);
	/**
	 * Exit a parse tree produced by {@link UnifiedParser#paramList}.
	 * @param ctx the parse tree
	 */
	void exitParamList(UnifiedParser.ParamListContext ctx);
	/**
	 * Enter a parse tree produced by {@link UnifiedParser#parameter}.
	 * @param ctx the parse tree
	 */
	void enterParameter(UnifiedParser.ParameterContext ctx);
	/**
	 * Exit a parse tree produced by {@link UnifiedParser#parameter}.
	 * @param ctx the parse tree
	 */
	void exitParameter(UnifiedParser.ParameterContext ctx);
	/**
	 * Enter a parse tree produced by {@link UnifiedParser#genericParams}.
	 * @param ctx the parse tree
	 */
	void enterGenericParams(UnifiedParser.GenericParamsContext ctx);
	/**
	 * Exit a parse tree produced by {@link UnifiedParser#genericParams}.
	 * @param ctx the parse tree
	 */
	void exitGenericParams(UnifiedParser.GenericParamsContext ctx);
	/**
	 * Enter a parse tree produced by {@link UnifiedParser#genericParam}.
	 * @param ctx the parse tree
	 */
	void enterGenericParam(UnifiedParser.GenericParamContext ctx);
	/**
	 * Exit a parse tree produced by {@link UnifiedParser#genericParam}.
	 * @param ctx the parse tree
	 */
	void exitGenericParam(UnifiedParser.GenericParamContext ctx);
	/**
	 * Enter a parse tree produced by {@link UnifiedParser#typeConstraint}.
	 * @param ctx the parse tree
	 */
	void enterTypeConstraint(UnifiedParser.TypeConstraintContext ctx);
	/**
	 * Exit a parse tree produced by {@link UnifiedParser#typeConstraint}.
	 * @param ctx the parse tree
	 */
	void exitTypeConstraint(UnifiedParser.TypeConstraintContext ctx);
	/**
	 * Enter a parse tree produced by {@link UnifiedParser#whereClause}.
	 * @param ctx the parse tree
	 */
	void enterWhereClause(UnifiedParser.WhereClauseContext ctx);
	/**
	 * Exit a parse tree produced by {@link UnifiedParser#whereClause}.
	 * @param ctx the parse tree
	 */
	void exitWhereClause(UnifiedParser.WhereClauseContext ctx);
	/**
	 * Enter a parse tree produced by {@link UnifiedParser#whereConstraint}.
	 * @param ctx the parse tree
	 */
	void enterWhereConstraint(UnifiedParser.WhereConstraintContext ctx);
	/**
	 * Exit a parse tree produced by {@link UnifiedParser#whereConstraint}.
	 * @param ctx the parse tree
	 */
	void exitWhereConstraint(UnifiedParser.WhereConstraintContext ctx);
	/**
	 * Enter a parse tree produced by {@link UnifiedParser#structDecl}.
	 * @param ctx the parse tree
	 */
	void enterStructDecl(UnifiedParser.StructDeclContext ctx);
	/**
	 * Exit a parse tree produced by {@link UnifiedParser#structDecl}.
	 * @param ctx the parse tree
	 */
	void exitStructDecl(UnifiedParser.StructDeclContext ctx);
	/**
	 * Enter a parse tree produced by {@link UnifiedParser#structMember}.
	 * @param ctx the parse tree
	 */
	void enterStructMember(UnifiedParser.StructMemberContext ctx);
	/**
	 * Exit a parse tree produced by {@link UnifiedParser#structMember}.
	 * @param ctx the parse tree
	 */
	void exitStructMember(UnifiedParser.StructMemberContext ctx);
	/**
	 * Enter a parse tree produced by {@link UnifiedParser#enumDecl}.
	 * @param ctx the parse tree
	 */
	void enterEnumDecl(UnifiedParser.EnumDeclContext ctx);
	/**
	 * Exit a parse tree produced by {@link UnifiedParser#enumDecl}.
	 * @param ctx the parse tree
	 */
	void exitEnumDecl(UnifiedParser.EnumDeclContext ctx);
	/**
	 * Enter a parse tree produced by {@link UnifiedParser#enumVariant}.
	 * @param ctx the parse tree
	 */
	void enterEnumVariant(UnifiedParser.EnumVariantContext ctx);
	/**
	 * Exit a parse tree produced by {@link UnifiedParser#enumVariant}.
	 * @param ctx the parse tree
	 */
	void exitEnumVariant(UnifiedParser.EnumVariantContext ctx);
	/**
	 * Enter a parse tree produced by {@link UnifiedParser#enumVariantParams}.
	 * @param ctx the parse tree
	 */
	void enterEnumVariantParams(UnifiedParser.EnumVariantParamsContext ctx);
	/**
	 * Exit a parse tree produced by {@link UnifiedParser#enumVariantParams}.
	 * @param ctx the parse tree
	 */
	void exitEnumVariantParams(UnifiedParser.EnumVariantParamsContext ctx);
	/**
	 * Enter a parse tree produced by {@link UnifiedParser#enumVariantParam}.
	 * @param ctx the parse tree
	 */
	void enterEnumVariantParam(UnifiedParser.EnumVariantParamContext ctx);
	/**
	 * Exit a parse tree produced by {@link UnifiedParser#enumVariantParam}.
	 * @param ctx the parse tree
	 */
	void exitEnumVariantParam(UnifiedParser.EnumVariantParamContext ctx);
	/**
	 * Enter a parse tree produced by {@link UnifiedParser#interfaceDecl}.
	 * @param ctx the parse tree
	 */
	void enterInterfaceDecl(UnifiedParser.InterfaceDeclContext ctx);
	/**
	 * Exit a parse tree produced by {@link UnifiedParser#interfaceDecl}.
	 * @param ctx the parse tree
	 */
	void exitInterfaceDecl(UnifiedParser.InterfaceDeclContext ctx);
	/**
	 * Enter a parse tree produced by {@link UnifiedParser#interfaceMember}.
	 * @param ctx the parse tree
	 */
	void enterInterfaceMember(UnifiedParser.InterfaceMemberContext ctx);
	/**
	 * Exit a parse tree produced by {@link UnifiedParser#interfaceMember}.
	 * @param ctx the parse tree
	 */
	void exitInterfaceMember(UnifiedParser.InterfaceMemberContext ctx);
	/**
	 * Enter a parse tree produced by {@link UnifiedParser#functionSig}.
	 * @param ctx the parse tree
	 */
	void enterFunctionSig(UnifiedParser.FunctionSigContext ctx);
	/**
	 * Exit a parse tree produced by {@link UnifiedParser#functionSig}.
	 * @param ctx the parse tree
	 */
	void exitFunctionSig(UnifiedParser.FunctionSigContext ctx);
	/**
	 * Enter a parse tree produced by {@link UnifiedParser#implDecl}.
	 * @param ctx the parse tree
	 */
	void enterImplDecl(UnifiedParser.ImplDeclContext ctx);
	/**
	 * Exit a parse tree produced by {@link UnifiedParser#implDecl}.
	 * @param ctx the parse tree
	 */
	void exitImplDecl(UnifiedParser.ImplDeclContext ctx);
	/**
	 * Enter a parse tree produced by {@link UnifiedParser#implMember}.
	 * @param ctx the parse tree
	 */
	void enterImplMember(UnifiedParser.ImplMemberContext ctx);
	/**
	 * Exit a parse tree produced by {@link UnifiedParser#implMember}.
	 * @param ctx the parse tree
	 */
	void exitImplMember(UnifiedParser.ImplMemberContext ctx);
	/**
	 * Enter a parse tree produced by {@link UnifiedParser#actorDecl}.
	 * @param ctx the parse tree
	 */
	void enterActorDecl(UnifiedParser.ActorDeclContext ctx);
	/**
	 * Exit a parse tree produced by {@link UnifiedParser#actorDecl}.
	 * @param ctx the parse tree
	 */
	void exitActorDecl(UnifiedParser.ActorDeclContext ctx);
	/**
	 * Enter a parse tree produced by {@link UnifiedParser#actorMember}.
	 * @param ctx the parse tree
	 */
	void enterActorMember(UnifiedParser.ActorMemberContext ctx);
	/**
	 * Exit a parse tree produced by {@link UnifiedParser#actorMember}.
	 * @param ctx the parse tree
	 */
	void exitActorMember(UnifiedParser.ActorMemberContext ctx);
	/**
	 * Enter a parse tree produced by {@link UnifiedParser#typeAlias}.
	 * @param ctx the parse tree
	 */
	void enterTypeAlias(UnifiedParser.TypeAliasContext ctx);
	/**
	 * Exit a parse tree produced by {@link UnifiedParser#typeAlias}.
	 * @param ctx the parse tree
	 */
	void exitTypeAlias(UnifiedParser.TypeAliasContext ctx);
	/**
	 * Enter a parse tree produced by {@link UnifiedParser#constantDecl}.
	 * @param ctx the parse tree
	 */
	void enterConstantDecl(UnifiedParser.ConstantDeclContext ctx);
	/**
	 * Exit a parse tree produced by {@link UnifiedParser#constantDecl}.
	 * @param ctx the parse tree
	 */
	void exitConstantDecl(UnifiedParser.ConstantDeclContext ctx);
	/**
	 * Enter a parse tree produced by {@link UnifiedParser#type}.
	 * @param ctx the parse tree
	 */
	void enterType(UnifiedParser.TypeContext ctx);
	/**
	 * Exit a parse tree produced by {@link UnifiedParser#type}.
	 * @param ctx the parse tree
	 */
	void exitType(UnifiedParser.TypeContext ctx);
	/**
	 * Enter a parse tree produced by {@link UnifiedParser#typeList}.
	 * @param ctx the parse tree
	 */
	void enterTypeList(UnifiedParser.TypeListContext ctx);
	/**
	 * Exit a parse tree produced by {@link UnifiedParser#typeList}.
	 * @param ctx the parse tree
	 */
	void exitTypeList(UnifiedParser.TypeListContext ctx);
	/**
	 * Enter a parse tree produced by {@link UnifiedParser#statement}.
	 * @param ctx the parse tree
	 */
	void enterStatement(UnifiedParser.StatementContext ctx);
	/**
	 * Exit a parse tree produced by {@link UnifiedParser#statement}.
	 * @param ctx the parse tree
	 */
	void exitStatement(UnifiedParser.StatementContext ctx);
	/**
	 * Enter a parse tree produced by {@link UnifiedParser#letStatement}.
	 * @param ctx the parse tree
	 */
	void enterLetStatement(UnifiedParser.LetStatementContext ctx);
	/**
	 * Exit a parse tree produced by {@link UnifiedParser#letStatement}.
	 * @param ctx the parse tree
	 */
	void exitLetStatement(UnifiedParser.LetStatementContext ctx);
	/**
	 * Enter a parse tree produced by {@link UnifiedParser#varStatement}.
	 * @param ctx the parse tree
	 */
	void enterVarStatement(UnifiedParser.VarStatementContext ctx);
	/**
	 * Exit a parse tree produced by {@link UnifiedParser#varStatement}.
	 * @param ctx the parse tree
	 */
	void exitVarStatement(UnifiedParser.VarStatementContext ctx);
	/**
	 * Enter a parse tree produced by {@link UnifiedParser#regionStatement}.
	 * @param ctx the parse tree
	 */
	void enterRegionStatement(UnifiedParser.RegionStatementContext ctx);
	/**
	 * Exit a parse tree produced by {@link UnifiedParser#regionStatement}.
	 * @param ctx the parse tree
	 */
	void exitRegionStatement(UnifiedParser.RegionStatementContext ctx);
	/**
	 * Enter a parse tree produced by {@link UnifiedParser#exprStatement}.
	 * @param ctx the parse tree
	 */
	void enterExprStatement(UnifiedParser.ExprStatementContext ctx);
	/**
	 * Exit a parse tree produced by {@link UnifiedParser#exprStatement}.
	 * @param ctx the parse tree
	 */
	void exitExprStatement(UnifiedParser.ExprStatementContext ctx);
	/**
	 * Enter a parse tree produced by {@link UnifiedParser#returnStatement}.
	 * @param ctx the parse tree
	 */
	void enterReturnStatement(UnifiedParser.ReturnStatementContext ctx);
	/**
	 * Exit a parse tree produced by {@link UnifiedParser#returnStatement}.
	 * @param ctx the parse tree
	 */
	void exitReturnStatement(UnifiedParser.ReturnStatementContext ctx);
	/**
	 * Enter a parse tree produced by {@link UnifiedParser#ifStatement}.
	 * @param ctx the parse tree
	 */
	void enterIfStatement(UnifiedParser.IfStatementContext ctx);
	/**
	 * Exit a parse tree produced by {@link UnifiedParser#ifStatement}.
	 * @param ctx the parse tree
	 */
	void exitIfStatement(UnifiedParser.IfStatementContext ctx);
	/**
	 * Enter a parse tree produced by {@link UnifiedParser#loopStatement}.
	 * @param ctx the parse tree
	 */
	void enterLoopStatement(UnifiedParser.LoopStatementContext ctx);
	/**
	 * Exit a parse tree produced by {@link UnifiedParser#loopStatement}.
	 * @param ctx the parse tree
	 */
	void exitLoopStatement(UnifiedParser.LoopStatementContext ctx);
	/**
	 * Enter a parse tree produced by {@link UnifiedParser#whileStatement}.
	 * @param ctx the parse tree
	 */
	void enterWhileStatement(UnifiedParser.WhileStatementContext ctx);
	/**
	 * Exit a parse tree produced by {@link UnifiedParser#whileStatement}.
	 * @param ctx the parse tree
	 */
	void exitWhileStatement(UnifiedParser.WhileStatementContext ctx);
	/**
	 * Enter a parse tree produced by {@link UnifiedParser#forStatement}.
	 * @param ctx the parse tree
	 */
	void enterForStatement(UnifiedParser.ForStatementContext ctx);
	/**
	 * Exit a parse tree produced by {@link UnifiedParser#forStatement}.
	 * @param ctx the parse tree
	 */
	void exitForStatement(UnifiedParser.ForStatementContext ctx);
	/**
	 * Enter a parse tree produced by {@link UnifiedParser#switchStatement}.
	 * @param ctx the parse tree
	 */
	void enterSwitchStatement(UnifiedParser.SwitchStatementContext ctx);
	/**
	 * Exit a parse tree produced by {@link UnifiedParser#switchStatement}.
	 * @param ctx the parse tree
	 */
	void exitSwitchStatement(UnifiedParser.SwitchStatementContext ctx);
	/**
	 * Enter a parse tree produced by {@link UnifiedParser#caseClause}.
	 * @param ctx the parse tree
	 */
	void enterCaseClause(UnifiedParser.CaseClauseContext ctx);
	/**
	 * Exit a parse tree produced by {@link UnifiedParser#caseClause}.
	 * @param ctx the parse tree
	 */
	void exitCaseClause(UnifiedParser.CaseClauseContext ctx);
	/**
	 * Enter a parse tree produced by {@link UnifiedParser#breakStatement}.
	 * @param ctx the parse tree
	 */
	void enterBreakStatement(UnifiedParser.BreakStatementContext ctx);
	/**
	 * Exit a parse tree produced by {@link UnifiedParser#breakStatement}.
	 * @param ctx the parse tree
	 */
	void exitBreakStatement(UnifiedParser.BreakStatementContext ctx);
	/**
	 * Enter a parse tree produced by {@link UnifiedParser#continueStatement}.
	 * @param ctx the parse tree
	 */
	void enterContinueStatement(UnifiedParser.ContinueStatementContext ctx);
	/**
	 * Exit a parse tree produced by {@link UnifiedParser#continueStatement}.
	 * @param ctx the parse tree
	 */
	void exitContinueStatement(UnifiedParser.ContinueStatementContext ctx);
	/**
	 * Enter a parse tree produced by {@link UnifiedParser#blockStatement}.
	 * @param ctx the parse tree
	 */
	void enterBlockStatement(UnifiedParser.BlockStatementContext ctx);
	/**
	 * Exit a parse tree produced by {@link UnifiedParser#blockStatement}.
	 * @param ctx the parse tree
	 */
	void exitBlockStatement(UnifiedParser.BlockStatementContext ctx);
	/**
	 * Enter a parse tree produced by {@link UnifiedParser#tryStatement}.
	 * @param ctx the parse tree
	 */
	void enterTryStatement(UnifiedParser.TryStatementContext ctx);
	/**
	 * Exit a parse tree produced by {@link UnifiedParser#tryStatement}.
	 * @param ctx the parse tree
	 */
	void exitTryStatement(UnifiedParser.TryStatementContext ctx);
	/**
	 * Enter a parse tree produced by {@link UnifiedParser#pattern}.
	 * @param ctx the parse tree
	 */
	void enterPattern(UnifiedParser.PatternContext ctx);
	/**
	 * Exit a parse tree produced by {@link UnifiedParser#pattern}.
	 * @param ctx the parse tree
	 */
	void exitPattern(UnifiedParser.PatternContext ctx);
	/**
	 * Enter a parse tree produced by {@link UnifiedParser#patternList}.
	 * @param ctx the parse tree
	 */
	void enterPatternList(UnifiedParser.PatternListContext ctx);
	/**
	 * Exit a parse tree produced by {@link UnifiedParser#patternList}.
	 * @param ctx the parse tree
	 */
	void exitPatternList(UnifiedParser.PatternListContext ctx);
	/**
	 * Enter a parse tree produced by {@link UnifiedParser#fieldPattern}.
	 * @param ctx the parse tree
	 */
	void enterFieldPattern(UnifiedParser.FieldPatternContext ctx);
	/**
	 * Exit a parse tree produced by {@link UnifiedParser#fieldPattern}.
	 * @param ctx the parse tree
	 */
	void exitFieldPattern(UnifiedParser.FieldPatternContext ctx);
	/**
	 * Enter a parse tree produced by {@link UnifiedParser#fieldPatternList}.
	 * @param ctx the parse tree
	 */
	void enterFieldPatternList(UnifiedParser.FieldPatternListContext ctx);
	/**
	 * Exit a parse tree produced by {@link UnifiedParser#fieldPatternList}.
	 * @param ctx the parse tree
	 */
	void exitFieldPatternList(UnifiedParser.FieldPatternListContext ctx);
	/**
	 * Enter a parse tree produced by {@link UnifiedParser#expr}.
	 * @param ctx the parse tree
	 */
	void enterExpr(UnifiedParser.ExprContext ctx);
	/**
	 * Exit a parse tree produced by {@link UnifiedParser#expr}.
	 * @param ctx the parse tree
	 */
	void exitExpr(UnifiedParser.ExprContext ctx);
	/**
	 * Enter a parse tree produced by {@link UnifiedParser#caseExpr}.
	 * @param ctx the parse tree
	 */
	void enterCaseExpr(UnifiedParser.CaseExprContext ctx);
	/**
	 * Exit a parse tree produced by {@link UnifiedParser#caseExpr}.
	 * @param ctx the parse tree
	 */
	void exitCaseExpr(UnifiedParser.CaseExprContext ctx);
	/**
	 * Enter a parse tree produced by {@link UnifiedParser#primary}.
	 * @param ctx the parse tree
	 */
	void enterPrimary(UnifiedParser.PrimaryContext ctx);
	/**
	 * Exit a parse tree produced by {@link UnifiedParser#primary}.
	 * @param ctx the parse tree
	 */
	void exitPrimary(UnifiedParser.PrimaryContext ctx);
	/**
	 * Enter a parse tree produced by {@link UnifiedParser#lambdaExpr}.
	 * @param ctx the parse tree
	 */
	void enterLambdaExpr(UnifiedParser.LambdaExprContext ctx);
	/**
	 * Exit a parse tree produced by {@link UnifiedParser#lambdaExpr}.
	 * @param ctx the parse tree
	 */
	void exitLambdaExpr(UnifiedParser.LambdaExprContext ctx);
	/**
	 * Enter a parse tree produced by {@link UnifiedParser#asyncExpr}.
	 * @param ctx the parse tree
	 */
	void enterAsyncExpr(UnifiedParser.AsyncExprContext ctx);
	/**
	 * Exit a parse tree produced by {@link UnifiedParser#asyncExpr}.
	 * @param ctx the parse tree
	 */
	void exitAsyncExpr(UnifiedParser.AsyncExprContext ctx);
	/**
	 * Enter a parse tree produced by {@link UnifiedParser#structExpr}.
	 * @param ctx the parse tree
	 */
	void enterStructExpr(UnifiedParser.StructExprContext ctx);
	/**
	 * Exit a parse tree produced by {@link UnifiedParser#structExpr}.
	 * @param ctx the parse tree
	 */
	void exitStructExpr(UnifiedParser.StructExprContext ctx);
	/**
	 * Enter a parse tree produced by {@link UnifiedParser#fieldInitList}.
	 * @param ctx the parse tree
	 */
	void enterFieldInitList(UnifiedParser.FieldInitListContext ctx);
	/**
	 * Exit a parse tree produced by {@link UnifiedParser#fieldInitList}.
	 * @param ctx the parse tree
	 */
	void exitFieldInitList(UnifiedParser.FieldInitListContext ctx);
	/**
	 * Enter a parse tree produced by {@link UnifiedParser#fieldInit}.
	 * @param ctx the parse tree
	 */
	void enterFieldInit(UnifiedParser.FieldInitContext ctx);
	/**
	 * Exit a parse tree produced by {@link UnifiedParser#fieldInit}.
	 * @param ctx the parse tree
	 */
	void exitFieldInit(UnifiedParser.FieldInitContext ctx);
	/**
	 * Enter a parse tree produced by {@link UnifiedParser#listExpr}.
	 * @param ctx the parse tree
	 */
	void enterListExpr(UnifiedParser.ListExprContext ctx);
	/**
	 * Exit a parse tree produced by {@link UnifiedParser#listExpr}.
	 * @param ctx the parse tree
	 */
	void exitListExpr(UnifiedParser.ListExprContext ctx);
	/**
	 * Enter a parse tree produced by {@link UnifiedParser#mapExpr}.
	 * @param ctx the parse tree
	 */
	void enterMapExpr(UnifiedParser.MapExprContext ctx);
	/**
	 * Exit a parse tree produced by {@link UnifiedParser#mapExpr}.
	 * @param ctx the parse tree
	 */
	void exitMapExpr(UnifiedParser.MapExprContext ctx);
	/**
	 * Enter a parse tree produced by {@link UnifiedParser#keyValue}.
	 * @param ctx the parse tree
	 */
	void enterKeyValue(UnifiedParser.KeyValueContext ctx);
	/**
	 * Exit a parse tree produced by {@link UnifiedParser#keyValue}.
	 * @param ctx the parse tree
	 */
	void exitKeyValue(UnifiedParser.KeyValueContext ctx);
	/**
	 * Enter a parse tree produced by {@link UnifiedParser#setExpr}.
	 * @param ctx the parse tree
	 */
	void enterSetExpr(UnifiedParser.SetExprContext ctx);
	/**
	 * Exit a parse tree produced by {@link UnifiedParser#setExpr}.
	 * @param ctx the parse tree
	 */
	void exitSetExpr(UnifiedParser.SetExprContext ctx);
	/**
	 * Enter a parse tree produced by {@link UnifiedParser#tupleExpr}.
	 * @param ctx the parse tree
	 */
	void enterTupleExpr(UnifiedParser.TupleExprContext ctx);
	/**
	 * Exit a parse tree produced by {@link UnifiedParser#tupleExpr}.
	 * @param ctx the parse tree
	 */
	void exitTupleExpr(UnifiedParser.TupleExprContext ctx);
	/**
	 * Enter a parse tree produced by {@link UnifiedParser#block}.
	 * @param ctx the parse tree
	 */
	void enterBlock(UnifiedParser.BlockContext ctx);
	/**
	 * Exit a parse tree produced by {@link UnifiedParser#block}.
	 * @param ctx the parse tree
	 */
	void exitBlock(UnifiedParser.BlockContext ctx);
	/**
	 * Enter a parse tree produced by {@link UnifiedParser#argList}.
	 * @param ctx the parse tree
	 */
	void enterArgList(UnifiedParser.ArgListContext ctx);
	/**
	 * Exit a parse tree produced by {@link UnifiedParser#argList}.
	 * @param ctx the parse tree
	 */
	void exitArgList(UnifiedParser.ArgListContext ctx);
	/**
	 * Enter a parse tree produced by {@link UnifiedParser#literal}.
	 * @param ctx the parse tree
	 */
	void enterLiteral(UnifiedParser.LiteralContext ctx);
	/**
	 * Exit a parse tree produced by {@link UnifiedParser#literal}.
	 * @param ctx the parse tree
	 */
	void exitLiteral(UnifiedParser.LiteralContext ctx);
	/**
	 * Enter a parse tree produced by {@link UnifiedParser#identifier}.
	 * @param ctx the parse tree
	 */
	void enterIdentifier(UnifiedParser.IdentifierContext ctx);
	/**
	 * Exit a parse tree produced by {@link UnifiedParser#identifier}.
	 * @param ctx the parse tree
	 */
	void exitIdentifier(UnifiedParser.IdentifierContext ctx);
}