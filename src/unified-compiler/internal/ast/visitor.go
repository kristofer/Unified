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
// func (v *ASTBuilder) VisitItem(ctx *parser.ItemContext) interface{} {
// 	// Check which type of item we have
// 	if moduleCtx := ctx.ModuleDecl(); moduleCtx != nil {
// 		return v.VisitModuleDecl(moduleCtx.(*parser.ModuleDeclContext))
// 	}
// 	// Add handlers for other item types...

// 	return nil
// }

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

// Update the VisitItem method to handle more item types
func (v *ASTBuilder) VisitItem(ctx *parser.ItemContext) interface{} {
	// Check which type of item we have
	if moduleCtx := ctx.ModuleDecl(); moduleCtx != nil {
		return v.VisitModuleDecl(moduleCtx.(*parser.ModuleDeclContext))
	}
	if funcCtx := ctx.FunctionDecl(); funcCtx != nil {
		return v.VisitFunctionDecl(funcCtx.(*parser.FunctionDeclContext))
	}
	if structCtx := ctx.StructDecl(); structCtx != nil {
		return v.VisitStructDecl(structCtx.(*parser.StructDeclContext))
	}
	// Add more item types as needed...

	return nil
}

// VisitFunctionDecl builds a FunctionDecl node
func (v *ASTBuilder) VisitFunctionDecl(ctx *parser.FunctionDeclContext) interface{} {
	// Check if function is public
	isPublic := ctx.PUB() != nil

	// Get function name
	name := ctx.Identifier().GetText()

	// Get parameters
	var params []*Parameter
	if paramListCtx := ctx.ParamList(); paramListCtx != nil {
		params = v.processParamList(paramListCtx.(*parser.ParamListContext))
	}

	// Get return type if specified
	var returnType Type
	if typeCtx := ctx.Type_(); typeCtx != nil {
		returnType = v.VisitType_(typeCtx.(*parser.TypeContext)).(Type)
	}

	// Get generic parameters if specified
	var genericParams []*GenericParam
	if genericParamsCtx := ctx.GenericParams(); genericParamsCtx != nil {
		genericParams = v.processGenericParams(genericParamsCtx.(*parser.GenericParamsContext))
	}

	// Get where clause if specified
	var whereConstraints []*WhereConstraint
	if whereClauseCtx := ctx.WhereClause(); whereClauseCtx != nil {
		whereConstraints = v.processWhereClause(whereClauseCtx.(*parser.WhereClauseContext))
	}

	// Get function body
	body := v.VisitBlock(ctx.Block().(*parser.BlockContext)).(*Block)

	return &FunctionDecl{
		Name:             name,
		IsPublic:         isPublic,
		Parameters:       params,
		ReturnType:       returnType,
		GenericParams:    genericParams,
		WhereConstraints: whereConstraints,
		Body:             body,
		Position:         v.getPosition(ctx),
	}
}

// Helper method to process parameter list
func (v *ASTBuilder) processParamList(ctx *parser.ParamListContext) []*Parameter {
	params := []*Parameter{}

	for _, paramCtx := range ctx.AllParameter() {
		param := v.VisitParameter(paramCtx.(*parser.ParameterContext)).(*Parameter)
		params = append(params, param)
	}

	return params
}

// VisitParameter builds a Parameter node
func (v *ASTBuilder) VisitParameter(ctx *parser.ParameterContext) interface{} {
	// Handle self parameter specially
	if ctx.SELF(0) != nil {
		isMut := ctx.MUT() != nil
		isRef := ctx.BIT_AND() != nil

		return &Parameter{
			IsSelf:      true,
			IsReference: isRef,
			IsMutable:   isMut,
			Position:    v.getPosition(ctx),
		}
	}

	// Regular parameter with name and type
	name := ctx.Identifier().GetText()
	paramType := v.VisitType_(ctx.Type_().(*parser.TypeContext)).(Type) // Note the underscore

	return &Parameter{
		Name:     name,
		Type:     paramType,
		Position: v.getPosition(ctx),
	}
}

// func (v *ASTBuilder) VisitParameter(ctx *parser.ParameterContext) interface{} {
// 	// Handle self parameter specially
// 	if ctx.SELF(0) != nil {
// 		isMut := ctx.MUT() != nil
// 		isRef := ctx.BIT_AND() != nil

// 		return &Parameter{
// 			IsSelf:      true,
// 			IsReference: isRef,
// 			IsMutable:   isMut,
// 			Position:    v.getPosition(ctx),
// 		}
// 	}

// 	// Regular parameter with name and type
// 	name := ctx.Identifier().GetText()
// 	paramType := v.VisitType(ctx.Type().(*parser.TypeContext)).(Type)

// 	return &Parameter{
// 		Name:     name,
// 		Type:     paramType,
// 		Position: v.getPosition(ctx),
// 	}
// }

// VisitStructDecl builds a StructDecl node
func (v *ASTBuilder) VisitStructDecl(ctx *parser.StructDeclContext) interface{} {
	// Check if struct is public
	isPublic := ctx.PUB() != nil

	// Get struct name
	name := ctx.Identifier().GetText()

	// Process generic parameters if any
	var genericParams []*GenericParam
	if genericParamsCtx := ctx.GenericParams(); genericParamsCtx != nil {
		genericParams = v.processGenericParams(genericParamsCtx.(*parser.GenericParamsContext))
	}

	// Process struct members
	members := []*StructMember{}
	for _, memberCtx := range ctx.AllStructMember() {
		member := v.VisitStructMember(memberCtx.(*parser.StructMemberContext))
		if member != nil {
			members = append(members, member.(*StructMember))
		}
	}

	return &StructDecl{
		Name:          name,
		IsPublic:      isPublic,
		GenericParams: genericParams,
		Members:       members,
		Position:      v.getPosition(ctx),
	}
}

// VisitStructMember builds a StructMember node
func (v *ASTBuilder) VisitStructMember(ctx *parser.StructMemberContext) interface{} {
	// Check if it's a field or a method
	if ctx.Identifier() != nil {
		// It's a field
		isPublic := ctx.PUB() != nil
		name := ctx.Identifier().GetText()
		fieldType := v.VisitType_(ctx.Type_().(*parser.TypeContext)).(Type)

		return &StructMember{
			Name:     name,
			IsPublic: isPublic,
			Type:     fieldType,
			IsMethod: false,
			Method:   nil,
			Position: v.getPosition(ctx),
		}
	} else if ctx.FunctionDecl() != nil {
		// It's a method
		method := v.VisitFunctionDecl(ctx.FunctionDecl().(*parser.FunctionDeclContext)).(*FunctionDecl)

		return &StructMember{
			IsMethod: true,
			Method:   method,
			Position: v.getPosition(ctx),
		}
	}

	return nil
}

// VisitType handles type references
func (v *ASTBuilder) VisitType_(ctx *parser.TypeContext) interface{} {
	// Simple/generic type
	if ctx.Identifier(0) != nil {
		typeName := ctx.Identifier(0).GetText()

		// Check for generic type arguments
		var typeArgs []Type
		if ctx.TypeList() != nil {
			typeArgs = v.processTypeList(ctx.TypeList().(*parser.TypeListContext))
		}

		return &TypeReference{
			Name:     typeName,
			TypeArgs: typeArgs,
			Position: v.getPosition(ctx),
		}
	}

	// Function type
	if ctx.FN() != nil {
		var paramTypes []Type
		if ctx.TypeList() != nil {
			paramTypes = v.processTypeList(ctx.TypeList().(*parser.TypeListContext))
		}

		returnType := v.VisitType_(ctx.Type_(0).(*parser.TypeContext)).(Type)

		return &FunctionType{
			ParamTypes: paramTypes,
			ReturnType: returnType,
			Position:   v.getPosition(ctx),
		}
	}

	// Other type forms...
	// Add handling for reference types, tuple types, etc.

	return &TypeReference{
		Name:     "unknown",
		Position: v.getPosition(ctx),
	}
}

// Helper to process type lists
func (v *ASTBuilder) processTypeList(ctx *parser.TypeListContext) []Type {
	types := []Type{}

	for _, typeCtx := range ctx.AllType_() {
		t := v.VisitType_(typeCtx.(*parser.TypeContext)).(Type)
		types = append(types, t)
	}

	return types
}

// Helper to process generic parameters
func (v *ASTBuilder) processGenericParams(ctx *parser.GenericParamsContext) []*GenericParam {
	params := []*GenericParam{}

	for _, paramCtx := range ctx.AllGenericParam() {
		param := v.VisitGenericParam(paramCtx.(*parser.GenericParamContext)).(*GenericParam)
		params = append(params, param)
	}

	return params
}

// VisitGenericParam builds a GenericParam node
func (v *ASTBuilder) VisitGenericParam(ctx *parser.GenericParamContext) interface{} {
	name := ctx.Identifier().GetText()

	var constraints []Type
	if ctx.TypeConstraint() != nil {
		constraints = v.processTypeConstraint(ctx.TypeConstraint().(*parser.TypeConstraintContext))
	}

	return &GenericParam{
		Name:        name,
		Constraints: constraints,
		Position:    v.getPosition(ctx),
	}
}

// Helper to process type constraints
func (v *ASTBuilder) processTypeConstraint(ctx *parser.TypeConstraintContext) []Type {
	constraints := []Type{}

	for _, typeCtx := range ctx.AllType_() {
		t := v.VisitType_(typeCtx.(*parser.TypeContext)).(Type)
		constraints = append(constraints, t)
	}

	return constraints
}

// Helper to process where clauses
func (v *ASTBuilder) processWhereClause(ctx *parser.WhereClauseContext) []*WhereConstraint {
	constraints := []*WhereConstraint{}

	for _, constraintCtx := range ctx.AllWhereConstraint() {
		constraint := v.VisitWhereConstraint(constraintCtx.(*parser.WhereConstraintContext)).(*WhereConstraint)
		constraints = append(constraints, constraint)
	}

	return constraints
}

// VisitWhereConstraint builds a WhereConstraint node
func (v *ASTBuilder) VisitWhereConstraint(ctx *parser.WhereConstraintContext) interface{} {
	subjectType := v.VisitType_(ctx.Type_().(*parser.TypeContext)).(Type)
	constraints := v.processTypeConstraint(ctx.TypeConstraint().(*parser.TypeConstraintContext))

	return &WhereConstraint{
		SubjectType: subjectType,
		Constraints: constraints,
		Position:    v.getPosition(ctx),
	}
}

// VisitBlock builds a Block node
func (v *ASTBuilder) VisitBlock(ctx *parser.BlockContext) interface{} {
	statements := []Statement{}

	for _, stmtCtx := range ctx.AllStatement() {
		stmt := v.VisitStatement(stmtCtx.(*parser.StatementContext))
		if stmt != nil {
			statements = append(statements, stmt.(Statement))
		}
	}

	// Process optional trailing expression
	var expr Expression
	if exprCtx := ctx.Expr(); exprCtx != nil {
		expr = v.VisitExpr(exprCtx.(*parser.ExprContext)).(Expression)
	}

	return &Block{
		Statements: statements,
		Expression: expr,
		Position:   v.getPosition(ctx),
	}
}

// VisitExpr builds an Expression node
func (v *ASTBuilder) VisitExpr(ctx *parser.ExprContext) interface{} {
	// Handle primary expressions (identifiers, literals, etc.)
	if primary := ctx.Primary(); primary != nil {
		return v.VisitPrimary(primary.(*parser.PrimaryContext))
	}

	// Binary operations (handle according to child count and operator)
	if len(ctx.AllExpr()) == 2 {
		// Binary operations
		left := v.VisitExpr(ctx.Expr(0).(*parser.ExprContext)).(Expression)
		right := v.VisitExpr(ctx.Expr(1).(*parser.ExprContext)).(Expression)

		// Check the operator type
		if ctx.PLUS() != nil {
			return &BinaryExpr{
				Left:     left,
				Operator: OperatorAdd,
				Right:    right,
				Position: v.getPosition(ctx),
			}
		}
		// Add similar cases for other binary operators...
	}

	// Unary operations
	if len(ctx.AllExpr()) == 1 && (ctx.PLUS() != nil || ctx.MINUS() != nil || ctx.NOT() != nil) {
		operand := v.VisitExpr(ctx.Expr(0).(*parser.ExprContext)).(Expression)
		var op OperatorType

		if ctx.PLUS() != nil {
			op = OperatorUnaryPlus
		} else if ctx.MINUS() != nil {
			op = OperatorUnaryMinus
		} else if ctx.NOT() != nil {
			op = OperatorNot
		}

		return &UnaryExpr{
			Operator: op,
			Operand:  operand,
			Position: v.getPosition(ctx),
		}
	}

	// Add cases for method calls, field access, etc.

	// And more cases...

	return nil
}

// VisitPrimary handles primary expressions
func (v *ASTBuilder) VisitPrimary(ctx *parser.PrimaryContext) interface{} {
	if literalCtx := ctx.Literal(); literalCtx != nil {
		return v.VisitLiteral(literalCtx.(*parser.LiteralContext))
	}
	if identifierCtx := ctx.Identifier(); identifierCtx != nil {
		return &Identifier{
			Name:     identifierCtx.GetText(),
			Position: v.getPosition(ctx),
		}
	}
	// Handle other primary expressions...

	return nil
}

// VisitLiteral builds a Literal node
func (v *ASTBuilder) VisitLiteral(ctx *parser.LiteralContext) interface{} {
	if ctx.IntLiteral() != nil {
		return &Literal{
			Kind:     LiteralInt,
			Value:    ctx.IntLiteral().GetText(),
			Position: v.getPosition(ctx),
		}
	}
	if ctx.FloatLiteral() != nil {
		return &Literal{
			Kind:     LiteralFloat,
			Value:    ctx.FloatLiteral().GetText(),
			Position: v.getPosition(ctx),
		}
	}
	if ctx.StringLiteral() != nil {
		return &Literal{
			Kind:     LiteralString,
			Value:    ctx.StringLiteral().GetText(),
			Position: v.getPosition(ctx),
		}
	}
	if ctx.CharLiteral() != nil {
		return &Literal{
			Kind:     LiteralChar,
			Value:    ctx.CharLiteral().GetText(),
			Position: v.getPosition(ctx),
		}
	}
	if ctx.BoolLiteral() != nil {
		return &Literal{
			Kind:     LiteralBool,
			Value:    ctx.BoolLiteral().GetText(),
			Position: v.getPosition(ctx),
		}
	}
	if ctx.NullLiteral() != nil {
		return &Literal{
			Kind:     LiteralNull,
			Value:    "null",
			Position: v.getPosition(ctx),
		}
	}

	return nil
}

// VisitStatement handles all statement types
func (v *ASTBuilder) VisitStatement(ctx *parser.StatementContext) interface{} {
	// Check which type of statement we have
	if letCtx := ctx.LetStatement(); letCtx != nil {
		return v.VisitLetStatement(letCtx.(*parser.LetStatementContext))
	}
	if varCtx := ctx.VarStatement(); varCtx != nil {
		return v.VisitVarStatement(varCtx.(*parser.VarStatementContext))
	}
	if regionCtx := ctx.RegionStatement(); regionCtx != nil {
		return v.VisitRegionStatement(regionCtx.(*parser.RegionStatementContext))
	}
	if exprCtx := ctx.ExprStatement(); exprCtx != nil {
		return v.VisitExprStatement(exprCtx.(*parser.ExprStatementContext))
	}
	if returnCtx := ctx.ReturnStatement(); returnCtx != nil {
		return v.VisitReturnStatement(returnCtx.(*parser.ReturnStatementContext))
	}
	if ifCtx := ctx.IfStatement(); ifCtx != nil {
		return v.VisitIfStatement(ifCtx.(*parser.IfStatementContext))
	}
	if loopCtx := ctx.LoopStatement(); loopCtx != nil {
		return v.VisitLoopStatement(loopCtx.(*parser.LoopStatementContext))
	}
	if whileCtx := ctx.WhileStatement(); whileCtx != nil {
		return v.VisitWhileStatement(whileCtx.(*parser.WhileStatementContext))
	}
	if forCtx := ctx.ForStatement(); forCtx != nil {
		return v.VisitForStatement(forCtx.(*parser.ForStatementContext))
	}
	if switchCtx := ctx.SwitchStatement(); switchCtx != nil {
		return v.VisitSwitchStatement(switchCtx.(*parser.SwitchStatementContext))
	}
	if breakCtx := ctx.BreakStatement(); breakCtx != nil {
		return v.VisitBreakStatement(breakCtx.(*parser.BreakStatementContext))
	}
	if continueCtx := ctx.ContinueStatement(); continueCtx != nil {
		return v.VisitContinueStatement(continueCtx.(*parser.ContinueStatementContext))
	}
	if blockCtx := ctx.BlockStatement(); blockCtx != nil {
		return v.VisitBlockStatement(blockCtx.(*parser.BlockStatementContext))
	}
	return nil
}

// VisitLetStatement builds a let statement node
func (v *ASTBuilder) VisitLetStatement(ctx *parser.LetStatementContext) interface{} {
	name := ctx.Identifier().GetText()
	isMutable := ctx.MUT() != nil

	var varType Type
	if typeCtx := ctx.Type_(); typeCtx != nil {
		varType = v.VisitType_(typeCtx.(*parser.TypeContext)).(Type)
	}

	expr := v.VisitExpr(ctx.Expr().(*parser.ExprContext)).(Expression)

	return &LetStatement{
		Name:      name,
		IsMutable: isMutable,
		Type:      varType,
		Value:     expr,
		Position:  v.getPosition(ctx),
	}
}

// VisitReturnStatement builds a return statement node
func (v *ASTBuilder) VisitReturnStatement(ctx *parser.ReturnStatementContext) interface{} {
	var expr Expression
	if exprCtx := ctx.Expr(); exprCtx != nil {
		expr = v.VisitExpr(exprCtx.(*parser.ExprContext)).(Expression)
	}

	return &ReturnStatement{
		Value:    expr,
		Position: v.getPosition(ctx),
	}
}

// VisitExprStatement builds an expression statement node
func (v *ASTBuilder) VisitExprStatement(ctx *parser.ExprStatementContext) interface{} {
	expr := v.VisitExpr(ctx.Expr().(*parser.ExprContext)).(Expression)

	return &ExprStatement{
		Expression: expr,
		Position:   v.getPosition(ctx),
	}
}

// VisitVarStatement builds a variable statement node
func (v *ASTBuilder) VisitVarStatement(ctx *parser.VarStatementContext) interface{} {
	name := ctx.Identifier().GetText()
	varType := v.VisitType_(ctx.Type_().(*parser.TypeContext)).(Type)

	var expr Expression
	if exprCtx := ctx.Expr(); exprCtx != nil {
		expr = v.VisitExpr(exprCtx.(*parser.ExprContext)).(Expression)
	}

	return &VarStatement{
		Name:     name,
		Type:     varType,
		Value:    expr,
		Position: v.getPosition(ctx),
	}
}

// VisitRegionStatement builds a region statement node
func (v *ASTBuilder) VisitRegionStatement(ctx *parser.RegionStatementContext) interface{} {
	name := ctx.Identifier().GetText()
	body := v.VisitBlock(ctx.Block().(*parser.BlockContext)).(*Block)

	return &RegionStatement{
		Name:     name,
		Body:     body,
		Position: v.getPosition(ctx),
	}
}

// VisitIfStatement builds an if statement node
func (v *ASTBuilder) VisitIfStatement(ctx *parser.IfStatementContext) interface{} {
	// Process main if condition and block
	condition := v.VisitExpr(ctx.Expr(0).(*parser.ExprContext)).(Expression)
	thenBlock := v.VisitBlock(ctx.Block(0).(*parser.BlockContext)).(*Block)

	// Process else-if branches
	var elseIfBranches []*IfBranch

	// Count how many else-if branches we have
	elseIfCount := 0
	for i := 0; i < len(ctx.AllELSE()) && i < len(ctx.AllIF())-1; i++ {
		if ctx.ELSE(i) != nil && i+1 < len(ctx.AllIF()) {
			elseIfCount++
		}
	}

	// Process each else-if branch
	for i := 0; i < elseIfCount; i++ {
		// The condition is the expression after the first IF
		exprIndex := i + 1  // Skip the first expression which belongs to the main if
		blockIndex := i + 1 // Skip the first block which belongs to the main if

		elseIfCondition := v.VisitExpr(ctx.Expr(exprIndex).(*parser.ExprContext)).(Expression)
		elseIfBlock := v.VisitBlock(ctx.Block(blockIndex).(*parser.BlockContext)).(*Block)

		elseIfBranches = append(elseIfBranches, &IfBranch{
			Condition: elseIfCondition,
			Body:      elseIfBlock,
		})
	}

	// Process else block if present
	var elseBlock *Block
	if len(ctx.AllELSE()) > len(ctx.AllIF())-1 { // There's one more ELSE than IF-1
		// The last block is the else block
		elseBlockIndex := len(ctx.AllBlock()) - 1
		elseBlock = v.VisitBlock(ctx.Block(elseBlockIndex).(*parser.BlockContext)).(*Block)
	}

	return &IfStatement{
		Condition:      condition,
		ThenBlock:      thenBlock,
		ElseIfBranches: elseIfBranches,
		ElseBlock:      elseBlock,
		Position:       v.getPosition(ctx),
	}
}

// VisitLoopStatement builds a loop statement node
func (v *ASTBuilder) VisitLoopStatement(ctx *parser.LoopStatementContext) interface{} {
	var label string
	if ctx.Identifier() != nil {
		label = ctx.Identifier().GetText()
	}

	body := v.VisitBlock(ctx.Block().(*parser.BlockContext)).(*Block)

	return &LoopStatement{
		Label:    label,
		Body:     body,
		Position: v.getPosition(ctx),
	}
}

// VisitWhileStatement builds a while statement node
func (v *ASTBuilder) VisitWhileStatement(ctx *parser.WhileStatementContext) interface{} {
	var label string
	if ctx.Identifier() != nil {
		label = ctx.Identifier().GetText()
	}

	condition := v.VisitExpr(ctx.Expr().(*parser.ExprContext)).(Expression)
	body := v.VisitBlock(ctx.Block().(*parser.BlockContext)).(*Block)

	return &WhileStatement{
		Label:     label,
		Condition: condition,
		Body:      body,
		Position:  v.getPosition(ctx),
	}
}

// VisitForStatement builds a for statement node
func (v *ASTBuilder) VisitForStatement(ctx *parser.ForStatementContext) interface{} {
	var label string
	var iterVar string

	// Check if we have a label (which would be followed by a colon)
	if ctx.COLON() != nil {
		label = ctx.Identifier(0).GetText()
		iterVar = ctx.Identifier(1).GetText() // Iterator is the second identifier
	} else {
		// No label, iterator is the first identifier
		iterVar = ctx.Identifier(0).GetText()
	}

	iterable := v.VisitExpr(ctx.Expr().(*parser.ExprContext)).(Expression)
	body := v.VisitBlock(ctx.Block().(*parser.BlockContext)).(*Block)

	return &ForStatement{
		Label:    label,
		Variable: iterVar,
		Iterable: iterable,
		Body:     body,
		Position: v.getPosition(ctx),
	}
}

// VisitSwitchStatement builds a switch statement node
func (v *ASTBuilder) VisitSwitchStatement(ctx *parser.SwitchStatementContext) interface{} {
	value := v.VisitExpr(ctx.Expr().(*parser.ExprContext)).(Expression)

	// Process case clauses
	cases := []*CaseClause{}
	for _, caseCtx := range ctx.AllCaseClause() {
		clause := v.VisitCaseClause(caseCtx.(*parser.CaseClauseContext)).(*CaseClause)
		cases = append(cases, clause)
	}

	return &SwitchStatement{
		Value:    value,
		Cases:    cases,
		Position: v.getPosition(ctx),
	}
}

// VisitCaseClause builds a case clause node
func (v *ASTBuilder) VisitCaseClause(ctx *parser.CaseClauseContext) interface{} {
	pattern := v.VisitPattern(ctx.Pattern().(*parser.PatternContext)).(*Pattern)

	// Handle statement-style or block-style body
	var statements []Statement
	var block *Block

	if ctx.Statement() != nil {
		stmt := v.VisitStatement(ctx.Statement().(*parser.StatementContext)).(Statement)
		statements = []Statement{stmt}
	} else if ctx.Block() != nil {
		block = v.VisitBlock(ctx.Block().(*parser.BlockContext)).(*Block)
		statements = block.Statements
	}

	return &CaseClause{
		Pattern:    pattern,
		Statements: statements,
		Position:   v.getPosition(ctx),
	}
}

// VisitPattern builds a pattern node
func (v *ASTBuilder) VisitPattern(ctx *parser.PatternContext) interface{} {
	// Implement pattern matching based on the pattern rule
	// This is a simplified version - expand based on your grammar
	if ctx.Identifier() != nil && ctx.LPAREN() == nil && ctx.LBRACE() == nil {
		// Variable pattern
		return &Pattern{
			Kind:     PatternVariable,
			Value:    ctx.Identifier().GetText(),
			Position: v.getPosition(ctx),
		}
	} else if ctx.UNDERSCORE() != nil {
		// Wildcard pattern
		return &Pattern{
			Kind:     PatternWildcard,
			Position: v.getPosition(ctx),
		}
	} else if ctx.Literal() != nil {
		// Literal pattern
		lit := v.VisitLiteral(ctx.Literal().(*parser.LiteralContext)).(*Literal)
		return &Pattern{
			Kind:     PatternLiteral,
			Literal:  lit,
			Position: v.getPosition(ctx),
		}
	}

	// Default case
	return &Pattern{
		Kind:     PatternUnknown,
		Position: v.getPosition(ctx),
	}
}

// VisitBreakStatement builds a break statement node
func (v *ASTBuilder) VisitBreakStatement(ctx *parser.BreakStatementContext) interface{} {
	var label string
	if ctx.Identifier() != nil {
		label = ctx.Identifier().GetText()
	}

	return &BreakStatement{
		Label:    label,
		Position: v.getPosition(ctx),
	}
}

// VisitContinueStatement builds a continue statement node
func (v *ASTBuilder) VisitContinueStatement(ctx *parser.ContinueStatementContext) interface{} {
	var label string
	if ctx.Identifier() != nil {
		label = ctx.Identifier().GetText()
	}

	return &ContinueStatement{
		Label:    label,
		Position: v.getPosition(ctx),
	}
}

// VisitBlockStatement builds a block statement node
func (v *ASTBuilder) VisitBlockStatement(ctx *parser.BlockStatementContext) interface{} {
	block := v.VisitBlock(ctx.Block().(*parser.BlockContext)).(*Block)

	return &BlockStatement{
		Block:    block,
		Position: v.getPosition(ctx),
	}
}

// VisitTryStatement builds a try statement node
// func (v *ASTBuilder) VisitTryStatement(ctx *parser.TryStatementContext) interface{} {
// 	body := v.VisitBlock(ctx.Block().(*parser.BlockContext)).(*Block)

// 	return &TryStatement{
// 		Body:     body,
// 		Position: v.getPosition(ctx),
// 	}
// }

// Statement node types
type LetStatement struct {
	Name      string
	IsMutable bool
	Type      Type
	Value     Expression
	Position  Position
}

func (s *LetStatement) Pos() Position  { return s.Position }
func (s *LetStatement) statementNode() {}
